package webhook

import (
	"encoding/json"

	"github.com/farmerx/gorsa"
)

type WebhookWithdrawConfirmReq struct {
	Time      uint64 `json:"time"`       // 必填，当前时间戳
	Charset   string `json:"charset"`    // 必填，编码格式，无特殊情况，传参数utf-8
	Version   string `json:"version"`    // 必填，接口版本号，无特殊情况，传参数v2
	RequestId string `json:"request_id"` // 必填，请求唯一标识
	FromUid   string `json:"from_uid"`   // 必填，转出用户ID
	ToAddress string `json:"to_address"` // 必填，转入用户地址
	Amount    string `json:"amount"`     // 必填，提现金额，包含提现手续费；手续费需要在商户后台配置
	Symbol    string `json:"symbol"`     // 必填，提现币种
	CheckSum  string `json:"check_sum"`  // 必填，随机校验码，第三方原样返回此字段平台认为成功
}

type WebhookWithdrawConfirmResponse struct {
	CheckSum string `json:"check_sum"` // 必填 请求参数中的check_sum
	Time     uint64 `json:"time"`      // 必填，当前时间戳
}

func GetWebhookWithdrawConfirmRequestData(custodyPubKey string, dataStr string) (*WebhookWithdrawConfirmReq, error) {
	err := gorsa.RSA.SetPublicKey(custodyPubKey)
	if err != nil {
		return nil, err
	}

	jsonB, err := gorsa.RSA.PubKeyDECRYPT([]byte(dataStr))
	if err != nil {
		return nil, err
	}
	var data *WebhookWithdrawConfirmReq
	err = json.Unmarshal(jsonB, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
