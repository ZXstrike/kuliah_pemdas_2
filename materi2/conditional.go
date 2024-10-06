package materi2

import (
	"fmt"
)

func IfElse() {
	var a = 10
	var b = 5

	if a > b {
		fmt.Println("a lebih besar dari b")
	} else if a < b {
		fmt.Println("a lebih kecil dari b")
	} else {
		fmt.Println("a sama dengan b")
	}
}

func SwitchCase() {
	var nilai = 90

	switch nilai {
	case 90:
		fmt.Println("Nilai anda A")
	case 80:
		fmt.Println("Nilai anda B")
	case 70:
		fmt.Println("Nilai anda C")
	case 60:
		fmt.Println("Nilai anda D")
	default:
		fmt.Println("Nilai anda E")
	}
}

func SwitchCaseFallthrough() {
	var nilai = 90

	switch nilai {
	case 90:
		fmt.Println("Nilai anda A")
		fallthrough
	case 80:
		fmt.Println("Nilai anda B")
		fallthrough
	case 70:
		fmt.Println("Nilai anda C")
		fallthrough
	case 60:
		fmt.Println("Nilai anda D")
	default:
		fmt.Println("Nilai anda E")
	}
}

func SwitchCaseWithExpression() {
	var nilai = 90

	switch {
	case nilai > 90:
		fmt.Println("Nilai anda A")
	case nilai > 80:
		fmt.Println("Nilai anda B")
	case nilai > 70:
		fmt.Println("Nilai anda C")
	case nilai > 60:
		fmt.Println("Nilai anda D")
	default:
		fmt.Println("Nilai anda E")
	}
}

func SwitchCaseWithMultipleExpression() {
	var nilai = 90

	switch {
	case nilai > 90:
		fmt.Println("Nilai anda A")
	case nilai > 80:
		fmt.Println("Nilai anda B")
	case nilai > 70, nilai == 70:
		fmt.Println("Nilai anda C")
	case nilai > 60:
		fmt.Println("Nilai anda D")
	default:
		fmt.Println("Nilai anda E")
	}
}

func NestedConditional() {
	var a = 10
	var b = 5

	if a == 10 {
		if b == 5 {
			fmt.Println("a adalah 10 dan b adalah 5")
		} else {
			fmt.Println("a adalah 10 tapi b bukan 5")
		}
	} else {
		fmt.Println("a bukan 10")
	}
}

func ConditionalAssignment() {
	var a = 10
	var b = 5

	var hasil = a + b

	if hasil == 15 {
		fmt.Println("Hasil dari a + b adalah 15")
	} else {
		fmt.Println("Hasil dari a + b bukan 15")
	}
}
