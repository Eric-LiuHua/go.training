
package example

import(
	"fmt"
)

//是否质数
func IsPrime(n int) bool  {
	if n<=1 {
		return false
	}else{
		for i := 2; i < n; i++ {
			if n % i ==0{
				return false 
			}
		}
		return true
	}
}

//练习，0 到n 有多少个质数
func AllPrimeNumbers(n int)  {
	for i := 0; i < n; i++ {
		if IsPrime(i){
			fmt.Printf("[%d] is prime. \n",i)
		}
	}	

}

//判断是否水仙花。百位立方+十位立方+个位立方 等于原数
func IsSXH(n int) bool  {
	if n<100 ||n >=1000{
		return false 
	}

	first :=n%10
	second := (n/10)%10
	three := n/100


	return ((first* first*first) +(second *second*second)+(three*three*three)) == n

}


//统计各类字符出现的次数
func CharCount(s string) (num int,en int,space int, other int)  {
	
	//因为字符串是utf8，需要rune 转切片
	utfchars := []rune(s)

	for i := 0; i < len(utfchars); i++ {
		if(utfchars[i]>='a'&&utfchars[i]<='z')||(utfchars[i]>='A'&&utfchars[i]<='Z'){
			en++
			continue
		}else if(utfchars[i]>='0'&&utfchars[i]<='9'){
			num++
			continue
		}else if(utfchars[i]==' '){
			space++
			continue
		}
		other++
	}

	return
}