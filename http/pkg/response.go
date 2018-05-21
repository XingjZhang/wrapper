package httpWrapper

import (
//	"encoding/json"
//	"encoding/xml"
	"io"
//	"io/ioutil"
	"net/http"
)

type HttpResponse struct {
	rawurl         string
	body           io.ReadCloser
	header         http.Header
	status         string
	statuscode     int
}


