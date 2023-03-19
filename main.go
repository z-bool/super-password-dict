package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"strings"
)

var WeakNum = []string{
	"1",
	"2",
	"12",
	"11",
	"22",
	"000",
	"0000",
	"000000",
	"123",
	"1234",
	"123456",
	"123123",
	"111",
	"1111",
	"111111",
	"888",
	"8888",
	"888888",
	"520",
	"1314",
	"1999",
	"2000",
	"2016",
	"2018",
	"2019",
	"2020",
	"2021",
	"2022",
	"2023",
}

var WeakString = []string{
	"abc",
	"Abc",
	"ABC",
	"Aa",
	"abcd",
	"Abcd",
	"admin",
	"Admin",
	"pass",
	"Pass",
	"passwd",
	"Passwd",
	"password",
	"Password",
	"admin",
	"Admin",
	"user",
	"test",
	"love",
	"super",
}

var weakPasswordWithoutChar = []string{}
var weakPassword = []string{}

type Answers struct {
	ChineseName    string // 中文名
	HuaMing        string // 花名小写
	Xing           string // 姓全拼(首字母大写)
	TuoFenTailMing string // 名全拼(末尾字小写)
	TuoFenMing     string // 名(驼峰)
	XingShouZiMu   string // 姓首字母(代码获取)
	MingShouZiMu   string //名首字母
	GongSiQuanChen string // 公司小写全称
	GongSiJianChen string // 公司中文简称
	GongSiDomain   string // 公司域名
}

// the questions to ask
var qs = []*survey.Question{
	{
		Name:   "ChineseName",
		Prompt: &survey.Input{Message: "职员姓名中文全称,如:王小明(比较少见)"},
	},
	{
		Name:   "HuaMing",
		Prompt: &survey.Input{Message: "公司花名,如:adai"},
	},
	{
		Name:   "Xing",
		Prompt: &survey.Input{Message: "职员姓(首字母大写),如:Wang"},
	},
	{
		Name:   "TuoFenMing",
		Prompt: &survey.Input{Message: "驼峰命名法的名,如:XiaoMing"},
	},
	{
		Name:   "TuoFenTailMing",
		Prompt: &survey.Input{Message: "名首字母大写,如:Xiaoming"},
	},
	{
		Name:   "MingShouZiMu",
		Prompt: &survey.Input{Message: "名首字母,如:xm"},
	},
	{
		Name:   "GongSiQuanChen",
		Prompt: &survey.Input{Message: "公司名拼音,如qiangshengkeji/qiangshen(强盛科技)"},
	},
	{
		Name:   "GongSiJianChen",
		Prompt: &survey.Input{Message: "公司名简拼,如qs(强盛科技)"},
	},
	{
		Name:   "GongSiDomain",
		Prompt: &survey.Input{Message: "公司主域名,如baidu.com"},
	},
}

var result = []string{
	"Abc123!@#",
	"@bcd1234",
	"abc123!@#",
	"Abc123!@#",
	"#EDC4rfv",
	"abcABC123",
	"ABCabc123",
	"1qaz!@#$",
	"QAZwsx123",
	"Pa$$w0rd",
	"P@ssw0rd",
	"P@$$word",
	"P@$$word123",
	"!QAZ2wsx",
	"!QAZ3edc",
	"2wsx#EDC",
	"1!qaz2@wsx",
	"1q2w3e4r",
	"1234abcd",
	"1234qwer",
	"1qaz!QAZ",
	"1qaz2wsx",
	"1qaz@WSX",
	"1qaz@WSX#EDC",
	"!q2w3e4r",
	"1234QWER",
	"QWER!@#$",
	"P@ssw0rd",
	"1qaz@WSX#EDC",
	"p@ssw0rd",
}

// 写入文件
func WriteWeakPassword(weakPassword string) {
	f, err := os.OpenFile("weakPassword.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data := []byte(weakPassword)
	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
}

// 拼接原始弱口令拼接
func CompareWeakPassword() {
	// 1. 第1步把原数据传进去
	for _, num := range WeakNum {
		weakPasswordWithoutChar = append(weakPasswordWithoutChar, num)
		if len(num) >= 6 {
			result = append(result, num)
		}
	}
	for _, str := range WeakString {
		weakPasswordWithoutChar = append(weakPasswordWithoutChar, str)
		if len(str) >= 6 {
			result = append(result, str)
		}
	}
	// 2. 第2步拼接str+num
	for _, num := range WeakNum {
		for _, str := range WeakString {
			sumStr := str + num
			weakPasswordWithoutChar = append(weakPasswordWithoutChar, sumStr)
			result = append(result, str+"@"+num)
			result = append(result, str+"#"+num)
		}
	}
	// 3. 第3步将拼接完的不带字符的带上字符作为一个总的弱口令
	for _, str := range weakPasswordWithoutChar {
		weakPassword = append(weakPassword, str)
	}
	for _, str := range weakPasswordWithoutChar {
		weakPassword = append(weakPassword, str+"!")
		weakPassword = append(weakPassword, str+".")
		weakPassword = append(weakPassword, str+"#")
	}
	// 4. 第4步将弱口令结果导入到result
	for _, str := range weakPassword {
		if len(str) >= 6 {
			result = append(result, str)
		}
	}
}

// 姓名与域名/公司名/公司简拼的处理
func usernameCompare(username string, domain string, gongSiMing string, gongSiJianPin string) {
	if username != "" {
		if domain != "" {
			if len(domain) >= 6 {
				result = append(result, domain)
			}
			// 1. 先加入用户名与域名的拼接
			result = append(result, domain+"@"+username)
			result = append(result, domain+"#"+username)
			result = append(result, domain+"#"+username+"!")
			result = append(result, domain+"@"+username+"!")
			result = append(result, domain+"@"+username+"#")
			result = append(result, username+"@"+domain)
			result = append(result, username+"#"+domain)
			result = append(result, username+"@"+domain+"!")
			result = append(result, username+"@"+domain+"#")
			result = append(result, username+"#"+domain+"!")
			// 2. 域名去掉后缀.com诸如此类的拼接结果
			domainStr := strings.Split(domain, ".")[0]
			result = append(result, domainStr+"@"+username)
			result = append(result, domainStr+"@"+username+"!")
			result = append(result, domainStr+"@"+username+".")
			result = append(result, domainStr+"@"+username+"#")
			result = append(result, domainStr+"#"+username)
			result = append(result, domainStr+"#"+username+"!")
			result = append(result, domainStr+"#"+username+".")
			result = append(result, username+"#"+domainStr)
			result = append(result, username+"#"+domainStr+"!")
			result = append(result, username+"#"+domainStr+".")
			result = append(result, username+"@"+domainStr)
			result = append(result, username+"@"+domainStr+"!")
			result = append(result, username+"@"+domainStr+".")
			result = append(result, username+"@"+domainStr+"#")
			for _, str := range WeakNum {
				result = append(result, domainStr+"@"+username+str)
				result = append(result, domainStr+"@"+username+str+"!")
				result = append(result, domainStr+"@"+username+str+".")
				result = append(result, domainStr+"@"+username+str+"#")
				result = append(result, domainStr+"#"+username+str)
				result = append(result, domainStr+"#"+username+str+"!")
				result = append(result, domainStr+"#"+username+str+".")
				result = append(result, username+"@"+domainStr+str)
				result = append(result, username+"@"+domainStr+str+"!")
				result = append(result, username+"@"+domainStr+str+".")
				result = append(result, username+"@"+domainStr+str+"#")
				result = append(result, username+"#"+domainStr+str)
				result = append(result, username+"#"+domainStr+str+"!")
				result = append(result, username+"#"+domainStr+str+".")
			}
		}
		// 公司名拼音
		if gongSiMing != "" {
			if len(gongSiMing) >= 6 {
				result = append(result, gongSiMing)
			}
			// 名字和公司名的拼接
			result = append(result, gongSiMing+"@"+username)
			result = append(result, gongSiMing+"#"+username)
			result = append(result, gongSiMing+"#"+username+"!")
			result = append(result, gongSiMing+"@"+username+"!")
			result = append(result, gongSiMing+"@"+username+"#")
			result = append(result, username+"@"+gongSiMing)
			result = append(result, username+"#"+gongSiMing)
			result = append(result, username+"@"+gongSiMing+"!")
			result = append(result, username+"@"+gongSiMing+"#")
			result = append(result, username+"#"+gongSiMing+"!")
			for _, str := range WeakNum {
				result = append(result, gongSiMing+"@"+username+str)
				result = append(result, gongSiMing+"@"+username+str+"!")
				result = append(result, gongSiMing+"@"+username+str+".")
				result = append(result, gongSiMing+"@"+username+str+"#")
				result = append(result, gongSiMing+"#"+username+str)
				result = append(result, gongSiMing+"#"+username+str+"!")
				result = append(result, gongSiMing+"#"+username+str+".")
				result = append(result, username+"@"+gongSiMing+str)
				result = append(result, username+"@"+gongSiMing+str+"!")
				result = append(result, username+"@"+gongSiMing+str+".")
				result = append(result, username+"@"+gongSiMing+str+"#")
				result = append(result, username+"#"+gongSiMing+str)
				result = append(result, username+"#"+gongSiMing+str+"!")
				result = append(result, username+"#"+gongSiMing+str+".")
			}
		}
		// 公司简称
		if gongSiJianPin != "" {
			if len(gongSiJianPin) >= 6 {
				result = append(result, gongSiJianPin)
			}
			// 名字和公司简拼的拼接
			result = append(result, gongSiJianPin+"@"+username)
			result = append(result, gongSiJianPin+"#"+username)
			result = append(result, gongSiJianPin+"#"+username+"!")
			result = append(result, gongSiJianPin+"@"+username+"!")
			result = append(result, gongSiJianPin+"@"+username+"#")
			result = append(result, gongSiJianPin+"@"+gongSiMing)
			result = append(result, gongSiJianPin+"#"+gongSiMing)
			result = append(result, gongSiJianPin+"@"+gongSiMing+"!")
			result = append(result, gongSiJianPin+"@"+gongSiMing+"#")
			result = append(result, gongSiJianPin+"#"+gongSiMing+"!")
			for _, str := range WeakNum {
				result = append(result, gongSiJianPin+"@"+username+str)
				result = append(result, gongSiJianPin+"@"+username+str+"!")
				result = append(result, gongSiJianPin+"@"+username+str+".")
				result = append(result, gongSiJianPin+"@"+username+str+"#")
				result = append(result, gongSiJianPin+"#"+username+str)
				result = append(result, gongSiJianPin+"#"+username+str+"!")
				result = append(result, gongSiJianPin+"#"+username+str+".")
				result = append(result, username+"@"+gongSiJianPin+str)
				result = append(result, username+"@"+gongSiJianPin+str+"!")
				result = append(result, username+"@"+gongSiJianPin+str+".")
				result = append(result, username+"@"+gongSiJianPin+str+"#")
				result = append(result, username+"#"+gongSiJianPin+str)
				result = append(result, username+"#"+gongSiJianPin+str+"!")
				result = append(result, username+"#"+gongSiJianPin+str+".")
			}
		}
		// 名字+数字
		for _, str := range WeakNum {
			result = append(result, username+str)
			result = append(result, username+str+".")
			result = append(result, username+str+"!")
			result = append(result, username+str+"#")
		}
		// 名字+@/#+弱口令
		for _, str := range weakPassword {
			if len(username) >= 6 {
				result = append(result, username)
			}
			result = append(result, username+"@"+str)
			result = append(result, username+"#"+str)
		}
	}

}

// 中文名字弱口令组合
func ChineseCompare(username string) {
	for _, str := range weakPassword {
		result = append(result, username+str)
		result = append(result, username+"@"+str)
		result = append(result, username+"#"+str)
	}
}

//	func main() {
//		CompareWeakPassword()
//		usernameCompare("zhangsan", "baidu.com", "baidu", "bd")
//		for _, str := range result {
//			WriteWeakPassword(str + "\r\n")
//		}
//	}
func init() {
	fmt.Println(`

 ________  ___  ___  ________  _______   ________          ________  ___  ________ _________   
|\   ____\|\  \|\  \|\   __  \|\  ___ \ |\   __  \        |\   ___ \|\  \|\   ____\\___   ___\ 
\ \  \___|\ \  \\\  \ \  \|\  \ \   __/|\ \  \|\  \       \ \  \_|\ \ \  \ \  \___\|___ \  \_| 
 \ \_____  \ \  \\\  \ \   ____\ \  \_|/_\ \   _  _\       \ \  \ \\ \ \  \ \  \       \ \  \  
  \|____|\  \ \  \\\  \ \  \___|\ \  \_|\ \ \  \\  \|       \ \  \_\\ \ \  \ \  \____   \ \  \ 
    ____\_\  \ \_______\ \__\    \ \_______\ \__\\ _\        \ \_______\ \__\ \_______\  \ \__\
   |\_________\|_______|\|__|     \|_______|\|__|\|__|        \|_______|\|__|\|_______|   \|__|
   \|_________|                                                                                
                                                                                               
                                   欢迎使用阿呆超级字典生成器，关注公众号:阿呆攻防了解更多                
`)
}

func main() {
	// 结果写入到结构体
	answer := &Answers{}
	var usernameQuanPinXiaoXie string
	var usernameTuoFeng string
	var usernameTuoFengTailMing string
	var usernameXingQuanChenDaXieMingJianChenXiaoXie string
	var usernameXingQuanChenXiaoXieMingJianChenXiaoXie string
	var usernameXingJianPinDaXieMingQuanPinXiaoXie string
	var usernameXingJianPinXiaoXieMingQuanPinXiaoXie string
	var usernameJianPin string

	// 执行提问
	err := survey.Ask(qs, answer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 中文弱口令
	if answer.ChineseName != "" {
		ChineseCompare(answer.ChineseName)
	}
	// 姓名全拼
	if answer.Xing != "" {
		answer.XingShouZiMu = answer.Xing[0:1] //姓缩写
		if answer.MingShouZiMu != "" {
			// 姓全拼+名首字母
			usernameXingQuanChenDaXieMingJianChenXiaoXie = answer.Xing + answer.MingShouZiMu
			usernameXingQuanChenXiaoXieMingJianChenXiaoXie = strings.ToLower(answer.Xing + answer.MingShouZiMu)
			// 姓名简拼
			usernameJianPin = strings.ToLower(answer.XingShouZiMu) + strings.ToLower(answer.MingShouZiMu)
		}

		if answer.TuoFenMing != "" {
			usernameQuanPinXiaoXie = strings.ToLower(answer.Xing) + strings.ToLower(answer.TuoFenMing)
			usernameTuoFeng = answer.Xing + answer.TuoFenMing
			// 姓简拼名全拼
			usernameXingJianPinDaXieMingQuanPinXiaoXie = strings.ToUpper(answer.XingShouZiMu) + strings.ToLower(answer.TuoFenMing)
			usernameXingJianPinXiaoXieMingQuanPinXiaoXie = strings.ToLower(answer.XingShouZiMu) + strings.ToLower(answer.TuoFenMing)
		}
		if answer.TuoFenTailMing != "" {
			usernameQuanPinXiaoXie = strings.ToLower(answer.Xing) + strings.ToLower(answer.TuoFenTailMing)
			usernameTuoFengTailMing = answer.Xing + answer.TuoFenTailMing
			// 姓简拼名全拼
			usernameXingJianPinDaXieMingQuanPinXiaoXie = strings.ToUpper(answer.XingShouZiMu) + strings.ToLower(answer.TuoFenTailMing)
			usernameXingJianPinXiaoXieMingQuanPinXiaoXie = strings.ToLower(answer.XingShouZiMu) + strings.ToLower(answer.TuoFenTailMing)
		}
	}
	CompareWeakPassword()
	if usernameQuanPinXiaoXie != "" {
		usernameCompare(usernameQuanPinXiaoXie, answer.GongSiDomain, answer.GongSiQuanChen, answer.GongSiJianChen)
	}
	if usernameTuoFeng != "" {
		usernameCompare(usernameTuoFeng, answer.GongSiDomain, answer.GongSiQuanChen, answer.GongSiJianChen)
	}
	if usernameTuoFengTailMing != "" {
		usernameCompare(usernameTuoFengTailMing, answer.GongSiDomain, answer.GongSiQuanChen, answer.GongSiJianChen)
	}
	if usernameXingQuanChenDaXieMingJianChenXiaoXie != "" {
		usernameCompare(usernameXingQuanChenDaXieMingJianChenXiaoXie, answer.GongSiDomain, answer.GongSiQuanChen, answer.GongSiJianChen)
	}
	if usernameXingQuanChenXiaoXieMingJianChenXiaoXie != "" {
		usernameCompare(usernameXingQuanChenXiaoXieMingJianChenXiaoXie, answer.GongSiDomain, answer.GongSiQuanChen, answer.GongSiJianChen)
	}
	if usernameXingJianPinDaXieMingQuanPinXiaoXie != "" {
		usernameCompare(usernameXingJianPinDaXieMingQuanPinXiaoXie, answer.GongSiDomain, answer.GongSiQuanChen, answer.GongSiJianChen)
	}
	if usernameXingJianPinXiaoXieMingQuanPinXiaoXie != "" {
		usernameCompare(usernameXingJianPinXiaoXieMingQuanPinXiaoXie, answer.GongSiDomain, answer.GongSiQuanChen, answer.GongSiJianChen)
	}
	if usernameJianPin != "" {
		usernameCompare(strings.ToLower(usernameJianPin), answer.GongSiDomain, answer.GongSiQuanChen, answer.GongSiJianChen)
		usernameCompare(strings.ToUpper(usernameJianPin), answer.GongSiDomain, answer.GongSiQuanChen, answer.GongSiJianChen)
	}
	if answer.HuaMing != "" {
		usernameCompare(strings.ToLower(answer.HuaMing), answer.GongSiDomain, answer.GongSiQuanChen, answer.GongSiJianChen)
	}
	if answer.GongSiDomain != "" {
		for _, str := range weakPassword {
			result = append(result, answer.GongSiDomain+str)
			result = append(result, answer.GongSiDomain+"@"+str)
			result = append(result, answer.GongSiDomain+"#"+str)
		}
	}
	if answer.GongSiQuanChen != "" {
		for _, str := range weakPassword {
			result = append(result, strings.ToLower(answer.GongSiQuanChen)+str)
			result = append(result, strings.ToLower(answer.GongSiQuanChen)+"@"+str)
			result = append(result, strings.ToLower(answer.GongSiQuanChen)+"#"+str)
		}
	}
	if answer.GongSiJianChen != "" {
		for _, str := range weakPassword {
			result = append(result, strings.ToLower(answer.GongSiJianChen)+str)
			result = append(result, strings.ToLower(answer.GongSiJianChen)+"@"+str)
			result = append(result, strings.ToLower(answer.GongSiJianChen)+"#"+str)
			result = append(result, strings.ToUpper(answer.GongSiJianChen)+str)
			result = append(result, strings.ToUpper(answer.GongSiJianChen)+"@"+str)
			result = append(result, strings.ToUpper(answer.GongSiJianChen)+"#"+str)
		}
	}
	for _, str := range result {
		WriteWeakPassword(str + "\r\n")
	}
	fmt.Println("字典生成完毕!!!")

}
