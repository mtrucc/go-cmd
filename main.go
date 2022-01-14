package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func pingSite(site string) {
	fmt.Println(site)
	sysType := runtime.GOOS

	cmd := &exec.Cmd{}

	if sysType == "linux" {
		// LINUX系统
		cmd = exec.Command("ping", "-c", "4", site)
	}

	if sysType == "windows" {
		// windows系统
		cmd = exec.Command("ping", site)
	}

	var cmdOut, cmdErr bytes.Buffer
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmdStdout := ConvertByte2String([]byte(cmdOut.String()), "GB18030")
	cmdStderr := ConvertByte2String([]byte(cmdErr.String()), "GB18030")
	fmt.Println("cmdStdout:", cmdStdout)
	fmt.Println("cmdStderr:", cmdStderr)
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

func main() {
	fmt.Println("Test")
	go pingSite("www.baidu.com")
	go pingSite("www.qq.com")
	go pingSite("www.hao123.com")
	//go pingSite("www.net.com")
	//go pingSite("www.bing.com")

	select {}
}
