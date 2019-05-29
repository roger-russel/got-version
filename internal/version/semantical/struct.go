package semantical

import "regexp"

// Regex have regex of all stuffs
var Regex = struct {
	Rgx        *regexp.Regexp
	Major      int8
	Minor      int8
	Patch      int8
	Tag        int8
	TagVersion int8
	Rows       []int8
}{
	Rgx:        regexp.MustCompile("v\\.?([\\d]+).([\\d]+).([\\d]+)(-([A-Za-z-]).([\\d]+))?"),
	Major:      1,
	Minor:      2,
	Patch:      3,
	Tag:        5,
	TagVersion: 6,
	Rows:       []int8{1, 2, 3},
}

type TagMap map[int8]int

func MakeTagMap() TagMap {

	t := make(TagMap)

	for _, v := range Regex.Rows {
		t[v] = 0
	}

	return t

}
