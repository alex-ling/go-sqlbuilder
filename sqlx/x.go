package sqlx

import (
	"bytes"
	"fmt"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func EscapeName(name string) string {
	if len(name) == 0 {
		panic("Empty name")
	}
	return "`" + name + "`"
}

func EscapeValue(val interface{}) string {
	if str, ok := val.(string); ok {
		if "?" == str {
			return str
		} else {
			return "'" + str + "'"
		}
	} else {
		return fmt.Sprintf("%v", val)
	}
}

func Map0(items []interface{}, mapper func(interface{}) string) []string {
	newItems := make([]string, len(items))
	for i, v := range items {
		newItems[i] = mapper(v)
	}
	return newItems
}

func Map(items []string, mapper func(string) string) []string {
	newItems := make([]string, len(items))
	for i, v := range items {
		newItems[i] = mapper(v)
	}
	return newItems
}

func opEscape(name string, op string, value interface{}) string {
	return EscapeName(name) + op + EscapeValue(value)
}

func opIgnore(name string, op string, value string) string {
	return EscapeName(name) + op + value
}

func makeSQL(buffer *bytes.Buffer) string {
	buffer.WriteByte(';')
	return buffer.String()
}
