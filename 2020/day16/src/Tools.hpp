#ifndef TOOLS_H
#define TOOLS_H

#include <iostream>
#include <regex>
#include <string>
#include <fstream>
#include <sstream>
#include <vector>
#include <unordered_map>

using namespace std;

struct Range {
    int max;
    int min;

    Range(int a, int b) { max = b; min = a; };
};

enum { RULES, OWN, NEARBY, SET };

vector<int> readTicket(string line);
void readRule(unordered_map<string, vector<Range>>& rules, string line);
void readFile(unordered_map<string, vector<Range>>& rules,
    vector<vector<int>>& tickets, string filename);


#endif
