package main

import "os"
import "fmt"
import "time"


func main() {

	fmt.Print("請輸入日期 (YYYY-MM-DD)：")
	var targetDate string
	fmt.Scanln(&targetDate)

	target, err := time.Parse("2006-01-02", targetDate)
	if err != nil {
		fmt.Println("日期解析錯誤：", err)
		return
	}

	for {

		now := time.Now()


		diff := target.Sub(now)


		if diff <= 0 {
			fmt.Println("目標日期已到達！")
			break
		}


		days := int(diff.Hours() / 24)
		hours := int(diff.Hours()) % 24
		minutes := int(diff.Minutes()) % 60
		seconds := int(diff.Seconds()) % 60


		fmt.Fprintf(os.Stdout, "\r距離 %s 還剩下：%02d天 %02d小時 %02d分鐘 %02d秒", targetDate, days, hours, minutes, seconds)
		os.Stdout.Sync()



		time.Sleep(time.Second)
	}

	fmt.Println() 
}
