package parser

import (
	"fmt"
	"sort"
	"strings"
)

/*
 * ColorEntry 供解析使用的颜色条目（中文名、英文名、位号等）
 */
type ColorEntry struct {
	Id    string
	X     int
	Y     int
	Name  string
	Color string
	Desc  string
}

/*
 * ColorMatch 解析结果：位号 + 颜色 id + 换色位置(1-5) 及简要信息
 */
type ColorMatch struct {
	PositionNo string
	Id         string
	X          int
	Y          int
	Name       string
	Color      string
	Desc       string
	Slot       int // 换色位置 1-5，0 表示未指定
}

type keywordCandidate struct {
	runes []rune
	entry ColorEntry
}

/*
 * Parse 在文本中做最长匹配，找出所有出现的颜色名（中文名或英文名），
 * 按在文本中首次出现的顺序返回位号和颜色 id。
 * 同一颜色在文本中出现多次只记录一次（按首次出现）。
 */
func Parse(text string, entries []ColorEntry) []ColorMatch {
	if text == "" || len(entries) == 0 {
		return nil
	}
	textRunes := []rune(strings.TrimSpace(text))
	if len(textRunes) == 0 {
		return nil
	}

	// 构建关键词列表：每个颜色的 Name、Color 都作为可匹配关键词（去重、去空）
	seenKw := make(map[string]struct{})
	var candidates []keywordCandidate
	for _, e := range entries {
		for _, kw := range []string{e.Name, e.Color} {
			kw = strings.TrimSpace(kw)
			if kw == "" {
				continue
			}
			key := kw
			if _, ok := seenKw[key]; ok {
				continue
			}
			seenKw[key] = struct{}{}
			candidates = append(candidates, keywordCandidate{
				runes: []rune(kw),
				entry: e,
			})
		}
	}
	// 按关键词长度降序，优先匹配更长的词（如先 "勃艮第红" 再 "红色"）
	sort.Slice(candidates, func(i, j int) bool {
		return len(candidates[i].runes) > len(candidates[j].runes)
	})

	// 已加入结果的 id，同一 id 只保留首次匹配
	added := make(map[string]struct{})
	var result []ColorMatch
	n := len(textRunes)
	for i := 0; i < n; {
		found := false
		for _, c := range candidates {
			kw := c.runes
			end := i + len(kw)
			if end > n {
				continue
			}
			if runeSliceEqual(textRunes[i:end], kw) {
				if _, ok := added[c.entry.Id]; !ok {
					added[c.entry.Id] = struct{}{}
					slot := findSlotInContext(textRunes, i, end)
					result = append(result, ColorMatch{
						PositionNo: positionNo(c.entry.X, c.entry.Y),
						Id:         c.entry.Id,
						X:          c.entry.X,
						Y:          c.entry.Y,
						Name:       c.entry.Name,
						Color:      c.entry.Color,
						Desc:       c.entry.Desc,
						Slot:       slot,
					})
				}
				i = end
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}
	return result
}

func positionNo(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

// 换色位置 1-5 的常见表述：X号位、X号、第X位、位置X、X位
var slotPatterns = buildSlotPatterns()

func buildSlotPatterns() []struct {
	pattern []rune
	slot    int
} {
	var out []struct {
		pattern []rune
		slot    int
	}
	for slot := 1; slot <= 5; slot++ {
		digit := rune('0' + slot)
		// X号位、X号、第X位、位置X、X位
		out = append(out, struct {
			pattern []rune
			slot    int
		}{[]rune{digit, '号', '位'}, slot})
		out = append(out, struct {
			pattern []rune
			slot    int
		}{[]rune{digit, '号'}, slot})
		out = append(out, struct {
			pattern []rune
			slot    int
		}{[]rune{'第', digit, '位'}, slot})
		out = append(out, struct {
			pattern []rune
			slot    int
		}{[]rune{'位', '置', digit}, slot})
		out = append(out, struct {
			pattern []rune
			slot    int
		}{[]rune{digit, '位'}, slot})
	}
	return out
}

// findSlotInContext 在颜色匹配位置前后查找换色位置 1-5，优先取紧挨在颜色前的号位表述
func findSlotInContext(textRunes []rune, colorStart, colorEnd int) int {
	const window = 15
	n := len(textRunes)
	start := colorStart - window
	if start < 0 {
		start = 0
	}
	end := colorEnd + window
	if end > n {
		end = n
	}
	// 优先在颜色前找：从 colorStart-1 向左扫
	bestBefore := 0
	bestBeforeDist := -1
	for p := colorStart - 1; p >= start && p >= 0; p-- {
		for _, sp := range slotPatterns {
			q := p + len(sp.pattern)
			if q <= colorStart && q <= n && runeSliceEqual(textRunes[p:q], sp.pattern) {
				dist := colorStart - q
				if bestBeforeDist < 0 || dist < bestBeforeDist {
					bestBeforeDist = dist
					bestBefore = sp.slot
				}
			}
		}
	}
	if bestBefore > 0 {
		return bestBefore
	}
	// 再在颜色后找：从 colorEnd 向右扫
	for p := colorEnd; p < end; p++ {
		for _, sp := range slotPatterns {
			q := p + len(sp.pattern)
			if q <= n && runeSliceEqual(textRunes[p:q], sp.pattern) {
				return sp.slot
			}
		}
	}
	return 0
}

func runeSliceEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
