/**
 * Day ol - 2
 * jarviliam
 */

#include <array>
#include <fstream>
#include <iostream>
#include <vector>

void buildStrings(std::vector<std::string> &output) {
  std::ifstream file("./input.txt");
  if (file.is_open()) {
    std::string entry;
    while (std::getline(file, entry)) {
      output.push_back(entry);
    };
  };
};

int part2(std::vector<std::string> inputs) {
  int valid = 0;
  for (auto i : inputs) {
    // Retrive Minimum Val
    int delimPos = i.find("-");
    int minVal = std::stoi(i.substr(0, delimPos));
    i.erase(0, delimPos + 1);

    // Retrieve Max
    delimPos = i.find(" ");
    int maxVal = std::stoi(i.substr(0, delimPos));
    i.erase(0, delimPos + 1);

    std::cout << i << std::endl;

    // Retrieve Pass
    delimPos = i.find(":");
    char val = i[delimPos - 1];
    i.erase(0, delimPos + 1);

    //Big ol XOR
    if ((i[minVal] == val) ^ (i[maxVal] == val)) {
      valid++;
    }
  }
  return valid;
};

int main() {
  std::vector<std::string> inputs;
  buildStrings(inputs);

  int valid = 0;
  // Loop through each entry
  for (auto i : inputs) {

    // Format: 4-5 m : mmmmm
    // Retrive Minimum Val
    int delimPos = i.find("-");
    int minVal = std::stoi(i.substr(0, delimPos));
    i.erase(0, delimPos + 1);

    // Retrieve Max
    delimPos = i.find(" ");
    int maxVal = std::stoi(i.substr(0, delimPos));
    i.erase(0, delimPos + 1);

    // Retrieve Pass
    delimPos = i.find(":");
    char val = i[delimPos - 1];
    i.erase(0, delimPos + 1);

    int cnt = 0;
    for (auto ch : i) {
      if (ch == val) {
        cnt++;
      };
    };
    if (cnt >= minVal && cnt <= maxVal) {
      valid++;
    };
  };
  std::cout << "Answers : " << valid << std::endl;
  std::cout << "Input Size : " << inputs.size() << std::endl;

  std:: cout << part2(inputs) << std::endl;
};
