package analyze

import (
	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/object"
	"time"
)

type Analyze struct {
	vv         api.VideoSearchResponse
	yearVideos map[int][]object.VideoVideo
}

func NewAnalyze(vv api.VideoSearchResponse) *Analyze {
	return &Analyze{
		vv:         vv,
		yearVideos: map[int][]object.VideoVideo{},
	}
}

func (a Analyze) SplitByYears() error {
	for _, i := range a.vv.Items {
		year := time.Unix(int64(i.Date), 0).
			Year()
		a.yearVideos[year] = append(a.yearVideos[year], i)
	}
	return nil
}

type StatYears struct {
	Likes    map[int]int
	Views    map[int]int
	Comments map[int]int
	Counts   map[int]int
}

func NewStatYears(analyze Analyze) *StatYears {
	sy := &StatYears{
		Likes:    map[int]int{},
		Views:    map[int]int{},
		Comments: map[int]int{},
		Counts:   map[int]int{},
	}
	sy.countingLikes(analyze)
	sy.countingViews(analyze)
	sy.countingComments(analyze)
	sy.counting(analyze)
	return sy
}

func (s StatYears) countingLikes(analyze Analyze) {

	for year, videos := range analyze.yearVideos {
		for _, video := range videos {
			s.Likes[year] += video.Likes.Count
		}
	}
}

func (s StatYears) countingViews(analyze Analyze) {

	for year, videos := range analyze.yearVideos {
		for _, video := range videos {
			s.Views[year] += video.Views
		}
	}
}
func (s StatYears) countingComments(analyze Analyze) {

	for year, videos := range analyze.yearVideos {
		for _, video := range videos {
			s.Comments[year] += video.Comments
		}
	}
}

func (s StatYears) counting(analyze Analyze) {
	for year, videos := range analyze.yearVideos {
		s.Counts[year] = len(videos)
	}
}

type BIMetrics struct {
	Activity map[int]int
	Passion  map[int]int
}

func NewBIMetrics(st StatYears) *BIMetrics {
	bi := &BIMetrics{
		Activity: make(map[int]int),
		Passion:  make(map[int]int),
	}
	bi.countingPassion(st)
	bi.countingActivity(st)
	return bi
}

func (b BIMetrics) countingActivity(stat StatYears) {
	for year, _ := range stat.Counts {
		b.Activity[year] = (stat.Likes[year] + stat.Comments[year]) / stat.Counts[year]
	}
}

func (b BIMetrics) countingPassion(stat StatYears) {
	for year, _ := range stat.Counts {
		b.Passion[year] = stat.Views[year] / stat.Counts[year]
	}
}
