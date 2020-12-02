// Same as the other alternate but getting data from file rather than
// standard input. Did it for funsies and to see how much faster it is
// on the big input
#include <iostream>
#include <sstream>
#include <fstream>
#include <string>

using namespace std;

// Counts character and checks if the required char is within the range
bool isValidPart1(string password, char req, int first, int second) {
  int min = first, max = second;
  int count = 0;

  for (char c : password) {
    if (c == req)
      count++;
  }

  return count <= max && count >= min;
}

// This is just xor
bool isValidPart2(string password, char req, int first, int second) {
  return !(password[first-1] == req) != !(password[second-1] == req);
}

// This reads the file and counts each part's interpretation of valid passes
void parseInputsAndCount(int &c1, int &c2, string fname) {
  string temp;
  int f, s;
  string pass;
  char req;
  ifstream in;
  in.open(fname);

  while (in.good()) {

    in >> f;
    in.get();
    in >> s >> req;
    in.get();
    in >> pass;

    if (isValidPart1(pass, req, f, s)) c1++;
    if (isValidPart2(pass, req, f, s)) c2++;
  }

  in.close();
}

int main(int argc, char *argv[]) {
  int count1 = 0, count2 = 0;
  parseInputsAndCount(count1, count2, argv[1]);
  
  cout << count1 << '\n' << count2 << '\n';

  return 0;
}