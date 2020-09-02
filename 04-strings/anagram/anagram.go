package anagram

//determine if string is anagram of second string
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[int]int)
	for i := range s {
		m[int(s[i])]++
	}

	for i := range t {
		m[int(t[i])]--
		if m[int(t[i])] < 0 {
			return false
		}
	}
	return true
}

//this is slower
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	arr := make([]int, 26)

// 	for _, v := range s {
// 		arr[v-'a'] = arr[v-'a'] + 1
// 	}

// 	for _, v := range t {

// 		if arr[v-'a'] <= 0 {
// 			return false
// 		} else {
// 			arr[v-'a'] = arr[v-'a'] - 1
// 		}
// 	}

// 	for _, v := range s {
// 		if arr[v-'a'] != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
