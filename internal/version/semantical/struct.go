package semantical

import "regexp"

const (
	rgx        string = `v\.?([\d]+).([\d]+).([\d]+)(-([A-Za-z-]).([\d]+))?`
	major      int8   = 1
	minor      int8   = 2
	patch      int8   = 3
	tag        int8   = 5
	tagVersion int8   = 6
)

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
	Rgx:        regexp.MustCompile(rgx),
	Major:      major,
	Minor:      minor,
	Patch:      patch,
	Tag:        tag,
	TagVersion: tagVersion,
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
