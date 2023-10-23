package spot

import (
	"testing"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSpotAccountAssetsService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)

	data := `{
		"code":"00000",
		"msg":"success",
		"requestTime":1698067287632,
		"data":[
		  {
			  "coinId":10012,
			  "coinName":"usdt",
			  "available":"0",
			  "frozen":"0",
			  "lock":"0",
			  "uTime":"1622697148"
		  }
	  ]
	  }`
	uri := constants.SpotAccount + "/assets-lite"
	params := internal.NewParams()
	mockClient.EXPECT().DoGet(uri, params).Return(data, nil)

	service := NewSpotAccountGetAccountAssetsLiteService(mockClient)
	response, err := service.Do()
	expect := SpotAccountGetAccountAssetsLiteResponse{
		CommonResponse: common.CommonResponse{
			Code:        "00000",
			Msg:         "success",
			RequestTime: 1698067287632,
		},
		Data: []SpotAccountGetAccountAssetsLiteData{
			{
				CoinId:    10012,
				CoinName:  "usdt",
				Available: "0",
				Frozen:    "0",
				Lock:      "0",
				UTime:     "1622697148",
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expect, response)
}

func TestSpotAccountClient_Bills(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uri := constants.SpotAccount + "/bills"
	data := `{
		"code":"00000",
		"msg":"success",
		"requestTime": 1698067287632,
		"data":[{
			"cTime":"1622697148",
			"coinId":22,
			"coinName":"usdt",
			"groupType":"deposit",
			"bizType":"transfer-in", 
			"quantity":"1",
			"balance": "1",
			"fees":"0",
			"billId":"1291"
	  }]
	}`
	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().DoPost(uri, `{"coinId":"","groupType":"","bizType":"","after":"","before":"","limit":""}`).Return(data, nil)

	expect := SpotAccountGetBillsServiceResponse{
		CommonResponse: common.CommonResponse{
			Code:        "00000",
			Msg:         "success",
			RequestTime: 1698067287632,
		},
		Data: []SpotAccountGetBillsServiceData{
			{
				CTime:     "1622697148",
				CoinId:    22,
				CoinName:  "usdt",
				GroupType: "deposit",
				BizType:   "transfer-in",
				Quantity:  "1",
				Balance:   "1",
				Fees:      "0",
				BillId:    "1291",
			},
		},
	}
	service := NewSpotAccountGetBillsService(mockClient)
	res, err := service.Do()

	assert.NoError(t, err)
	assert.Equal(t, expect, res)
}
