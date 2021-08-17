package webmsg

import (
	"encoding/json"
	"testing"
)

func TestWebBindMsg(t *testing.T)  {
	msg:=`
{
  "issue_addr": "qXgQwbo/T6hPSEJCoDP8duWeBRc=",
  "user_addr": "bJKcMWKaOnMrWmvPYJrJrDjA6U8nquAmvclfluPEJzo=",
  "n_days": 5,
  "random_id": "j33+UQlXmDOCalKntduMoVa1h2nCJJwLdbbmh2QqQGw=",
  "signature": "jQc2U0GZfeobfMeqha21LrVkwo8uH/3T96Bqa63pEZU27409DzInKGtk/ceu4R7FP7p1aZR5einlVOFZRkxsTxs="
}`


lb := &LicenseBind{}

err:=json.Unmarshal([]byte(msg),lb)

if err!=nil{
	panic(err)
}



}
