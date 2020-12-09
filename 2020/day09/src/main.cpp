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
// improved from last version I had by shamelessly taking inspiration from others
tuple<int,int> findRange(int target, const vector<int>& data) {
    int low = 0, sum = data[low];

    // Loop grows range on high end and adds to sum
    for (int high = 1; high < data.size(); high++) {
        sum += data[high];

        // If sum is greater than target, shrink range from bottom until it isn't
        while (sum > target) {
            sum -= data[low];
            low++;
        }

        if (sum == target)
            return  {low, high};
    }

    return {-1, -1};
}

// Given the preamble checks that entry is a sum of two of it's numbers
bool isSumOf(int num, const set<int>& list) {
    auto high = list.lower_bound(num);
    auto low =  list.begin();

    // High and low converge while checking low against num minus high
    while (*high > *low) {
        if (*low == num - *high)
            return true;
        else if (*low > num - *high)
            high--;
        else
            low++;
    }

    return false;
}

// Filname is passed as first argument, preamble size is second argument
int main(int argc, char *argv[]) {
    if (argc < 3) return EXIT_FAILURE;

    ifstream infile(argv[1]);
    queue<int> dataq;
    vector<int> data;
    set<int> preamble;
    int entry, rule_breaker, preamble_size = stoi(argv[2]);

    // Read input while checking for the invalid entry
    while (infile >> entry) {
        dataq.push(entry);
        data.push_back(entry);
        preamble.insert(entry);

        // Continue if preamble isn't filled yet
        if (dataq.size() <= preamble_size)
            continue;

        // Check for rule break. If not found remove first entry from preamble
        if (!isSumOf(dataq.back(), preamble)) {
            rule_breaker = dataq.back();
            cout << dataq.back() << " breaks rule\n";
            break;
        }

        preamble.erase(dataq.front());
        dataq.pop();
    }

    auto [low_i, high_i] = findRange(rule_breaker, data);
    if (low_i != -1)
        cout << weaknessFromRange(low_i, high_i, data) << "\n";
    else
        cout << "Encryption weakness not found\n";

    return EXIT_SUCCESS;
}
