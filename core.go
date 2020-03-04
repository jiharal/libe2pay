package e2pay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// CoreGW is a core of payment gateway e2pay.
type CoreGW struct {
	Client Client
}

const (
	grantType             = "client_credentials"
	grantTypeRefreshToken = "refresh_token"

	authH2H                           = "rest/oauth/token"
	hostVerifyUsername                = "/b2b/customer/auth/verifyUsername"
	hostCustomerRegistration          = "/b2b/customer/register"
	hostCompleteRegistrationH2H       = "/b2b/customer/register/confirm"
	hostResendSMSToken                = "/b2b/customer/resend/registration/sms"
	hostCustomerAccount               = "/b2b/customer/me/account"
	hostCustomerTransactionInquiry    = "/b2b/customer/me/transaction/inquiry"
	hostCustomerTransactionRequest    = "/b2b/customer/me/transaction/request"
	hostCustomerTransferoutInquiry    = "/b2b/customer/me/ibft/inquiry"
	hostCustomerTransactionRequestH2H = "/b2b/customer/me/transaction/request"
	hostCustomerTransactionConfirm    = "/b2b/customer/me/transaction/confirm"
)

// Call is ...
func (gw *CoreGW) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = gw.Client.Host + path
	return gw.Client.Call(method, path, header, body, v)
}

// AuthH2H is used to authorize host to host.
func (gw *CoreGW) AuthH2H() (res Token, err error) {
	data := url.Values{}
	data.Set("client_id", gw.Client.ClientID)
	data.Set("client_secret", gw.Client.SecretKey)
	data.Set("grant_type", grantType)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err = gw.Call(http.MethodPost, authH2H, headers, strings.NewReader(data.Encode()), &res)
	if err != nil {
		return
	}
	return
}

// RefreshToken is ...
func (gw *CoreGW) RefreshToken(token, redirectURI string) (res RefreshToken, err error) {
	data := url.Values{}
	data.Set("client_id", gw.Client.ClientID)
	data.Set("client_secret", gw.Client.SecretKey)
	data.Set("grant_type", grantTypeRefreshToken)
	data.Set("refresh_token", token)
	data.Set("redirect_uri", redirectURI)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err = gw.Call(http.MethodPost, authH2H, headers, strings.NewReader(data.Encode()), &res)
	if err != nil {
		return
	}
	return
}

// VirifyUsername is ...
func (gw *CoreGW) VirifyUsername(username, token string) (res VerifyUsername, err error) {
	uri := fmt.Sprintf("%s?username=%s", hostVerifyUsername, username)
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	err = gw.Call(http.MethodGet, uri, headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// CustomerRegistration is ...
func (gw *CoreGW) CustomerRegistration(data *CustomerRegistrationRequest, token string) (res CustomerRegistrationResponse, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	log.Println("Customer reg", string(jsonData))
	reqData := bytes.NewBuffer(jsonData)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	err = gw.Call(http.MethodPost, hostCustomerRegistration, headers, reqData, &res)
	if err != nil {
		return
	}
	return
}

// CompleteRegistrationH2H is ...
func (gw *CoreGW) CompleteRegistrationH2H(data *CompleteRegistrationH2HRequest, token string) (err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	log.Println("Complete customer reg", string(jsonData))
	reqData := bytes.NewBuffer(jsonData)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	err = gw.Call(http.MethodPost, hostCompleteRegistrationH2H, headers, reqData, nil)
	if err != nil {
		return
	}
	return
}

// ResendSMSToken is ...
func (gw *CoreGW) ResendSMSToken(req *ResendSMSTokenRequest, token string) (err error) {
	data := ResendSMSTokenRequest{
		Username: req.Username,
		Phone:    req.Phone,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	log.Println("Resend SMS Token", string(jsonData))
	reqData := bytes.NewBuffer(jsonData)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	err = gw.Call(http.MethodPost, hostResendSMSToken, headers, reqData, nil)
	if err != nil {
		return
	}
	return
}

// MerchantRegistrationByEmail is ...
func (gw *CoreGW) MerchantRegistrationByEmail(req *MerchantRegistrationRequest, token string) (res MerchantRegistration, err error) {
	merchantReq := MerchantRegistrationRequest{
		Phone:     req.Phone,
		Name:      req.Name,
		Email:     req.Email,
		PartnerID: req.PartnerID,
		CountryID: req.CountryID,
	}
	reqJSON, err := json.Marshal(merchantReq)
	if err != nil {
		return
	}
	log.Println("request: ", string(reqJSON))
	requestReader := bytes.NewBuffer(reqJSON)

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	err = gw.Call(http.MethodPost, authH2H, headers, requestReader, &res)
	if err != nil {
		return
	}
	return
}

// CustomerAccount is ...
func (gw *CoreGW) CustomerAccount(token string) (res CustomerAccountResponse, err error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	err = gw.Call(http.MethodGet, hostCustomerAccount, headers, nil, &res)
	fmt.Println(res)
	if err != nil {
		return
	}
	return
}

// CustomerTransactionInquiry is ...
func (gw *CoreGW) CustomerTransactionInquiry(accountID, phone, token string) (res CustomerTransactionInquiryResponse, err error) {
	uri := fmt.Sprintf("%s?accountID=%s&phone=%s", hostCustomerTransactionInquiry, accountID, phone)
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	err = gw.Call(http.MethodGet, uri, headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// CustomerTransactionRequest is ...
func (gw *CoreGW) CustomerTransactionRequest(req *CustomerTransactionRequest, token string) (res CustomerTransactionRequestResponse, err error) {
	data := CustomerTransactionRequest{
		AccountSrc:        req.AccountSrc,
		AccountDst:        req.AccountDst,
		Amount:            req.Amount,
		TransactionCodeID: req.TransactionCodeID,
		Description:       req.Description,
		ClientRef:         req.ClientRef,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	log.Println("CustomerTransactionRequest", string(jsonData))
	reqData := bytes.NewBuffer(jsonData)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	err = gw.Call(http.MethodPost, hostCustomerTransactionRequest, headers, reqData, nil)
	if err != nil {
		return
	}
	return
}

// CustomerTransferoutInquiry is ...
func (gw *CoreGW) CustomerTransferoutInquiry(bankID, accountID, token string) (res CustomerTransferoutInquiryResponse, err error) {
	uri := fmt.Sprintf("%s?bankID=%s&accountID=%s", hostCustomerTransferoutInquiry, bankID, accountID)
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	err = gw.Call(http.MethodPost, uri, headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// CustomerTransactionRequestH2H is ...
func (gw *CoreGW) CustomerTransactionRequestH2H(req CustomerTransactionRequestH2HRequest, token string) (res CustomerTransactionRequestH2HResponse, err error) {
	data := CustomerTransactionRequestH2HRequest{
		AccountSrc:        req.AccountSrc,
		AccountDst:        req.AccountDst,
		Amount:            req.Amount,
		TransactionCodeID: req.TransactionCodeID,
		Description:       req.Description,
		ClientRef:         req.ClientRef,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	log.Println("CustomerTransactionRequest", string(jsonData))
	reqData := bytes.NewBuffer(jsonData)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	err = gw.Call(http.MethodPost, hostCustomerTransactionRequestH2H, headers, reqData, &res)
	if err != nil {
		return
	}
	return
}

// CustomerTransactionConfirm is ...
func (gw *CoreGW) CustomerTransactionConfirm(req *CustomerTransactionConfirmRequest, token string) (res CustomerTransactionConfirmResponse, err error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return
	}
	log.Println("CustomerTransactionRequest", string(jsonData))
	reqData := bytes.NewBuffer(jsonData)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	err = gw.Call(http.MethodPost, hostCustomerTransactionConfirm, headers, reqData, &res)
	if err != nil {
		return
	}
	return
}
