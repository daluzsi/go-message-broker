package domain

type Authorization struct {
	AuthorizationCategory     string  `json:"authorizationCategory"`
	AccountId                 uint64  `json:"accountId"`
	CardId                    uint64  `json:"cardId"`
	CustomerId                uint64  `json:"customerId"`
	AuthorizationCode         string  `json:"authorizationCode"`
	Caller                    string  `json:"caller"`
	Cid                       string  `json:"cid"`
	LocalAmount               float64 `json:"localAmount"`
	Mti                       string  `json:"mti"`
	AuthorizationResponseCode int32   `json:"authorizationResponseCode"`
	CustomResponseCode        string  `json:"customResponseCode"`
	MerchantName              string  `json:"merchantName"`
}
