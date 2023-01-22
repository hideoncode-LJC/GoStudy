package split

import "strings"

func Split(s, seq string) (ret []string) {
	/*
		index(串，子串)
		返回字串在串中第一个出现的位置
		如果串中不存在字串，就返回 -1
	*/
	idx := strings.Index(s, seq)
	// while(i > -1)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+1:]
		idx = strings.Index(s, seq)
	}
	ret = append(ret, s)
	return ret
}
