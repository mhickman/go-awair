package awair

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserService_ListDevices(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/self/devices", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

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

	devices, _, err := client.User.ListDevices(context.Background())

	assert.NoError(t, err)

	if len(devices) != 1 {
		t.Errorf("expected len(devices) to be 1, was %d", len(devices))
		t.FailNow()
	}

	device := devices[0]

	assert.Equal(t, expectedDevice, device)
}

func TestUserService_ListDevices_Empty(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/self/devices", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

		_, _ = fmt.Fprint(w, "{}")
	})

	devices, _, err := client.User.ListDevices(context.Background())

	if len(devices) > 0 {
		t.Error("expected no devices")
	}

	if err != nil {
		t.Error(err)
	}
}

func TestUserService_GetUserInfo(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/self", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

		fmt.Fprint(w, `{
    "dobDay": 1,
    "usages": [
        {
            "scope": "USER_DEVICE_LIST",
            "usage": 14
        },
        {
            "scope": "USER_INFO",
            "usage": 1
        }
    ],
    "tier": "Hobbyist",
    "email": "the@email.com",
    "dobYear": 1955,
    "permissions": [
        {
            "scope": "FIFTEEN_MIN",
            "quota": 100
		}
    ],
    "dobMonth": 2,
    "sex": "UNKNOWN",
    "lastName": "Smith",
    "firstName": "John",
    "id": "12345"
}`)
	})

	expectedUserInfo := &UserInfo{
		ID:        stringPtr("12345"),
		FirstName: stringPtr("John"),
		LastName:  stringPtr("Smith"),
		DOBDay:    intPtr(1),
		DOBMonth:  intPtr(2),
		DOBYear:   intPtr(1955),
		Usages: []*Usage{
			{
				Scope: stringPtr("USER_DEVICE_LIST"),
				Usage: int32Ptr(14),
			},
			{
				Scope: stringPtr("USER_INFO"),
				Usage: int32Ptr(1),
			},
		},
		Tier:  stringPtr("Hobbyist"),
		Email: stringPtr("the@email.com"),
	}

	info, _, err := client.User.GetUserInfo(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedUserInfo, info)
}

func TestUserService_ListDeviceAPIUsages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/self/devices/device-type/device-id/api-usages", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

		_, _ = fmt.Fprint(w, `{
  "usages": [
    {
      "scope": "FIFTEEN_MIN",
      "usage": 1
    }
  ]
}`)
	})

	expectedUsages := []*Usage{
		{
			Scope: stringPtr("FIFTEEN_MIN"),
			Usage: int32Ptr(1),
		},
	}

	usages, _, err := client.User.ListDeviceAPIUsages(context.Background(), "device-type", "device-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsages, usages)
}
