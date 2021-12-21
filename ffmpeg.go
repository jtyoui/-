// Package main
// @Time  : 2021/12/20 下午5:28
// @Author: Jtyoui@qq.com
// @note  : 将其它格式转为mp4
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

type fileMessage string

var wg = &sync.WaitGroup{}

// toMP4 自动转
func (f fileMessage) toMP4() {
	defer wg.Done()
	if f.isFile() {
		suffix := f.suffix()
		name := suffix + "ToMP4"
		value := reflect.ValueOf(f)
		m := value.MethodByName(name)
		if m.IsValid() && !m.IsNil() {
			m.Call(nil)
		}
	}
}

func (f fileMessage) suffix() string {
	name := path.Ext(f.String())[1:]
	return strings.ToUpper(name[:1]) + name[1:]
}

func (f fileMessage) prefix() string {
	base := path.Base(f.String())
	return base[:len(base)-len(f.suffix())-1]
}

func (f fileMessage) dir() string {
	return path.Dir(f.String())
}

func (f fileMessage) isFile() bool {
	s, err := os.Stat(f.String())
	if err != nil {
		return false
	}
	return !s.IsDir()
}

func (f fileMessage) String() string {
	return string(f)
}

func (f fileMessage) RmvbToMP4() {
	fileName := f.dir() + "/" + f.prefix() + ".mp4"
	cmdArguments := []string{"-i", f.String(), "-c:v", "libx264", "-strict", "-2", fileName}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Remove(fileName)
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}
}

func files(dir string) []string {
	var f []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		f = append(f, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return f
}

func main() {
	dir := flag.String("d", "", "视频文件夹")
	flag.Parse()
	ls := files(*dir)
	wg.Add(len(ls))
	for _, file := range ls {
		message := fileMessage(file)
		go message.toMP4()
	}
	wg.Wait()
}
