#include "Tools.h"

// Helper for canContain. Calls canContain to populate a hash map tracking
// which bags can contain target and then counts the number of bags that can
int containsHelper(string target,
  unordered_map<string, unordered_map<string,int>*>& rules)
{
  unordered_map<string, bool> checked;
  for (auto rule : rules)
    if (checked.find(rule.first) != checked.end())
      continue;
    else
      checked[rule.first] = canContain(target, rule.first, checked, rules);

  int count = 0;
  for (auto bag : checked) {
    if (bag.second) {
      //printHelper(checked, rules, bag.first);
      //cout << '\n';
      count++;
    }
  }

  return count;
}

// Recursive function for finding number of bags that can contain the target,
// uses a helper function to start and count the found bags
bool canContain(string target, string bag,
  unordered_map<string, bool>& checked,
  const unordered_map<string, unordered_map<string,int>*>& rules)
{
  if (rules.at(bag) == nullptr) {
    return false;
  }

  auto rule = rules.at(bag);
  bool found = false;

  // If target is directly inside current bag
  if (rule->find(target) != rule->end())
    return true;

  // Checks all bags inside current bag
  for (auto r: *rule) {

    // If we have checked bag don't check again
    if (checked.find(r.first) != checked.end()) {
      found |= checked[r.first];
      continue;
    }

    checked[r.first] = canContain(target, r.first, checked, rules);
    found |= checked[r.first];
  }

  return found;
}

// Recursive function for counting all bags
int totalBags(string bag,
  const unordered_map<string, unordered_map<string,int>*>& rules) 
{
  if (rules.at(bag) == nullptr)
    return 1;

  int count = 0;
  for (auto r : *rules.at(bag))
    count += r.second * totalBags(r.first, rules);

  return count + 1;
}

// Helper for print a line of all the sub bags in checked that are true
// Used in contHelper
void printHelper(unordered_map<string,bool>& checked, 
  unordered_map<string, unordered_map<string,int>*>& rules,
  string bag) 
{
  if (checked[bag] == true)
    cout << bag << " ";
  else
    return;

  for (auto rule : *rules[bag])
    printHelper(checked, rules, rule.first);
}


// Uses regex to parse line and fill hash map
void parseLine(string line, unordered_map<string,
  unordered_map<string,int>*>& results) 
{
  regex bag("(.*) bags contain");
  regex contains("(\\d+) ([\\w ]*) bag[s]?");
  sregex_iterator end;
  smatch match;

  if (regex_search(line, match, bag)) {
    string bag_rule = match[1];

    auto rit = sregex_iterator(line.begin(), line.end(), contains);
    auto temp = new unordered_map<string, int>;

    // If there are no matches for second regex search (iterator is at end)
    // then set the rule to null
    if (rit == end) {
      results[bag_rule] = nullptr;
      return;
    }

    // Goes through captured groups and sets the rules accordingly
    while(rit != end) {
      for (int i = 1; i < rit->size(); i += 2) {
        (*temp)[(*rit)[i+1]] = stoi((*rit)[i]);
      }
      rit++;
    }
    results[bag_rule] = temp;
  }
}

// Given a filled rules hash map, it will sort and print all the contents
void printRules(const unordered_map<string,
  unordered_map<string,int>*>& rules)
{
  vector<string> bags;
  for (auto rule : rules)
    bags.push_back(rule.first);

  sort(bags.begin(), bags.end());

  for (auto bag : bags) {
    cout << bag << ":\n";
    if (rules.at(bag) != nullptr)
      for (auto r: *rules.at(bag))
        cout << "\t" << r.first << ": " << r.second << "\n";
  }
}
