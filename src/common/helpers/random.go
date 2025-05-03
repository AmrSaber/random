package helpers

import (
	"crypto/rand"
	"fmt"
	"math"

	"github.com/AmrSaber/random/v3/src/common"
)

// returns a random floating point number created with
// cryptographically secure pseudo random generator (crypto).
func GetSecureFloat() float64 {
	// Generate 7 random bytes
	bytes := make([]byte, 7)
	rand.Read(bytes)

	// Shift 5 bits from first byte by 5 to the right
	randomValue := float64(bytes[0]%(1<<5)) / float64(1<<5)

	// For each of the following 6 bytes, add its value and shift it 8 bits to the right
	for _, byte := range bytes[1:] {
		randomValue = (randomValue + float64(byte)) / float64(1<<8)
	}

	return randomValue
}

// returns a random integer in the given range,
// both start and end are inclusive.
func GetRandomIntInRange(min int, max int) int {
	if min > max {
		common.Fail("min must be less than or equal to max")
	}

	return int(math.Floor(GetSecureFloat()*float64(max-min+1))) + min
}

// returns a random element from the given slice.
func GetRandomElement[T any](arr []T) T {
	if len(arr) == 0 {
		common.Fail("array is empty")
	}

	randomIndex := GetRandomIntInRange(0, len(arr)-1)
	return arr[randomIndex]
}

// returns boolean values for the given type.
func getBooleanTypeValues(boolType string) []any {
	switch boolType {
	case common.BOOLEAN_TYPE_NUMERIC:
		return []any{1, 0}
	case common.BOOLEAN_TYPE_YES_NO:
		return []any{"yes", "no"}
	case common.BOOLEAN_TYPE_TRUE_FALSE:
		return []any{true, false}
	default:
		common.Fail(fmt.Sprintf("unknown boolean type [%s]", boolType))
		return nil
	}
}

// returns a random boolean of the specified type.
func GetRandomBoolean(boolType string) any {
	values := getBooleanTypeValues(boolType)
	return GetRandomElement(values)
}

// returns a shuffled array with the elements [start .. end]
// This uses a variation of fisher-yates algorithm
func GetShuffledArray(start, end int) []int {
	shuffled := make([]int, 0, end-start+1)

	for value := start; value <= end; value++ {
		randomIndex := GetRandomIntInRange(0, len(shuffled))

		if randomIndex == len(shuffled) {
			shuffled = append(shuffled, value)
		} else {
			shuffled = append(shuffled, shuffled[randomIndex])
			shuffled[randomIndex] = value
		}
	}

	return shuffled
}

// returns a shuffled copy of the input slice
func Shuffle[T any](arr []T) []T {
	shuffledIndexes := GetShuffledArray(0, len(arr)-1)
	result := make([]T, len(arr))

	for i, idx := range shuffledIndexes {
		result[i] = arr[idx]
	}

	return result
}

// returns an array of valid characters for the given type
func getValidCharacters(strType string) []rune {
	switch strType {
	case common.STRING_TYPE_ASCII:
		return append(common.ASCII_LETTERS, common.NUMBERS...)
	case common.STRING_TYPE_NUMBERS:
		return common.NUMBERS
	case common.STRING_TYPE_LETTERS:
		return common.ASCII_LETTERS
	case common.STRING_TYPE_EXTENDED:
		return append(append(common.ASCII_LETTERS, common.NUMBERS...), []rune("+-_$#/@!")...)
	case common.STRING_TYPE_HEX:
		return common.HEX_DIGITS
	case common.STRING_TYPE_BASE_64:
		return append(append(common.ASCII_LETTERS, common.NUMBERS...), []rune("+/")...)
	default:
		common.Fail(fmt.Sprintf("unknown type [%s]", strType))
		return nil
	}
}

// returns a random string of the given type and length
func GetRandomString(strType string, length int) string {
	valid := getValidCharacters(strType)
	randomLetters := make([]rune, 0, length)

	for i := 0; i < length; i++ {
		randomLetters = append(randomLetters, GetRandomElement(valid))
	}

	// Base64 string must have length that is a multiple of 4
	if strType == common.STRING_TYPE_BASE_64 {
		for len(randomLetters)%4 != 0 {
			randomLetters = append(randomLetters, '=')
		}
	}

	return string(randomLetters)
}
