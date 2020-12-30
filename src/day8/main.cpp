#include <fstream>
#include <iostream>
#include <unordered_set>
#include <vector>

void loadIt(std::vector<std::string> &input) {
  std::ifstream file("./input.txt");
  if (file.is_open()) {
    std::string entry;
    while (std::getline(file, entry)) {
      input.emplace_back(entry);
    };
  };
};
int main() {
  std::vector<std::string> instructions;
  loadIt(instructions);

  int accum = 0;
  std::unordered_set<int> jmpList = {};

  std::vector<std::string>::iterator itr;

  for (itr = instructions.begin(); itr <= instructions.end(); itr++) {

    int space = (*itr).find(" ");
    std::string inst = (*itr).substr(0, space);
    int val = std::stoi((*itr).substr(space + 1, (*itr).length()));
    std::cout << inst << std::endl;

    if (jmpList.find(itr - instructions.begin()) != jmpList.end()) {
      // This instruction has already be ran
      std::cout << "Found " << accum << std::endl;
      break;
    }
    if (inst == "acc") {
      accum += val;
    } else if (inst == "jmp") {
      std::cout << "Jumping " << accum << " " << val << std::endl;
      // Emplace Before Advancing Itr
      jmpList.emplace(itr - instructions.begin());
      std::advance(itr, val - 1);
      continue;
    }
    jmpList.emplace(itr - instructions.begin());
  }
};
