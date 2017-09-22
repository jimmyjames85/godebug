package godebug

import (
	"fmt"
	"log"
)

// Yell is a wrapper around log.Printf used to make text stand out in
// bold yello by using escape codes from `tput`
func Yell(format string, a ...interface{}) {
	log.Printf("[1m[33m%s[m", fmt.Sprintf(format, a...))
}
