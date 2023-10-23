package common

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/arensusu/bitget-golang-sdk-api/config"
	"github.com/arensusu/bitget-golang-sdk-api/constants"
	"github.com/arensusu/bitget-golang-sdk-api/internal"
)

type BitgetRestClient struct {
	ApiKey       string
	ApiSecretKey string
	Passphrase   string
	BaseUrl      string
	HttpClient   *http.Client
	Signer       *Signer
}

type CommonResponse struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}

func (p *BitgetRestClient) Init() *BitgetRestClient {
	creds := config.GetDefaultCredentials()
	return p.InitWithCreds(creds)
}
func (p *BitgetRestClient) InitWithCreds(creds *config.ApiCreds) *BitgetRestClient {
	p.ApiKey = creds.ApiKey
	p.ApiSecretKey = creds.SecretKey
	p.Passphrase = creds.PASSPHRASE

	p.BaseUrl = config.BaseUrl
	p.Signer = new(Signer).Init(creds.SecretKey)
	p.HttpClient = &http.Client{
		Timeout: time.Duration(config.TimeoutSecond) * time.Second,
	}

	return p
}

func (p *BitgetRestClient) DoPost(uri string, params string) (string, error) {
	timesStamp := internal.TimesStamp()
	//body, _ := internal.BuildJsonParams(params)

	sign := p.Signer.Sign(constants.POST, uri, params, timesStamp)
	requestUrl := config.BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(constants.POST, requestUrl, buffer)

	internal.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)
	if err != nil {
		return "", err
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)

	if response.StatusCode != 200 {
		return "", fmt.Errorf("StatusCode: %d, error: %v", response.StatusCode, responseBodyString)
	}

	return responseBodyString, err
}

func (p *BitgetRestClient) DoGet(uri string, params map[string]string) (string, error) {
	timesStamp := internal.TimesStamp()
	body := internal.BuildGetParams(params)

	sign := p.Signer.Sign(constants.GET, uri, body, timesStamp)

	requestUrl := p.BaseUrl + uri + body

	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	internal.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)

	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)

	if response.StatusCode != 200 {
		return "", fmt.Errorf("StatusCode: %d, error: %v", response.StatusCode, responseBodyString)
	}

	return responseBodyString, err
}
