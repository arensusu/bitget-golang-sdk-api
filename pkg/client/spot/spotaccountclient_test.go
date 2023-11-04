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

func TestSpotAccountGetAssetsLiteService(t *testing.T) {
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

	param := GetAssetsLiteParam{}
	service := NewSpotAccountService(mockClient)
	response, err := service.GetAssetsLite(param)
	expect := GetAssetsLiteResponse{
		CommonResponse: common.CommonResponse{
			Code: "00000",
			Msg:  "success",
			//RequestTime: 1698067287632,
		},
		Data: []GetAssetsLiteData{
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
