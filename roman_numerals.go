package kata

type RomanNumeralConverter interface {
	ConvertToRoman(i int) string
}

type RomanNumerals struct{}

func (r RomanNumerals) ConvertToRoman( input int) (result string) {
	result = ""
	return result
}
