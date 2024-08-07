package main

import (
	"fmt"
	"math/rand"
)

func UserInput() (guess int, yes bool) {
	yes = true
	fmt.Println("请输入你猜的数字")
	// Scanf 简化代码实现
	// 但如果用户输入的非int 会 输出多次异常提示
	// 用Scan 不读入换行符
	_, err := fmt.Scan(&guess)
	if err != nil {
		fmt.Println("非法输入，你输入的不是一个int")
		yes = false
	}
	return guess, yes
}
func main() {
	fmt.Println("你好呀，开始我们的猜数字游戏")
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println(secretNumber)
	for {
		userInput, yes := UserInput()
		if yes != true {
			continue
		}
		if userInput == secretNumber {
			fmt.Println("恭喜你猜对了，游戏结束")
			break
		} else if userInput > secretNumber {
			fmt.Println("太大了，重新猜")
		} else if userInput < secretNumber {
			fmt.Println("太小了重新猜测")
		}
	}
}
