package main

// 字符串模式匹配
// https://leetcode-cn.com/problems/pattern-matching-lcci/
//
// 参考 https://leetcode-cn.com/problems/pattern-matching-lcci/solution/goshuang-100-by-oldjon/


func patternMatching(pattern string, value string) bool {

	// 首先处理特殊情况
	if len(pattern) == 0 {
		if len(value) == 0 {
			return true
		}else{
			return false
		}
	}

	pattern, anum, bnum, bstart := countandconvertab(pattern)

	// 首先需要检查pattern是否能等分成len(pattern)个单一字符
	if len(pattern) == anum {
		return canslip(value, anum)
	}
	// 然后检查当pattern中只有一种字母时的情况
	if len(value) > 0 && (canslip(value,anum) || canslip(value,bnum)){
		return true
	}

	abmap := make(map[rune]string)

	for i := 1; i <= len(value)/anum; i++ {
		// i 即为a的长度
		// 假定a为value[:i]
		abmap['a'] = value[:i]
		// 取b
		if (len(value)-i*anum)%bnum != 0 {
			continue
		}
		// 通过value的长度减去全部a所占的长度，即全部b的长度，除以b的个数，得到b的长度
		blen := (len(value) - i*anum) / bnum
		// value头部n个a串后的blen个字符即为b
		abmap['b'] = value[bstart*i : bstart*i+blen]
		// 生成abmap中ab组成的字符串
		p := ""
		for _, v := range pattern {
			p = p + abmap[v]
		}
		if value == p {
			return true
		}
	}
	return false
}

func countandconvertab(p string) (op string,anum,bnum,bstart int){
	// 计算a和b的个数 并将原pattern转换为以a开头的字符串
	// 同时返回b的起始位置
	startWithB := false
	if p[0] == 'b'{
		startWithB = true
	}

	ob := make([]rune, len(p))

	for i := range p {
		if startWithB {
			if p[i] == 'b' {
				ob[i] = 'a'
			} else {
				ob[i] = 'b'
			}
		} else {
			ob[i] = rune(p[i])
		}
		if ob[i] == 'a'{
			anum ++
		} else {
			bnum ++
			if bstart == 0 {
				bstart = i
			}
		}
	}
	return string(ob), anum, bnum, bstart
}

func canslip(s string,n int)bool{
	// 判断字符能否被整除
	if len(s) %n !=0{
		return false
	}
	intv:=len(s)/n
	i := intv
	a := s[:intv]
	for i < len(s) {
		if s[i:i+intv] != a {
			return false
		}
		i+=intv
	}
	return true
}