package goboringavatars

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// hashCode computes a hash code for a given string
func hashCode(name string) int {

	var hash int32 = 0

	for i := 0; i < len(name); i++ {
		char := name[i]
		hash = (hash << 5) - hash + int32(char)
	}

	if hash < 0 {
		hash = -hash
	}

	return int(hash)

}

// getDigit returns the nth digit of a number
func getDigit(number, ntn int) int {
	return (number / int(math.Pow(10, float64(ntn)))) % 10
}

// getBoolean returns a boolean based on the nth digit of a number
func getBoolean(number, ntn int) bool {
	return getDigit(number, ntn)%2 == 0
}

// GetUnit computes a unit value based on a number, range, and index
func getUnit(number, ra, index int) float64 {

	value := float64(number % ra)
	if index != 0 && getDigit(number, index)%2 == 0 {
		return -value
	}
	return value
}

// GetRandomColor returns a random color from the provided slice
func getRandomColor(number int, colors []string) string {
	return colors[number%len(colors)]
}

// GetContrast determines the contrast color (black or white) for the given hex color.
// It returns an error if the input is not a valid hex color.
func getContrast(hexcolor string) (string, error) {
	// Normalize the hex color string
	hexcolor = strings.TrimPrefix(hexcolor, "#")
	if len(hexcolor) != 6 {
		return "", errors.New("invalid hex color format")
	}

	// Parse the RGB components from the hex color
	r, err := strconv.ParseInt(hexcolor[0:2], 16, 64)
	if err != nil {
		return "", fmt.Errorf("invalid red component: %w", err)
	}
	g, err := strconv.ParseInt(hexcolor[2:4], 16, 64)
	if err != nil {
		return "", fmt.Errorf("invalid green component: %w", err)
	}
	b, err := strconv.ParseInt(hexcolor[4:6], 16, 64)
	if err != nil {
		return "", fmt.Errorf("invalid blue component: %w", err)
	}

	// Calculate the YIQ luminance value
	yiq := ((r * 299) + (g * 587) + (b * 114)) / 1000

	// Determine the contrast color
	if yiq >= 128 {
		return "#000000", nil
	} else {
		return "#FFFFFF", nil
	}
}
