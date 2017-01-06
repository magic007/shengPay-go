package main

import (
	"shengPay" //导入库
	"fmt"	
	"strconv"
	"math/rand"
	
)

/** 盛付通支付
 * Created by Magic <zhongguovu@gmail.com>.
 * User: Magic
 * Date: 2016/12/26
 * Time: 11:35
 */

func main() {

	fmt.Println("程序开始运行.......")

	fmt.Println("/********************************************************************")
	fmt.Println("***本程序测试使用测试商户号:“107537”")
	fmt.Println("***如需修改，请修改代码shengPay.go SENDER_ID 参数改为:“申请商户号”")	
	fmt.Println("********************************************************************/")


	var orderNO string
	
	fmt.Print("请输入一个测试订单号,6个字符以上: ")
	fmt.Scanln(&orderNO)
	if orderNO==""   {
		unixNano:=rand.Intn(10)
		orderNO = strconv.Itoa(unixNano)			
	}
	fmt.Printf("订单号：%s，订单生成中......\n",orderNO)
	
	money := 0.01
	productName := "测试产品名称"
	result, err := shengPay.Pay(orderNO, money, productName)
	if err != nil {
		fmt.Println("错误", err)
	}

	
	fmt.Printf("返回结果:%+v\n",result.Return)

	fmt.Println("(********************************************************************")
	fmt.Println("订单号：",orderNO,",支付地址", result.Return.Extension.Ext3)

	fmt.Println("/********************************************************************")
	fmt.Println("***如需微信支付，请把代码shengPay.go PAY_CHANNEL 参数改为:“wp”")
	fmt.Println("***请把支付地址用手机浏览器打开，或在线生成二维码,用手机支付宝APP扫描")	
	fmt.Println("********************************************************************/")
	for {}
}
