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

/**
 * Dynamic Programming
 */
long long int research(int i, const std::vector<int> &inst,
             std::unordered_map<int, long long int> &storage) {
  if (i == inst.size() - 1) {
    return 1;
  }
  // If Value Exists in the storage, return it (memoized)
  if (storage.find(i) != storage.end()) {
    return storage[i];
  }
  long long int total = 0;
  for (int j = i + 1; j < inst.size(); j++) {
    // Valid Charge Diff
    if (inst[j] - inst[i] <= 3) {
      total += research(j, inst, storage);
    }
  };
  storage[i] = total;
  return total;
};

int main() {
  std::set<int> st;
  loadItSet(st);
  auto pt1 = part1(st);
  std::cout << pt1 << std::endl;
  std::vector<int> pt2(st.begin(), st.end());
  pt2.insert(pt2.begin(), 0);
  pt2.emplace_back(pt2.back() + 3);
  for(auto i : pt2){
      std::cout<< i << std::endl;
  }
  std::unordered_map<int, long long int> memo;
  std::cout << research(0, pt2, memo) << std::endl;
};
