package osx

import (
	"fmt"
	"os"
	"time"
)

//OS 基础函数
func TestOsx() {
	hostName, err := os.Hostname()
	fmt.Printf("本机系统内核提供的主机名 os.Hostname():%s,error:%s \n", hostName, err)
	fmt.Println("预定义变量, 保存命令行参数 os.Args:", os.Args)
	fmt.Println("调用者所在进程的进程ID os.Getegid:", os.Getegid())
	fmt.Println("环境变量：")
	envs := os.Environ()
	for k, v := range envs {
		fmt.Println(k, v)
	}
	fmt.Println("单个环境变量 os.Getenv(\"PATH\")：", os.Getenv("PATH"))
	dir, err := os.Getwd()
	fmt.Println("当前目录  os.Getwd()：", dir, err)
	derr := os.Mkdir(dir+"/new1", 0755)
	fmt.Println("os.Mkdir(dir+\"/new1\", 0755):", derr)
	derr2 := os.Mkdir(dir+"/new2", 0755)
	fmt.Println("os.Mkdir(dir+\"/new2\", 0755):", derr2)
	rerr := os.Remove(dir + "/new1")
	rerr2 := os.Remove(dir + "/new2")
	fmt.Println("os.Remove(dir+\"/new1\") ", rerr)
	fmt.Println("os.Remove(dir+\"/new2\") ", rerr2)
	tmpdir := os.TempDir()
	fmt.Println("os.TempDir() ", tmpdir)
	file := dir + "/new"
	var fh *os.File
	fi, _ := os.Stat(file)
	if fi == nil {
		fh, _ = os.Create(file) // 文件不存在就创建
	} else {
		fh, _ = os.OpenFile(file, os.O_RDWR, 0666) // 文件存在就打开
	}
	w := []byte("hello go language" + time.Now().String())
	n, err := fh.Write(w)
	fmt.Println(n, err)
	// 设置下次读写位置
	ret, err := fh.Seek(0, 0)
	fmt.Println("当前文件指针位置", ret, err)
	b := make([]byte, 128)
	n, err = fh.Read(b)
	fmt.Println(n, err, string(b))
	defer func() {
		fh.Close()
	}()

	os.Remove(file)
	os.Remove(tmpdir)
	//终止
	os.Exit(1)
}
