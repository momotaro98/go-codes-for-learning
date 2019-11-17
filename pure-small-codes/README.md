# Shift operation
https://play.golang.org/p/hrEQNM7546Z

# 制御構造

## For statement

```go
package main
import "fmt"
// aを逆に並び替える
// ① i++, j-- というのは"文"なので複数変数をfor文で扱うときはこのようにする
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 { // ①
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a) // [5 4 3 2 1]
}
```

## Switch statement

```go
package main
import "fmt"
func retInterface() interface{} {
	v := 100
	return &v
}
func main() {
	interfaceValue := retInterface()
	switch t := interfaceValue.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T は型を出力する
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // pointer to integer 100
	}
}
```
