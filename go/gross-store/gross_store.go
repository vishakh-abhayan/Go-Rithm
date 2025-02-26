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
	if  _, ok := bill[item];!ok {
		return false
	}
	if  _, ok := units[unit];!ok {
		return false
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
