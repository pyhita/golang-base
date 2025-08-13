package main

import "math"

// 练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。

func max(values ...int) int {
	t := math.MinInt64
	if len(values) == 0 {
		return t
	}

	for _, value := range values {
		if value > t {
			t = value
		}
	}
	return t
}

func min(values ...int) int {
	t := math.MaxInt64
	if len(values) == 0 {
		return t
	}

	for _, value := range values {
		if value < t {
			t = value
		}
	}
	return t
}
