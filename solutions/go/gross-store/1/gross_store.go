package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen":3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    };
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    if unitQuantity, ok := units[unit]; ok {
        _, isInBill := bill[item];
        if isInBill{
            bill[item] += unitQuantity;
        } else{
            bill[item] = unitQuantity;
        }
        
        return true;
    }

    return false;
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billItem, inBill := bill[item]; 
    unitQuantity, inUnit := units[unit];
    balanceQuantity := billItem - unitQuantity;
    if !inBill || !inUnit || balanceQuantity < 0 {
        return false;
    }
    if balanceQuantity == 0  {
        delete(bill, item);
        return true;
    }
    bill[item] = balanceQuantity;
    
    return true;
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billItem, inBill := bill[item];
    if !inBill {
        return 0, false;
    }
    return billItem, true;
}
