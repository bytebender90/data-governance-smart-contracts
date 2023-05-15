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
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (g4rnrSNn *G1Y_7pz_) aZuYW6KY(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var cSW1wSO3 error
	var vWNUn3N4 string
	var w5VhCLka string
	if  /*line xYDcoaMU.go:1*/ len(ktsi1_SQ) != 5 {
		return  /*line m6ImKjV9.go:1*/ shim.Error( /*line AWYsLdSf.go:1*/ func() string {
			fullData :=  /*line AWYsLdSf.go:1*/ []byte("\xb3\xda!\x88\xbc3+fm\xb4T~{)\"+\xb0l\x8f\xec\x98\x0f_\xe5\xa0\x03I\xe3\xa4C\xfcݐB@k\x0fh5X)f\xe3J\xc6z\xa2ܲ\\e\x92\x8c\xdbV8\x1b\xae\x9bk\xe9R\xd7x\x94G\x91#`U\xb4\xef\xff\xce:\xb71n꾥\xcdaM&\x04\xd6?\v\xff\x1d\x81\x18\xd4\xf6\x18\x03b/\x91\xad}\xc7$׀s\a\xae\xf8\xab89!\x02\xd1\x0f\xbeQ\xb1\b \xb3\x00\xa5q")
			var data []byte
			data =  /*line AWYsLdSf.go:1*/ append(data, fullData[124]^fullData[19], fullData[101]-fullData[36], fullData[104]-fullData[82], fullData[122]-fullData[61], fullData[85]+fullData[37], fullData[16]-fullData[65], fullData[78]+fullData[45], fullData[88]^fullData[6], fullData[81]^fullData[52], fullData[110]-fullData[112], fullData[22]-fullData[109], fullData[107]-fullData[51], fullData[92]+fullData[69], fullData[95]-fullData[0], fullData[94]+fullData[63], fullData[29]+fullData[76], fullData[79]+fullData[97], fullData[67]+fullData[15], fullData[56]^fullData[77], fullData[11]+fullData[71], fullData[42]-fullData[91], fullData[7]+fullData[89], fullData[99]^fullData[27], fullData[9]+fullData[17], fullData[3]-fullData[14], fullData[26]^fullData[84], fullData[55]-fullData[44], fullData[106]+fullData[100], fullData[75]-fullData[118], fullData[13]^fullData[49], fullData[64]+fullData[1], fullData[48]+fullData[119], fullData[121]^fullData[10], fullData[102]+fullData[46], fullData[93]+fullData[58], fullData[54]^fullData[111], fullData[47]-fullData[4], fullData[25]^fullData[125], fullData[18]+fullData[86], fullData[33]^fullData[5], fullData[115]^fullData[28], fullData[50]-fullData[123], fullData[31]^fullData[57], fullData[21]^fullData[12], fullData[34]-fullData[72], fullData[20]-fullData[38], fullData[68]^fullData[96], fullData[105]+fullData[23], fullData[43]+fullData[40], fullData[35]+fullData[120], fullData[73]-fullData[24], fullData[30]+fullData[103], fullData[41]-fullData[113], fullData[66]^fullData[60], fullData[90]^fullData[8], fullData[117]^fullData[53], fullData[62]^fullData[70], fullData[108]-fullData[74], fullData[59]-fullData[114], fullData[83]+fullData[2], fullData[39]^fullData[87], fullData[116]^fullData[98], fullData[80]+fullData[32])
			return  /*line AWYsLdSf.go:1*/ string(data)
		}())
	}
	cVj_Mj6s := ktsi1_SQ[0]
	eviEA3sp := ktsi1_SQ[1]
	olYMzSWK := ktsi1_SQ[2]
	pbLLcfzS := ktsi1_SQ[3]
	ezgSKu0t :=  /*line HEJOoNZF.go:1*/ strings.Split(ktsi1_SQ[4],  /*line zvzU2r1G.go:1*/ func() string {
		data :=  /*line zvzU2r1G.go:1*/ []byte("y")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line zvzU2r1G.go:1*/ byte(i) +  /*line zvzU2r1G.go:1*/ byte(positions[i]^positions[i+1]) + 179
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line zvzU2r1G.go:1*/ string(data)
	}())
	mzjt9rLK :=  /*line Gf5y3oXv.go:1*/ func() string {
		seed :=  /*line Gf5y3oXv.go:1*/ byte(207)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Gf5y3oXv.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/  /*line Gf5y3oXv.go:1*/ fnc(157)(9)(4)(12)(224)(22)(15)(203)(54)(232)(22)(250)(240)
		return  /*line Gf5y3oXv.go:1*/ string(data)
	}()
	nTXZkyTM, cSW1wSO3 :=  /*line ISY1RPdr.go:1*/ g4rnrSNn.dHINeXgX(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line fEdFpV4z.go:1*/ shim.Error( /*line IKjzx_dv.go:1*/ cSW1wSO3.Error())
	}
	                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    

	aZUvqwea, cSW1wSO3 :=  /*line HtHzs5sC.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line scdlsvMZ.go:1*/ shim.Error( /*line okAaOx2z.go:1*/ cSW1wSO3.Error())
	}

	tbtwirHa := eviEA3sp +  /*line JtK3Gw4F.go:1*/ func() string {
		data :=  /*line JtK3Gw4F.go:1*/ []byte("d")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line JtK3Gw4F.go:1*/ byte(i) +  /*line JtK3Gw4F.go:1*/ byte(positions[i]^positions[i+1]) + 53
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line JtK3Gw4F.go:1*/ string(data)
	}() + aZUvqwea
	                                
	lYWrPDrH, cSW1wSO3 :=  /*line BK0KksRN.go:1*/ n4c7mRtG.GetState(tbtwirHa)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line HoLHvxgp.go:1*/ func() string {
			key :=  /*line HoLHvxgp.go:1*/ []byte("\x1c\xf6\xeem\r'\xa7\xdeno\xaco\x15@\xf0\a\xbe\xad\xc5;\bb\x04\x92\xccp\xd8\x10T\xc9_\xd8]\xc4\xef\x11$Q\xbc\x94\xbav\x89`\x9d\xa9q\xeb\xe5\xca\xe4\x17M\xb7\xa4x\x85")
			data :=  /*line HoLHvxgp.go:1*/ []byte("gԫ\x1f\u007fH\xd5\xfcTM\xea\x0e|,\x95c\x9e٪\x1bk\na\xf1\xa7P\xb9e \xa10\xaa4\xbe\x8eeM>Ҵ\xdc\x19\xfb@\xe9\xc1\x14ˁ\xab\x90v>\xd2\xd0B\xa5")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line HoLHvxgp.go:1*/ string(data)
		}() + eviEA3sp +  /*line TkWgUBF6.go:1*/ func() string {
			var data []byte
			i := 1
			decryptKey := 164
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 1:
					i = 2
					data =  /*line TkWgUBF6.go:1*/ append(data, 194)
				case 3:
					for y := range data {
						data[y] = data[y] -  /*line TkWgUBF6.go:1*/ byte(decryptKey^y)
					}
					i = 0
				case 2:
					data =  /*line TkWgUBF6.go:1*/ append(data, 30)
					i = 3
				}
			}
			return  /*line TkWgUBF6.go:1*/ string(data)
		}()
		return  /*line PeN6mY79.go:1*/ shim.Error(w5VhCLka)
	} else if (lYWrPDrH == nil) && (olYMzSWK == "" || olYMzSWK ==  /*line sBcENJLg.go:1*/ func() string {
		seed :=  /*line sBcENJLg.go:1*/ byte(212)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line sBcENJLg.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line sBcENJLg.go:1*/ fnc(244)
		return  /*line sBcENJLg.go:1*/ string(data)
	}()) {
		w5VhCLka =  /*line P1wNCOz5.go:1*/ func() string {
			fullData :=  /*line P1wNCOz5.go:1*/ []byte("'J\xael!\xd4?w\xbb\x9cRp\v\xf3\xc8\xc7e\xf6\x1e\x16\xed\x91\xdc\xeen\xf9\xffGМ`\xa8ܑ\xe9怏U\xb3\xa8M&\x93\x04\xe9\x10T\x8c\n\xce%\x19E1\x96h\x95\x92\x80S\xc9\x01bB\x16y\x91!\xb6݄\x9cc\xa7u\x00Zm\xb6ē\xf2\xc5i\xe8\xf5\xa9\xdex)\x9c\xe2\xa9\x13]\xa8\xbcv\xe9\xf7&e\x9d\xe0g0\xa2\xfc\xff\x88\xb3E\xe2) [9\xb8\xb2\x11e\xbe\x96\xbeg\xe1g\xc5\\G\xff>|\xe1G\xb9o\x80\n\xbc\vV\xc8\xd8.x\x90\xaa<F\xc0Zw(\x90\xa8\xf3\xc03\x952\x82\xb4LR\xb1\xa1w\x8c\x02\x01\f\xf3\xc0\xfc\xa2\xac")
			var data []byte
			data =  /*line P1wNCOz5.go:1*/ append(data, fullData[24]-fullData[173], fullData[67]-fullData[137], fullData[111]+fullData[58], fullData[113]^fullData[147], fullData[1]+fullData[154], fullData[90]+fullData[150], fullData[50]-fullData[129], fullData[70]^fullData[26], fullData[75]+fullData[83], fullData[166]^fullData[43], fullData[85]^fullData[97], fullData[116]-fullData[13], fullData[61]+fullData[9], fullData[66]+fullData[74], fullData[131]-fullData[37], fullData[57]+fullData[28], fullData[2]-fullData[149], fullData[123]-fullData[114], fullData[59]+fullData[34], fullData[27]-fullData[5], fullData[151]-fullData[41], fullData[76]^fullData[84], fullData[157]-fullData[71], fullData[127]-fullData[25], fullData[79]^fullData[55], fullData[99]^fullData[36], fullData[54]+fullData[64], fullData[162]^fullData[176], fullData[155]+fullData[88], fullData[121]-fullData[17], fullData[115]-fullData[177], fullData[125]^fullData[130], fullData[30]-fullData[175], fullData[11]+fullData[86], fullData[120]^fullData[168], fullData[172]+fullData[95], fullData[16]-fullData[100], fullData[18]+fullData[135], fullData[170]+fullData[63], fullData[103]+fullData[21], fullData[10]-fullData[161], fullData[47]+fullData[82], fullData[4]-fullData[158], fullData[171]+fullData[56], fullData[169]^fullData[104], fullData[81]-fullData[145], fullData[98]+fullData[23], fullData[14]-fullData[40], fullData[163]+fullData[174], fullData[107]-fullData[159], fullData[3]-fullData[164], fullData[48]-fullData[101], fullData[144]+fullData[33], fullData[19]-fullData[96], fullData[73]-fullData[109], fullData[46]^fullData[106], fullData[167]^fullData[143], fullData[78]+fullData[62], fullData[124]-fullData[38], fullData[136]-fullData[53], fullData[72]-fullData[133], fullData[118]+fullData[93], fullData[141]-fullData[31], fullData[0]-fullData[80], fullData[148]+fullData[8], fullData[51]-fullData[119], fullData[128]^fullData[69], fullData[108]-fullData[32], fullData[49]^fullData[89], fullData[6]^fullData[77], fullData[165]-fullData[126], fullData[29]^fullData[45], fullData[20]^fullData[110], fullData[42]-fullData[39], fullData[156]^fullData[22], fullData[140]-fullData[91], fullData[35]+fullData[138], fullData[142]^fullData[117], fullData[105]+fullData[12], fullData[112]^fullData[102], fullData[122]-fullData[152], fullData[153]-fullData[65], fullData[160]^fullData[134], fullData[139]-fullData[87], fullData[7]-fullData[44], fullData[146]-fullData[94], fullData[15]-fullData[60], fullData[68]+fullData[52], fullData[132]+fullData[92])
			return  /*line P1wNCOz5.go:1*/ string(data)
		}() + eviEA3sp +  /*line G2oXEkRb.go:1*/ func() string {
			key :=  /*line G2oXEkRb.go:1*/ []byte("&\u007f")
			data :=  /*line G2oXEkRb.go:1*/ []byte("\x04\x02")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line G2oXEkRb.go:1*/ string(data)
		}()
		return  /*line qaP1kENa.go:1*/ shim.Error(w5VhCLka)
	}

	b4dMB1zL, cSW1wSO3 :=  /*line GMf4t7Lo.go:1*/ jmSY84co(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line Ox_A8Dr6.go:1*/ shim.Error( /*line MZaIM36u.go:1*/ cSW1wSO3.Error())
	}

	zzhdvyqZ := &DatasetMetadataResponse{}
	cSW1wSO3 =  /*line zFLTQbzG.go:1*/ json.Unmarshal( /*line zyvk6O9S.go:1*/ []byte(b4dMB1zL), zzhdvyqZ)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line tUiFexmM.go:1*/ func() string {
			data :=  /*line tUiFexmM.go:1*/ []byte("\x81\"\xa4\x038kr0piFEQNed\x9a/o]/Tc\xc2[a{\x8a\xb1ON")
			positions := [...]byte{9, 20, 25, 5, 27, 28, 4, 8, 16, 23, 2, 0, 19, 7, 16, 23, 16, 21, 7, 0, 17, 26, 12, 2, 24, 0, 23, 7, 28, 13, 23, 11, 3, 16}
			for i := 0; i < 34; i += 2 {
				localKey :=  /*line tUiFexmM.go:1*/ byte(i) +  /*line tUiFexmM.go:1*/ byte(positions[i]^positions[i+1]) + 240
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line tUiFexmM.go:1*/ string(data)
		}() + eviEA3sp +  /*line yrPRY9Tt.go:1*/ func() string {
			var data []byte
			i := 2
			decryptKey := 179
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					data =  /*line yrPRY9Tt.go:1*/ append(data, 112)
					i = 3
				case 1:
					for y := range data {
						data[y] = data[y] +  /*line yrPRY9Tt.go:1*/ byte(decryptKey^y)
					}
					i = 0
				case 3:
					data =  /*line yrPRY9Tt.go:1*/ append(data, 202)
					i = 1
				}
			}
			return  /*line yrPRY9Tt.go:1*/ string(data)
		}()
		return  /*line cBWpFTzM.go:1*/ shim.Error(w5VhCLka)
	}

	                                                                  
	if  /*line AirpfbhS.go:1*/ nSRaDwQg(ezgSKu0t, zzhdvyqZ.CustomAccessRights) {
		 /*line P3B91C1W.go:1*/ fmt.Println( /*line TYXmcTXH.go:1*/ func() string {
			seed :=  /*line TYXmcTXH.go:1*/ byte(67)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line TYXmcTXH.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/  /*line TYXmcTXH.go:1*/ fnc(132)(55)(110)(222)(185)(115)(232)(199)(134)(31)(47)(25)(117)(252)(246)(237)(213)(168)(3)(71)(144)(32)(66)(146)(36)(245)(60)(111)(220)(185)(126)(251)
			return  /*line TYXmcTXH.go:1*/ string(data)
		}())
	} else {
		w5VhCLka =  /*line KgXwtv5l.go:1*/ func() string {
			var data []byte
			i := 1
			decryptKey := 30
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					i = 14
					data =  /*line KgXwtv5l.go:1*/ append(data, "X\t\x1e"...,
					)
				case 12:
					i = 8
					data =  /*line KgXwtv5l.go:1*/ append(data, "E\x03_H"...,
					)
				case 5:
					i = 15
					data =  /*line KgXwtv5l.go:1*/ append(data, "Zke"...,
					)
				case 15:
					data =  /*line KgXwtv5l.go:1*/ append(data, "z|*t"...,
					)
					i = 13
				case 0:
					i = 9
					data =  /*line KgXwtv5l.go:1*/ append(data, "WT\x1aZ"...,
					)
				case 7:
					i = 12
					data =  /*line KgXwtv5l.go:1*/ append(data, "qUOO"...,
					)
				case 3:
					i = 0
					data =  /*line KgXwtv5l.go:1*/ append(data, 75)
				case 1:
					data =  /*line KgXwtv5l.go:1*/ append(data, "W\x0f"...,
					)
					i = 10
				case 14:
					data =  /*line KgXwtv5l.go:1*/ append(data, 7)
					i = 7
				case 10:
					i = 11
					data =  /*line KgXwtv5l.go:1*/ append(data, "k]"...,
					)
				case 9:
					data =  /*line KgXwtv5l.go:1*/ append(data, "WVSD"...,
					)
					i = 6
				case 11:
					i = 4
					data =  /*line KgXwtv5l.go:1*/ append(data, "ZF"...,
					)
				case 8:
					i = 3
					data =  /*line KgXwtv5l.go:1*/ append(data, 77)
				case 6:
					data =  /*line KgXwtv5l.go:1*/ append(data, "C\x11@"...,
					)
					i = 5
				case 13:
					i = 2
					for y := range data {
						data[y] = data[y] ^  /*line KgXwtv5l.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return  /*line KgXwtv5l.go:1*/ string(data)
		}()
		return  /*line jihepK4c.go:1*/ shim.Error(w5VhCLka)
	}

	                                                                               
	_ra7bbD0, cSW1wSO3 :=  /*line Cvlu03SJ.go:1*/ g4rnrSNn.y4dHt0L7(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line wlK9qCRI.go:1*/ shim.Error( /*line kEqLudwt.go:1*/ cSW1wSO3.Error())
	}

	b2UcdGWD, cSW1wSO3 :=  /*line YKXrlsT1.go:1*/ g4rnrSNn.hDKHUUS5(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line igzzzE45.go:1*/ shim.Error( /*line nUmpUvUz.go:1*/ cSW1wSO3.Error())
	}

	sM8l5NjA, cSW1wSO3 :=  /*line AQQ11cPB.go:1*/ g4rnrSNn.lu_3VIbC(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line RRuqjbMR.go:1*/ shim.Error( /*line QyEnssK6.go:1*/ cSW1wSO3.Error())
	}

	if sM8l5NjA ==  /*line buNvW2LS.go:1*/ func() string {
		key :=  /*line buNvW2LS.go:1*/ []byte("I\xa1{\x80\xec")
		data :=  /*line buNvW2LS.go:1*/ []byte("/\xc0\x17\xf3\x89")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line buNvW2LS.go:1*/ string(data)
	}() {
		vWNUn3N4, cSW1wSO3 =  /*line C8R_HBhn.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
		if cSW1wSO3 != nil {
			return  /*line FN9_5tSo.go:1*/ shim.Error( /*line oBEQ7Puw.go:1*/ cSW1wSO3.Error())
		}

	} else if sM8l5NjA ==  /*line YIkaNxEZ.go:1*/ func() string {
		key :=  /*line YIkaNxEZ.go:1*/ []byte("\xa2f\x1b\xa5")
		data :=  /*line YIkaNxEZ.go:1*/ []byte("\xd6\x14n\xc0")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line YIkaNxEZ.go:1*/ string(data)
	}() {
		vWNUn3N4, cSW1wSO3 =  /*line RiQZJj3w.go:1*/ g4rnrSNn.rae8UYAQ(n4c7mRtG)
		if cSW1wSO3 != nil {
			return  /*line JFYSdFHc.go:1*/ shim.Error( /*line SFjbglbi.go:1*/ cSW1wSO3.Error())
		}

	} else {
		return  /*line HrydKAx6.go:1*/ shim.Error( /*line PNO57apd.go:1*/ func() string {
			data :=  /*line PNO57apd.go:1*/ []byte("\xe9nknҘnj\xf9\x88t\xf4\xce\xc2\x05\xb3u\xdc\v\x1f")
			positions := [...]byte{5, 0, 14, 8, 5, 8, 14, 18, 19, 12, 14, 9, 7, 0, 7, 13, 5, 11, 5, 0, 17, 14, 15, 12, 15, 7, 5, 17, 13, 4}
			for i := 0; i < 30; i += 2 {
				localKey :=  /*line PNO57apd.go:1*/ byte(i) +  /*line PNO57apd.go:1*/ byte(positions[i]^positions[i+1]) + 53
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line PNO57apd.go:1*/ string(data)
		}())
	}

	aRL4DUka, cSW1wSO3 :=  /*line rgUA8dbc.go:1*/ g4rnrSNn.zSzRgzTF(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line zZ7CBNle.go:1*/ shim.Error( /*line _JkdWUXC.go:1*/ cSW1wSO3.Error())
	}

	                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 

	aZuYW6KY := RequestAccess{mzjt9rLK, cVj_Mj6s, eviEA3sp, zzhdvyqZ.UsernameOfProvider, zzhdvyqZ.OrgOfProvider, zzhdvyqZ.DatasetDescription, olYMzSWK,
		ezgSKu0t, _ra7bbD0, b2UcdGWD, vWNUn3N4, aZUvqwea, aRL4DUka, nTXZkyTM, pbLLcfzS}
	hvFCuRzl, cSW1wSO3 :=  /*line bYzXY6g4.go:1*/ json.Marshal(aZuYW6KY)
	if cSW1wSO3 != nil {
		return  /*line WDZWmGig.go:1*/ shim.Error( /*line ztzdKOcS.go:1*/ cSW1wSO3.Error())
	}

	                        
	cSW1wSO3 =  /*line tS71sRDl.go:1*/ n4c7mRtG.PutState(cVj_Mj6s, hvFCuRzl)
	if cSW1wSO3 != nil {
		return  /*line dQ3e8C8p.go:1*/ shim.Error( /*line HNly1aCN.go:1*/ cSW1wSO3.Error())
	}

	xHTOVzVs :=  /*line CupUSNJf.go:1*/ func() string {
		seed :=  /*line CupUSNJf.go:1*/ byte(159)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line CupUSNJf.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/  /*line CupUSNJf.go:1*/ fnc(198)(10)(26)(251)(164)(90)(231)(24)(244)(16)(246)(15)(170)(64)(27)(175)(90)(249)(229)(23)(254)(238)(13)(188)(98)
		return  /*line CupUSNJf.go:1*/ string(data)
	}() + eviEA3sp +  /*line C2a8kEau.go:1*/ func() string {
		fullData :=  /*line C2a8kEau.go:1*/ []byte("\xc0\x02^w\xd7=\xc3\xc94\x1b\xe9+q\"\x84\x80Z)\x9c\xfa\x15\x9d\xf1\xfd<\xf9\xaa\xa1\x1dn\xa9\x95Pa\x94\xf1\xa4\x90")
		var data []byte
		data =  /*line C2a8kEau.go:1*/ append(data, fullData[36]-fullData[14], fullData[31]^fullData[23], fullData[37]^fullData[22], fullData[28]^fullData[29], fullData[10]-fullData[7], fullData[13]-fullData[0], fullData[17]+fullData[24], fullData[15]-fullData[9], fullData[19]^fullData[34], fullData[30]+fullData[3], fullData[12]+fullData[1], fullData[11]^fullData[2], fullData[27]^fullData[6], fullData[26]-fullData[5], fullData[16]-fullData[35], fullData[20]^fullData[33], fullData[21]+fullData[4], fullData[18]^fullData[25], fullData[32]^fullData[8])
		return  /*line C2a8kEau.go:1*/ string(data)
	}()
	q8ZF8zaz :=  /*line mNsdnNiO.go:1*/ []byte(xHTOVzVs)
	return  /*line GuoGJIBA.go:1*/ shim.Success(q8ZF8zaz)
}

func (g4rnrSNn *G1Y_7pz_) islTrX3n(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var cSW1wSO3 error
	var vWNUn3N4 string
	var w5VhCLka string

	if  /*line G4VrFslQ.go:1*/ len(ktsi1_SQ) != 5 {
		return  /*line HRPUQA_M.go:1*/ shim.Error( /*line EFWZM4Lw.go:1*/ func() string {
			fullData :=  /*line EFWZM4Lw.go:1*/ []byte("\x00\xb1t\xdcz\xde\xc6\xf2\xbaTS\xd7\\ \xbc\x17~Z\x96\xa2\xffF\x19\x8a]\x91\xf3\x9c\x12\xffØ.Q\xde3V\x87\xf9\xb9\xad\xec\xb0K\x85\bf\xc4C\xe6L\x90M(\xdc\r\xa3f\xd5߽Ʉ_\xefj*C\xec\a\xd0@\xc9\xfb\xdf\x1b\x96\x10J\xab\xbe\x05\u007fm\xd5.\xd1 U\x8f=߂\xa7\xb3WC?\x9c²*\x1a^\xe0s&\x89I\x9d \x87$9L\x91dj\x92\xb2\xe8%\xb2\xafrB2A\"\xe6o\x17\x1cb^/\x9b\r")
			var data []byte
			data =  /*line EFWZM4Lw.go:1*/ append(data, fullData[18]^fullData[59], fullData[90]^fullData[10], fullData[43]-fullData[84], fullData[127]^fullData[13], fullData[119]+fullData[8], fullData[65]+fullData[20], fullData[36]-fullData[7], fullData[125]+fullData[5], fullData[93]^fullData[129], fullData[79]-fullData[113], fullData[75]+fullData[114], fullData[86]-fullData[12], fullData[22]^fullData[2], fullData[37]-fullData[128], fullData[52]-fullData[74], fullData[41]^fullData[31], fullData[71]+fullData[104], fullData[70]-fullData[92], fullData[66]^fullData[63], fullData[0]+fullData[83], fullData[110]-fullData[80], fullData[132]+fullData[108], fullData[38]-fullData[111], fullData[1]^fullData[25], fullData[3]+fullData[23], fullData[109]-fullData[32], fullData[121]^fullData[95], fullData[116]+fullData[14], fullData[47]+fullData[19], fullData[9]-fullData[91], fullData[73]+fullData[105], fullData[115]-fullData[85], fullData[101]^fullData[103], fullData[123]^fullData[6], fullData[137]^fullData[133], fullData[54]+fullData[118], fullData[45]^fullData[53], fullData[99]+fullData[42], fullData[76]^fullData[26], fullData[126]^fullData[48], fullData[135]^fullData[17], fullData[67]-fullData[34], fullData[87]-fullData[40], fullData[15]-fullData[56], fullData[81]-fullData[94], fullData[33]-fullData[68], fullData[39]+fullData[60], fullData[49]^fullData[107], fullData[61]-fullData[134], fullData[77]+fullData[88], fullData[30]+fullData[16], fullData[120]-fullData[44], fullData[29]^fullData[98], fullData[72]+fullData[27], fullData[78]-fullData[11], fullData[100]-fullData[97], fullData[122]-fullData[62], fullData[21]-fullData[106], fullData[58]-fullData[51], fullData[82]-fullData[69], fullData[117]^fullData[102], fullData[124]-fullData[55], fullData[50]+fullData[131], fullData[28]^fullData[46], fullData[4]+fullData[64], fullData[24]^fullData[35], fullData[112]^fullData[96], fullData[89]-fullData[130], fullData[136]-fullData[57])
			return  /*line EFWZM4Lw.go:1*/ string(data)
		}())
	}
	cVj_Mj6s := ktsi1_SQ[0]
	eviEA3sp := ktsi1_SQ[1]
	olYMzSWK := ktsi1_SQ[2]
	pbLLcfzS := ktsi1_SQ[3]
	ezgSKu0t :=  /*line Yes0h6qA.go:1*/ strings.Split(ktsi1_SQ[4],  /*line SN8lm2C8.go:1*/ func() string {
		key :=  /*line SN8lm2C8.go:1*/ []byte("\x88")
		data :=  /*line SN8lm2C8.go:1*/ []byte("\xa4")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line SN8lm2C8.go:1*/ string(data)
	}())
	mzjt9rLK :=  /*line aBgTxZaz.go:1*/ func() string {
		data :=  /*line aBgTxZaz.go:1*/ []byte("Neq}~silR[vozedAc>eoQ")
		positions := [...]byte{17, 5, 0, 20, 12, 3, 12, 7, 5, 7, 0, 9, 6, 3, 3, 4, 19, 19, 17, 19, 9, 20}
		for i := 0; i < 22; i += 2 {
			localKey :=  /*line aBgTxZaz.go:1*/ byte(i) +  /*line aBgTxZaz.go:1*/ byte(positions[i]^positions[i+1]) + 246
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line aBgTxZaz.go:1*/ string(data)
	}()

	aZUvqwea, cSW1wSO3 :=  /*line WNHPMnzL.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line o4CN1XHk.go:1*/ shim.Error( /*line Y0rmElE8.go:1*/ cSW1wSO3.Error())
	}

	syHhHGpa := eviEA3sp +  /*line QXiOzXRX.go:1*/ func() string {
		key :=  /*line QXiOzXRX.go:1*/ []byte("\xcb")
		data :=  /*line QXiOzXRX.go:1*/ []byte("\xe4")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line QXiOzXRX.go:1*/ string(data)
	}() + aZUvqwea
	tQcQN8RB, cSW1wSO3 :=  /*line z8dzRuOF.go:1*/ n4c7mRtG.GetState(syHhHGpa)	                    
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line UX41xcLb.go:1*/ func() string {
			var data []byte
			i := 17
			decryptKey := 183
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					data =  /*line UX41xcLb.go:1*/ append(data, 232)
					i = 18
				case 20:
					i = 21
					data =  /*line UX41xcLb.go:1*/ append(data, "\x91\x8d\x93\xc2"...,
					)
				case 13:
					data =  /*line UX41xcLb.go:1*/ append(data, "\x8c\x8c\xd7"...,
					)
					i = 15
				case 3:
					i = 1
					data =  /*line UX41xcLb.go:1*/ append(data, "\x90\x9a"...,
					)
				case 0:
					data =  /*line UX41xcLb.go:1*/ append(data, "\xa4\xb0\xaa\xad"...,
					)
					i = 10
				case 5:
					data =  /*line UX41xcLb.go:1*/ append(data, "\x8d\x82\x86"...,
					)
					i = 13
				case 1:
					data =  /*line UX41xcLb.go:1*/ append(data, "\x94\x93"...,
					)
					i = 16
				case 12:
					i = 14
					data =  /*line UX41xcLb.go:1*/ append(data, "\xa1\xbf\xec"...,
					)
				case 15:
					i = 3
					data =  /*line UX41xcLb.go:1*/ append(data, "\x82\x9a\xd4"...,
					)
				case 19:
					data =  /*line UX41xcLb.go:1*/ append(data, "\x96\x8a\xae"...,
					)
					i = 9
				case 16:
					i = 2
					data =  /*line UX41xcLb.go:1*/ append(data, "\x94ޜ\x89"...,
					)
				case 7:
					data =  /*line UX41xcLb.go:1*/ append(data, "\xa0\xb7\xa5"...,
					)
					i = 4
				case 21:
					data =  /*line UX41xcLb.go:1*/ append(data, "\xd5̫"...,
					)
					i = 5
				case 10:
					i = 12
					data =  /*line UX41xcLb.go:1*/ append(data, "\xaf\xe0\xa9"...,
					)
				case 14:
					data =  /*line UX41xcLb.go:1*/ append(data, "\xbf\xa2\xac"...,
					)
					i = 6
				case 18:
					data =  /*line UX41xcLb.go:1*/ append(data, "\xb3\xb7\xa1\xb5"...,
					)
					i = 7
				case 17:
					i = 20
					data =  /*line UX41xcLb.go:1*/ append(data, "\x9cĠ\x96"...,
					)
				case 11:
					i = 8
					for y := range data {
						data[y] = data[y] ^  /*line UX41xcLb.go:1*/ byte(decryptKey^y)
					}
				case 9:
					data =  /*line UX41xcLb.go:1*/ append(data, 188)
					i = 0
				case 2:
					i = 19
					data =  /*line UX41xcLb.go:1*/ append(data, "\x8f\x92"...,
					)
				case 4:
					data =  /*line UX41xcLb.go:1*/ append(data, "\xea\xff"...,
					)
					i = 11
				}
			}
			return  /*line UX41xcLb.go:1*/ string(data)
		}() + syHhHGpa +  /*line mxH1M8A5.go:1*/ func() string {
			fullData :=  /*line mxH1M8A5.go:1*/ []byte("p\r\x1c\x06")
			var data []byte
			data =  /*line mxH1M8A5.go:1*/ append(data, fullData[3]+fullData[2], fullData[1]+fullData[0])
			return  /*line mxH1M8A5.go:1*/ string(data)
		}()
		return  /*line lEQrHget.go:1*/ shim.Error(w5VhCLka)
	} else if tQcQN8RB != nil {
		a8ZPVup7 := AuthorizedUsers{}
		cSW1wSO3 =  /*line YOb_cSgQ.go:1*/ json.Unmarshal(tQcQN8RB, &a8ZPVup7)	                               
		if cSW1wSO3 != nil {
			return  /*line Qw0n6wKZ.go:1*/ shim.Error( /*line MzLKFnyb.go:1*/ cSW1wSO3.Error())
		}
		if  /*line SUn9lX9y.go:1*/ nSRaDwQg(ezgSKu0t, a8ZPVup7.Users.CustomAccessRights) {
			 /*line bQfP3D4J.go:1*/ fmt.Println( /*line xELqK5vV.go:1*/ func() string {
				seed :=  /*line xELqK5vV.go:1*/ byte(113)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data =  /*line xELqK5vV.go:1*/ append(data, x^seed); seed += x; return fnc }
				 /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/  /*line xELqK5vV.go:1*/ fnc(48)(209)(2)(6)(21)(255)(252)(227)(12)(13)(227)(73)(209)(246)(10)(247)(21)(226)(81)(163)(6)(8)(22)(250)(240)(83)(180)(19)(234)(31)(226)(11)
				return  /*line xELqK5vV.go:1*/ string(data)
			}())
		} else {
			w5VhCLka =  /*line P75CPL6n.go:1*/ func() string {
				data :=  /*line P75CPL6n.go:1*/ []byte("|\"PrFoM\"V\".PRn\x11@vs\t\u007fo0\"aZP``q\tvR\xc5h>\x99\x1b}")
				positions := [...]byte{17, 28, 12, 26, 35, 32, 24, 22, 27, 21, 11, 10, 28, 30, 4, 25, 19, 8, 36, 14, 22, 22, 14, 0, 25, 6, 31, 11, 26, 6, 2, 32, 29, 27, 15, 31, 27, 8, 35, 0, 11, 10, 16, 2, 34, 24, 18, 22}
				for i := 0; i < 48; i += 2 {
					localKey :=  /*line P75CPL6n.go:1*/ byte(i) +  /*line P75CPL6n.go:1*/ byte(positions[i]^positions[i+1]) + 247
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return  /*line P75CPL6n.go:1*/ string(data)
			}()
			return  /*line s5s1RTMK.go:1*/ shim.Error(w5VhCLka)
		}
	}

	nTXZkyTM, cSW1wSO3 :=  /*line ld59iZY0.go:1*/ g4rnrSNn.dHINeXgX(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line jkqtbKtB.go:1*/ shim.Error( /*line GDttfk4B.go:1*/ cSW1wSO3.Error())
	}

	                                                                               
	_ra7bbD0, cSW1wSO3 :=  /*line sc1hwGVZ.go:1*/ g4rnrSNn.y4dHt0L7(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line bwzOGJHO.go:1*/ shim.Error( /*line oSnrxhtz.go:1*/ cSW1wSO3.Error())
	}

	b2UcdGWD, cSW1wSO3 :=  /*line P9qvPrVj.go:1*/ g4rnrSNn.hDKHUUS5(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line xXuPYkgb.go:1*/ shim.Error( /*line dzCn3Med.go:1*/ cSW1wSO3.Error())
	}

	sM8l5NjA, cSW1wSO3 :=  /*line l6FVme0F.go:1*/ g4rnrSNn.lu_3VIbC(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line zd3eacE1.go:1*/ shim.Error( /*line ElCwcqTg.go:1*/ cSW1wSO3.Error())
	}

	if sM8l5NjA ==  /*line Mo9SYvuS.go:1*/ func() string {
		data :=  /*line Mo9SYvuS.go:1*/ []byte("f˧Ɠ")
		positions := [...]byte{4, 1, 3, 2, 2, 4}
		for i := 0; i < 6; i += 2 {
			localKey :=  /*line Mo9SYvuS.go:1*/ byte(i) +  /*line Mo9SYvuS.go:1*/ byte(positions[i]^positions[i+1]) + 201
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line Mo9SYvuS.go:1*/ string(data)
	}() {
		vWNUn3N4, cSW1wSO3 =  /*line qAZ_CrIt.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
		if cSW1wSO3 != nil {
			return  /*line B17qG0m6.go:1*/ shim.Error( /*line JY2Rm6dV.go:1*/ cSW1wSO3.Error())
		}

	} else if sM8l5NjA ==  /*line w14K85V_.go:1*/ func() string {
		data :=  /*line w14K85V_.go:1*/ []byte("\xa4r\xf0\xf0")
		positions := [...]byte{0, 3, 2, 2, 2, 0}
		for i := 0; i < 6; i += 2 {
			localKey :=  /*line w14K85V_.go:1*/ byte(i) +  /*line w14K85V_.go:1*/ byte(positions[i]^positions[i+1]) + 190
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line w14K85V_.go:1*/ string(data)
	}() {
		vWNUn3N4, cSW1wSO3 =  /*line RAn4OUIm.go:1*/ g4rnrSNn.rae8UYAQ(n4c7mRtG)
		if cSW1wSO3 != nil {
			return  /*line tFS7IFtB.go:1*/ shim.Error( /*line As9ReS6E.go:1*/ cSW1wSO3.Error())
		}

	} else {
		return  /*line K4fYCVCX.go:1*/ shim.Error( /*line A8KTFY6B.go:1*/ func() string {
			fullData :=  /*line A8KTFY6B.go:1*/ []byte("\xf9\xfe\x00<6\x8c\xc9Dh.{\xd6:\r\xbcc\xe2\xd9*\x99\xf2\vvV\x96q\x88\xa2v\xa5\xcc\xf8\x94\xe7\xde\x03e4\xa1t")
			var data []byte
			data =  /*line A8KTFY6B.go:1*/ append(data, fullData[30]^fullData[19], fullData[16]^fullData[5], fullData[6]+fullData[27], fullData[18]^fullData[7], fullData[26]+fullData[33], fullData[38]+fullData[11], fullData[12]+fullData[37], fullData[23]^fullData[22], fullData[29]-fullData[3], fullData[21]+fullData[15], fullData[39]^fullData[2], fullData[36]-fullData[4], fullData[17]^fullData[14], fullData[28]-fullData[1], fullData[34]+fullData[24], fullData[20]+fullData[9], fullData[8]+fullData[13], fullData[10]+fullData[31], fullData[0]-fullData[32], fullData[35]^fullData[25])
			return  /*line A8KTFY6B.go:1*/ string(data)
		}())
	}

	aRL4DUka, cSW1wSO3 :=  /*line uUG4KpP8.go:1*/ g4rnrSNn.zSzRgzTF(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line beA1g6zs.go:1*/ shim.Error( /*line tuyYHLMP.go:1*/ cSW1wSO3.Error())
	}

	                                                   
	thOUuGet, cSW1wSO3 :=  /*line M63WdD_q.go:1*/ fvIQVpVM(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line TzCllkNO.go:1*/ shim.Error( /*line ll3wKz7A.go:1*/ cSW1wSO3.Error())
	}

	islTrX3n := RequestRevokeAccessByUser{mzjt9rLK, cVj_Mj6s, eviEA3sp, thOUuGet, olYMzSWK,
		ezgSKu0t, _ra7bbD0, b2UcdGWD, vWNUn3N4, aZUvqwea, aRL4DUka, nTXZkyTM, pbLLcfzS}
	czwjgpuG, cSW1wSO3 :=  /*line I8RQkrPJ.go:1*/ json.Marshal(islTrX3n)
	if cSW1wSO3 != nil {
		return  /*line K8_Axc8f.go:1*/ shim.Error( /*line H1iWenzF.go:1*/ cSW1wSO3.Error())
	}

	                        
	cSW1wSO3 =  /*line uv_bE1Kl.go:1*/ n4c7mRtG.PutState(cVj_Mj6s, czwjgpuG)
	if cSW1wSO3 != nil {
		return  /*line tRlV7o9W.go:1*/ shim.Error( /*line noI89h4v.go:1*/ cSW1wSO3.Error())
	}

	xHTOVzVs :=  /*line lCaoGdzJ.go:1*/ func() string {
		key :=  /*line lCaoGdzJ.go:1*/ []byte("3^6T*\x96\xbc\xfe\xcb\x1b<\xf8\x87\xff\xc2\xf5\xe2\xf8\xc4\x00@<\xcc\n'\xff\x84f\x15\x87V\x05")
		data :=  /*line lCaoGdzJ.go:1*/ []byte("&\x11?\x1e\xf6ܩx\xa4P)(\xebf\xaf\x80\x83{\xb0 43TZ:u\xdd\rP\xed\xe4\x1b")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line lCaoGdzJ.go:1*/ string(data)
	}() + eviEA3sp +  /*line cjNrsJPL.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 209
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 8:
				data =  /*line cjNrsJPL.go:1*/ append(data, 188)
				i = 3
			case 7:
				data =  /*line cjNrsJPL.go:1*/ append(data, "\xaf\xb3\xac"...,
				)
				i = 8
			case 9:
				i = 1
				data =  /*line cjNrsJPL.go:1*/ append(data, "\xab\xb8"...,
				)
			case 1:
				data =  /*line cjNrsJPL.go:1*/ append(data, "쯫\xaa"...,
				)
				i = 4
			case 6:
				i = 7
				data =  /*line cjNrsJPL.go:1*/ append(data, "\xb1\xb6\xa6\xa8"...,
				)
			case 4:
				i = 6
				data =  /*line cjNrsJPL.go:1*/ append(data, "\xae\xe1"...,
				)
			case 3:
				i = 0
				data =  /*line cjNrsJPL.go:1*/ append(data, 190)
			case 0:
				for y := range data {
					data[y] = data[y] ^  /*line cjNrsJPL.go:1*/ byte(decryptKey^y)
				}
				i = 5
			case 2:
				i = 9
				data =  /*line cjNrsJPL.go:1*/ append(data, "\xe8\xa1"...,
				)
			}
		}
		return  /*line cjNrsJPL.go:1*/ string(data)
	}()
	q8ZF8zaz :=  /*line tSqyB_M_.go:1*/ []byte(xHTOVzVs)
	return  /*line eHTH13uu.go:1*/ shim.Success(q8ZF8zaz)
}

func (g4rnrSNn *G1Y_7pz_) c22g4rkk(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	if  /*line QJCT7A4R.go:1*/ len(ktsi1_SQ) != 2 {
		return  /*line VUyowxbq.go:1*/ shim.Error( /*line v8ZO6Y0a.go:1*/ func() string {
			key :=  /*line v8ZO6Y0a.go:1*/ []byte("\x153a\x83\xc9\x0e\x12\xc2\xd4Z\xab\x93g\x97\xaei\xbbl\xe0\x9d\xfb?\xfc\xf0\x92#\x9fs\xbd\x8cDv\x13h+x\xb5\x90\x18R\nd\xdbo7\xcevr\xe6\x10\xb7\xff\xb7\xe3\x19&\xf6\xdc_ݕ4\xc4e;")
			data :=  /*line v8ZO6Y0a.go:1*/ []byte("4;\x02\xec\xa9dS\xa1\xa0\xc6\xc3\xe2\x06˷\te\x03\x86\x83f3k\x85\xdbB\xcf\x01\xb6\xa2\xdc\xcfe\b:\xeb\xbf\xd9V\x15\x16\xceS\xb1\x1b\x97\xfb\x03\u007fc\xbdJ\x8d=\r\xfa~\x8d\x0e\x88\xde@\x9d\b5")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line v8ZO6Y0a.go:1*/ string(data)
		}())
	}

	hSHnto8d := ktsi1_SQ[0]
	tIeFhn0D := ktsi1_SQ[1]

	w8EOizcx, cSW1wSO3 :=  /*line Ah4vMjKw.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line Qb_F1xls.go:1*/ shim.Error( /*line AFeQKxcl.go:1*/ cSW1wSO3.Error())
	}

	mzjt9rLK :=  /*line E095rbFE.go:1*/ func() string {
		key :=  /*line E095rbFE.go:1*/ []byte("\xc7'\xeby\x8e- \xd3\xff\x15qVh\xf5=")
		data :=  /*line E095rbFE.go:1*/ []byte("\b\x9c_\xe1\xfd\x9f\x89Mdy\xc6\xc9\xcdg\xb0")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line E095rbFE.go:1*/ string(data)
	}()
	qo8GOXik := RequestAccess{}
	mGZijNgp, cSW1wSO3 :=  /*line FYjEEL95.go:1*/ n4c7mRtG.GetState(hSHnto8d)	                                      
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line BR92IKBq.go:1*/ func() string {
			var data []byte
			i := 8
			decryptKey := 177
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 1:
					data =  /*line BR92IKBq.go:1*/ append(data, "//"...,
					)
					i = 18
				case 5:
					i = 16
					data =  /*line BR92IKBq.go:1*/ append(data, 36)
				case 14:
					i = 6
					data =  /*line BR92IKBq.go:1*/ append(data, 202)
				case 15:
					data =  /*line BR92IKBq.go:1*/ append(data, "\xfc\xff\xfc\xff"...,
					)
					i = 12
				case 6:
					data =  /*line BR92IKBq.go:1*/ append(data, "\x16\x15!\xce"...,
					)
					i = 19
				case 8:
					data =  /*line BR92IKBq.go:1*/ append(data, 54)
					i = 9
				case 10:
					i = 11
					data =  /*line BR92IKBq.go:1*/ append(data, 18)
				case 12:
					data =  /*line BR92IKBq.go:1*/ append(data, "\x12\x13\xbd\x04"...,
					)
					i = 3
				case 2:
					i = 14
					data =  /*line BR92IKBq.go:1*/ append(data, " \x18"...,
					)
				case 18:
					data =  /*line BR92IKBq.go:1*/ append(data, "\xe0\xed"...,
					)
					i = 4
				case 3:
					data =  /*line BR92IKBq.go:1*/ append(data, "\x02\x06\xb1"...,
					)
					i = 13
				case 11:
					data =  /*line BR92IKBq.go:1*/ append(data, "\x17\f"...,
					)
					i = 7
				case 4:
					data =  /*line BR92IKBq.go:1*/ append(data, "\xd6\xf7\x13 "...,
					)
					i = 5
				case 19:
					data =  /*line BR92IKBq.go:1*/ append(data, "\x15\t"...,
					)
					i = 10
				case 17:
					i = 1
					data =  /*line BR92IKBq.go:1*/ append(data, 49)
				case 9:
					i = 17
					data =  /*line BR92IKBq.go:1*/ append(data, "\xde\xfe,"...,
					)
				case 7:
					i = 15
					data =  /*line BR92IKBq.go:1*/ append(data, "\x1b\x19\xc6"...,
					)
				case 13:
					for y := range data {
						data[y] = data[y] +  /*line BR92IKBq.go:1*/ byte(decryptKey^y)
					}
					i = 0
				case 16:
					data =  /*line BR92IKBq.go:1*/ append(data, "\x1a\x1a\xcb"...,
					)
					i = 2
				}
			}
			return  /*line BR92IKBq.go:1*/ string(data)
		}() + hSHnto8d +  /*line u7cpu_0c.go:1*/ func() string {
			fullData :=  /*line u7cpu_0c.go:1*/ []byte("m\x10(\xfa")
			var data []byte
			data =  /*line u7cpu_0c.go:1*/ append(data, fullData[2]+fullData[3], fullData[0]^fullData[1])
			return  /*line u7cpu_0c.go:1*/ string(data)
		}()
		return  /*line IXqCnDUn.go:1*/ shim.Error(w5VhCLka)
	} else if mGZijNgp == nil {
		w5VhCLka =  /*line xYhWrwWH.go:1*/ func() string {
			data :=  /*line xYhWrwWH.go:1*/ []byte("{\xc2e,\xe9or\":\x91\xaf\xba{m ѫj\xdcexi\x94[s\xe9\x03V\xbc\xaee@\xe71\x88o\x95nv\x9dn\xcc\xd4et\xbf\xa7&f6\x94\xd0nqxewtaD")
			positions := [...]byte{38, 56, 18, 2, 15, 25, 4, 13, 42, 47, 49, 12, 54, 56, 58, 26, 46, 54, 56, 41, 50, 4, 51, 58, 22, 36, 15, 3, 49, 3, 11, 16, 58, 2, 17, 23, 10, 17, 33, 1, 48, 13, 38, 9, 29, 28, 58, 39, 42, 49, 9, 18, 41, 26, 15, 15, 12, 45, 23, 2, 46, 41, 32, 42, 11, 52, 11, 40, 1, 18, 37, 12, 58, 31, 37, 27, 54, 28, 25, 48, 47, 13, 34, 34, 42, 25, 15, 37}
			for i := 0; i < 88; i += 2 {
				localKey :=  /*line xYhWrwWH.go:1*/ byte(i) +  /*line xYhWrwWH.go:1*/ byte(positions[i]^positions[i+1]) + 156
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line xYhWrwWH.go:1*/ string(data)
		}() + hSHnto8d +  /*line XW9zcOlz.go:1*/ func() string {
			data :=  /*line XW9zcOlz.go:1*/ []byte("-\xd2")
			positions := [...]byte{0, 1}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line XW9zcOlz.go:1*/ byte(i) +  /*line XW9zcOlz.go:1*/ byte(positions[i]^positions[i+1]) + 79
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line XW9zcOlz.go:1*/ string(data)
		}()
		return  /*line WJD2B3XD.go:1*/ shim.Error(w5VhCLka)

	}

	cSW1wSO3 =  /*line IwQ5TOut.go:1*/ json.Unmarshal(mGZijNgp, &qo8GOXik)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line B4GFLbB_.go:1*/ func() string {
			fullData :=  /*line B4GFLbB_.go:1*/ []byte("Ř\xa3\xba\xe3K14\x06\xa8\xe3\xed\x958\xbb\xbe\xed\b\xf4\xc8C3\x1ab\xa8֯A{\"\x91\xb4,\x8fk\xd3\x02xΰ\x88\xa8\x96\xad\xd1C\u0603\xa7G|1\xb7\xa7\x8e\xdb(\x94\x8fݷ#")
			var data []byte
			data =  /*line B4GFLbB_.go:1*/ append(data, fullData[35]+fullData[41], fullData[33]^fullData[43], fullData[59]^fullData[1], fullData[60]^fullData[0], fullData[27]+fullData[51], fullData[19]+fullData[48], fullData[30]^fullData[10], fullData[34]+fullData[52], fullData[36]^fullData[13], fullData[44]-fullData[26], fullData[39]+fullData[42], fullData[22]+fullData[49], fullData[53]^fullData[38], fullData[58]-fullData[61], fullData[24]-fullData[20], fullData[23]^fullData[8], fullData[11]+fullData[21], fullData[17]^fullData[50], fullData[45]+fullData[32], fullData[40]^fullData[9], fullData[54]+fullData[25], fullData[37]+fullData[16], fullData[14]^fullData[46], fullData[31]^fullData[55], fullData[29]-fullData[15], fullData[7]+fullData[6], fullData[47]^fullData[2], fullData[12]-fullData[5], fullData[28]^fullData[56], fullData[4]-fullData[57], fullData[3]^fullData[18])
			return  /*line B4GFLbB_.go:1*/ string(data)
		}() + hSHnto8d +  /*line _APLcR0Q.go:1*/ func() string {
			fullData :=  /*line _APLcR0Q.go:1*/ []byte(" \x01\x02|")
			var data []byte
			data =  /*line _APLcR0Q.go:1*/ append(data, fullData[0]^fullData[2], fullData[3]+fullData[1])
			return  /*line _APLcR0Q.go:1*/ string(data)
		}()
		return  /*line ynVsW98_.go:1*/ shim.Error(w5VhCLka)
	}
	eviEA3sp := qo8GOXik.Dataset_Name
	hcl4L2y6 := qo8GOXik.Dataset_Provider
	olYMzSWK := qo8GOXik.Permission
	ezgSKu0t := qo8GOXik.CustomAccessRights
	cVKu8ZN2 := qo8GOXik.Name
	gN6Bx3KM := qo8GOXik.Surname
	dJwgW2jL := qo8GOXik.Organization
	igyLs9co := qo8GOXik.Username
	hzVxjOjd := qo8GOXik.Email
	hGBwBh29 := qo8GOXik.Role
	deTJzFce := qo8GOXik.Description
	e5loS483 := qo8GOXik.DatasetProviderOrg
	                                            
	qFOnXzGg, cSW1wSO3 :=  /*line RGjCszbF.go:1*/ fvIQVpVM(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line Fei0TEwk.go:1*/ shim.Error( /*line D5NzS0hs.go:1*/ cSW1wSO3.Error())
	}
	if w8EOizcx != qFOnXzGg {
		w5VhCLka =  /*line cu6GHMfA.go:1*/ func() string {
			data :=  /*line cu6GHMfA.go:1*/ []byte("ZMKrr@j\":d^\"t KuYhJ\x1d\bzA@\x1a\bf{\x1d4U ")
			positions := [...]byte{18, 19, 18, 23, 27, 26, 18, 16, 22, 14, 6, 6, 25, 26, 20, 24, 14, 2, 23, 20, 28, 1, 26, 9, 10, 15, 23, 30, 27, 10, 11, 20, 19, 16, 0, 5, 29, 0, 26, 18}
			for i := 0; i < 40; i += 2 {
				localKey :=  /*line cu6GHMfA.go:1*/ byte(i) +  /*line cu6GHMfA.go:1*/ byte(positions[i]^positions[i+1]) + 14
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line cu6GHMfA.go:1*/ string(data)
		}() + w8EOizcx +  /*line mMJEgrSI.go:1*/ func() string {
			seed :=  /*line mMJEgrSI.go:1*/ byte(190)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line mMJEgrSI.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/  /*line mMJEgrSI.go:1*/ fnc(216)(249)(253)(172)(92)(245)(253)(231)(30)(238)(13)(188)(98)
			return  /*line mMJEgrSI.go:1*/ string(data)
		}() + eviEA3sp +  /*line uC9L26wv.go:1*/ func() string {
			key :=  /*line uC9L26wv.go:1*/ []byte("I\x1f")
			data :=  /*line uC9L26wv.go:1*/ []byte("kb")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line uC9L26wv.go:1*/ string(data)
		}()
		return  /*line BGMD8LJx.go:1*/ shim.Error(w5VhCLka)
	}
	syHhHGpa := eviEA3sp +  /*line eZ8jjbJ4.go:1*/ func() string {
		fullData :=  /*line eZ8jjbJ4.go:1*/ []byte("}\xb2")
		var data []byte
		data =  /*line eZ8jjbJ4.go:1*/ append(data, fullData[1]+fullData[0])
		return  /*line eZ8jjbJ4.go:1*/ string(data)
	}() + igyLs9co

	b4dMB1zL, cSW1wSO3 :=  /*line oT3BGerd.go:1*/ jmSY84co(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line URTyNozG.go:1*/ shim.Error( /*line qwEydife.go:1*/ cSW1wSO3.Error())
	}

	zzhdvyqZ := &DatasetMetadataResponse{}
	cSW1wSO3 =  /*line zgMKMF3S.go:1*/ json.Unmarshal( /*line zfGMSNba.go:1*/ []byte(b4dMB1zL), zzhdvyqZ)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line oGARuXQH.go:1*/ func() string {
			var data []byte
			i := 0
			decryptKey := 250
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 7:
					data =  /*line oGARuXQH.go:1*/ append(data, "«\xd0\xec"...,
					)
					i = 3
				case 5:
					i = 1
					data =  /*line oGARuXQH.go:1*/ append(data, "\xe5\xef\xec\xec"...,
					)
				case 10:
					data =  /*line oGARuXQH.go:1*/ append(data, 176)
					i = 12
				case 1:
					i = 8
					for y := range data {
						data[y] = data[y] -  /*line oGARuXQH.go:1*/ byte(decryptKey^y)
					}
				case 11:
					data =  /*line oGARuXQH.go:1*/ append(data, "\xc7\xf5"...,
					)
					i = 13
				case 6:
					i = 4
					data =  /*line oGARuXQH.go:1*/ append(data, 248)
				case 0:
					i = 11
					data =  /*line oGARuXQH.go:1*/ append(data, "\xfb\xa3"...,
					)
				case 3:
					data =  /*line oGARuXQH.go:1*/ append(data, "\xf5\xf9\xf3\xf3"...,
					)
					i = 10
				case 12:
					data =  /*line oGARuXQH.go:1*/ append(data, "\x05\x01\xb3\xf8"...,
					)
					i = 9
				case 2:
					data =  /*line oGARuXQH.go:1*/ append(data, "\xfc\xfe\xba"...,
					)
					i = 5
				case 4:
					data =  /*line oGARuXQH.go:1*/ append(data, 169)
					i = 7
				case 9:
					i = 2
					data =  /*line oGARuXQH.go:1*/ append(data, "\xfa\xf9\x06"...,
					)
				case 13:
					data =  /*line oGARuXQH.go:1*/ append(data, "\xf6\xf4"...,
					)
					i = 6
				}
			}
			return  /*line oGARuXQH.go:1*/ string(data)
		}() + eviEA3sp +  /*line LzrJjhEh.go:1*/ func() string {
			seed :=  /*line LzrJjhEh.go:1*/ byte(73)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line LzrJjhEh.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line LzrJjhEh.go:1*/  /*line LzrJjhEh.go:1*/ fnc(217)(91)
			return  /*line LzrJjhEh.go:1*/ string(data)
		}()
		return  /*line ozXiWJaV.go:1*/ shim.Error(w5VhCLka)
	}

	                                                                  
	if  /*line Ahy1976I.go:1*/ nSRaDwQg(ezgSKu0t, zzhdvyqZ.CustomAccessRights) {
		 /*line Dluxt8th.go:1*/ fmt.Println( /*line emot6MwP.go:1*/ func() string {
			var data []byte
			i := 10
			decryptKey := 104
			for counter := 0; i != 4; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 10:
					i = 8
					data =  /*line emot6MwP.go:1*/ append(data, 116)
				case 3:
					data =  /*line emot6MwP.go:1*/ append(data, "\x81\x8a\x89"...,
					)
					i = 12
				case 2:
					i = 15
					data =  /*line emot6MwP.go:1*/ append(data, "\x91\x96\x96"...,
					)
				case 0:
					i = 11
					data =  /*line emot6MwP.go:1*/ append(data, 166)
				case 15:
					i = 6
					data =  /*line emot6MwP.go:1*/ append(data, 161)
				case 13:
					i = 2
					data =  /*line emot6MwP.go:1*/ append(data, "\x9eJ\x9b"...,
					)
				case 6:
					i = 9
					data =  /*line emot6MwP.go:1*/ append(data, 159)
				case 12:
					i = 13
					data =  /*line emot6MwP.go:1*/ append(data, "\x8a\x97"...,
					)
				case 8:
					data =  /*line emot6MwP.go:1*/ append(data, "\xa2\xa1\xa2"...,
					)
					i = 0
				case 7:
					i = 3
					data =  /*line emot6MwP.go:1*/ append(data, "\x8fA"...,
					)
				case 9:
					i = 4
					for y := range data {
						data[y] = data[y] -  /*line emot6MwP.go:1*/ byte(decryptKey^y)
					}
				case 14:
					data =  /*line emot6MwP.go:1*/ append(data, "\x9d\x9c\xae\x9e"...,
					)
					i = 5
				case 11:
					data =  /*line emot6MwP.go:1*/ append(data, "\xa6\xa7"...,
					)
					i = 14
				case 5:
					i = 1
					data =  /*line emot6MwP.go:1*/ append(data, "X\xa2\xb3\xb0"...,
					)
				case 1:
					data =  /*line emot6MwP.go:1*/ append(data, "\xb0\x92"...,
					)
					i = 7
				}
			}
			return  /*line emot6MwP.go:1*/ string(data)
		}())
	} else {
		w5VhCLka =  /*line vGBb3uZW.go:1*/ func() string {
			data :=  /*line vGBb3uZW.go:1*/ []byte("{qEr\xefo\xac\x96\xeefW\xefe\x83g\x98\xe2줚\xf3m\xcc,\xa7Q\x8d\x9c\x00 \x83i\xe0hWs\xdd}")
			positions := [...]byte{22, 26, 25, 28, 34, 16, 1, 24, 26, 20, 12, 30, 12, 28, 36, 12, 32, 6, 9, 19, 23, 36, 27, 8, 11, 4, 13, 36, 22, 20, 16, 16, 18, 8, 32, 18, 24, 30, 7, 15, 17, 18, 19, 19}
			for i := 0; i < 44; i += 2 {
				localKey :=  /*line vGBb3uZW.go:1*/ byte(i) +  /*line vGBb3uZW.go:1*/ byte(positions[i]^positions[i+1]) + 92
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line vGBb3uZW.go:1*/ string(data)
		}()
		return  /*line fdVADA8r.go:1*/ shim.Error(w5VhCLka)
	}

	var iECyqXpw []string
	if olYMzSWK ==  /*line xg4OPfjY.go:1*/ func() string {
		fullData :=  /*line xg4OPfjY.go:1*/ []byte("\xa0M\xef\x12\x18\x8b\xf7X")
		var data []byte
		data =  /*line xg4OPfjY.go:1*/ append(data, fullData[3]-fullData[0], fullData[4]+fullData[1], fullData[7]-fullData[6], fullData[5]^fullData[2])
		return  /*line xg4OPfjY.go:1*/ string(data)
	}() {
		iECyqXpw = []string{ /*line sH6qnHjS.go:1*/ func() string {
			var data []byte
			i := 0
			decryptKey := 64
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 5
					for y := range data {
						data[y] = data[y] -  /*line sH6qnHjS.go:1*/ byte(decryptKey^y)
					}
				case 0:
					i = 2
					data =  /*line sH6qnHjS.go:1*/ append(data, 183)
				case 4:
					i = 1
					data =  /*line sH6qnHjS.go:1*/ append(data, 168)
				case 1:
					data =  /*line sH6qnHjS.go:1*/ append(data, 170)
					i = 3
				case 2:
					i = 4
					data =  /*line sH6qnHjS.go:1*/ append(data, 169)
				}
			}
			return  /*line sH6qnHjS.go:1*/ string(data)
		}()}
	} else if olYMzSWK ==  /*line Znha8pYk.go:1*/ func() string {
		fullData :=  /*line Znha8pYk.go:1*/ []byte("\xc8=\xe2[\xd4$\xdb\xe6u\x97K\x82")
		var data []byte
		data =  /*line Znha8pYk.go:1*/ append(data, fullData[0]-fullData[3], fullData[5]^fullData[10], fullData[7]^fullData[11], fullData[1]-fullData[4], fullData[6]-fullData[8], fullData[2]+fullData[9])
		return  /*line Znha8pYk.go:1*/ string(data)
	}() {
		iECyqXpw = []string{ /*line WOhsmfsp.go:1*/ func() string {
			fullData :=  /*line WOhsmfsp.go:1*/ []byte("1z\xeb+C\xebLO")
			var data []byte
			data =  /*line WOhsmfsp.go:1*/ append(data, fullData[4]^fullData[0], fullData[1]+fullData[2], fullData[6]-fullData[5], fullData[7]^fullData[3])
			return  /*line WOhsmfsp.go:1*/ string(data)
		}(),  /*line Md0dHMo2.go:1*/ func() string {
			seed :=  /*line Md0dHMo2.go:1*/ byte(230)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line Md0dHMo2.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line Md0dHMo2.go:1*/  /*line Md0dHMo2.go:1*/  /*line Md0dHMo2.go:1*/  /*line Md0dHMo2.go:1*/  /*line Md0dHMo2.go:1*/  /*line Md0dHMo2.go:1*/ fnc(135)(2)(245)(5)(253)(19)
			return  /*line Md0dHMo2.go:1*/ string(data)
		}()}
	} else if olYMzSWK ==  /*line z0xF9CZs.go:1*/ func() string {
		key :=  /*line z0xF9CZs.go:1*/ []byte("m\xdes\xfbg]\x93")
		data :=  /*line z0xF9CZs.go:1*/ []byte("\x03\x87\xffx\x02\x16\xe1")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line z0xF9CZs.go:1*/ string(data)
	}() {
		iECyqXpw = []string{ /*line CrOHZeEr.go:1*/ func() string {
			var data []byte
			i := 5
			decryptKey := 44
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 4
					data =  /*line CrOHZeEr.go:1*/ append(data, 131)
				case 1:
					for y := range data {
						data[y] = data[y] -  /*line CrOHZeEr.go:1*/ byte(decryptKey^y)
					}
					i = 0
				case 5:
					data =  /*line CrOHZeEr.go:1*/ append(data, 146)
					i = 2
				case 2:
					data =  /*line CrOHZeEr.go:1*/ append(data, 134)
					i = 3
				case 4:
					i = 1
					data =  /*line CrOHZeEr.go:1*/ append(data, 135)
				}
			}
			return  /*line CrOHZeEr.go:1*/ string(data)
		}(),  /*line XBV1oaT4.go:1*/ func() string {
			key :=  /*line XBV1oaT4.go:1*/ []byte("\x98\xd1ì\xf6\x14")
			data :=  /*line XBV1oaT4.go:1*/ []byte("\xf5\xbe\xa7Őm")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line XBV1oaT4.go:1*/ string(data)
		}(),  /*line eS5ME9SJ.go:1*/ func() string {
			fullData :=  /*line eS5ME9SJ.go:1*/ []byte("\rK\xab\xb3g\u0092\xe0!\xc0\xf48Q\x8f")
			var data []byte
			data =  /*line eS5ME9SJ.go:1*/ append(data, fullData[12]^fullData[8], fullData[10]-fullData[13], fullData[6]^fullData[7], fullData[3]^fullData[9], fullData[5]^fullData[2], fullData[11]^fullData[1], fullData[4]+fullData[0])
			return  /*line eS5ME9SJ.go:1*/ string(data)
		}()}
	} else {
		 /*line YBGcDL4Q.go:1*/ fmt.Println( /*line X7DOsTHi.go:1*/ func() string {
			seed :=  /*line X7DOsTHi.go:1*/ byte(49)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line X7DOsTHi.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/  /*line X7DOsTHi.go:1*/ fnc(29)(33)(5)(172)(84)(244)(253)(187)(85)(254)(2)(236)(11)(180)(80)(245)(13)(251)(252)(10)(0)(246)(6)(255)
			return  /*line X7DOsTHi.go:1*/ string(data)
		}())
	}

	tQcQN8RB, cSW1wSO3 :=  /*line iFazKwoo.go:1*/ n4c7mRtG.GetState(syHhHGpa)	                    
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line ZjHyzrWK.go:1*/ func() string {
			seed :=  /*line ZjHyzrWK.go:1*/ byte(164)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line ZjHyzrWK.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/  /*line ZjHyzrWK.go:1*/ fnc(31)(229)(237)(7)(14)(25)(53)(26)(76)(128)(36)(99)(206)(159)(55)(109)(150)(128)(251)(167)(145)(39)(75)(148)(48)(21)(107)(234)(211)(154)(59)(121)(233)(227)(173)(109)(207)(164)(71)(64)(198)(149)(45)(8)(100)(188)(117)(165)(142)(25)(69)(119)(0)(242)(243)(172)(62)
			return  /*line ZjHyzrWK.go:1*/ string(data)
		}() + syHhHGpa +  /*line w7S8YyuZ.go:1*/ func() string {
			data :=  /*line w7S8YyuZ.go:1*/ []byte("\x93}")
			positions := [...]byte{0, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line w7S8YyuZ.go:1*/ byte(i) +  /*line w7S8YyuZ.go:1*/ byte(positions[i]^positions[i+1]) + 143
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line w7S8YyuZ.go:1*/ string(data)
		}()
		return  /*line BGpInoHK.go:1*/ shim.Error(w5VhCLka)
	} else if tQcQN8RB == nil && (olYMzSWK == "" || olYMzSWK ==  /*line JY2tOjZw.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 255
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				for y := range data {
					data[y] = data[y] -  /*line JY2tOjZw.go:1*/ byte(decryptKey^y)
				}
				i = 0
			case 2:
				i = 1
				data =  /*line JY2tOjZw.go:1*/ append(data, 30)
			}
		}
		return  /*line JY2tOjZw.go:1*/ string(data)
	}()) {
		w5VhCLka =  /*line T8qEv6XW.go:1*/ func() string {
			fullData :=  /*line T8qEv6XW.go:1*/ []byte("\xe8\xa3\x11u\x1e\xb1\x92\x9d\xde\xe3\xc1\xff\xa8\xdc.\x82m\x04\x8f\xe5\xe8\x1bU\x94\xd3\x03\xaa\x80\x9bqb\x8cX\u0379\x13Z\xbb\xda*\xefa\xe4wuR\xe5\x16\xf3I\x9d\xdaU\x02a\xd1\xcb\x16b\xfa\xa5z\xb6\xb5'\xd0Bxj\xaa\x06\xf43\n\x13\x8d&3_ֺ\x17ᖵ\xbc\xe9\x8f\xe2\xfa\x8a\xd3}\xd64e\xa9r\u007f͋\xdc\x1e\x8f\xab\x10\x8d\x15\x8bvD|\xd7\xed\xdb\x064Q\x93_XT\xc5\x16\xcc&\x19\xfd>\xc4#\xae\xe2\xeao\xfa\xf3 \xb9\nSu\xbc\xb7\xe5\xb2\xf0[\xc3\f\x9ev\xf6\xac")
			var data []byte
			data =  /*line T8qEv6XW.go:1*/ append(data, fullData[24]+fullData[12], fullData[149]-fullData[133], fullData[116]-fullData[40], fullData[13]+fullData[83], fullData[82]^fullData[118], fullData[131]+fullData[10], fullData[34]^fullData[56], fullData[5]-fullData[87], fullData[6]-fullData[32], fullData[53]+fullData[137], fullData[79]-fullData[15], fullData[0]-fullData[27], fullData[63]^fullData[65], fullData[141]^fullData[22], fullData[64]-fullData[143], fullData[134]-fullData[139], fullData[62]+fullData[142], fullData[26]+fullData[148], fullData[106]^fullData[42], fullData[109]-fullData[25], fullData[120]-fullData[19], fullData[21]^fullData[97], fullData[84]^fullData[38], fullData[144]^fullData[100], fullData[76]+fullData[59], fullData[11]+fullData[68], fullData[95]^fullData[123], fullData[58]-fullData[66], fullData[147]+fullData[35], fullData[28]^fullData[71], fullData[31]+fullData[20], fullData[61]^fullData[36], fullData[55]-fullData[16], fullData[153]+fullData[138], fullData[72]^fullData[52], fullData[140]+fullData[47], fullData[23]^fullData[135], fullData[129]-fullData[119], fullData[69]+fullData[80], fullData[105]+fullData[102], fullData[48]-fullData[91], fullData[96]+fullData[7], fullData[98]+fullData[132], fullData[8]-fullData[3], fullData[103]^fullData[9], fullData[101]-fullData[43], fullData[30]^fullData[115], fullData[89]^fullData[51], fullData[29]-fullData[127], fullData[94]-fullData[122], fullData[128]+fullData[88], fullData[104]^fullData[33], fullData[111]+fullData[113], fullData[39]^fullData[110], fullData[44]^fullData[2], fullData[50]-fullData[92], fullData[136]+fullData[151], fullData[67]^fullData[57], fullData[114]^fullData[145], fullData[93]+fullData[150], fullData[112]+fullData[49], fullData[1]+fullData[99], fullData[45]+fullData[130], fullData[86]^fullData[108], fullData[117]-fullData[46], fullData[78]+fullData[73], fullData[4]-fullData[37], fullData[70]^fullData[125], fullData[121]-fullData[146], fullData[85]+fullData[60], fullData[75]-fullData[126], fullData[18]-fullData[14], fullData[90]-fullData[81], fullData[41]+fullData[17], fullData[107]^fullData[54], fullData[124]^fullData[152], fullData[77]^fullData[74])
			return  /*line T8qEv6XW.go:1*/ string(data)
		}() + eviEA3sp +  /*line TdBo53tn.go:1*/ func() string {
			data :=  /*line TdBo53tn.go:1*/ []byte("\xcc}")
			positions := [...]byte{0, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line TdBo53tn.go:1*/ byte(i) +  /*line TdBo53tn.go:1*/ byte(positions[i]^positions[i+1]) + 86
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line TdBo53tn.go:1*/ string(data)
		}()
		return  /*line L_fthOFV.go:1*/ shim.Error(w5VhCLka)
	} else if tQcQN8RB != nil {
		a8ZPVup7 := AuthorizedUsers{}
		cSW1wSO3 =  /*line BKtxZRe2.go:1*/ json.Unmarshal(tQcQN8RB, &a8ZPVup7)	                               
		if cSW1wSO3 != nil {
			return  /*line PgezDjNB.go:1*/ shim.Error( /*line fV2FGTgt.go:1*/ cSW1wSO3.Error())
		}

		if olYMzSWK == "" || olYMzSWK ==  /*line fc0En9kh.go:1*/ func() string {
			var data []byte
			i := 2
			decryptKey := 32
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					i = 1
					data =  /*line fc0En9kh.go:1*/ append(data, 65)
				case 1:
					i = 0
					for y := range data {
						data[y] = data[y] -  /*line fc0En9kh.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return  /*line fc0En9kh.go:1*/ string(data)
		}() {
			iECyqXpw = a8ZPVup7.Users.Permission
		}

		 /*line Q7FaFrLy.go:1*/ fmt.Println( /*line o911_niD.go:1*/ func() string {
			key :=  /*line o911_niD.go:1*/ []byte("\xba\x86\xb5Q\x8c\x0e\xf5&\f\xc59\xb2\xe7Y\xe2\xab U\x11,\xd8\xfcO\x1f\xaabE")
			data :=  /*line o911_niD.go:1*/ []byte("\xfb\xf6%\xb6\xfar\x15\x89\x818\xad!TyC\x0e\x83\xba\x84\x9f\xf8n\xb8\x86\x12ָ")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line o911_niD.go:1*/ string(data)
		}())
		if  /*line BQ8qiiba.go:1*/ len(ezgSKu0t) > 0 {
			ezgSKu0t =  /*line tsepeUGg.go:1*/ _urwHavD(ezgSKu0t, a8ZPVup7.Users.CustomAccessRights)
		}

	}
	                                                                                                                                                                                                                                                                                                                                                                                                                                                              

	im2M6_4m := &CredentialsAuthorizedUsers{iECyqXpw, ezgSKu0t, cVKu8ZN2, gN6Bx3KM, dJwgW2jL, igyLs9co, hzVxjOjd, hGBwBh29}
	                                                                  
	l5TiVBv5 := &AuthorizedUsers{mzjt9rLK, eviEA3sp, hcl4L2y6, e5loS483, deTJzFce, im2M6_4m}
	niQzOH6M, cSW1wSO3 :=  /*line f6xwb4yQ.go:1*/ json.Marshal(l5TiVBv5)
	if cSW1wSO3 != nil {
		return  /*line FmsgrdXW.go:1*/ shim.Error( /*line rrgnakfC.go:1*/ cSW1wSO3.Error())
	}
	                        
	cSW1wSO3 =  /*line ORDgSYF1.go:1*/ n4c7mRtG.PutState(syHhHGpa, niQzOH6M)
	if cSW1wSO3 != nil {
		return  /*line Ew1Dyayd.go:1*/ shim.Error( /*line daqKf_Np.go:1*/ cSW1wSO3.Error())
	}

	cSW1wSO3 =  /*line zu_WRLqz.go:1*/ n4c7mRtG.DelState(hSHnto8d)	                                                                
	if cSW1wSO3 != nil {
		return  /*line eKMCQInF.go:1*/ shim.Error( /*line KgpDxn3k.go:1*/ func() string {
			key :=  /*line KgpDxn3k.go:1*/ []byte("\xa1\x96\x9bs\xdb#\x0f\xedC\xb5\xb6\xe8\x15\xcc\x02=\x8e\xcdT\xbb\xf9\x16\x1e\x1do\xd5,\xf0\x8e\xf3\x19\xdd")
			data :=  /*line KgpDxn3k.go:1*/ []byte("\xe7\xf7\x04\xdf@\x87/a\xb2\xd5+Xy-v\xa2\xae?\xb9,n{\x91\x91\x8f6\x8fS\xf3f\x8c\x17")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line KgpDxn3k.go:1*/ string(data)
		}() +  /*line Mcpf4qDx.go:1*/ cSW1wSO3.Error())
	}

	               
	xwK9eDsS, cSW1wSO3 :=  /*line Fh3UjmWl.go:1*/ wjnibpdu(n4c7mRtG, []string{eviEA3sp, w8EOizcx, igyLs9co, dJwgW2jL, hSHnto8d,  /*line AuQLq1jg.go:1*/ func() string {
		fullData :=  /*line AuQLq1jg.go:1*/ []byte("\n\xed\xc02ZTi\xb0̙\xce1")
		var data []byte
		data =  /*line AuQLq1jg.go:1*/ append(data, fullData[1]+fullData[5], fullData[11]+fullData[3], fullData[6]^fullData[0], fullData[9]+fullData[8], fullData[2]+fullData[7], fullData[10]-fullData[4])
		return  /*line AuQLq1jg.go:1*/ string(data)
	}(), tIeFhn0D, olYMzSWK,  /*line g0Nc39fn.go:1*/ strings.Join(ezgSKu0t,  /*line DPTmS09I.go:1*/ func() string {
		fullData :=  /*line DPTmS09I.go:1*/ []byte("(\xfc")
		var data []byte
		data =  /*line DPTmS09I.go:1*/ append(data, fullData[0]-fullData[1])
		return  /*line DPTmS09I.go:1*/ string(data)
	}()), e5loS483, deTJzFce})
	if cSW1wSO3 != nil {
		return  /*line ZP_3vS6t.go:1*/ shim.Error( /*line gEywpYpg.go:1*/ cSW1wSO3.Error())
	}
	 /*line kVgnmjf_.go:1*/ fmt.Println(xwK9eDsS)

	xHTOVzVs :=  /*line Qd6Hdz2x.go:1*/ func() string {
		seed :=  /*line Qd6Hdz2x.go:1*/ byte(51)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Qd6Hdz2x.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line Qd6Hdz2x.go:1*/  /*line Qd6Hdz2x.go:1*/  /*line Qd6Hdz2x.go:1*/  /*line Qd6Hdz2x.go:1*/  /*line Qd6Hdz2x.go:1*/  /*line Qd6Hdz2x.go:1*/  /*line Qd6Hdz2x.go:1*/  /*line Qd6Hdz2x.go:1*/  /*line Qd6Hdz2x.go:1*/ fnc(103)(242)(233)(85)(191)(250)(230)(27)(164)
		return  /*line Qd6Hdz2x.go:1*/ string(data)
	}() + cVKu8ZN2 +  /*line C_3OTfOQ.go:1*/ func() string {
		seed :=  /*line C_3OTfOQ.go:1*/ byte(142)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line C_3OTfOQ.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line C_3OTfOQ.go:1*/ fnc(174)
		return  /*line C_3OTfOQ.go:1*/ string(data)
	}() + gN6Bx3KM +  /*line I48RfL7Q.go:1*/ func() string {
		data :=  /*line I48RfL7Q.go:1*/ []byte("\xbf\ay\r \x13e\xc1\xe3\x96\x0e\v\x14!\x11r\xd0\r\xc5%\xf3")
		positions := [...]byte{18, 0, 16, 5, 11, 14, 2, 18, 1, 10, 2, 7, 13, 19, 9, 7, 20, 7, 18, 17, 20, 8, 12, 5, 2, 3, 17, 7}
		for i := 0; i < 28; i += 2 {
			localKey :=  /*line I48RfL7Q.go:1*/ byte(i) +  /*line I48RfL7Q.go:1*/ byte(positions[i]^positions[i+1]) + 147
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line I48RfL7Q.go:1*/ string(data)
	}() +  /*line D_FfbFzE.go:1*/ func() string {
		fullData :=  /*line D_FfbFzE.go:1*/ []byte("\xff\xa9\xf5@\x1e!!\x90\u007f\xca+l\xb7:\xcd>\x06*V\xc6r\xc4")
		var data []byte
		data =  /*line D_FfbFzE.go:1*/ append(data, fullData[16]^fullData[20], fullData[19]^fullData[1], fullData[15]^fullData[4], fullData[17]+fullData[13], fullData[14]-fullData[11], fullData[9]-fullData[18], fullData[3]+fullData[5], fullData[21]^fullData[12], fullData[7]-fullData[10], fullData[2]+fullData[8], fullData[6]+fullData[0])
		return  /*line D_FfbFzE.go:1*/ string(data)
	}() + eviEA3sp
	q8ZF8zaz :=  /*line OWPu4XYZ.go:1*/ []byte(xHTOVzVs)

	return  /*line GmwBoAfa.go:1*/ shim.Success(q8ZF8zaz)
}

func (g4rnrSNn *G1Y_7pz_) if1PvepX(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	var cSW1wSO3 error
	if  /*line hRw8rxZE.go:1*/ len(ktsi1_SQ) != 5 {
		return  /*line vORZBzm8.go:1*/ shim.Error( /*line IVZ0iRS0.go:1*/ func() string {
			fullData :=  /*line IVZ0iRS0.go:1*/ []byte("a\xcd\xedo\xb4\x8a\xc4\xce\bf\x143\xfe\xe9yQc\xa7V\f\xed\xb7\x1f\xd8\xfa\xbf\xee\t\xde\xf1\xb1\xe9Ii\xea>\x1eLBUk\x83\x11mʬ\xc5\xed\xb5\xd0\xfdL^\xd3\x1a`\x1d\xf4\x9b\x01G\x82\xb6Ұ\x8b\xceb\xd5\xdd\xfb\xbci\xf9\x94\x84x\x8bk\xd7k\x9akN\x1f\xc0ht\xb8\xa60\xc9\r\xf8\xba\xbf5O\xf3F4$~\x01\x90z\x1bbRŅ\xcct\v39\x85(\x95\x97VX\xc9\x06\xaad\xc9ૂ\xcd-z\xfa;\n")
			var data []byte
			data =  /*line IVZ0iRS0.go:1*/ append(data, fullData[53]^fullData[81], fullData[126]^fullData[17], fullData[103]+fullData[67], fullData[129]+fullData[47], fullData[23]^fullData[124], fullData[118]+fullData[69], fullData[4]+fullData[30], fullData[70]+fullData[86], fullData[111]-fullData[121], fullData[131]^fullData[92], fullData[16]+fullData[113], fullData[43]+fullData[8], fullData[40]^fullData[123], fullData[60]+fullData[106], fullData[125]^fullData[59], fullData[51]^fullData[35], fullData[127]^fullData[85], fullData[122]+fullData[89], fullData[6]-fullData[52], fullData[100]^fullData[10], fullData[87]+fullData[2], fullData[38]+fullData[90], fullData[112]+fullData[98], fullData[15]+fullData[101], fullData[3]+fullData[12], fullData[74]^fullData[29], fullData[84]+fullData[97], fullData[19]^fullData[76], fullData[26]+fullData[116], fullData[128]^fullData[77], fullData[11]-fullData[1], fullData[95]^fullData[49], fullData[130]^fullData[25], fullData[93]+fullData[117], fullData[110]-fullData[22], fullData[78]^fullData[36], fullData[94]-fullData[37], fullData[66]-fullData[80], fullData[115]+fullData[134], fullData[104]^fullData[73], fullData[9]+fullData[27], fullData[88]+fullData[62], fullData[33]^fullData[32], fullData[58]^fullData[31], fullData[45]^fullData[91], fullData[57]+fullData[61], fullData[102]^fullData[42], fullData[135]+fullData[0], fullData[68]^fullData[64], fullData[109]^fullData[75], fullData[48]-fullData[108], fullData[24]-fullData[119], fullData[107]-fullData[50], fullData[13]+fullData[5], fullData[72]^fullData[54], fullData[55]^fullData[83], fullData[34]-fullData[44], fullData[132]-fullData[96], fullData[7]-fullData[18], fullData[46]-fullData[39], fullData[120]^fullData[114], fullData[56]+fullData[99], fullData[133]+fullData[105], fullData[65]+fullData[28], fullData[41]^fullData[20], fullData[63]-fullData[82], fullData[79]-fullData[21], fullData[14]+fullData[71])
			return  /*line IVZ0iRS0.go:1*/ string(data)
		}())
	}

	eviEA3sp := ktsi1_SQ[0]
	olYMzSWK := ktsi1_SQ[1]
	igyLs9co := ktsi1_SQ[2]
	tIeFhn0D := ktsi1_SQ[3]
	ezgSKu0t :=  /*line TVLKRxPl.go:1*/ strings.Split(ktsi1_SQ[4],  /*line kwHRe8zy.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 86
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 0
				for y := range data {
					data[y] = data[y] +  /*line kwHRe8zy.go:1*/ byte(decryptKey^y)
				}
			case 2:
				i = 1
				data =  /*line kwHRe8zy.go:1*/ append(data, 213)
			}
		}
		return  /*line kwHRe8zy.go:1*/ string(data)
	}())

	w8EOizcx, cSW1wSO3 :=  /*line HdjvAiUz.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line XeacsCar.go:1*/ shim.Error( /*line B8nS3LbR.go:1*/ cSW1wSO3.Error())
	}

	b4dMB1zL, cSW1wSO3 :=  /*line gKbrxfq3.go:1*/ jmSY84co(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line jyJyN6B5.go:1*/ shim.Error( /*line EJXZO35N.go:1*/ cSW1wSO3.Error())
	}

	zzhdvyqZ := &DatasetMetadataResponse{}
	cSW1wSO3 =  /*line YxNAT26Z.go:1*/ json.Unmarshal( /*line EjbWOLi7.go:1*/ []byte(b4dMB1zL), zzhdvyqZ)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line z6iMQeyN.go:1*/ func() string {
			data :=  /*line z6iMQeyN.go:1*/ []byte("\x94G\xf7\xee\xfc\xc5r\xd5:\xb2\xf4ai\x02ed\xb4th\xa1\x1d\xa2Ъd\xefw\xb1SO\x0e")
			positions := [...]byte{21, 25, 22, 19, 18, 0, 20, 23, 27, 25, 10, 22, 9, 16, 21, 2, 1, 5, 10, 30, 26, 5, 20, 3, 25, 25, 30, 10, 18, 5, 4, 13, 30, 0, 2, 7}
			for i := 0; i < 36; i += 2 {
				localKey :=  /*line z6iMQeyN.go:1*/ byte(i) +  /*line z6iMQeyN.go:1*/ byte(positions[i]^positions[i+1]) + 73
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line z6iMQeyN.go:1*/ string(data)
		}() + eviEA3sp +  /*line CauAZ2kF.go:1*/ func() string {
			seed :=  /*line CauAZ2kF.go:1*/ byte(196)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line CauAZ2kF.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line CauAZ2kF.go:1*/  /*line CauAZ2kF.go:1*/ fnc(230)(215)
			return  /*line CauAZ2kF.go:1*/ string(data)
		}()
		return  /*line iUTUxEzF.go:1*/ shim.Error(w5VhCLka)
	}

	                                                                  
	if  /*line XHi3CMnI.go:1*/ nSRaDwQg(ezgSKu0t, zzhdvyqZ.CustomAccessRights) {
		 /*line rU8ASySA.go:1*/ fmt.Println( /*line VAtucd5X.go:1*/ func() string {
			seed :=  /*line VAtucd5X.go:1*/ byte(23)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line VAtucd5X.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/  /*line VAtucd5X.go:1*/ fnc(42)(47)(0)(2)(253)(1)(2)(247)(248)(19)(241)(187)(67)(18)(254)(1)(251)(254)(179)(65)(2)(0)(2)(14)(0)(173)(82)(247)(254)(1)(12)(255)
			return  /*line VAtucd5X.go:1*/ string(data)
		}())
	} else {
		w5VhCLka =  /*line HpIiqGGu.go:1*/ func() string {
			fullData :=  /*line HpIiqGGu.go:1*/ []byte("\a\xaf\xad\xf2~>C\xd4!)\xd0\x18\x0f\xf6\x9c>R\x87\tY\xb1\xfe\f\xbc/\x8dV\xc8\xe1\f\xd3\b\xa2]>Ҝ\xf9hY\xc1\x18Y\xe7\xd6\x19\x11\xfe\x9aE\xda\xda\x02\xae=j\xfcDH\x99d\x8d\x15p\xa5\xc6O\x0fJ\a\x1e+_/\x11i")
			var data []byte
			data =  /*line HpIiqGGu.go:1*/ append(data, fullData[49]^fullData[5], fullData[6]-fullData[8], fullData[33]^fullData[41], fullData[32]+fullData[10], fullData[37]-fullData[17], fullData[30]+fullData[36], fullData[1]-fullData[54], fullData[56]-fullData[50], fullData[35]+fullData[38], fullData[9]-fullData[0], fullData[66]+fullData[31], fullData[22]^fullData[4], fullData[71]+fullData[57], fullData[63]-fullData[52], fullData[39]-fullData[3], fullData[34]^fullData[70], fullData[60]^fullData[69], fullData[25]-fullData[11], fullData[2]+fullData[65], fullData[51]+fullData[48], fullData[44]+fullData[59], fullData[27]+fullData[64], fullData[74]+fullData[12], fullData[55]-fullData[18], fullData[68]-fullData[43], fullData[46]-fullData[53], fullData[29]+fullData[42], fullData[67]-fullData[14], fullData[24]-fullData[23], fullData[40]^fullData[28], fullData[19]+fullData[45], fullData[20]-fullData[58], fullData[47]+fullData[75], fullData[15]^fullData[26], fullData[62]+fullData[72], fullData[21]^fullData[61], fullData[13]^fullData[7], fullData[16]^fullData[73])
			return  /*line HpIiqGGu.go:1*/ string(data)
		}()
		return  /*line QALzLHWg.go:1*/ shim.Error(w5VhCLka)
	}

	                                            
	qFOnXzGg, cSW1wSO3 :=  /*line Gh0Kqkci.go:1*/ fvIQVpVM(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line tp7vwdS2.go:1*/ shim.Error( /*line aUSzmiaP.go:1*/ cSW1wSO3.Error())
	}
	if qFOnXzGg != w8EOizcx {
		return  /*line Jrryji62.go:1*/ shim.Error( /*line FlfD0sSt.go:1*/ func() string {
			data :=  /*line FlfD0sSt.go:1*/ []byte("\xf9o\xe6\xc1}\xc3\xca\x06\xffr\xaez\x1d\xb2&to\x1dr\"\x13\xb0\xe8e zcc\xfds\xc0")
			positions := [...]byte{4, 21, 19, 30, 12, 30, 6, 14, 22, 3, 12, 3, 28, 4, 17, 21, 12, 10, 5, 14, 13, 13, 20, 22, 19, 10, 8, 7, 10, 13, 13, 0, 25, 2, 12, 2}
			for i := 0; i < 36; i += 2 {
				localKey :=  /*line FlfD0sSt.go:1*/ byte(i) +  /*line FlfD0sSt.go:1*/ byte(positions[i]^positions[i+1]) + 64
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line FlfD0sSt.go:1*/ string(data)
		}())
	}
	rHDGEtgU := eviEA3sp +  /*line EihNDRQd.go:1*/ func() string {
		fullData :=  /*line EihNDRQd.go:1*/ []byte("\xf78")
		var data []byte
		data =  /*line EihNDRQd.go:1*/ append(data, fullData[1]+fullData[0])
		return  /*line EihNDRQd.go:1*/ string(data)
	}() + igyLs9co
	tQcQN8RB, cSW1wSO3 :=  /*line zRPMPEbq.go:1*/ n4c7mRtG.GetState(rHDGEtgU)	                    
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line _rejrzAW.go:1*/ func() string {
			seed :=  /*line _rejrzAW.go:1*/ byte(202)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line _rejrzAW.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/  /*line _rejrzAW.go:1*/ fnc(177)(89)(145)(23)(14)(229)(29)(174)(0)(24)(20)(7)(4)(29)(235)(29)(182)(56)(235)(79)(221)(243)(235)(26)(248)(171)(87)(248)(241)(30)(251)(253)(229)(11)(29)(237)(239)(26)(225)(80)(166)(9)(29)(172)(76)(236)(21)(165)(78)(25)(229)(23)(254)(238)(13)(188)(98)
			return  /*line _rejrzAW.go:1*/ string(data)
		}() + rHDGEtgU +  /*line PL_j6Y03.go:1*/ func() string {
			data :=  /*line PL_j6Y03.go:1*/ []byte("\xadR")
			positions := [...]byte{1, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line PL_j6Y03.go:1*/ byte(i) +  /*line PL_j6Y03.go:1*/ byte(positions[i]^positions[i+1]) + 207
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line PL_j6Y03.go:1*/ string(data)
		}()
		return  /*line wNjrRxTv.go:1*/ shim.Error(w5VhCLka)
	} else if tQcQN8RB == nil {
		w5VhCLka =  /*line r1Hqz6z8.go:1*/ func() string {
			var data []byte
			i := 12
			decryptKey := 250
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 23
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xc1\xcf"...,
					)
				case 25:
					data =  /*line r1Hqz6z8.go:1*/ append(data, "ӈ"...,
					)
					i = 6
				case 9:
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\x80\xc5\xc3\xd7"...,
					)
					i = 4
				case 4:
					data =  /*line r1Hqz6z8.go:1*/ append(data, 221)
					i = 22
				case 18:
					i = 3
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xb7\xb6"...,
					)
				case 24:
					i = 20
					data =  /*line r1Hqz6z8.go:1*/ append(data, 103)
				case 1:
					i = 0
					data =  /*line r1Hqz6z8.go:1*/ append(data, 132)
				case 0:
					i = 9
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xd9\xce\xcc"...,
					)
				case 11:
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xc5\xd5\xe7\xd3"...,
					)
					i = 25
				case 14:
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xc8\xca\xc8u"...,
					)
					i = 2
				case 15:
					data =  /*line r1Hqz6z8.go:1*/ append(data, 188)
					i = 17
				case 7:
					i = 18
					data =  /*line r1Hqz6z8.go:1*/ append(data, 97)
				case 13:
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xc1\xba\xb8"...,
					)
					i = 15
				case 23:
					data =  /*line r1Hqz6z8.go:1*/ append(data, "~\xc8\xcb"...,
					)
					i = 10
				case 2:
					i = 19
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xb7\xccĹ"...,
					)
				case 5:
					i = 21
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xb2\x99"...,
					)
				case 21:
					for y := range data {
						data[y] = data[y] -  /*line r1Hqz6z8.go:1*/ byte(decryptKey^y)
					}
					i = 8
				case 17:
					i = 24
					data =  /*line r1Hqz6z8.go:1*/ append(data, "m~"...,
					)
				case 12:
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xc7o\x93"...,
					)
					i = 13
				case 6:
					i = 16
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xcf\xd9"...,
					)
				case 16:
					i = 1
					data =  /*line r1Hqz6z8.go:1*/ append(data, 221)
				case 22:
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\xf0\xe3\xf3"...,
					)
					i = 5
				case 10:
					data =  /*line r1Hqz6z8.go:1*/ append(data, 121)
					i = 14
				case 20:
					data =  /*line r1Hqz6z8.go:1*/ append(data, "\x9a\xaf\xa5"...,
					)
					i = 7
				case 19:
					i = 11
					data =  /*line r1Hqz6z8.go:1*/ append(data, 193)
				}
			}
			return  /*line r1Hqz6z8.go:1*/ string(data)
		}() + rHDGEtgU +  /*line C5JYj_pc.go:1*/ func() string {
			fullData :=  /*line C5JYj_pc.go:1*/ []byte("\xc0e\xe2\xe2")
			var data []byte
			data =  /*line C5JYj_pc.go:1*/ append(data, fullData[3]^fullData[0], fullData[2]-fullData[1])
			return  /*line C5JYj_pc.go:1*/ string(data)
		}()
		return  /*line DG56Unok.go:1*/ shim.Error(w5VhCLka)
	}
	a8ZPVup7 := AuthorizedUsers{}
	cSW1wSO3 =  /*line j6jSF9I1.go:1*/ json.Unmarshal(tQcQN8RB, &a8ZPVup7)	                               
	if cSW1wSO3 != nil {
		return  /*line DqIB7OsI.go:1*/ shim.Error( /*line mJwJnsjL.go:1*/ cSW1wSO3.Error())
	}

	if (olYMzSWK ==  /*line GosIzXER.go:1*/ func() string {
		data :=  /*line GosIzXER.go:1*/ []byte("xva\x97")
		positions := [...]byte{1, 3, 1, 0}
		for i := 0; i < 4; i += 2 {
			localKey :=  /*line GosIzXER.go:1*/ byte(i) +  /*line GosIzXER.go:1*/ byte(positions[i]^positions[i+1]) + 16
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line GosIzXER.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line CMkiGyyN.go:1*/ func() string {
		data :=  /*line CMkiGyyN.go:1*/ []byte("mYdcX{")
		positions := [...]byte{1, 5, 5, 1, 3, 1, 4, 1}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line CMkiGyyN.go:1*/ byte(i) +  /*line CMkiGyyN.go:1*/ byte(positions[i]^positions[i+1]) + 44
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line CMkiGyyN.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line TnVt4zjw.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 27
		for counter := 0; i != 6; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 4
				data =  /*line TnVt4zjw.go:1*/ append(data, 55)
			case 2:
				data =  /*line TnVt4zjw.go:1*/ append(data, 53)
				i = 0
			case 0:
				i = 6
				for y := range data {
					data[y] = data[y] +  /*line TnVt4zjw.go:1*/ byte(decryptKey^y)
				}
			case 5:
				i = 2
				data =  /*line TnVt4zjw.go:1*/ append(data, 55)
			case 3:
				i = 7
				data =  /*line TnVt4zjw.go:1*/ append(data, 55)
			case 7:
				data =  /*line TnVt4zjw.go:1*/ append(data, 57)
				i = 8
			case 4:
				data =  /*line TnVt4zjw.go:1*/ append(data, 45)
				i = 3
			case 8:
				data =  /*line TnVt4zjw.go:1*/ append(data, 44)
				i = 5
			}
		}
		return  /*line TnVt4zjw.go:1*/ string(data)
	}()) && ((ezgSKu0t[0] != "") || (ezgSKu0t[0] !=  /*line q2x_DQYK.go:1*/ func() string {
		seed :=  /*line q2x_DQYK.go:1*/ byte(36)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line q2x_DQYK.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line q2x_DQYK.go:1*/ fnc(252)
		return  /*line q2x_DQYK.go:1*/ string(data)
	}())) {
		var iECyqXpw []string
		if olYMzSWK ==  /*line Pu47v6AN.go:1*/ func() string {
			seed :=  /*line Pu47v6AN.go:1*/ byte(249)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line Pu47v6AN.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line Pu47v6AN.go:1*/  /*line Pu47v6AN.go:1*/  /*line Pu47v6AN.go:1*/  /*line Pu47v6AN.go:1*/ fnc(121)(243)(252)(3)
			return  /*line Pu47v6AN.go:1*/ string(data)
		}() {
			cSW1wSO3 =  /*line bx32l4Dj.go:1*/ n4c7mRtG.DelState(rHDGEtgU)	                             
			if cSW1wSO3 != nil {
				return  /*line zVc_7yXU.go:1*/ shim.Error( /*line QajAWgtM.go:1*/ func() string {
					var data []byte
					i := 10
					decryptKey := 71
					for counter := 0; i != 2; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 6:
							i = 9
							data =  /*line QajAWgtM.go:1*/ append(data, "\xe8\xea"...,
							)
						case 8:
							i = 3
							data =  /*line QajAWgtM.go:1*/ append(data, "\xe7\u05cd\xdf"...,
							)
						case 4:
							i = 2
							for y := range data {
								data[y] = data[y] -  /*line QajAWgtM.go:1*/ byte(decryptKey^y)
							}
						case 1:
							i = 8
							data =  /*line QajAWgtM.go:1*/ append(data, 213)
						case 5:
							i = 1
							data =  /*line QajAWgtM.go:1*/ append(data, "\x94\xdb\xdb\xdd"...,
							)
						case 3:
							data =  /*line QajAWgtM.go:1*/ append(data, "\xe3\xcf"...,
							)
							i = 0
						case 0:
							i = 4
							data =  /*line QajAWgtM.go:1*/ append(data, "\xddͥ"...,
							)
						case 9:
							data =  /*line QajAWgtM.go:1*/ append(data, "\xde\xdc"...,
							)
							i = 7
						case 7:
							data =  /*line QajAWgtM.go:1*/ append(data, "\x9b\xee"...,
							)
							i = 11
						case 11:
							i = 5
							data =  /*line QajAWgtM.go:1*/ append(data, 228)
						case 10:
							i = 6
							data =  /*line QajAWgtM.go:1*/ append(data, "\xc3\xdd"...,
							)
						}
					}
					return  /*line QajAWgtM.go:1*/ string(data)
				}() +  /*line dNAZfdi5.go:1*/ cSW1wSO3.Error())
			}
		} else {
			if olYMzSWK ==  /*line YFKBTSO9.go:1*/ func() string {
				fullData :=  /*line YFKBTSO9.go:1*/ []byte("w\x89\xed\xfa\xe0\xdbD\x9c)\x8cz\v")
				var data []byte
				data =  /*line YFKBTSO9.go:1*/ append(data, fullData[8]^fullData[6], fullData[10]-fullData[11], fullData[5]-fullData[0], fullData[4]^fullData[1], fullData[3]^fullData[7], fullData[2]+fullData[9])
				return  /*line YFKBTSO9.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line IWAYrzpQ.go:1*/ func() string {
					key :=  /*line IWAYrzpQ.go:1*/ []byte("C\xf4|L")
					data :=  /*line IWAYrzpQ.go:1*/ []byte("1\x91\x1d(")
					for i, b := range key {
						data[i] = data[i] ^ b
					}
					return  /*line IWAYrzpQ.go:1*/ string(data)
				}()}
			} else if olYMzSWK ==  /*line VgpcrHia.go:1*/ func() string {
				var data []byte
				i := 6
				decryptKey := 54
				for counter := 0; i != 5; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 3:
						i = 5
						for y := range data {
							data[y] = data[y] +  /*line VgpcrHia.go:1*/ byte(decryptKey^y)
						}
					case 1:
						i = 3
						data =  /*line VgpcrHia.go:1*/ append(data, 76)
					case 8:
						i = 0
						data =  /*line VgpcrHia.go:1*/ append(data, 70)
					case 2:
						i = 4
						data =  /*line VgpcrHia.go:1*/ append(data, 63)
					case 0:
						i = 2
						data =  /*line VgpcrHia.go:1*/ append(data, 70)
					case 4:
						data =  /*line VgpcrHia.go:1*/ append(data, 72)
						i = 1
					case 6:
						data =  /*line VgpcrHia.go:1*/ append(data, 66)
						i = 7
					case 7:
						data =  /*line VgpcrHia.go:1*/ append(data, 54)
						i = 8
					}
				}
				return  /*line VgpcrHia.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line cAHrkDol.go:1*/ func() string {
					data :=  /*line cAHrkDol.go:1*/ []byte("[=\x10d")
					positions := [...]byte{0, 1, 2, 0, 1, 2}
					for i := 0; i < 6; i += 2 {
						localKey :=  /*line cAHrkDol.go:1*/ byte(i) +  /*line cAHrkDol.go:1*/ byte(positions[i]^positions[i+1]) + 94
						data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
					}
					return  /*line cAHrkDol.go:1*/ string(data)
				}(),  /*line pmzddS7E.go:1*/ func() string {
					key :=  /*line pmzddS7E.go:1*/ []byte("\xe2\x9a\x16À\x81")
					data :=  /*line pmzddS7E.go:1*/ []byte("O\tz,\xe6\xfa")
					for i, b := range key {
						data[i] = data[i] - b
					}
					return  /*line pmzddS7E.go:1*/ string(data)
				}()}
			}

			a8ZPVup7.Users.Permission = iECyqXpw

			jGct6K0c :=  /*line gwJizNFo.go:1*/ bszIYEJU(a8ZPVup7.Users.CustomAccessRights, ezgSKu0t)
			a8ZPVup7.Users.CustomAccessRights = jGct6K0c

			niQzOH6M, cSW1wSO3 :=  /*line ZTuFYLf2.go:1*/ json.Marshal(a8ZPVup7)
			if cSW1wSO3 != nil {
				return  /*line BT0c82N0.go:1*/ shim.Error( /*line zONEYaZm.go:1*/ cSW1wSO3.Error())
			}

			                        
			cSW1wSO3 =  /*line XpgdZKNZ.go:1*/ n4c7mRtG.PutState(rHDGEtgU, niQzOH6M)
			if cSW1wSO3 != nil {
				return  /*line NC5ltLbX.go:1*/ shim.Error( /*line wa5H4ALv.go:1*/ cSW1wSO3.Error())
			}
		}

	} else if (olYMzSWK ==  /*line p8Mp5uyy.go:1*/ func() string {
		fullData :=  /*line p8Mp5uyy.go:1*/ []byte("J7\v.g \xabD")
		var data []byte
		data =  /*line p8Mp5uyy.go:1*/ append(data, fullData[2]+fullData[4], fullData[1]+fullData[3], fullData[6]-fullData[0], fullData[5]+fullData[7])
		return  /*line p8Mp5uyy.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line yql94_P8.go:1*/ func() string {
		data :=  /*line yql94_P8.go:1*/ []byte("Zogi\xf3\xe0")
		positions := [...]byte{0, 2, 0, 0, 5, 4, 2, 2}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line yql94_P8.go:1*/ byte(i) +  /*line yql94_P8.go:1*/ byte(positions[i]^positions[i+1]) + 129
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line yql94_P8.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line K_ZR0M9o.go:1*/ func() string {
		key :=  /*line K_ZR0M9o.go:1*/ []byte("\\\xd0B\x12\x82U6")
		data :=  /*line K_ZR0M9o.go:1*/ []byte(",\xb50a\xeb&B")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line K_ZR0M9o.go:1*/ string(data)
	}()) {
		var iECyqXpw []string
		if olYMzSWK ==  /*line Rp3kz2fg.go:1*/ func() string {
			data :=  /*line Rp3kz2fg.go:1*/ []byte("eea=")
			positions := [...]byte{3, 0, 3, 3}
			for i := 0; i < 4; i += 2 {
				localKey :=  /*line Rp3kz2fg.go:1*/ byte(i) +  /*line Rp3kz2fg.go:1*/ byte(positions[i]^positions[i+1]) + 76
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line Rp3kz2fg.go:1*/ string(data)
		}() {
			cSW1wSO3 =  /*line BK8rQvSo.go:1*/ n4c7mRtG.DelState(rHDGEtgU)	                             
			if cSW1wSO3 != nil {
				return  /*line IwDtRGvY.go:1*/ shim.Error( /*line _imOBfVR.go:1*/ func() string {
					data :=  /*line _imOBfVR.go:1*/ []byte("rmil\ai\x02t<f\x18e\x15\x15te AtFte:")
					positions := [...]byte{13, 1, 4, 12, 0, 8, 19, 17, 9, 12, 9, 4, 12, 13, 19, 13, 8, 0, 1, 6, 19, 5, 10, 19, 17, 8}
					for i := 0; i < 26; i += 2 {
						localKey :=  /*line _imOBfVR.go:1*/ byte(i) +  /*line _imOBfVR.go:1*/ byte(positions[i]^positions[i+1]) + 74
						data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
					}
					return  /*line _imOBfVR.go:1*/ string(data)
				}() +  /*line soNhiCM2.go:1*/ cSW1wSO3.Error())
			}
		} else {
			if olYMzSWK ==  /*line V92HEccX.go:1*/ func() string {
				key :=  /*line V92HEccX.go:1*/ []byte("\xd5E\x94\"\xbd\xdd")
				data :=  /*line V92HEccX.go:1*/ []byte("B\xb4\xf8\x8b#V")
				for i, b := range key {
					data[i] = data[i] - b
				}
				return  /*line V92HEccX.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line IEWl_RUl.go:1*/ func() string {
					data :=  /*line IEWl_RUl.go:1*/ []byte("r\xeaw\xef")
					positions := [...]byte{1, 2, 3, 1}
					for i := 0; i < 4; i += 2 {
						localKey :=  /*line IEWl_RUl.go:1*/ byte(i) +  /*line IEWl_RUl.go:1*/ byte(positions[i]^positions[i+1]) + 134
						data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
					}
					return  /*line IEWl_RUl.go:1*/ string(data)
				}()}
			} else if olYMzSWK ==  /*line XYsFoXbS.go:1*/ func() string {
				seed :=  /*line XYsFoXbS.go:1*/ byte(90)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data =  /*line XYsFoXbS.go:1*/ append(data, x-seed); seed += x; return fnc }
				 /*line XYsFoXbS.go:1*/  /*line XYsFoXbS.go:1*/  /*line XYsFoXbS.go:1*/  /*line XYsFoXbS.go:1*/  /*line XYsFoXbS.go:1*/  /*line XYsFoXbS.go:1*/  /*line XYsFoXbS.go:1*/ fnc(202)(137)(31)(63)(116)(242)(229)
				return  /*line XYsFoXbS.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line PSYX7KYK.go:1*/ func() string {
					var data []byte
					i := 3
					decryptKey := 243
					for counter := 0; i != 2; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 4:
							data =  /*line PSYX7KYK.go:1*/ append(data, 83)
							i = 5
						case 3:
							i = 0
							data =  /*line PSYX7KYK.go:1*/ append(data, 98)
						case 1:
							i = 2
							for y := range data {
								data[y] = data[y] -  /*line PSYX7KYK.go:1*/ byte(decryptKey^y)
							}
						case 5:
							i = 1
							data =  /*line PSYX7KYK.go:1*/ append(data, 87)
						case 0:
							i = 4
							data =  /*line PSYX7KYK.go:1*/ append(data, 86)
						}
					}
					return  /*line PSYX7KYK.go:1*/ string(data)
				}(),  /*line Jh3V5S4c.go:1*/ func() string {
					data :=  /*line Jh3V5S4c.go:1*/ []byte("m\xff\x8ai\x86y")
					positions := [...]byte{2, 4, 4, 4, 2, 2, 1, 1}
					for i := 0; i < 8; i += 2 {
						localKey :=  /*line Jh3V5S4c.go:1*/ byte(i) +  /*line Jh3V5S4c.go:1*/ byte(positions[i]^positions[i+1]) + 106
						data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
					}
					return  /*line Jh3V5S4c.go:1*/ string(data)
				}()}
			}

			a8ZPVup7.Users.Permission = iECyqXpw

			niQzOH6M, cSW1wSO3 :=  /*line _nGPtbEy.go:1*/ json.Marshal(a8ZPVup7)
			if cSW1wSO3 != nil {
				return  /*line Fm5QVy0o.go:1*/ shim.Error( /*line tmTMht19.go:1*/ cSW1wSO3.Error())
			}

			                        
			cSW1wSO3 =  /*line fjo7CsfZ.go:1*/ n4c7mRtG.PutState(rHDGEtgU, niQzOH6M)
			if cSW1wSO3 != nil {
				return  /*line EyDXBJCL.go:1*/ shim.Error( /*line H0h7PKQK.go:1*/ cSW1wSO3.Error())
			}
		}

	} else if (ezgSKu0t[0] != "") || (ezgSKu0t[0] !=  /*line PQ1FTBd4.go:1*/ func() string {
		data :=  /*line PQ1FTBd4.go:1*/ []byte("\xa9")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line PQ1FTBd4.go:1*/ byte(i) +  /*line PQ1FTBd4.go:1*/ byte(positions[i]^positions[i+1]) + 137
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line PQ1FTBd4.go:1*/ string(data)
	}()) {

		jGct6K0c :=  /*line khwl9cmF.go:1*/ bszIYEJU(a8ZPVup7.Users.CustomAccessRights, ezgSKu0t)
		a8ZPVup7.Users.CustomAccessRights = jGct6K0c

		niQzOH6M, cSW1wSO3 :=  /*line pyDkuajI.go:1*/ json.Marshal(a8ZPVup7)
		if cSW1wSO3 != nil {
			return  /*line zkqTq1ng.go:1*/ shim.Error( /*line Q2Q2FXCw.go:1*/ cSW1wSO3.Error())
		}

		                        
		cSW1wSO3 =  /*line FqBPEdXI.go:1*/ n4c7mRtG.PutState(rHDGEtgU, niQzOH6M)
		if cSW1wSO3 != nil {
			return  /*line l71xkBRk.go:1*/ shim.Error( /*line PPm3zMFF.go:1*/ cSW1wSO3.Error())
		}

	}

	                                     
	dJEF9Lt2, cSW1wSO3 :=  /*line teJHQDok.go:1*/ t94C_D8c(n4c7mRtG, []string{eviEA3sp, w8EOizcx, igyLs9co, a8ZPVup7.Users.Organization, olYMzSWK, tIeFhn0D,  /*line M4TqrXyn.go:1*/ strings.Join(ezgSKu0t,  /*line nqcZWcXz.go:1*/ func() string {
		seed :=  /*line nqcZWcXz.go:1*/ byte(63)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line nqcZWcXz.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line nqcZWcXz.go:1*/ fnc(19)
		return  /*line nqcZWcXz.go:1*/ string(data)
	}()), a8ZPVup7.DatasetProviderOrg, a8ZPVup7.Description})
	if cSW1wSO3 != nil {
		return  /*line XGz5HEo0.go:1*/ shim.Error( /*line J_Rg34Fe.go:1*/ cSW1wSO3.Error())
	}
	 /*line V5ojrfmG.go:1*/ fmt.Println(dJEF9Lt2)

	xHTOVzVs :=  /*line yIYHdGU4.go:1*/ func() string {
		data :=  /*line yIYHdGU4.go:1*/ []byte("wE4 d_\x14~ h\x16sv{\\m#jr\x11\x12i\x06e\x04JTo\nNdea\x1d\x03+t7")
		positions := [...]byte{0, 25, 10, 37, 26, 7, 4, 21, 21, 2, 16, 25, 1, 12, 16, 29, 33, 13, 17, 20, 29, 6, 19, 13, 28, 13, 15, 22, 10, 35, 25, 24, 34, 31, 22, 26, 33, 23, 24, 12, 14, 2, 32, 21, 33, 15, 21, 13, 5, 13, 34, 28, 35, 17}
		for i := 0; i < 54; i += 2 {
			localKey :=  /*line yIYHdGU4.go:1*/ byte(i) +  /*line yIYHdGU4.go:1*/ byte(positions[i]^positions[i+1]) + 5
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line yIYHdGU4.go:1*/ string(data)
	}() + eviEA3sp
	q8ZF8zaz :=  /*line LkeDTKWI.go:1*/ []byte(xHTOVzVs)
	return  /*line v4kXrnjy.go:1*/ shim.Success(q8ZF8zaz)

}

func (g4rnrSNn *G1Y_7pz_) ziwzP1ja(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	if  /*line L4Dn0ptI.go:1*/ len(ktsi1_SQ) != 2 {
		return  /*line qzRAOvn1.go:1*/ shim.Error( /*line BSDlVWVf.go:1*/ func() string {
			data :=  /*line BSDlVWVf.go:1*/ []byte("ȸ\xcco\xb3L\xe9J\xa3 a\xbdm\xd8e\x18 \xa2\xb8\xaa\xf4(\xd4:mekts8\x9f\rrZfu8\x94ri\xa7Qɥ\xb6l\xeb\xe4\xe2\xdbe\x88\xaa\xca\xfft\xb8Nڸ\xc8\xd3\xfc(\xb5n\xeaV\xe5")
			positions := [...]byte{48, 10, 4, 29, 58, 30, 62, 36, 63, 30, 20, 7, 33, 15, 23, 31, 53, 2, 33, 64, 59, 57, 62, 46, 13, 26, 63, 54, 47, 15, 8, 58, 20, 46, 41, 68, 21, 6, 37, 11, 11, 33, 4, 66, 54, 33, 5, 41, 44, 26, 11, 22, 33, 51, 41, 49, 20, 8, 43, 30, 13, 49, 6, 40, 59, 49, 48, 31, 17, 0, 1, 22, 52, 31, 61, 60, 56, 38, 67, 44, 42, 41, 4, 18, 19, 59}
			for i := 0; i < 86; i += 2 {
				localKey :=  /*line BSDlVWVf.go:1*/ byte(i) +  /*line BSDlVWVf.go:1*/ byte(positions[i]^positions[i+1]) + 82
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line BSDlVWVf.go:1*/ string(data)
		}())
	}
	hSHnto8d := ktsi1_SQ[0]
	tIeFhn0D := ktsi1_SQ[1]

	lWKzd4Pe, cSW1wSO3 :=  /*line uY_DQQGz.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line Iz03vyUC.go:1*/ shim.Error( /*line zDFKicS4.go:1*/ cSW1wSO3.Error())
	}
	qo8GOXik := RequestAccess{}
	mGZijNgp, cSW1wSO3 :=  /*line Ui4jWQ1M.go:1*/ n4c7mRtG.GetState(hSHnto8d)	                                      
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line BNBdMmzw.go:1*/ func() string {
			fullData :=  /*line BNBdMmzw.go:1*/ []byte("\x87\xaa\x1cv9\xddvB\xd5\xe3\x98\xd3B\xe8\xf4\xbaU\x94\x83b\xffhg\x9aW;\xfb\xc1\x85-V\xc1\x84\xee\x00\xff\x95\xd3I\x14\x06n\x0f\xfc\xe7 \xa9HE\t\xe7\xb33>\xba#\xd8R$\x8a\x90\x10V\x81q\xbbb\n\xaf\xaf\b/\f@\rѷ'B\xfe\x1c\xe6 \xf4\xf0D\x06E\xe5<\x0e\xe9ɚXj[\xdd\x1e\xb4\xc9\xeas\xe2u{|_")
			var data []byte
			data =  /*line BNBdMmzw.go:1*/ append(data, fullData[31]^fullData[54], fullData[75]-fullData[68], fullData[59]-fullData[48], fullData[80]+fullData[30], fullData[65]^fullData[92], fullData[73]+fullData[71], fullData[43]+fullData[3], fullData[70]-fullData[81], fullData[44]^fullData[97], fullData[58]^fullData[86], fullData[28]+fullData[27], fullData[36]^fullData[83], fullData[101]^fullData[18], fullData[90]^fullData[66], fullData[1]-fullData[87], fullData[94]-fullData[14], fullData[96]^fullData[105], fullData[103]-fullData[41], fullData[84]-fullData[63], fullData[49]-fullData[91], fullData[99]+fullData[51], fullData[16]+fullData[61], fullData[11]-fullData[107], fullData[88]+fullData[25], fullData[100]-fullData[24], fullData[7]^fullData[77], fullData[15]-fullData[38], fullData[102]^fullData[40], fullData[95]^fullData[42], fullData[60]^fullData[9], fullData[72]-fullData[10], fullData[64]+fullData[69], fullData[13]-fullData[0], fullData[106]+fullData[50], fullData[104]+fullData[33], fullData[78]+fullData[55], fullData[8]-fullData[19], fullData[39]^fullData[22], fullData[6]-fullData[62], fullData[32]-fullData[98], fullData[26]^fullData[17], fullData[67]+fullData[21], fullData[34]+fullData[82], fullData[57]+fullData[45], fullData[20]-fullData[93], fullData[52]+fullData[53], fullData[4]+fullData[89], fullData[23]^fullData[35], fullData[76]-fullData[85], fullData[5]^fullData[46], fullData[47]+fullData[56], fullData[2]-fullData[37], fullData[12]-fullData[79], fullData[29]-fullData[74])
			return  /*line BNBdMmzw.go:1*/ string(data)
		}() + hSHnto8d +  /*line hzkGd_lE.go:1*/ func() string {
			data :=  /*line hzkGd_lE.go:1*/ []byte("\"\xfb")
			positions := [...]byte{1, 1}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line hzkGd_lE.go:1*/ byte(i) +  /*line hzkGd_lE.go:1*/ byte(positions[i]^positions[i+1]) + 130
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line hzkGd_lE.go:1*/ string(data)
		}()
		return  /*line EGXEAgqc.go:1*/ shim.Error(w5VhCLka)
	} else if mGZijNgp == nil {
		w5VhCLka =  /*line IZmUomfk.go:1*/ func() string {
			var data []byte
			i := 4
			decryptKey := 27
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					data =  /*line IZmUomfk.go:1*/ append(data, ":-"...,
					)
					i = 1
				case 7:
					data =  /*line IZmUomfk.go:1*/ append(data, "{v"...,
					)
					i = 13
				case 6:
					data =  /*line IZmUomfk.go:1*/ append(data, "t&w"...,
					)
					i = 12
				case 16:
					i = 15
					data =  /*line IZmUomfk.go:1*/ append(data, 110)
				case 15:
					i = 2
					data =  /*line IZmUomfk.go:1*/ append(data, "iuk"...,
					)
				case 10:
					data =  /*line IZmUomfk.go:1*/ append(data, "dsK\x1e"...,
					)
					i = 9
				case 11:
					data =  /*line IZmUomfk.go:1*/ append(data, "isc"...,
					)
					i = 14
				case 4:
					i = 5
					data =  /*line IZmUomfk.go:1*/ append(data, "d<"...,
					)
				case 1:
					i = 7
					data =  /*line IZmUomfk.go:1*/ append(data, "4Q"...,
					)
				case 12:
					i = 10
					data =  /*line IZmUomfk.go:1*/ append(data, "arw"...,
					)
				case 3:
					i = 11
					data =  /*line IZmUomfk.go:1*/ append(data, "z-"...,
					)
				case 13:
					data =  /*line IZmUomfk.go:1*/ append(data, 97)
					i = 8
				case 14:
					data =  /*line IZmUomfk.go:1*/ append(data, "z|"...,
					)
					i = 6
				case 5:
					i = 16
					data =  /*line IZmUomfk.go:1*/ append(data, 88)
				case 9:
					for y := range data {
						data[y] = data[y] ^  /*line IZmUomfk.go:1*/ byte(decryptKey^y)
					}
					i = 0
				case 8:
					i = 3
					data =  /*line IZmUomfk.go:1*/ append(data, "1~`"...,
					)
				}
			}
			return  /*line IZmUomfk.go:1*/ string(data)
		}() + hSHnto8d +  /*line Hdzui7zh.go:1*/ func() string {
			var data []byte
			i := 1
			decryptKey := 159
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 1:
					data =  /*line Hdzui7zh.go:1*/ append(data, 186)
					i = 3
				case 2:
					for y := range data {
						data[y] = data[y] ^  /*line Hdzui7zh.go:1*/ byte(decryptKey^y)
					}
					i = 0
				case 3:
					i = 2
					data =  /*line Hdzui7zh.go:1*/ append(data, 228)
				}
			}
			return  /*line Hdzui7zh.go:1*/ string(data)
		}()
		return  /*line p50Gbd1r.go:1*/ shim.Error(w5VhCLka)
	}

	cSW1wSO3 =  /*line G4mTBSF6.go:1*/ json.Unmarshal(mGZijNgp, &qo8GOXik)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line kv5zMVZh.go:1*/ func() string {
			data :=  /*line kv5zMVZh.go:1*/ []byte("q\"\xe2\xd1r\xe0\xb5\xd4:\x13\xefaSle\x1d \xcco\xc2\xfee\xbbo\x91\\ FVOp")
			positions := [...]byte{28, 28, 15, 7, 19, 12, 20, 25, 30, 9, 30, 0, 15, 30, 19, 22, 28, 5, 28, 10, 2, 10, 3, 20, 28, 27, 9, 24, 28, 27, 17, 0, 30, 7, 17, 6}
			for i := 0; i < 36; i += 2 {
				localKey :=  /*line kv5zMVZh.go:1*/ byte(i) +  /*line kv5zMVZh.go:1*/ byte(positions[i]^positions[i+1]) + 136
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line kv5zMVZh.go:1*/ string(data)
		}() + hSHnto8d +  /*line YClJOF6O.go:1*/ func() string {
			var data []byte
			i := 1
			decryptKey := 11
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					for y := range data {
						data[y] = data[y] ^  /*line YClJOF6O.go:1*/ byte(decryptKey^y)
					}
					i = 2
				case 1:
					data =  /*line YClJOF6O.go:1*/ append(data, 42)
					i = 3
				case 3:
					data =  /*line YClJOF6O.go:1*/ append(data, 116)
					i = 0
				}
			}
			return  /*line YClJOF6O.go:1*/ string(data)
		}()
		return  /*line MVce9QjC.go:1*/ shim.Error(w5VhCLka)
	}

	if lWKzd4Pe != qo8GOXik.Dataset_Provider {
		return  /*line C4Gh23WM.go:1*/ shim.Error( /*line E69wuLsI.go:1*/ func() string {
			key :=  /*line E69wuLsI.go:1*/ []byte("1Z\xe0\xd7,\xe0\xa53_E\"\xe2*\x1aY\xbdSuG\xaeH\x82\x94\x89G\x8f'9<\x8c\xa5\x0f")
			data :=  /*line E69wuLsI.go:1*/ []byte("\u007f\xc9T\xf7\x8dU\x19\x9bη\x8b\\\x8f~y1\u0095\xab\x13\xb4\xe7\b\xeeg\x01\x8c\xaa\xb1\xf1\x18\x83")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line E69wuLsI.go:1*/ string(data)
		}())

	}
	cSW1wSO3 =  /*line XjDJzkG2.go:1*/ n4c7mRtG.DelState(hSHnto8d)	                                                                
	if cSW1wSO3 != nil {
		return  /*line FvdhFe7D.go:1*/ shim.Error( /*line e5ceCfEQ.go:1*/ func() string {
			fullData :=  /*line e5ceCfEQ.go:1*/ []byte("\x91lg\x92\x97\ns8m\xa2\x87p>\x03\xf7\x97L\xb1\x11\xb0(N\xb3\xfa\xa7\xe3\x13\x8c\xd8w\x9c\xd9s\xbb\xb76\xa5\"\x82\xdf~AR\xee\xd8@\xf3\x16\x94\x1b\x8d+V\xbf\xf9\xc5\x12q \xfc_\xfcc-")
			var data []byte
			data =  /*line e5ceCfEQ.go:1*/ append(data, fullData[2]+fullData[39], fullData[32]-fullData[56], fullData[41]^fullData[20], fullData[33]+fullData[17], fullData[50]+fullData[28], fullData[53]+fullData[36], fullData[27]+fullData[48], fullData[7]^fullData[16], fullData[34]^fullData[44], fullData[0]-fullData[57], fullData[46]^fullData[4], fullData[3]-fullData[63], fullData[24]+fullData[55], fullData[10]-fullData[37], fullData[15]^fullData[25], fullData[26]+fullData[42], fullData[38]^fullData[9], fullData[62]^fullData[18], fullData[21]^fullData[51], fullData[52]+fullData[49], fullData[47]+fullData[60], fullData[12]-fullData[31], fullData[8]-fullData[23], fullData[40]-fullData[5], fullData[45]-fullData[58], fullData[6]+fullData[43], fullData[19]+fullData[22], fullData[14]+fullData[1], fullData[30]^fullData[54], fullData[29]+fullData[59], fullData[13]+fullData[11], fullData[35]-fullData[61])
			return  /*line e5ceCfEQ.go:1*/ string(data)
		}() +  /*line iUpcIzxG.go:1*/ cSW1wSO3.Error())
	}
	                                                      
	               
	xwK9eDsS, cSW1wSO3 :=  /*line nI4_25IN.go:1*/ wjnibpdu(n4c7mRtG, []string{qo8GOXik.Dataset_Name, lWKzd4Pe, qo8GOXik.Username, qo8GOXik.Organization, hSHnto8d,  /*line p9_lPMfY.go:1*/ func() string {
		key :=  /*line p9_lPMfY.go:1*/ []byte("HH\xa4\n")
		data :=  /*line p9_lPMfY.go:1*/ []byte("\f-\xcas")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line p9_lPMfY.go:1*/ string(data)
	}(), tIeFhn0D, qo8GOXik.Permission,  /*line HbbjaOop.go:1*/ strings.Join(qo8GOXik.CustomAccessRights,  /*line Ser2W92y.go:1*/ func() string {
		fullData :=  /*line Ser2W92y.go:1*/ []byte("\xd1[")
		var data []byte
		data =  /*line Ser2W92y.go:1*/ append(data, fullData[0]+fullData[1])
		return  /*line Ser2W92y.go:1*/ string(data)
	}()), qo8GOXik.DatasetProviderOrg, qo8GOXik.Description})
	if cSW1wSO3 != nil {
		return  /*line CwgXRHhn.go:1*/ shim.Error( /*line JI3JNzDe.go:1*/ cSW1wSO3.Error())
	}
	 /*line Q546ho5l.go:1*/ fmt.Println(xwK9eDsS)
	   

	xHTOVzVs :=  /*line RYIyuUdW.go:1*/ func() string {
		key :=  /*line RYIyuUdW.go:1*/ []byte("\xb6X\xa9\x05\x01W\xb52\x87")
		data :=  /*line RYIyuUdW.go:1*/ []byte("\x9c\r\xc8pd\x1c\xbf\b\x99")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line RYIyuUdW.go:1*/ string(data)
	}() + qo8GOXik.RequestID +  /*line Azk2W0W6.go:1*/ func() string {
		fullData :=  /*line Azk2W0W6.go:1*/ []byte("rxh@5Z\xc3[kF`\x05\xf7~\x88\xe3\xa2\x00\x15\x17/ޑ\xfe\x9a`]אo")
		var data []byte
		data =  /*line Azk2W0W6.go:1*/ append(data, fullData[10]^fullData[3], fullData[21]-fullData[1], fullData[17]^fullData[0], fullData[5]+fullData[18], fullData[23]-fullData[22], fullData[13]+fullData[16], fullData[8]-fullData[12], fullData[27]-fullData[29], fullData[11]^fullData[25], fullData[14]-fullData[2], fullData[9]+fullData[20], fullData[15]+fullData[28], fullData[24]-fullData[4], fullData[19]+fullData[7], fullData[6]+fullData[26])
		return  /*line Azk2W0W6.go:1*/ string(data)
	}() + qo8GOXik.Name +  /*line hf1W0AbQ.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 81
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 0
				data =  /*line hf1W0AbQ.go:1*/ append(data, 207)
			case 0:
				for y := range data {
					data[y] = data[y] +  /*line hf1W0AbQ.go:1*/ byte(decryptKey^y)
				}
				i = 2
			}
		}
		return  /*line hf1W0AbQ.go:1*/ string(data)
	}() + qo8GOXik.Surname +  /*line CLOyveNS.go:1*/ func() string {
		key :=  /*line CLOyveNS.go:1*/ []byte("5]\xe3\xfdAoq\x11W\f7\x82D\x88O\x0f\xe5S")
		data :=  /*line CLOyveNS.go:1*/ []byte("\x155\x82\x8ea\r\x14t9,S\xe7(\xed;j\x81s")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line CLOyveNS.go:1*/ string(data)
	}() +  /*line QGs2G1Rn.go:1*/ func() string {
		var data []byte
		i := 0
		decryptKey := 83
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				i = 9
				data =  /*line QGs2G1Rn.go:1*/ append(data, 33)
			case 7:
				i = 2
				data =  /*line QGs2G1Rn.go:1*/ append(data, "~n\u007ft"...,
				)
			case 4:
				data =  /*line QGs2G1Rn.go:1*/ append(data, "k)"...,
				)
				i = 5
			case 6:
				data =  /*line QGs2G1Rn.go:1*/ append(data, 49)
				i = 8
			case 2:
				i = 6
				data =  /*line QGs2G1Rn.go:1*/ append(data, 130)
			case 5:
				data =  /*line QGs2G1Rn.go:1*/ append(data, "ll"...,
				)
				i = 7
			case 8:
				for y := range data {
					data[y] = data[y] -  /*line QGs2G1Rn.go:1*/ byte(decryptKey^y)
				}
				i = 1
			case 9:
				data =  /*line QGs2G1Rn.go:1*/ append(data, "frt%"...,
				)
				i = 3
			case 3:
				data =  /*line QGs2G1Rn.go:1*/ append(data, "xo"...,
				)
				i = 4
			}
		}
		return  /*line QGs2G1Rn.go:1*/ string(data)
	}() + qo8GOXik.Dataset_Name
	q8ZF8zaz :=  /*line HQlwYZcz.go:1*/ []byte(xHTOVzVs)
	return  /*line Fi5Yj8wB.go:1*/ shim.Success(q8ZF8zaz)
}

                                             
func (g4rnrSNn *G1Y_7pz_) dpTyuH84(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	                  
	lWKzd4Pe, cSW1wSO3 :=  /*line mRT61vop.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line cqk0nONo.go:1*/ shim.Error( /*line weQIIrs1.go:1*/ cSW1wSO3.Error())
	}
	qVH0siY7 :=  /*line Fo55KC7c.go:1*/ fmt.Sprintf( /*line HbNuz0J_.go:1*/ func() string {
		fullData :=  /*line HbNuz0J_.go:1*/ []byte("\xe1}\x04\xf0\xff\x1c\xca\x05\xab\xb7Q\x14F\x9f|\x95\x9c\xadt\xb8\xd7p\x9eN\xfd\xf7\x98\xa73\xe5\xce&9D\xaf\xde\xfa\xf1\xeaXd#_k\x94\x8e\x1c\xba\xf5\xd5\xdb\x1f\xdb.UV9M\x9a\x8a\x94g\x0e\x80\xe1$n\xca\x02\x87\x11\xb6\x86mT\x13\xb8\x95\xfa=dIƞg\xf5>\xfc\xf7$\xd3O$ptL'q\x17\xb1p\x13jAP\x19\x02\x8e\xdfҪf\f\x03\xa5M\x00\x87\u03a24e\xba\xa8\xfc\x9d\x93N\xaf\xff*o")
		var data []byte
		data =  /*line HbNuz0J_.go:1*/ append(data, fullData[128]-fullData[120], fullData[91]+fullData[90], fullData[129]+fullData[94], fullData[78]-fullData[77], fullData[20]-fullData[43], fullData[19]+fullData[17], fullData[39]-fullData[48], fullData[126]+fullData[0], fullData[52]+fullData[60], fullData[93]^fullData[68], fullData[76]+fullData[102], fullData[92]-fullData[38], fullData[87]^fullData[117], fullData[3]^fullData[109], fullData[61]^fullData[113], fullData[49]^fullData[122], fullData[6]-fullData[84], fullData[125]+fullData[9], fullData[88]^fullData[107], fullData[22]-fullData[53], fullData[8]-fullData[12], fullData[42]^fullData[1], fullData[115]-fullData[75], fullData[83]-fullData[14], fullData[130]+fullData[98], fullData[10]+fullData[65], fullData[27]-fullData[28], fullData[81]+fullData[51], fullData[32]-fullData[67], fullData[66]-fullData[124], fullData[70]-fullData[123], fullData[89]-fullData[110], fullData[50]+fullData[59], fullData[55]+fullData[62], fullData[116]^fullData[54], fullData[79]^fullData[127], fullData[24]^fullData[26], fullData[85]^fullData[69], fullData[35]+fullData[15], fullData[86]-fullData[46], fullData[36]-fullData[118], fullData[16]+fullData[72], fullData[103]+fullData[41], fullData[112]^fullData[73], fullData[30]^fullData[47], fullData[111]-fullData[7], fullData[131]^fullData[5], fullData[4]^fullData[58], fullData[2]^fullData[100], fullData[33]-fullData[29], fullData[108]^fullData[34], fullData[31]+fullData[95], fullData[45]^fullData[64], fullData[121]^fullData[101], fullData[104]^fullData[56], fullData[99]-fullData[57], fullData[97]^fullData[11], fullData[82]-fullData[74], fullData[119]^fullData[63], fullData[114]^fullData[13], fullData[71]^fullData[44], fullData[96]-fullData[106], fullData[80]-fullData[37], fullData[21]-fullData[23], fullData[105]+fullData[40], fullData[18]-fullData[25])
		return  /*line HbNuz0J_.go:1*/ string(data)
	}(), lWKzd4Pe)
	bxM55Mmo, cSW1wSO3 :=  /*line aT9OJMjG.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line _2LkTsdN.go:1*/ shim.Error( /*line BFqkB3zg.go:1*/ cSW1wSO3.Error())
	}
	return  /*line aZNpJrj6.go:1*/ shim.Success(bxM55Mmo)
}

                                               
func (g4rnrSNn *G1Y_7pz_) yHTyS5ce(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	nTXZkyTM, cSW1wSO3 :=  /*line AW6nMwkG.go:1*/ g4rnrSNn.dHINeXgX(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line GPM3qUzo.go:1*/ shim.Error( /*line R79351ks.go:1*/ cSW1wSO3.Error())
	}
	if nTXZkyTM !=  /*line OpRrP9zK.go:1*/ func() string {
		fullData :=  /*line OpRrP9zK.go:1*/ []byte("\aP\xc11i9]S\xa6:")
		var data []byte
		data =  /*line OpRrP9zK.go:1*/ append(data, fullData[1]^fullData[3], fullData[2]-fullData[6], fullData[8]-fullData[5], fullData[9]^fullData[7], fullData[0]^fullData[4])
		return  /*line OpRrP9zK.go:1*/ string(data)
	}() {
		return  /*line f93i08fY.go:1*/ shim.Error( /*line Gjk2quHJ.go:1*/ func() string {
			key :=  /*line Gjk2quHJ.go:1*/ []byte(";\xc4˜\xd63\xd2L\x81!\x87\x87R\xfd\xeb\xfd\xba\xf7\x91\r}j#\x89\x86Bt\"\xf5\x9c\x107\x95>\x99$\x9d\xf0\xb6\x1e4\x8e\x82\x05m\xfcK\xe8\x1e\x15")
			data :=  /*line Gjk2quHJ.go:1*/ []byte("\x14\xaa\xa1\xddJ<\xa0\x1b\x9f@\xdd\xe6\x17q5f\xa7w\x8fd\xf8\xfbO\xf0\x9a.\xf1Px\xcdc<\xd41\xd5O\x83v\xb9T\xec\xe1\xf0b\xb3y(}T^")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line Gjk2quHJ.go:1*/ string(data)
		}())
	}

	lBdZ8ktW, cSW1wSO3 :=  /*line Fj1wUUCh.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line b4bjxgvn.go:1*/ shim.Error( /*line VhwAq8ao.go:1*/ cSW1wSO3.Error())
	}

	qVH0siY7 :=  /*line NFxzIA6B.go:1*/ fmt.Sprintf( /*line hAYXxjDK.go:1*/ func() string {
		key :=  /*line hAYXxjDK.go:1*/ []byte("\x13\x19\x06\x1f&\xf5\x81\xb7\x9b\xbf\xb2ie\xbc\xf9\x18<\xd4D\x9f\a\xe6\x150\x87\x8e\x89\xd39\x98r{iZG~\xabI\x17\xfc\xdb\xd3\xfc'\xceBq\xaf\a\xe38\a\x84\x06#\xf5O\x97\x84\xeaT\xfc\xbb\x93h0MF")
		data :=  /*line hAYXxjDK.go:1*/ []byte("h;uzJ\x90\xe2\xc3\xf4͐S\x1e\x9e\x9dw_\x80=\xefb\xc4/\x12\xc6\xfb\xfd\xbbV\xea\x1b\x01\f>\x12\r\xce;d\xde\xf7\xf1\x89T\xab0\x02\x81h\x91_f\xeaoY\x94;\xfe\xeb\x84vƙ\xb6\x1b\x120;")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line hAYXxjDK.go:1*/ string(data)
	}(), lBdZ8ktW)
	bxM55Mmo, cSW1wSO3 :=  /*line IMvMobT2.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line UGlcZ7G6.go:1*/ shim.Error( /*line troL4EpD.go:1*/ cSW1wSO3.Error())
	}
	return  /*line D1SsEFyB.go:1*/ shim.Success(bxM55Mmo)
}

                                             
func (g4rnrSNn *G1Y_7pz_) q2ZIG7Ai(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line uJxa9rwU.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line GRlZMO_e.go:1*/ shim.Error( /*line wOkMNyLW.go:1*/ func() string {
			var data []byte
			i := 9
			decryptKey := 75
			for counter := 0; i != 12; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					data =  /*line wOkMNyLW.go:1*/ append(data, "\xed\xf9\xa4"...,
					)
					i = 3
				case 13:
					data =  /*line wOkMNyLW.go:1*/ append(data, "%\x1c"...,
					)
					i = 5
				case 16:
					data =  /*line wOkMNyLW.go:1*/ append(data, "\n\x02\x0f"...,
					)
					i = 18
				case 6:
					data =  /*line wOkMNyLW.go:1*/ append(data, "\xbd\v\x05"...,
					)
					i = 8
				case 3:
					i = 19
					data =  /*line wOkMNyLW.go:1*/ append(data, "\xf5\xfb\xee"...,
					)
				case 5:
					data =  /*line wOkMNyLW.go:1*/ append(data, "\x14\x11\x1d"...,
					)
					i = 14
				case 1:
					for y := range data {
						data[y] = data[y] -  /*line wOkMNyLW.go:1*/ byte(decryptKey^y)
					}
					i = 12
				case 15:
					data =  /*line wOkMNyLW.go:1*/ append(data, "\xfa\xf0"...,
					)
					i = 4
				case 0:
					data =  /*line wOkMNyLW.go:1*/ append(data, 213)
					i = 1
				case 19:
					data =  /*line wOkMNyLW.go:1*/ append(data, "\xe2\xe8\xf4"...,
					)
					i = 6
				case 7:
					data =  /*line wOkMNyLW.go:1*/ append(data, "\x04\xbe\xb3\xd7"...,
					)
					i = 13
				case 9:
					data =  /*line wOkMNyLW.go:1*/ append(data, "\xd6\xfa\xf2\xfd"...,
					)
					i = 17
				case 8:
					data =  /*line wOkMNyLW.go:1*/ append(data, 190)
					i = 2
				case 2:
					data =  /*line wOkMNyLW.go:1*/ append(data, 250)
					i = 16
				case 10:
					i = 11
					data =  /*line wOkMNyLW.go:1*/ append(data, "\xf9\x05"...,
					)
				case 18:
					i = 10
					data =  /*line wOkMNyLW.go:1*/ append(data, 2)
				case 11:
					i = 7
					data =  /*line wOkMNyLW.go:1*/ append(data, 10)
				case 14:
					data =  /*line wOkMNyLW.go:1*/ append(data, "\x11\x19\x11\xc5"...,
					)
					i = 0
				case 17:
					i = 15
					data =  /*line wOkMNyLW.go:1*/ append(data, 251)
				}
			}
			return  /*line wOkMNyLW.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]

	                  
	lWKzd4Pe, cSW1wSO3 :=  /*line n4LGvuwz.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line BhYDs2If.go:1*/ shim.Error( /*line tg1TaAO3.go:1*/ cSW1wSO3.Error())
	}
	qVH0siY7 :=  /*line v9xqh1fn.go:1*/ fmt.Sprintf( /*line rZrOwazN.go:1*/ func() string {
		fullData :=  /*line rZrOwazN.go:1*/ []byte("3%\x0f\xe0D>\xba'\xd38\x1d,]\b\xa1\xc0\xe22S\xd0oeB\x19pL\xc5m\x90z_\x03˄A\"\x00\xa6sM\x9aw\xb1`Ok[P}\x88\xa2\xb55Rf6\x12o\xd6\x18\x93Grh\x81\x1e\x0045$\xce#\x02\xecI\xba{\x84\xfcx\xe7:x<\xf9\a\x87M\x85Jq\xd6@\xee\xe7\xear\x11\x8e\x02Iߦ\n%,J\x8f\x02\x05\xe5k\xaf,,w{\xb3e\x1e;\xb9\x8f\u007f\xf6\x18}\xf4ׂ\v\a\x10\xd8\xf9m\xa9F\x14\xc4\xf3?\x0e\tz\x80<\xa7\x9a\xea\x1c$2bg\xbe\x1a\x1b\x99ʐ\xe2s\xbdd\x16\x19\xa9\x17(\xba\x1e")
		var data []byte
		data =  /*line rZrOwazN.go:1*/ append(data, fullData[65]^fullData[21], fullData[171]-fullData[78], fullData[151]+fullData[44], fullData[14]^fullData[139], fullData[118]-fullData[134], fullData[61]^fullData[35], fullData[159]-fullData[154], fullData[167]-fullData[68], fullData[27]^fullData[108], fullData[147]+fullData[32], fullData[9]^fullData[156], fullData[95]+fullData[47], fullData[52]-fullData[6], fullData[49]+fullData[148], fullData[51]+fullData[112], fullData[36]^fullData[57], fullData[110]-fullData[129], fullData[8]^fullData[86], fullData[127]+fullData[88], fullData[16]-fullData[96], fullData[107]+fullData[91], fullData[142]^fullData[113], fullData[75]^fullData[145], fullData[164]^fullData[137], fullData[55]+fullData[130], fullData[48]^fullData[13], fullData[135]^fullData[166], fullData[25]+fullData[150], fullData[115]^fullData[125], fullData[62]-fullData[66], fullData[136]^fullData[15], fullData[102]-fullData[105], fullData[98]+fullData[128], fullData[82]-fullData[138], fullData[54]-fullData[97], fullData[11]^fullData[30], fullData[1]^fullData[92], fullData[158]-fullData[7], fullData[93]-fullData[76], fullData[132]^fullData[152], fullData[71]^fullData[2], fullData[67]-fullData[56], fullData[85]+fullData[12], fullData[19]^fullData[42], fullData[60]^fullData[94], fullData[124]+fullData[45], fullData[163]^fullData[70], fullData[74]^fullData[114], fullData[50]^fullData[58], fullData[162]+fullData[73], fullData[101]-fullData[90], fullData[29]^fullData[157], fullData[117]+fullData[170], fullData[119]^fullData[116], fullData[53]^fullData[24], fullData[4]-fullData[103], fullData[20]-fullData[39], fullData[160]-fullData[111], fullData[100]^fullData[81], fullData[43]^fullData[22], fullData[123]^fullData[18], fullData[106]^fullData[63], fullData[155]+fullData[37], fullData[144]+fullData[80], fullData[140]+fullData[64], fullData[133]^fullData[121], fullData[46]+fullData[59], fullData[165]^fullData[38], fullData[126]-fullData[143], fullData[40]^fullData[26], fullData[5]+fullData[17], fullData[0]+fullData[141], fullData[122]+fullData[3], fullData[79]-fullData[72], fullData[87]^fullData[69], fullData[104]^fullData[34], fullData[31]+fullData[153], fullData[28]^fullData[161], fullData[89]-fullData[169], fullData[146]-fullData[99], fullData[120]^fullData[23], fullData[83]-fullData[168], fullData[149]-fullData[41], fullData[10]+fullData[109], fullData[33]-fullData[131], fullData[84]^fullData[77])
		return  /*line rZrOwazN.go:1*/ string(data)
	}(), eviEA3sp, lWKzd4Pe)
	bxM55Mmo, cSW1wSO3 :=  /*line T2_HeSmz.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line fGqyjqsX.go:1*/ shim.Error( /*line vMltVQF0.go:1*/ cSW1wSO3.Error())
	}
	return  /*line KtnwRKCC.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) vJc5XjHG(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line rzynI1PY.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line RlhWK7dh.go:1*/ shim.Error( /*line wRCvo4yC.go:1*/ func() string {
			key :=  /*line wRCvo4yC.go:1*/ []byte("~|\xb2λ\x87_\x9a\xbc\xfd\x03\x16\xae\xaafW\xf2Mzζ\x12+\x10\xad\xe9\xda\xdf\xf9\xe8\xa0\rû\xabz\xf2u\xff1\xa8\xc0")
			data :=  /*line wRCvo4yC.go:1*/ []byte("7\x12ѡ\xc9\xf5:\xf9\xc8\xddmc\xc3\xc8\x03%\xd2\"\x1c\xee\xd7`Le\xc0\x8c\xb4\xab\x8aƀH\xbb\xcb\xce\x19\x86\x1c\x91V\x88\xf1")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line wRCvo4yC.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]

	                  
	ml5MYLdh, cSW1wSO3 :=  /*line GVHXjWSy.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line Igbw7gIB.go:1*/ shim.Error( /*line jP62Kg0t.go:1*/ cSW1wSO3.Error())
	}
	qVH0siY7 :=  /*line aiichxBC.go:1*/ fmt.Sprintf( /*line yeC_rnuf.go:1*/ func() string {
		var data []byte
		i := 13
		decryptKey := 148
		for counter := 0; i != 28; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 23:
				i = 27
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xee\xe6\xe3\xf5"...,
				)
			case 34:
				data =  /*line yeC_rnuf.go:1*/ append(data, "ϟ"...,
				)
				i = 26
			case 27:
				data =  /*line yeC_rnuf.go:1*/ append(data, 225)
				i = 12
			case 26:
				i = 0
				data =  /*line yeC_rnuf.go:1*/ append(data, "\x96\x99\xcd\xca"...,
				)
			case 7:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xea\xe8\xf5"...,
				)
				i = 30
			case 4:
				i = 9
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xb7\xb0"...,
				)
			case 6:
				i = 24
				data =  /*line yeC_rnuf.go:1*/ append(data, 174)
			case 8:
				for y := range data {
					data[y] = data[y] ^  /*line yeC_rnuf.go:1*/ byte(decryptKey^y)
				}
				i = 28
			case 24:
				i = 16
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xa1\xa8\xe8\xf1"...,
				)
			case 32:
				i = 2
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xc7\xd1"...,
				)
			case 18:
				i = 20
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xdb\xcc\xc2"...,
				)
			case 15:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xc3\xf1\xd6"...,
				)
				i = 32
			case 30:
				i = 10
				data =  /*line yeC_rnuf.go:1*/ append(data, 245)
			case 22:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xc8\xcc\xde\xca"...,
				)
				i = 18
			case 19:
				i = 11
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xf5\xc3"...,
				)
			case 21:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xda\xd4"...,
				)
				i = 29
			case 33:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xf1\xa9\xec\xe6"...,
				)
				i = 19
			case 17:
				i = 1
				data =  /*line yeC_rnuf.go:1*/ append(data, 247)
			case 3:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\x8b\x9c\x9a"...,
				)
				i = 34
			case 1:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xb1\xaa\xb3\xdf"...,
				)
				i = 7
			case 31:
				i = 8
				data =  /*line yeC_rnuf.go:1*/ append(data, 168)
			case 14:
				data =  /*line yeC_rnuf.go:1*/ append(data, "֒"...,
				)
				i = 3
			case 10:
				i = 15
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xe9\xf1\xe3\xc3"...,
				)
			case 12:
				i = 33
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xfd\xae\xb7"...,
				)
			case 0:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xa3\xb5\xb7\xeb"...,
				)
				i = 4
			case 20:
				data =  /*line yeC_rnuf.go:1*/ append(data, 232)
				i = 21
			case 29:
				data =  /*line yeC_rnuf.go:1*/ append(data, 223)
				i = 14
			case 5:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xec\xa5\xf5\xa9"...,
				)
				i = 31
			case 9:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xa5\xb3\xa0"...,
				)
				i = 6
			case 2:
				i = 25
				data =  /*line yeC_rnuf.go:1*/ append(data, "Ӄ"...,
				)
			case 16:
				data =  /*line yeC_rnuf.go:1*/ append(data, 234)
				i = 5
			case 25:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\x82\x8d"...,
				)
				i = 22
			case 11:
				i = 17
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xed\xe5"...,
				)
			case 13:
				data =  /*line yeC_rnuf.go:1*/ append(data, "\xfd\xa5\xf7\xe0"...,
				)
				i = 23
			}
		}
		return  /*line yeC_rnuf.go:1*/ string(data)
	}(), eviEA3sp, ml5MYLdh)
	bxM55Mmo, cSW1wSO3 :=  /*line EkpybXiv.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line jgdM2J2X.go:1*/ shim.Error( /*line CVmcbAIH.go:1*/ cSW1wSO3.Error())
	}
	return  /*line YN_TUz3O.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) m60retx4(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	uiea5UXJ, cSW1wSO3 :=  /*line KtZseIYZ.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line _VhrEaie.go:1*/ shim.Error( /*line Zp8lV68p.go:1*/ cSW1wSO3.Error())
	}

	qVH0siY7 :=  /*line dNONhnqM.go:1*/ fmt.Sprintf( /*line HlzaEIOF.go:1*/ func() string {
		var data []byte
		i := 25
		decryptKey := 222
		for counter := 0; i != 26; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 19:
				data =  /*line HlzaEIOF.go:1*/ append(data, 100)
				i = 18
			case 20:
				data =  /*line HlzaEIOF.go:1*/ append(data, "OX"...,
				)
				i = 4
			case 9:
				data =  /*line HlzaEIOF.go:1*/ append(data, "HxY"...,
				)
				i = 15
			case 4:
				data =  /*line HlzaEIOF.go:1*/ append(data, "HUY"...,
				)
				i = 23
			case 16:
				data =  /*line HlzaEIOF.go:1*/ append(data, "k}iu"...,
				)
				i = 7
			case 3:
				i = 21
				data =  /*line HlzaEIOF.go:1*/ append(data, 120)
			case 8:
				data =  /*line HlzaEIOF.go:1*/ append(data, 101)
				i = 24
			case 21:
				i = 9
				data =  /*line HlzaEIOF.go:1*/ append(data, "TJ"...,
				)
			case 12:
				i = 14
				data =  /*line HlzaEIOF.go:1*/ append(data, "\x0e\x15"...,
				)
			case 1:
				data =  /*line HlzaEIOF.go:1*/ append(data, "Vaa"...,
				)
				i = 6
			case 5:
				data =  /*line HlzaEIOF.go:1*/ append(data, "y#\""...,
				)
				i = 8
			case 2:
				i = 0
				data =  /*line HlzaEIOF.go:1*/ append(data, "PQFR"...,
				)
			case 22:
				data =  /*line HlzaEIOF.go:1*/ append(data, 110)
				i = 16
			case 6:
				data =  /*line HlzaEIOF.go:1*/ append(data, "z|b"...,
				)
				i = 3
			case 15:
				i = 11
				data =  /*line HlzaEIOF.go:1*/ append(data, "NZ"...,
				)
			case 10:
				i = 17
				data =  /*line HlzaEIOF.go:1*/ append(data, "~:#"...,
				)
			case 27:
				i = 26
				for y := range data {
					data[y] = data[y] ^  /*line HlzaEIOF.go:1*/ byte(decryptKey^y)
				}
			case 14:
				i = 28
				data =  /*line HlzaEIOF.go:1*/ append(data, "\x10\x16C"...,
				)
			case 24:
				data =  /*line HlzaEIOF.go:1*/ append(data, "q|H"...,
				)
				i = 19
			case 25:
				i = 29
				data =  /*line HlzaEIOF.go:1*/ append(data, 117)
			case 13:
				data =  /*line HlzaEIOF.go:1*/ append(data, "\x04\v\x06"...,
				)
				i = 2
			case 17:
				data =  /*line HlzaEIOF.go:1*/ append(data, 52)
				i = 1
			case 28:
				data =  /*line HlzaEIOF.go:1*/ append(data, "\x1332w"...,
				)
				i = 27
			case 7:
				i = 5
				data =  /*line HlzaEIOF.go:1*/ append(data, "&?"...,
				)
			case 29:
				i = 22
				data =  /*line HlzaEIOF.go:1*/ append(data, "-\u007fhf"...,
				)
			case 18:
				i = 10
				data =  /*line HlzaEIOF.go:1*/ append(data, 106)
			case 0:
				data =  /*line HlzaEIOF.go:1*/ append(data, "R\x10J"...,
				)
				i = 20
			case 23:
				i = 12
				data =  /*line HlzaEIOF.go:1*/ append(data, "TS\x15"...,
				)
			case 11:
				data =  /*line HlzaEIOF.go:1*/ append(data, 90)
				i = 13
			}
		}
		return  /*line HlzaEIOF.go:1*/ string(data)
	}(), uiea5UXJ)

	bxM55Mmo, cSW1wSO3 :=  /*line VXPMSXHf.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line NSJ_UuvR.go:1*/ shim.Error( /*line bkkcS4tV.go:1*/ cSW1wSO3.Error())
	}

	return  /*line Uw04k3fw.go:1*/ shim.Success(bxM55Mmo)
}

                                                 
func (g4rnrSNn *G1Y_7pz_) lv99L1hf(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line A6r8FFdz.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line N1lRyO3m.go:1*/ shim.Error( /*line yaykdt93.go:1*/ func() string {
			data :=  /*line yaykdt93.go:1*/ []byte("InvD\vX]3A\x93\x17\xb2PNnt \x00]T\\er")
			positions := [...]byte{9, 11, 10, 9, 17, 10, 11, 10, 7, 10, 3, 4, 3, 17, 11, 5, 20, 10, 12, 9, 6, 19, 5, 18, 3, 13}
			for i := 0; i < 26; i += 2 {
				localKey :=  /*line yaykdt93.go:1*/ byte(i) +  /*line yaykdt93.go:1*/ byte(positions[i]^positions[i+1]) + 199
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line yaykdt93.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]
	lWKzd4Pe, cSW1wSO3 :=  /*line KuhtxRq2.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line y4KOofet.go:1*/ shim.Error( /*line CLmmTBMM.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line gLWoizPB.go:1*/ fmt.Sprintf( /*line HE30_lCg.go:1*/ func() string {
		fullData :=  /*line HE30_lCg.go:1*/ []byte("\x94\xca\xfcI\vᙗ?\xcd\b`\x1fߓ9r\x01\r\f\xd4b9\xbf\x102\x03\x99\x18\xd9_\x9c\x19\xb2\x85\x05\x1aǷ\x0e@i\x12\xc0\x8d\xcfJ\x01\U000e6e97\xd9\xcf'\f\xb1\xd4e\xdb\xd1K\x02\x81\x1fo\x9d\xe2T\xd5R\x12\xb8\x1c\xa7n\x12^\x1d\x8b\xa2\xbas\x93\x8b~\x9f\xdf\x04\x1bco\x10\xa4\xe8\ṛ\xc37zBo[\x93a\xf4\xe5\xe0\xe3\xa3K\x84\x0e\xbc^{\xc7\x1b\x9e+\x05sS\xbf\x1c3_\x84\x90\xe0\x97\xa6u\xa7\x97\xdb\xf3\xa5.\xdfgh\x17\xa8\x85[g2n\x96\x15L\x96\x19\xd8\xf4\x1d\xbb\xfecA\xbb\xdbL\x86\x0fX")
		var data []byte
		data =  /*line HE30_lCg.go:1*/ append(data, fullData[15]+fullData[101], fullData[117]+fullData[146], fullData[69]-fullData[21], fullData[50]^fullData[13], fullData[19]^fullData[11], fullData[108]+fullData[34], fullData[79]+fullData[155], fullData[35]+fullData[65], fullData[162]-fullData[164], fullData[72]^fullData[1], fullData[52]+fullData[3], fullData[109]^fullData[29], fullData[71]-fullData[135], fullData[110]^fullData[63], fullData[7]^fullData[137], fullData[84]-fullData[125], fullData[24]^fullData[122], fullData[51]^fullData[98], fullData[80]^fullData[136], fullData[156]-fullData[112], fullData[32]+fullData[152], fullData[148]^fullData[92], fullData[99]^fullData[95], fullData[157]^fullData[8], fullData[133]^fullData[54], fullData[62]+fullData[160], fullData[10]-fullData[131], fullData[43]-fullData[111], fullData[73]-fullData[38], fullData[121]+fullData[75], fullData[165]-fullData[76], fullData[77]-fullData[78], fullData[88]^fullData[147], fullData[153]-fullData[126], fullData[139]^fullData[61], fullData[53]+fullData[93], fullData[96]^fullData[23], fullData[46]^fullData[142], fullData[36]+fullData[42], fullData[104]^fullData[56], fullData[66]-fullData[22], fullData[48]+fullData[149], fullData[143]-fullData[97], fullData[118]^fullData[100], fullData[154]-fullData[49], fullData[5]+fullData[128], fullData[151]+fullData[30], fullData[115]^fullData[47], fullData[166]^fullData[105], fullData[106]-fullData[14], fullData[91]+fullData[159], fullData[60]+fullData[0], fullData[123]+fullData[45], fullData[114]+fullData[85], fullData[33]-fullData[129], fullData[94]^fullData[9], fullData[20]+fullData[86], fullData[74]^fullData[145], fullData[119]-fullData[16], fullData[158]^fullData[6], fullData[141]^fullData[26], fullData[40]-fullData[140], fullData[132]-fullData[25], fullData[81]+fullData[134], fullData[68]+fullData[64], fullData[57]-fullData[102], fullData[127]^fullData[120], fullData[18]+fullData[70], fullData[116]^fullData[4], fullData[59]-fullData[41], fullData[67]-fullData[82], fullData[55]-fullData[150], fullData[103]+fullData[39], fullData[138]+fullData[124], fullData[27]^fullData[2], fullData[113]-fullData[31], fullData[161]^fullData[90], fullData[12]+fullData[89], fullData[17]-fullData[87], fullData[144]^fullData[44], fullData[83]+fullData[130], fullData[37]^fullData[107], fullData[58]+fullData[28], fullData[167]-fullData[163])
		return  /*line HE30_lCg.go:1*/ string(data)
	}(), eviEA3sp, lWKzd4Pe)
	bxM55Mmo, cSW1wSO3 :=  /*line paYzCPcm.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line BlgRSVjT.go:1*/ shim.Error( /*line fWW4Yasz.go:1*/ cSW1wSO3.Error())
	}

	return  /*line AtDAoWBX.go:1*/ shim.Success(bxM55Mmo)
}

                                                           
func (g4rnrSNn *G1Y_7pz_) z8Fbum9p(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	lWKzd4Pe, cSW1wSO3 :=  /*line HkRVUykl.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line MEJ3ydbm.go:1*/ shim.Error( /*line owpxS1cr.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line JgnvIgq8.go:1*/ fmt.Sprintf( /*line CVuUSzzh.go:1*/ func() string {
		seed :=  /*line CVuUSzzh.go:1*/ byte(185)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line CVuUSzzh.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/  /*line CVuUSzzh.go:1*/ fnc(194)(167)(81)(242)(7)(249)(254)(17)(251)(3)(176)(24)(65)(167)(66)(11)(244)(241)(37)(247)(245)(189)(24)(232)(48)(19)(12)(4)(240)(14)(1)(205)(34)(0)(2)(14)(0)(175)(10)(246)(66)(253)(19)(237)(18)(242)(15)(235)(17)(2)(253)(7)(243)(251)(1)(13)(176)(24)(232)(3)(78)(175)(91)(0)
		return  /*line CVuUSzzh.go:1*/ string(data)
	}(), lWKzd4Pe)
	bxM55Mmo, cSW1wSO3 :=  /*line HicnQCDn.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line hiJUzW51.go:1*/ shim.Error( /*line Iwoq2e07.go:1*/ cSW1wSO3.Error())
	}

	return  /*line KNu02ron.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) ldHTDH3Q(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	bHnVZUAf, cSW1wSO3 :=  /*line rJSTDLWI.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line GtsosJYE.go:1*/ shim.Error( /*line z1aIPY9H.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line FmwqAXCL.go:1*/ fmt.Sprintf( /*line CBvMkOnZ.go:1*/ func() string {
		var data []byte
		i := 10
		decryptKey := 166
		for counter := 0; i != 8; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 5:
				i = 9
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\t\xfc\n"...,
				)
			case 6:
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\xf9\xafȱ"...,
				)
				i = 17
			case 20:
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\x06\x15\x17"...,
				)
				i = 21
			case 12:
				i = 13
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\xac\b"...,
				)
			case 0:
				i = 8
				for y := range data {
					data[y] = data[y] +  /*line CBvMkOnZ.go:1*/ byte(decryptKey^y)
				}
			case 19:
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\xce\xf7\v"...,
				)
				i = 11
			case 14:
				data =  /*line CBvMkOnZ.go:1*/ append(data, 35)
				i = 15
			case 1:
				data =  /*line CBvMkOnZ.go:1*/ append(data, ",\xd4\x17"...,
				)
				i = 14
			case 22:
				i = 5
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\xbcǾ\n"...,
				)
			case 15:
				i = 2
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\x10\x02("...,
				)
			case 16:
				i = 22
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\x01\x04\x13\f"...,
				)
			case 18:
				i = 19
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\xcc\xe5"...,
				)
			case 21:
				i = 16
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\xe5\x00"...,
				)
			case 9:
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\xff\xf3\x00"...,
				)
				i = 6
			case 10:
				data =  /*line CBvMkOnZ.go:1*/ append(data, 56)
				i = 7
			case 13:
				i = 0
				data =  /*line CBvMkOnZ.go:1*/ append(data, 9)
			case 7:
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\xe02%"...,
				)
				i = 3
			case 4:
				data =  /*line CBvMkOnZ.go:1*/ append(data, "$(\xd9\xf2"...,
				)
				i = 1
			case 17:
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\xb5\xfc"...,
				)
				i = 12
			case 3:
				data =  /*line CBvMkOnZ.go:1*/ append(data, "%\x1f\x1e0"...,
				)
				i = 4
			case 2:
				data =  /*line CBvMkOnZ.go:1*/ append(data, " \x0e"...,
				)
				i = 18
			case 11:
				i = 20
				data =  /*line CBvMkOnZ.go:1*/ append(data, "\x18\x1d"...,
				)
			}
		}
		return  /*line CBvMkOnZ.go:1*/ string(data)
	}(), bHnVZUAf)
	bxM55Mmo, cSW1wSO3 :=  /*line USANnzQz.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line oEpKNuC6.go:1*/ shim.Error( /*line DBg1xwCS.go:1*/ cSW1wSO3.Error())
	}

	return  /*line UxiK88TV.go:1*/ shim.Success(bxM55Mmo)
}

                                   

func (g4rnrSNn *G1Y_7pz_) x52PQYyS(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	ml5MYLdh, cSW1wSO3 :=  /*line wrec7IaN.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line tqq0cahs.go:1*/ shim.Error( /*line q9jc6zLs.go:1*/ cSW1wSO3.Error())
	}
	                  
	qVH0siY7 :=  /*line AK3EN6zg.go:1*/ fmt.Sprintf( /*line mUCiTCNO.go:1*/ func() string {
		key :=  /*line mUCiTCNO.go:1*/ []byte("\xd0\xe0\xd4oy\"\xe7\xab\xd4\xf6t\xea\xc5\xc64\xf4\x1ad\x10\xf7\xea\xfe\xc0\xc0b\xf2I1\\\xa6\x9f\x1b\n{\f\xb2y\xce\xf7\x92\xe8\xb8\xec5W\xa0\xb7c\xbat#\\Yφ\"\xab\xbfIَQ*\t")
		data :=  /*line mUCiTCNO.go:1*/ []byte("\xab§\n\x15G\x84\u07fb\x84Vо\xe4P\x9by0i\x87\x8f\xdc\xfa\xe20\x978D9\xd5\xeb/X\x1ez\xdd\x12\xab\x93ӋۉF$\x82\x9bA\xcf\aF.7\xae\xebG\x89\x85k\xfc\xfdsWt")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line mUCiTCNO.go:1*/ string(data)
	}(), ml5MYLdh)
	bxM55Mmo, cSW1wSO3 :=  /*line hZJpp2d6.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line C6AwWjNf.go:1*/ shim.Error( /*line fCw1bRVt.go:1*/ cSW1wSO3.Error())
	}

	return  /*line oAlY0Xcy.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) yjAEmnnC(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line a9zEAv0E.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line pm5_nPmC.go:1*/ shim.Error( /*line sok3RlPb.go:1*/ func() string {
			fullData :=  /*line sok3RlPb.go:1*/ []byte("h\xc7\xc2n}jύ\xdf\xdc\xf8 -,;)\t\xc7>\xeaU\xb9ww\x86\x13*\xa3\xead|P,\xbe\x10\xf4jC[~\xa0\xb0n\n\x99\xe4\xccG\x01\xe77ٞ_\xfe\xc1I\x10\xbc\x10\xf8\xb2\x19SN\x8a)1\xc0X\xa2\xd6O-Rv\x01\x94\t\xa3\u007fQ\xe6@ \x99њ\xef2\xa2\x84p^")
			var data []byte
			data =  /*line sok3RlPb.go:1*/ append(data, fullData[47]-fullData[54], fullData[22]-fullData[78], fullData[43]^fullData[30], fullData[7]-fullData[32], fullData[6]^fullData[27], fullData[19]+fullData[80], fullData[14]+fullData[66], fullData[2]+fullData[93], fullData[33]-fullData[4], fullData[36]-fullData[10], fullData[62]-fullData[61], fullData[38]-fullData[82], fullData[42]-fullData[76], fullData[48]^fullData[29], fullData[64]^fullData[84], fullData[28]-fullData[75], fullData[41]+fullData[92], fullData[11]^fullData[3], fullData[21]+fullData[58], fullData[83]^fullData[73], fullData[60]+fullData[5], fullData[1]^fullData[70], fullData[51]+fullData[44], fullData[17]-fullData[85], fullData[34]+fullData[59], fullData[23]-fullData[89], fullData[90]+fullData[71], fullData[35]-fullData[91], fullData[74]^fullData[50], fullData[25]+fullData[31], fullData[86]+fullData[79], fullData[24]^fullData[88], fullData[45]+fullData[65], fullData[77]-fullData[12], fullData[81]-fullData[67], fullData[37]-fullData[8], fullData[18]^fullData[53], fullData[46]-fullData[69], fullData[55]+fullData[40], fullData[72]-fullData[9], fullData[49]+fullData[39], fullData[52]-fullData[26], fullData[56]-fullData[15], fullData[87]-fullData[13], fullData[0]^fullData[16], fullData[68]-fullData[63], fullData[20]+fullData[57])
			return  /*line sok3RlPb.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]
	ml5MYLdh, cSW1wSO3 :=  /*line dWNsCco0.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line H1lB0hNY.go:1*/ shim.Error( /*line QP4mgyYm.go:1*/ cSW1wSO3.Error())
	}
	                  
	qVH0siY7 :=  /*line Zms2bVFU.go:1*/ fmt.Sprintf( /*line BPtmd21p.go:1*/ func() string {
		data :=  /*line BPtmd21p.go:1*/ []byte("\xb1\"\xace2e\x87\xd8ؿ\xcd\x11%%\xc5; 5\"\x01NzgN\x81\xce_uG[\u009e\xe10<o\xd1\xf9\xb5@\xb1cms\x1b\"\x10\"\xf3\xb4e\xf6\x8d\xb9\xcd\tN:\"\x1d\xd3\",\x8e)`\xd1a\xfdet\xd7nP!\xb3\xce\xee\xf1%\x93\"P\xba")
		positions := [...]byte{53, 36, 73, 75, 26, 60, 39, 44, 40, 2, 80, 75, 55, 69, 74, 53, 14, 53, 28, 82, 31, 8, 16, 60, 17, 36, 54, 2, 7, 39, 30, 68, 59, 21, 14, 65, 71, 51, 32, 21, 15, 16, 28, 32, 23, 56, 32, 21, 12, 4, 46, 28, 39, 77, 34, 66, 37, 18, 75, 54, 24, 78, 6, 51, 0, 13, 51, 19, 11, 20, 73, 37, 14, 48, 31, 53, 76, 44, 22, 38, 76, 68, 32, 25, 42, 68, 49, 83, 82, 33, 63, 64, 29, 31, 6, 26, 10, 66, 2, 63, 17, 82, 52, 69, 33, 55, 8, 13, 38, 40, 68, 9}
		for i := 0; i < 112; i += 2 {
			localKey :=  /*line BPtmd21p.go:1*/ byte(i) +  /*line BPtmd21p.go:1*/ byte(positions[i]^positions[i+1]) + 17
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line BPtmd21p.go:1*/ string(data)
	}(), ml5MYLdh, eviEA3sp)
	bxM55Mmo, cSW1wSO3 :=  /*line MSUJmwFd.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line l2qKezAZ.go:1*/ shim.Error( /*line REAz45QZ.go:1*/ cSW1wSO3.Error())
	}

	return  /*line F8eDilDs.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) qBChpnHz(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	lWKzd4Pe, cSW1wSO3 :=  /*line c_fY0MNr.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line uosdAZhj.go:1*/ shim.Error( /*line nZnXLlYC.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line idPALQ3p.go:1*/ fmt.Sprintf( /*line A68ONjrf.go:1*/ func() string {
		var data []byte
		i := 17
		decryptKey := 215
		for counter := 0; i != 7; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 8:
				data =  /*line A68ONjrf.go:1*/ append(data, 212)
				i = 2
			case 30:
				i = 24
				data =  /*line A68ONjrf.go:1*/ append(data, "\xf1\xeb\xe8\xe3"...,
				)
			case 15:
				i = 20
				data =  /*line A68ONjrf.go:1*/ append(data, 157)
			case 1:
				i = 27
				data =  /*line A68ONjrf.go:1*/ append(data, "\x9a\xcd\xcb\xdf"...,
				)
			case 23:
				data =  /*line A68ONjrf.go:1*/ append(data, "p\x89"...,
				)
				i = 16
			case 22:
				i = 18
				data =  /*line A68ONjrf.go:1*/ append(data, "|\xce\xc1\xc9"...,
				)
			case 18:
				i = 8
				data =  /*line A68ONjrf.go:1*/ append(data, "\xc3\xc2"...,
				)
			case 20:
				for y := range data {
					data[y] = data[y] +  /*line A68ONjrf.go:1*/ byte(decryptKey^y)
				}
				i = 7
			case 6:
				i = 30
				data =  /*line A68ONjrf.go:1*/ append(data, "\xcb\xdf"...,
				)
			case 4:
				data =  /*line A68ONjrf.go:1*/ append(data, "x\xbb"...,
				)
				i = 29
			case 9:
				data =  /*line A68ONjrf.go:1*/ append(data, "蘣"...,
				)
				i = 1
			case 27:
				data =  /*line A68ONjrf.go:1*/ append(data, "\xcd\xe0"...,
				)
				i = 28
			case 24:
				data =  /*line A68ONjrf.go:1*/ append(data, "\xe3\xc1"...,
				)
				i = 14
			case 19:
				data =  /*line A68ONjrf.go:1*/ append(data, "\xa7\xb4\xb9"...,
				)
				i = 13
			case 21:
				i = 25
				data =  /*line A68ONjrf.go:1*/ append(data, "\xd4\xd2\xda"...,
				)
			case 28:
				data =  /*line A68ONjrf.go:1*/ append(data, "\xd3\xe3\xcf\xd1"...,
				)
				i = 21
			case 11:
				i = 15
				data =  /*line A68ONjrf.go:1*/ append(data, "@\x9c"...,
				)
			case 16:
				data =  /*line A68ONjrf.go:1*/ append(data, "r\x93"...,
				)
				i = 19
			case 14:
				data =  /*line A68ONjrf.go:1*/ append(data, "\xd4\xd5\xd8\xe7"...,
				)
				i = 9
			case 5:
				i = 23
				data =  /*line A68ONjrf.go:1*/ append(data, "\xbc\xb2"...,
				)
			case 10:
				data =  /*line A68ONjrf.go:1*/ append(data, 208)
				i = 4
			case 2:
				i = 12
				data =  /*line A68ONjrf.go:1*/ append(data, 192)
			case 17:
				i = 22
				data =  /*line A68ONjrf.go:1*/ append(data, 212)
			case 13:
				i = 6
				data =  /*line A68ONjrf.go:1*/ append(data, "\xaa\xb9\xbb|"...,
				)
			case 25:
				i = 0
				data =  /*line A68ONjrf.go:1*/ append(data, "\xce\xca\xcc\xda"...,
				)
			case 0:
				i = 3
				data =  /*line A68ONjrf.go:1*/ append(data, ";T=A"...,
				)
			case 29:
				i = 26
				data =  /*line A68ONjrf.go:1*/ append(data, "Ǭ"...,
				)
			case 26:
				data =  /*line A68ONjrf.go:1*/ append(data, "\x9e\xc4"...,
				)
				i = 5
			case 12:
				data =  /*line A68ONjrf.go:1*/ append(data, "\xc4u\x8e"...,
				)
				i = 10
			case 3:
				i = 11
				data =  /*line A68ONjrf.go:1*/ append(data, 144)
			}
		}
		return  /*line A68ONjrf.go:1*/ string(data)
	}(), lWKzd4Pe)
	bxM55Mmo, cSW1wSO3 :=  /*line P2Yr8_Vq.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line NDs3zeE4.go:1*/ shim.Error( /*line Mrrzq0ov.go:1*/ cSW1wSO3.Error())
	}

	return  /*line Wnyf344x.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) gA2iac8z(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line DvbocMBb.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line eYtbCQFG.go:1*/ shim.Error( /*line AYhwfs2q.go:1*/ func() string {
			seed :=  /*line AYhwfs2q.go:1*/ byte(124)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line AYhwfs2q.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/  /*line AYhwfs2q.go:1*/ fnc(205)(37)(8)(235)(11)(253)(251)(188)(33)(49)(245)(14)(248)(248)(9)(6)(172)(46)(39)(248)(245)(3)(13)(188)(242)(37)(51)(248)(245)(254)(17)(245)(5)(249)(185)(68)(253)(19)(237)(18)(242)(15)(172)(78)(243)(12)(248)
			return  /*line AYhwfs2q.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]
	lWKzd4Pe, cSW1wSO3 :=  /*line bWNHdO9f.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line Dmfh7X_R.go:1*/ shim.Error( /*line kc_YdAZj.go:1*/ cSW1wSO3.Error())
	}
	                  
	qVH0siY7 :=  /*line OIAH7evu.go:1*/ fmt.Sprintf( /*line UPry7g5E.go:1*/ func() string {
		var data []byte
		i := 25
		decryptKey := 237
		for counter := 0; i != 7; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 22:
				data =  /*line UPry7g5E.go:1*/ append(data, "\r&\v"...,
				)
				i = 35
			case 23:
				i = 19
				data =  /*line UPry7g5E.go:1*/ append(data, 244)
			case 12:
				i = 18
				data =  /*line UPry7g5E.go:1*/ append(data, "[YiW"...,
				)
			case 33:
				i = 27
				data =  /*line UPry7g5E.go:1*/ append(data, 1)
			case 10:
				i = 11
				data =  /*line UPry7g5E.go:1*/ append(data, "\xfe\xf1\xfd\xe9"...,
				)
			case 37:
				data =  /*line UPry7g5E.go:1*/ append(data, 58)
				i = 0
			case 8:
				data =  /*line UPry7g5E.go:1*/ append(data, "\xf0.\xd6"...,
				)
				i = 24
			case 26:
				i = 31
				data =  /*line UPry7g5E.go:1*/ append(data, "\x1c."...,
				)
			case 6:
				i = 1
				data =  /*line UPry7g5E.go:1*/ append(data, "\xfa\xfb"...,
				)
			case 0:
				data =  /*line UPry7g5E.go:1*/ append(data, "\x1f#n\x1e"...,
				)
				i = 30
			case 16:
				data =  /*line UPry7g5E.go:1*/ append(data, "\x04&"...,
				)
				i = 38
			case 3:
				i = 26
				data =  /*line UPry7g5E.go:1*/ append(data, "#'!"...,
				)
			case 1:
				i = 14
				data =  /*line UPry7g5E.go:1*/ append(data, "\xfa\t\x06"...,
				)
			case 25:
				i = 3
				data =  /*line UPry7g5E.go:1*/ append(data, ":\xe20"...,
				)
			case 38:
				data =  /*line UPry7g5E.go:1*/ append(data, "\x1e\x10"...,
				)
				i = 17
			case 20:
				data =  /*line UPry7g5E.go:1*/ append(data, "\x13\r\x06"...,
				)
				i = 33
			case 31:
				data =  /*line UPry7g5E.go:1*/ append(data, "&*\xd7"...,
				)
				i = 8
			case 35:
				i = 28
				data =  /*line UPry7g5E.go:1*/ append(data, "\x0fZ\n"...,
				)
			case 39:
				data =  /*line UPry7g5E.go:1*/ append(data, "\xd6\xf1\x05"...,
				)
				i = 20
			case 11:
				i = 15
				data =  /*line UPry7g5E.go:1*/ append(data, "\xf7\xfa"...,
				)
			case 21:
				i = 10
				data =  /*line UPry7g5E.go:1*/ append(data, "\x01\xef"...,
				)
			case 28:
				data =  /*line UPry7g5E.go:1*/ append(data, "bc"...,
				)
				i = 13
			case 15:
				i = 34
				data =  /*line UPry7g5E.go:1*/ append(data, 244)
			case 9:
				i = 5
				data =  /*line UPry7g5E.go:1*/ append(data, 8)
			case 29:
				i = 22
				data =  /*line UPry7g5E.go:1*/ append(data, "QZS"...,
				)
			case 13:
				i = 7
				for y := range data {
					data[y] = data[y] +  /*line UPry7g5E.go:1*/ byte(decryptKey^y)
				}
			case 36:
				data =  /*line UPry7g5E.go:1*/ append(data, "\xf9\r\x16\x1b"...,
				)
				i = 9
			case 5:
				i = 39
				data =  /*line UPry7g5E.go:1*/ append(data, "\x17\x15"...,
				)
			case 4:
				i = 29
				data =  /*line UPry7g5E.go:1*/ append(data, "YeQ]"...,
				)
			case 34:
				i = 23
				data =  /*line UPry7g5E.go:1*/ append(data, "\xfc\xec\xe8\xe6"...,
				)
			case 17:
				data =  /*line UPry7g5E.go:1*/ append(data, 206)
				i = 2
			case 19:
				data =  /*line UPry7g5E.go:1*/ append(data, 33)
				i = 37
			case 24:
				data =  /*line UPry7g5E.go:1*/ append(data, "\x15!\x12"...,
				)
				i = 16
			case 27:
				i = 6
				data =  /*line UPry7g5E.go:1*/ append(data, "\xfd\xdb"...,
				)
			case 2:
				i = 36
				data =  /*line UPry7g5E.go:1*/ append(data, "\xe3\xcc"...,
				)
			case 14:
				data =  /*line UPry7g5E.go:1*/ append(data, "\xb6\xbd"...,
				)
				i = 32
			case 30:
				data =  /*line UPry7g5E.go:1*/ append(data, "%\x1c"...,
				)
				i = 12
			case 32:
				data =  /*line UPry7g5E.go:1*/ append(data, "\xb4\xf3\xf1"...,
				)
				i = 21
			case 18:
				i = 4
				data =  /*line UPry7g5E.go:1*/ append(data, 102)
			}
		}
		return  /*line UPry7g5E.go:1*/ string(data)
	}(), lWKzd4Pe, eviEA3sp)
	bxM55Mmo, cSW1wSO3 :=  /*line o1SY7DLe.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line lyuCyjNw.go:1*/ shim.Error( /*line grUmd0A6.go:1*/ cSW1wSO3.Error())
	}

	return  /*line z_KBvjno.go:1*/ shim.Success(bxM55Mmo)
}
