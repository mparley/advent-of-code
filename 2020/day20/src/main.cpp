#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <bitset>
#include "Tile.hpp"

using namespace std;

unsigned int reverseBits(unsigned int n, int num_bits=10) {
    unsigned int ret = 0;
    for (int i = 0; i < num_bits; i++) {
        ret |= ((n >> i) & 1) << (num_bits-1-i);
    }
    return ret;
}

int main(int argc, char *argv[]) {
    if (argc < 2) return EXIT_FAILURE;

    ifstream infile(argv[1]);
    string line, tile_init;
    int tile_id;
    bool print = false;
    unordered_map<int, Tile*> tiles;
    unordered_map<unsigned int, unordered_set<int>> sides;
    unordered_map<int, int> unique_sides;

    while (infile.good()) {
        getline(infile, line);
        if (line.empty() || !infile.good()) {
            Tile* t = new Tile(tile_id, tile_init);
            tiles[tile_id] = t;
            tile_id = 0;
            tile_init = "";
        } else if (line[0] == 'T') {
            tile_id = stoi(line.substr(5,4));
        } else {
            tile_init += line;
        }
    }

    for (auto tile : tiles) {
        sides[tile.second->top()].insert(tile.first);
        sides[tile.second->right()].insert(tile.first);
        sides[tile.second->bottom()].insert(tile.first);
        sides[tile.second->left()].insert(tile.first);
    }

    for (auto it = sides.begin(); it != sides.end(); ) {
        int rside = reverseBits((*it).first);
        if (sides.find(rside) != sides.end()) {
            sides[rside].merge((*it).second);
            it = sides.erase(it);
        } else
            it++;
    }

    if (print) {
        for (auto side : sides) {
            cout << side.first << ": ";
            for (auto tile : side.second)
                cout << tile << " ";
            cout << "\n";
        }
        cout << "-------------------------\n";
    }

    for (auto side : sides)
        if (side.second.size() == 1)
            unique_sides[*(side.second.begin())]++;

    if (print) {
        for (auto tile : unique_sides) {
            cout << tile.first << ": " << tile.second << "\n";
        }
        cout << "-------------------------\n";
    }

    uint64_t corner_product = 1;
    for (auto tile : unique_sides)
        if (tile.second == 2)
            corner_product *= tile.first;

    for (auto tile: tiles) {
        if (print) {
            cout << tile.second->id() << ":\n";
            int top = tile.second->top(), left = tile.second->left(),
                right = tile.second->right(), bottom = tile.second->bottom();
            bitset<10> btop(top), bleft(left), bright(right), bbottom(bottom);
            cout << "\ttop: " << btop << " (" << top << ")\n";
            cout << "\tright: " << bright << " (" << right << ")\n";
            cout << "\tleft: " << bleft << " (" << left << ")\n";
            cout << "\tbottom: " << bbottom << " (" << bottom << ")\n\n";
        }
        delete tile.second;
    }
    tiles.clear();
    if (print) cout << "-------------------------\n";

    cout << "\nPart 1: " << corner_product << "\n";

    return EXIT_SUCCESS;
}
