package manager

import (
	"bytes"
	"encoding/json"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
)

func httpGet(url string, res interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("http get failed,url=%v", url)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Errorf("http get resp is nil,err=%v,uri=%v", err, url)
		return err
	}
	if err := json.Unmarshal(body, res); err != nil {
		log.Errorf("http get unmarshal error,err=%v,url=%v,resp=%v", err, url, string(body))
		return err
	}
	return nil
}

func httpPost(url string, body interface{}, res interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(b)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", buf)
	if err != nil {
		log.Errorf("http post fail,err=%v,uri=%v,body=%v", err, url, string(b))
		return err
	}

	defer resp.Body.Close()
	bodyS, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
		// handle error
	}
	if err := json.Unmarshal(bodyS, res); err != nil {
		log.Errorf("json unmarshal err,err=%v", err)
		return err
	}
	return nil
}
