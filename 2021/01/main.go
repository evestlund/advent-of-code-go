package main

import "fmt"

func main() {
	Something();
}

func Something() (int, error) {
	return fmt.Println("Hello, 世界")
}
