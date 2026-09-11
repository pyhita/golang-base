package errors

import (
	"errors"
	"fmt"
	"testing"
)

// 定义测试用的自定义错误类型
type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code=%d msg=%s", e.Code, e.Msg)
}

// TestIs 测试 errors.Is 函数
func TestIs(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		target   error
		expected bool
	}{
		{
			name:     "相同错误",
			err:      errors.New("test error"),
			target:   errors.New("test error"),
			expected: false, // 两个不同的错误实例
		},
		{
			name:     "同一错误实例",
			err:      errTest,
			target:   errTest,
			expected: true,
		},
		{
			name:     "包装后相同错误",
			err:      fmt.Errorf("wrap: %w", errTest),
			target:   errTest,
			expected: true,
		},
		{
			name:     "多层包装后相同错误",
			err:      fmt.Errorf("wrap1: %w", fmt.Errorf("wrap2: %w", errTest)),
			target:   errTest,
			expected: true,
		},
		{
			name:     "不同错误",
			err:      errors.New("error A"),
			target:   errors.New("error B"),
			expected: false,
		},
		{
			name:     "nil 错误",
			err:      nil,
			target:   errTest,
			expected: false,
		},
		{
			name:     "目标为 nil",
			err:      errTest,
			target:   nil,
			expected: false,
		},
		{
			name:     "包装后不同错误",
			err:      fmt.Errorf("wrap: %w", errors.New("other")),
			target:   errTest,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Is(tt.err, tt.target)
			if result != tt.expected {
				t.Errorf("Is(%v, %v) = %v, want %v", tt.err, tt.target, result, tt.expected)
			}
		})
	}
}

// TestAs 测试 errors.As 函数
func TestAs(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
		validate func(*testing.T, error)
	}{
		{
			name:     "直接类型匹配",
			err:      &MyError{Code: 100, Msg: "test"},
			expected: true,
			validate: func(t *testing.T, e error) {
				if myErr, ok := e.(*MyError); ok {
					if myErr.Code != 100 {
						t.Errorf("expected code 100, got %d", myErr.Code)
					}
				}
			},
		},
		{
			name:     "包装后类型匹配",
			err:      fmt.Errorf("wrap: %w", &MyError{Code: 200, Msg: "wrapped"}),
			expected: true,
			validate: func(t *testing.T, e error) {
				if myErr, ok := e.(*MyError); ok {
					if myErr.Code != 200 {
						t.Errorf("expected code 200, got %d", myErr.Code)
					}
				}
			},
		},
		{
			name:     "多层包装后类型匹配",
			err:      fmt.Errorf("wrap1: %w", fmt.Errorf("wrap2: %w", &MyError{Code: 300, Msg: "multi"})),
			expected: true,
		},
		{
			name:     "类型不匹配",
			err:      errors.New("standard error"),
			expected: false,
		},
		{
			name:     "nil 错误",
			err:      nil,
			expected: false,
		},
		{
			name:     "包装后类型不匹配",
			err:      fmt.Errorf("wrap: %w", errors.New("other")),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var target *MyError
			result := As(tt.err, &target)

			if result != tt.expected {
				t.Errorf("As(%v, &target) = %v, want %v", tt.err, result, tt.expected)
				return
			}

			if tt.validate != nil && result {
				tt.validate(t, target)
			}
		})
	}
}

// TestAsWithCode 测试 errors.As 与 withCode 类型
func TestAsWithCode(t *testing.T) {
	err1 := WithCode(404, "not found")
	err2 := WrapC(500, err1, "database error")

	// 测试 As 能否找到 withCode 类型
	var target *withCode
	if !As(err2, &target) {
		t.Error("As failed to find *withCode type")
	}

	if target.code != 500 {
		t.Errorf("expected code 500, got %d", target.code)
	}
}

// TestIsWithCode 测试 errors.Is 与 withCode 类型
func TestIsWithCode(t *testing.T) {
	baseErr := errors.New("base error")
	err1 := WrapC(404, baseErr, "not found")
	err2 := Wrap(err1, "wrapped again")

	if !Is(err2, baseErr) {
		t.Error("Is failed to find base error through wrapping")
	}

	// 测试不存在的错误
	anotherErr := errors.New("another error")
	if Is(err2, anotherErr) {
		t.Error("Is incorrectly matched different error")
	}
}

// BenchmarkIs 基准测试 Is 函数
func BenchmarkIs(b *testing.B) {
	err := fmt.Errorf("wrap: %w", errTest)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Is(err, errTest)
	}
}

// BenchmarkAs 基准测试 As 函数
func BenchmarkAs(b *testing.B) {
	err := fmt.Errorf("wrap: %w", &MyError{Code: 100, Msg: "test"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var target *MyError
		As(err, &target)
	}
}

// 定义全局错误变量用于测试
var errTest = errors.New("test error")
