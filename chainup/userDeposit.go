package chainup

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/guanfantech/chainup-sdk/api"
)

// 同步充值记录
type SyncDepositListResDataItem struct {
	Id            int    `json:"id"`            // 充值唯一id
	Uid           int    `json:"uid"`           // 充值用户id
	Symbol        string `json:"symbol"`        // 币种
	Amount        string `json:"amount"`        // 充值金额
	CreatedAt     uint64 `json:"created_at"`    // 创建时间
	UpdatedAt     uint64 `json:"updated_at"`    // 修改时间
	Txid          string `json:"txid"`          // 区块链交易ID
	Confirmations int    `json:"confirmations"` // 区块链确认数
	AddressFrom   string `json:"address_from"`  // 来源地址
	AddressTo     string `json:"address_to"`    // 充值到帐地址
	Status        int    `json:"status"`        // 0待确认，1 成功，2 失败，4 待KYT验证，5 待人工审核(KYT风险等级过高)，6 待人工审核(KYT充值熔断)
}

func (c *Client) SyncDepositList(maxId int64) (resData *[]SyncDepositListResDataItem, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
		MaxId   int64  `json:"maxId"`   // 必填 返回大于id的100条充值记录数据
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		MaxId:   maxId,
	}
	_, originResponse, err := c.newRequest(api.SyncDepositList, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 批量获取充值记录
type GetDepositListResDataItem struct {
	Id            int    `json:"id"`            // 充值唯一id
	Uid           int    `json:"uid"`           // 充值 用户id
	Symbol        string `json:"symbol"`        // 币种
	Amount        string `json:"amount"`        // 充值金额
	CreatedAt     uint64 `json:"created_at"`    // 创建时间,时间戳
	UpdatedAt     uint64 `json:"updated_at"`    // 修改时间，时间戳
	Txid          string `json:"txid"`          // 区块链交易ID
	Confirmations int    `json:"confirmations"` // 区块链确认数
	AddressFrom   string `json:"address_from"`  // 来源地址
	AddressTo     string `json:"address_to"`    // 充值到帐地址
	Status        int    `json:"status"`        // 0待确认，1 成功，2 失败，4 待KYT验证，5 待人工审核(KYT风险等级过高)，6 待人工审核(KYT充值熔断)
}

func (c *Client) GetDepositList(ids []int) (resData *[]GetDepositListResDataItem, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
		Ids     string `json:"ids"`     // 必填	多个充值id使用逗号隔开，最多100个id
	}
	idsStr := []string{}
	for _, v := range ids {
		id := strconv.FormatInt(int64(v), 10)
		idsStr = append(idsStr, id)
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		Ids:     strings.Join(idsStr, ","),
	}
	_, originResponse, err := c.newRequest(api.GetDepositList, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 获取归集矿工费
// 账户类型的币种充值后需要进行归集，UTXO无归集矿工费费用
type SyncMinerFeeListResDataItem struct {
	Id              int    `json:"id"`               // 必填 归集唯一ID
	Symbol          string `json:"symbol"`           // 必填 币种
	Amount          string `json:"amount"`           // 必填 归集金额
	Fee             string `json:"fee"`              // 必填 归集手续费
	CreatedAt       uint64 `json:"created_at"`       // 必填 创建时间
	UpdatedAt       uint64 `json:"updated_at"`       // 必填 修改时间
	Txid            string `json:"txid"`             // 必填 区块链交易ID
	Confirmations   int    `json:"confirmations"`    // 必填 区块链确认数
	Status          int    `json:"status"`           // 必填 0待确认，1 成功，2 失败，4 待KYT验证，5 待人工审核(KYT风险等级过高)，6 待人工审核(KYT充值熔断)
	AddressTo       string `json:"address_to"`       // 必填 充值到账地址
	AddressFrom     string `json:"address_from"`     // 必填 充值发送地址
	Txid_type       string `json:"txid_type"`        // 必填 0 链上交易，1 联盟转账交易
	Base_symbol     string `json:"base_symbol"`      // 必填 主链币名称
	ContractAddress string `json:"contract_address"` // 必填 币种合约地址
	Email           string `json:"email"`            // 必填 邮箱
}

func (c *Client) SyncMinerFeeList(maxId int64) (resData *[]SyncMinerFeeListResDataItem, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
		MaxId   int64  `json:"maxId"`   // 必填 返回大于 id 的 100 条归集矿 工费记录数据
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		MaxId:   maxId,
	}
	_, originResponse, err := c.newRequest(api.SyncMinerFeeList, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
