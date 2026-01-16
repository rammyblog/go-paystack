package helpers

import (
	"fmt"
	"strings"

	"github.com/rammyblog/go-paystack/internal/types"
)

func AddQueryToUrl(url string, queries ...types.QueryType) string {
	for _, query := range queries {
		if strings.Contains(url, "?") {
			url += fmt.Sprintf("&%s=%s", query.Key, query.Value)
		} else {
			url += fmt.Sprintf("?%s=%s", query.Key, query.Value)
		}
	}
	return url
}
