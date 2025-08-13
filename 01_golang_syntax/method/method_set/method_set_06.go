package main

// 方法集合示例六：当某个结构体嵌入了接口了，它也就实现了该接口所定义的方法，这一点经常用在单元测试中

type Result struct {
	Count int
}

func (r Result) Int() int { return r.Count }

type Rows []struct{}

type Stmt interface {
	Close() error
	NumInput() int
	Exec(stmt string, args ...string) (Result, error)
	Query(args []string) (Rows, error)
}

// 返回男性员工总数
func MaleCount(s Stmt) (int, error) {
	result, err := s.Exec("select count(*) from employee_tab where gender=?", "1")
	if err != nil {
		return 0, err
	}

	return result.Int(), nil
}
