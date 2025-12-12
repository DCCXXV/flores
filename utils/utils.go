package utils

import "math/rand/v2"

var default_tiles = []string{"#", "$", "@", "%", "&"}

func GenerateRandomTile(random *rand.Rand, tiles ...string) string {
	if len(tiles) == 0 {
		tiles = default_tiles
	}
	return tiles[random.IntN(len(tiles))]
}

func GenerateRandomString(random *rand.Rand, length int, tiles ...string) string {
	var res string
	for range length {
		res += GenerateRandomTile(random, tiles...)
	}
	return res
}
