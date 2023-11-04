package bitget

import (
	"net/http"

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

func (c *Client) SetHttpClient(client *http.Client) *Client {
	c.client.HttpClient = client
	return c
}

// mix
func (c *Client) NewMixAccountGetAccountListService() *mix.MixAccountGetAccountListService {
	return mix.NewMixAccountGetAccountListService(c.client)
}

func (c *Client) MixMarket() *mix.MixMarketService {
	return mix.NewMixMarketService(c.client)
}

// spot
func (c *Client) NewSpotAccountGetAccountAssetsLiteService() *spot.SpotAccountGetAccountAssetsLiteService {
	return spot.NewSpotAccountGetAccountAssetsLiteService(c.client)
}
