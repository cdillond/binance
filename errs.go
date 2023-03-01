package binance

import (
	"encoding/json"
	"errors"
	"fmt"
)

type RespErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (r RespErr) Error() string {
	return codeToError(r.Code).Error()
}

func parseRespErr(b []byte) error {
	var e RespErr
	err := json.Unmarshal(b, &e)
	if err != nil {
		return fmt.Errorf("%w %v", ErrRespParse, string(b))
	}
	return e
}

var (
	ErrUnknown           = errors.New("- 1000 UNKNOWN")
	ErrDisconnect        = errors.New("- 1001 DISCONNECTED")
	ErrUnauth            = errors.New("- 1002 UNAUTHORIZED")
	ErrTooManyReqs       = errors.New("- 1003 TOO_MANY_REQUESTS")
	ErrUnexpected        = errors.New("- 1006 UNEXPECTED_RESP")
	ErrTimeOut           = errors.New("- 1007 TIMEOUT")
	ErrBusy              = errors.New("- 1008 SERVER_BUSY")
	ErrOrderComp         = errors.New("- 1014 UNKNOWN_ORDER_COMPOSITION")
	ErrTooManyOrd        = errors.New("- 1015 TOO_MANY_ORDERS")
	ErrServiceDown       = errors.New("- 1016 SERVICE_SHUTTING_DOWN")
	ErrUnsuported        = errors.New("- 1020 UNSUPPORTED_OPERATION")
	ErrTimestamp         = errors.New("- 1021 INVALID_TIMESTAMP")
	ErrSignature         = errors.New("- 1022 INVALID_SIGNATURE")
	ErrIllChar           = errors.New("- 1100 ILLEGAL_CHARS")
	ErrTooManyPar        = errors.New("- 1101 TOO_MANY_PARAMETERS")
	ErrBadParam          = errors.New("- 1102 MANDATORY_PARAM_EMPTY_OR_MALFORMED")
	ErrUnknownParam      = errors.New("- 1103 UNKNOWN_PARAM")
	ErrUnreadParams      = errors.New("- 1104 UNREAD_PARAMETERS")
	ErrEmptyParam        = errors.New("- 1105 PARAM_EMPTY")
	ErrParamNR           = errors.New("- 1106 PARAM_NOT_REQUIRED")
	ErrParamOverflow     = errors.New("- 1108 PARAM_OVERFLOW")
	ErrBadPrecision      = errors.New("- 1111 BAD_PRECISION")
	ErrNoDepth           = errors.New("- 1112 NO_DEPTH")
	ErrTIFNR             = errors.New("- 1114 TIF_NOT_REQUIRED")
	ErrBadTIF            = errors.New("- 1115 INVALID_TIF")
	ErrBadOrderType      = errors.New("- 1116 INVALID_ORDER_TYPE")
	ErrBadSide           = errors.New("- 1117 INVALID_SIDE")
	ErrEmptyNewClOId     = errors.New("- 1118 EMPTY_NEW_CL_ORD_ID")
	ErrEmptyOrigClOId    = errors.New("- 1119 EMPTY_ORG_CL_ORD_ID")
	ErrBadInterval       = errors.New("- 1120 BAD_INTERVAL")
	ErrBadSymbol         = errors.New("- 1121 BAD_SYMBOL")
	ErrBadListenKey      = errors.New("- 1125 INVALID_LISTEN_KEY")
	ErrMoreThanXXHrs     = errors.New("- 1127 MORE_THAN_XX_HOURS")
	ErrBadParamsCombo    = errors.New("- 1128 OPTIONAL_PARAMS_BAD_COMBO")
	ErrReqBadParam       = errors.New("- 1130 INVALID_PARAMETER")
	ErrBadJson           = errors.New("- 1135 INVALID_JSON")
	ErrNewOrderReject    = errors.New("- 2010 NEW_ORDER_REJECTED")
	ErrCancelReject      = errors.New("- 2011 CANCEL_REJECTED")
	ErrNoSuchOrder       = errors.New("- 2013 NO_SUCH_ORDER")
	ErrBadApiKey         = errors.New("- 2014 BAD_API_KEY_FMT")
	ErrBadMbxKey         = errors.New("- 2015 REJECTED_MBX_KEY")
	ErrTradingWindow     = errors.New("- 2016 NO_TRADING_WINDOW")
	ErrPartialFail       = errors.New("- 2021 Order cancel-replace partially failed")
	ErrCancelReplaceFail = errors.New("- 2022 Order cancel-replace failed")
	ErrMsgReceived       = errors.New("- 1010 ERROR_MSG_RECEIVED")
	ErrWsUknownProperty  = errors.New("0 unknown property")
	ErrWsBadValueType    = errors.New("1 invalid value type")
	ErrWsBadRequest      = errors.New("2 invalid request")
	ErrWsBadJson         = errors.New("3 invalid json")

	ErrRespParse = errors.New("client could not parse error response from binance api")
)

func codeToError(i int) error {
	switch i {
	case -1000:
		return ErrUnknown
	case -1001:
		return ErrDisconnect
	case -1002:
		return ErrUnauth
	case -1003:
		return ErrTooManyReqs
	case -1006:
		return ErrUnexpected
	case -1007:
		return ErrTimeOut
	case -1008:
		return ErrBusy
	case -1010:
		return ErrMsgReceived
	case -1014:
		return ErrOrderComp
	case -1015:
		return ErrTooManyOrd
	case -1016:
		return ErrServiceDown
	case -1020:
		return ErrUnsuported
	case -1021:
		return ErrTimestamp
	case -1022:
		return ErrSignature
	case -1100:
		return ErrIllChar
	case -1101:
		return ErrTooManyPar
	case -1102:
		return ErrBadParam
	case -1103:
		return ErrUnknownParam
	case -1104:
		return ErrUnreadParams
	case -1105:
		return ErrEmptyParam
	case -1106:
		return ErrParamNR
	case -1108:
		return ErrParamOverflow
	case -1111:
		return ErrBadPrecision
	case -1112:
		return ErrNoDepth
	case -1114:
		return ErrTIFNR
	case -1115:
		return ErrBadTIF
	case -1116:
		return ErrBadOrderType
	case -1117:
		return ErrBadSide
	case -1118:
		return ErrEmptyNewClOId
	case -1119:
		return ErrEmptyOrigClOId
	case -1120:
		return ErrBadInterval
	case -1121:
		return ErrBadSymbol
	case -1125:
		return ErrBadListenKey
	case -1127:
		return ErrMoreThanXXHrs
	case -1128:
		return ErrBadParamsCombo
	case -1130:
		return ErrReqBadParam
	case -1135:
		return ErrBadJson
	case -2010:
		return ErrNewOrderReject
	case -2011:
		return ErrCancelReject
	case -2013:
		return ErrNoSuchOrder
	case -2014:
		return ErrBadApiKey
	case -2015:
		return ErrBadMbxKey
	case -2016:
		return ErrTradingWindow
	case -2021:
		return ErrPartialFail
	case -2022:
		return ErrCancelReplaceFail
	case 0:
		return ErrWsUknownProperty
	case 1:
		return ErrWsBadValueType
	case 2:
		return ErrWsBadRequest
	case 3:
		return ErrWsBadJson
	default:
		return ErrRespParse
	}
}
