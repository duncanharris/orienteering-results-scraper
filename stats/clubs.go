package stats

import "strings"

func isBOClub(s string) bool {
	_, ok := clubs[strings.ToUpper(s)]
	return ok
}

var (
	s = struct{}{}

	// from British Orienteering sign-up form
	clubs = map[string]struct{}{
		"AIRE":   s,
		"AROS":   s,
		"AUOC":   s,
		"AYROC":  s,
		"BADO":   s,
		"BAOC":   s,
		"BASOC":  s,
		"BKO":    s,
		"BL":     s,
		"BOF":    s,
		"BOK":    s,
		"BUMC":   s,
		"CHIG":   s,
		"CLARO":  s,
		"CLOK":   s,
		"CLYDE":  s,
		"CUOC":   s,
		"DEE":    s,
		"DEVON":  s,
		"DFOK":   s,
		"DRONGO": s,
		"DUOC":   s,
		"DVO":    s,
		"EBOR":   s,
		"ECKO":   s,
		"ELO":    s,
		"EPOC":   s,
		"ERYRI":  s,
		"ESOC":   s,
		"EUOC":   s,
		"FERMO":  s,
		"FVO":    s,
		"GMOA":   s,
		"GO":     s,
		"GRAMP":  s,
		"HALO":   s,
		"HAVOC":  s,
		"HH":     s,
		"HOC":    s,
		"INT":    s,
		"INVOC":  s,
		"OK":     s,
		"JOK":    s,
		"KERNO":  s,
		"KFO":    s,
		"LEI":    s,
		"LOC":    s,
		"LOG":    s,
		"LOK":    s,
		"LUOC":   s,
		"LVO":    s,
		"MA":     s,
		"MAROC":  s,
		"MDOC":   s,
		"MOR":    s,
		"MV":     s,
		"NATO":   s,
		"NGOC":   s,
		"NN":     s,
		"NOC":    s,
		"NOR":    s,
		"NWO":    s,
		"NWOC":   s,
		"OD":     s,
		"OUOC":   s,
		"PFO":    s,
		"POTOC":  s,
		"QO":     s,
		"RAFO":   s,
		"RNOC":   s,
		"RR":     s,
		"SARUM":  s,
		"SAX":    s,
		"SBOC":   s,
		"SELOC":  s,
		"SHUOC":  s,
		"SLOW":   s,
		"SMOC":   s,
		"SN":     s,
		"SO":     s,
		"SOC":    s,
		"SOLWAY": s,
		"SOS":    s,
		"SROC":   s,
		"STAG":   s,
		"STUOC":  s,
		"SUFFOC": s,
		"SWOC":   s,
		"SYO":    s,
		"TAY":    s,
		"TINTO":  s,
		"TVOC":   s,
		"UBOC":   s,
		"WAOC":   s,
		"WAROC":  s,
		"WCH":    s,
		"WCOC":   s,
		"WIGHTO": s,
		"WIM":    s,
		"WRE":    s,
		"WSX":    s,
	}
)
