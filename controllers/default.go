package controllers

import (
	"1td/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

type MainController struct {
	beego.Controller
}

//http://ad.etongdai.com/orange/info?date_from=2016-05-01&date_to=2016-05-30&sign=6a46bbd71bd2fc9cf17fccd82b4e6c72
func (c *MainController) Get() {
	start_time := c.GetString("from")
	end_time := c.GetString("to")
	token := models.Md5(start_time + end_time)
	url := "http://ad.etongdai.com/orange/info?date_from=" + start_time + "&date_to=" + end_time + "&sign=" + token
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			var content []map[string]interface{}
			json.Unmarshal(body, &content)
			c.Data["content"] = content
			c.TplName = "index.tpl"
		}
	}
}
