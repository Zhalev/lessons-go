package hw_2

import (
	"sort"
	"strings"
)

func getTop10Words(text string) (res map[string]int) {

	type Words struct {
		Name  string
		Count int
	}
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, "!", "")
	text = strings.ReplaceAll(text, "?", "")
	text = strings.ReplaceAll(text, ":", "")
	ws := strings.Split(text, " ")
	ms := map[string]int{}
	for _, w := range ws {
		//if _, ok:=s[word]; ok {
		ms[w] = ms[w] + 1
	}
	aa := []Words{}
	for k, c := range ms {
		//if _, ok:=s[word]; ok {
		wor := Words{k, c}
		aa = append(aa, wor)
	}
	sort.SliceStable(aa, func(i, j int) bool {
		return aa[i].Count > aa[j].Count
	})
	aa = aa[:10]
	//fmt.Printf("%v", aa)
	res = make(map[string]int, len(aa))

	for _, v := range aa {
		res[v.Name] = v.Count
	}

	return res

}
