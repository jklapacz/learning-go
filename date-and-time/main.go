package main
import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Println(current.String())
	fmt.Println("MM-DD-YYYY : ", current.Format("01-02-2006"))
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", current.Format("2006-01-02 15:04:05"))
	// add or subtract dates
	nextMonthDate := current.AddDate(0, 0, 24)
	fmt.Println("MM-DD-YYYY : ", nextMonthDate.Format("01-02-2006"))

	// parse date and time from strings
	str := "2018-03-12T11:45:26.371Z"
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t, err)

}
