package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF
	TokenIndent
	TokenDash     // -
	TokenColon    // :
	TokenComma    // ,
	TokenLBracket // [
	TokenRBracket // ]
	TokenLBrace   // {
	TokenRBrace   // }
	TokenString
	TokenNumber
	TokenBool
	TokenNull
	TokenComment
	TokenAnchor    // &
	TokenAlias     // *
	TokenTag       // !!
	TokenDirective // %YAML, %TAG
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
	Col   int
}

func (tt TokenType) String() string {
	switch tt {
	case TokenError:
		return "Error"
	case TokenEOF:
		return "EOF"
	case TokenIndent:
		return "Indent"
	case TokenDash:
		return "Dash"
	case TokenColon:
		return "Colon"
	case TokenComma:
		return "Comma"
	case TokenLBracket:
		return "LBracket"
	case TokenRBracket:
		return "RBracket"
	case TokenLBrace:
		return "LBrace"
	case TokenRBrace:
		return "RBrace"
	case TokenString:
		return "String"
	case TokenNumber:
		return "Number"
	case TokenBool:
		return "Bool"
	case TokenNull:
		return "Null"
	case TokenComment:
		return "Comment"
	case TokenAnchor:
		return "Anchor"
	case TokenAlias:
		return "Alias"
	case TokenTag:
		return "Tag"
	case TokenDirective:
		return "Directive"
	default:
		return "Unknown"
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%s (%d:%d): %q", t.Type.String(), t.Line, t.Col, t.Value)
}

type parser struct {
	reader  *bufio.Reader
	pos     int
	line    int
	col     int
	lastCol int
}

func NewParser(r io.Reader) *parser {
	return &parser{
		reader: bufio.NewReader(r),
		line:   1,
		col:    0,
	}
}

// readRune reads next rune
func (l *parser) readRune() (rune, error) {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		return 0, err
	}

	l.pos++
	l.col++
	if r == '\n' {
		l.line++
		l.lastCol = l.col - 1
		l.col = 0
	}

	return r, nil
}

// unreadRune rollbacks to previous rune
func (l *parser) unreadRune() error {
	err := l.reader.UnreadRune()
	if err != nil {
		return err
	}

	l.pos--
	if l.col == 0 {
		l.line--
		l.col = l.lastCol
	} else {
		l.col--
	}

	return nil
}

// skipWhitespace skips whitespace runes
func (l *parser) skipWhitespace() error {
	for {
		r, err := l.readRune()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if !unicode.IsSpace(r) {
			return l.unreadRune()
		}
	}
}

// NextToken returns next token
func (l *parser) NextToken() (Token, error) {
	err := l.skipWhitespace()
	if err != nil {
		return Token{TokenError, err.Error(), l.line, l.col}, err
	}

	r, err := l.readRune()
	if err != nil {
		if err == io.EOF {
			return Token{TokenEOF, "", l.line, l.col}, nil
		}
		return Token{TokenError, err.Error(), l.line, l.col}, err
	}

	switch r {
	case '-':
		// check begin of multiline literal
		nextR, err := l.readRune()
		if err != nil && err != io.EOF {
			return Token{TokenError, err.Error(), l.line, l.col}, err
		}
		if err == nil && nextR == ' ' {
			return Token{TokenDash, "-", l.line, l.col - 1}, nil
		}
		if err == nil {
			_ = l.unreadRune()
		}
		return Token{TokenDash, "-", l.line, l.col - 1}, nil
	case ':':
		return Token{TokenColon, ":", l.line, l.col - 1}, nil
	case ',':
		return Token{TokenComma, ",", l.line, l.col - 1}, nil
	case '[':
		return Token{TokenLBracket, "[", l.line, l.col - 1}, nil
	case ']':
		return Token{TokenRBracket, "]", l.line, l.col - 1}, nil
	case '{':
		return Token{TokenLBrace, "{", l.line, l.col - 1}, nil
	case '}':
		return Token{TokenRBrace, "}", l.line, l.col - 1}, nil
	case '#':
		return l.readComment()
	case '&':
		return l.readAnchor()
	case '*':
		return l.readAlias()
	case '!':
		return l.readTag()
	case '%':
		return l.readDirective()
	case '"', '\'':
		return l.readString(r)
	default:
		// check begin of number
		if unicode.IsDigit(r) || r == '-' || r == '+' || r == '.' {
			_ = l.unreadRune()
			return l.readNumber()
		}

		// check keywords (true, false, null)
		if unicode.IsLetter(r) {
			_ = l.unreadRune()
			return l.readKeyword()
		}

		return Token{TokenError, fmt.Sprintf("unexpected character: %c", r), l.line, l.col}, nil
	}
}

func (l *parser) readComment() (Token, error) {
	var sb strings.Builder
	sb.WriteRune('#')

	for {
		r, err := l.readRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return Token{TokenError, err.Error(), l.line, l.col}, err
		}

		if r == '\n' {
			_ = l.unreadRune()
			break
		}

		sb.WriteRune(r)
	}

	return Token{TokenComment, sb.String(), l.line, l.col - sb.Len()}, nil
}

func (l *parser) readAnchor() (Token, error) {
	var sb strings.Builder
	sb.WriteRune('&')

	for {
		r, err := l.readRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return Token{TokenError, err.Error(), l.line, l.col}, err
		}

		if !isAnchorChar(r) {
			_ = l.unreadRune()
			break
		}

		sb.WriteRune(r)
	}

	return Token{TokenAnchor, sb.String(), l.line, l.col - sb.Len()}, nil
}

// readAlias reads an alias (*alias)
func (l *parser) readAlias() (Token, error) {
	var sb strings.Builder
	sb.WriteRune('*')

	for {
		r, err := l.readRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return Token{TokenError, err.Error(), l.line, l.col}, err
		}

		if !isAnchorChar(r) {
			_ = l.unreadRune()
			break
		}

		sb.WriteRune(r)
	}

	return Token{TokenAlias, sb.String(), l.line, l.col - sb.Len()}, nil
}

// readTag reads a tag (!tag)
func (l *parser) readTag() (Token, error) {
	var sb strings.Builder
	sb.WriteRune('!')

	// check double ! (!!str)
	nextR, err := l.readRune()
	if err != nil && err != io.EOF {
		return Token{TokenError, err.Error(), l.line, l.col}, err
	}
	if err == nil && nextR == '!' {
		sb.WriteRune('!')
	} else if err == nil {
		_ = l.unreadRune()
	}

	for {
		r, err := l.readRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return Token{TokenError, err.Error(), l.line, l.col}, err
		}

		if !isTagChar(r) {
			_ = l.unreadRune()
			break
		}

		sb.WriteRune(r)
	}

	return Token{TokenTag, sb.String(), l.line, l.col - sb.Len()}, nil
}

// readDirective reds directives (%YAML, %TAG)
func (l *parser) readDirective() (Token, error) {
	var sb strings.Builder
	sb.WriteRune('%')

	for {
		r, err := l.readRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return Token{TokenError, err.Error(), l.line, l.col}, err
		}

		if !isDirectiveChar(r) {
			_ = l.unreadRune()
			break
		}

		sb.WriteRune(r)
	}

	return Token{TokenDirective, sb.String(), l.line, l.col - sb.Len()}, nil
}

// readString reads quoted string
func (l *parser) readString(quote rune) (Token, error) {
	var sb strings.Builder
	sb.WriteRune(quote)
	escape := false

	for {
		r, err := l.readRune()
		if err != nil {
			if err == io.EOF {
				return Token{TokenError, "unclosed string", l.line, l.col}, nil
			}
			return Token{TokenError, err.Error(), l.line, l.col}, err
		}

		sb.WriteRune(r)

		if r == '\\' && !escape {
			escape = true
			continue
		}

		if r == quote && !escape {
			break
		}

		escape = false
	}

	return Token{TokenString, sb.String(), l.line, l.col - sb.Len()}, nil
}

func (l *parser) readNumber() (Token, error) {
	var sb strings.Builder
	hasDot := false
	hasExp := false

	for {
		r, err := l.readRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return Token{TokenError, err.Error(), l.line, l.col}, err
		}

		if unicode.IsDigit(r) {
			sb.WriteRune(r)
			continue
		}

		switch r {
		case '.':
			if hasDot || hasExp {
				_ = l.unreadRune()
				return Token{TokenNumber, sb.String(), l.line, l.col - sb.Len()}, nil
			}
			hasDot = true
			sb.WriteRune(r)
		case 'e', 'E':
			if hasExp {
				_ = l.unreadRune()
				return Token{TokenNumber, sb.String(), l.line, l.col - sb.Len()}, nil
			}
			hasExp = true
			sb.WriteRune(r)

			// check +/-
			nextR, err := l.readRune()
			if err != nil && err != io.EOF {
				return Token{TokenError, err.Error(), l.line, l.col}, err
			}
			if err == nil && (nextR == '+' || nextR == '-') {
				sb.WriteRune(nextR)
			} else if err == nil {
				_ = l.unreadRune()
			}
		case '-', '+':
			// check signs before or after e/E
			if sb.Len() > 0 && (sb.String()[sb.Len()-1] != 'e' && sb.String()[sb.Len()-1] != 'E') {
				_ = l.unreadRune()
				return Token{TokenNumber, sb.String(), l.line, l.col - sb.Len()}, nil
			}
			sb.WriteRune(r)
		default:
			_ = l.unreadRune()
			return Token{TokenNumber, sb.String(), l.line, l.col - sb.Len()}, nil
		}
	}

	return Token{TokenNumber, sb.String(), l.line, l.col - sb.Len()}, nil
}

// readKeyword reads a keyword (true, false, null)
func (l *parser) readKeyword() (Token, error) {
	var sb strings.Builder

	for {
		r, err := l.readRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return Token{TokenError, err.Error(), l.line, l.col}, err
		}

		if !unicode.IsLetter(r) {
			_ = l.unreadRune()
			break
		}

		sb.WriteRune(r)
	}

	value := sb.String()
	switch value {
	case "true", "false":
		return Token{TokenBool, value, l.line, l.col - len(value)}, nil
	case "null", "~", "Null", "NULL":
		return Token{TokenNull, value, l.line, l.col - len(value)}, nil
	default:
		return Token{TokenString, value, l.line, l.col - len(value)}, nil
	}
}

// isAnchorChar checks rune allow in anchor/alias
func isAnchorChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-'
}

// isTagChar checks rune allow in tag
func isTagChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' || r == '/' || r == '#' || r == ';' || r == ':' || r == '?' || r == '@' || r == '&' || r == '=' || r == '+' || r == '$' || r == ',' || r == '~' || r == '%' || r == '*' || r == '[' || r == ']'
}

// isDirectiveChar checks rune allow in directive
func isDirectiveChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' || r == '/' || r == '#' || r == ';' || r == ':' || r == '?' || r == '@' || r == '&' || r == '=' || r == '+' || r == '$' || r == ',' || r == '~' || r == '%' || r == '*' || r == '[' || r == ']'
}

func main() {
	input := `---
# YAML document example
person:
  name: "John Doe"
  age: 30
  married: true
  children: null
  hobbies:
    - hiking
    - reading
  address: &addr
    street: 123 Main St
    city: Anytown
  other_address: *addr
  !!str custom: "value"
`

	lexer := NewParser(strings.NewReader(input))

	for {
		token, err := lexer.NextToken()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fmt.Println(token)

		if token.Type == TokenEOF {
			break
		}
	}
}
