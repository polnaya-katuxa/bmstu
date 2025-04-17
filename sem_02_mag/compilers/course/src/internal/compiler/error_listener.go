package compiler

import "github.com/antlr4-go/antlr/v4"

type customErrorListener struct {
	*antlr.DefaultErrorListener
	IsError bool
}

func (c *customErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.IsError = true
}
