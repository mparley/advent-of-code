#include <unordered_map>
#include <iostream>
#include <fstream>
#include <string>
#include "Tools.h"

using namespace std;

// Filename is first argument
int main(int argc, char* argv[]) {
  if (argc < 2)
    return EXIT_FAILURE;

  ifstream infile(argv[1]);
  string line;
  unordered_map<string, unordered_map<string,int>*> rules;

  while (getline(infile, line)) {
    parseLine(line, rules);
  }

  cout << containsHelper("shiny gold", rules) << '\n';
  cout << totalBags("shiny gold", rules) - 1 << '\n';

  // Clean up
  for (auto it = rules.begin(); it != rules.end(); it++) {
    delete it->second;
  }
  rules.clear();

  return EXIT_SUCCESS;
}
