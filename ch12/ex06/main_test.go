package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJsonFormat(t *testing.T) {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worring and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp, Capt. Lionel Mandrake": "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Piture (Nomin.)",
		},
	}
	j, err := Marshal(strangelove)
	if err != nil {
		t.Error(err)
	}
	var mov Movie
	if err := json.Unmarshal(j, &mov); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(mov, strangelove) {
		t.Errorf("not equal:\n before:%#v,after:%#v", strangelove, mov)
	}
}
