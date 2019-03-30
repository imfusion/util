package util

import "os"

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
