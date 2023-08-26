package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type finderStruct struct {
	sync.Mutex
	items []*FsInfo
}

type FsInfo struct {
	Name  string
	Path  string
	IsDir bool
}

var finderObj *finderStruct = &finderStruct{
	items: make([]*FsInfo, 0),
}

func Find(pattern, root string, jumpHiddenItem bool) ([]*FsInfo, error) {
	err := find(pattern, root, jumpHiddenItem, func(foo, path string, info fs.FileInfo) (bool, error) {
		//正则判断
		s := filepath.Base(path)
		pattern = strings.ToLower(pattern)
		s = strings.ToLower(s)
		matched, err := regexp.MatchString(pattern, s)
		if err != nil {
			return false, err
		}
		return matched, nil
	})
	items := finderObj.items
	finderObj.items = make([]*FsInfo, 0)
	return items, err
}

func find(foo, root string, jumpHiddenItem bool, f func(foo, path string, info fs.FileInfo) (bool, error)) error {
	return _find(foo, root, true, jumpHiddenItem, f)
}

var wg = sync.WaitGroup{}

func _find(foo, root string, main bool, jumpHiddenItem bool, f func(foo, path string, info fs.FileInfo) (bool, error)) error {

	dirEntry, err := os.ReadDir(root)
	if err != nil {
		//权限问题，跳过，直接返回
		if !main {
			wg.Done()
		}
		return nil
	}
	for _, entry := range dirEntry {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		path := filepath.Join(root, info.Name())
		// //忽略隐藏文件，排除项目
		if jumpHiddenItem && strings.HasPrefix(info.Name(), ".") {
			continue
		}

		//若是文件夹
		if info.IsDir() {
			wg.Add(1)
			go func(foo, path string, f func(foo string, path string, info fs.FileInfo) (bool, error)) {
				err := _find(foo, path, false, jumpHiddenItem, f)
				if err != nil {
					fmt.Println("Error: ", err)
					os.Exit(1)
				}
			}(foo, path, f)
		}
		//接口函数
		flag, err := f(foo, path, info)
		if err != nil {
			return err
		}
		if flag {
			finderObj.Lock()
			finderObj.items = append(finderObj.items, &FsInfo{info.Name(), path, info.IsDir()})
			finderObj.Unlock()
		}
	}
	if main { //若主线程则等待，待子进程全部返回后打印
		wg.Wait()
		displayItems()
	} else { //否则副线程完毕
		wg.Done()
	}
	return nil
}

func displayItems() {
	sort.Slice(finderObj.items, func(i, j int) bool {
		return finderObj.items[i].Name < finderObj.items[j].Name
	})
}
