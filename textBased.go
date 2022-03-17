package main

import "fmt"

// Prints out intro text
func intro() {
	fmt.Println("Welcome to the ASCII People Printer Application for Golang! (APPAG)")
	fmt.Println("Within this program, you'll have the chance to \nmake some super neat people out of basic text!")
	fmt.Println("This program will take your inputs and create a neat little fella for you!")
	fmt.Println("\nNOTES:\n - Any improper inputs will create a 'default' person.\n - All persons made through this program are claimed by APPAGCO for use in NFT creation.")
	fmt.Println("====================================================================")
}

// Deals with user input and calls to build a person
func interact() {
	fmt.Println("Let's start! There are 5 different attributes to control!\n")
	fmt.Println("First off, your person needs a name!\nWhat's their first name?:")
	var first string
	fmt.Scanln(&first)

	fmt.Println("And of course, they need a last name!\nWhat would that be?:")
	var last string
	fmt.Scanln(&last)

	fmt.Println("Great! Next we need their sex: man, woman or non-binary?\nPlease enter 'M', 'W', or 'NB':")
	var sex string
	fmt.Scanln(&sex)
	if sex != "M" && sex != "W" && sex != "NB" {
		fmt.Println("--- UNKNOWN INPUT: DEFAULTING TO NON-BINARY ---")
		sex = "NB"
	} else {
		fmt.Print("Great! ")
	}

	fmt.Println("Are they short, medium, or tall?\nPlease enter 'S', 'M', or 'T':")
	var height string
	fmt.Scanln(&height)
	if height != "S" && height != "M" && height != "T" {
		fmt.Println("--- UNKNOWN INPUT: DEFAULTING TO MEDIUM HEIGHT ---")
		height = "M"
	} else {
		fmt.Print("Great! ")
	}

	fmt.Println("Finally, input any one character to be their head!:")
	var head string
	fmt.Scanln(&head)
	if len(head) != 1 {
		fmt.Println("--- UNKNOWN INPUT: DEFAULTING TO '0' AS HEAD ---")
		head = "0"
	} else {
		fmt.Println("Great!")
	}

	buildAPerson(first, last, sex, height, head)
}

// Prints out the cool person the user made
func buildAPerson(first string, last string, sex string, height string, head string) {
	fmt.Println("Let's build you a person!\n\n")
	fmt.Println("====================================================================")
	fmt.Println(" ,")
	fmt.Printf("(_%s \n", head)

	if sex == "NB" {
		sex = "3"
	}

	// Switch for heights
	if height == "S" {
		fmt.Printf("  %s)\nS", sex)
	}
	if height == "M" {
		fmt.Printf("  %s)\n  |\n", sex)
	}
	if height == "T" {
		fmt.Printf("  %s)\n  |\n  |\n", sex)
	}

	fmt.Println(" / \\")
	fmt.Println()
	fmt.Println("Introducing.....", first, last, "!")
	fmt.Println()
	fmt.Println("====================================================================")

}

// Prints closing remarks
func close() {
	fmt.Println()
	fmt.Println("Well, that's about it!\nThanks for using my program!")
}

func main() {
	intro()
	interact()
	close()
}
