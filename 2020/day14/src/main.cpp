#include <iostream>
#include <fstream>
#include <string>
#include <unordered_map>
#include <bitset>
#include <vector>

using namespace std;

enum class IType {MASK, MEM};

// Struct for each input line, for simplicity it can be either mask or mem
struct Instruction {
    IType type;
    uint64_t mem_loc;
    uint64_t value;
    string mask;

    Instruction(uint64_t m, uint64_t v) {
        type = IType::MEM; mem_loc = m; value = v;
    };

    Instruction(string m) {
        type = IType::MASK; mask = m; mem_loc = 0; value = 0;
    };
};

// Part 1 solution
uint64_t initV1 (const vector<Instruction>& instructions) {
    bitset<36> keep(0), mask(0);
    unordered_map<uint64_t, uint64_t> memory;

    for (auto instruction : instructions) {

        // If instruction is a mask, reset bitsets and reconstruct them
        if (instruction.type == IType::MASK) {
            keep.reset(), mask.reset();
            for (int i = 0; i < instruction.mask.length(); i++) {
                int bit_pos = instruction.mask.length() - 1 - i; // reversed
                if (instruction.mask[i] == 'X')
                    keep.set(bit_pos);
                else
                    mask[bit_pos] = instruction.mask[i] == '1';
            }

        // If instruction is MEM then apply the mask to the value and store
        } else {
            bitset<36> value(instruction.value);
            value &= keep; // & with 'keep' to make the bits 'mask' changes 0
            value |= mask;
            memory[instruction.mem_loc] = value.to_ullong();
        }
    }

    uint64_t sum = 0;
    for (auto addr : memory) {
        sum += addr.second;
    }

    return sum;
}

// Recursively set all possible memory locations with the value
void setFloatMemory(bitset<36> f_bits, int f_pos,
                    unordered_map<uint64_t, uint64_t>& mem,
                    bitset<36> mem_loc, uint64_t value)
{
    // We reached the end when there are no more 'floating' bits to check
    if (f_bits.none()) {
        mem[mem_loc.to_ullong()] = value;
        return;
    }

    // Start at fpos so we don't have to start again. Go through the f_bit set
    // until we see a floating position, then we make two recursive calls for
    // the two branches of memory locations
    for (int i = f_pos; i < f_bits.size(); i++) {
        if (f_bits.test(i)) {
            f_bits.reset(i);
            setFloatMemory(f_bits, i+1, mem, mem_loc, value);
            mem_loc.flip(i);
            setFloatMemory(f_bits, i+1, mem, mem_loc, value);
            break; // We don't want to continue since it's handled by recursion
        }
    }

    return;
}

// Part 2's solution similar to part 1 except the memory is changed, and we
// don't need a keep bitset since a 0 in mask will leave an unchanged bit anyway
uint64_t initV2 (const vector<Instruction>& instructions) {
    bitset<36> mask(0), floating(0);
    unordered_map<uint64_t, uint64_t> memory;

    for (auto instruction : instructions) {

        // If MASK instruction same as part 1 with floating instead of keep
        // zero out bitsets and construct them based on the mask string
        if (instruction.type == IType::MASK) {
            mask.reset(); floating.reset();
            for (int i = 0; i < instruction.mask.length(); i++) {
                int bit_pos = instruction.mask.length() - 1 - i;
                if (instruction.mask[i] == 'X')
                    floating.set(bit_pos);
                else
                    mask[bit_pos] = instruction.mask[i] == '1';
            }

        // Or the memory_location with the mask and pass it in to setFloatMemory
        // to recursively set all possible memory locations
        } else {
            bitset<36> mem(instruction.mem_loc);
            mem |= mask;
            setFloatMemory(floating, 0, memory, mem, instruction.value);
        }
    }

    uint64_t sum = 0;
    for (auto addr : memory) {
        sum += addr.second;
    }

    return sum;
}

// Helper function to read the file and populate a vector of instructions
void readFile(string filename, vector<Instruction>& instructions) {
    ifstream infile(filename);
    string line;

    while (getline(infile, line)) {
        if (line[1] == 'a') {
            string mask = line.substr(line.length()-36, 36);
            instructions.emplace_back(Instruction(mask));
        } else {
            int pos = line.find("[");
            uint64_t mem_loc = stoull(line.substr(pos + 1));
            pos = line.find("=");
            uint64_t value =  stoull(line.substr(pos+2));
            instructions.emplace_back(Instruction(mem_loc, value));
        }
    }
}

// Filename is passed as first argument
int main(int argc, char *argv[]) {
    if (argc < 2) return EXIT_FAILURE;

    vector<Instruction> instructions;
    readFile(argv[1], instructions);

    cout << initV1(instructions) << "\n";
    cout << initV2(instructions) << endl;

    return EXIT_SUCCESS;
}
