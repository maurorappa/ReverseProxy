package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	open_sesame bool = false
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func setup_logging(conf *config) {

	f, err := os.OpenFile(conf.Logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}
}

func main() {

	confPath := flag.String("conf", "rp.conf", "Configuration file")
	verbose := flag.Bool("verbose", false, "Enable logging")
	flag.Parse()

	conf, err := loadConfig(*confPath)
	if err != nil {
		log.Fatalln(err)
	}

	if *verbose {
		conf.Verbose = true
	}
	setup_logging(conf)
	fmt.Printf("Connecting to %s...\n", conf.Backend)
	_, err = http.Get(conf.Backend)
	if err != nil {
		fmt.Println("Cannot connect!")
		os.Exit(3)
	}
	origin, _ := url.Parse(conf.Backend)
	open_path := randSeq(conf.Path_len)
	close_path := randSeq(conf.Path_len)
	fmt.Printf("Open sesame path will be: %s\n", open_path)
	fmt.Printf("Close sesame path will be: %s\n", close_path)

	director := func(req *http.Request) {
		//req.Header.Add("X-Forwarded-Host", req.Host)
		//req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host

		log.WithFields(log.Fields{
			"IP":     req.RemoteAddr,
			"Method": req.Method,
			"URL":    req.URL.Path,
		}).Info("Req")
		if req.URL.Path == "/"+open_path+"/" {
			open_sesame = true
			fmt.Printf("Called Open sesame path\n")
		}
		if req.URL.Path == "/"+close_path+"/" {
			open_sesame = false
			fmt.Printf("Called close sesame path \n")
		}
	}

	proxy := &httputil.ReverseProxy{Director: director}
	fmt.Printf("Listening to %s\n", conf.Listen)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf("%v\n",open_sesame)
		if inArray(r.Method, conf.Forbidden_methods) {
			log.WithFields(log.Fields{
				"IP":     r.RemoteAddr,
				"Method": r.Method,
				"URL":    r.URL.Path,
			}).Warn("Security")
			w.WriteHeader(401)
		} else if strings.Contains(r.URL.Path, conf.Path) && !open_sesame {
			fmt.Println("closed!")
			log.WithFields(log.Fields{
				"IP":     r.RemoteAddr,
				"Method": r.Method,
				"URL":    r.URL.Path,
			}).Warn("Admin")
			w.WriteHeader(404)
		} else {
			proxy.ServeHTTP(w, r)
		}
	})

	log.Fatal(http.ListenAndServe(conf.Listen, limit(conf,mux)))
}
