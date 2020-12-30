#include <fstream>
#include <iostream>
#include <set>
#include <tuple>
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

std::tuple<std::string, int> parse(const std::string x) {
  int space = x.find(" ");
  std::string inst = x.substr(0, space);
  int val = std::stoi(x.substr(space + 1, x.length()));
  return {inst, val};
};

int main() {
  std::vector<std::string> instructions;
  loadIt(instructions);

  int accum = 0;
  std::unordered_set<int> jmpList = {};
  std::vector<int> instList = {};

  std::vector<std::string>::iterator itr;

  for (itr = instructions.begin(); itr <= instructions.end(); itr++) {

    // C++17 Thank you.
    auto [inst, val] = parse((*itr));

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
      instList.emplace_back(itr - instructions.begin());
      jmpList.emplace(itr - instructions.begin());
      std::advance(itr, val - 1);
      continue;
    } else {
      if ((instructions.end() - itr) == val - 1) {
        std::cout << "Should Jump from Nop " << std::endl;
      }
    }
    jmpList.emplace(itr - instructions.begin());
    instList.emplace_back(itr - instructions.begin());
  }

  std::vector<int>::reverse_iterator rIt;
  // TODO(jarviliam): Reverse Loop -> If Jump -> Attempt to finish by changing
  // to Nop / Vice Versa
  for (rIt = instList.rbegin(); rIt != instList.rend(); rIt++) {
    auto [inst, val] = parse(instructions[(*rIt)]);
    std::cout << instructions[*rIt] << std::endl;
    if (inst == "jmp") {
      std::cout << "Reverse Jump " << (*rIt) << std::endl;
    };
  };
};
