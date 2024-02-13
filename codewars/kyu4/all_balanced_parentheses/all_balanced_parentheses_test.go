package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5426d7a2c2c7784365000783

type ParenthesisGenerator struct {
	result  []string
	current string
	n       int
}

func (p *ParenthesisGenerator) generateParenthesisHelper(open, close int) {
	if len(p.current) == 2*p.n {
		p.result = append(p.result, p.current)
		return
	}

	if open < p.n {
		p.current += "("
		p.generateParenthesisHelper(open+1, close)
		p.current = p.current[:len(p.current)-1]
	}

	if close < open {
		p.current += ")"
		p.generateParenthesisHelper(open, close+1)
		p.current = p.current[:len(p.current)-1]
	}
}

func BalancedParens(n int) []string {
	p := ParenthesisGenerator{n: n}
	p.generateParenthesisHelper(0, 0)
	return p.result
}

func TestBalancedParens(t *testing.T) {
	assert.Equal(t, []string{""}, BalancedParens(0))
	assert.Equal(t, []string{"()"}, BalancedParens(1))
	assert.Equal(t, []string{"(())", "()()"}, BalancedParens(2))
	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, BalancedParens(3))
	assert.Equal(t, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}, BalancedParens(4))
}
