package Str

func isIsomorphic(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}
	map1, map2 := map[byte]byte{}, map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if map1[x] > 0 && map1[x] != y || map2[y] > 0 && map2[y] != x {
			return false
		}
		map1[x] = y
		map2[y] = x
	}
	return true
}
