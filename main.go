package main

import (
	"embed"
	"flag"
	"log"

	"github.com/Ericwyn/pancors"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed mcp-partner/dist
var assets embed.FS

//go:embed mcp-partner/public
var iconAssets embed.FS

const DeskVersion = "v0.0.1"

func main() {
	go func() {
		pancors.RunPancorsServ(":36875")
	}()

	flag.Parse()

	// 在启动前设置环境变量，尝试禁用硬件加速来验证是否是渲染引擎问题
	//os.Setenv("WEBKIT_DISABLE_COMPOSITING_MODE", "1")

	// Create an instance of the app structure
	app := NewApp()

	//// ---------------------- 新增菜单逻辑开始 ----------------------
	//appMenu := menu.NewMenu()
	//
	//toolsMenu := appMenu.AddSubmenu("Help")
	//toolsMenu.AddText("Press Ctrl/Cmd+Shift+F12 Open devtools", nil, nil)
	//toolsMenu.AddText("Version: "+DeskVersion, nil, nil)
	//// 添加分割线
	//toolsMenu.AddSeparator()

	// ---------------------- 新增菜单逻辑结束 ----------------------

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "mcp-partner",
		Width:     1400,
		Height:    900,
		MinWidth:  1024,
		MinHeight: 768,
		//MaxWidth:          1480,
		//MaxHeight:         800,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            assets,
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		Bind: []interface{}{
			app,
		},
		Linux: &linux.Options{
			ProgramName: "mcp-partner",
			Icon:        readIconBytes("mcp-partner/public/icon_128px.png"),
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "mcp-partner",
				Message: "",
				Icon:    readIconBytes("mcp-partner/public/icon_128px.png"),
			},
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func readIconBytes(iconNameFile string) []byte {
	iconBytes, err := iconAssets.ReadFile(iconNameFile)
	if err != nil {
		log.Fatal(err)
	}
	return iconBytes
}
