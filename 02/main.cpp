#include <cstdio>
#include <iostream>
#include <unordered_map>

// I did this out of obligation.

class Requirement {
  int from, to;
  char letter;
  std::string password;

  Requirement(std::string line) {
    auto pos = line.find("-");
    auto fromstr = line.substr(0, pos);
    from = std::stoi(fromstr);

    auto from = 0;
  };
};

int main() {
  int from, to;
  char letter;
  char password;

  std::scanf("%d-%d %c: %s\n", &from, &to, &letter, &password);
}
