package persist

import (
	"encoding/json"
	"gengycSrc/lianjiaSpider/conf"
	"gengycSrc/lianjiaSpider/engine"
	"gengycSrc/lianjiaSpider/model"
	"log"
	"net/http"
	"strings"
)

func ItemSaver()(chan engine.Item,error)  {

	out := make(chan engine.Item)
	go func() {
		itemCount :=0
		for  {
			item :=<-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			saveData(item.Payload)
		}
	}()

	return out,nil
}

func saveData(profile model.Profile) {
	bodyJson, _ := json.Marshal(profile)
	requestUrl := conf.BaseUrl + "add"
	r, err := http.Post(
		requestUrl,
		"application/json",
		strings.NewReader(string(bodyJson)),
	)
	if err != nil {
		log.Printf("err happened when trying to save data:%v", err)
	}

	defer r.Body.Close()
}