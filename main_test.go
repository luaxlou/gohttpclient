package gohttpclient

import (
	"testing"
	"log"
	"net/http"
)

func TestGet(t *testing.T) {
	_, s, err := Get("https://www.baidu.com").Exec().String()

	if err != nil {
		t.Fatal(err)
	}

	log.Println(s)
}

func TestPost(t *testing.T) {

	res := Post("http://jd.com/pageNotExists").Exec()

	i, s, err := res.String()

	if i != http.StatusOK {
		t.Fatal("status code:", i)
	}

	if err != nil {
		t.Fatal("error:", err)
	}

	log.Println(s)
}
