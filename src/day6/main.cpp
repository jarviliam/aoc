#include <algorithm>
#include <fstream>
#include <iostream>
#include <unordered_map>
#include <unordered_set>
#include <vector>

void loadIt(std::vector<std::vector<char>> &input) {
  std::ifstream file("./input.txt");
  if (file.is_open()) {
    std::string entry;
    std::vector<char> temp;
    while (std::getline(file, entry)) {
      if (entry == "") {
        input.emplace_back(temp);
        temp.clear();
        continue;
      }
      for (auto i : entry) {
        temp.push_back(i);
      }
    };
    // Last Vec
    input.emplace_back(temp);
  };
};

int part2() {
  std::ifstream file("./input.txt");
  int ans = 0;
  // std::vector<std::
  if (file.is_open()) {
    std::string entry;
    int ppl = 0;
    std::unordered_map<char, int> vals;
    while (std::getline(file, entry)) {
      if (entry == "") {
        int localAns = 0;
        /*std::find_if(vals.begin(), vals.end(), [&ppl]() {

        });*/
        for (auto i : vals) {
          if (i.second == ppl) {
            ans++;
          }
        };
        vals.clear();
        ppl = 0;
        continue;
      }
      ppl++;
      for (auto i : entry) {
        auto it = vals.find(i);
        if (vals.end() != it) {
          it->second++;
        } else {
          vals[i] = 1;
        }
      }
    };
    // Last one
    for (auto i : vals) {
      if (i.second == ppl) {
        ans++;
      }
    }
  };
  return ans;
};

int main() {
  std::vector<std::vector<char>> lanes;
  loadIt(lanes);
  int sum = 0;
  for (auto i : lanes) {
    std::sort(i.begin(), i.end());
    int unique = std::unique(i.begin(), i.end()) - i.begin();
    sum += unique;
    std::cout << i.size() << "Unqiue " << unique << std::endl;
  }
  std::cout << "Ans" << sum << std::endl;
  std::cout << part2() << std::endl;
};
