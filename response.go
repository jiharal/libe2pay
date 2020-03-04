package e2pay

import "time"

type (

	// Token is ...
	Token struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresID    string `json:"expires_id"`
		RefreshToken string `json:"tefresh_token"`
		Scope        string `json:"scope"`
	}

	// RefreshToken is ...
	RefreshToken struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   string `json:"expires_in"`
		Scope       string `json:"scope"`
	}

	// MerchantRegistration response
	MerchantRegistration struct {
		MerchantRegistrationID string `json:"merchant_registration_id"`
		TokenPrefix            string `json:"token_prefix"`
		Username               string `json:"username"`
		AccountGroupID         string `json:"account_group_id"`
	}

	// VerifyUsername is ...
	VerifyUsername struct {
		Exists bool `json:"exists"`
	}

	// CustomerRegistrationResponse struct
	CustomerRegistrationResponse struct {
		ExpiryMinutes    int    `json:"expiryMinutes"`
		ExpiryTimestamp  string `json:"expiryTimestamp"`
		RequestTimestamp string `json:"requestTimestamp"`
		Phone            string `json:"phone"`
		Active           bool   `json:"active"`
		TokenPrefix      string `json:"tokenPrefix"`
		ID               string `json:"id"`
		TokenType        string `json:"tokenType"`
		Attempt          int    `json:"attempt"`
		Username         string `json:"username"`
	}

	// CustomerAccountResponse is ...
	CustomerAccountResponse struct {
		UpgradeStatus         string    `json:"upgradeStatus"`
		CustomerStatus        string    `json:"customerStatus"`
		AccountID             string    `json:"accountId"`
		RegistrationTimestamp time.Time `json:"registrationTimestamp"`
		AccountTypeID         string    `json:"accountTypeId"`
		Balance               float64   `json:"balance"`
		AccountName           string    `json:"accountName"`
		AccountGroupID        string    `json:"accountGroupId"`
		Phone                 string    `json:"phone"`
		AccountGroupName      string    `json:"accountGroupName"`
		AccountTypeName       string    `json:"accountTypeName"`
	}

	// CustomerTransactionInquiryResponse is ...
	CustomerTransactionInquiryResponse struct {
		AccountID   string `json:"accountId"`
		AccountName string `json:"accountName"`
		AccountType string `json:"accountType"`
	}

	// CustomerTransactionRequestResponse struct
	CustomerTransactionRequestResponse struct {
		AccountSrc           string  `json:"accountSrc"`
		AccountDst           string  `json:"accountDst"`
		Amount               float64 `json:"amount"`
		TransactionCodeID    string  `json:"transactionCodeId"`
		Description          string  `json:"description"`
		TransactionRequestID string  `json:"transactionRequestId"`
	}

	// CustomerTransferoutInquiryResponse is ...
	CustomerTransferoutInquiryResponse struct {
		AccountID string  `json:"accountId"`
		Name      string  `json:"name"`
		InquiryID string  `json:"inquiryId"`
		FeeAmount float64 `json:"feeAmount"`
	}

	// CustomerTransactionRequestH2HResponse is ...
	CustomerTransactionRequestH2HResponse struct {
		AccountSrc           string  `json:"accountSrc"`
		AccountDst           string  `json:"accountDst"`
		Amount               float64 `json:"amount"`
		TransactionCodeID    string  `json:"transactionCodeId"`
		Description          string  `json:"description"`
		TransactionRequestID string  `json:"transactionRequestId"`
	}

	// CustomerTransactionConfirmResponse is ...
	CustomerTransactionConfirmResponse struct {
		AccountSrcID         string    `json:"accountSrcId"`
		AccountDstID         string    `json:"accountDstId"`
		Amount               float64   `json:"amount"`
		TransactionCodeID    string    `json:"transactionCodeId"`
		Description          string    `json:"description"`
		ClientRef            string    `json:"clientRef"`
		ClientTimestamp      time.Time `json:"clientTimestamp"`
		TransactionTimestamp time.Time `json:"transactionTimestamp"`
		JournalID            string    `json:"journalId"`
		AccountDstName       string    `json:"accountDstName"`
		AccountSrcName       string    `json:"accountSrcName"`
	}
)
