package code

const SecretKey  = "-----BEGIN PUBLIC KEY-----D48Tc13U59E4IU5kuAmNcVB59cnFEgkaRO0wejPRatyq7iLD3tazGZQEsFMI5UmonKLPSVyovVmStplCyWiOgvzCluqB6k1E2hNs-----END PUBLIC KEY-----"
type countryCode struct {
	ChineseNameSimple  map[string]string
	EnglishName  map[string]string
}


var CountryCode = countryCode{
	ChineseNameSimple: map[string]string{
		"CHN":"中国",
		"HKG":"中国香港",
		"NZL":"新西兰",
		"AUS":"澳大利亚",
		"USA":"美国",
	},
}