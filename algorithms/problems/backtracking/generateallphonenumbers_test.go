package backtracking

import (
	"fmt"
	"sgago/thestudyguide/col/queue"
	"testing"
)

/*

Generate All Letter Combinations from a Phone Number
https://algo.monster/problems/letter_combinations_of_phone_number

Given phone number digits, generate all the words you
can spell using:
1: N/A
2: A B C
3: D E F
4: G H I
5: J K L
6: M N O
7: P Q R S
8: T U V
9: W X Y Z

Input: "56"
Output: ["jm","jn","jo","km","kn","ko","lm","ln","lo"]

========================================

So, we can tackle this with either looping or recursion.
*/

func TestGenerateAllPhoneNumbers_56(t *testing.T) {
	actual := generateAllPhoneNumbersLooping("56")

	fmt.Println(actual)
}

func TestGenerateAllPhoneNumbers_919(t *testing.T) {
	actual := generateAllPhoneNumbersLooping("919Z")

	fmt.Println(actual)
}

type phoneNumState struct {
	num string
	idx int
}

var (
	numToChar = map[string][]string{
		"2": {"A", "B", "C"},
		"3": {"D", "E", "F"},
		"4": {"G", "H", "I"},
		"5": {"J", "K", "L"},
		"6": {"M", "N", "O"},
		"7": {"P", "Q", "R", "S"},
		"8": {"T", "U", "V"},
		"9": {"W", "X", "Y", "Z"},
	}
)

func generateAllPhoneNumbersLooping(phone string) []string {
	n := len(phone)
	ans := []string{}

	q := queue.New[phoneNumState](4 * n)

	num := string(phone[0])

	// Load initial states
	for _, letter := range numToChar[num] {
		q.EnqTail(phoneNumState{
			num: letter,
			idx: 0,
		})
	}

	for !q.Empty() {
		curr := q.DeqHead()

		if len(curr.num) == n {
			ans = append(ans, curr.num)

			continue
		}

		num := string(phone[curr.idx+1])

		if _, ok := numToChar[num]; !ok {
			// For unrecognized values like 1 or Z, just add a ? I guess?
			// Oh, the phone numbers will only be 2-9 from problem constraints
			// So, so bonus error handling I guess lol
			q.EnqTail(phoneNumState{
				num: curr.num + "?",
				idx: curr.idx + 1,
			})

			continue
		}

		for _, letter := range numToChar[num] {
			q.EnqTail(phoneNumState{
				num: curr.num + letter,
				idx: curr.idx + 1,
			})
		}
	}

	return ans
}

func TestGenerateAllPhoneNumbersRecursion_56(t *testing.T) {
	actual := GenerateAllPhoneNumbersRecursion("56")
	fmt.Println(actual)
}

func TestGenerateAllPhoneNumbersRecursion_919Z(t *testing.T) {
	actual := GenerateAllPhoneNumbersRecursion("919Z")

	fmt.Println(actual)
}

func GenerateAllPhoneNumbersRecursion(phone string) []string {
	ans := make([]string, 0)
	generateAllPhoneNumbersRecursion(phone, "", &ans)
	return ans
}

func generateAllPhoneNumbersRecursion(phone string, curr string, ans *[]string) {
	if len(curr) == len(phone) {
		*ans = append(*ans, curr)
		return
	}

	num := string(phone[len(curr)])

	if _, ok := numToChar[num]; !ok {
		generateAllPhoneNumbersRecursion(phone, curr+"?", ans)
		return
	}

	for _, letter := range numToChar[num] {
		generateAllPhoneNumbersRecursion(phone, curr+letter, ans)
	}
}
