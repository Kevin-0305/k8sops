package utils

import (
	"fmt"
	"gin-vue-admin/model"
)

func GitClone(pro model.Project, proWorkSpace string, LogCh chan<- string) (string, error) {
	cmd := fmt.Sprintf("git clone -b master %s %s", pro.WareHouse, proWorkSpace)
	LogCh <- fmt.Sprintf("第一步：检出代码：", cmd)
	cmdOut, err := CmdSync(cmd, LogCh)
	LogCh <- fmt.Sprintf(string(cmdOut))
	if err != nil {
		LogCh <- fmt.Sprintf("第一步：检出代码执行失败：", err)
		return "", err
	}
	return string(cmdOut), nil
}

func GitPull(pro model.Project, proWorkSpace string, LogCh chan<- string) (string, error) {
	gitpull := fmt.Sprintf("cd %s; git pull origin %s; git log --oneline -2", proWorkSpace, "master")
	cmdOut, err := CmdSync(gitpull, LogCh)
	LogCh <- fmt.Sprintf(string(cmdOut))
	if err != nil {
		LogCh <- fmt.Sprintf("拉取代码失败：", err)
		return "", err
	}
	return string(cmdOut), nil
}
