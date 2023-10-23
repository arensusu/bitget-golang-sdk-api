package spot

import (
	"encoding/json"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common/domain"
)

type SpotAccountGetAccountAssetsLiteService struct {
	Client domain.RestClient
	coin   *string
}

func NewSpotAccountGetAccountAssetsLiteService(c domain.RestClient) *SpotAccountGetAccountAssetsLiteService {
	return &SpotAccountGetAccountAssetsLiteService{Client: c}
}

type SpotAccountGetAccountAssetsLiteResponse struct {
	common.CommonResponse
	Data []SpotAccountGetAccountAssetsLiteData `json:"data"`
}

type SpotAccountGetAccountAssetsLiteData struct {
	CoinId    int    `json:"coinId"`
	CoinName  string `json:"coinName"`
	Available string `json:"available"`
	Frozen    string `json:"frozen"`
	Lock      string `json:"lock"`
	UTime     string `json:"uTime"`
}

func (s *SpotAccountGetAccountAssetsLiteService) Coin(coin string) *SpotAccountGetAccountAssetsLiteService {
	s.coin = &coin
	return s
}

/**
 * Obtain account assets
 * @return ResponseResult
 */
func (s *SpotAccountGetAccountAssetsLiteService) Do() (SpotAccountGetAccountAssetsLiteResponse, error) {
	var res SpotAccountGetAccountAssetsLiteResponse

	params := internal.NewParams()
	if s.coin != nil {
		params["coin"] = *s.coin
	}

	uri := constants.SpotAccount + "/assets-lite"
	resp, err := s.Client.DoGet(uri, params)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal([]byte(resp), &res); err != nil {
		return res, err
	}

	return res, nil

}

type SpotAccountGetTransferRecordsService struct {
	Client   domain.RestClient
	coinId   *string
	fromType *string
	limit    *string
	after    *string
	before   *string
}

func NewSpotAccountGetTransferRecordsService(c domain.RestClient) *SpotAccountGetTransferRecordsService {
	return &SpotAccountGetTransferRecordsService{Client: c}
}

func (s *SpotAccountGetTransferRecordsService) CoinId(id string) *SpotAccountGetTransferRecordsService {
	s.coinId = &id
	return s
}

func (s *SpotAccountGetTransferRecordsService) FromType(fromType string) *SpotAccountGetTransferRecordsService {
	s.fromType = &fromType
	return s
}

func (s *SpotAccountGetTransferRecordsService) Limit(limit string) *SpotAccountGetTransferRecordsService {
	s.limit = &limit
	return s
}

func (s *SpotAccountGetTransferRecordsService) After(after string) *SpotAccountGetTransferRecordsService {
	s.after = &after
	return s
}

func (s *SpotAccountGetTransferRecordsService) Before(before string) *SpotAccountGetTransferRecordsService {
	s.before = &before
	return s
}

type SpotAccountGetTransferRecordsServiceResponse struct {
	common.CommonResponse
	Data []SpotAccountGetTransferRecordsServiceData `json:"data"`
}

type SpotAccountGetTransferRecordsServiceData struct {
	CoinName   string `json:"coinName"`
	Status     string `json:"status"`
	ToType     string `json:"toType"`
	ToSymbol   string `json:"toSymbol"`
	FromType   string `json:"fromType"`
	FromSymbol string `json:"fromSymbol"`
	Amount     string `json:"amount"`
	TradeTime  string `json:"tradeTime"`
	ClientOid  string `json:"clientOid"`
	TransferId string `json:"transferId"`
}

/**
 * Obtain transfer records
 * @param coinId
 * @param fromType
 * @param limit
 * @param after
 * @param before
 * @return ResponseResult
 */
func (s *SpotAccountGetTransferRecordsService) Do() (SpotAccountGetTransferRecordsServiceResponse, error) {
	var res SpotAccountGetTransferRecordsServiceResponse

	params := internal.NewParams()
	if s.coinId != nil {
		params["coinId"] = *s.coinId
	}
	if s.fromType != nil {
		params["fromType"] = *s.fromType
	}
	if s.limit != nil {
		params["limit"] = *s.limit
	}
	if s.after != nil {
		params["after"] = *s.after
	}
	if s.before != nil {
		params["before"] = *s.before
	}

	uri := constants.SpotAccount + "/transferRecords"

	resp, err := s.Client.DoGet(uri, params)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal([]byte(resp), &res); err != nil {
		return res, err
	}

	return res, nil
}

type SpotAccountGetBillsParam struct {
	CoinId    string `json:"coinId"`
	GroupType string `json:"groupType"`
	BizType   string `json:"bizType"`
	After     string `json:"after"`
	Before    string `json:"before"`
	Limit     string `json:"limit"`
}

type SpotAccountGetBillsService struct {
	Client domain.RestClient
	SpotAccountGetBillsParam
}

func NewSpotAccountGetBillsService(c domain.RestClient) *SpotAccountGetBillsService {
	return &SpotAccountGetBillsService{Client: c}
}

func (s *SpotAccountGetBillsService) CoinId(id string) *SpotAccountGetBillsService {
	s.SpotAccountGetBillsParam.CoinId = id
	return s
}

func (s *SpotAccountGetBillsService) GroupType(groupType string) *SpotAccountGetBillsService {
	s.SpotAccountGetBillsParam.GroupType = groupType
	return s
}

func (s *SpotAccountGetBillsService) BizType(bizType string) *SpotAccountGetBillsService {
	s.SpotAccountGetBillsParam.BizType = bizType
	return s
}

func (s *SpotAccountGetBillsService) After(id string) *SpotAccountGetBillsService {
	s.SpotAccountGetBillsParam.After = id
	return s
}

func (s *SpotAccountGetBillsService) Before(before string) *SpotAccountGetBillsService {
	s.SpotAccountGetBillsParam.Before = before
	return s
}

func (s *SpotAccountGetBillsService) Limit(limit string) *SpotAccountGetBillsService {
	s.SpotAccountGetBillsParam.Limit = limit
	return s
}

type SpotAccountGetBillsServiceResponse struct {
	common.CommonResponse
	Data []SpotAccountGetBillsServiceData `json:"data"`
}

type SpotAccountGetBillsServiceData struct {
	CTime     string `json:"cTime"`
	CoinId    string `json:"coinId"`
	CoinName  string `json:"coinName"`
	GroupType string `json:"groupType"`
	BizType   string `json:"bizType"`
	Quantity  string `json:"quantity"`
	Balance   string `json:"balance"`
	Fees      string `json:"fees"`
	BillId    string `json:"billId"`
}

/**
 * Get the bill flow
 * @param SpotAccountGetBillsParam
 * @return ResponseResult
 */
func (s *SpotAccountGetBillsService) Do() (SpotAccountGetBillsServiceResponse, error) {
	var res SpotAccountGetBillsServiceResponse

	postBody, jsonErr := internal.ToJson(s.SpotAccountGetBillsParam)
	if jsonErr != nil {
		return res, jsonErr
	}

	uri := constants.SpotAccount + "/bills"

	resp, err := s.Client.DoPost(uri, postBody)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal([]byte(resp), &res); err != nil {
		return res, err
	}

	return res, nil
}
