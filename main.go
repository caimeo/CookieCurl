// mirthapi project main.go
package main

import (
	"flag"

	"github.com/caimeo/console"
	"github.com/caimeo/iniflags"
	"github.com/caimeo/stickyjar/curljar"
)

var verboseMode = flag.Bool("verbose", false, "Verbose trace output.")
var debugMode = flag.Bool("debug", false, "Debug trace output.")
var cookieFileP = flag.String("file", "", "Cookiejar file")
var cookieFile string

var t console.Console

func main() {
	iniflags.SetConfigFile(".settings")
	iniflags.SetAllowMissingConfigFile(true)
	iniflags.Parse()

	cookieFile = *cookieFileP

	t := console.New(*verboseMode, *debugMode)

	t.Always("Cookie Curl")
	t.Always(cookieFile)

	cj, _ := curljar.New(cookieFile, nil)

	t.Verbose(cj.String())

}
