package members

type Member struct {
	Name  string
	Title string
	Age   int16
	Sex   string
}

type Fam struct {
	Family []Member
}

func AddFamily() Fam {
	var family []Member

	family = append(family, Member{Name: "Anna", Title: "Mother", Age: 47, Sex: "Female"})
	family = append(family, Member{Name: "Yuriy", Title: "Father", Age: 55, Sex: "Male"})
	family = append(family, Member{Name: "Biba", Title: "Brother", Age: 15, Sex: "male"})
	family = append(family, Member{Name: "Boba", Title: "Brother", Age: 14, Sex: "male"})
	return Fam{Family: family}
}