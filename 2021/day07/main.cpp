#include <iostream>
#include <vector>
#include <algorithm>
#include <cmath>

using namespace std;

int fuel(const vector<int>& crabs) {
  int median = crabs[crabs.size()/2];

  int fuel = 0;
  for (auto crab : crabs) {
    fuel += abs(crab - median);
  }

  return fuel;
}

int fuel2(const vector<int>& crabs) {
  double average = 0;
  for (auto crab : crabs)
    average += crab;
  average /= crabs.size();

  int fuel[] = {0, 0};
  for (auto crab : crabs) {
    int dist1 = abs(crab - floor(average));
    int dist2 = abs(crab - ceil(average));
    fuel[0] += (dist1 * (dist1 +1 )) / 2;
    fuel[1] += (dist2 * (dist2 +1 )) / 2;
  }

  if (fuel[0] < fuel[1])
    return fuel[0];

  return fuel[1];
}

int main() {
  int in;
  vector<int> crabs;

  while (cin >> in) {
    cin.get();
    crabs.push_back(in);
  }

  sort(crabs.begin(),crabs.end());

  cout << "Part 1 fuel: " << fuel(crabs) << "\n";
  cout << "Part 2 fuel: " << fuel2(crabs) << "\n";

  return EXIT_SUCCESS;
}