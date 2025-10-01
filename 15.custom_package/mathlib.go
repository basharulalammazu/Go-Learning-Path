package custom_package

// Here function start with capital letter so it is exported
// and can be accessed from other packages
func Add(num1 int, num2 int) int {
	return num1 + num2
}