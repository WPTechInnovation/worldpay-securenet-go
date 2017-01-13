package cardnotpresent

// ChargeTokenResponse is a response to a request to charge a token when card is not present
type ChargeTokenResponse struct {
	Result       string      `json:"result"`
	ResponseCode string      `json:"responseCode"`
	Messages     string      `json:"message"`
	Transaction  interface{} `json:"transaction"`
}
