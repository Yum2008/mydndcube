package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
)

func Min4(a [4]uint8) uint8 {
	minIdx := uint8(0)
	if a[1] < a[minIdx] {
		minIdx = 1
	}
	if a[2] < a[minIdx] {
		minIdx = 2
	}
	if a[3] < a[minIdx] {
		minIdx = 3
	}
	return minIdx
}

func zacherk(val []uint8, id int) string {
	s := color.New(color.CrossedOut)
	var str string
	var u uint8
	for i, v := range val {
		if i == id {
			str += s.Sprint(v)
		} else {
			u += v
			str += fmt.Sprint(v)
		}
		str += " "
	}
	str += "= " + strconv.Itoa(int(u))
	return str
}

func roll_pers() {
	r := rand.New(rand.NewSource(rand.Int63()))
	var sclad [6][4]uint8

	for i := 0; i < 24; i++ {
		sclad[i/4][i%4] = 1 + uint8(r.Intn(6))
	}
	s := uint8(0)
	var ids [6]uint8
	for i, row := range sclad {
		for _, v := range row {
			s += v
		}
		ids[i] = Min4(row)
	}
	if s < 70 {
		roll_pers()
		return
	}

	str := ""

	for i, row := range sclad {
		str += zacherk(row[:], int(ids[i])) + "   "
		if i&1 == 1 {
			fmt.Print(str + "\n")
			str = ""
		}
	}
	fmt.Print("Сумма значений: ", s, "\n")
}

func roll_k100(n int) {
	r := rand.New(rand.NewSource(rand.Int63()))
	str := ""
	for i := 0; i < n; i++ {
		str += strconv.Itoa((r.Intn(10))*10+(1+r.Intn(10))) + " "
	}
	fmt.Println(str)
}

func roll(t, n int) {
	k := 0

	switch t {
	case 1:
		k = 20
	case 2:
		k = 12
	case 3:
		k = 10
	case 4:
		k = 8
	case 5:
		k = 6
	case 6:
		k = 4
	case 7:
		roll_k100(n)
		return
	}
	r := rand.New(rand.NewSource(rand.Int63()))

	str := ""
	all := 0; m := 100; mm := 0
	for i := 0; i < n; i++ {
		q := 1 + r.Intn(k)
		all += q
		str += strconv.Itoa(q) + " "
		if q<m{m=q}
		if q>mm{mm=q}
	}
	fmt.Println("Бросок К"+ strconv.Itoa(k),"х"+strconv.Itoa(n))
	fmt.Println(str, "Сумма:", all)
	fmt.Println("Минимальное:", m, "Максимальное:",mm)
}

func pause() (bool) {
	var agree bool
	huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Еще бросок?").
				Affirmative("Да").
				Negative("Нет").
				Value(&agree),
		),
	).Run()
	return agree
}

func main() {
	for {
		var i int
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().
					Title("Выбери бросок").
					Options(	
						huh.NewOption("Нароллить статы персу", 0),
						huh.NewOption("Бросок к20", 1),
						huh.NewOption("Бросок к12", 2),
						huh.NewOption("Бросок к10", 3),
						huh.NewOption("Бросок к8", 4),
						huh.NewOption("Бросок к6", 5),
						huh.NewOption("Бросок к4", 6),
						huh.NewOption("Бросок к100", 7),
					).
					Value(&i),
			),
		).Run()
		if err != nil {
			return
		}

		if i == 0 {
			roll_pers()

		} else {
			var ns string
			err := huh.NewForm(
				huh.NewGroup(
				huh.NewInput().
				Title("Сколько бросков совершить?").
				Validate(func(input string) error {
					_, err := strconv.Atoi(input)
					if err != nil {
						return err
					}
					return nil
				}).
				Value(&ns),
			),
			).Run()
			if err != nil {
				fmt.Printf("Ты чета не то ввел")
				return
			}
			n, _ := strconv.Atoi(ns)
			roll(i, n)
		}

		again := pause()
		if !again {
			break
		}
	}

	fmt.Print("\n\nНажмите Enter для выхода...")
	fmt.Scanln()
}