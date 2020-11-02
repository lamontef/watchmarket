package fixer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInitProvider(t *testing.T) {
	provider := InitProvider("demo.api", "key", "USD")
	assert.NotNil(t, provider)
	assert.Equal(t, "demo.api", provider.client.api)
	assert.Equal(t, "key", provider.client.key)
	assert.Equal(t, "fixer", provider.id)
	assert.Equal(t, "USD", provider.currency)
}

func TestProvider_GetProvider(t *testing.T) {
	provider := InitProvider("demo.api", "key", "USD")
	assert.Equal(t, "fixer", provider.GetProvider())
}

func createMockedAPI() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/latest", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedResponse); err != nil {
			panic(err)
		}
	})

	return r
}

var (
	mockedResponse = `{ "success": true, "timestamp": 1589030105, "base": "USD", "date": "2020-05-09", "rates": { "AED": 3.67315, "AFN": 76.203991, "ALL": 114.903989, "AMD": 484.110403, "ANG": 1.795199, "AOA": 551.905041, "ARS": 66.472813, "AUD": 1.530402, "AWG": 1.8, "AZN": 1.70397, "BAM": 1.803839, "BBD": 2.019276, "BDT": 84.99451, "BGN": 1.80549, "BHD": 0.378178, "BIF": 1906, "BMD": 1, "BND": 1.413261, "BOB": 6.895677, "BRL": 5.732404, "BSD": 1.000046, "BTC": 0.000103, "BTN": 75.58479, "BWP": 12.144339, "BYN": 2.435844, "BYR": 19600, "BZD": 2.015864, "CAD": 1.40175, "CDF": 1810.000362, "CHF": 0.970982, "CLF": 0.029928, "CLP": 825.803912, "CNY": 7.074204, "COP": 3896.5, "CRC": 568.93705, "CUC": 1, "CUP": 26.5, "CVE": 102.00036, "CZK": 25.132604, "DJF": 177.720394, "DKK": 6.880904, "DOP": 55.040393, "DZD": 128.33601, "EGP": 15.563556, "ERN": 15.000358, "ETB": 33.703876, "EUR": 0.911536, "FJD": 2.25304, "FKP": 0.805997, "GBP": 0.806062, "GEL": 3.220391, "GGP": 0.805997, "GHS": 5.74504, "GIP": 0.805997, "GMD": 51.18039, "GNF": 9450.000355, "GTQ": 7.705603, "GYD": 209.34286, "HKD": 7.76695, "HNL": 25.030389, "HRK": 6.938593, "HTG": 107.33289, "HUF": 323.140388, "IDR": 14773.435, "ILS": 3.506704, "IMP": 0.805997, "INR": 75.50404, "IQD": 1190, "IRR": 42105.000352, "ISK": 146.240386, "JEP": 0.805997, "JMD": 142.83448, "JOD": 0.709504, "JPY": 106.67604, "KES": 106.050385, "KGS": 78.903801, "KHR": 4110.000351, "KMF": 453.750384, "KPW": 900.000306, "KRW": 1219.803792, "KWD": 0.30935, "KYD": 0.833556, "KZT": 421.99093, "LAK": 9005.000349, "LBP": 1512.763039, "LKR": 186.51808, "LRD": 198.503775, "LSL": 18.390382, "LTL": 2.95274, "LVL": 0.60489, "LYD": 1.420381, "MAD": 9.825039, "MDL": 17.831157, "MGA": 3800.000347, "MKD": 56.826776, "MMK": 1395.135304, "MNT": 2794.123058, "MOP": 7.984837, "MRO": 357.00003, "MUR": 39.710379, "MVR": 15.503741, "MWK": 737.503739, "MXN": 23.672604, "MYR": 4.334039, "MZN": 68.080377, "NAD": 18.530377, "NGN": 390.000344, "NIO": 34.403725, "NOK": 10.217039, "NPR": 120.93559, "NZD": 1.629196, "OMR": 0.383447, "PAB": 1.000138, "PEN": 3.399039, "PGK": 3.430375, "PHP": 50.495039, "PKR": 159.650375, "PLN": 4.20475, "PYG": 6531.670104, "QAR": 3.641038, "RON": 4.453404, "RSD": 108.455038, "RUB": 73.40369, "RWF": 937.5, "SAR": 3.756262, "SBD": 8.267992, "SCR": 17.168052, "SDG": 55.325038, "SEK": 9.771904, "SGD": 1.412704, "SHP": 0.805997, "SLL": 9860.000339, "SOS": 583.000338, "SRD": 7.458038, "STD": 22051.386135, "SVC": 8.751671, "SYP": 514.451644, "SZL": 18.52037, "THB": 32.02037, "TJS": 10.265663, "TMT": 3.51, "TND": 2.912504, "TOP": 2.31435, "TRY": 7.089104, "TTD": 6.757574, "TWD": 29.857038, "TZS": 2314.203635, "UAH": 26.838273, "UGX": 3800.322804, "USD": 1, "UYU": 43.139569, "UZS": 10109.000335, "VEF": 9.987504, "VND": 23400.5, "VUV": 119.848296, "WST": 2.78472, "XAF": 604.99802, "XAG": 0.06474, "XAU": 0.000587, "XCD": 2.70255, "XDR": 0.734807, "XOF": 605.000332, "XPF": 110.4036, "YER": 250.375037, "ZAR": 18.350904, "ZMK": 9001.203593, "ZMW": 18.576435, "ZWL": 322.000001 } }`
	wantedRates    = `[{"currency":"AED","provider":"fixer","rate":0.2722458925,"timestamp":1589030105},{"currency":"AFN","provider":"fixer","rate":0.0131226723,"timestamp":1589030105},{"currency":"ALL","provider":"fixer","rate":0.0087029181,"timestamp":1589030105},{"currency":"AMD","provider":"fixer","rate":0.0020656445,"timestamp":1589030105},{"currency":"ANG","provider":"fixer","rate":0.5570413085,"timestamp":1589030105},{"currency":"AOA","provider":"fixer","rate":0.0018119059,"timestamp":1589030105},{"currency":"ARS","provider":"fixer","rate":0.0150437443,"timestamp":1589030105},{"currency":"AUD","provider":"fixer","rate":0.6534230875,"timestamp":1589030105},{"currency":"AWG","provider":"fixer","rate":0.5555555556,"timestamp":1589030105},{"currency":"AZN","provider":"fixer","rate":0.5868647922,"timestamp":1589030105},{"currency":"BAM","provider":"fixer","rate":0.5543732007,"timestamp":1589030105},{"currency":"BBD","provider":"fixer","rate":0.4952270022,"timestamp":1589030105},{"currency":"BDT","provider":"fixer","rate":0.0117654658,"timestamp":1589030105},{"currency":"BGN","provider":"fixer","rate":0.5538662635,"timestamp":1589030105},{"currency":"BHD","provider":"fixer","rate":2.6442574661,"timestamp":1589030105},{"currency":"BIF","provider":"fixer","rate":0.000524659,"timestamp":1589030105},{"currency":"BMD","provider":"fixer","rate":1,"timestamp":1589030105},{"currency":"BND","provider":"fixer","rate":0.7075833834,"timestamp":1589030105},{"currency":"BOB","provider":"fixer","rate":0.1450183934,"timestamp":1589030105},{"currency":"BRL","provider":"fixer","rate":0.1744468813,"timestamp":1589030105},{"currency":"BSD","provider":"fixer","rate":0.9999540021,"timestamp":1589030105},{"currency":"BTC","provider":"fixer","rate":9708.7378640777,"timestamp":1589030105},{"currency":"BTN","provider":"fixer","rate":0.013230175,"timestamp":1589030105},{"currency":"BWP","provider":"fixer","rate":0.0823428924,"timestamp":1589030105},{"currency":"BYN","provider":"fixer","rate":0.4105353216,"timestamp":1589030105},{"currency":"BYR","provider":"fixer","rate":0.0000510204,"timestamp":1589030105},{"currency":"BZD","provider":"fixer","rate":0.4960652107,"timestamp":1589030105},{"currency":"CAD","provider":"fixer","rate":0.7133939718,"timestamp":1589030105},{"currency":"CDF","provider":"fixer","rate":0.0005524861,"timestamp":1589030105},{"currency":"CHF","provider":"fixer","rate":1.029885209,"timestamp":1589030105},{"currency":"CLF","provider":"fixer","rate":33.4135257952,"timestamp":1589030105},{"currency":"CLP","provider":"fixer","rate":0.0012109412,"timestamp":1589030105},{"currency":"CNY","provider":"fixer","rate":0.1413586603,"timestamp":1589030105},{"currency":"COP","provider":"fixer","rate":0.0002566406,"timestamp":1589030105},{"currency":"CRC","provider":"fixer","rate":0.0017576637,"timestamp":1589030105},{"currency":"CUC","provider":"fixer","rate":1,"timestamp":1589030105},{"currency":"CUP","provider":"fixer","rate":0.0377358491,"timestamp":1589030105},{"currency":"CVE","provider":"fixer","rate":0.009803887,"timestamp":1589030105},{"currency":"CZK","provider":"fixer","rate":0.039788953,"timestamp":1589030105},{"currency":"DJF","provider":"fixer","rate":0.0056268162,"timestamp":1589030105},{"currency":"DKK","provider":"fixer","rate":0.1453297416,"timestamp":1589030105},{"currency":"DOP","provider":"fixer","rate":0.0181684749,"timestamp":1589030105},{"currency":"DZD","provider":"fixer","rate":0.0077920453,"timestamp":1589030105},{"currency":"EGP","provider":"fixer","rate":0.0642526682,"timestamp":1589030105},{"currency":"ERN","provider":"fixer","rate":0.0666650756,"timestamp":1589030105},{"currency":"ETB","provider":"fixer","rate":0.029670178,"timestamp":1589030105},{"currency":"EUR","provider":"fixer","rate":1.097049376,"timestamp":1589030105},{"currency":"FJD","provider":"fixer","rate":0.4438447609,"timestamp":1589030105},{"currency":"FKP","provider":"fixer","rate":1.2406994071,"timestamp":1589030105},{"currency":"GBP","provider":"fixer","rate":1.2405993584,"timestamp":1589030105},{"currency":"GEL","provider":"fixer","rate":0.3105213001,"timestamp":1589030105},{"currency":"GGP","provider":"fixer","rate":1.2406994071,"timestamp":1589030105},{"currency":"GHS","provider":"fixer","rate":0.1740631919,"timestamp":1589030105},{"currency":"GIP","provider":"fixer","rate":1.2406994071,"timestamp":1589030105},{"currency":"GMD","provider":"fixer","rate":0.0195387335,"timestamp":1589030105},{"currency":"GNF","provider":"fixer","rate":0.0001058201,"timestamp":1589030105},{"currency":"GTQ","provider":"fixer","rate":0.129775697,"timestamp":1589030105},{"currency":"GYD","provider":"fixer","rate":0.0047768527,"timestamp":1589030105},{"currency":"HKD","provider":"fixer","rate":0.1287506679,"timestamp":1589030105},{"currency":"HNL","provider":"fixer","rate":0.0399514366,"timestamp":1589030105},{"currency":"HRK","provider":"fixer","rate":0.1441214379,"timestamp":1589030105},{"currency":"HTG","provider":"fixer","rate":0.0093168087,"timestamp":1589030105},{"currency":"HUF","provider":"fixer","rate":0.0030946302,"timestamp":1589030105},{"currency":"IDR","provider":"fixer","rate":0.0000676891,"timestamp":1589030105},{"currency":"ILS","provider":"fixer","rate":0.2851680667,"timestamp":1589030105},{"currency":"IMP","provider":"fixer","rate":1.2406994071,"timestamp":1589030105},{"currency":"INR","provider":"fixer","rate":0.0132443244,"timestamp":1589030105},{"currency":"IQD","provider":"fixer","rate":0.0008403361,"timestamp":1589030105},{"currency":"IRR","provider":"fixer","rate":0.0000237501,"timestamp":1589030105},{"currency":"ISK","provider":"fixer","rate":0.0068380563,"timestamp":1589030105},{"currency":"JEP","provider":"fixer","rate":1.2406994071,"timestamp":1589030105},{"currency":"JMD","provider":"fixer","rate":0.0070011107,"timestamp":1589030105},{"currency":"JOD","provider":"fixer","rate":1.4094353238,"timestamp":1589030105},{"currency":"JPY","provider":"fixer","rate":0.0093741762,"timestamp":1589030105},{"currency":"KES","provider":"fixer","rate":0.0094294801,"timestamp":1589030105},{"currency":"KGS","provider":"fixer","rate":0.0126736607,"timestamp":1589030105},{"currency":"KHR","provider":"fixer","rate":0.000243309,"timestamp":1589030105},{"currency":"KMF","provider":"fixer","rate":0.0022038549,"timestamp":1589030105},{"currency":"KPW","provider":"fixer","rate":0.0011111107,"timestamp":1589030105},{"currency":"KRW","provider":"fixer","rate":0.000819804,"timestamp":1589030105},{"currency":"KWD","provider":"fixer","rate":3.2325844513,"timestamp":1589030105},{"currency":"KYD","provider":"fixer","rate":1.1996794457,"timestamp":1589030105},{"currency":"KZT","provider":"fixer","rate":0.0023697192,"timestamp":1589030105},{"currency":"LAK","provider":"fixer","rate":0.0001110494,"timestamp":1589030105},{"currency":"LBP","provider":"fixer","rate":0.0006610421,"timestamp":1589030105},{"currency":"LKR","provider":"fixer","rate":0.0053614105,"timestamp":1589030105},{"currency":"LRD","provider":"fixer","rate":0.0050376876,"timestamp":1589030105},{"currency":"LSL","provider":"fixer","rate":0.0543762495,"timestamp":1589030105},{"currency":"LTL","provider":"fixer","rate":0.338668491,"timestamp":1589030105},{"currency":"LVL","provider":"fixer","rate":1.6531931426,"timestamp":1589030105},{"currency":"LYD","provider":"fixer","rate":0.7040364522,"timestamp":1589030105},{"currency":"MAD","provider":"fixer","rate":0.1017807665,"timestamp":1589030105},{"currency":"MDL","provider":"fixer","rate":0.0560816104,"timestamp":1589030105},{"currency":"MGA","provider":"fixer","rate":0.0002631579,"timestamp":1589030105},{"currency":"MKD","provider":"fixer","rate":0.0175973383,"timestamp":1589030105},{"currency":"MMK","provider":"fixer","rate":0.0007167764,"timestamp":1589030105},{"currency":"MNT","provider":"fixer","rate":0.000357894,"timestamp":1589030105},{"currency":"MOP","provider":"fixer","rate":0.1252373718,"timestamp":1589030105},{"currency":"MRO","provider":"fixer","rate":0.0028011202,"timestamp":1589030105},{"currency":"MUR","provider":"fixer","rate":0.0251823333,"timestamp":1589030105},{"currency":"MVR","provider":"fixer","rate":0.0645005615,"timestamp":1589030105},{"currency":"MWK","provider":"fixer","rate":0.0013559253,"timestamp":1589030105},{"currency":"MXN","provider":"fixer","rate":0.0422429235,"timestamp":1589030105},{"currency":"MYR","provider":"fixer","rate":0.230731657,"timestamp":1589030105},{"currency":"MZN","provider":"fixer","rate":0.0146885203,"timestamp":1589030105},{"currency":"NAD","provider":"fixer","rate":0.0539654428,"timestamp":1589030105},{"currency":"NGN","provider":"fixer","rate":0.0025641003,"timestamp":1589030105},{"currency":"NIO","provider":"fixer","rate":0.02906662,"timestamp":1589030105},{"currency":"NOK","provider":"fixer","rate":0.0978757153,"timestamp":1589030105},{"currency":"NPR","provider":"fixer","rate":0.0082688644,"timestamp":1589030105},{"currency":"NZD","provider":"fixer","rate":0.6137996902,"timestamp":1589030105},{"currency":"OMR","provider":"fixer","rate":2.6079223465,"timestamp":1589030105},{"currency":"PAB","provider":"fixer","rate":0.999862019,"timestamp":1589030105},{"currency":"PEN","provider":"fixer","rate":0.2942008021,"timestamp":1589030105},{"currency":"PGK","provider":"fixer","rate":0.2915133185,"timestamp":1589030105},{"currency":"PHP","provider":"fixer","rate":0.0198039257,"timestamp":1589030105},{"currency":"PKR","provider":"fixer","rate":0.0062636871,"timestamp":1589030105},{"currency":"PLN","provider":"fixer","rate":0.2378262679,"timestamp":1589030105},{"currency":"PYG","provider":"fixer","rate":0.0001531002,"timestamp":1589030105},{"currency":"QAR","provider":"fixer","rate":0.2746469551,"timestamp":1589030105},{"currency":"RON","provider":"fixer","rate":0.224547335,"timestamp":1589030105},{"currency":"RSD","provider":"fixer","rate":0.0092204108,"timestamp":1589030105},{"currency":"RUB","provider":"fixer","rate":0.0136232933,"timestamp":1589030105},{"currency":"RWF","provider":"fixer","rate":0.0010666667,"timestamp":1589030105},{"currency":"SAR","provider":"fixer","rate":0.2662221112,"timestamp":1589030105},{"currency":"SBD","provider":"fixer","rate":0.1209483512,"timestamp":1589030105},{"currency":"SCR","provider":"fixer","rate":0.0582477266,"timestamp":1589030105},{"currency":"SDG","provider":"fixer","rate":0.0180749989,"timestamp":1589030105},{"currency":"SEK","provider":"fixer","rate":0.1023342022,"timestamp":1589030105},{"currency":"SGD","provider":"fixer","rate":0.7078623689,"timestamp":1589030105},{"currency":"SHP","provider":"fixer","rate":1.2406994071,"timestamp":1589030105},{"currency":"SLL","provider":"fixer","rate":0.0001014199,"timestamp":1589030105},{"currency":"SOS","provider":"fixer","rate":0.0017152649,"timestamp":1589030105},{"currency":"SRD","provider":"fixer","rate":0.1340835217,"timestamp":1589030105},{"currency":"STD","provider":"fixer","rate":0.0000453486,"timestamp":1589030105},{"currency":"SVC","provider":"fixer","rate":0.1142638931,"timestamp":1589030105},{"currency":"SYP","provider":"fixer","rate":0.0019438173,"timestamp":1589030105},{"currency":"SZL","provider":"fixer","rate":0.0539946016,"timestamp":1589030105},{"currency":"THB","provider":"fixer","rate":0.0312301201,"timestamp":1589030105},{"currency":"TJS","provider":"fixer","rate":0.0974121204,"timestamp":1589030105},{"currency":"TMT","provider":"fixer","rate":0.2849002849,"timestamp":1589030105},{"currency":"TND","provider":"fixer","rate":0.3433471679,"timestamp":1589030105},{"currency":"TOP","provider":"fixer","rate":0.432086763,"timestamp":1589030105},{"currency":"TRY","provider":"fixer","rate":0.1410615502,"timestamp":1589030105},{"currency":"TTD","provider":"fixer","rate":0.1479821013,"timestamp":1589030105},{"currency":"TWD","provider":"fixer","rate":0.0334929406,"timestamp":1589030105},{"currency":"TZS","provider":"fixer","rate":0.0004321141,"timestamp":1589030105},{"currency":"UAH","provider":"fixer","rate":0.0372602216,"timestamp":1589030105},{"currency":"UGX","provider":"fixer","rate":0.0002631355,"timestamp":1589030105},{"currency":"USD","provider":"fixer","rate":1,"timestamp":1589030105},{"currency":"UYU","provider":"fixer","rate":0.0231805747,"timestamp":1589030105},{"currency":"UZS","provider":"fixer","rate":0.0000989217,"timestamp":1589030105},{"currency":"VEF","provider":"fixer","rate":0.1001251163,"timestamp":1589030105},{"currency":"VND","provider":"fixer","rate":0.0000427341,"timestamp":1589030105},{"currency":"VUV","provider":"fixer","rate":0.0083438817,"timestamp":1589030105},{"currency":"WST","provider":"fixer","rate":0.359102531,"timestamp":1589030105},{"currency":"XAF","provider":"fixer","rate":0.001652898,"timestamp":1589030105},{"currency":"XAG","provider":"fixer","rate":15.4464009886,"timestamp":1589030105},{"currency":"XAU","provider":"fixer","rate":1703.5775127768,"timestamp":1589030105},{"currency":"XCD","provider":"fixer","rate":0.3700209062,"timestamp":1589030105},{"currency":"XDR","provider":"fixer","rate":1.3609015701,"timestamp":1589030105},{"currency":"XOF","provider":"fixer","rate":0.0016528917,"timestamp":1589030105},{"currency":"XPF","provider":"fixer","rate":0.0090576757,"timestamp":1589030105},{"currency":"YER","provider":"fixer","rate":0.0039940084,"timestamp":1589030105},{"currency":"ZAR","provider":"fixer","rate":0.0544932282,"timestamp":1589030105},{"currency":"ZMK","provider":"fixer","rate":0.0001110963,"timestamp":1589030105},{"currency":"ZMW","provider":"fixer","rate":0.0538316421,"timestamp":1589030105},{"currency":"ZWL","provider":"fixer","rate":0.0031055901,"timestamp":1589030105}]`
)
