// Copyright (c) 2016 ECS Team, Inc. - All Rights Reserved
// https://github.com/ECSTeam/cloudfoundry-top-plugin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"time"
	"os"
	"regexp"
	"strings"
	"github.com/simonleung8/flags"
	"bytes"
	"encoding/json"
)


// Events represents Buildpack Usage CLI interface
type Events struct{}

// Metadata is the data retrived from the response json
type Metadata struct {
	GUID string `json:"guid"`
}

// Inputs represent the parsed input args
type Inputs struct {
	fromDate time.Time
	toDate   time.Time
	isCsv    bool
	isJson   bool
}

// GetMetadata provides the Cloud Foundry CLI with metadata to provide user about how to use `get-events` command
func (c *Events) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "get-events",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "get-events",
				HelpText: "Get microservice events (by akoranne@ecsteam.com)",
				UsageDetails: plugin.Usage {
					Usage: UsageText(),
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
	var ins Inputs

	switch args[0] {
	case "get-events":
		ins = c.buildClientOptions(args)
	case "example-alternate-command":
	default:
		return
	}

	orgs := c.GetOrgs(cli)
	spaces := c.GetSpaces(cli)
	apps := c.GetAppData(cli)
	events := c.GetEventsData(cli, ins)
	results := c.FilterResults(cli, ins, orgs, spaces, apps, events)
	if (ins.isCsv) {
		c.EventsInCSVFormat(results)
	} else {
		c.EventsInJsonFormat(results)
	}
}


func Usage(code int) {
	fmt.Println("\nUsage: ", UsageText())
	os.Exit(code)
}

func UsageText() (string) {
	usage := "cf get-events [options]" +
		"\n    where options include: " +
		"\n       --today                  : get all events for today (till now)" +
		"\n       --yesterday              : get events for yesterday ownwards (till now)" +
		"\n       --yesterday-on           : get events from yesterday only" +
		"\n       --all                    : get all events (defaults to last 90 days)" +
		"\n       --json                   : list output in json format (default is csv)\n" +
		"\n       --frdt <yyyymmdd>        : get events from given date onwards (till now)" +
		"\n       --frdtm <yyyymmddhhmmss> : get events from given date and time onwards (till now)" +
		"\n       --todt <yyyymmdd>        : get events till given date" +
		"\n       --todtm <yyyymmddhhmmss> : get events till given date and time\n" +
		"\n       --frdt <yyyymmdd> --todt <yyyymmdd>" +
		"\n       --frdtm <yyyymmddhhmmss> --todtm <yyyymmddhhmmss>"
	return usage
}

func GetStartOfDay(today time.Time) (time.Time) {
	var now = fmt.Sprintf("%s", today.Format("2006-01-02"))
	t, _ := time.Parse(time.RFC3339, now+"T00:00:00Z")
	return t
}

func GetEndOfDay(today time.Time) (time.Time) {
	var now = fmt.Sprintf("%s", today.Format("2006-01-02"))
	t, _ := time.Parse(time.RFC3339, now+"T23:59:59Z")
	return t
}

// sanitize data by replacing \r, and \n with ';'
func sanitize(data string) (string) {
	var re = regexp.MustCompile(`\r?\n`)
	var str = re.ReplaceAllString(data, ";")
	str = strings.Replace(str, ";;", ";", 1)
	return str;
}

// read arguments passed for the plugin
func (c *Events) buildClientOptions(args[] string) (Inputs) {
	fc := flags.New()
	fc.NewBoolFlag("all", "all", " get all events (defaults to last 90 days)")
	fc.NewBoolFlag("today", "today", "get all events for today (till now)")
	fc.NewBoolFlag("yesterday", "yest", "get events from yesterday only")
	fc.NewBoolFlag("yesterday-on", "yon", "get events for yesterday ownwards (till now)")
	fc.NewStringFlag("frdt", "frdt", "get events from given date onwards (till now)")
	fc.NewStringFlag("frdtm", "frdtm", "get events from given date and time onwards (till now)")
	fc.NewStringFlag("todt", "todt", "get events till given date")
	fc.NewStringFlag("todtm", "todtm", "get events till given date and time")
	fc.NewBoolFlag("json", "js", "list output in json format (default is csv)")
	err := fc.Parse(args[1:]...)

	if err != nil {
		fmt.Println("\n Receive error reading arguments ... ", err)
		Usage(1)
	}

	today := time.Now()

	var ins Inputs
	ins.isCsv = true
	ins.isJson = false
	ins.fromDate = GetStartOfDay(today)
	ins.toDate = time.Now()

	if (fc.IsSet("all")) {
		nintyDays := time.Hour * -(24 * 90)
		ins.fromDate  = today.Add(nintyDays) // today - 90  days
	}
	if (fc.IsSet("today")) {
		ins.fromDate  = GetStartOfDay(today)
	}
	if (fc.IsSet("yesterday")) {
		oneDay := time.Hour * -24
		ins.fromDate  = GetStartOfDay(today.Add(oneDay)) // today - 1 day
		ins.toDate = GetEndOfDay(ins.fromDate )
	}
	if (fc.IsSet("yesterday-on")) {
		oneDay := time.Hour * -24
		ins.fromDate  = GetStartOfDay(today.Add(oneDay)) // today - 1 day
	}
	if (fc.IsSet("frdt")) {
		var value = fc.String("frdt")
		const layout = "20060102"        // yyyymmdd
		t, err := time.Parse(layout, value)
		// fmt.Println("-------> (1) filter date - ", t, filterDate, err)
		if err != nil {
			fmt.Println("Error: Failed to parse given date - ", value)
			fmt.Println(err)
			Usage(1)
		} else {
			// filterDate = fmt.Sprintf("%s", t.Format("2006-01-02"))
			ins.fromDate  = t
		}
	}
	if (fc.IsSet("frdtm")) {
		var value = fc.String("frdtm")
		const layout = "20060102150405"        // yyyymmddhhmmss
		t, err := time.Parse(layout, value)
		// fmt.Println("-------> (1) filter date - ", t, filterDate, err)
		if err != nil {
			fmt.Println("Error: Failed to parse given date - ", value)
			fmt.Println(err)
			Usage(1)
		} else {
			// filterDate = fmt.Sprintf("%s", t.Format("2006-01-02"))
			ins.fromDate  = t
		}
	}
	if (fc.IsSet("todt")) {
		var value = fc.String("todt")
		const layout = "20060102150405"        // yyyymmdd
		t, err := time.Parse(layout, value+"235959")
		// fmt.Println("-------> (1) filter date - ", t, filterDate, err)
		if err != nil {
			fmt.Println("Error: Failed to parse given date - ", value)
			fmt.Println(err)
			Usage(1)
		} else {
			// filterDate = fmt.Sprintf("%s", t.Format("2006-01-02"))
			ins.toDate = t
		}
	}
	if (fc.IsSet("todtm")) {
		var value = fc.String("todtm")
		const layout = "20060102150405"        // yyyymmddhhmmss
		t, err := time.Parse(layout, value)

		// fmt.Println("-------> (1) filter date - ", t, filterDate, err)
		if err != nil {
			fmt.Println("Error: Failed to parse given date - ", value)
			fmt.Println(err)
			Usage(1)
		} else {
			// filterDate = fmt.Sprintf("%s", t.Format("2006-01-02"))
			ins.toDate = t
		}
	}

	if (fc.IsSet("json")) {
		ins.isJson = true
		ins.isCsv = false
	}
	// fmt.Println("-------> (1) ins - ", ins.fromDate, ins.toDate)

	return ins
}

// prints the results as a csv text to console
func (c Events) EventsInCSVFormat(results OutputResults) {
	fmt.Println("")
	fmt.Printf(results.Comment)

	//  "20161212", "dr", "lab", "app", "pcf-status", "pcf-status",  "app.crash", "crashed", "2 error(s) occurred:\n\n* 2 error(s) occurred:\n\n* Exited with status 255 (out of memory)\n* cancelled\n* 1 error(s) occurred:\n\n* cancelled"
	//  "2016-12-09T21:44:46Z", "demo", "sandbox", "app", "test-nodejs", "admin", "app.update", "stopped", ""

	fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s\n", "DATE", "ORG", "SPACE", "ACTEE-TYPE", "ACTEE-NAME", "ACTOR", "EVENT TYPE", "DETAILS")
	for _, val := range results.Resources  {
		var mdata = sanitize(fmt.Sprintf("%+v", val.Entity.Metadata))
		fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s\n",
			val.Entity.Timestamp, val.Entity.Org, val.Entity.Space,
			val.Entity.ActeeType, val.Entity.ActeeName, val.Entity.ActorName, val.Entity.Type, mdata)
	}

}

// prints the results as a json text to console
func (c Events) EventsInJsonFormat(results OutputResults) {
	var out bytes.Buffer
	b, _ := json.Marshal(results)
	err := json.Indent(&out, b, "", "\t")
	if err != nil {
		fmt.Println(" Recevied error formatting json output.")
	} else {
		fmt.Println(out.String())
	}
}

