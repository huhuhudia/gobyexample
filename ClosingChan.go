package main

import "fmt"

func main(){
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func(){
		for{
			i, more := <- jobs
			if more {
				fmt.Println("received job", i)
			}else{
				fmt.Println("recieved all jobs")
				done <- true
				return
			}
		}
	}()

	for j:=1; j<=3; j++{
		jobs<-j
		fmt.Println("send job",j)
	}

	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}