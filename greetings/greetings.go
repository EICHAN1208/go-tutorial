package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you , %v!",
		"Hail, %v! Well met!",
	}

	// 現在の時刻を使って疑似乱数のシードを設定する
	// これを行わないと、毎回同じシーケンスの疑似乱数が生成されるため、
	// 本当のランダムな動きを得ることができない
	rand.Seed(time.Now().UnixNano())

	return formats[rand.Intn(len(formats))]
}
