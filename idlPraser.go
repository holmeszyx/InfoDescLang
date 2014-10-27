package idl

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strings"
)

// 空属性
var ErrEmptyAttr = errors.New("Attribute is empty")

// 不是属性，不是以"-" 开头
var ErrNotAttr = errors.New("Not attribute")

// 不是信息头, 不是以 ":"结尾
var ErrNotInfo = errors.New("Not information header")

// 信息头
var ErrEmptyInfo = errors.New("Information header is empty")

// 空行
var ErrEmptyLine = errors.New("Empty line")

var ErrNoClosedRawString = errors.New("Raw string not end with \"")

// 信息解析器
type IdlPraser interface {

	// 解析一个信息
	ParseInfo() (info *Information, e error)

	// 解析一个属性
	ParserAttribute() (attr *Attribute, e error)
}

type SimpleIdlParser struct {
	reader *bufio.Reader
	// 解析了的行
	parsedLine int
}

func NewSimpleIdlParser(r io.Reader) *SimpleIdlParser {
	parser := &SimpleIdlParser{}
	parser.reader = bufio.NewReader(r)
	return parser
}

func (s *SimpleIdlParser) ParseInfomations() []*Information {
	infos := make([]*Information, 0, 16)
	var err error = nil
	for err != io.EOF {
		info, e := s.ParseInfo()
		if info != nil {
			infos = append(infos, info)
		}
		err = e
	}
	return infos
}

// 一行已解析
func (s *SimpleIdlParser) lineParsed() int {
	s.parsedLine++
	return s.parsedLine
}

func (s *SimpleIdlParser) ParseInfo() (info *Information, e error) {
	line, err := readLine(s.reader)
	if err != nil {
		return nil, err
	}
	//log.Println("info,", line)
	s.lineParsed()

	firstNoSpace := getFirstNoSpaceIndex(line)
	if firstNoSpace > -1 {
		if line[firstNoSpace] == byte(INFO_SUFFIX) {
			log.Printf("information can not start with ':' or empty information: line %d: %s", s.parsedLine, line)
			return nil, ErrEmptyInfo
		}

		line = strings.TrimSpace(line[firstNoSpace:])
		suffixIndex := strings.LastIndex(line, string(INFO_SUFFIX))
		if suffixIndex == -1 {
			log.Printf("information not end with ':': line %d: %s", s.parsedLine, line)
			return nil, ErrNotInfo
		}

		line = strings.TrimSpace(line[:suffixIndex])
		line, err = getRawString(line)
		if err != nil {
			return nil, err
		}

		info = &Information{line, make([]*Attribute, 0, 8)}

		for {
			attr, err := s.ParserAttribute()
			if attr != nil {
				info.Attrs.Add(attr)
			}

			if err == ErrEmptyLine || err == io.EOF {
				if err == io.EOF {
					e = io.EOF
				}
				break
			}
		}

		return info, e
	} else {
		return nil, ErrEmptyLine
	}

	return
}

func (s *SimpleIdlParser) ParserAttribute() (attr *Attribute, e error) {
	line, err := readLine(s.reader)
	if err != nil {
		return nil, err
	}
	//log.Println("attr,", line)
	s.lineParsed()

	firstNoSpace := getFirstNoSpaceIndex(line)
	if firstNoSpace > -1 {
		if line[firstNoSpace] != byte(ATTR_PREFIX) {
			log.Printf("Attribute not start with '-': line %d: %s", s.parsedLine, line)
			return nil, ErrNotAttr
		}

		line = line[firstNoSpace+1:]
		line = strings.TrimSpace(line)
		line, err = getRawString(line)
		if err != nil {
			return nil, err
		}
		if len(line) == 0 {
			e = ErrEmptyAttr
			return nil, e
		}

		attr = &Attribute{"", line}
		return
	} else {
		return nil, ErrEmptyLine
	}

	return
}

// 读取一行
func readLine(read *bufio.Reader) (string, error) {
	buff := make([]byte, 0, 4096)
	line, prefix, e := read.ReadLine()
	if e != nil {
		return "", e
	}
	buff = append(buff, line...)

	for ; prefix; line, prefix, e = read.ReadLine() {
		if e != nil {
			return string(buff[:]), e
		}
	}

	return string(buff[:]), nil
}

// 获取原始的string
// 如果是 " str " (由引号括着)则为取得 ' str ' (没有引号),
// 如果是  str  则为得到 'str' (没有引号)
func getRawString(spaceTrimedStr string) (string, error) {
	i := strings.Index(spaceTrimedStr, "\"")
	if i == 0 {
		if spaceTrimedStr[len(spaceTrimedStr)-1] == byte('"') {
			return strings.Trim(spaceTrimedStr, "\""), nil
		}
		return "", ErrNoClosedRawString
	} else {
		return spaceTrimedStr, nil
	}
}
