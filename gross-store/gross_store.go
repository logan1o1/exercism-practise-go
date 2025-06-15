package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	units["half_of_a_dozen"] = 6
	units["quarter_of_a_dozen"] = 3

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitVal, exists := units[unit]
	if exists {
		bill[item] += unitVal
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitVal, uintExists := units[unit]
	itemVal, itemExists := bill[item]

	if !uintExists || !itemExists {
		return false
	}

	if newQuant := itemVal - unitVal; newQuant == 0 {
		delete(bill, item)
	} else if newQuant < 0 {
		return false
	} else {
		bill[item] = newQuant
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemval, itemExist := bill[item]
	if !itemExist {
		return 0, false
	}
	return itemval, true
}
