package algoritmosexercicios

import "fmt"

func minimumBribes(q []int32) {
	total := 0

	for i := 0; i < len(q); i++ {
		// veririfica se uma pessoa avançou mais que duas posicoes
		// exemplo :  se q[i]=5 estiver na posicao 1=i+1 , isso nao seria possivel
		// pois 5-1 = 4 , este deveria trocar 4 vezes
		if q[i]-int32(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		start := max(0, int(q[i])-2)

		for j := start; j < i; j++ {
			
			if q[j] > q[i] {
				total++
			}
		}
	}

	fmt.Println(total)
}