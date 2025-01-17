package main

import (
	"bufio" // for files
	"fmt"
	"log" // for files
	"os" // for files
	 "strings"
	 "strconv" // for converting string to int
)

func main() {
	recipeList := getRecipeList()
	fmt.Println("Welcome to the Recipe Book!")
	choice := 10
	for choice != 0 {
		fmt.Print("\nHere are your options:\n1) View your recipies\n2) Add a recipe\n3) Edit a recipe\n4) See what you can bake\n0) Quit\nEnter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
			case 1: {
				for i := 0; i < len(recipeList); i++ {
					printRecipe(recipeList[i])
				}
			}
			case 2: {
				var newRecipe Recipe
				fmt.Print("Enter the name of the baked good: ")
				reader := bufio.NewReader(os.Stdin)
				n, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println(err)
				}
				newRecipe.name = n
				fmt.Println("Enter the ingredients (measurement unit ingredient):");
				ingredientCount := 1;
				for {
					var m float32
					var u string
					var i string
					fmt.Printf("%v) ", ingredientCount)
					fmt.Scanf("%g %s", &m, &u)
					reader := bufio.NewReader(os.Stdin)
					i, err := reader.ReadString('\n')
					if err != nil {
						fmt.Println(err)
					}
					ingr := Ingredient{i, m, u}
					if ingr.name != "\n" {
						ingredientCount++;
							newRecipe.ingredientList = append(newRecipe.ingredientList, ingr)
					} else {
						break
					}
				}
				fmt.Print("Enter the baking temperature: ")
				var temp int
				fmt.Scanf("%d", &temp)
					newRecipe.bakingTemp = temp
				fmt.Print("Enter the baking time: ")
				var time int
				fmt.Scanf("%d", &time)
				newRecipe.bakingTime = time
				recipeList = append(recipeList, newRecipe)
			}
			case 3:
			case 4:
			case 0: {
				writeRecipesToFile(recipeList)
				fmt.Println("Thank you for using my program.")
			}
		}
	}
}

func getRecipeList() []Recipe { // read line by line https://www.geeksforgeeks.org/how-to-read-a-file-line-by-line-to-string-in-golang/
	var rList []Recipe
	file, err := os.Open("recipes.txt")
	if err != nil {
		log.Fatalf("failed to open: %v", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		if strings.Index(scanner.Text(), "Name: ") != -1 {
			var rec Recipe
			rec.name = strings.Split(scanner.Text(), "Name: ")[1]
			rList = append(rList, rec)
		} else if strings.Index(scanner.Text(), "Bake temp: ") != -1 {
			bTempString := strings.Split(scanner.Text(), "Bake temp: ")[1]
			bTemp, err := strconv.Atoi(bTempString)
			if err != nil {
				fmt.Println("failed to convert string to int")
			}
			rList[len(rList) - 1].bakingTemp = int(bTemp)
		} else if strings.Index(scanner.Text(), "Bake time: ") != -1 {
			bTimeString := strings.Split(scanner.Text(), "Bake time: ")[1]
			bTime, err := strconv.Atoi(bTimeString)
			if err != nil {
				fmt.Println("failed to convert string to int")
			}
			rList[len(rList) - 1].bakingTime = int(bTime)
		} else if strings.Index(scanner.Text(), " - ") != -1 {
			ingredientString := strings.Split(scanner.Text(), " - ")[1]
			rList[len(rList) - 1].ingredientList = append(rList[len(rList) - 1].ingredientList, toIngredient(ingredientString))
		}
		text = nil // clears the array
		text = append(text, scanner.Text()) // appends the next line
	}
	file.Close()
	return rList
}

func recipeToString(r Recipe) string {
	s := "Name: " + r.name + "\nBake time: " + strconv.Itoa(r.bakingTime) + "\nBake temp: " + strconv.Itoa(r.bakingTemp) + "\nInredients:\n"
	for i := 0; i < len(r.ingredientList); i++ {
		s = s + " - " + ingredientToString(r.ingredientList[i])
	}
	return s
}

func ingredientToString(i Ingredient) string {
	var s string
	s = strconv.FormatFloat(float64(i.measure), 'f', -1, 32) + " " + i.measureUnit + " " + i.name + "\n"
	return s
}

func printRecipe(r Recipe) {
	fmt.Printf("\n%v\nBake time: %v minutes\nBake temp: %v F\nInredients:\n", r.name, r.bakingTime, r.bakingTemp)
	for i := 0; i < len(r.ingredientList); i++ {
		fmt.Print(" - ")
		printIngredient(r.ingredientList[i])
	}
}

func printIngredient(i Ingredient) {
	fmt.Println(i.measure, i.measureUnit, i.name)
}

func toIngredient(iString string) Ingredient {
	var i Ingredient
	list := strings.Split(iString, " ")
	m, err := strconv.ParseFloat(list[0], 32)
	if err != nil {
		fmt.Println("failed to convert string to int")
	}
	i.measure = float32(m)
	i.measureUnit = list[1]
	i.name = list[2]
	for word := 3; word < len(list); word++ {
		i.name = i.name + " " + list[word]
	}
	return i
}

func writeRecipesToFile(rList []Recipe) {
	file, err1 := os.Create("recipes.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	for i := 0; i < len(rList); i++ {
		l, err2 := file.WriteString(recipeToString(rList[i]))
		if err2 != nil {
			fmt.Println(l)
			fmt.Println(err2)
		}
	}
	err3 := file.Close()
	if err3 != nil {
		fmt.Println(err3)
	}
}

type Ingredient struct {
	name string
	measure float32
	measureUnit string
}

type Recipe struct {
	name string
	ingredientList []Ingredient
	bakingTemp int
	bakingTime int
}