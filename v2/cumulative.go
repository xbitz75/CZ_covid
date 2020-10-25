package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const graphDir = ".\\graphs"

type cumulativeJSON struct {
	Modified time.Time
	Source   string
	Data     []cumulative
}

type cumulative struct {
	Datum                        date
	Kumulativni_pocet_nakazenych int
	Kumulativni_pocet_vylecenych int
	Kumulativni_pocet_umrti      int
	Kumulativni_pocet_testu      int
}

func (c cumulative) String() string {
	return fmt.Sprintf(
		"Datum: %v\nKumulativni pocet nakazenych: %v\nKumulativni pocet vylecenych: %v\nKumulativni pocet umrti: %v\nKumulativni pocet testu: %v\n",
		c.Datum, c.Kumulativni_pocet_nakazenych, c.Kumulativni_pocet_vylecenych, c.Kumulativni_pocet_umrti, c.Kumulativni_pocet_testu)
}

func getCumulativeCollection() cumulativeColection {
	data := getData("https://onemocneni-aktualne.mzcr.cz/api/v2/covid-19/nakazeni-vyleceni-umrti-testy.json")
	return parseCumulative(data)
}

func parseCumulative(data []byte) cumulativeColection {
	var parsedData cumulativeJSON
	if err := json.Unmarshal(data, &parsedData); err != nil {
		fmt.Println(err)
	}

	return cumulativeColection{data: parsedData.Data}
}

type cumulativeColection struct {
	data []cumulative
}

func (c *cumulativeColection) plotDeaths() {
	var deaths []int
	for _, d := range c.data {
		deaths = append(deaths, d.Kumulativni_pocet_umrti)
	}

	c.plot("Covid umrti", "Pocet umrti", deaths, ".\\cumulativeDeaths.png")
}

func (c *cumulativeColection) plotTested() {
	var tested []int
	for _, t := range c.data {
		tested = append(tested, t.Kumulativni_pocet_testu)
	}

	c.plot("Covid testovani", "Pocet testovanych", tested, ".\\cumulativeTested.png")
}

func (c *cumulativeColection) plotRecovered() {
	var recovered []int
	for _, r := range c.data {
		recovered = append(recovered, r.Kumulativni_pocet_vylecenych)
	}

	c.plot("Covid vyleceni", "Pocet vylecenych", recovered, ".\\cumulativeRecovered.png")
}

func (c *cumulativeColection) plotInfected() {
	var infected []int
	for _, i := range c.data {
		infected = append(infected, i.Kumulativni_pocet_nakazenych)
	}

	c.plot("Covid nakazeni", "Pocet nakazenych", infected, ".\\cumulativeInfected.png")
}

func (c *cumulativeColection) plot(title string, yLabel string, inputData []int, filename string) {
	xticks := plot.TimeTicks{Format: "2006-01-02"}

	generateDeathsData := func() plotter.XYs {
		const (
			hour = 1
			min  = 1
			sec  = 1
			nsec = 1
		)
		pts := make(plotter.XYs, len(c.data))
		for i := range pts {
			currentDataDate := c.data[i].Datum
			date := time.Date(currentDataDate.year, time.Month(currentDataDate.month), currentDataDate.day, hour, min, sec, nsec, time.UTC).Unix()
			pts[i].X = float64(date)
			pts[i].Y = float64(inputData[i])
		}
		return pts
	}

	data := generateDeathsData()

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = title
	p.X.Tick.Marker = xticks
	p.Y.Label.Text = yLabel
	p.Add(plotter.NewGrid())

	line, err := plotter.NewLine(data)
	if err != nil {
		log.Panic(err)
	}
	line.Color = color.RGBA{R: 255, A: 255}

	p.Add(line)

	filepath := fmt.Sprintf("%s\\%s", graphDir, filename)
	if _, err := os.Stat(graphDir); os.IsNotExist(err) {
		os.Mkdir(graphDir, 0755)
	}

	err = p.Save(10*vg.Centimeter, 5*vg.Centimeter, filepath)
	if err != nil {
		log.Panic(err)
	}
}
