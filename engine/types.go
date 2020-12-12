package engine

import "gengycSrc/lianjiaSpider/model"

type Parser interface {
	Parse (url string) ParseResults
}

type ParserFunc func(url string) ParseResults

type ProfileParser struct {

}

type FuncParser struct {
	parser ParserFunc
	name  string //函数名
}

type Item struct {
	Id string
	Url string
	Type string
	Payload model.Profile
}

type Request struct {
	Url string
	Parser Parser
}

type ParseResults struct {
	Requests []Request
	Items []Item
}

func (f *FuncParser) Parse(url string) ParseResults  {
	return f.parser(url)
}

func NewFuncParser(p ParserFunc,name string) *FuncParser  {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}