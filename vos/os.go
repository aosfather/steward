package vos

import (
	"fmt"
	"time"
	"github.com/aosfather/steward/bots"
)

/**
  语音控制流程
1、检测唤醒词
2、如果被唤醒，开始语音识别
3、对于识别的结果进行处理(A 类流程)，中途反馈执行状态
4、对于B类，通过中断来检测是否有唤醒的行为
 */

 //控制系统API
 type OsApi interface {
 	Say(text string) //使用声音提醒用户
 	IsWakeUp(voice []byte)bool //是否唤醒
 	//调用bot
    CallBot(botname string,paramters ... string)(string,error)

 }


 type OS struct {

    cmdbots []CmdBot
 	recognizer RecognizeFunction
 }

 func (this *OS)Say(text string){
 	//使用声音说话
 }

 //主循环
 func (this *OS) loop() {
 	for {

 		//检测唤醒
 		if this.detectWakeup() {
 			data:=this.recordVoice(8) //录下8秒的指令

 			if this.recognizer!=nil {
 				text:=this.recognizer(data)

 				if !this.doCommand(text) {
                    this.talk(text)
				}


			}


           continue
		}
 		fmt.Println("没有唤醒....")


	}

 }

 //录音
 func (this *OS) recordVoice(second int)[]byte {

 	return nil
 }

 //检测是否唤醒
 func (this *OS) detectWakeup() bool{

    time.Sleep(time.Second*1)
 	return true
 }

 //执行命令
 func (this *OS) doCommand(text string) bool {
 	for _,cmdbot:=range this.cmdbots{
 		istrigger,paramters:=cmdbot.IsTrigger(text)
 		if istrigger {
 			if cmdbot.IsRunning() {
 				fmt.Println("running")
 				this.Say("指令还在运行中")
			}else {
				go this.runbot(cmdbot,paramters)
			}
			return true
		}
	}

 	return false
 }

 func (this *OS) runbot(bot CmdBot,paramters...string) {
 	fmt.Println(bot.Run(paramters))


 }

 //聊天
 func (this *OS) talk(text string) {
    result,err:=bots.RunTulingTalkBot("主人",text)
    if err!=nil {
    	result="我不懂你说的"
	}

	fmt.Println(result)
	this.Say(result)


 }
