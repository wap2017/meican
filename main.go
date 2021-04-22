package main

import (
	"flag"
	"github.com/howeyc/gopass"
	"github.com/wap2017/meican/api"
	"github.com/wap2017/meican/config"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fi := flag.String("fi", "", "floor_index")
	u := flag.String("u", "", "username")
	p := flag.String("p", "", "password")
	dd := flag.String("dd", "", "dislike_dish")
	dr := flag.String("dr", "", "dislike_restaurant")
	pd := flag.String("pd", "", "prefer_dish")
	pr := flag.String("pr", "", "prefer_restaurant")
	loc := flag.String("loc", "", "location")
	flag.Parse()

	//log.Printf("%v", *f)
	if *fi != "0" && *fi != "1" && *fi != "2" ||
		*u == "" || *loc == "" {
		log.Println("=========================================")
		log.Println("Usage: ")
		log.Println("\tmeican_robot -u ***@qw.com -f 0 -loc 星辉")
		log.Println("\tmeican_robot -u ***@qw.com -f 0 -loc 星辉 -p 123123")
		log.Println("\tmeican_robot -u ***@qw.com -f 1 -dd \"辣|8勺\"   -loc 星辉 -p 123123")
		log.Println("\tmeican_robot -u ***@qw.com -f 2 -dr \"抄手|丽华\" -loc 星辉 -p 123123")
		log.Println("\tmeican_robot -u ***@qw.com -f 1 -dd \"辣|8勺\"   -dr \"抄手|丽华\" -pd \"猪扒|鸡扒\" -loc 星辉 -p 123123 ")
		log.Println("\tmeican_robot -u ***@qw.com -f 1 -dd \"辣|8勺\"   -dr \"抄手|丽华\" -pd \"猪扒|鸡扒\" -pr \"便当|煲仔饭\" -loc 星辉 -p 123123 ")
		log.Println("Params: ")
		log.Println("\t-u    username ***@qw.com")
		log.Println("\t-p    password 123123")
		log.Println("\t-fi    floor 楼层下标，（17-19）就是（0-2） （11）就是（0）")
		log.Println("\t-dd   dislike_dish  \"辣|酸\"")
		log.Println("\t-dr   dislike_restaurant  \"抄手|丽华\"")
		log.Println("\t-pd   prefer_dish  \"猪扒|汤饭|寿司\"")
		log.Println("\t-pr   prefer_restaurant  \"便当|煲仔饭\"")
		log.Println("\t-loc location \"星辉|高志\"")
		log.Println("============================================")
		return
	}

	pwd := *p
	if pwd == "" {
		pwdB, err := gopass.GetPasswdPrompt("请输入密码:", true, os.Stdin, os.Stdout)
		if err != nil {
			log.Fatalf("GetPasswdPrompt failed err=%v", err)
		}
		pwd = string(pwdB)
	}
	if pwd == "" {
		log.Println("密码为空,退出")
		return
	}

	floorInd, e := strconv.ParseUint(*fi, 10, 32)
	if e != nil {
		log.Fatalf("ParseUint failed,e=%v", e)
	}
	conf := config.Config{
		DislikeDishWordList:       strings.Split(*dd, "|"),
		DislikeRestaurantWordList: strings.Split(*dr, "|"),
		PreferDishWordList:        strings.Split(*pd, "|"),
		PreferRestaurantWordList:  strings.Split(*pr, "|"),
		//Floor:                     uint32(floor),
		FloorInd: uint32(floorInd),
	}
	app := api.NewMeiCan(conf)
	finalResult := app.RobotOrder(*u, pwd, *loc)

	log.Println("===============")
	log.Println("最终点餐结果:")
	log.Println(finalResult)

}
