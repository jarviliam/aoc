/**
 *
 * Extremely Gross.
 * Should have just used Binary
 * As the Question suggested
 */

#include <fstream>
#include <iostream>
#include <math.h>
#include <unordered_set>
#include <vector>

void loadIt(std::vector<std::string> &input) {
  std::ifstream file("./input.txt");
  if (file.is_open()) {
    std::string entry;
    while (std::getline(file, entry)) {
      if (entry == "")
        continue;
      input.push_back(entry);
    };
  };
}

struct calcu {
  int high = 127;
  int low = 0;
  int answer = 0;
  int colHigh = 7;
  int colLow = 0;
  int colAnswer = 0;

  void calculateRow(char type) {
    if (high == low)
      return;
    if (type == 'F') {
      if (high - low == 1) {
        answer = low;
      } else {
        if (high % 2 != 0)
          high++;
        high -= ((high - low) / 2) + 1;
      };
    } else {
      if (high - low == 1) {
        answer = high;
      } else {
        low += ceil((high - low) / 2) + 1;
      };
    }
  };
  void calculateColumn(char type) {
    if (colHigh == colLow) {
      return;
    };
    if (type == 'L') {
      if (colHigh - colLow == 1) {
        colAnswer = colLow;
      } else {
        if (high % 2 != 0)
          high++;
        colHigh -= ((colHigh - colLow) / 2) + 1;
      }
    } else {
      if (colHigh - colLow == 1) {
        colAnswer = colHigh;
      } else {
        colLow += ceil((colHigh - colLow) / 2) + 1;
      }
    }
  };

  int operator()(std::string line) {
    std::string::iterator itr;
    for (itr = line.begin(); itr < line.end(); itr++) {
      if (*itr == 'F' || *itr == 'B') {
        calculateRow(*itr);
      } else {
        calculateColumn(*itr);
      };
    };
    return answer * 8 + colAnswer;
  }
};

int main() {
  std::vector<std::string> lanes;
  std::unordered_set<int> seats;

  std::string example = "FBFBBFFRLR";
  loadIt(lanes);
  int highest = 0;
  for (auto i : lanes) {
    int ans = calcu()(i);
    seats.emplace(ans);
    if (ans > highest) {
      highest = ans;
    }
  }
  std::cout << "Highest : " << highest << std::endl;
  for (int i = 0; i < 127; i++) {
    for (int j = 0; j < 8; j++) {
      const int id = i * 8 + j;
      if ((seats.find(id + 1) != seats.end()) &&
          (seats.find(id - 1) != seats.end()) &&
          (seats.find(id) == seats.end())) {
        std::cout << "MY SEAT: " << id << std::endl;
      }
    }
  }
};
