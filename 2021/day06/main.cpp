#include <iostream>
#include <vector>

using namespace std;

unsigned long pass_days (vector<unsigned long> school, int days) {
    for (int day = 0, index = 0; day < days; day++) {
        school[(index+7)%9] += school[index];
        index = (index + 1) % 9;
    }

    unsigned long sum = 0;
    for (auto& fish : school)
        sum += fish;
    return sum;
}

int main() {
    int fish_in, index = 0;
    vector<unsigned long> school(9, 0);

    while (cin >> fish_in) {
        cin.get();
        school[fish_in]++;
    }

    cout << pass_days(school, 80) << "\n" << pass_days(school, 256) << "\n";

    return EXIT_SUCCESS;
}