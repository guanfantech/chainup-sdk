package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/guanfantech/chainup-sdk/api"
	"github.com/pkg/errors"

	"github.com/farmerx/gorsa"
)

type EnvType string

const (
	EnvDev  EnvType = "Dev"
	EnvProd EnvType = "Prod"
)

var chainupUrl = map[EnvType]string{
	EnvDev:  "https://openapi.hicoin.vip/api/v2",
	EnvProd: "https://openapi.hicoin.vip/api/v2",
}

// Client type
// 關於公私鑰詳情請看文檔
// https://custodydocs.chainup.com/zh/latest/API-WaaS-V2/api_description.html#api-api
type Client struct {
	env           EnvType
	url           string
	appId         string       // 請創建錢包後獲取
	custodyPubKey string       // Custody系統公鑰；請創建錢包後從Custody系統獲取
	pubKey        string       // 客戶公鑰；自主生成；請創建錢包後配置到Custody系統
	privateKey    string       // 客戶私鑰；自主生成、保存
	httpClient    *http.Client // default http.DefaultClient
}

// ClientOption type
type ClientOption func(*Client) error

var chainup *Client

func New(env EnvType, appId, custodyPubKey, pubKey, privateKey string, cOptions ...ClientOption) (*Client, error) {
	// check environment
	if env != EnvDev && env != EnvProd {
		return nil, errors.New("invalid environment")
	}
	// check params
	if appId == "" || custodyPubKey == "" || pubKey == "" || privateKey == "" {
		return nil, errors.New("invalid params")
	}

	// 設定chainup的公鑰進行解密
	if err := gorsa.RSA.SetPublicKey(custodyPubKey); err != nil {
		log.Fatalln(`set public key :`, err)
	}

	// 設定自己的私鑰進行加密
	if err := gorsa.RSA.SetPrivateKey(privateKey); err != nil {
		log.Fatalln(`set private key :`, err)
	}

	chainup = &Client{
		env:           env,
		url:           chainupUrl[env],
		appId:         appId,
		custodyPubKey: custodyPubKey,
		pubKey:        pubKey,
		privateKey:    privateKey,
		httpClient:    http.DefaultClient,
	}
	return chainup, nil
}

type OriginResponse struct {
	Code string `json:"code"` // 状态码
	Msg  string `json:"msg"`  // 响应结果说明
	Data string `json:"data"` // 具体响应数据，数据结构定义如下
}

// NewRequest method
func (c *Client) newRequest(route api.RouteDetail, requestData interface{}) (*http.Response, *OriginResponse, error) {
	dataByte, err := json.Marshal(requestData)
	if err != nil {

		return nil, nil, errors.WithStack(err)
	}

	req, err := http.NewRequest(route.Method, c.url+"/api/v2"+route.Path, nil)
	if err != nil {

		return nil, nil, errors.WithStack(err)
	}
	prienctypt, err := gorsa.RSA.PriKeyENCTYPT(dataByte)
	if err != nil {

		return nil, nil, errors.WithStack(err)
	}
	base64Data := base64.RawURLEncoding.EncodeToString(prienctypt)

	q := req.URL.Query()
	q.Add("app_id", c.appId)
	q.Add("data", string(base64Data))
	req.URL.RawQuery = q.Encode()

	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("X-LINE-ChannelId", c.channelID)
	// req.Header.Set("X-LINE-Authorization-Nonce", nounce)

	// hash := hmac.New(sha256.New, []byte(c.channelSecret))
	// hash.Write([]byte(message))
	// req.Header.Set("X-LINE-Authorization", base64.StdEncoding.EncodeToString(hash.Sum(nil)))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("response.StatusCode:", resp.StatusCode)
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	// if resp.StatusCode != 200 {
	fmt.Println("response Body:", string(body))
	// }

	response := &struct {
		Data string `json:"data"`
	}{}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	dataEnctypt, err := base64.RawURLEncoding.DecodeString(response.Data)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	resDataByte, err := gorsa.RSA.PubKeyDECRYPT(dataEnctypt)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	// Debug
	fmt.Println("api response: ", string(resDataByte))

	originResData := &OriginResponse{}

	err = json.Unmarshal(resDataByte, originResData)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return resp, originResData, err
}
