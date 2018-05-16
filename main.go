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
	"runtime"
	"strings"
	"time"
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

// logging control flags
var (
	Tracing    = true
	Developing = true
)

/*
main starts this program
*/
func main() {

	fmt.Printf("\nProgram:\n")
	fmt.Printf("  Name    : %s\n", progName)
	fmt.Printf("  Release : %s - %s\n", progVersion, progDate)
	fmt.Printf("  Purpose : %s\n", progPurpose)
	fmt.Printf("  Info    : %s\n", progInfo)

	// initialize logger
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	// set log output
	file, err := os.OpenFile(os.Args[0]+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

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

	// trace example
	rogueOne()

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

/*
trace traces a function
- no input parameters
resulting format:
2018/05/16 12:06:21.146105 main.go:161: "level=TRACE", "entry=main.rogueOne", "file=/Users/klaus/go/src/klaus/logformat/main.go", "line=174"
2018/05/16 12:06:21.467578 main.go:164: "level=TRACE", "exit=main.rogueOne", "duration=321.38664ms"
*/
func trace() func() {

	callerName := "NO_CALLER"
	callerFile := "NO_FILE"
	callerLine := -1

	fpcs := make([]uintptr, 1)

	// get caller parameters (skip two levels)
	n := runtime.Callers(2, fpcs)
	if n != 0 {
		caller := runtime.FuncForPC(fpcs[0] - 1)
		if caller == nil {
			callerName = "NIL_CALLER"
		} else {
			callerName = caller.Name()
			callerFile, callerLine = caller.FileLine(fpcs[0] - 1)
		}
	}

	start := time.Now()
	log.Println(logformat(TRACE, "entry", callerName, "file", callerFile, "line", callerLine))

	return func() {
		log.Println(logformat(TRACE, "exit", callerName, "duration", time.Since(start)))
	}
}

/*
rogueOne does some work
*/
func rogueOne() {

	if Tracing {
		defer trace()()
	}

	// ... some work ...
	time.Sleep(321 * time.Millisecond)
}
