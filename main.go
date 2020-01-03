package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("\nWelcome to Drum Machine!")
	printMenu()
}

/*
 * Prints the main menu with 5 possible options (4 patterns and exit)
 * Validates the input by looping back to printMenu() as default case
 */
func printMenu() {
	fmt.Println("\nPlease select from the following options:")
	fmt.Println("\n1) Play Four on the floor (a pattern in 4/4 time)\n" +
		"2) Play Gravity by John Mayer (a pattern in 6/8 time)\n" +
		"3) Play Take Five by Dave Brubeck (a pattern in 5/4, a compound meter)\n" +
		"4) Play More Cowbell (it's just like it sounds)\n" +
		"5) Exit")
	fmt.Print("\nWhat would you like to do? (Please enter number 1-5): ")
	input, err := getUserInput()
	if err != nil {
		os.Exit(0) // We already output the error in the func, so just exit here.
	}
	switch input {
	case "1":
		play4OTF()
	case "2":
		playGravity()
	case "3":
		playTake5()
	case "4":
		playCowbell()
	case "5":
		exitDM()
	default:
		fmt.Println("\nI'm sorry. I didn't understand your input.")
		printMenu()
	}
}

/*
 * First, asks for user input for tempo specification for Four on the Floor pattern.
 * Calls validateTempo to validate
 * Calls to struct files to get data for Four on the Floor Pattern
 * Calls playPattern to play the prepared pattern.
 */
func play4OTF() {
	fmt.Println("\nYou've selected to play Four on the Floor!")
	fmt.Println("\nAt what tempo would you like this pattern to play?")
	fmt.Print("Please enter a tempo between 40-200 (Largo to Presto):")
	tempostr, err := getUserInput()
	if err != nil {
		os.Exit(0) // We already output the error in the func, so just exit here.
	}
	tempo := validateTempo(tempostr)
	if tempo == -1 {
		play4OTF() // error message is printed in func, so just try again here
	}
	ptrn := prepare4OTF(tempo)
	playPattern(ptrn)
	printMenu()
}

/*
 * First, asks for user input for tempo specification for Gravity pattern.
 * Calls validateTempo to validate
 * Calls to struct files to get data for Gravity Pattern
 * Calls playPattern to play the prepared pattern.
 */
func playGravity() {
	fmt.Println("\nYou've selected to play Gravity!")
	fmt.Println("\nAt what tempo would you like this pattern to play?")
	fmt.Print("Please enter a tempo between 40-200 (Largo to Presto):")
	tempostr, err := getUserInput()
	if err != nil {
		os.Exit(0) // We already output the error in the func, so just exit here.
	}
	tempo := validateTempo(tempostr)
	if tempo == -1 {
		playGravity() // error message is printed in func, so just try again here
	}
	ptrn := prepareGravity(tempo)
	playPattern(ptrn)
	printMenu()

}

/*
 * First, asks for user input for tempo specification for Take Five pattern.
 * Calls validateTempo to validate
 * Calls to struct files to get data for Take Five Pattern
 * Calls playPattern to play the prepared pattern.
 */
func playTake5() {
	fmt.Println("\nYou've selected to play Take Five!")
	fmt.Println("\nAt what tempo would you like this pattern to play?")
	fmt.Print("Please enter a tempo between 40-200 (Largo to Presto):")
	tempostr, err := getUserInput()
	if err != nil {
		os.Exit(0) // We already output the error in the func, so just exit here.
	}
	tempo := validateTempo(tempostr)
	if tempo == -1 {
		playTake5() // error message is printed in func, so just try again here
	}
	ptrn := prepareTake5(tempo)
	playPattern(ptrn)
	printMenu()
}

/*
 * First, asks for user input for tempo specification for We Need More Cowbell pattern.
 * Calls validateTempo to validate
 * Calls to struct files to get data for We Need More Cowbell Pattern
 * Calls playPattern to play the prepared pattern.
 */
func playCowbell() {
	fmt.Println("\nYou've selected to play We Need More Cowbell!")
	fmt.Println("\nAt what tempo would you like this pattern to play?")
	fmt.Print("Please enter a tempo between 40-200 (Largo to Presto):")
	tempostr, err := getUserInput()
	if err != nil {
		os.Exit(0) // We already output the error in the func, so just exit here.
	}
	tempo := validateTempo(tempostr)
	if tempo == -1 {
		playCowbell() // error message is printed in func, so just try again here
	}
	ptrn := prepareCowbell(tempo)
	playPattern(ptrn)
	printMenu()
}

/*
 * Reusable input capture and error checking (we do this a lot)
 * return: user input as a string or empty string on error.
 */
func getUserInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", err
	}
	return input, err
}

/*
 * First, verifies that string input can be an int
 * Then, verifies that int is within given range.
 * I chose to limit the range to 40-200 as that incorporates most musical tempi from Largo to Presto.
 * return: provided tempo string as an int, or -1 on error.
 */
func validateTempo(tempostr string) int {
	var tempo int
	var err error
	if tempo, err = strconv.Atoi(tempostr); err != nil {
		log.Println("\nYour input must be an integer. Please try again.")
		return -1
	}
	if !(tempo >= 40) || !(tempo <= 200) {
		fmt.Println("The tempo you entered was not between 40 and 200 (Largo - Presto). Please try again.")
		return -1
	}
	return tempo
}

/*
 * First, displays the pattern info
 * Then, determines the number of milliseconds to wait between playing each beat.
 * Loops through instrumentation slice to print contents in time determined above.
 */
func playPattern(ptrn Pattern) {
	fmt.Println("\n---Pattern Info---")
	fmt.Printf("Name: %s\nTime Signature: %d/%d\nTempo: %d\nNumber of Bars: %d\n",
		ptrn.name, ptrn.timeSig[0], ptrn.timeSig[1], ptrn.tempo, ptrn.bars)
	fmt.Println()

	/* This is a super ugly line of code. It's concise, but it needs some explanation:
	 * Given a BPM specification, we first need to determine how many beats per second (float64(ptrn.tempo) / 60.0)
	 * i.e., 120 BPM = 2 BPS
	 * Then, we determine how many seconds are between each beat (((1.0 / (float64(ptrn.tempo) / 60.0))
	 * i.e., 2 BPS = .5 second between beats
	 * Then, we need to determine what value we place on the beat (float64(ptrn.divsPerBar) / float64(ptrn.timeSig[0]))
	 * i.e., a 4/4 bar with 8 divisions means each division represents 1/2 of a beat.
	 * Finally, we multiply the whole thing by 1000 to get milliseconds. All ints are cast as floats for precision.
	*/
	tempoInt := 1000 * ((1.0 / (float64(ptrn.tempo) / 60.0)) / (float64(ptrn.divsPerBar) / float64(ptrn.timeSig[0])))
	n := ptrn.bars
	for n > 0 {
		for i, beat := range ptrn.instPerBeat {
			if i%ptrn.divsPerBar == 0 {
				fmt.Println()
			}
			if len(beat) > 1 {
				fmt.Print("[")
				for j, s := range beat {
					fmt.Print(s)
					if j != len(beat)-1 {
						fmt.Print("+")
					}
				}
				fmt.Print("]")
			} else if len(beat) == 0 {
				fmt.Print("[x]")
			} else {
				fmt.Print(beat)
			}
			time.Sleep(time.Duration(tempoInt) * time.Millisecond)
		}
		n--
	}
	fmt.Println()
}

/*
 * Handles program exit logic
 */
func exitDM() {
	fmt.Print("\nAre you sure you want to exit (y/n)? ")
	input, err := getUserInput()
	if err != nil {
		os.Exit(0) // We already output the error in the func, so just exit here.
	}
	switch strings.ToLower(input) {
	case "y":
		fmt.Println("\nThanks for using Drum Machine!\nBye! :)")
		fmt.Println()
		os.Exit(0)
	case "n":
		fmt.Println("\nNo problem!")
		printMenu()
	default:
		fmt.Println("\nI'm sorry. I didn't understand your input.\nPlease enter y for yes or n for no.")
		exitDM()
	}
}
