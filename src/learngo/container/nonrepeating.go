package main

import (
	"fmt"
)

// leetcode 3.无重复字符的最长子串
// 我第一次的做法
func lengthOfLongestSubstring1(s string) int {
	mp := make(map[uint8]int)
	start := 1
	maxLength := 0
	length := 0
	for i := range s {
		if _, ok := mp[s[i]]; ok && mp[s[i]] >= start {
			t := start
			start = mp[s[i]] + 1
			length = length - (start - t)
			mp[s[i]] = i + 1
			length++
		} else {
			mp[s[i]] = i + 1
			length++
		}
		if maxLength < length {
			maxLength = length
		}
	}
	return maxLength
}

// ccmouse做法
func lengthOfLongestSubstring2(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

// 国际版
func lengthOfLongestSubstring3(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
func main() {
	res := lengthOfLongestSubstring1("abcabcbb")
	fmt.Println(res)

	res = lengthOfLongestSubstring1("bbbbb")
	fmt.Println(res)

	res = lengthOfLongestSubstring1("p1232dakew")
	fmt.Println(res)

	res = lengthOfLongestSubstring2("abcabcbb")
	fmt.Println(res)

	res = lengthOfLongestSubstring2("bbbbb")
	fmt.Println(res)

	res = lengthOfLongestSubstring2("p1232dakew")
	fmt.Println(res)

	res = lengthOfLongestSubstring2("一二三二一")
	fmt.Println(res) // 5，说明只支持ASCII字符，这个算法不能很好的支持中文

	res = lengthOfLongestSubstring3("一二三二一") // 3
	fmt.Println(res)

	fmt.Println(
		lengthOfLongestSubstring3("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
