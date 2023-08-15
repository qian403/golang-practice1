# Golang學習筆記
程式存放於[Github存儲庫](https://github.com/qian403/golang-practice1)
同步於[hackmd]()
## 簡介


## GO!

## 基本語法&輸入與輸出
這是程式一開始的時候
```go
package main //宣告main套件
import "fmt" //宣告fmt套件

func main() {
	//coding here

}
```

### 輸入與輸出
```go
package main
import "fmt"
func main() {

    fmt.Print("Hello World") //不換行output
	fmt.Println("Hello World") //換行output
	fmt.Printf("Hello World") 
}
```
要注意的是在go中和c++不同的是要在給變數賦值前加上&才能輸入
```go
package main 
import "fmt"

func main() {
	//coding here
	var a int  //宣告變數
	fmt.Scan(&a) //輸入變數a
	fmt.Println(a) //輸出變數a
	fmt.Println(&a) //輸出變數a的記憶體位置
}
```
### 變數

```go
package main
import "fmt"
func main(){
    var name1 string = "John" //宣告變數並賦值
	var int1 int = 11 
	var float1 float32 = 1.2
	var bool1 bool = true
	fmt.Println(name1, int1, float1, bool1)
}
```
### if判斷式

```go 
package main
import "fmt"
func main(){
	var aa int = 11
	if aa > 10 {
		fmt.Println("aa > 10")
	}
}
```
```go
package main
import "fmt"
func main(){
	var aa int = 11 //宣告變數並賦值
	if aa > 10 { //判斷式
		fmt.Println("aa > 10") //aa大於10則輸出
	}else if aa > 5 {
		fmt.Println("aa > 5") //aa大於5則輸出
	}else {
		fmt.Println("aa < 5") //aa小於5則輸出
	}
	
}