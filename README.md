# logformat (snippet)
A good logging concept is to provide machine readable logs. In order to achieve this, it's recommended to log everything
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

## Result
```
2018/05/16 06:58:50.470799 main.go:92: "level=INFO", "message=map build ok", "id=42", "fruits=[apple banana]"
2018/05/16 06:58:50.471016 main.go:99: "level=DEVEL", "message=processing startet", "person={Sean 50}", "hash=map[foo:11 bar:22]"
2018/05/16 06:58:50.471083 main.go:104: "level=METRIC", "cpu-user=23.69", "cpu-sys=33.64", "cpu-idle=42.65"
```
