#include <vector>
#include <iostream>
#include <fstream>


int part2(){
    int ans = 0;
    return ans;
};

void loadIt(std::vector<std::string>& input){
  std::ifstream file("./input.txt");
  if (file.is_open()) {
    std::string entry;
    while (std::getline(file, entry)) {
      input.push_back(entry);
    };
  };
}

int main () {
    std::vector<std::string> lanes;
    loadIt(lanes);

    for(auto i : lanes){
        std::cout << i << std::endl;
    }
};
