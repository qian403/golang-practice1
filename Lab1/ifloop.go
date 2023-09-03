package main 
import "fmt"

func main() {
	//coding here


	// var a int 
	// fmt.Scan(&a)
	// fmt.Println("你考了", a, "分")
	// if a >= 60 {
	// 	fmt.Println("你及格了!!!!")
	// } else {
	// 	fmt.Println("你不及格，給我重考!!!!")
	// }
		// if true{
		// 	fmt.Println("Go")
		// }else{
		// 	fmt.Println("NO GO")
		// }
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


}