package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the class of slices")

	var fruitList = []string{"Tomato", "Apple", "Peach"}
	fmt.Printf("The type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // 1 included but 3 not included
	fmt.Println(fruitList)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 946
	highscores[2] = 465
	highscores[3] = 898
	// highscores[4] = 234

	highscores = append(highscores, 458, 789, 125)

	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// how to remove a element form a index based slices
	var courses = []string{"reactjs", "javascript", "golang", "swift", "Python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
