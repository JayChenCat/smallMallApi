package service

import (
	"SmallMall/utils/common"
	"SmallMall/utils/erromsg"
	_ "SmallMall/utils/erromsg"
	"SmallMall/utils/goredis"
	"SmallMall/viewModels"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

//发送电子邮件（支持群发，传入的收件人账号为多个）方法
//@ recipientsList 收件人账号字符数组
func SendEmailCode(uid int64, recipientsList []string) int {
	msg, rcode := SendEmail(recipientsList)
	redis_key := "email-code"
	//有效时间为负数时就说明验证码过期了
	expire := goredis.GetExpire(redis_key)
	fmt.Println(expire)
	ok := goredis.Exists(redis_key)
	//redis不存在emailCode的验证码的值，就向redis中存入有效期为3分钟的验证码
	if !ok {
		if msg == erromsg.ERROR_EMAIL_RegisteredSUCC {
			//设置验证码有效期时间3分钟
			succ := goredis.Set(redis_key, rcode, 3*time.Minute)
			if succ {
				return erromsg.ERROR_EMAIL_RegisteredSUCC
			}
			return erromsg.ERROR_EMAIL_RegisteredFAIL
		}
		return msg
	}
	//提示验证码已存在，请勿重复发送
	return erromsg.ERROR_EMAIL_EXIST
	/*emailSendCode, msgCode := models.ByUserIdGetCode(uid)
	//数据库中没有发送邮箱验证码记录
	if msgCode == erromsg.ERROR_EMAIL_NOEXIST {
		return SendAndInsertCodeInfo(uid, recipientsList)
	}
	//数据库中有发送邮箱验证码记录，判断该验证码是否在有效期内
	addTime := emailSendCode.AddTime
	currTime := time.Now()
	difference := currTime.Sub(addTime)
	expire := difference.Minutes()
	_expireTime := emailSendCode.Expire
	//过期后就清除验证码记录
	if expire > _expireTime {
		DeleteCodeInfo(uid)
		return SendAndInsertCodeInfo(uid, recipientsList)
	}
	return msgCode*/
}

//发送邮箱验证码并入库存储
func SendAndInsertCodeInfo(uid int64, recipientsList []string) int {
	//保存最新验证码记录
	msg, rcode := SendEmail(recipientsList)
	//判断发往收件人邮箱的验证码是否成功,成功后才存入验证码记录
	if msg == erromsg.ERROR_EMAIL_RegisteredSUCC {
		codeInfo := viewModels.CodeInfoViewModel{
			//Id:      100000279878,
			UId:     uid,
			Code:    rcode,
			Expire:  3,
			AddTime: time.Now(),
		}
		code := EditCodeInfo(codeInfo)
		if code == erromsg.SUCCSE {
			return erromsg.ERROR_EMAIL_RegisteredSUCC
		}
		return code
	}
	return msg
}

func SendEmail(recipientsList []string) (int, string) {
	var mailConf viewModels.MailBoxConf
	mailConf.Title = "美物小程序后台管理系统验证码通知"
	//发送的邮箱内容，但是也可以通过下面的html代码作为邮件内容
	// mailConf.Body = "Sweet"
	//支持群发，只需填写多个人的邮箱即可，发送人使用的是QQ邮箱，所以接收人也必须都要是QQ邮箱
	mailConf.RecipientsList = recipientsList //[]string{"邮箱账号1", "邮箱账号2"} 收件人账号字符数组
	mailConf.Sender = "2602719914@qq.com"    //`邮箱账号`
	//QQ邮箱要填写授权码，网易邮箱则直接填写自己的邮箱密码，授权码获得方法在下面
	mailConf.SenderPassword = "coqguizhcskmebib" //"填写自己QQ邮箱授权码"
	//下面是官方邮箱提供的SMTP服务地址和端口
	// QQ邮箱：SMTP服务器地址：smtp.qq.com（端口：587）
	// 雅虎邮箱: SMTP服务器地址：smtp.yahoo.com（端口：587）
	// 163邮箱：SMTP服务器地址：smtp.163.com（端口：25）
	// 126邮箱: SMTP服务器地址：smtp.126.com（端口：25）
	// 新浪邮箱: SMTP服务器地址：smtp.sina.com（端口：25）
	mailConf.SMTPAddr = `smtp.qq.com`
	mailConf.SMTPPort = 25
	//产生六位数随机验证码
	_rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	rcode := fmt.Sprintf("%06v", _rand.Int31n(1000000))
	//发送的内容
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%s,为了保证账号安全，验证码有效期为3分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, rcode)
	//调用发送邮箱第三方插件
	m := gomail.NewMessage()
	// 增加发件人别名
	m.SetHeader("From", m.FormatAddress(mailConf.Sender, "美物小程序后台管理系统管理员"))
	//说明:如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面方法转码
	//m.SetHeader("From", "test"+"<"+mailConf.Sender+">") //增加发件人别名
	m.SetHeader(`To`, mailConf.RecipientsList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)
	// m.Attach("./Dockerfile") //添加附件
	d := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SenderPassword)
	//关闭ssl协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	if err != nil {
		common.WriteLog(fmt.Sprintf("向用户绑定的邮箱发送验证码错误:%s", err.Error()))
		return erromsg.ERROR_EMAIL_RegisteredFAIL, "" //邮箱发送失败
	}
	return erromsg.ERROR_EMAIL_RegisteredSUCC, rcode
}
