package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// DataPoint representa un punto de datos con latitud, longitud, tipo de delito y fecha/hora.
type DataPoint struct {
	Latitude  float64 `json:"latitud"`
	Longitude float64 `json:"longitud"`
	CrimeType string  `json:"tipo_delito"`
	DateTime  string  `json:"fecha_hora"`
}

// KMeansResult representa el resultado de la agrupación k-means.
type KMeansResult struct {
	Centroids []DataPoint `json:"centroids"`
}

// KMeans algoritmo de agrupación
func KMeans(data []DataPoint, k int) KMeansResult {
	centroids := make([]DataPoint, k)
	for i := 0; i < k; i++ {
		centroids[i] = data[rand.Intn(len(data))]
	}
	return KMeansResult{Centroids: centroids}
}

func loadDataFromURL(url string) ([]DataPoint, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []DataPoint
	reader := csv.NewReader(resp.Body)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		lat, _ := strconv.ParseFloat(record[0], 64)
		lon, _ := strconv.ParseFloat(record[1], 64)
		data = append(data, DataPoint{
			Latitude:  lat,
			Longitude: lon,
			CrimeType: record[2],
			DateTime:  record[3],
		})
	}
	return data, nil
}

func handleCluster(w http.ResponseWriter, r *http.Request) {
	data, err := loadDataFromURL("https://raw.githubusercontent.com/cesar6793/Final_concurrente/main/delitos.csv")
	if err != nil {
		http.Error(w, "Failed to load data", http.StatusInternalServerError)
		return
	}

	result := KMeans(data, 4)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/cluster", handleCluster)
	log.Println("API server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
