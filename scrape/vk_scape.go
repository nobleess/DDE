package scrape

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v3/api"
	"sync"
)

/*

api.Params {
	sort:0,

}

*/

type VScrape struct {
	videos api.VideoSearchResponse
	client *api.VK
	values api.Params
}

func NewVScrape(client *api.VK, values api.Params) VScrape {
	return VScrape{
		client: client,
		values: values,
	}
}

func (vs *VScrape) Scrape() (err error) {
	// only 500 first https://dev.vk.com/ru/method/video.search
	//buff := make([]api.VideoSearchResponse, 0, 50)
	const limit = 20
	wg := sync.WaitGroup{}
	mu := &sync.Mutex{}
	for offset := 0; offset < 100; offset += limit {
		wg.Add(1)
		go func() {
			defer wg.Done()
			params := api.Params{}
			params["q"] = vs.values["q"]
			params["offset"] = offset
			params["count"] = limit
			var vv api.VideoSearchResponse
			vv, err = vs.client.VideoSearch(vs.values)
			if err != nil {
				return
			}
			mu.Lock()
			vs.videos.Items = append(vs.videos.Items, vv.Items...)
			mu.Unlock()
		}()
	}
	wg.Wait()
	//len(vs.videos.Items))
	return nil
}

func (vs VScrape) Prom(query string) error {
	if vs.values == nil {
		return fmt.Errorf("no values set")
	}
	vs.values["q"] = query // set prom
	return nil
}

func (vs VScrape) Sort(sort uint) error {
	if vs.values == nil {
		return fmt.Errorf("no values set")
	}
	if sort > 4 {
		return fmt.Errorf("sort must be <= 4")
	}
	vs.values["sort"] = sort // set sort
	return nil
}

func (vs VScrape) GetVideo() api.VideoSearchResponse {
	return vs.videos
}
