package util

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Exist 判断路径/文件是否存在
func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsDir 判断是否是文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	return err == nil && s.IsDir()
}

// IsFile 判断是否是文件
func IsFile(path string) bool {
	s, err := os.Stat(path)
	return err == nil && !s.IsDir()
}

// MkDir 创建目录，如果目录存在则直接返回
func MkDir(path string) error {
	if !Exist(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

// CurrentPath 获取当前主程序的路径
func CurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}

	i := strings.LastIndex(path, string(os.PathSeparator))

	return string(path[0:i]), err
}

// ReadLine 按行读取文件返回所有行的 slice 切片
func ReadLine(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	var lines []string
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	return lines, nil
}
