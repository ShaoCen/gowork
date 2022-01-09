package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "net/http/pprof"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Info("Starting http server...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("entering root handler")

	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		w.Header().Set(k, strings.Join(v, ","))
	}

	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.WriteHeader(200)

	fmt.Println(fmt.Sprintf("Remote Addr=%s, Reponse Code=%s\n", r.RemoteAddr, "200"))
	glog.Info(fmt.Sprintf("Remote Addr=%s, Reponse Code=%s\n", r.RemoteAddr, "200"))
}
