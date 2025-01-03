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
	segmentor := NewSegmentor(englishCorpus)

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"basic", args{"WhatIsTheWeatherliketoday"}, []string{"what", "is", "the", "weather", "like", "today"}},
		{"tricky 1", args{"friendstossfrisbees"}, []string{"friends", "toss", "frisbees"}},
		{"tricky 2", args{"crispairof"}, []string{"crisp", "air", "of"}},
		{"tricky 3", args{"theresachill"}, []string{"theres", "a", "chill"}},
		{"complex 1", args{"thecrispairofafallmorningcombinedwiththesightofvibrantorangeandredleavesmakesforoneofthemostinvigoratingstartstoadaythegroundbeneathyourfeetiscoveredinasoftcarpetoffallenleavesandeachstepbringsthesatisfyingcrunchofautumnunderfootasthecoolbreezecarrieswithitthescentofwoodsmokeandapplesthere"}, []string{"the", "crisp", "air", "of", "a", "fall", "morning", "combined", "with", "the", "sight", "of", "vibrant", "orange", "and", "red", "leaves", "makes", "for", "one", "of", "the", "most", "invigorating", "starts", "to", "a", "day", "the", "ground", "beneath", "your", "feet", "is", "covered", "in", "a", "soft", "carpet", "of", "fallen", "leaves", "and", "each", "step", "brings", "the", "satisfying", "crunch", "of", "autumn", "underfoot", "as", "the", "cool", "breeze", "carries", "with", "it", "the", "scent", "of", "wood", "smoke", "and", "apples", "there"}},
		{"complex 2", args{"thesoundoflaughterechoingthroughacrowdedparkonasunnyafternoonisoneoflifessimplepleasuresfamiliesgatherforpicnicsfriendstossfrisbeesandchildrenrunfreelythroughthegrasstheirgigglesfillingtheairthesunshinesdowncastinglongshadowsasthedaybeginstowinddownandthescentoffreshlycutgrassmingleswith"}, []string{"the", "sound", "of", "laughter", "echoing", "through", "a", "crowded", "park", "on", "a", "sunny", "afternoon", "is", "one", "of", "lifes", "simple", "pleasures", "families", "gather", "for", "picnics", "friends", "toss", "frisbees", "and", "children", "run", "freely", "through", "the", "grass", "their", "giggles", "filling", "the", "air", "the", "sun", "shines", "down", "casting", "long", "shadows", "as", "the", "day", "begins", "to", "wind", "down", "and", "the", "scent", "of", "freshly", "cut", "grass", "mingles", "with"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := segmentor.Segment(tt.args.Text)
			assert.Equalf(t, tt.want, got, fmt.Sprintf("Segmentor.Segment() = %v, want %v\n", got, tt.want))
		})
	}
}
