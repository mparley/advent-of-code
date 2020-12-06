#include <iostream>
#include <string>
#include <bitset>
#include <fstream>

using namespace std;

// Filename is passed in as first argument
int main(int argc, char* argv[]) {
  if (argc < 2)
    return EXIT_FAILURE;

  ifstream infile(argv[1]);
  string line;
  bitset<26> bs1, bs2, curr; // Using bitsets for alphabet
  int count1 = 0, count2 = 0;

  bs2.set();

  // Reads file and operates on bitsets
  // Part 1 bitset needs to OR with line entries
  // Part 2 bitset needs to AND with line entries
  while (infile.good()) {
    getline(infile, line);
    curr.reset();

    // On empty line or end of file, update counts
    // reset bits.
    if (line.empty() || !infile.good()) {
      count1 += bs1.count();
      count2 += bs2.count();
      bs1.reset();
      bs2.set();

    // Read characters, and bitwise operations
    } else {
      for (char c : line)
        curr.set(c - 'a', 1);

      bs1 |= curr;
      bs2 &= curr;
    }
  }

  printf("%d %d\n", count1, count2);

  return EXIT_SUCCESS;
}
