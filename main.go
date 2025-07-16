package main

import (
	"org.mutantcat.MagicLink/config"
	"os"
)

// MaginLink 是一个系统命令行补全工具
// By；mutantcat.org
// 功能：可以通过类似busybox那样的符号链接的方式补全指令
// 1、借助busybox等工具补全指令
// 2、使用bash补全脚本（文件夹中对应的脚本文件作为指令结果）
// 3、返回死数据补全操作（指定文本文件未知作为补全位置）
func main() {
	// 首先进行参数获取和处理
	command := make([]string, 0)
	exec_name := config.GetExecName()
	exec_path := config.GetExecPath()

	// 先处理一下argv[0]的问题
	// 这里需要进行一种检测，就是如果我是自身程序开头的就跳过这个参数作为command
	// 如果是从符号链接来的就将argv[0]作为第一个参数
	if os.Args[0] == exec_name || os.Args[0] == exec_path {
		// 如果只有一个可执行程序参数就打印Logo啥的
		if len(os.Args) == 1 {
			println("MagicLink - 一款多功能Linux系统命令行补全工具")
			println("___  ___            _      _     _       _    \n|  \\/  |           (_)    | |   (_)     | |   \n| .  . | __ _  __ _ _  ___| |    _ _ __ | | __\n| |\\/| |/ _` |/ _` | |/ __| |   | | '_ \\| |/ /\n| |  | | (_| | (_| | | (__| |___| | | | |   < \n\\_|  |_/\\__,_|\\__, |_|\\___\\_____/_|_| |_|_|\\_\\\n               __/ |                          \n              |___/                           ")
			println("- Usage: " + exec_name + " [command] [options]")
			println("- Usage: 使用sudo ln -s magiclink xxx 来将xxx指令挂载到系统上")
			println("- Usage: 可以将程序放置到/bin目录下，或者将程序放置到PATH环境变量中")
			println("- Usage: 将需要补全的指令放置到本程序二进制程序同级的/mlink文件夹或需要执行指令的文件夹下的/mlink文件夹")
			println("- Usage: /mlink下的文件夹分为exe（可执行文件）、sh（bash脚本）、static（文本文件）等")
			println("- Usage: 若上述文件夹中有重名脚本则会按static > bash > exe的顺序依次执行")
			println("- Usage: 直接在需要指令的文件夹下的/mlink文件夹中的脚本优先级更高")
			println("- Usage: 文件名即是指令名且可挂载全局系统，详情见 https://www.mutantcat.org/software/magiclink")
			println("- Info : 本产品由异猫工作群（www.mutantcat.org）提供维护支持")
			println("- Version: V1.0.20250716")
			return
		}
		// 如果是自身程序开头的就跳过这个参数
		command = append(command, os.Args[1:]...)
	} else {
		// 如果不是自身程序开头的就将argv[0]作为第一个参数
		command = append(command, os.Args[0])
		command = append(command, os.Args[1:]...)
	}

	// 进行指令处理
	MakeCase(command)
}
