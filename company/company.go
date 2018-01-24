package company

type Company struct {
	Name    string
	Website string
}

func New(name, website string) Company {
	return Company{Name: name, Website: website}
}

func (c Company) IsCool() bool {
	if c.Name == "TheLions" {
		return true
	}

	return false
}
