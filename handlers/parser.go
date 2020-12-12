package handlers

import (
	"gengycSrc/lianjiaSpider/engine"
	"github.com/gocolly/colly"
	"log"
)

func Parser(url string) engine.ParseResults{

	c := NewCollyCollector()
	result := engine.ParseResults{}
	c.OnHTML("ul[class='sellListContent']", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, item *colly.HTMLElement) {
			href := item.ChildAttr("div[class='info clear'] > div[class='title'] > a","href")
			//log.Printf("Parser is executed with url:%s,and we get the href:%s",url,href)

			result.Requests = append(result.Requests,engine.Request{
				Url: href,
				Parser: engine.NewFuncParser(ParseProfile,"ParseProfile"),
			})
		})
	})
	err := c.Visit(url)
	if err!=nil{
		log.Printf("There is an error happened when we visit this url:%s\n,and the err is:%v\n",url,err)
	}

	return result
}
