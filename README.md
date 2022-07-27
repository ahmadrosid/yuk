# yuk - Programming Language
Programming language that compile to go.

## What the hell is this is about?
I love and hate golang. I love rust safety, but I can not use it at my work.
Golang is fine, it works it fast and enough for the day-to-day work.

But often I found myself still need to fix some stupids bug which is dealing with panic issue because of nil pointer error. 

I also hate how to stupid I'm when writing code in Go when it comes to error handling. I don't want to write repetitive code, I just want to finish my task.

The goal for this language is to prevent common bug because the unsafely of go, and just getting thing done and be productive.

## The Idea
The goal is to nil safety and be more productive. I want to write important things not the boilerplate.
```go
package yuk

import "fmt"
import "encoding/json"

type Name string
struct User(First TokenType, Last string)
struct Post (
	Title        string      `json:"input"`
	CreatedBy    User        `json:"created_by"`
	CreatedAt    Date        `json:"created_at"`
	UpdatedAt    Date
)
struct Wrong(count int)

func ReturnFunc() string {
    return "hello"
}

// this is comment
func main() {
    var name = "Ahmad Rosid"

    switch name {
        '=' => {
            var you = "rock"
            return tok
        },
        '!' => {
            var me = "lads"
            break
        },
        _ => {
            var nooo = "no"
            break
        }
    }

    var data = map(string, interface) {
        "name": 1,
        "date": "2022-07-27"
    }

    if true {
        var ok = "works"
    } else {
        var ok = "not works"
    }
}
```

**Roadmap**
- [x] Easier to create struct
- [x] Simple switch statement
  - [x] Handle multiple case statement
- [x] Shortcut map
- [ ] Anonymous struct
- [ ] String extentions `"some".len()`, `some.is_empty()`
- [ ] Array extentions `[1,2,3].len()`, `arr.is_empty()`
- [ ] Anonymous struct
- [ ] Easier to implement struct
- [ ] Mutable and immutable struct implementation
- [ ] Macro
- [ ] Typechecker
- [ ] Unsafe
