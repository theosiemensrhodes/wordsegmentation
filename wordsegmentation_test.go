package wordsegmentation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	c "github.com/theosiemensrhodes/wordsegmentation/corpus"
)

func TestSegmentor_Segment(t *testing.T) {
	type args struct {
		Text string
	}
	englishCorpus := c.NewEnglishCorpus()
	segmentor := NewSegmentor(englishCorpus, 12)

	tests := []struct {
		name string
		args args
		want []string
	}{
		// {"basic", args{"WhatIsTheWeatherliketoday?"}, []string{"what", "is", "the", "weather", "like", "today"}},
		// {"basic with spaces", args{"click me next"}, []string{"click", "me", "next"}},
		{"complex", args{"THESOUNDOFLAUGHTERECHOINGTHROUGHACROWDEDPARKONASUNNYAFTERNOONISONEOFLIFESSIMPLEPLEASURESFAMILIESGATHERFORPICNICSFRIENDSTOSSFRISBEESANDCHILDRENRUNFREELYTHROUGHTHEGRASSTHEIRGIGGLESFILLINGTHEAIRTHESUNSHINESDOWNCASTINGLONGSHADOWSASTHEDAYBEGINSTOWINDDOWNANDTHESCENTOFFRESHLYCUTGRASSMINGLESWITH"}, []string{"the", "sound", "of", "laughter", "echoing", "through", "a", "crowded", "park", "on", "a", "sunny", "afternoon", "is", "one", "of", "life", "s", "simple", "pleasures", "families", "gather", "for", "picnics", "friends", "toss", "frisbees", "and", "children", "run", "freely", "through", "the", "grass", "their", "giggles", "filling", "the", "air", "the", "sun", "shines", "down", "casting", "long", "shadows", "as", "the", "day", "begins", "to", "wind", "down", "and", "the", "scent", "of", "freshly", "cut", "grass", "mingles", "with"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := segmentor.Segment(tt.args.Text)
			assert.Equalf(t, tt.want, got, fmt.Sprintf("Segmentor.Segment() = %v, want %v\n", got, tt.want))
		})
	}
}
