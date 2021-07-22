package lc0071

import (
	path2 "path"
	"strings"
)

func simplifyPath(path string) string {
	var stdPath []rune
	stdPath = append(stdPath, '/')

	tokens := strings.Split(path, "/")

	for _, t := range tokens {
		switch t {
		case "":
		case ".":
		case "..":
			// go back
			if len(stdPath) > 1 {
				// $$ 想清楚i的含义
				// 这里，stdPath[i]会被去掉，所以i必须>=1
				// for ; i > 0; i-- { 会导致i降到0所以出错

				// $$ 循环体内的i永远符合条件，
				// 但循环体外用到i的地方会多一次++/-- !!!
				i := len(stdPath) - 1
				//for ; i > 0; i-- {
				for ; i > 1; i-- {
					if stdPath[i] == '/' {
						break
					}
				}
				stdPath = stdPath[:i]
			}
		default:
			// double slash
			if len(stdPath) > 1 {
				stdPath = append(stdPath, '/')
			}
			for _, r := range t {
				stdPath = append(stdPath, r)
			}
		}
	}

	return string(stdPath)
}

func simplifyPath2(path string) string {
	ret := []string{}
	for _, v := range strings.Split(path, "/") {
		switch v {
		case "":
			break
		case ".":
			break
		case "..":
			if l := len(ret); l > 0 {
				ret = ret[:l-1]
			}
		default:
			ret = append(ret, v)
		}
	}
	return "/" + strings.Join(ret, "/")
}

func simplifyPath3(path string) string {
	return path2.Clean(path)
}

// func main() {
// 	p := "/home/"
// 	fmt.Println(simplifyPath(p))
// }
