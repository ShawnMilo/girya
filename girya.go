package girya

import "fmt"

const max = 5
const min = 1

// LbsToKg converts pounds to kilograms
func LbsToKg(lbs float64) float64 {
	return lbs / 2.20462
}

// KgToLbs converts kilograms to pounds
func KgToLbs(lbs float64) float64 {
	return lbs * 2.20462
}

// Kettlebell describes the kettlebell
func Kettlebell() string {
	return "Kettlebell == strength!"
}

// Strength describes strength
func Strength() string {
	return "Strength == kettlebell!"
}

// PrintLadders prints out a list of the sets in a ladder.
func PrintLadders(rungs int) {
	if rungs < min || rungs > max {
		fmt.Printf("%d-rung ladders are invalid -- minimum %d, maximum %d.\n", rungs, min, max)
		return
	}

	for i := min; i <= rungs; i++ {
		fmt.Printf("%d reps with your non-dominant side, then %d reps with your dominant side. Rest.\n", i, i)
	}
	fmt.Printf("This is one %d-rung ladder. Start with three %d-rung ladders and work your way up to five.\n", rungs, rungs)
	if rungs < max {
		fmt.Printf("Work your way up to %d-rung ladders, then start over at three-rung ladders with a heavier weight if desired.\n", max)
	}
}

// TotalMass returns the total
func TotalMass(kg, ladders, rungs int) string {
	mass := float64(kg) * factorial(rungs) * float64(ladders)
	return fmt.Sprintf("Total mass lifted in workout: %.2f kg (%.2f lbs).", mass, KgToLbs(mass))
}

func factorial(i int) float64 {
	var f float64
	for n := i; n >= 1; n-- {
		f += float64(n)
	}
	return f
}
