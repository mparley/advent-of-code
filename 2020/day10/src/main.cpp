#include <iostream>
#include <fstream>
#include <algorithm>
#include <unordered_map>
#include <vector>

using namespace std;

// Filename passed as first argument
int main(int argc, char *argv[]) {
    if (argc < 2) return EXIT_FAILURE;

    ifstream infile(argv[1]);
    vector<int> joltages;
    int next_joltage = 0, diff3 = 0, diff1 = 0;
    unordered_map<int, unsigned long> paths;
    joltages.push_back(0);

    while (infile >> next_joltage) {
        joltages.push_back(next_joltage);
    }
    infile.close();

    // Sort the adapters and count the differences of 1s and 3s for part 1
    sort(joltages.begin(), joltages.end());
    for (int i=0; i < joltages.size()-1; i++) {
        int diff = joltages[i+1] - joltages[i];
        if (diff == 3) diff3++;
        else if (diff == 1) diff1++;
    }

    // Add the built in adapter and it's difference, and print out part 1 answer
    joltages.push_back(*(joltages.end()-1) + 3);
    diff3++;
    cout << diff1 * diff3 << "\n";

    // This was the hard part
    // Spent a lot of time trying to do something with graphs and dfs til
    // I realized you can just go through the sorted vector adding the immediate
    // 3 lower joltage adapters' path counts (if they exist) to get the current
    // adapter's path count.
    paths[joltages[0]] = 1;
    for (auto joltage : joltages) {
        paths[joltage] += paths[joltage-1] + paths[joltage-2] + paths[joltage-3];
    }

    cout << paths[joltages.back()] << "\n";

    return 0;
}
