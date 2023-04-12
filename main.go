package main

import (
	"desafios/easy"
)

func main() {
	// words := "aab,cab,abba,abacus,scabs,scab,bacca,ab"
	// letters := "aabbscb"

	// easy.EncontraPalavras(letters, words)
	// // ===============================
	// nums := []int{3, 3}
	// target := 6

	// result := easy.TwoSum(nums, target)
	// fmt.Println(result)
	// ===============================

	// x := -12222221
	// fmt.Println("É palindromo: ", easy.IsPalindrome(x))

	// =====================================================
	// s := "IV"
	// fmt.Println(easy.RomanToInt(s))

	// =====================================================
	// strs := []string{"flower", "flower", "flower", "flower"}
	// fmt.Println("Prefixo:", easy.LongestCommonPrefix(strs))

	// =====================================================
	// stringValid := "()[]{}"
	// fmt.Println("É uma string valida?", easy.IsValid(stringValid))

	// =====================================================
	list1 := &easy.ListNode{Val: 1}
	current := list1
	current.Next = &easy.ListNode{Val: 2}
	current = current.Next
	current.Next = &easy.ListNode{Val: 4}

	list2 := &easy.ListNode{Val: 1}
	current = list2
	current.Next = &easy.ListNode{Val: 3}
	current = current.Next
	current.Next = &easy.ListNode{Val: 4}

	easy.MergeTwoLists(list1, list2)
	easy.PrintLinkedList(list1)

	// =====================================================

}
