// connect
wp := WhatsappConn(&WpConn{
		ApiKey:   "1N0PbCli93vRTHny2pINg-tkq8yUPedq",
		Password: "123456",
	})
  
//save device
	deviceResp, err := wp.AddDevice()

//Devices list
	devices, err := wp.DeviceList(1)

//Send Message
	res, err := wp.SendMessage(&ISendMessage{
		Identifier: "905452716912.0:91@s.whatsapp.net",
		To:         "+905452716912",
		Message:    "test message",
	})
