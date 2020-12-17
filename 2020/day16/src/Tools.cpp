#include "Tools.hpp"

// This reads a rule line using regex and puts it in the rules map
void readRule(unordered_map<string, vector<Range>>& rules, string line) {
    regex re("([\\w ]*):? (\\d*)-(\\d*)");
    auto rIt = sregex_iterator(line.begin(), line.end(), re);
    auto end = sregex_iterator();
    vector<Range> rule;
    string rule_class;

    if (rIt != end)
        rule_class = (*rIt)[1];
    else return;

    rule.reserve(distance(rIt, end));
    while (rIt != end) {
        rule.emplace_back(stoi((*rIt)[2]), stoi((*rIt)[3]));
        rIt++;
    }

    rules[rule_class] = rule;
}

// This reads a ticket line
vector<int> readTicket(string line) {
    vector<int> ret;
    stringstream ss(line);
    string temp;

    while (getline(ss, temp, ','))
        ret.push_back(stoi(temp));

    return ret;
}

// This reads the file and populates the hash map and vector of tickets
// the mode uses an enum declared in Tools.h
void readFile(unordered_map<string, vector<Range>>& rules,
    vector<vector<int>>& tickets, string filename)
{
    int mode = RULES;
    ifstream infile(filename);
    string line;

    while (getline(infile, line)) {
        if (line.empty()) {
            mode = SET;
            continue;
        }

        switch (mode) {
            case RULES:
                readRule(rules, line);
                break;
            case OWN:
                tickets.insert(tickets.begin(), readTicket(line));
                break;
            case NEARBY:
                tickets.push_back(readTicket(line));
                break;
            case SET:
                mode = (line.find("your") != string::npos) ? OWN : NEARBY;
                break;
            default:
                break;
        }
    }
}
