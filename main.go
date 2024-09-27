package main

import "fmt"

func main() {
	fmt.Println("Please fill in the Number from the list bellow: ")

	fmt.Println("1. Celcius To Farenheit")
	fmt.Println("2. Celcius To Kelvin")
	fmt.Println("3. Farenheit To Celcius")
	fmt.Println("4. Farenheit To Kelvin")
	fmt.Println("5. Kelvin To Celcius")
	fmt.Println("6. Kelvin To Farenheit")

	println("list number: ")

	var list int
	fmt.Scanf("%d", &list)

	for list < 1 || list > 6 {
		fmt.Println("Your choice not available, please fill with right choice: ")
		fmt.Scanf("%d", &list)
	}

	var inputSuhu float64
	fmt.Println("fill temp.: ")
	fmt.Scanf("%f", &inputSuhu)

	var outputSuhu float64
	switch list {
	case 1:
		outputSuhu = CelciusToFahrenheit(inputSuhu)
	case 2:
		outputSuhu = CelciusToKelvin(inputSuhu)
	case 3:
		outputSuhu = FarenheitToCelcius(inputSuhu)
	case 4:
		outputSuhu = FarenheitToKelvin(inputSuhu)
	case 5:
		outputSuhu = KelvinToCelcius(inputSuhu)
	case 6:
		outputSuhu = KelvinToFarenheit(inputSuhu)
	default:
		fmt.Println("your choice not valid")
		return
	}

	fmt.Printf("Temp. is: %.2f\n", outputSuhu)
}

func CelciusToFahrenheit(celcius float64) float64 {
	farenheit := (9.0 / 5.0 * celcius) + 32

	return farenheit
}

func CelciusToKelvin(celcius float64) float64 {
	kelvin := celcius + 273.15

	return kelvin
}

func FarenheitToCelcius(farenheit float64) float64 {
	celcius := (farenheit - 32) * (5.0 / 9.0)

	return celcius
}

func FarenheitToKelvin(farenheit float64) float64 {
	kelvin := (farenheit + 459.67) * (5.0 / 9.0)

	return kelvin
}

func KelvinToCelcius(kelvin float64) float64 {
	celcius := kelvin - 273.15

	return celcius
}

func KelvinToFarenheit(kelvin float64) float64 {
	farenheit := (kelvin * 9.0 / 5.0) - 459.67

	return farenheit
}
