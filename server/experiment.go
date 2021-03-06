package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// experiments holds the list of all experiments
var experiments ExperimentSlice

const maxSaves = 2

var saveDir string
var savingInterval time.Duration
var serverPort string

var saveQueue = make(chan string, maxSaves+1)

func setupConfig() {
	viper.SetConfigName("./config/config") // name of config file (without extension)
	viper.SetConfigType("yaml")            // REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	saveDir = viper.GetString("server.saveDirectory")
	savingInterval = (time.Duration(viper.GetInt64("server.snapShotInterVal")) * time.Second)
	serverPort = viper.GetString("server.Port")

	log.Println("Loaded the confing from the config file")
}

func periodicSave() {
	ticker := time.NewTicker(1 * savingInterval)
	for range ticker.C {
		log.Println("Attempting to snapshot the current state...")
		saveExperimentsToFile()
		log.Println("Managed to save the state... Resuming...")
	}
}

func handleExperimentSubmission(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("Serving a GET request...")
		expID := r.URL.Query().Get("id")
		log.Println("\tDetected query parameter id = ", expID)
		exp, _, err := findExperimentByID(&experiments, expID)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			resJSON, err := json.Marshal(*exp)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.Write(resJSON)
		}
	case "POST":
		log.Println("Serving a POST request...")

		// retrieve the message body
		var exp Experiment

		err := json.NewDecoder(r.Body).Decode(&exp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		exp.ID = uuid.New().String()
		exp.Date = time.Now().String()
		exp.Metrics = make(map[string][]float32)
		experiments.Experiments = append(experiments.Experiments, exp)
		experiments.ExperimentCount++
		w.Write([]byte("Added an experiment to the experiment list\n"))

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}

func handleDeleteExperiment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("Serving a DELETE request...")
		expID := r.URL.Query().Get("id")
		log.Println("\tDetected query parameter id = ", expID)
		err := deleteExperiment(&experiments, expID)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Fprintf(w, "Deleted the experiment %s", expID)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func experimentsToJSON() ([]byte, error) {
	experimentJSON, err := json.Marshal(experiments)
	return experimentJSON, err
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func loadExperimentsFromFile() {
	// fileList is already sorted
	exist, _ := exists(saveDir)
	if !exist {
		os.MkdirAll(saveDir, 0777)
	}
	fileList, err := ioutil.ReadDir(saveDir)
	if err != nil {
		log.Panic(err)
	}

	if len(fileList) == 0 {
		// no files found
		experiments = ExperimentSlice{
			ExperimentCount: 0,
		}
		return
	}
	// if doesn't work, try also the ModTime
	// snapshot is the last entry
	latestSnap := fileList[len(fileList)-1]

	file, err := ioutil.ReadFile(saveDir + "/" + latestSnap.Name())
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal([]byte(file), &experiments)
	if err != nil {
		log.Panic(err)
	}
}

func saveExperimentsToFile() {
	t := time.Now().Format("20060102150405")
	fileData, err := json.MarshalIndent(experiments, "", " ")
	if err != nil {
		log.Panic(err)
	}
	filename := saveDir + "/" + "snapshot-" + t + ".json"
	err = ioutil.WriteFile(filename, fileData, 0644)
	if err != nil {
		log.Panic(err)
	}
	saveQueue <- filename
	if len(saveQueue) >= maxSaves {
		// cut the unnecessary saves
		pop := <-saveQueue
		err := os.Remove(pop)
		log.Println("Removed old snapshot", pop)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func handleListExperiments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("Serving a GET ALL request...")

		expJ, err := experimentsToJSON()
		if err != nil {
			log.Panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(expJ)
	}
}

func main() {
	setupConfig()
	initServer()
	// connectMongo()

	// loadExperimentsFromFile()
	// go periodicSave()

	// mux := http.NewServeMux()
	// mux.HandleFunc("/experiment", handleExperimentSubmission)
	// mux.HandleFunc("/experiments", handleListExperiments)
	// mux.HandleFunc("/delete", handleDeleteExperiment)

	// log.Println("Starting server on :4000...")
	// err := http.ListenAndServe(":"+serverPort, mux)
	// log.Fatal(err)
}
