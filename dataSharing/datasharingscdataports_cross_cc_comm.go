//line :1
package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

               
func jmSY84co(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) (string, error) {
	if  /*line c5CzduLS.go:1*/ len(ktsi1_SQ) != 1 {
		return "",  /*line nyxXNOX5.go:1*/ fmt.Errorf( /*line zD_iwzv5.go:1*/ func() string {
			var data []byte
			i := 14
			decryptKey := 2
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					i = 4
					data =  /*line zD_iwzv5.go:1*/ append(data, "\x9cE\x94"...,
					)
				case 15:
					i = 6
					data =  /*line zD_iwzv5.go:1*/ append(data, 140)
				case 9:
					i = 16
					data =  /*line zD_iwzv5.go:1*/ append(data, "i\u007f,"...,
					)
				case 3:
					i = 0
					for y := range data {
						data[y] = data[y] +  /*line zD_iwzv5.go:1*/ byte(decryptKey^y)
					}
				case 13:
					data =  /*line zD_iwzv5.go:1*/ append(data, "\x88\x84@"...,
					)
					i = 7
				case 8:
					i = 5
					data =  /*line zD_iwzv5.go:1*/ append(data, 141)
				case 7:
					i = 1
					data =  /*line zD_iwzv5.go:1*/ append(data, "7]"...,
					)
				case 11:
					data =  /*line zD_iwzv5.go:1*/ append(data, 132)
					i = 10
				case 14:
					i = 18
					data =  /*line zD_iwzv5.go:1*/ append(data, "Lrdq"...,
					)
				case 1:
					i = 12
					data =  /*line zD_iwzv5.go:1*/ append(data, "\x8d\x86\x80\u007f"...,
					)
				case 5:
					data =  /*line zD_iwzv5.go:1*/ append(data, "\x87=\x82"...,
					)
					i = 11
				case 10:
					i = 15
					data =  /*line zD_iwzv5.go:1*/ append(data, "\x98\x82\x95"...,
					)
				case 18:
					data =  /*line zD_iwzv5.go:1*/ append(data, "yzj"...,
					)
					i = 9
				case 16:
					i = 17
					data =  /*line zD_iwzv5.go:1*/ append(data, 106)
				case 17:
					data =  /*line zD_iwzv5.go:1*/ append(data, "|v\x85"...,
					)
					i = 2
				case 2:
					data =  /*line zD_iwzv5.go:1*/ append(data, "zs\x81"...,
					)
					i = 13
				case 12:
					i = 8
					data =  /*line zD_iwzv5.go:1*/ append(data, "\x8d\x83"...,
					)
				case 4:
					data =  /*line zD_iwzv5.go:1*/ append(data, "\x8c\x99\x8e"...,
					)
					i = 3
				}
			}
			return  /*line zD_iwzv5.go:1*/ string(data)
		}())
	}
	utUwSJgF := ktsi1_SQ[0]

	y32DtQHo := []string{ /*line CuLeAuO1.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 203
		for counter := 0; i != 8; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				data =  /*line CuLeAuO1.go:1*/ append(data, "w\x85\x81"...,
				)
				i = 0
			case 3:
				i = 5
				data =  /*line CuLeAuO1.go:1*/ append(data, "\x91}"...,
				)
			case 7:
				for y := range data {
					data[y] = data[y] +  /*line CuLeAuO1.go:1*/ byte(decryptKey^y)
				}
				i = 8
			case 2:
				data =  /*line CuLeAuO1.go:1*/ append(data, "gitk"...,
				)
				i = 7
			case 9:
				i = 6
				data =  /*line CuLeAuO1.go:1*/ append(data, 96)
			case 4:
				data =  /*line CuLeAuO1.go:1*/ append(data, 141)
				i = 9
			case 0:
				data =  /*line CuLeAuO1.go:1*/ append(data, "\x83\u007f"...,
				)
				i = 3
			case 5:
				i = 2
				data =  /*line CuLeAuO1.go:1*/ append(data, "]\x93"...,
				)
			case 1:
				data =  /*line CuLeAuO1.go:1*/ append(data, "\x89\x8c{\x87"...,
				)
				i = 4
			}
		}
		return  /*line CuLeAuO1.go:1*/ string(data)
	}(), utUwSJgF}
	fWJzYwn4 :=  /*line Fau_hJr5.go:1*/ make([][]byte,  /*line U2j8gEfN.go:1*/ len(y32DtQHo))
	for gGz2tP6V, vChBf6MJ := range y32DtQHo {
		fWJzYwn4[gGz2tP6V] =  /*line lfsq6mb9.go:1*/ []byte(vChBf6MJ)
	}

	szAT31uU :=  /*line HYkOyYxa.go:1*/ n4c7mRtG.InvokeChaincode( /*line UrAnEBEk.go:1*/ func() string {
		fullData :=  /*line UrAnEBEk.go:1*/ []byte("\xe0\xd2\xf6E\xc1&^?\xf3N\xe5\x81\"k\xa2/\x80TC(")
		var data []byte
		data =  /*line UrAnEBEk.go:1*/ append(data, fullData[19]+fullData[3], fullData[18]^fullData[5], fullData[17]-fullData[0], fullData[15]^fullData[9], fullData[10]^fullData[11], fullData[13]+fullData[2], fullData[1]+fullData[14], fullData[7]+fullData[12], fullData[8]-fullData[16], fullData[4]-fullData[6])
		return  /*line UrAnEBEk.go:1*/ string(data)
	}(), fWJzYwn4,  /*line wnlbNtaj.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 196
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 4:
				i = 1
				for y := range data {
					data[y] = data[y] -  /*line wnlbNtaj.go:1*/ byte(decryptKey^y)
				}
			case 2:
				i = 7
				data =  /*line wnlbNtaj.go:1*/ append(data, 45)
			case 6:
				data =  /*line wnlbNtaj.go:1*/ append(data, 35)
				i = 3
			case 7:
				data =  /*line wnlbNtaj.go:1*/ append(data, 51)
				i = 5
			case 0:
				i = 6
				data =  /*line wnlbNtaj.go:1*/ append(data, 39)
			case 5:
				data =  /*line wnlbNtaj.go:1*/ append(data, 51)
				i = 0
			case 3:
				i = 4
				data =  /*line wnlbNtaj.go:1*/ append(data, 47)
			}
		}
		return  /*line wnlbNtaj.go:1*/ string(data)
	}())
	if szAT31uU.Status != shim.OK {
		return "",  /*line PCyLeThz.go:1*/ fmt.Errorf( /*line F4z76uHT.go:1*/ func() string {
			seed :=  /*line F4z76uHT.go:1*/ byte(91)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line F4z76uHT.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/  /*line F4z76uHT.go:1*/ fnc(29)(25)(248)(229)(11)(29)(182)(56)(235)(79)(207)(248)(224)(23)(5)(161)(65)(11)(15)(20)(255)(243)(236)(11)(31)(185)(31)(20)(241)(23)(233)(23)(249)(231)(77)(220)(249)(253)(172)(117)(200)(1)(23)(233)(23)(249)(231)(77)(211)(227)(22)(233)(65)(144)(7)(40)(27)(170)(81)(247)(14)(229)(29)(182)(98)(129)(86)
			return  /*line F4z76uHT.go:1*/ string(data)
		}(), szAT31uU.Message)
	}

	return  /*line KtNl_2Oh.go:1*/ string(szAT31uU.Payload), nil
}

                              
func fvIQVpVM(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) (string, error) {
	if  /*line zUnwGOhl.go:1*/ len(ktsi1_SQ) != 1 {
		return "",  /*line QLaJ5Blz.go:1*/ fmt.Errorf( /*line qytw3qtq.go:1*/ func() string {
			key :=  /*line qytw3qtq.go:1*/ []byte("\xf0i'\u007f\xaf\x96\xc6\xf1\x05P>\u007fu\xf4\xf8\xa5\xdd.\xeb\xb0\f\x1c\xb4\x0e\u007fvFd)\xe4\xde0\xf0l\xdaE\xa3\x03\x91\bAo\\\x95\xa0e\x98L\xb3(\xb8\xe1D\x14\xe1]\xc8NV?\xb4i\x1b\xddC2")
			data :=  /*line qytw3qtq.go:1*/ []byte("9\u05ca\xee!\b+Typ\x9f\xf1\xdcie\nK\xa2^\xde,e\".\xee\xe8\xaaɛ\x04R\x9f\x10\xd5H\xbb\x12n\xf6(\x8e\xd4\xd0\xf6\x04\xc6\f\xadӍ0Q\xa9wU\xc66\xb5v\xa3\x15\xdd|P\xa8\xa6")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line qytw3qtq.go:1*/ string(data)
		}())
	}
	utUwSJgF := ktsi1_SQ[0]

	y32DtQHo := []string{ /*line g7MgDV0S.go:1*/ func() string {
		key :=  /*line g7MgDV0S.go:1*/ []byte("\xb7\xa7x:\x1eՁ\x94^˝\xe3}m\b\xca\xfb{\xf5")
		data :=  /*line g7MgDV0S.go:1*/ []byte("\xb0\xbe\xfc\x13G\x9f\xe0\xd0\x03\xa9\xc4m\xf5\x02n\x9fi\xea}")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line g7MgDV0S.go:1*/ string(data)
	}(), utUwSJgF}
	fWJzYwn4 :=  /*line THEjO0ks.go:1*/ make([][]byte,  /*line Ahak4Nbq.go:1*/ len(y32DtQHo))
	for gGz2tP6V, vChBf6MJ := range y32DtQHo {
		fWJzYwn4[gGz2tP6V] =  /*line yLShFsw1.go:1*/ []byte(vChBf6MJ)
	}

	szAT31uU :=  /*line E7iyok62.go:1*/ n4c7mRtG.InvokeChaincode( /*line nxwkV7rg.go:1*/ func() string {
		var data []byte
		i := 11
		decryptKey := 196
		for counter := 0; i != 6; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 4:
				data =  /*line nxwkV7rg.go:1*/ append(data, 153)
				i = 9
			case 1:
				i = 4
				data =  /*line nxwkV7rg.go:1*/ append(data, 155)
			case 11:
				i = 2
				data =  /*line nxwkV7rg.go:1*/ append(data, 160)
			case 0:
				i = 10
				data =  /*line nxwkV7rg.go:1*/ append(data, 174)
			case 3:
				i = 5
				data =  /*line nxwkV7rg.go:1*/ append(data, 165)
			case 9:
				i = 8
				data =  /*line nxwkV7rg.go:1*/ append(data, 169)
			case 8:
				data =  /*line nxwkV7rg.go:1*/ append(data, 151)
				i = 0
			case 10:
				i = 7
				data =  /*line nxwkV7rg.go:1*/ append(data, 159)
			case 2:
				data =  /*line nxwkV7rg.go:1*/ append(data, 153)
				i = 3
			case 5:
				i = 1
				data =  /*line nxwkV7rg.go:1*/ append(data, 147)
			case 7:
				for y := range data {
					data[y] = data[y] +  /*line nxwkV7rg.go:1*/ byte(decryptKey^y)
				}
				i = 6
			}
		}
		return  /*line nxwkV7rg.go:1*/ string(data)
	}(), fWJzYwn4,  /*line IpI4L5Nr.go:1*/ func() string {
		data :=  /*line IpI4L5Nr.go:1*/ []byte("ghofeh")
		positions := [...]byte{5, 2, 2, 4, 4, 5, 3, 1}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line IpI4L5Nr.go:1*/ byte(i) +  /*line IpI4L5Nr.go:1*/ byte(positions[i]^positions[i+1]) + 2
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line IpI4L5Nr.go:1*/ string(data)
	}())
	if szAT31uU.Status != shim.OK {
		return "",  /*line Fozl4VIV.go:1*/ fmt.Errorf( /*line neWxjRdX.go:1*/ func() string {
			data :=  /*line neWxjRdX.go:1*/ []byte("\x84\xad\xff\x86\xa0\xcb\xfb\xfa\x92\x8cq\xe4e\xa7yT\x96W\xbei\x1dc\x87dۀ\x84e\xe4\xd2da\xa5a\xaa\xb9\x86r\xa9\x0ee\xecaF\xbdt\xc2\x02\x94\xceo\xd0\xee\f\xa9.\xc7\u008b\x05\x1b \xaf\xea\t&_J \xf3s")
			positions := [...]byte{49, 24, 28, 60, 51, 57, 63, 41, 6, 43, 69, 65, 0, 34, 35, 28, 48, 52, 64, 8, 5, 20, 2, 63, 47, 53, 54, 8, 39, 4, 36, 56, 11, 46, 18, 4, 5, 53, 67, 53, 26, 66, 1, 29, 38, 4, 53, 67, 52, 63, 13, 32, 59, 69, 7, 6, 51, 15, 3, 53, 43, 7, 59, 63, 1, 44, 51, 9, 17, 6, 47, 46, 22, 62, 54, 22, 16, 1, 58, 58, 47, 57, 64, 20, 2, 67, 25, 55}
			for i := 0; i < 88; i += 2 {
				localKey :=  /*line neWxjRdX.go:1*/ byte(i) +  /*line neWxjRdX.go:1*/ byte(positions[i]^positions[i+1]) + 110
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line neWxjRdX.go:1*/ string(data)
		}(), szAT31uU.Message)
	}

	return  /*line nWRnK6YX.go:1*/ string(szAT31uU.Payload), nil
}

                            
func imxpv0H5(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) (string, error) {
	if  /*line Bv26qTSG.go:1*/ len(ktsi1_SQ) != 1 {
		return "",  /*line WTyHRNqw.go:1*/ fmt.Errorf( /*line n8k_m7rb.go:1*/ func() string {
			seed :=  /*line n8k_m7rb.go:1*/ byte(237)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line n8k_m7rb.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/  /*line n8k_m7rb.go:1*/ fnc(164)(255)(243)(236)(29)(254)(239)(26)(231)(90)(181)(251)(227)(18)(20)(232)(27)(228)(7)(85)(240)(137)(39)(80)(175)(29)(232)(17)(247)(92)(172)(235)(79)(215)(251)(230)(25)(228)(22)(169)(127)(212)(241)(23)(233)(23)(249)(231)(77)(223)(225)(10)(225)(6)(31)(227)(3)(23)(167)(74)(25)(229)(23)(254)(238)(13)
			return  /*line n8k_m7rb.go:1*/ string(data)
		}())
	}
	utUwSJgF := ktsi1_SQ[0]

	y32DtQHo := []string{ /*line XrCeFLVO.go:1*/ func() string {
		key :=  /*line XrCeFLVO.go:1*/ []byte("\x0f{\xbe!E\xae\x14v\xce\xec1\xbfм\x1c\x1e\xdc2\x89\xc4\xed")
		data :=  /*line XrCeFLVO.go:1*/ []byte("X\xea\xb6, \xc6M\ue4c80\x8c\x95\xbd[Q\x962˝z")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line XrCeFLVO.go:1*/ string(data)
	}(), utUwSJgF}
	fWJzYwn4 :=  /*line SMKFLQzS.go:1*/ make([][]byte,  /*line UKdoU7pw.go:1*/ len(y32DtQHo))
	for gGz2tP6V, vChBf6MJ := range y32DtQHo {
		fWJzYwn4[gGz2tP6V] =  /*line KDKBW68D.go:1*/ []byte(vChBf6MJ)
	}

	szAT31uU :=  /*line TH24TzOg.go:1*/ n4c7mRtG.InvokeChaincode( /*line EGA6zrnn.go:1*/ func() string {
		key :=  /*line EGA6zrnn.go:1*/ []byte("\xab\xc4@\xc7\xed1L\xdbv\xf2")
		data :=  /*line EGA6zrnn.go:1*/ []byte("\x18)\xb4(Q\x92\xc0<\xe9U")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line EGA6zrnn.go:1*/ string(data)
	}(), fWJzYwn4,  /*line pXxBJHhM.go:1*/ func() string {
		var data []byte
		i := 3
		decryptKey := 237
		for counter := 0; i != 4; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				for y := range data {
					data[y] = data[y] +  /*line pXxBJHhM.go:1*/ byte(decryptKey^y)
				}
				i = 4
			case 2:
				data =  /*line pXxBJHhM.go:1*/ append(data, 122)
				i = 6
			case 3:
				data =  /*line pXxBJHhM.go:1*/ append(data, 124)
				i = 7
			case 7:
				data =  /*line pXxBJHhM.go:1*/ append(data, 130)
				i = 0
			case 6:
				i = 5
				data =  /*line pXxBJHhM.go:1*/ append(data, 114)
			case 0:
				i = 2
				data =  /*line pXxBJHhM.go:1*/ append(data, 134)
			case 5:
				data =  /*line pXxBJHhM.go:1*/ append(data, 126)
				i = 1
			}
		}
		return  /*line pXxBJHhM.go:1*/ string(data)
	}())
	if szAT31uU.Status != shim.OK {
		return "",  /*line gYbh4gvP.go:1*/ fmt.Errorf( /*line uyziviHj.go:1*/ func() string {
			key :=  /*line uyziviHj.go:1*/ []byte("\xbc\x8c\xcf\xc0\xf73\x18\xa6\xbd\xbe\x8a\xa4\x1fq\xa7\xa2\xad\x1f\xfd\x06\xe7Z\xad\xa3c\xab|{\xedoX\xa7rB\x93T\xd4\x14\xcd(B{\x82\x9b;\xb0,g0\xf2\u007f.\x1a\xe0\xafcyz\xe4\xe3q7\x03<")
			data :=  /*line uyziviHj.go:1*/ []byte("\xfa\xed\xa6\xac\x92W8\xd2Ҟ\xfb\xd1z\x03ނ\xcew\x9co\x899\xc2\xc7\x06\x8b1\x1e\x99\x0e<\xc6\x06#\xb32\xbbf\xedc'\x02\xf5\xf4I\xd4x\x06W\xdc_iu\x94\x8f\x06\v\b\x8b\x91K\x17&O")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line uyziviHj.go:1*/ string(data)
		}(), szAT31uU.Message)
	}

	return  /*line ugXAVonA.go:1*/ string(szAT31uU.Payload), nil
}

                          
func grz7gdfX(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) (string, error) {
	if  /*line BT6SsErL.go:1*/ len(ktsi1_SQ) != 1 {
		return "",  /*line RraTIJo3.go:1*/ fmt.Errorf( /*line ZziPMa6i.go:1*/ func() string {
			key :=  /*line ZziPMa6i.go:1*/ []byte("\xbbu\xcc?\xab\xe8\x00A\x98\xb9Z\xdf\t'bp\xa4\xd4\xd9\xf8\x90\xca\xd4\x01\xf6\xe5\xf3tW\x13d\x82A\xf6\x04TtJ\x1a\xc02\xff\xb7d\xc5\b\xc7P\x89jسD\xa4t\xe5OE\xb5\n\xfe\x9b\xc9u\xbc\x1d")
			data :=  /*line ZziPMa6i.go:1*/ []byte("\x8e\xf9\x970Ǌe\"\xdcg\a\x93^N\v\xf5ʠ\x9a6\x90\u007f\x9a\x1fy\x8dq\xf1\x1b\r\x10\xed\xdfsj\"\xfb!K`\x1bf\xbd\xfd\x9fY\xad\x11\x97\xfb\xa0\xbd!\xbf\x00\x84\x1f\"kZc٘\xfe\xa9W")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line ZziPMa6i.go:1*/ string(data)
		}())
	}
	utUwSJgF := ktsi1_SQ[0]

	y32DtQHo := []string{ /*line ZtMj62QI.go:1*/ func() string {
		data :=  /*line ZtMj62QI.go:1*/ []byte("g\xddt\x80\x10\xf5Jޕ\xec\nY\x8ddpo\xecxt")
		positions := [...]byte{17, 8, 16, 12, 7, 1, 6, 11, 16, 5, 12, 9, 4, 10, 9, 10, 8, 4, 17, 9, 11, 11, 3, 6, 6, 8}
		for i := 0; i < 26; i += 2 {
			localKey :=  /*line ZtMj62QI.go:1*/ byte(i) +  /*line ZtMj62QI.go:1*/ byte(positions[i]^positions[i+1]) + 111
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line ZtMj62QI.go:1*/ string(data)
	}(), utUwSJgF}
	fWJzYwn4 :=  /*line F7syRuuT.go:1*/ make([][]byte,  /*line J5Fmoocv.go:1*/ len(y32DtQHo))
	for gGz2tP6V, vChBf6MJ := range y32DtQHo {
		fWJzYwn4[gGz2tP6V] =  /*line vRB5CD9B.go:1*/ []byte(vChBf6MJ)
	}

	szAT31uU :=  /*line z4YAORj8.go:1*/ n4c7mRtG.InvokeChaincode( /*line xPmUCLtw.go:1*/ func() string {
		var data []byte
		i := 4
		decryptKey := 207
		for counter := 0; i != 8; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 11:
				i = 0
				data =  /*line xPmUCLtw.go:1*/ append(data, 227)
			case 5:
				i = 7
				data =  /*line xPmUCLtw.go:1*/ append(data, 201)
			case 0:
				i = 3
				data =  /*line xPmUCLtw.go:1*/ append(data, 209)
			case 1:
				i = 5
				data =  /*line xPmUCLtw.go:1*/ append(data, 216)
			case 3:
				i = 6
				data =  /*line xPmUCLtw.go:1*/ append(data, 205)
			case 2:
				data =  /*line xPmUCLtw.go:1*/ append(data, 205)
				i = 1
			case 9:
				i = 11
				data =  /*line xPmUCLtw.go:1*/ append(data, 211)
			case 4:
				i = 9
				data =  /*line xPmUCLtw.go:1*/ append(data, 218)
			case 7:
				i = 8
				for y := range data {
					data[y] = data[y] +  /*line xPmUCLtw.go:1*/ byte(decryptKey^y)
				}
			case 10:
				data =  /*line xPmUCLtw.go:1*/ append(data, 223)
				i = 2
			case 6:
				i = 10
				data =  /*line xPmUCLtw.go:1*/ append(data, 203)
			}
		}
		return  /*line xPmUCLtw.go:1*/ string(data)
	}(), fWJzYwn4,  /*line b_xJGlMb.go:1*/ func() string {
		key :=  /*line b_xJGlMb.go:1*/ []byte("P\x1f\xd0\xdb\xebi")
		data :=  /*line b_xJGlMb.go:1*/ []byte("\x17M\x9f\x87v\x03")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line b_xJGlMb.go:1*/ string(data)
	}())
	if szAT31uU.Status != shim.OK {
		return "",  /*line nq59oRKt.go:1*/ fmt.Errorf( /*line PdTKzote.go:1*/ func() string {
			seed :=  /*line PdTKzote.go:1*/ byte(57)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line PdTKzote.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/  /*line PdTKzote.go:1*/ fnc(127)(217)(248)(229)(11)(29)(182)(56)(235)(79)(207)(248)(224)(23)(5)(161)(65)(11)(15)(20)(255)(243)(236)(11)(31)(185)(31)(20)(241)(23)(233)(23)(249)(231)(77)(220)(249)(253)(172)(93)(251)(244)(244)(23)(230)(27)(228)(90)(238)(251)(216)(251)(170)(81)(247)(14)(229)(29)(182)(98)(129)(86)
			return  /*line PdTKzote.go:1*/ string(data)
		}(), szAT31uU.Message)
	}

	return  /*line p0dixnjq.go:1*/ string(szAT31uU.Payload), nil
}

                                    
func wjnibpdu(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) (string, error) {
	if  /*line ihANI7_T.go:1*/ len(ktsi1_SQ) != 11 {
		return "",  /*line lR0r9XeG.go:1*/ fmt.Errorf( /*line NK01xMVD.go:1*/ func() string {
			seed :=  /*line NK01xMVD.go:1*/ byte(227)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line NK01xMVD.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/  /*line NK01xMVD.go:1*/ fnc(170)(227)(19)(236)(29)(254)(239)(26)(231)(90)(181)(251)(227)(18)(20)(232)(27)(228)(7)(85)(240)(137)(39)(80)(175)(29)(232)(17)(247)(92)(172)(235)(79)(215)(251)(230)(25)(228)(22)(169)(94)(255)(232)(16)(226)(27)(164)(77)(13)(242)(17)(230)(31)(227)(3)(23)(167)(31)(124)(233)(196)(23)(225)(27)(236)(6)
			return  /*line NK01xMVD.go:1*/ string(data)
		}())
	}
	utUwSJgF := ktsi1_SQ[0]
	bRd2au7r := ktsi1_SQ[1]
	bSzIZJnz := ktsi1_SQ[2]
	dJwgW2jL := ktsi1_SQ[3]
	hSHnto8d := ktsi1_SQ[4]
	xhl9F7R0 := ktsi1_SQ[5]
	tIeFhn0D := ktsi1_SQ[6]
	olYMzSWK := ktsi1_SQ[7]
	ezgSKu0t := ktsi1_SQ[8]
	iVUeYYCJ := ktsi1_SQ[9]
	deTJzFce := ktsi1_SQ[10]

	y32DtQHo := []string{ /*line EqFxFG2C.go:1*/ func() string {
		seed :=  /*line EqFxFG2C.go:1*/ byte(193)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line EqFxFG2C.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/  /*line EqFxFG2C.go:1*/ fnc(172)(2)(255)(251)(11)(251)(3)(207)(34)(0)(2)(14)(0)(223)(19)(12)(4)(240)(14)(1)
		return  /*line EqFxFG2C.go:1*/ string(data)
	}(), utUwSJgF, bRd2au7r, bSzIZJnz, dJwgW2jL, hSHnto8d, xhl9F7R0, tIeFhn0D, olYMzSWK, ezgSKu0t, iVUeYYCJ, deTJzFce}
	g2qanrhL :=  /*line Ldh0b7F2.go:1*/ make([][]byte,  /*line umOlk8Ql.go:1*/ len(y32DtQHo))
	for gGz2tP6V, vChBf6MJ := range y32DtQHo {
		g2qanrhL[gGz2tP6V] =  /*line DE2_tlGI.go:1*/ []byte(vChBf6MJ)
	}

	szAT31uU :=  /*line khk9dWc5.go:1*/ n4c7mRtG.InvokeChaincode( /*line SwMmbkiX.go:1*/ func() string {
		var data []byte
		i := 0
		decryptKey := 69
		for counter := 0; i != 7; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 5:
				i = 7
				for y := range data {
					data[y] = data[y] +  /*line SwMmbkiX.go:1*/ byte(decryptKey^y)
				}
			case 1:
				data =  /*line SwMmbkiX.go:1*/ append(data, 28)
				i = 4
			case 6:
				i = 3
				data =  /*line SwMmbkiX.go:1*/ append(data, 35)
			case 2:
				i = 1
				data =  /*line SwMmbkiX.go:1*/ append(data, 25)
			case 0:
				i = 6
				data =  /*line SwMmbkiX.go:1*/ append(data, 31)
			case 4:
				data =  /*line SwMmbkiX.go:1*/ append(data, 42)
				i = 5
			case 3:
				data =  /*line SwMmbkiX.go:1*/ append(data, 24)
				i = 2
			}
		}
		return  /*line SwMmbkiX.go:1*/ string(data)
	}(), g2qanrhL,  /*line RYzJZTX1.go:1*/ func() string {
		var data []byte
		i := 0
		decryptKey := 213
		for counter := 0; i != 4; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 4
				for y := range data {
					data[y] = data[y] +  /*line RYzJZTX1.go:1*/ byte(decryptKey^y)
				}
			case 6:
				data =  /*line RYzJZTX1.go:1*/ append(data, 142)
				i = 2
			case 3:
				i = 7
				data =  /*line RYzJZTX1.go:1*/ append(data, 154)
			case 2:
				data =  /*line RYzJZTX1.go:1*/ append(data, 154)
				i = 1
			case 5:
				i = 3
				data =  /*line RYzJZTX1.go:1*/ append(data, 150)
			case 7:
				data =  /*line RYzJZTX1.go:1*/ append(data, 142)
				i = 6
			case 0:
				i = 5
				data =  /*line RYzJZTX1.go:1*/ append(data, 144)
			}
		}
		return  /*line RYzJZTX1.go:1*/ string(data)
	}())
	if szAT31uU.Status != shim.OK {
		return "",  /*line uL5NVH9B.go:1*/ fmt.Errorf( /*line YbSAfRWR.go:1*/ func() string {
			data :=  /*line YbSAfRWR.go:1*/ []byte("f\x87ى\v\xaeh\x90o\xd7\xcenvo~V\xb1cx\x84\xa1\xcb%\xb8\x04\xad l\x8d\xa5ge\xfa.\"Go\xc9\xc2e\xa4\xc2K\xa2\xd0w\xb4\xc2")
			positions := [...]byte{41, 3, 44, 47, 34, 46, 10, 0, 6, 22, 38, 4, 37, 29, 28, 18, 7, 45, 47, 16, 40, 14, 0, 7, 29, 43, 34, 42, 25, 43, 34, 33, 14, 1, 21, 18, 25, 2, 3, 42, 16, 7, 16, 15, 23, 23, 18, 19, 38, 24, 19, 2, 32, 19, 9, 19, 45, 9, 4, 23, 38, 40, 24, 19, 5, 25, 20, 44}
			for i := 0; i < 68; i += 2 {
				localKey :=  /*line YbSAfRWR.go:1*/ byte(i) +  /*line YbSAfRWR.go:1*/ byte(positions[i]^positions[i+1]) + 237
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line YbSAfRWR.go:1*/ string(data)
		}(), szAT31uU.Payload)
	}
	return  /*line zVbSGWmp.go:1*/ func() string {
		data :=  /*line zVbSGWmp.go:1*/ []byte("\x12n\x97o(\x98\xa8\xactar\xbbs")
		positions := [...]byte{4, 0, 4, 2, 0, 2, 5, 0, 7, 11, 2, 6, 6, 7, 6, 6}
		for i := 0; i < 16; i += 2 {
			localKey :=  /*line zVbSGWmp.go:1*/ byte(i) +  /*line zVbSGWmp.go:1*/ byte(positions[i]^positions[i+1]) + 36
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line zVbSGWmp.go:1*/ string(data)
	}(), nil
}

                                          
func t94C_D8c(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) (string, error) {
	if  /*line GEQ3O9G5.go:1*/ len(ktsi1_SQ) != 9 {
		return "",  /*line IMFh5gpn.go:1*/ fmt.Errorf( /*line VoQw9kRM.go:1*/ func() string {
			key :=  /*line VoQw9kRM.go:1*/ []byte("\x82\x9c\xff\xe5\xf6\x93\x06:\x9b\xdd\x0f\x1fN\xea\x0fZ\x15=\x13%~\xe41\x1f\xb1\xcco\xc8*މQVW\xf4g\xaf\x17\x9a\x8d\u05f9\x99\x058\xc2t%\xbffe;\x96\xc7$v\xdb\xdc\xc5h\xf2\xd2Ƴ\xa9V=\xf9\x89\x1a\xe6\xa2(x\x84\x80\xf9\x9b\xa2\xc8\xfa,J\xfe")
			data :=  /*line VoQw9kRM.go:1*/ []byte("\xc7\xd2d\x8a|\xdf_)\xd9CRS\x19\x8b^\vY7`\t\xa2e=\x01\xbe\xa6\xf5\x9dHB\xeb\x1e\xca\x12z\x0f\xc0T˓\x95\xb6\xceb-\xb0\xacA\xb0\f\xbb7ϯK\xf5\x8a\x88[\xf9q\x91\x9f\xc0\xca\xca(\u007f\xe7K}\xd2A\xf6\xe3\xa0@\x85ԙrI\x1bu")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line VoQw9kRM.go:1*/ string(data)
		}())
	}

	utUwSJgF := ktsi1_SQ[0]
	bRd2au7r := ktsi1_SQ[1]
	bSzIZJnz := ktsi1_SQ[2]
	dJwgW2jL := ktsi1_SQ[3]
	xdlIGP1Q := ktsi1_SQ[4]
	tIeFhn0D := ktsi1_SQ[5]
	ezgSKu0t := ktsi1_SQ[6]
	iVUeYYCJ := ktsi1_SQ[7]
	deTJzFce := ktsi1_SQ[8]

	y32DtQHo := []string{ /*line p5QmszlW.go:1*/ func() string {
		data :=  /*line p5QmszlW.go:1*/ []byte("m\xa8\"\xc8:\xdb\xcd\xdf+v\xc7k4.A#\xd8Ks\xc7U\xc8\xd1\xcd")
		positions := [...]byte{1, 17, 10, 22, 2, 17, 2, 13, 4, 8, 2, 7, 17, 15, 3, 12, 1, 5, 19, 3, 6, 23, 16, 21, 17, 17, 8, 4}
		for i := 0; i < 28; i += 2 {
			localKey :=  /*line p5QmszlW.go:1*/ byte(i) +  /*line p5QmszlW.go:1*/ byte(positions[i]^positions[i+1]) + 128
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line p5QmszlW.go:1*/ string(data)
	}(), utUwSJgF, bRd2au7r, bSzIZJnz, dJwgW2jL, xdlIGP1Q, tIeFhn0D, ezgSKu0t, iVUeYYCJ, deTJzFce}
	g2qanrhL :=  /*line G6CIzN0H.go:1*/ make([][]byte,  /*line GakZ_FuI.go:1*/ len(y32DtQHo))
	for gGz2tP6V, vChBf6MJ := range y32DtQHo {
		g2qanrhL[gGz2tP6V] =  /*line YTlJZcAh.go:1*/ []byte(vChBf6MJ)
	}

	szAT31uU :=  /*line L4411QBF.go:1*/ n4c7mRtG.InvokeChaincode( /*line Xh7n5jIU.go:1*/ func() string {
		seed :=  /*line Xh7n5jIU.go:1*/ byte(137)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Xh7n5jIU.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line Xh7n5jIU.go:1*/  /*line Xh7n5jIU.go:1*/  /*line Xh7n5jIU.go:1*/  /*line Xh7n5jIU.go:1*/  /*line Xh7n5jIU.go:1*/  /*line Xh7n5jIU.go:1*/ fnc(245)(237)(210)(164)(70)(153)
		return  /*line Xh7n5jIU.go:1*/ string(data)
	}(), g2qanrhL,  /*line DzsR4yrk.go:1*/ func() string {
		fullData :=  /*line DzsR4yrk.go:1*/ []byte("lG\xf3o\xe3\x03\x96\xd6w&\bU")
		var data []byte
		data =  /*line DzsR4yrk.go:1*/ append(data, fullData[3]-fullData[10], fullData[4]-fullData[8], fullData[5]+fullData[0], fullData[11]-fullData[2], fullData[1]^fullData[9], fullData[7]+fullData[6])
		return  /*line DzsR4yrk.go:1*/ string(data)
	}())
	if szAT31uU.Status != shim.OK {
		return "",  /*line wezujxFd.go:1*/ fmt.Errorf( /*line CnuelPKC.go:1*/ func() string {
			fullData :=  /*line CnuelPKC.go:1*/ []byte("\x87\x15\xe4\xa7\xe2\xe1T\x15\x17_<J\xe5\x92\xf7<\xec\x80\xd1\xf7\xfaF\xc6\x10X\xa8\x93\xce^\x10\xb5\xa4\xf2E[\x9e\x80\xae\xf860\xb2\xb1J\x18j\xa4\xc97\xab˵\x13\xdai~\xaa\x83[ޔP6Y\x14\xed\x9cv;\xc3-\xe9R\x03\x1cD\xfc\xc7\xe9\x17\x85\xe6\xe8\x9f5\x1cY\xacI\x88qŰ4&\x9a")
			var data []byte
			data =  /*line CnuelPKC.go:1*/ append(data, fullData[65]+fullData[63], fullData[51]-fullData[6], fullData[80]^fullData[16], fullData[25]-fullData[15], fullData[4]+fullData[57], fullData[38]-fullData[60], fullData[55]^fullData[28], fullData[72]-fullData[59], fullData[27]-fullData[9], fullData[88]^fullData[54], fullData[17]-fullData[8], fullData[81]^fullData[89], fullData[3]^fullData[18], fullData[24]-fullData[71], fullData[85]-fullData[42], fullData[20]^fullData[83], fullData[47]^fullData[78], fullData[2]^fullData[0], fullData[53]^fullData[41], fullData[68]+fullData[94], fullData[91]+fullData[46], fullData[69]+fullData[49], fullData[26]-fullData[40], fullData[50]^fullData[31], fullData[33]-fullData[5], fullData[12]+fullData[36], fullData[79]-fullData[14], fullData[32]^fullData[35], fullData[86]^fullData[62], fullData[70]^fullData[43], fullData[34]^fullData[10], fullData[52]-fullData[37], fullData[93]^fullData[21], fullData[66]+fullData[13], fullData[1]^fullData[84], fullData[61]+fullData[19], fullData[58]+fullData[64], fullData[73]+fullData[90], fullData[74]-fullData[76], fullData[92]+fullData[30], fullData[45]^fullData[44], fullData[87]+fullData[22], fullData[39]-fullData[77], fullData[82]^fullData[95], fullData[11]-fullData[29], fullData[56]+fullData[67], fullData[23]+fullData[7], fullData[75]^fullData[48])
			return  /*line CnuelPKC.go:1*/ string(data)
		}(), szAT31uU.Payload)
	}
	return  /*line Bw_TmL0O.go:1*/ func() string {
		data :=  /*line Bw_TmL0O.go:1*/ []byte("i\xdc$\xcc\xd1\xd5G->\xd3r0\xe8")
		positions := [...]byte{2, 4, 8, 11, 5, 12, 4, 9, 3, 3, 1, 8, 6, 7, 6, 5, 7, 11}
		for i := 0; i < 18; i += 2 {
			localKey :=  /*line Bw_TmL0O.go:1*/ byte(i) +  /*line Bw_TmL0O.go:1*/ byte(positions[i]^positions[i+1]) + 85
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line Bw_TmL0O.go:1*/ string(data)
	}(), nil
}

                                          
func baNjNl2z(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) (string, error) {
	if  /*line yJxLHYVq.go:1*/ len(ktsi1_SQ) != 12 {
		return "",  /*line E2Q04mKR.go:1*/ fmt.Errorf( /*line U_Cmeqlb.go:1*/ func() string {
			key :=  /*line U_Cmeqlb.go:1*/ []byte("\x1c#\xb44\xbc\x1cD)\xbb9\xbe\x18-\xcc\xc1\x15AI\x88\xf4\xb4\xd2!\x91\xd3w\xae\xd9\xe3\xfd[3\x0e\tpR\x87\x8aKe\xd3Ԇp6&\x83RN\x8f\x1cA}u\x03\xc2q\x05\xc9{\xe4\x95\xdbD\x95\xbadz喆L\xc9];\x0eG\xa2\xc3")
			data :=  /*line U_Cmeqlb.go:1*/ []byte("e\x91\x17\xa3.\x8e\xa9\x8c/Y\x1f\x8a\x94A.z\xaf\xbd\xfb\x14\x1aA\x93\xb19\xec\x1c<Wfʡ.|\xd5\xc6\xd9\xef\xc1\xd4>9걙\x89\xe8\xc5\xc1ގ\xa8\xf0\xbeq(\xe0y8\xe7S\xfcB\xa9\a脿]\x06\xeb\xaf=Ʃug\xd3\xf5")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line U_Cmeqlb.go:1*/ string(data)
		}())
	}

	utUwSJgF := ktsi1_SQ[0]
	bRd2au7r := ktsi1_SQ[1]
	kvhgzQIY := ktsi1_SQ[2]
	xdlIGP1Q := ktsi1_SQ[3]
	pbLLcfzS := ktsi1_SQ[4]
	hOgpzUCn := ktsi1_SQ[5]
	ipXT6jpu := ktsi1_SQ[6]
	wSxUCaJ0 := ktsi1_SQ[7]
	v5vcwK31 := ktsi1_SQ[8]
	ezgSKu0t := ktsi1_SQ[9]
	iVUeYYCJ := ktsi1_SQ[10]
	deTJzFce := ktsi1_SQ[11]

	y32DtQHo := []string{ /*line zL14yLK1.go:1*/ func() string {
		fullData :=  /*line zL14yLK1.go:1*/ []byte("\xd5\x10\xf9B:5\x900\x10\xe4B\xa0\x029\x00lR)\x80qt\x9a\xb0±\xd5\r\xb7\xb7\x05Ǧ=RV<\x99\xe2\x1fGQ\xf4\xd2E\x9d*")
		var data []byte
		data =  /*line zL14yLK1.go:1*/ append(data, fullData[31]+fullData[30], fullData[19]-fullData[12], fullData[13]+fullData[5], fullData[36]-fullData[7], fullData[41]^fullData[18], fullData[33]^fullData[32], fullData[34]-fullData[9], fullData[8]^fullData[10], fullData[39]-fullData[37], fullData[35]+fullData[4], fullData[20]-fullData[29], fullData[22]-fullData[43], fullData[28]-fullData[16], fullData[0]^fullData[24], fullData[1]^fullData[40], fullData[14]-fullData[44], fullData[2]^fullData[21], fullData[25]+fullData[6], fullData[38]^fullData[15], fullData[45]-fullData[27], fullData[3]+fullData[26], fullData[42]^fullData[11], fullData[17]-fullData[23])
		return  /*line zL14yLK1.go:1*/ string(data)
	}(), utUwSJgF, bRd2au7r, kvhgzQIY, xdlIGP1Q, pbLLcfzS, hOgpzUCn, ipXT6jpu, wSxUCaJ0, v5vcwK31, ezgSKu0t, iVUeYYCJ, deTJzFce}
	g2qanrhL :=  /*line k0ldkiw7.go:1*/ make([][]byte,  /*line scLIuq0r.go:1*/ len(y32DtQHo))
	for gGz2tP6V, vChBf6MJ := range y32DtQHo {
		g2qanrhL[gGz2tP6V] =  /*line ZV1DmybJ.go:1*/ []byte(vChBf6MJ)
	}

	szAT31uU :=  /*line xyftG8ZX.go:1*/ n4c7mRtG.InvokeChaincode( /*line K4zawo8b.go:1*/ func() string {
		fullData :=  /*line K4zawo8b.go:1*/ []byte("\xc0\xd7\xf6]=:eά\xa7[\xc0")
		var data []byte
		data =  /*line K4zawo8b.go:1*/ append(data, fullData[0]^fullData[8], fullData[4]-fullData[7], fullData[5]^fullData[3], fullData[9]+fullData[11], fullData[10]-fullData[2], fullData[1]-fullData[6])
		return  /*line K4zawo8b.go:1*/ string(data)
	}(), g2qanrhL,  /*line rz0KoI7k.go:1*/ func() string {
		fullData :=  /*line rz0KoI7k.go:1*/ []byte("I\x92:/j+\v\x04h&\x96U")
		var data []byte
		data =  /*line rz0KoI7k.go:1*/ append(data, fullData[10]-fullData[3], fullData[7]^fullData[8], fullData[11]^fullData[2], fullData[0]^fullData[5], fullData[6]^fullData[4], fullData[1]-fullData[9])
		return  /*line rz0KoI7k.go:1*/ string(data)
	}())
	if szAT31uU.Status != shim.OK {
		return "",  /*line _KCz0DOF.go:1*/ fmt.Errorf( /*line BNHd4qQo.go:1*/ func() string {
			data :=  /*line BNHd4qQo.go:1*/ []byte("\x15a{2s\xb1U\xf6\u007f 5\x01y<Ge$c\xb2a\x1en\xd6\xe6de\xb5\xf8ov\xc66O\xe7 \xf8\xceb9err\xfcr\x16\xdd\xc68")
			positions := [...]byte{38, 32, 13, 31, 38, 2, 16, 45, 26, 8, 38, 5, 23, 37, 18, 45, 3, 35, 0, 12, 18, 22, 44, 14, 47, 7, 16, 6, 4, 37, 20, 27, 3, 11, 29, 4, 10, 0, 4, 42, 33, 7, 30, 36, 8, 37, 23, 11, 46, 20, 38, 18}
			for i := 0; i < 52; i += 2 {
				localKey :=  /*line BNHd4qQo.go:1*/ byte(i) +  /*line BNHd4qQo.go:1*/ byte(positions[i]^positions[i+1]) + 69
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line BNHd4qQo.go:1*/ string(data)
		}(), szAT31uU.Payload)
	}
	return  /*line dIgVzfvf.go:1*/ func() string {
		fullData :=  /*line dIgVzfvf.go:1*/ []byte("\x02Ӗ,E\r\x9d˾X\x9e\xd1U\xdd)\xa3%l?t\xf5\x86'D\xfbN")
		var data []byte
		data =  /*line dIgVzfvf.go:1*/ append(data, fullData[22]^fullData[25], fullData[17]+fullData[0], fullData[7]-fullData[12], fullData[24]+fullData[19], fullData[18]+fullData[3], fullData[9]+fullData[5], fullData[10]^fullData[8], fullData[13]+fullData[2], fullData[15]+fullData[11], fullData[23]^fullData[16], fullData[4]-fullData[1], fullData[6]-fullData[14], fullData[21]^fullData[20])
		return  /*line dIgVzfvf.go:1*/ string(data)
	}(), nil
}

                                    
func o8Z_JkaH(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) (string, error) {
	if  /*line DIID0Lyd.go:1*/ len(ktsi1_SQ) != 14 {
		return "",  /*line AzecHaEP.go:1*/ fmt.Errorf( /*line FYD1Dpuk.go:1*/ func() string {
			seed :=  /*line FYD1Dpuk.go:1*/ byte(16)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line FYD1Dpuk.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/  /*line FYD1Dpuk.go:1*/ fnc(57)(37)(245)(12)(3)(0)(243)(254)(17)(172)(65)(17)(245)(14)(248)(248)(9)(6)(255)(187)(242)(41)(37)(178)(79)(3)(242)(1)(13)(174)(84)(251)(177)(73)(5)(8)(249)(252)(250)(187)(76)(3)(248)(0)(254)(13)(174)(69)(19)(248)(245)(254)(17)(245)(5)(249)(185)(17)(3)(236)(86)(235)(11)(9)(240)(14)
			return  /*line FYD1Dpuk.go:1*/ string(data)
		}())
	}
	utUwSJgF := ktsi1_SQ[0]
	bRd2au7r := ktsi1_SQ[1]
	mdzhNJEk := ktsi1_SQ[2]
	hSHnto8d := ktsi1_SQ[3]
	xhl9F7R0 := ktsi1_SQ[4]
	pbLLcfzS := ktsi1_SQ[5]
	hOgpzUCn := ktsi1_SQ[6]
	ipXT6jpu := ktsi1_SQ[7]
	wSxUCaJ0 := ktsi1_SQ[8]
	v5vcwK31 := ktsi1_SQ[9]
	olYMzSWK := ktsi1_SQ[10]
	ezgSKu0t := ktsi1_SQ[11]
	iVUeYYCJ := ktsi1_SQ[12]
	deTJzFce := ktsi1_SQ[13]

	y32DtQHo := []string{ /*line vujN1qHD.go:1*/ func() string {
		fullData :=  /*line vujN1qHD.go:1*/ []byte("R\xf7\x13\x06h2\xf9tu\x00\b\x9f,\x90$v\x15v\xf1\x84\xf6{B\xff\x8a\xb6\xce\xfd\xe5\x04Ci\x88\xb2\xb3\xfc\x84\xa0\x97\v\x92\xfdA\xdeK\xc7\xe8\xa6")
		var data []byte
		data =  /*line vujN1qHD.go:1*/ append(data, fullData[32]+fullData[28], fullData[33]-fullData[30], fullData[6]^fullData[38], fullData[9]+fullData[31], fullData[37]-fullData[12], fullData[8]-fullData[3], fullData[29]-fullData[40], fullData[43]^fullData[11], fullData[15]^fullData[16], fullData[44]-fullData[46], fullData[27]+fullData[4], fullData[41]-fullData[24], fullData[7]+fullData[23], fullData[17]-fullData[14], fullData[39]-fullData[47], fullData[25]^fullData[45], fullData[36]^fullData[18], fullData[2]+fullData[0], fullData[10]^fullData[21], fullData[22]+fullData[5], fullData[35]^fullData[34], fullData[19]^fullData[20], fullData[13]^fullData[1], fullData[42]-fullData[26])
		return  /*line vujN1qHD.go:1*/ string(data)
	}(), utUwSJgF, bRd2au7r, mdzhNJEk, hSHnto8d, xhl9F7R0, pbLLcfzS, hOgpzUCn, ipXT6jpu, wSxUCaJ0, v5vcwK31, olYMzSWK, ezgSKu0t, iVUeYYCJ, deTJzFce}
	g2qanrhL :=  /*line JS6zf26r.go:1*/ make([][]byte,  /*line iIzkYXXY.go:1*/ len(y32DtQHo))
	for gGz2tP6V, vChBf6MJ := range y32DtQHo {
		g2qanrhL[gGz2tP6V] =  /*line YjPwXXNy.go:1*/ []byte(vChBf6MJ)
	}

	szAT31uU :=  /*line wj82RrKG.go:1*/ n4c7mRtG.InvokeChaincode( /*line VTOJyLQO.go:1*/ func() string {
		var data []byte
		i := 4
		decryptKey := 14
		for counter := 0; i != 6; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				for y := range data {
					data[y] = data[y] -  /*line VTOJyLQO.go:1*/ byte(decryptKey^y)
				}
				i = 6
			case 3:
				i = 7
				data =  /*line VTOJyLQO.go:1*/ append(data, 135)
			case 7:
				data =  /*line VTOJyLQO.go:1*/ append(data, 149)
				i = 1
			case 5:
				data =  /*line VTOJyLQO.go:1*/ append(data, 150)
				i = 2
			case 0:
				i = 3
				data =  /*line VTOJyLQO.go:1*/ append(data, 140)
			case 4:
				i = 5
				data =  /*line VTOJyLQO.go:1*/ append(data, 146)
			case 2:
				i = 0
				data =  /*line VTOJyLQO.go:1*/ append(data, 139)
			}
		}
		return  /*line VTOJyLQO.go:1*/ string(data)
	}(), g2qanrhL,  /*line biTd3TZK.go:1*/ func() string {
		data :=  /*line biTd3TZK.go:1*/ []byte("g\xc0\xc1Ǝl")
		positions := [...]byte{1, 1, 4, 3, 2, 4, 1, 4}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line biTd3TZK.go:1*/ byte(i) +  /*line biTd3TZK.go:1*/ byte(positions[i]^positions[i+1]) + 203
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line biTd3TZK.go:1*/ string(data)
	}())
	if szAT31uU.Status != shim.OK {
		return "",  /*line JG8Mq_x3.go:1*/ fmt.Errorf( /*line LkiLKHBU.go:1*/ func() string {
			key :=  /*line LkiLKHBU.go:1*/ []byte("<L\xc9_\x81\x9c\xb8\x87\x1f!\x8e\x85\x1c\u007fj\xa7\x9d%P9\xe6\x1b\xfcJ_\x99\x1eM\x8229@ _¾\x00\xf6\a\x13`ǧ\xe2Tal\x93\xddw*/!-\xee\x93!\xc4I\x10\xa62\xde")
			data :=  /*line LkiLKHBU.go:1*/ []byte("\x82\xad2\xcb\xe6\x00\xd8\xfb\x8eA\xf7\xf3\x92\xee\xd5\f\xbd\x88\xb8\x9aO\x89_\xb9\xc3\xfe>\xb9\U00059825\x92\u007f(-r\x16v\x85\xc7\xe7\bE\xb7\xc6\xdf\x06\v\x97q\x9e\x95MS\x05\x933\xbbJ\xc6WQ")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line LkiLKHBU.go:1*/ string(data)
		}(), szAT31uU.Message)
	}
	return  /*line eNUoQzVj.go:1*/ func() string {
		key :=  /*line eNUoQzVj.go:1*/ []byte("N\x02\xa7\x81\xc6\xe7ny\xa6Y\x13lU")
		data :=  /*line eNUoQzVj.go:1*/ []byte("\xb7p\x1d\xf01L\x8e\xec\x1a\xba\x85\xe0\xc8")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line eNUoQzVj.go:1*/ string(data)
	}(), nil
}

                                    
func _YWFxzXg(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) (string, error) {
	if  /*line Mkyj7n5L.go:1*/ len(ktsi1_SQ) != 5 {
		return "",  /*line xNpvnJNJ.go:1*/ fmt.Errorf( /*line kJrzMxS9.go:1*/ func() string {
			data :=  /*line kJrzMxS9.go:1*/ []byte(".\xb6\x89\xcb\xec\x83ectYDrx\x95\xb2~\x19\x13ET `n\xd1\xcbiwW\x8ei\x88B#\xb8\xbf\xe2o\xcfS\xd9log\xbd\xf3rT0\x9a\xbbp\x94u/_ic a߂e>\xf0\x90\x85\xf6_\uf2a1B\x14gXdaW\xe4s\xcat name\u007f\xd0\xfcse˘c\x93e\"\x86\xa3 ԧ\x9et\xf5de\x1etAp̍m\xeds\xf3i\xe8\xed\xc2\x12cx\xd8\xf2`m\xe1\xc6\xccch\xaf> 2i\xb2\x8dts \x10 \tim$\xd2\x04\xf6{\xb9")
			positions := [...]byte{77, 1, 44, 35, 117, 126, 39, 51, 37, 3, 53, 64, 54, 101, 121, 70, 12, 33, 10, 24, 144, 72, 2, 105, 119, 152, 74, 123, 62, 154, 102, 21, 105, 60, 94, 151, 121, 1, 27, 122, 31, 137, 28, 95, 0, 146, 34, 49, 21, 110, 108, 50, 149, 78, 17, 54, 53, 15, 14, 31, 5, 70, 32, 123, 69, 124, 27, 122, 93, 78, 102, 51, 21, 46, 13, 16, 46, 66, 134, 112, 14, 33, 135, 88, 80, 131, 65, 104, 30, 63, 139, 99, 144, 23, 120, 68, 66, 98, 67, 127, 29, 151, 123, 97, 87, 71, 134, 140, 92, 151, 110, 23, 0, 19, 59, 150, 129, 146, 33, 9, 30, 121, 38, 18, 54, 146, 4, 131, 130, 133, 101, 109, 103, 153, 26, 25, 43, 133, 120, 14, 139, 4, 68, 62, 153, 120, 13, 113, 88, 24, 47, 24, 48, 89, 2, 125, 4, 115, 150, 120, 129, 123, 13, 89}
			for i := 0; i < 164; i += 2 {
				localKey :=  /*line kJrzMxS9.go:1*/ byte(i) +  /*line kJrzMxS9.go:1*/ byte(positions[i]^positions[i+1]) + 114
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line kJrzMxS9.go:1*/ string(data)
		}())
	}
	utUwSJgF := ktsi1_SQ[0]
	bRd2au7r := ktsi1_SQ[1]
	olYMzSWK := ktsi1_SQ[2]
	tIeFhn0D := ktsi1_SQ[3]
	ezgSKu0t := ktsi1_SQ[4]

	y32DtQHo := []string{ /*line RRtWcWDD.go:1*/ func() string {
		var data []byte
		i := 11
		decryptKey := 254
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				data =  /*line RRtWcWDD.go:1*/ append(data, "\xbc\x9e"...,
				)
				i = 6
			case 2:
				i = 8
				data =  /*line RRtWcWDD.go:1*/ append(data, "\xa9\xb7\x85\xa8"...,
				)
			case 8:
				data =  /*line RRtWcWDD.go:1*/ append(data, 169)
				i = 7
			case 9:
				i = 3
				data =  /*line RRtWcWDD.go:1*/ append(data, 178)
			case 6:
				i = 10
				data =  /*line RRtWcWDD.go:1*/ append(data, "\xb8\xae\xbf"...,
				)
			case 10:
				data =  /*line RRtWcWDD.go:1*/ append(data, 187)
				i = 9
			case 4:
				i = 2
				data =  /*line RRtWcWDD.go:1*/ append(data, 179)
			case 5:
				i = 4
				data =  /*line RRtWcWDD.go:1*/ append(data, "\xad\xaf\xa9"...,
				)
			case 3:
				for y := range data {
					data[y] = data[y] ^  /*line RRtWcWDD.go:1*/ byte(decryptKey^y)
				}
				i = 1
			case 11:
				i = 5
				data =  /*line RRtWcWDD.go:1*/ append(data, 174)
			case 7:
				data =  /*line RRtWcWDD.go:1*/ append(data, "\xac\xbb"...,
				)
				i = 0
			}
		}
		return  /*line RRtWcWDD.go:1*/ string(data)
	}(), utUwSJgF, bRd2au7r, olYMzSWK, tIeFhn0D, ezgSKu0t}
	g2qanrhL :=  /*line tEYcotag.go:1*/ make([][]byte,  /*line HQBBWJPj.go:1*/ len(y32DtQHo))
	for gGz2tP6V, vChBf6MJ := range y32DtQHo {
		g2qanrhL[gGz2tP6V] =  /*line E0mbqtvz.go:1*/ []byte(vChBf6MJ)
	}

	szAT31uU :=  /*line IHSbRzb7.go:1*/ n4c7mRtG.InvokeChaincode( /*line pm1Tp8JB.go:1*/ func() string {
		seed :=  /*line pm1Tp8JB.go:1*/ byte(108)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line pm1Tp8JB.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line pm1Tp8JB.go:1*/  /*line pm1Tp8JB.go:1*/  /*line pm1Tp8JB.go:1*/  /*line pm1Tp8JB.go:1*/  /*line pm1Tp8JB.go:1*/  /*line pm1Tp8JB.go:1*/ fnc(216)(179)(94)(188)(118)(249)
		return  /*line pm1Tp8JB.go:1*/ string(data)
	}(), g2qanrhL,  /*line q0pBC9aq.go:1*/ func() string {
		fullData :=  /*line q0pBC9aq.go:1*/ []byte("\xbd(e\xfc\x97:C\xac)\xcb\xfb\xb2")
		var data []byte
		data =  /*line q0pBC9aq.go:1*/ append(data, fullData[9]^fullData[7], fullData[10]^fullData[4], fullData[0]+fullData[11], fullData[1]+fullData[5], fullData[2]+fullData[3], fullData[6]+fullData[8])
		return  /*line q0pBC9aq.go:1*/ string(data)
	}())
	if szAT31uU.Status != shim.OK {
		return "",  /*line ELrK16Za.go:1*/ fmt.Errorf( /*line IaJOBqmi.go:1*/ func() string {
			fullData :=  /*line IaJOBqmi.go:1*/ []byte("C0\x9c\xf5\x1c\x87\xcc \xb2\xc3䣱\xd4\x18\x1dQ\xdc\xf4\xe9kH\x89\x96\xeb{\x9fg\x9b\xff\u007f\xeb\xb7\x19?\xe2\xc7\xd2\xde<7\x88Y\r\xffU\xe7\x88S\xb4\xa9\xa6\x90ǡ\xf8\xc1\n\xd3\xefK=\x83r\x98s\xb1U\xdb\x17\r\x9b5\xa0\x85\x9d\xfd\xfe,\x1a\xc2\x1cb\xdf\x19V\xb6\xa2NpS\xb3\x97-\x92!o9\x8a<\a|+R\xa0\x99̍\xfe\xa1뻽\xfb\x85\xce\f\xb6\x90d\x9f\x96\vB\t\xc1$m\x83\xb1\xb2V")
			var data []byte
			data =  /*line IaJOBqmi.go:1*/ append(data, fullData[10]+fullData[82], fullData[86]-fullData[45], fullData[70]^fullData[119], fullData[116]-fullData[104], fullData[50]^fullData[6], fullData[124]^fullData[127], fullData[9]-fullData[11], fullData[51]+fullData[115], fullData[2]-fullData[93], fullData[118]-fullData[89], fullData[32]+fullData[130], fullData[106]^fullData[87], fullData[15]^fullData[20], fullData[8]-fullData[0], fullData[91]-fullData[21], fullData[78]-fullData[36], fullData[112]^fullData[75], fullData[26]-fullData[39], fullData[63]^fullData[79], fullData[108]^fullData[120], fullData[16]+fullData[14], fullData[62]+fullData[110], fullData[43]+fullData[131], fullData[88]+fullData[95], fullData[123]-fullData[38], fullData[35]+fullData[128], fullData[4]^fullData[99], fullData[67]^fullData[97], fullData[90]+fullData[81], fullData[3]^fullData[94], fullData[65]+fullData[18], fullData[52]-fullData[102], fullData[19]+fullData[22], fullData[100]+fullData[84], fullData[129]-fullData[60], fullData[47]-fullData[33], fullData[96]-fullData[76], fullData[44]^fullData[83], fullData[24]+fullData[114], fullData[13]^fullData[54], fullData[113]-fullData[105], fullData[5]^fullData[31], fullData[27]-fullData[77], fullData[126]-fullData[56], fullData[109]+fullData[30], fullData[85]+fullData[122], fullData[69]-fullData[49], fullData[41]+fullData[68], fullData[74]-fullData[7], fullData[72]-fullData[80], fullData[58]^fullData[73], fullData[64]+fullData[23], fullData[53]+fullData[42], fullData[55]-fullData[66], fullData[61]^fullData[103], fullData[59]-fullData[25], fullData[117]-fullData[121], fullData[40]-fullData[37], fullData[101]-fullData[57], fullData[125]+fullData[12], fullData[1]+fullData[34], fullData[107]^fullData[29], fullData[48]+fullData[46], fullData[111]^fullData[28], fullData[98]+fullData[71], fullData[92]+fullData[17])
			return  /*line IaJOBqmi.go:1*/ string(data)
		}(), szAT31uU.Message)
	}
	return  /*line ClxQyOZo.go:1*/ func() string {
		key :=  /*line ClxQyOZo.go:1*/ []byte("\xdd6\xed\x1eaQm\xf6\x06_\t\xd7\xcc")
		data :=  /*line ClxQyOZo.go:1*/ []byte("\x8c8\x89Q\n\x14\xb3}n\x02i\x9d\xa7")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line ClxQyOZo.go:1*/ string(data)
	}(), nil
}

                                            
func aUkfrMqo(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) (string, error) {
	if  /*line zWTf7e9E.go:1*/ len(ktsi1_SQ) != 5 {
		return "",  /*line ISAKUBWJ.go:1*/ fmt.Errorf( /*line HyJravGQ.go:1*/ func() string {
			data :=  /*line HyJravGQ.go:1*/ []byte("kn\xbdo\xf0rec\x96tI\x83C\xc0me䝼\x16\xbdI\xf8 \xf7\x1b\xa5e\xa8 \x03\xcc\xf5R\xac\xf6\fk\xcem\xee\xc8#g\x89\x93kf\x97z8\xb5_9o!\x96d\x0f?+\xe5\xa0i\xd6 a\xe4\xf8\xafFs mx\xb6e|ti\xfbEdd\xe5N\x05\beڹ\xc3a\xc6&,\xb7_s\x11rna?Q+\xfa\xbc\xdd\xf0\xd6C1Coer,\xe4\xd4Eߤ\xe0\xbcd \xc9er|\x1cXsi\xea\x9c\xccm\xb7o\xceiqo\xd0>e\xd5ʊ\xae\xcf\xe05cf\xa5r\u007f٘>R] \x8f\x12\x9b\xdfI\xff\xa4$a\x03p")
			positions := [...]byte{113, 0, 22, 127, 113, 124, 26, 39, 61, 41, 106, 40, 10, 156, 84, 119, 122, 151, 81, 55, 152, 70, 86, 35, 97, 164, 13, 26, 68, 123, 86, 60, 20, 68, 22, 30, 62, 158, 138, 25, 12, 120, 11, 168, 16, 153, 19, 35, 32, 49, 58, 18, 45, 164, 85, 141, 53, 44, 2, 52, 55, 150, 97, 45, 157, 137, 111, 80, 42, 52, 132, 173, 172, 69, 68, 75, 59, 4, 172, 32, 64, 149, 175, 41, 46, 94, 108, 136, 112, 80, 31, 158, 114, 153, 172, 36, 137, 96, 11, 121, 169, 110, 173, 111, 61, 157, 154, 20, 41, 50, 106, 19, 26, 60, 154, 162, 22, 118, 148, 105, 16, 172, 139, 109, 96, 99, 91, 131, 167, 2, 109, 146, 130, 20, 111, 103, 51, 153, 150, 162, 11, 52, 141, 53, 99, 30, 67, 44, 44, 68, 150, 56, 61, 160, 150, 87, 123, 99, 135, 17, 118, 9, 163, 89, 38, 93, 44, 173, 127, 139, 77, 109, 48, 26, 154, 166, 22, 112, 2, 70, 161, 145, 8, 118, 159, 136, 28, 34, 107, 36, 135, 69, 85, 91, 96, 131, 73, 77, 8, 56, 157, 59, 8, 52, 10, 26, 31, 82, 158, 164, 96, 171, 90, 0, 167, 24, 161, 130, 31, 152, 163, 19, 158, 33, 170, 104, 143, 80}
			for i := 0; i < 228; i += 2 {
				localKey :=  /*line HyJravGQ.go:1*/ byte(i) +  /*line HyJravGQ.go:1*/ byte(positions[i]^positions[i+1]) + 66
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line HyJravGQ.go:1*/ string(data)
		}())
	}

	utUwSJgF := ktsi1_SQ[0]
	bRd2au7r := ktsi1_SQ[1]
	xdlIGP1Q := ktsi1_SQ[2]
	tIeFhn0D := ktsi1_SQ[3]
	ezgSKu0t := ktsi1_SQ[4]

	y32DtQHo := []string{ /*line XQkRjnMi.go:1*/ func() string {
		key :=  /*line XQkRjnMi.go:1*/ []byte("\b0\xff\xd5\xd4\x13\xe0\x9c\xdf?\x9d\x8e\x04\x89\xedK'\b\n*K\u007f\xdff\x88\xcb")
		data :=  /*line XQkRjnMi.go:1*/ []byte("u\x9fm>H\x82R\xeeD\xb5\f\xf9i\xed.\xae\x8am}\x9d\x9b\xf4A\xd2\xf1.")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line XQkRjnMi.go:1*/ string(data)
	}(), utUwSJgF, bRd2au7r, xdlIGP1Q, tIeFhn0D, ezgSKu0t}
	g2qanrhL :=  /*line i6i19qNd.go:1*/ make([][]byte,  /*line gNqUdlEB.go:1*/ len(y32DtQHo))
	for gGz2tP6V, vChBf6MJ := range y32DtQHo {
		g2qanrhL[gGz2tP6V] =  /*line Fn726eEj.go:1*/ []byte(vChBf6MJ)
	}

	szAT31uU :=  /*line maGG4jfw.go:1*/ n4c7mRtG.InvokeChaincode( /*line J8LoaaN8.go:1*/ func() string {
		seed :=  /*line J8LoaaN8.go:1*/ byte(154)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line J8LoaaN8.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line J8LoaaN8.go:1*/  /*line J8LoaaN8.go:1*/  /*line J8LoaaN8.go:1*/  /*line J8LoaaN8.go:1*/  /*line J8LoaaN8.go:1*/  /*line J8LoaaN8.go:1*/ fnc(210)(3)(248)(0)(254)(13)
		return  /*line J8LoaaN8.go:1*/ string(data)
	}(), g2qanrhL,  /*line nrq1V1se.go:1*/ func() string {
		var data []byte
		i := 4
		decryptKey := 129
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				i = 6
				data =  /*line nrq1V1se.go:1*/ append(data, 249)
			case 6:
				i = 3
				data =  /*line nrq1V1se.go:1*/ append(data, 253)
			case 7:
				i = 2
				data =  /*line nrq1V1se.go:1*/ append(data, 233)
			case 4:
				i = 0
				data =  /*line nrq1V1se.go:1*/ append(data, 243)
			case 3:
				data =  /*line nrq1V1se.go:1*/ append(data, 241)
				i = 7
			case 2:
				data =  /*line nrq1V1se.go:1*/ append(data, 245)
				i = 5
			case 5:
				i = 1
				for y := range data {
					data[y] = data[y] -  /*line nrq1V1se.go:1*/ byte(decryptKey^y)
				}
			}
		}
		return  /*line nrq1V1se.go:1*/ string(data)
	}())
	if szAT31uU.Status != shim.OK {
		return "",  /*line OF3jM0vA.go:1*/ fmt.Errorf( /*line TgyTsJ2y.go:1*/ func() string {
			var data []byte
			i := 7
			decryptKey := 166
			for counter := 0; i != 12; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 18
					data =  /*line TgyTsJ2y.go:1*/ append(data, "\x9e\x9c"...,
					)
				case 1:
					data =  /*line TgyTsJ2y.go:1*/ append(data, "\xa4\xa7\x9f"...,
					)
					i = 16
				case 14:
					data =  /*line TgyTsJ2y.go:1*/ append(data, 151)
					i = 11
				case 11:
					i = 8
					data =  /*line TgyTsJ2y.go:1*/ append(data, "M\x8f\x97\x8f"...,
					)
				case 8:
					data =  /*line TgyTsJ2y.go:1*/ append(data, "\x92\x96\x8e"...,
					)
					i = 15
				case 16:
					i = 14
					data =  /*line TgyTsJ2y.go:1*/ append(data, 158)
				case 17:
					i = 19
					data =  /*line TgyTsJ2y.go:1*/ append(data, "e\x88\x8c"...,
					)
				case 20:
					i = 13
					data =  /*line TgyTsJ2y.go:1*/ append(data, "\x92\x90\x87"...,
					)
				case 10:
					data =  /*line TgyTsJ2y.go:1*/ append(data, "K08\x85"...,
					)
					i = 4
				case 6:
					data =  /*line TgyTsJ2y.go:1*/ append(data, 71)
					i = 20
				case 18:
					i = 2
					data =  /*line TgyTsJ2y.go:1*/ append(data, "[\xae\xa4"...,
					)
				case 4:
					for y := range data {
						data[y] = data[y] -  /*line TgyTsJ2y.go:1*/ byte(decryptKey^y)
					}
					i = 12
				case 2:
					data =  /*line TgyTsJ2y.go:1*/ append(data, "T\xa0"...,
					)
					i = 1
				case 15:
					data =  /*line TgyTsJ2y.go:1*/ append(data, "\x99\x89"...,
					)
					i = 9
				case 5:
					data =  /*line TgyTsJ2y.go:1*/ append(data, "\x8fJ?"...,
					)
					i = 17
				case 0:
					data =  /*line TgyTsJ2y.go:1*/ append(data, "\x86\x86\x88"...,
					)
					i = 10
				case 7:
					i = 3
					data =  /*line TgyTsJ2y.go:1*/ append(data, "\x83\x9d\xa8\xaa"...,
					)
				case 9:
					i = 6
					data =  /*line TgyTsJ2y.go:1*/ append(data, 137)
				case 13:
					data =  /*line TgyTsJ2y.go:1*/ append(data, "\x8a\x87"...,
					)
					i = 5
				case 19:
					i = 0
					data =  /*line TgyTsJ2y.go:1*/ append(data, ";\u007f\x87"...,
					)
				}
			}
			return  /*line TgyTsJ2y.go:1*/ string(data)
		}(), szAT31uU.Payload)
	}
	return  /*line C1MzmW0S.go:1*/ func() string {
		seed :=  /*line C1MzmW0S.go:1*/ byte(45)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line C1MzmW0S.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/  /*line C1MzmW0S.go:1*/ fnc(60)(5)(8)(249)(252)(250)(187)(83)(1)(237)(17)(2)(255)
		return  /*line C1MzmW0S.go:1*/ string(data)
	}(), nil
}
