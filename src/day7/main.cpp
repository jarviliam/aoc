#include <algorithm>
#include <deque>
#include <fstream>
#include <iostream>
#include <unordered_map>
#include <unordered_set>
#include <vector>

void loadIt(std::vector<std::string> &input) {
  std::ifstream file("./test.txt");
  if (file.is_open()) {
    std::string entry;
    while (std::getline(file, entry)) {
      input.emplace_back(entry);
    };
  };
};

std::vector<std::string> split(std::string &s, std::string delimiter) {
  std::vector<std::string> list;
  size_t pos = 0;
  std::string token;
  while ((pos = s.find(delimiter)) != std::string::npos) {
    token = s.substr(0, pos);
    list.push_back(token);
    s.erase(0, pos + delimiter.length());
  }
  list.push_back(s);
  return list;
}

int part1(std::vector<std::string> l) {
  int ans;

  std::unordered_map<std::string, std::vector<std::string>> list;
  for (auto i : l) {
    if (i != "") {
      auto words = split(i, " ");
      std::string holder = words[0] + words[1] + words[2];
      holder.erase(holder.size() - 1, 1);
      if (words[words.size() - 3] == "no") {
        continue;
      };
      int idx = 4;
      while (idx < words.size()) {
        std::string bag =
            words[idx] + words[idx + 1] + words[idx + 2] + words[idx + 3];
        // Clean bag
        if (bag[bag.size() - 1] == '.') {
          bag.erase(bag.size() - 1, 1);
        }
        if (bag[bag.size() - 1] == ',') {
          bag.erase(bag.size() - 1, 1);
        }
        if (bag[bag.size() - 1] == 's') {
          bag.erase(bag.size() - 1, 1);
        }
        if (std::isdigit(bag[0])) {
          bag.erase(0, 1);
        }
        if (list.find(bag) == list.end()) {
          list[bag] = std::vector<std::string>{};
        };
        list[bag].emplace_back(holder);
        idx += 4;
      };
    }
  };

  // BFS
  std::unordered_set<std::string> seen;
  std::string t = "shinygoldbag";
  std::deque<std::string> dq{t};
  while (dq.size() > 0) {
    std::string x = dq[0];
    dq.pop_front();
    if (seen.find(x) != seen.end()) {
      continue;
    }
    seen.insert(x);

    for (auto y : list[x]) {
      dq.emplace_back(y);
    }
  };
  ans = seen.size() - 1;
  return ans;
}
int main() {
  std::vector<std::string> instructions;
  loadIt(instructions);
  std::cout << part1(instructions) << std::endl;
};
