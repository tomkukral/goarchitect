package goarchitect

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Real interface implementation
type RealOs struct {
}

func (o RealOs) Getenv(name string) string {
	return os.Getenv(name)
}

func (o RealOs) LogFatal(v ...interface{}) {
	log.Fatal(v)
}

func (o RealOs) HttpDo(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(req)
}

// Test inteface implementation
type TestOs struct {
	Args        []string
	EmptyConfig bool
	Status      string
	Body        string
}

func (o TestOs) Getenv(name string) string {
	var r string

	fmt.Println(o)

	if o.EmptyConfig {
		r = ""
	} else {
		r = fmt.Sprintf("valueof:%s", name)
	}

	return r
}

func (o TestOs) LogFatal(v ...interface{}) {
	o.Status = "log.Fatal"
}

func (o TestOs) HttpDo(req *http.Request) (*http.Response, error) {

	header := make(http.Header, 0)
	header.Add("Content-Type", "application/json")

	t := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(o.Body)),
		ContentLength: int64(len(o.Body)),
		Request:       req,
		Header:        header,
	}

	return t, nil
}
