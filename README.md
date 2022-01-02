# fromorc-go

commandline tool that will get you simplified trail conditions from the morc
website trail conditions https://trails.morcmtb.org.

this is a golang version of my [shell-script version of fromorc](https://github.com/bikeonastick/fromorc). same functionality.

## Install/run

1. clone this repo 
1. run `go run .`

## Features

### List all trails with status

```
./fromorc.sh
```

#### Outputs

* trail name
* emoji to indicate status: open(👍) or closed(👎)
* emoji to indicate how fresh the status is:
  * ✅ no more than two days old 
  * 🤞 between 2 days and a week
  * 💩 older than a week

## Acknowledgment

Taking inspiration from [dtanner](https://gist.github.com/dtanner/54b10ef8932b026afec0398495b5b2b5).



