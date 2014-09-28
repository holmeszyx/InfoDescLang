package idl

import (
	"io"
	"bufio"
	"log"
	"strings"
	"errors"
)

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

func NewSimpleIdlParser(r *io.Reader) *SimpleIdlParser{
	parser := &SimpleIdlParser{}
	parser.reader = bufio.NewReader(r)
	return parser
}

// 一行已解析
func (s *SimpleIdlParser) lineParsed() int {
	s.parsedLine ++
	return s.parsedLine
}

func (s *SimpleIdlParser) ParseInfo() (info *Information, e error){

	return
}

func (s *SimpleIdlParser) ParserAttribute() (attr *Attribute, e error) {
	line, err := readLine(s.reader)
	if err != nil {
		return nil, err
	}
	s.lineParsed()

	firstNoSpace := getFirstNoSpaceIndex(line)
	if firstNoSpace > 0 {
		if line[firstNoSpace] != byte(ATTR_PREFIX) {
			log.Fatalf("Attribute not start with '-': line %d: %s", s.parsedLine, line)
		}

		line = line[firstNoSpace+1:]
		line = strings.TrimSpace(line)
		if (len(line) == 0) {
			e = errors.New("Attribute is empty")
			return nil, e
		}
	}else {
		return nil, errors.New("Empty line")
	}

	return
}

// 读取一行
func readLine(read *bufio.Reader)(string, error){
	buff := make([]byte, 0, 4096)
	line, prefix, e := read.ReadLine()
	if e != nil{
		return "", e
	}
	buff = append(buff, line)

	for ; prefix; line, prefix, e = read.ReadLine(){
		if e != nil{
			return string(buff[:]), e
		}
	}

	return string(buff[:]), nil
}
