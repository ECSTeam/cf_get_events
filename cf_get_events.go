package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"time"
	"os"
	"regexp"
	"strings"
)

// Events represents Buildpack Usage CLI interface
type Events struct{}

// Metadata is the data retrived from the response json
type Metadata struct {
	GUID string `json:"guid"`
}

// GetMetadata provides the Cloud Foundry CLI with metadata to provide user about how to use `get-events` command
func (c *Events) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "get-events",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 5,
			Build: 20161227,
		},
		Commands: []plugin.Command{
			{
				Name:     "get-events",
				HelpText: "Get microservice events (by akoranne@ecsteam.com)",
				UsageDetails: plugin.Usage{
					Usage: "cf get-events --today\n   cf get-events --yesterday\n   cf get-events --date <yyyy-Moon-dd>\n   cf get-events --all\n ",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(Events))
}

// Run is what is executed by the Cloud Foundry CLI when the get-events command is specified
func (c Events) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "get-events" {

		orgs := c.GetOrgs(cli)
		spaces := c.GetSpaces(cli)
		apps := c.GetAppData(cli)
		today := time.Now()

		// var filterDate = fmt.Sprintf("%s", today.Format("2006-01-02"))
		var filterDate = GetStartOfDay(today)
		if len(args) == 2 {
			if args[1] == "--all" {
				// request for all events
				// this will be costly..
				nintyDays := time.Hour * -(24 * 90)
				filterDate = today.Add(nintyDays) // today - 90  days
			} else if args[1] == "--yesterday" {
				oneDay := time.Hour * -24
				filterDate = GetStartOfDay(today.Add(oneDay)) // today - 1 day
				// filterDate = fmt.Sprintf("%s", yesterday.Format("2006-01-02"))
			} else {
				// all other switches default to today
				// to avoid runaway
			}
			events := c.GetEventsData(cli, filterDate)
			c.EventsInCSVFormat(filterDate, orgs, spaces, apps, events)
		} else if len(args) == 3 {
			// fmt.Println("------->  (0) totals args - ", len(args), ",", args[1], ",", args[2])
			// a filter date was passed in. Use that.
			if args[1] == "--date" {
				const layout = "2006-Jan-02"
				t, err := time.Parse(layout, args[2])
				// fmt.Println("-------> (1) filter date - ", t, filterDate, err)
				if err != nil {
					fmt.Println("Error: Failed to parse given date - ", args[2])
					fmt.Println(err)
					Usage(1)
				} else {
					// filterDate = fmt.Sprintf("%s", t.Format("2006-01-02"))
					filterDate = t
				}
			}
			//	fmt.Println("--------> (3) calling getEvents with filter date - ", filterDate)
			events := c.GetEventsData(cli, filterDate)
			c.EventsInCSVFormat(filterDate, orgs, spaces, apps, events)
		} else {
			fmt.Println("\nMissing one or more arguments ... ")
			Usage(0)
		}
	}
}


func Usage(code int) {
	fmt.Println("")
	fmt.Println("Usage: cf get-events --today")
	fmt.Println("       cf get-events --yesterday")
	fmt.Println("       cf get-events --all")
	fmt.Println("       cf get-events --date <yyyy-mm-dd>")
	fmt.Println("             where: filter date in <yyyy-mm-dd>")
	os.Exit(code)
}

func GetStartOfDay(today time.Time) (time.Time) {
	var now = fmt.Sprintf("%s", today.Format("2006-01-02"))
	t, _ := time.Parse(time.RFC3339, now+"T00:00:00Z")
	return t
}

func StringToDate(dtStr string) (time.Time) {
	// t, _ := time.Parse(time.RFC3339, dtStr+"T00:00:00Z")
	t, _ := time.Parse(time.RFC3339, dtStr)
	return t
}


// PrintInMarkDownFormat prints the buildpack data to console
func (c Events) EventsInCSVFormat(filterDate time.Time, orgs map[string]string, spaces map[string]SpaceSearchEntity, apps AppSearchResults, events EventSearchResults) {

	fmt.Println("")
	fmt.Printf("Following events were recorded from '%s' \n\n", filterDate)

	//  "20161212", "dr", "lab", "app", "pcf-status", "pcf-status",  "app.crash", "crashed", "2 error(s) occurred:\n\n* 2 error(s) occurred:\n\n* Exited with status 255 (out of memory)\n* cancelled\n* 1 error(s) occurred:\n\n* cancelled"
	//  "2016-12-09T21:44:46Z", "demo", "sandbox", "app", "test-nodejs", "admin", "app.update", "stopped", ""

	fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s\n",
		"DATE", "ORG", "SPACE", "ACTEE-TYPE", "ACTEE-NAME", "ACTOR", "EVENT TYPE", "DETAILS")

	//fmt.Println("# of entities: ", len(events.Resources))

	for _, val := range events.Resources  {

		evTmsp, _ := time.Parse(time.RFC3339, val.Entity.Timestamp)
		// fmt.Println("timestamps: ", evTmsp.Nanosecond(), filterDate.Nanosecond(), )

		if (evTmsp.Before(filterDate)) {
			// all events are retrieved in descending order.
			// we processed all events that are filterDate onwards
			// reached older events, break out
			fmt.Println("event date: ", evTmsp, "filterDate: ", filterDate )
			break
		}

		space := spaces[val.Entity.SpaceGUID]
		spaceName := space.Name
		orgName := orgs[space.OrgGUID]

		var mdata = sanitize(fmt.Sprintf("%+v", val.Entity.Metadata))
		fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s\n",
			val.Entity.Timestamp, orgName, spaceName,
			val.Entity.ActeeType, val.Entity.ActeeName, val.Entity.ActorName, val.Entity.Type, mdata)
	}

}


func sanitize(data string) (string) {
	var re = regexp.MustCompile(`\r?\n`)
	var str = re.ReplaceAllString(data, ";")
	str = strings.Replace(str, ";;", ";", 1)
	return str;
}
