#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <deque>
#include <queue>
#include <sstream>

using namespace std;

// Got ahead of myself and added minus and division
uint64_t operate(uint64_t l, uint64_t r, char op) {
    uint64_t ret;
    switch (op) {
        case '+':
            ret = l + r;
            break;
        case '-':
            ret = l - r;
            break;
        case '*':
            ret = l * r;
            break;
        case '/':
            ret = l / r;
            break;
        default:
            ret = 0;
            break;
    }

    //cout << l << " " << op << " " << r << " = " << ret << "\n";
    return ret;
}


// This runs the operations
uint64_t runOps(deque<uint64_t>& values, queue<char>& operations) {
    uint64_t l = values.front();
    values.pop_front();

    while (!values.empty() && !operations.empty()) {
        l = operate(l, values.front(), operations.front());
        operations.pop();
        values.pop_front();
    }

    return l;
}

// Part 1, pretty dumb assums valid input
uint64_t evalLine(stringstream& ss) {
    deque<uint64_t> vals;
    queue<char> operations;
    char c;

    while (ss >> c) {
        if (isdigit(c)) {
            vals.push_back(c-'0');
            continue;
        }

        switch (c) {
            case '(':
               vals.push_back(evalLine(ss));
               break;
            case ')':
                return runOps(vals, operations);
            case '+':
            case '*':
                operations.push(c);
                break;
            default:
                break;
        }
    }

    return runOps(vals, operations);
}

// Part 2 I basically just changed vals from queue to deque and
// handled the add operation right there.
uint64_t evalLineAdvanced(stringstream& ss) {
    deque<uint64_t> vals;
    queue<char> operations;
    char c;

    while (ss >> c) {
        if (isdigit(c)) {
            vals.push_back(c-'0');
            continue;
        }

        switch (c) {
            case '(':
               vals.push_back(evalLineAdvanced(ss));
               break;
            case ')':
                return runOps(vals, operations);
                break;
            case '+': {
                ss >> c;
                uint64_t back = vals.back(), right;
                vals.pop_back();
                if (c == '(') {
                    right = evalLineAdvanced(ss);
                } else {
                    right = c - '0';
                }
                vals.push_back(operate(back, right, '+'));
                break;
            }
            case '*':
                operations.push(c);
                break;
            default:
                break;
        }
    }

    return runOps(vals, operations);
}

// Filename passed as the first argument
int main(int argc, char *argv[]) {
    if (argc < 2) return EXIT_FAILURE;

    ifstream infile(argv[1]);
    vector<string> lines;
    string line;

    // That empty line screwed me over for a good while
    while (getline(infile, line))
        if (!line.empty())
            lines.push_back(line);
    infile.close();

    uint64_t sum = 0, sum2 = 0;
    for (auto line: lines) {
        stringstream ss(line), ss2(line);
        sum += evalLine(ss);
        sum2 += evalLineAdvanced(ss2);
    }

    cout << sum << "\n" << sum2 << endl;
    return EXIT_SUCCESS;
}
