/**
 * Yuck
 */

#include <algorithm>
#include <fstream>
#include <iostream>
#include <unordered_map>
#include <unordered_set>
#include <vector>

typedef std::vector<std::unordered_map<std::string, std::string>> choo;

typedef bool (*myfunc)(std::string x);

bool isfourDigits(std::string x) { return x.size() == 4; }
bool inrange(int x, int l, int m) { return x >= l && x <= m; };
bool isvalidheight(std::string x) {
  std::string ty = x.substr(x.length() - 2, x.length());
  std::string val = x.substr(0, x.length() - 2);
  if (ty != "cm" && ty != "in") {
    return false;
  };
  if (ty == "cm") {
    return inrange(std::stoi(val), 150, 193);
  } else {
    return inrange(std::stoi(val), 59, 76);
  };
};
bool validHairColor(std::string x) {
  std::string vals = x.substr(1, x.length());
  if (x[0] != '#' || vals.length() != 6) {
    return false;
  };
  if (std::all_of(vals.begin(), vals.end(), ::isxdigit)) {
    return true;
  };
  return false;
};

/**
 * A better way would be to pass requirements as functions (i.e. pass the
 * functions that return the valid conditions as arguments) Then just loop
 * through each argument and run each function and update the valid accordingly
 */
 template<typename... Args> bool f(Args &&...args);
struct funco {
  bool valid = true;
 bool checker(std::forward<f>(args)...) {
    valid = valid && args("test");
    return valid;
  };
};
bool is_number(const std::string &s) {
  return !s.empty() && std::find_if(s.begin(), s.end(), [](unsigned char c) {
                         return !std::isdigit(c);
                       }) == s.end();
}

/**
 * This is extremely gross
 */
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
        // Not 4digits
        if (!isfourDigits(i[k])) {
          valid--;
          break;
        } else {
          // Not In range
          if (!inrange(std::stoi(i[k]), 1920, 2002)) {
            valid--;
            break;
          };
        }
      } else if (k == "eyr") {
        // Not 4 Dig
        if (!isfourDigits(i[k])) {
          valid--;
          break;
        } else {
          // Not in range
          if (!inrange(std::stoi(i[k]), 2020, 2030)) {
            valid--;
            break;
          };
        }
      } else if (k == "iyr") {
        // Not four dig
        if (!isfourDigits(i[k])) {
          valid--;
          break;
        } else {
          // Not in range
          if (!inrange(std::stoi(i[k]), 2010, 2020)) {
            valid--;
            break;
          };
        }
      } else if (k == "hgt") {
        // Not valid
        if (!isvalidheight(i[k])) {
          valid--;
          break;
        };
      } else if (k == "hcl") {
        // Not valid
        if (!validHairColor(i[k])) {
          valid--;
          break;
        }
      } else if (k == "ecl") {
        // Not in et
        std::unordered_set<std::string> types = {"amb", "blu", "brn", "gry",
                                                 "grn", "hzl", "oth"};
        if (types.find(i[k]) == types.end()) {
          valid--;
          break;
        }
      } else {
        // Not number or its length != 9
        if (!is_number(i[k]) || i[k].length() != 9) {
          valid--;
          break;
        }
      };
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
