#include <iostream>
#include <string>
#include <utility>
#include <unordered_map>
#include <vector>
#include <boost/functional/hash.hpp>

using namespace std;

struct Line {
    int x1, y1, x2, y2;

    bool hor() {
        return y1 == y2;
    }

    bool vert() {
        return x1 == x2;
    }

    friend ostream &operator<<(ostream &out, const Line &L) {
        out << L.x1 << "," << L.y1 << " -> " << L.x2 << "," << L.y2;
        return out;
    }

    friend istream &operator>>(istream &in, Line &L) {
        string delim;
        in >> L.x1; in.get();
        in >> L.y1 >> delim >> L.x2; in.get();
        in >> L.y2;
        return in;
    }
};

int main(int argc, char* argv[]) {
    Line l;
    vector<Line> lines;
    unordered_map<pair<int,int>,int,boost::hash<pair<int,int>>> points;

    while (cin >> l) {
        lines.push_back(l);
        //cout << l << "\n";
    }

    for (auto line : lines) {
        //if (!(line.hor() || line.vert())) continue;

        int x1 = line.x1, y1 = line.y1;
        int x2 = line.x2, y2 = line.y2;
        int xd = 1, yd = 1;

        if (x1 > x2) xd = -1;
        if (y1 > y2) yd = -1;

        while (x1 != x2 || y1 != y2) {
            points[{x1,y1}]++;
            if (x1 != x2) x1 += xd;
            if (y1 != y2) y1 += yd;
        }
        points[{x2,y2}]++;
    }

    int count = 0;
    for (auto point : points) {
        if (point.second >= 2)
            count++;
    }

    cout << count << "\n";

    return EXIT_SUCCESS;
}