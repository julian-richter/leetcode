package main

// Given a string `s`, find the length of the longest substring without duplicate characters.

/*
ok so the longest substring without duplicates...

first thought: maybe check every possible substring?
but that's like... check substring starting at 0, then 1, then 2...
for each starting point, check all possible endings
that are nested loops again, which is crazy slow

the tricky part: how do I know if something's a duplicate?
hash map again, store each character and maybe its position
when I see a character already in the map, that's my duplicate

example: "abcabcbb"
- a b c → all good, length 3
- hit 'a' again → move a left pointer past the first 'a'
- keep going, track max length seen

in C I'd probably use an array for ASCII chars (256 slots),
but Go's maps are easier and handle Unicode too
Rob Pike loves Unicode, so Go probably handles it well

the window slides through the string once
each character gets added and maybe removed once
so it scales linearly with string length, way better than checking all substrings
*/

func LengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	maxLength := 0
	left := 0

	for right, char := range s {
		// if we've seen this char, and it's in our current window
		if lastIndex, exists := seen[char]; exists && lastIndex >= left {
			// move the left pointer past the duplicate
			left = lastIndex + 1
		}

		// update the character's position
		seen[char] = right

		// calculate the current window size
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
