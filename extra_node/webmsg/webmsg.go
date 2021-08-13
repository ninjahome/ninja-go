package webmsg

import "encoding/json"

const (
	Success = iota
	ParseJsonErr
	ConnectionErr
	CallContractErr
	OtherErr
)

type LicenseBind struct {
	IssueAddr []byte `json:"issue_addr"`
	UserAddr  []byte `json:"user_addr"`
	NDays     int32  `json:"n_days"`
	RandomId  []byte `json:"random_id"`
	Signature []byte `json:"signature"`
}

type LicenseResult struct {
	ResultCode    int32  `json:"result_code"`
	ResultMessage string `json:"result_message"`
	Tx            []byte `json:"tx"`
}

func LicenseResultPack(code int32, msg string, tx []byte) []byte {
	lr := &LicenseResult{
		ResultCode:    code,
		ResultMessage: msg,
		Tx:            tx,
	}

	j, _ := json.Marshal(lr)

	return j
}
