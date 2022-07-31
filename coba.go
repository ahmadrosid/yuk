package yuk

import "fmt"
import "encoding/json"

type Name string
type User struct {
	First TokenType
	Last  string
}
type Post struct {
	Title     string `json:"input"`
	CreatedBy User   `json:"created_by"`
	CreatedAt Date   `json:"created_at"`
	UpdatedAt Date
}
type Wrong struct {
	count int
}

func ReturnFunc() string {
	return "hello"
}
func main() {
	var name = "Ahmad Rosid"

	lock := true

	switch name {
	case '=':
		{
			var you = "rock"

			return tok
		}
	case '!':
		{
			var me = "lads"

			break
		}
	default:
		{
			var nooo = "no"

			break
		}
	}

	var data = map[string]interface{}{
		"name": 1,
		"date": "2022-07-27",
	}

	if true {
		var ok = "works"
	} else {
		var ok = "not works"
	}
}
