package api

import (
	"encoding/json"
	"github.com/paykassa-dev/golang-api-sdk/dto"
	"net/http"
	"net/url"
)

const apiBaseUrl = "https://paykassa.app/api/0.5/index.php"

type PaymentApiInterface interface {
	CheckBalance(request *dto.CheckBalanceRequest) *dto.CheckBalanceResponse
	MakePayment(request *dto.MakePaymentRequest) *dto.MakePaymentResponse
	SetApiId(id string)
	SetApiKey(key string)
}

type PaymentApi struct {
	apiId  string
	apiKey string
}

func NewPaymentApi(apiId string, apiKey string) PaymentApiInterface {
	return &PaymentApi{apiId: apiId, apiKey: apiKey}
}

func (p PaymentApi) CheckBalance(request *dto.CheckBalanceRequest) *dto.CheckBalanceResponse {
	response, err := p.makeHttpRequest("api_get_shop_balance", request)

	if err != nil {
		return &dto.CheckBalanceResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	responseObject := &dto.CheckBalanceResponse{}
	err = json.NewDecoder(response.Body).Decode(responseObject)

	if err != nil {
		return &dto.CheckBalanceResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	return responseObject
}

func (p PaymentApi) MakePayment(request *dto.MakePaymentRequest) *dto.MakePaymentResponse {
	response, err := p.makeHttpRequest("api_payment", request)

	if err != nil {
		return &dto.MakePaymentResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	responseObject := &dto.MakePaymentResponse{}
	err = json.NewDecoder(response.Body).Decode(responseObject)

	if err != nil {
		return &dto.MakePaymentResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	return responseObject
}

func (p *PaymentApi) makeHttpRequest(endpoint string, request dto.Request) (*http.Response, error) {
	payload := p.createRequestPayload(endpoint, request)
	return http.PostForm(apiBaseUrl, payload)
}

func (p *PaymentApi) createRequestPayload(endpoint string, request dto.Request) url.Values {
	data := request.Normalize()
	payload := make(url.Values, len(data)+3)

	for key, value := range data {
		payload[key] = []string{value}
	}

	payload["func"] = []string{endpoint}
	payload["api_id"] = []string{p.apiId}
	payload["api_key"] = []string{p.apiKey}

	return payload
}

func (p *PaymentApi) SetApiId(id string) {
	p.apiId = id
}

func (p *PaymentApi) SetApiKey(key string) {
	p.apiKey = key
}
