// Each Station holds an array of devices.

package mqtttestclient

import (
	"math/rand"
	"time"
)

// Station Holds various devices as part of a station
type Station struct {
	name    string
	Devices []*Device
}

// CreateRandomStations creates an array of Random Stations
func CreateRandomStations(count int) []*Station {
	rand.Seed(time.Now().Unix())
	var randomStationArray []*Station
	for i := 0; i < count; i++ {
		randomStation := randomStationPrefix[rand.Intn(len(randomStationPrefix))] + " " + randomStationSuffix[rand.Intn(len(randomStationSuffix))]
		randomStationArray = append(randomStationArray, &Station{name: randomStation})
	}
	return randomStationArray
}
