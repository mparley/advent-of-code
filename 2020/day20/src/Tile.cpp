#include "Tile.hpp"

Tile::Tile(int id, std::string init_str) {
    id_ = id;
    for (int i = 0; i < init_str.length(); i++)
        pixels_[i/10][i%10] = init_str[i];
}

Tile::~Tile() {}

int Tile::id() { return id_; }

unsigned int Tile::top() {
    unsigned int out = 0;
    for (int i = 0; i < 10; i++) {
        out |= (pixels_[0][i] == '#');
        if (i < 9) out = out << 1;
    }
    return out;
}

unsigned int Tile::bottom() {
    unsigned int out = 0;
    for (int i = 0; i < 10; i++) {
        out |= (pixels_[9][i] == '#');
        if (i < 9) out = out << 1;
    }
    return out;
}

unsigned int Tile::left() {
    unsigned int out = 0;
    for (int i = 0; i < 10; i++) {
        out |= (pixels_[i][0] == '#');
        if (i < 9) out = out << 1;
    }
    return out;
}

unsigned int Tile::right() {
    unsigned int out = 0;
    for (int i = 0; i < 10; i++) {
        out |= (pixels_[i][9] == '#');
        if (i < 9) out = out << 1;
    }
    return out;
}
