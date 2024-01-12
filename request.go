package paystack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

const (
	// User agent used when communicating with the Paystack API.
	userAgent = "Mozilla/5.0 (Unknown; Linux) AppleWebKit/538.1 (KHTML, like Gecko) Chrome/v1.0.0 Safari/538.1"
)

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(data)
	return err
}

func postResource(ctx context.Context, c *Client, url string, body interface{}, res interface{}) error {
	reqUrl, _ := c.BaseUrl.Parse(url)
	buf, err := json.Marshal(body)
	if err != nil {
		return err
	}

	payload := bytes.NewBuffer(buf)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl.String(), payload)
	if err != nil {
		return err
	}

	return doReq(c, req, res)
}

func putResource(ctx context.Context, c *Client, url string, body interface{}, res interface{}) error {
	reqUrl, _ := c.BaseUrl.Parse(url)
	if body == nil {
		body = `{}`
	}

	buf, err := json.Marshal(body)
	if err != nil {
		return err
	}

	payload := bytes.NewBuffer(buf)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, reqUrl.String(), payload)
	if err != nil {
		return err
	}
	return doReq(c, req, res)
}

func getResource(ctx context.Context, c *Client, url string, res interface{}) error {
	reqUrl, _ := c.BaseUrl.Parse(url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return err
	}

	return doReq(c, req, res)
}

func deleteResource(ctx context.Context, c *Client, url string, res interface{}) error {
	reqUrl, _ := c.BaseUrl.Parse(url)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, reqUrl.String(), nil)
	if err != nil {
		return err
	}

	return doReq(c, req, res)
}

func doReq(c *Client, req *http.Request, res interface{}) error {

	if req.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	req.Header.Set("User-Agent", userAgent)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error processing request - %+v", err)
	}

	err = parseAPIResponse(c, resp, res)
	if err != nil {
		return err
	}

	return nil
}

func parseAPIResponse(c *Client, resp *http.Response, resultPtr interface{}) error {

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("error while reading the response bytes - %+v", err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("error closing response body - ", err)
		}
	}()

	var response APIResponse

	err = json.Unmarshal(b, &response)
	if err != nil {
		return fmt.Errorf("error while unmarshalling the response bytes %+v ", err)
	}

	if status, _ := response["status"].(bool); !status || resp.StatusCode >= 400 {
		// c.log.Error(fmt.Sprintln("Paystack response: %v\n", resp))
		return newAPIError(resp, response)
	}

	// c.log.Info(fmt.Sprintln("Paystack response: %v\n", response["data"]))
	// looking for a more betterway
	if data, ok := response["data"]; ok {
		switch t := response["data"].(type) {

		case map[string]interface{}:
			return mapstruct(data, resultPtr)
			// i
		default:
			// if response is an array
			_ = t
			return mapstruct(response, resultPtr)
		}
	}
	// if response data does not contain data key, map entire response to v
	return mapstruct(response, resultPtr)
}
