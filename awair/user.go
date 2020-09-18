package awair

import (
	"context"
	"fmt"
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

// Usage represents usages of an API endpoint.
type Usage struct {
	Scope *string `json:"scope,omitempty"`
	Usage *int32  `json:"usage,omitempty"`
}

// UserInfo contains information about the requested user.
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

// GetUserInfo returns info about the user.
// It uses GET /v1/users/self
func (u *UserService) GetUserInfo(ctx context.Context) (*UserInfo, *http.Response, error) {
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

type listDeviceAPIUsagesResponse struct {
	Usages []*Usage `json:"usages,omitempty"`
}

// ListDeviceAPIUsages returns the Usages for a service deviceType + deviceID pair.
// It uses GET v1/users/self/devices/{{device_type}}/{{device_id}}/api-usages
func (u *UserService) ListDeviceAPIUsages(ctx context.Context, deviceType, deviceID string) ([]*Usage, *http.Response, error) {
	url := fmt.Sprintf("users/self/devices/%v/%v/api-usages", deviceType, deviceID)
	req, err := u.client.NewRequest(http.MethodGet, url)

	if err != nil {
		return nil, nil, err
	}

	usages := new(listDeviceAPIUsagesResponse)
	resp, err := u.client.Do(ctx, req, usages)

	if err != nil {
		return nil, nil, err
	}

	return usages.Usages, resp, err
}
