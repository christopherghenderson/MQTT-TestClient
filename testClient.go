// A Struct built around the Paho Client, provides extended functionality like aquiring a Sequence number,
// and holds the unambigious remainder of our topic. (Group/EdgeNode shiz)

package mqtttestclient

import (
	"fmt"
	"reflect"
	"time"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	spb "github.com/christopherghenderson/sparkplugbProtobufGo"
	"github.com/golang/protobuf/proto"
)

// TestClient class
type TestClient struct {
	Opts       *MQTT.ClientOptions
	PahoClient MQTT.Client
	Message    MQTT.Message
	Seq        uint64
	Namespace  string
	GroupID    string
	EdgeNodeID string
	Stations   []*Station
}

// GetSequence adds 1 to sequence and gets the sequence number (cap of 255 then returns to 0)
func (testClient *TestClient) GetSequence() uint64 {
	testClient.Seq = testClient.Seq + 1
	if testClient.Seq > 255 {
		testClient.Seq = 0
	}
	return testClient.Seq
}

// NewClient Creates a client
func NewClient(ip string, port string, clientID string) *TestClient {
	opts := MQTT.NewClientOptions().AddBroker("tcp://" + ip + ":" + port)
	opts.SetClientID(clientID)
	//opts.SetDefaultPublishHandler(messageHandler)
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		//fmt.Printf("Connected to server\n")
	}
	return &TestClient{
		Opts:       opts,
		PahoClient: client,
	}
}

// SendMessage Sends a message
func (testClient TestClient) SendMessage(topic string, message interface{}) {
	testClient.PahoClient.Publish(topic, 0, false, message)
}

// WriteNBirth writes the NBirth message
func (testClient *TestClient) WriteNBirth() {
	payLoad := new(spb.Payload)
	payLoad.Timestamp = proto.Uint64(uint64(time.Now().Unix()) * 1000)
	payLoad.Seq = proto.Uint64(0)
	// Create Metric
	metric := new(spb.Payload_Metric)
	// Add Device Name
	metric.Name = proto.String("bdSeq")
	// Add Value
	flt := new(spb.Payload_Metric_IntValue)
	flt.IntValue = 0
	metric.Value = flt

	// Add Datatype
	metric.Datatype = proto.Uint32(8)
	// Create Random Timestamp (subtract from current unix time)
	metric.Timestamp = proto.Uint64(uint64(time.Now().Unix()))
	// Append Metric to Payload
	payLoad.Metrics = append(payLoad.Metrics, metric)

	// Publish
	topic := testClient.Namespace + "/" + testClient.GroupID + "/" + "NBIRTH" + "/" + testClient.EdgeNodeID
	marshalledPayload, _ := proto.Marshal(payLoad)
	testClient.SendMessage(topic, marshalledPayload)
}

// WriteData writes a message for all devices in Devices.  msgType can be either DDATA or DBIRTH.
func (testClient *TestClient) WriteData(station *Station, msgType string) {
	payLoad := new(spb.Payload)
	payLoad.Timestamp = proto.Uint64(uint64(time.Now().Unix() * 1000))
	payLoad.Seq = proto.Uint64(testClient.GetSequence())
	//fmt.Println("Broadcast Station '" + station.name + "' Devices:" + fmt.Sprint(len(station.Devices)))
	for _, device := range station.Devices {
		// Create Metric
		metric := new(spb.Payload_Metric)
		// Add Device Name
		if device.analDigital == "analog" {
			metric.Name = proto.String("analog/" + device.deviceID)
		} else {
			metric.Name = proto.String("multistate/" + device.deviceID)
		}
		// Add Value
		if device.Value != nil {
			if fmt.Sprint(reflect.TypeOf(device.Value)) == "float32" {
				flt := new(spb.Payload_Metric_FloatValue)
				flt.FloatValue = device.Value.(float32)
				metric.Value = flt
			} else {
				bl := new(spb.Payload_Metric_BooleanValue)
				bl.BooleanValue = device.Value.(bool)
				metric.Value = bl
			}
		}
		// Add Datatype
		metric.Datatype = proto.Uint32(device.dataType)
		// Create Random Timestamp (subtract from current unix time)
		metric.Timestamp = proto.Uint64(uint64(time.Now().Unix() * 1000))

		// Create PropertySet
		propertySet := new(spb.Payload_PropertySet)
		// Value
		propertyValArray := []*spb.Payload_PropertyValue{}
		propertyValRef := new(spb.Payload_PropertyValue_IntValue)
		propertyValRef.IntValue = 192
		propertyVal := new(spb.Payload_PropertyValue)
		propertyVal.Type = proto.Uint32(3)
		propertyVal.Value = propertyValRef
		propertyValArray = append(propertyValArray, propertyVal)
		// Add to PropertySet
		propertySet.Keys = []string{"Quality"}
		propertySet.Values = propertyValArray
		metric.Properties = propertySet

		// Append Metric to Payload
		payLoad.Metrics = append(payLoad.Metrics, metric)
	}
	// Publish
	topic := testClient.Namespace + "/" + testClient.GroupID + "/" + msgType + "/" + testClient.EdgeNodeID + "/" + station.name
	marshalledPayload, _ := proto.Marshal(payLoad)
	testClient.SendMessage(topic, marshalledPayload)
}
