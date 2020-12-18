#include <iostream>
#include <functional>
#include <fstream>
#include <string>
#include <vector>
#include <deque>
#include <unordered_set>

using namespace std;

// Structs for 3d and 4 points
struct Point3 {
    int x, y, z;
    Point3(int a, int b, int c) { x = a; y = b; z = c; };
};

struct Point4 {
    int x, y, z, w;
    Point4(int a, int b, int c, int d) { x = a; y = b; z = c; w = d; };
};

// Both need compare operators for hashing
bool operator==(const Point3& l, const Point3& r) {
    return (l.x == r.x) && (l.y == r.y) && (l.z == r.z);
}

bool operator==(const Point4& l, const Point4& r) {
    return (l.x == r.x) && (l.y == r.y) && (l.z == r.z) && (l.w == r.w);
}

// Simple custom hashing for point structs
namespace std {
    template<>
    struct hash<Point3> {
        size_t operator()(const Point3 &p) const {
            size_t h1 = hash<int>{}(p.x);
            size_t h2 = hash<int>{}(p.y);
            size_t h3 = hash<int>{}(p.z);

            return h1 ^ (h2 << 1) ^ (h3 << 2);
        }
    };

    template<>
    struct hash<Point4> {
        size_t operator()(const Point4 &p) const {
            size_t h1 = hash<int>{}(p.x);
            size_t h2 = hash<int>{}(p.y);
            size_t h3 = hash<int>{}(p.z);
            size_t h4 = hash<int>{}(p.w);

            return h1 ^ (h2 << 1) ^ (h3 << 2) ^ (h4 << 3);
        }
    };
}

// Basically overload all these functions to handle both 3d and 4d points
void readFile(string filename, unordered_set<Point3>& space) {
    ifstream infile(filename);
    string line;
    int z = 0;

    for (int y = 0; getline(infile, line); y++) {
        for (int x = 0; x < line.length(); x++)
            if (line[x] == '#')
                space.insert(Point3(x, y, z));
    }
}

void readFile(string filename, unordered_set<Point4>& space) {
    ifstream infile(filename);
    string line;
    int w = 0, z = 0;

    for (int y = 0; getline(infile, line); y++) {
        for (int x = 0; x < line.length(); x++)
            if (line[x] == '#')
                space.insert(Point4(x, y, z, w));
    }
}


// Goes through each neighbor and checks if they are alive, returns true to signify point
// being checked needs to be alive
bool checkNeighbors(Point3 p, unordered_set<Point3>& space, bool active) {
    int count = 0;
    for (int z = p.z - 1; z <= p.z + 1; z++)
        for (int y = p.y - 1; y <= p.y + 1; y++)
            for (int x = p.x - 1; x <= p.x + 1; x++) {
                Point3 temp(x, y, z);
                if (p == temp) continue; // Don't check itself
                if (space.find(temp) != space.end())
                    count++;
                if (count > 3) // neither active nor inactive can be alive w/ 4+
                    return false;
            }
    if (active)
        return count == 3 || count == 2;
    else
        return count == 3;
}

bool checkNeighbors(Point4 p, unordered_set<Point4>& space, bool active) {
    int count = 0;
    for (int w = p.w-1; w <= p.w+1; w++)
        for (int z = p.z-1; z <= p.z+1; z++)
            for (int y = p.y-1; y <= p.y+1; y++)
                for (int x = p.x-1; x <= p.x+1; x++) {
                    Point4 temp(x, y, z, w);
                    if (p == temp) continue;
                    if (space.find(temp) != space.end())
                        count++;
                    if (count > 3)
                        return false;
                }
    if (active)
        return count == 3 || count == 2;
    else
        return count == 3;
}

// This runs one cycle of the life game
void cycle(unordered_set<Point3>& space) {
    unordered_set<Point3> to_check, next_set;

    // Goes through every alive point and adds all their neighbors to the set
    // to be checked
    for (auto point: space) {
        for (int z = point.z-1; z <= point.z+1; z++)
            for (int y = point.y-1; y <= point.y+1; y++)
                for (int x = point.x-1; x <= point.x+1; x++)
                    to_check.insert(Point3(x, y, z));
    }

    // Goes through the to_check set adding the alive points to the next space
    for (auto point: to_check) {
        bool active = space.find(point) != space.end();
        if (checkNeighbors(point, space, active))
            next_set.insert(point);
    }

    space = next_set;
}

void cycle(unordered_set<Point4>& space) {
    unordered_set<Point4> to_check, next_set;
    for (auto point: space) {
        for (int w = point.w-1; w <= point.w+1; w++)
            for (int z = point.z-1; z <= point.z+1; z++)
                for (int y = point.y-1; y <= point.y+1; y++)
                    for (int x = point.x-1; x <= point.x+1; x++)
                        to_check.insert(Point4(x, y, z, w));
    }

    for (auto point: to_check) {
        bool active = space.find(point) != space.end();
        if (checkNeighbors(point, space, active))
            next_set.insert(point);
    }

    space = next_set;
}

// Runs the n number of cycles of either 3d or 4d points
template <typename T>
int runCycles(int cycles, unordered_set<T>& space) {
    for (int i = 0; i < cycles; i++) {
        cycle(space);
    }
    return space.size();
}

// Pass filename as first argument and number of cycles as second
int main(int argc, char *argv[]) {
    if (argc < 3) return EXIT_FAILURE;

    unordered_set<Point3> space3d;
    unordered_set<Point4> space4d;
    readFile(argv[1], space3d);
    readFile(argv[1], space4d);
    int cycles = stoi(argv[2]);

    cout << runCycles(cycles, space3d) << endl;
    cout << runCycles(cycles, space4d) << endl;

    return EXIT_SUCCESS;
}
