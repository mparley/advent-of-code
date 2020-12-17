#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include "Tools.hpp"
#include "Solution.hpp"

using namespace std;

int main(int argc, char *argv[]) {
    if (argc < 2) return EXIT_FAILURE;

    unordered_map<string, vector<Range>> rules;
    vector<vector<int>> tickets;
    readFile(rules, tickets, argv[1]);

    cout << findErrorRate(rules, tickets) << endl;
    cout << departureProduct(rules, tickets) << endl;

    return 0;
}
