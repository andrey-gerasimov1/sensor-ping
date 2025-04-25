
package main

import (
    "encoding/json"
    "testing"
)

func TestPingFormat(t *testing.T) {
    sensor := NewSensor("test-sensor")
    output := sensor.Ping()

    var payload Payload
    err := json.Unmarshal([]byte(output), &payload)
    if err != nil {
        t.Fatalf("Ping output is not valid JSON: %v", err)
    }

    if payload.SensorID != "test-sensor" {
        t.Errorf("Expected sensor_id to be 'test-sensor', got %s", payload.SensorID)
    }
}
