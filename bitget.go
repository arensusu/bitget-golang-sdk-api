package bitget

import (
	"github.com/arensusu/bitget-golang-sdk-api/internal/common"
	"github.com/arensusu/bitget-golang-sdk-api/pkg/client/mix"
	"github.com/arensusu/bitget-golang-sdk-api/pkg/client/spot"
)

// client
type Client struct {
	client *common.BitgetRestClient
}

func NewClient() *Client {
	bc := new(common.BitgetRestClient).Init()

	return &Client{
		client: bc,
	}
}

// mix
func (c *Client) MixAccount() *mix.MixAccountService {
	return mix.NewMixAccountService(c.client)
}

func (c *Client) MixMarket() *mix.MixMarketService {
	return mix.NewMixMarketService(c.client)
}

// spot
func (c *Client) SpotAccount() *spot.SpotAccountService {
	return spot.NewSpotAccountService(c.client)
}
