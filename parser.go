%{
// プログラムのヘッダを指定
package main
import (
  "os"
  "fmt"
)
%}

%union {
// 自作言語で使うトークンの宣言
  num int
  str string
}

// プログラムの構成要素を指定
%type<num> program expr
%type<str> program show
%token<num> NUMBER
%token<str> STRING

// 演算の優先度の指定
%left '+','-'
%left '*','/'

%%
// 文法規則を指定
// 文法要素
//   : 定義1 {Go言語のプログラム1}
//   | 定義2 {Go言語のプログラム2}
program // program要素の定義
  // expr 入力を字句解析
  // $$はメタ変数。文法を解析した結果を$$に格納。$$には%unionのメンバのどれかが入る。
  : expr
  {
    $$ = $1 // 引数の1つ目を取得
    yylex.(*Lexer).result = $$ // 計算結果をmain関数に渡すため、Lexer構造体に用意したresultに値を格納。
  }

expr // expr要素の定義
  : NUMBER
  | expr '+' expr { $$ = $1 + $3 }
  | expr '-' expr { $$ = $1 - $3 }
  | expr '*' expr { $$ = $1 * $3 }
  | expr '/' expr { $$ = $1 / $3 }
  : STRING
  | '(' expr ')' { $$ = $2 }
%%

// 最低限必要な構造体を定義
type Lexer struct {
  src    string
  index  int
  result int
}

// Lex 字句解析
// ここでトークン（最小限の要素）を一つずつ返す
func (p *Lexer) Lex(lval *yySymType) int {
  for p.index < len(p.src) {
    c := p.src[p.index]
    p.index++
    if c == '+' { return int(c) }
    if c == '-' { return int(c) }
    if c == '*' { return int(c) }
    if c == '/' { return int(c) }
    if '0' <= c && c <= '9' {
      lval.num = int(c - '0') // ?
      return NUMBER
    }
    if c == 'm' && p.src[p.index+1] == 'u' {
      continue
    }
    if c == 'u' && p.src[p.index-1] == 'm' {
      return ''
    }
  }
  return -1
}

// エラー報告用
func (p *Lexer) Error(e string) {
  fmt.Println("[error] " + e)
}

// メイン関数
func main() {
  if len(os.Args) <= 1 { return }
  lexer := &Lexer{src: os.Args[1], index:0}
  // goyaccでは%%で囲んだ規則部分をもとに字句解析器yyParse関数の中身を定義してくれる。
  // yyParseの中でLex, Error関数を呼んでくれる。
  // TODO: 後でyyParseの中身をみて、外部で定義した関数の読み込み方法を学ぶ
  yyParse(lexer)
  println("計算結果:", lexer.result)
}
