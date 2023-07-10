package httpserver

import (
	"net"
	"net/http"
	"sync"

	"cqserver/golibs/logger"
	"github.com/gorilla/mux"
)

type HttpServer struct {
	listener net.Listener
	closing  bool
	Router   *mux.Router
	wg       sync.WaitGroup
}

func NewHttpServer() *HttpServer {
	server := &HttpServer{Router: mux.NewRouter()}
	server.Router.StrictSlash(true)
	return server
}

func (this *HttpServer) Start(netAddr string) error {
	l, err := net.Listen("tcp", netAddr)
	if err != nil {
		return err
	}
	this.listener = l

	this.wg.Add(1)
	go func() {
		this.serve()
		this.wg.Done()
	}()
	return nil
}

func (this *HttpServer) Stop() {
	this.closing = true
	if this.listener != nil {
		this.listener.Close()
	}
	this.wg.Wait()
}

func (this *HttpServer) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	this.Router.HandleFunc(pattern, handler)
}

func (this *HttpServer) Handle(pattern string, handler http.Handler) {
	this.Router.Handle(pattern, handler)
}

func (this *HttpServer) serve() {
	err := http.Serve(this.listener, this.Router)
	if !this.closing && err != nil {
		logger.Info("http serve error: " + err.Error())
	}
}

var DefaultHttpServer = NewHttpServer()

func Start(netAddr string) error {
	return DefaultHttpServer.Start(netAddr)
}

func Stop() {
	DefaultHttpServer.Stop()
}

func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	DefaultHttpServer.HandleFunc(pattern, handler)
}
func Handle(pattern string, handler http.Handler) {
	DefaultHttpServer.Handle(pattern, handler)
}

func Router() *mux.Router {
	return DefaultHttpServer.Router
}
