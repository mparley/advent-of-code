// Have to comment as I barely wrapped my head around it myself

// Disclaimer
// I didn't come up with this without help from a couple friends
// so I guess technically I'm filtered but it is kind of similar to
// the fish problem. Not sure I would have finished this one without
// the advice and heavy hints I got :(

#include <iostream>
#include <unordered_map>
#include <map>

using namespace std;

// Mapping all possible combinations of 3 rolls to the number of
// universes each combination of rolls will create
const unordered_map<int,int> DiracCounts = {
  {3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1},
};

struct BoardState {
  int pos[2];
  int score[2];
  bool turn;

  BoardState(int p1, int p2) {
    pos[0] = p1;
    pos[1] = p2;
    score[0] = 0;
    score[1] = 0;
    turn = false;
  }

  // Need this to use a key in map
  bool operator<(const BoardState& r) const {
    return (pos[0] < r.pos[0]) 
    || (pos[0] == r.pos[0] && pos[1] < r.pos[1])
    || (pos[0] == r.pos[0] && pos[1] == r.pos[1] && score[0] < r.score[0])
    || (pos[0] == r.pos[0] && pos[1] == r.pos[1] && score[0] == r.score[0] 
        && score[1] < r.score[1])
    || (pos[0] == r.pos[0] && pos[1] == r.pos[1] && score[0] == r.score[0] 
        && score[1] == r.score[1] && turn < r.turn);
  }
};

struct WinCount {
  uint64_t player[2];

  WinCount() {
    player[0] = 0;
    player[1] = 0;
  }

  WinCount(uint64_t p1, uint64_t p2) {
    player[0] = p1;
    player[1] = p2;
  }

  // For adding the result of a Dirac call times the amount
  // of universes giving that result to the total
  void Add(const WinCount& r, uint64_t c) {
    player[0] += r.player[0] * c;
    player[1] += r.player[1] * c;
  }
};

map<BoardState,WinCount> BoardCache;


// Meat of part 2 solution. Recursive function that takes
// a board state, progresses it and makes calls on the next
// state for all the universes it created
WinCount Dirac(BoardState b)
{
  if (BoardCache.find(b) != BoardCache.end())
    return BoardCache[b];

  WinCount wins(0,0);

  // Goes through all the created universes per DiraCounts map
  for (auto [roll, count] : DiracCounts) {
    BoardState next = b;
    next.turn = !b.turn;
    next.pos[b.turn] = ((next.pos[b.turn] + roll - 1) % 10) + 1;
    next.score[b.turn] += next.pos[b.turn];

    // If the next board has a win add the universes for this roll
    // to the output winner, else make the recursive call and add
    // the results to the output.
    if (next.score[b.turn] >= 21)
      wins.player[b.turn] += count;
    else
      wins.Add(Dirac(next), count); // universes compound see Win.Add()
  }

  // Add result of current state to cache and return
  BoardCache[b] = wins;
  return wins;
}

struct DDie {
  int rolls;
  int roll;

  DDie() : rolls(0), roll(1) {}

  int Roll() {
    rolls++;
    roll = ((roll) % 10) + 1;
    return roll;
  }

  int Roll(int times) {
    int total = 0;
    for (int i = 0; i < times; i++)
      total += Roll();
    return total;
  }
};

int Part1(BoardState b) {
  DDie d;

  do {
    b.pos[b.turn] = ((b.pos[b.turn] + d.Roll(3) - 1) % 10) + 1;
    b.score[b.turn] += b.pos[b.turn];
    b.turn = !b.turn;
  } while (b.score[!b.turn] < 1000);

  return d.rolls * min(b.score[0],b.score[1]);
}

uint64_t Part2(BoardState b) {
  WinCount w = Dirac(b);
  return min(w.player[0], w.player[1]);
}

int main(int argc, char** argv) {
  BoardState b(stoi(argv[1]), stoi(argv[2]));

  cout << Part1(b) << "\n";
  cout << Part2(b) << "\n";

  return EXIT_SUCCESS;
}