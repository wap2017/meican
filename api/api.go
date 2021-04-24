package api

import (
	"encoding/json"
	"fmt"
	"github.com/wap2017/meican/config"
	"github.com/wap2017/meican/module"
	"github.com/wap2017/meican/utils"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type MeiCan struct {
	baseUrl    string
	baseParams map[string]string
	cookies    []*http.Cookie

	finalResult string
	conf        config.Config
}

func NewMeiCan(conf config.Config) *MeiCan {
	return &MeiCan{
		baseUrl: "https://meican.com/",
		baseParams: map[string]string{
			// 感觉像是一段时间不变
			"client_id":     "Xqr8w0Uk4ciodqfPwjhav5rdxTaYepD",
			"client_secret": "vD11O6xI9bG3kqYRu9OyPAHkRGxLh4E",
		},
		finalResult: "\n",
		conf:        conf,
	}
}

func (mc *MeiCan) Login(username, password string) bool {
	m := map[string]string{"remember": "true"}
	requestUrl := mc.buildUrl("preference/preorder/api/v2.0/oauth/token", m)
	response, err := http.PostForm(requestUrl, url.Values{
		"grant_type":             []string{"password"},
		"meican_credential_type": []string{"password"},
		"username":               []string{username},
		"password":               []string{password},
		"username_type":          []string{"username"},
	})
	if err != nil {
		log.Printf("err=%v", err)
		return false
	}
	if response.StatusCode != 200 {
		log.Printf("login falied,statusCode=%v", response.StatusCode)
		return false
	}
	defer response.Body.Close()

	cookies := response.Cookies() //遍历cookies
	mc.cookies = cookies
	//for _, cookie := range cookies {
	//	log.Print("cookie:", cookie)
	//}
	return true
}

func (mc *MeiCan) ShowOrders(beginDate, endDate string) *module.OrderRsp {
	partUrl := "preorder/api/v2.1/calendaritems/list"
	m := map[string]string{
		"withOrderDetail": "false",
		"beginDate":       beginDate,
		"endDate":         endDate,
	}
	response, err := mc.Get(partUrl, m)
	defer response.Body.Close()

	if err != nil {
		log.Fatalf("err=%v", err)
	}
	if response.StatusCode != 200 {
		log.Fatalf("ShowOrders falied,statusCode=%v", response.StatusCode)
	} else {
		d, e := ioutil.ReadAll(response.Body)
		if e != nil {
			log.Fatalf("ReadAll failed,e=%v", e)
		}
		//log.Printf("resp.body=%v", string(d))
		rsp := &module.OrderRsp{}
		if e := json.Unmarshal(d, rsp); e != nil {
			log.Fatalf("Unmarshal failed,e=%v", e)
		}
		return rsp
	}
	return nil
}

func (mc *MeiCan) ShowRestaurants(tabUniqueId string, targetTime string) *module.RestaurantRsp {
	partUrl := "preorder/api/v2.1/restaurants/list"
	m := map[string]string{
		"tabUniqueId": tabUniqueId,
		"targetTime":  targetTime,
	}
	response, err := mc.Get(partUrl, m)
	defer response.Body.Close()

	if err != nil {
		log.Fatalf("err=%v", err)
	}
	if response.StatusCode != 200 {
		log.Fatalf("ShowRestaurants falied,statusCode=%v", response.StatusCode)
	} else {
		d, e := ioutil.ReadAll(response.Body)
		if e != nil {
			log.Fatalf("ReadAll failed,e=%v", e)
		}
		//log.Printf("resp.body=%v", string(d))
		rsp := &module.RestaurantRsp{}
		if e := json.Unmarshal(d, rsp); e != nil {
			log.Fatalf("Unmarshal failed,e=%v", e)
		}
		return rsp
	}
	return nil
}

func (mc *MeiCan) ShowDishes(tabUniqueId string, targetTime string, restId string) *module.DishRsp {
	partUrl := "preorder/api/v2.1/restaurants/show"
	//log.Printf("targetTime=%v", targetTime)
	m := map[string]string{
		"tabUniqueId":        tabUniqueId,
		"targetTime":         strings.Replace(targetTime, " ", "+", 1),
		"restaurantUniqueId": restId,
	}
	response, err := mc.Get(partUrl, m)
	defer response.Body.Close()

	if err != nil {
		log.Fatalf("err=%v", err)
	}
	if response.StatusCode != 200 {
		log.Fatalf("ShowDishes falied,statusCode=%v", response.StatusCode)
	} else {
		d, e := ioutil.ReadAll(response.Body)
		if e != nil {
			log.Fatalf("ReadAll failed,e=%v", e)
		}
		//log.Printf("resp.body=%v", string(d))
		rsp := &module.DishRsp{}
		if e := json.Unmarshal(d, rsp); e != nil {
			log.Fatalf("Unmarshal failed,e=%v", e)
		}
		return rsp
	}
	return nil
}

func (mc *MeiCan) AddOrder(tabUniqueId string, addressUniqueId string, targetTime string,
	dishId uint64) *module.AddOrderRsp {

	partUrl := "preorder/api/v2.1/orders/add"
	//log.Printf("targetTime=%v", targetTime)

	m := map[string]string{
		"corpAddressRemark":   "",
		"corpAddressUniqueId": addressUniqueId,
		"order":               fmt.Sprintf("[{\"count\":1,\"dishId\":%v}]", dishId),
		"remarks":             fmt.Sprintf("[{\"dishId\":\"%v\",\"remark\":\"\"}]", dishId),
		"tabUniqueId":         tabUniqueId,
		"targetTime":          strings.Replace(targetTime, "+", " ", 1),
		"userAddressUniqueId": addressUniqueId,
	}

	//log.Printf("args=%v", m)

	response, err := mc.Post(partUrl, m)
	defer response.Body.Close()

	if err != nil {
		log.Fatalf("err=%v", err)
	}
	if response.StatusCode != 200 {
		log.Fatalf("ShowDishes falied,statusCode=%v", response.StatusCode)
	} else {
		d, e := ioutil.ReadAll(response.Body)
		if e != nil {
			log.Fatalf("ReadAll failed,e=%v", e)
		}
		//log.Printf("resp.body=%v", string(d))
		rsp := &module.AddOrderRsp{}
		if e := json.Unmarshal(d, rsp); e != nil {
			log.Fatalf("Unmarshal failed,e=%v", e)
		}
		return rsp
	}
	return nil
}

func (mc *MeiCan) OrderOneCalendar(order module.DateListItem, calendar module.CalendarItem, conf config.Config, bloomFilter map[string]bool) (string, bool) {

	tabUniqueId := calendar.UserTab.UniqueId
	addressList := calendar.UserTab.Corp.AddressList
	targetTime := time.Unix(int64(calendar.TargetTime/1000), 0).Format("2006-01-02+15:04")
	restaurantRsp := mc.ShowRestaurants(tabUniqueId, targetTime)

	log.Printf("正在检索餐厅...🍴")
	candidateList := make([]*module.CandidateRestAndDish, 0)

	for _, rest := range restaurantRsp.RestaurantList {
		if filterName, match := utils.InDislike(rest.Name, conf.DislikeRestaurantWordList); match {
			log.Printf("根据关键字【%v】过滤不喜欢的餐厅:【%v】...跳过...🦘", filterName, rest.Name)
			continue
		}

		dishRsp := mc.ShowDishes(tabUniqueId, targetTime, rest.UniqueId)

		log.Printf("正在餐厅【%v】搜索佳肴...🥘", rest.Name)
		for _, dish := range dishRsp.DishList {
			if !dish.IsSection {
				// 跳过加钱的
				if dish.PriceString != "15" {
					continue
				}

				// 跳过今天刚吃过的
				if bloomFilter[dish.Name] {
					continue
				}

				if filterName, match := utils.InDislike(dish.Name, conf.DislikeDishWordList); match {
					log.Printf("根据关键字【%v】过滤不喜欢的菜品:【%v】...跳过...🦘", filterName, dish.Name)
					continue
				}

				candidateList = append(candidateList, &module.CandidateRestAndDish{
					RestaurantName:  rest.Name,
					DishName:        dish.Name,
					TabUniqueId:     tabUniqueId,
					AddressUniqueId: addressList[conf.FloorInd].UniqueId,
					TargetTime:      targetTime,
					DishId:          dish.ID,
					Address:         addressList[conf.FloorInd].Address,
					PriceString:     dish.PriceString,
					Title:           calendar.Title,
				})

			}
		}
	}

	//先选喜欢的菜
	for _, c := range candidateList {
		if dishName, match := utils.InPrefer(c.DishName, conf.PreferDishWordList); match {
			log.Printf("根据喜欢的菜品【%v】下单:%v:%v:%v:%v...🍱",
				dishName, c.Title, c.RestaurantName, c.DishName, c.PriceString)
			addOrderRsp := mc.AddOrder(c.TabUniqueId, c.AddressUniqueId, c.TargetTime, c.DishId)
			if addOrderRsp.Status == "SUCCESSFUL" {
				line := fmt.Sprintf("下单成功🥳 根据喜欢的菜品【%v】下单:%v:%v:%v:%v:%v...🍱",
					dishName, order.Date, c.Title, c.RestaurantName, c.DishName, c.PriceString)
				mc.finalResult += line + "\n"
				log.Printf("下单成功🥳")
				return c.DishName, true
			} else {
				panic("下单失败🌚")
			}
		}
	}

	//再选喜欢的餐厅
	for _, c := range candidateList {
		if restName, match := utils.InPrefer(c.RestaurantName, conf.PreferRestaurantWordList); match {
			log.Printf("根据喜欢的餐厅【%v】下单:%v:%v:%v:%v...🍱",
				restName, c.Title, c.RestaurantName, c.DishName, c.PriceString)
			addOrderRsp := mc.AddOrder(c.TabUniqueId, c.AddressUniqueId, c.TargetTime, c.DishId)
			if addOrderRsp.Status == "SUCCESSFUL" {
				line := fmt.Sprintf("下单成功🥳 根据喜欢的餐厅【%v】下单:%v:%v:%v:%v:%v...🍱",
					restName, order.Date, c.Title, c.RestaurantName, c.DishName, c.PriceString)
				mc.finalResult += line + "\n"
				log.Printf("下单成功🥳")
				return c.DishName, true
			} else {
				panic("下单失败🌚")
			}
		}
	}

	//没有就瞎几把乱选
	if len(candidateList) > 0 {
		rand.Seed(time.Now().Unix())
		c := candidateList[rand.Int31n(int32(len(candidateList)))]
		log.Printf("随机下单:%v:%v:%v:%v...🍱", c.Title, c.RestaurantName, c.DishName, c.PriceString)
		addOrderRsp := mc.AddOrder(c.TabUniqueId, c.AddressUniqueId, c.TargetTime, c.DishId)
		if addOrderRsp.Status == "SUCCESSFUL" {
			line := fmt.Sprintf("下单成功🥳 随机下单:%v:%v:%v:%v:%v...🍱",
				order.Date, c.Title, c.RestaurantName, c.DishName, c.PriceString)
			mc.finalResult += line + "\n"
			log.Printf("下单成功🥳")
			return c.DishName, true
		} else {
			log.Printf("addOrderRsp=%+v", addOrderRsp)
			panic("下单失败🌚")
		}
	}

	return "", false
}

func (mc *MeiCan) RobotOrder(username, password string) string {
	log.Printf("正在登陆...🤓")
	isLogin := mc.Login(username, password)
	if !isLogin {
		return "账号名或者密码错误☹️"
	}

	log.Printf("正在查看订单...🔖")
	orders := mc.ShowOrders(utils.GetStartDateAndEndDate())
	log.Printf("正在检索日期...📆")

	mc.finalResult = "\n"
	bloomFilter := make(map[string]bool)

	for _, order := range orders.DateList {
		t, e := time.Parse("2006-01-02", order.Date)
		if e != nil {
			panic(e)
		}

		//跳过周末
		if t.Weekday() == time.Saturday || t.Weekday() == time.Sunday {
			continue
		}

		log.Printf("正在检索上下午...🌔")

		for _, calendar := range order.CalendarItemList {

			// 跳过已关闭
			if calendar.Status == "CLOSED" {
				continue
			}

			//location: 高志|星辉
			if strings.Contains(calendar.UserTab.Name, mc.conf.Location) {
				//跳过已点餐
				log.Printf("date:%v title:%v", order.Date, calendar.Title)
				if calendar.CorpOrderUser.UniqueId != "" {
					dishName := calendar.CorpOrderUser.RestaurantItemList[0].DishItemList[0].Dish.Name
					line := fmt.Sprintf("当前时间date:%v title:%v 已点餐:【%v】...跳过...🦘", order.Date, calendar.Title, dishName)
					mc.finalResult += line + "\n"
					log.Println(line)
					continue
				}

				if dishName, done := mc.OrderOneCalendar(order, calendar, mc.conf, bloomFilter); done {
					bloomFilter[dishName] = true
				}
			}

		}
	}
	return mc.finalResult
}
