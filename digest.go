package monero

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	connectTimeOut   = time.Duration(10 * time.Second)
	readWriteTimeout = time.Duration(20 * time.Second)
	userAgent        = "AtScale"
)

const (
	nc = "00000001"
)

// func (c *CallClient) Test(method string, req, rep interface{}) error {
// 	// client := DefaultTimeoutClient()
// 	// jar := &myjar{}
// 	// jar.jar = make(map[string][]*http.Cookie)
// 	// client.Jar = jar
// 	var reqest *http.Request
// 	var resp *http.Response
// 	client := &http.Client{}
// 	reqest, _ = http.NewRequest("POST", c.endpoint, EncodeClientRequest(method, req))
// 	// reqest, _ = http.NewRequest("POST", c.endpoint, nil)
// 	headers := http.Header{
// 		"User-Agent":      []string{"AtScale"},
// 		"Accept":          []string{"*/*"},
// 		"Accept-Encoding": []string{"identity"},
// 		"Connection":      []string{"Keep-Alive"},
// 		"Host":            []string{reqest.Host},
// 		"Content-Type":    []string{"application/json"},
// 	}
// 	reqest.Header = headers
// 	resp, err := client.Do(reqest)
// 	if err != nil {
// 		return err
// 	}
// 	// defer resp.Body.Close()
// 	io.Copy(ioutil.Discard, resp.Body)
// 	if resp.StatusCode == http.StatusUnauthorized {
// 		var authorization map[string]string = DigestAuthParams(resp)
// 		log.Println("authorization", authorization)
// 		realmHeader := authorization["realm"]
// 		qopHeader := authorization["qop"]
// 		nonceHeader := authorization["nonce"]
// 		algorithm := authorization["algorithm"]
// 		realm := realmHeader
// 		// A1
// 		h := md5.New()
// 		A1 := fmt.Sprintf("%s:%s:%s", c.username, realm, c.password)
// 		io.WriteString(h, A1)
// 		HA1 := hex.EncodeToString(h.Sum(nil))

// 		// A2
// 		h = md5.New()
// 		A2 := fmt.Sprintf("POST:%s", "/json_rpc")
// 		io.WriteString(h, A2)
// 		HA2 := hex.EncodeToString(h.Sum(nil))

// 		// response
// 		cnonce := RandomKey()
// 		response := H(strings.Join([]string{HA1, nonceHeader, nc, cnonce, qopHeader, HA2}, ":"))
// 		// Digest qop="auth",algorithm=MD5,realm="monero-rpc",nonce="Xv95vUKvFx+kxW0S4YR/fA==",stale=false,
// 		// Digest qop="auth",algorithm=MD5-sess,realm="monero-rpc",nonce="Xv95vUKvFx+kxW0S4YR/fA==",stale=false
// 		// Digest username="user", realm="monero-rpc", nonce="U8V3x/9EcdCD2sOkky4e5g==", uri="/", algorithm=MD5, response="624cb3e625748e2f0c2fd1ba053e438f", qop=auth, nc=00000004, cnonce="d555659c77291e68"
// 		AuthHeader := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", algorithm="%s", response="%s", qop=%s, nc=%s, cnonce="%s"`,
// 			c.username, realmHeader, nonceHeader, "/json_rpc", algorithm, response, qopHeader, nc, cnonce)
// 		log.Println("inDigest", resp.Header.Get("WWW-authenticate"))
// 		log.Println("toDigest", AuthHeader)
// 		// reqest.Header.Set("Authorization", AuthHeader)
// 		reqests, _ := http.NewRequest("POST", c.endpoint, EncodeClientRequest(method, req))
// 		headers := http.Header{
// 			"User-Agent":      []string{"AtScale"},
// 			"Accept":          []string{"*/*"},
// 			"Accept-Encoding": []string{"identity"},
// 			"Connection":      []string{"Keep-Alive"},
// 			"Host":            []string{reqest.Host},
// 			"Authorization":   []string{AuthHeader},
// 			"Content-Type":    []string{"application/json"},
// 		}
// 		reqests.Header = headers
// 		resp, err := client.Do(reqests)
// 		if err != nil {
// 			log.Println("Do2222", err)
// 			return err
// 		}
// 		defer resp.Body.Close()
// 		data, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Println("read body error:", err)
// 		}
// 		log.Println("read body data2222:", string(data), resp.StatusCode)
// 		return DecodeClientResponse(resp.Body, rep)
// 	}
// 	data, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println("read body error:", err)

// 	}
// 	log.Println("read body data11111:", string(data), resp.StatusCode)
// 	return DecodeClientResponse(resp.Body, rep)
// }

func Auth(username string, password string, uri string) (bool, error) {
	client := DefaultTimeoutClient()
	jar := &myjar{}
	jar.jar = make(map[string][]*http.Cookie)
	client.Jar = jar
	var req *http.Request
	var resp *http.Response
	var err error
	req, err = http.NewRequest("POST", uri, nil)
	if err != nil {
		log.Println("NewRequest")
		return false, err
	}
	headers := http.Header{
		"User-Agent":      []string{userAgent},
		"Accept":          []string{"*/*"},
		"Accept-Encoding": []string{"identity"},
		"Connection":      []string{"Keep-Alive"},
		"Host":            []string{req.Host},
	}
	// headers := http.Header{
	// 	"User-Agent":      []string{userAgent},
	// 	"Accept":          []string{"*/*"},
	// 	"Accept-Encoding": []string{"identity"},
	// 	"Connection":      []string{"Keep-Alive"},
	// 	"Host":            []string{req.Host},
	// 	"Authorization":   []string{AuthHeader},
	// }
	req.Header = headers

	resp, err = client.Do(req)
	if err != nil {
		log.Println("Do")
		return false, err
	}
	// you HAVE to read the whole body and then close it to reuse the http connection
	// otherwise it *could* fail in certain environments (behind proxy for instance)
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	log.Println("Do11111")
	if resp.StatusCode == http.StatusUnauthorized {
		var authorization map[string]string = DigestAuthParams(resp)
		log.Println("authorization", authorization)
		realmHeader := authorization["realm"]
		qopHeader := authorization["qop"]
		nonceHeader := authorization["nonce"]
		// opaqueHeader := authorization["opaque"]
		algorithm := authorization["algorithm"]
		realm := realmHeader
		// A1
		h := md5.New()
		A1 := fmt.Sprintf("%s:%s:%s", username, realm, password)
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
		// Digest qop="auth",algorithm=MD5,realm="monero-rpc",nonce="Xv95vUKvFx+kxW0S4YR/fA==",stale=false,
		// Digest qop="auth",algorithm=MD5-sess,realm="monero-rpc",nonce="Xv95vUKvFx+kxW0S4YR/fA==",stale=false
		// Digest username="user", realm="monero-rpc", nonce="U8V3x/9EcdCD2sOkky4e5g==", uri="/", algorithm=MD5, response="624cb3e625748e2f0c2fd1ba053e438f", qop=auth, nc=00000004, cnonce="d555659c77291e68"
		AuthHeader := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", algorithm="%s", response="%s", qop=%s, nc=%s, cnonce="%s"`,
			username, realmHeader, nonceHeader, "/json_rpc", algorithm, response, qopHeader, nc, cnonce)
		log.Println("inDigest", resp.Header.Get("WWW-authenticate"))
		log.Println("toDigest", AuthHeader)
		headers := http.Header{
			"User-Agent":      []string{userAgent},
			"Accept":          []string{"*/*"},
			"Accept-Encoding": []string{"identity"},
			"Connection":      []string{"Keep-Alive"},
			"Host":            []string{req.Host},
			"Authorization":   []string{AuthHeader},
		}
		//req, err = http.NewRequest("GET", uri, nil)
		req.Header = headers
		resp, err = client.Do(req)
		if err != nil {
			log.Println("Do2222")
			return false, err
		}
		defer resp.Body.Close()
		log.Println("444333", resp.StatusCode)
	} else {
		return false, fmt.Errorf("response status code should have been 401, it was %v", resp.StatusCode)
	}
	log.Println("333", resp.StatusCode)
	return resp.StatusCode == http.StatusOK, err
}

type myjar struct {
	jar map[string][]*http.Cookie
}

func (p *myjar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	p.jar[u.Host] = cookies
}

func (p *myjar) Cookies(u *url.URL) []*http.Cookie {
	return p.jar[u.Host]
}

func DefaultTimeoutClient() *http.Client {
	return NewTimeoutClient(connectTimeOut, readWriteTimeout)
}

func NewTimeoutClient(cTimeout time.Duration, rwTimeout time.Duration) *http.Client {
	certLocation := os.Getenv("atscale_http_sslcert")
	keyLocation := os.Getenv("atscale_http_sslkey")
	disableKeepAlives := os.Getenv("atscale_disable_keepalives")
	disableKeepAlivesBool := false
	if disableKeepAlives == "true" {
		disableKeepAlivesBool = true
	}

	// default
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	if len(certLocation) > 0 && len(keyLocation) > 0 {
		// Load client cert if available
		cert, err := tls.LoadX509KeyPair(certLocation, keyLocation)
		if err == nil {
			tlsConfig = &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
		} else {
			fmt.Printf("Error loading X509 Key Pair:%v\n", err)
		}
	}
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:   tlsConfig,
			DisableKeepAlives: disableKeepAlivesBool,
			Dial:              timeoutDialer(cTimeout, rwTimeout),
		},
	}
}

func timeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		if rwTimeout > 0 {
			conn.SetDeadline(time.Now().Add(rwTimeout))
		}
		return conn, nil
	}
}

func DigestAuthParams(r *http.Response) map[string]string {
	s := strings.SplitN(r.Header.Get("Www-Authenticate"), " ", 2)
	if len(s) != 2 || s[0] != "Digest" {
		return nil
	}

	result := map[string]string{}
	for _, kv := range strings.Split(s[1], ",") {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) != 2 {
			continue
		}
		result[strings.Trim(parts[0], "\" ")] = strings.Trim(parts[1], "\" ")
	}
	return result
}

func RandomKey() string {
	k := make([]byte, 8)
	for bytes := 0; bytes < len(k); {
		n, err := rand.Read(k[bytes:])
		if err != nil {
			panic("rand.Read() failed")
		}
		bytes += n
	}
	return base64.StdEncoding.EncodeToString(k)
}

func H(data string) string {
	digest := md5.New()
	digest.Write([]byte(data))
	return hex.EncodeToString(digest.Sum(nil))
}
