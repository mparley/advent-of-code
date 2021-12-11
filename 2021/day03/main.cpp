#include <iostream>
#include <string>
#include <vector>

using namespace std;

unsigned int convert(string val) {
    unsigned int r = 0;
    for (int i = 0; i < val.size(); i++) {
        if (val[i] == '1')
            r |= (1 << (val.size()-1-i));
    }
    return r;
}

unsigned int invert(unsigned int value, unsigned int num_bits) {
    unsigned int v = ~value;
    unsigned int mask = 0;
    for (int i = 0; i < num_bits; i++) {
        mask |= (1 << i);
    }

    return v & mask;
}

string rating(vector<string> lines, int num0, int num1, bool most_common = true) {
    int p = 0, num_bits = lines[0].size();

    while (lines.size() > 1) {
        char target = (most_common == (num1 >= num0)) ? '1' : '0';
        num0 = 0;
        num1 = 0;

        for (auto it = lines.begin(); it != lines.end() && p < num_bits - 1;) {
            if ((*it)[p] == target) {
                if ((*it)[p+1] == '1')
                    num1++;
                else if ((*it)[p+1] == '0')
                    num0++;
                it++;
            } else {
                it = lines.erase(it);
            }
        }
        p++;
    }

    return lines[0];
}

int main() {
    vector<string> lines;
    string line;
    cin >> line;
    lines.push_back(line);

    const size_t num_bits = line.size();
    vector<unsigned int> bit_freq(num_bits, 0);

    while (cin >> line) {
        lines.push_back(line);
        for (int i = 0; i < num_bits; i++)
            bit_freq[i] += (line[i] == '1');
    }

    for (int i = 0; i < num_bits; i++)
        line[i] = (bit_freq[i] >= (lines.size() / 2)) ? '1' : '0';

    unsigned int gamma = convert(line);
    cout << gamma * invert(gamma, num_bits) << endl;

    string oxy_str = rating(lines,lines.size()-bit_freq[0],bit_freq[0]);
    unsigned int oxy = convert(oxy_str);

    string co2_str = rating(lines,lines.size()-bit_freq[0],bit_freq[0],false);
    unsigned int co2 = convert(co2_str);

    cout << oxy << " * " << co2 << " = " << oxy * co2 << endl;
    
    return EXIT_SUCCESS;
}