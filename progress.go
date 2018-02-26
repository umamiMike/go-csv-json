package main

import (
	"gopkg.in/cheggaaa/pb.v2"
	"time"
)

func simpleProgress() {
	count := 1000
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond * 2)
	}
	bar.Finish()
}
