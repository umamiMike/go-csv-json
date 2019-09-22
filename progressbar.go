package main

import "gopkg.in/cheggaaa/pb.v2"

func incrementBar(bar *pb.ProgressBar, c chan int) {
	for {
		inc := <-c
		if inc == 1 {
			bar.Increment()
		}
	}
}
