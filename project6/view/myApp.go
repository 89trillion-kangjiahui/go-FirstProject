package view

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/golang/protobuf/proto"

	"project6/global"
	"project6/response"
	"project6/service"
)

//设置界面
func SetView() *fyne.Container {
	//用户名输入框
	nameEntry := widget.NewEntry()
	//用户列表
	ul := widget.NewLabel("user_list:")
	//内容列表
	cl := widget.NewVBox()
	//说话内容输入框
	multiEntry := widget.NewEntry()
	multiEntry.SetPlaceHolder("please enter\nyour description")
	multiEntry.MultiLine = true

	//获得用户名输入行
	Line1 := getLine1(nameEntry)
	//获得连接管理行
	serverNewLabel, statusNewLabel, Line2 := getLine2(nameEntry, ul)
	//获得用户列表，聊天内容
	Line3 := getLine3(serverNewLabel, statusNewLabel, ul, cl)
	//获得输入聊天内容行
	Line4 := getLine4(multiEntry)

	//将所有的组件进行布局
	content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		Line1, Line2, Line3, multiEntry, Line4)
	return content
}

//获得用户名输入行
func getLine1(nameEntry *widget.Entry) *fyne.Container {
	Line1 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		widget.NewLabel("UserName:"),
		nameEntry)
	return Line1
}

//获得连接管理行
func getLine2(nameEntry *widget.Entry, ul *widget.Label) (*widget.Label, *widget.Label, *fyne.Container) {
	serverNewLabel := widget.NewLabel("")
	statusNewLabel := widget.NewLabel("no")
	serverBox := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		widget.NewLabel("Server:"),
		serverNewLabel)
	LoginBtn := widget.NewButton("login", func() {
		service.LoginService(nameEntry)
	})
	LogoutBtn := widget.NewButton("exit", func() {
		//更新用户的连接状态和ip地址,用户列表，以及用户内容
		serverNewLabel.SetText("")
		statusNewLabel.SetText("no")
		ul.SetText("")
		service.ExitService()
	})
	statusBox := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		widget.NewLabel("status:"), statusNewLabel)

	Line2 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		serverBox,
		layout.NewSpacer(),
		LoginBtn,
		LogoutBtn,
		layout.NewSpacer(),
		statusBox)
	return serverNewLabel, statusNewLabel, Line2
}

//获得用户列表，聊天内容
func getLine3(serverNewLabel, statusNewLabel, ul *widget.Label, cl *widget.Box) *fyne.Container {
	//重新渲染聊天内容，登录信息，用户列表, 用户连接状态，以及用户的ip地址。
	go func() {
		for {
			select {
			case c := <-global.Connection.ReadChan:
				var ret response.Data
				proto.Unmarshal(c, &ret)
				if ret.Type == "talk" {
					//输出用户消息
					cl.Append(widget.NewLabel(ret.Content))
				} else if ret.Type == "login" {
					//更新用户的连接状态和ip地址
					serverNewLabel.SetText(ret.Ip)
					statusNewLabel.SetText("yes")
					//输出系统消息：xxx登录了
					cl.Append(widget.NewLabel(ret.Content))
				} else if ret.Type == "exit" {
					//输出系统消息：xxx下线了
					cl.Append(widget.NewLabel(ret.Content))
				} else if ret.Type == "user_list" {
					var str = "user_list\n"
					for _, s := range ret.Userlist {
						str += s + "\n"
					}
					ul.SetText(str)
				}
			}
		}
	}()
	container := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		ul,
		layout.NewSpacer(),
		cl,
		layout.NewSpacer())
	return container
}

//获得输入聊天内容行
func getLine4(multiEntry *widget.Entry) *fyne.Container {
	SendBtn := widget.NewButton("send", func() {
		service.SendService(multiEntry)
	})
	Line4 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		SendBtn)
	return Line4
}
