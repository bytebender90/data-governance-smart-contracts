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

func (hslgON00 *EMhVAN9S) vH1NU2MF(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	var xREbPiGc error
	var w1e7h2B3 string
	if  /*line Bvyk2n12.go:1*/ len(eFz_TaL1) != 14 {
		return  /*line dY_ReMeA.go:1*/ shim.Error( /*line xWGNeEmP.go:1*/ func() string {
			var data []byte
			i := 17
			decryptKey := 104
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 29:
					data =  /*line xWGNeEmP.go:1*/ append(data, "_\xa3\xa9"...,
					)
					i = 24
				case 16:
					data =  /*line xWGNeEmP.go:1*/ append(data, "\x8b\xac"...,
					)
					i = 0
				case 11:
					data =  /*line xWGNeEmP.go:1*/ append(data, "\x834\x84\x85"...,
					)
					i = 26
				case 4:
					i = 16
					data =  /*line xWGNeEmP.go:1*/ append(data, "\x8a\x95\x9c"...,
					)
				case 21:
					data =  /*line xWGNeEmP.go:1*/ append(data, "\x9f\x99\x9f"...,
					)
					i = 28
				case 18:
					i = 20
					data =  /*line xWGNeEmP.go:1*/ append(data, "\xa5\x9d"...,
					)
				case 12:
					i = 6
					data =  /*line xWGNeEmP.go:1*/ append(data, "\u007f~t"...,
					)
				case 6:
					data =  /*line xWGNeEmP.go:1*/ append(data, "qu d"...,
					)
					i = 25
				case 28:
					data =  /*line xWGNeEmP.go:1*/ append(data, "m\x92"...,
					)
					i = 5
				case 5:
					data =  /*line xWGNeEmP.go:1*/ append(data, "\x91\x86\x93\x96"...,
					)
					i = 1
				case 27:
					data =  /*line xWGNeEmP.go:1*/ append(data, "y\u007fj"...,
					)
					i = 7
				case 23:
					i = 18
					data =  /*line xWGNeEmP.go:1*/ append(data, "\x98\x95\xa9\x9d"...,
					)
				case 10:
					i = 12
					data =  /*line xWGNeEmP.go:1*/ append(data, 121)
				case 25:
					data =  /*line xWGNeEmP.go:1*/ append(data, "tl"...,
					)
					i = 15
				case 26:
					data =  /*line xWGNeEmP.go:1*/ append(data, "\x97\x91"...,
					)
					i = 21
				case 7:
					data =  /*line xWGNeEmP.go:1*/ append(data, "í"...,
					)
					i = 14
				case 13:
					data =  /*line xWGNeEmP.go:1*/ append(data, "\xa6\xb3"...,
					)
					i = 8
				case 15:
					i = 19
					data =  /*line xWGNeEmP.go:1*/ append(data, "ytk\x87"...,
					)
				case 17:
					i = 10
					data =  /*line xWGNeEmP.go:1*/ append(data, "Rvn"...,
					)
				case 8:
					for y := range data {
						data[y] = data[y] -  /*line xWGNeEmP.go:1*/ byte(decryptKey^y)
					}
					i = 2
				case 19:
					data =  /*line xWGNeEmP.go:1*/ append(data, "\x8c\x8eH"...,
					)
					i = 3
				case 1:
					i = 4
					data =  /*line xWGNeEmP.go:1*/ append(data, 116)
				case 20:
					i = 27
					data =  /*line xWGNeEmP.go:1*/ append(data, 105)
				case 9:
					data =  /*line xWGNeEmP.go:1*/ append(data, "t\x84|\x81"...,
					)
					i = 11
				case 0:
					i = 22
					data =  /*line xWGNeEmP.go:1*/ append(data, 172)
				case 24:
					i = 23
					data =  /*line xWGNeEmP.go:1*/ append(data, 160)
				case 22:
					i = 29
					data =  /*line xWGNeEmP.go:1*/ append(data, "\x8a\xac\xa4\xaf"...,
					)
				case 14:
					data =  /*line xWGNeEmP.go:1*/ append(data, "\xbb\xc3"...,
					)
					i = 13
				case 30:
					i = 9
					data =  /*line xWGNeEmP.go:1*/ append(data, "b\x94\x8c"...,
					)
				case 3:
					data =  /*line xWGNeEmP.go:1*/ append(data, 61)
					i = 30
				}
			}
			return  /*line xWGNeEmP.go:1*/ string(data)
		}())
	}
	cXmPo8Kd := eFz_TaL1[0]
	dwbLZnck := eFz_TaL1[1]
	eSpmNp6n := eFz_TaL1[2]
	hHY_irCR := eFz_TaL1[3]
	fvUjTcZo := eFz_TaL1[4]
	xrSSF6uG := eFz_TaL1[5]
	_Y4gkQi4 := eFz_TaL1[6]
	gLRzXpRk := eFz_TaL1[7]
	aPvOa68E := eFz_TaL1[8]
	uOH_DeF4 := eFz_TaL1[9]
	gBYxQKrs := eFz_TaL1[10]
	g3ka7rbx :=  /*line FYMt23Iq.go:1*/ strings.Split(eFz_TaL1[11],  /*line xaulX4tW.go:1*/ func() string {
		key :=  /*line xaulX4tW.go:1*/ []byte("\xb3")
		data :=  /*line xaulX4tW.go:1*/ []byte("\x9f")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line xaulX4tW.go:1*/ string(data)
	}())
	p5F1lbum := eFz_TaL1[12]
	clszPHCk := eFz_TaL1[13]
	_0uGwnuQ :=  /*line ai2MAekh.go:1*/ func() string {
		key :=  /*line ai2MAekh.go:1*/ []byte("^w\x92\xad\xfc\x1a$\xd2@\x89\xf3\x15wz\x011")
		data :=  /*line ai2MAekh.go:1*/ []byte("\xaa\xe6\xf9\x14a\x8ce5\xa3\xeef\x88\xc6\xech\xa4")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line ai2MAekh.go:1*/ string(data)
	}()
	cKWTKAnj :=  /*line pIaArFUR.go:1*/ func() string {
		var data []byte
		i := 3
		decryptKey := 210
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				i = 1
				data =  /*line pIaArFUR.go:1*/ append(data, "\x97\xb8\xa7"...,
				)
			case 0:
				data =  /*line pIaArFUR.go:1*/ append(data, "\xb4\x8f\xb3"...,
				)
				i = 6
			case 1:
				for y := range data {
					data[y] = data[y] ^  /*line pIaArFUR.go:1*/ byte(decryptKey^y)
				}
				i = 5
			case 3:
				i = 4
				data =  /*line pIaArFUR.go:1*/ append(data, "\x80\xa2\xa9"...,
				)
			case 4:
				i = 7
				data =  /*line pIaArFUR.go:1*/ append(data, "\xa8\xad\xbb\x8b"...,
				)
			case 7:
				data =  /*line pIaArFUR.go:1*/ append(data, "\xa8\xa7\xa0\xb5"...,
				)
				i = 0
			case 6:
				i = 2
				data =  /*line pIaArFUR.go:1*/ append(data, "\xa5\xb0"...,
				)
			}
		}
		return  /*line pIaArFUR.go:1*/ string(data)
	}()

	if  /*line irzO2pxa.go:1*/ len(cXmPo8Kd) <= 0 {
		return  /*line Nw4yZToy.go:1*/ shim.Error( /*line hs9uhn3H.go:1*/ func() string {
			key :=  /*line hs9uhn3H.go:1*/ []byte("Y\xd5\f\xe4O\xe9\a0\x9f\x1dk\xe5\xdf3\xf4\x9e\x1f\xb0\x04Y\xa5q](\x00\xb6\xeb0\x9a4]jT\x96")
			data :=  /*line hs9uhn3H.go:1*/ []byte("\xeb\x8ch}$|m\xf0\xceX\b\x8fA/q\x82Bpj\x16ɼ\bEp\xbe\x8e\xf0\xd9@\x15\xff\x1a\xd1")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line hs9uhn3H.go:1*/ string(data)
		}())
	}

	if  /*line z0dT4L7A.go:1*/ len(dwbLZnck) <= 0 {
		return  /*line d7CEAKKd.go:1*/ shim.Error( /*line UA0DrhNq.go:1*/ func() string {
			seed :=  /*line UA0DrhNq.go:1*/ byte(3)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line UA0DrhNq.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/  /*line UA0DrhNq.go:1*/ fnc(88)(206)(142)(41)(78)(143)(42)(76)(83)(245)(225)(124)(72)(146)(33)(73)(133)(5)(11)(35)(244)(53)(114)(226)(197)(54)(174)(95)(121)(51)(37)(152)(49)(97)(129)(58)(124)(251)(250)(249)(153)(133)(11)(20)(31)(67)(127)
			return  /*line UA0DrhNq.go:1*/ string(data)
		}())
	}

	if  /*line mz44eWWV.go:1*/ len(eSpmNp6n) <= 0 {
		return  /*line c86Ude8P.go:1*/ shim.Error( /*line cgoSOGzx.go:1*/ func() string {
			fullData :=  /*line cgoSOGzx.go:1*/ []byte("7\x93V$\br\x01\xe1\x8e]c\xbd\xd4\xf0\xe43[¯\xff\xc9z\xcaP)\x9a\xbf\x89\xdd#\xd4\xd3G\x92O\x8d쳆(Ҿ$\xe13\xaeGW?\x06D|\xdf\xca\xf7\x9e\xacT\xc7\x00LG:9\x14\x1b\xab\xa2\xf4\xfd\xfd\x1f\xae0)@Ua")
			var data []byte
			data =  /*line cgoSOGzx.go:1*/ append(data, fullData[70]-fullData[45], fullData[29]+fullData[34], fullData[73]+fullData[0], fullData[58]+fullData[25], fullData[68]+fullData[21], fullData[11]^fullData[12], fullData[35]^fullData[54], fullData[53]^fullData[66], fullData[67]+fullData[40], fullData[77]^fullData[4], fullData[16]+fullData[64], fullData[46]^fullData[24], fullData[72]^fullData[8], fullData[17]-fullData[76], fullData[39]^fullData[9], fullData[22]-fullData[47], fullData[61]-fullData[31], fullData[50]-fullData[42], fullData[65]+fullData[32], fullData[10]^fullData[49], fullData[30]+fullData[60], fullData[19]^fullData[55], fullData[71]^fullData[48], fullData[57]^fullData[62], fullData[2]^fullData[63], fullData[1]^fullData[69], fullData[20]^fullData[14], fullData[38]+fullData[52], fullData[41]+fullData[18], fullData[36]-fullData[51], fullData[3]^fullData[23], fullData[56]-fullData[15], fullData[6]-fullData[7], fullData[43]+fullData[33], fullData[44]-fullData[26], fullData[59]+fullData[5], fullData[75]^fullData[74], fullData[28]^fullData[37], fullData[13]-fullData[27])
			return  /*line cgoSOGzx.go:1*/ string(data)
		}())
	}
	if  /*line waKf26Q0.go:1*/ len(fvUjTcZo) <= 0 {
		return  /*line B3f3BN1Y.go:1*/ shim.Error( /*line jbGuAR5r.go:1*/ func() string {
			data :=  /*line jbGuAR5r.go:1*/ []byte("S\xdaa\xe8uI\x01\xf9\xdfN\xe8 \b\x98\x18\xfb~\xe1\u007f\xfe\xe0em\xdd\xf1t\xb2;tr\xfcn$")
			positions := [...]byte{7, 27, 32, 8, 14, 7, 16, 26, 8, 6, 26, 5, 12, 23, 18, 20, 23, 19, 19, 25, 30, 24, 10, 15, 19, 3, 18, 20, 9, 17, 20, 13, 31, 31, 7, 26, 31, 31, 1, 17}
			for i := 0; i < 40; i += 2 {
				localKey :=  /*line jbGuAR5r.go:1*/ byte(i) +  /*line jbGuAR5r.go:1*/ byte(positions[i]^positions[i+1]) + 94
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line jbGuAR5r.go:1*/ string(data)
		}())
	}

	tqOONhpi :=  /*line dbuUGYJF.go:1*/ func() string {
		fullData :=  /*line dbuUGYJF.go:1*/ []byte("\v5ѐ\x10%C\x98cL$\xcd\x19\x93\xd9J,@\x9aDHzF\x19")
		var data []byte
		data =  /*line dbuUGYJF.go:1*/ append(data, fullData[2]+fullData[3], fullData[23]^fullData[21], fullData[15]+fullData[12], fullData[5]+fullData[17], fullData[4]^fullData[8], fullData[14]+fullData[18], fullData[19]+fullData[0], fullData[22]+fullData[16], fullData[6]^fullData[10], fullData[13]-fullData[20], fullData[11]+fullData[7], fullData[1]^fullData[9])
		return  /*line dbuUGYJF.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line xJMvDTms.go:1*/ func() string {
		seed :=  /*line xJMvDTms.go:1*/ byte(29)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line xJMvDTms.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line xJMvDTms.go:1*/ fnc(50)
		return  /*line xJMvDTms.go:1*/ string(data)
	}() + dwbLZnck +  /*line IbAMFXza.go:1*/ func() string {
		seed :=  /*line IbAMFXza.go:1*/ byte(45)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line IbAMFXza.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line IbAMFXza.go:1*/ fnc(2)
		return  /*line IbAMFXza.go:1*/ string(data)
	}() + eSpmNp6n
	Rzr3OlkK := LoggerAccessOrgsCounter{}
	sdHID17m := -1
	t40HkFnP, xREbPiGc :=  /*line Bw8sxuus.go:1*/ fOOTaXmK.GetState(tqOONhpi)	                                     
	if xREbPiGc != nil {
		w1e7h2B3 =  /*line K9PsTKLp.go:1*/ func() string {
			var data []byte
			i := 5
			decryptKey := 92
			for counter := 0; i != 3; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 22:
					data =  /*line K9PsTKLp.go:1*/ append(data, 86)
					i = 23
				case 21:
					i = 2
					data =  /*line K9PsTKLp.go:1*/ append(data, 123)
				case 15:
					data =  /*line K9PsTKLp.go:1*/ append(data, "\x19\x0e"...,
					)
					i = 9
				case 1:
					data =  /*line K9PsTKLp.go:1*/ append(data, 74)
					i = 22
				case 5:
					i = 21
					data =  /*line K9PsTKLp.go:1*/ append(data, "G\x1f"...,
					)
				case 23:
					data =  /*line K9PsTKLp.go:1*/ append(data, 72)
					i = 15
				case 6:
					data =  /*line K9PsTKLp.go:1*/ append(data, "]WW"...,
					)
					i = 19
				case 8:
					data =  /*line K9PsTKLp.go:1*/ append(data, 54)
					i = 11
				case 11:
					for y := range data {
						data[y] = data[y] ^  /*line K9PsTKLp.go:1*/ byte(decryptKey^y)
					}
					i = 3
				case 13:
					data =  /*line K9PsTKLp.go:1*/ append(data, "EB"...,
					)
					i = 17
				case 10:
					i = 18
					data =  /*line K9PsTKLp.go:1*/ append(data, 15)
				case 4:
					data =  /*line K9PsTKLp.go:1*/ append(data, 65)
					i = 10
				case 17:
					i = 20
					data =  /*line K9PsTKLp.go:1*/ append(data, "SR\x02"...,
					)
				case 16:
					data =  /*line K9PsTKLp.go:1*/ append(data, "^\vEF"...,
					)
					i = 13
				case 0:
					i = 7
					data =  /*line K9PsTKLp.go:1*/ append(data, 96)
				case 14:
					i = 12
					data =  /*line K9PsTKLp.go:1*/ append(data, 117)
				case 9:
					i = 6
					data =  /*line K9PsTKLp.go:1*/ append(data, "\x17pVY"...,
					)
				case 20:
					data =  /*line K9PsTKLp.go:1*/ append(data, 75)
					i = 14
				case 18:
					data =  /*line K9PsTKLp.go:1*/ append(data, "OL"...,
					)
					i = 16
				case 7:
					data =  /*line K9PsTKLp.go:1*/ append(data, ":}{g"...,
					)
					i = 8
				case 19:
					data =  /*line K9PsTKLp.go:1*/ append(data, "\fY"...,
					)
					i = 4
				case 12:
					data =  /*line K9PsTKLp.go:1*/ append(data, "njpj"...,
					)
					i = 0
				case 2:
					data =  /*line K9PsTKLp.go:1*/ append(data, 77)
					i = 1
				}
			}
			return  /*line K9PsTKLp.go:1*/ string(data)
		}() + tqOONhpi +  /*line iNUeFgjk.go:1*/ func() string {
			fullData :=  /*line iNUeFgjk.go:1*/ []byte(":H\xb7\xda")
			var data []byte
			data =  /*line iNUeFgjk.go:1*/ append(data, fullData[1]+fullData[3], fullData[2]-fullData[0])
			return  /*line iNUeFgjk.go:1*/ string(data)
		}()
		return  /*line d3NwBUzo.go:1*/ shim.Error(w1e7h2B3)
	} else if t40HkFnP == nil {
		sdHID17m = 0
		m3ZWz2be := &LoggerRevokedAccessCounterOrg{cKWTKAnj, cXmPo8Kd, sdHID17m}
		t40HkFnP, xREbPiGc =  /*line ER7zi4sh.go:1*/ json.Marshal(m3ZWz2be)
		if xREbPiGc != nil {
			return  /*line y8PJMu0G.go:1*/ shim.Error( /*line vLOkqIDO.go:1*/ xREbPiGc.Error())
		}

	} else {
		xREbPiGc =  /*line uDwbZu2J.go:1*/ json.Unmarshal(t40HkFnP, &Rzr3OlkK)	                               
		if xREbPiGc != nil {
			return  /*line O52nYbQs.go:1*/ shim.Error( /*line Lxhej0O2.go:1*/ xREbPiGc.Error())
		}
		sdHID17m = Rzr3OlkK.Count + 1
		Rzr3OlkK.Count = sdHID17m
		t40HkFnP, _ =  /*line Dz7uDA91.go:1*/ json.Marshal(Rzr3OlkK)
	}

	                                               
	xREbPiGc =  /*line whU_mt8_.go:1*/ fOOTaXmK.PutState(tqOONhpi, t40HkFnP)
	if xREbPiGc != nil {
		return  /*line JZVoiCg6.go:1*/ shim.Error( /*line FlTUyJLW.go:1*/ xREbPiGc.Error())
	}

	BogxjCz3 := &LoggerAccessOrgs{_0uGwnuQ, cXmPo8Kd, dwbLZnck, p5F1lbum, clszPHCk, eSpmNp6n, xrSSF6uG, gBYxQKrs, g3ka7rbx, fvUjTcZo, hHY_irCR, sdHID17m,
		_Y4gkQi4, gLRzXpRk, aPvOa68E, uOH_DeF4}
	EugKyXCJ, xREbPiGc :=  /*line abYz0Haz.go:1*/ json.Marshal(BogxjCz3)
	if xREbPiGc != nil {
		return  /*line evSRj9tF.go:1*/ shim.Error( /*line You4pIWk.go:1*/ xREbPiGc.Error())
	}

	iPrzUCH_ :=  /*line qQQHqNYh.go:1*/ func() string {
		seed :=  /*line qQQHqNYh.go:1*/ byte(230)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line qQQHqNYh.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line qQQHqNYh.go:1*/  /*line qQQHqNYh.go:1*/  /*line qQQHqNYh.go:1*/  /*line qQQHqNYh.go:1*/  /*line qQQHqNYh.go:1*/  /*line qQQHqNYh.go:1*/  /*line qQQHqNYh.go:1*/  /*line qQQHqNYh.go:1*/  /*line qQQHqNYh.go:1*/ fnc(123)(2)(0)(2)(14)(0)(220)(35)(245)
		return  /*line qQQHqNYh.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line JguCb5zL.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 152
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				data =  /*line JguCb5zL.go:1*/ append(data, 201)
				i = 2
			case 2:
				i = 0
				for y := range data {
					data[y] = data[y] -  /*line JguCb5zL.go:1*/ byte(decryptKey^y)
				}
			}
		}
		return  /*line JguCb5zL.go:1*/ string(data)
	}() + dwbLZnck + eSpmNp6n +  /*line Cys_hHAb.go:1*/ func() string {
		key :=  /*line Cys_hHAb.go:1*/ []byte("\x89")
		data :=  /*line Cys_hHAb.go:1*/ []byte("\xb8")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line Cys_hHAb.go:1*/ string(data)
	}() +  /*line sUF5mDPi.go:1*/ strconv.Itoa(sdHID17m)
	                        
	xREbPiGc =  /*line gCvZqb78.go:1*/ fOOTaXmK.PutState(iPrzUCH_, EugKyXCJ)
	if xREbPiGc != nil {
		return  /*line Ezg0HTQF.go:1*/ shim.Error( /*line XD3gqTWp.go:1*/ xREbPiGc.Error())
	}

	auAEXfA1 :=  /*line BFszcBN2.go:1*/ func() string {
		var data []byte
		i := 7
		decryptKey := 48
		for counter := 0; i != 3; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 16:
				i = 13
				data =  /*line BFszcBN2.go:1*/ append(data, "\xd1\xd2"...,
				)
			case 4:
				data =  /*line BFszcBN2.go:1*/ append(data, "\x86\xc9\xc5\xd7"...,
				)
				i = 14
			case 0:
				data =  /*line BFszcBN2.go:1*/ append(data, 227)
				i = 12
			case 12:
				data =  /*line BFszcBN2.go:1*/ append(data, 231)
				i = 6
			case 14:
				data =  /*line BFszcBN2.go:1*/ append(data, "\xc3\xd4"...,
				)
				i = 9
			case 5:
				for y := range data {
					data[y] = data[y] -  /*line BFszcBN2.go:1*/ byte(decryptKey^y)
				}
				i = 3
			case 2:
				data =  /*line BFszcBN2.go:1*/ append(data, 233)
				i = 1
			case 7:
				i = 8
				data =  /*line BFszcBN2.go:1*/ append(data, "\xc2\xea"...,
				)
			case 8:
				data =  /*line BFszcBN2.go:1*/ append(data, "\xe2\xdd\xed"...,
				)
				i = 0
			case 15:
				data =  /*line BFszcBN2.go:1*/ append(data, "Ә}"...,
				)
				i = 5
			case 13:
				data =  /*line BFszcBN2.go:1*/ append(data, "\xd1\xd2"...,
				)
				i = 11
			case 6:
				data =  /*line BFszcBN2.go:1*/ append(data, "ߗ\xbe\xe4"...,
				)
				i = 2
			case 1:
				i = 16
				data =  /*line BFszcBN2.go:1*/ append(data, "\xe6ב"...,
				)
			case 9:
				i = 15
				data =  /*line BFszcBN2.go:1*/ append(data, 197)
			case 11:
				data =  /*line BFszcBN2.go:1*/ append(data, "\xdfފ\xcf"...,
				)
				i = 10
			case 10:
				i = 4
				data =  /*line BFszcBN2.go:1*/ append(data, "\xd7\xd9"...,
				)
			}
		}
		return  /*line BFszcBN2.go:1*/ string(data)
	}() + cXmPo8Kd
	pPXX7o_t :=  /*line Qw5Aruuj.go:1*/ []byte(auAEXfA1)
	return  /*line TdbJYHj3.go:1*/ shim.Success(pPXX7o_t)

}

func (hslgON00 *EMhVAN9S) dQMbWmvo(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	var w1e7h2B3 string
	var xREbPiGc error
	if  /*line brWEsRfS.go:1*/ len(eFz_TaL1) != 12 {
		return  /*line E6zmQ2El.go:1*/ shim.Error( /*line B5uEkpgG.go:1*/ func() string {
			key :=  /*line B5uEkpgG.go:1*/ []byte("1\x8b\x85\xe5\x91\xdd\x14\x91\x1d\x1b\xad\xa44\x1eaz\xf9\xa4\xdb\\\xafROp\xf9\xc0\xa0\xf2\x1e\xa9\xc4~\xbcA\x94\x8b\x82\x0e\xfb\x95\xbf{\xf0l\xfc\u007f\x97\x0e\x06o\x03%\x15%\x0fk\xbb\x1eY\xad\x1c\x05\xc1\xe4㶝\xe3\xad\xe7£j")
			data :=  /*line B5uEkpgG.go:1*/ []byte("\x18\xe3ފ\xe1\x95Q\xd2W\x05\xb4\xce3W\f\xebuИ\xd2q\xf4&\xfej\xb4\xc9}Pw\xa9\xf1\xb2(\xe0\xe4\xf0Dj\xe1\xb0\xf0u\xf8E\xe4\xccWm\x04LMR\xfbV\r\xb5G\n\xc7Mi\xa6<N|\x83\x93\xb4\x85\xb3\xc2\t")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line B5uEkpgG.go:1*/ string(data)
		}())
	}
	cXmPo8Kd := eFz_TaL1[0]
	dwbLZnck := eFz_TaL1[1]
	eSpmNp6n := eFz_TaL1[2]
	beT6BA9K := eFz_TaL1[3]
	xrSSF6uG := eFz_TaL1[4]
	_Y4gkQi4 := eFz_TaL1[5]
	gLRzXpRk := eFz_TaL1[6]
	aPvOa68E := eFz_TaL1[7]
	uOH_DeF4 := eFz_TaL1[8]
	g3ka7rbx :=  /*line BoTQnFob.go:1*/ strings.Split(eFz_TaL1[9],  /*line VG7aJIkr.go:1*/ func() string {
		key :=  /*line VG7aJIkr.go:1*/ []byte("\xc5")
		data :=  /*line VG7aJIkr.go:1*/ []byte("\xe9")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line VG7aJIkr.go:1*/ string(data)
	}())
	p5F1lbum := eFz_TaL1[10]
	clszPHCk := eFz_TaL1[11]
	_0uGwnuQ :=  /*line AYJey2eg.go:1*/ func() string {
		data :=  /*line AYJey2eg.go:1*/ []byte("X\fgzQ\rSdjoJ\\DVsW\\srKrms")
		positions := [...]byte{21, 13, 8, 6, 16, 5, 13, 7, 11, 21, 19, 14, 10, 18, 3, 6, 11, 1, 15, 4, 10, 12, 16, 11, 14, 0}
		for i := 0; i < 26; i += 2 {
			localKey :=  /*line AYJey2eg.go:1*/ byte(i) +  /*line AYJey2eg.go:1*/ byte(positions[i]^positions[i+1]) + 21
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line AYJey2eg.go:1*/ string(data)
	}()
	cKWTKAnj :=  /*line SShjmKtL.go:1*/ func() string {
		seed :=  /*line SShjmKtL.go:1*/ byte(143)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line SShjmKtL.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/  /*line SShjmKtL.go:1*/ fnc(219)(217)(170)(84)(166)(89)(146)(55)(127)(247)(234)(206)(155)(19)(72)(144)(34)(82)(164)(32)(90)(200)(102)(239)(211)(178)
		return  /*line SShjmKtL.go:1*/ string(data)
	}()

	if  /*line ETDQvVDP.go:1*/ len(cXmPo8Kd) <= 0 {
		return  /*line i9hzMzv0.go:1*/ shim.Error( /*line Hdhrmvf9.go:1*/ func() string {
			data :=  /*line Hdhrmvf9.go:1*/ []byte(".aV\x1aKv~\x16\x0euYZ\\cU ad\x06on\x02eE\\tFL~Zrin\x1e")
			positions := [...]byte{33, 0, 14, 4, 14, 26, 6, 2, 18, 33, 28, 17, 7, 27, 23, 5, 17, 24, 6, 0, 17, 27, 7, 13, 11, 29, 23, 14, 10, 4, 6, 12, 33, 18, 3, 27, 12, 7, 26, 13, 12, 21, 12, 8}
			for i := 0; i < 44; i += 2 {
				localKey :=  /*line Hdhrmvf9.go:1*/ byte(i) +  /*line Hdhrmvf9.go:1*/ byte(positions[i]^positions[i+1]) + 0
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line Hdhrmvf9.go:1*/ string(data)
		}())
	}

	if  /*line CmrNumUR.go:1*/ len(eSpmNp6n) <= 0 {
		return  /*line GjTn16Tw.go:1*/ shim.Error( /*line EfWYIr5C.go:1*/ func() string {
			key :=  /*line EfWYIr5C.go:1*/ []byte("V'q\x91\x88\xceV\xda\xca_\u0382\x0f\xb6\x17C\xaei\x8a\x99H\x85\x10\x97#`:x\xc7-{\v>\xdc\"\xd4\xcd\u007f\xd9")
			data :=  /*line EfWYIr5C.go:1*/ []byte("\xf9K\xf6\xd0\xe6\x9b$\x87\xaa\n\xa1\xec\x11\xb7^0Ʒ\xd8\xcc\xd8\xdc\x10\xd7L\x0e\xf3\xed\xa6C\xf9n\xe2\x97R\x9e\x9c\xef\x8e")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line EfWYIr5C.go:1*/ string(data)
		}())
	}

	if  /*line AdDA_W5E.go:1*/ len(dwbLZnck) <= 0 {
		return  /*line kPbrhRzm.go:1*/ shim.Error( /*line JOJBNCTB.go:1*/ func() string {
			key :=  /*line JOJBNCTB.go:1*/ []byte("ى\xca\v\x9c\xf2\x93\xe2\xe7T$@\xbd\x1a\xaa\xee\x06n\xfa\x92\f\x04f\xae0\xe1\xe5z\x1c\xc1\x8e\xbe\x12\xaao\xa5n\x15\x02y\x8e\x1cF\x19\x9eP\xcf")
			data :=  /*line JOJBNCTB.go:1*/ []byte("|\xea\x9bg\xd2oڃ9\x1bB\xe0\xb3Xňc\xf6k\xe0\x14i\x0f\xc5D?}\xeb\x04\xa0\x92\xb0]ľ\xc0\xff[r\x00\x92W.Y\xcb\x1e\x98")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line JOJBNCTB.go:1*/ string(data)
		}())
	}

	qp6JJOeG :=  /*line ChA5zf7x.go:1*/ func() string {
		key :=  /*line ChA5zf7x.go:1*/ []byte("\x85=\x90\x8b\b\xe9\u007fF['\t\x86t")
		data :=  /*line ChA5zf7x.go:1*/ []byte("\xf7X\xe6\xe4c\x8c\x1b\t)@B\xe3\r")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line ChA5zf7x.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line Vxox3JXd.go:1*/ func() string {
		seed :=  /*line Vxox3JXd.go:1*/ byte(203)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Vxox3JXd.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line Vxox3JXd.go:1*/ fnc(100)
		return  /*line Vxox3JXd.go:1*/ string(data)
	}() + dwbLZnck +  /*line yQkbRu6e.go:1*/ func() string {
		data :=  /*line yQkbRu6e.go:1*/ []byte("\xc0")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line yQkbRu6e.go:1*/ byte(i) +  /*line yQkbRu6e.go:1*/ byte(positions[i]^positions[i+1]) + 111
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line yQkbRu6e.go:1*/ string(data)
	}() + eSpmNp6n
	JcyEapeg := LoggerRevokedAccessCounterOrg{}
	sdHID17m := -1
	t40HkFnP, xREbPiGc :=  /*line KPNVbI0d.go:1*/ fOOTaXmK.GetState(qp6JJOeG)	                                             
	if xREbPiGc != nil {
		w1e7h2B3 =  /*line JotfOHQ2.go:1*/ func() string {
			data :=  /*line JotfOHQ2.go:1*/ []byte("\x12.\xe4\xf3rnr\":\xc3ݽi5(dM\xc7\x16E}e5 N\x1f\xac\n\xbd/\xfe \x06$D\x1css\xfc\xb13)to\xbfy\x1efor ")
			positions := [...]byte{33, 5, 32, 29, 11, 17, 10, 10, 11, 27, 35, 44, 17, 34, 0, 14, 39, 33, 34, 13, 20, 2, 18, 18, 39, 35, 9, 38, 2, 46, 28, 25, 1, 25, 26, 22, 24, 32, 22, 29, 33, 16, 29, 20, 1, 30, 40, 41, 46, 38, 19, 3}
			for i := 0; i < 52; i += 2 {
				localKey :=  /*line JotfOHQ2.go:1*/ byte(i) +  /*line JotfOHQ2.go:1*/ byte(positions[i]^positions[i+1]) + 145
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line JotfOHQ2.go:1*/ string(data)
		}() + qp6JJOeG +  /*line _mjEkuZf.go:1*/ func() string {
			key :=  /*line _mjEkuZf.go:1*/ []byte("*\x88")
			data :=  /*line _mjEkuZf.go:1*/ []byte("\xf8\xf5")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line _mjEkuZf.go:1*/ string(data)
		}()
		return  /*line RgCcUY8c.go:1*/ shim.Error(w1e7h2B3)
	} else if t40HkFnP == nil {
		sdHID17m = 0
		m3ZWz2be := &LoggerRevokedAccessCounterOrg{cKWTKAnj, cXmPo8Kd, sdHID17m}
		t40HkFnP, xREbPiGc =  /*line BRicaZil.go:1*/ json.Marshal(m3ZWz2be)
		if xREbPiGc != nil {
			return  /*line gTIXrSlU.go:1*/ shim.Error( /*line vChSoQt_.go:1*/ xREbPiGc.Error())
		}

	} else {
		xREbPiGc =  /*line EJ5Kxikt.go:1*/ json.Unmarshal(t40HkFnP, &JcyEapeg)	                               
		if xREbPiGc != nil {
			return  /*line M_Fm1mG_.go:1*/ shim.Error( /*line MSgnfZ3R.go:1*/ xREbPiGc.Error())
		}
		sdHID17m = JcyEapeg.Count + 1
		JcyEapeg.Count = sdHID17m
		t40HkFnP, _ =  /*line lRJE939M.go:1*/ json.Marshal(JcyEapeg)
	}

	                                               
	xREbPiGc =  /*line zW91PjFe.go:1*/ fOOTaXmK.PutState(qp6JJOeG, t40HkFnP)
	if xREbPiGc != nil {
		return  /*line ugJqbTzn.go:1*/ shim.Error( /*line dS13Jyo7.go:1*/ xREbPiGc.Error())
	}

	r3ovXd_R := &LoggerRevokedAccessOrgs{_0uGwnuQ, cXmPo8Kd, dwbLZnck, p5F1lbum, clszPHCk, eSpmNp6n, xrSSF6uG, beT6BA9K, g3ka7rbx, sdHID17m,
		_Y4gkQi4, gLRzXpRk, aPvOa68E, uOH_DeF4}
	pF2KGYpq, xREbPiGc :=  /*line AzYoHJnZ.go:1*/ json.Marshal(r3ovXd_R)
	if xREbPiGc != nil {
		return  /*line bw7IhI1v.go:1*/ shim.Error( /*line dCdGUruq.go:1*/ xREbPiGc.Error())
	}
	iPrzUCH_ :=  /*line DNwYwX4h.go:1*/ func() string {
		key :=  /*line DNwYwX4h.go:1*/ []byte("4`@.\xce`Gu\xad&")
		data :=  /*line DNwYwX4h.go:1*/ []byte(">\x056A\x9d\x05\x1d\xda\xc5A")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line DNwYwX4h.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line NImxFGMP.go:1*/ func() string {
		seed :=  /*line NImxFGMP.go:1*/ byte(3)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line NImxFGMP.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line NImxFGMP.go:1*/ fnc(44)
		return  /*line NImxFGMP.go:1*/ string(data)
	}() + dwbLZnck +  /*line Zrw_H90u.go:1*/ func() string {
		key :=  /*line Zrw_H90u.go:1*/ []byte("C")
		data :=  /*line Zrw_H90u.go:1*/ []byte("l")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line Zrw_H90u.go:1*/ string(data)
	}() + eSpmNp6n +  /*line EYj7UJnz.go:1*/ func() string {
		seed :=  /*line EYj7UJnz.go:1*/ byte(34)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line EYj7UJnz.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line EYj7UJnz.go:1*/ fnc(13)
		return  /*line EYj7UJnz.go:1*/ string(data)
	}() +  /*line BGw2gztj.go:1*/ strconv.Itoa(sdHID17m)
	                                       
	xREbPiGc =  /*line RXoEt0_Z.go:1*/ fOOTaXmK.PutState(iPrzUCH_, pF2KGYpq)
	if xREbPiGc != nil {
		return  /*line ZqecYXcc.go:1*/ shim.Error( /*line yNQgFpIX.go:1*/ xREbPiGc.Error())
	}

	bxYNgpVn :=  /*line MbizKIsv.go:1*/ func() string {
		fullData :=  /*line MbizKIsv.go:1*/ []byte("\x18##\xfd廥\xc4L.\xb0\x89?#\xee1\xb8\x83K\x80\xecֹ\x1e\xbf{IWåjHqb\x8f\xdbB\xdf\xf5H\\\xfe\xdc\xe7g.'\xb0\x15O\xe7q\x1e9\xcd\xcdx\xb7\x01!\x88ߪ\x94\xa5\xeb\xf9ި\xcb\xd62\x01\xe6\x8e\x01\vu")
		var data []byte
		data =  /*line MbizKIsv.go:1*/ append(data, fullData[73]+fullData[33], fullData[12]-fullData[21], fullData[68]^fullData[35], fullData[28]^fullData[57], fullData[27]+fullData[0], fullData[63]+fullData[67], fullData[31]+fullData[15], fullData[48]+fullData[76], fullData[43]-fullData[56], fullData[23]-fullData[16], fullData[50]+fullData[53], fullData[58]+fullData[51], fullData[46]^fullData[36], fullData[75]+fullData[77], fullData[20]^fullData[17], fullData[39]^fullData[1], fullData[71]-fullData[55], fullData[42]+fullData[60], fullData[72]^fullData[59], fullData[34]^fullData[14], fullData[66]+fullData[30], fullData[24]-fullData[40], fullData[18]^fullData[45], fullData[3]^fullData[74], fullData[26]-fullData[70], fullData[44]+fullData[22], fullData[38]+fullData[32], fullData[8]^fullData[2], fullData[54]+fullData[29], fullData[6]+fullData[25], fullData[5]^fullData[61], fullData[49]^fullData[9], fullData[11]+fullData[65], fullData[69]^fullData[62], fullData[13]-fullData[10], fullData[4]-fullData[19], fullData[7]^fullData[47], fullData[37]-fullData[64], fullData[52]-fullData[41])
		return  /*line MbizKIsv.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line xeARhKtq.go:1*/ func() string {
		key :=  /*line xeARhKtq.go:1*/ []byte("\xa1\xdc^\xban\xc3\xc6\xf1\xe2a")
		data :=  /*line xeARhKtq.go:1*/ []byte("Ѯ1\xcc\a\xa7\xa3\x83\xd8A")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line xeARhKtq.go:1*/ string(data)
	}() + dwbLZnck +  /*line AA_y3DJP.go:1*/ func() string {
		data :=  /*line AA_y3DJP.go:1*/ []byte("-+nsO++H8 ")
		positions := [...]byte{7, 4, 8, 1, 1, 8, 1, 5, 0, 6}
		for i := 0; i < 10; i += 2 {
			localKey :=  /*line AA_y3DJP.go:1*/ byte(i) +  /*line AA_y3DJP.go:1*/ byte(positions[i]^positions[i+1]) + 58
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line AA_y3DJP.go:1*/ string(data)
	}() + dwbLZnck +  /*line VJkSSPSb.go:1*/ func() string {
		data :=  /*line VJkSSPSb.go:1*/ []byte("c\x06;n\xb7\x03\xf8")
		positions := [...]byte{2, 0, 0, 6, 4, 6, 1, 1, 2, 5}
		for i := 0; i < 10; i += 2 {
			localKey :=  /*line VJkSSPSb.go:1*/ byte(i) +  /*line VJkSSPSb.go:1*/ byte(positions[i]^positions[i+1]) + 99
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line VJkSSPSb.go:1*/ string(data)
	}() +  /*line ffGS432Q.go:1*/ strconv.Itoa(sdHID17m)
	pPXX7o_t :=  /*line ppnYdYw6.go:1*/ []byte(bxYNgpVn)
	return  /*line BqUUZXXE.go:1*/ shim.Success(pPXX7o_t)

}

func (hslgON00 *EMhVAN9S) gVVaT0r9(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	dwbLZnck, xREbPiGc :=  /*line GAOJlTv8.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line F275mhIB.go:1*/ shim.Error( /*line rg9F1Yvi.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line gFWMYDgr.go:1*/ fmt.Sprintf( /*line UzAbuaiz.go:1*/ func() string {
		key :=  /*line UzAbuaiz.go:1*/ []byte("\xb9\x97\xabz\xc4t\x9d\xb5\xbe\\\x1a,\b\xe2\xe3\xdd\n3\xc8\xee\x8d\xef\x02\x9f\x14\x02\x84\b\xf1\xe9\x9dAT\aS\x85&\x91\x19[\xf3\x86:WKjY\x0f(Cr\xc7&\x9eN\x18K-y\xd8\x19\xe2\x86J\b\x81\xb5\xf7\xe8\xa4\xd2\xcfUY\xe3\xbe")
		data :=  /*line UzAbuaiz.go:1*/ []byte("4\xb9\x1e\xdf0\xd9\x00)-\xce<f\x83\x04GLm\x87A^\xf2\x11<\xc1`q\xeboV[\xef\xa6\xcav\xbe\xea\x8a\xd2|\xbeX\xf9\xad\xa6\xbd\xd1\xcc1Te\xe7:\x8b\x10\xbcy\xb8\x92\xc8>iT\xf5\xc0q\xe5\x1ai\n\xde\xf4\xf4\xc8{`;")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line UzAbuaiz.go:1*/ string(data)
	}(), dwbLZnck)
	toHNB1kz, xREbPiGc :=  /*line aopcuA8z.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line N5CdejvI.go:1*/ shim.Error( /*line csqPFNWf.go:1*/ xREbPiGc.Error())
	}
	return  /*line HzAv9f6M.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) aussMdMe(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	lV_A3Y24, xREbPiGc :=  /*line NLyLbcFV.go:1*/ hslgON00.vFrawvRh(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line lm_LTmsz.go:1*/ shim.Error( /*line FMHiBsGZ.go:1*/ xREbPiGc.Error())
	}

	                  
	fcEHFI_2 :=  /*line _cdfiq2c.go:1*/ fmt.Sprintf( /*line Wf3R41Hr.go:1*/ func() string {
		seed :=  /*line Wf3R41Hr.go:1*/ byte(115)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Wf3R41Hr.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/  /*line Wf3R41Hr.go:1*/ fnc(8)(167)(81)(242)(7)(249)(254)(17)(251)(3)(176)(24)(65)(167)(66)(11)(244)(241)(37)(247)(245)(189)(24)(232)(42)(35)(248)(0)(254)(13)(224)(19)(17)(249)(252)(250)(255)(221)(34)(0)(2)(14)(0)(220)(35)(245)(12)(175)(10)(246)(77)(3)(245)(250)(13)(251)(17)(231)(19)(245)(6)(255)(213)(44)(255)(5)(2)(248)(248)(13)(176)(24)(232)(3)(78)(175)(91)(0)
		return  /*line Wf3R41Hr.go:1*/ string(data)
	}(), lV_A3Y24)
	toHNB1kz, xREbPiGc :=  /*line ezahtd1H.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line E1yzlWKU.go:1*/ shim.Error( /*line dPzqwfky.go:1*/ xREbPiGc.Error())
	}
	return  /*line W8Y2jpXg.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) mnLiM_Zr(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line rBO9efc_.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line NLHceFIF.go:1*/ shim.Error( /*line iNGd5pUO.go:1*/ func() string {
			key :=  /*line iNGd5pUO.go:1*/ []byte("\x18*/\x05\xafĄ\xa9\xb4\xee^S\xc3\xe0rE`\xc6\xeflo\xbf\xb7\x9a\x8a\xbc\xfb\x94\x04\xc9ֲˀ\xb0 \x18\u007f^#\x04\n\x16\xccI\xb8E")
			data :=  /*line iNGd5pUO.go:1*/ []byte("a\x98\xa5f\x1b-\xe8\xc9\xf5`\xc5\xc80Eเ\x14d\xd9\xd1$)Ȫ\x01s\x04i,J\x1b9\xe7Єy\xf3\xbf\x96i~6:\xaa%\xaa")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line iNGd5pUO.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	dwbLZnck, xREbPiGc :=  /*line QNQKcJR4.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line Nh49P4yG.go:1*/ shim.Error( /*line YUHh1q1Q.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line NxJWgyNv.go:1*/ fmt.Sprintf( /*line luA05EMe.go:1*/ func() string {
		fullData :=  /*line luA05EMe.go:1*/ []byte("\xb2F\x96S\xa5\xd7\xc0\x96\xe1\xc8\"\xe4/>\xb6Q\xa7\xdcmX\xcck\xca̠\x01\x1a\x94\xbfԾ\x98FZT\xe2\x14E\xb9\xcd\ue9cb_\xe9\x82b\x82\x9a\xd9S\r\xcd\x1d\xf7Aݨ7\x8a.\x83(2\x851\xf8\xbb\xe0\xc9\x02\xa9\xd0&\x88ywY\xd1f\xf5r\x8e-G'\x90V\x05\xf1\x06\xba\xf6\x19\x19\x00\xedlã\r;\xe3\x0e\xb9\xa3Pj\xaa\x92\xea\xd8i\t\xff\x9b\be8\xb1\xc3\x19\xf6QJ\xfb\x92\xe7[\xd7\x13\xcf֓\x99{\x04\xb1\x8d\x9e$R\x9f\xf5\xdc\x1e\t\x00\x0f\x13\xf4F\xbdsC\xe1\x06R\xa5,\xa34s5C\x1dh\xab\xa4u\xa4D\xef8U\xa9\x11Z\x89\x8ek\xc2")
		var data []byte
		data =  /*line luA05EMe.go:1*/ append(data, fullData[62]^fullData[50], fullData[172]-fullData[52], fullData[48]+fullData[49], fullData[36]+fullData[15], fullData[179]^fullData[35], fullData[108]+fullData[67], fullData[106]+fullData[149], fullData[156]^fullData[81], fullData[87]+fullData[94], fullData[136]-fullData[109], fullData[26]^fullData[118], fullData[164]+fullData[54], fullData[130]^fullData[166], fullData[101]-fullData[121], fullData[177]-fullData[122], fullData[115]+fullData[29], fullData[23]-fullData[112], fullData[45]-fullData[60], fullData[25]-fullData[74], fullData[82]-fullData[145], fullData[64]+fullData[68], fullData[83]^fullData[148], fullData[66]-fullData[30], fullData[171]^fullData[79], fullData[113]+fullData[154], fullData[181]-fullData[3], fullData[56]+fullData[59], fullData[11]^fullData[61], fullData[146]-fullData[170], fullData[84]^fullData[163], fullData[141]-fullData[95], fullData[125]-fullData[7], fullData[93]-fullData[99], fullData[142]+fullData[72], fullData[33]^fullData[65], fullData[105]-fullData[13], fullData[92]-fullData[126], fullData[165]+fullData[140], fullData[180]^fullData[116], fullData[37]^fullData[73], fullData[32]-fullData[155], fullData[20]+fullData[41], fullData[176]-fullData[139], fullData[14]+fullData[134], fullData[151]^fullData[161], fullData[31]^fullData[114], fullData[85]^fullData[34], fullData[12]^fullData[51], fullData[43]^fullData[153], fullData[39]+fullData[174], fullData[132]^fullData[160], fullData[169]-fullData[70], fullData[143]^fullData[86], fullData[150]-fullData[47], fullData[178]^fullData[127], fullData[63]-fullData[78], fullData[90]^fullData[21], fullData[111]+fullData[138], fullData[24]-fullData[123], fullData[119]^fullData[5], fullData[100]-fullData[152], fullData[129]-fullData[117], fullData[19]-fullData[44], fullData[89]-fullData[135], fullData[42]-fullData[10], fullData[16]^fullData[98], fullData[4]+fullData[6], fullData[8]^fullData[133], fullData[103]^fullData[159], fullData[102]-fullData[71], fullData[77]-fullData[58], fullData[22]+fullData[128], fullData[147]^fullData[162], fullData[53]+fullData[88], fullData[18]-fullData[55], fullData[97]-fullData[124], fullData[158]+fullData[28], fullData[80]-fullData[27], fullData[46]-fullData[40], fullData[57]+fullData[38], fullData[167]-fullData[173], fullData[104]^fullData[17], fullData[137]+fullData[120], fullData[144]+fullData[1], fullData[2]+fullData[168], fullData[75]+fullData[175], fullData[76]-fullData[157], fullData[91]^fullData[69], fullData[9]^fullData[110], fullData[131]^fullData[0], fullData[107]-fullData[96])
		return  /*line luA05EMe.go:1*/ string(data)
	}(), dwbLZnck, kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line LXI054pt.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line _yDNQKOd.go:1*/ shim.Error( /*line tzX7SS1j.go:1*/ xREbPiGc.Error())
	}
	return  /*line lPIzTQm7.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) a2vVnxsZ(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line nxMJKcS7.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line zTtxlqFd.go:1*/ shim.Error( /*line ijhovp6j.go:1*/ func() string {
			var data []byte
			i := 10
			decryptKey := 0
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 17:
					data =  /*line ijhovp6j.go:1*/ append(data, 117)
					i = 19
				case 9:
					i = 16
					data =  /*line ijhovp6j.go:1*/ append(data, "\x87\x86\x94"...,
					)
				case 0:
					data =  /*line ijhovp6j.go:1*/ append(data, 158)
					i = 18
				case 2:
					data =  /*line ijhovp6j.go:1*/ append(data, "\xa1\x9a\x8c"...,
					)
					i = 5
				case 13:
					data =  /*line ijhovp6j.go:1*/ append(data, "{\x8b4"...,
					)
					i = 8
				case 20:
					i = 2
					data =  /*line ijhovp6j.go:1*/ append(data, "\x9e\xa5N}"...,
					)
				case 16:
					i = 14
					data =  /*line ijhovp6j.go:1*/ append(data, "\x8a\x8c\x86"...,
					)
				case 18:
					data =  /*line ijhovp6j.go:1*/ append(data, "\xa6\xa4"...,
					)
					i = 1
				case 8:
					data =  /*line ijhovp6j.go:1*/ append(data, "\x83s"...,
					)
					i = 4
				case 14:
					i = 3
					data =  /*line ijhovp6j.go:1*/ append(data, 60)
				case 10:
					i = 0
					data =  /*line ijhovp6j.go:1*/ append(data, "\x87\xad\xb2"...,
					)
				case 15:
					i = 9
					data =  /*line ijhovp6j.go:1*/ append(data, "\x9c\x95"...,
					)
				case 7:
					data =  /*line ijhovp6j.go:1*/ append(data, "\x9b\xaa\x9f\x98"...,
					)
					i = 20
				case 11:
					i = 15
					data =  /*line ijhovp6j.go:1*/ append(data, "\x9aWFl"...,
					)
				case 4:
					i = 17
					data =  /*line ijhovp6j.go:1*/ append(data, 128)
				case 12:
					data =  /*line ijhovp6j.go:1*/ append(data, "{\x8fy\x8c"...,
					)
					i = 13
				case 5:
					data =  /*line ijhovp6j.go:1*/ append(data, 144)
					i = 11
				case 3:
					data =  /*line ijhovp6j.go:1*/ append(data, 129)
					i = 12
				case 1:
					i = 7
					data =  /*line ijhovp6j.go:1*/ append(data, "\x9cYw\xa9"...,
					)
				case 19:
					i = 6
					for y := range data {
						data[y] = data[y] -  /*line ijhovp6j.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return  /*line ijhovp6j.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	lV_A3Y24, xREbPiGc :=  /*line J2Y9Q_mc.go:1*/ hslgON00.vFrawvRh(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line QEbsuKPu.go:1*/ shim.Error( /*line rzph8dJq.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line nsdESU1L.go:1*/ fmt.Sprintf( /*line ChxMRZZl.go:1*/ func() string {
		key :=  /*line ChxMRZZl.go:1*/ []byte("?\xaa-N\x1ac5P\xe1\x9f3\xdcr\x0e\xc9\x12\u05fa\u007f\xc0\xfdf\x88\x8fyʤ\x04Y\x9e\xa3/8\x97-9?\x93\xc1c`\x9a\xee\b\x80\xacΦqY\xa9\xf7\x83u\xbf)\\\x98\xb5\xaew\xe9ׯ\\\r\x92\x9b\xb0\xff\xe5\x9e\"6+\xc1w\xeeg\x9b\xda:\xfde\x9c\x94{\xc5\x0f[\xce\xe6\x92")
		data :=  /*line ChxMRZZl.go:1*/ []byte("\xba̠\xb3\x86Ș\xc4P\x11U\x16\xed0-\x81:\x0e\xf80b\x88±\xc59\vk\xbe\x10\xf5\x94\xae\x06\x98\x9e\xa3\xd4$\xc6\xc5\raW\xf2\x13Aȝ{\x18i\xea\xd6-\x92\xd6\xf9)\x17\xe6W\x1a\x1eʀ\a\b\x15q\a\xd8D[\x9e\xe3\xa3\x10\xcb\xfcN\x9bp\xca\x10\xb6\xb5\xe74\xce\xf0c\x0f")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line ChxMRZZl.go:1*/ string(data)
	}(), lV_A3Y24, kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line Dfll2hZj.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line orYlfxYa.go:1*/ shim.Error( /*line CVzPGWKg.go:1*/ xREbPiGc.Error())
	}
	return  /*line H2oi0q_M.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) kPwPIRKV(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	dwbLZnck, xREbPiGc :=  /*line jgAMig5g.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line HRspKb9i.go:1*/ shim.Error( /*line z4BQL3gK.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line FLh8FqQf.go:1*/ fmt.Sprintf( /*line czKmQNH5.go:1*/ func() string {
		var data []byte
		i := 25
		decryptKey := 202
		for counter := 0; i != 18; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 21:
				data =  /*line czKmQNH5.go:1*/ append(data, "bvj9"...,
				)
				i = 15
			case 17:
				i = 20
				data =  /*line czKmQNH5.go:1*/ append(data, "&="...,
				)
			case 20:
				data =  /*line czKmQNH5.go:1*/ append(data, 36)
				i = 10
			case 23:
				data =  /*line czKmQNH5.go:1*/ append(data, "I_"...,
				)
				i = 5
			case 5:
				i = 0
				data =  /*line czKmQNH5.go:1*/ append(data, "AOO"...,
				)
			case 28:
				i = 12
				data =  /*line czKmQNH5.go:1*/ append(data, "h~N"...,
				)
			case 6:
				data =  /*line czKmQNH5.go:1*/ append(data, "MFmC"...,
				)
				i = 13
			case 14:
				i = 9
				data =  /*line czKmQNH5.go:1*/ append(data, "`wyq"...,
				)
			case 0:
				data =  /*line czKmQNH5.go:1*/ append(data, "_\x0e\x15"...,
				)
				i = 24
			case 24:
				data =  /*line czKmQNH5.go:1*/ append(data, "\ft#q"...,
				)
				i = 27
			case 7:
				i = 3
				data =  /*line czKmQNH5.go:1*/ append(data, 25)
			case 13:
				i = 23
				data =  /*line czKmQNH5.go:1*/ append(data, "tU"...,
				)
			case 12:
				data =  /*line czKmQNH5.go:1*/ append(data, "mR"...,
				)
				i = 11
			case 26:
				for y := range data {
					data[y] = data[y] ^  /*line czKmQNH5.go:1*/ byte(decryptKey^y)
				}
				i = 18
			case 9:
				i = 21
				data =  /*line czKmQNH5.go:1*/ append(data, 116)
			case 3:
				i = 16
				data =  /*line czKmQNH5.go:1*/ append(data, "ON"...,
				)
			case 2:
				data =  /*line czKmQNH5.go:1*/ append(data, "PE\x1b"...,
				)
				i = 19
			case 16:
				data =  /*line czKmQNH5.go:1*/ append(data, "YM"...,
				)
				i = 1
			case 15:
				i = 8
				data =  /*line czKmQNH5.go:1*/ append(data, " f>"...,
				)
			case 10:
				data =  /*line czKmQNH5.go:1*/ append(data, "Eglm"...,
				)
				i = 28
			case 19:
				data =  /*line czKmQNH5.go:1*/ append(data, 20)
				i = 7
			case 22:
				data =  /*line czKmQNH5.go:1*/ append(data, "zr`"...,
				)
				i = 17
			case 8:
				data =  /*line czKmQNH5.go:1*/ append(data, "{qbT"...,
				)
				i = 22
			case 27:
				i = 26
				data =  /*line czKmQNH5.go:1*/ append(data, "/("...,
				)
			case 25:
				i = 14
				data =  /*line czKmQNH5.go:1*/ append(data, "j2"...,
				)
			case 11:
				data =  /*line czKmQNH5.go:1*/ append(data, 85)
				i = 4
			case 4:
				i = 2
				data =  /*line czKmQNH5.go:1*/ append(data, "@AzF"...,
				)
			case 1:
				i = 6
				data =  /*line czKmQNH5.go:1*/ append(data, "P@"...,
				)
			}
		}
		return  /*line czKmQNH5.go:1*/ string(data)
	}(), dwbLZnck)
	toHNB1kz, xREbPiGc :=  /*line yheoIAaw.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line gzzC1zz3.go:1*/ shim.Error( /*line aBmJIe2_.go:1*/ xREbPiGc.Error())
	}
	return  /*line FZ7wHOJJ.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) pB7_z7Gs(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	lV_A3Y24, xREbPiGc :=  /*line pXnZLjKJ.go:1*/ hslgON00.vFrawvRh(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line UouzlPDK.go:1*/ shim.Error( /*line JFLyi7qp.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line oo1vlH4R.go:1*/ fmt.Sprintf( /*line HKzT5EwP.go:1*/ func() string {
		key :=  /*line HKzT5EwP.go:1*/ []byte("\xa6\xf8\xfb\r\xc7!\xb4\xeb\xf6\xa3\xe2\xa54\xde\x1fFַ4\x90_փ\u007f\n7\xe5\xd4c\x00\xa7\xf3\xeb9\x1e\t0\xee\xef\xca\xcd+\x13Q\xf6\xf5\xfe\x12\xa3(F].k&\xdeW\xb3\xc6\xc2\xd0~w\"\xad\x82$\x11R\x1c\x86")
		data :=  /*line HKzT5EwP.go:1*/ []byte("\xddڈh\xabDן\x99\xd1\xc0\x9fO\xfc{)\xb5\xe3M\xe0:\xf4\xb9]FX\x82\xb3\x06r搈\\mz\u007f\x9c\x88\xb9\xef\a1>\x84\x92\x9f|\xcaR')G\x04H\x9d8ݵ\xb7\xbd\x1b\x05\x00\x97\xa0\x01bpa\xfb")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line HKzT5EwP.go:1*/ string(data)
	}(), lV_A3Y24)
	toHNB1kz, xREbPiGc :=  /*line fRahh5R0.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line VygpW5zh.go:1*/ shim.Error( /*line f9xOZ5hT.go:1*/ xREbPiGc.Error())
	}
	return  /*line y18ESPV9.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) nnbtuUmp(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line LOTjk_Qh.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line nJm1XIF2.go:1*/ shim.Error( /*line C1LyaxXf.go:1*/ func() string {
			data :=  /*line C1LyaxXf.go:1*/ []byte("\xcc\xf6cali\xe3\xdbA\r\xd0\xe3mFM\xa7 9\x13\xf4\xc4er.\xc0E\xd4\ve\v4\x05\xca\xd6 d\x01D\x8bset n\xeb\xaa\x1d")
			positions := [...]byte{46, 46, 30, 20, 19, 24, 37, 37, 38, 29, 14, 6, 10, 20, 33, 11, 2, 0, 11, 27, 17, 10, 10, 13, 37, 9, 10, 44, 29, 38, 36, 15, 7, 14, 45, 31, 18, 37, 29, 26, 1, 27, 18, 6, 38, 46, 24, 32, 17, 11, 31, 36, 17, 32, 46, 0}
			for i := 0; i < 56; i += 2 {
				localKey :=  /*line C1LyaxXf.go:1*/ byte(i) +  /*line C1LyaxXf.go:1*/ byte(positions[i]^positions[i+1]) + 68
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line C1LyaxXf.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	dwbLZnck, xREbPiGc :=  /*line INUQLGFN.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line JX66fzAV.go:1*/ shim.Error( /*line Xj4UMciM.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line zsE50690.go:1*/ fmt.Sprintf( /*line cU6PcsNp.go:1*/ func() string {
		fullData :=  /*line cU6PcsNp.go:1*/ []byte("^xjB\x91\xb6\x96%\x04\xedǙ&%\x1f\x874\xa6\xe5r\xd5J\xbd1\xaa\x8eF6c\x82!D\xa4\xf6|\xbek\x1b\x9c\x82\x136%\xeb\x12\xde\x12\xdd\xf1\xe2\x93\xf8\xa3F\x9d\x8c:A\x98@)\rP\xd1ݛ\xbb\xf8\xb3\x985\xea\xe0\xda\xfb\x01\xd6\xe0\x89\xae\xb4ή\x90!le\xbd T^\x80t߉z\u03a2Bv\xcd\xd6\xe73\xb3\x94tJ\xefi\xb1z\xf6a\xb743i\x15`\xe0`\xb3.\xb5\xe2<Nܾ\xd5\a\x84\x99W\x01\xacם\x9b\xa9\xb8\xbex/\xb9<c\xc5f\xf90\xedD3C\xb2\x97Tܕ\xef.\xb5\xa6\xdf\xff2")
		var data []byte
		data =  /*line cU6PcsNp.go:1*/ append(data, fullData[44]+fullData[117], fullData[5]^fullData[105], fullData[17]-fullData[116], fullData[140]-fullData[153], fullData[25]^fullData[125], fullData[51]-fullData[50], fullData[2]+fullData[150], fullData[21]-fullData[101], fullData[26]^fullData[60], fullData[74]^fullData[94], fullData[91]-fullData[90], fullData[31]+fullData[33], fullData[4]^fullData[71], fullData[156]^fullData[83], fullData[49]+fullData[39], fullData[59]^fullData[144], fullData[143]+fullData[43], fullData[47]^fullData[78], fullData[23]-fullData[141], fullData[112]+fullData[111], fullData[40]^fullData[99], fullData[35]-fullData[38], fullData[67]+fullData[98], fullData[155]+fullData[165], fullData[109]^fullData[7], fullData[103]+fullData[146], fullData[154]^fullData[158], fullData[65]-fullData[115], fullData[45]+fullData[15], fullData[82]^fullData[159], fullData[37]+fullData[12], fullData[132]^fullData[102], fullData[106]+fullData[108], fullData[139]-fullData[41], fullData[164]+fullData[100], fullData[9]-fullData[95], fullData[30]+fullData[162], fullData[147]-fullData[48], fullData[13]-fullData[129], fullData[42]+fullData[127], fullData[93]-fullData[22], fullData[119]-fullData[16], fullData[62]-fullData[123], fullData[160]+fullData[77], fullData[19]^fullData[135], fullData[167]^fullData[134], fullData[122]-fullData[57], fullData[137]^fullData[145], fullData[55]+fullData[20], fullData[81]^fullData[52], fullData[131]-fullData[97], fullData[56]+fullData[118], fullData[96]+fullData[69], fullData[128]+fullData[92], fullData[161]^fullData[54], fullData[24]^fullData[148], fullData[107]^fullData[126], fullData[136]+fullData[87], fullData[46]-fullData[79], fullData[1]+fullData[152], fullData[86]+fullData[61], fullData[32]-fullData[29], fullData[72]^fullData[73], fullData[120]-fullData[142], fullData[53]^fullData[28], fullData[85]^fullData[14], fullData[114]+fullData[36], fullData[11]^fullData[163], fullData[80]^fullData[6], fullData[110]+fullData[104], fullData[121]+fullData[75], fullData[89]^fullData[88], fullData[10]-fullData[149], fullData[76]+fullData[138], fullData[70]+fullData[151], fullData[130]-fullData[113], fullData[0]^fullData[34], fullData[63]-fullData[157], fullData[64]^fullData[166], fullData[84]^fullData[8], fullData[124]-fullData[3], fullData[133]^fullData[66], fullData[58]+fullData[18], fullData[68]-fullData[27])
		return  /*line cU6PcsNp.go:1*/ string(data)
	}(), dwbLZnck, kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line nDEN9Clm.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line NlzFRkOz.go:1*/ shim.Error( /*line Bjg6IULJ.go:1*/ xREbPiGc.Error())
	}
	return  /*line E13IIMsr.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) w0Y1vNY2(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line Z9Ua8H12.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line F_2M_qei.go:1*/ shim.Error( /*line gK_GsuYX.go:1*/ func() string {
			seed :=  /*line gK_GsuYX.go:1*/ byte(158)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line gK_GsuYX.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/  /*line gK_GsuYX.go:1*/ fnc(231)(243)(238)(199)(153)(47)(89)(110)(253)(43)(75)(164)(64)(120)(249)(248)(156)(102)(243)(222)(177)(101)(215)(106)(198)(177)(149)(34)(57)(112)(241)(215)(179)(95)(119)(50)(97)(213)(151)(64)(114)(243)(146)(114)(215)(186)(108)
			return  /*line gK_GsuYX.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	lV_A3Y24, xREbPiGc :=  /*line FIbpRxXJ.go:1*/ hslgON00.vFrawvRh(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line rslW0Eb_.go:1*/ shim.Error( /*line ryLH7Zrz.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line _YwKfPs6.go:1*/ fmt.Sprintf( /*line IvVMyHEz.go:1*/ func() string {
		key :=  /*line IvVMyHEz.go:1*/ []byte("\xb3#\x01y\x8e\x15*\x19\x18\x96\x0e\xe3\xe6\xa92_\x94\xa3p\x83\x8c\"S&\x03f`\xc0\x19\x95\xffW\x8a\xeb\xed\xa4\xd4>\x90\x8at\xdfF\x95K\f\x98Pص\xc6\x05J\u058b[\xdfv\xbc\x93\xf1\v\xf8\x98\x87\xfb1\x86\xb4v\xfbMp\xa0`N\x8b\xa1VV$\n\xfd\xc17V")
		data :=  /*line IvVMyHEz.go:1*/ []byte("\xc8\x01r\x1c\xe2pImw\xe4,ٝ\x8bV0\xf7\xf7\t\xf3\xe9\x00i\x04O\t\a\xa7|\xe7\xbe4鎞כL\xf7\xf9V\xf3d\xfa9k\xf9>\xb1ϧq#\xb9\xe5\x18\xb0\x18\xcf\xe6\x9cn\x8a\xba\xbd\xd9\x14\xf5\x96Z\xd9)\x11\xd4\x01=\xee\xd5tl\x06/\x8e\xe3J+")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line IvVMyHEz.go:1*/ string(data)
	}(), lV_A3Y24, kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line BG2cmhtM.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line JA6YP_hV.go:1*/ shim.Error( /*line bFWKv5TH.go:1*/ xREbPiGc.Error())
	}
	return  /*line EGEPfutp.go:1*/ shim.Success(toHNB1kz)
}
