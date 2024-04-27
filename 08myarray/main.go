package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("The Fruit List is :", fruitList)
	fmt.Println("The Fruit List is :", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Mushrooms"}
	fmt.Println("The vegy list is :", vegList)
	fmt.Println("The vegy list is :", len(vegList))
}
