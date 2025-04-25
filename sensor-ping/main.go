
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    sensors := []Sensor{
        NewSensor("sensor-1"),
        NewSensor("sensor-2"),
        NewSensor("sensor-3"),
    }

    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        for _, sensor := range sensors {
            payload := sensor.Ping()
            fmt.Println(payload)
        }
    }
}
