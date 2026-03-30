#!/usr/bin/env python
import argparse
import sys

def main():
    # Create parser for arguments
    parser = argparse.ArgumentParser()
    parser.add_argument("threshold", type=float, help="Input threshold")
    parser.add_argument("limit", type=float, help="Sum limit")
    args = parser.parse_args()
    
    # Parse input from STDIN
    values = []
    for line in sys.stdin:
        value = float(line.strip())
        values.append(value)
    
    # Compute sum
    vsum, deltas = compute(args.threshold, args.limit, values)
    
    # Iterate over transformed inputs
    for delta in deltas:
        # Print transformed value
        print(f"{delta:.1f}")
    
    # Print final sum
    print(f"{vsum:.1f}")

def compute(threshold, limit, values):
    # Sum of input values
    vsum = 0.0
    
    # Output values
    deltas = []

    # Iterate over input values
    for value in values:
        # Transform value and update sum
        delta = min(limit - vsum, max(0.0, value - threshold))
        vsum += delta
        
        # Store transformed value
        deltas.append(delta)

    return vsum, deltas

if __name__ == "__main__":
    main()
