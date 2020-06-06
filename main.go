// 要件

// Stringerインタフェース
// String() stringメソッドを持つインタフェースを作る
// そして3つ以上Stringerインタフェースを実装する型を作る

// インタフェースを受け取る関数
// Stringerインタフェースを引数で受け取る関数を作る
// 受け取った値を上記の3つの具象型によって分岐し、具象型の型名と値を表示する
package main

import "fmt"

// ここでインタフェースを作成
type Stringer interface {
	String() string
}

// typeで
type I int
func (n I) String() string {
	return "I"
}

type B bool
func (n B) String() string {
	return "B"
}

type S string
func(n S) String() string {
	return "S"
}

func F(s Stringer) {
	switch v := s.(type) {
	case I:
		fmt.Println(int(v), "I")
	case B:
		fmt.Println(bool(v), "B")
	case S:
		fmt.Println(string(v), "S")
	}
}

func main() {
	F(I(100))
	F(B(true))
	F(S("hello"))
}
