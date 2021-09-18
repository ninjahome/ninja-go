package client

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/filter"
	"github.com/btcsuite/goleveldb/leveldb/opt"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"strconv"
	"strings"
)

type ChatHistoryStore struct {
	db *leveldb.DB
	dbPath string
	historyInterval int64
	walletAddr string
}

const(
	//account key
	HistoryIntervalKey = "HistoryIntervalKey_%s"
	//from address, unixTime
	ChatMessageKey = "ChatMessage_1_%s_%s"

)

func NewStore(dbPath,walletAddr string) (*ChatHistoryStore,error)  {
	db,err:=leveldb.OpenFile(dbPath, &opt.Options{
		Filter:         filter.NewBloomFilter(10),
		ErrorIfMissing: false,
	})
	if err!=nil{
		return nil,err
	}
	return &ChatHistoryStore{
		dbPath: dbPath,
		db: db,
		walletAddr: strings.ToLower(walletAddr),
	},nil
}

func (chs *ChatHistoryStore)SetWalletAddr(addr string)  {
	chs.walletAddr = addr
}

func (chs *ChatHistoryStore)SetHistoryInterval(timeInterval int64) error {
	key:=fmt.Sprintf(HistoryIntervalKey,chs.walletAddr)

	ti:=strconv.FormatInt(timeInterval,10)

	if err:=chs.db.Put([]byte(key),[]byte(ti),&opt.WriteOptions{Sync: true});err!=nil{
		return err
	}

	if timeInterval < 0{
		chs.historyInterval = timeInterval
		return nil
	}

	if timeInterval < chs.historyInterval{
		//todo ... delete old history
	}
	chs.historyInterval = timeInterval

	return nil
}

func (chs *ChatHistoryStore)Load() error {
	key:=fmt.Sprintf(HistoryIntervalKey,chs.walletAddr)

	if data,err:=chs.db.Get([]byte(key),nil);err!=nil{
		return err
	}else{
		var ti int64

		if ti, err = strconv.ParseInt(string(data),10,64);err!=nil{
			return err
		}else{
			chs.historyInterval = ti
		}
	}

	return nil
}

func (chs *ChatHistoryStore)SaveChatMessage(msg *pbs.WSCryptoMsg, encryptData []byte) error {

	buf:=make([]byte,8)
	binary.BigEndian.PutUint64(buf,uint64(msg.UnixTime))


	key:=fmt.Sprintf(ChatMessageKey,msg.From,hex.EncodeToString(buf))

	if err:=chs.db.Put([]byte(key),encryptData,&opt.WriteOptions{Sync: true});err!=nil{
		return err
	}

	return nil
}

func (chs *ChatHistoryStore)SaveGroupChatMsg(gmsg *pbs.WSCryptoGroupMsg, groupId string, encryptData []byte) error  {
	
	return nil
}


