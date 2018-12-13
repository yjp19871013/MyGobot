package orvibo

import (
	"net/url"
	"sort"
	"strings"
)

func CalSign(uri string, method string, params map[string]string, appKey string) string {
	uriEncoded := url.QueryEscape(uri)

	keys := make([]string, 0)
	for key := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	paramsStr := ""
	for _, key := range keys {
		value := params[key]
		paramsStr += key + "=" + value + "&"
	}

	strLen := len(paramsStr)
	if strLen > 0 {
		paramsStr = paramsStr[:strLen-1]
	}

	paramsStrEncoded := url.QueryEscape(paramsStr)
	source := strings.ToUpper(method) + "&" + uriEncoded + "&" + paramsStrEncoded
	signKey := appKey + "&"

	return GetHmacSha1(signKey, source)
}
