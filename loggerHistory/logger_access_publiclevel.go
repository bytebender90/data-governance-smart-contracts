/*
Copyright 2023 Centre for Research and Technology Hellas (CERTH)

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS “AS IS” AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
//line :1
package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (hslgON00 *EMhVAN9S) em7oAAEw(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	var xREbPiGc error
	var w1e7h2B3 string
	if  /*line BA5A6u9c.go:1*/ len(eFz_TaL1) != 5 {
		return  /*line Q3LvRVcs.go:1*/ shim.Error( /*line _K308lH6.go:1*/ func() string {
			seed :=  /*line _K308lH6.go:1*/ byte(122)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line _K308lH6.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/  /*line _K308lH6.go:1*/ fnc(195)(171)(75)(162)(71)(142)(15)(28)(73)(62)(189)(139)(11)(36)(64)(120)(249)(248)(239)(153)(36)(107)(255)(247)(234)(229)(193)(135)(7)(199)(182)(147)(44)(86)(158)(247)(52)(113)(229)(120)(32)(101)(183)(120)(237)(212)(101)(22)(37)(91)(165)(81)(86)(237)(220)(184)(114)(242)(228)(117)(50)(101)(212)(169)(77)(157)(65)(41)(151)(65)(122)(233)(208)(177)(87)(179)(95)(119)(3)(241)(56)(91)(193)(139)(6)(26)
			return  /*line _K308lH6.go:1*/ string(data)
		}())
	}
	cXmPo8Kd := eFz_TaL1[0]
	dwbLZnck := eFz_TaL1[1]
	gBYxQKrs := eFz_TaL1[2]
	xrSSF6uG := eFz_TaL1[3]
	g3ka7rbx :=  /*line qBAMrab5.go:1*/ strings.Split(eFz_TaL1[4],  /*line jmlIBBAJ.go:1*/ func() string {
		fullData :=  /*line jmlIBBAJ.go:1*/ []byte("\xef\xc3")
		var data []byte
		data =  /*line jmlIBBAJ.go:1*/ append(data, fullData[0]^fullData[1])
		return  /*line jmlIBBAJ.go:1*/ string(data)
	}())
	_0uGwnuQ :=  /*line lt0CtONi.go:1*/ func() string {
		fullData :=  /*line lt0CtONi.go:1*/ []byte("\xe0\x93\xfa\xce\xf1Ի]\xa7h5\x91\xe1\xb0\x1fٻ\x84m\xbbCҼ\xf5\x1b\xb9ׁ\xc0X\xd1\t\\j@T")
		var data []byte
		data =  /*line lt0CtONi.go:1*/ append(data, fullData[6]+fullData[11], fullData[20]-fullData[5], fullData[28]^fullData[8], fullData[19]-fullData[35], fullData[17]^fullData[12], fullData[34]-fullData[3], fullData[15]+fullData[9], fullData[7]-fullData[2], fullData[14]-fullData[22], fullData[31]+fullData[32], fullData[29]+fullData[24], fullData[0]-fullData[18], fullData[30]-fullData[27], fullData[33]-fullData[23], fullData[4]^fullData[1], fullData[26]^fullData[16], fullData[13]+fullData[25], fullData[10]-fullData[21])
		return  /*line lt0CtONi.go:1*/ string(data)
	}()
	cKWTKAnj :=  /*line J7CFnZZD.go:1*/ func() string {
		key :=  /*line J7CFnZZD.go:1*/ []byte("\x17mp\xefJ\xdaEG\x1dV\x86\xf6\x8e!\xdf\x17\xc1\x02h\x99J")
		data :=  /*line J7CFnZZD.go:1*/ []byte("c\xdc\xd7V\xafL\x86\xaa\x80\xbb\xf9iޖA\x83*e\xb3\xfe\xc3")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line J7CFnZZD.go:1*/ string(data)
	}()

	if  /*line vM4kyGxv.go:1*/ len(cXmPo8Kd) <= 0 {
		return  /*line ojlHyHyB.go:1*/ shim.Error( /*line cHpJaNj1.go:1*/ func() string {
			fullData :=  /*line cHpJaNj1.go:1*/ []byte("*\xfcMJ w̘3^\xb5\xb80\x95|E\xed\xddI\xfcc\x10\x91\xdc+\xc15\x8b \x89\x98\xed\xbb\x0e\xa9JV\xb9T\x96f9\xdczQ\\\r\x11<)0\xaeo\xb1\x94i\x91xx?\bZq\xab\xeb\xc3\x1f\x8a")
			var data []byte
			data =  /*line cHpJaNj1.go:1*/ append(data, fullData[52]^fullData[24], fullData[38]+fullData[46], fullData[19]+fullData[57], fullData[56]-fullData[12], fullData[21]^fullData[20], fullData[6]^fullData[34], fullData[25]-fullData[2], fullData[55]-fullData[18], fullData[22]^fullData[1], fullData[53]-fullData[48], fullData[32]+fullData[11], fullData[54]-fullData[28], fullData[31]+fullData[8], fullData[5]+fullData[64], fullData[37]^fullData[42], fullData[43]-fullData[61], fullData[39]-fullData[26], fullData[10]-fullData[13], fullData[66]^fullData[62], fullData[30]-fullData[49], fullData[50]^fullData[9], fullData[29]-fullData[45], fullData[4]^fullData[15], fullData[65]-fullData[36], fullData[58]-fullData[60], fullData[35]+fullData[0], fullData[47]-fullData[7], fullData[63]-fullData[27], fullData[3]^fullData[41], fullData[40]+fullData[33], fullData[51]^fullData[23], fullData[16]+fullData[14], fullData[44]^fullData[59], fullData[67]+fullData[17])
			return  /*line cHpJaNj1.go:1*/ string(data)
		}())
	}

	if  /*line s8UC3rif.go:1*/ len(dwbLZnck) <= 0 {
		return  /*line nzb2ZavQ.go:1*/ shim.Error( /*line pPX3UJY4.go:1*/ func() string {
			var data []byte
			i := 9
			decryptKey := 69
			for counter := 0; i != 17; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 8:
					data =  /*line pPX3UJY4.go:1*/ append(data, 71)
					i = 2
				case 12:
					data =  /*line pPX3UJY4.go:1*/ append(data, "[\xa9\x9fX"...,
					)
					i = 4
				case 4:
					i = 5
					data =  /*line pPX3UJY4.go:1*/ append(data, "\xaf\xb0\xac"...,
					)
				case 19:
					i = 7
					data =  /*line pPX3UJY4.go:1*/ append(data, "O\x8fM"...,
					)
				case 20:
					i = 19
					data =  /*line pPX3UJY4.go:1*/ append(data, "J\x8b\x8d"...,
					)
				case 3:
					i = 14
					data =  /*line pPX3UJY4.go:1*/ append(data, "\xa5\x96\xa2\xa5"...,
					)
				case 10:
					i = 18
					data =  /*line pPX3UJY4.go:1*/ append(data, "\x8a\x88"...,
					)
				case 16:
					data =  /*line pPX3UJY4.go:1*/ append(data, 141)
					i = 10
				case 2:
					data =  /*line pPX3UJY4.go:1*/ append(data, "\x93\x9a\x97\x9f"...,
					)
					i = 20
				case 6:
					i = 16
					data =  /*line pPX3UJY4.go:1*/ append(data, "\x89\x8d;\x8d"...,
					)
				case 1:
					data =  /*line pPX3UJY4.go:1*/ append(data, 134)
					i = 13
				case 14:
					data =  /*line pPX3UJY4.go:1*/ append(data, 151)
					i = 15
				case 0:
					i = 17
					for y := range data {
						data[y] = data[y] -  /*line pPX3UJY4.go:1*/ byte(decryptKey^y)
					}
				case 11:
					data =  /*line pPX3UJY4.go:1*/ append(data, ">u\x84\x86"...,
					)
					i = 6
				case 7:
					data =  /*line pPX3UJY4.go:1*/ append(data, "\x9a\x82\x80"...,
					)
					i = 11
				case 9:
					i = 3
					data =  /*line pPX3UJY4.go:1*/ append(data, 136)
				case 5:
					data =  /*line pPX3UJY4.go:1*/ append(data, "\xb2\x8c"...,
					)
					i = 1
				case 15:
					data =  /*line pPX3UJY4.go:1*/ append(data, "\xa2\x99"...,
					)
					i = 12
				case 13:
					data =  /*line pPX3UJY4.go:1*/ append(data, "\x86\x92"...,
					)
					i = 8
				case 18:
					data =  /*line pPX3UJY4.go:1*/ append(data, "\x8c\x84"...,
					)
					i = 0
				}
			}
			return  /*line pPX3UJY4.go:1*/ string(data)
		}())
	}

	pchk2ap7 :=  /*line DvI7HX3S.go:1*/ func() string {
		var data []byte
		i := 3
		decryptKey := 232
		for counter := 0; i != 7; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 10:
				for y := range data {
					data[y] = data[y] -  /*line DvI7HX3S.go:1*/ byte(decryptKey^y)
				}
				i = 7
			case 9:
				data =  /*line DvI7HX3S.go:1*/ append(data, 45)
				i = 5
			case 5:
				i = 2
				data =  /*line DvI7HX3S.go:1*/ append(data, 64)
			case 3:
				data =  /*line DvI7HX3S.go:1*/ append(data, 48)
				i = 8
			case 6:
				i = 11
				data =  /*line DvI7HX3S.go:1*/ append(data, 61)
			case 1:
				data =  /*line DvI7HX3S.go:1*/ append(data, 48)
				i = 0
			case 0:
				data =  /*line DvI7HX3S.go:1*/ append(data, 49)
				i = 4
			case 2:
				i = 10
				data =  /*line DvI7HX3S.go:1*/ append(data, 245)
			case 8:
				i = 1
				data =  /*line DvI7HX3S.go:1*/ append(data, 49)
			case 11:
				i = 9
				data =  /*line DvI7HX3S.go:1*/ append(data, 20)
			case 4:
				data =  /*line DvI7HX3S.go:1*/ append(data, 62)
				i = 6
			}
		}
		return  /*line DvI7HX3S.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line yz4g3zIk.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 43
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				data =  /*line yz4g3zIk.go:1*/ append(data, 4)
				i = 0
			case 0:
				for y := range data {
					data[y] = data[y] +  /*line yz4g3zIk.go:1*/ byte(decryptKey^y)
				}
				i = 2
			}
		}
		return  /*line yz4g3zIk.go:1*/ string(data)
	}() + dwbLZnck
	fAxwz5lN := LoggerAccessPublicCounter{}
	sdHID17m := -1
	atKj4ikX, xREbPiGc :=  /*line l_2TjYhw.go:1*/ fOOTaXmK.GetState(pchk2ap7)	                                            
	if xREbPiGc != nil {
		w1e7h2B3 =  /*line BRwt2BD7.go:1*/ func() string {
			var data []byte
			i := 9
			decryptKey := 152
			for counter := 0; i != 12; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 15:
					i = 5
					data =  /*line BRwt2BD7.go:1*/ append(data, "\v\xbaѸ"...,
					)
				case 16:
					i = 2
					data =  /*line BRwt2BD7.go:1*/ append(data, 30)
				case 11:
					data =  /*line BRwt2BD7.go:1*/ append(data, 239)
					i = 1
				case 17:
					i = 4
					data =  /*line BRwt2BD7.go:1*/ append(data, "(\xd5"...,
					)
				case 6:
					data =  /*line BRwt2BD7.go:1*/ append(data, "\xe8(11"...,
					)
					i = 3
				case 1:
					data =  /*line BRwt2BD7.go:1*/ append(data, 253)
					i = 13
				case 9:
					i = 8
					data =  /*line BRwt2BD7.go:1*/ append(data, "\x1a\xc0"...,
					)
				case 7:
					data =  /*line BRwt2BD7.go:1*/ append(data, "\xfe\xf6\xf4\xaf"...,
					)
					i = 14
				case 5:
					i = 7
					data =  /*line BRwt2BD7.go:1*/ append(data, "\xdb\xf5\xfc"...,
					)
				case 8:
					i = 0
					data =  /*line BRwt2BD7.go:1*/ append(data, "\xe2\x0e"...,
					)
				case 14:
					data =  /*line BRwt2BD7.go:1*/ append(data, "\x02\xfc\xac\xf2"...,
					)
					i = 11
				case 2:
					i = 17
					data =  /*line BRwt2BD7.go:1*/ append(data, 38)
				case 4:
					i = 12
					for y := range data {
						data[y] = data[y] -  /*line BRwt2BD7.go:1*/ byte(decryptKey^y)
					}
				case 10:
					i = 6
					data =  /*line BRwt2BD7.go:1*/ append(data, "\xe9\xf6\xf5\xa1"...,
					)
				case 3:
					data =  /*line BRwt2BD7.go:1*/ append(data, "+-3\xd9"...,
					)
					i = 16
				case 0:
					data =  /*line BRwt2BD7.go:1*/ append(data, "\r\t"...,
					)
					i = 15
				case 13:
					i = 10
					data =  /*line BRwt2BD7.go:1*/ append(data, "\xa8\xe8\xe9\xe8"...,
					)
				}
			}
			return  /*line BRwt2BD7.go:1*/ string(data)
		}() + pchk2ap7 +  /*line opRrSZ3i.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 170
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					for y := range data {
						data[y] = data[y] +  /*line opRrSZ3i.go:1*/ byte(decryptKey^y)
					}
					i = 1
				case 3:
					i = 0
					data =  /*line opRrSZ3i.go:1*/ append(data, 116)
				case 0:
					data =  /*line opRrSZ3i.go:1*/ append(data, 206)
					i = 2
				}
			}
			return  /*line opRrSZ3i.go:1*/ string(data)
		}()
		return  /*line cnFeOkdf.go:1*/ shim.Error(w1e7h2B3)
	} else if atKj4ikX == nil {
		sdHID17m = 0
		ezyyoaqN := &LoggerRevokedAccessCounterPublic{cKWTKAnj, cXmPo8Kd, sdHID17m}
		atKj4ikX, xREbPiGc =  /*line jfLvvOqI.go:1*/ json.Marshal(ezyyoaqN)
		if xREbPiGc != nil {
			return  /*line _J4E1K_h.go:1*/ shim.Error( /*line tArOWT7n.go:1*/ xREbPiGc.Error())
		}

	} else {
		xREbPiGc =  /*line qyTkARXz.go:1*/ json.Unmarshal(atKj4ikX, &fAxwz5lN)	                               
		if xREbPiGc != nil {
			return  /*line __i34T0u.go:1*/ shim.Error( /*line zyz6s9NY.go:1*/ xREbPiGc.Error())
		}
		sdHID17m = fAxwz5lN.Count + 1
		fAxwz5lN.Count = sdHID17m
		atKj4ikX, _ =  /*line Jw2Yk09y.go:1*/ json.Marshal(fAxwz5lN)
	}
	                                               
	xREbPiGc =  /*line AMdGAH2i.go:1*/ fOOTaXmK.PutState(pchk2ap7, atKj4ikX)
	if xREbPiGc != nil {
		return  /*line wDqUZCvi.go:1*/ shim.Error( /*line RAK8BeBB.go:1*/ xREbPiGc.Error())
	}

	f0KssXLM := &LoggerAccessPublic{_0uGwnuQ, cXmPo8Kd, dwbLZnck, gBYxQKrs, xrSSF6uG, g3ka7rbx, sdHID17m}
	z5YH5tSP, xREbPiGc :=  /*line b3YNGSVz.go:1*/ json.Marshal(f0KssXLM)
	if xREbPiGc != nil {
		return  /*line L33Z0JZU.go:1*/ shim.Error( /*line ug9z9TkT.go:1*/ xREbPiGc.Error())
	}

	em7oAAEw :=  /*line w_tTwHih.go:1*/ func() string {
		var data []byte
		i := 0
		decryptKey := 223
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 7:
				data =  /*line w_tTwHih.go:1*/ append(data, 89)
				i = 4
			case 3:
				i = 6
				data =  /*line w_tTwHih.go:1*/ append(data, 100)
			case 1:
				i = 7
				data =  /*line w_tTwHih.go:1*/ append(data, 88)
			case 6:
				data =  /*line w_tTwHih.go:1*/ append(data, 33)
				i = 5
			case 5:
				for y := range data {
					data[y] = data[y] -  /*line w_tTwHih.go:1*/ byte(decryptKey^y)
				}
				i = 2
			case 0:
				data =  /*line w_tTwHih.go:1*/ append(data, 85)
				i = 1
			case 4:
				data =  /*line w_tTwHih.go:1*/ append(data, 92)
				i = 8
			case 8:
				i = 3
				data =  /*line w_tTwHih.go:1*/ append(data, 99)
			}
		}
		return  /*line w_tTwHih.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line p0FqzLWl.go:1*/ func() string {
		seed :=  /*line p0FqzLWl.go:1*/ byte(175)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line p0FqzLWl.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line p0FqzLWl.go:1*/ fnc(128)
		return  /*line p0FqzLWl.go:1*/ string(data)
	}() + dwbLZnck +  /*line oGd5BWzp.go:1*/ func() string {
		key :=  /*line oGd5BWzp.go:1*/ []byte("c")
		data :=  /*line oGd5BWzp.go:1*/ []byte("\xcc")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line oGd5BWzp.go:1*/ string(data)
	}() +  /*line oMbOWdy3.go:1*/ strconv.Itoa(sdHID17m)
	                        
	xREbPiGc =  /*line nxn171w7.go:1*/ fOOTaXmK.PutState(em7oAAEw, z5YH5tSP)
	if xREbPiGc != nil {
		return  /*line BDjMp5NX.go:1*/ shim.Error( /*line V9rSsbGN.go:1*/ xREbPiGc.Error())
	}

	auAEXfA1 :=  /*line seuMRI5w.go:1*/ func() string {
		key :=  /*line seuMRI5w.go:1*/ []byte("\xc1\x11\xf3~\xd8\b\xab2\x0e\xbc\x81\xe46_\xa7\xdd\x06\xd6S\xe4\xd1BM\x89\x93\x9b\xc48\xe6\x13F\xa6\x91\xe86\a;YP\x14\xb4\x86")
		data :=  /*line seuMRI5w.go:1*/ []byte("\x04}X\xdfJq\x19\x99.\x04\xf0Y\xa9\xc4\xc7M{8\xbfM4b\xae\xec\xf6\x007\xab\x06y\xb5\x18\xb1L\x97{\x9c̵\x88\xee\xa6")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line seuMRI5w.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line f7lP_UeT.go:1*/ func() string {
		data :=  /*line f7lP_UeT.go:1*/ []byte("proY1.e<'(")
		positions := [...]byte{5, 3, 4, 7, 5, 7, 3, 4, 8, 9, 8, 7, 9, 7}
		for i := 0; i < 14; i += 2 {
			localKey :=  /*line f7lP_UeT.go:1*/ byte(i) +  /*line f7lP_UeT.go:1*/ byte(positions[i]^positions[i+1]) + 148
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line f7lP_UeT.go:1*/ string(data)
	}() + dwbLZnck +  /*line cFJ21dRr.go:1*/ func() string {
		seed :=  /*line cFJ21dRr.go:1*/ byte(83)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line cFJ21dRr.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/  /*line cFJ21dRr.go:1*/ fnc(35)(19)(251)(233)(4)(2)(0)(26)(226)(1)(74)(154)
		return  /*line cFJ21dRr.go:1*/ string(data)
	}() + gBYxQKrs
	pPXX7o_t :=  /*line GHkmzs_H.go:1*/ []byte(auAEXfA1)
	return  /*line QDsJStFz.go:1*/ shim.Success(pPXX7o_t)

}

func (hslgON00 *EMhVAN9S) dFYSS6WB(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	var w1e7h2B3 string
	var xREbPiGc error
	if  /*line nh9E14yZ.go:1*/ len(eFz_TaL1) != 5 {
		return  /*line BrPXeygR.go:1*/ shim.Error( /*line _4idMXoh.go:1*/ func() string {
			seed :=  /*line _4idMXoh.go:1*/ byte(40)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line _4idMXoh.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/  /*line _4idMXoh.go:1*/ fnc(33)(37)(245)(12)(3)(0)(243)(254)(17)(172)(65)(17)(245)(14)(248)(248)(9)(6)(255)(187)(242)(44)(35)(248)(0)(254)(13)(174)(51)(240)(221)(70)(9)(3)(174)(77)(2)(255)(251)(11)(251)(3)(247)(5)(249)(185)(82)(243)(17)(249)(252)(250)(255)(188)(65)(2)(0)(2)(14)(0)(173)(70)(9)(3)(174)(80)(5)(237)(10)(253)(250)(189)(65)(2)(0)(2)(14)(0)(173)(69)(19)(248)(245)(254)(17)(245)(5)(249)(185)(21)(235)(86)(235)(11)(9)(240)(14)
			return  /*line _4idMXoh.go:1*/ string(data)
		}())
	}
	cXmPo8Kd := eFz_TaL1[0]
	dwbLZnck := eFz_TaL1[1]
	beT6BA9K := eFz_TaL1[2]
	xrSSF6uG := eFz_TaL1[3]
	g3ka7rbx :=  /*line uX5etrsi.go:1*/ strings.Split(eFz_TaL1[4],  /*line tf8r95cL.go:1*/ func() string {
		key :=  /*line tf8r95cL.go:1*/ []byte("\xa1")
		data :=  /*line tf8r95cL.go:1*/ []byte("\x8d")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line tf8r95cL.go:1*/ string(data)
	}())
	_0uGwnuQ :=  /*line TKSnvfHv.go:1*/ func() string {
		data :=  /*line TKSnvfHv.go:1*/ []byte("\xf2Bg@7rLE1\x03keC\xcecc\x0essP\a\xe3l-\xa0")
		positions := [...]byte{9, 24, 13, 21, 8, 8, 20, 7, 0, 16, 16, 13, 13, 16, 9, 21, 0, 9, 23, 21, 24, 6, 12, 3, 4, 7, 23, 13, 1, 1}
		for i := 0; i < 30; i += 2 {
			localKey :=  /*line TKSnvfHv.go:1*/ byte(i) +  /*line TKSnvfHv.go:1*/ byte(positions[i]^positions[i+1]) + 183
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line TKSnvfHv.go:1*/ string(data)
	}()
	cKWTKAnj :=  /*line HZAQK_iq.go:1*/ func() string {
		seed :=  /*line HZAQK_iq.go:1*/ byte(200)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line HZAQK_iq.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/  /*line HZAQK_iq.go:1*/ fnc(20)(75)(142)(28)(54)(121)(210)(183)(127)(247)(234)(206)(155)(19)(72)(144)(34)(82)(164)(37)(111)(203)(160)(61)(116)(208)(186)(136)
		return  /*line HZAQK_iq.go:1*/ string(data)
	}()

	if  /*line aK1a7Yzp.go:1*/ len(cXmPo8Kd) <= 0 {
		return  /*line rs7jXqTh.go:1*/ shim.Error( /*line dUJUkFBo.go:1*/ func() string {
			var data []byte
			i := 15
			decryptKey := 107
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 1:
					i = 12
					data =  /*line dUJUkFBo.go:1*/ append(data, "sw"...,
					)
				case 15:
					i = 8
					data =  /*line dUJUkFBo.go:1*/ append(data, 96)
				case 2:
					data =  /*line dUJUkFBo.go:1*/ append(data, "\x89\x8b0"...,
					)
					i = 1
				case 12:
					data =  /*line dUJUkFBo.go:1*/ append(data, "3m-"...,
					)
					i = 7
				case 10:
					data =  /*line dUJUkFBo.go:1*/ append(data, "xty\u007f"...,
					)
					i = 13
				case 11:
					i = 0
					for y := range data {
						data[y] = data[y] -  /*line dUJUkFBo.go:1*/ byte(decryptKey^y)
					}
				case 9:
					data =  /*line dUJUkFBo.go:1*/ append(data, "tl\xaa\xa4"...,
					)
					i = 11
				case 5:
					i = 14
					data =  /*line dUJUkFBo.go:1*/ append(data, "\x8b~\x8e"...,
					)
				case 4:
					data =  /*line dUJUkFBo.go:1*/ append(data, "~v6"...,
					)
					i = 3
				case 13:
					i = 9
					data =  /*line dUJUkFBo.go:1*/ append(data, "'su"...,
					)
				case 3:
					data =  /*line dUJUkFBo.go:1*/ append(data, 111)
					i = 10
				case 8:
					data =  /*line dUJUkFBo.go:1*/ append(data, "~\x92\x80"...,
					)
					i = 5
				case 14:
					i = 6
					data =  /*line dUJUkFBo.go:1*/ append(data, ";\x81"...,
					)
				case 7:
					data =  /*line dUJUkFBo.go:1*/ append(data, 124)
					i = 4
				case 6:
					data =  /*line dUJUkFBo.go:1*/ append(data, 138)
					i = 2
				}
			}
			return  /*line dUJUkFBo.go:1*/ string(data)
		}())
	}

	if  /*line PXGfo4of.go:1*/ len(dwbLZnck) <= 0 {
		return  /*line pzD_CRr6.go:1*/ shim.Error( /*line hAvDtOyR.go:1*/ func() string {
			seed :=  /*line hAvDtOyR.go:1*/ byte(39)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line hAvDtOyR.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/  /*line hAvDtOyR.go:1*/ fnc(124)(22)(30)(73)(142)(15)(42)(76)(83)(245)(225)(124)(72)(146)(33)(73)(133)(5)(11)(35)(244)(53)(114)(226)(197)(54)(174)(95)(121)(51)(37)(152)(49)(97)(129)(58)(124)(251)(250)(249)(153)(133)(11)(20)(31)(67)(127)
			return  /*line hAvDtOyR.go:1*/ string(data)
		}())
	}

	l6ajnYn1 :=  /*line CLGa74TX.go:1*/ func() string {
		seed :=  /*line CLGa74TX.go:1*/ byte(47)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line CLGa74TX.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line CLGa74TX.go:1*/  /*line CLGa74TX.go:1*/  /*line CLGa74TX.go:1*/  /*line CLGa74TX.go:1*/  /*line CLGa74TX.go:1*/  /*line CLGa74TX.go:1*/  /*line CLGa74TX.go:1*/  /*line CLGa74TX.go:1*/  /*line CLGa74TX.go:1*/  /*line CLGa74TX.go:1*/  /*line CLGa74TX.go:1*/ fnc(93)(233)(3)(23)(228)(22)(237)(61)(214)(240)(86)
		return  /*line CLGa74TX.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line QNBeDIFg.go:1*/ func() string {
		key :=  /*line QNBeDIFg.go:1*/ []byte("\x9f")
		data :=  /*line QNBeDIFg.go:1*/ []byte("\xb0")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line QNBeDIFg.go:1*/ string(data)
	}() + dwbLZnck
	tB0W0Bax := LoggerRevokedAccessCounterPublic{}
	sdHID17m := -1
	a4oCVVI8, xREbPiGc :=  /*line yUabW9ZL.go:1*/ fOOTaXmK.GetState(l6ajnYn1)	                                             
	if xREbPiGc != nil {
		w1e7h2B3 =  /*line _F7tTZGq.go:1*/ func() string {
			fullData :=  /*line _F7tTZGq.go:1*/ []byte("Y\xf7\x860\x93\xc2\xf0\xc61'xlt+\xd89\x1f\xa3k(\x87\xedAY\xe2#\xd4JO\xa2\xf4\x01)\xd2\xcd\xd4\xdd\x1a\x85\xfeKw\x14\\\x04\xdbz1\x83\x8eX\xe3aS\xcd\xcd~b\x93\xa2\b\xbe\xed\xe1cA{\u007f\uf51c/\xed\xd9\x03pc\xc7\xd1\x05\xaf$\xf4\x9b\x8eAk\xcf\xeb\u007fl3\xcc?\xe8\xfe\xf6<\x87 ^\xdb")
			var data []byte
			data =  /*line _F7tTZGq.go:1*/ append(data, fullData[25]+fullData[50], fullData[1]+fullData[13], fullData[73]+fullData[90], fullData[61]^fullData[92], fullData[85]+fullData[8], fullData[101]-fullData[11], fullData[36]-fullData[86], fullData[26]^fullData[96], fullData[55]-fullData[4], fullData[64]-fullData[22], fullData[98]-fullData[65], fullData[38]-fullData[81], fullData[51]-fullData[46], fullData[10]+fullData[30], fullData[59]^fullData[77], fullData[29]^fullData[7], fullData[87]-fullData[80], fullData[78]+fullData[17], fullData[74]-fullData[69], fullData[53]-fullData[91], fullData[45]-fullData[12], fullData[33]+fullData[58], fullData[75]^fullData[44], fullData[34]^fullData[21], fullData[27]-fullData[14], fullData[49]^fullData[88], fullData[19]^fullData[100], fullData[48]-fullData[42], fullData[97]+fullData[71], fullData[15]-fullData[35], fullData[60]+fullData[43], fullData[18]-fullData[40], fullData[66]^fullData[37], fullData[31]^fullData[57], fullData[63]-fullData[56], fullData[76]-fullData[39], fullData[70]-fullData[32], fullData[82]+fullData[89], fullData[5]^fullData[24], fullData[23]^fullData[47], fullData[3]^fullData[0], fullData[62]+fullData[2], fullData[83]-fullData[9], fullData[6]+fullData[67], fullData[41]^fullData[79], fullData[95]^fullData[20], fullData[16]^fullData[93], fullData[84]^fullData[94], fullData[28]^fullData[99], fullData[52]-fullData[68], fullData[54]^fullData[72])
			return  /*line _F7tTZGq.go:1*/ string(data)
		}() + l6ajnYn1 +  /*line xztKoiu5.go:1*/ func() string {
			seed :=  /*line xztKoiu5.go:1*/ byte(67)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line xztKoiu5.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line xztKoiu5.go:1*/  /*line xztKoiu5.go:1*/ fnc(97)(217)
			return  /*line xztKoiu5.go:1*/ string(data)
		}()
		return  /*line MevqXcwV.go:1*/ shim.Error(w1e7h2B3)
	} else if a4oCVVI8 == nil {
		sdHID17m = 0
		f0KssXLM := &LoggerRevokedAccessCounterPublic{cKWTKAnj, cXmPo8Kd, sdHID17m}
		a4oCVVI8, xREbPiGc =  /*line aisWT134.go:1*/ json.Marshal(f0KssXLM)
		if xREbPiGc != nil {
			return  /*line HZHF2WdB.go:1*/ shim.Error( /*line vwJFFddg.go:1*/ xREbPiGc.Error())
		}

	} else {
		xREbPiGc =  /*line DZ3T_4t2.go:1*/ json.Unmarshal(a4oCVVI8, &tB0W0Bax)	                               
		if xREbPiGc != nil {
			return  /*line GG57PFKJ.go:1*/ shim.Error( /*line vlM9mGUM.go:1*/ xREbPiGc.Error())
		}
		sdHID17m = tB0W0Bax.Count + 1
		tB0W0Bax.Count = sdHID17m
		a4oCVVI8, _ =  /*line R1ezfIkd.go:1*/ json.Marshal(tB0W0Bax)
	}

	                                            
	xREbPiGc =  /*line Scrdw2vX.go:1*/ fOOTaXmK.PutState(l6ajnYn1, a4oCVVI8)
	if xREbPiGc != nil {
		return  /*line dHqw_Wic.go:1*/ shim.Error( /*line fsiv7SzU.go:1*/ xREbPiGc.Error())
	}

	r3ovXd_R := &LoggerRevokedAccessPublic{_0uGwnuQ, cXmPo8Kd, dwbLZnck, xrSSF6uG, beT6BA9K, g3ka7rbx, sdHID17m}
	pF2KGYpq, xREbPiGc :=  /*line QNdCbM7Y.go:1*/ json.Marshal(r3ovXd_R)
	if xREbPiGc != nil {
		return  /*line YWEuidUZ.go:1*/ shim.Error( /*line OzoaqoNr.go:1*/ xREbPiGc.Error())
	}
	blnYsXkH :=  /*line a8JWHM8D.go:1*/ func() string {
		key :=  /*line a8JWHM8D.go:1*/ []byte("\xb9\xc1\u007fK\x10)AP")
		data :=  /*line a8JWHM8D.go:1*/ []byte("\xb9\xa4\xf7$[<#\xdf")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line a8JWHM8D.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line XiNeLQD7.go:1*/ func() string {
		seed :=  /*line XiNeLQD7.go:1*/ byte(90)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line XiNeLQD7.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line XiNeLQD7.go:1*/ fnc(117)
		return  /*line XiNeLQD7.go:1*/ string(data)
	}() + dwbLZnck +  /*line V8zXBnBf.go:1*/ func() string {
		fullData :=  /*line V8zXBnBf.go:1*/ []byte("\xf9\xca")
		var data []byte
		data =  /*line V8zXBnBf.go:1*/ append(data, fullData[0]-fullData[1])
		return  /*line V8zXBnBf.go:1*/ string(data)
	}() +  /*line ePTbx5O1.go:1*/ strconv.Itoa(sdHID17m)
	                                       
	xREbPiGc =  /*line EilgrVvL.go:1*/ fOOTaXmK.PutState(blnYsXkH, pF2KGYpq)
	if xREbPiGc != nil {
		return  /*line lv8uN9lM.go:1*/ shim.Error( /*line IDyK0ail.go:1*/ xREbPiGc.Error())
	}

	bxYNgpVn :=  /*line BV3Q94Hz.go:1*/ func() string {
		seed :=  /*line BV3Q94Hz.go:1*/ byte(25)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line BV3Q94Hz.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/  /*line BV3Q94Hz.go:1*/ fnc(97)(227)(208)(161)(61)(125)(1)(169)(161)(57)(44)(170)(71)(159)(55)(106)(206)(155)(242)(52)(109)(199)(152)(45)(84)(101)(11)(24)(48)(98)(210)(164)(245)(48)(105)(213)(88)(244)(229)(221)(167)(96)(178)(115)(172)(62)
		return  /*line BV3Q94Hz.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line vjMq9NEJ.go:1*/ func() string {
		data :=  /*line vjMq9NEJ.go:1*/ []byte("Grihi^\\n:\x10")
		positions := [...]byte{2, 9, 0, 7, 3, 5, 7, 9, 0, 6, 2, 5, 5, 9}
		for i := 0; i < 14; i += 2 {
			localKey :=  /*line vjMq9NEJ.go:1*/ byte(i) +  /*line vjMq9NEJ.go:1*/ byte(positions[i]^positions[i+1]) + 30
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line vjMq9NEJ.go:1*/ string(data)
	}() + dwbLZnck +  /*line e78qRoer.go:1*/ func() string {
		data :=  /*line e78qRoer.go:1*/ []byte("\x1e\x96u\xb8\xc4  ")
		positions := [...]byte{0, 4, 1, 3, 1, 0, 3, 5}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line e78qRoer.go:1*/ byte(i) +  /*line e78qRoer.go:1*/ byte(positions[i]^positions[i+1]) + 166
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line e78qRoer.go:1*/ string(data)
	}() +  /*line WQzFSDk3.go:1*/ strconv.Itoa(sdHID17m)
	pPXX7o_t :=  /*line dS2NRIAm.go:1*/ []byte(bxYNgpVn)
	return  /*line Gyow9Tas.go:1*/ shim.Success(pPXX7o_t)

}

func (hslgON00 *EMhVAN9S) r4NAdg6P(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	dwbLZnck, xREbPiGc :=  /*line Xt4M69kK.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line UxSPkxcA.go:1*/ shim.Error( /*line OdEyI7kb.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line BPgCiayK.go:1*/ fmt.Sprintf( /*line pbqdCObn.go:1*/ func() string {
		data :=  /*line pbqdCObn.go:1*/ []byte("_Amhd\xf9N,A\xa1\xfd|{X\x0f\"\xbe\x9dy\x854\xd8N\xa5g\x00[\x83\x87r\xdabvo[\xe6d\xf8cc\x91\x1fpPu[l4cn\xb4<8s\xc9JcAm\bO\x82P\x17\xf3v){c4\u007f\x06N%9\x01}^")
		positions := [...]byte{3, 9, 14, 66, 9, 59, 11, 69, 40, 9, 70, 20, 16, 22, 77, 57, 23, 54, 5, 10, 37, 21, 4, 67, 14, 27, 54, 24, 27, 0, 28, 16, 50, 35, 41, 72, 61, 30, 16, 48, 20, 47, 22, 45, 30, 19, 50, 61, 45, 13, 55, 75, 26, 60, 17, 27, 56, 30, 67, 22, 31, 8, 0, 74, 47, 1, 67, 55, 57, 42, 49, 21, 74, 68, 26, 71, 13, 67, 66, 5, 7, 17, 67, 37, 13, 25, 1, 40, 59, 28, 20, 30, 31, 55, 64, 61, 21, 17, 15, 52, 2, 51, 67, 21, 35, 59, 63, 30, 31, 31, 6, 5, 34, 45}
		for i := 0; i < 114; i += 2 {
			localKey :=  /*line pbqdCObn.go:1*/ byte(i) +  /*line pbqdCObn.go:1*/ byte(positions[i]^positions[i+1]) + 186
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line pbqdCObn.go:1*/ string(data)
	}(), dwbLZnck)
	toHNB1kz, xREbPiGc :=  /*line BzGXs9Ps.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line fP8G0EMQ.go:1*/ shim.Error( /*line ivBqLD7z.go:1*/ xREbPiGc.Error())
	}
	return  /*line yFk1rNXp.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) mpjqK08z(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line oCK1P7US.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line P3BlJhtd.go:1*/ shim.Error( /*line rFAFesNx.go:1*/ func() string {
			seed :=  /*line rFAFesNx.go:1*/ byte(225)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line rFAFesNx.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/  /*line rFAFesNx.go:1*/ fnc(104)(37)(8)(235)(11)(253)(251)(188)(33)(49)(245)(14)(248)(248)(9)(6)(172)(46)(39)(248)(245)(3)(13)(188)(242)(37)(51)(248)(245)(254)(17)(245)(5)(249)(185)(68)(253)(19)(237)(18)(242)(15)(172)(78)(243)(12)(248)
			return  /*line rFAFesNx.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	dwbLZnck, xREbPiGc :=  /*line F0IP3_gu.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line lzPnBWNN.go:1*/ shim.Error( /*line W7cs4ii3.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line c31GmeRk.go:1*/ fmt.Sprintf( /*line HbNiu2nt.go:1*/ func() string {
		fullData :=  /*line HbNiu2nt.go:1*/ []byte("\xd7dy8\x16y\u007f\x00q\xf6_{\xc6\xc3\xd8d\xb9h\x14\xe4\t\x90\x8a\xf0\xf0r\xe9\xec\"\xebL\xd3\x10]d\x9et\xc8_\xc4AC\xa0\a\xbb\xbe\x03\xba\xd7|\x1ac\xb8\xe4\xec\x14dB\x90\xc1\x12\x18\xd5}6\x18\xc3\x16+B\x84\x84\x17L\x03J+\xf1NE^:1\x13\xfevToc\x97\xd5z\ta7\xabK\x06\xfa\x1d\xa1\xa9Q\x1c\xda\xe0#\xb7\xfd\xea\xa3t\xd8\x1be\xff\xce͢Y\xd8\fP\xaeF0N\x9e\xc1[-\xeb@\xaa\xb0dt\xa3&\x99OOl5=̈\xa1\u0091W`\xbaà8\x02{2\"Q\xab2 <\x00\x86;\x8a\xffg\xa1I\x85*\anb\x88\xf0hhU\x15\x87\xcb")
		var data []byte
		data =  /*line HbNiu2nt.go:1*/ append(data, fullData[88]+fullData[61], fullData[35]+fullData[70], fullData[2]-fullData[97], fullData[161]+fullData[47], fullData[178]+fullData[53], fullData[105]-fullData[157], fullData[11]^fullData[65], fullData[54]+fullData[146], fullData[48]^fullData[52], fullData[34]^fullData[4], fullData[116]+fullData[86], fullData[140]+fullData[29], fullData[110]+fullData[14], fullData[123]+fullData[136], fullData[109]-fullData[166], fullData[144]+fullData[158], fullData[26]+fullData[91], fullData[89]^fullData[153], fullData[129]^fullData[28], fullData[49]-fullData[121], fullData[80]+fullData[43], fullData[60]-fullData[23], fullData[139]+fullData[171], fullData[124]^fullData[1], fullData[78]+fullData[84], fullData[175]^fullData[180], fullData[133]-fullData[41], fullData[114]^fullData[156], fullData[3]^fullData[33], fullData[57]^fullData[125], fullData[162]+fullData[163], fullData[141]+fullData[67], fullData[62]+fullData[100], fullData[90]^fullData[152], fullData[184]^fullData[27], fullData[55]+fullData[102], fullData[149]+fullData[31], fullData[63]+fullData[39], fullData[103]^fullData[6], fullData[83]-fullData[134], fullData[182]-fullData[24], fullData[160]+fullData[159], fullData[94]+fullData[164], fullData[92]^fullData[119], fullData[59]-fullData[30], fullData[22]+fullData[112], fullData[120]-fullData[142], fullData[19]+fullData[173], fullData[42]+fullData[13], fullData[21]-fullData[176], fullData[37]+fullData[135], fullData[51]-fullData[40], fullData[131]+fullData[168], fullData[58]-fullData[99], fullData[9]+fullData[87], fullData[25]-fullData[165], fullData[8]-fullData[46], fullData[108]+fullData[15], fullData[132]+fullData[130], fullData[143]^fullData[122], fullData[107]-fullData[17], fullData[96]+fullData[113], fullData[71]+fullData[145], fullData[177]-fullData[179], fullData[154]-fullData[82], fullData[0]^fullData[147], fullData[18]-fullData[95], fullData[167]^fullData[38], fullData[72]+fullData[126], fullData[50]^fullData[181], fullData[128]+fullData[93], fullData[77]^fullData[185], fullData[45]+fullData[56], fullData[12]+fullData[10], fullData[85]-fullData[74], fullData[170]^fullData[79], fullData[138]-fullData[98], fullData[69]^fullData[151], fullData[118]+fullData[148], fullData[73]+fullData[183], fullData[117]^fullData[16], fullData[64]^fullData[150], fullData[36]+fullData[169], fullData[76]+fullData[81], fullData[111]^fullData[7], fullData[106]+fullData[115], fullData[75]-fullData[32], fullData[5]+fullData[101], fullData[66]-fullData[127], fullData[174]+fullData[172], fullData[68]-fullData[20], fullData[137]+fullData[104], fullData[155]-fullData[44])
		return  /*line HbNiu2nt.go:1*/ string(data)
	}(), dwbLZnck, kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line AcVTGGTs.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line YiwiZFl0.go:1*/ shim.Error( /*line rQ7_PQ8n.go:1*/ xREbPiGc.Error())
	}
	return  /*line _hKSFoSf.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) ebv_m8Xg(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	                  
	fcEHFI_2 :=  /*line mc8ps3_w.go:1*/ func() string {
		key :=  /*line mc8ps3_w.go:1*/ []byte("c $]\xd4>\x119\xa6\xa9\xff\x04\xbe\xdd\xcad\xc7e\xbdh\xf8ŲW\x05\xa5\a\x1b\xbf\x838\xdf\xe4#\xad\xcfj\xa2\x1f\x93=\x8a\x1e\xdd\xf7ܭ\n\xa1J\a\x12")
		data :=  /*line mc8ps3_w.go:1*/ []byte("\x18\x02O\b\x98'R;\xc9\xc9#6\xbdE\x9a\v\x9c\xef\xbc\bm]\x88\xcbG\xca`L\xa6\xef\x1a\x86\x92L\xbe\x96\xfa\x9fD\xd0(\xe9Us~\x86\xbf_\xc2\xd8vk")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line mc8ps3_w.go:1*/ string(data)
	}()
	toHNB1kz, xREbPiGc :=  /*line XFHw6sx4.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line vGlCseQx.go:1*/ shim.Error( /*line wKKYb0pk.go:1*/ xREbPiGc.Error())
	}
	return  /*line W618gpR1.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) kMnPPFc7(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line _br1qY3W.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line zfohtfnD.go:1*/ shim.Error( /*line UBQOyNBO.go:1*/ func() string {
			key :=  /*line UBQOyNBO.go:1*/ []byte("\fC\xbf\x88\x83EE!c\x83(\x1cή7\xf5`\xf7\x14\x9e\x10I\x97Z\x89I\x81\x99;\x8e\xa3ĺ,I\x03\xb1\xb7\x12l\x86\xb3`;\xaf\x9bK")
			data :=  /*line UBQOyNBO.go:1*/ []byte("=+\xb7\xd9\xe9$\x1f\xff\xde\xef?Y\x9f\xb77\u007f\xc0Wa\xcfR\x1c\xdbԗ\xfc\xf7\xd7*\xd5ѥ\xb4;\xd7a\xb0\xbdO\a\xdf\xc1\xc03\xb2\xd2\x1a")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line UBQOyNBO.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	                  
	fcEHFI_2 :=  /*line ExAacn5q.go:1*/ fmt.Sprintf( /*line wxH9P7vT.go:1*/ func() string {
		key :=  /*line wxH9P7vT.go:1*/ []byte("\xb0Ҍ%\x1f\x1c\xabi\x92Y\x96\v\xb5\x9e\xa8\xf3G\xab\x8eȑ#\xf6^\f\x8c\xb5\x96\x88\x17-\x85I\x94m\x05\xb6\xdcH\xa9\xc9)\xees*\xa5\xfe\x89}\x8e\x1dڸ!\xdf@%\xb0\x8d\xbdy\xfbw\xcf\x000\x06")
		data :=  /*line wxH9P7vT.go:1*/ []byte("\xcbP\xe7@MI\xb8\v\xdd\x19\x8c/Ƅ\xbc|\x1c\xa9\xeb\xa8\xd4\xffD\xc4@\xe3\xb2\xd1\xdd[%\xe0-\xdb\xfe`\xaee\x1b\xba\x9cJ\x85\xddK\xbdn\xe0\xe6\x94\x0fH\xac@\x95!N\xb5\xe7e\xc1'\xae\xa4\"Mw")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line wxH9P7vT.go:1*/ string(data)
	}(), kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line AmqpVaRw.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line TGWL3DxY.go:1*/ shim.Error( /*line OtBu1TH7.go:1*/ xREbPiGc.Error())
	}
	return  /*line O5RikIvT.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) f78r3kP3(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	                  
	fcEHFI_2 :=  /*line N3M144dR.go:1*/ func() string {
		seed :=  /*line N3M144dR.go:1*/ byte(152)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line N3M144dR.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/  /*line N3M144dR.go:1*/ fnc(227)(89)(167)(30)(245)(235)(26)(231)(21)(253)(174)(0)(65)(89)(176)(235)(12)(47)(211)(13)(239)(91)(238)(224)(238)(255)(232)(16)(226)(27)(197)(42)(16)(230)(26)(240)(35)(227)(27)(248)(229)(18)(161)(89)(0)
		return  /*line N3M144dR.go:1*/ string(data)
	}()
	toHNB1kz, xREbPiGc :=  /*line z8fGaCB8.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line Q5zhWhlD.go:1*/ shim.Error( /*line Gbna51uB.go:1*/ xREbPiGc.Error())
	}
	return  /*line kxx60Orp.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) mJnWi4R9(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line zSzINukJ.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line RaL6yrxd.go:1*/ shim.Error( /*line Fu14MSYx.go:1*/ func() string {
			seed :=  /*line Fu14MSYx.go:1*/ byte(71)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line Fu14MSYx.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/  /*line Fu14MSYx.go:1*/ fnc(2)(37)(8)(235)(11)(253)(251)(188)(33)(49)(245)(14)(248)(248)(9)(6)(172)(46)(39)(248)(245)(3)(13)(188)(242)(37)(51)(248)(245)(254)(17)(245)(5)(249)(185)(68)(253)(19)(237)(18)(242)(15)(172)(78)(243)(12)(248)
			return  /*line Fu14MSYx.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	                  
	fcEHFI_2 :=  /*line XTaA5xs8.go:1*/ fmt.Sprintf( /*line ZPog2NF0.go:1*/ func() string {
		data :=  /*line ZPog2NF0.go:1*/ []byte("u\"s.lecE\x1cr(:{Ud\\\x1a\x1dJpe\xda\ag\x00os,\f\x19A\x80\x0f\\DsHub\x84\xfac%F\"%\nha/\",2#&Es\"z6")
		positions := [...]byte{17, 43, 28, 47, 16, 22, 16, 50, 0, 13, 42, 32, 13, 22, 43, 13, 21, 58, 39, 29, 55, 13, 58, 8, 27, 3, 18, 49, 17, 45, 7, 28, 59, 33, 23, 53, 22, 55, 23, 10, 34, 54, 59, 28, 15, 54, 23, 55, 24, 51, 39, 52, 24, 26, 47, 55, 31, 49, 29, 40, 43, 36, 49, 8, 24, 26, 32, 46}
		for i := 0; i < 68; i += 2 {
			localKey :=  /*line ZPog2NF0.go:1*/ byte(i) +  /*line ZPog2NF0.go:1*/ byte(positions[i]^positions[i+1]) + 25
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line ZPog2NF0.go:1*/ string(data)
	}(), kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line ocSySu7c.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line BKPEV907.go:1*/ shim.Error( /*line z1RmBIFp.go:1*/ xREbPiGc.Error())
	}
	return  /*line NxxPHzfK.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) eWGFpIkz(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	dwbLZnck, xREbPiGc :=  /*line ymMlT2pN.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line fU7s4iOX.go:1*/ shim.Error( /*line FvjW9QWQ.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line FK8PgzU7.go:1*/ fmt.Sprintf( /*line qTgNzoso.go:1*/ func() string {
		var data []byte
		i := 23
		decryptKey := 192
		for counter := 0; i != 12; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 28:
				data =  /*line qTgNzoso.go:1*/ append(data, "\"#\x14$"...,
				)
				i = 0
			case 20:
				i = 15
				data =  /*line qTgNzoso.go:1*/ append(data, "\t\xaf\xf4\xfe"...,
				)
			case 9:
				data =  /*line qTgNzoso.go:1*/ append(data, "\x16\xf6"...,
				)
				i = 1
			case 15:
				data =  /*line qTgNzoso.go:1*/ append(data, "\xf5\xe5\r\x03"...,
				)
				i = 29
			case 23:
				data =  /*line qTgNzoso.go:1*/ append(data, "\xfd\xa3"...,
				)
				i = 8
			case 3:
				data =  /*line qTgNzoso.go:1*/ append(data, "\xea\xeb"...,
				)
				i = 22
			case 1:
				i = 18
				data =  /*line qTgNzoso.go:1*/ append(data, "\x1a\n\x13\x13"...,
				)
			case 18:
				i = 5
				data =  /*line qTgNzoso.go:1*/ append(data, 12)
			case 24:
				i = 20
				data =  /*line qTgNzoso.go:1*/ append(data, "\xae\xc5"...,
				)
			case 16:
				i = 26
				data =  /*line qTgNzoso.go:1*/ append(data, "\x15 "...,
				)
			case 17:
				for y := range data {
					data[y] = data[y] +  /*line qTgNzoso.go:1*/ byte(decryptKey^y)
				}
				i = 12
			case 0:
				i = 16
				data =  /*line qTgNzoso.go:1*/ append(data, 31)
			case 21:
				i = 28
				data =  /*line qTgNzoso.go:1*/ append(data, 208)
			case 6:
				i = 13
				data =  /*line qTgNzoso.go:1*/ append(data, "2$\"\""...,
				)
			case 13:
				data =  /*line qTgNzoso.go:1*/ append(data, "2\xe1"...,
				)
				i = 4
			case 10:
				data =  /*line qTgNzoso.go:1*/ append(data, 242)
				i = 3
			case 5:
				data =  /*line qTgNzoso.go:1*/ append(data, "\xce\xd7"...,
				)
				i = 21
			case 8:
				data =  /*line qTgNzoso.go:1*/ append(data, "\xf7\xe8"...,
				)
				i = 10
			case 29:
				i = 14
				data =  /*line qTgNzoso.go:1*/ append(data, 251)
			case 25:
				data =  /*line qTgNzoso.go:1*/ append(data, "\xe6\b\x03\x02"...,
				)
				i = 19
			case 27:
				data =  /*line qTgNzoso.go:1*/ append(data, "\x05\x06\x17"...,
				)
				i = 9
			case 22:
				i = 24
				data =  /*line qTgNzoso.go:1*/ append(data, "\xfb\xf9\xfb"...,
				)
			case 4:
				i = 7
				data =  /*line qTgNzoso.go:1*/ append(data, 252)
			case 26:
				i = 2
				data =  /*line qTgNzoso.go:1*/ append(data, "\x1b\x04\x1e"...,
				)
			case 2:
				i = 6
				data =  /*line qTgNzoso.go:1*/ append(data, "\a,("...,
				)
			case 19:
				data =  /*line qTgNzoso.go:1*/ append(data, "\x03\x0f\xe1\x02"...,
				)
				i = 27
			case 14:
				i = 25
				data =  /*line qTgNzoso.go:1*/ append(data, "\xb7ҹ"...,
				)
			case 11:
				data =  /*line qTgNzoso.go:1*/ append(data, "BE"...,
				)
				i = 17
			case 7:
				data =  /*line qTgNzoso.go:1*/ append(data, "\xe3\xe96\xe8"...,
				)
				i = 11
			}
		}
		return  /*line qTgNzoso.go:1*/ string(data)
	}(), dwbLZnck)
	toHNB1kz, xREbPiGc :=  /*line uDZr2Hyz.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line C5gs6fAW.go:1*/ shim.Error( /*line Vhv9s2oA.go:1*/ xREbPiGc.Error())
	}
	return  /*line OI49pLE6.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) mq0UtCn8(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line g7Ubwrlf.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line N9V9T5Lb.go:1*/ shim.Error( /*line qlMgc175.go:1*/ func() string {
			data :=  /*line qlMgc175.go:1*/ []byte("Bn\t~l\x03\x0f\x16Ar[ymVtB P_mp\x00r[\x186xcr\x1fV\x17nbU\x01tJ!\nJM nJx`")
			positions := [...]byte{14, 36, 0, 34, 44, 6, 0, 46, 11, 30, 39, 3, 13, 20, 37, 10, 20, 0, 24, 38, 8, 18, 27, 11, 17, 0, 18, 24, 5, 31, 5, 38, 20, 2, 24, 36, 37, 38, 44, 20, 41, 33, 25, 36, 6, 29, 15, 33, 27, 20, 33, 39, 33, 40, 23, 3, 31, 6, 39, 15, 45, 34, 35, 21, 37, 34, 28, 25, 41, 7}
			for i := 0; i < 70; i += 2 {
				localKey :=  /*line qlMgc175.go:1*/ byte(i) +  /*line qlMgc175.go:1*/ byte(positions[i]^positions[i+1]) + 240
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line qlMgc175.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	dwbLZnck, xREbPiGc :=  /*line JxjM9OIj.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line ROyj8N9P.go:1*/ shim.Error( /*line ErO_BbLK.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line hOACxDzX.go:1*/ fmt.Sprintf( /*line ziGgznPL.go:1*/ func() string {
		seed :=  /*line ziGgznPL.go:1*/ byte(79)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line ziGgznPL.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/  /*line ziGgznPL.go:1*/ fnc(44)(167)(81)(242)(7)(249)(254)(17)(251)(3)(176)(24)(65)(167)(66)(11)(244)(241)(37)(247)(245)(189)(24)(232)(42)(35)(248)(0)(254)(13)(207)(34)(0)(2)(14)(0)(221)(37)(237)(10)(253)(250)(191)(10)(246)(83)(254)(242)(13)(252)(243)(12)(248)(234)(23)(234)(34)(253)(7)(243)(251)(1)(13)(176)(24)(232)(3)(78)(175)(10)(246)(66)(253)(19)(237)(18)(242)(15)(174)(24)(232)(3)(78)(175)(91)(0)
		return  /*line ziGgznPL.go:1*/ string(data)
	}(), dwbLZnck, kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line BsFcPvIa.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line YAO4KeI6.go:1*/ shim.Error( /*line SGR1O7L5.go:1*/ xREbPiGc.Error())
	}
	return  /*line u8Inplhn.go:1*/ shim.Success(toHNB1kz)
}
