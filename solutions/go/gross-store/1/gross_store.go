package gross
import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    pos := map[string]int{}

    pos["quarter_of_a_dozen"] = 3
    pos["half_of_a_dozen"] = 6
    pos["dozen"] = 12
    pos["small_gross"] = 120
    pos["gross"] = 144
    pos["great_gross"] = 1728
    
    return pos
    
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    _, unit_exists := units[unit]
    if (unit_exists == false){
        return unit_exists
    }
	_, item_exists := bill[item]
    if (item_exists == false){
        bill[item] = units[unit]
        return true
    }
    bill[item] += units[unit]
    return true
    
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Debug
    fmt.Println(bill)

    
    // Item not found in bill
	_, item_exists := bill[item]
    if (item_exists == false){
        return item_exists
    }

    // Unit not found in Units
    _, unit_exists := units[unit]
    if (unit_exists == false){
        return unit_exists
    }

    // Item qtty < 0
    if (bill[item] - units[unit] < 0) {
        return false
    }

    // Item qtty == 0
    if (bill[item]  - units[unit] == 0) {
        delete(bill, item)
        fmt.Println(bill)
        return true
    }

    // Removal is valid
    bill[item] -= units[unit]
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, item_exists := bill[item]

    if(value > 0 && item_exists == true) {
        return bill[item], true
    }

    return value, item_exists
}
