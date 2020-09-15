package awair

import (
	"context"
	"net/http"
)

// UserService has methods that correspond to the User API
// https://docs.developer.getawair.com/#0d57dee7-dc26-4bd5-877c-dc7f887712b1
type UserService service

// Device represents an Awair device.
type Device struct {
	Name       *string `json:"name,omitempty"`
	MACAddress *string `json:"macAddress,omitempty"`
	Preference *string `json:"preference,omitempty"`
	RoomType   *string `json:"roomType,omitempty"`
	DeviceType *string `json:"deviceType,omitempty"`
	SpaceType  *string `json:"spaceType,omitempty"`
	DeviceUUID *string `json:"deviceUUID,omitempty"`
	DeviceID   *int32  `json:"deviceId,omitempty"`
}

type listDevicesResponse struct {
	Devices []*Device `json:"devices"`
}

// ListDevices returns a list of devices belonging to the user.
// It uses /v1/users/self/devicesResponse
func (d *UserService) ListDevices(ctx context.Context) ([]*Device, *http.Response, error) {
	req, err := d.client.NewRequest(http.MethodGet, "users/self/devices")

	if err != nil {
		return nil, nil, err
	}

	devicesResponse := new(listDevicesResponse)
	resp, err := d.client.Do(ctx, req, devicesResponse)

	if err != nil {
		return nil, nil, err
	}

	return devicesResponse.Devices, resp, err
}
