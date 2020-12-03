#include <string>
#include <fstream>
#include <iostream>
#include <vector>

using namespace std;

// Counts the trees on the path, wraps the x value
int treesInPath(vector<string> lines, int dx, int dy) {
  int x = 0, y = 0;
  int tree_count = 0;
  int len = lines[0].length();

  while (y < lines.size()) {
    if (lines[y][x] == '#')
      tree_count++;

    x = (x + dx) % len;
    y += dy;
  }

  return tree_count;
}

// The filename of the input is the first argument, the following args are
// for pairs of steps for the slope - right then down, for as many slopes
// you wanna check
int main(int argc, char* argv[]) {
  if (argc < 2)
    return -1;

  ifstream infile(argv[1]);
  string line;
  vector<string> lines;
  int tree_product = 1;

  while (getline(infile, line)) {
    lines.push_back(line);
  }

  for (int i = 2; i < argc - 1; i += 2)
    tree_product *= treesInPath(lines, atoi(argv[i]), atoi(argv[i+1]));

  cout << tree_product << '\n';

  return 0;
}