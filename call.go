package monero

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

// ----------------------------------------------------------------------------
// Request and Response
// ----------------------------------------------------------------------------

// clientRequest represents a JSON-RPC request sent by a client.
type clientRequest struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`

	// A String containing the name of the method to be invoked.
	Method string `json:"method"`

	// Object to pass as request parameter to the method.
	Params interface{} `json:"params"`

	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	Id uint64 `json:"id"`
}

// clientResponse represents a JSON-RPC response returned to a client.
type clientResponse struct {
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
}

type CallClient struct {
	endpoint string
	username string
	password string
}

func NewCallClient(endpoint, username, password string) *CallClient {
	return &CallClient{endpoint, username, password}
}

func (c *CallClient) Daemon(method string, req, rep interface{}) error {
	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", c.endpoint, EncodeClientRequest(method, req))
	reqest.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(reqest)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return DecodeClientResponse(resp.Body, rep)
}

func (c *CallClient) Wallet(method string, req, rep interface{}) error {
	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", c.endpoint, EncodeClientRequest(method, req))
	resp, err := client.Do(reqest)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	if resp.StatusCode == http.StatusUnauthorized {
		var authorization map[string]string = DigestAuthParams(resp)
		// log.Println("authorization", authorization)
		realmHeader := authorization["realm"]
		qopHeader := authorization["qop"]
		nonceHeader := authorization["nonce"]
		algorithm := authorization["algorithm"]
		realm := realmHeader
		// A1
		h := md5.New()
		A1 := fmt.Sprintf("%s:%s:%s", c.username, realm, c.password)
		io.WriteString(h, A1)
		HA1 := hex.EncodeToString(h.Sum(nil))

		// A2
		h = md5.New()
		A2 := fmt.Sprintf("POST:%s", "/json_rpc")
		io.WriteString(h, A2)
		HA2 := hex.EncodeToString(h.Sum(nil))

		// response
		cnonce := RandomKey()
		response := H(strings.Join([]string{HA1, nonceHeader, nc, cnonce, qopHeader, HA2}, ":"))
		AuthHeader := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", algorithm="%s", response="%s", qop=%s, nc=%s, cnonce="%s"`,
			c.username, realmHeader, nonceHeader, "/json_rpc", algorithm, response, qopHeader, nc, cnonce)
		reqests, _ := http.NewRequest("POST", c.endpoint, EncodeClientRequest(method, req))
		headers := http.Header{
			"User-Agent":      []string{"AtScale"},
			"Accept":          []string{"*/*"},
			"Accept-Encoding": []string{"identity"},
			"Connection":      []string{"Keep-Alive"},
			"Host":            []string{reqest.Host},
			"Authorization":   []string{AuthHeader},
			"Content-Type":    []string{"application/json"},
		}
		reqests.Header = headers
		resp, err := client.Do(reqests)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		// data, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Println("read body error:", err)
		// }
		// log.Println("read body data2222:", string(data), resp.StatusCode)
		return DecodeClientResponse(resp.Body, rep)
	}
	return DecodeClientResponse(resp.Body, rep)
}

// EncodeClientRequest encodes parameters for a JSON-RPC client request.
func EncodeClientRequest(method string, args interface{}) *bytes.Reader {
	c := &clientRequest{
		Version: "2.0",
		Method:  method,
		Params:  args,
		Id:      uint64(rand.Int63()),
	}
	data, _ := json.Marshal(c)
	return bytes.NewReader(data)

}

// DecodeClientResponse decodes the response body of a client request into
// the interface reply.
func DecodeClientResponse(r io.Reader, reply interface{}) error {
	var c clientResponse
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		log.Println("read Decode Error:", c)
		return err
	}
	// log.Println("read body Result:", string(*c.Result))
	if c.Error != nil {
		jsonErr := &Error{}
		if err := json.Unmarshal(*c.Error, jsonErr); err != nil {
			log.Println("read Error Error:", string(*c.Error))
			return &Error{
				Code:    E_SERVER,
				Message: string(*c.Error),
			}
		}
		log.Println("read body Error:", string(*c.Error))
		return jsonErr
	}

	if c.Result == nil {
		return ErrNullResult
	}
	// log.Println("read body Result:", string(*c.Result))
	return json.Unmarshal(*c.Result, reply)
}
