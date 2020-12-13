#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <climits>

using namespace std;

/* Frankly, I got filtered by part 02 of this puzzle. I think I might have learned
 * the chinese remainder theorem in college, but I didn't remember anything
 * about it and I defenitely wouldn't know to use it here without help from  and
 * discussions with others. Basically, these two functions are straight from
 * rosetta code  */
uint64_t mulInv(uint64_t a, uint64_t b) {
    uint64_t b0 = b, x0 = 0, x1 = 1;

    if (b == 1) return 1;

    while (a > 1) {
        uint64_t q = a / b;
        uint64_t a_mod_b = a % b;
        a = b;
        b = a_mod_b;

        uint64_t x_q_x = x1 - q * x0;
        x1 = x0;
        x0 = x_q_x;
    }

    if (x1 < 0)
        x1 += b0;

    return x1;
}

uint64_t chRemainder(const vector<int>& n, const vector<int>& a) {
    uint64_t prod = 1, sum = 0;

    for (int i : n)
        prod *= i;

    for (int i = 0; i < n.size(); i++) {
        uint64_t p = prod / n[i];
        sum += a[i] * mulInv(p, n[i]) * p;
    }

    return sum % prod;
}

// Filename is passed in as first argument. It expects two lines
int main(int argc, char *argv[]) {
    if (argc < 2) return EXIT_FAILURE;

    ifstream infile(argv[1]);
    int earliest;
    vector<int> buses, offsets;
    string temp;

    infile >> earliest;
    infile.get();

    // Read the second line. Offsets are negative from 0
    for (int i = 0; getline(infile, temp, ','); i--) {
        if (temp[0] == 'x')
            continue;
        buses.push_back(stoi(temp));
        offsets.push_back(i);
    }

    // Part 1 is just simple math to find minutes before earliest time and
    // finding the lowest
    int id, wait = INT_MAX;
    for (auto bus : buses) {
        int before_depart = ((earliest / bus) +1 ) * bus - earliest;
        if (wait > before_depart) {
            wait = before_depart;
            id = bus;
        }
    }

    cout << wait * id << '\n';
    cout << chRemainder(buses, offsets) << '\n';

    return EXIT_SUCCESS;
}
