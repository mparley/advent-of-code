#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <climits>
#include <unordered_map>

using namespace std;

string parseInput(string fname, unordered_map<string,char>& rules) {
    string polymer;
    ifstream infile(fname);

    getline(infile,polymer);

    while (infile.good()) {
        string line;
        getline(infile,line);

        rules[line.substr(0,2)] = line[line.size()-1];
    }
    infile.close();
    return polymer;
}

int main(int argc, char *argv[]) {
    unordered_map<string, uint64_t> pairs;
    unordered_map<char, uint64_t> counts;
    unordered_map<string, char> rules;

    string polymer = parseInput(argv[1], rules);
    int num_its = (argc > 2) ? stoi(argv[2]) : 10;


    // Separate the template into pairs map and count the elements
    char prev = polymer[0];
    counts[polymer[0]]++;

    for (int i = 1; i < polymer.size(); i++) {
        string pr(2, prev);
        pr[1] = polymer[i];
        pairs[pr]++;
        counts[polymer[i]]++;
        prev = polymer[i];
    }

    // Each iteration go through pairs map and apply rules to turn
    // one pair into two new pairs. Count the added character
    for (int i = 0; i < num_its; i++) {
        unordered_map<string, uint64_t> nMap;
        for (auto pr : pairs) {
            if (pr.second > 0) {
                for (int j = 0; j < 2; j++) {
                    string nPair = pr.first;
                    nPair[j] = rules[pr.first];
                    nMap[nPair] += pr.second;
                }
                counts[rules[pr.first]] += pr.second;
            }
        }
        pairs = nMap;
    }

    // Find most common and least common and print solution
    uint64_t most = 0, least = UINT64_MAX;
    for (auto e : counts) {
        if (most < e.second) most = e.second;
        if (least > e.second) least = e.second;
    }
    cout << most << " - " << least << " = " << most - least << endl;


    return EXIT_SUCCESS;
}
