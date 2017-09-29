package requests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
)

func init() {
}

func Request(method, url, token string, body []byte) ([]byte, int) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Debug().Msgf("response Status: %+v", resp.Status)
		panic(err)
	}
	defer resp.Body.Close()

	//log.Debug("response Headers:", resp.Header)
	rBody, _ := ioutil.ReadAll(resp.Body)
	//log.Debug("response Body:", string(rBody))
	return rBody, resp.StatusCode
}
