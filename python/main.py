#!/usr/bin/env python

def to_palindrome(string):

    odd_chars = 0
    odd_char = ''
    freq: dict[str, int] = { }

    for c in string:
        if c in freq:
            freq[c] += 1
        else:
            freq[c] = 1


    for c in freq:
        if freq[c]%2 != 0:
            odd_chars += 1
            odd_char = c

    print("stat:", freq, odd_char, odd_chars)

    if odd_chars > 1:
        return False

    n_string: str = ""

    for c in freq.keys():
        n_string += c*(freq[c]//2)

    if odd_chars:
        n_string += odd_char

    second_half = n_string[::-1]

    n_string += second_half

    return n_string


def main():
    str_in = input("Sting: ")
    string = to_palindrome(str_in)
    print(string if string else "Cannot form palindorme")

if __name__ == "__main__":
    main()
