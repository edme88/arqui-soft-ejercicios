package operaciones

import "testing"

func TestSuma(t *testing.T){
	tabla := []struct{
		a int
		b int
		c int
	}{
		{1,2,3},
		{2,3,5},
		{25,12,37},
	}

	for _, item := range tabla{
		total := Suma(item.a, item.b)

		if total != item.c{
			t.Errorf("Suma incorrecta. Se esperaba %d, total %d", item.c, total)
		}
	}
	// total := Suma(2,5)

	// if total != 7 {
	// 	t.Errorf("Suma incorrecta. Se esperaba %d, total %d", 7, total)
	// }
}

func TestGetMax(t *testing.T){
	tabla := []struct{
		a int
		b int
		c int
	}{
		{1,2,2},
		{2,3,3},
		{45,12,45},
	}

	for _, item := range tabla{
		total := GetMax(item.a, item.b)

		if total != item.c{
			t.Errorf("Maximo incorrecto. Se esperaba %d, total %d", item.c, total)
		}
	}
	// total := Suma(2,5)

	// if total != 7 {
	// 	t.Errorf("Suma incorrecta. Se esperaba %d, total %d", 7, total)
	// }
}