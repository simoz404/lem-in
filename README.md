# Lem-in

A digital ant farm simulation written in **Go**. This project reads a file describing a colony of ants, rooms, and tunnels, and calculates the most efficient way to move `n` ants from a start room to an end room while avoiding traffic jams.

![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go)
![Algorithm](https://img.shields.io/badge/Algorithm-Graph%20Theory-purple)
![License](https://img.shields.io/badge/License-MIT-green)

## 📖 Table of Contents
- [About the Project](#about-the-project)
- [How it Works](#how-it-works)
- [Input Format](#input-format)
- [Output Format](#output-format)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Error Handling](#error-handling)

## About the Project
Lem-in is an algorithmic project that focuses on **Graph Theory** and **Pathfinding**. The goal is to find the optimal flow of ants through a network of rooms and tunnels.

The challenge lies in the constraints:
1.  **One Ant per Room:** Only one ant can occupy a room at a time (except for `##start` and `##end`).
2.  **Traffic Management:** You must find not necessarily the *shortest* path, but the *quickest* combination of paths to move all ants without bottlenecks.

## How it Works
The program parses a map file to construct a graph. It then uses pathfinding algorithms (typically BFS or concepts similar to Edmonds-Karp/Max-Flow) to determine the best set of non-conflicting paths. Finally, it simulates the movement turn-by-turn.

## Input Format
The program takes a file passed as an argument. The file structure is as follows:

1.  **Number of Ants:** Single integer.
2.  **Rooms:** `Name X Y` (e.g., `RoomA 2 3`).
    * Special commands `##start` and `##end` signal the entry and exit rooms.
3.  **Links:** `Name1-Name2` (Defines a tunnel between two rooms).
4.  **Comments:** Lines starting with `#` are ignored.

**Example Input:**
```text
3
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
