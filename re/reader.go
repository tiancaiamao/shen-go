package re

import (
	"bufio"
	"io"
	"strconv"
	"unicode"
)

type SexpReader struct {
	reader *bufio.Reader
	buf    []rune
	// 'extended' reader is used by cora, which handles reader macro ' and [,
	// and expect ; as comment
	extended bool
}

func NewSexpReader(r io.Reader, extended bool) *SexpReader {
	return &SexpReader{
		reader:   bufio.NewReader(r),
		extended: extended,
	}
}

func (r *SexpReader) Read() (Obj, error) {
	b, err := peekFirstRune(r.reader)
	if err != nil {
		return Nil, err
	}

	if r.extended {
		switch b {
		case rune(';'):
			b, _, err = r.reader.ReadRune()
			if err != nil {
				return Nil, err
			}
			for b != '\n' {
				b, _, err = r.reader.ReadRune()
				if err != nil {
					return Nil, err
				}
			}
			return r.Read()
		}
	}

	switch b {
	case rune('('):
		return r.readSexp()
	case rune('"'):
		return r.readString()
	}

	r.resetBuf()
	r.appendBuf(b)
	b, _, err = r.reader.ReadRune()
	for err == nil {
		if r.notSymbolChar(b) {
			r.reader.UnreadRune()
			break
		}
		r.appendBuf(b)
		b, _, err = r.reader.ReadRune()
	}

	return tokenToObj(string(r.buf)), err
}

func (r *SexpReader) readString() (Obj, error) {
	r.resetBuf()
	b, _, err := r.reader.ReadRune()
	for err == nil && b != rune('"') {
		r.appendBuf(b)
		b, _, err = r.reader.ReadRune()
	}
	return String(r.buf), err
}

func (r *SexpReader) readSexp() (Obj, error) {
	ret := Nil
	b, err := peekFirstRune(r.reader)
	for err == nil && b != ')' {
		var obj Obj
		r.reader.UnreadRune()
		obj, err = r.Read()
		if err == nil {
			ret = cons(obj, ret)
			b, err = peekFirstRune(r.reader)
		}
	}
	return reverse(ret), err
}

func (r *SexpReader) resetBuf() {
	r.buf = r.buf[:0]
}

func (r *SexpReader) appendBuf(b rune) {
	r.buf = append(r.buf, b)
}

func peekFirstRune(r *bufio.Reader) (rune, error) {
	b, _, err := r.ReadRune()
	for err == nil && unicode.IsSpace(b) {
		b, _, err = r.ReadRune()
	}
	return b, err
}

func (r *SexpReader) notSymbolChar(c rune) bool {
	if unicode.IsSpace(c) {
		return true
	}
	switch c {
	case '(', '"', ')':
		return true
	case '[', ']':
		return r.extended
	}
	return false
}

func tokenToObj(str string) Obj {
	switch str {
	case "true":
		return True
	case "false":
		return False
	}
	if v, err := strconv.ParseFloat(str, 64); err == nil {
		return MakeNumber(v)
	}
	return MakeSymbol(str)
}
