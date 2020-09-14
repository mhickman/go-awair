package awair

import (
	"context"
	"net/http"
)

type DevicesService service

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
	Devices []Device `json:"devices"`
}

func (d *DevicesService) List(ctx context.Context) ([]Device, *http.Response, error) {
	req, err := d.client.NewRequest("GET", "v1/users/self/devices")

	if err != nil {
		return nil, nil, err
	}

	devices := new(listDevicesResponse)
	resp, err := d.client.Do(ctx, req, devices)

	if err != nil {
		return nil, nil, err
	}

	var deviceResp []Device
	if devices != nil {
		deviceResp = devices.Devices
	}
	return deviceResp, resp, err
}
