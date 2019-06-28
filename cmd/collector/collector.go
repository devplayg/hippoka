package main

import (
	"compress/gzip"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	u := "http://abc.com"
	parsedURL, err := url.Parse(u)
	if err != nil {
	}

	req := &http.Request{
		Method:     "GET",
		URL:        parsedURL,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		//Header:     hdr,
		//Body:       nil,
		Host: parsedURL.Host,
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	contentEncoding := strings.ToLower(res.Header.Get("Content-Encoding"))
	fmt.Printf("content-encodeing: %s\n", contentEncoding)
	fmt.Printf("uncompressed: %v\n", res.Uncompressed)

	var bodyReader io.Reader = res.Body
	if !res.Uncompressed && (strings.Contains(contentEncoding, "gzip") || (contentEncoding == "" && strings.Contains(strings.ToLower((res.Header.Get("Content-Type"))), "gzip"))) {
		bodyReader, err = gzip.NewReader(bodyReader)
		if err != nil {
			panic(err)
		}
	}

	_, err = ioutil.ReadAll(bodyReader)
	if err != nil {
		panic(err)
	}
	spew.Dump(body)

	//if err != nil || resp.StatusCode >= 500 {
	//return resp, err
	//}

	//spew.Dump(res.Body)
}
