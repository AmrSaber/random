package common

import (
	"crypto/rand"
	"fmt"
	"math"
)

// GetSecureFloat returns a cryptographically secure random floating-point number.
func GetSecureFloat() float64 {
	bytes := make([]byte, 7)
	rand.Read(bytes)

	randomValue := float64(bytes[0]%(1<<5)) / float64(1<<5)
	for _, b := range bytes[1:] {
		randomValue = (randomValue + float64(b)) / float64(1<<8)
	}

	return randomValue
}

// GetRandomIntInRange returns a random integer within the inclusive [lo, hi] range.
func GetRandomIntInRange(lo int, hi int) int {
	if lo > hi {
		Fail("lo must be less than or equal to hi")
	}

	return int(math.Floor(GetSecureFloat()*float64(hi-lo+1))) + lo
}

// GetRandomElement returns a random element from the given slice.
func GetRandomElement[T any](arr []T) T {
	if len(arr) == 0 {
		Fail("array is empty")
	}

	randomIndex := GetRandomIntInRange(0, len(arr)-1)
	return arr[randomIndex]
}

func getBooleanTypeValues(boolType string) []any {
	switch boolType {
	case BooleanTypeNumeric:
		return []any{1, 0}
	case BooleanTypeYesNo:
		return []any{"yes", "no"}
	case BooleanTypeTrueFalse:
		return []any{true, false}
	default:
		Fail(fmt.Sprintf("unknown boolean type [%s]", boolType))
		return nil
	}
}

// GetRandomBoolean returns a random boolean representation of the given type.
func GetRandomBoolean(boolType string) any {
	values := getBooleanTypeValues(boolType)
	return GetRandomElement(values)
}

// GetShuffledArray returns the numbers in [start, end] in random order.
func GetShuffledArray(start, end int) []int {
	shuffled := make([]int, 0, end-start+1)

	for value := start; value <= end; value++ {
		randomIndex := GetRandomIntInRange(0, len(shuffled))
		if randomIndex == len(shuffled) {
			shuffled = append(shuffled, value)
			continue
		}

		shuffled = append(shuffled, shuffled[randomIndex])
		shuffled[randomIndex] = value
	}

	return shuffled
}

// Shuffle returns a shuffled copy of the input slice.
func Shuffle[T any](arr []T) []T {
	shuffledIndexes := GetShuffledArray(0, len(arr)-1)
	result := make([]T, len(arr))

	for i, idx := range shuffledIndexes {
		result[i] = arr[idx]
	}

	return result
}

func getValidCharacters(strType string) []rune {
	switch strType {
	case StringTypeASCII:
		return append(ASCIILetters, Numbers...)
	case StringTypeNumbers:
		return Numbers
	case StringTypeLetters:
		return ASCIILetters
	case StringTypeExtended:
		return append(append(ASCIILetters, Numbers...), []rune("+-_$#/@!")...)
	case StringTypeHex:
		return HexDigits
	case StringTypeBase64:
		return append(append(ASCIILetters, Numbers...), []rune("+/")...)
	default:
		Fail(fmt.Sprintf("unknown type [%s]", strType))
		return nil
	}
}

// GetRandomString returns a random string of the given type and length.
func GetRandomString(strType string, length int) string {
	valid := getValidCharacters(strType)
	randomLetters := make([]rune, 0, length)

	for range length {
		randomLetters = append(randomLetters, GetRandomElement(valid))
	}

	if strType == StringTypeBase64 {
		for len(randomLetters)%4 != 0 {
			randomLetters = append(randomLetters, '=')
		}
	}

	return string(randomLetters)
}
