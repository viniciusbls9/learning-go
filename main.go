package main

import (
	"fmt"
)

func main() {
	result, err := soma(1, 20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result, err)
	}

}

func soma (a, b int) (int, error) {
	if a + b > 10 {
		return 0, fmt.Errorf("Soma maior que 10")
	}

	return a + b, nil
}