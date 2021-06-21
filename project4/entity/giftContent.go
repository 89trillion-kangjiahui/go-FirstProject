package entity


type GiftContent struct {
	GmtCreate  string            `json:"gmtCreate"`  //创建时间
	CreateUser string            `json:"createUser"` //创建用户
	Describe   string            `json:"describe"`   //礼品描述
	GiftList   []Gift            `json:"giftList"`   //礼品内容列表
	ReceiveNum int            `json:"receiveNum"` //可领取次数
	UsefulDate string            `json:"usefulDate"` //有效期
	AlreadyNum int            `json:"alreadyNum"` //已领取次数
	DrawList   map[string]string `json:"drawList"`   //领取列表
	DrawId     string            `json:"drawId"`     //指定用户的Id
	CodeType   string            `json:"codeType"`   //礼品码类型 1:指定用户一次性消耗，2：不指定用户限制兑换次数，3：不限用户，不限兑换次数
}
