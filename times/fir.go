package times

import (
	"fmt"
	"time"
)

func MyTime() {

	t1 := time.Now()
	fmt.Println("Now time is ", t1)
	fmt.Printf("Type is %T\n", t1)

	ts := t1.Format("2006-01-02 15:04:05")
	fmt.Println("ts = ", ts)
	fmt.Printf("Type is %T\n", ts)

}
