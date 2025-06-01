package easy

// RomanToInteger MCMXCIV
func RomanToInteger(s string) int {
	r1 := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	r2 := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			if num, ok := r2[string(s[i])+string(s[i+1])]; ok {
				result += num
				i++
				continue
			}
		}
		result += r1[string(s[i])]
	}
	return result
}
