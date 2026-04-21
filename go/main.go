package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func toPalindrome(str string) (bool, string) {

	freq := make(map[rune]int)
	for _, c := range str {
		freq[c]++
	}

	oddCount := 0
	var oddChar rune
	for char, count := range freq {
		if count%2 == 1 {
			oddCount++
			oddChar = char
		}
	}

	if oddCount > 1 {
		return false, ""
	}

	var half strings.Builder
	half.Grow(len(str) / 2)

	for char, count := range freq {
		half.WriteString(strings.Repeat(string(char), count/2))
	}

	result := half.String()
	if oddCount == 1 {
		result += string(oddChar)
	}

	reverseHalf := []rune(half.String())
	for i, j := 0, len(reverseHalf)-1; i < j; i, j = i+1, j-1 {
		reverseHalf[i], reverseHalf[j] = reverseHalf[j], reverseHalf[i]
	}
	result += string(reverseHalf)

	return true, result
}

func isValidEmail(email string) bool {

	if len(email) < 6 || len(email) > 30 {
		return false
	}

	atIndex := strings.Index(email, "@")
	if atIndex == -1 || strings.Count(email, "@") != 1 {
		return false
	}

	username := email[:atIndex]
	// _ = email[atIndex:]

	if len(username) > 0 && (username[0] == '.' || username[len(username)-1] == '.') {
		return false
	}

	for i, c := range email {
		if !(unicode.IsLetter(c) && (c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') ||
			unicode.IsDigit(c) ||
			c == '.' || c == '+' || c == '@') {
			return false
		}

		if i > 0 && i < atIndex && c == '.' && email[i-1] == '.' {
			return false
		}
	}

	return true
}

func sanitizeEmail(email string) string {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return email
	}

	username := email[:atIndex]
	domain := email[atIndex:]

	var cleanUsername strings.Builder
	cleanUsername.Grow(len(username))

	for _, c := range username {
		if c == '.' {
			continue
		}
		if c == '+' {
			break
		}
		cleanUsername.WriteRune(c)
	}

	return cleanUsername.String() + domain
}

func countUniqueEmails(emails []string) int {
	uniqueEmails := make(map[string]bool)

	for _, email := range emails {
		if isValidEmail(email) {
			cleanEmail := sanitizeEmail(email)
			uniqueEmails[cleanEmail] = true
		}
	}
	return len(uniqueEmails)
}

func isFullSquare(num int) bool {
	if num < 0 {
		return false
	}
	root := int(math.Sqrt(float64(num)))
	return root*root == num
}

func handlePalindrome(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: program -p <string1> <string2> ...")
		return
	}

	for _, str := range args {
		possible, palindrome := toPalindrome(str)
		if possible {
			fmt.Printf("[*] %s -> %s\n", str, palindrome)
		} else {
			fmt.Printf("[!] Cannot turn string '%s' into palindrome!\n", str)
		}
	}
}

func handleEmail(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: program -e <email1> <email2> ...")
		return
	}

	fmt.Println("[*] Your emails:")
	for _, email := range args {
		fmt.Printf("\t- %s\n", email)
	}

	uniqueCount := countUniqueEmails(args)
	fmt.Printf("[*] Of them unique are: %d\n", uniqueCount)
}

func handleSquare(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: program -s <num1> <num2> ...")
		return
	}

	for _, arg := range args {
		var num int
		_, err := fmt.Sscanf(arg, "%d", &num)
		if err != nil {
			fmt.Printf("[!] Invalid number: %s\n", arg)
			continue
		}

		if num < 0 {
			fmt.Printf("[!] Can't take a square from a negative number: %d\n", num)
		} else if isFullSquare(num) {
			fmt.Printf("[+] %d is a full square!\n", num)
		} else {
			fmt.Printf("[-] %d isn't a full square!\n", num)
		}
	}
}

func printUsage(programName string) {
	fmt.Printf("Usage: %s [-eps] ARGS\n", programName)
	fmt.Println("Options:")
	fmt.Println("  -p\tTurn the string to a palindrome, if possible")
	fmt.Println("  -e\tCount valid emails in ARGS")
	fmt.Println("  -s\tCount full squares in ARGS")
}

func main() {
	if len(os.Args) < 3 {
		printUsage(os.Args[0])
		os.Exit(1)
	}

	option := os.Args[1]
	args := os.Args[2:]

	switch option {
	case "-p":
		handlePalindrome(args)
	case "-e":
		handleEmail(args)
	case "-s":
		handleSquare(args)
	default:
		printUsage(os.Args[0])
		os.Exit(1)
	}
}
