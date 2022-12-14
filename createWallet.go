package main

import (
	"encoding/json"
	"time"

	"github.com/guanfantech/chainup-sdk/api"
)

// 用户手机注册
type CreateUserResData struct {
	Uid int64 `json:"uid"` // 用户在钱包服务的唯一标识
}

func (c *Client) CreateUser(country, mobile string) (resData *CreateUserResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
		Country string `json:"country"` // 必填	国家编号，如：86表示中国
		Mobile  string `json:"mobile"`  // 必填	手机号
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		Country: country,
		Mobile:  mobile,
	}
	_, originResponse, err := c.newRequest(api.CreateUser, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 用户邮箱注册
type RegisterEmailResData struct {
	Uid int64 `json:"uid"` // 用户在钱包服务的唯一标识
}

func (c *Client) RegisterEmail(email string) (resData *RegisterEmailResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
		Email   string `json:"email"`   // 邮箱或虚拟账号，确保其唯一性,最多支持100字符
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		Email:   email,
	}
	_, originResponse, err := c.newRequest(api.RegisterEmail, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}

// 获取用户指定币种充币地址
// 获取用户指定币账户地址。如果没有地址，则给用户分配一个地址，此时不应生成账户，账户还是按需生成。

type GetDepositAddressResData struct {
	Uid     int64  `json:"uid"`     // 用户在钱包服务的唯一标识
	Address string `json:"address"` // 币种账户地址
}

func (c *Client) GetDepositAddress(uid string, symbol string) (resData *GetDepositAddressResData, err error) {
	type reqData struct {
		Time    int64  `json:"time"`    // 必填	当前时间戳
		Charset string `json:"charset"` // 必填	编码格式，无特殊情况，传参数utf-8
		Version string `json:"version"` // 必填	接口版本号，无特殊情况，传参数v2
		Uid     string `json:"uid"`     // 用户ID
		Symbol  string `json:"symbol"`  // 币种
	}
	req := &reqData{
		Time:    time.Now().Unix(),
		Charset: "utf-8",
		Version: "v2",
		Uid:     uid,
		Symbol:  symbol,
	}
	_, originResponse, err := c.newRequest(api.GetDepositAddress, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(originResponse.Data), resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
