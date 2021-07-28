package websocket

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"

	pbs "github.com/ninjahome/ninja-go/pbs/websocket"

	"google.golang.org/protobuf/proto"
)

const(
	groupMsgDBKeyHead = "GroupMessageDBKey1_%s_%s"
	groupMsgDBKeyEnd = "GroupMessageDBKey2"

	groupMsgReceiverDBKeyHead = "GroupMsgReceiverDBKey1_%s_%s"
	groupMsgReceiverDBKeyEnd = "GroupMsgReceiverDBKey2_%s_ffffffffffffffff"
)

func int64Decimal2comparableString(i int64) string {
	buf:=make([]byte, 8)

	binary.BigEndian.PutUint64(buf,uint64(i))

	return hex.EncodeToString(buf)
}

func msg2Hash(msg *pbs.WSCryptoGroupMsg) string {
	data,_:=proto.Marshal(msg)

	s:=sha256.New()
	s.Write(data)

	h:=s.Sum(nil)

	return hex.EncodeToString(h)
}

func groupMsgDbKey(msg *pbs.WSCryptoGroupMsg) string {
	return fmt.Sprintf(groupMsgDBKeyHead,msg2Hash(msg),int64Decimal2comparableString(msg.UnixTime))
}

func SaveGroupMsg(db *leveldb.DB,msg *pbs.WSCryptoGroupMsg) (groupKey string,err error) {
	key:=groupMsgDbKey(msg)

	return key,db.Put([]byte(key),msg.MustData(),&opt.WriteOptions{Sync: true})
}

func GetGroupMsg(db *leveldb.DB, key []byte) (*pbs.WSCryptoGroupMsg,error)  {
	msg:=&pbs.WSCryptoGroupMsg{}

	if data,err:=db.Get(key,nil);err!=nil{
		return nil,err
	}else{
		if err = proto.Unmarshal(data,msg);err!=nil{
			return nil, err
		}
	}

	return msg,nil
}

func receiverGroupMsgDBKey(receiver string, msgTime int64) string {
	return fmt.Sprintf(groupMsgReceiverDBKeyHead,receiver,int64Decimal2comparableString(msgTime))
}

func StartKey(receiver string) string {
	return receiverGroupMsgDBKey(receiver,0)
}

func EndKey(receiver string) string  {
	return fmt.Sprintf(groupMsgReceiverDBKeyEnd,receiver)
}

func SaveReceiverGroupMsg(db *leveldb.DB, receiver string, groupMsgKey []byte, msgTime int64)  error{
	key:=receiverGroupMsgDBKey(receiver,msgTime)

	return db.Put([]byte(key),groupMsgKey,&opt.WriteOptions{Sync: true})
}










