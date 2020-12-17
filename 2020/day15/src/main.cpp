#include <iostream>
#include <unordered_map>
#include <fstream>
#include <string>

using namespace std;

// Pass filename and target turn as arguments
int main(int argc, char *argv[]) {
    if (argc < 3) return EXIT_FAILURE;

    ifstream infile(argv[1]);
    unordered_map<int,int> turns;
    uint64_t turn = 1, last_num;
    uint64_t target_num = stoull(argv[2]);
    string temp;

    // Read numbers in and mark their turns in the hash map
    while (getline(infile, temp, ',')) {
        uint64_t number = stoull(temp);

        cout << turn << ": " << number << "\n";

        turns[number] = turn;
        last_num = number;
        turn++;
    }

    // Pretty dumb solution but basically loop x number of times to reach target.
    // Each loop, if a previous turn of the number is recorded, get the age,
    // update last turn, and set the number variable to the age for the next
    // loop. If not update the turn and set it to 0.
    while (turn <= target_num) {
        if (turns.find(last_num) != turns.end()) {
            uint64_t age = turn - 1 - turns[last_num];
            turns[last_num] = turn - 1;
            last_num = age;
        } else {
            turns[last_num] = turn - 1;
            last_num = 0;
        }

        turn++;
    }

    cout << last_num << "\n";

    return EXIT_SUCCESS;
}
