package thirdhttp

import (
	"context"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

func Curl[T any](ctx context.Context, uri string, reqData interface{}, headers map[string]string) *T {
	client := HttpPoolsGet()
	defer HttpPoolsPut(client)
	htc := client.R() // 设置上下文
	if ctx != nil {
		htc = htc.SetContext(ctx)
	}
	if headers != nil {
		htc.SetHeaders(headers)
	}
	var rps *resty.Response
	var err error
	if reqData != nil {
		rps, err = htc.SetBody(reqData).Post(uri)
	} else {
		rps, err = htc.Get(uri)
	}
	if err != nil {
		// 日志记录
		return nil
	}
	var rt T
	if err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(rps.Body(), &rt); err != nil {
		// 日志记录
		return nil
	}
	return &rt
}

func Curl2[T any](ctx context.Context, uri string, reqData interface{}, headers map[string]string) (*T, error) {
	client := HttpPoolsGet()
	defer HttpPoolsPut(client)
	htc := client.R()
	if ctx != nil {
		htc = htc.SetContext(ctx)
	}
	if headers != nil {
		htc.SetHeaders(headers)
	}
	var rps *resty.Response
	var err error
	if reqData != nil {
		rps, err = htc.SetBody(reqData).Post(uri)
	} else {
		rps, err = htc.Get(uri)
	}
	if err != nil {
		// 日志记录
		return nil, err
	}
	var rt T
	if err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(rps.Body(), &rt); err != nil {
		// 日志记录
		return nil, err
	}
	return &rt, nil
}
