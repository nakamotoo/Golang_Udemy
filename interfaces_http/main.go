package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 自前でWriter Interfaceをもつ型を実装してみる
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*
		// 99999要素を格納できるbyte sliceを作っておく
		bs := make([]byte, 99999)
		// ReaderインターフェースのRead()を実行
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/
	// 上のやり方は骨が折れるので、もっと簡単に同じことを実現したい (下の一行で実現 !!!)
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// Writer Interfaceは、Write()メソッドをもつ型に適応されるので、logWriter用に実装
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
