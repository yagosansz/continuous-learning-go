package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitsToScore := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return unitsToScore
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	emptyBill := make(map[string]int)
	return emptyBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	} else {
		itemQuantity, itemExists := GetItem(bill, item)
		if itemExists {
			bill[item] = itemQuantity + unitValue
		} else {
			bill[item] = unitValue
		}
		return true
	}

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	itemQuantity, itemExists := GetItem(bill, item)
	if !unitExists || !itemExists {
		return false
	} else {
		newQuantity := itemQuantity - unitValue
		switch {
		case newQuantity < 0:
			return false
		case newQuantity == 0:
			delete(bill, item)
			return true
		default:
			bill[item] = newQuantity
			return true
		}
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemQuantity, itemExists := bill[item]
	return itemQuantity, itemExists
}
