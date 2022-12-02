#include <iostream>
#include <fstream>
#include <vector>
#include <utility>
#include <unordered_set>
#include <algorithm>
#include <sstream>

using namespace std;

struct Point3 {
    long long x, y, z;

    Point3() : x(0), y(0), z(0) {}
    Point3(long long x, long long y, long long z) : x(x), y(y), z(z) {}

  bool operator==(const Point3& r) const {
      return (x == r.x) && (y == r.y) && (z == r.z);
}
};

struct Cube {
  Point3 mn;
  Point3 mx;
  bool pos;

  Cube(Point3 from, Point3 to) {
    mn = from;
    mx = to;
    pos = true;
  }

  Cube(Point3 from, Point3 to, bool p) {
    mn = from;
    mx = to;
    pos = p;
  }

  pair<bool, Cube> GetOverlap(const Cube& o) {
    bool ret = true;
    Point3 from(max(mn.x, o.mn.x), max(mn.y, o.mn.y), max(mn.z, o.mn.z));
    Point3 to(min(mx.x, o.mx.x), min(mx.y, o.mx.y), min(mx.z, o.mx.z));
    if (from.x > to.x || from.y > to.y || from.z > to.z) ret = false;
    return {ret, Cube(from, to, true)};
  }

  bool Intersects(const Cube& o) const {
    return mn.x <= o.mx.x && o.mn.x <= mx.x &&
           mn.y <= o.mx.y && o.mn.y <= mx.y &&
           mn.z <= o.mx.z && o.mn.z <= mx.z;  
  }

  bool Contains(const Cube& o) const {
    return mn.x <= o.mn.x && o.mx.x <= mx.x &&
           mn.y <= o.mn.y && o.mx.y <= mx.y &&
           mn.z <= o.mn.z && o.mx.z <= mx.z;  
  }

  Cube Intersection(const Cube& o) const {
    Point3 from(max(mn.x, o.mn.x), max(mn.y, o.mn.y), max(mn.z, o.mn.z));
    Point3 to(min(mx.x, o.mx.x), min(mx.y, o.mx.y), min(mx.z, o.mx.z));
    return Cube(from, to);
  }

  long long volume() const {
    return (mx.x - mn.x + 1) * (mx.y - mn.y + 1) * (mx.z - mn.z + 1);
  }
};


bool CubeInRange(Cube c, Point3 from, Point3 to) {
  Cube c2(from, to);
  return c2.Contains(c);
}

int main(int argc, char** argv) {
  ifstream infile(argv[1]);
  string line;

  vector<Cube> inputs;
  vector<Cube> on_set;
  vector<Cube> off_set;

  while (getline(infile, line)) {
    Point3 from(0,0,0);
    Point3 to(0,0,0);
    stringstream ss(line);
    string ins;

    ss >> ins >> from.x >> to.x >> from.y >> to.y >> from.z >> to.z;

    bool p = (ins == "on");
    inputs.push_back({from, to, p});
  }
  infile.close();

  vector<Cube> cubes;
  for (auto input : inputs) {
    // if (!CubeInRange(in_cubes[i], Point3(-50,-50,-50), Point3(50,50,50))) continue;
    vector<Cube> to_add;

    for (auto c : cubes) {
      if (c.Intersects(input)) {
        Cube inter = c.Intersection(input);
        inter.pos = !c.pos;
        to_add.push_back(inter);
      }
    }

    cubes.insert(cubes.end(), to_add.begin(), to_add.end());

    if (input.pos)
      cubes.push_back(input);
  }

  long long total = 0;
  for (auto c : cubes) {
    if (c.pos)
      total += c.volume();
    else
      total -= c.volume();
  }
  cout << total << endl;


  return EXIT_SUCCESS;
}