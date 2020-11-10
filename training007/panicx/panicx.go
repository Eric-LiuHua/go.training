
import (
    "errors"
    "fmt"
    "time"
)

//panic     导致关键流程出现不可修复性错误的情况
//error     其他情况使用
//recover      捕获panic 的异常信息
func TestPanic1() {

    defer func() {
        fmt.Printf("正常捕获异常是在defer的匿名函数中，recover():%v \n", recover())
    }()

    defer func() {
        func() {
            fmt.Printf("间接调用defer 捕获不到异常，recover():%v \n", recover())
        }()
    }()

    defer fmt.Printf("相当直接调用，捕获不到异常，recover():%v \n", recover())

    //error 的用法，只是
    panic(errors.New("error  TestPanic1 异常1"))
}

func TestPanic2() {

    ThrowsPanic(genErr)
}
func genErr() {
    fmt.Printf("%v 正常语句!\n", time.Now())
    defer func() {
        fmt.Printf("%v defer 正常语句 !\n", time.Now())
        panic("第二个错误 recover 只能看到最后一个panic ")
    }()
    panic("第一个错误，不影响panic 之后的defer 执行")
}
//用于捕获f对异常
func ThrowsPanic(f func()) (b bool) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("%v, 捕获到异常, %v \n", time.Now(), r)
            b = true
        }
    }()
    f()
    return
}