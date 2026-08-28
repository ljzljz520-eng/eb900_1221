package report

import "warehouse5s/internal/domain"

func StatusLabel(status string) string {
	switch status {
	case "draft":
		return "待审核"
	case "review":
		return "审核中"
	case "approved":
		return "已批准"
	case "archived":
		return "已归档"
	default:
		return "未知"
	}
}
func ItemLabel(item domain.Item) string {
	if item.Complete {
		return item.Area + ":完成"
	}
	return item.Area + ":待整改"
}
func AreaLabel1(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel2(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel3(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel4(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel5(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel6(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel7(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel8(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel9(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel10(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel11(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel12(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel13(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel14(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel15(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel16(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel17(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel18(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel19(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel20(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel21(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel22(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel23(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel24(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel25(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel26(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel27(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel28(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel29(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel30(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel31(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel32(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel33(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel34(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel35(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel36(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel37(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel38(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel39(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel40(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel41(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel42(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel43(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel44(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel45(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel46(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel47(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel48(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel49(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel50(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel51(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel52(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel53(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel54(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel55(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel56(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel57(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel58(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel59(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel60(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel61(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel62(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel63(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel64(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel65(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel66(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel67(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel68(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel69(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel70(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel71(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel72(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel73(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel74(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel75(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel76(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel77(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel78(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel79(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel80(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel81(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel82(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel83(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel84(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel85(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel86(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel87(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel88(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel89(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel90(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel91(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel92(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel93(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel94(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel95(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel96(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel97(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel98(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel99(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel100(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel101(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel102(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel103(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel104(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel105(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
func AreaLabel106(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-1" {
			n++
		}
	}
	return n
}
func AreaLabel107(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-2" {
			n++
		}
	}
	return n
}
func AreaLabel108(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-3" {
			n++
		}
	}
	return n
}
func AreaLabel109(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-4" {
			n++
		}
	}
	return n
}
func AreaLabel110(items []domain.Item) int {
	n := 0
	for _, item := range items {
		if item.Area == "area-5" {
			n++
		}
	}
	return n
}
