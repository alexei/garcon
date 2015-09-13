package main

import (
	"flag"
	"github.com/facebookgo/grace/gracehttp"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		bind    = flag.String("bind", "127.0.0.1:8080", "Bind")
		cwd, _  = os.Getwd()
		path    = flag.String("path", cwd, "Root")
		prefix  = flag.String("prefix", "", "Prefix")
		logfile = flag.String("log", "", "Log")
	)
	flag.Parse()

	if *logfile != "" {
		f, err := os.OpenFile(*logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	log.Printf("Starting Garçon")

	mux := http.NewServeMux()
	mux.Handle("/", loggingHandler(http.StripPrefix(*prefix, http.FileServer(http.Dir(*path)))))
	gracehttp.Serve(
		&http.Server{Addr: *bind, Handler: mux},
	)
}

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var l loggingResponseWriter = &responseLogger{w: w}
		h.ServeHTTP(l, r)
		log.Printf("%s %s %s %s %d %d", r.RemoteAddr, r.Method, r.URL, r.Proto, l.Status(), l.Size())
	})
}

type loggingResponseWriter interface {
	http.ResponseWriter
	Status() int
	Size() int
}

type responseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

func (l *responseLogger) Header() http.Header {
	return l.w.Header()
}

func (l *responseLogger) Write(b []byte) (int, error) {
	size, err := l.w.Write(b)
	l.size += size
	return size, err
}

func (l *responseLogger) WriteHeader(s int) {
	l.w.WriteHeader(s)
	l.status = s
}

func (l *responseLogger) Status() int {
	return l.status
}

func (l *responseLogger) Size() int {
	return l.size
}
