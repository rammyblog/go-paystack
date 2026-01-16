package paystack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	_http "net/http"

	"github.com/mitchellh/mapstructure"
	"github.com/rammyblog/go-paystack/internal/http"
	"github.com/rammyblog/go-paystack/internal/types"
)

const (
	// User agent used when communicating with the Paystack API.
	userAgent = "go-paystack/1.0.0"
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

func (c *Client) PostResource(ctx context.Context, url string, body interface{}, res interface{}) error {
	reqUrl, err := c.BaseUrl.Parse(url)
	if err != nil {
		return fmt.Errorf("invalid URL %q: %w", url, err)
	}
	buf, err := json.Marshal(body)
	if err != nil {
		return err
	}

	payload := bytes.NewBuffer(buf)
	req, err := _http.NewRequestWithContext(ctx, _http.MethodPost, reqUrl.String(), payload)
	if err != nil {
		return err
	}

	return c.doReq(req, res)
}

func (c *Client) PutResource(ctx context.Context, url string, body interface{}, res interface{}) error {
	reqUrl, err := c.BaseUrl.Parse(url)
	if err != nil {
		return fmt.Errorf("invalid URL %q: %w", url, err)
	}
	if body == nil {
		body = `{}`
	}

	buf, err := json.Marshal(body)
	if err != nil {
		return err
	}

	payload := bytes.NewBuffer(buf)
	req, err := _http.NewRequestWithContext(ctx, _http.MethodPut, reqUrl.String(), payload)
	if err != nil {
		return err
	}
	return c.doReq(req, res)
}

func (c *Client) GetResource(ctx context.Context, url string, res interface{}) error {
	reqUrl, err := c.BaseUrl.Parse(url)
	if err != nil {
		return fmt.Errorf("invalid URL %q: %w", url, err)
	}

	req, err := _http.NewRequestWithContext(ctx, _http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return err
	}

	return c.doReq(req, res)
}

func (c *Client) DeleteResource(ctx context.Context, url string, res interface{}) error {
	reqUrl, err := c.BaseUrl.Parse(url)
	if err != nil {
		return fmt.Errorf("invalid URL %q: %w", url, err)
	}

	req, err := _http.NewRequestWithContext(ctx, _http.MethodDelete, reqUrl.String(), nil)
	if err != nil {
		return err
	}

	return c.doReq(req, res)
}

func (c *Client) doReq(req *_http.Request, res interface{}) error {

	if req.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	req.Header.Set("User-Agent", userAgent)

	c.logger.WithContext(ctx).Debug("sending request",
		"method", req.Method,
		"url", req.URL.String(),
		"headers", req.Header,
	)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		c.logger.Error("request failed",
			"error", err,
			"method", req.Method,
			"url", req.URL.String(),
		)
		return fmt.Errorf("error processing request - %+v", err)
	}

	defer resp.Body.Close()

	err = c.parseAPIResponse(resp, res)
	if err != nil {
		c.logger.Error("failed to parse response",
			"error", err,
			"status", resp.Status,
		)
		return err
	}

	return nil
}

func (c *Client) parseAPIResponse(resp *_http.Response, resultPtr interface{}) error {
	var response types.APIResponse
	err := json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return fmt.Errorf("error while unmarshalling the response bytes %+v ", err)
	}

	if status, _ := response["status"].(bool); !status || resp.StatusCode >= 400 {
		return http.NewAPIError(resp, response)
	}

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
