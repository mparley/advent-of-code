#include <iostream>
#include <vector>
#include <fstream>
#include <string>

using namespace std;

// Counts occupied seats
int countOccupied(vector<vector<char>>& seats) {
    int seatcount = 0;
    for (auto row : seats)
        for (char col : row)
            if (col == '#')
                seatcount++;
    return seatcount;
}

// Prints the seats 2d vector
void printSeats(const vector<vector<char>>& seats) {
    for (auto row : seats) {
        for (char c : row)
            cout << c;
        cout << '\n';
    }
    cout << '\n';
}

// Inclusive in range
bool inBounds(int val, int min, int max) {
    return val >= min && val <= max;
}

// Generalized helper for part 2 visible function
bool lookDir(int dx, int dy, int x, int y, const vector<vector<char>>& seats) {
    int xmax = seats[0].size()-1;
    int ymax = seats.size()-1;

    for (int ix = x+dx, iy = y+dy;
         inBounds(ix, 0, xmax) && inBounds(iy, 0, ymax);
         ix += dx, iy+= dy)
    {
        if (seats[iy][ix] == '#')
            return true;
        else if (seats[iy][ix] == 'L')
            break;
    }

    return false;
}

// Helper for part 1 making sure checked x, y are in bounds
bool lookAt(int x, int y, const vector<vector<char>>& seats) {
    if (inBounds(x, 0, seats[0].size()-1)
        && inBounds(y, 0, seats.size()-1))
    {
        return seats[y][x] == '#';
    } else
        return false;
}

// Part 2 check for seats in line of sight
int visible(int x, int y, const vector<vector<char>>& seats) {
    int count = 0;
    count += lookDir(1, 0, x, y, seats); // right
    count += lookDir(-1, 0, x, y, seats); // left
    count += lookDir(0, 1, x, y, seats); // down
    count += lookDir(0, -1, x, y, seats); // up
    count += lookDir(1, 1, x, y, seats); // down right
    count += lookDir(-1, -1, x, y, seats); // up left
    count += lookDir(1, -1, x, y, seats); // up right
    count += lookDir(-1, 1, x, y, seats); // down left
    return count;
}

// Part 1 check of each immediately adjacent seat
int adjacent(int x, int y, const vector<vector<char>>& seats) {
    int count = 0;
    count += lookAt(x+1, y, seats); // right
    count += lookAt(x-1, y, seats); // left
    count += lookAt(x, y+1, seats); // down
    count += lookAt(x, y-1, seats); // up
    count += lookAt(x+1, y+1, seats); // down right
    count += lookAt(x-1, y-1, seats); // up left
    count += lookAt(x+1, y-1, seats); // up right
    count += lookAt(x-1, y+1, seats); // down left
    return count;
}

// Part 1 version of people arrive function
bool peopleArrive(vector<vector<char>>& seats) {
    int width = seats[0].size();
    int height = seats.size();
    bool changed = false;
    vector<vector<char>> copy = seats; // copy so altering doesn't alter check

    for (int i = 0; i < height; i++) {
        for (int j = 0; j < width; j++) {
            if (seats[i][j] == 'L' && adjacent(j, i, copy) == 0) {
                seats[i][j] = '#';
                changed = true;
            } else if (seats[i][j] == '#' && adjacent(j, i, copy) >= 4) {
                seats[i][j] = 'L';
                changed = true;
            }
        }
    }

    //printSeats(seats);

    return changed;
}

// Part 2 people arrived function. Basically the same except using visible()
bool peopleArriveRevised(vector<vector<char>>& seats) {
    int width = seats[0].size();
    int height = seats.size();
    bool changed = false;
    vector<vector<char>> copy = seats;

    for (int i = 0; i < height; i++) {
        for (int j = 0; j < width; j++) {
            if (seats[i][j] == 'L' && visible(j, i, copy) == 0) {
                seats[i][j] = '#';
                changed = true;
            } else if (seats[i][j] == '#' && visible(j, i, copy) >= 5) {
                seats[i][j] = 'L';
                changed = true;
            }
        }
    }

    //printSeats(seats);

    return changed;
}


// Filename is passed as first argument
int main(int argc, char *argv[]) {
    if (argc < 2) return EXIT_FAILURE;

    ifstream infile(argv[1]);
    vector<vector<char>> seats, seats2;
    string line;
    char c;

    while (getline(infile, line)) {
        vector<char> temp;
        for (char c : line)
            temp.push_back(c);

        seats.push_back(temp);
        seats2.push_back(temp);
    }

    //printSeats(seats);
    while (peopleArrive(seats)) {}
    cout << countOccupied(seats) << "\n";

    while (peopleArriveRevised(seats2)) {}
    cout << countOccupied(seats2) << "\n";

    return EXIT_SUCCESS;
}
