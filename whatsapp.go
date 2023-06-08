package whatsapptoplumesaj

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://api.whatsapptoplumesaj.com"
)

type WhatsappClient struct {
	apiKey string
}

func NewWhatsappClient(apiKey string) *WhatsappClient {
	return &WhatsappClient{
		apiKey: apiKey,
	}
}

func (wp *WhatsappClient) AddDevice() (AddDeviceResponse, error) {

	url := baseURL + "/whatsapp/customer/login"
	data := map[string]string{
		"api_key": wp.apiKey,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return AddDeviceResponse{}, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return AddDeviceResponse{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AddDeviceResponse{}, err
	}
	var response AddDeviceResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return AddDeviceResponse{}, err
	}
	code := FirstNumber(resp.StatusCode)
	if code != 2 {
		return AddDeviceResponse{}, errors.New(response.Description)
	}
	return response, err
}

func (wp *WhatsappClient) SendMessage(msg *ISendMessage) (SendMessageResponse, error) {
	url := baseURL + "/whatsapp/customer/sendmessage/single"
	requestBody := struct {
		ApiKey     string `json:"api_key"`
		Phone      string `json:"phone"`
		Message    string `json:"message"`
		Identifier string `json:"identifier"`
		Type       uint   `json:"type"`
	}{
		ApiKey:     wp.apiKey,
		Phone:      msg.To,
		Message:    msg.Message,
		Identifier: msg.Identifier,
		Type:       1,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return SendMessageResponse{}, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return SendMessageResponse{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SendMessageResponse{}, err
	}
	var response SendMessageResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return SendMessageResponse{}, err
	}
	code := FirstNumber(resp.StatusCode)
	if code != 2 {
		return SendMessageResponse{}, errors.New(response.Description)
	}
	return response, nil
}
func (wp *WhatsappClient) SendMultiMessage(msg *ISendMultiMessageReq) (SendMessageResponse, error) {
	url := baseURL + "/campaign/customer"
	var isNow bool
	if len(msg.TimePost) > 0 {
		isNow = false
	} else {
		isNow = true
	}
	requestBody := ISendMultiMessage{
		ApiKey:    wp.apiKey,
		Name:      msg.Name,
		DeviceID:  msg.DeviceID,
		Numbers:   msg.Numbers,
		Message:   msg.Message,
		TimePost:  msg.TimePost,
		Type:      1,
		Now:       isNow,
		SendSpeed: msg.SendSpeed,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return SendMessageResponse{}, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return SendMessageResponse{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SendMessageResponse{}, err
	}
	var response SendMessageResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return SendMessageResponse{}, err
	}
	code := FirstNumber(resp.StatusCode)
	if code != 2 {
		return SendMessageResponse{}, errors.New(response.Description)
	}
	return response, nil
}

func (wp *WhatsappClient) DeviceList(reqBody IDeviceListBody) (IDeviceList, error) {
	url := baseURL + "/whatsapp/customer/devices"
	requestBody := struct {
		ApiKey string `json:"api_key"`
		Page   int    `json:"page"`
		Type   int    `json:"type"`
	}{
		ApiKey: wp.apiKey,
		Page:   reqBody.Page,
		Type:   1,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return IDeviceList{}, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return IDeviceList{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IDeviceList{}, err
	}
	var response IDeviceList
	err = json.Unmarshal(body, &response)
	if err != nil {
		return IDeviceList{}, err
	}
	code := FirstNumber(resp.StatusCode)
	if code != 2 {
		return IDeviceList{}, errors.New(response.Description)
	}
	return response, nil
}

func (wp *WhatsappClient) DeviceLoginControl(registrationID string) (bool, error) {
	url := baseURL + "/whatsapp/customer/ws/logincontrol?registrationId=" + registrationID
	requestBody := struct {
		ApiKey string `json:"api_key"`
	}{
		ApiKey: wp.apiKey,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return false, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var response IDeviceLoginControl
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}
	code := FirstNumber(resp.StatusCode)
	if code != 2 {
		return false, errors.New(response.Description)
	} else {
		return true, nil
	}
}
