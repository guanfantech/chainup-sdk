package chainup_sdk

import (
	"encoding/json"
	"time"

	"github.com/guanfantech/chainup-sdk/api"
)

// 查詢用戶信息
type GetUserInfoResData struct {
	Uid      int64  `json:"uid"`      // 用户在钱包服务的唯一标识
	Nickname string `json:"nickname"` // 用户昵称
}

func (c *Client) GetUserInfo(country, mobile, email string) (resData *GetUserInfoResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`              // 必填	当前时间戳
		Charset string `json:"charset"`           // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"`           // 必填	接口版本号，无特殊情况，传参数v2
		Country string `json:"country,omitempty"` // 可选	国家编号，mobile不为空时，该字段必填。如：86
		Mobile  string `json:"mobile,omitempty"`  // 可选	手机号，手机和邮箱需要保证其中之一不能为空
		Email   string `json:"email,omitempty"`   // 可选	邮箱，手机和邮箱需要保证其中之一不能为空
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		Country: country,
		Mobile:  mobile,
		Email:   email,
	}
	_, originResponse, err := c.newRequest(api.GetUserInfo, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 获取支持的币种列表
type GetCoinListResData []GetCoinListResDataItem
type GetCoinListResDataItem struct {
	Symbol              string `json:"symbol"`               // 币种（调用提币的接口，及任何查询接口时务必使用此字段返回的值）
	Icon                string `json:"icon"`                 // 币种icon
	RealSymbol          string `json:"real_symbol"`          // 币种链上名称
	Decimals            string `json:"decimals"`             // 精度
	Name                string `json:"name"`                 // 币种全称
	BaseSymbol          string `json:"base_symbol"`          // 主链币币名
	ContractAddress     string `json:"contract_address"`     // 合约地址
	DepositConfirmation string `json:"deposit_confirmation"` // 币种充值确认数
	Explorer            string `json:"explorer"`             // 区块浏览器
	AddressRegex        string `json:"address_regex"`        // 地址正则
	AddressTagRegex     string `json:"address_tag_regex"`    // tag正则
	SupportMemo         string `json:"support_memo"`         // 是否支持memo，0不支持1支持
	SupportToken        string `json:"support_token"`        // 是否支持token币，0不支持1支持,主链币才有值,代币为空
}

func (c *Client) GetCoinList() (resData *GetCoinListResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
	}
	_, originResponse, err := c.newRequest(api.GetCoinList, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 获取指定用户的账户信息
type GetByUidAndSymbolResData struct {
	NormalBalance  string `json:"normal_balance"`  // 正常账户余额
	LockBalance    string `json:"lock_balance"`    // 冻结账户余额
	DepositAddress string `json:"deposit_address"` // 币种对应的充值地址
}

func (c *Client) GetByUidAndSymbol(uid, symbol string) (resData *GetByUidAndSymbolResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
		Uid     string `json:"uid"`     // 必填	用户ID
		Symbol  string `json:"symbol"`  // 必填	币种
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		Uid:     uid,
		Symbol:  symbol,
	}
	_, originResponse, err := c.newRequest(api.GetByUidAndSymbol, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 获取商户归集账户余额
// 開啟商戶資金自動歸集之後，商戶可以通過該接口種獲取商戶歸集賬戶餘額

type GetCompanyBySymbolResData struct {
	Symbol  string `json:"symbol"`  // 币种名称
	Balance string `json:"balance"` // 归集账户余额
}

func (c *Client) GetCompanyBySymbol(symbol string) (resData *GetCompanyBySymbolResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
		Symbol  string `json:"symbol"`  // 必填	币种
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		Symbol:  symbol,
	}
	_, originResponse, err := c.newRequest(api.GetCompanyBySymbol, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
