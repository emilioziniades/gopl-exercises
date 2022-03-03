package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
	})
	MonthOld, LessYearOld, MoreYearOld := false, false, false
	today := time.Now()
	MonthAgo := today.AddDate(0, -1, 0)
	YearAgo := today.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		if !MonthOld {
			fmt.Println("\nLESS THAN A MONTH OLD")
			MonthOld = true
		}
		if item.CreatedAt.Before(MonthAgo) && !LessYearOld {
			fmt.Println("\nLESS THAN A YEAR OLD")
			LessYearOld = true
		}
		if item.CreatedAt.Before(YearAgo) && !MoreYearOld {
			fmt.Println("\nMORE THAN A YEAR OLD")
			MoreYearOld = true
		}
		fmt.Printf("#%-5d %v %9.9s %.55s\n",
			item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
}
