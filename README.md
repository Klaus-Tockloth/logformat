# logformat()
A good logging concept is to provide human and machine readable logs. In order to achieve this, it's recommended to log everything
as key/value pairs. That is what this snippet does. It can be used with the standard Go logger.

## Usage
```
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
```

## Output
```
2018/05/16 06:58:53.470799 main.go:92: "level=INFO", "message=map build ok", "id=42", "fruits=[apple banana]"
2018/05/16 06:58:53.471016 main.go:99: "level=DEVEL", "message=processing startet", "person={Sean 50}", "hash=map[foo:11 bar:22]"
2018/05/16 06:58:53.471083 main.go:104: "level=METRIC", "cpu-user=23.69", "cpu-sys=33.64", "cpu-idle=42.65"
```

# trace()
During development it's sometimes helpful to visualize the program flow (calling chain). The trace() function uses the runtime in order to
get the caller. It calls logformat() for outputting function start and end (including filename, linenumber and duration). The usage is very simple.
Just call "defer trace()()" at the beginning of your function. Control tracing with a flag. Don't use tracing in production.

## Usage
```
func rogueOne() {
	if Tracing {
		defer trace()()
	}
	time.Sleep(321 * time.Millisecond)
}
```

## Output
```
2018/05/16 12:06:21.146105 main.go:161: "level=TRACE", "entry=main.rogueOne", "file=/Users/klaus/go/src/klaus/logformat/main.go", "line=174"
2018/05/16 12:06:21.467578 main.go:164: "level=TRACE", "exit=main.rogueOne", "duration=321.38664ms"
```