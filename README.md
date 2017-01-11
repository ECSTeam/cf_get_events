# Cloud Foundry Get Events CLI Plugin

Cloud Foundry plugin to view events for applications or long running processes - LRPs.

## Install

```
$ go get github.com/ECSTeam/cf_get_events
$ cf install-plugin $GOPATH/bin/cf_get_events
```

## Motivation 

_Why do I need this plugin?_ 

In a large organization, a cloud foundry foundation can have hundreds of application instances and LRPs. 
Different pipelines can push changes through out the day and night. 
The `get-events` plugin allows the platform operator to get a quick snapshot of all the 
LRP events that took place today, or since yesterday, or from a particular date. 

If a LRP crashes, that LRP will be restarted. That is one big benefit of cloud foundry platform. 
However, this resilience can also mask services that crash frequently. 
The `get-events` plug-in will highlight such LRPs.

Using [cf_scripts/app_profiler](https://github.com/ECSTeam/cf_scripts/tree/master/app_profiler)
the platform operator can script forwarding the plugin output to `Splunk` or `Statsd` based event logger. 
This will help capture events across time and understand event patterns.


## Usage

```
 $> cf get-events --help
NAME:
   get-events - Get events for applications / long running processes (LRPs) (by akoranne@ecsteam.com)

Usage: cf get-events [options]
    where options include:
       --today                  : get all events for today (till now)
       --yesterday              : get events for yesterday ownwards (till now)
       --yesterday-on           : get events from yesterday only
       --all                    : get all events (defaults to last 90 days)
       --json                   : list output in json format (default is csv)

       --from <yyyymmdd>        : get events from given date onwards (till now)
       --from <yyyymmddhhmmss>  : get events from given date and time onwards (till now)
       --to <yyyymmdd>          : get events till given date
       --to <yyyymmddhhmmss>    : get events till given date and time

       --from <yyyymmdd> --to <yyyymmdd>
       --from <yyyymmddhhmmss> --to <yyyymmddhhmmss>
```

## Access 

The `get-events` plugin will show events for the orgs and spaces that the current user has access too.

If you want events, across all orgs, and spaces the plugin user will 
need `cloud controller admin` access to get all events. To do that, please follow the steps below.

```
   $ uaac token client get admin -s <MyAdminPassword>
   $ uaac user add event_plugin_user -p welcome1 --emails <event_plugin_user@mydomain.com>
   $ uaac member add cloud_controller.admin event_plugin_user   
```

## Sample Output

```
$> date
	Tue Jan 10 15:19:24 CST 2017
```


```
 $> cf get-events --today

	Following events were recorded from '2017-01-10 00:00:00 +0000 UTC', to '2017-01-10 15:17:58.879052387 -0600 CST'
	
	DATE,ORG,SPACE,ACTEE-TYPE,ACTEE-NAME,ACTOR,EVENT TYPE,DETAILS
	2017-01-10T21:11:05Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T20:54:23Z,dr,lab,,test,test,app.crash,{Instance:3abeb474-684b-4d87-6e20-65140c5e5755 Index:11 ExitDescription:2 error(s) occurred:;* exceeded 30s timeout;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T20:54:15Z,dr,lab,,test,test,app.crash,{Instance:7f14cc73-86e3-48b8-44a2-a4f77e8a00d5 Index:41 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 1;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T20:54:13Z,dr,lab,,test,test,app.crash,{Instance:6b3e526f-c93d-40f7-6e6b-02d0944745b7 Index:45 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 1;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T19:56:59Z,dr,lab,,test,test,app.crash,{Instance:dc60b699-3cb4-423e-5560-a301daa741ef Index:21 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 1;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T19:56:52Z,sandbox,lots-of-apps,,testApp100,testApp100,app.crash,{Instance:7f66b752-c904-4cc5-7894-de6c867fcc6a Index:0 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 1;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T19:41:19Z,sandbox,lots-of-apps,,top-interactive078,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-10T19:41:03Z,sandbox,lots-of-apps,,top-interactive078,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T19:41:03Z,sandbox,lots-of-apps,,top-interactive078,admin,audit.app.map-route,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T19:41:02Z,sandbox,lots-of-apps,,top-interactive078,admin,audit.app.create,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-10T19:30:49Z,sandbox,lots-of-apps,,plugins,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T19:30:06Z,sandbox,lots-of-apps,,plugins,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T19:29:59Z,sandbox,lots-of-apps,,plugins,admin,audit.app.create,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-10T18:30:56Z,dr,lab,,test,test,app.crash,{Instance:54d3580d-8eb9-4568-4cec-07715c4f025e Index:13 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 2;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T09:49:33Z,dr,lab,,test,test,app.crash,{Instance:eccbec7c-c52b-41b9-5bc9-3e6c88fe95cb Index:79 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 1;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T07:54:07Z,sandbox,cftop,,top,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-10T07:54:07Z,sandbox,cftop,,top,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}

```


```
$> cf get-events --yesterday

	Following events were recorded from '2017-01-09 00:00:00 +0000 UTC', to '2017-01-09 23:59:59 +0000 UTC'
	
	DATE,ORG,SPACE,ACTEE-TYPE,ACTEE-NAME,ACTOR,EVENT TYPE,DETAILS
	2017-01-09T22:26:30Z,sandbox,cftop,,top-interactive,top-interactive,app.crash,{Instance:e40dea11-a43f-4b78-4d21-b2dff4008e73 Index:0 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 1;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T22:25:54Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T22:25:53Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T22:25:51Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T22:22:23Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T22:22:22Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T22:22:17Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T22:21:27Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T22:19:06Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T22:19:05Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T22:19:04Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T22:17:56Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T22:17:56Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T22:17:55Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T22:16:30Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T22:16:29Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T22:16:29Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T22:15:05Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T22:15:05Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T22:15:04Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T22:04:51Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T22:04:50Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T22:04:50Z,dr,lab,,exampleb,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T19:34:16Z,demo,sandbox,,simple-rules,admin,audit.app.delete-request,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T19:20:14Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:5b177376-c39a-48b5-4297-6819f7a9cf8d Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T19:04:11Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:6168ca8d-8cb4-45c5-7da9-d263f99d4952 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:48:06Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:b42593f4-e6ec-4584-53b1-83475e623cbc Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:32:01Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:ec2be0a8-0353-4f0b-5dd0-2fb77e18f91c Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:23:58Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:744076eb-2b13-4382-745a-b9565618595c Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:19:27Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:5bb9a0fe-ae31-44f8-7919-226d5d6385e2 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:16:56Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:b5491fae-c1f0-421e-69ea-ff30c9369d9b Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:15:26Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:b3cdc42e-29dd-48c2-4fe4-d5db72b55eb7 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:14:47Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:5d449220-4e2b-4628-44f2-0bd3376c644f Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:14:46Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:851af908-6194-4739-65d6-a51017323dcf Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:14:38Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:2e841433-7428-4c9d-732d-8291f28d5d68 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* 1 error(s) occurred:;;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:13:47Z,demo,sandbox,,simple-rules,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T18:13:45Z,demo,sandbox,,simple-rules,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T18:13:37Z,demo,sandbox,,simple-rules,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T18:13:10Z,demo,sandbox,,simple-rules,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T18:12:11Z,demo,sandbox,,simple-rules,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T18:10:34Z,demo,sandbox,,simple-rules,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T18:10:22Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:1bf23031-54c9-490d-54bf-994bf038b777 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:05:50Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:f691871f-d14e-406b-4489-6b8ad7ba2c95 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:03:20Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:7c001fc1-452d-4e28-48bb-0ef3dd3c138e Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:01:49Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:773fba02-2e21-458c-58d6-78499e775b0a Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:00:54Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:4a465cf2-181b-4324-7b4b-156ecd558a78 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* 1 error(s) occurred:;;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:00:52Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:7f11c9af-bf1b-476c-6eff-ba8a50f55ef9 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T18:00:44Z,demo,sandbox,,simple-rules,simple-rules,app.crash,{Instance:06fa5dc8-27a6-4233-6ff7-5020380bc19a Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 1;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-09T17:59:50Z,demo,sandbox,,simple-rules,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T17:59:17Z,demo,sandbox,,simple-rules,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T17:59:17Z,demo,sandbox,,simple-rules,admin,audit.app.map-route,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T17:59:15Z,demo,sandbox,,simple-rules,admin,audit.app.create,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T15:28:33Z,sandbox,cftop,,top-interactive,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T15:28:33Z,sandbox,cftop,,top-interactive,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T15:28:29Z,sandbox,cftop,,top,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T15:28:20Z,sandbox,cftop,,top-interactive,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	
```

```
$> cf get-events --frdtm 20170110190000

	Following events were recorded from '2017-01-10 19:00:00 +0000 UTC', to '2017-01-10 15:26:03.134703417 -0600 CST'
	
	DATE,ORG,SPACE,ACTEE-TYPE,ACTEE-NAME,ACTOR,EVENT TYPE,DETAILS
	2017-01-10T21:19:38Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T21:19:19Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T21:11:05Z,dr,lab,,test,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T20:54:23Z,dr,lab,,test,test,app.crash,{Instance:3abeb474-684b-4d87-6e20-65140c5e5755 Index:11 ExitDescription:2 error(s) occurred:;* exceeded 30s timeout;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T20:54:15Z,dr,lab,,test,test,app.crash,{Instance:7f14cc73-86e3-48b8-44a2-a4f77e8a00d5 Index:41 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 1;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T20:54:13Z,dr,lab,,test,test,app.crash,{Instance:6b3e526f-c93d-40f7-6e6b-02d0944745b7 Index:45 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 1;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T19:56:59Z,dr,lab,,test,test,app.crash,{Instance:dc60b699-3cb4-423e-5560-a301daa741ef Index:21 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 1;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T19:56:52Z,sandbox,lots-of-apps,,testApp100,testApp100,app.crash,{Instance:7f66b752-c904-4cc5-7894-de6c867fcc6a Index:0 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 1;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
	2017-01-10T19:41:19Z,sandbox,lots-of-apps,,top-interactive078,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-10T19:41:03Z,sandbox,lots-of-apps,,top-interactive078,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T19:41:03Z,sandbox,lots-of-apps,,top-interactive078,admin,audit.app.map-route,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T19:41:02Z,sandbox,lots-of-apps,,top-interactive078,admin,audit.app.create,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-10T19:30:49Z,sandbox,lots-of-apps,,plugins,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T19:30:06Z,sandbox,lots-of-apps,,plugins,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-10T19:29:59Z,sandbox,lots-of-apps,,plugins,admin,audit.app.create,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
```

```
$> cf get-events --frdtm 20170109150000 --todtm 20170109160000

	Following events were recorded from '2017-01-09 15:00:00 +0000 UTC', to '2017-01-09 16:00:00 +0000 UTC'
	
	DATE,ORG,SPACE,ACTEE-TYPE,ACTEE-NAME,ACTOR,EVENT TYPE,DETAILS
	2017-01-09T15:28:33Z,sandbox,cftop,,top-interactive,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
	2017-01-09T15:28:33Z,sandbox,cftop,,top-interactive,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
	2017-01-09T15:28:29Z,sandbox,cftop,,top,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
	2017-01-09T15:28:20Z,sandbox,cftop,,top-interactive,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
```


## Uninstall

```
$ cf uninstall-plugin get-events
```

