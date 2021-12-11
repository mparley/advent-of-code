#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <queue>
#include <utility>
#include <boost/functional/hash/hash.hpp>

using namespace std;

int basinSize(pair<int,int> pos, const vector<string>& cave_map) {
  unordered_set<pair<int,int>,boost::hash<pair<int,int>>> visited;
  queue<pair<int,int>> q;

  q.push(pos);

  while (!q.empty()) {
    int r = q.front().first, c = q.front().second;
    q.pop();

    if (cave_map[r][c] == '9') continue;
    if (visited.find({r,c}) != visited.end()) continue;

    if (r != 0) q.push({r-1,c});
    if (r != cave_map.size() - 1) q.push({r+1,c});
    if (c != 0) q.push({r,c-1});
    if (c != cave_map[r].size()-1) q.push({r,c+1});

    visited.insert({r,c});
  }
  
  return visited.size();
}

int main(int argc, char** argv) {
  if (argc < 2) {
    cout << "Specify file\n";
    return EXIT_FAILURE;
  }

  vector<string> lines;
  unordered_map<pair<int,int>,int,boost::hash<pair<int,int>>> risk_map;
  string iline;

  ifstream infile(argv[1]);
  while (getline(infile,iline)) {
    lines.push_back(iline);
  }
  infile.close();

  for (int i = 0; i < lines.size(); i++) {
    for (int j = 0; j < lines[i].size(); j++) {
      if (i != 0 && lines[i-1][j] <= lines[i][j]) continue;
      if (i != lines.size()-1 && lines[i+1][j] <= lines[i][j]) continue;
      if (j != 0 && lines[i][j-1] <= lines[i][j]) continue;
      if (j != lines[i].size()-1 && lines[i][j+1] <= lines[i][j]) continue;
      risk_map[{i,j}] = (lines[i][j] - '0') + 1;
    }
  }

  int risk_sum = 0;
  for (auto r : risk_map)
    risk_sum += r.second;

  cout << risk_sum << "\n";

  priority_queue<int> pq;
  for (auto r : risk_map) {
    pq.push(basinSize(r.first, lines));
  }

  int solution = 1;
  for (int i = 0; i < 3; i++) {
    solution *= pq.top();
    pq.pop();
  }

  cout << solution << "\n";

  return EXIT_SUCCESS;
}