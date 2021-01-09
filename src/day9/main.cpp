#include <algorithm>
#include <fstream>
#include <iostream>
#include <set>
#include <vector>

void loadIt(std::vector<unsigned long int> &input) {
  std::ifstream file("./input.txt");
  if (file.is_open()) {
    std::string entry;
    while (std::getline(file, entry)) {
      input.emplace_back(std::stoul(entry));
    };
  };
};
std::pair<unsigned long int, std::vector<unsigned long int>::iterator>
part1(int threshold, std::vector<unsigned long int> &instructions) {
  std::vector<unsigned long int> prevVals = {};
  std::vector<unsigned long int>::iterator i;
  std::vector<unsigned long int>::iterator final;
  unsigned long int answer;
  for (i = instructions.begin(); i < instructions.end(); i++) {
    int index = i - instructions.begin();
    if (index < threshold) {
      prevVals.emplace_back(*i);
      continue;
    };
    bool found = false;
    for (auto x : prevVals) {
      int target = *i - x;
      auto res = std::find(prevVals.begin(), prevVals.end(), target);
      if (res != prevVals.end()) {
        found = true;
        break;
      }
    }
    if (found == false) {
      answer = *i;
      final = i;
      break;
    }
    prevVals.erase(prevVals.begin());
    prevVals.emplace_back(*i);
  };
  return std::pair<unsigned long int, std::vector<unsigned long int>::iterator>(
      answer, final);
}
unsigned long int part2(std::vector<unsigned long int>::iterator itr,
                        std::vector<unsigned long int> &instructions) {
  unsigned long int answer = 0;
  unsigned long int pt1Ans = *itr;
  for (auto i = instructions.begin(); i < itr; i++) {
    if (answer != 0) {
      break;
    }
    unsigned long int temp = pt1Ans - *i;
    std::set<unsigned long int> history{*i};
    for (auto j = i + 1; j < itr; j++) {
      history.emplace(*j);
      temp -= *j;
      if (temp == 0) {
        answer = (*history.begin() + *history.rbegin());
      } else if (temp < 0) {
        break;
      }
    };
  };
  return answer;
}

int main() {
  std::vector<unsigned long int> instructions;
  loadIt(instructions);
  const int threshold = 25;
  auto pt1 = part1(threshold, instructions);
  auto itr = pt1.second;
  auto x = part2(itr, instructions);
  std::cout << pt1.first << std::endl;
  std::cout << x << std::endl;
};
