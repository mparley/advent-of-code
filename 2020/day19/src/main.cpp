#include <iostream>
#include <string>
#include <fstream>
#include <unordered_map>
#include <vector>
#include <regex>
#include <sstream>

using namespace std;

// Recursively translate the numbers to a or b based on rules,
// output a regex string
string translate(int rule_num, const unordered_map<int, string>& rules,
    unordered_map<int,string>& lookup, int max_depth = 0, int depth = 0)
{
    stringstream ss(rules.at(rule_num));
    string out = "", temp;

    while (ss >> temp) {
        if (temp[0] == '|')
            out += temp;

        // If we find a " we know there is only one element and its a letter
        else if (temp[0] == '"')
            out += temp[1];

        // If not a single a/b or | then we handle a rule number
        else {
            int pos = rules.at(stoi(temp)).find('"');

            // First check the saved values for the rule
            if (lookup.find(stoi(temp)) != lookup.end())
                out += lookup[stoi(temp)];

            // Then see if the rule is a single a or b and add to lookup
            else if (pos != string::npos) {
                out += rules.at(stoi(temp))[pos+1];
                lookup[stoi(temp)] = rules.at(stoi(temp))[pos+1];

            // Else we make the recursive call
            } else {
                if (!max_depth == 0 && depth > max_depth) continue; //Stops loop
                out += "("
                    + translate(stoi(temp), rules, lookup, max_depth, depth+1)
                    + ")";
            }
        }
    }

    // At this point our out *should* be done so we add it to lookup if not
    // already saved
    if (lookup.find(rule_num) == lookup.end()) {
        bool has_digit = false;
        for (char c: out)
            if (isdigit(c)) {
                has_digit =true;
                break;
            }

        if (!has_digit)
            lookup[rule_num] = "(" + out + ")";
    }

    return out;
}


// Counts the matches using regex and translate
int countExactMatches(int rule, const vector<string>& messages,
    const unordered_map<int, string>& rules, unordered_map<int, string>& lookup,
    int max_depth = 100)
{
    string ex = translate(rule, rules, lookup, max_depth);
    regex reg(ex);
    int count = 0;

    for (auto message: messages) {
        if (regex_match(message, reg))
            count++;
    }

    return count;
}

// Filename is 1st arg, rule to match is 2nd arg
int main(int argc, char *argv[]) {
    if (argc < 3) return EXIT_FAILURE;

    ifstream infile(argv[1]);
    int to_match = stoi(argv[2]);
    unordered_map<int, string> rules, lookup;
    vector<string> messages;
    string line = ":^)";

    // Read rules first
    while (getline(infile, line) && !line.empty()) {
        stringstream ss(line);
        string temp;
        int rule_num;

        ss >> rule_num >> temp;
        getline(ss, temp);
        rules[rule_num] = temp;
    }

    // Then read messages, get max length of messages to use for depth check
    int max_length = 0;
    while (getline(infile, line)) {
        messages.push_back(line);
        if (max_length < line.size())
            max_length = line.size();
    }
    infile.close();

    // Part 1 call
    cout << countExactMatches(to_match, messages, rules, lookup) << "\n";

    // Change rules and lookup. Rule 8 is simple and can be done with regex
    // 11 is bullshit and can't be done with regex since the number of repeated
    // 42s and 31s need to match
    rules[8] = "42 | 42 8";
    rules[11] = "42 31 | 42 11 31";
    lookup[8] = "(" + lookup[42] + "+)";

    // Clean lookup of rule 11 instances
    for (auto it = lookup.begin(); it != lookup.end();)
        if ((*it).first == 11 || (*it).second.find("11") != string::npos)
            it = lookup.erase(it);
        else
            it++;

    // Part 2 call, give the max_depth as the max_length/10 could change to
    // whatever necessary. max_length is too big alone and will hit the
    // _GLIBCXX_REGEX_STATE_LIMIT
    cout << countExactMatches(to_match, messages, rules, lookup, max_length/10) << "\n";

    return EXIT_SUCCESS;
}
