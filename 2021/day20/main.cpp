#include <iostream>
#include <fstream>
#include <vector>
#include <bitset>

using namespace std;

bool GetBool(char c) { return (c == '#') ? 1 : 0; }

char GetChar(bool b) { return (b) ? '#' : ' '; }

struct InfiniteImage {
  vector<vector<bool>> actual;
  bool inf_bit;

  InfiniteImage(vector<vector<bool>> v) 
    : actual(v) {}

  InfiniteImage(int width, int height)
    : actual(vector<vector<bool>>( height, vector<bool>(width) )) {}

  int height() const { return actual.size(); }
  int width() const { return actual[0].size(); }

  bool at(int x, int y) const {
    if (x < width() && x >= 0 && y < height() && y >= 0)
      return actual[y][x];
    else 
      return inf_bit;
  }

  void set(int x, int y, bool b) {
    actual[y][x] = b;
  }

  int e_index(int j, int i) const {
    int index = 0;
    for (int k = 0; k < 9; k++) {
      int x = j - 1 + (k % 3);
      int y = i - 1 + (k / 3);
      index |= at(x,y) << (8-k);
    }
    return index;
  }

  int count() const {
    int sum = 0;
    for (auto row : actual)
      for (auto col : row)
        if (col) sum++;
    return sum;
  }
};

ostream& operator<<(ostream& out, const InfiniteImage& ii) {
  for (auto row : ii.actual) {
    for (auto col : row)
      out << GetChar(col);
    cout << "\n";
  }
  return out;
}

InfiniteImage Enhance(string algo, const InfiniteImage& input) {
  InfiniteImage output(input.width()+2, input.height()+2);

  for (int i = -1; i <= input.height(); i++) {
    for (int j = -1; j <= input.width(); j++) {
      int index = input.e_index(j, i);
      output.set(j+1, i+1, GetBool(algo[index]));
    }
  }

  output.inf_bit = GetBool(algo[input.e_index(-10,-10)]);
  return output;
}

int main(int argc, char** argv) {
  string line, algo;
  int iterations = (argc > 2) ? stoi(argv[2]) : 2;

  ifstream infile(argv[1]);
  getline(infile, algo);

  vector<vector<bool>> input_image;
  while (getline(infile,line)) {
    if (line.empty()) continue;
    vector<bool> row;
    for (char c : line)
      row.push_back(GetBool(c));
    input_image.push_back(row);
  }

  InfiniteImage image(input_image);

  for (int i = 0; i < iterations; i++)
    image = Enhance(algo, image);

  cout << image << "\n";
  cout << image.count() << "\n";

  return EXIT_SUCCESS;
}