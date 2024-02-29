// package main

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/url"
// )

// type Request struct {
// 	Method string
// 	Url    *url.URL
// }

// func (req *Request) encode() ([]byte, error) {
// 	bytes, err := json.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return bytes, nil
// }

// func (req *Request) decode(bytes []byte) error {
// 	return json.Unmarshal(bytes, req)
// }

//	func (req *Request) convertToHttp() *http.Request {
//		return &http.Request{
//			Method: req.Method,
//			URL:    req.Url,
//		}
//	}
package main
