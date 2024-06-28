package loadingAnimation

import (
	"fmt"
	"time"
)

func LoadingAnimation() {
	// Define the characters of the animation
	chars := []string{"|", "/", "-", "\\"}

	// Loop to show the animation
	for i := 0; i < 100; i++ {
		for _, char := range chars {
			fmt.Print("\r" + char)
			time.Sleep(100 * time.Microsecond) // Adjust the speed of the animation
		}
	}
}