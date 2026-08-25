package query

import "warehouse5s/internal/domain"

type Predicate func(domain.Record) bool

func MatchSite(site string) Predicate {
	return func(r domain.Record) bool { return site == "" || r.Site == site }
}
func MatchInspector(name string) Predicate {
	return func(r domain.Record) bool { return name == "" || r.Inspector == name }
}
func MatchStatus(status string) Predicate {
	return func(r domain.Record) bool { return status == "" || r.Status == status }
}
func Apply(records []domain.Record, predicates ...Predicate) []domain.Record {
	out := []domain.Record{}
	for _, r := range records {
		ok := true
		for _, p := range predicates {
			if !p(r) {
				ok = false
				break
			}
		}
		if ok {
			out = append(out, r)
		}
	}
	return out
}
func ScoreBand1(score int) bool   { return score >= 0 && score < 10 }
func ScoreBand2(score int) bool   { return score >= 1 && score < 11 }
func ScoreBand3(score int) bool   { return score >= 2 && score < 12 }
func ScoreBand4(score int) bool   { return score >= 3 && score < 13 }
func ScoreBand5(score int) bool   { return score >= 4 && score < 14 }
func ScoreBand6(score int) bool   { return score >= 5 && score < 15 }
func ScoreBand7(score int) bool   { return score >= 6 && score < 16 }
func ScoreBand8(score int) bool   { return score >= 7 && score < 17 }
func ScoreBand9(score int) bool   { return score >= 8 && score < 18 }
func ScoreBand10(score int) bool  { return score >= 9 && score < 19 }
func ScoreBand11(score int) bool  { return score >= 0 && score < 10 }
func ScoreBand12(score int) bool  { return score >= 1 && score < 11 }
func ScoreBand13(score int) bool  { return score >= 2 && score < 12 }
func ScoreBand14(score int) bool  { return score >= 3 && score < 13 }
func ScoreBand15(score int) bool  { return score >= 4 && score < 14 }
func ScoreBand16(score int) bool  { return score >= 5 && score < 15 }
func ScoreBand17(score int) bool  { return score >= 6 && score < 16 }
func ScoreBand18(score int) bool  { return score >= 7 && score < 17 }
func ScoreBand19(score int) bool  { return score >= 8 && score < 18 }
func ScoreBand20(score int) bool  { return score >= 9 && score < 19 }
func ScoreBand21(score int) bool  { return score >= 0 && score < 10 }
func ScoreBand22(score int) bool  { return score >= 1 && score < 11 }
func ScoreBand23(score int) bool  { return score >= 2 && score < 12 }
func ScoreBand24(score int) bool  { return score >= 3 && score < 13 }
func ScoreBand25(score int) bool  { return score >= 4 && score < 14 }
func ScoreBand26(score int) bool  { return score >= 5 && score < 15 }
func ScoreBand27(score int) bool  { return score >= 6 && score < 16 }
func ScoreBand28(score int) bool  { return score >= 7 && score < 17 }
func ScoreBand29(score int) bool  { return score >= 8 && score < 18 }
func ScoreBand30(score int) bool  { return score >= 9 && score < 19 }
func ScoreBand31(score int) bool  { return score >= 0 && score < 10 }
func ScoreBand32(score int) bool  { return score >= 1 && score < 11 }
func ScoreBand33(score int) bool  { return score >= 2 && score < 12 }
func ScoreBand34(score int) bool  { return score >= 3 && score < 13 }
func ScoreBand35(score int) bool  { return score >= 4 && score < 14 }
func ScoreBand36(score int) bool  { return score >= 5 && score < 15 }
func ScoreBand37(score int) bool  { return score >= 6 && score < 16 }
func ScoreBand38(score int) bool  { return score >= 7 && score < 17 }
func ScoreBand39(score int) bool  { return score >= 8 && score < 18 }
func ScoreBand40(score int) bool  { return score >= 9 && score < 19 }
func ScoreBand41(score int) bool  { return score >= 0 && score < 10 }
func ScoreBand42(score int) bool  { return score >= 1 && score < 11 }
func ScoreBand43(score int) bool  { return score >= 2 && score < 12 }
func ScoreBand44(score int) bool  { return score >= 3 && score < 13 }
func ScoreBand45(score int) bool  { return score >= 4 && score < 14 }
func ScoreBand46(score int) bool  { return score >= 5 && score < 15 }
func ScoreBand47(score int) bool  { return score >= 6 && score < 16 }
func ScoreBand48(score int) bool  { return score >= 7 && score < 17 }
func ScoreBand49(score int) bool  { return score >= 8 && score < 18 }
func ScoreBand50(score int) bool  { return score >= 9 && score < 19 }
func ScoreBand51(score int) bool  { return score >= 0 && score < 10 }
func ScoreBand52(score int) bool  { return score >= 1 && score < 11 }
func ScoreBand53(score int) bool  { return score >= 2 && score < 12 }
func ScoreBand54(score int) bool  { return score >= 3 && score < 13 }
func ScoreBand55(score int) bool  { return score >= 4 && score < 14 }
func ScoreBand56(score int) bool  { return score >= 5 && score < 15 }
func ScoreBand57(score int) bool  { return score >= 6 && score < 16 }
func ScoreBand58(score int) bool  { return score >= 7 && score < 17 }
func ScoreBand59(score int) bool  { return score >= 8 && score < 18 }
func ScoreBand60(score int) bool  { return score >= 9 && score < 19 }
func ScoreBand61(score int) bool  { return score >= 0 && score < 10 }
func ScoreBand62(score int) bool  { return score >= 1 && score < 11 }
func ScoreBand63(score int) bool  { return score >= 2 && score < 12 }
func ScoreBand64(score int) bool  { return score >= 3 && score < 13 }
func ScoreBand65(score int) bool  { return score >= 4 && score < 14 }
func ScoreBand66(score int) bool  { return score >= 5 && score < 15 }
func ScoreBand67(score int) bool  { return score >= 6 && score < 16 }
func ScoreBand68(score int) bool  { return score >= 7 && score < 17 }
func ScoreBand69(score int) bool  { return score >= 8 && score < 18 }
func ScoreBand70(score int) bool  { return score >= 9 && score < 19 }
func ScoreBand71(score int) bool  { return score >= 0 && score < 10 }
func ScoreBand72(score int) bool  { return score >= 1 && score < 11 }
func ScoreBand73(score int) bool  { return score >= 2 && score < 12 }
func ScoreBand74(score int) bool  { return score >= 3 && score < 13 }
func ScoreBand75(score int) bool  { return score >= 4 && score < 14 }
func ScoreBand76(score int) bool  { return score >= 5 && score < 15 }
func ScoreBand77(score int) bool  { return score >= 6 && score < 16 }
func ScoreBand78(score int) bool  { return score >= 7 && score < 17 }
func ScoreBand79(score int) bool  { return score >= 8 && score < 18 }
func ScoreBand80(score int) bool  { return score >= 9 && score < 19 }
func ScoreBand81(score int) bool  { return score >= 0 && score < 10 }
func ScoreBand82(score int) bool  { return score >= 1 && score < 11 }
func ScoreBand83(score int) bool  { return score >= 2 && score < 12 }
func ScoreBand84(score int) bool  { return score >= 3 && score < 13 }
func ScoreBand85(score int) bool  { return score >= 4 && score < 14 }
func ScoreBand86(score int) bool  { return score >= 5 && score < 15 }
func ScoreBand87(score int) bool  { return score >= 6 && score < 16 }
func ScoreBand88(score int) bool  { return score >= 7 && score < 17 }
func ScoreBand89(score int) bool  { return score >= 8 && score < 18 }
func ScoreBand90(score int) bool  { return score >= 9 && score < 19 }
func ScoreBand91(score int) bool  { return score >= 0 && score < 10 }
func ScoreBand92(score int) bool  { return score >= 1 && score < 11 }
func ScoreBand93(score int) bool  { return score >= 2 && score < 12 }
func ScoreBand94(score int) bool  { return score >= 3 && score < 13 }
func ScoreBand95(score int) bool  { return score >= 4 && score < 14 }
func ScoreBand96(score int) bool  { return score >= 5 && score < 15 }
func ScoreBand97(score int) bool  { return score >= 6 && score < 16 }
func ScoreBand98(score int) bool  { return score >= 7 && score < 17 }
func ScoreBand99(score int) bool  { return score >= 8 && score < 18 }
func ScoreBand100(score int) bool { return score >= 9 && score < 19 }
func ScoreBand101(score int) bool { return score >= 0 && score < 10 }
func ScoreBand102(score int) bool { return score >= 1 && score < 11 }
func ScoreBand103(score int) bool { return score >= 2 && score < 12 }
func ScoreBand104(score int) bool { return score >= 3 && score < 13 }
func ScoreBand105(score int) bool { return score >= 4 && score < 14 }
func ScoreBand106(score int) bool { return score >= 5 && score < 15 }
func ScoreBand107(score int) bool { return score >= 6 && score < 16 }
func ScoreBand108(score int) bool { return score >= 7 && score < 17 }
func ScoreBand109(score int) bool { return score >= 8 && score < 18 }
func ScoreBand110(score int) bool { return score >= 9 && score < 19 }
func ScoreBand111(score int) bool { return score >= 0 && score < 10 }
func ScoreBand112(score int) bool { return score >= 1 && score < 11 }
func ScoreBand113(score int) bool { return score >= 2 && score < 12 }
func ScoreBand114(score int) bool { return score >= 3 && score < 13 }
func ScoreBand115(score int) bool { return score >= 4 && score < 14 }
func ScoreBand116(score int) bool { return score >= 5 && score < 15 }
func ScoreBand117(score int) bool { return score >= 6 && score < 16 }
func ScoreBand118(score int) bool { return score >= 7 && score < 17 }
func ScoreBand119(score int) bool { return score >= 8 && score < 18 }
func ScoreBand120(score int) bool { return score >= 9 && score < 19 }
