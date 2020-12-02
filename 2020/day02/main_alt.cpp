// Alternate, more optimized version of my solution
// No more unneccesary hash maps or structs
#include <iostream>
#include <sstream>
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

bool isValidPart2(string password, char req, int first, int second) {
  return (password[first-1] == req) != (password[second-1] == req);
}

// This reads the input and counts each part's interpretation of valid passes
// This time it does the checking while reading input
void parseInputsAndCount(int &c1, int &c2) {
  string temp;
  int f, s;
  string pass;
  char req;

  while (getline(cin, temp)) {
    stringstream ss(temp);

    ss >> f;
    ss.get();
    ss >> s >> req;
    ss.get();
    ss >> pass;

    if (isValidPart1(pass, req, f, s)) c1++;
    if (isValidPart2(pass, req, f, s)) c2++;
  }
}

int main(int argc, char *argv[]) {
  int count1 = 0, count2 = 0;
  parseInputsAndCount(count1, count2);
  
  cout << count1 << '\n' << count2 << '\n';

  return 0;
}