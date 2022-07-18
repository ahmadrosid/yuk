# yuk - Programming Language
Programming language that compile to go.

## What the hell this is about?
I love and hate golang. I love rust safety, but I can not use it at my work.
Golang is fine, it works it fast and enough for the day-to-day work.

But often I found myself still need to fix some stupids bug which is dealing with panic issue because of nil pointer error. 

I also hate how to stupid I'm when writing code in Go when it comes to error handling. I don't want to write repetitive code, I just want to finish my task.

The goal for this language is to prevent common bug because the unsafely of go, and just getting thing done and be productive.

## The Idea
The goal is to nil safety and be more productive. I want to write important things not the boilerplate.
```go
package main

import fmt
import encoding/json

type TokenType string
struct Token(Type TypeToken, Literal string)

struct Lexer (
	input        string
	position     int
	readPosition int
	ch           byte
){
    func NewLexer(input string) *Lexer {
        self := Lexer(input)
        self.readChar()
        return self
    }

    func NexToken() Token {
        var tok Token

        switch self.ch {
            '=' =>  {

            },
            '!' => {

            }
        }
        
        self.readChar()
        return tok
    }
}

var keyword = map(string, TokenType) {
    "fn": FUNCTION
}

func main() {
    fmt.Println("Hello world!")
}
```
