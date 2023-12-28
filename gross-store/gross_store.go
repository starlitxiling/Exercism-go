package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	// panic("Please implement the Units() function")
	foo := map[string]int{}
	foo["quarter_of_a_dozen"] = 3
	foo["half_of_a_dozen"] = 6
	foo["dozen"] = 12
	foo["small_gross"] = 120
	foo["gross"] = 144
	foo["great_gross"] = 1728

	return foo
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	// panic("Please implement the NewBill() function")
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the AddItem() function")
	_, exists := units[unit]
	if !exists {
		return false
	}
	_, inBill := bill[item]
	if inBill {
		bill[item] += units[unit]
	} else {
		bill[item] = units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the RemoveItem() function")
	currencyQty, itemExists := bill[item]
	if !itemExists {
		return false
	}
	unitQty, unitExists := units[unit]
	if !unitExists {
		return false
	}
	newQty := currencyQty - unitQty
	if newQty < 0 {
		return false
	}
	if newQty == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQty
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	// panic("Please implement the GetItem() function")
	qty, exists := bill[item]
	if !exists {
		return 0, false
	}
	return qty, true
}
