// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

// Package gomaze contains utilities for macOS.
package gomaze

type (
	mazeSquare struct {
		value int
	}
	mazeLine  [10]mazeSquare
	mazeArray [10]mazeLine
)

// var a = {
//     [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
//     [1, 0, 1, 0, 1, 0, 0, 0, 0, 1],
//     [1, 0, 1, 0, 1, 0, 0, 0, 0, 1],
//     [1, 0, 1, 0, 1, 1, 1, 1, 0, 1],
//     [1, 0, 1, 0, 0, 0, 0, 1, 0, 1],
//     [1, 0, 1, 0, 0, 0, 0, 1, 0, 1],
//     [1, 0, 0, 0, 0, 0, 0, 1, 0, 1],
//     [1, 0, 1, 0, 0, 0, 0, 0, 0, 1],
//     [1, 0, 1, 0, 0, 0, 0, 0, 0, 1],
//     [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
// }
