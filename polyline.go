package main

import (
	"fmt"
	"github.com/twpayne/gopolyline/polyline"
	"github.com/gershwinlabs/gokml"
	"io/ioutil"
)

func main() {

	test:= "e`miGhmocNpD{TeLm|@wTcaBce@tJ{bAhb@ci@nFeVsHas@}YsWkc@hA}k@w[e}@ke@_Q}f@}b@y}@p_@qiBtJc`BxRs_A~A{CweByOgnB{q@gaFeh@m~D}Xo`Ay`@mwBw\\atEbJat@`PoaAkJqh@o{AurCmrBcmEgaAynCoa@srAwVgvBas@s`Foq@ccF~Ai~Bt@mfCuzBq}NKutBxt@_jBhKcx@aTapDik@qsIobA}}Ca\\kfFqa@}gGgnA_aKaoEmn\\ypAswJ{RonEkYgnD{|@cmGnZmiByTqxCmt@{bCcMclByp@qtBwUqkCvJecC~Rkt@dKoeAsoAarLms@czKsc@{`Egd@imC}c@czAgf@uiBmWsgAeZm[_`@qqAgZwyCc`@a{@sq@s|@yL{tAwc@kjDex@mmCwz@wbBsp@wnDwt@k}BcDkhDcoAgpDydCwyG{aAqpD_tBssM_hBe`KscAsxEaiAwyHjB}cApWmfB{cAy`OutCo`a@k\\wyBqv@ozCL{lAgOiyBgXahBmCkaCgUs}GmMiwIs_@kfR_SepR_Maq@tAel@veAusDzd@ezD`Pe{B{Cqs@qg@wmAo`@oaCo|@chEej@w|@ir@mpAqp@uhBuXg|Bkc@idIok@u_Oop@_tDua@igFpKcyDiw@kjEs|A}fFix@grCuKycAh@c|D}j@{_EkcAioBg^}rA{z@wvAor@qw@uj@woA_g@_k@o|@uU{iBm|@ojBwvAomBunBkrBywBi`@ssA}MwsAlIylAeH{f@cp@k`BmZsGaqAZec@wUcqAyzBuc@ec@a~@q_@yyAs_DkkDctFu|AsgCuoAqhBqgJwoLkc@ss@wWozA}h@qyAoxDgrGcp@ydA_q@i`@_j@uZ}]oh@wmEuhI_tAcgCuuC_pFsjByfEkfB{hEqtAi`FufDwoKe}D}nMyi@qbBuQ{Mk~@cRkQuU}gCutIgiAcyDusFe{RwwAguHt@{m@to@e`Cva@yhBtHgjBs@yvFgs@oaDaM{r@qE_qEqTe{B_wAszEiiBmiFmsEi~LkpAe~BwkAoqDcfDolLsn@ed@kuAcpDuxAi_Dmn@szCcnFwoNkpDevEmn@_`As}@whCkuDmyLqeAgiDqq@}jAelA{gBkcAmwBmVkkDam@mwBs_Ayx@_[ooAw`A_wDqf@ebDel@akFo_Ai{F_YqoAaBysAdL}qFwGipAhCm|EpHywClWcjAT{n@cn@w{AgcB{_Egq@{j@sYkMai@w_Aiw@_uAwIoO_PkIkF~MmHn@gCiD"
	points, err := polyline.Decode(test, 2)

	if err != nil {
		return
	}

	kml := gokml.NewKML("teste")
	//style := gokml.NewStyle("linestringStyle", 127, 255, 255, 0)
	linestring := gokml.NewLineString()
	placemark := gokml.NewPlacemark("Rota", "route between two points", linestring)
	//placemark.SetStyle("linestringStyle")
	kml.AddFeature(placemark)
	//kml.AddFeature(style)

	for i := 0; i < len(points); i = i+2 {
		point := gokml.NewPoint(points[i],points[i+1],0)
		linestring.AddPoint(point)
	}

	kmlString := kml.Render()

	err = ioutil.WriteFile("teste.kml", []byte(kmlString), 0644)
	check(err)

	fmt.Printf("%+v\n", kmlString)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
