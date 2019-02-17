package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type EmailData struct {
	ReceiverAddress   string
	NickName          string
	Subject			  string
	Body			  string
}


const (
	Select_all_user = "查找全部用户"
)

type Claims struct {
	Appid string `json:"appid"`
	// recommended having
	jwt.StandardClaims
}

func Token_auth(signedToken, secret string) (string, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		//fmt.Println(reflect.TypeOf(claims.StandardClaims.ExpiresAt))
		//return claims.Appid, err
		return claims.Appid, err
	}
	return "", err
}

//配置session，缓存用
var GlobalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 300,
		ProviderConfig: "./tmp",
	}
	GlobalSessions, _ = session.NewManager("memory",sessionConfig)
	go GlobalSessions.GC()
}

func SendEMail(email EmailData) interface{} {
	// 邮箱地址
	UserEmail := beego.AppConfig.String("user_email")
	// 端口号，:25也行
	Mail_Smtp_Port := beego.AppConfig.String("mail_smtp_port")
	//邮箱的授权码，去邮箱自己获取
	Mail_Password := beego.AppConfig.String("mail_password")
	// 此处填写SMTP服务器
	Mail_Smtp_Host := beego.AppConfig.String("mail_smtp_host")

	fmt.Println(UserEmail)
	fmt.Println(Mail_Smtp_Port)
	fmt.Println(Mail_Password)
	fmt.Println(Mail_Smtp_Host)
	auth := smtp.PlainAuth("", UserEmail, Mail_Password, Mail_Smtp_Host)

	fmt.Println(1)

	to := []string{email.ReceiverAddress} 	// "quartzant@163.com"
	nickname := email.NickName
	user := UserEmail

	fmt.Println(2)

	subject := email.Subject
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := email.Body



	fmt.Println(3)
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(Mail_Smtp_Host+Mail_Smtp_Port, auth, user, to, msg)
	fmt.Println(4)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
	fmt.Println(5)
	fmt.Println(err)
	return err
}


//产生随机字符串
func GetRandomNum(l int) string {
	str := "0123456789" //abcdefghijklmnopqrstuvwxyz
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//写入缓存
func WriteToLog (message string, err error) {
	logs.SetLogger(logs.AdapterFile, `{"filename":"log////test.log"}`)  // ！！！！
	//an official log.Logger
	l := logs.GetLogger()

	l.Println("this is a message in err of" + message + ". the error is", err)
}


//十进制换成其它
var TenToAny = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4"}
func DecimalToAny(num, n int) string {
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		if 76 > remainder && remainder > 9 {
			remainder_string = TenToAny[remainder]
		} else {
			remainder_string = strconv.Itoa(remainder)
		}
		new_num_str = remainder_string + new_num_str
		num = num / n
	}
	return new_num_str
}