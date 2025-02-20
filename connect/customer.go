package connect

import (
	"encoding/json"
	"fmt"

	"github.com/theNoobExpert/icicibreeze/pkg/config"
	"github.com/theNoobExpert/icicibreeze/pkg/utils"
)

type CustomerDetailsResponse struct {
	Status  int             `json:"Status"`
	Error   *string         `json:"Error"` // Pointer to handle null values
	Success CustomerDetails `json:"Success"`
}

type CustomerDetails struct {
	ExgTradeDate     ExchangeTradeDate `json:"exg_trade_date"`
	ExgStatus        ExchangeStatus    `json:"exg_status"`
	SegmentsAllowed  SegmentsAllowed   `json:"segments_allowed"`
	IDirectUserid    string            `json:"idirect_userid"`
	SessionToken     string            `json:"session_token"`
	IDirectUserName  string            `json:"idirect_user_name"`
	IDirectLastLogin string            `json:"idirect_lastlogin_time"`
}

type ExchangeTradeDate struct {
	NSE string `json:"NSE"`
	BSE string `json:"BSE"`
	FNO string `json:"FNO"`
	NDX string `json:"NDX"`
}

type ExchangeStatus struct {
	NSE string `json:"NSE"`
	BSE string `json:"BSE"`
	FNO string `json:"FNO"`
	NDX string `json:"NDX"`
}

type SegmentsAllowed struct {
	Trading     string `json:"Trading"`
	Equity      string `json:"Equity"`
	Derivatives string `json:"Derivatives"`
	Currency    string `json:"Currency"`
}

func (brc *BreezeConnect) GetCustomerDetails() (*CustomerDetailsResponse, error) {
	payload, _ := json.Marshal(map[string]string{
		"SessionToken": brc.ApiSessionKey,
		"AppKey":       brc.AppKey,
	})

	body := string(payload)

	request := &BreezeRequest{
		Method:  config.HTTP_GET,
		URL:     config.API_URL + string(config.ENDPOINT_CUST_DETAILS),
		Body:    body,
		Headers: brc.GenerateHeaders(body, ""),
	}

	response, err := brc.MakeRequest(request)
	if err != nil {
		return nil, fmt.Errorf("error while getting customer details: %w", err)
	}

	var customerDetails CustomerDetailsResponse
	err = json.Unmarshal(response, &customerDetails)
	if err != nil {
		return nil, fmt.Errorf("error parsing customer details: %w", err)
	}

	return &customerDetails, nil
}

//////////////////////////////////////////////////////////////////

func (brc *BreezeConnect) InitSessionToken() (*CustomerDetailsResponse, error) {
	customerDetails, err := brc.GetCustomerDetails()
	if err != nil {
		return nil, fmt.Errorf("error while getting customer details: %w", err)
	}

	brc.ApiSessionToken = customerDetails.Success.SessionToken
	brc.IsClientInitialized = true

	err = utils.Validate.Struct(brc)
	if err != nil {
		return nil, fmt.Errorf("breeze client validation error: %w", err)
	}

	return customerDetails, nil
}

//////////////////////////////////////////////////////////////////

type CustomerDematResponse struct {
	Status  int                    `json:"Status"`
	Error   json.RawMessage        `json:"Error"`
	Success []CustomerDematDetails `json:"Success"`
}

type CustomerDematDetails struct {
	StockCode              string `json:"stock_code"`
	StockISIN              string `json:"stock_ISIN"`
	Quantity               string `json:"quantity"`
	DematTotalBulkQuantity string `json:"demat_total_bulk_quantity"`
	DematAvailQuantity     string `json:"demat_avail_quantity"`
	BlockedQuantity        string `json:"blocked_quantity"`
	DematAllocatedQuantity string `json:"demat_allocated_quantity"`
}

func (brc *BreezeConnect) GetDematHoldings() (*CustomerDematResponse, error) {

	response, err := brc.MakeRequestWithTokens(config.HTTP_GET, config.ENDPOINT_DEMAT_HOLDING, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error while getting demat holdings: %w", err)
	}

	var dematHoldings CustomerDematResponse
	err = json.Unmarshal(response, &dematHoldings)
	if err != nil {
		return nil, fmt.Errorf("error parsing demat details response: %w. %s", err, string(response))
	}

	return &dematHoldings, nil
}

//////////////////////////////////////////////////////////////////

type CustomerFundsResponse struct {
	Status  int             `json:"Status"`
	Error   json.RawMessage `json:"Error"`
	Success FundDetails     `json:"Success"`
}

type FundDetails struct {
	BankAccount           string  `json:"bank_account"`
	TotalBankBalance      float64 `json:"total_bank_balance"`
	AllocatedEquity       float64 `json:"allocated_equity"`
	AllocatedFNO          float64 `json:"allocated_fno"`
	AllocatedCommodity    float64 `json:"allocated_commodity"`
	AllocatedCurrency     float64 `json:"allocated_currency"`
	BlockByTradeEquity    float64 `json:"block_by_trade_equity"`
	BlockByTradeFNO       float64 `json:"block_by_trade_fno"`
	BlockByTradeCommodity float64 `json:"block_by_trade_commodity"`
	BlockByTradeCurrency  float64 `json:"block_by_trade_currency"`
	BlockByTradeBalance   float64 `json:"block_by_trade_balance"`
	UnallocatedBalance    string  `json:"unallocated_balance"` // String in response body
}

func (brc *BreezeConnect) GetFunds() (*CustomerFundsResponse, error) {

	response, err := brc.MakeRequestWithTokens(config.HTTP_GET, config.ENDPOINT_FUND, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error while getting customer funds: %w", err)
	}

	var customerFunds CustomerFundsResponse
	err = json.Unmarshal(response, &customerFunds)
	if err != nil {
		return nil, fmt.Errorf("error parsing customer funds response: %w. %s", err, string(response))
	}

	return &customerFunds, nil
}
