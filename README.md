# govoipms
Voip.ms API library written in Go utilizing the JSON API available from here: https://voip.ms/m/apidocs.php

The implementation is currently incomplete. All of the API functions in accounts.go, cdr.go, clients.go, and general.go are complete. Most of the read functions and several of the Set* functions in dids.go are complete. Nothing in fax.go or voicemail.go has been worked on other than the associated types being roughed in. Unfortunately I can no longer spend the effort to complete this project at this time as I've I started on a Python implementation instead. The next version might utilize Voip.MS's SOAP API instead of the JSON API to see if it speeds implementation time.

## Usage

```
v1c := govoipms.NewV1Client("https://voip.ms/api/v1/rest.php", "email", "password", true)
general := v1c.NewGeneralAPI()
log.Println(general.GetBalance(true))
```

See examples/main.go for more details.
