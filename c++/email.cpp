#include <cctype>
#include <cstddef>
#include <string>
#include <unordered_set>
#include <vector>

std::string sanitizeEmail(const std::string &email) {

	size_t at_index = email.find('@');

	std::string uname = email.substr(0, at_index);
	std::string domain = email.substr(at_index);

	std::string s_uname;
	for (char c : uname) {
		if (c == '.') {
			continue;
		} else if (c == '+') {
			break;
		} else {
			s_uname += c;
		}
	}

	return s_uname + domain;
}

bool isValidEmail(const std::string &email) {

	if ((email.length() < 6) || (email.length() > 30)) {
		return false;
	}

	if (email.front() == '.') {
		return false;
	}

	if ((email.find('@') == std::string::npos) ||
		(email.find_first_of('@') != email.find_last_of('@'))) {
		return false;
	}

	for (char c : email) {
		if (!(std::isalnum(c) || c == '.' || c == '+' || c == '@')) {
			return false;
		}
	}
	return true;
}

int countUniqEmais(const std::vector<std::string> &emails) {
	std::unordered_set<std::string> uniq_emails;

	for (const auto &email : emails) {
		if (isValidEmail(email)) {
			uniq_emails.insert(sanitizeEmail(email));
		}
	}

	return uniq_emails.size();
}
