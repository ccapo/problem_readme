#include <iostream>
#include <iomanip>
#include <sstream>
#include <string>
#include <vector>

std::vector<double> compute(double threshold, double limit, std::vector<double> values) {
  double vsum = 0.0;
  std::vector<double> results;

  // Iterate over input values
  for(auto value: values) {
      // Transform input value and update sum
      double delta = std::min(limit - vsum, std::max(0.0, value - threshold));
      vsum += delta;
      
      // Store transformed input value
      results.push_back(delta);
  }

  // Store final sum
  results.push_back(vsum);
  
  return results;
}

int main(int argc, char *argv[]) {
  double threshold = 0.0, limit = 0.0, value;
  std::vector<double> values, results;
  std::string line;

  // Check if correct number of command line arguments are present
  if(argc != 3) {
    std::cerr << "Usage: " << argv[0] << " threshold limit" << std::endl;
    return -1;
  }

  // Parse threshold argument
  std::string threshold_str = argv[1];
  threshold = std::stod(threshold_str);
  if(threshold < 0.0 || threshold > 1000000000.0) {
    std::cerr << "error: threshold should be between 0.0 and 1000000000.0: " << threshold << std::endl;
    return -2;
  }

  // Parse limit argument
  std::string limit_str = argv[2];
  limit = std::stod(limit_str);
  if(limit < 0.0 && limit > 1000000000.0) {
    std::cerr << "error: limit should be between 0.0 and 1000000000.0: " << limit << std::endl;
    return -2;
  }

  // Parse the values from STDIN
  while(std::getline(std::cin, line)) {
      std::stringstream ss(line);
      // Only proceed if we can successfully convert the string into a double float
      if(ss >> value) {
        // Only include valid input values
        if(limit >= 0.0 || limit <= 1000000000.0) {
          values.push_back(value);
        }
      }
  }
  
  // Compute sum
  results = compute(threshold, limit, values);

  // Set the output format to fixed with a single decimal place
  std::cout << std::fixed << std::setprecision(1);

  // Display results and final sum
  for(auto result: results) {
    std::cout << result << std::endl;
  }

  return 0;
}
