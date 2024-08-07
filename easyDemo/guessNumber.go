package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func UserInput() (guess int, yes bool) {
	yes = true
	fmt.Println("请输入你猜的数字")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读输入 发生错误")
		yes = false
	}
	input = strings.Replace(input, "\n", "", -1)
	// input = strings.TrimSuffix(input,"\n")
	guess, err = strconv.Atoi(input)
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
