package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func eliminar(s []float64, indice int) []float64 {
	return append(s[:indice], s[indice+1:]...)
}

func Pearson(persona1, persona2 []float64, elementoVacio Par) (resultado float64) {
	persona1 = eliminar(persona1, elementoVacio.posJ)
	persona2 = eliminar(persona2, elementoVacio.posJ)

	suma := 0.0
	for _, v := range persona1 {
		suma += v
	}
	promedioP1 := suma / float64(len(persona1))
	suma = 0
	for _, v := range persona2 {
		suma += v
	}
	promedioP2 := suma / float64(len(persona2))

	num := 0.0
	for i := 0; i < len(persona1); i++ {
		num += ((persona1[i] - promedioP1) * (persona2[i] - promedioP2))
	}
	aux1 := 0.0
	aux2 := 0.0
	for i := 0; i < len(persona1); i++ {
		aux1 += math.Pow(float64(persona1[i]-promedioP1), 2)
		aux2 += math.Pow(float64(persona2[i]-promedioP2), 2)
	}
	den := math.Sqrt(aux1) * math.Sqrt(aux2)
	resultado = float64(num) / den
	
	return
}

func Coseno(persona1, persona2 []float64, elementoVacio Par) (resultado float64) {
	persona1 = eliminar(persona1, elementoVacio.posJ)
	persona2 = eliminar(persona2, elementoVacio.posJ)

	num := 0.0
	for i := 0; i < len(persona1); i++ {
		num += (persona1[i] * persona2[i])
	}
	aux1 := 0.0
	aux2 := 0.0
	for i := 0; i < len(persona1); i++ {
		aux1 += math.Pow(float64(persona1[i]), 2)
		aux2 += math.Pow(float64(persona2[i]), 2)
	}
	den := math.Sqrt(aux1) * math.Sqrt(aux2)
	resultado = float64(num) / den
	
	return
}

func Euclideo(persona1, persona2 []float64, elementoVacio Par) (resultado float64) {
	persona1 = eliminar(persona1, elementoVacio.posJ)
	persona2 = eliminar(persona2, elementoVacio.posJ)

	resultado = 0.0
	for i := 0; i < len(persona1); i++ {
		resultado += math.Pow(float64(persona1[i])-float64(persona2[i]), 2)
	}
	resultado = math.Sqrt(resultado)
	
	return
}

func PrediccionSimple(matriz [][]float64, valoresMetricos []float64, vecinosSeleccionados []int, elementoVacio Par) (resultado float64) {
	num, den := 0.0, 0.0

	for i := 0; i < len(vecinosSeleccionados); i++ {
		num += valoresMetricos[vecinosSeleccionados[i]] * matriz[vecinosSeleccionados[i]][elementoVacio.posJ]
		den += math.Abs(valoresMetricos[vecinosSeleccionados[i]])
	}
	resultado = (num / den)
	return
}

func PrediccionMedia(matriz [][]float64, valoresMetricos []float64, vecinosSeleccionados []int, elementoVacio Par) (resultado float64) {
	num, den := 0.0, 0.0
	suma := 1.0
	for _, v := range matriz[elementoVacio.posI] {
		suma += v
	}
	promedioUsuario := suma / float64(len(matriz[elementoVacio.posI])-1)
	promediosUsuarios := 0.0
	for i := 0; i < len(vecinosSeleccionados); i++ {
		suma = 0.0
		for _, v := range matriz[vecinosSeleccionados[i]] {
			suma += v
		}
		promediosUsuarios = suma / float64(len(matriz[vecinosSeleccionados[i]]))
		num += valoresMetricos[vecinosSeleccionados[i]] * (matriz[vecinosSeleccionados[i]][elementoVacio.posJ] - promediosUsuarios)
		den += math.Abs(valoresMetricos[vecinosSeleccionados[i]])
	}
	resultado = promedioUsuario + (num / den)
	return
}

type Par struct {
	posI int
	posJ int
}

func main() {
	nombreArchivo := flag.String("nombre", "tabla.txt", "Nombre del archivo.")
	metrica := flag.String("metrica", "CP", "Métrica seleccionada. Las opciones son:\n1. CP (Correlación de Pearson).\n2. DC (Distancia coseno).\n3. DE (Distancia Euclídea).")
	vecinos := flag.Int("vecinos", 3, "Número de vecinos a considerar.")
	prediccion := flag.String("prediccion", "DM", "Tipo de predicción:\n1. PS (Predicción simple).\n2. DM (Diferencia con la media).")
	flag.Parse()
	fmt.Println("Nombre del archivo: ", *nombreArchivo)
	fmt.Println("Métrica seleccionada: ", *metrica)
	fmt.Println("Número de vecinos a considerar: ", *vecinos)
	fmt.Println("Tipo de predicción: ", *prediccion)

	archivo, err := os.Open(*nombreArchivo)

	if err != nil {
		fmt.Printf("Error al abrir el archivo %s: %v", *nombreArchivo, err)
		log.Fatal(err)
	}
	defer archivo.Close()

	bufer := make([]byte, 1)
	var matrizAux [][]string
	var filaAux []string
	for {
		byte, err := archivo.Read(bufer)
		caracter := string(bufer[:byte])
		if caracter == "-" {
			caracter = "-1"
		}
		if caracter != " " && caracter != "\n" {
			filaAux = append(filaAux, caracter)
		}
		if caracter == "\n" {
			matrizAux = append(matrizAux, filaAux)
			filaAux = nil
		}
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error al leer el contenido: %v", err)
				log.Fatal(err)
			}
			matrizAux = append(matrizAux, filaAux)
			break
		}
	}

	var matriz [][]float64
	var fila []float64
	var elementoVacio Par


	for i := 0; i < len(matrizAux); i++ {
		for j := 0; j < len(matrizAux[i]); j++ {
			numero, err := strconv.ParseFloat(matrizAux[i][j], 64)
			if err != nil {
				fmt.Printf("Error al convertir el número: %v", err)
				log.Fatal(err)
			}
			if numero == -1 {
				elementoVacio = Par{i, j}
			}
			fila = append(fila, numero)
		}
		matriz = append(matriz, fila)
		fila = nil
	}

		var metricaFunc func([]float64, []float64, Par) float64

		switch *metrica {
		case "CP":
			metricaFunc = Pearson
		case "DC":
			metricaFunc = Coseno
		case "DE":
			metricaFunc = Euclideo
		default:
			log.Fatalf("Métrica '%s' no reconocida.", *metrica)
		}
	
		valoresMetricos := make([]float64, len(matriz))
	
		for i := 0; i < len(matriz); i++ {
			if i != elementoVacio.posI {
				valoresMetricos[i] = metricaFunc(matriz[elementoVacio.posI], matriz[i], elementoVacio)
			} else {
				valoresMetricos[i] = -math.MaxFloat64
			}
		}
	
		vecinosSeleccionados := make([]int, 0)
	
		for i := 0; i < *vecinos; i++ {
			maxValor := -math.MaxFloat64
			maxPos := -1
	
			for j := 0; j < len(valoresMetricos); j++ {
				if valoresMetricos[j] > maxValor {
					maxValor = valoresMetricos[j]
					maxPos = j
				}
			}
	
			if maxPos == -1 {
				break
			}
	
			vecinosSeleccionados = append(vecinosSeleccionados, maxPos)
			valoresMetricos[maxPos] = -math.MaxFloat64
		}
	
		var resultado float64
	
		switch *prediccion {
		case "PS":
			resultado = PrediccionSimple(matriz, valoresMetricos, vecinosSeleccionados, elementoVacio)
		case "DM":
			resultado = PrediccionMedia(matriz, valoresMetricos, vecinosSeleccionados, elementoVacio)
		default:
			log.Fatalf("Tipo de predicción '%s' no reconocida.", *prediccion)
		}
	
		fmt.Printf("El valor predicho para el usuario %d en la posición %d es: %f\n", elementoVacio.posI+1, elementoVacio.posJ+1, resultado)
}