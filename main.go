package main

import (
	"flag"
	"log"
	"mp.52tt.com/meican/api"
	"mp.52tt.com/meican/config"
	"strconv"
	"strings"
)

func main() {

	f := flag.String("f", "", "floor")
	u := flag.String("u", "", "username")
	p := flag.String("p", "", "password")
	dd := flag.String("dd", "", "dislike_dish")
	dr := flag.String("dr", "", "dislike_restaurant")
	pd := flag.String("pd", "", "prefer_dish")
	pr := flag.String("pr", "", "prefer_restaurant")
	flag.Parse()
	//log.Printf("%v", *f)
	if *f != "17" && *f != "18" && *f != "19" ||
		*u == "" || *p == "" {
		log.Println("=========================================")
		log.Println("Usage: ")
		log.Println("\tmeican_robot -u ***@qw.com -p 123123 -f 17")
		log.Println("\tmeican_robot -u ***@qw.com -p 123123 -f 18 -dd \"辣|8勺\"")
		log.Println("\tmeican_robot -u ***@qw.com -p 123123 -f 19 -dr \"抄手|丽华\"")
		log.Println("\tmeican_robot -u ***@qw.com -p 123123 -f 18 -dd \"辣|8勺\"  -dr \"抄手|丽华\" -pd \"猪扒|鸡扒\"")
		log.Println("\tmeican_robot -u ***@qw.com -p 123123 -f 18 -dd \"辣|8勺\"  -dr \"抄手|丽华\" -pd \"猪扒|鸡扒\" -pr \"便当|煲仔饭\"")
		log.Println("Params: ")
		log.Println("\t-u    username ***@qw.com")
		log.Println("\t-p    password 123123")
		log.Println("\t-f    floor 17")
		log.Println("\t-dd   dislike_dish  \"辣|酸\"")
		log.Println("\t-dr   dislike_restaurant  \"抄手|丽华\"")
		log.Println("\t-pd   prefer_dish  \"猪扒|汤饭|寿司\"")
		log.Println("\t-pr   prefer_restaurant  \"便当|煲仔饭\"")
		log.Println("============================================")
		return
	}

	floor, e := strconv.ParseUint(*f, 10, 32)
	if e != nil {
		log.Fatalf("ParseUint failed,e=%v", e)
	}
	conf := config.Config{
		DislikeDishWordList:       strings.Split(*dd, "|"),
		DislikeRestaurantWordList: strings.Split(*dr, "|"),
		PreferDishWordList:        strings.Split(*pd, "|"),
		PreferRestaurantWordList:  strings.Split(*pr, "|"),
		Floor:                     uint32(floor),
	}
	app := api.NewMeiCan(conf)
	finalResult := app.RobotOrder(*u, *p)

	log.Println("===============")
	log.Println("最终点餐结果:")
	log.Println(finalResult)

}
