// A program that will flip a coin and output it
// By: IceKoda

package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
)

// Declares global count for results
var headc = 0
var tailc = 0

func main() {
	usagemsg := "Usage: coinflip [rounds]"

	// Checks to see if any arguments are passed. If not, flips coin once
	if len(os.Args) == 1 {
		coin()


		
	// The first argument checks for best x of y flips
	} else if len(os.Args) == 2 { 

		// Converts first argument to integer
		setting, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(usagemsg)
			os.Exit(0)
		}


		// checks if greater than 3 and odd number
		if setting >= 3 && setting % 2 == 1 {
			for i := 0; i < setting; i++ {
				coin()
			}
			var winner string
			if headc > tailc {
				winner = "HEADS"
			} else {
				winner = "TAILS"
			}
			// Outputs total count of each
			fmt.Println("\n===RESULTS=== \nHeads: ", headc , "\nTails: ", tailc, "\n\nWinner:", winner)
		} else {
			fmt.Println("Must be an odd integer greater than or equal to 3")
			os.Exit(0)
		}
	} else {
		fmt.Println(usagemsg)
	}
}

func coin() {
	coin := flip()
	if coin == 1 {
		fmt.Println("Heads")
		headc += 1
	} else {
		fmt.Println("Tails")
		tailc += 1
	}
}
// Flips coin, and returns heads or tails
func flip() int{
	return check(getInt())
}



// Gets integer [0..100]
func getInt() int {
	s := rand.NewSource(time.Now().UnixNano())
	num := rand.New(s)
	
	randnum := num.Intn(100)
	
	return randnum
}

// Checks to see if heads or tails
func check(num int) int {
	output := 0
	if num > 50 {
		output = 1	
	} else if num < 50 {
		output = 0
	} else {
		// Retries the function until output != 50
		output = flip()	
	}
	return output
}
