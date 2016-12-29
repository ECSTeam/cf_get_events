# Cloud Foundry Get Events CLI Plugin

Cloud Foundry plugin to view events associated with microservices.

## Install

```
$ go get github.com/ECSTeam/cf_get_events
$ cf install-plugin $GOPATH/bin/cf_get_events
```

## Usage

```
 $> cf get-events --help
NAME:
   get-events - Get microservice events (by akoranne@ecsteam.com)

USAGE:
   cf get-events --today
   cf get-events --yesterday
   cf get-events --all
   cf get-events --date <yyyymmdd>
   cf get-events --datetime <yyyymmddhhmmss>

```

## Sample Output

```
 $> cf get-events --today

Following events were recorded from '2016-12-19 00:00:00 +0000 UTC'

DATE,ORG,SPACE,ACTEE-TYPE,ACTEE-NAME,ACTOR,EVENT TYPE,DETAILS
2016-12-20T05:23:43Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.delete-request,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-20T05:22:49Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
2016-12-20T05:22:43Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-20T05:22:43Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.map-route,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-20T05:22:42Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.create,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
2016-12-20T04:18:08Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:3683a577-4900-4b33-572e-d5e879ba189a Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* 1 error(s) occurred:;;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T04:02:03Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:f1d38235-0de5-4910-6752-cc244d5b1f19 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T03:46:00Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:66fd0d21-bcef-49af-4401-8a5d772c7c3a Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T03:29:56Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:9f72f8bb-f6a9-4d91-7d9c-5998666a090d Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T03:13:52Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:cf40567c-89af-4793-5827-ab70541383e4 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T02:57:48Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:9874e0fa-7af4-4cc0-7941-56486512686f Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T02:41:44Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:e6d66320-726d-409c-72b4-a6059325cd4f Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T02:31:32Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.delete-request,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-20T02:30:47Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STARTED Recursive:}}
2016-12-20T02:30:46Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State:STOPPED Recursive:}}
2016-12-20T02:30:39Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-20T02:25:41Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:d8a84502-0550-4148-6125-f855c4556643 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T02:09:37Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:d270894a-8db7-4821-57d7-1ab7eb959b39 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T01:53:34Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:732c37fa-29ee-409b-467c-d4a5347a9a3e Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T01:37:29Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:23554368-bff1-4dfc-794d-a8090a973bcc Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T01:21:24Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:0075ab51-6c21-4b4a-5e3b-4e534713cb98 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T01:05:22Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:eddc18ae-0bc8-4d26-6648-2dcdefe3c7f4 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T00:49:17Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:dbc3bace-e443-4cf8-6480-94e83ec09c6e Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T00:33:14Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:f984efbd-05c0-4076-7f5a-eef87c6b4ea2 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T00:17:08Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:e1f68c67-b561-4414-5b99-2a74d4ae3d76 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
2016-12-20T00:01:05Z,demo,sandbox,,misbehaving_app,misbehaving_app,app.crash,{Instance:f0dbc217-88fd-43d4-6510-ddaaebad68c7 Index:0 ExitDescription:2 error(s) occurred:;* 2 error(s) occurred:;;* Exited with status 66;* cancelled;* 1 error(s) occurred:;;* cancelled Reason:CRASHED Request:{State: Recursive:}} $>
```


```
$> date
Wed Dec 28 13:31:41 CST 2016

$> cf get-events --yesterday

Following events were recorded from '2016-12-27 00:00:00 +0000 UTC'

DATE,ORG,SPACE,ACTEE-TYPE,ACTEE-NAME,ACTOR,EVENT TYPE,DETAILS
2016-12-28T17:00:08Z,demo,sandbox,,java-kill,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:54:27Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:54:11Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:53:00Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-27T19:17:55Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-27T19:17:55Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.map-route,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-27T00:24:13Z,dr,lab,,testApp01,testApp01,app.crash,{Instance:1a4b839b-f7b3-44e2-7db3-96c228ad483b Index:0 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 2;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
```

```
$> cf get-events --date 20161225

Following events were recorded from '2016-12-25 00:00:00 +0000 UTC'

DATE,ORG,SPACE,ACTEE-TYPE,ACTEE-NAME,ACTOR,EVENT TYPE,DETAILS
2016-12-28T17:00:08Z,demo,sandbox,,java-kill,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:54:27Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:54:11Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:53:00Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-27T19:17:55Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-27T19:17:55Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.map-route,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-27T00:24:13Z,dr,lab,,testApp01,testApp01,app.crash,{Instance:1a4b839b-f7b3-44e2-7db3-96c228ad483b Index:0 ExitDescription:2 error(s) occurred:;* 1 error(s) occurred:;;* Exited with status 2;* 2 error(s) occurred:;;* cancelled;* cancelled Reason:CRASHED Request:{State: Recursive:}}
```

```
$> cf get-events --datetime 20161227190000

Following events were recorded from '2016-12-27 19:00:00 +0000 UTC'

DATE,ORG,SPACE,ACTEE-TYPE,ACTEE-NAME,ACTOR,EVENT TYPE,DETAILS
2016-12-28T17:00:08Z,demo,sandbox,,java-kill,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:54:27Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:54:11Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:53:00Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-27T19:17:55Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.update,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-27T19:17:55Z,sandbox,lots-of-apps,,testApp000,admin,audit.app.map-route,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
```


```
$> cf get-events --datetime 20161227191800

Following events were recorded from '2016-12-27 19:18:00 +0000 UTC'

DATE,ORG,SPACE,ACTEE-TYPE,ACTEE-NAME,ACTOR,EVENT TYPE,DETAILS
2016-12-28T17:00:08Z,demo,sandbox,,java-kill,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:54:27Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:54:11Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
2016-12-28T16:53:00Z,dr,lab,,pcf-status,admin,audit.app.ssh-authorized,{Instance: Index:0 ExitDescription: Reason: Request:{State: Recursive:}}
```

## Uninstall

```
$ cf uninstall-plugin get-events
```

## Motivation 

_Why do I need this plugin?_ 

In a large organization, a cloud coundry foundation can have hundreds of microservices. 
Different pipelines can push changes through out the day and night. 
The `get-events` plugin allows the platform operator to get a quick snapshot of all the 
service events that took place today, or since yesterday, or from a particular date. 

If a micro-service crashes, the service will be restarted. That is one big benefit of cloud foundry platform. 
However, this resilience can also mask services that crash frequently. The `get-events` plug-in will highlight such services.

Using [cf_scripts/app_profiler](https://github.com/ECSTeam/cf_scripts/tree/master/app_profiler)
the platform operator can script forwarding the plugin output to `Splunk` or `Statsd` based event logger. 
This will help capture events across time and understand event patterns.
