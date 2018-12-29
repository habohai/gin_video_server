package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path"
)

// GetSize 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist 检查文件是否存在， 文件存在返回true
func CheckExist(src string) (bool, error) {
	_, err := os.Stat(src)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	exist, err := CheckExist(src)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	if err := MkDir(src); err != nil {
		return err
	}

	return nil
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// IsToOpen 会打开文件，文件路径不存在会创建, 当文件不存在时打开失败
func IsToOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	fileFullPath := src + fileName
	ok, err := CheckExist(fileFullPath)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("file not exist")
	}

	f, err := Open(fileFullPath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

// DeleteFile 会打开文件，文件路径不存在会创建, 当文件不存在时打开失败
func DeleteFile(fileName, filePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	fileFullPath := src + fileName

	err = os.Remove(fileFullPath)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}

	return nil
}

// MustOpen 一定会打开文件，文件路径，文件不存在都会创建
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("111Fail to OpenFile :%v", err)
	}

	return f, nil
}
