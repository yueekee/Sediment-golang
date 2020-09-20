package main

import (
	"fmt"
	"strings"
)

/*
给你一个字符串 text ，该字符串由若干被空格包围的单词组成。每个单词由一个或者多个小写英文字母组成，并且两个单词之间至少存在一个空格。题目测试用例保证 text 至少包含一个单词 。

请你重新排列空格，使每对相邻单词之间的空格数目都 相等 ，并尽可能 最大化 该数目。如果不能重新平均分配所有空格，请 将多余的空格放置在字符串末尾 ，这也意味着返回的字符串应当与原 text 字符串的长度相等。

返回 重新排列空格后的字符串 。
*/

func main() {
	text := "  this   is  a sentence "
	//text := "  this "
	//text := " practice   makes   perfect"
	//text := "hello   world"
	//text := "  walks  udp package   into  bar a"
	//text := "a"

	spaces := reorderSpaces(text)
	fmt.Println(text)
	fmt.Println(spaces)
}

// func reorderSpaces1(text string) string {
// 	fmt.Println("len(text):", len(text))
// 	trimSpace := strings.TrimSpace(text)
// 	if trimSpace == "" {
// 		fmt.Println("newText1:", len(text))
// 		return text
// 	}

// 	num := 0
// 	for i := 0; i < len(text); i++ {
// 		if text[i] == ' ' {
// 			num++
// 		}
// 	}
// 	fmt.Println("num:", num)

// 	split := strings.Split(text, " ")
// 	for i := 0; i < len(split); i++ {
// 		if split[i] == "" {
// 			if i == len(split)-1 {
// 				split = split[:len(split)-1]
// 			} else {
// 				split = append(split[:i], split[i+1:]...)
// 			}
// 			i--
// 		}
// 	}

// 	newText := ""
// 	if len(split) == 1 {
// 		newText = split[0]
// 		for i := 0; i < num; i++ {
// 			newText += " "
// 		}
// 		fmt.Println("newText2:", len(newText))
// 		return newText
// 	}

// 	preSpaceNum := num / (len(split) - 1)
// 	leftNum := num % (len(split) - 1)
// 	fmt.Println("preSpaceNum:", preSpaceNum)
// 	fmt.Println("leftNum:", leftNum)

// 	for i, s := range split {
// 		newText += s
// 		if i == len(split) - 1 {
// 			continue
// 		}
// 		for i := 0; i < preSpaceNum; i++ {
// 			newText += " "
// 		}
// 	}
// 	for i := 0; i < leftNum; i++ {
// 		newText += " "
// 	}

// 	fmt.Println("newText3:", len(newText))

// 	return newText

// }

func reorderSpaces(text string) string {
	trimSpace := strings.TrimSpace(text)
	if trimSpace == "" {
		return text
	}

	num := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			num++
		}
	}

	split := strings.Split(text, " ")
	var newSplit []string
	for i := 0; i < len(split); i++ {
		if split[i] != "" {
			newSplit = append(newSplit, split[i])
		}
	}

	if len(split) == 1 {
		return split[0] + strings.Repeat(" ", num)
	}

	preSpace := strings.Repeat(" ", num/(len(split)-1))
	leftSpace := strings.Repeat(" ", num%(len(split)-1))

	return strings.Join(newSplit, preSpace) + leftSpace
}
