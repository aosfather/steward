package vos

import "testing"

func mytestrecorgniz(data []byte)string {
	return "test"
}
func TestOS_Say(t *testing.T) {
	o:=OS{}
	o.recognizer=mytestrecorgniz
	o.cmdbots=append(o.cmdbots,&SimpleCmdBot{"test",[]string{"test"},"","echo","haha %s",false})
	o.loop()
}
