package practice

/*
Problem:
Ref: https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/

Given:
- A string contains digits from 2-9

Return:
- All possible letter combinations that the number could represent

Constraints:
- digits length: [1, 4]
- digits[i]: [2, 9]
- mapping of digits to letters is in telephone buttons
*/

/*
First approach:
I will iterate all the digits.
For each digit, I will use all the possible letters to create combinations
*/

func letterCombinations(digits string) []string {
	var combinations []string
	if len(digits) == 0 {
		return combinations
	}

	possibleLetters := getMappingLetters(digits[0])
	combs := letterCombinations(digits[1:])
	if len(combs) == 0 {
		for _, letter := range possibleLetters {
			combinations = append(combinations, string(letter))
		}
	} else {
		for _, comb := range combs {
			for _, letter := range possibleLetters {
				combinations = append(combinations, string(letter)+comb)
			}
		}
	}

	return combinations
}

func getMappingLetters(digit byte) []byte {
	switch digit {
	case '2':
		return []byte{'a', 'b', 'c'}
	case '3':
		return []byte{'d', 'e', 'f'}
	case '4':
		return []byte{'g', 'h', 'i'}
	case '5':
		return []byte{'j', 'k', 'l'}
	case '6':
		return []byte{'m', 'n', 'o'}
	case '7':
		return []byte{'p', 'q', 'r', 's'}
	case '8':
		return []byte{'t', 'u', 'v'}
	case '9':
		return []byte{'w', 'x', 'y', 'z'}
	}

	return nil
}
