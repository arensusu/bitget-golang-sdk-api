package mix

import (
	"fmt"
	"testing"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common/mocks"
	"github.com/arensusu/bitget-golang-sdk-api/pkg/model/mix/account"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMixAccountGetAccountList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uri := constants.MixAccount + "/accounts"
	params := map[string]string{"productType": "umcbl"}
	data := `{
		"code":"00000",
		"msg":"success",
		"requestTime":1698067287632,
		"data":[
			{
				"marginCoin":"USDT",
				"locked":"0",
				"available":"3868.76261019",
				"crossMaxAvailable":"265.53341019",
				"fixedMaxAvailable":"265.53341019",
				"maxTransferOut":"265.53341019",
				"equity":"3859.83161019",
				"usdtEquity":"3859.831610196897",
				"btcEquity":"0.126083574363",
				"crossRiskRate":"0.095866548951",
				"unrealizedPL":"-8.93100000056",
				"bonus":"0"
			}
		]
	}`
	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().DoGet(uri, params).Return(data, nil)

	expect := MixAccountGetAccountListReponse{
		CommonResponse: common.CommonResponse{
			Code:        "00000",
			Msg:         "success",
			RequestTime: 1698067287632,
		},
		Data: []MixAccountGetAccountListData{
			{
				MarginCoin:        "USDT",
				Locked:            "0",
				Available:         "3868.76261019",
				CrossMaxAvailable: "265.53341019",
				FixedMaxAvailable: "265.53341019",
				MaxTransferOut:    "265.53341019",
				Equity:            "3859.83161019",
				UsdtEquity:        "3859.831610196897",
				BtcEquity:         "0.126083574363",
				CrossRiskRate:     "0.095866548951",
				UnrealizedPL:      "-8.93100000056",
				Bonus:             "0",
			},
		},
	}

	service := NewMixAccountGetAccountListService(mockClient)
	res, err := service.ProductType("umcbl").Do()

	assert.NoError(t, err)
	assert.Equal(t, expect, res)
}

func TestMixAccountClient_GetAccount(t *testing.T) {
	client := new(MixAccountClient).Init()
	resp, err := client.Account("BTCUSDT_UMCBL", "USDT")
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestMixAccountClient_SetLeverage(t *testing.T) {
	client := new(MixAccountClient).Init()
	req := account.SetLeveragerReq{Symbol: "BTCUSDT_UMCBL", MarginCoin: "USDT", Leverage: "10"}

	resp, err := client.SetLeverage(req)

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestMixAccountClient_SetMargin(t *testing.T) {
	client := new(MixAccountClient).Init()
	req := account.SetMarginReq{Symbol: "BTCUSDT_UMCBL", MarginCoin: "USDT", HoldSide: "long", Amount: "10"}

	resp, err := client.SetMargin(req)

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}
func TestMixAccountClient_SetMarginMode(t *testing.T) {
	client := new(MixAccountClient).Init()
	req := account.SetMarginModeReq{Symbol: "BTCUSDT_UMCBL", MarginCoin: "USDT", MarginMode: "fixed"}

	resp, err := client.SetMarginMode(req)

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestMixAccountClient_OpenCount(t *testing.T) {
	client := new(MixAccountClient).Init()
	req := account.OpenCountReq{Symbol: "BTCUSDT_UMCBL", MarginCoin: "USDT", OpenPrice: "3000", OpenAmount: "99999", Leverage: "20"}

	resp, err := client.OpenCount(req)

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}
