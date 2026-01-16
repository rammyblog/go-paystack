package helpers

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/rammyblog/go-paystack/internal/types"
)

func AddQueryToUrl(baseUrl string, queries ...types.QueryType) string {
	for _, query := range queries {
		if strings.Contains(baseUrl, "?") {
			baseUrl += fmt.Sprintf("&%s=%s", query.Key, url.QueryEscape(query.Value))
		} else {
			baseUrl += fmt.Sprintf("?%s=%s", query.Key, url.QueryEscape(query.Value))
		}
	}
	return baseUrl
}
