# Golang學習筆記
程式存放於[Github存儲庫](https://github.com/qian403/golang-practice1)
同步於[hackmd]()
## 簡介


## GO!
Go執行程式的方式有兩種，一種是直接執行，另一種是編譯後執行
直接執行，main.go要換成檔案名稱
```
go run main.go
```
編譯後執行，會產生一個可執行檔，可以直接執行
```
go build main.go
```


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
```
可以把下面的```if true```改成```if false```試試看，觀察看看會發生什麼事
```go
package main
import "fmt"
func main(){
    		if true{
			fmt.Println("Go")
		}else{
			fmt.Println("NO GO")
		}
}
```
#### 練習
##### ATM機器
要寫一個小程式，去限制使用者提領的金額上限，如果超過上限就輸出```提領失敗```，如果沒有超過上限就輸出```提領成功```

```go
package main
import "fmt"
func main(){
		var money int //宣告變數
		fmt.Print("輸入要提領的金額:") 
		fmt.Scan(&money) //輸入變數money
		if money <= 10000 { //判斷式，如果money小於等於10000則輸出提領成功，否則輸出提領失敗
			fmt.Println("提領成功")
		}else{
				fmt.Println("提領失敗")
		}
}
```
附加條件:最低提領金額為1000元，如果低於1000元則輸出```提領失敗```
解答:
```go
package main
import "fmt"
func main(){
    		var money int
		fmt.Print("輸入要提領的金額:")
		fmt.Scan(&money)
		if money <100{
			fmt.Println("提領金額太少")
		}else if money >10000{
			fmt.Println("提領金額太多")
		}else{
			fmt.Println("提領成功")
		}
        fmt.Println("程式結束") //程式結束觀察if迴圈的執行順序
}
```


### for迴圈
![Image](https://i.imgur.com/165juLv.png)
![Image](https://i.imgur.com/0G43hh4.png)
可以先試試看，下面的程式會無限輸出```Go```，要停止的話可以按```ctrl+c```
```go
package main
import "fmt"
func main(){
    for true{
        fmt.Println("Go")
    }
}
```


```go
package main
import "fmt"
func main() {
	fmt.Println("作法一")
	var i int = 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
////////////////////////////////////////	
	fmt.Println("作法二")
	var x int
	for x=1;x<=10;x++ {
		fmt.Println(x)
		
	}

}
```
#### 練習
##### 1~50的總和
```go
package main
import "fmt"
func main(){
	var x int
	var result int = 0
	for x<=50 {
		result += x
		x++

	}
	fmt.Println(result)
}
```
##### 九九乘法表
```go
package main
import "fmt"
func main(){
    var i int
    var j int
    for i=1;i<=9;i++{
        for j=1;j<=9;j++{
            fmt.Printf("%d*%d=%d\t",i,j,i*j)
        }
        fmt.Println()
    }
}
```
##### 1~100的奇數和
```~~go~~
package main
import "fmt"
func main(){
    var i int
    var result int = 0
    for i=1;i<=100;i++{
        if i%2==1{
            result += i
        }
    }
    fmt.Println(result)
}
```
#### break與continue


