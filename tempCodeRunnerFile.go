	target, err := time.Parse("2006-01-02", targetDate)
	if err != nil {
		fmt.Println("日期解析錯誤：", err)
		return
	}