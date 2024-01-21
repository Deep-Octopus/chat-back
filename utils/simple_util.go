package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateAccountNumber 生成指定位数的
func GenerateAccountNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	var accountNumber string
	for i := 0; i < length; i++ {
		digit := rand.Intn(10)
		accountNumber += fmt.Sprintf("%d", digit)
	}
	return accountNumber
}
