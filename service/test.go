package service

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	// 创建
	u1 := uuid.NewV4().String()
	fmt.Printf("UUIDv4: %s\n", u1)

	// 解析
	u2, err := uuid.FromString(u1)
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}
