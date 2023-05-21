package whatsapptoplumesaj

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var baseUrl = "https://api.whatsapptoplumesaj.com"

type IWhatsapp interface {
	AddDevice() (AddDeviceResponse, error)
	SendMessage(msg *ISendMessage) (SendMessageResponse, error)
	DeviceList(page int) (IDeviceList, error)
}

type WpConn struct {
	ApiKey   string `json:"api_key"`
	Password string `json:"password"`
}

func WhatsappConn(data *WpConn) IWhatsapp {
	return WpConn{ApiKey: data.ApiKey, Password: data.Password}
}

func (wp WpConn) AddDevice() (AddDeviceResponse, error) {
	var response AddDeviceResponse
	url := baseUrl + "/whatsapp/customer/login"
	data := map[string]string{
		"api_key":  wp.ApiKey,
		"password": wp.Password,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return response, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (wp WpConn) SendMessage(msg *ISendMessage) (SendMessageResponse, error) {
	var response SendMessageResponse
	url := baseUrl + "/whatsapp/customer/sendmessage/single"
	type requestBody struct {
		ApiKey     string `json:"api_key"`
		Password   string `json:"password"`
		Phone      string `json:"phone"`
		Message    string `json:"message"`
		Identifier string `json:"identifier"`
		Type       uint   `json:"type"`
	}
	jsonData, err := json.Marshal(requestBody{
		ApiKey:     wp.ApiKey,
		Password:   wp.Password,
		Phone:      msg.To,
		Message:    msg.Message,
		Identifier: msg.Identifier,
		Type:       1,
	})
	if err != nil {
		return response, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (wp WpConn) DeviceList(page int) (IDeviceList, error) {
	var response IDeviceList
	url := baseUrl + "/whatsapp/customer/devices"
	type requestBody struct {
		ApiKey   string `json:"api_key"`
		Password string `json:"password"`
		Page     int    `json:"page"`
	}
	jsonData, err := json.Marshal(requestBody{
		ApiKey:   wp.ApiKey,
		Password: wp.Password,
		Page:     page,
	})
	if err != nil {
		return response, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, err
}
