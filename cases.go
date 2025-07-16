package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func MakeCase(command []string) {
	switch command[0] {
	// 创建
	case "makelink":
		println("Creating symbolic link...")

		break
	// 检查busybox指令是否存在
	case "busybox_check":
		checkBusybox()
		break
	// 安装busybox指令
	case "busybox_install":
		installBusybox()
		break
	// 挂载busybox指令
	case "busybox_mount":
		mountBusybox()
		break
	default:
		doStatic(command)
		doSH(command)
		doExe(command)
		break
	}

}

// doStatic 处理静态补全
func doStatic(command []string) {
	// 如果不是以上指令，先看一下当前程序目录下的./mlink/static下的无拓展名文本文件名，并返回相应的内容
	staticProxy := "./mlink/static/" + command[0] //制定到写死的文本文件
	//println("正在处理静态补全: " + staticProxy)
	// 检查文件是否存在
	if _, err := os.Stat(staticProxy); os.IsNotExist(err) {
		// 如果当前目录下不存在 就去可执行文件所在的文件夹里面获取
		execPath, err := os.Executable()
		if err != nil {
			return
		}
		staticProxy = filepath.Join(filepath.Dir(execPath), "mlink", "static", command[0])
		// 检查文件是否存在
		if _, err := os.Stat(staticProxy); os.IsNotExist(err) {
			return
		}
	}
	// 如果文件存在，读取文件内容并返回
	file, err := os.Open(staticProxy)
	if err != nil {
		return
	}
	defer file.Close()
	content, err := os.ReadFile(staticProxy)
	if err != nil {
		return
	}
	fmt.Println(string(content))
}

// doSH 处理shell脚本补全
func doSH(command []string) {
	// 检查当前目录下是否有对应的shell脚本文件
	shellProxy := "./mlink/sh/" + command[0] + ".sh" // 指定到写死的shell脚本名
	//println("正在处理shell脚本补全: " + shellProxy)
	// 检查文件是否存在
	if _, err := os.Stat(shellProxy); os.IsNotExist(err) {
		// 如果当前目录下不存在 就去可执行文件所在的文件夹里面获取
		execPath, err := os.Executable()
		if err != nil {
			return
		}
		shellProxy = filepath.Join(filepath.Dir(execPath), "mlink", "sh", command[0]+".sh")
		// 检查文件是否存在
		if _, err := os.Stat(shellProxy); os.IsNotExist(err) {
			return
		}
	}

	args := command[1:] // 获取除第一个参数外的所有参数
	// 将脚本路径加在 args 前面
	if len(args) == 0 {
		args = []string{shellProxy} // 如果没有参数，就只执行脚本
	} else {
		args = append([]string{shellProxy}, args...) // 将脚本路径作为第一个参数
	}
	// 如果文件存在，尝试执行它
	cmd := exec.Command("sh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return
	}
}

// 处理可执行程序（shell脚本）补全为系统指令 这个文件可能是带后缀名的 所以检查的时候要检查一下后缀名的文件
func doExe(command []string) {
	// 检查当前目录下是否有可执行文件
	execProxy := "./mlink/exe/" + command[0] // 指定到写死的可执行文件名
	//println("正在处理可执行补全: " + execProxy)
	// 检查文件是否存在
	if _, err := os.Stat(execProxy); os.IsNotExist(err) {
		// 如果当前目录下不存在 就去可执行文件所在的文件夹里面获取
		execPath, err := os.Executable()
		if err != nil {
			return
		}
		execProxy = filepath.Join(filepath.Dir(execPath), "mlink", "exe", command[0])
		// 检查文件是否存在
		if _, err := os.Stat(execProxy); os.IsNotExist(err) {
			return
		}
	}

	// 检查文件是否是可执行的
	fileInfo, err := os.Stat(execProxy)
	if err != nil {
		return
	}
	if !fileInfo.Mode().IsRegular() || (fileInfo.Mode()&0111) == 0 {
		fmt.Printf("❌ %s 不是一个可执行文件（可能权限不足）\n", execProxy)
		return
	}

	// 如果文件存在，尝试执行它
	cmd := exec.Command(execProxy, command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return
	}
}

// checkBusybox 检查busybox是否存在并可用
func checkBusybox() {
	fmt.Println("正在检查busybox...")

	// 首先检查busybox是否在PATH中
	path, err := exec.LookPath("busybox")
	if err != nil {
		fmt.Println("❌ busybox未找到在系统PATH中")
		fmt.Println("提示：您可以运行 'busybox_install' 来安装busybox")
		return
	}

	fmt.Printf("✅ busybox找到路径: %s\n", path)

	// 尝试执行busybox来验证它是否工作正常
	cmd := exec.Command("busybox", "--help")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("❌ busybox执行失败: %v\n", err)
		return
	}

	// 解析busybox版本信息
	lines := strings.Split(string(output), "\n")
	if len(lines) > 0 {
		fmt.Printf("✅ busybox版本: %s\n", strings.TrimSpace(lines[0]))
	}

	// 获取busybox支持的命令列表
	listCmd := exec.Command("busybox", "--list")
	listOutput, err := listCmd.Output()
	if err != nil {
		fmt.Printf("⚠️  无法获取busybox命令列表: %v\n", err)
	} else {
		commands := strings.Fields(string(listOutput))
		fmt.Printf("✅ busybox支持 %d 个命令\n", len(commands))
		fmt.Println("常用命令包括: ls, cat, grep, find, sed, awk 等")
	}
}

// installBusybox 安装busybox
func installBusybox() {
	fmt.Println("正在尝试安装busybox...")

	// 首先检查操作系统
	if runtime.GOOS != "linux" {
		fmt.Printf("❌ 当前操作系统 (%s) 不支持自动安装busybox\n", runtime.GOOS)
		fmt.Println("请手动从以下地址下载: https://busybox.net/downloads/")
		return
	}

	// 检查是否有root权限
	if os.Geteuid() != 0 {
		fmt.Println("⚠️  安装busybox需要root权限")
		fmt.Println("请使用 sudo 运行此命令")
		return
	}

	// 尝试不同的包管理器
	packageManagers := []struct {
		name    string
		check   string
		install string
	}{
		{"apt", "apt", "apt update && apt install -y busybox"},
		{"yum", "yum", "yum install -y busybox"},
		{"dnf", "dnf", "dnf install -y busybox"},
		{"pacman", "pacman", "pacman -S --noconfirm busybox"},
		{"zypper", "zypper", "zypper install -y busybox"},
		{"apk", "apk", "apk add busybox"},
	}

	var selectedPM *struct {
		name    string
		check   string
		install string
	}

	// 检查哪个包管理器可用
	for i := range packageManagers {
		pm := &packageManagers[i]
		if _, err := exec.LookPath(pm.check); err == nil {
			selectedPM = pm
			break
		}
	}

	if selectedPM == nil {
		fmt.Println("❌ 未找到支持的包管理器")
		fmt.Println("请手动安装busybox:")
		fmt.Println("1. 从 https://busybox.net/downloads/ 下载二进制文件")
		fmt.Println("2. 或者从源码编译安装")
		return
	}

	fmt.Printf("📦 使用 %s 包管理器安装busybox...\n", selectedPM.name)

	// 执行安装命令
	cmd := exec.Command("sh", "-c", selectedPM.install)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("❌ 安装失败: %v\n", err)
		fmt.Println("请尝试手动安装busybox:")
		fmt.Printf("  %s\n", selectedPM.install)
		return
	}

	fmt.Println("✅ busybox安装完成！")

	// 验证安装
	fmt.Println("正在验证安装...")
	checkBusybox()
}

// 挂载busybox指令 使用busybox --install -s /bin
func mountBusybox() {
	fmt.Println("正在挂载busybox指令...")

	// 首先检查busybox是否在PATH中
	path, err := exec.LookPath("busybox")
	if err != nil {
		fmt.Println("❌ busybox未找到在系统PATH中")
		fmt.Println("提示：您可以运行 'busybox_install' 来安装busybox")
		return
	}

	fmt.Printf("✅ busybox找到路径: %s\n", path)

	// 执行挂载命令
	cmd := exec.Command("busybox", "--install", "-s", "/bin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("❌ 挂载失败: %v\n", err)
		return
	}

	fmt.Println("✅ busybox指令挂载完成！")
}
