package api

import "net/http"

type RouteDetail struct {
	Path   string
	Method string
}

var (
	//User-related operations
	CreateUser    = RouteDetail{Path: "/user/createUser", Method: http.MethodPost}
	RegisterEmail = RouteDetail{Path: "/user/registerEmail", Method: http.MethodPost}
	GetUserInfo   = RouteDetail{Path: "/user/info", Method: http.MethodGet}
	GetCoinList   = RouteDetail{Path: "/user/getCoinList", Method: http.MethodGet}

	//Account related operations api
	GetDepositAddress  = RouteDetail{Path: "/account/getDepositAddress", Method: http.MethodPost}
	GetByUidAndSymbol  = RouteDetail{Path: "/account/getByUidAndSymbol", Method: http.MethodGet}
	GetCompanyBySymbol = RouteDetail{Path: "/account/getCompanyBySymbol", Method: http.MethodGet}

	Transfer         = RouteDetail{Path: "/account/transfer", Method: http.MethodPost}
	TransferCheck    = RouteDetail{Path: "/same-with-withdraw-check-sum", Method: http.MethodPost}
	GetTransferList  = RouteDetail{Path: "/account/transferList", Method: http.MethodGet}
	SyncTransferList = RouteDetail{Path: "/account/syncTransferList", Method: http.MethodGet}

	//Deposit-related operations api
	SyncDepositList  = RouteDetail{Path: "/billing/syncDepositList", Method: http.MethodGet}
	GetDepositList   = RouteDetail{Path: "/billing/depositList", Method: http.MethodGet}
	SyncMinerFeeList = RouteDetail{Path: "/billing/syncMinerFeeList", Method: http.MethodGet}
	MinerFeeList     = RouteDetail{Path: "/billing/minerFeeList", Method: http.MethodGet}
	Withdraw         = RouteDetail{Path: "/billing/withdraw", Method: http.MethodPost}
	SyncWithdrawList = RouteDetail{Path: "/billing/syncWithdrawList", Method: http.MethodGet}
	GetWithdrawList  = RouteDetail{Path: "/billing/withdrawList", Method: http.MethodGet}
)
