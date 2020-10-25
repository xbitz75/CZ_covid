package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type overviewJSON struct {
	Modified time.Time
	Source   string
	Data     []overview
}

type overview struct {
	Datum                                date
	Provedene_testy_celkem               int
	Potvrzene_pripady_celkem             int
	Aktivni_pripady                      int
	Vyleceni                             int
	Umrti                                int
	Aktualne_hospitalizovani             int
	Provedene_testy_vcerejsi_den         int
	Potvrzene_pripady_vcerejsi_den       int
	Potvrzene_pripady_dnesni_den         int
	Provedene_testy_vcerejsi_den_datum   date
	Potvrzene_pripady_vcerejsi_den_datum date
	Potvrzene_pripady_dnesni_den_datum   date
}

func (o overview) String() string {
	return fmt.Sprintf(
		"Datum: %v\nProvedene testy celkem: %v\nPotvrzene pripady celkem: %v\nAktivni pripady: %v\nVyleceni: %v\nUmrti: %v\nAktualne hospitalizovani: %v\nProvedene testy vcerejsi den: %v\nPotvrzene pripady vcerejsi den: %v\nPotvrzene pripady dnesni den: %v\nProvedene testy vcerejsi den datum: %v\nPotvrzene pripady vcerejsi den datum: %v\nPotvrzene pripady dnesni den datum: %v",
		o.Datum, o.Provedene_testy_celkem, o.Potvrzene_pripady_celkem, o.Aktivni_pripady, o.Vyleceni, o.Umrti, o.Aktualne_hospitalizovani, o.Provedene_testy_vcerejsi_den, o.Potvrzene_pripady_vcerejsi_den, o.Potvrzene_pripady_dnesni_den, o.Provedene_testy_vcerejsi_den_datum, o.Potvrzene_pripady_vcerejsi_den_datum, o.Potvrzene_pripady_dnesni_den_datum)
}

func getOverview() overview {
	data := getData("https://onemocneni-aktualne.mzcr.cz/api/v2/covid-19/zakladni-prehled.json")
	return parseOverview(data)
}

func parseOverview(data []byte) overview {
	var parsedData overviewJSON
	if err := json.Unmarshal(data, &parsedData); err != nil {
		fmt.Println(err)
	}

	return parsedData.Data[0]
}
