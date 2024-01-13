package gadgets

type Gadgets struct {
	Type   string
	Age    int16
	Model  string
	Colour string
	Price  string
}

type Gadget struct {
	Gadget []Gadgets
}

func AddGadgets() Gadget {
	var gadgets []Gadgets
	gadgets = append(gadgets, Gadgets{Type: "Telephone", Age: 1, Colour: "Black", Model: "Xiaomi", Price: "2999 p"})
	gadgets = append(gadgets, Gadgets{Type: "Comp", Age: 1, Colour: "Black", Model: "Xiaomi", Price: "22999 p"})
	gadgets = append(gadgets, Gadgets{Type: "tv", Age: 1, Colour: "Black", Model: "Xiaomi", Price: "20999 p"})

	return Gadget{Gadget: gadgets}
}