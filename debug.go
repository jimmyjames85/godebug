package godebug

import (
	"fmt"
	"log"
)

const (
	bold   = "[1m"
	red    = "[31m"
	green  = "[32m"
	yellow = "[33m"
	blue   = "[34m"
	purple = "[35m"
	cyan   = "[36m"
	reset  = "[m"
)

// Red is a wrapper around log.Printf used to make text stand out in
// bold red text by using escape codes from `tput`
func Redf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, red, s, reset)
}

// Green is a wrapper around log.Printf used to make text stand out in
// bold green text by using escape codes from `tput`
func Greenf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, green, s, reset)
}

// Yellow is a wrapper around log.Printf used to make text stand out
// in bold yellow text by using escape codes from `tput`
func Yellowf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, yellow, s, reset)
}

// Blue is a wrapper around log.Printf used to make text stand out in
// bold blue text by using escape codes from `tput`
func Bluef(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, blue, s, reset)
}

// Purple is a wrapper around log.Printf used to make text stand out
// in bold purple text by using escape codes from `tput`
func Purplef(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, purple, s, reset)
}

// Cyan is a wrapper around log.Printf used to make text stand out in
// bold cyan text by using escape codes from `tput`
func Cyanf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, cyan, s, reset)
}
