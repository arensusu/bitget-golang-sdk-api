package mix

import (
	"encoding/json"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common/domain"
)

type MixMarketClient struct {
	BitgetRestClient domain.RestClient
}

func NewMixMarketClient(c domain.RestClient) *MixMarketClient {
	return &MixMarketClient{BitgetRestClient: c}
}

/**
 * Contract information
 * @param productType
 * @return ResponseResult
 */
func (p *MixMarketClient) Contracts(productType string) (string, error) {

	params := internal.NewParams()
	params["productType"] = productType

	uri := constants.MixMarket + "/contracts"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err

}

/**
 * Deep market
 * @param symbol
 * @param limit
 * @return ResponseResult
 */
func (p *MixMarketClient) Depth(symbol string, limit string) (string, error) {

	params := internal.NewParams()
	params["symbol"] = symbol
	params["limit"] = limit

	uri := constants.MixMarket + "/depth"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}

type GetTickersResponse struct {
	common.CommonResponse
	Data []GetTickerData
}

type GetTickerData struct {
	Symbol             string `json:"symbol"`
	Last               string `json:"last"`
	BestAsk            string `json:"bestAsk"`
	BestBid            string `json:"bestBid"`
	BidSz              string `json:"bidSz"`
	AskSz              string `json:"askSz"`
	High24h            string `json:"high24h"`
	Low24h             string `json:"low24h"`
	Timestamp          string `json:"timestamp"`
	PriceChangePercent string `json:"priceChangePercent"`
	BaseVolume         string `json:"baseVolume"`
	QuoteVolume        string `json:"quoteVolume"`
	UsdtVolume         string `json:"usdtVolume"`
	OpenUtc            string `json:"openUtc"`
	ChgUtc             string `json:"chgUtc"`
	IndexPrice         string `json:"indexPrice"`
	FundingRate        string `json:"fundingRate"`
	HoldingAmount      string `json:"holdingAmount"`
	DeliveryStatus     string `json:"deliveryStatus"`
}

/**
 * Acquisition of single ticker market
 * @param productType
 * @return ResponseResult
 */
func (s *MixMarketClient) GetTickers(productType string) (GetTickersResponse, error) {
	var res GetTickersResponse

	params := internal.NewParams()
	params["productType"] = productType

	uri := constants.MixMarket + "/tickers"

	resp, err := s.BitgetRestClient.DoGet(uri, params)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal([]byte(resp), &res); err != nil {
		return res, err
	}

	return res, nil
}

type GetTickerResponse struct {
	common.CommonResponse
	Data GetTickerData
}

/**
 * Deep market
 * @param symbol
 * @return ResponseResult
 */
func (s *MixMarketClient) GetTicker(symbol string) (GetTickerResponse, error) {
	var res GetTickerResponse

	params := internal.NewParams()
	params["symbol"] = symbol

	uri := constants.MixMarket + "/ticker"

	resp, err := s.BitgetRestClient.DoGet(uri, params)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal([]byte(resp), &res); err != nil {
		return res, err
	}

	return res, nil
}

/**
 * Obtain transaction details
 * @param symbol
 * @param limit
 * @return ResponseResult
 */
func (p *MixMarketClient) Fills(symbol string, limit string) (string, error) {

	params := internal.NewParams()
	params["symbol"] = symbol
	params["limit"] = limit

	uri := constants.MixMarket + "/fills"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}

/**
 * Obtain K line data
 * @param symbol
 * @param granularity (Category of k line)
 * @param startTime
 * @param endTime
 * @return ResponseResult
 */
func (p *MixMarketClient) Candles(symbol string, granularity string, startTime string, endTime string) (string, error) {

	params := internal.NewParams()
	params["symbol"] = symbol
	params["granularity"] = granularity
	params["startTime"] = startTime
	params["endTime"] = endTime

	uri := constants.MixMarket + "/candles"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}

/*
*
获取币种指数。
*/
func (p *MixMarketClient) Index(symbol string) (string, error) {

	params := internal.NewParams()
	params["symbol"] = symbol

	uri := constants.MixMarket + "/index"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}

type GetNextFundingTimeResponse struct {
	common.CommonResponse
	Data GetNextFundingTimeData
}

type GetNextFundingTimeData struct {
	Symbol      string `json:"symbol"`
	FundingTime string `json:"fundingTime"`
	RatePeriod  string `json:"ratePeriod"`
}

/**
 * Get the next settlement time of the contract
 * @param symbol
 * @return ResponseResult
 */
func (s *MixMarketClient) GetNextFundingTime(symbol string) (GetNextFundingTimeResponse, error) {
	var res GetNextFundingTimeResponse

	params := internal.NewParams()
	params["symbol"] = symbol

	uri := constants.MixMarket + "/funding-time"

	resp, err := s.BitgetRestClient.DoGet(uri, params)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(resp), &res)

	return res, err
}

/**
 * Get contract tag price
 * @param symbol
 * @return ResponseResult
 */
func (p *MixMarketClient) MarkPrice(symbol string) (string, error) {

	params := internal.NewParams()
	params["symbol"] = symbol

	uri := constants.MixMarket + "/mark-price"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}

/**
 * Get historical fund rate
 * @param symbol
 * @param pageSize
 * @param pageNo
 * @param nextPage
 * @return ResponseResult
 */
func (p *MixMarketClient) HistoryFundRate(symbol string, pageSize string, pageNo string, nextPage string) (string, error) {

	params := internal.NewParams()
	params["symbol"] = symbol
	if len(pageSize) > 0 {
		params["pageSize"] = pageSize
	}
	if len(pageNo) > 0 {
		params["pageNo"] = pageNo
	}
	if len(nextPage) > 0 {
		params["nextPage"] = nextPage
	}

	uri := constants.MixMarket + "/history-fundRate"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}

/**
 * Get the current fund rate
 * @param symbol
 * @return ResponseResult
 */
func (p *MixMarketClient) CurrentFundRate(symbol string) (string, error) {

	params := internal.NewParams()
	params["symbol"] = symbol

	uri := constants.MixMarket + "/current-fundRate"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}

/**
 * Obtain the total position of the platform
 * @param symbol
 * @return ResponseResult
 */
func (p *MixMarketClient) OpenInterest(symbol string) (string, error) {

	params := internal.NewParams()
	params["symbol"] = symbol

	uri := constants.MixMarket + "/open-interest"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}
