#include <iostream>
#include <iomanip>
#include <sstream>
#include <string>

int main(int argc, char *argv[]) {
  double threshold = 0.0, limit = 0.0, vsum = 0.0, value;
  std::string line;

  // Check if correct number of command line arguments are present
  if(argc != 3) {
    std::cerr << "Usage: " << argv[0] << " threshold limit" << std::endl;
    return -1;
  }

  // Parse threshold argument
  std::string threshold_str = argv[1];
  threshold = std::stod(threshold_str);

  // Parse limit argument
  std::string limit_str = argv[2];
  limit = std::stod(limit_str);
  
  // Set the output format to fixed with a single decimal place
  std::cout << std::fixed << std::setprecision(1);

  // Parse the values from STDIN
  while(std::getline(std::cin, line)) {
      std::stringstream ss(line);
      if(ss >> value) {
        // Transform input value and update sum
        double delta = std::min(limit - vsum, std::max(0.0, value - threshold));
        vsum += delta;
        std::cout << delta << std::endl;
      }
  }

  std::cout << vsum << std::endl;

  return 0;
}
