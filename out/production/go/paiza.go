package main

import (
	"bufio"
	"os"
	"fmt"
)

func scan() string{
	reader := bufio.NewReader(os.Stdin);
	input, _ := reader.ReadString('\n');
	return input;
}
func main() {
	test:=scan();
	fmt.Printf(test);
}
