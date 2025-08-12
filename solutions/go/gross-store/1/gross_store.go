package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen":	3,
        "half_of_a_dozen":	6,
        "dozen":	12,
        "small_gross":	120,
        "gross":	144,
        "great_gross":	1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, isExist := units[unit]
	if !isExist {
		return false
	}

	if billValue, isExist := bill[item]; !isExist {
		bill[item] = unitValue
	} else {
		bill[item] = billValue + unitValue
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, isExist := units[unit]
	if !isExist {
		return false
	}

    billValue, isExist := bill[item]
	if !isExist {
        return false
    }

    newQuantity := billValue - unitValue
	if newQuantity < 0 {
    	return false
    }
    
    if newQuantity == 0 {
    	delete(bill, item)
    } else {
        bill[item] = newQuantity
    }

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    q, isExist := bill[item]
	return q, isExist
}
