package utils

import "github.com/gobwas/glob"

func StringGlobSlice(s []string, val string) (bool, error) {
	g, err := glob.Compile(val)
	if err != nil {
		return false, err
	}

	for _, v := range s {
		if g.Match(v) {
			return true, nil
		}
	}

	return false, nil
}
