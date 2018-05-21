package httpWrapper
import (
//	"context"
//	"crypto/tls"
//	"encoding/json"
//	"encoding/xml"
//	"io"
//	"net"
	"net/http"
//	"net/url"
	"sync"
//	"time"
)
type basicAuth struct {
	 userName     string
	 passWord     string
 }
type HttpClient struct {
	cli         *http.Client
	auth        basicAuth
	pool        *sync.Pool
	cookies     []*http.Cookie
	headers     map[string]string
}

func NewClient() *HttpClient {
	client := &http.Client{
		Transport: http.DefaultTransport,
	}
	return &HttpClient{
		cli      : client,
		auth     : basicAuth{},
		pool     : nil,
		cookies  : make([]*http.Cookie, 0),
		headers   : make(map[string]string),
	}
}

