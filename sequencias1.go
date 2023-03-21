package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Print("Qual o primeiro nuemro? ")
	fmt.Scan(&num1)
	fmt.Print("Qual o segundo nuemro? ")
	fmt.Scan(&num2)
	fmt.Print("Qual o terceiro nuemro? ")
	fmt.Scan(&num3)

	fmt.Println("A soma dos seus tres numeros é de:", num1+num2+num3)

}
