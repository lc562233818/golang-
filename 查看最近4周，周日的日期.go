package main

import (
	"fmt"
	"time"
)

func keyDay() {
	now := time.Now()
	sub := 7 - int(now.Weekday()) //周一
	fmt.Println(sub)
	interval := sub
	if sub == 0 { //是否为周日
		interval = 7
	}
	//firstEr := now.Add(24 * time.Duration(interval) * time.Hour)
	//fmt.Println(firstEr.Format("2006-01-02"))
	for i := 0; i < 4; i++ {
		firstEr := now.Add(24 * time.Duration(interval*i) * time.Hour)
		fmt.Println(firstEr.Format("2006-01-02"))
	}
}
func main() {
	keyDay()
}
