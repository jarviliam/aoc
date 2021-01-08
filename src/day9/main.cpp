#include <algorithm>
#include <fstream>
#include <iostream>
#include <vector>

void loadIt(std::vector<unsigned long int> &input) {
  std::ifstream file("./input.txt");
  if (file.is_open()) {
    std::string entry;
    while (std::getline(file, entry)) {
      std::cout << entry << std::endl;
      input.emplace_back(std::stoul(entry));
    };
  };
};
std::pair<unsigned long int, std::vector<unsigned long int>::iterator>
part1(const int &threshold, std::vector<unsigned long int> &instructions) {
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
        std::cout << "Found Target" << *res << std::endl;
        found = true;
        break;
      }
    }
    if (found == false) {
      std::cout << "Did not add " << *i << std::endl;
      answer = *i;
      final = i;
      break;
    }
    prevVals.erase(prevVals.begin());
    prevVals.emplace_back(*i);
  };
  return std::pair<unsigned long, std::vector<unsigned long int>::iterator>(
      answer, final);
}

int main() {
  std::vector<unsigned long int> instructions;
  loadIt(instructions);
  const int threshold = 25;
  auto pt1 = part1(25, instructions);
  std::cout << pt1.first << std::endl;
  std::cout << "part 2" << part2(pt1.second, instructions) << std::endl;
};

unsigned long int part2(std::vector<unsigned long int>::iterator itr,
                        std::vector<unsigned long int> &instructions) {
  unsigned long int answer;
  unsigned long int pt1Ans = *itr;
  std::cout << pt1Ans << std::endl;
  for (auto i = instructions.begin(); i < itr; i++) {
  };

  return answer;
}
