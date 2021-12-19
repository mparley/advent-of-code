#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <queue>

using namespace std;

struct SnailPair {
  int val;
  SnailPair *l, *r;

  SnailPair() : val(-1), l(nullptr), r(nullptr) {}
  SnailPair(int val) : val(val), l(nullptr), r(nullptr) {}
  SnailPair(SnailPair *l, SnailPair *r) : val(-1), l(l), r(r) {}
};

struct SnailTree {
  SnailPair *root;

  SnailTree() : root(nullptr) {}

  SnailTree(string s) {
    root = ParseString(s);
  }

  SnailPair* ParseString(string s) {
    if (s.size() == 1)
      return new SnailPair(stoi(s));

    int open_count = 1, comma = 1;
    for (; comma < s.size()-1; comma++) {
      if (s[comma] == ',' && open_count == 1) break;
      else if (s[comma] == '[') open_count++;
      else if (s[comma] == ']') open_count--;
    }

    return new SnailPair(
      ParseString(s.substr(1,comma-1)), 
      ParseString(s.substr(comma+1,s.size()-1-(comma+1))));
  }

  void PrintTree(SnailPair* p) {
    if (p->val != -1) {
      cout << p->val;
      return;
    }

    cout << "[";
    PrintTree(p->l);
    cout << ",";
    PrintTree(p->r);
    cout << "]";
    return;
  }

  void PrintTree() {
    PrintTree(root);
    cout << "\n";
  }

  void clean(SnailPair *p) {
    if (p->l != nullptr) {
      clean(p->l);
      delete p->l;
      p->l = nullptr;
    }

    if (p->r != nullptr) {
      clean(p->r);
      delete p->r;
      p->r = nullptr;
    }

    return;
  }

  ~SnailTree() {
    clean(root);
    delete root;
  }

  void AddPair(SnailPair *p) {
    if (root == nullptr) {
      root = p;
      return;
    }

    SnailPair *new_root = new SnailPair(root,p);
    root = new_root;

    while (true) {
      vector<SnailPair*> v;
      if (Explode(v)) continue;
      if (Split(v)) continue;
      break;
    }
  }

  void AddPair(string sp) {
    AddPair(ParseString(sp));
  }

  void ExplodeHelper(SnailPair *p, vector<SnailPair*> &v, 
    bool &found, int depth = 1)
  {
    if (p->l == nullptr && p->r == nullptr) {
      v.push_back(p);
      return;
    }

    // When finding pair at depth, add the parent to v and skip the nums
    if (!found && depth > 4 && p->l->val != -1 && p->r->val != -1) {
      v.push_back(p);
      found = true;
      return;
    }

    if (p->l != nullptr)
      ExplodeHelper(p->l, v, found, depth+1);
    if (p->r != nullptr)
      ExplodeHelper(p->r, v, found, depth+1);
  }

  bool Explode(vector<SnailPair*> &v) {
    bool found = false;
    ExplodeHelper(root, v, found);
    
    if (!found) return false;

    int i = 0;
    for (; i < v.size(); i++)
      if (v[i]->val == -1) break;

    if (i > 0) v[i-1]->val += v[i]->l->val;
    if (i < v.size()-1) v[i+1]->val += v[i]->r->val;

    clean(v[i]);
    v[i]->val = 0;
    return true;
  }

  bool Split(vector<SnailPair*> &v) {
    SnailPair* p = nullptr;
    for (auto leaf : v)
      if (leaf->val >= 10) {
        p = leaf;
        break;
      }

    if (p == nullptr) return false;

    p->l = new SnailPair(p->val / 2);
    p->r = new SnailPair((p->val / 2) + (p->val % 2));
    p->val = -1;
    return true;
  }

  int MagnitudeHelper(SnailPair* p) {
    if (p->val != -1) return p->val;
    return (3 * MagnitudeHelper(p->l)) + (2 * MagnitudeHelper(p->r));
  }

  int magnitude() { return MagnitudeHelper(root); }
};

int main(int argc, char* argv[]) {
  string line;
  ifstream infile(argv[1]);
  vector<string> inputs;
  SnailTree st;

  while (getline(infile,line)) {
    inputs.push_back(line);
    st.AddPair(line);
  }

  cout << st.magnitude() << "\n";

  int mag = 0;
  for (int i = 0; i < inputs.size(); i++) {
    for (int j = 0; j < inputs.size(); j++) {
      if (j == i) continue;
      SnailTree st2(inputs[i]);
      st2.AddPair(inputs[j]);
      int curr = st2.magnitude();
      if (mag < curr) mag = curr;
    }
  }

  cout << mag << "\n";

  return EXIT_SUCCESS;
}