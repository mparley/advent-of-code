#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <utility>
#include <cmath>
#include <unordered_set>
#include <boost/functional/hash.hpp>

using namespace std;

void printPaper(const unordered_set<pair<int,int>,boost::hash<pair<int,int>>> &paper) {
    int w = 0, h = 0;
    vector<string> output;
    for (auto dot : paper) {
        int x = dot.first, y = dot.second;

        if (y > h-1) {
            h = y+1;
            output.resize(h);
        }

        if (x > w-1) w = x+1;
        if (w > output[y].size()) output[y].resize(w, ' ');

        output[y][x] = '#';
    }
    
    for (auto line : output)
        cout << line << "\n";
}


void foldPaper(
    unordered_set<pair<int,int>,boost::hash<pair<int,int>>> &paper, 
    const pair<char,int> &fold
) {
    for (auto dot = paper.begin(); dot != paper.end();) {
        int pos[] = { (*dot).first, (*dot).second };
        int i = fold.first - 'x';
        
        if (pos[i] > fold.second) {
            pos[i] = fold.second - abs(pos[i] - fold.second);
            paper.insert({pos[0],pos[1]});
            dot = paper.erase(dot);
        } else
            dot++;
    }
}

int main(int argc, char** argv) {
    vector<pair<char,int>> folds;
    unordered_set<pair<int,int>, boost::hash<pair<int,int>>> dots;
    
    ifstream infile(argv[1]);
    while (infile.good()) {
        string line;
        getline(infile,line);

        if (line[0] == 'f') {
            folds.push_back({line[11],stoi(line.substr(13))});

        } else if (!line.empty()) {
            auto comma = line.find(',');
            int x = stoi(line.substr(0,comma)), y = stoi(line.substr(comma+1));
            dots.insert({x,y});
        }
    }
    infile.close();

    // Part 1
    foldPaper(dots,folds[0]);
    cout << "Part 1:\nSum after first fold: " << dots.size() << "\n";

    // Part 2
    for (int i = 1; i < folds.size(); i++)
        foldPaper(dots, folds[i]);

    cout << "\nPart 2:\n";
    printPaper(dots);

    return EXIT_SUCCESS;
}

