package awair

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

const testResponse = `
	{
    "devices": [
        {
            "name": "Living room",
            "macAddress": "70226B130EC1",
            "preference": "ALLERGY",
            "roomType": "LIVING_ROOM",
            "deviceType": "awair-element",
            "spaceType": "HOME",
            "deviceUUID": "awair-element_8631",
            "deviceId": 8631
        }
    ]
}
`

func TestDeviceSerialization(t *testing.T) {
	r := strings.NewReader(testResponse)

	var devicesResp listDevicesResponse
	err := json.NewDecoder(r).Decode(&devicesResp)

	fmt.Println(err)
	fmt.Println("hello World")
}
