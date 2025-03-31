package api

import (
	"encoding/json"
	"github.com/paykassa-dev/golang-api-sdk/dto"
	"net/http"
	"net/url"
	"strconv"
)

const sciBaseUrl = "https://paykassa.app/sci/0.4/index.php"

type MerchantApiInterface interface {
	CheckPayment(request *dto.CheckPaymentRequest) *dto.CheckPaymentResponse
	CheckTransaction(request *dto.CheckTransactionRequest) *dto.CheckTransactionResponse
	GenerateAddress(request *dto.GenerateAddressRequest) *dto.GenerateAddressResponse
	GetPaymentUrl(request *dto.GetPaymentUrlRequest) *dto.GetPaymentUrlResponse
	SetTest(test bool)
}

type MerchantApi struct {
	sciId  string
	sciKey string
	test   bool
}

func NewMerchantApi(sciId string, sciKey string) MerchantApiInterface {
	return &MerchantApi{sciId: sciId, sciKey: sciKey, test: false}
}

func (m *MerchantApi) SetTest(test bool) {
	m.test = test
}

func (m *MerchantApi) CheckPayment(request *dto.CheckPaymentRequest) *dto.CheckPaymentResponse {
	response, err := m.makeHttpRequest("sci_confirm_order", request)

	if err != nil {
		return &dto.CheckPaymentResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	responseObject := &dto.CheckPaymentResponse{}
	err = json.NewDecoder(response.Body).Decode(responseObject)

	if err != nil {
		return &dto.CheckPaymentResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	return responseObject
}

func (m *MerchantApi) CheckTransaction(request *dto.CheckTransactionRequest) *dto.CheckTransactionResponse {
	response, err := m.makeHttpRequest("sci_confirm_transaction_notification", request)

	if err != nil {
		return &dto.CheckTransactionResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	responseObject := &dto.CheckTransactionResponse{}
	err = json.NewDecoder(response.Body).Decode(responseObject)

	if err != nil {
		return &dto.CheckTransactionResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	return responseObject
}

func (m *MerchantApi) GenerateAddress(request *dto.GenerateAddressRequest) *dto.GenerateAddressResponse {
	response, err := m.makeHttpRequest("sci_create_order_get_data", request)

	if err != nil {
		return &dto.GenerateAddressResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	responseObject := &dto.GenerateAddressResponse{}
	err = json.NewDecoder(response.Body).Decode(responseObject)

	if err != nil {
		return &dto.GenerateAddressResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	return responseObject
}

func (m *MerchantApi) GetPaymentUrl(request *dto.GetPaymentUrlRequest) *dto.GetPaymentUrlResponse {
	response, err := m.makeHttpRequest("sci_create_order", request)

	if err != nil {
		return &dto.GetPaymentUrlResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	responseObject := &dto.GetPaymentUrlResponse{}
	err = json.NewDecoder(response.Body).Decode(responseObject)

	if err != nil {
		return &dto.GetPaymentUrlResponse{
			Error:   true,
			Message: err.Error(),
		}
	}

	return responseObject
}

func (m *MerchantApi) SetSciId(id string) {
	m.sciId = id
}

func (m *MerchantApi) SetSciKey(key string) {
	m.sciKey = key
}

func (m *MerchantApi) makeHttpRequest(endpoint string, request dto.Request) (*http.Response, error) {
	payload := m.createRequestPayload(endpoint, request)
	return http.PostForm(sciBaseUrl, payload)
}

func (m *MerchantApi) createRequestPayload(endpoint string, request dto.Request) url.Values {
	data := request.Normalize()

	data.Set("func", endpoint)
	data.Set("sci_id", m.sciId)
	data.Set("sci_key", m.sciKey)
	data.Set("test", strconv.FormatBool(m.test))

	return data
}
