package covidfetcher

import (
	"backend-covid-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jasonlvhit/gocron"
)

var CovidStore *models.CovidStore

func fetchCovidDataGouvFr(datetime string) []models.CovidReport {
	covidReport := []models.CovidReport{}
	url := fmt.Sprintf("https://dashboard.covid19.data.gouv.fr/data/date-%s.json", datetime)

	resp, err := http.Get(url)
	if err != nil {
		return covidReport
	}

	// il se peut que les informations ne soient pas mises à jour toutes les 24h
	if resp.StatusCode == 404 {
		// 119 départements
		for i := 0; i < 119; i++ {
			cr := models.CovidReport{
				Date: datetime,
			}
			covidReport = append(covidReport, cr)
		}
		return covidReport
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return covidReport
	}
	if err = json.Unmarshal(body, &covidReport); err != nil {
		log.Fatalln(err)
	}
	return covidReport
}

func fetchLastSevenDaysCovidData() chan []models.CovidReport {
	ch := make(chan []models.CovidReport, 7)
	go func() {
		for i := -24; i >= -168; i = i - 24 {
			currentTime := time.Now().Add(time.Duration(i) * time.Hour).Format("2006-01-02")
			data := fetchCovidDataGouvFr(currentTime)
			ch <- data
		}
		close(ch)
	}()
	return ch
}

func refreshData(covStore *models.CovidStore) {
	oneDayBefore := time.Now().Add(time.Duration(-24) * time.Hour).Format("2006-01-02")
	data := fetchCovidDataGouvFr(oneDayBefore)

	lastCurrentTime := time.Now().Add(time.Duration(-192) * time.Hour).Format("2006-01-02")
	covStore.Lock()
	defer covStore.Unlock()

	if _, exists := covStore.LastSevenDataDays[lastCurrentTime]; !exists {
		return
	}
	delete(covStore.LastSevenDataDays, lastCurrentTime)
	datetime := data[0].Date
	covStore.LastSevenDataDays[datetime] = data
}

func CovidDataBackup() {
	covStore := models.CovidStore{}
	covStore.LastSevenDataDays = make(map[string][]models.CovidReport, 7)
	datas := fetchLastSevenDaysCovidData()
	for v := range datas {
		datetime := v[0].Date
		covStore.LastSevenDataDays[datetime] = v
	}
	CovidStore = &covStore

	go func() {
		if err := gocron.Every(1).Day().At("6:00").Do(refreshData, &covStore); err != nil {
			log.Println("error on job refreshing")
		}
		<-gocron.Start()
	}()
}

