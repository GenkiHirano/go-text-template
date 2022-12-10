package main

import (
	"fmt"
	"go-text-template/sample"
)

func main() {
	fmt.Println("<< 通常パターン >>")
	sample.Sample1()
	fmt.Println("<< 構造体を使ったパターン >>")
	sample.Sample2()
	fmt.Println("<< 雛形ファイルと連携したパターン >>")
	sample.Sample3()
	fmt.Println("<< 雛形ファイル内での制御構文パターン >>")
	sample.Sample4()
	fmt.Println("<< 自作関数を使ったパターン >>")
	sample.Sample5()
}
