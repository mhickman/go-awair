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
// It uses GET /v1/users/self/devicesResponse
func (u *UserService) ListDevices(ctx context.Context) ([]*Device, *http.Response, error) {
	req, err := u.client.NewRequest(http.MethodGet, "users/self/devices")

	if err != nil {
		return nil, nil, err
	}

	devicesResponse := new(listDevicesResponse)
	resp, err := u.client.Do(ctx, req, devicesResponse)

	if err != nil {
		return nil, nil, err
	}

	return devicesResponse.Devices, resp, err
}

type Usage struct {
	Scope *string `json:"scope,omitempty"`
	Usage *int32  `json:"usage,omitempty"`
}

type UserInfo struct {
	ID *string `json:"id,omitempty"`

	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`

	DOBDay   *int `json:"dobDay,omitempty"`
	DOBMonth *int `json:"dobMonth,omitempty"`
	DOBYear  *int `json:"dobYear,omitempty"`

	Usages []*Usage `json:"usages,omitempty"`
	Tier   *string  `json:"tier,omitempty"`
	Email  *string  `json:"email,omitempty"`
}

// UserInfo returns info about the user.
// It uses GET /v1/users/self
func (u *UserService) UserInfo(ctx context.Context) (*UserInfo, *http.Response, error) {
	req, err := u.client.NewRequest(http.MethodGet, "users/self")

	if err != nil {
		return nil, nil, err
	}

	userInfo := new(UserInfo)
	resp, err := u.client.Do(ctx, req, userInfo)

	if err != nil {
		return nil, nil, err
	}

	return userInfo, resp, err
}
