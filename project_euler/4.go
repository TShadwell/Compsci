package main
import(
	"strconv"
	"fmt"
)
/*

The highest number possible to get from two natural 3 digit numbers would be
999*999 = 998, 001
The lowest would be
100*100 = 10000

Meaning any palindrome expressed as three digit numbers, x
must be
10000<p<998,001

Since palindromes must be mirrors of themselves, the plaindrome must have 6
digits, since we can neither exceed to 8 digits, nor descend to 4

So, our search area is effectively
abc,bca

100 to 999
*/
func reverse (str string)string{
	runes:=[]rune(str)
	for forward, back :=0, len(runes)-1; forward < back; forward, back = forward+1, back-1{
		runes[forward], runes[back] = runes[back], runes[forward]
	}
	return string(runes)
}

func isPalindrome(sVal string)bool{
		if len(sVal)==6 && sVal[:3] == reverse(sVal[3:]){
			return true
		}
		return false
}
func main(){
	checked := make(map[uint32]bool)
	palindr:=uint32(100001)
	for i:=uint32(999);i>0;i--{
		if checked[i] ==true{
			continue
		} else{
			for q:=uint32(999);q>0;q--{
				if checked[q] == true{
					continue
				} else{
					value:=i*q
					sVal:=strconv.Itoa(int(value))
					if isPalindrome(sVal) && value > palindr{
						palindr = value
					}
				}
			}
			checked[i] =true
		}
	}
	fmt.Println(palindr)
}
