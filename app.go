package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

var current, _ = os.UserHomeDir()

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) FuseSearch(input string, jumpHiddleItem bool) []*FsInfo {
	infos, err := Find(input, current, jumpHiddleItem)
	if err != nil {
		return nil
	}
	return infos
}

func (a *App) AbsoluteSearch(input string, jumpHiddleItem bool) []*FsInfo {
	dir := filepath.Dir(input)
	// 若是文件夹则保持原样
	info, err := os.Stat(input)
	if err == nil {
		if info.IsDir() {
			dir = input
		}
	}
	// 自动切换当前目录
	a.ChangeCurrentDir(dir)
	entry, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	r := make([]*FsInfo, 0, len(entry))
	for _, item := range entry {
		path := filepath.Join(dir, item.Name())
		if jumpHiddleItem && strings.HasPrefix(filepath.Base(path), ".") {
			continue
		}
		if strings.HasPrefix(path, input) {
			info := &FsInfo{
				Name:  item.Name(),
				Path:  path,
				IsDir: item.IsDir(),
			}
			r = append(r, info)
		}
	}
	return r
}

func (a *App) BackToLastDir() string {
	a.ChangeCurrentDir(filepath.Dir(current))
	return a.GetCurrentDir()
}

func (a *App) ChangeCurrentDir(to string) {
	current = to
}

func (a *App) GetCurrentDir() string {
	return current
}

func (a *App) Open(path string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", path)
	case "windows":
		cmd = exec.Command("cmd",`\c`,"start", "", path)
	case "darwin":
		cmd = exec.Command("open", path)
	}
	go cmd.Start()
}

func (a *App) GetBase64ImageSrc(path string) string {
	base, _ := convertToBase64(path, 150)
	// 添加前缀
	return "data:image/jpeg;base64," + base
}
