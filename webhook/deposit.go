package webhook

import (
	"encoding/json"

	"github.com/farmerx/gorsa"
)

// 注：Custody系统的异步回调是当每笔订单最终态时才会触发，每日最多发送5次；
// 定时任务：总计回调五次
// 通知时间：第一次1s, 第二次2min，第三次8min，第四次32min，第五次128min
// 回调逻辑：
// 如果回调成功，更新回调状态；
// 如果回调失败，继续回调，更新下次回调间隔时间；
// 当回调失败达到5次，停止回调
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
	CreatedAt     uint64 `json:"created_at"`    // 必填 创建时间
	UpdatedAt     uint64 `json:"updated_at"`    // 必填 修改时间
	Txid          string `json:"txid"`          // 必填 区块链交易ID
	Confirmations string `json:"confirmations"` // 必填 区块链确认数
	Status        string `json:"status"`        // 必填 充值状态 0待确认，1 成功，2 失败，4 待KYT验证，5 待人工审核(KYT风险等级过高)，6 待人工审核(KYT充值熔断)
}

// SUCCESS表示成功，FAILURE表示失败 （注意此处返回参数无需进行加密）

func GetDepositRequestData(custodyPubKey string, dataStr string) (*Deposit, error) {
	err := gorsa.RSA.SetPublicKey(custodyPubKey)
	if err != nil {
		return nil, err
	}
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
