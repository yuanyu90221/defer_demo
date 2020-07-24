# golang defer 特性(延遲處理)

## introduction

defer 是 golang裡面一個 preserve word

是一個語法糖 

將defer 之後的func 延遲執行 直到所在 func content執行結束(return or panic) 才執行

多個 defer 同時存在時 執行順序有 LIFO(Last In First Out)特性

## example
```golang
package main 

func deferFun() {
	fmt.Println("deferFunc Before return")
}
func deferFun1() {
	fmt.Println("deferFunc1 Before return")
}
func main() {
	defer deferFun()
	defer deferFun1()
	panic(fmt.Errorf("error out"))
}

```
```golang
// output 順序是
deferFunc1 Before return
deferFunc Before return
Panic: error out
```

## 需要注意的點
如果defer的func 有帶入參數則當時代入的值也會被傳入 
因此在invoke時 所使用的值 就是傳入時的值
```golang
import "fmt"

func main() {
	i := 10
	defer fmt.Println(i)
	i++
	fmt.Println(i)
}

```
https://golang.org/doc/effective_go.html#defer

https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/

## Advantage of use defer?

1 如果把釋放資源的邏輯放在defer之後,可以確保在離開這段func之前一定會執行到

2 defer寫法可以讓 釋放resource的邏輯 寫起來比較簡便

需要注意的是

當defer 後面function有待入參數 該參數 會維持當初帶入時的值

因此如果是多個defer的使用 需要注意順序索引發的正確性問題
## when to use?

1 DB connection 的 close

2 file stream 的close

## test

```golang
package main
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```

```golang
// output
// entering: b
// in b
// entering: a
// in a
// leaving: a
// leaving: b
```

## panic 

panic是 golang中用來直接拋錯的機制

## recover

recover是對應於panic 用來接下 panic的機制

可以把panic 接下成error

重點是 recover一定要在 defer的func 內才能使用

[golang defer特性](https://hackmd.io/7F5_kLPqRoObMbQhsnt61w?view)