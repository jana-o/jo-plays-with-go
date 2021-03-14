package main

import "fmt"

func main() {

	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}

	fmt.Println(countItemsMatchingRule(items, "color", "silver"))
}

func countItemsMatchingRule(items [][]string, ruleKey string, ruleValue string) int {
	hm := map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}
	result := 0

	for i := range items {
		if items[i][hm[ruleKey]] == ruleValue {
			result++
		}
		fmt.Println(hm)

	}

	return result
}

func countItemsMatchingRule2(items [][]string, ruleKey string, ruleValue string) int {
	result := 0
	for _, item := range items {
		if ruleKey == "type" {
			if ruleValue == item[0] {
				result++
			}
		} else if ruleKey == "color" {
			if ruleValue == item[1] {
				result++
			}
		} else {
			if ruleValue == item[2] {
				result++
			}
		}
	}
	return result
}
