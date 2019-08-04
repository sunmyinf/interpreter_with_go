package lexer

type Lexer struct {
	input        string /
	position     int    // 入力における現在の位置
	readPosition int    // これから読み込む位置
	ch           byte   // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0	
	} else {
		l.ch = l.input[.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
