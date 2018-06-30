package main

import (
	"cmgateserver/gateserver"
	"cmgateserver/gui"
	"fmt"
	"github.com/bianchengxiaobei/cmgo/log4g"
	"github.com/lxn/walk"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
)

var (
	cmd      = "url.dll,FileProtocolHandler"
	runDll32 = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
)
var (
	startGateAction   walk.EventHandler
	settingGateAction walk.EventHandler
	linkAction        walk.LinkLabelLinkEventHandler
)
var server *gateserver.GateServer
var window gui.GateWindow

func main() {
	log4g.LoadConfig("/logConfig.txt")
	startGateAction = StartOrCloseGate
	settingGateAction = SettingGate
	linkAction = OnClickLink
	window = gui.CreateGateWindow()
	window.Win.Init = func() {
		window.AddTabPage(gui.CreateTabPage("开启"))
		window.AddTabPage(gui.CreateTabPage("设置"))
		window.AddTabPage(gui.CreateTabPage("日志"))

		connectLabel := window.CreateConnectLabel("网关未开启")
		connectButton := window.CreateConnectButton("开启", startGateAction)
		window.AddTabPageContent(0, connectLabel)
		window.AddTabPageContent(0, connectButton)
		//设置
		settingBtn := window.CreateSettingButton("网关session设置", ClickGateSessionConfig)
		window.AddTabPageContent(1, settingBtn)
		//日志
		if logview, err := window.CreateLogView(); err != nil {
			panic(err)
		} else {
			log.SetOutput(logview)
		}
	}
	window.SetSize(800, 600)
	window.SetLayout(gui.VBoxLayout)
	window.AddChildWidget(gui.CreateLinkLabel(`网关服务器github地址: <a id="github" href="https://github.com/bianchengxiaobei/cmgo">链接</a>`, linkAction))
	window.AddChildWidget(window.CreateTabWidget())

	server = gateserver.NewGateServer()
	server.Init("gateBaseConfig.txt", "gateSesssionConfig.txt", "innnerSessionConfig.txt")
	window.SetTitle(fmt.Sprintf(server.Name, server.Id))
	window.Run()
}
func StartOrCloseGate() {
	if server.IsRunning == false {
		server.Run()
		window.ConnectButton.SetText("停止")
		window.ConnectLabel.SetText("网关已经启动!")
	} else {
		server.Close()
		window.ConnectButton.SetText("开启")
		window.ConnectLabel.SetText("网关未开启!")
	}

}
func SettingGate() {
	fmt.Println("设置")
}
func OnClickLink(link *walk.LinkLabelLink) {
	fmt.Println(link.URL())
	cmdExe := exec.Command(runDll32, cmd, link.URL())
	if err := cmdExe.Start(); err != nil {
		fmt.Println(err)
	}
}
func ClickGateSessionConfig() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("fefeqqq3")
			return
		}
	}()
	if window.SettingDialog == nil {
		dialog := window.CreateSettingDialog("设置")
		window.SettingDialogSecond = gui.CreateDialog(dialog)
		dialog.SetLayout(walk.NewVBoxLayout())
		dialog.SetSize(walk.Size{300, 300})
		//fmt.Println(dialog)
		window.SettingDialogSecond.Init = func() {
			fmt.Println("fefef")
			com, _ := walk.NewComposite(dialog)
			com.SetLayout(walk.NewGridLayout())
			dialog.Children().Add(com)
			v := reflect.ValueOf(server.UserClientServer.SessionConfig).Elem()
			t := reflect.TypeOf(server.UserClientServer.SessionConfig).Elem()
			count := v.NumField()
			for i := 0; i < count; i++ {
				f := v.Field(i)
				k := t.Field(i)
				switch f.Kind() {
				case reflect.String:
					label, _ := walk.NewLabel(com)
					label.SetText(k.Name)
					line, _ := walk.NewLineEdit(com)
					line.SetText(f.String())
					fmt.Println(k.Name)
					com.Children().Add(label)
					com.Children().Add(line)
				case reflect.Int:
					label, _ := walk.NewLabel(com)
					label.SetText(k.Name)
					num, _ := walk.NewNumberEdit(com)
					num.SetValue(float64(f.Int()))
					com.Children().Add(label)
					com.Children().Add(num)
				case reflect.Bool:
					label, _ := walk.NewLabel(com)
					label.SetText(k.Name)
					check, _ := walk.NewCheckBox(com)
					check.SetChecked(f.Bool())
					com.Children().Add(label)
					com.Children().Add(check)
				}
			}
		}
	}
	window.SettingDialogSecond.Run(*window.Win.AssignTo)
}
