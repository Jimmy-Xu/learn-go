package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

var (
	method            int
	filename, content string
)

func main() {
	flag.IntVar(&method, "method", 0, "method of write file, 0,1,2,3")
	flag.StringVar(&filename, "filename", "test.log", "filename to write")
	flag.StringVar(&content, "content", "", "content to write to file")
	flag.Parse()

	content = fmt.Sprintf("%s %s %d\n", time.Now().Format("2006-01-02 15:04:05.000"), content, method)

	switch method {
	case 0:
		Write0(filename, content)
    case 1:
		Write1(filename, content)
    case 2:
		Write2(filename, content)
    case 3:
		Write3(filename, content)
    default:
		fmt.Println("method is invlaid")
		os.Exit(1)
	}

}

func Write0(filename, content string) {
	var f *os.File
	var err error

	if CheckFileExist(filename) { //文件存在
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModePerm) //打开文件
		if err != nil {
			fmt.Println("file open fail", err)
			return
		}
	} else { //文件不存在
		f, err = os.Create(filename) //创建文件
		if err != nil {
			fmt.Println("file create fail")
			return
		}
	}

	//将文件写进去
	n, err1 := io.WriteString(f, content)
	if err1 != nil {
		fmt.Println("write error: ", err1)
		return
	}
	fmt.Println("写入的字节数是：", n)

	f.Close()
}

func Write1(filename, content string) {
	var d = []byte(content)
	err := ioutil.WriteFile(filename, d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
	fmt.Println("write success")
}

func Write2(filename, content string) {
	var d1 = []byte(content)

	f, err3 := os.Create(filename) //创建文件
	if err3 != nil {
		fmt.Println("create file fail")
	}
	defer f.Close()
	n2, err3 := f.Write(d1) //写入文件(字节数组)
	fmt.Printf("写入 %d 个字节n", n2)

	f.Sync()
}

func Write3(filename, content string) {
	f, err3 := os.Create(filename) //创建文件
	if err3 != nil {
		fmt.Println("create file fail")
	}
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, err3 := w.WriteString(content)
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()
}

func CheckFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("%s not exist\n", filename)
		return false
	}
	fmt.Printf("%s is exist\n", filename)
	return true
}
