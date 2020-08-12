package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// for keeping things easy to read and type-safe
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}
	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	screen.Clear()

	for {
		screen.MoveTopLeft()

		times := time.Now()
		hour, min, second := times.Hour(), times.Minute(), times.Second()
		clock := []placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[second/10], digits[second%10],
		}

		for line := range clock[0] {
			// Print a line for each placeholder in clock
			for index, digit := range clock {
				// Colon blink on every two seconds.
				// + On each sec divisible by two, prints an empty line
				// + Otherwise: prints the current pixel
				next := clock[index][line]
				if digit == colon && second%2 == 0 {
					next = "   "
				}
				// Print the next line and,
				// give it enough space for the next placeholder
				fmt.Print(next, "  ")
			}

			// After each line of a placeholder, print a newline
			fmt.Println()
		}

		// pause for 1 second
		time.Sleep(time.Second)
	}

}
