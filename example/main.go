package main

import (
	"time"
)

type APIKey struct {
	ID        string
	UpdatedAt time.Time
}

// mq 批处理命令
func getAPIKeys() (APIKey, error) {
	var err error
	return APIKey{}, err
}

// 批量写入数据库
func saveToDB(apiKeys *[]APIKey) error {
	var err error
	//err=db.save(apiKeys)
	return err
}

func main() {
	//todo 100批处理之后执行一次批量插入
	//todo 每30s执行一次批量插入
	akChan := make(chan []APIKey)
	ticker := time.NewTicker(30 * time.Second)

	go func() {
		for range ticker.C {
			apiKeys, err := GetApiKeys()
			if err != nil {
				//错误处理,并通知异步回滚
			}
			akChan <- apiKeys
		}
	}()
	for {
		select {
		case apiKeys := <-akChan:
			err := saveToDB(&apiKeys)
			//事务回滚
			if err != nil {
				//回滚
			}
		}
	}
}

func GetApiKeys() ([]APIKey, error) {
	aks := make([]APIKey, 100)
	for i := 0; i < 100; i++ {
		key, err := getAPIKeys()
		if err != nil {
			return nil, err
		}
		aks = append(aks, key)
	}
	return aks, nil
}
