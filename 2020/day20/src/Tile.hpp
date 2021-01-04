#ifndef __TILE_H_
#define __TILE_H_

#include <string>

class Tile {
    private:
        char pixels_[10][10];
        int id_;

    public:
        Tile(int id, std::string init_str);
        ~Tile();

        int id();
        unsigned int top();
        unsigned int bottom();
        unsigned int left();
        unsigned int right();
};

#endif // __TILE_H_
