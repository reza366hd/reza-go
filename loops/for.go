package loops

import "fmt"

func Tmp() {
	score := 0
	for i := 11; i < 10; i++ {
		fmt.Println("score : ", score)
		fmt.Println("i : ", i)
		score += 1
		fmt.Println("score + i = ", score)
	}
}
