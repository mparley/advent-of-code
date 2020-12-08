#include "Tools.h"

// Reads the file and populates the vector of instructions
void readFile(std::string filename, std::vector<Instruction>& code_lines) {
  std::ifstream infile(filename); 

  // This is just for converting string to operation code (int)
  std::unordered_map<std::string,int> opsMap = {
    {"acc", Ops::ACC}, {"jmp", Ops::JMP}, {"nop", Ops::NOP}
  };

  std::string op;
  int val;

  while (infile >> op >> val) {
    Instruction line = { opsMap.at(op), val };
    code_lines.push_back(line);
  }
  infile.close();
}

// Helper for printing current line while running instructions in BootCheck
// uncomment out in there to see it
void printCodeLine(int line, int acc, const Instruction& in) {
  std::cout << "line: " << line << "\tacc: " << acc << "\top: " << in.operation
    << "\tval: " << in.value << "\n";
}

// Helper to switch JMP to NOP and vice versa
void switchOp(Instruction& in) {
  if (in.operation == Ops::ACC)
    return;
  in.operation = in.operation == Ops::JMP ? Ops::NOP : Ops::JMP;
}
