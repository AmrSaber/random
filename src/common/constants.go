package common

// Boolean type constants
const (
	BOOLEAN_TYPE_TRUE_FALSE string = "true_false"
	BOOLEAN_TYPE_NUMERIC    string = "numeric"
	BOOLEAN_TYPE_YES_NO     string = "yes_no"
)

// String type constants
const (
	STRING_TYPE_HEX      string = "hex"
	STRING_TYPE_ASCII    string = "ascii"
	STRING_TYPE_BASE_64  string = "base64"
	STRING_TYPE_NUMBERS  string = "numbers"
	STRING_TYPE_LETTERS  string = "letters"
	STRING_TYPE_EXTENDED string = "extended"
)

var STRING_TYPES = []string{
	STRING_TYPE_HEX,
	STRING_TYPE_ASCII,
	STRING_TYPE_NUMBERS,
	STRING_TYPE_LETTERS,
	STRING_TYPE_EXTENDED,
	STRING_TYPE_BASE_64,
}

// Possible string values
var (
	ASCII_LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	NUMBERS       = []rune("0123456789")
	HEX_DIGITS    = []rune("0123456789abcdef")
)

const DEFAULT_DELIMITER = " "
const DEFAULT_STRING_LENGTH = 20

const (
	ID_TYPE_UUID4 = "uuidv4"
	ID_TYPE_UUID7 = "uuidv7"
	ID_TYPE_NANO  = "nanoid"
)

var ID_TYPES = []string{
	ID_TYPE_UUID4,
	ID_TYPE_UUID7,
	ID_TYPE_NANO,
}
