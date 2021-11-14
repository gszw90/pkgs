package fileHander

import "os"

// IsFileExisted 文件或者目录是否存在
func IsFileExisted(filePath string) (ok bool, err error) {
	_, err = os.Stat(filePath)
	if err == nil {
		ok = true
		return
	}
	if os.IsNotExist(err) {
		err = nil
	}
	return
}

// CreateDir 创建一级或者多级目录
func CreateDir(dirPath string) (err error) {
	ok, err := IsFileExisted(dirPath)
	if ok {
		return err
	}
	if err != nil {
		return err
	}
	err = os.MkdirAll(dirPath, os.ModePerm)
	return err
}

// IsDir 判断是否为目录
func IsDir(dirPath string) (ok bool) {
	fi, err := os.Stat(dirPath)
	if err != nil {
		return
	}
	if fi.IsDir() {
		ok = true
	}
	return
}
