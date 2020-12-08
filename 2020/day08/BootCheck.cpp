#include "BootCheck.h"

// Returns the accumulator before the loop is detected and populates history
// so we can use it in fix code. If no loop is found it will return final acc
int accBeforeLoop(std::unordered_set<int>& history,
  const std::vector<Instruction>& code_lines)
{
  int accumulator = 0;
  for (int i=0, off=1; i < code_lines.size() && i >= 0; i += off) {
    //printCodeLine(i, accumulator, code_lines[i]);
    
    // Check history for loop
    if (history.find(i) != history.end())
      break;

    switch (code_lines[i].operation) {

      // ACC will add to the accumulator and set the offset to 1 for next line
      case Ops::ACC:
        accumulator += code_lines[i].value;
        off = 1;
        break;

      // JMP just sets the offset so next instruction will be +/- it's value 
      case Ops::JMP:
        off = code_lines[i].value;
        break;

      // NOP or anything else will reset the offset so it goes to next line
      default:
        off = 1;
        break;
    }

    history.insert(i);
  }

  return accumulator;
}

// Goes through every JMP and NOP in history and tries a run with those
// operations switched. If a run is found with the last instruction in
// it's history, we've fixed it so the accumulator is returned. If nothing is
// fixed return -1
int fixCode(const std::unordered_set<int>& history,
  std::vector<Instruction>& code_lines) 
{
  int accumulator = 0;
  for (auto line : history) {
    
    // Skip ACC operations
    if (code_lines[line].operation == Ops::ACC)
      continue;

    // Creates an empty "alt history", switches the operation, and runs it
    std::unordered_set<int> alt_history;
    switchOp(code_lines[line]);
    accumulator = accBeforeLoop(alt_history, code_lines);

    // If last line found in history, it finished instructions
    if (alt_history.find(code_lines.size()-1) != alt_history.end())
      return accumulator;

    // Reset for next run
    switchOp(code_lines[line]);
  }

  return -1;
}
