package furnitures

type Furniture struct {
	Type     string
	Age      int16
	Colour   string
	Material string
	Price    string
}

type FurnitureSet struct {
	FurnitureSet []Furniture
}

func AddFurnitureSet() FurnitureSet {
	var furn []Furniture
	furn = append(furn, Furniture{Type: "Stul", Age: 1, Colour: "black", Material: "plastic", Price: "999 p"})
	furn = append(furn, Furniture{Type: "krovat'", Age: 3, Colour: "black", Material: "wood", Price: "10000 p"})
	furn = append(furn, Furniture{Type: "tv", Age: 10, Colour: "Black", Material: "Plastic and glass", Price: "bezcenno"})
	furn = append(furn, Furniture{Type: "shkaf", Age: 3, Colour: "black", Material: "wood", Price: "8000 p"})
	furn = append(furn, Furniture{Type: "stol", Age: 4, Colour: "Black", Material: "wood", Price: "6000 p"})

	return FurnitureSet{FurnitureSet: furn}
}