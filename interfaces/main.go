package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// printGreetingは完全にアルゴリズムが一緒だけど、引数の型が違うから別々にしないといけないのか？ => interfacesで解決しよう!
/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
*/

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	// spanishとは全然違うアルゴリズムで挨拶を生成する とする
	return "Hi There!"
}

// sb(eb)を使ってないなら、省略化
func (spanishBot) getGreeting() string {
	return "Hola!"
}
