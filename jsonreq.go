package gohttpclient

import (
	"encoding/json"
)

func (c *GoHttpClient) addApplicationJsonHeader() {
	c.req.Header.Set("Content-Type", "application/json")

}

//Render result with json
func (c *GoHttpClient) RenderJSON(resObj interface{}) (int, error) {

	return c.StatusCode(), json.Unmarshal(c.body, resObj)

}

//Start with post a json object body
func PostBody(url string, reqObj interface{}) *GoHttpClient {

	bs, err := json.Marshal(reqObj)

	c := Raw(url, bs)

	if err != nil {
		return c
	}

	c.addApplicationJsonHeader()

	return c
}
