package company

import "testing"

func TestCompanyIsCool(t *testing.T) {
	c := New("TheLions", "https://www.thelions.com")

	if !c.IsCool() {
		t.Error("The Company TheLions is definitely cool.. something must be wrong")
	}
}
func TestCompanyIsNotCool(t *testing.T) {
	c := New("SomeOtherCompany", "https://something.com")

	if c.IsCool() {
		t.Error("The Company SomeOtherCompany isn't cool")
	}
}
