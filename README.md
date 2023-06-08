
## Download
```golang
go get github.com/vatanyazilim/whatsapp-api-go
```



## Usage/Examples

```golang
// Connect
 wp := WhatsappConn("1N0PbCli93vRTHny2pINg-tkq8yUPedq")
```


```golang
// Save Device 
deviceResp, err := wp.AddDevice()
```

```golang
// Device Login Cotrol 
wp.DeviceLoginControl("1157316539")
```

```golang
// Device List 
devices, err := wp.DeviceList(IDeviceListBody{
Page: 1,
})
```


```golang
// Send Message
res, err := wp.SendMessage(&ISendMessage{
   Identifier: "905452716912.0:91@s.whatsapp.net",
   To: "+905452716912", 
   Message: "test message",
   })
})
```
```golang
//multi Message
res, err := wp.SendMultiMessage(&ISendMultiMessageReq{
 	Name:      "Test Message Title",
	DeviceID:  301,
 	Numbers:   "+905452716912,+905452716912,+905452716912,+905452716912",
 	TimePost:  "",
 	SendSpeed: 1,
        Message:   "test messages",
 })
```

