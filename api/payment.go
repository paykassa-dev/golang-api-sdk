package api

import (
	"encoding/json"
	"github.com/paykassa-dev/golang-api-sdk/dto"
	"net/http"
	"net/url"
	"strconv"
)

const apiBaseUrl = "https://paykassa.app/api/0.9/index.php"

type PaymentApiInterface interface {
	CheckBalance(request *dto.CheckBalanceRequest) *dto.CheckBalanceResponse
	MakePayment(request *dto.MakePaymentRequest) *dto.MakePaymentResponse
	GetTxidsOfInvoices(request *dto.GetTxidsOfInvoicesRequest) *dto.GetTxidsOfInvoicesResponse
	SetTest(test bool)
}

type PaymentApi struct {
	apiId  string
	apiKey string
	test   bool
}

func NewPaymentApi(apiId string, apiKey string) PaymentApiInterface {
	return &PaymentApi{apiId: apiId, apiKey: apiKey, test: false}
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

	data.Set("func", endpoint)
	data.Set("api_id", p.apiId)
	data.Set("api_key", p.apiKey)
	data.Set("test", strconv.FormatBool(p.test))

	return data
}

func (p *PaymentApi) SetApiId(id string) {
	p.apiId = id
}

func (p *PaymentApi) SetApiKey(key string) {
	p.apiKey = key
}

func (r *PaymentApi) SetTest(test bool) {
	r.test = test
}

func (p PaymentApi) GetTxidsOfInvoices(request *dto.GetTxidsOfInvoicesRequest) *dto.GetTxidsOfInvoicesResponse {
	response, err := p.makeHttpRequest("api_get_shop_txids", request)

	if err != nil {
		return &dto.GetTxidsOfInvoicesResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	responseObject := &dto.GetTxidsOfInvoicesResponse{}
	err = json.NewDecoder(response.Body).Decode(responseObject)

	if err != nil {
		return &dto.GetTxidsOfInvoicesResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	return responseObject
}
