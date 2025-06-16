# Flood Fill Algorithm

This project was developed as an assignment for the Foundations of Algorithm Design and Analysis course at PUC Minas by **Guilherme Leroy**.

## Problem Introduction

The flood fill algorithm is a fundamental computer graphics and image processing technique used to determine and fill connected regions of similar values in a multi-dimensional array (typically representing pixels in an image).

**The problem can be stated as:**

Given a 2D matrix representing an image, a starting point, and a new color/value, fill all connected cells that have the same original value as the starting point with the new value.

## Applications

The flood fill algorithm has numerous practical applications:

- **Paint Bucket Tool**: Fill enclosed areas in image editing software
- **Game Development**: Flood fill for terrain generation and area detection
- **Medical Imaging**: Region segmentation and analysis
- **Geographic Information Systems**: Area calculation and classification
- **Pathfinding**: Identifying connected components in grid-based games
- **Image Processing**: Background removal and object isolation

## How the Flood Fill Algorithm Works

This project implements an **iterative queue-based** approach using Breadth-First Search (BFS) to avoid stack overflow issues commonly associated with recursive implementations.

### Algorithm Overview

The algorithm works in the following steps:

1. **Initialization**: Start with a matrix where cells contain either available spaces (0) or blocked areas (1)

2. **Region Discovery**: Scan the entire matrix to find connected regions of available cells

3. **Iterative Filling**: For each available region found:
   - Add the starting coordinate to a queue
   - Mark the starting cell with the new color
   - While the queue is not empty:
     - Dequeue a coordinate from the front
     - Check all 4 adjacent coordinates (up, down, left, right)
     - For each valid adjacent cell with the original color:
       - Fill it with the new color
       - Add it to the queue for further processing

### Key Features

- **Queue-based BFS**: Uses a breadth-first search approach for systematic region filling
- **Boundary Checking**: Ensures coordinates stay within matrix bounds
- **Color Validation**: Only fills cells that match the original target color
- **Connected Component Detection**: Automatically identifies and colors separate regions with different values

### Algorithm Complexity

- **Time Complexity**: O(m × n) where m and n are the matrix dimensions
- **Space Complexity**: O(m × n) in worst case for the queue storage

## Project Setup and Execution

This project is built using Go and demonstrates the flood fill algorithm on a sample 5×5 matrix, but it can also work with larger matrix datasets.

### Prerequisites

- Go 1.16 or higher

### Steps to Run the Project

1. Clone the repository or download the source code
2. Navigate to the project directory
3. Run the application:

```bash
go run GOFloodFill.go
```

### Expected Output

The program will display:

1. The initial matrix with 0s (available cells) and 1s (blocked cells)
2. The final matrix after flood fill, where each connected region of 0s is filled with a unique color (2, 3, 4, etc.)

### Sample Execution

```
Initial Matrix:
0 1 0 0 1
0 1 1 0 1
0 0 0 1 0
1 1 0 1 0
0 0 1 0 0

Final Matrix after Iterative Flood Fill:
2 1 3 3 1
2 1 1 3 1
2 2 2 1 4
1 1 2 1 4
5 5 1 6 6
```

Each number represents a different connected region, making it easy to visualize the separate areas that were identified and filled by the algor
