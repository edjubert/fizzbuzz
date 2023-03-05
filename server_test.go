package main_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edjubert/leboncoin/constants"
	"github.com/edjubert/leboncoin/controllers"
	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/types"
)

type MockRedis struct {
	FakePing  func() error
	FakeClose func() error

	FakeUpdateScore      func(context.Context, types.Params) error
	FakeSaveParamsAndMsg func(context.Context, types.Params, string) error
	FakeGetMsgFromParams func(context.Context, types.Params) (string, error)
	FakeMostUsedParams   func(context.Context) (types.Params, float64, error)
}

func (m *MockRedis) Ping() error                                                { return nil }
func (m *MockRedis) Close() error                                               { return nil }
func (m *MockRedis) UpdateScore(ctx context.Context, params types.Params) error { return nil }
func (m *MockRedis) SaveParamsAndMsg(ctx context.Context, params types.Params, msg string) error {
	return nil
}
func (m *MockRedis) GetMsgFromParams(ctx context.Context, params types.Params) (string, error) {
	if m.FakeGetMsgFromParams != nil {
		return m.FakeGetMsgFromParams(ctx, params)
	}

	return "", nil
}
func (m *MockRedis) GetMostUsedParams(ctx context.Context) (types.Params, float64, error) {
	return types.Params{}, 0, nil
}

func getMockedRedis() redis.Cache {
	var DB redis.Cache
	DB = &MockRedis{
		FakeGetMsgFromParams: func(context.Context, types.Params) (string, error) {
			return "", nil
		},
	}

	return DB
}

func TestFizzBuzz(t *testing.T) {
	w := httptest.NewRecorder()
	uri := constants.FIZZBUZZ

	t.Run("FizzBuzz with valid params", func(t *testing.T) {
		r := strings.NewReader(`{"int1":3,"int2":5,"limit":100,"str1":"fizz","str2":"buzz"}`)
		req, err := http.NewRequest(http.MethodPost, uri, r)
		if err != nil {
			t.Fatal(err)
		}

		redisCache := getMockedRedis()
		controllers.FizzBuzz(w, req, redisCache)

		res := w.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		fileName := "mocks/3_5_100_fizz_buzz"
		mockedFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			t.Fatalf("Could not open file %s", fileName)
		}
		if string(data) != string(mockedFile) {
			t.Error("Wrong fizzbuzz output")
		}
	})

	t.Run("FizzBuzz with missing param", func(t *testing.T) {
		r := strings.NewReader(`{"int1":3,"int2":5,"limit":100,"str1":"fizz","str2":"buzz"}`)
		req, err := http.NewRequest(http.MethodPost, uri, r)
		if err != nil {
			t.Fatal(err)
		}

		redisCache := getMockedRedis()
		controllers.FizzBuzz(w, req, redisCache)

		res := w.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if string(data) != "" {
			t.Errorf("Output should be:\n''\n\nhave:\n'%s'", string(data))
		}
	})

	t.Run("FizzBuzz with no param", func(t *testing.T) {
		r := strings.NewReader(`{}`)
		req, err := http.NewRequest(http.MethodPost, uri, r)
		if err != nil {
			t.Fatal(err)
		}

		redisCache := getMockedRedis()
		controllers.FizzBuzz(w, req, redisCache)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("Does not returned 200 -> [%d]: %s", res.StatusCode, res.Status)
		}

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if string(data) != "" {
			t.Errorf("Output should be:\n''\n\nhave:\n'%s'", string(data))
		}
	})

	t.Run("Method GET should not be implemented", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			t.Fatal(err)
		}

		controllers.FizzBuzz(w, req, nil)

		res := w.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if strings.Contains(string(data), "Not implemented") {
			t.Errorf("%s method should not be implemented", http.MethodGet)
		}
	})
}
