package reverse

func Reverse(num int) int {
	reverse := 0
	for num != 0 {
		reverse *= 10
		reverse += num % 10
		num /= 10
	}
	return reverse
}
