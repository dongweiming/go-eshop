package eshop

const ( //
	US = iota
	HK
	JP
	BR // 巴西 Brasil
	RU // 俄罗斯 Russia
	PL // 波兰 Poland
	MX // 墨西哥 México
	PE // 秘鲁 Peru
	CO // 哥伦比亚 Colombia
	ZA // 南非 South Africa
	CA // 加拿大 Canada
	AR // 阿根廷 Argentina
	CL // 智利 Chile
	AU // 澳大利亚 Australia

	AlgoliaID = "U3B6GR4UA3"
	AlgoliaKey = "c4da8be7fd29f0f5bfa42920b0a99dc7"

	PerPage = 200

	K_NORMAL = iota
	K_DLC // 占位
)

const (
	Americas = iota
	Asia
	Europe
)

var (
	Region = map[int][]int {
		Americas: []int{US, BR, CA, MX, PE, CO, AR, CL},
		Asia: []int{HK, JP, AU},
		Europe: []int{RU, PL, ZA},
	}
	AlgoliaIndexMap = map[int]string {
		US: "ncom_game_en_us",
		BR: "ncom_game_pt_br",
		MX: "ncom_game_es_mx",
		CA: "ncom_game_en_ca",
	}
	CountryMap = map[string]int {
			"US": US,
			"HK": HK,
			"JP": JP,
			"BR": BR,
			"RU": RU,
			"PL": PL,
			"MX": MX,
			"PE": PE,
			"CO": CO,
			"ZA": ZA,
			"CA": CA,
			"AR": AR,
			"CL": CL,
			"AU": AU,
	}
)
