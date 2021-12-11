#include <iostream>
#include <fstream>
#include <string>
#include <stack>
#include <vector>
#include <unordered_map>
#include <algorithm>

using namespace std;

const unordered_map<char,char> MATCHES = {{'(',')'}, {'[',']'}, {'{','}'}, {'<','>'}};
const unordered_map<char,int> ESCORES = {{')',3}, {']',57}, {'}',1197}, {'>',25137}};
const unordered_map<char,int> CSCORES = {{')',1}, {']',2}, {'}',3}, {'>',4}};

int main(int argc, char** argv) {

  unsigned long err_score = 0;
  vector<unsigned long> comp_scores;
  ifstream infile(argv[1]);

  while (infile.good()) {
    stack<char> s;
    string line;
    bool error = false;
    getline(infile,line);

    for (int i = 0; i < line.size(); i++) {
      if (MATCHES.find(line[i]) != MATCHES.end()) {
        s.push(line[i]);
      } else if (line[i] == MATCHES.at(s.top())) {
        s.pop();
      } else {
        error = true;
        err_score += ESCORES.at(line[i]);
        break;
      }
    }

    if (error) continue;

    line.clear();
    while (!s.empty()) {
      line += MATCHES.at(s.top());
      s.pop();
    }

    unsigned long score = 0;
    for (char c : line) {
      score *= 5;
      score += CSCORES.at(c);
    }

    comp_scores.push_back(score);
  }
  infile.close();

  cout << err_score << "\n";

  sort(comp_scores.begin(),comp_scores.end());
  cout << comp_scores[comp_scores.size()/2] << "\n";

  return EXIT_SUCCESS;
}