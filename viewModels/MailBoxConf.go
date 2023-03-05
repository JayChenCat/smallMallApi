package viewModels

//邮件配置
type MailBoxConf struct {
	Title          string   //邮件标题
	Content        string   //邮件内容
	RecipientsList []string //收件人列表
	Sender         string   //发件人账号
	SenderPassword string   // 发件人密码，QQ邮箱这里配置授权码
	SMTPAddr       string   // SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPPort       int      // SMTP端口 QQ邮箱是25
}
