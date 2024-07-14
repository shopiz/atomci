package version

import (
	"strconv"
	"strings"
)

// Next 下一个版本号生成器
func Next(version string) string {
	versions := strings.Split(version, ".")
	if len(versions) < 2 {
		versions[1] = "0"
	}
	last := len(versions) - 1
	v, _ := strconv.Atoi(versions[last])
	versions[last] = strconv.Itoa(v + 1)
	return strings.Join(versions, ".")
}
