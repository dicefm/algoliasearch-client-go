package algoliasearch

func checkSettings(settings map[string]interface{}) error {
	for k, v := range settings {
		switch k {
		case "altCorrections":
			if _, ok := v.([]Alternative); !ok {
				return invalidType(k, "[]Alternative")
			}

		case "synonyms":
			if _, ok := v.([][]string); !ok {
				return invalidType(k, "[][]string")
			}

		case "attributesForDistinct",
			"attributesForFaceting",
			"attributesToIndex",
			"numericAttributesToIndex",
			"ranking",
			"slaves",
			"unretrievableAttributes",
			"disableTypoToleranceOnAttributes",
			"disableTypoToleranceOnWords",
			"attributesToHighlight",
			"attributesToRetrieve",
			"attributesToSnippet",
			"optionalWords":
			if _, ok := v.([]string); !ok {
				return invalidType(k, "[]string")
			}

		case "allowCompressionOfIntegerArray",
			"advancedSyntax",
			"allowTyposOnNumericTokens",
			"ignorePlurals",
			"removeStopWords",
			"replaceSynonymsInHighlight":
			if _, ok := v.(bool); !ok {
				return invalidType(k, "bool")
			}

		case "distinct",
			"hitsPerPage",
			"maxValuesPerFacet",
			"minProximity",
			"minWordSizefor1Typo",
			"minWordSizefor2Typos":
			if _, ok := v.(int); !ok {
				return invalidType(k, "int")
			}

		case "placeholders":
			if _, ok := v.(map[string][]string); !ok {
				return invalidType(k, "map[string][]string")
			}

		case "separatorsToIndex",
			"highlightPostTag",
			"highlightPreTag",
			"queryType",
			"snippetEllipsisText",
			"typoTolerance":
			if _, ok := v.(string); !ok {
				return invalidType(k, "string")
			}

		default:
		}
	}

	return nil
}