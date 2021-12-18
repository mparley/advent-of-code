#include <iostream>
#include <vector>
#include <string>

using namespace std;

vector<int> ParseInput() {
  vector<int> out;
  string s;
  cin >> s >> s;

  while (cin >> s) {
    int start = s.find("=");
    int mid = s.find("..");
    out.push_back(stoi(s.substr(start+1,mid-start)));
    out.push_back(stoi(s.substr(mid+2)));
  }

  return out;
}

struct Velocity {
  int x, y;
  Velocity(int x, int y) : x(x), y(y) {}
};

bool TestVelocity(Velocity v, const vector<int> &input) {
  for (int x = 0, y = 0; x <= input[1] && y >= input[2];) {
    x += v.x;
    y += v.y;
    if (v.x > 0) v.x--;
    v.y--;

    if ((x >= input[0] && x <= input[1]) && (y >= input[2] && y <= input[3]))
      return true;
  }

  return false;
}

int main() {
  vector<int> input = ParseInput();
  Velocity v(0,0);

  for (int sumation = 0; sumation < input[0];) {
    v.x++;
    sumation = (v.x*(v.x+1))/2;
  }

  v.y = (input[2]+1) * -1;
  int peak = (v.y * (v.y+1))/2;

  cout << "vel: " << v.x << "," << v.y 
    << "\npeak: " << peak << "\n";

  Velocity min(v.x,input[2]), max(input[1],v.y);

  int total = 0;
  for (int x = min.x; x <= max.x; x++)
    for (int y = min.y; y <= max.y; y++)
      total += TestVelocity(Velocity(x, y), input);

  cout << "total: " << total << "\n";

  return EXIT_SUCCESS;
}