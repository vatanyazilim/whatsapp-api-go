package main

type WpConn struct {
	ApiKey string `json:"api_key"`
}
type AddDeviceResponse struct {
	Data struct {
		Code           string `json:"code"`
		RegistrationId int    `json:"registrationId"`
	}
	Description string `json:"description"`
}

type SendMessageResponse struct {
	Description string `json:"description"`
}

type ISendMessage struct {
	Identifier string `json:"identifier"`
	To         string `json:"to"`
	Message    string `json:"message"`
}

type ISendMultiMessage struct {
	ApiKey    string `json:"api_key"`
	Name      string `json:"name" form:"name" validate:"required"`
	DeviceID  uint   `json:"device_id"  form:"device_id" validate:"required"`
	Numbers   string `json:"numbers" form:"numbers" validate:"required"`
	Message   string `json:"message" form:"message" validate:"required"` // number1,number2,number3
	TimePost  string `json:"time_post" form:"time_post" validate:"required"`
	Type      uint   `json:"type" form:"type" validate:"required"` //1- message 2-image 3-doc
	Now       bool   `json:"now" form:"now" validate:"omitempty"`
	SendSpeed uint   `json:"send_speed" form:"send_speed" validate:"omitempty"` //1-slow 2- medium 3-fast
}
type ISendMultiMessageReq struct {
	Name      string `json:"name" form:"name" validate:"required"`
	DeviceID  uint   `json:"device_id"  form:"device_id" validate:"required"`
	Numbers   string `json:"numbers" form:"numbers" validate:"required"`
	Message   string `json:"message" form:"message" validate:"required"` // number1,number2,number3
	TimePost  string `json:"time_post" form:"time_post" validate:"required"`
	SendSpeed uint   `json:"send_speed" form:"send_speed" validate:"omitempty"` //1-slow 2- medium 3-fast
}

type IDeviceList struct {
	Description string
	Data        struct {
		Records           []Device
		AdditionalData    uint
		From              uint
		To                uint
		CountOfTotalItems uint
	}
}
type IDeviceLoginControl struct {
	Description string
	Data        string
}

type Device struct {
	ID              uint   `json:"id"`
	Identifier      string `json:"identifier"`
	RegistirationID int64  `json:"registiration_id"`
	Platform        string `json:"platform"`
	BusinessName    string `json:"business_name"`
	Username        string `json:"username"`
	State           uint   `json:"state"`
	ProfileUrl      string `json:"profile_url"`
}

type IDeviceListBody struct {
	Page int
}
