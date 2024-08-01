# rubik
42 School Project


## Description

This project is a Rubik's Cube solver written in Go


## Installation


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
