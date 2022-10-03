# Golang 面试准备
> 包含经典代码实现，阅读资料以及其他资源的汇总

![License](https://img.shields.io/github/license/songquanpeng/go-interview?color=brightgreen)
[![Test](https://github.com/songquanpeng/go-interview/actions/workflows/test.yml/badge.svg)](https://github.com/songquanpeng/go-interview/actions/workflows/test.yml)
![Report](https://goreportcard.com/badge/github.com/songquanpeng/go-interview)

## 基本语法
### 关键字
+ [ ] var / const / type
+ [ ] struct
+ [ ] interface
+ [ ] if / else
+ [ ] for / continue / break
+ [ ] switch / case / default / fallthrough
+ [ ] func / return / defer
+ [ ] map
+ [ ] range
+ [ ] go
+ [ ] select / chan
+ [ ] goto
+ [ ] import / package

### 预定义标识符
#### 类型
+ [ ] int / int8 / int32 / int64
+ [ ] uint / uint8 / uint32 / uint64 / uintptr
+ [ ] float32 / float64
+ [ ] rune / string
+ [ ] bool
+ [ ] any
+ [ ] byte
+ [ ] complex64 / complex128
+ [ ] error
+ [ ] comparable

#### 常量
+ [ ] true / false
+ [ ] iota

#### 零值
+ [ ] nil

#### 内置函数
+ [ ] append / cap / len
+ [ ] make / new / copy / delete
+ [ ] panic / recover
+ [ ] close
+ [ ] complex / real / imag
+ [ ] print / println

### 其他
+ [ ] unsafe.Pointer

## 数据结构
+ [ ] 数组
+ [ ] 切片
+ [ ] 字典
+ [ ] 集合
+ [ ] 字符串
+ [ ] 结构体

## 代码实现
### 标准库
+ [ ] sync
  + [ ] sync.Map
  + [ ] sync.Locker
  + [ ] sync.Mutex
  + [ ] sync.Cond
  + [ ] sync.Once
  + [ ] sync.Pool
  + [ ] sync.RWMutex
  + [ ] sync.WaitGroup

### 常见错误
+ [ ] 循环 & 闭包

### 并发编程
+ [ ] 协程池
  + [x] [固定大小，无参任务函数，不能等待所有任务结束](./concurrent/pool/pool.go)
  + [x] [固定大小，无参任务函数，**可以**等待所有任务结束](./concurrent/pool2/pool.go)
  + [ ] **动态**大小，**有参**任务函数，可以等待所有任务结束

## 阅读资料
+ [ ] [《Go 入门指南》](https://github.com/unknwon/the-way-to-go_ZH_CN)
+ [ ] [《Go 语言高级编程》](https://chai2010.cn/advanced-go-programming-book/)
+ [ ] [《Go 语言设计与实现》](https://draveness.me/golang/)
+ [ ] [《Go 语言高性能编程》](https://geektutu.com/post/high-performance-go.html)
+ [ ] https://www.bookstack.cn/explore?cid=10&tab=popular
+ [ ] https://github.com/lifei6671/interview-go
+ [ ] https://github.com/xiaobaiTech/golangFamily
+ [ ] https://github.com/iswbm/golang-interview
+ [ ] https://github.com/cnymw/GolangStudy
+ [ ] https://github.com/yqchilde/Golang-Interview
+ [ ] https://github.com/menggggggg/go-interview
+ [ ] https://github.com/pibigstar/go-demo/tree/master/interview
+ [ ] https://www.yuque.com/devhg/golang
+ [ ] https://geektutu.com/post/qa-golang.html
+ [ ] [《Golang 入门到精通教程》](https://geekr.dev/golang-tutorial)

## 参考
+ [Golang：循环下的闭包](https://studygolang.com/articles/14696)
+ [一文告诉你神奇的Go内建函数源码在哪里](https://tonybai.com/2020/12/17/where-is-the-source-of-builtin-functions/)
+ [Golang Builtin Documentation](https://pkg.go.dev/builtin)