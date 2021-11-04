package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	printFileContents(file)

}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

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

	printfFile("basic/branch/abc.txt")
	s := `abc"d
	kkk
	123
	
	p`
	printFileContents(strings.NewReader(s))
	// forever()
}
