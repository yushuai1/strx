package strx

import (
	"strings"
)

func New(x string) *str {
	n := str(x)
	return &n
}

type str string

// bool to string
// return "true" or "false"
func BoolToString(source bool) string {
	if source {
		return "true"
	} else {
		return "false"
	}
}

// 判断是否是空，是空返回true，不是空返回false
func IsEmpty(source string) bool {
	if strings.Trim(source, " ") == "" {
		return true
	}
	return false
}

// 判断是否不是空，是空返回false，不是空返回true
func IsNotEmpty(source string) bool {
	return !IsEmpty(source)
}

// 没有连续的空格
func (s *str) NoConsecutiveSpaces() *str {
	var a, b string
	a = string(*s)
	for {
		b = strings.Replace(a, "  ", " ", -1)

		if a == b {
			s = New(strings.Trim(a, " "))
			return s
		}

		a = b
	}

}

// 字符串清除
func (s *str) Clean(cleanStr ...string) *str {
	return s.ReplaceM("", cleanStr...)
}

// 字符串替换右侧
func (s *str) TrimRight(trimStr string) *str {
	s = New(strings.TrimRight(s.Val(), trimStr))
	return s
}

// 字符串替换左侧
func (s *str) TrimLeft(trimStr string) *str {
	s = New(strings.TrimLeft(s.Val(), trimStr))
	return s
}

// 字符串替换两端
func (s *str) Trim(trimStr string) *str {
	s = New(strings.Trim(s.Val(), trimStr))
	return s
}

// 字符串替换
func (s *str) Replace(oldStr string, newStr string) *str {
	s = New(strings.Replace(s.Val(), oldStr, newStr, -1))
	return s
}

// 字符串替换，将所有的oldStr替换为newStr
func (s *str) ReplaceM(newStr string, oldStr ...string) *str {
	*s = str(ReplaceM(s.Val(), newStr, oldStr...))
	return s
}

// 字符串替换，将所有的oldStr替换为newStr
func ReplaceM(sourceStr string, newStr string, oldStr ...string) (resultStr string) {
	for i := 0; i < len(oldStr); i++ {
		sourceStr = strings.Replace(sourceStr, oldStr[i], newStr, -1)
	}
	return sourceStr
}

// 返回第一次出现 start与第一次出现end之间的字符串
func (s *str) Between(start string, end string) *str {
	resultStr := Between(s.Val(), start, end)
	s = New(resultStr)
	return s
}

// 返回第一次出现 start与第一次出现end之间的字符串
func Between(sourceStr, startStr, endStr string) string {
	startIndex := strings.Index(sourceStr, startStr)
	endIndex := strings.Index(sourceStr, endStr)
	return sourceStr[startIndex+len(startStr) : endIndex]
}

// 字符串添加
func (s *str) Append(appendStr string) *str {
	newStr := s.Val() + appendStr
	s = New(newStr)
	return s
}

// 获得string
func (s *str) Val() string {
	return string(*s)
}

// 获得字符串倒序
func (s *str) Reverse() *str {
	runes := []rune(s.Val())
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	s = New(string(runes))
	return s
}
