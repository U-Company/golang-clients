package clients

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	json "github.com/json-iterator/go"
)

type Method interface {
	GetPath() (string, error)
	GetMethod() (string, error) // see: http.MethodGet, http.MethodPost ...
	GetHeader() map[string]string
	GetCookies() map[string]string
	GetQueryParams() map[string]string
	GetBody() (io.Reader, error)
	GetAcceptStatusCodes() []StatusCode

	ResponseProcess(body io.ReadCloser, h http.Header, s StatusCode) (*Response, error)
}

type BaseMethod struct {
	Path        string
	Count       int
	Method      string // see: http.MethodGet, http.MethodPost ...
	Headers     map[string]string
	Cookies     map[string]string
	QueryParams map[string]string
	Body        *string
	AcceptStatusCodes []StatusCode
}

func NewBaseMethod(path string, countArgs int, args ...string) *BaseMethod {
	if len(args) != countArgs {
		panic(CountArgsError)
	}
	as := make([]interface{}, 0, len(args))
	for _, a := range args {
		as = append(as, a)
	}
	p := fmt.Sprintf(path, as...)
	return &BaseMethod{
		Path:        p,
		Method:      http.MethodGet,
		Headers:     map[string]string{},
		Cookies:     map[string]string{},
		QueryParams: map[string]string{},
		AcceptStatusCodes: []StatusCode{http.StatusOK},
	}
}

func (m *BaseMethod) GetPath() (string, error) {
	if len(m.Path) == 0 {
		return "", EmptyPathError
	}
	return m.Path, nil
}

func (m *BaseMethod) GetMethod() (string, error) {
	if len(m.Method) == 0 {
		return "", EmptyMethodError
	}
	return m.Method, nil
}

func (m *BaseMethod) GetHeader() map[string]string {
	return m.Headers
}

func (m *BaseMethod) GetCookies() map[string]string {
	return m.Cookies
}

func (m *BaseMethod) GetQueryParams() map[string]string {
	return m.QueryParams
}

func (m *BaseMethod) GetBody() (io.Reader, error) {
	if m.Method == http.MethodGet && m.Body != nil {
		return nil, NotEmptyBodyError
	}
	if m.Body == nil {
		return nil, nil
	}
	return strings.NewReader(*m.Body), nil
}

func (m *BaseMethod) GetAcceptStatusCodes() []StatusCode {
	return m.AcceptStatusCodes
}

func (m *BaseMethod) ResponseProcess(body io.ReadCloser, h http.Header, s StatusCode) (*Response, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	if err := body.Close(); err != nil {
		return nil, err
	}
	api := map[string]interface{}{}
	err = json.Unmarshal(b, &api)
	if err != nil {
		return nil, err
	}
	return &Response{API: api, StatusCode: s}, nil
}
