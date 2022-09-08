package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// eraseLastLines erases the last line in the command-line interface
// using ANSI escape codes.
//
// ANSI escape codes are standardized text patterns interpreted by
// terminals and terminal emulators that can be used to manipulate
// the cursor's position, font styling and do other actions on the
// command line interface.
func eraseLastLines(numLineToErase int) {
	if numLineToErase == 0 {
		panic("Can't erase 0 lines.")
	} else if numLineToErase < 0 {
		panic("Number of lines to erase must be positive.")
	}

	for i := 1; i <= numLineToErase; i++ {
		fmt.Print("\x1b[1F") // F = line above; E = line bellow; int = Number of lines
		fmt.Print("\x1b[2K") // K = Erase in line; 2 = Option for "clear entire line"
	}
}

// suspenseGenerator generates a slice of pseudo-random integer values
// between 1 and 3 that emulates the decision process of your hypothetical
// opponent. Its only parameter corresponds to the length of the generated slice.
func suspenseGenerator(suspenseDuration int) []int {
	if suspenseDuration <= 0 {
		panic("Generated slice can't be of length <= 0")
	}

	// Preventing it from generating the same numbers always
	randSrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSrc)

	generateValidNumber := func(randInstance *rand.Rand) int {
		return randInstance.Intn(3) + 1
	}

	slice := make([]int, suspenseDuration)
	for index := range slice {
		slice[index] = generateValidNumber(rnd)
		// If it's the first iteration, skip the rest of the loop
		if index == 0 {
			continue
		}

		// If the generated value is equals to the previous one,
		// keep generating until they are different.
		for {
			if slice[index] == slice[index-1] {
				slice[index] = generateValidNumber(rnd)
			} else {
				break
			}
		}
	}

	return slice
}

// checkWinner takes weapon codes to determine if the user won based
// in the hypothetical opponent's weapon.
// 1 = Rock,
// 2 = Paper,
// 3 = Scissors.
func checkWinner(weaponCode, opponentsWeaponCode int) bool {
	if weaponCode < 1 && weaponCode > 3 {
		panic("Invalid weapon code for user.")
	} else if opponentsWeaponCode < 1 && opponentsWeaponCode > 3 {
		panic("Invalid weapon code for opponent.")
	}

	switch {
	case weaponCode == 1 && opponentsWeaponCode == 3:
		return true
	case weaponCode == 2 && opponentsWeaponCode == 1:
		return true
	case weaponCode == 3 && opponentsWeaponCode == 2:
		return true
	default:
		return false
	}
}

// weaponCodeToStr converts the weapon code (1 <= x <= 3) into
// the weapon's name.
// 1 = Rock,
// 2 = Paper,
// 3 = Scissors.
func weaponCodeToStr(weaponCode int) string {
	if weaponCode < 1 && weaponCode > 3 {
		panic("Invalid weapon code.")
	}
	switch weaponCode {
	case 1:
		return "Rock"
	case 2:
		return "Paper"
	default:
		return "Scissors"
	}
}

func main() {
	fmt.Println("******************************")
	fmt.Println("* Rock, Paper, and Scissors! *")
	fmt.Println("******************************")
	fmt.Println("\n 1) Rock\n 2) Paper\n 3) Scissors")

	scanWeapon := func() string {
		var weapon string
		fmt.Scanln(&weapon)
		return weapon
	}

	// Input validation.
	// If the code make it through the forever loop it means
	// the user has inputted a valid value.
	fmt.Print("\nChoose your weapon: ")
	nextNumLinesToErase := 1
	chosenOption := scanWeapon()
	for {
		if chosenOption == "" {
			eraseLastLines(nextNumLinesToErase)
			fmt.Println("Empty input. Try again.")
			fmt.Print("Choose your weapon: ")

			nextNumLinesToErase = 2
			chosenOption = scanWeapon()
		} else if !strings.ContainsAny(chosenOption, "123") || len(chosenOption) > 1 {
			eraseLastLines(nextNumLinesToErase)
			fmt.Println("Input is not a valid option. Try again.")
			fmt.Print("Choose your weapon: ")

			nextNumLinesToErase = 2
			chosenOption = scanWeapon()
		} else {
			fmt.Println() // New line after valid input
			break
		}
	}

	// Loop through a generated slice that represents the deciding process
	// of your hypothetic opponent. For each option of weapon check if it
	// matches the generated value. If yes, draw a nice arrow before it.
	// If not, just print it indented.
	//
	// Erase the last 3 lines and repeat the above process to emulate a
	// random-decision mechanic.
	fmt.Println("Waiting for opponent's decision:")

	options := [3]string{"1) Rock", "2) Paper", "3) Scissors"}
	opponentDecisions := suspenseGenerator(5)

	for opponentIndex, opponentValue := range opponentDecisions {
		// for each in options...
		for optionIndex, optionValue := range options {
			if opponentValue-1 == optionIndex {
				fmt.Printf(" > %s\n", optionValue)
			} else {
				fmt.Printf("   %s\n", optionValue)
			}
		}

		time.Sleep(1e9) // Wait for 1 seconds
		if opponentIndex < len(opponentDecisions)-1 {
			eraseLastLines(3) // Erase last 3 lines (printed options)
		}
	}

	opponentWeapon := opponentDecisions[len(opponentDecisions)-1] // int
	userWeapon, userErr := strconv.Atoi(chosenOption)             // Converting string input to int
	if userErr != nil {                                           // Error handling conversion
		panic(userErr)
	}

	fmt.Println()
	fmt.Printf("Match: %s vs. %s\n\n", weaponCodeToStr(userWeapon), weaponCodeToStr(opponentWeapon))

	// Print winner
	if userWeapon == opponentWeapon {
		fmt.Println("Draw!")
	} else if checkWinner(userWeapon, opponentWeapon) {
		fmt.Println("You Won! Congratulations!")
	} else {
		fmt.Println("Your opponent won!")
	}
}
