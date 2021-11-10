package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

func GetHttpRequestTime() {
	var httpMetric http_metric

	httpMetric.datetime = time.Now().Local().String()
	start := time.Now()
	req, _ := http.NewRequest("GET", "https://www.google.com", nil)

	trace := &httptrace.ClientTrace{
		DNSDone: func(di httptrace.DNSDoneInfo) {
			httpMetric.dnsResolutionDone = time.Since(start).Milliseconds()
		},
		TLSHandshakeDone: func(cs tls.ConnectionState, e error) {
			httpMetric.tlsHandshakeDone = time.Since(start).Milliseconds()
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", httpMetric)
}
