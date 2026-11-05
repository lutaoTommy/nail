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
	runes      []rune
	lowerRunes []rune
	entry      ColorEntry
}

type slotPattern struct {
	pattern []rune
	slot    int
}

/*
 * Parse 在文本中做最长匹配，找出所有出现的颜色名（中文名或英文名），
 * 按在文本中首次出现的顺序返回位号和颜色 id。
 * 同一颜色同一槽位在文本中出现多次只记录一次（按首次出现）。
 */
func Parse(text string, entries []ColorEntry) []ColorMatch {
	if text == "" || len(entries) == 0 {
		return nil
	}
	rawText := strings.TrimSpace(text)
	textRunes := []rune(rawText)
	if len(textRunes) == 0 {
		return nil
	}
	textLower := strings.ToLower(rawText)
	textLowerRunes := []rune(textLower)

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
				runes:      []rune(kw),
				lowerRunes: []rune(strings.ToLower(kw)),
				entry:      e,
			})
		}
	}
	// 按关键词长度降序，优先匹配更长的词（如先 "勃艮第红" 再 "红色"）
	sort.Slice(candidates, func(i, j int) bool {
		return len(candidates[i].runes) > len(candidates[j].runes)
	})

	// 已加入结果的 (id, slot) 对，同一对只记录一次
	added := make(map[string]map[int]struct{})
	var result []ColorMatch
	n := len(textRunes)
	var prevColorEnd int // 上一个颜色匹配结束位置，用于限定槽位搜索范围
	for i := 0; i < n; {
		found := false
		for _, c := range candidates {
			kw := c.lowerRunes
			end := i + len(kw)
			if end > n {
				continue
			}
			if runeSliceEqual(textLowerRunes[i:end], kw) {
				slots, afterEnd := findSlotsInContext(textLowerRunes, i, end, prevColorEnd, candidates)
				if len(slots) == 0 {
					slots = []int{0}
				}
				if _, ok := added[c.entry.Id]; !ok {
					added[c.entry.Id] = make(map[int]struct{})
				}
				for _, slot := range slots {
					if _, ok := added[c.entry.Id][slot]; !ok {
						added[c.entry.Id][slot] = struct{}{}
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
				}
				prevColorEnd = afterEnd
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

// 换色位置 1-5 的常见表述：
//
//	数字式：X号位、X号、第X位、位置X、X位（X = 1-5 阿拉伯数字）
//	中式：一~五号、一~五号指甲、第一~五指、第X指甲
//	指名称：拇指/大拇指、食指、中指、无名指、小指/小拇指
var slotPatterns = sortSlotPatterns(buildSlotPatterns())

func sortSlotPatterns(patterns []slotPattern) []slotPattern {
	sort.Slice(patterns, func(i, j int) bool {
		return len(patterns[i].pattern) > len(patterns[j].pattern)
	})
	return patterns
}

func buildSlotPatterns() []slotPattern {
	var out []slotPattern
	// 阿拉伯数字模式
	for slot := 1; slot <= 5; slot++ {
		digit := rune('0' + slot)
		out = append(out, slotPattern{[]rune{digit, '号', '位'}, slot})
		out = append(out, slotPattern{[]rune{digit, '号'}, slot})
		out = append(out, slotPattern{[]rune{'第', digit, '位'}, slot})
		out = append(out, slotPattern{[]rune{'位', '置', digit}, slot})
		out = append(out, slotPattern{[]rune{digit, '位'}, slot})
	}
	// 中文数字模式：一号、二号 ... 五号
	chineseDigits := []rune{'一', '二', '三', '四', '五'}
	for slot := 1; slot <= 5; slot++ {
		cd := chineseDigits[slot-1]
		out = append(out, slotPattern{[]rune{cd, '号'}, slot})
		out = append(out, slotPattern{[]rune{cd, '号', '指', '甲'}, slot})
		out = append(out, slotPattern{[]rune{'第', cd, '指', '甲'}, slot})
		out = append(out, slotPattern{[]rune{'第', cd, '位'}, slot})
		out = append(out, slotPattern{[]rune{'第', cd, '个'}, slot})
	}
	// 手指名称模式
	out = append(out, slotPattern{[]rune{'拇', '指'}, 1})
	out = append(out, slotPattern{[]rune{'大', '拇', '指'}, 1})
	out = append(out, slotPattern{[]rune{'食', '指'}, 2})
	out = append(out, slotPattern{[]rune{'中', '指'}, 3})
	out = append(out, slotPattern{[]rune{'无', '名', '指'}, 4})
	out = append(out, slotPattern{[]rune{'小', '指'}, 5})
	out = append(out, slotPattern{[]rune{'小', '拇', '指'}, 5})
	// 英文数字模式
	engNums := []string{"first", "second", "third", "fourth", "fifth"}
	for slot := 1; slot <= 5; slot++ {
		out = append(out, slotPattern{[]rune("slot " + engNums[slot-1]), slot})
		out = append(out, slotPattern{[]rune("nail " + engNums[slot-1]), slot})
		out = append(out, slotPattern{[]rune("finger " + engNums[slot-1]), slot})
		out = append(out, slotPattern{[]rune("position " + engNums[slot-1]), slot})
	}
	for slot := 1; slot <= 5; slot++ {
		digit := rune('0' + slot)
		out = append(out, slotPattern{[]rune{'s', 'l', 'o', 't', ' ', digit}, slot})
		out = append(out, slotPattern{[]rune{'n', 'a', 'i', 'l', ' ', digit}, slot})
		out = append(out, slotPattern{[]rune{'f', 'i', 'n', 'g', 'e', 'r', ' ', digit}, slot})
		out = append(out, slotPattern{[]rune{'p', 'o', 's', 'i', 't', 'i', 'o', 'n', ' ', digit}, slot})
	}
	// 英文裸序数词
	for slot := 1; slot <= 5; slot++ {
		out = append(out, slotPattern{[]rune("the " + engNums[slot-1]), slot})
		out = append(out, slotPattern{[]rune(engNums[slot-1]), slot})
	}
	// 英文序数缩写 1st, 2nd, 3rd, 4th, 5th
	engOrdAbbr := []string{"1st", "2nd", "3rd", "4th", "5th"}
	for slot := 1; slot <= 5; slot++ {
		out = append(out, slotPattern{[]rune("the " + engOrdAbbr[slot-1]), slot})
		out = append(out, slotPattern{[]rune(engOrdAbbr[slot-1]), slot})
	}
	// 英文手指名称
	out = append(out, slotPattern{[]rune("thumb"), 1})
	out = append(out, slotPattern{[]rune("index"), 2})
	out = append(out, slotPattern{[]rune("middle"), 3})
	out = append(out, slotPattern{[]rune("ring"), 4})
	out = append(out, slotPattern{[]rune("pinky"), 5})
	out = append(out, slotPattern{[]rune("little"), 5})
	return out
}

// findSlotsInContext 在颜色匹配位置前后查找所有不重叠的换色位置 1-5。
// prevColorEnd 是上一个已匹配颜色的结束位置，用于防止槽位窜到前一个颜色。
// candidates 用于在后向扫描中检测下一个颜色名的起始位置，避免槽位窜到下一个分句。
// 返回值 (slots, afterEnd)：slots 为找到的槽位，afterEnd 为后向扫描实际停止（下一个颜色起始/窗口末尾）的位置。
func findSlotsInContext(textRunes []rune, colorStart, colorEnd, prevColorEnd int, candidates []keywordCandidate) ([]int, int) {
	const window = 25
	n := len(textRunes)
	start := colorStart - window
	if start < 0 {
		start = 0
	}
	if prevColorEnd > start {
		start = prevColorEnd
	}
	windowEnd := colorEnd + window
	if windowEnd > n {
		windowEnd = n
	}

	found := make(map[int]struct{})

	// 在颜色前左到右扫描（非重叠最长匹配）
	i := start
	for i < colorStart {
		matched := false
		for _, sp := range slotPatterns {
			q := i + len(sp.pattern)
			if q <= colorStart && q <= n && runeSliceEqual(textRunes[i:q], sp.pattern) {
				found[sp.slot] = struct{}{}
				i = q
				matched = true
				break
			}
		}
		if !matched {
			i++
		}
	}

	// 颜色前已有槽位 → 表达已完整，不再向后搜
	// 颜色前无槽位才后向搜索，并在遇到下一个颜色名时停止
	afterEnd := colorEnd
	if len(found) == 0 {
		i = colorEnd
		for i < windowEnd {
			// 后向扫描遇到下一个颜色名立即停止，防止窜到下一个分句
			if isNextColorStart(textRunes, i, candidates) {
				afterEnd = i
				break
			}
			matched := false
			for _, sp := range slotPatterns {
				q := i + len(sp.pattern)
				if q <= n && runeSliceEqual(textRunes[i:q], sp.pattern) {
					found[sp.slot] = struct{}{}
					i = q
					matched = true
					break
				}
			}
			if !matched {
				i++
			}
		}
		if afterEnd == colorEnd {
			afterEnd = i
		}
	}

	result := make([]int, 0, len(found))
	for s := range found {
		result = append(result, s)
	}
	sort.Ints(result)
	return result, afterEnd
}

// isNextColorStart 检查 textRunes[i:] 是否以某个颜色名为开头（不区分大小写）
func isNextColorStart(textRunes []rune, i int, candidates []keywordCandidate) bool {
	n := len(textRunes)
	for _, c := range candidates {
		q := i + len(c.lowerRunes)
		if q <= n && runeSliceEqual(textRunes[i:q], c.lowerRunes) {
			return true
		}
	}
	return false
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
