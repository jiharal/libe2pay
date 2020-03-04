package e2pay

type (
	// MerchantRegistrationRequest is ...
	MerchantRegistrationRequest struct {
		Phone     string
		Name      string
		Email     string
		PartnerID string
		CountryID string
	}

	// CustomerRegistrationRequest is ...
	CustomerRegistrationRequest struct {
		Phone      string `json:"phone"`
		Name       string `json:"name"`
		BirthDate  string `json:"birthDate"`
		BirthPlace string `json:"birthPlace"`
		Email      string `json:"email"`
		Gender     bool   `json:"gender"`
		CountryID  string `json:"countryId"`
		SourceID   string `json:"sourceId"`
	}

	// CompleteRegistrationH2HRequest is ...
	CompleteRegistrationH2HRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}

	// ResendSMSTokenRequest is ..
	ResendSMSTokenRequest struct {
		Username string
		Phone    string
	}

	// CustomerTransactionRequest struct
	CustomerTransactionRequest struct {
		AccountSrc        string
		AccountDst        string
		Amount            float64
		TransactionCodeID string
		Description       string
		ClientRef         string
	}

	// CustomerTransactionRequestH2HRequest is ...
	CustomerTransactionRequestH2HRequest struct {
		AccountSrc        string
		AccountDst        string
		Amount            float64
		TransactionCodeID string
		Description       string
		ClientRef         string
	}

	// CustomerTransactionConfirmRequest is ...
	CustomerTransactionConfirmRequest struct {
		AccountSrc           string
		AccountDst           string
		Amount               float64
		TransactionCodeID    string
		TransactionRequestID string
		Password             string
		ClientRef            string
	}
)
