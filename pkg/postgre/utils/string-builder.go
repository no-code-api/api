package utils

import (
	"fmt"
	"strings"
)

type StringBuilder struct {
	sb strings.Builder
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{}
}

func (sb *StringBuilder) Append(text string) *StringBuilder {
	sb.sb.WriteString(text)
	return sb
}

func (sb *StringBuilder) AppendNewLine() *StringBuilder {
	sb.sb.WriteString("\n")
	return sb
}

func (sb *StringBuilder) AppendLine(text string) *StringBuilder {
	sb.sb.WriteString(text)
	sb.sb.WriteString("\n")
	return sb
}

func (sb *StringBuilder) AppendFormat(text string, args ...any) *StringBuilder {
	sb.sb.WriteString(fmt.Sprintf(text, args...))
	return sb
}

func (sb *StringBuilder) String() string {
	return sb.sb.String()
}
