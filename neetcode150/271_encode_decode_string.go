package neetcode150

import "strconv"

type Solution struct{}

//// Input: ["we","say",":","yes"]
//// 2|we3|say3|yes
//// Output: ["we","say",":","yes"]
//// Input: ["we", "say", ":", "yes", "!@#$%^&*()"]

// Encode takes a slice of strings and encodes them into a single string.
// Pattern Used: Delimiter + Length Encoding
// - For each string, we prefix it with its length and a delimiter ('|').
// - Example: ["hello", "world"] -> "5|hello5|world"
func (s *Solution) Encode(strs []string) string {
	result := ""

	// Loop through each string in the input
	for _, v := range strs {
		// Convert length of the string to string form
		// Add delimiter "|" + the string itself
		result += strconv.Itoa(len(v)) + "|" + v
	}
	return result
}

// Decode takes the encoded string and reconstructs the original slice of strings.
// - It reads characters until it finds a '|', which marks the end of a length.
// - Converts that length into an integer.
// - Reads exactly that many characters after '|' to extract the string.
func (s *Solution) Decode(str string) []string {
	var result []string
	lenString := "" // will temporarily hold the digits representing length

	// Loop through the encoded string character by character
	for i := 0; i < len(str); i++ {
		// If we encounter a '|', that means the length part has ended
		if str[i] == '|' {
			// Convert collected length string to an integer
			currentLen, _ := strconv.Atoi(lenString)

			// Extract substring of exactly 'currentLen' characters
			word := str[i+1 : i+1+currentLen]
			result = append(result, word)

			// Reset length collector
			lenString = ""

			// Move index forward to skip the characters we just processed
			i += currentLen
			continue
		}

		// If it's a digit (part of length), keep appending
		lenString += string(str[i])
	}
	return result
}
