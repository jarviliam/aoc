#include <algorithm>
#include <fstream>
#include <iostream>
#include <set>
#include <unordered_map>
#include <vector>

void loadItSet(std::set<int> &input) {
  std::ifstream file("./input.txt");
  if (file.is_open()) {
    std::string entry;
    while (std::getline(file, entry)) {
      input.emplace(std::stoi(entry));
    };
  };
};
int part1(std::set<int> &instructions) {
  int device = *instructions.rbegin() + 3;
  std::unordered_map<int, int> map;
  int start = 0;
  int diff = 0;
  std::set<int>::iterator itr;
  for (itr = instructions.begin(); itr != instructions.end(); itr++) {
    auto next = std::next(itr, 1);
    std::cout << "Next : " << *next << " Curr : " << *itr << std::endl;
    if ((*next - *itr) > 3 || (*next - *itr) < 0) {
      std::cout << "error" << *itr << std::endl;
      break;
    }
    std::cout << "Diff : " << (*next - *itr) << std::endl;
    map[(*next - *itr)]++;
    diff += (*next - *itr);
  }

  // Add 1 For 0 entry
  for (auto &i : map) {
    std::cout << '{' << i.first << ", " << i.second << '}' << std::endl;
  }

  int answer = ((map[3] + 1) * (map[1] + 1));
  return answer;
}

int main() {
  std::set<int> st;
  loadItSet(st);
  auto pt1 = part1(st);
  std::cout << pt1 << std::endl;
};
