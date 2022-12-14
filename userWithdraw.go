package chainup_sdk

import (
	"encoding/json"
	"time"

	"github.com/guanfantech/chainup-sdk/api"
)

// 发起提现
// 如果提现地址为联盟内部地址，则不会上链；否则需要上链。另外，为了避免非法的提现请求，钱包服务在接收到第三方应用的提现请求时，会调用第三方提供的[提现二次确认接口]。
type WithdrawResData struct {
	Status string `json:"status"` // 必填	0 风控审核中，1 支付中，2 审核拒绝，4 失败，5 成功， 6 已撤销，7 待KYT验证， 8 待人工审核(KYT风险等级过高)
}

func (c *Client) Withdraw(requestId, fromUid, toAddress, amount, symbol string) (resData *WithdrawResData, err error) {
	type reqData struct {
		Time      int64  `json:"time"`       // 必填	当前时间戳
		Charset   string `json:"charset"`    // 必填	编码格式，无特殊情况，传参数utf-8
		Version   string `json:"version"`    // 必填	接口版本号，无特殊情况，传参数v2
		RequestId string `json:"request_id"` // 必填 请求唯一标识，最多支持64位
		FromUid   string `json:"from_uid"`   // 必填 转出用户ID
		ToAddress string `json:"to_address"` // 必填 转入用户地址，memo类型，使用"_"进行拼接，如: eos_24545
		Amount    string `json:"amount"`     // 必填 提现金额，包含提现手续费；手续费需要在商户后台配置
		Symbol    string `json:"symbol"`     // 必填 提现币种
	}
	req := &reqData{
		Time:      time.Now().Unix(),
		Charset:   "utf-8",
		Version:   "v2",
		RequestId: requestId,
		FromUid:   fromUid,
		ToAddress: toAddress,
		Amount:    amount,
		Symbol:    symbol,
	}
	_, originResponse, err := c.newRequest(api.Withdraw, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 同步提现记录
type SyncWithdrawListResData []SyncWithdrawListResDataItem
type SyncWithdrawListResDataItem struct {
	RequestId         string `json:"request_id"`          // 必填	请求id,
	Id                int    `json:"id"`                  // 必填	提现id
	Uid               int    `json:"uid"`                 // 必填	提现用户id
	Symbol            string `json:"symbol"`              // 必填	币种
	Amount            string `json:"amount"`              // 必填	提现金额
	WithdrawFeeSymbol string `json:"withdraw_fee_symbol"` // 必填	提现手续费币种
	WithdrawFee       string `json:"withdraw_fee"`        // 必填	提现手续费
	FeeSymbol         string `json:"fee_symbol"`          // 必填	挖矿手续费币种
	RealFee           string `json:"real_fee"`            // 必填	矿工费
	CreatedAt         string `json:"created_at"`          // 必填	创建时间
	UpdatedAt         string `json:"updated_at"`          // 必填	修改时间
	AddressFrom       string `json:"address_from"`        // 必填	来源地址
	AddressTo         string `json:"address_to"`          // 必填	到账地址
	Txid              string `json:"txid"`                // 必填	区块链交易ID
	Confirmations     int    `json:"confirmations"`       // 必填	区块链确认数
	SaasStatus        int    `json:"saas_status"`         // 必填	平台审核状态:0 未审核，1 已审核，2 审核拒绝
	CompanyStatus     int    `json:"company_status"`      // 必填	商户审核状态:0 未审核，1 已审核，2 审核拒绝
	Status            int    `json:"status"`              // 必填	提现状态: 0 风控审核中，1 支付中，2 审核拒绝，4 失败，5 成功， 6 已撤销，7 待KYT验证， 8 待人工审核(KYT风险等级过高)
}

func (c *Client) SyncWithdrawList(maxId string) (resData *SyncWithdrawListResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
		MaxId   string `json:"max_Id"`  // 必填	返回大于id的100条充值记录数据
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		MaxId:   maxId,
	}
	_, originResponse, err := c.newRequest(api.SyncWithdrawList, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 批量获取提现记录
type GetWithdrawListResData []GetWithdrawListResDataItem
type GetWithdrawListResDataItem struct {
	Request_id        string `json:"request_id"`          // 必填 请求id
	Id                int    `json:"id"`                  // 必填 提现id
	Uid               int    `json:"uid"`                 // 必填 提现用户id
	Symbol            string `json:"symbol"`              // 必填 币种
	Amount            string `json:"amount"`              // 必填 提现金额
	WithdrawFeeSymbol string `json:"withdraw_fee_symbol"` // 必填 提现手续费币种
	WithdrawFee       string `json:"withdraw_fee"`        // 必填 提现手续费
	FeeSymbol         string `json:"fee_symbol"`          // 必填 挖矿手续费币种
	RealFee           string `json:"real_fee"`            // 必填 矿工费
	CreatedAt         string `json:"created_at"`          // 必填 创建时间,
	UpdatedAt         string `json:"updated_at"`          // 必填 修改时间
	AddressFrom       string `json:"address_from"`        // 必填 来源地址
	AddressTo         string `json:"address_to"`          // 必填 到账地址
	Txid              string `json:"txid"`                // 必填 区块链交易ID
	Confirmations     int    `json:"confirmations"`       // 必填 区块链确认数
	SaasStatus        int    `json:"saas_status"`         // 必填 平台审核状态
	CompanyStatus     int    `json:"company_status"`      // 必填 商户审核状态
	Status            int    `json:"status"`              // 必填 提现状态: 0 风控审核中，1 支付中，2 审核拒绝，4 失败，5 成功， 6 已撤销，7 待KYT验证， 8 待人工审核(KYT风险等级过高)
}

func (c *Client) GetWithdrawList(ids string) (resData *GetWithdrawListResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
		Ids     string `json:"ids"`     // 必填	多个request_id使用逗号隔开，最多100个request_id
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		Ids:     ids,
	}
	_, originResponse, err := c.newRequest(api.GetWithdrawList, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
