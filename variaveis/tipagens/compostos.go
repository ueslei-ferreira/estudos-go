package variaveis

import "fmt"

func Notas() {
	alunos := [3]string{"ueslei", "fulano", "ciclano"}

	notasUeslei := []float64{9.5, 8.7, 9.9}
	notasFulano := []float64{8.5, 8.7, 8.9}
	notasCiclano := []float64{7.5, 7.7, 7.9}

	var MediaUeslei float64
	var MediaFulano float64
	var MediaCiclano float64

	for i := 0; i < len(alunos); i++ {
		MediaUeslei += notasUeslei[i]
		MediaFulano += notasFulano[i]
		MediaCiclano += notasCiclano[i]
	}
	MediaUeslei = MediaUeslei / float64(len(notasUeslei))
	MediaFulano = MediaFulano / float64(len(notasFulano))
	MediaCiclano = MediaCiclano / float64(len(notasCiclano))

	medias := map[string]float64{
		alunos[0]: MediaUeslei,
		alunos[1]: MediaFulano,
		alunos[2]: MediaCiclano,
	}

	for i := 0; i < len(alunos); i++ {
		fmt.Printf("%s: %.2f \n", alunos[i], medias[alunos[i]])
	}

}
