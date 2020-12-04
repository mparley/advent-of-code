#include <string>
#include <unordered_map>
#include <unordered_set>
#include <algorithm>
#include <bitset>
#include <fstream>
#include <sstream>
#include <iostream>

using namespace std;

enum { BYR, IYR, EYR, HGT, HCL, ECL, PID, CID };

// Filename is passed as argument
int main(int argc, char* argv[]) {
  if (argc != 2)
    return 1;
  
  unordered_map<string, int> fields = {
    {"byr", BYR}, {"iyr", IYR}, {"eyr", EYR}, {"hgt", HGT}, {"hcl", HCL},
    {"ecl", ECL}, {"pid", PID}, {"cid", CID},
  };
  
  unordered_set<string> eye_colors= {
    "amb", "blu", "brn", "gry", "grn", "hzl", "oth",
  };

  ifstream infile(argv[1]);
  string line;
  bitset<8> found, valid, required(string("01111111"));
  int count_found = 0, count_valid = 0;

  while (infile.good()) {
    getline(infile, line);
    stringstream ss(line);
    string temp;

    while (ss >> temp) {
      string prop = temp.substr(0, 3);
      string val = temp.substr(4);

      if (fields.find(prop) != fields.end()) {

        // Set the corresponding bit in the found bitset
        found.set(fields[prop]);

        // Big ol' switch to set the bits in valid bitset
        switch (fields[prop]) {

          // Following three cases check the years' requirements
          case BYR:
            if (val.length() == 4 && (stoi(val) >= 1920 && stoi(val) <= 2002))
              valid.set(BYR);
            break;

          case IYR:
            if (val.length() == 4 && (stoi(val) >= 2010 && stoi(val) <= 2020))
              valid.set(IYR);
            break;

          case EYR:
            if (val.length() == 4 && (stoi(val) >= 2020 && stoi(val) <= 2030))
              valid.set(EYR);
            break;

          // Height check
          case HGT: {
            if (val.length() <= 2) break; //To protect from passing empty string to stoi

            int ival = stoi(val.substr(0, val.length()-2));
            string unit = val.substr(val.length()-2);

            if (unit == "in"  && ival >= 59 && ival <= 76)
              valid.set(HGT);
            else if (unit == "cm" && ival >= 150 && ival <= 193)
              valid.set(HGT);
            break;
          }

          // Hair color
          case HCL:
            // This checks that first char is #, the length is correct, and then
            // checks every character to make sure it's a hex digit
            if (val[0] == '#' && val.length() == 7 
              && all_of(val.begin()+1, val.end(),
              [](unsigned char c){ return ::isxdigit(c); }))
              valid.set(HCL);
            break;

          // Eye color, just sees if val is in the eye color set
          case ECL:
            if (eye_colors.find(val) != eye_colors.end())
              valid.set(ECL);
            break;

          // PID case checks length and that all characters are digits
          case PID:
            if (val.length() == 9
              && all_of(val.begin(), val.end(),
              [](unsigned char c){ return ::isdigit(c); }))
              valid.set(PID);
            break;

          default:
            break;
        }
      }
    }

    // At empty line or end of file checks the bitsets against the required
    // bitset, increments the counts, and resets the bitsets for the next entry
    if (line.empty() || !infile.good()) {
      if ((found & required) == required)
        count_found++;

      if ((valid & required) == required)
        count_valid++;

      found.reset();
      valid.reset();
    }
  }

  cout << count_found << " " << count_valid << '\n';

  return 0;
}