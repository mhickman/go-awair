package awair

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDevicesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/self/devices", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET got %s", r.Method)
		}

		fmt.Fprint(w, `{
    "devices": [
        {
            "name": "device name",
            "macAddress": "DEADBEEF",
            "preference": "ALLERGY",
            "roomType": "LIVING_ROOM",
            "deviceType": "awair-element",
            "spaceType": "HOME",
            "deviceUUID": "awair-element_9999",
            "deviceId": 9999
        }
    ]
}`)
	})

	expectedDevice := &Device{
		Name:       stringPtr("device name"),
		MACAddress: stringPtr("DEADBEEF"),
		Preference: stringPtr("ALLERGY"),
		RoomType:   stringPtr("LIVING_ROOM"),
		DeviceType: stringPtr("awair-element"),
		SpaceType:  stringPtr("HOME"),
		DeviceUUID: stringPtr("awair-element_9999"),
		DeviceID:   int32Ptr(9999),
	}

	devices, _, err := client.Devices.List(context.Background())

	assert.NoError(t, err)

	if len(devices) != 1 {
		t.Errorf("expected len(devices) to be 1, was %d", len(devices))
		t.FailNow()
	}

	device := devices[0]

	assert.Equal(t, *expectedDevice.Name, *device.Name)
}

func TestDevicesService_List_Empty(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/self/devices", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET got %s", r.Method)
		}

		fmt.Fprint(w, "{}")
	})

	devices, _, err := client.Devices.List(context.Background())

	if len(devices) > 0 {
		t.Error("expected no devices")
	}

	if err != nil {
		t.Error(err)
	}
}
