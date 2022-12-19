package webhook

import (
	"encoding/json"

	"github.com/farmerx/gorsa"
)

type Deposit struct {
	Charset       string `json:"charset"`       // 必填 编码格式，无特殊情况，传参数utf-8
	Version       string `json:"version"`       // 必填 接口版本号，无特殊情况，传参数v2
	Side          string `json:"side"`          // 必填 通知类型， 充值通知：deposit， 提现通知： withdraw
	NotifyTime    string `json:"notify_time"`   // 必填 通知时间
	Id            string `json:"id"`            // 必填 充值id
	Uid           string `json:"uid"`           // 必填 提现用户id
	Symbol        string `json:"symbol"`        // 必填 币种
	Amount        string `json:"amount"`        // 必填 提现金额
	AddressTo     string `json:"address_to"`    // 必填 充值地址
	CreatedAt     string `json:"created_at"`    // 必填 创建时间
	UpdatedAt     string `json:"updated_at"`    // 必填 修改时间
	Txid          string `json:"txid"`          // 必填 区块链交易ID
	Confirmations string `json:"confirmations"` // 必填 区块链确认数
	Status        string `json:"status"`        // 必填 充值状态 0待确认，1 成功，2 失败，4 待KYT验证，5 待人工审核(KYT风险等级过高)，6 待人工审核(KYT充值熔断)
}

// SUCCESS表示成功，FAILURE表示失败 （注意此处返回参数无需进行加密）

func GetDepositRequestData(dataStr string) (*Deposit, error) {
	jsonB, err := gorsa.RSA.PubKeyDECRYPT([]byte(dataStr))
	if err != nil {
		return nil, err
	}
	var data *Deposit
	err = json.Unmarshal(jsonB, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
