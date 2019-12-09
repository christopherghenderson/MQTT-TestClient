// A Device holds a value, and a data type.
// This device can be made to hold a random value.

package mqtttestclient

import (
	"fmt"
	"math/rand"
	"time"
)

// Device holds a device with a topic and value
type Device struct {
	analDigital string
	deviceID    string
	Value       interface{}
	dataType    uint32
}

// CreateRandomDevice creates a random device with specified topic parameters
func CreateRandomDevice(dataType string) *Device {
	deviceID := randomDevicePrefix[rand.Intn(len(randomDevicePrefix))] + "-" + fmt.Sprint(rand.Intn(899)+100)
	device := Device{analDigital: dataType, deviceID: deviceID}
	return &device
}

// CreateRandomDevices creates random devices of type, adds to Devices
func CreateRandomDevices(count int) []*Device {
	var devices []*Device
	for i := 0; i < count; i++ {
		randomMessageType := randomMessageType[rand.Intn(len(randomMessageType))]
		devices = append(devices, CreateRandomDevice(randomMessageType))
	}
	return devices
}

// SetValueToRandom sets the device to a random variable on a timer.  Max ignored on non-analog.
func (device *Device) SetValueToRandom(timer int, max int, millisecondBetweenValuechange int) {
	// Set initial value
	if device.analDigital == "analog" {
		device.Value = float32(rand.Intn(max))
		device.dataType = 9
	} else {
		device.Value = rand.Float32() < 0.5
		device.dataType = 11
	}
	// Set to random value every X seconds
	ticker := time.NewTicker(time.Duration(millisecondBetweenValuechange) * time.Millisecond)
	go func() {
		for t := range ticker.C {
			_ = t
			if device.analDigital == "analog" {
				device.Value = float32(rand.Intn(max))
				device.dataType = 9
			} else {
				device.Value = rand.Float32() < 0.5
				device.dataType = 11
			}
		}
	}()
}
