package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (p *PkgModules) Get(ctx context.Context, fullURI string, headers map[string]interface{}, params map[string]interface{}) ([]byte, error) {
	queryValues := make(url.Values)
	for key, value := range params {
		queryValues.Add(key, value.(string))
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", fullURI, queryValues.Encode()), nil)
	if err != nil {
		log.Println("[Pkg - Get] Error new get request, err:", err)
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value.(string))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("[Pkg - Get] Error doing request, err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[Pkg - Get] Error reading response body, err:", err)
		return nil, err
	}

	return data, nil
}

func (p *PkgModules) Post(ctx context.Context, fullURI string, headers map[string]interface{}, params map[string]interface{}) ([]byte, error) {
	form, _ := json.Marshal(params)
	req, err := http.NewRequest(http.MethodPost, fullURI, bytes.NewReader(form))
	if err != nil {
		log.Println("[Pkg - Post] Error new post request, err:", err)
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value.(string))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("[Pkg - Post] Error doing request, err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[Pkg - Post] Error reading response body, err:", err)
		return nil, err
	}

	return data, nil
}
