package utils

import "math/rand/v2"

var default_tiles = []string{"#", "$", "@", "%", "&"}

func GenerateRandomTile(tiles ...string) string {
	if len(tiles) == 0 {
		tiles = default_tiles
	}
	return tiles[rand.IntN(len(tiles))]
}

func GenerateRandomString(length int, tiles ...string) string {
	var res string
	for range length {
		res += GenerateRandomTile(tiles...)
	}
	return res
}
