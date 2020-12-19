#include <fstream>
#include <iostream>
#include <unordered_map>
#include <unordered_set>
#include <vector>

typedef std::vector<std::unordered_map<std::string, std::string>> choo;

bool isfourDigits(std::string x) { return x.size() == 4; }
bool inrange(int x, int l, int m) { return x >= l && x <= m; };
int part2(choo passports, std::vector<std::string> list) {
  int valid = passports.size();
  for (auto i : passports) {
    for (auto k : list) {
      if (i.find(k) == i.end() && k != "cid") {
        valid--;
        break;
      }
      // DONE
      if (k == "cid")
        continue;
      // DONE
      else if (k == "byr") {
        if (!isfourDigits(i[k])) {
          valid--;
          break;
        } else {
          if (!inrange(std::stoi(i[k]), 1920, 2002)) {
            valid--;
            break;
          };
        }
      } else if (k == "eyr") {
        if (!isfourDigits(i[k])) {
          valid--;
          break;
        } else {
          if (!inrange(std::stoi(i[k]), 2020, 2030)) {
            valid--;
            break;
          };
        }
      } else if (k == "iyr") {
        if (!isfourDigits(i[k])) {
          valid--;
          break;
        } else {
          if (!inrange(std::stoi(i[k]), 2010, 2020)) {
            valid--;
            break;
          };
        }
      } else if (k == "hgt") {

      } else if (k == "hcl") {
          if((i[k].substr(0,1) != '#') || 

      } else if (k == "ecl") {

      } else {
      }
    };
  };
  return valid;
};

void loadEntry(std::string en,
               std::unordered_map<std::string, std::string> &map) {
  int colon = en.find(":");
  map[en.substr(0, colon)] = en.substr(colon + 1, en.length() - 1);
};

void loadIt(choo &input) {
  std::ifstream file("./input.txt");
  if (file.is_open()) {
    std::string entry;
    int pp = 0;
    std::unordered_map<std::string, std::string> test;
    input.emplace_back(test);
    while (std::getline(file, entry)) {
      if (entry.empty()) {
        pp++;
        std::unordered_map<std::string, std::string> test2;
        input.emplace_back(test2);
      } else {
        while (entry.length() != 0) {
          int spaceDelim = entry.find(" ");
          if (spaceDelim != std::string::npos) {
            std::string s = entry.substr(0, spaceDelim);
            loadEntry(s, input[pp]);
            entry.erase(0, spaceDelim + 1);
          } else {
            loadEntry(entry, input[pp]);
            entry.erase(0, entry.length());
          }
        };
      }
    };
    std::cout << input.size() << std::endl;
  };
}

int part1(choo passports, std::vector<std::string> list) {
  int valid = passports.size();
  for (auto i : passports) {
    for (auto k : list) {
      if (k == "cid")
        continue;
      if (i.find(k) == i.end()) {
        valid--;
        break;
      }
    };
  };
  return valid;
}

int main() {
  choo passports{};
  std::vector<std::string> list = {"byr", "iyr", "eyr", "hgt",
                                   "hcl", "ecl", "pid", "cid"};
  loadIt(passports);

  std::cout << "Passports : " << passports.size() << std::endl;

  std::cout << "Valid :" << part1(passports, list) << std::endl;
  std::cout << "Crazy :" << part2(passports, list) << std::endl;
};
