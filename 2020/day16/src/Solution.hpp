#ifndef SOLUTION_H
#define SOLUTION_H

#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <string>
#include <algorithm>
#include <boost/multiprecision/cpp_int.hpp>
#include "Tools.hpp"

using namespace std;
using boost::multiprecision::cpp_int;

bool valid(const vector<Range>& rule, int val);

int findErrorRate(const unordered_map<string,vector<Range>>& rules,
    vector<vector<int>>& tickets);

cpp_int departureProduct(const unordered_map<string, vector<Range>>& rules,
    const vector<vector<int>>& tickets);


#endif
