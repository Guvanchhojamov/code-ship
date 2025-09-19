package daily

/*
3136. Valid Word
A word is considered valid if:
- It contains a minimum of 3 characters.
- It contains only digits (0-9), and English letters (uppercase and lowercase).
- It includes at least one vowel.
- It includes at least one consonant.

You are given a string word.
Return true if word is valid, otherwise, return false.

Notes:

'a', 'e', 'i', 'o', 'u', and their uppercases are vowels.
A consonant is an English letter that is not a vowel.


Example 1:

Input: word = "234Adas"

Output: true

Explanation:

This word satisfies the conditions.

Example 2:

Input: word = "b3"

Output: false

Explanation:

The length of this word is fewer than 3, and does not have a vowel.

Example 3:

Input: word = "a3$e"

Output: false

Explanation:

This word contains a '$' character and does not have a consonant.

Constraints:

1 <= word.length <= 20
word consists of English uppercase and lowercase letters, digits, '@', '#', and '$'.
*/
/*
 well what we need to do is:
 A word is considered valid if:
- It contains a minimum of 3 characters.
- It contains only digits (0-9), and English letters (uppercase and lowercase).
- It includes at least one vowel.
- It includes at least one consonant.
 len(s)>=3
 0<a<9 OR aA<s<zZ
 vowel count >= 1
 consonant count >=1

 brute force:
  -iterate over chars and check each option.
  tc: O(n)
  sc: O(n)
*/

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	var hasVowel bool
	var hasConsonant bool

	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	for _, char := range word {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return false
		}

		if vowels[char] {
			hasVowel = true
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			hasConsonant = true
		}
	}

	return hasVowel && hasConsonant
}
