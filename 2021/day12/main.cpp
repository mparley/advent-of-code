#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <unordered_map>
#include <unordered_set>

using namespace std;

class Cave {
private:
  string name_;
  bool big_;
  vector<Cave*> adjacent_;
  
public:
  Cave(string name) {
    name_ = name;
    big_ = isupper(name[0]) ? true : false;
  }

  ~Cave() {}

  string name() { return name_; }
  bool big() { return big_; }

  vector<Cave*> GetAdjacent() { return adjacent_; }

  void AddPath(Cave* adj_cave) {
    adjacent_.push_back(adj_cave);
  }
};

void findPathsHelper(vector<Cave*> path, vector<string>& output, 
  unordered_set<string> small_caves, Cave* cur, Cave* target, 
  int srevisit_used = -1)
{
  path.push_back(cur);

  if (cur->name() == target->name()) {
    string out;
    for (auto c : path)
      out += c->name() + ",";
    output.push_back(out);
    return;
  }

  if (!cur->big()) {
    if (small_caves.find(cur->name()) != small_caves.end()) {
      if (srevisit_used != 0 || cur->name() == "start")
        return;
      else
        srevisit_used = 1;
    }
    else {
      small_caves.insert(cur->name());
    }
  }

  for (auto cave : cur->GetAdjacent()) {
    findPathsHelper(path, output, small_caves, cave, target, srevisit_used);
  }
}

vector<string> findPaths(Cave* start, Cave* end, int part = 1) {
  vector<Cave*> path;
  vector<string> outputs;
  unordered_set<string> small_caves;

  path.push_back(start);
  small_caves.insert(start->name());

  for (auto a : start->GetAdjacent()) {
    findPathsHelper(path, outputs, small_caves, a, end, part-2);
  }

  return outputs;
}

int main(int argc, char** argv) {
  unordered_map<string,Cave*> caves;
  
  ifstream infile(argv[1]);
  while (infile.good()) {
    string line, names[2];
    getline(infile,line);

    auto delim = line.find('-');
    names[0] = line.substr(0,delim);
    names[1] = line.substr(delim+1);

    for (auto n : names)
      if (caves.find(n) == caves.end())
        caves[n] = new Cave(n);

    caves[names[0]]->AddPath(caves[names[1]]);
    caves[names[1]]->AddPath(caves[names[0]]);
  }
  infile.close();

  /*
  for (auto c : caves) {
    cout << c.second->name() << ": ";
    for (auto a : c.second->GetAdjacent()) {
      cout << a->name() << ",";
    }
    cout << "\n";
  }
  */
  
  vector<string> paths = findPaths(caves["start"], caves["end"]);
  cout << paths.size() << "\n";
  paths = findPaths(caves["start"], caves["end"],2);
  cout << paths.size() << "\n";

  for (auto c : caves)
    delete c.second;
  caves.clear();

  return EXIT_SUCCESS;
}