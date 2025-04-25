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
		total:= Suma(item.a, item.b)

		if total != item.c{
			t.Errorf("Suma incorrecta. Tiene %d se esperaba %d", total, item.c)
		}
	}

	// total := Suma(2,5)

	// if total != 7{
	// 	t.Errorf("Suma incorrecto. Tiene %d se esperaba %d", total, 7)
	// }
}

func TestGetMax(t *testing.T){
	tabla := []struct{
		a int
		b int
		c int
	}{
		{1,2,2},
	}

	for _, item := range tabla{
		total:= GetMax(item.a, item.b)

		if total != item.c{
			t.Errorf("Mayor incorrecto. Devolvio %d se esperaba %d", total, item.c)
		}
	}
}