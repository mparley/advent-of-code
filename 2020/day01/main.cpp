#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

/* Starts at the ends of the sorted list and works inwards depending on the
difference of the ends from the target sum */
int find_two_expenses(vector<int> expenses, int target_sum) {
  auto big = expenses.end()-1;
  auto small = expenses.begin();

  while (big > small) {
    // Finds the difference against the high value
    int diff = target_sum - *big;

    // Found if small value is the diff, makes sure they are different entries
    if (diff == *small && big != small)
      return *big * *small;

    // Moves small up to or past diff
    while (*small < diff && small < big)
      small++;

    // Gets diff against low value
    diff = target_sum - *small;

    // Moves high value down until at or below diff
    while (diff < *big && big > small)
      big--;
  }

  // Failed to find value
  return -1;
}

/* Spent a long time trying to be clever but ended up giving up and going
with this brute force solution with nested loops. Binary search helps a bit */
int find_three_expenses(vector<int> expenses, int target_sum) {
  for (auto it1 = expenses.begin(); it1 < expenses.end(); it1++ ) {
    for (auto it2 = it1 + 1; it2 < expenses.end(); it2++ ) {
      int diff = target_sum - (*it1 + *it2);
      if (diff > 0) {
        auto it3 = it2 + 1;
        if (binary_search(it3, expenses.end(), diff))
          return *it1 * *it2 * diff;
      }
    }
  }

  return -1;
}

/* The target sum must be passed as an argument, the data is given
via stdin so pipe that shit in */
int main(int argc, char *argv[]) {
  if (argc <= 1) {
    cout << "Pass the target sum as an argument\n";
    return -1;
  }

  int target = atoi(argv[1]);

  vector<int> expense_entries;
  int expense;
  while (cin >> expense)
    expense_entries.push_back(expense);

  // Pre sort data so you don't have to sort twice
  sort(expense_entries.begin(), expense_entries.end());

  int part1 = find_two_expenses(expense_entries, target);
  cout << "Result of part 1: " << part1 << '\n';
  long part2 = find_three_expenses(expense_entries, target);
  cout << "Result of part 2: " << part2 << '\n';

  return 0;
}
