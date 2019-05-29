package tag

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/roger-russel/got-version/internal/version/semantical"
)

//GetLastTag the last tag
func GetLastTag(pipe *bufio.Reader) string {

	lastTag := semantical.MakeTagMap()

	for {

		input, err := pipe.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		curTag := semantical.Regex.Rgx.FindAllSubmatch([]byte(input), 1)

		check(lastTag, curTag[0])

	}

	return fmt.Sprintf("v.%v.%v.%v",
		lastTag[semantical.Regex.Major],
		lastTag[semantical.Regex.Minor],
		lastTag[semantical.Regex.Patch],
	)

}

func check(last semantical.TagMap, cur [][]byte) {

	for _, v := range semantical.Regex.Rows {

		iLast, _ := last[v]
		iCur, _ := strconv.Atoi(string(cur[v]))

		switch true {
		case iLast < iCur:
			updateLast(last, cur)
			break
		case iLast > iCur:
			break
		}
	}
}

func updateLast(last semantical.TagMap, cur [][]byte) {

	for i := range last {
		last[i], _ = strconv.Atoi(string(cur[i]))
	}

}
