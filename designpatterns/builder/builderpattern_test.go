package builder
import (
	"reflect"
	"testing"
)

func TestBuilderPatter(t *testing.T){
	t.Run("Test Normal Builder", func(t *testing.T) {
		normalBuilder := getBuilder("normal")
		director := newDirector(normalBuilder)
		normalHouse := director.buildHouse()
		if !reflect.DeepEqual(normalHouse.doorType, "Wooden Door") {
			t.Errorf("got: %v, want: %v", normalHouse.doorType, "Wooden Door")
		}
		if !reflect.DeepEqual(normalHouse.windowType, "Wooden Window") {
			t.Errorf("got: %v, want: %v", normalHouse.windowType, "Wooden Window")
		}
		if !reflect.DeepEqual(normalHouse.floor, 2) {
			t.Errorf("got: %v, want: %v", normalHouse.floor, 2)
		}
	})
	t.Run("Test Igloo Builder", func(t *testing.T) {
		iglooBuilder := getBuilder("igloo")
		director := newDirector(iglooBuilder)
		iglooHouse := director.buildHouse()
		if !reflect.DeepEqual(iglooHouse.doorType, "Snow Door") {
			t.Errorf("got: %v, want: %v", iglooHouse.doorType, "Snow Door")
		}
		if !reflect.DeepEqual(iglooHouse.windowType, "Snow Window") {
			t.Errorf("got: %v, want: %v", iglooHouse.windowType, "Snow Window")
		}
		if !reflect.DeepEqual(iglooHouse.floor, 1) {
			t.Errorf("got: %v, want: %v", iglooHouse.floor, 1)
		}
	})

}


