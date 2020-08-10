package gohttpclient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type GoHttpClient struct {
	req *http.Request

	body []byte

	err error

	executed bool

	debug bool

	statusCode int
}

func New() *GoHttpClient {
	c := &GoHttpClient{}

	return c
}

//Start with get
func Get(url string) *GoHttpClient {

	c := New()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		c.err = err
		return c
	}

	c.req = req

	return c

}

//Start with post
func Post(url string) *GoHttpClient {

	c := New()

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		c.err = err
		return c
	}

	c.req = req

	return c
}

//Start with Raw
func Raw(url string, bs []byte) *GoHttpClient {

	c := New()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bs))

	if err != nil {
		c.err = err
		return c
	}

	c.req = req

	return c

}

//Add query k,v
func (c *GoHttpClient) Query(k, v string) *GoHttpClient {

	q := c.req.URL.Query()

	q.Add(k, v)

	c.req.URL.RawQuery = q.Encode()
	return c

}

//Add post form k,v
func (c *GoHttpClient) Form(k, v string) *GoHttpClient {

	c.req.ParseForm()

	c.req.Form.Add(k, v)

	return c

}

//Start Request
func (c *GoHttpClient) Exec() *GoHttpClient {

	c.executed = true

	if c.req == nil {
		return c
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(c.req)

	if err != nil {
		c.err = err
		return c
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if c.debug {
		log.Println(string(body))
	}

	c.statusCode = resp.StatusCode

	if err != nil {
		c.err = err
		return c
	}

	c.body = body

	return c
}

func (c *GoHttpClient) GetError() error {

	if !c.executed {
		return errors.New("please do exec first")

	}

	if c.err != nil {

		return c.err
	}

	if c.body == nil {

		return errors.New("body is nil")
	}

	return nil
}

//render result with string
func (c *GoHttpClient) String() (int, string, error) {

	err := c.GetError()

	if err != nil {
		return c.StatusCode(), "", err
	}

	return c.StatusCode(), string(c.body), nil
}

func (c *GoHttpClient) StatusCode() int {

	if !c.executed {
		return 0

	}

	return c.statusCode
}

func (c *GoHttpClient) Bytes() (int, []byte, error) {

	return c.StatusCode(), c.body, c.GetError()
}

func (c *GoHttpClient) Debug() *GoHttpClient {

	c.debug = true
	return c
}
