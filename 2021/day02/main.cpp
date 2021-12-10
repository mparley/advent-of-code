#include <iostream>
#include <string>

int main() {

    std::string command;
    int value, hpos = 0, depth1 = 0, depth2 = 0, aim = 0;

    while(std::cin >> command) {
        std::cin >> value;

        if (command == "forward") {
            hpos += value;
            depth2 += aim * value;
        } else {
            if (command == "up") value *= -1;
            depth1 += value;
            aim += value;
        }
    }

    std::cout << "Part 1\n";
    std::cout << "Horizontal position: " << hpos << "\n";
    std::cout << "Depth: " << depth1 << "\n";
    std::cout << "Solution: " << depth1 * hpos << "\n\n";

    std::cout << "Part 2\n";
    std::cout << "Horizontal position: " << hpos << "\n";
    std::cout << "Depth: " << depth2 << "\n";
    std::cout << "Solution: " << depth2 * hpos << std::endl;

    return EXIT_SUCCESS;
}