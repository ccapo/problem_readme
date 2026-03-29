#include <iostream>
#include <iomanip>
#include <sstream>
#include <string>
#include <vector>

void compute(double threshold, double limit, std::vector<double> values) {
  double vsum = 0.0;

  // Set the output format to fixed with a single decimal place
  std::cout << std::fixed << std::setprecision(1);
  
  // Iterate over input values
  for(auto value: values) {
      // Transform input value and update sum
      double delta = std::min(limit - vsum, std::max(0.0, value - threshold));
      vsum += delta;
      
      // Print transformed input value
      std::cout << delta << std::endl;
  }

  // Print final sum
  std::cout << vsum << std::endl;
}

int main(int argc, char *argv[]) {
  double threshold = 0.0, limit = 0.0, value;
  std::vector<double> values;
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

  // Parse the values from STDIN
  while(std::getline(std::cin, line)) {
      std::stringstream ss(line);
      if(ss >> value) {
        values.push_back(value);
      }
  }
  
  // Compute sum
  compute(threshold, limit, values);

  return 0;
}
