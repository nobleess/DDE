package flow

import (
	"DDE/analyze"
	"DDE/orm"
	"DDE/scrape"
	"fmt"
	api "github.com/SevereCloud/vksdk/v3/api"
	"os"

	"log"
)

func Flow(proms []string) {
	token := os.Getenv("TOKEN")
	client := api.NewVK(token)
	if client == nil {
		log.Fatalln("Failed to connect to VK")
	}
	vs := scrape.NewVScrape(client, api.Params{})

	for _, prom := range proms {

		prom := orm.Prom{
			Promm: prom,
		}
		_ = vs.Prom(prom.Promm)
		//vs.Sort(0)

		db, err := orm.NewDB("dashboard/sources/localdb/localdb")
		if err != nil {
			log.Fatalln(err)
		}
		db.Migration()
		err = db.Conn.Create(&prom).Error
		if err != nil {
			fmt.Println(err)
			continue
		} // save prom

		err = vs.Scrape()
		if err != nil {
			log.Fatalln(err)
		}

		analyz := analyze.NewAnalyze(vs.GetVideo())
		err = analyz.SplitByYears()
		if err != nil {
			log.Fatalln(err)
		}
		statYear := analyze.NewStatYears(*analyz)
		biMetrics := analyze.NewBIMetrics(*statYear)

		//
		//// save statyaer
		l, v, c, count := orm.StatYearsAdapter(statYear, prom.ID)

		if err := db.Conn.Create(l).Error; err != nil {
			log.Fatalln(err)
		}

		if err := db.Conn.Create(v).Error; err != nil {
			log.Fatalln(err)
		}
		if err := db.Conn.Create(c).Error; err != nil {
			log.Fatalln(err)
		}
		if err := db.Conn.Create(count).Error; err != nil {
			log.Fatalln(err)
		}

		activity, passion := orm.BIAdapter(*biMetrics, prom.ID)
		db.Conn.Create(activity)
		db.Conn.Create(passion)

	}
}
