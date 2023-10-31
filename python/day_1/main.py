#!/usr/bin/env python3

"""
Main file for the project. This file process input of a csv file with headers name,age and occupation and returns  a list of users between ages of 25 and 35 and return as newline list in terminal.
"""

import sys
import os
from typing import List

def process_file(filename: str) -> List[str]:
    #takes a filename with input in lines, break the input in lines based on a blank line ("\n\n") into string array and return the array
    with open(filename) as f:
        lines = f.read().split("\n\n")
    
    return lines

def findTotalCalories(unit: str) -> int:
    # take unit and split its content into lines variable, splitting unit by "\n" and trimming any spaces
    lines = unit.strip().split("\n")
    total = 0
    for line in lines:
        try:
            total += int(line)
        except ValueError:
            print(f"Error: {line} is not a number")
            sys.exit(1)
    return total


if __name__ == "__main__":
    # take input from terminal and pass it to the program
    filename = sys.argv[1]
    if not os.path.exists(filename):
        print(f"Error: File '{filename}' not found")
        sys.exit(1)

    topElves = int(sys.argv[2])

    units = process_file(filename)

    totalCalories = []
    for unit in units:
        totalCalories.append(findTotalCalories(unit))

    totalCalories = sorted(totalCalories)
    topTotal = totalCalories[-topElves:]

    finalTotal = sum(topTotal)
        
    # we print the top elf that has the highest number of calories
    print(f"The top {topElves} elves have {finalTotal} calories")