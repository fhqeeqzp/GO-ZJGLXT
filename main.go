package main

import (
	"context"
	"embed"
	"fmt"

	"go-wails-admin/internal/config"
	"go-wails-admin/internal/database"
	"go-wails-admin/internal/models"
	"go-wails-admin/internal/services"
	"go-wails-admin/internal/utils"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:             "工程造价管理系统",
		Width:             1280,
		Height:            768,
		MinWidth:          800,
		MinHeight:         600,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         true, // 无边框
		BackgroundColour:  &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:   windows.RGB(27, 38, 54),
				DarkModeTitleText:  windows.RGB(255, 255, 255),
				DarkModeBorder:     windows.RGB(27, 38, 54),
				LightModeTitleBar:  windows.RGB(240, 240, 240),
				LightModeTitleText: windows.RGB(0, 0, 0),
				LightModeBorder:    windows.RGB(240, 240, 240),
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// App struct
type App struct {
	ctx         context.Context
	config      *config.Config
	userService *services.UserService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}
	a.config = cfg

	// Initialize database
	if err := database.InitDB(&cfg.Database); err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		return
	}

	// Initialize services
	a.userService = services.NewUserService()

	fmt.Println("Application started successfully!")
}

// shutdown is called when the app is closing
func (a *App) shutdown(ctx context.Context) {
	fmt.Println("Application is shutting down...")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, Welcome to Go Wails Admin!", name)
}

// GetUserList returns paginated user list
func (a *App) GetUserList(page, pageSize int) utils.Response {
	users, total, err := a.userService.GetUserList(page, pageSize)
	if err != nil {
		return utils.Error(err.Error())
	}

	return utils.Success(map[string]interface{}{
		"list":  users,
		"total": total,
	})
}

// GetUserByID returns a user by ID
func (a *App) GetUserByID(id uint) utils.Response {
	user, err := a.userService.GetUserByID(id)
	if err != nil {
		return utils.Error(err.Error())
	}
	return utils.Success(user)
}

// CreateUser creates a new user
func (a *App) CreateUser(user models.User) utils.Response {
	if err := a.userService.CreateUser(&user); err != nil {
		return utils.Error(err.Error())
	}
	return utils.Success(user)
}

// UpdateUser updates an existing user
func (a *App) UpdateUser(user models.User) utils.Response {
	if err := a.userService.UpdateUser(&user); err != nil {
		return utils.Error(err.Error())
	}
	return utils.Success(user)
}

// DeleteUser deletes a user by ID
func (a *App) DeleteUser(id uint) utils.Response {
	if err := a.userService.DeleteUser(id); err != nil {
		return utils.Error(err.Error())
	}
	return utils.Success(nil)
}

// GetConfig returns current configuration
func (a *App) GetConfig() utils.Response {
	return utils.Success(a.config)
}

// ==================== 窗口控制方法 ====================

// WindowMinimize 最小化窗口
func (a *App) WindowMinimize() {
	runtime.WindowMinimise(a.ctx)
}

// WindowMaximize 最大化/恢复窗口
func (a *App) WindowMaximize() {
	if runtime.WindowIsMaximised(a.ctx) {
		runtime.WindowUnmaximise(a.ctx)
	} else {
		runtime.WindowMaximise(a.ctx)
	}
}

// WindowClose 关闭窗口
func (a *App) WindowClose() {
	runtime.Quit(a.ctx)
}

// WindowIsMaximised 检查窗口是否已最大化
func (a *App) WindowIsMaximised() bool {
	return runtime.WindowIsMaximised(a.ctx)
}

// WindowSetAlwaysOnTop 设置窗口置顶
func (a *App) WindowSetAlwaysOnTop(alwaysOnTop bool) {
	runtime.WindowSetAlwaysOnTop(a.ctx, alwaysOnTop)
}

// WindowSetSize 设置窗口大小
func (a *App) WindowSetSize(width, height int) {
	runtime.WindowSetSize(a.ctx, width, height)
}

// WindowCenter 窗口居中
func (a *App) WindowCenter() {
	runtime.WindowCenter(a.ctx)
}
