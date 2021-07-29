package node

import (
	"context"
	"fmt"
	badger "github.com/ipfs/go-ds-badger"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	CNM "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	dis "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-pubsub"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/ninjahome/ninja-go/node/worker"
	"github.com/ninjahome/ninja-go/service/contact"
	"github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/wallet"
	"path/filepath"
	"runtime"
	"time"
)

const (
	DefaultMaxUserNo         = 1 << 10
	MsgNoPerUser             = 1 << 6
	MsgAverageSize           = 1 << 8
	DefaultMaxMessageSize    = DefaultMaxUserNo * MsgNoPerUser * MsgAverageSize
	DefaultP2pPort           = 9999
	DefaultWorkerTryTimeOut  = 10 * time.Second
	DefaultOutboundQueueSize = 1 << 6
	DefaultIMThreadNo        = MsgNoPerUser * DefaultMaxUserNo
	DefaultLowConn           = 1 << 5
	DefaultHighConn          = 1 << 10
	DefaultConnGrace         = time.Minute

	DHTPrefix                    = "ninja"
	MainChain             ChanID = 1
	TestChain             ChanID = 2
	P2pOnLineValidateTime        = 2 * time.Second

	P2pChanDebug = "/0.1/Global/TEST"

	P2pChanUserOnOffLine  = "/0.1/Global/user/on_offline"
	P2pChanImmediateMsg   = "/0.1/Global/message/immediate"
	P2pChanUnreadMsg      = "/0.1/Global/message/unread"
	P2pChanContactOperate = "/0.1/Global/contact/operate"

	StreamContactQuery  = "/0.1/Global/contact/query"
	StreamSyncOnline    = "/0.1/Rendezvous/user/onlineSet"
	StreamSyncDevTokens = "/0.1/Rendezvous/user/deviceTokens"
)

type ChanID int

func (c ChanID) String() string {
	switch c {
	case MainChain:
		return "main network"
	case TestChain:
		return "test network"
	}
	return ""
}

var (
	MainP2pBoots = []string{"/ip4/0.0.0.0/tcp/9999/p2p/12D3KooWH1vt62wMAzSBHaAhH273MV8hnNuwF7jrDWptGzGFzPNe"}
	//TestP2pBoots = []string{"/ip4/167.179.78.33/tcp/9999/p2p/12D3KooWJ9jcvDTJGWFkjRtNLcrsQrJafTiE6mJ68hAcfbi4zp2y",
	//	"/ip4/198.13.44.159/tcp/9999/p2p/12D3KooWCRSAwwpEV96Zz1v4WiGFpeE34PZb6jZRc3yJkRrXz1Ww"}
	//TestP2pBoots = []string{
	//	"/ip4/39.99.198.143/tcp/9999/p2p/12D3KooWQgbCevCip25pjC2GZ6erWt2HuWuu2dsfctXddSxe5Bm2",
	//	"/ip4/47.113.87.58/tcp/9999/p2p/12D3KooWR2dFKCkiQCAzabKqSMR8WprH2BEJxkTnX8ZGRoHkZa1v",

	TestP2pBoots = []string{
		"/ip4/118.186.203.36/tcp/19999/p2p/12D3KooWA86Y1TSX3aMe9AbXh4Gc3ThaMBQis23sdyavS3n58SpE",
		"/ip4/39.99.198.143/tcp/19999/p2p/12D3KooWJRtcwoS3jJgBiDYtBoFgtz88ipUTWNQLHVCQRFT8uA3K",
		"/ip4/47.113.87.58/tcp/19999/p2p/12D3KooW9rMrpU6by3DdvSTuUoRR9YKrBAJM1hv78v2AfZGZPJGa",
	}
)

type pubSubConfig struct {
	MaxMsgSize         utils.ByteSize `json:"max_msg_size"`
	MaxOutQueuePerPeer int            `json:"out_queue_size"`
	MaxOnLineThread    int            `json:"online_topic_threads"`
	MaxIMTopicThread   int            `json:"im_topic_threads"`
}

func (c *pubSubConfig) String() string {
	s := fmt.Sprintf("\n\t******************Pub Sub****************")
	s += fmt.Sprintf("\n\t*max message size:\t\t%s\t*", c.MaxMsgSize)
	s += fmt.Sprintf("\n\t*max out chan per peer:\t\t%d\t*", c.MaxOutQueuePerPeer)
	s += fmt.Sprintf("\n\t*max online topic thread:\t%d\t*", c.MaxOnLineThread)
	s += fmt.Sprintf("\n\t*max IM topic thread:\t\t%d\t*", c.MaxIMTopicThread)
	s += fmt.Sprintf("\n\t*****************************************\n")
	return s
}

type dhtConfig struct {
	DataStoreFile string   `json:"cache_dir"`
	Boots         []string `json:"bootstrap"`
}

func (c *dhtConfig) String() string {
	s := fmt.Sprintf("\n\t******************DHT********************")
	s += fmt.Sprintf("\n\t*dht cache dir:%s", c.DataStoreFile)
	s += fmt.Sprintf("\n\t*boot strap nodes:%d", len(c.Boots))
	for _, boot := range c.Boots {
		s += fmt.Sprintf("\n\t%s", boot)
	}
	s += fmt.Sprintf("\n\t******************************************\n")
	return s
}

type connManagerConfig struct {
	LowWater  int           `json:"conn.low"`
	HighWater int           `json:"conn.high"`
	GraceTime time.Duration `json:"conn.grace"`
}

func (c *connManagerConfig) String() string {
	s := fmt.Sprintf("\n\t***********Connection Manager*************")
	s += fmt.Sprintf("\n\t*connection low :\t%d", c.LowWater)
	s += fmt.Sprintf("\n\t*connection high :\t%d", c.HighWater)
	s += fmt.Sprintf("\n\t*connection grace :\t%s", c.GraceTime)
	s += fmt.Sprintf("\n\t******************************************\n")
	return s

}

type Config struct {
	SrvPort            int16         `json:"port"`
	WorkerStartTimeOut time.Duration `json:"worker.start.try"`
	ChainID            ChanID
	P2oLogOpen         bool               `json:"p2pLog"`
	PsConf             *pubSubConfig      `json:"pub_sub"`
	DHTConf            *dhtConfig         `json:"dht"`
	ConnMngConf        *connManagerConfig `json:"connManager"`
	SrvPost            bool               `json:"srv_post,omitempty"`
}

func (c Config) String() string {
	s := fmt.Sprintf("\n----------------------Node Config-----------------------")
	s += fmt.Sprintf("\nworker timeout:\t\t%s", c.WorkerStartTimeOut)
	s += fmt.Sprintf("\nchain id:\t\t%d", c.ChainID)
	s += fmt.Sprintf("\nchain name:\t\t%s", c.ChainID.String())
	s += fmt.Sprintf("\np2p log open:\t\t%t", c.P2oLogOpen)
	s += fmt.Sprintf("\nnode service port:\t%d\n", c.SrvPort)
	s += fmt.Sprintf(c.PsConf.String())
	s += fmt.Sprintf(c.DHTConf.String())
	s += fmt.Sprintf(c.ConnMngConf.String())
	s += fmt.Sprintf("\n-------------------------------------------------------\n")
	return s
}

var _nodeConfig *Config = nil

func DefaultConfig(isMain bool, base string) *Config {
	var (
		isOpen  bool
		boots   []string
		dhtDir  string
		chainID ChanID
	)
	if isMain {
		isOpen = false
		boots = MainP2pBoots
		chainID = MainChain
		dhtDir = filepath.Join(base, string(filepath.Separator), "dht_cache")
	} else {
		isOpen = true
		boots = TestP2pBoots
		chainID = TestChain
		dhtDir = filepath.Join(base, string(filepath.Separator), "dht_cache_test")
	}

	return &Config{
		SrvPort:            DefaultP2pPort,
		WorkerStartTimeOut: DefaultWorkerTryTimeOut,
		ChainID:            chainID,
		P2oLogOpen:         isOpen,
		PsConf: &pubSubConfig{
			MaxMsgSize:         DefaultMaxMessageSize,
			MaxOutQueuePerPeer: DefaultOutboundQueueSize,
			MaxOnLineThread:    DefaultMaxUserNo,
			MaxIMTopicThread:   DefaultIMThreadNo,
		},
		DHTConf: &dhtConfig{
			DataStoreFile: dhtDir,
			Boots:         boots,
		},
		ConnMngConf: &connManagerConfig{
			LowWater:  DefaultLowConn,
			HighWater: DefaultHighConn,
			GraceTime: DefaultConnGrace,
		},
	}
}

func InitConfig(c *Config) {
	_nodeConfig = c
}

func (c *Config) initStreamWorker(h host.Host) {
	h.SetStreamHandler(StreamSyncOnline, websocket.Inst().OnlineMapQuery)
	h.SetStreamHandler(StreamSyncDevTokens, websocket.Inst().DevtokensQuery)
	h.SetStreamHandler(StreamContactQuery, contact.Inst().ContactQueryFromP2pNetwork)
}

func GetSrvPost() bool {
	return _nodeConfig.getSrvPost()
}

func (c *Config) getSrvPost() bool {
	return c.SrvPost
}

func (c *Config) initOptions() []libp2p.Option {

	systemTopics = map[string]worker.TopicReader{
		P2pChanUserOnOffLine:  websocket.Inst().OnOffLineForP2pNetwork,
		P2pChanImmediateMsg:   websocket.Inst().ImmediateMsgForP2pNetwork,
		P2pChanUnreadMsg:      websocket.Inst().UnreadMsgFromP2pNetwork,
		P2pChanContactOperate: contact.Inst().ContactOperationFromP2pNetwork,
	}

	if c.P2oLogOpen {
		log.SetAllLoggers(log.LevelInfo)
	}

	listenAddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", _nodeConfig.SrvPort))
	if err != nil {
		panic(err)
	}

	activeKey := wallet.Inst().KeyInUsed()
	if activeKey == nil {
		panic("no valid key right now")
	}
	key, err := activeKey.CastEd25519Key()
	if err != nil {
		panic(err)
	}

	//id,_:=peer.IDFromPrivateKey(key)

	var addressFactory func(addrs []ma.Multiaddr) []ma.Multiaddr

	var externalIP string
	externalIP, err = utils.GetExternalIP()
	if err != nil {
		utils.LogInst().Warn().Str("get external ip address", err.Error())
	} else {
		var extMultiAddr ma.Multiaddr

		extMultiAddr, err = ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d", externalIP, _nodeConfig.SrvPort))
		if err != nil {
			utils.LogInst().Warn().Str("create external multiaddr error", err.Error())
		} else {
			addressFactory = func(addrs []ma.Multiaddr) []ma.Multiaddr {
				if extMultiAddr != nil {
					addrs = append(addrs, extMultiAddr)
				}
				return addrs
			}
		}
	}

	connManager := CNM.NewConnManager(c.ConnMngConf.LowWater,
		c.ConnMngConf.HighWater,
		c.ConnMngConf.GraceTime)
	if addressFactory == nil {
		return []libp2p.Option{
			libp2p.ConnectionManager(connManager),
			libp2p.ListenAddrs(listenAddr),
			libp2p.Identity(key),
			libp2p.EnableNATService(),
			libp2p.ForceReachabilityPublic(),
		}
	} else {
		return []libp2p.Option{
			libp2p.ConnectionManager(connManager),
			libp2p.ListenAddrs(listenAddr),
			libp2p.Identity(key),
			libp2p.EnableNATService(),
			libp2p.ForceReachabilityPublic(),
			libp2p.AddrsFactory(addressFactory),
		}
	}

}

func (c *Config) pubSubOpts(disc discovery.Discovery) []pubsub.Option {
	return []pubsub.Option{
		pubsub.WithValidateQueueSize(c.PsConf.MaxOnLineThread),
		pubsub.WithPeerOutboundQueueSize(c.PsConf.MaxOutQueuePerPeer),
		pubsub.WithValidateWorkers(runtime.NumCPU() * 2),
		pubsub.WithValidateThrottle(c.PsConf.MaxOnLineThread + c.PsConf.MaxIMTopicThread),
		pubsub.WithMaxMessageSize(int(c.PsConf.MaxMsgSize)),
		pubsub.WithDiscovery(disc),
	}
}

func (c *Config) dhtOpts() ([]dht.Option, error) {
	ds, err := badger.NewDatastore(c.DHTConf.DataStoreFile, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot open Badger data store at %s, err:%s",
			c.DHTConf.DataStoreFile, err)
	}
	peers := make([]peer.AddrInfo, 0)

	for _, id := range c.DHTConf.Boots {
		addr, err := ma.NewMultiaddr(id)
		if err != nil {
			utils.LogInst().Warn().Str("invalid boot id", id)
			continue
		}
		peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			utils.LogInst().Warn().Str("parse failed for boot id", id)
			continue
		}
		peers = append(peers, *peerInfo)
	}
	if len(peers) == 0 {
		return nil, fmt.Errorf("no invalid bootstrap node")
	}

	return []dht.Option{
		dht.Datastore(ds),
		dht.ProtocolPrefix(DHTPrefix),
		dht.BootstrapPeers(peers...),
	}, nil
}

//func notForSelfValidator(_ context.Context, peer peer.ID, _ *pubsub.Message) pubsub.ValidationResult {
//	return pubsub.ValidationAccept
//}
//
//func initTopicValidators(ps *pubsub.PubSub) error {
//
//	err := ps.RegisterTopicValidator(P2pChanUserOnOffLine,
//		notForSelfValidator,
//		pubsub.WithValidatorTimeout(P2pOnLineValidateTime),
//		pubsub.WithValidatorConcurrency(_nodeConfig.PsConf.MaxOnLineThread))
//
//	if err != nil {
//		return err
//	}
//
//	err = ps.RegisterTopicValidator(P2pChanImmediateMsg,
//		notForSelfValidator,
//		pubsub.WithValidatorConcurrency(_nodeConfig.PsConf.MaxIMTopicThread))
//	if err != nil {
//		return err
//	}
//
//	if err := ps.RegisterTopicValidator(P2pChanUnreadMsg,
//		notForSelfValidator); err != nil {
//		return err
//	}
//	return nil
//}

func newPubSub(ctx context.Context, h host.Host) (*pubsub.PubSub, error) {
	dhtOpts, err := _nodeConfig.dhtOpts()

	kademliaDHT, err := dht.New(ctx, h, dhtOpts...)
	if err != nil {
		return nil, err
	}

	disc := dis.NewRoutingDiscovery(kademliaDHT)

	psOption := _nodeConfig.pubSubOpts(disc)

	ps, err := pubsub.NewGossipSub(ctx, h, psOption...)
	if err != nil {
		return nil, err
	}

	if err := kademliaDHT.Bootstrap(ctx); err != nil {
		return nil, err
	}

	//if err := initTopicValidators(ps); err != nil {
	//	return nil, err
	//}

	return ps, nil
}
