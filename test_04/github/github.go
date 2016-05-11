package github

import (
	"time"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
)

type IssusesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time
	Body      string
}

const IssuesURL = "https://api.github.com/search/issues"

func SearchIssues(terms []string) (*IssusesSearchResult, error) {
	q := strings.Join(terms, "&")
	fmt.Println(IssuesURL + "?q=" + q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed:%s", resp.StatusCode)
	}
	var result IssusesSearchResult
	//fmt.Println(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

}
