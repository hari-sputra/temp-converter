package main

import "fmt"

type Celcius struct {
	suhu float64
}
type Farenheit struct {
	suhu float64
}
type Kelvin struct {
	suhu float64
}

func (c Celcius) toCelcius() float64 {
	return c.suhu
}

func (c Celcius) toFarenheit() float64 {
	return (9.0 / 5.0 * c.suhu) + 32
}

func (c Celcius) toKelvin() float64 {
	return c.suhu + 273.15
}

func (f Farenheit) toFarenheit() float64 {
	return f.suhu
}

func (f Farenheit) toCelcius() float64 {
	return (f.suhu - 32) * (5.0 / 9.0)
}

func (f Farenheit) toKelvin() float64 {
	return (f.suhu + 459.67) * (5.0 / 9.0)
}

func (k Kelvin) toKelvin() float64 {
	return k.suhu
}

func (k Kelvin) toCelcius() float64 {
	return k.suhu - 273.15
}

func (k Kelvin) toFarenheit() float64 {
	return (k.suhu * 9.0 / 5.0) - 459.67
}

type tempCalc interface {
	toCelcius() float64
	toFarenheit() float64
	toKelvin() float64
}

func main() {
	fmt.Println("Please fill in the Number from the first temp. list bellow: ")

	fmt.Println("1. Celcius")
	fmt.Println("2. Farenheit")
	fmt.Println("3. Kelvin ")

	println("first list number: ")

	var first_list int
	fmt.Scanf("%d", &first_list)

	for first_list < 1 || first_list > 6 {
		fmt.Println("Your choice not available, please fill with right choice: ")
		fmt.Scanf("%d", &first_list)
	}

	fmt.Println("Please fill in the Number from the Last temp. list bellow: ")

	fmt.Println("1. Celcius")
	fmt.Println("2. Farenheit")
	fmt.Println("3. Kelvin ")

	println("last list number: ")

	var last_list int
	fmt.Scanf("%d", &last_list)

	for last_list < 1 || last_list > 6 {
		fmt.Println("Your choice not available, please fill with right choice: ")
		fmt.Scanf("%d", &last_list)
	}

	var temp float64
	fmt.Println("fill temp: ")
	fmt.Scanf("%f", &temp)

	var interfaceTemp tempCalc
	switch first_list {
	case 1:
		interfaceTemp = Celcius{temp}
	case 2:
		interfaceTemp = Farenheit{temp}
	case 3:
		interfaceTemp = Kelvin{temp}
	default:
		fmt.Println("your choice not valid")
		return
	}

	var finalTemp float64
	switch last_list {
	case 1:
		finalTemp = interfaceTemp.toCelcius()
	case 2:
		finalTemp = interfaceTemp.toFarenheit()
	case 3:
		finalTemp = interfaceTemp.toKelvin()
	default:
		fmt.Println("your choice not valid")
		return
	}

	fmt.Printf("Temp. is: %.2f\n", finalTemp)
}
