package newsapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const apiKeyHeader = "X-Api-Key"
const baseURL = "https://newsapi.org/v2"
const timeoutLimit = 3 * time.Second

var ErrMissingAPIKey = errors.New("Api Key Missing")
var ErrNotEnoughArguments = errors.New("Not enough arguments")

type NewsAPIClient struct {
	apiKey string
	client *http.Client
}

func New(apiKey string) (*NewsAPIClient, error) {
	if len(apiKey) == 0 {
		return nil, ErrMissingAPIKey
	}

	cli := NewsAPIClient{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: timeoutLimit,
		},
	}
	return &cli, nil
}

func prepareURL(baseUrl string, args ...string) (preparedUrl string) {
	var params string
	for i, param := range args {
		params += param
		if i < len(args)-1 {
			params += "&"
		}
	}
	url := baseUrl + url.PathEscape(params)
	return url
}

func parseOptions(opts Options) url.Values {

	urlValues := url.Values{}

	v := reflect.ValueOf(opts)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i).Name
		value := v.Field(i)

		fieldLower := strings.ToLower(string(field))

		switch value.Kind() {

		case reflect.String:
			if len(value.String()) > 0 {
				urlValues.Add(fieldLower, value.String())
			}

		case reflect.Int64:
			intString := strconv.FormatInt(value.Int(), 10)
			if len(intString) > 0 {
				urlValues.Add(fieldLower, value.String())
			}
		}
	}

	return urlValues
}

func (c *NewsAPIClient) prepareRequest(endpoint string, opts Options) (*http.Request, error) {
	urlValues := parseOptions(opts)
	url := fmt.Sprintf("%s/%s?%s", baseURL, endpoint, urlValues.Encode())
	return http.NewRequest(http.MethodGet, url, nil)

}

func (c *NewsAPIClient) doRequest(req *http.Request) (*NewsAPIResponse, error) {
	req.Header.Add(apiKeyHeader, c.apiKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	news := new(NewsAPIResponse)

	err = json.NewDecoder(resp.Body).Decode(news)
	return news, err

}

func (c *NewsAPIClient) GetTopHeadlines(opts Options) (*NewsAPIResponse, error) {
	req, err := c.prepareRequest("top-headlines", opts)
	if err != nil {
		return nil, err
	}
	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, err

}

func (c *NewsAPIClient) GetEverything(opts Options) (*NewsAPIResponse, error) {
	req, err := c.prepareRequest("everything", opts)
	if err != nil {
		return nil, err
	}
	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, err

}
