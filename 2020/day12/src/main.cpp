#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <cmath>

using namespace std;

enum Dir { N, E, S, W };

// Struct for instructions
struct Nav {
    char action;
    int value;
    Nav(char c, int i) { action = c; value = i; }
};

// Struct for positions
struct Loc {
    int x, y;
    Loc (int in_x, int in_y) { x = in_x; y = in_y; }
    Loc () { x = 0; y = 0; }
};

// Part 1
int navigatePart1 (const vector<Nav>& instructions) {
    char action;
    int value, dir = Dir::E;
    char dirc[4] = {'N', 'E', 'S', 'W'};
    Loc pos(0, 0);

    for (auto instruction : instructions) {
        action = instruction.action;
        value = instruction.value;

        // If forward we just switch in the direction
        if (action == 'F') action = dirc[dir];

        switch (action) {
            case 'N':
                pos.y -= value;
                break;
            case 'S':
                pos.y += value;
                break;
            case 'E':
                pos.x += value;
                break;
            case 'W':
                pos.x -= value;
                break;
            case 'L':
                dir = ((dir - (value / 90)) % 4 + 4) % 4;
                break;
            case 'R':
                dir = (dir + (value / 90)) % 4;
                break;
            default:
                break;
        }
    }

    return abs(pos.x) + abs(pos.y);
}

// Helper for rotating 90 degrees
void rotate90(int value, Loc& coord) {
    int temp, times = abs(value / 90), dir = value < 0 ? -1 : 1;
    for (int i = 0; i < (times % 4); i++) {
        temp = coord.x;
        coord.x = -1 * coord.y * dir;
        coord.y = temp * dir;
    }
}

// Part 2
int navigatePart2 (const vector<Nav>& instructions) {
    char action;
    int value;
    Loc w(10, -1), pos(0, 0);

    for (auto instruction : instructions) {
        action = instruction.action;
        value = instruction.value;

        switch (action) {
            case 'N':
                w.y -= value;
                break;
            case 'S':
                w.y += value;
                break;
            case 'E':
                w.x += value;
                break;
            case 'W':
                w.x -= value;
                break;
            case 'L':
                rotate90(-1 * value, w);
                break;
            case 'R':
                rotate90(value, w);
                break;
            case 'F':
                pos.x += value * w.x;
                pos.y += value * w.y;
                break;
            default:
                break;
        }
    }

    return abs(pos.x) + abs(pos.y);
}

// Filename passed as first argument
int main(int argc, char *argv[]) {
    if (argc < 2) return EXIT_FAILURE;

    ifstream infile(argv[1]);
    string line;
    vector<Nav> instructions;

    // We could do parts 1 and 2 while parsing input but I wanted both
    // parts in one file so I put it into a vector
    while (getline(infile, line)) {
        instructions.emplace_back(Nav(line[0], stoi(line.substr(1))));
    }

    cout << navigatePart1(instructions) << '\n';
    cout << navigatePart2(instructions) << '\n';

    return EXIT_SUCCESS;
}
