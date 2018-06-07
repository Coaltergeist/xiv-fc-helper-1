package structs

// MBListing represents a listing on the marketboard
type MBListing struct {
	IsListed bool
	Sale     struct {
		Price    string
		Quantity string
		Total    string
	}
	HQSale struct {
		Price    string
		Quantity string
		Total    string
	}
	LastSoldFor struct {
		Price    string
		Quantity string
		Total    string
	}
	DateSold string
	Buyer    string
}
