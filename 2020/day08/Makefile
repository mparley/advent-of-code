CC := g++
SRCDIR := src
BUILDDIR := build
TARGET := a.out
 
SOURCES := $(shell find $(SRCDIR) -type f -name *.cpp)
OBJECTS := $(patsubst $(SRCDIR)/%,$(BUILDDIR)/%,$(SOURCES:.cpp=.o))
CFLAGS := -w -std=c++17 -O3
INC := -I .

$(TARGET): $(OBJECTS)
	$(CC) $^ -o $(TARGET)

$(BUILDDIR)/%.o: $(SRCDIR)/%.$(SRCEXT)
	@mkdir -p build
	$(CC) $(CFLAGS) $(INC) -c -o $@ $<

clean:
	rm -r $(BUILDDIR) $(TARGET);

.PHONY: clean
