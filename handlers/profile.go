package handlers

import (
	"gengycSrc/lianjiaSpider/engine"
	"gengycSrc/lianjiaSpider/model"
	"gengycSrc/lianjiaSpider/utils"
	"github.com/gocolly/colly"
	"log"
	"regexp"
)

//提取id

var idUrlRe = regexp.MustCompile(`https://changzhou.lianjia.com/ershoufang/pg([\d]+)/`)
var profile model.Profile

func ParseProfile(url string) engine.ParseResults  {
	result := engine.ParseResults{}

	c := NewCollyCollector()
	c.OnHTML("div[class='overview'] > div[class='content']", func(e *colly.HTMLElement) {
		price := utils.HandleStrings(e.ChildText("div[class='price '] > span[class='total']"))
		room := utils.HandleStrings(e.ChildText("div[class='houseInfo'] > div[class='room']"))
		typeHouse := utils.HandleStrings(e.ChildText("div[class='houseInfo'] > div[class='type']"))
		area := utils.HandleStrings(e.ChildText("div[class='houseInfo'] > div[class='area']"))
		communityName := utils.HandleStrings(e.ChildText("div[class='aroundInfo'] > div[class='communityName'] > a[class='info ']"))
		areaName := utils.HandleStrings(e.ChildText("div[class='aroundInfo'] > div[class='areaName'] > span[class='info']"))
		visitTime := utils.HandleStrings(e.ChildText("div[class='aroundInfo'] > div[class='visitTime'] > span[class='info']"))
		profile.Price = price + "万"
		profile.HouseInfo = model.HouseInfo{Room: room, Type: typeHouse, Area: area}
		profile.AroundInfo = model.AroundInfo{CommunityName: communityName, AreaName: areaName, VisitTime: visitTime}
	})
	c.OnHTML("div[class='sellDetailHeader']", func(e *colly.HTMLElement) {
		title := utils.HandleStrings(e.ChildText("div[class='content'] > div[class='title'] > h1"))
		profile.Title =title
		result.Items = []engine.Item{
			{
				Url: url,
				Type: "lianjia",
				Payload:profile,
				Id:extractString([]byte(url),idUrlRe),
			},
		}
	})
	//fmt.Println(url,profile)
	err := c.Visit(url)
	if err!=nil{
		log.Printf("There is an error happened when we visit this detail url:%s\n,and the err is:%v\n",url,err)
	}
	return result
}

type ProfileParser struct {

}

func (p *ProfileParser) Parse(url string) engine.ParseResults{
	return ParseProfile(url)
}

func NewProfileParser() *ProfileParser  {
	return &ProfileParser{}
}
func extractString(contents []byte, re *regexp.Regexp) string  {
	match := re.FindSubmatch(contents)
	if len(match) >=2 {
		return string(match[1])
	} else {
		return ""
	}

}