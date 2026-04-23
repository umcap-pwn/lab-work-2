fun canFormPalindrome(s: String): Boolean {
    val charCount = mutableMapOf<Char, Int>()
    for (char in s) {
        charCount[char] = charCount.getOrDefault(char, 0) + 1
    }
    val oddCount = charCount.values.count { it % 2 != 0 }
    return oddCount <= 1
}

fun rearrangeToPalindrome(s: String): String? {
    if (!canFormPalindrome(s)) return null

    val charCount = mutableMapOf<Char, Int>()
    for (char in s) {
        charCount[char] = charCount.getOrDefault(char, 0) + 1
    }

    val half = StringBuilder()
    var middle = ""

    for ((char, count) in charCount) {
        if (count % 2 != 0) {
            middle = char.toString()
        }
        repeat(count / 2) {
            half.append(char)
        }
    }

    return half.toString() + middle + half.reverse().toString()
}

fun main() {
    print("Введите строку: ")
    val input = readLine() ?: ""

    val result = rearrangeToPalindrome(input)

    if (result != null) {
        println("Palindrome: $result")
    } else {
        println("Cannot form palindrome")
    }
}
