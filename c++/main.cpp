#include "email.cpp"
#include "palindrome.cpp"
#include <cmath>
#include <cstdlib>
#include <iostream>
#include <ostream>
#include <string>
#include <vector>

bool isFullSquare(int num) {
	return (num == std::pow(std::floor(std::sqrt(num)), 2));
}
bool toPalindrome(std::string &str);
int countUniqEmais(const std::vector<std::string> &emails);

void email(int argc, char *argv[]) {

	std::vector<std::string> strv;

	std::cout << "[*] Your emails:" << std::endl;

	for (int i = 2; i < argc; ++i) {
		strv.push_back(argv[i]);
		std::cout << "\t- " << argv[i] << "\n";
	}

	std::cout << "[*] Of them unique are: " << countUniqEmais(strv)
			  << std::endl;
}

void palindrome(int argc, char *argv[]) {

	for (int i = 2; i < argc; ++i) {

		std::string str{argv[i]};
		bool possible = toPalindrome(str);

		if (possible) {
			std::cout << "[*] " << str << '\n';
		} else {
			std::cout << "[!] Cannot turn string into palindrome!\n";
		}
	}
	return;
}

void sqare(int argc, char *argv[]) {

	for (int i = 2; i < argc; ++i) {
		int n = std::stoi(argv[i]);
		if (n < 0) {
			std::cout << "[!] Can't take a square from a negative number!\n";
		} else {
			std::cout << (isFullSquare(n) ? "[+] " : "[-] ") << argv[i]
					  << (isFullSquare(n) ? " is" : " isn't")
					  << " a full square!\n";
		}
	}
}

void err_usage(int argc, char **argv) {
	std::cout << "Usage: " << argv[0] << " [-eps] ARGS" << std::endl;
	std::cout << "Options:\n"
				 "  -p\tTurn the string to a palindrome, if possible\n"
				 "  -e\tCount valid emails in ARGS\n"
				 "  -s\tCount full squares in ARGS\n";
	exit(1);
}

int main(int argc, char *argv[]) {

	if (argc < 3)
		err_usage(argc, argv);

	switch (argv[1][1]) {
	case 'p':
		palindrome(argc, argv);
		break;
	case 'e':
		email(argc, argv);
		break;
	case 's':
		sqare(argc, argv);
		break;
	default:
		err_usage(argc, argv);
	}

	return 0;
}
