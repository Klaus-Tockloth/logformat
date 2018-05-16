/*
Purpose:
- key/value log formatter

Description:
- Simple key/value log formatter.

Releases:
- 1.0.0 - 2018/05/16 : initial release

Author:
- Klaus Tockloth

Copyright and license:
- Copyright (c) 2018 Klaus Tockloth
- MIT license

Contact (eMail):
- freizeitkarte@googlemail.com

Remarks:
- NN

Links:
- NN
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// general program info
var (
	progName    = os.Args[0]
	progVersion = "1.0.0"
	progDate    = "2018/05/16"
	progPurpose = "key/value log formatter"
	progInfo    = "Simple key/value log formatter."
)

// logging constants
const (
	MSG    string = "message"
	TRACE  string = "TRACE"
	DEVEL  string = "DEVEL"
	INFO   string = "INFO"
	METRIC string = "METRIC"
)

/*
init initializes this program
*/
func init() {

	// initialize logger
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

/*
main starts this program
*/
func main() {

	fmt.Printf("\nProgram:\n")
	fmt.Printf("  Name    : %s\n", progName)
	fmt.Printf("  Release : %s - %s\n", progVersion, progDate)
	fmt.Printf("  Purpose : %s\n", progPurpose)
	fmt.Printf("  Info    : %s\n\n", progInfo)

	id := 42
	fruits := []string{"apple", "banana"}
	log.Println(logformat(INFO, MSG, "map build ok", "id", id, "fruits", fruits))

	person := struct {
		name string
		age  int
	}{"Sean", 50}
	hash := map[string]int{"foo": 11, "bar": 22}
	log.Println(logformat(DEVEL, MSG, "processing startet", "person", person, "hash", hash))

	user := 23.69
	sys := 33.64
	idle := 42.65
	log.Println(logformat(METRIC, "cpu-user", user, "cpu-sys", sys, "cpu-idle", idle))

	fmt.Printf("\n")
	os.Exit(0)
}

/*
logformat formats all log elements
- level = log level (mandatory)
- kvPairs = additional key/value pairs to log (optional)
resulting format:
- level=INFO", "message=map build ok", "id=42", "fruits=[apple banana]"
- level=DEVEL", "message=processing startet", "person={Sean 50}", "hash=map[foo:11 bar:22]"
- level=METRIC", "cpu-user=23.69", "cpu-sys=33.64", "cpu-idle=42.65"
*/
func logformat(level string, kvPairs ...interface{}) string {

	logElements := []string{}
	logElements = append(logElements, fmt.Sprintf("%q", fmt.Sprintf("level=%v", level)))

	if len(kvPairs)%2 != 0 {
		// error: key or value missing
		logElements = append(logElements, fmt.Sprintf("%q", "internal-log-format-error=key or value missing"))
	} else {
		for index := 0; index < len(kvPairs); index += 2 {
			logElements = append(logElements, fmt.Sprintf("%q", fmt.Sprintf("%v=%v", kvPairs[index], kvPairs[index+1])))
		}
	}

	return strings.Join(logElements, ", ")
}
