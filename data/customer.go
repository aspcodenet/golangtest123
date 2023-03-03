package data

type Customer struct {
	Name string
	Age  int
}

func GetAllCustomers() []Customer {
	//1. Det finns arrayer i GO
	//2. En array är STATISK
	// var c = new List<HockeyPlayer>
	// c.Append(new HockeyPlayer)
	// var allCustomer = [5]Customer{}
	// allCustomer[0] = Customer{Name: "Stefan", Age: 50}
	// allCustomer[1] = Customer{Name: "Kerstin", Age: 49}
	// 5

	// slice
	var allCustomerSlice = []Customer{}
	//var allCustomerSlice = make([]Customer, 0, 50)
	//allCustomerSlice.hjkasf
	allCustomerSlice = append(allCustomerSlice, Customer{Name: "Stefan", Age: 50})
	allCustomerSlice = append(allCustomerSlice, Customer{Name: "Kerstin", Age: 49})
	allCustomerSlice = append(allCustomerSlice, Customer{Name: "Oliver", Age: 14})
	return allCustomerSlice
}
