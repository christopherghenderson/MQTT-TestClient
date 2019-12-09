//**********************************
// MQTT Test Client:
// Create Test Station/Devices
// Publish random values
//**********************************

package mqtttestclient

import (
	"time"
)

var testClient *TestClient

func main() {}

// CreateTestClient creates a test client and populated it with random devices/points
func CreateTestClient(brokerIP string, port string, clientID string) *TestClient {
	testClient := NewClient(brokerIP, port, clientID)
	testClient.Namespace = "spBv1.0"
	testClient.GroupID = "SparkyDevices"
	testClient.EdgeNodeID = "Oaspark"
	return testClient
}

// CreateRandomStations Return array of random stations with random devices with random values
func (testClient TestClient) CreateRandomStations(stationsNum int, pointsNum int, millisecondBetweenValuechange int) []*Station {
	// Create an array of Random Stations, populate with random devices
	randomStationArray := CreateRandomStations(stationsNum)
	for _, station := range randomStationArray {
		station.Devices = CreateRandomDevices(pointsNum)
	}
	// Set Devices to Random values
	for _, station := range randomStationArray {
		for _, device := range station.Devices {
			device.SetValueToRandom(1000, 10000, millisecondBetweenValuechange)
		}
	}
	return randomStationArray
}

// PublishPoints a
func (testClient TestClient) PublishPoints(millsecondBetweenBroadcast int) {
	// Publish Node Birth
	testClient.WriteNBirth()

	// Publish DBirth (synchronously)
	for _, station := range testClient.Stations {
		testClient.WriteData(station, "DBIRTH")
	}
	time.Sleep(time.Duration(millsecondBetweenBroadcast) * time.Millisecond)

	//&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	// MAIN LOOP (Publishes Ddata)
	//&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	for {
		for _, station := range testClient.Stations {
			testClient.WriteData(station, "DDATA")
			//Prevent Sequence from being out of order
			time.Sleep(4 * time.Millisecond)
		}
		time.Sleep(time.Duration(millsecondBetweenBroadcast) * time.Millisecond)
	}
	//&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
}
