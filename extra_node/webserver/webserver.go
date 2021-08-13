package webserver

import (
	"context"
	"fmt"
	"github.com/ninjahome/ninja-go/extra_node/config"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"time"
)

type WebProxyServer struct {
	listenAddr string
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

func NewWebServer(networkAddr string) *WebProxyServer {
	ws := WebProxyServer{
		listenAddr: networkAddr,
		quit:       make(chan struct{}, 8),
	}

	return ws.init()

}

func (ws *WebProxyServer) init() *WebProxyServer {
	rh := &RegexpHander{
		routes: make([]*route, 0),
	}

	rh.HandleFunc("license/add", ws.addLicense)
	//rh.HandleFunc("pushmessage", ws.proxyFunc)

	server := &http.Server{
		Handler: rh,
	}

	ws.server = server

	return ws
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

		writer.WriteHeader(500)
		fmt.Fprintf(writer, "server")

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

func StartWebDaemon() {

	c := config.GetExtraConfig()

	webServer = NewWebServer(c.ListenAddr)

	fmt.Println("start proxy at ", webServer.listenAddr, "  ...")

	webServer.Start()
}

func StopWebDaemon() {
	webServer.Shutdown()

	fmt.Println("stop proxy ...")
}
