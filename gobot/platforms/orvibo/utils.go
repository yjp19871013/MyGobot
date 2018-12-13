package orvibo

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/satori/go.uuid"
)

func PostJson(uri string, jsonStr string) (string, error) {
	resp, err := http.Post(uri, "application/json", strings.NewReader(jsonStr))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetJson(uri string) (string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetUUID() string {
	u, err := uuid.NewV4()
	if err != nil {
		return ""
	}

	return strings.Replace(u.String(), "-", "", -1)
}

func ParamsMapToString(params map[string]string) string {
	str := ""
	for key, value := range params {
		str += key + "=" + value + "&"
	}

	strLen := len(str)
	if strLen > 0 {
		str = str[:strLen-1]
	}

	return str
}

func GetHmacSha1(key string, str string) string {
	keyBytes := []byte(key)
	mac := hmac.New(sha1.New, keyBytes)
	mac.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
