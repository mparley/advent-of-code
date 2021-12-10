#include <iostream>
#include <deque>

int main() {

  std::deque<int> depth_window;
  int depth, c1 = 0, c2 = 0, prev_sum = 0;
  bool first_window = true;
  std::cin >> depth;
  depth_window.push_front(depth);

  while (std::cin >> depth) {
    if (depth > depth_window.front()) c1++;

    depth_window.push_front(depth);
    if (depth_window.size() > 3)
      depth_window.pop_back();
    
    if (depth_window.size() == 3) {
      int sum = 0;
      for (auto d : depth_window)
        sum += d;

      //std::cout << sum << " " << prev_sum << std::endl;

      if (first_window)
        first_window = false;
      else
        if (sum > prev_sum) c2++;

      prev_sum = sum;
    }
  }

  std::cout << c1 << std::endl;
  std::cout << c2 << std::endl;

  return EXIT_SUCCESS;
}
