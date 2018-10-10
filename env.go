package strastic

import (
	"os"
	"strings"
)

func GetStrasticEnv(prefix string) map[string]string {
	res := map[string]string{}

	for _, pair := range os.Environ() {
		if strings.Index(pair, prefix) != 0 {
			continue
		}

		v := strings.Split(pair, "=")

		res[strings.Replace(v[0], prefix, "", 1)] = v[1]
	}

	return res
}
