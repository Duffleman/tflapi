# tflapi

Quick example to get it started, will have to write something cool up

```go
import (
    "github.com/duffleman/tflapi"
)

client := tflapi.NewClient("", "", nil)

stations, err := client.Line.ListStationsOnLine("district", false)
if err != nil {
    panic(err)
}

for _, station := range stations {
    if station.ID == "940GZZLUWPL" {
        pretty.Println(station)
    }
}
```
