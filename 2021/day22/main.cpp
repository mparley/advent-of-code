#include <iostream>
#include <fstream>
#include <vector>
#include <utility>
#include <unordered_set>
#include <algorithm>
#include <sstream>

using namespace std;

struct Point3 {
    int x, y, z;

    Point3() : x(0), y(0), z(0) {}
    Point3(int x, int y, int z) : x(x), y(y), z(z) {}

  bool operator==(const Point3& r) const {
      return (x == r.x) && (y == r.y) && (z == r.z);
}
};

struct Cube {
  Point3 mn;
  Point3 mx;

  Cube(Point3 from, Point3 to) {
    mn = from;
    mx = to;
  }

  bool Intersects(const Cube& o) const {
    return mn.x < o.mx.x && o.mn.x < mx.x &&
           mn.y < o.mx.y && o.mn.y < mx.y &&
           mn.z < o.mx.z && o.mn.z < mx.z;  
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

  long volume() const {
    long v = (mx.x - mn.x + 1) * (mx.y - mn.y + 1) * (mx.z - mn.z + 1);
    return v;
  }

  bool operator==(const Cube& o) const {
    return mn == o.mn && mx == o.mx;
  }


  void SplitBy(const Cube& i, vector<Cube>& pieces) const {
    // Cube i = Intersection(o);
    if (mx.x > i.mx.x) {
      Point3 from = mx;
      from.x = i.mx.x;
      pieces.push_back({ from, mx });
    }

    if (mx.y > i.mx.y) {
      Point3 from = mx;
      from.y = i.mx.y;
      pieces.push_back({ from, mx });
    }

    if (mx.z > i.mx.z) {
      Point3 from = mx;
      from.z = i.mx.z;
      pieces.push_back({ from, mx });
    }

    if (mn.x < i.mn.x) {
      Point3 to = mn;
      to.x = i.mn.x;
      pieces.push_back({ mn, to });
    }

    if (mn.y < i.mn.y) {
      Point3 to = mn;
      to.y = i.mn.y;
      pieces.push_back({ mn, to });
    }

    if (mn.z < i.mn.z) {
      Point3 to = mn;
      to.z = i.mn.z;
      pieces.push_back({ mn, to });
    }
  }
};

ostream& operator<<(ostream& o, const Cube& c) {
  o << c.mn.x << "," << c.mn.y << "," << c.mn.z << "->" 
    << c.mx.x << "," << c.mx.y << "," << c.mx.z;
  return o;
}

bool CubeInRange(Cube c, Point3 from, Point3 to) {
  Cube c2(from, to);
  return c2.Contains(c);
}

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
    struct hash<Cube> {
        size_t operator()(const Cube &c) const {
            size_t h1 = hash<Point3>{}(c.mn);
            size_t h2 = hash<Point3>{}(c.mx);

            return h1 ^ (h2 << 1);
        }
    };
}

int main(int argc, char** argv) {
  ifstream infile(argv[1]);
  string line;

  vector<Cube> in_cubes;
  vector<string> instruction;
  unordered_set<Cube> on_set;
  unordered_set<Cube> off_set;

  while (getline(infile, line)) {
    Point3 from(0,0,0);
    Point3 to(0,0,0);
    stringstream ss(line);
    string ins;

    ss >> ins >> from.x >> to.x >> from.y >> to.y >> from.z >> to.z;

    in_cubes.push_back({from, to});
    instruction.push_back(ins);
  }
  infile.close();

  cout << instruction.size() << "\n";
  for (int i = 0; i < instruction.size(); i++) {
    if (!CubeInRange(in_cubes[i], Point3(-50,-50,-50), Point3(50,50,50))) continue;
    cout << "INSTRUCTION: " << i << "\t";
    Cube cube = in_cubes[i];
    vector<Cube> pieces;

    if (instruction[i] == "on") {
      cout << "ON:\n";
      // on_set.insert(cube);
      // cout << on_set.size() << "\n";

      bool add = true;
      for (auto it = on_set.begin(); it != on_set.end();) {
        if (cube.Contains(*it)) {
          it = on_set.erase(it);
        } else if ((*it).Intersects(cube)) {
          (*it).SplitBy(cube,pieces);
          // cube.SplitBy(*it, pieces);
          // pieces.push_back((*it).Intersection(cube));
          it = on_set.erase(it);
        } else {
          it++;
        }
      }

      if (add) on_set.insert(cube);

      for (auto piece : pieces) {
        on_set.insert(piece);
      }

      long total = 0;
      for (auto& c : on_set) {
        total += c.volume();
        cout << c << endl;
      }
      cout << total << endl;



    } else if (instruction[i] == "off") {
      cout << "OFF:\n";
      for (auto it = on_set.begin(); it != on_set.end();) {
        if (*it == cube) break;
        else if (cube.Contains(*it)) {
          it = on_set.erase(it);
        } else if ((*it).Intersects(cube)) {
          (*it).SplitBy(cube,pieces);
          it = on_set.erase(it);
        } else
          it++;
      }

      for (auto piece : pieces)
        on_set.insert(piece);

      long total = 0;
      for (auto& c : on_set) {
        total += c.volume();
        cout << c << endl;
      }
      cout << total << endl;

      // cout << on_set.size() << "\n";

    } else
      cout << "What?\n";
  }

  unordered_set<Point3> on_points;
  for (auto c : on_set) {
    for (int i = c.mn.x; i <= c.mx.x; i++)
    for (int j = c.mn.y; j <= c.mx.y; j++)
    for (int k = c.mn.z; k <= c.mx.z; k++)
      on_points.insert(Point3(i,j,k));
  }

  cout << on_points.size() << endl;


  return EXIT_SUCCESS;
}