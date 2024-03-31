package spacial_queries

import (
	"testing"
)

var points = []Point[string]{
	{58.000000, 56.316666, "Perm"},
	{54.733334, 56.000000, "Ufa"},
	{47.233334, 39.700001, "Rostov on Don"},
	{54.983334, 73.366669, "Omsk"},
	{55.796391, 49.108891, "Kazan"},
	{56.833332, 60.583332, "Yekaterinburg"},
	{59.937500, 30.308611, "St Petersburg"},
	{58.200348, 68.256332, "Tobolsk"},
	{54.766323, 83.086037, "Berdsk"},
	{54.715424, 20.509207, "Kaliningrad"},
	{57.629971, 39.872799, "Yaroslavl"},
	{50.272778, 127.540405, "Blagoveshchensk"},
	{46.358803, 48.059937, "Astrakhan"},
	{67.496780, 64.060638, "Vorkuta"},
	{45.039268, 38.987221, "Krasnodar"},
	{51.776272, 55.099594, "Orenburg"},
	{42.966633, 47.512630, "Makhachkala"},
	{52.723598, 41.442307, "Tambov"},
	{53.241505, 50.221245, "Samara"},
	{51.592365, 45.960804, "Saratov"},
	{57.161297, 65.525017, "Tyumen"},
	{55.164440, 61.436844, "Chelyabinsk"},
	{62.035454, 129.675476, "Yakutsk"},
	{59.410412, 56.791721, "Berezniki"},
	{55.688713, 37.901073, "Lyubertsy"},
	{66.530426, 66.613708, "Salekhard"},
	{56.143063, 40.410934, "Vladimir"},
	{43.588348, 39.729996, "Sochi"},
	{64.542465, 40.537319, "Arkhangelsk"},
	{57.767193, 40.976257, "Kostroma"},
	{48.700001, 44.516666, "Volgograd"},
	{61.666668, 50.816666, "Syktyvkar"},
	{56.633331, 47.866669, "Yoshkar-Ola"},
	{59.566666, 150.800003, "Magadan"},
	{55.950001, 38.049999, "Fryazino"},
	{59.222340, 39.882431, "Vologda"},
	{55.751244, 37.618423, "Moscow"},
	{55.018803, 82.933952, "Novosibirsk"},
	{55.854476, 38.441852, "Noginsk"},
	{55.920244, 37.991474, "Schelkovo"},
}

func TestUnit_GeohashBtree_Set_Ok(t *testing.T) {
	r := NewGeohashBTree[string]()

	for _, point := range points {
		r.Set(&point)
	}

	println(r)
}

func TestUnit_GeohashBtree_RangeSearch_Ok(t *testing.T) {
	r := NewGeohashBTree[string]()

	for _, point := range points {
		r.Set(&point)
	}

	requiredItems := map[string]struct{}{
		"Lyubertsy": {},
		"Moscow":    {},
		"Noginsk":   {},
		"Schelkovo": {},
		"Fryazino":  {},
	}

	r.RangeSearch(
		&RangeSearchInput{
			minLat: 54.247425,
			minLng: 35.718781,
			maxLat: 57.197908,
			maxLng: 40.873665,
		}, func(point Point[string]) bool {
			_, ok := requiredItems[point.v]
			if !ok {
				t.Errorf("extra item in return: %s", point.v)
			}

			delete(requiredItems, point.v)

			return true
		},
	)

	if len(requiredItems) != 0 {
		t.Errorf("items not found: %v", requiredItems)
	}
}
