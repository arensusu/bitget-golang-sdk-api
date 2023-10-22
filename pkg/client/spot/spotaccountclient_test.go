package spot

import (
	"fmt"
	"testing"

	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/internal/common/mocks"
	"github.com/arensusu/bitget-golang-sdk-api/pkg/client/spot/domain"
	"github.com/arensusu/bitget-golang-sdk-api/pkg/model/spot/account"
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
	uri := constants.SpotAccount + "/assets"
	params := internal.NewParams()
	mockClient.EXPECT().DoGet(uri, params).Return(data, nil)

	service := NewSpotGetAccountAssetsService(mockClient)
	response, err := service.Do()
	expect := domain.SpotGetAccountAssetsResponse{
		CommonResponse: common.CommonResponse{
			Code: "00000",
			Msg:  "success",
		},
		Data: []domain.SpotGetAccountAssetsData{
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
	client := new(SpotAccountClient).Init()

	req := account.BillsReq{CoinId: "1", Before: "777031099461570560"}

	resp, err := client.Bills(req)

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestSpotAccountClient_TransferRecords(t *testing.T) {
	client := new(SpotAccountClient).Init()

	resp, err := client.TransferRecords("2", "USDT_MIX", "10", "", "")

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}
