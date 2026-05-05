package main

import "fmt"

func main() {
	var filho_mais_novo, total_de_filhos int
	fmt.Scan(&filho_mais_novo, &total_de_filhos)
	for i := 0; i < total_de_filhos; i++ {
			fmt.Printf("%d\n", filho_mais_novo)
		filho_mais_novo+=2
	}

}
