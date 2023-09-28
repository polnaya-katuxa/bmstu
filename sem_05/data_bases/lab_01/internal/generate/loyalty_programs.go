package generate

type LoyaltyProgram struct {
	ID          int     `csv:"id"`
	Name        string  `csv:"name"`
	Design      string  `csv:"string"`
	Percent     int     `csv:"cashback_percent"`
	MinPurchase float32 `csv:"minimum_purchase_sum"`
}

func LoaltyPrograms() []LoyaltyProgram {
	return []LoyaltyProgram{
		{
			ID:          1,
			Name:        "Welcome",
			Design:      "computers.club/images/welcome.png",
			Percent:     0,
			MinPurchase: 0,
		},
		{
			ID:          2,
			Name:        "Beginner",
			Design:      "computers.club/images/beginner.png",
			Percent:     5,
			MinPurchase: 3000,
		},
		{
			ID:          3,
			Name:        "Pilot",
			Design:      "computers.club/images/pilot.png",
			Percent:     7,
			MinPurchase: 6000,
		},
		{
			ID:          4,
			Name:        "Silver",
			Design:      "computers.club/images/silver.png",
			Percent:     10,
			MinPurchase: 10000,
		},
		{
			ID:          5,
			Name:        "Gold",
			Design:      "computers.club/images/gold.png",
			Percent:     15,
			MinPurchase: 20000,
		},
		{
			ID:          6,
			Name:        "Platinum",
			Design:      "computers.club/images/platinum.png",
			Percent:     25,
			MinPurchase: 50000,
		},
	}
}
