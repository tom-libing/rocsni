package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

var RedBall []int
var BlueBall []int

func init(){
	RedBall = make([]int,0,33)
	for v:=1;v<=33;v++{
		RedBall = append(RedBall,v)
	}
	BlueBall = make([]int,0,16)
	for v:=1;v<=16;v++{
		BlueBall = append(BlueBall,v)
	}
}

func main(){
	red1,bucket1 := lottery(RedBall,32)
	fmt.Println(len(bucket1))
	fmt.Println("红色球:"+strconv.Itoa(red1))
	red2,bucket2 := lottery(bucket1,31)
	fmt.Println(len(bucket2))
	fmt.Println("红色球:"+strconv.Itoa(red2))
	red3,bucket3 := lottery(bucket2,30)
	fmt.Println(len(bucket3))
	fmt.Println("红色球:"+strconv.Itoa(red3))
	red4,bucket4 := lottery(bucket3,29)
	fmt.Println(len(bucket4))
	fmt.Println("红色球:"+strconv.Itoa(red4))
	red5,bucket5 := lottery(bucket4,28)
	fmt.Println(len(bucket5))
	fmt.Println("红色球:"+strconv.Itoa(red5))
	red6,bucket6 := lottery(bucket5,27)
	fmt.Println(len(bucket6))
	fmt.Println("红色球:"+strconv.Itoa(red6))
	blue1,bucket7 := lottery(BlueBall,15)
	fmt.Println(len(bucket7))
	fmt.Println("蓝色球:"+strconv.Itoa(blue1))
}

func lottery (Ball []int,keynum int) (int,[]int){
	unix := time.Now().Unix()
	bucket := make([]int,0)
	for{
		key := rand.Intn(keynum)
		if time.Now().After(time.Unix(unix+10,0)){
			for k,v:=range Ball {
				if k!=key {
					bucket = append(bucket, v)
				}
			}
			return Ball[key],bucket
		}
	}
}
