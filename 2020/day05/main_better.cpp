// When I solved the problem initially, I missed the ABSOLUTELY
// OBVIOUS fact that the seat inputs were LITERALLY JUST BINARY
// I'm actually really embarassed I somehow missed this
// This is the better solution
#include <iostream>
#include <string>
#include <fstream>
#include <algorithm>
#include <vector>

using namespace std;

// Goes through the input and converts it to an int via bitshifting
int getSeat(const string& line) {
  int id = 0;

  for (int i = 0; i < line.length(); i++)
    if (line[i] == 'B' || line[i] == 'R')
      id |= 1 << (line.length() - i - 1);

  return id;
}

// Input passed as first argument
int main(int argc, char* argv[]) {
  if (argc != 2)
    return EXIT_FAILURE;

  ifstream infile(argv[1]);
  string line;
  vector<int> ids;
  int prev_id = 0;

  // Read file and get ids 
  while (getline(infile, line)) {
    int id = getSeat(line);
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
