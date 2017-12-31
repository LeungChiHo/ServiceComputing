package main

import (
	"fmt"
	"time"
)

func asynchronous() [10]Info {
	start := time.Now()
	customer := GetCustomerDetails()
	destinations := GetDestinations(customer)
	var infos [10]Info

	quotes := [10]chan Quoting{}
	weathers := [10]chan Weather{}

	//创建 channel Quoting
	for i, _ := range quotes {
		quotes[i] = make(chan Quoting)
	}

	//创建 channel Weather
	for i, _ := range weathers {
		weathers[i] = make(chan Weather)
	}

	for index, dest := range destinations {
		i := index
		d := dest
		go func() {
			// 从 channel GetQuote(d) 接收 Quoting，并保存到 quotes[i] 中
			quotes[i] <- GetQuote(d)
		}()

		go func() {
			// 从 channel GetWeather(d) 接收 Weather，并保存到 weathers[i] 中
			weathers[i] <- GetWeather(d)
		}()
	}

	//等待，直到从 channel 上接收Quote和Weather。
	for index, dest := range destinations {
		infos[index] = Info{Destinations:dest, Quote:<-quotes[index], Weather:<-weathers[index]}
	}

	fmt.Println(time.Since(start))
	return infos
}

func synchronous() [10]Info {
	start := time.Now()
	customer := GetCustomerDetails()
	destinations := GetDestinations(customer)
	var infos [10]Info
	for index, dest := range destinations {
		q := GetQuote(dest)
		w := GetWeather(dest)
		infos[index] = Info{Destinations:dest, Quote:q, Weather:w}
	}	
	fmt.Println(time.Since(start))
	return infos
}


func main() {
	var choice int
	for {
		fmt.Print("请输入模式编号(0异步,1同步,2退出)：")
		fmt.Scanln(&choice)
		if choice == 0 {
			asynchronous()
		} else if choice == 1 {
			synchronous()
		} else if choice == 2 {
			break
		}
	}
}
