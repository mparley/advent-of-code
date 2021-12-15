#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <utility>
#include <climits>
#include <queue>
#include <unordered_map>
#include <unordered_set>
#include <boost/functional/hash.hpp>

using namespace std;

struct Point {
    uint x, y;

    Point(uint ix, uint iy)
        : x(ix), y(iy) {}

    bool operator==(const Point &o) const {
        return (x == o.x && y == o.y);
    }
};

struct PointHash {
    size_t operator()(Point const& p) const noexcept {
        boost::hash<pair<uint,uint>> pair_hash;
        return pair_hash({p.x, p.y});
    }
};

class PointComp {

    unordered_map<Point,uint,PointHash>* dist_map;

public:

    PointComp(unordered_map<Point,uint,PointHash> *dists) {
        dist_map = dists;
    }

    bool operator()(const Point &l, const Point &r) const { 
        return dist_map->at(l) > dist_map->at(r);
    }
};

vector<Point> getNeighbors(Point p, uint tx, uint ty) {
    vector<Point> out;

    if (p.x > 0) out.push_back(Point(p.x-1, p.y));
    if (p.y > 0) out.push_back(Point(p.x, p.y-1));
    if (p.x < tx) out.push_back(Point(p.x+1, p.y));
    if (p.y < ty) out.push_back(Point(p.x, p.y+1));

    return out;
}

vector<vector<uint>> parseInput(string fname) {
    ifstream infile(fname);
    vector<vector<uint>> cave;

    for (string line; getline(infile,line);) {
        vector<uint> iline;
        for (int j = 0; j < line.size(); j++)
            iline.push_back(line[j]-'0');
        cave.push_back(iline);
    }
    infile.close();

    return cave;
}

uint risk(const vector<vector<uint>> &cave, Point p) {
    uint height = cave.size(), width = (*cave.begin()).size();
    uint frame = (p.x / width) + (p.y / height);
    uint cx = p.x % width, cy = p.y % height;
    return (cave[cy][cx] - 1 + frame) % 9 + 1;
}


int main(int argc, char* argv[]) {
    vector<vector<uint>> cave = parseInput(argv[1]);
    unordered_map<Point,uint,PointHash> rdists = { { Point(0,0), 0 }};
    unordered_set<Point,PointHash> visited;
    priority_queue<Point,vector<Point>,PointComp> pq((PointComp(&rdists)));
    uint factor = (argc > 2) ? stoi(argv[2]) : 1;
    uint target_x = (cave[0].size() * factor) - 1;
    uint target_y = (cave.size() * factor) - 1;

    pq.push({0,0});

    while (!pq.empty()) {
        Point u = pq.top();
        visited.insert(u);
        pq.pop();

        for (auto v : getNeighbors(u, target_x, target_y)) {
            if (visited.find(v) != visited.end()) continue;

            if (rdists.find(v) == rdists.end())
                rdists[v] = UINT_MAX;

            uint alt = rdists[u] + risk(cave,v);
            if (alt < rdists[v]) {
                rdists[v] = alt;
                pq.push(v);
            }
        }
    }

    cout << rdists[Point(target_x,target_y)] << endl;

    return EXIT_SUCCESS;
}
