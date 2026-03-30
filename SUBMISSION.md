# Submission Notes
My submission contains three implementations of the same solution: Python and Golang inside a Docker containers, and a standalone C++ implementation.
In the implementations, I parsed all the input values (i.e. command line arguments, and the values from stdin) and performed the computations separately. I chose to store all the input and output values in memory for two reasons:
1. The number of values was going to be less than a hundred
2. It made it easier to create unit tests

## Python
This is my main submission, and can be built and executed using Docker. I included a few unit tests, which are executed during the image build process.

## C++
This is separate implementation in the `cpp` folder, and can be compiled with the `compile` script which calls `make`.

## Golang
This is an extra implementation in the `golang` folder, and can be built and executed using Docker. I included a few unit tests, which are executed during the build process.
