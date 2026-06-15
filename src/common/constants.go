// Package common provides shared constants, helpers, and state for the random CLI.
package common

// BooleanTypeTrueFalse identifies the true/false boolean style.
const BooleanTypeTrueFalse string = "true_false"

// BooleanTypeNumeric identifies the numeric boolean style.
const BooleanTypeNumeric string = "numeric"

// BooleanTypeYesNo identifies the yes/no boolean style.
const BooleanTypeYesNo string = "yes_no"

// StringTypeHex identifies the hexadecimal string type.
const StringTypeHex string = "hex"

// StringTypeASCII identifies the ASCII alphanumeric string type.
const StringTypeASCII string = "ascii"

// StringTypeBase64 identifies the base64 string type.
const StringTypeBase64 string = "base64"

// StringTypeNumbers identifies the numeric-only string type.
const StringTypeNumbers string = "numbers"

// StringTypeLetters identifies the letter-only string type.
const StringTypeLetters string = "letters"

// StringTypeExtended identifies the extended symbol string type.
const StringTypeExtended string = "extended"

// StringTypes enumerates all supported string generation types.
var StringTypes = []string{
	StringTypeHex,
	StringTypeASCII,
	StringTypeNumbers,
	StringTypeLetters,
	StringTypeExtended,
	StringTypeBase64,
}

// ASCIILetters lists all ASCII letters used for random string generation.
var ASCIILetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Numbers lists all numeric digits used for random string generation.
var Numbers = []rune("0123456789")

// HexDigits lists all hexadecimal digits used for random string generation.
var HexDigits = []rune("0123456789abcdef")

// DefaultDelimiter is the default separator for list outputs.
const DefaultDelimiter = " "

// DefaultStringLength is the default length used for generated strings.
const DefaultStringLength = 20

// IDTypeUUID4 identifies UUID version 4 IDs.
const IDTypeUUID4 = "uuidv4"

// IDTypeUUID7 identifies UUID version 7 IDs.
const IDTypeUUID7 = "uuidv7"

// IDTypeNano identifies NanoID IDs.
const IDTypeNano = "nanoid"

// IDTypes enumerates all supported identifier generation types.
var IDTypes = []string{
	IDTypeUUID4,
	IDTypeUUID7,
	IDTypeNano,
}

// RootOptions captures global flags shared by the CLI commands.
type RootOptions struct {
	Count int
}
