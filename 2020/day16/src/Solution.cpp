#include "Solution.hpp"

// Checks if a range is valid
bool valid(const vector<Range>& rule, int val) {
    for (auto r: rule) {
        if (val >= r.min && val <= r.max) {
            return true;
        }
    }
    return false;
}

// Part 1 solution, also removes tickets for part 2
int findErrorRate(const unordered_map<string,vector<Range>>& rules,
    vector<vector<int>>& tickets)
{
    int sum = 0;
    vector<int> bad_indexes;

    // Loops through all tickets after the first ticket (yours) and checks its
    // validity. If invalid, it tracks the index and adds the value to sum
    for (int i = 1; i < tickets.size(); i++) {
        for (auto val: tickets[i]) {
            bool found_valid = false;

            // This checks the value against all the rules if it matches any
            // we set found to true and break
            for (auto rule: rules) {
                if (valid(rule.second, val)) {
                    found_valid = true;
                    break;
                }
            }

            // This is only executed when we've checked against all rules
            // and none of them matched
            if (!found_valid) {
                sum += val;
                bad_indexes.push_back(i);
                break;
            }
        }
    }

    // We reverse sort and delete the bad tickets
    sort(bad_indexes.rbegin(), bad_indexes.rend());
    for (auto i: bad_indexes)
        tickets.erase(tickets.begin()+i);

    return sum;
}

// Part 2 solution, we need to use boost because eric decided we needed to go
// fuck ourselves for using c++ instead of python.
cpp_int departureProduct(const unordered_map<string, vector<Range>>& rules,
    const vector<vector<int>>& tickets)
{
    // Possible rules tracks the set of rules for each ticket value position
    // the indexes coorespond
    vector<unordered_set<string>*> possible_rules;
    unordered_set<string> base;

    // Populate base set with the class names
    for (auto r : rules)
        base.insert(r.first);

    // Populate the set for each field
    for (int i=0; i < tickets[0].size(); i++)
        possible_rules.emplace_back(new unordered_set(base));

    // This goes through each ticket value and removes rules that don't fit the
    // values. Luckily after this we find that at least one value can only be one
    // class. This makes it easier to eliminate possiblities.
    for (auto t: tickets) {
        for (int i=0; i < t.size(); i++) {
            if (possible_rules[i]->size() == 1) continue;
            for (auto it=possible_rules[i]->begin(); it != possible_rules[i]->end();) {
                if (!valid(rules.at(*it), t[i])) {
                    it = possible_rules[i]->erase(it);
                }
                else
                    it++;
            }
        }
    }

    // Here we go through the possible rules sets and if the set only contains
    // one element that must be the rule so we remove that rule from the other
    // sets and add it to this field map.
    unordered_map<string, int> fields;
    while (fields.size() < tickets[0].size()) {
        for (int i=0; i < possible_rules.size(); i++) {
            if (possible_rules[i]->empty()) continue;

            if (possible_rules[i]->size() == 1) {
                string rule = *possible_rules[i]->begin();
                fields[rule] = i;

                for (auto p : possible_rules)
                    p->erase(rule);
            }
        }
    }

    for (auto p: possible_rules) {
        delete p;
    }
    possible_rules.clear();

    // Go through your own ticket's departure fields and multiply the values
    cpp_int ret = 1;
    for (auto f: fields)
        if (f.first.find("departure") != string::npos) {
            ret *= (uint64_t)tickets[0][f.second];
        }

    return ret;
}
