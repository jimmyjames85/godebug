package godebug

import (
	"encoding/json"
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

// Sredf is a wrapper around fmt.Sprintf used to make text stand out
// in bold red text by using escape codes from `tput`
func Sredf(format string, a ...interface{}) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s%s%s%s", bold, red, s, reset)
}

// Sgreenf is a wrapper around fmt.Sprintf used to make text stand out
// in bold green text by using escape codes from `tput`
func Sgreenf(format string, a ...interface{}) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s%s%s%s", bold, green, s, reset)
}

// Syellowf is a wrapper around fmt.Sprintf used to make text stand
// out in bold yellow text by using escape codes from `tput`
func Syellowf(format string, a ...interface{}) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s%s%s%s", bold, yellow, s, reset)
}

// Sbluef is a wrapper around fmt.Sprintf used to make text stand out
// in bold blue text by using escape codes from `tput`
func Sbluef(format string, a ...interface{}) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s%s%s%s", bold, blue, s, reset)
}

// Spurplef is a wrapper around fmt.Sprintf used to make text stand
// out in bold purple text by using escape codes from `tput`
func Spurplef(format string, a ...interface{}) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s%s%s%s", bold, purple, s, reset)
}

// Scyanf is a wrapper around fmt.Sprintf used to make text stand out
// in bold cyan text by using escape codes from `tput`
func Scyanf(format string, a ...interface{}) string {
	s := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s%s%s%s", bold, cyan, s, reset)
}

// Redf is a wrapper around log.Printf used to make text stand out in
// bold red text by using escape codes from `tput`
func Redf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, red, s, reset)
}

// Greenf is a wrapper around log.Printf used to make text stand out
// in bold green text by using escape codes from `tput`
func Greenf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, green, s, reset)
}

// Yellowf is a wrapper around log.Printf used to make text stand out
// in bold yellow text by using escape codes from `tput`
func Yellowf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, yellow, s, reset)
}

// Bluef is a wrapper around log.Printf used to make text stand out in
// bold blue text by using escape codes from `tput`
func Bluef(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, blue, s, reset)
}

// Purplef is a wrapper around log.Printf used to make text stand out
// in bold purple text by using escape codes from `tput`
func Purplef(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, purple, s, reset)
}

// Cyanf is a wrapper around log.Printf used to make text stand out in
// bold cyan text by using escape codes from `tput`
func Cyanf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s%s", bold, cyan, s, reset)
}

// ToJSON returns the JSON form of obj. If unable to Marshal obj, a JSON error message is returned
// with the %#v formatted string of the object
func ToJSON(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		return fmt.Sprintf(`{"error":"failed to marshal into JSON","obj":%q}`, fmt.Sprintf("%#v", obj))
	}
	return string(b)
}
