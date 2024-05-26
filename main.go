package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
)

type TiledMap struct {
	Layers []TiledLayer `json:"layers"`
}

type TiledLayer struct {
	Data []int `json:"data"`
}

type Map struct {
	Layers []Layer `json:"layers"`
}

type Layer struct {
	Data [][]int `json:"data"`
}

func main() {
	data, err := os.ReadFile("./maps/full-map-example.json")
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier JSON : %v", err)
	}

	// Unmarshal
	var tiledMap TiledMap
	err = json.Unmarshal(data, &tiledMap)
	if err != nil {
		log.Fatalf("Erreur lors du parsing du JSON : %v", err)
	}

	layers := ProcessMap(tiledMap)

	content := Map{}

	for _, layerData := range layers {
		layer := Layer{
			Data: layerData,
		}
		content.Layers = append(content.Layers, layer)
	}

	// Marshal
	jsonData, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de l'encodage en JSON:", err)
		return
	}

	file, err := os.Create("slices.json")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
		return
	}

	fmt.Println("Données JSON écrites dans slices.json avec succès.")
}

func ProcessMap(tiledMap TiledMap) (layers [][][]int) {
	layersAmount := len(tiledMap.Layers)

	for i := 0; i < layersAmount; i++ {
		var buffer [][]int

		layer := tiledMap.Layers[i]
		border := int(math.Sqrt(float64(len(layer.Data))))

		if border <= 0 {
			return
		}

		fmt.Println("Layer", i, ":", border, "x", border)

		for i := 0; i < border; i++ {
			slice := layer.Data[i*border : (i+1)*border]
			buffer = append(buffer, slice)
		}

		layers = append(layers, buffer)
	}

	return layers
}
