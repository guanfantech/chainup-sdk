package webhook

type Withdraw struct {
	Charset           string `json:"charset"`             // 必填 编码格式，无特殊情况，传参数utf-8
	Version           string `json:"version"`             // 必填 接口版本号，无特殊情况，传参数v2
	Side              string `json:"side"`                // 必填 通知类型， 充值通知：deposit， 提现通知： withdraw
	NotifyTime        string `json:"notify_time"`         // 必填 通知时间
	RequestId         string `json:"request_id"`          // 必填 提现请求ID，对应提现接口中的request_id
	Id                string `json:"id"`                  // 必填 提现id
	Uid               string `json:"uid"`                 // 必填 提现用户id
	Symbol            string `json:"symbol"`              // 必填 币种
	Amount            string `json:"amount"`              // 必填 提现金额
	WithdrawFeeSymbol string `json:"withdraw_fee_symbol"` // 必填 提现手续费币种
	WithdrawFee       string `json:"withdraw_fee"`        // 必填 提现手续费
	FeeSymbol         string `json:"fee_symbol"`          // 必填 挖矿手续费币种
	RealFee           string `json:"real_fee"`            // 必填 矿工费
	AddressTo         string `json:"address_to"`          // 必填 充值地址
	CreatedAt         string `json:"created_at"`          // 必填 创建时间
	UpdatedAt         string `json:"updated_at"`          // 必填 修改时间
	Txid              string `json:"txid"`                // 必填 区块链交易ID
	Confirmations     string `json:"confirmations"`       // 必填 区块链确认数
	Status            string `json:"status"`              // 必填 提现状态: 0 风控审核中，1 支付中，2 审核拒绝，4 失败，5 成功， 6 已撤销，7 待KYT验证， 8 待人工审核(KYT风险等级过高)
}

// SUCCESS表示成功，FAILURE表示失败 （注意此处返回参数无需进行加密）
