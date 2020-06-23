package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	vegeta "github.com/tsenart/vegeta/lib"
)

type LoginReqBody struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProductID  string `json:"product_id"`
	CategoryID string `json:"category_id"`
	GroupID    string `json:"group_id"`
}

type MyTarget struct {
	Method string
	URL    string
	Body   interface{}
	Header http.Header
}

var userTokenMap = make(map[string]string) // user_id:token pair

func NewScenarioTargeter() vegeta.Targeter {
	// この例では2つだけのエンドポイントをリクエストするシナリオ
	targets := []MyTarget{
		{
			Method: "POST",
			URL:    "http://localhost:18000/v2/students/login",
			Body: LoginReqBody{
				Email:      "ikeda2020-06-24-02@example.com",
				Password:   "Password",
				ProductID:  "002",
				CategoryID: "01",
				GroupID:    "01",
			},
			Header: map[string][]string{
				"Content-Type":     []string{"application/json"},
				"X-Api-Key":        []string{"WV5CZjAjH75R3CdVO5zHIIyWguPWKtZ1UycVGsN9"},
				"X-Transaction-ID": []string{"vegeta-lib-test"},
			},
		},
		{
			Method: "GET",
			URL:    "http://localhost:18000/v2/students/student/1000020/profile",
			Body:   nil,
			Header: map[string][]string{
				"Content-Type":     []string{"application/json"},
				"X-Api-Key":        []string{"WV5CZjAjH75R3CdVO5zHIIyWguPWKtZ1UycVGsN9"},
				"X-Transaction-ID": []string{"vegeta-lib-test"},
			},
		},
	}

	const userID = "1000020" // この例では同一ユーザのみ
	// [NOTE!] 期限付きトークン!! Need to be replaced
	userTokenMap[userID] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTI5NzgyNjUsImlhdCI6MTU5Mjg5MTg2NSwicHJvZHVjdF9pZCI6IjAwMiIsImFjdG9yX3VzZXJfY2F0ZWdvcnlfaWQiOiIwMSIsImFjdG9yX3VzZXJfZ3JvdXBfaWQiOiIwMSIsImFjdG9yX3VzZXJfaWQiOiIxMDAwMDIwIn0.UHohv1NWqiLyJ01l3qiF2e8QGzHUkAtLQubuOzFxDaU"

	var mu sync.Mutex
	var loopIdx int

	return func(tgt *vegeta.Target) error {
		// この例ではloopIdxのためにスレッドセーフにする必要がある
		mu.Lock()
		defer mu.Unlock()

		// エンドポイントを順にリクエストする(シナリオ)ための調整
		defer func() {
			if loopIdx >= len(targets)-1 {
				loopIdx = 0
			} else {
				loopIdx++
			}
		}()

		// Map to vegeta.Target
		myTarget := targets[loopIdx]
		tgt.Method = myTarget.Method
		tgt.URL = myTarget.URL

		tgt.Header = myTarget.Header
		if val, ok := userTokenMap[userID]; ok {
			tgt.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", val)}
		}

		if myTarget.Body != nil {
			if body, err := json.Marshal(myTarget.Body); err != nil {
				return err
			} else {
				tgt.Body = body
			}
		}

		return nil
	}
}
