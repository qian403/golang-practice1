package main
import "fmt"
func main(){


	fmt.Print("Hello World") //不換行output
	fmt.Println("Hello World1") 
	fmt.Println("Hello World2") //換行output

	var a int  //宣告變數
	fmt.Scan(&a) //輸入變數a
	fmt.Println(a) //輸出變數a
	fmt.Println(&a) //輸出變數a的記憶體位置	
	var name1 string = "John" //宣告變數並賦值
	var int1 int = 11 
	var float1 float32 = 1.2
	var bool1 bool = true
	fmt.Println(name1, int1, float1, bool1)
}