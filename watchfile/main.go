package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/hpcloud/tail"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	var source, target string
	var reopen, poll, pipe bool
	flag.StringVar(&source, "source", "test.log", "source log file name")
	flag.StringVar(&target, "target", "test-utf8.log", "target log file name")
	flag.BoolVar(&reopen, "reopen", false, "reopen")
	flag.BoolVar(&poll, "poll", false, "poll")
	flag.BoolVar(&pipe, "pipe", false, "pipe")
	flag.Parse()

	//准备输入文件
	t, err := tail.TailFile(source, tail.Config{
		Follow: true,
		ReOpen: reopen,
		Poll:   poll,
		Pipe:   pipe,
		Location: &tail.SeekInfo{
			Whence: os.SEEK_END,
		},
	})
	if err != nil {
		panic(err)
	}

	var eol string
	switch runtime.GOOS {
	case "windows":
		eol = "\r\n"
	default:
		eol = "\n"
	}

	if target == "" {
		fmt.Printf("--target can not be empty")
		return
	}

	var bufferData Buffer = Buffer{
		target: target,
	}

	//start flush
	log.Printf("start flush")
	bufferData.Flush()

	//监视源文件
	for line := range t.Lines {

		//gbk 转 utf8
		buf, err := GbkToUtf8([]byte(line.Text))
		if err != nil {
			log.Println("gbk to utf8 error: ", err)
			return
		}

		//输出到buffer
		bufferData.Write(buf)
		bufferData.Write([]byte(eol))
	}

}

///////////////////////////////////////////////////////////////////////////
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

///////////////////////////////////////////////////////////////////////////
func CheckFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Printf("%s not exist", filename)
		return false
	}
	return true
}

///////////////////////////////////////////////////////////////////////////
// Buffer is a goroutine safe bytes.Buffer
type Buffer struct {
	target string
	buffer bytes.Buffer
	mutex  sync.Mutex
}

// Write appends the contents of p to the buffer, growing the buffer as needed. It returns
// the number of bytes written.
func (s *Buffer) Write(p []byte) (n int, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.buffer.Write(p)
}

// String returns the contents of the unread portion of the buffer
// as a string.  If the Buffer is a nil pointer, it returns "<nil>".
func (s *Buffer) String() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.buffer.String()
}

// read all data from cache , reset buffer after read
func (s *Buffer) ReadAll() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	rlt := s.buffer.String()
	s.buffer.Reset()
	return rlt
}

func (s *Buffer) Flush() {
	s.writeLog(s.ReadAll())
	time.AfterFunc(time.Duration(30*time.Second), s.Flush)
}

func (s *Buffer) writeLog(content string) {
	var f *os.File
	var err error

	if len(content) == 0 {
		//log.Printf("no data to write, ignore")
		return
	}

	filename := s.target

	if CheckFileExist(filename) { //文件存在
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModePerm) //打开文件
		if err != nil {
			log.Println("file open fail", err)
			return
		}
	} else { //文件不存在
		f, err = os.Create(filename) //创建文件
		if err != nil {
			log.Println("file create fail")
			return
		}
	}

	//将文件写进去
	n, err1 := io.WriteString(f, content)
	if err1 != nil {
		log.Printf("write error: %v", err1)
		return
	}

	log.Printf("write buffer to file, size:%d", n)

	f.Close()
}
