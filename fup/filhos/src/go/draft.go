package main

import "fmt"

func main() {
	var filho_mais_novo, total_de_filhos, outro_filho int
	fmt.Scan(&filho_mais_novo, &total_de_filhos)
	for i := filho_mais_novo; i < total_de_filhos; i++ {
			fmt.Printf("%d\n", i)
		outro_filho+= 1
	}

}
