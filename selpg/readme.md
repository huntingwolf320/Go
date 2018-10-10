## 设计说明
代码使用pflag包，注意测试命令行规范。

测试使用样例说明：

hello文件夹是个会输出200行只含有c:[行号]的测试程序

data为200行只含有l:[行号],且以‘\n’分行的测试数据文件

fdata为200行只含有c:[行号]，且以'\f'分页的测试数据文件


## 使用测试结果
按照使用selpg章节的测试命令
(除了第一条测试，其他命令设置页大小--l 5节省输出)
1.$ selpg --s 1 --e 1 data
``` 
$ selpg --s 1 --e 1 data
l:1
l:2
l:3
l:4
l:5
l:6
l:7
l:8
l:9
l:10
l:11
l:12
l:13
l:14
l:15
l:16
l:17
l:18
l:19
l:20
l:21
l:22
l:23
l:24
l:25
l:26
l:27
l:28
l:29
l:30
l:31
l:32
l:33
l:34
l:35
l:36
l:37
l:38
l:39
l:40
l:41
l:42
l:43
l:44
l:45
l:46
l:47
l:48
l:49
l:50
l:51
l:52
l:53
l:54
l:55
l:56
l:57
l:58
l:59
l:60
l:61
l:62
l:63
l:64
l:65
l:66
l:67
l:68
l:69
l:70
l:71
l:72
``` 
2.$ selpg --s 1 --e 1 < data
```
$ selpg --s 1 --e 1 --l 5 < data
l:1
l:2
l:3
l:4
l:5
```
3.$ other_command | selpg -s10 -e20
```
$ hello | selpg --s 5 --e 6 --l 5
c:21
c:22
c:23
c:24
c:25
c:26
c:27
c:28
c:29
c:30
```
4.$ selpg -s10 -e20 input_file >output_file
```
$ selpg --s 1 --e 2 --l 5 data >out
//out文件内容
l:1
l:2
l:3
l:4
l:5
l:6
l:7
l:8
l:9
l:10
```
5.$ selpg -s10 -e20 input_file 2>error_file
```
$ selpg --s 2 --e 1 --l 5 data 2>err
屏幕无输出
//err文件内容
The number of start page should be greater than or equal to the end page's!
```
6.$ selpg -s10 -e20 input_file >output_file 2>error_file
2>error_file
```
$ selpg --s 39 --e 41 --l 5 data >out 2>err
屏幕无输出
//out文件内容
l:191
l:192
l:193
l:194
l:195
l:196
l:197
l:198
l:199
//err文件内容
The range of page is illegal!
```
7.$ selpg -s10 -e20 input_file >output_file 2>/dev/null
```
$ selpg --s 39 --e 41 --l 5 data >out 2>/dev/null
屏幕无输出
//out文件内容
l:191
l:192
l:193
l:194
l:195
l:196
l:197
l:198
l:199
//dev/null文件内容为空
```
8.$ selpg -s10 -e20 input_file >/dev/null
```
$ selpg --s 39 --e 41 --l 5 data >/dev/null
The range of page is illegal!
//dev/null文件内容为空
```
9.$ selpg -s10 -e20 input_file | other_command
```
$ selpg --s 1 --e 2 --l 5 data | wc
 10      10      41
```
10.$ selpg -s10 -e20 input_file 2>error_file | other_command
```
$ selpg --s 39 --e 41 --l 5 data 2>err | wc
 9      9      54
 //err文件内容
 The range of page is illegal!
```
11.$ selpg -s10 -e20 -l66 input_file

此测试上面做过同类的不再重复

12.$ selpg -s10 -e20 -f input_file
```
$ selpg --s 1 --e 5 --f fdata
c:1
c:2
c:3
c:4
c:5
```
13.$ selpg -s10 -e20 -dlp1 input_file
由于没有打印机，命令执行出现错误，但已经显示利用管道执行了lp命令
```
$ selpg --s 1 --e 2 --d lp1 data
/usr/bin/lp: The printer or class does not exist.
```
