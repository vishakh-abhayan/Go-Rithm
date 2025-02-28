package gross

// Units stores the Gross Store unit measurements.
func Units()(grossUnit map[string]int) {
	grossUnit = map[string]int{
		"quarter_of_a_dozen" : 3,
		"half_of_a_dozen" : 6,
		"dozen" : 12,
		"small_gross" : 120,
		"gross" : 144,
		"great_gross" : 1728,
	}	
	return 
}

// NewBill creates a new bill.
func NewBill()(emptyBill map[string]int) {
	emptyBill = map[string]int{}
	return 
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if  val, ok := units[unit];	ok {
		bill[item] += val
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemQty, itemExists := bill[item]
	unitQty, unitExists := units[unit]
	if !itemExists || !unitExists || itemQty < unitQty {
		return false
	}
	if  bill[item] - units[unit] == 0 {
		delete(bill, item)
	}
	if  bill[item] - units[unit] > 0 {
		bill[item] -= units[unit]		
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, ok := bill[item];
	return val, ok
}
