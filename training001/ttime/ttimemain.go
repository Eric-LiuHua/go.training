package ttime

import(
	"fmt"
	"time"
	"math/rand"
)


func TestTime()  {
	now := time.Now()
	fmt.Println("time.Now():",now)
	
	year := now.Year()
	month :=now.Month()
	day :=now.Day()

	hour:= now.Hour()
	minute:=now.Minute()
	second:=now.Second()

	fmt.Printf("year:%v,month:%v,day:%v,hour:%v,minute:%v,second:%v,\n",year,month,day,hour,minute,second)

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minute,second)


	//1970 到现在到秒数 ，毫秒数
	fmt.Printf("TimeStamp:%d \n",now.Unix())

	fmt.Printf("ToTime:%v \n",ToTime(1593917280))
}

//时间戳转time
func ToTime(TimeStamp int64) time.Time {
	return time.Unix(TimeStamp,0)
}
//横杠格式化
func ToGString(now time.Time) string  {
	return	fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d\n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
}
//斜杠格式化
func ToxXString(now time.Time) string  {
	return	fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d\n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
}

//定时器学习
func TestTicker() chan bool{
	//每五秒
	ticker:=time.NewTicker(5*time.Second)
	stopChan :=make(chan bool)
	go func(ticker *time.Ticker) {
		//方法压栈，下面for 终止之后就倒序执行。
		defer ticker.Stop()
		for{
			select {
			case tmp:= <- ticker.C:
				fmt.Println("ticker running at :",ToGString(tmp))
			case stop:= <-stopChan:
				if stop{
					fmt.Println("ticker stop at :",ToGString(time.Now()))
					return
				}
			}
		}
	}(ticker)

	return stopChan
}

//time 部分常量
func TestConst()  {
	fmt.Printf("time.Nanosecond :%d \n",time.Nanosecond)
	fmt.Printf("time.Microsecond :%d \n",time.Microsecond)
	fmt.Printf("time.Millisecond :%d \n",time.Millisecond)
	fmt.Printf("time.second :%d \n",time.Second)
}

//日期格式化很特殊，
//目标很特殊，必须是 2006-01-02 15(3):04:05 ,填错就格式化不一样来
func TestFormat()  {
	fmt.Println("............... TestFormat start ...............")
	now:=time.Now()
	fmt.Printf("time.Format(now) :%s \n",now.Format("2006-01-02 15:04:05"))
	fmt.Printf("格式化练习1:time.Format(now) :%s \n",now.Format("2006/01/02 15:04:05"))
	fmt.Printf("格式化练习2:time.Format(now) :%s \n",ToxXString(now))
}


//练习耗时
func TestCost()  {
	fmt.Println("............... TestCost start ...............")
	start :=time.Now().UnixNano()
	//设置种子，不然每次都会随机成0
	rand.Seed(time.Now().Unix())
	times:= 5
	for i := 0; i < times; i++ {
		time.Sleep(time.Millisecond* time.Duration(rand.Intn(times)))
	}
	end	:=time.Now().UnixNano()

	fmt.Printf("............... TestCost end :%d us\n",((end-start)/1000))
}
