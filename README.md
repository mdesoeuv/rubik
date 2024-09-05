# Rubik
42 School Project


## Description

This project is a Rubik's Cube solver written in Go


## Installation
You can run `go install github.com/mdesoeuv/rubik`

## Usage

### CLI

```bash
go run . [flags] "<Moves>"
```

### Flags

- `--tui` / `-tui` : Launch the program in TUI mode
- `--profile` / `-profile` : Launch the program with pprof enabled (read the profile file with `go tool pprof -http localhost:8080 profile.prof`)
- `--help` / `-help` : Display the help message

### Examples

```bash
go run . "R U R' U'"
```

```bash
go run . -tui "R U R' U'"
```

```bash
go run . -profile "R U R' U'"
```

## Moves Notation

The `FRUBLD` notation is used to describe the moves
Each letter represents a face of the cube that can be rotated clockwise by default

### Base moves

- `F` : Front face clockwise rotation
- `R` : Right face clockwise rotation
- `U` : Up face clockwise rotation
- `B` : Back face clockwise rotation
- `L` : Left face clockwise rotation
- `D` : Down face clockwise rotation

### Inverted moves

The counter-clockwise rotation of a face is denoted by a `'` after the face letter

- `F'` : Front face counter-clockwise rotation
- `R'` : Right face counter-clockwise rotation
- `U'` : Up face counter-clockwise rotation
- `B'` : Back face counter-clockwise rotation
- `L'` : Left face counter-clockwise rotation
- `D'` : Down face counter-clockwise rotation

### Double moves

The double rotation of a face is denoted by a `2` after the face letter
Double moves are equivalent to two consecutive clockwise moves but are accounted for 1 move in the move count

- `F2` : Front face double rotation
- `R2` : Right face double rotation
- `U2` : Up face double rotation
- `B2` : Back face double rotation
- `L2` : Left face double rotation
- `D2` : Down face double rotation

# Thistlethwaite's algorithm

This program uses Thistlethwaite's algorithm to solve the rubik's cube.
As the reference we used [Jaap's website](https://www.jaapsch.net/puzzles/thistle.htm) that has web version
of a letter written on a typewritter by Thistlethwaite himself.

## Principals

The algorithm works by using group theory, the core idea being that
> Limiting the moves used to shuffle a solved cube only leads to cubes solvable by those moves.
> Thus creating a growing sequence move groups can provide a sequence of cube groups of grownig size.
> This way from a scrambled cube one can work their way up to smaller and smaller groups eventually getting to the solved cube.

## The groups used by Thistlethwaite

The letter describes 5 groups of varying sizes:
G0=<L, R, F, B, U, D>	4.33e19
G1=<L, R, F, B, U2,D2>	2.11e16
G2=<L, R, F2,B2,U2,D2>	1.95e10
G3=<L2,R2,F2,B2,U2,D2>	6.63e5
G4=I	                1

### Group G0
G0 is the biggest group containing every possible rubik's cube shuffle, at this stage all moves
are authorized to get to the next: G1.

### Group G1
G1 is the group that does not require quater bottom or top spins to be solved.
It is described as solving the orientation of the edges. In fact preventing quarter rotation
on U & D prevents the inversion of an edge.

### Group G2
G2 is where this time quarter rotations of the front and back side, this corresponds to solving
corner rotations this time. But these moves limitations after G2 will prevent edges between right & left
to get out, this means that the right left slice must also hold the correct edges.

### Group G3
This is the most difficult group, here, edges and corners must be in their correct orbit.
An orbit being the locations corner/edge can get into using only half turns, for edges that means
being in the correct slice, and for corners to be in the correct group of four (on each face they are at the opposite corner).
However even if corners are in the correct orbit not all permutation of them can be solved using
only half-turns.
To find the solvable 96 permutation one can apply all possible shuffles on corner permutations available
by doing only half-turns

### Group G4
Well done you solved the rubiks, in this group no moves are allowed as the only cube is the solved cube
