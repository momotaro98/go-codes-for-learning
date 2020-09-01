package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sync"

	vegeta "github.com/tsenart/vegeta/lib"
)

type LoginReqBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginTarget struct {
	Method string
	URL    string
	Header http.Header
	Body   interface{}
}

var (
	usersNum   int
	loginCount int
	countMutex sync.Mutex
)

// CSVファイルから全アカウントのクレデンシャル情報をロード
func LoadLoginUsers(filePath string) (targets []*LoginTarget, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string

	for {
		line, err = reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		targets = append(targets, &LoginTarget{
			Method: http.MethodPost,
			URL:    "http://localhost:8090/login",
			Header: map[string][]string{
				"Content-Type": {"application/json"},
			},
			Body: LoginReqBody{
				Email:    line[0],
				Password: line[1],
			},
		})
	}

	return targets, nil
}

// ログインリクエスト
func NewLoginTargeter(userFilePath string) vegeta.Targeter {
	targets, err := LoadLoginUsers(userFilePath)
	if err != nil {
		panic(err)
	}

	usersNum = len(targets)

	var (
		index int
		mu    sync.Mutex
	)

	return func(tgt *vegeta.Target) error {
		mu.Lock()
		defer mu.Unlock()
		defer func() {
			index++
		}()

		if index >= len(targets) {
			return vegeta.ErrNoTargets
		}
		target := targets[index]
		tgt.Method = target.Method
		tgt.URL = target.URL
		tgt.Header = target.Header
		if target.Body != nil {
			if body, err := json.Marshal(target.Body); err != nil {
				return err
			} else {
				tgt.Body = body
			}
		}

		return nil
	}
}

func PassToken(resBody []byte, tokenChn chan<- string) {
	if len(resBody) == 0 {
		return
	}

	type LoginRes struct {
		JwtToken string `json:"jwt_token"`
	}
	var loginRes LoginRes
	json.Unmarshal(resBody, &loginRes)
	tokenChn <- loginRes.JwtToken

	// 全アカウント分終了時にトークン用のチャネルを閉じるための制御
	countMutex.Lock()
	defer countMutex.Unlock()
	loginCount++
	if loginCount == usersNum {
		close(tokenChn)
	}
}
