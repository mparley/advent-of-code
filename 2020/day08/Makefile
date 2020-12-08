SRC = $(wildcard ./*cpp)
OBJ = $(SRC:.cpp=.o)
CC = g++
CXXFLAGS = -w -std=c++17 -O3
NAME = a.out

$(NAME): $(OBJ)
	$(CC) -o $@ $^

.PHONY: clean
clean:
	rm -f $(OBJ) $(NAME)
