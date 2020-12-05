#include <iostream>
#include <string>
#include <fstream>
#include <algorithm>
#include <vector>

using namespace std;

// Recursively get row or column based on string of binary directions
int BinaryGet(int low, int high, string seat_string) {
  if (seat_string.empty())
    return high;

  if (seat_string[0] == 'B' || seat_string[0] == 'R') {
    int new_low = ((high - low) / 2) + low;
    return BinaryGet(new_low, high, seat_string.substr(1));
  }

  else if (seat_string[0] == 'F' || seat_string[0] == 'L') {
    int new_high = ((high - low) / 2) + low;
    return BinaryGet(low, new_high, seat_string.substr(1));
  }
}

// Input passed as first argument
int main(int argc, char* argv[]) {
  if (argc != 2)
    return EXIT_FAILURE;

  ifstream infile(argv[1]);
  string line;
  vector<int> ids;
  int prev_id = 0;

  // Read file and get rows, cols and ids
  while (getline(infile, line)) {
    int row = BinaryGet(0, 127, line.substr(0, 7));
    int col = BinaryGet(0, 7, line.substr(7));
    int id = (row * 8) + col;

    ids.push_back(id);
  }

  // Sort IDs to check for missing seat
  sort(ids.begin(), ids.end());
  prev_id = ids[0];
  for (auto id : ids) {
    if (id - prev_id == 2)
      break;
    prev_id = id;
  }

  cout << "Highest ID: " << *(ids.end()-1) << "\nMy seat ID: " << prev_id + 1 << "\n";

  return EXIT_SUCCESS;
}