#ifndef BOOT_CHECK_H
#define BOOT_CHECK_H

#include <unordered_set>
#include <vector>
#include "Tools.h"

int accBeforeLoop(std::unordered_set<int>& history,
  const std::vector<Instruction>& code_lines);

int fixCode(const std::unordered_set<int>& history,
  std::vector<Instruction>& code_lines);

#endif
