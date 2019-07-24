package zlog

import (
	"testing"
)

func TestStdZLog(t *testing.T) {

	//测试 默认debug输出
	Debug("lt_go_frame debug content1")
	Debug("lt_go_frame debug content2")

	Debugf(" lt_go_frame debug a = %d\n",10)

	//设置log标记位，加上长文件名称 和 微秒 标记
	ResetFlags(BitDate|BitLongFile|BitLevel)
	Info("lt_go_frame info content")

	//设置日志前缀，主要标记当前日志模块
	SetPrefix("MODULE")
	Error("lt_go_frame error content")

	//添加标记位
	AddFlag(BitShortFile|BitTime)
	Stack(" Zinx Stack! ")

	//设置日志写入文件
	SetLogFile("./log", "testfile.log")
	Debug("===> lt_go_frame debug content ~~666")
	Debug("===> lt_go_frame debug content ~~888")
	Error("===> lt_go_frame Error!!!! ~~~555~~~")

	//关闭debug调试
	CloseDebug()
	Debug("===> 我不应该出现~！")
	Debug("===> 我不应该出现~！")
	Error("===> lt_go_frame Error  after debug close !!!!")

}
