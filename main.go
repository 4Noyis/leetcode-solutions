package main

import "fmt"

func main() {

	s := convert("PAYPALISHIRING", 3)
	fmt.Println(s)
	// a := longestPalindrome("aasbabc")
	//nums1 := [4]int{-2, 4, -1, 3}
	//nums2 := [3]int{-5, 0, 2}
	//answer := findMedianSortedArrays(nums1[:], nums2[:])

}

// #1
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil

}

// #2
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummyHead.Next
}

// 6.8.2025 #3
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	b := []byte(s)
	maxLength := 0

	// Try starting from each position
	for i := 0; i < len(b); i++ {
		var subString []byte

		// Build substring starting from position i
		for j := i; j < len(b); j++ {
			currentLetter := b[j]

			// Check if current letter already exists in substring
			found := false
			for k := 0; k < len(subString); k++ {
				if subString[k] == currentLetter {
					found = true
					break
				}
			}

			if found {
				// Found duplicate, stop building this substring
				break
			} else {
				// Add current letter to substring
				subString = append(subString, currentLetter)
			}
		}

		// Update max length if current substring is longer
		if len(subString) > maxLength {
			maxLength = len(subString)
		}
	}

	fmt.Printf("Longest substring length for '%s': %d\n", s, maxLength)
	return maxLength
}

// 6.8.2025 #4
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var merged []int

	merged = append(merged, nums1...)
	merged = append(merged, nums2...)
	fmt.Println(merged)

	// sorting
	for i := 0; i < len(merged); i++ {
		for j := i + 1; j < len(merged); j++ {
			if merged[i] > merged[j] {
				merged[i], merged[j] = merged[j], merged[i]
			}
		}
	}
	fmt.Println(merged)

	if len(merged)%2 == 0 {
		mid1 := len(merged)/2 - 1
		mid2 := len(merged) / 2
		return (float64(merged[mid1]) + float64(merged[mid2])) / 2
	} else {
		mid := len(merged) / 2
		return float64(merged[mid])
	}

}

// 7.8.2025 #5
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	b := []byte(s)
	maxPalindrome := string(b[0:1]) // Start with first character

	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] == b[j] {
				// Found matching characters, check if substring between them is a palindrome
				substring := b[i : j+1]

				isPalin := true
				left := 0
				right := len(substring) - 1
				for left < right {
					if substring[left] != substring[right] {
						isPalin = false
						break
					}
					left++
					right--
				}

				if isPalin && len(substring) > len(maxPalindrome) {
					maxPalindrome = string(substring)
				}
			}
		}
	}

	return maxPalindrome
}

// 10.8.2025 #6
func convert(s string, numRows int) string {

	if numRows <= 1 {
		return s
	}

	var firstArray []string
	var secondArray []string

	chunkSize := numRows + (numRows - 2) // Total chunk size: numRows + diagonal chars

	// Process string in chunks
	for i := 0; i < len(s); i += chunkSize {
		// Get the current chunk
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunk := s[i:end]

		// First part: first numRows characters
		if len(chunk) >= numRows {
			firstArray = append(firstArray, chunk[:numRows])
		} else if len(chunk) > 0 {
			// Handle incomplete chunk
			firstArray = append(firstArray, chunk)
		}

		// Second part: the diagonal characters (numRows-2 characters)
		diagonalCount := numRows - 2
		if len(chunk) > numRows && diagonalCount > 0 {
			diagonalStart := numRows
			diagonalEnd := diagonalStart + diagonalCount
			if diagonalEnd > len(chunk) {
				diagonalEnd = len(chunk)
			}
			if diagonalStart < len(chunk) {
				secondArray = append(secondArray, chunk[diagonalStart:diagonalEnd])
			}
		}
	}

	rows := make([][]rune, numRows)

	// Reconstruct the zigzag pattern
	var result []rune
	maxLen := len(firstArray)
	if len(secondArray) > maxLen {
		maxLen = len(secondArray)
	}

	for i := 0; i < maxLen; i++ {
		// Add characters from firstArray (vertical part)
		if i < len(firstArray) {
			for j, char := range firstArray[i] {
				if j < numRows {
					rows[j] = append(rows[j], char)
				}
			}
		}

		// Add characters from secondArray (diagonal part)
		if i < len(secondArray) && numRows > 2 {
			// Place diagonal characters in reverse order (bottom-up)
			diagonalChars := []rune(secondArray[i])
			for j, char := range diagonalChars {
				rowIndex := numRows - 2 - j // Place from second-to-last row upward
				if rowIndex >= 1 && rowIndex < numRows-1 {
					rows[rowIndex] = append(rows[rowIndex], char)
				}
			}
		}
	}

	// Read rows to create final zigzag result
	for _, row := range rows {
		result = append(result, row...)
	}

	finalResult := string(result)
	fmt.Printf("Zigzag result: %s\n", finalResult)

	return finalResult

}

// #9
func isPalindrome(x int) bool {

	switch {
	case x < 0:
		return false
	case x >= 0: // Changed from x > 0 to handle x = 0
		result := 0
		n := x // Use := for cleaner declaration

		for n > 0 {
			result = result*10 + n%10 // Use n instead of x
			n = n / 10
		}

		return x == result // Simplified return
	}
	return false
}
