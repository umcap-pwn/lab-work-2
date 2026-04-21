#include <algorithm>
#include <array>
#include <string>

bool toPalindrome(std::string &str) {

  std::array<int, 256> freq{};

  for (unsigned char c : str) {
    freq[c]++;
  }

  int odd_count = std::count_if(freq.begin(), freq.end(),
                                [](int count) { return count & 1; });

  if (odd_count > 1) {
    return false;
  }

  std::string half;
  half.reserve(str.size() / 2);

  char middle = 0;

  for (int i = 0; i < 256; ++i) {
    if (freq[i] & 1) {
      middle = static_cast<char>(i);
    }
    half.append(freq[i] / 2, static_cast<char>(i));
  }

  str = half;
  if (middle) {
    str += middle;
  }

  std::reverse(half.begin(), half.end());
  str.append(half);

  return true;
}
