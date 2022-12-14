package chainup_sdk

import (
	"encoding/json"
	"time"

	"github.com/guanfantech/chainup-sdk/api"
)

// Custody商户转账
// Custody内部商户互相转账
type TransferResData struct {
	Receipt string `json:"receipt"` // 转账唯一凭证
}

func (c *Client) Transfer(requestId, amount, symbol, to, remark string) (resData *TransferResData, err error) {
	type reqData struct {
		Time      int64  `json:"time"`       // 必填	当前时间戳
		Charset   string `json:"charset"`    // 必填	编码格式，无特殊情况，传参数utf-8
		Version   string `json:"version"`    // 必填	接口版本号，无特殊情况，传参数v2
		RequestId string `json:"request_id"` // 必填	请求唯一标识，最多支持64位
		Symbol    string `json:"symbol"`     // 必填	转账币种，Custody获取币种名称
		Amount    string `json:"amount"`     // 必填	转账数量，包含转账手续费
		To        string `json:"to"`         // 必填	转入商户，目前填写商户APPID
		Remark    string `json:"remark"`     // 选填	备注字段
	}
	req := &reqData{
		Time:      time.Now().Unix(),
		Charset:   "utf-8",
		Version:   "v2",
		RequestId: requestId,
		Symbol:    symbol,
		Amount:    amount,
		To:        to,
		Remark:    remark,
	}
	_, originResponse, err := c.newRequest(api.Transfer, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 转账信息异步确认

type TransferCheckResData struct {
	Receipt string `json:"receipt"` // 转账唯一凭证
}

func (c *Client) TransferCheck(requestId, amount, symbol, to, checkSum, remark string) (resData *TransferCheckResData, err error) {
	type reqData struct {
		Time      int64  `json:"time"`       // 必填	当前时间戳
		Charset   string `json:"charset"`    // 必填	编码格式，无特殊情况，传参数utf-8
		Version   string `json:"version"`    // 必填	接口版本号，无特殊情况，传参数v2
		RequestId string `json:"request_id"` // 必填	请求唯一标识，最多支持64位
		Symbol    string `json:"symbol"`     // 必填	转账币种，Custody获取币种名称
		Amount    string `json:"amount"`     // 必填	转账数量，包含转账手续费
		To        string `json:"to"`         // 必填	转入商户，目前填写商户APPID
		CheckSum  string `json:"check_sum"`  // 必填 随机校验码，第三方原样返回此字段平台认为成功
		Remark    string `json:"remark"`     // 选填	备注字段
	}
	req := &reqData{
		Time:      time.Now().Unix(),
		Charset:   "utf-8",
		Version:   "v2",
		RequestId: requestId,
		Symbol:    symbol,
		Amount:    amount,
		To:        to,
		CheckSum:  checkSum,
		Remark:    remark,
	}
	_, originResponse, err := c.newRequest(api.TransferCheck, req)
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
type GetTransferListResData []GetTransferListResDataItem
type GetTransferListResDataItem struct {
	Id        string `json:"id"`         // 请求唯一标识，最多支持64位
	Symbol    string `json:"symbol"`     // 币种
	Amount    string `json:"amount"`     // 转账数量，包含转账手续费
	From      string `json:"from"`       // 转出商户，转出商户APPID
	To        string `json:"to"`         // 转入商户，转入商户APPID
	CreatedAt string `json:"created_at"` // 创建时间
	RequestId string `json:"request_id"` // 三方ID
	Receipt   string `json:"receipt"`    // 转账凭证
	Rremark   string `json:"remark"`     // 最大支持32字符
}

func (c *Client) GetTransferList(ids, idsType string) (resData *GetTransferListResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`     // 必填 当前时间戳
		Charset string `json:"charset"`  // 必填 编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"`  // 必填 接口版本号，无特殊情况，传参数v2
		Ids     string `json:"ids"`      // 必填 请求唯一标识,多个之间用英文逗号分割，最多100个
		IdsType string `json:"ids_type"` // 必填 request_id：请求ID（默认）；receipt：转账凭证
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		Ids:     ids,
		IdsType: idsType,
	}
	_, originResponse, err := c.newRequest(api.GetTransferList, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 同步转账记录
// 同步所有转账记录（分页）
type SyncTransferListResData []SyncTransferListResDataItem
type SyncTransferListResDataItem struct {
	Id        string `json:"id"`         // 请求唯一标识，最多支持64位
	Symbol    string `json:"symbol"`     // 币种
	Amount    string `json:"amount"`     // 转账数量，包含转账手续费
	From      string `json:"from"`       // 转出商户，转出商户APPID
	To        string `json:"to"`         // 转入商户，转入商户APPID
	CreatedAt string `json:"created_at"` // 创建时间
	RequestId string `json:"request_id"` // 三方ID
	Receipt   string `json:"receipt"`    // 转账凭证
	Rremark   string `json:"remark"`     // 最大支持32字符
}

func (c *Client) SyncTransferList(maxId string) (resData *SyncTransferListResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填 当前时间戳
		Charset string `json:"charset"` // 必填 编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填 接口版本号，无特殊情况，传参数v2
		MaxId   string `json:"max_id"`  // 必填	返回大于id的100条转账记录数据
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		MaxId:   maxId,
	}
	_, originResponse, err := c.newRequest(api.SyncTransferList, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
