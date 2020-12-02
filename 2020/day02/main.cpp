// This is the version I solved the problems with
// After I realized steps I took during part 1 in preparation for a more complex
// part 2 that never happened, I made a more optimized version in main_alt.cpp
// where I don't bother store the data before handling it.
#include <iostream>
#include <sstream>
#include <unordered_map>
#include <vector>
#include <string>

using namespace std;

struct PassEntry {
  int first;
  int second;
  char req;
  string password;
};

// Counts each character and checks if the required char is within the range
bool isValidPart1(PassEntry p) {
  unordered_map<char, int> char_counts;
  int min = p.first, max = p.second;

  // Lmao I think I thought I might have had to track multiple chars in part 2
  // Should have just used a single int to count occurances. See main_alt.cpp
  for (char c : p.password) {
    char_counts[c]++;
  }

  return char_counts[p.req] <= max && char_counts[p.req] >= min;
}

// This is just xor
bool isValidPart2(PassEntry p) {
  return !(p.password[p.first-1] == p.req) != !(p.password[p.second-1] == p.req);
}

// This really isn't needed - you can handle all the logic while reading
// the data. I just expected it might be useful for part 2 but it wasn't.
vector<PassEntry> parseInputs() {
  string temp;
  vector<PassEntry> pass_list;

  while (getline(cin, temp)) {
    PassEntry p;
    stringstream ss(temp);

    ss >> p.first;
    ss.get();
    ss >> p.second >> p.req;
    ss.get();
    ss >> p.password;

    pass_list.push_back(p);
  }

  return pass_list;
}

int main(int argc, char *argv[]) {
  vector<PassEntry> pass_list = parseInputs();
  int count1 = 0, count2 = 0;

  // Could have done while reading input, oh well
  for (auto p : pass_list) {
    if (isValidPart1(p))
      count1++;
    
    if (isValidPart2(p))
      count2++;
  }
  
  cout << count1 << '\n' << count2 << '\n';

  return 0;
}