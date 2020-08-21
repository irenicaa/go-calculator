package calculator

import (
	"github.com/irenicaa/go-calculator/evaluator"
	"github.com/irenicaa/go-calculator/tokenizer"
	"github.com/irenicaa/go-calculator/translator"
)

// Calculator ...
type Calculator struct {
	tokenizer  tokenizer.Tokenizer
	translator translator.Translator
	evaluator  evaluator.Evaluator
}
