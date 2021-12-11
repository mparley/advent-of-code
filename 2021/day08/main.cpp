#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <vector>
#include <unordered_map>
#include <bitset>

using namespace std;

//   0***:   1*:     2***:   3**:    4*:
//  aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
//  ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
//  gggg    ....    gggg    gggg    ....

//   5***:   6***:   7*:     8*:     9**:
//  aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
//  dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
//  gggg    gggg    ....    gggg    gggg

const unordered_map<int,int> UNIQUES = { {2,1}, {3,7}, {4,4}, {7,8} };

bitset<7> getBitset(string s) {
  bitset<7> out(0);
  for (int i = 0; i < s.size(); i++) {
    out.set(s[i] - 'a');
  }
  return out;
}

vector<bitset<7>> parseLine(string line) {
  vector<bitset<7>> out;
  istringstream iss(line);
  string temp;

  while (iss >> temp)
    out.push_back(getBitset(temp));

  return out;
}

int getNum(string signals, string outputs) {
  vector<bitset<7>> sig_sets = parseLine(signals);
  vector<bitset<7>> out_sets = parseLine(outputs);
  unordered_map<int,bitset<7>> found;
  
  //round 1: find uniques
  for (auto it = sig_sets.begin(); it != sig_sets.end();) {
    int count = it->count();
    if (UNIQUES.find(count) != UNIQUES.end()) {
      found[UNIQUES.at(count)] = *it;
      it = sig_sets.erase(it);
    } else
      it++;
  }
  
  //round 2: find 9, 3
  for (auto it = sig_sets.begin(); it != sig_sets.end();) {
    int count = it->count();
    if (count == 5 && ((*it | found[7]) == *it)) {
      found[3] = *it;
      it = sig_sets.erase(it);

    } else if (count == 6 && ((*it | found[4]) == *it)) {
      found[9] = *it;
      it = sig_sets.erase(it);

    } else
      it++;
  }

  //round 3: find rest
  for (auto it = sig_sets.begin(); it != sig_sets.end();) {
    int count = it->count();
    if (count == 5) {
      bitset<7> test = *it ^ found[9];
      if (test.count() == 1)
        found[5] = *it;
      else
        found[2] = *it;
      it = sig_sets.erase(it);

    } else if (count == 6) {
      if ((*it | found[7]) == *it)
        found[0] = *it;
      else 
        found[6] = *it;
      it = sig_sets.erase(it);

    } else
      it++;
  }

  unordered_map<bitset<7>,int> rfound;
  for (auto f : found)
    rfound[f.second] = f.first;

  int mult = 1000, ret = 0;
  for (auto seg : out_sets) {
    ret += rfound[seg] * mult;
    mult /= 10;
  }

  return ret;
}

int main(int argc, char** argv) {
  if (argc < 2) {
    cout << "Specify filename\n";
    return EXIT_FAILURE;
  }

  ifstream infile(argv[1]);
  string line, output, signal;
  vector<string> signals, outputs;

  while (getline(infile, line)) {
    signals.push_back(line.substr(0,59));
    outputs.push_back(line.substr(60));
  }
  infile.close();

  int tally = 0;
  for (auto output : outputs) {
    istringstream iss(output);
    string num;
    while(iss >> num)
      if (UNIQUES.find(num.size()) != UNIQUES.end())
        tally++;
  }

  cout << tally << "\n";

  unsigned long total = 0;
  for (int i = 0; i < signals.size(); i++) {
    total += getNum(signals[i], outputs[i]);
  }

  cout << total << "\n";

  return EXIT_SUCCESS;
}