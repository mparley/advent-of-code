#include <iostream>
#include <fstream>
#include <tuple>
#include <queue>
#include <vector>
#include <set>

using namespace std;

// Given the range and vector, find the sum of the smallest and largest numbers
int weaknessFromRange(int low, int high, const vector<int>& data) {
    int smallest = data[low], largest = data[low];
    for (int i = low; i <= high; i++) {
        if (data[i] < smallest) smallest = data[i];
        if (data[i] > largest) largest = data[i];
    }
    return smallest + largest;
}

// Finds the contiguous range of data entries that sum up to the target
tuple<int,int> findSumRange(int target, const vector<int>& data) {
    int low = 0, high = 1;
    int sum = data[low] + data[high];

    while (high < data.size()) {

        // While the sum is less than target we expand range
        while (sum < target) {
            high++;
            if (high >= data.size()) break; // Prevents segfault

            sum += data[high];
            if (sum == target)
                return {low, high};
        }

        // If we passed the target increase low, reset range and sum
        low++;
        high = low + 1;
        sum = data[low] + data[high];
    }

    return {-1, -1};
}

// Given the preamble checks that entry is a sum of two of it's numbers
bool isSumOf(int num, const set<int>& list) {
    auto high = list.lower_bound(num); // starts at target (or just after)
    auto low =  list.begin();
    bool found = false;

    // High and low converge
    // if low is the difference of num and high, num is valid
    // if low is greater than the difference, we move high back
    // if low is less than the difference, we move high forward
    while (*high > *low) {
        if (*low == num - *high) {
            found = true;
            break;
        } else if (*low > num - *high)
            high--;
        else
            low++;
    }

    return found;
}

// Filname is passed as first argument, preamble size is second argument
int main(int argc, char *argv[]) {
    if (argc < 3) return EXIT_FAILURE;

    ifstream infile(argv[1]);
    queue<int> dataq;
    vector<int> data;
    set<int> preamble;
    int entry, rule_breaker, preamble_size = stoi(argv[2]);
    bool found_rule_break = false;

    // Read input while checking for the invalid entry
    while (infile >> entry) {
        dataq.push(entry);
        data.push_back(entry);
        preamble.insert(entry);

        // Continue if preamble isn't filled yet
        if (dataq.size() <= preamble_size)
            continue;

        // Check for rule break. If not found remove first entry from preamble
        if (!found_rule_break) {
            if (!isSumOf(dataq.back(), preamble)) {
                rule_breaker = dataq.back();
                cout << dataq.back() << " breaks rule\n";
                found_rule_break = true;
            }

            preamble.erase(dataq.front());
            dataq.pop();
        }
    }

    auto [low_i, high_i] = findSumRange(rule_breaker, data);
    if (low_i != -1)
        cout << weaknessFromRange(low_i, high_i, data) << "\n";
    else
        cout << "Encryption weakness not found\n";

    return EXIT_SUCCESS;
}
