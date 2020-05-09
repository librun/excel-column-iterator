package iterator

const lenAlphabet = 25

var (
	convertMapIntToString = map[int]string{
		0:  "A",
		1:  "B",
		2:  "C",
		3:  "D",
		4:  "E",
		5:  "F",
		6:  "G",
		7:  "H",
		8:  "I",
		9:  "J",
		10: "K",
		11: "L",
		12: "M",
		13: "N",
		14: "O",
		15: "P",
		16: "Q",
		17: "R",
		18: "S",
		19: "T",
		20: "U",
		21: "V",
		22: "W",
		23: "X",
		24: "Y",
		25: "Z",
	}

	convertMapStringToInt = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
		"I": 8,
		"J": 9,
		"K": 10,
		"L": 11,
		"M": 12,
		"N": 13,
		"O": 14,
		"P": 15,
		"Q": 16,
		"R": 17,
		"S": 18,
		"T": 19,
		"U": 20,
		"V": 21,
		"W": 22,
		"X": 23,
		"Y": 24,
		"Z": 25,
	}
)

// ConvertIntToString function for convert from number column to string (excel)
func ConvertIntToString(in int) string {
	var result string

	if in < 1 {
		return result
	}

	num := in - 1
	for {
		denom := num % (lenAlphabet + 1)
		result = convertMapIntToString[denom] + result

		if num/(lenAlphabet+1) >= 1 {
			num = (num - 1 - denom) / (lenAlphabet + 1)
		} else {
			break
		}
	}

	return result
}

// ConvertStringToInt function for convert from string column name (excel) to number
func ConvertStringToInt(in string) int {
	return 1
}
