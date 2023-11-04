package mix

import (
	"fmt"
	"testing"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMixMarketGetSymbols(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uri := constants.MixMarket + "/contracts"
	params := map[string]string{"productType": "umcbl"}

	data := `{
		"code": "00000",
		"msg": "success",
		"requestTime": "1698998744571",
		"data":[
			{
				"baseCoin": "TIA",
				"buyLimitPriceRatio": "0.02",
				"feeRateUpRatio": "0.005",
				"limitOpenTime": "-1",
				"maintainTime": "",
				"makerFeeRate": "0.0002",
				"maxOrderNum": "400",
				"maxPositionNum": "150",
				"minTradeNum": "0.1",
				"minTradeUSDT": "5",
				"offTime": "-1",
				"openCostUpRatio": "0.01",
				"priceEndStep": "1",
				"pricePlace": "4",
				"quoteCoin": "USDT",
				"sellLimitPriceRatio": "0.02",
				"sizeMultiplier": "0.1",
				"supportMarginCoins": [
					"USDT"
				],
				"symbol": "TIAUSDT_UMCBL",
				"symbolName": "TIAUSDT",
				"symbolStatus": "normal",
				"symbolType": "perpetual",
				"takerFeeRate": "0.0006",
				"volumePlace": "1"
			}
		]
	}`

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().DoGet(uri, params).Return(data, nil)

	expect := GetSymbolsResponse{
		CommonResponse: common.CommonResponse{
			Code: "00000",
			Msg:  "success",
			//RequestTime: "1698998744571",
		},
		Data: []GetSymbolsData{
			{
				BaseCoin:            "TIA",
				BuyLimitPriceRatio:  "0.02",
				FeeRateUpRatio:      "0.005",
				LimitOpenTime:       "-1",
				MaintainTime:        "",
				MakerFeeRate:        "0.0002",
				MaxOrderNum:         "400",
				MaxPositionNum:      "150",
				MinTradeNum:         "0.1",
				MinTradeUSDT:        "5",
				OffTime:             "-1",
				OpenCostUpRatio:     "0.01",
				PriceEndStep:        "1",
				PricePlace:          "4",
				QuoteCoin:           "USDT",
				SellLimitPriceRatio: "0.02",
				SizeMultiplier:      "0.1",
				SupportMarginCoins:  []string{"USDT"},
				Symbol:              "TIAUSDT_UMCBL",
				SymbolName:          "TIAUSDT",
				SymbolStatus:        "normal",
				SymbolType:          "perpetual",
				TakerFeeRate:        "0.0006",
				VolumePlace:         "1",
			},
		},
	}

	service := NewMixMarketService(mockClient)
	res, err := service.GetSymbols(params["productType"])

	assert.NoError(t, err)
	assert.Equal(t, expect, res)

}

func TestMixMarketClient_Depth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	client := NewMixMarketService(mockClient)

	resp, err := client.GetDepth("BTCUSDT_UMCBL", "20")

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestMixMarketClient_Ticker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uri := constants.MixMarket + "/ticker"
	params := map[string]string{"symbol": "BTCUSDT_UMCBL"}

	data := `{
		"code": "00000",
		"msg": "success",
		"requestTime": "1698145377845",
		"data":{
				"askSz": "0.268",
				"baseVolume": "235231.756",
				"bestAsk": "34483.2",
				"bestBid": "34483.1",
				"bidSz": "3.043",
				"chgUtc": "0.04191",
				"deliveryStatus": "normal",
				"fundingRate": "0.000095",
				"high24h": "35929.3",
				"holdingAmount": "49626.455",
				"indexPrice": "34487.575151",
				"last": "34483.1",
				"low24h": "30392.3",
				"openUtc": "33093.7",
				"priceChangePercent": "0.11682",
				"quoteVolume": "7728953739.0099",
				"symbol": "BTCUSDT_UMCBL",
				"timestamp": "1698145377845",
				"usdtVolume": "7728953739.0099"
		}
	}`

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().DoGet(uri, params).Return(data, nil)

	expect := GetTickerResponse{
		CommonResponse: common.CommonResponse{
			Code: "00000",
			Msg:  "success",
			//RequestTime: "1698145377845",
		},
		Data: GetTickerData{
			AskSz:              "0.268",
			BaseVolume:         "235231.756",
			BestAsk:            "34483.2",
			BestBid:            "34483.1",
			BidSz:              "3.043",
			ChgUtc:             "0.04191",
			DeliveryStatus:     "normal",
			FundingRate:        "0.000095",
			High24h:            "35929.3",
			HoldingAmount:      "49626.455",
			IndexPrice:         "34487.575151",
			Last:               "34483.1",
			Low24h:             "30392.3",
			OpenUtc:            "33093.7",
			PriceChangePercent: "0.11682",
			QuoteVolume:        "7728953739.0099",
			Symbol:             "BTCUSDT_UMCBL",
			Timestamp:          "1698145377845",
			UsdtVolume:         "7728953739.0099",
		},
	}

	service := NewMixMarketService(mockClient)
	res, err := service.GetTicker(params["symbol"])

	assert.NoError(t, err)
	assert.Equal(t, expect, res)

}

func TestMixMarketGetTickers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uri := constants.MixMarket + "/tickers"
	params := map[string]string{"productType": "umcbl"}

	data := `{
		"code": "00000",
		"msg": "success",
		"requestTime": "1698145377845",
		"data": [
			{
				"askSz": "0.268",
				"baseVolume": "235231.756",
				"bestAsk": "34483.2",
				"bestBid": "34483.1",
				"bidSz": "3.043",
				"chgUtc": "0.04191",
				"deliveryStatus": "normal",
				"fundingRate": "0.000095",
				"high24h": "35929.3",
				"holdingAmount": "49626.455",
				"indexPrice": "34487.575151",
				"last": "34483.1",
				"low24h": "30392.3",
				"openUtc": "33093.7",
				"priceChangePercent": "0.11682",
				"quoteVolume": "7728953739.0099",
				"symbol": "BTCUSDT_UMCBL",
				"timestamp": "1698145377845",
				"usdtVolume": "7728953739.0099"
			}
			]
			}`

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().DoGet(uri, params).Return(data, nil)

	expect := GetTickersResponse{
		CommonResponse: common.CommonResponse{
			Code: "00000",
			Msg:  "success",
			//RequestTime: "1698145377845",
		},
		Data: []GetTickerData{
			{
				AskSz:              "0.268",
				BaseVolume:         "235231.756",
				BestAsk:            "34483.2",
				BestBid:            "34483.1",
				BidSz:              "3.043",
				ChgUtc:             "0.04191",
				DeliveryStatus:     "normal",
				FundingRate:        "0.000095",
				High24h:            "35929.3",
				HoldingAmount:      "49626.455",
				IndexPrice:         "34487.575151",
				Last:               "34483.1",
				Low24h:             "30392.3",
				OpenUtc:            "33093.7",
				PriceChangePercent: "0.11682",
				QuoteVolume:        "7728953739.0099",
				Symbol:             "BTCUSDT_UMCBL",
				Timestamp:          "1698145377845",
				UsdtVolume:         "7728953739.0099",
			},
		},
	}

	service := NewMixMarketService(mockClient)
	res, err := service.GetTickers(params["productType"])

	assert.NoError(t, err)
	assert.Equal(t, expect, res)

}

func TestMixMarketClient_Fills(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	client := NewMixMarketService(mockClient)

	resp, err := client.Fills("BTCUSDT_UMCBL", "20")

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestMixMarketClient_Candles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	client := NewMixMarketService(mockClient)

	resp, err := client.Candles("BTCUSDT_UMCBL", "60", "1629177891000", "1629181491000")

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestMixMarketClient_Index(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	client := NewMixMarketService(mockClient)

	resp, err := client.Index("BTCUSDT_UMCBL")

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestGetNextFundingTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uri := constants.MixMarket + "/funding-time"
	params := map[string]string{"symbol": "BTCUSDT_UMCBL"}

	data := `{
		"code": "00000",
		"msg": "success",
		"requestTime": 1698150826421,
		"data": {
			"symbol": "BTCUSDT_UMCBL",
			"fundingTime": "1698163200000",
			"ratePeriod": "8"
		}
	}`

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().DoGet(uri, params).Return(data, nil)

	expect := GetNextFundingTimeResponse{
		CommonResponse: common.CommonResponse{
			Code: "00000",
			Msg:  "success",
			//RequestTime: 1698150826421,
		},
		Data: GetNextFundingTimeData{
			Symbol:      "BTCUSDT_UMCBL",
			FundingTime: "1698163200000",
			RatePeriod:  "8",
		},
	}

	service := NewMixMarketService(mockClient)
	res, err := service.GetNextFundingTime(params["symbol"])

	assert.NoError(t, err)
	assert.Equal(t, expect, res)
}

func TestMixMarketClient_HistoryFundRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	client := NewMixMarketService(mockClient)

	resp, err := client.HistoryFundRate("BTCUSDT_UMCBL", "10", "1", "true")

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestMixMarketClient_CurrentFundRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	client := NewMixMarketService(mockClient)

	resp, err := client.CurrentFundRate("BTCUSDT_UMCBL")

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestMixMarketClient_OpenInterest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	client := NewMixMarketService(mockClient)

	resp, err := client.OpenInterest("BTCUSDT_UMCBL")

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestMixMarketClient_MarkPrice(t *testing.T) {

}
