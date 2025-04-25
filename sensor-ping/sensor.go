
package main

import (
    "encoding/json"
    "math/rand" 
    "time"
)


type Sensor struct {
    ID string
}

type Payload struct {
    SensorID  string    `json:"sensor_id"`
    Status    string    `json:"status"`
    Timestamp time.Time `json:"timestamp"`
}

func NewSensor(id string) Sensor {
    return Sensor{ID: id}
}

func (s Sensor) Ping() string {
    status := "ok"
    if rand.Intn(10) == 0 { 
        status = "fail"
    }

    payload := Payload{
        SensorID:  s.ID,
        Status:    status,
        Timestamp: time.Now(),
    }

    payloadBytes, _ := json.Marshal(payload)
    return string(payloadBytes)
}

