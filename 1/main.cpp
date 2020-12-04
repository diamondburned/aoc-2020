#include <fstream>
#include <iostream>
#include <stdio.h>
#include <vector>

int part1(std::vector<int> numbers) {
  for (auto i : numbers) {
    for (auto j : numbers) {
      if (i + j == 2020) {
        return i * j;
      }
    }
  }
  return -1;
}

int part2(std::vector<int> numbers) {
  for (auto i : numbers) {
    for (auto j : numbers) {
      for (auto k : numbers) {
        if (i + j + k == 2020) {
          return i * j * k;
        }
      }
    }
  }
  return -1;
}

int main() {
  std::ifstream f;
  f.open("input", std::ios::in);

  std::vector<int> numbers;
  int tmp;

  while (f >> tmp) {
    numbers.push_back(tmp);
  }

  std::cout << part1(numbers) << "\n" << part2(numbers) << "\n";

  return 0;
}
