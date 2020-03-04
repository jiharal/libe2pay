package e2pay

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type E2Pay struct {
	suite.Suite
}

func TestE2Pay(t *testing.T) {
	suite.Run(t, new(E2Pay))
}

func InitTest() *CoreGW {
	client := NewClient()
	client.Host = "_"
	client.ClientID = "_"
	client.SecretKey = "_"
	client.SourceID = "_"  // your source id
	client.PartnerID = "_" // your patner id
	client.LogLevel = 3    // your log level default 2

	gw := &CoreGW{
		Client: client,
	}
	return gw
}

func (s *E2Pay) TestAuthH2H() {
	gw := InitTest()
	resp, err := gw.AuthH2H()
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *E2Pay) TestVerifyUsername() {
	gw := InitTest()
	username := "082325600xxx"
	resp, _ := gw.AuthH2H()
	a, err := gw.VirifyUsername(username, resp.AccessToken)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), a)
}

func (s *E2Pay) TestCustomerRegistration() {
	gw := InitTest()
	auth, _ := gw.AuthH2H()
	data := CustomerRegistrationRequest{
		Phone:      "082325600996",
		Name:       "Jihar",
		BirthDate:  "1997-06-28",
		BirthPlace: "Waro",
		Email:      "jihar@tlab.co.id",
		Gender:     true,
		CountryID:  "ID",
		SourceID:   "01011",
	}
	resp, err := gw.CustomerRegistration(&data, auth.AccessToken)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *E2Pay) TestCompleteRegistration() {
	gw := InitTest()
	auth, _ := gw.AuthH2H()
	data := CompleteRegistrationH2HRequest{
		Username: "082325600996",
		Password: HastToMD5("1234"),
		Token:    "buH3738",
	}
	err := gw.CompleteRegistrationH2H(&data, auth.AccessToken)
	assert.Nil(s.T(), err)
	// assert.NotNil(s.T(), auth)
}

func (s *E2Pay) TestCustomerAccount() {
	gw := InitTest()
	auth, _ := gw.AuthH2H()
	data, err := gw.CustomerAccount(auth.AccessToken)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), data)
}
