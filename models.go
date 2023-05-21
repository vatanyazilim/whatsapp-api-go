package whatsapptoplumesaj

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
