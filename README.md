
## Download
```golang
go get github.com/vatanyazilim/whatsapp-api-go
```



## Usage/Examples

```golang
// Connect
 wp := WhatsappConn(&WpConn{
 ApiKey: "1N0PbCli93vRTHny2pINg-tkq8yUPedq",
  Password: "123456",
})
```


```golang
// Save Device 
deviceResp, err := wp.AddDevice()
})
```

```golang
// Send Message
 wp.SendMessage(&ISendMessage{
   Identifier: "905452716912.0:91@s.whatsapp.net",
   To: "+905452716912", 
   Message: "test message",
   })
})
```

