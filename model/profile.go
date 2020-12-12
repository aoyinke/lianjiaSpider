package model

type Profile struct {
	Price string
	Title string
	HouseInfo
	AroundInfo
}

type HouseInfo struct {
	Room string
	Type string
	Area string
}

type AroundInfo struct {
	CommunityName string
	AreaName string
	VisitTime string
}