package main

import (
	"encoding/json"
	"log"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

func main() {
	idx := "staging_doctors_mig"
	queries := []algoliasearch.IndexedQuery{
		{
			IndexName: idx,
			Params: algoliasearch.Map{
				"facetFilters":         []string{"locations.place_id:ChIJGTKkJjRdXz4RC5-FXb6asggE"},
				"hitsPerPage":          10,
				"page":                 1,
				"attributesToRetrieve": []string{"fullname"},
			},
		},
		{
			IndexName: idx,
			Params: algoliasearch.Map{
				"query":       "tes",
				"hitsPerPage": 3,
			},
		},
		// {
		// 	IndexName: "products",
		// 	Params: algoliasearch.Map{
		// 		"query":       "computer",
		// 		"hitsPerPage": 10,
		// 	},
		// },
	}

	client := algoliasearch.NewClient("6A1DZXM196", "a5cf84cbdf47ca86b90dd98a2fca9b60")

	res, err := client.MultipleQueries(queries, "stopIfEnoughMatches")
	log.Println(err)
	for _, v := range res {
		b, _ := json.Marshal(v.QueryRes.Hits)
		log.Printf("%+v\n", string(b))
	}
}
