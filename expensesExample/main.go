package main

import (
	"fmt"

	"github.com/k-stz/goscratch/expensesExample/expenses"
)

var records []expenses.Record = []expenses.Record{
	{Day: 1, Amount: 15, Category: "groceries"},
	{Day: 11, Amount: 300, Category: "utility-bills"},
	{Day: 12, Amount: 28, Category: "groceries"},
	{Day: 26, Amount: 300, Category: "university"},
	{Day: 28, Amount: 1300, Category: "rent"},
}

var period expenses.DaysPeriod = expenses.DaysPeriod{From: 1, To: 15}

// Day1Records only returns true for records that are from day 1
func Day1Records(r expenses.Record) bool {
	return r.Day == 1
}

// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"},
//   {Day: 11, Amount: 300, Category: "utility-bills"},
//   {Day: 12, Amount: 28, Category: "groceries"},
// ]

func main() {
	// filteredRecords := expenses.Filter(records, Day1Records)
	// fmt.Println("all Records:", records)
	// fmt.Println("after day1Records Filter:", filteredRecords)

	// records := expenses.Filter(records, expenses.ByDaysPeriod(period))
	// fmt.Println("after ByDayPeriod Filter", records)
	p1 := expenses.DaysPeriod{From: 1, To: 30}
	//p2 := expenses.DaysPeriod{From: 31, To: 60}

	records := []expenses.Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	res, error := expenses.CategoryExpenses(records, p1, "entertainment")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Result:", res)
}

// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"}
// ]
