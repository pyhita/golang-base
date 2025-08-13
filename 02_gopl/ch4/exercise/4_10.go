package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pyhita/golang-base/02_gopl/ch4/github"
)

type IssueType string

const (
	OneMonth IssueType = "<= 1 month"
	OneYear  IssueType = "<= 1 year"
	Others   IssueType = "> 1 year"
)

func makeType(createAt time.Time) IssueType {
	if isOneYear(createAt) {
		return OneYear
	}
	if isOneMonth(createAt) {
		return OneMonth
	}

	return Others
}

func isOneMonth(createAt time.Time) bool {
	now := time.Now()
	oneMonthLater := createAt.AddDate(0, 1, 0)

	return now.Before(oneMonthLater) || now.Equal(oneMonthLater)
}

func isOneYear(createAt time.Time) bool {
	now := time.Now()
	oneYearLater := createAt.AddDate(1, 0, 0)
	return now.Before(oneYearLater) || now.Equal(oneYearLater)
}

func isOthers(createAt time.Time) bool {
	return !isOneYear(createAt) && !isOneMonth(createAt)
}

// 练习 4.10： 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var m map[IssueType][]*github.Issue

	// 分类
	for _, item := range result.Items {
		t := makeType(item.CreatedAt)
		m[t] = append(m[t], item)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for t, times := range m {
		println("issue type ", t)
		for _, item := range times {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
