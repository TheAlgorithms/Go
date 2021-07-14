package abstractfactorry

import (
	"reflect"
	"testing"
)

// test client code and how  to use the factory

func TestAbstractFactory(t *testing.T) {

	t.Run("Test Nike Factory", func(t *testing.T) {
		nikeFactory, _ := getSportsFactory("nike")
		nikeShoe := nikeFactory.makeShoe()
		nikeShirt := nikeFactory.makeShirt()
		if !reflect.DeepEqual(nikeShoe.getLogo(), "nike") {
			t.Errorf("got: %v, want: %v", nikeShoe.getLogo(), "nike")
		}
		if !reflect.DeepEqual(nikeShoe.getSize(), 14) {
			t.Errorf("got: %v, want: %v", nikeShoe.getSize(), 14)
		}
		if !reflect.DeepEqual(nikeShirt.getLogo(), "nike") {
			t.Errorf("got: %v, want: %v", nikeShirt.getLogo(), "nike")
		}
		if !reflect.DeepEqual(nikeShirt.getSize(), 14) {
			t.Errorf("got: %v, want: %v", nikeShirt.getLogo(), 14)
		}
	})

	t.Run("Test Adidas Factory", func(t *testing.T) {
		adidasFactory, _ := getSportsFactory("adidas")
		adidasShoe := adidasFactory.makeShoe()
		adidasShirt := adidasFactory.makeShirt()

		if !reflect.DeepEqual(adidasShoe.getLogo(), "adidas") {
			t.Errorf("got: %v, want: %v", adidasShoe.getLogo(), "adidas")
		}
		if !reflect.DeepEqual(adidasShoe.getSize(), 14) {
			t.Errorf("got: %v, want: %v", adidasShoe.getSize(), 14)
		}
		if !reflect.DeepEqual(adidasShirt.getLogo(), "adidas") {
			t.Errorf("got: %v, want: %v", adidasShirt.getLogo(), "adidas")
		}
		if !reflect.DeepEqual(adidasShirt.getSize(), 14) {
			t.Errorf("got: %v, want: %v", adidasShirt.getLogo(), 14)
		}
	})

}



