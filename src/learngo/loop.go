package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	if result == "" {
		return "0"
	}
	return result
}

func printfFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/*func forever() {
	for { // 死循环
		fmt.Println("abc")
	}
}*/
func main() {
	fmt.Println(
		converToBin(5),  //101
		converToBin(13), //1101
		converToBin(0),
	)

	printfFile("abc.txt")
	// forever()
}
