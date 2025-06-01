package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInputFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("ошибка в числе '%s': %w", part, err)
			}
			nums = append(nums, num)
		}
	}

	return nums, scanner.Err()
}

func minSubsequenceCoveringAlphabet(nums []int) interface{} {
	mp := map[int]struct{}{}
	for i := 1; i <= 26; i++ {
		mp[i] = struct{}{}
	}

	count := map[int]int{}
	found := 0
	left := 0
	minLen := len(nums) + 1

	for right := 0; right < len(nums); right++ {
		val := nums[right]
		if _, ok := mp[val]; ok {
			count[val]++
			if count[val] == 1 {
				found++
			}
		}

		for found == 26 {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			lval := nums[left]
			if _, ok := mp[lval]; ok {
				count[lval]--
				if count[lval] == 0 {
					found--
				}
			}
			left++
		}
	}

	if minLen > len(nums) {
		return "NONE"
	}
	return minLen
}

func main() {
	nums, err := readInputFile("data_prog_contest_problem_2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	result := minSubsequenceCoveringAlphabet(nums)
	fmt.Println(result)
}
