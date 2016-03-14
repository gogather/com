package com

import (
	"errors"
	// "github.com/gogather/com/log"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func PathExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 读取文件
func ReadFileByte(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	return ioutil.ReadAll(fi)
}

// 读取文本文件
func ReadFileString(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}

	fd, err := ioutil.ReadAll(fi)

	return string(fd), err
}

// 读取文本文件
func ReadFile(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)

	return string(fd), err
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 字符串写入文件
func WriteFile(fullpath string, str string) error {
	data := []byte(str)
	return ioutil.WriteFile(fullpath, data, 0644)
}

// 创建文件夹
func Mkdir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

// Copyfile
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func MkdirWithCreatePath(fullpath string) error {
	dirArray := strings.Split(fullpath, string(filepath.Separator))
	if len(dirArray) < 1 {
		return errors.New("fullpath mistake")
	} else if len(dirArray) == 1 {
		Mkdir(fullpath)
		return nil
	}

	path := ""
	var err error = nil
	for i := 0; i < len(dirArray); i++ {
		dir := dirArray[i]
		path = path + string(filepath.Separator) + dir
		p := SubString(path, 1, len(path)-1)
		if !PathExist(p) {
			err = Mkdir(p)
		}
	}

	return err
}

// 写入文件，若目录不存在则自动创建
func WriteFileWithCreatePath(fullpath string, str string) error {
	fileDir := ""
	var err error = nil
	if !FileExist(fullpath) {
		fileDir = filepath.Dir(fullpath)
		err = MkdirWithCreatePath(fileDir)
	}
	err = WriteFile(fullpath, str)

	return err
}
