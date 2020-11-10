package interfaces

//数据类型是为了提交内存的利用率，因为不通的类型所需申请的内存大小不一致。
//如 bool 1个字节，int 可以根据运行系统觉得是int32 4个字节 还是int64 8个字节
type Shaper interface {
	Area() float64
}
