#include <iostream>
#include <unordered_set>
#include <vector>
#include "Tools.h"
#include "BootCheck.h"

using namespace std;

// Filename passed as first argument
int main(int argc, char *argv[]) {
  if (argc < 2)
      return EXIT_FAILURE;

  vector<Instruction> code_lines;
  unordered_set<int> history;

  readFile(argv[1], code_lines);

  cout << accBeforeLoop(history, code_lines) << '\n';
  cout << fixCode(history, code_lines) << '\n';

  return EXIT_SUCCESS;
}
