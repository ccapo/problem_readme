import argparse
import sys

def main():
    # Create parser for arguments
    parser = argparse.ArgumentParser()
    parser.add_argument("threshold", type=float, help="Input threshold")
    parser.add_argument("limit", type=float, help="Sum limit")
    args = parser.parse_args()
    
    # Compute sum
    compute(args.threshold, args.limit)

def compute(threshold, limit):
    # Sum of input values
    vsum = 0.0

    # Parse input from STDIN
    for line in sys.stdin:
        value = float(line.strip())
        delta = min(limit - vsum, max(0.0, value - threshold))
        vsum += delta
        print(f"{delta:.1f}")

    # Print final sum
    print(f"{vsum:.1f}")

if __name__ == "__main__":
    main()
