#include "algorithm"
#include "fstream"
#include "iostream"
#include "string"
#include "vector"

int part1(std::vector<std::string> inp) {
  int ans = 0;
  while (true) {
    // Deep Copy as we have to change lines
    std::vector<std::string> copy(inp);
    // For escaping loop
    bool change = false;
    // Col and Row size
    int cSize = inp[0].size();
    int rSize = inp.size();
    // map rows
    for (int r = 0; r != inp.size(); r++) {
      // map Columns
      for (int c = 0; c != inp[r].size(); c++) {
        int notOcc = 0;
        // Need to Check Cardinal Directions
        // -> -1 0 1 COL -1 0 1 ROW
        for (int i = -1; i <= 1; i++) {
          for (int j = -1; j <= 1; j++) {
            // If At Pos
            if (i == 0 && j == 0) {
              continue;
            }
            // Get Row Val and C Val
            int rr = r + i;
            int cc = c + j;
            // As long as its within the vector + string size
            // and is '#' increment notOcc
            if (0 <= rr && rr < rSize && 0 <= cc && cc < cSize &&
                inp[rr][cc] == '#') {
              notOcc += 1;
            };
          };
        };
        if (inp[r][c] == 'L') {
          // If there is no occupied around
          if (notOcc == 0) {
            // Change
            copy[r][c] = '#';
            // Notify Changed
            change = true;
          }
        } else if (inp[r][c] == '#' && notOcc >= 4) {
          // Lock Seat
          copy[r][c] = 'L';
          change = true;
        }
      };
    };
    // If there were no changes, It's donzo
    if (change == false) {
      break;
    }
    // For next iteration, replace previous
    inp = copy;
  };

  // Count Answer
  for (int i = 0; i != inp.size(); i++) {
      std::cout<< inp[i]<<std::endl;
    for (int j = 0; j != inp.size(); j++) {
      if (inp[i][j] == '#') {
        ans += 1;
      }
    }
  }
  return ans;
}

std::vector<std::string> inputs() {
  std::ifstream file("./input.txt");
  std::vector<std::string> input;
  if (file.is_open()) {
    std::string entry;
    while (std::getline(file, entry)) {
      input.emplace_back(entry);
    };
  };
  return input;
};

int main() {
  auto inp = inputs();
  std::cout << part1(inp) << "\n";
}
