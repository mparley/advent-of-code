#include <iostream>
#include <fstream>
#include <vector>
#include <climits>

using namespace std;

enum ReadState { START, LITERAL, OP };

struct Packet {
  uint8_t version, type_id; 
  uint16_t op_len;
  uint64_t value, length;
  bool op_mode;

  vector<Packet*> children;

  Packet(uint8_t v, uint8_t t)
    : version(v), type_id(t), value(0), op_len(0),
    length(6), op_mode(false) {}
};

uint64_t BitSlice(vector<bool> &vb, size_t &pos, int num) {
  uint64_t out = 0;
  for (int i = 0; i < num; i++)
    out |= (vb[pos+i] << (num - i - 1));
  pos += num;
  return out;
}

void ReadFile(string filename, vector<bool> &vb) {
  ifstream infile(filename);
  
  char c;
  while (infile >> c) {
    uint8_t val = (c >= 'A') ? (c - 'A') + 10 : c - '0';
    for (int i = 3; i >= 0; i--) {
      vb.push_back(val & (1 << i));
    }
  }
}

uint64_t EvaluateOP(uint8_t type_id, vector<Packet*> &children) {
  uint64_t val = 0;

  switch (type_id) {
    case 0: 
      for (auto c : children) val += c->value;
      return val;

    case 1:
      val = 1;
      for (auto c : children) val *= c->value;
      return val;

    case 2:
      val = UINT64_MAX;
      for (auto c : children)
        if (c->value < val) val = c->value;
      return val;

    case 3:
      for (auto c : children)
        if (c->value > val) val = c->value;
      return val;

    case 5:
      return children[0]->value > children[1]->value;

    case 6:
      return children[0]->value < children[1]->value;

    case 7:
      return children[0]->value == children[1]->value;

    default:
      return 0;
  }
}

Packet* ReadPackets(vector<Packet*> &pv, vector<bool> &vb, size_t &i) {
  ReadState rs = START;
  Packet* p = new Packet(0,0);
  size_t last_i = i;

  while (i < vb.size()) {

    switch(rs) {
      case START: {
        p->version = BitSlice(vb, i, 3);
        p->type_id = BitSlice(vb, i, 3);
        rs = (p->type_id == 4) ? LITERAL : OP;
        break;
      }

      case LITERAL: {
        vector<uint64_t> parts;
        while(vb[i])
          parts.push_back(BitSlice(vb,i,5));
        parts.push_back(BitSlice(vb,i,5));

        uint64_t value = 0;
        for (int j = 0; !parts.empty(); j++) {
          value |= (parts.back() & 0xF) << (j * 4);
          parts.pop_back();
        }

        p->value = value;
        p->length = i - last_i;
        pv.push_back(p);
        return p;
      }

      case OP: {
        bool op_mode = vb[i];
        i++;

        uint16_t op_len = 0;
        if (op_mode) {
          op_len = BitSlice(vb,i,11);
          while (p->children.size() < op_len)
            p->children.push_back(ReadPackets(pv, vb, i));

        } else {
          op_len = BitSlice(vb,i,15);
          size_t op_i = i;
          while (op_len > i - op_i)
            p->children.push_back(ReadPackets(pv, vb, i));
        }

        p->op_mode = op_mode;
        p->op_len = op_len;
        p->value = EvaluateOP(p->type_id, p->children);
        p->length = i - last_i;
        pv.push_back(p);
        return p;
      }

      default:
        return p;
    }
  }
}

int main(int argc, char* argv[]) {

  vector<bool> vb;
  ReadFile(argv[1], vb);

  size_t i = 0;
  vector<Packet*> pv;
  ReadPackets(pv, vb, i);

  uint64_t version_sum = 0;
  for (auto p : pv)
    version_sum += p->version;

  cout << version_sum << "\n";
  cout << pv.back()->value << "\n";

  for (auto p : pv)
    delete p;
  pv.clear();

  return EXIT_SUCCESS;
}