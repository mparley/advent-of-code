#include <iostream>
#include <string>
#include <vector>
#include <sstream>
#include <unordered_set>
#include <unordered_map>

using namespace std;

bool bingo(uint32_t marks, unordered_set<uint32_t> combos) {
    for (auto combo : combos) {
        if ((marks & combo) == combo)
            return true;
    }

    return false;
}

int main(int argc, char** argv) {
    int num;
    string balls, ball;
    unordered_set<uint32_t> bingo_combos;
    vector<vector<int>> cards;
    vector<int> card;

    for (int i = 0; i < 5; i++) {
        uint32_t t = 0;
        for (int j = 0; j < 5; j++) {
            t |= 1 << (j + (5 * i));
        }

        bingo_combos.insert(t);
    }

    for (int i = 0; i < 5; i++) {
        uint32_t t = 0;
        for (int j = 0; j < 5; j++) {
            t |= 1 << (i + (j * 5));
        }

        bingo_combos.insert(t);
    }

    bingo_combos.insert(0x88888);
    
    cin >> balls;
    for (int i = 0; cin >> num; i++) {
        card.push_back(num);
        if ((i != 0) && ((i+1) % 25 == 0)) {
            cards.push_back(card);
            card.clear();
        }
    }
    
    vector<uint32_t> marks(cards.size(),0);
    vector<int> winners;
    unordered_map<int, int> winning_num;
    istringstream iss(balls);

    while (getline(iss, ball, ',')) {
        num = stoi(ball);
        for (int i = 0; i < cards.size(); i++) {
            if (winning_num.find(i) != winning_num.end())
                continue;
            for (int j = 0; j < cards[i].size(); j++)
                if (cards[i][j] == num) {
                    marks[i] |= (1 << j);
                    break;
                }
        }

        for (int i = 0; i < marks.size(); i++) {
            if (winning_num.find(i) != winning_num.end())
                continue;
            if (bingo(marks[i], bingo_combos)) {
                winners.push_back(i);
                winning_num[i] = num;
            }
        }

       // if (!winners.empty())
       //     break;
    }

    cout << "The winning card is card " << winners.front() 
        << " with the number " << winning_num[winners.front()] << "\n";

    int sum = 0;
    for (int i = 0; i < cards[winners.front()].size(); i++) {
        if ((marks[winners.front()] & (1 << i)) != (1 << i))
            sum += cards[winners.front()][i];
    }

    cout << "The winning score is " << sum << " * " << winning_num[winners.front()] 
        << " = " << sum * winning_num[winners.front()] << "\n";


    cout << "The last card to win is card " << winners.back()
        << " with the number " << winning_num[winners.back()] << "\n";

    sum = 0;
    for (int i = 0; i < cards[winners.back()].size(); i++) {
        if ((marks[winners.back()] & (1 << i)) != (1 << i))
            sum += cards[winners.back()][i];
    }

    cout << "The last winning score is " << sum << " * " << winning_num[winners.back()] 
        << " = " << sum * winning_num[winners.back()] << "\n";


    return EXIT_SUCCESS;
}