package test

import "testing"

// 要测试的函数：两数相加
func Add(a, b int) int {
	return a + b
}

// 测试函数名：Test + 被测试函数名（大驼峰）
func TestAdd(t *testing.T) {
	// 调用函数
	res := Add(1, 2)

	// 断言结果
	if res != 3 {
		t.Errorf("Add(1,2) 期望: 3, 实际: %d", res)
	}
}
