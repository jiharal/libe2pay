# e2pay

this lib is used to integrate to e2pay

# How to use

```go
  client := NewClient()
  client.Host = "some-url"
  client.ClientID = "some-key"
  client.SecretKey = "some-secret"
  client.SourceID = "123"      // your source id
  client.PartnerID = "2201121" // your patner id
  client.LogLevel = 3          // your log level default 2
  gw := CoreGW{
  	Client: client,
  }
  resp, err := gw.AuthH2H()
```
