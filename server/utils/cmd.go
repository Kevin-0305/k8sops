package utils

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

//对字符进行转码
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

func handlerErr(errReader io.ReadCloser) {
	in := bufio.NewScanner(errReader)
	for in.Scan() {
		cmdRe := ConvertByte2String(in.Bytes(), "UTF8")
		fmt.Errorf(cmdRe)
	}
}

func CmdSync(userCmd string, LogCh chan<- string) (string, error) {
	var cmd *exec.Cmd
	//var outLog = ""
	// 执行单个shell命令时, 直接运行即可
	LogCh <- fmt.Sprintf("执行命令：", userCmd)
	cmd = exec.Command("bash", "-c", userCmd)
	//stdoutStderr, err := cmd.CombinedOutput()
	stdout, err := cmd.StdoutPipe()
	errReader, errr := cmd.StderrPipe()

	if errr != nil {
		fmt.Println(err)
		return "", nil
	}

	go handlerErr(errReader)

	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	cmd.Start()
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		cmdRe := ConvertByte2String(in.Bytes(), "UTF8")
		LogCh <- fmt.Sprintf(cmdRe)
		// outLog = outLog + cmdRe
	}
	cmd.Wait()
	return "", nil
}

func CmdExecute(userCmd string, LogCh chan<- string) (string, error) {
	var cmd *exec.Cmd
	//var outLog = ""
	LogCh <- fmt.Sprintf("执行命令：", userCmd)
	cmd = exec.Command("/bin/bash", "-c", userCmd)
	stdout, err := cmd.Output()
	if err != nil {
		LogCh <- fmt.Sprintf("error", err.Error()) 
		// outLog = outLog + err.Error()
	}
	LogCh <- string(stdout)
	//outLog = outLog + string(stdout)
	return "", nil
}
