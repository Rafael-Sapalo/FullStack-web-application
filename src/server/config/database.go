package config

import "fmt"

func IsDigit(c byte) bool {
	fmt.Println("isDigit() is called")
	return '0' <= c && c <= '9'
}
