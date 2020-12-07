#ifndef TOOLS_H
#define TOOLS_H

#include <string>
#include <regex>
#include <unordered_map>
#include <vector>
#include <algorithm>
#include <iostream>

using namespace std;

bool canContain(string target, string bag,
  unordered_map<string, bool>& checked,
  const unordered_map<string, unordered_map<string,int>*>& rules);

void printHelper(unordered_map<string,bool>& checked, 
  unordered_map<string, unordered_map<string,int>*>& rules,
  string bag);

int containsHelper(string target,
  unordered_map<string, unordered_map<string,int>*>& rules);

void parseLine(string line, unordered_map<string,
  unordered_map<string,int>*>& results);

void printRules(const unordered_map<string,
  unordered_map<string,int>*>& rules);

int totalBags(string bag,
  const unordered_map<string, unordered_map<string,int>*>& rules);

#endif
