# Golang-clients

This client intend of incapsulate some routines actions from user and beautifully code. All you need:

* Create the client to any service ("http://yandex.ru")
* Implement and create method by path ("/")
* Execute request to service

Implementation example method of Yandex:

    type DefaultMethod struct{ BaseMethod }

    func NewDefaultMethod() *DefaultMethod {
	      countArgs := 0
	      m := NewBaseMethod("/", countArgs)
	      m.Method = http.MethodGet
	      m.Headers = map[string]string{"cache-control": "no-cache"}
	      return &DefaultMethod{BaseMethod: *m}
    }
  
Creating the client:
  
    url := "http://yandex.ru"
    opts := clients.Options{}
    castoramaClient := clients.NewClientUrl(url, opts)

Request:

    m := NewCatalogBasePageMethod(u.Path)
    resp, err := p.client.Request(m)
    if err != nil {
        panic("internal error of client")
    }
    if resp.StatusCode != http.StatusOK {
        panic("not OK response from Yandex")
    }
 
# Task of client

* This client must provide a code higher abstraction.
* At the same time, it will improve code readability.
* In addition, at the client level, it will be possible to implement a common logic for all parsers / services. For 
example, you need prometheus to all services. This logic can be added to any parser. Perhaps a minimal changes in the 
parser / service code is required (add arguments for initialization, for example).
* Instead of classic client, we separate logic and request.

# Notice

We use [implement](https://github.com/U-Company/python-clients) the same client for python
