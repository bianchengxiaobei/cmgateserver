package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"errors"
	"unsafe"
	"syscall"
)

const (
	VBoxLayout = iota
	GridLayout
)

type GateWindow struct {
	Win MainWindow
	Tab           *walk.TabWidget
	//开启
	ConnectButton *walk.PushButton
	ConnectLabel  *walk.Label
	//设置
	SettingButton *walk.PushButton
	SettingDialog *walk.Dialog
	SettingDialogSecond Dialog
	//日志
	logView				*LogView
}
type LogView struct {
	walk.WidgetBase
}
func (window *GateWindow) SetTitle(name string) {
	window.Win.Title = name
}

func (window *GateWindow) SetSize(width int, height int) {
	window.Win.Size = Size{
		Width:  width,
		Height: height,
	}
}
func (window *GateWindow) SetLayout(layout int) {
	switch layout {
	case VBoxLayout:
		window.Win.Layout = VBox{
		}
	case GridLayout:
		window.Win.Layout = Grid{}
	}
}
func (window *GateWindow) AddMenuItem(item MenuItem) {
	window.Win.MenuItems = append(window.Win.MenuItems, item)
}
func (window *GateWindow) AddToolBar(tool ToolBar) {
	window.Win.ToolBar = tool
}
func (window *GateWindow) AddToolBarMenuItem(item MenuItem) {
	if &(window.Win.ToolBar) == nil {
		panic("没有ToolBar")
	}
	window.Win.ToolBar.Items = append(window.Win.ToolBar.Items, item)
}
func (window *GateWindow) AddChildWidget(widget Widget) {
	window.Win.Children = append(window.Win.Children, widget)
}
func (window *GateWindow) AddTabPage(page *walk.TabPage) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if page == nil {
		fmt.Printf("nil")
	}
	window.Tab.Pages().Add(page)
}
func (window *GateWindow) AddTabPageContent(index int, widget walk.Widget) {
	page := window.Tab.Pages().At(index)
	if page == nil {
		fmt.Println("page == nil")
		return
	}
	if page.Children() == nil {
		fmt.Println("pageChild == nil")
		return
	}
	if widget == nil {
		fmt.Println("wid == nil")
		return
	}
	if err := page.Children().Add(widget); err != nil {
		//fmt.Println(err)
	}
}
func (window *GateWindow) SetTabPageCurrentIndex(index int) {
	window.Tab.SetCurrentIndex(index)
}
func (window *GateWindow) CreateTabWidget() TabWidget {
	var err error
	if window.Tab == nil {
		if window.Win.AssignTo != nil {
			if window.Tab, err = walk.NewTabWidget(*window.Win.AssignTo); err != nil {
				fmt.Println(err)
			}
		}
	}
	if window.Tab == nil {
		fmt.Println("nil ")
	}
	return TabWidget{
		AssignTo: &window.Tab,
	}
}
func (window *GateWindow) CreateConnectButton(content string, action walk.EventHandler) *walk.PushButton {
	var err error
	if window.ConnectButton == nil {
		if window.ConnectButton, err = walk.NewPushButton(window.Tab.Pages().At(0)); err != nil {
			fmt.Println(err)
		}
	}
	window.ConnectButton.SetText(content)
	window.ConnectButton.Clicked().Attach(action)
	return window.ConnectButton
}
func (window *GateWindow)CreateSettingButton(content string, action walk.EventHandler)  *walk.PushButton{
	window.SettingButton = CreatePushButton(window.Tab.Pages().At(1),content,action)
	return window.SettingButton
}
func (window *GateWindow)CreateSettingDialog(title string)  *walk.Dialog{
	var err error
	if window.SettingDialog,err = walk.NewDialog(*window.Win.AssignTo);err != nil{
		fmt.Println(err)
		return nil
	}
	window.SettingDialog.SetTitle(title)
	return window.SettingDialog
}
func CreateDialog(dialog *walk.Dialog) Dialog{
	return Dialog{
		AssignTo:&dialog,
	}
}
func CreatePushButton(parent walk.Container,content string, action walk.EventHandler)  *walk.PushButton{
	if btn,err := walk.NewPushButton(parent);err != nil{
		fmt.Println(err)
	}else{
		btn.SetText(content)
		btn.Clicked().Attach(action)
		return btn
	}
	return nil
}
func (window *GateWindow) CreateConnectLabel(connect string) *walk.Label{
	var err error
	if window.ConnectLabel == nil{
		if window.ConnectLabel,err = walk.NewLabel(window.Tab.Pages().At(0));err != nil{
			fmt.Println(err)
		}
	}
	window.ConnectLabel.SetText(connect)
	//window.connectLabel.SetSize(walk.Size{500,400})
	window.ConnectLabel.SetTextColor(walk.RGB(1,0,0))
	return window.ConnectLabel
}
//创建日志
func (window *GateWindow) CreateLogView() (*LogView,error){
	window.logView = &LogView{}
	if err := walk.InitWidget(
		window.logView,
		window.Tab.Pages().At(2),
		"EDIT",
		win.WS_TABSTOP|win.WS_VISIBLE|win.WS_VSCROLL|win.ES_MULTILINE|win.ES_WANTRETURN,
		win.WS_EX_CLIENTEDGE); err != nil {
		return nil, err
	}
	if 0 == window.logView.SendMessage(win.EM_SETREADONLY, uintptr(win.BoolToBOOL(true)), 0) {
		return window.logView, errors.New("fail to call EM_SETREADONLY")
	}
	window.logView.SendMessage(win.EM_SETLIMITTEXT, 4294967295, 0)
	return window.logView,nil
}
func (lv *LogView) AppendLog(content string) {
	txtLen := int(lv.SendMessage(0x000E, uintptr(0), uintptr(0)))
	lv.SendMessage(win.EM_SETSEL, uintptr(txtLen), uintptr(txtLen))
	point,_ := syscall.UTF16PtrFromString(content)
	lv.SendMessage(win.EM_REPLACESEL, 0, uintptr(unsafe.Pointer(point)))
}
func (lv *LogView) Write(p []byte) (int, error) {
	lv.AppendLog(string(p) + "\r\n")
	return len(p), nil
}
func (*LogView) LayoutFlags() walk.LayoutFlags {
	return walk.ShrinkableHorz | walk.ShrinkableVert | walk.GrowableHorz | walk.GrowableVert | walk.GreedyHorz | walk.GreedyVert
}
func (*LogView) MinSizeHint() walk.Size {
	return walk.Size{20, 12}
}
func (*LogView) SizeHint() walk.Size {
	return walk.Size{600, 500}
}
func (window *GateWindow) Run() {
	window.Win.Run()
}
func CreateGateWindow() GateWindow {
	win,_ := walk.NewMainWindow()
	gateWin := GateWindow{
		Win: MainWindow{},
	}
	gateWin.Win.AssignTo = &win
	return gateWin
}
func CreateToolBar(style ToolBarButtonStyle) ToolBar {
	return ToolBar{
		ButtonStyle: style,
	}
}

func CreateMenuItemWithImg(name string, imgPath string, action walk.EventHandler) MenuItem {
	return Menu{
		Text:        name,
		Image:       imgPath,
		OnTriggered: action,
	}
}
func CreateMenuItem(name string, action walk.EventHandler) MenuItem {
	return Menu{
		Text:        name,
		OnTriggered: action,
	}
}
func CreateCheckBox(name string, content string, checked bool) CheckBox {
	return CheckBox{
		Name:    name,
		Text:    content,
		Checked: checked,
	}
}
func CreateLinkLabel(content string, action walk.LinkLabelLinkEventHandler) LinkLabel {
	return LinkLabel{
		Text:            content,
		//MaxSize:         Size{300, 0},
		OnLinkActivated: action,
	}
}
func CreateTabPage(title string) *walk.TabPage {
	tabPage, err := walk.NewTabPage()
	tabPage.SetTitle(title)
	tabPage.SetLayout(walk.NewVBoxLayout())
	if err != nil {
		panic(err)
	}
	return tabPage
}
