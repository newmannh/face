package fpp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	configPath = "apikey.cfg"
	baseUrl    = "https://apius.faceplusplus.com/v2/"
)

var (
	apiKey, apiSecret string
)

func GetRequest(path string, params map[string]string, respJson interface{}) error {
	url, err := RequestUrl(path, params)
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(respJson)
	fmt.Println(respJson)
	return err
}

func RequestUrl(path string, params map[string]string) (string, error) {
	parsedBaseUrl, err := url.Parse(baseUrl + path)
	if err != nil {
		return "", err
	}
	urlParams := url.Values{}
	credKey, credSecret, err := getCreds()
	if err != nil {
		return "", err
	}
	urlParams.Add("api_key", credKey)
	urlParams.Add("api_secret", credSecret)
	for k, v := range params {
		urlParams.Add(k, v)
	}
	parsedBaseUrl.RawQuery = urlParams.Encode()
	return parsedBaseUrl.String(), nil
}

func getCreds() (string, string, error) {
	if apiKey != "" && apiSecret != "" {
		return apiKey, apiSecret, nil
	}
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", "", err
	}
	body := string(bytes)
	if lines := strings.Split(body, "\n"); len(lines) > 2 {
		parseCred := func(line string) (string, error) {
			i1, i2 := strings.Index(line, "'"), strings.LastIndex(line, "'")
			if i1 > 0 && i2 > i1 {
				return line[i1+1 : i2], nil
			}
			return "", fmt.Errorf("Unable to parse credentials from file.")
		}
		if key, err := parseCred(lines[1]); err == nil {
			if secret, err := parseCred(lines[2]); err == nil {
				apiKey, apiSecret = key, secret
				return apiKey, apiSecret, nil
			}
		}
		return "", "", err
	}
	return "", "", fmt.Errorf("Unable to parse credential file %s", configPath)
}
