#ifndef TOOLS_H
#define TOOLS_H

#include <fstream>
#include <unordered_map>
#include <iostream>
#include <string>
#include <vector>

enum Ops { ACC, JMP, NOP };

struct Instruction {
  int operation;
  int value;
};

void readFile(std::string filename, std::vector<Instruction>& code_lines);

void printCodeLine(int line, int acc, const Instruction& in);

void switchOp(Instruction& in);

#endif
