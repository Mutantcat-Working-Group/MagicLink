<div align="center">
<img src="https://s2.loli.net/2025/07/16/KhZIYaePVlFEMm8.png" style="width:100px;" width="100"/>
<h2>魔链</h2>
</div>

### 功能简介

- 一款多功能Linux系统命令行补全工具
- 补全缺少的功能、指令、或直接返回某些信息
- 支持有逻辑和参数的shell脚本、支持执行程序和静态文本、工业操作简化
- 可用于软件行为欺骗、程序依赖临时补全

### 基础功能

- 可以先将二进制可执行程序放到bin目录或者环境变量所在的位置

- ![_20250123085332.png](https://s2.loli.net/2025/07/16/ea6wYMVC9glNSzk.png)

- 首先可以将想要自己补的Linux指令放在二进制文件所在文件夹下的mlink文件夹

- mlink文件夹结构大概如下，其中exe文件夹里面是可执行二进制程序，sh中存放的是shell脚本（可接收参数），static里面是没有后缀名的文本文件

- ```
  mlink
  ├── exe
  │   └── socat
  ├── sh
  │   ├── add.sh
  │   └── hello.sh
  └── static
      └── help
  ```

- 接下来，比如当我们使用magiclink hello时，系统将直接返回打印hello文件中的静态内容，这就是静态指令的使用方案（实现伪造、临时补充）

- 当我们使用magiclink socat的时候则会调起socat程序，使用magiclink add时则会调起这个shell脚本

- 除了静态文本的返回时，参数无效以外，其他操作的参数都是支持携带任意数量参数的

- 之后，我们可以使用ln -s magiclink /bin/xxx 将magiclink提供的某个指令挂载到全局并在将来可以仅用xxx这个关键字调用

- ![_20250123085332.png](https://s2.loli.net/2025/07/16/C37q2IP4dljBnks.png)

- 在这里是三个文件夹中的文件都支持挂载的，而且我们可以发现挂载后不再需要输入magiclink调用了，同时支持sudo传播

- 除此之外，当magiclink指令可全局访问了以后，我们可以在需要某些指令的局部文件夹中创建mlink文件夹，这个文件夹中定义的脚本或静态返回结果文件的优先级会更高，相当于没有当前文件夹指定的就走程序二进制文件同级位置的mlink文件夹里的，有的话先用当前文件夹中mlink指定的，从而实现动态伪装或补全

### BusyBox功能

- 为了方便更快补全指令，我们还提供了busybox的快捷功能

- magiclink busybox_check可以检查当前系统的busybox是否可以用
- magiclink busybox_install可以帮你安装busybox
- magiclink busybox_mount可以帮你用当前安装好的busybox补全的指令挂载
- 自己手动安装busybox也很方便，直接将对应平台的busybox二进制可执行程序放在/bin目录中即可
