package main 
import "fmt"

func main() {
	//coding here
	// fmt.Print("Hello World") //不換行output
	// fmt.Println("Hello World") //換行output
	// fmt.Printf("Hello World") 
	// var a int  //宣告變數
	// fmt.Scan(&a) //輸入變數a
	// fmt.Println(a) //輸出變數a
	// fmt.Println(&a) //輸出變數a的記憶體位置	
	// var name1 string = "John" //宣告變數並賦值
	// var int1 int = 11 
	// var float1 float32 = 1.2
	// var bool1 bool = true
	// fmt.Println(name1, int1, float1, bool1)

	var a int 
	fmt.Scan(&a)
	fmt.Println("你考了", a, "分")
	if a >= 60 {
		fmt.Println("你及格了!!!!")
	} else {
		fmt.Println("你不及格，給我重考!!!!")
	}

	
}