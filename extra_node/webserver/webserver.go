package webserver

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ninjahome/bls-wallet/bls"
	"github.com/ninjahome/ninja-go/extra_node/config"
	"github.com/ninjahome/ninja-go/extra_node/ethwallet"
	"github.com/ninjahome/ninja-go/extra_node/webmsg"

	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"time"
)

const (
	LicenseAddPath      = "/license/add"
	LicenseTransferPath = "/license/transfer"
	PushMessage         = "/ipush"
)

type WebProxyServer struct {
	listenAddr string
	quit       chan struct{}
	server     *http.Server
	wallet     ethwallet.Wallet
}

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHander struct {
	routes []*route
}

func (rh *RegexpHander) Handle(pattern string, handler http.Handler) {
	rh.routes = append(rh.routes, &route{pattern: regexp.MustCompilePOSIX(pattern), handler: handler})
}

func (rh *RegexpHander) HandleFunc(pattern string, handleFunc func(http.ResponseWriter, *http.Request)) {
	rh.routes = append(rh.routes, &route{pattern: regexp.MustCompilePOSIX(pattern), handler: http.HandlerFunc(handleFunc)})
}

func (rh *RegexpHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rh.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

func NewWebServer(networkAddr string, w ethwallet.Wallet) *WebProxyServer {
	ws := WebProxyServer{
		listenAddr: networkAddr,
		wallet:     w,
		quit:       make(chan struct{}, 8),
	}

	return ws.init()

}

func (ws *WebProxyServer) init() *WebProxyServer {
	rh := &RegexpHander{
		routes: make([]*route, 0),
	}

	rh.HandleFunc(LicenseAddPath, ws.addLicense)
	rh.HandleFunc(PushMessage, ws.pushMessage)
	rh.HandleFunc(LicenseTransferPath, ws.transferLicense)

	server := &http.Server{
		Handler: rh,
	}

	ws.server = server

	return ws
}

func (ws *WebProxyServer) pushMessage(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a post request")
		return
	}

	if contents, err := ioutil.ReadAll(request.Body); err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "read http body error")
		return
	} else {
		fmt.Println(string(contents))
		//todo...
	}
	return
}

func (ws *WebProxyServer) transferLicense(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a post request")
		return
	}

	if contents, err := ioutil.ReadAll(request.Body); err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "read http body error")
		return
	} else {

		fmt.Println(string(contents))

		tl := &webmsg.TransferLicense{}

		err = json.Unmarshal(contents, tl)
		if err != nil {
			writer.WriteHeader(200)
			writer.Write(webmsg.LicenseResultPack(webmsg.ParseJsonErr, "parse json error", nil))
			return
		}

		fmt.Println("from:", "0x"+hex.EncodeToString(tl.From))
		fmt.Println("to:", "0x"+hex.EncodeToString(tl.To))
		fmt.Println("nDays:", tl.NDays)
		fmt.Println("sig:", "0x"+hex.EncodeToString(tl.Signature))

		if b := ws.verifySignature(tl); !b {
			err = errors.New("sig not correct")
			writer.WriteHeader(200)
			writer.Write(webmsg.LicenseResultPack(webmsg.SigNatureNotCorrectErr, "signature not correct", nil))
			return
		}

		var tx []byte
		tx, err = ws._transferLicense(tl)
		if err != nil {
			fmt.Println("tx err", err)
			writer.WriteHeader(200)
			writer.Write(webmsg.LicenseResultPack(webmsg.CallContractErr, "call contract error", nil))
			return
		}

		writer.WriteHeader(200)
		writer.Write(webmsg.LicenseResultPack(webmsg.Success, "success", tx))
	}
}

func (ws *WebProxyServer) addLicense(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a post request")
		return
	}

	if contents, err := ioutil.ReadAll(request.Body); err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "read http body error")
		return
	} else {

		fmt.Println(string(contents))

		lb := &webmsg.LicenseBind{}

		err = json.Unmarshal(contents, lb)
		if err != nil {
			writer.WriteHeader(200)
			writer.Write(webmsg.LicenseResultPack(webmsg.ParseJsonErr, "parse json error", nil))
			return
		}

		fmt.Println("issue", hex.EncodeToString(lb.IssueAddr))
		fmt.Println("user:", hex.EncodeToString(lb.UserAddr))
		fmt.Println("randid:", hex.EncodeToString(lb.RandomId))
		fmt.Println("ndays:", lb.NDays)
		fmt.Println("signature:", hex.EncodeToString(lb.Signature))

		var tx []byte
		tx, err = ws.bind(lb)
		if err != nil {
			fmt.Println("tx err", err)
			writer.WriteHeader(200)
			writer.Write(webmsg.LicenseResultPack(webmsg.CallContractErr, "call contract error", nil))
			return
		}

		writer.WriteHeader(200)
		writer.Write(webmsg.LicenseResultPack(webmsg.Success, "success", tx))

	}
}

func (ws *WebProxyServer) bind(lb *webmsg.LicenseBind) (tx []byte, err error) {
	var (
		issueAddr common.Address
		userAddr  [32]byte
		randomId  [32]byte
	)
	copy(issueAddr[:], lb.IssueAddr)
	copy(userAddr[:], lb.UserAddr)
	copy(randomId[:], lb.RandomId)

	return Bind(issueAddr, userAddr, randomId, lb.NDays, lb.Signature, ws.wallet.SignKey())
}

func (ws *WebProxyServer) verifySignature(tl *webmsg.TransferLicense) bool {

	msg := make([]byte, 1024)
	n := copy(msg, tl.From)
	n += copy(msg[n:], tl.To)
	binary.BigEndian.PutUint32(msg[n:], uint32(tl.NDays))
	n += 4

	sig := &bls.Sign{}

	if err := sig.Deserialize(tl.Signature); err != nil {
		return false
	}
	p := &bls.PublicKey{}
	if err := p.Deserialize(tl.From); err != nil {
		fmt.Println(err)
		return false
	}

	return sig.VerifyByte(p, msg[:n])
}

func (ws *WebProxyServer) _transferLicense(tl *webmsg.TransferLicense) (tx []byte, err error) {
	var (
		from, to [32]byte
	)

	copy(from[:], tl.From)
	copy(to[:], tl.To)

	return TransferLicense(from, to, tl.NDays, ws.wallet.SignKey())
}

func (ws *WebProxyServer) Start() error {
	if l, err := net.Listen("tcp4", ws.listenAddr); err != nil {
		panic("start wss server failed")
	} else {
		return ws.server.Serve(l)
	}
}

func (ws *WebProxyServer) Shutdown() error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ws.server.Shutdown(ctx)
}

var webServer *WebProxyServer

func StartWebDaemon(w ethwallet.Wallet) {

	c := config.GetExtraConfig()

	webServer = NewWebServer(c.ListenAddr, w)

	fmt.Println("start proxy at ", webServer.listenAddr, "  ...")

	webServer.Start()
}

func StopWebDaemon() {
	webServer.Shutdown()

	fmt.Println("stop proxy ...")
}
