CPP = g++
CFLAGS = -O2 -s -w -I.

all: clean compute

compute: compute.cpp
	@$(CPP) $(CFLAGS) $< -o $@ $(LFLAGS)
	
clean:
	@rm -fr compute
