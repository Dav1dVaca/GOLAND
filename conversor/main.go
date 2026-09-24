package main

import (
	"fmt"

	"conversor/monedas"
)

func main() {
	var dolares float64

	fmt.Print("Ingrese la cantidad en dólares: ")
	fmt.Scanln(&dolares)

	fmt.Printf("%.2f USD = %.2f EUR\n",
		dolares,
		monedas.DolarAEuro(dolares),
	)

	fmt.Printf("%.2f USD = %.2f LB\n",
		dolares,
		monedas.DolarALibra(dolares),
	)

	fmt.Printf("%.2f USD = %.2f WON\n",
		dolares,
		monedas.DolarAWon(dolares),
	)

	fmt.Printf("%.2f USD = %.8f BTC\n",
		dolares,
		monedas.DolarABTC(dolares),
	)
}
