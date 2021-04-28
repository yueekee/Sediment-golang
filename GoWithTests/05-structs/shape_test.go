package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

/* 
使用Shape的interface，调用通过类型继承调用Shpe下不同的函数。
重构：表驱动测试
*/

func TestArea(t *testing.T) {
	// checkArea := func(t *testing.T, shape Shape, want float64){
	// 	t.Helper()
	// 	got := shape.Area()

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// }

// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		checkArea(t, rectangle, 100.0)
// 	})

// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

	areaTests := []struct{
		name string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, v := range areaTests {
		got := v.shape.Area()
		if got != v.want {
			t.Errorf("got %.2f want %.2f", got, v.want)
		}
	}
}
