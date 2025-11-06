package TwoSum

// Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

/*
ok so we need two numbers that add up to the target right

first thought: maybe just check every pair
but like... that's nested loops i'm not gonna play the loop game, i mean it works, but its not it
100 numbers = 100 * 100 = 10,000 checks
1000 numbers? Literally a million checks. That's crazy on a whole different level

i remember doing this in C with hash maps
the complement trick, for each number just find (target - number)
it's literally just a lookup problem

In C i would use a hash map and lookups are basically instant
Go's maps are probably the same thing because the creators of go wrote C before Go
they know hash maps are crazy so Go's implementation is probably crazy good

So instead of "find two numbers that sum to target"
maybe i should think about it like: "does the complement of this number exist?"

Single pass goes like:
  - see 2, need 7 for target 9 → check map, not there yet
  - see 7, need 2 for target 9 → check map, found it

With the map: 100 numbers = 100 checks, 1000 numbers = 1000 checks
scales linearly instead of exponentially - way better
the map uses memory (stores everything we've seen) but maybe speed matters more

C really drilled hash maps into my head and now it's paying off in Go
love it
*/
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for index, num := range nums {
		complement := target - num

		if complementIndex, exists := seen[complement]; exists {
			return []int{complementIndex, index}
		}

		seen[num] = index
	}

	return nil
}
