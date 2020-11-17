package models

import "sync"

type CovidReport struct {
	Deces                  int    `json:"deces"`
	Reanimation            int    `json:"reanimation"`
	Hospitalises           int    `json:"hospitalises"`
	Gueris                 int    `json:"gueris"`
	Date                   string `json:"date"`
	Code                   string `json:"code"`
	Nom                    string `json:"nom"`
	TauxOccupationReaColor string `json:"tauxOccupationReaColor"`
}

type CovidStore struct {
	LastSevenDataDays map[string][]CovidReport
	sync.Mutex
}
