/**
 * I really didnt read the question fully the first time
 * and it cost me...
 * how much did it cost me?
 * about 25mins
 * -Liam Jarvis
 */

#include <fstream>
#include <iostream>
#include <vector>

void loadIt(std::vector<std::string> &input) {
  std::ifstream file("./input.txt");
  if (file.is_open()) {
    std::string entry;
    while (std::getline(file, entry)) {
      input.push_back(entry);
    };
  };
}

typedef std::vector<std::string> weee;

int howManyTreesDidThisGuyHit(weee lanes, int r, int d) {
  int length = lanes[0].size();
  int height = lanes.size();
  int x = 0;
  int y = 0;
  int treeSmashes = 0;
  while (y < height - d) {
    x += r;
    if (x >= length) {
      x = x % length;
    };
    y += d;
    if (lanes[y][x] == '#') {
      treeSmashes++;
    }
  };
  return treeSmashes;
};

int main() {
  std::vector<std::string> lanes;
  loadIt(lanes);
  for (auto i : lanes) {
    std::cout << i << std::endl;
  }
  int getMeOuttaHere = howManyTreesDidThisGuyHit(lanes, 1, 1);
  int b = howManyTreesDidThisGuyHit(lanes, 3, 1);
  int c = howManyTreesDidThisGuyHit(lanes, 5, 1);
  int d = howManyTreesDidThisGuyHit(lanes, 7, 1);
  int e = howManyTreesDidThisGuyHit(lanes, 1, 2);
  std::cout << "Trees : " << getMeOuttaHere << std::endl;
  std::cout << "Trees : " << b << std::endl;
  std::cout << "Trees : " << c << std::endl;
  std::cout << "Trees : " << d << std::endl;
  std::cout << "Trees : " << e << std::endl;
};
