package proxy

import (
	"context"
	"fmt"
	"github.com/ninjahome/ninja-go/service/proxy/httputil"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"time"
)

type WebProxyServer struct {
	listenAddr string
	proxyAddr  []string
	quit       chan struct{}
	server     *http.Server
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

func NewWebServer(networkAddr string, proxyAddr []string) *WebProxyServer {
	ws := WebProxyServer{
		listenAddr: networkAddr,
		proxyAddr:  proxyAddr,
		quit:       make(chan struct{}, 8),
	}

	return ws.init()

}

func (ws *WebProxyServer) init() *WebProxyServer {
	rh := &RegexpHander{
		routes: make([]*route, 0),
	}

	rh.HandleFunc("license", ws.proxyFunc)
	//rh.HandleFunc("pushmessage", ws.proxyFunc)

	server := &http.Server{
		Handler: rh,
	}

	ws.server = server

	return ws
}

func (ws *WebProxyServer) proxyFunc(writer http.ResponseWriter, request *http.Request) {
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
		for i := 0; i < len(ws.proxyAddr); i++ {
			proxyUrl := ws.proxyAddr[i]+request.URL.Path

			fmt.Println("proxy url:", proxyUrl)

			var result string
			var code int
			result, code, err = httputil.NewHttpPost(nil, true, 2, 20).
				ProtectPost(proxyUrl, string(contents))
			if err != nil {
				fmt.Println("---->",err)
				continue
			}
			if code != 200 {
				fmt.Println("---->",code)
				continue
			}
			writer.WriteHeader(200)
			writer.Write([]byte(result))
			return
		}
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "proxy error")

	}
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

func StartProxyDaemon() {

	webServer = _proxyConfig.NewProxyWebServer()

	fmt.Println("start proxy at ", webServer.listenAddr, "  ...")

	webServer.Start()
}

func StopProxyDaemon() {
	webServer.Shutdown()

	fmt.Println("stop proxy ...")
}
