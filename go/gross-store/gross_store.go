package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	//panic("Please implement the Units() function")
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	//panic("Please implement the NewBill() function")
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	//panic("Please implement the AddItem() function")
	if unitValue, ok := units[unit]; ok {
		bill[item] = bill[item] + unitValue
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	//panic("Please implement the RemoveItem() function")
	if unitValue, ok := units[unit]; ok {
		if billValue, ok := bill[item]; ok && billValue-unitValue >= 0 {
			if newBillValue := billValue - unitValue; newBillValue != 0 {
				bill[item] = newBillValue
			} else {
				delete(bill, item)
			}
			return true
		}
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
