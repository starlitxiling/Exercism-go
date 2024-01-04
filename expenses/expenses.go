package expenses

import "fmt"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	// panic("Please implement the Filter function")
	var result []Record
	for _, j := range in {
		if predicate(j) {
			result = append(result, j)
		}
	}
	return result
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	// panic("Please implement the ByDaysPeriod function")
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	// panic("Please implement the ByCategory function")
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	// panic("Please implement the TotalByPeriod function")
	result := Filter(in, ByDaysPeriod(p))
	var total float64
	for _, r := range result {
		total += r.Amount
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	// panic("Please implement the CategoryExpenses function")
	categoryExists := false
	for _, r := range in {
		if r.Category == c {
			categoryExists = true
			break
		}
	}

	if !categoryExists {
		return 0, fmt.Errorf("unknown category %s", c)
	}

	filteredByCategory := Filter(in, ByCategory(c))
	filteredRecords := Filter(filteredByCategory, ByDaysPeriod(p))

	var total float64
	total = 0
	for _, r := range filteredRecords {
		total += r.Amount
	}
	return total, nil
}
