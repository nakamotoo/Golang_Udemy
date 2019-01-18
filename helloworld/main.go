// package = project = workspace
// package mainは特別。実行可能。他の名前のパッケージは実行できない。func main()が必ず含まれていないといけない。
package main

// stdlio的な。golang.org/pkgに詳細がある。
import "fmt"

func main() {
	fmt.Println("Hi there!")
}

// 実行方法: go run main.go
// go run: コンパイルして実行
// go build: ただコンパイルするだけ
