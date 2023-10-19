package http

import (
	"fmt"
	"net/http"

	"github.com/tpcarlsen-code/mon2http/storage"
)

const ctAJ = "application/json"
const ctTP = "text/plain"

type Server struct {
	port  int
	token string
	ss    *storage.Status
	vs    *storage.Values
}

func NewServer(port int, token string, ss *storage.Status, vs *storage.Values) *Server {
	return &Server{
		port:  port,
		token: token,
		ss:    ss,
		vs:    vs,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/status", s.status)
	http.HandleFunc("/metrics", s.metrics)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}

func (s *Server) checkToken(token string, r *http.Request) bool {
	if s.token == "" {
		return true
	}
	return r.URL.Query().Get("token") == token || r.Header.Get("authorization") == "Bearer "+token
}

func (s *Server) wantJSON(r *http.Request) bool {
	return r.URL.Query().Get("json") == "1" || r.URL.Query().Get("json") == "true"
}

func (s *Server) status(rw http.ResponseWriter, r *http.Request) {
	var ct string
	var b []byte
	if s.wantJSON(r) {
		ct = ctAJ
		b = s.ss.Get().Json()
	} else {
		ct = ctTP
		b = []byte(s.ss.Get().Txt())
	}

	rw.Header().Add("content-type", ct)
	rw.WriteHeader(200)
	rw.Write(b)
}

func (s *Server) metrics(rw http.ResponseWriter, r *http.Request) {
	var ct string
	var b []byte
	if s.wantJSON(r) {
		ct = ctAJ
		b = s.vs.Get().Json()
	} else {
		ct = ctTP
		b = []byte(s.vs.Get().Metrics())
	}

	rw.Header().Add("content-type", ct)
	rw.WriteHeader(200)
	rw.Write(b)
}
