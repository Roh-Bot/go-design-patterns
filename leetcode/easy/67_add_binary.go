package easy

// Input: a = "1010", b = "1011"
// Output: "10101"
func addBinary(a string, b string) string {
	//Make a byte array to store the result set
	var result []byte

	//initialize a carry and two pointers to iterate each binary strings
	carry := 0
	i, j := len(a)-1, len(b)-1

	// Loop in reverse until either of the strings has their start yet to be reached
	// or carry is greater than zero (for last 1+1 -> carry 1 equation edge case)
	for i >= 0 || j >= 0 || carry > 0 {
		// sum is a placeholder for carry for each iteration equation
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		//here we prepend the newly solved equation instead of appending
		//sum%2 for the case where sum can be 3
		result = append([]byte{byte(sum%2 + '0')}, result...)

		// if both strings have 1 in their current iteration
		// along with sum having value 1 then
		// carry will be 2. Hence, division by 2 as carry can never be 2
		carry = sum / 2
	}

	// convert the byte array to string and return
	return string(result)
}
