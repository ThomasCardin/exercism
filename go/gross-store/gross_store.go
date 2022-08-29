package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, found := units[unit]
	if found {
		_, found := bill[item]
		if found {
			bill[item] += value
			return true
		}

		bill[item] = value
		return true
	}

	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billValue, foundInBill := bill[item]
	unitsValue, foundInUnits := units[unit]
	if foundInBill && foundInUnits {
		switch {
		case billValue-unitsValue < 0:
			return false
		case billValue-unitsValue == 0:
			delete(bill, item)
			return true
		default:
			bill[item] -= unitsValue
			return true
		}
	}

	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, found := bill[item]
	if !found {
		return 0, false
	} else {
		return value, true
	}
}
