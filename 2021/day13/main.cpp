#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <utility>
#include <cmath>
#include <boost/dynamic_bitset.hpp>

using namespace std;

void printPaper(const vector<boost::dynamic_bitset<>> &paper, int mode = 0) {
    for (auto p : paper) {
        if (mode == 0)
            cout << p << "\n";
        else if (mode == 1) {
            for (int i = 0; i < p.size(); i++)
                cout << (p.test(i) ? '#' : ' ');
            cout << "\n";
        }
    }
}


void foldPaper(vector<boost::dynamic_bitset<>> &paper, const pair<char,int> &fold) {
    if (fold.first == 'y') {
        for (int i = fold.second+1; i < paper.size(); i++) {
            int newi = fold.second - abs(i - fold.second);
            paper[newi] = paper[newi] | paper[i];
            paper[i].reset();
        }
        paper.resize(fold.second);

    } else if (fold.first == 'x') {
        for (int i = 0; i < paper.size(); i++) {
            for (int j = fold.second+1; j < paper[i].size(); j++) {
                int newj = fold.second - abs(j - fold.second);
                paper[i][newj] |= paper[i][j];
            }
            paper[i].resize(fold.second);
        }
    }
}

int main(int argc, char** argv) {
    int width = 0, height = 0;
    string line;
    vector<pair<int,int>> dots;
    vector<pair<char,int>> folds;
    
    ifstream infile(argv[1]);
    while (getline(infile, line)) {
        if (line[0] == 'f') {
            folds.push_back({line[11],stoi(line.substr(13))});

        } else if (!line.empty()) {
            auto comma = line.find(',');
            int x = stoi(line.substr(0,comma)), y = stoi(line.substr(comma+1));
            if (x+1 > width) width = x+1;
            if (y+1 > height) height = y+1;
            dots.push_back({x,y});
        }
    }
    infile.close();

    // Populate paper with dots
    vector<boost::dynamic_bitset<>> paper;
    for (int i = 0; i < height; i++)
        paper.push_back(boost::dynamic_bitset<>(width, 0));

    for (auto dot : dots) {
        paper[dot.second].set(dot.first);
    }

    // Part 1
    foldPaper(paper,folds[0]);
    int sum = 0;
    for (auto p : paper)
        sum += p.count();

    cout << "Part 1:\nSum after first fold: " << sum << "\n";


    // Part 2
    for (auto it = folds.begin()+1; it != folds.end(); it++) {
        foldPaper(paper, *it);
    }

    cout << "\nPart 2:\n";
    printPaper(paper,1);

    return EXIT_SUCCESS;
}
