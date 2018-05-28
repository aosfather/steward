package vos

import (
	"github.com/aosfather/steward/bots"
	"fmt"
	"strings"
)


type CmdBot interface {
	IsTrigger(text string) (bool,string) //是否触发
    IsRunning()bool             //是否在执行
	Run(target...interface{})(string,error) //执行

}

/**
 系统默认的命令bot
 */
type SimpleCmdBot struct {
	Name string
	KeyWords []string
	Path string
	Cmd  string
	OptionTemplate string    //参数模板
	running bool
}

func (this *SimpleCmdBot) IsTrigger(text string) (bool,string) {
   for _,key :=range this.KeyWords{
	   if strings.HasPrefix(text,key) {
	   	return true,text[len(key):len(text)]
	   }
   }
    return false,text
   }

func (this *SimpleCmdBot)IsRunning()bool{
	return this.running
}

func (this *SimpleCmdBot) Run(target...interface{})(string,error) {
	fmt.Println(this.Name," is begin running ")
	option:=fmt.Sprintf(this.OptionTemplate,target...)
	this.running=true
	defer  this.stop()
	return bots.RunCmdBot(this.Path,this.Cmd,option)
}

func (this *SimpleCmdBot) stop(){
	this.running=false
	fmt.Println(this.Name," is end running ")
}




