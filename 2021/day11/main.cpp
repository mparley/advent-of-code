#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <queue>
#include <utility>

using namespace std;

vector<pair<int,int>> getAdjacent(int r, int c, int rmax, int cmax) {
  vector<pair<int,int>> out;

  if (c != 0)                  out.push_back({ r   , c-1 });
  if (r != 0 && c != 0)        out.push_back({ r-1 , c-1 });
  if (r != 0)                  out.push_back({ r-1 , c   });
  if (r != 0 && c != cmax)     out.push_back({ r-1 , c+1 });
  if (c != cmax)               out.push_back({ r   , c+1 });
  if (r != rmax && c != cmax)  out.push_back({ r+1 , c+1 });
  if (r != rmax)               out.push_back({ r+1 , c   });
  if (r != rmax && c != 0)     out.push_back({ r+1 , c-1 });

  return out;
}

unsigned int step(vector<string>& rows) {
  queue<pair<int,int>> q;
  unsigned int flash_count = 0;

  for (int row = 0; row < rows.size(); row++) {
    for (int col = 0; col < rows[col].size(); col++) {
      rows[row][col]++;
      if (rows[row][col] > '9')
        q.push({row,col});
    }
  }

  while (!q.empty()) {
    int r = q.front().first, c = q.front().second;
    q.pop();

    if (rows[r][c] != '0') {
      rows[r][c] = '0';
      flash_count++;
      
      vector<pair<int,int>> adj = getAdjacent(r,c,rows.size()-1,rows[r].size()-1);

      for (auto a : adj) {
        if (rows[a.first][a.second] != '0') {
          rows[a.first][a.second]++;
          if (rows[a.first][a.second] > '9')
            q.push({a.first,a.second});
        }
      }
    }
  }

  return flash_count;
}

int main(int argc, char** argv) {
  vector<string> rows;
  int steps = stoi(argv[2]);

  ifstream infile(argv[1]);
  while (infile.good()) {
    rows.push_back("");
    getline(infile,rows.back());
  }
  infile.close();

  unsigned long flashes = 0;
  int all_flash = -1;

  for (int i = 1; i <= steps; i++) {
    unsigned int t = step(rows);
    if (t == 100 && all_flash == -1) all_flash = i;
    flashes += t;
  }

  for (auto r : rows)
    cout << r << "\n";
  
  cout << "\n" << flashes << "\n";

  for (int i = steps+1; all_flash == -1; i++) {
    unsigned int t = step(rows);
    if (t == 100) all_flash = i;
  }

  cout << all_flash << "\n";

  return EXIT_SUCCESS;
}