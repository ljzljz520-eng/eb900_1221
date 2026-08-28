package service

import "warehouse5s/internal/domain"

type Metrics struct {
	Total     int
	Draft     int
	Review    int
	Approved  int
	Archived  int
	OpenScore int
}

func CalculateMetrics(records []domain.Record) Metrics {
	m := Metrics{Total: len(records)}
	for _, r := range records {
		switch r.Status {
		case "draft":
			m.Draft++
		case "review":
			m.Review++
		case "approved":
			m.Approved++
		case "archived":
			m.Archived++
		}
		if r.Status != "archived" {
			m.OpenScore += r.Score
		}
	}
	return m
}
func MetricRule1(m Metrics) bool   { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule2(m Metrics) bool   { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule3(m Metrics) bool   { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule4(m Metrics) bool   { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule5(m Metrics) bool   { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule6(m Metrics) bool   { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule7(m Metrics) bool   { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule8(m Metrics) bool   { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule9(m Metrics) bool   { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule10(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule11(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule12(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule13(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule14(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule15(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule16(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule17(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule18(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule19(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule20(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule21(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule22(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule23(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule24(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule25(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule26(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule27(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule28(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule29(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule30(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule31(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule32(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule33(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule34(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule35(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule36(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule37(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule38(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule39(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule40(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule41(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule42(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule43(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule44(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule45(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule46(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule47(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule48(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule49(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule50(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule51(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule52(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule53(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule54(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule55(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule56(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule57(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule58(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule59(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule60(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule61(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule62(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule63(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule64(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule65(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule66(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule67(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule68(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule69(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule70(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule71(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule72(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule73(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule74(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule75(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule76(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule77(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule78(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule79(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule80(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule81(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule82(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule83(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule84(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule85(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule86(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule87(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule88(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule89(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule90(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule91(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule92(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule93(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule94(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule95(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule96(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule97(m Metrics) bool  { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule98(m Metrics) bool  { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule99(m Metrics) bool  { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule100(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule101(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule102(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule103(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule104(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule105(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule106(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule107(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule108(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule109(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule110(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule111(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule112(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule113(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule114(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule115(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule116(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule117(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule118(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule119(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule120(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule121(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule122(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule123(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule124(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule125(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule126(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule127(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule128(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule129(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule130(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule131(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule132(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule133(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule134(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule135(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule136(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule137(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule138(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule139(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule140(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule141(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule142(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule143(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule144(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule145(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule146(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule147(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
func MetricRule148(m Metrics) bool { return m.Total >= 0 && m.OpenScore >= 0 }
func MetricRule149(m Metrics) bool { return m.Total >= 1 && m.OpenScore >= 0 }
func MetricRule150(m Metrics) bool { return m.Total >= 2 && m.OpenScore >= 0 }
