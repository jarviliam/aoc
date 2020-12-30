#include <algorithm>
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

struct Node {
  bool counted;
  std::string type;
  std::vector<Node> children;
  Node(std::string type, bool counted) : type(type),counted(counted) {}
};

// Type -> List Children
// X contain Y z bags, etc...
// Check no other bags
//
// Split at contain , number , bag
int main() {
  std::vector<std::string> instructions;
  std::unordered_map<std::string, Node> nodeList;
  loadIt(instructions);

  for (auto &i : instructions) {
    auto itr = i.find("contain");
    std::string key = i.substr(0, itr);
    i.erase(0, itr + 8);
    auto no = i.find("no");
    std::cout << i << std::endl;
    if (no != std::string::npos) {
      Node test = Node(key,false);
      nodeList.emplace(key, test);
    } else {
        while(i.find(",") == std::string::npos || i.find(".") == std::string::npos){
            int amount = std::stoi(i.substr(0,1));   
            i.erase(0,2);
            auto delim = i.find("bag");
            std::string type = i.substr(0,delim-1);
            i.erase(0,delim);
            auto commaDelim = i.find(",");
            if(amount > 1){
                //i.erase(0,
            }
        };
    }
  };
  for(auto i : nodeList){
      std::cout << i.first << std::endl;
  }
};
