package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string	// 封装为内部使用
	Date
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 10 {	// 用来确保没有过多的字符
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
