package main
import (
	"fmt"
	"strconv"
  )
func main(){
	num := 100
	say := "1"
	for i := 1; i <= 100; i++ {
		num = i
		if(num<=100){
			if(num%3 == 0){
				say = "Apa"
				if(num%5==0){
					say = "ApaBole"
				}
			}else if(num%5 == 0){
				say = "Bole"
			} else {
				say =  strconv.Itoa(num)
			}
		}
		fmt.Print(say)
		fmt.Print(", ")
	}
}