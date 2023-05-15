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

func (hslgON00 *EMhVAN9S) mdXE5msh(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	var w1e7h2B3 string
	var xREbPiGc error
	if  /*line Lf3qtATQ.go:1*/ len(eFz_TaL1) != 9 {
		return  /*line LNig0RIf.go:1*/ shim.Error( /*line e5oS_UVz.go:1*/ func() string {
			key :=  /*line e5oS_UVz.go:1*/ []byte("\xd4\x10\xef79\xd0\xe1|cF\x1e-\x00\x1e\xc8M3=|\x9e\xa0\xf3\xbb\xcb\xeb\xd0}\xa9v6ݣHR\x00L]\x06\xc1\v\xd0\xcb6ʶ\x12\xe2\x91u\x92\x8e-\x1ca9_G\xdcѢE\xf1\xee\x9f\x1a\x87h\xee\xea\xd3wӿw\x83\xa4\xe9\x9d\xf1")
			data :=  /*line e5oS_UVz.go:1*/ []byte("\x1d~R\xa6\xabBF\xdf\xd7f\u007f\x9fg\x935\xb2\xa1\xb1\xef\xcc\xc0?*2R5\xef\xc9\xc9y\xfd\t\xb7\xc4 \xb9\xcct*\u007f?=\x9f8\x1d2T\xf6\xeb\x01\xf9\x92\x80\x81\x9aªAD\x15eVf\x0f\u007f\xea\xdcWX:\x97\f\xdf\xed\xe4\x10^\x02d")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line e5oS_UVz.go:1*/ string(data)
		}())
	}

	cXmPo8Kd := eFz_TaL1[0]
	dwbLZnck := eFz_TaL1[1]
	nIZZng5T := eFz_TaL1[2]
	eBmFb7v9 := eFz_TaL1[3]
	beT6BA9K := eFz_TaL1[4]
	xrSSF6uG := eFz_TaL1[5]
	g3ka7rbx :=  /*line iZFSv2KD.go:1*/ strings.Split(eFz_TaL1[6],  /*line jTJECk3G.go:1*/ func() string {
		data :=  /*line jTJECk3G.go:1*/ []byte("7")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line jTJECk3G.go:1*/ byte(i) +  /*line jTJECk3G.go:1*/ byte(positions[i]^positions[i+1]) + 27
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line jTJECk3G.go:1*/ string(data)
	}())
	p5F1lbum := eFz_TaL1[7]
	clszPHCk := eFz_TaL1[8]
	_0uGwnuQ :=  /*line _A_Ef8iu.go:1*/ func() string {
		seed :=  /*line _A_Ef8iu.go:1*/ byte(25)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line _A_Ef8iu.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/  /*line _A_Ef8iu.go:1*/ fnc(51)(35)(248)(0)(254)(13)(224)(19)(17)(249)(252)(250)(255)(221)(34)(0)(2)(14)(0)
		return  /*line _A_Ef8iu.go:1*/ string(data)
	}()
	cKWTKAnj :=  /*line jeecx7_H.go:1*/ func() string {
		var data []byte
		i := 6
		decryptKey := 108
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 7:
				data =  /*line jeecx7_H.go:1*/ append(data, "\xf7\xf4\xcd\xec"...,
				)
				i = 4
			case 5:
				data =  /*line jeecx7_H.go:1*/ append(data, "\xfb\x11"...,
				)
				i = 0
			case 8:
				i = 9
				data =  /*line jeecx7_H.go:1*/ append(data, 252)
			case 2:
				i = 1
				for y := range data {
					data[y] = data[y] +  /*line jeecx7_H.go:1*/ byte(decryptKey^y)
				}
			case 0:
				data =  /*line jeecx7_H.go:1*/ append(data, "\v\x04\xff\x03"...,
				)
				i = 3
			case 10:
				data =  /*line jeecx7_H.go:1*/ append(data, 231)
				i = 5
			case 3:
				i = 7
				data =  /*line jeecx7_H.go:1*/ append(data, "\xe1\x00\x01\xe8"...,
				)
			case 6:
				data =  /*line jeecx7_H.go:1*/ append(data, "\xdf\x03\xf8\xf9"...,
				)
				i = 8
			case 9:
				data =  /*line jeecx7_H.go:1*/ append(data, 10)
				i = 10
			case 4:
				data =  /*line jeecx7_H.go:1*/ append(data, 1)
				i = 2
			}
		}
		return  /*line jeecx7_H.go:1*/ string(data)
	}()

	if  /*line LDKaUTWw.go:1*/ len(cXmPo8Kd) <= 0 {
		return  /*line RfAZQQ2B.go:1*/ shim.Error( /*line aY_FbPC5.go:1*/ func() string {
			data :=  /*line aY_FbPC5.go:1*/ []byte("Da\xb9A\xab5t\xb4m\xb4\xe4t\x9c$\xccc\x91 \xb8\xa7\xc6-emp[M \xd0\x03r\xf1\xad\x02")
			positions := [...]byte{33, 26, 13, 20, 14, 28, 26, 29, 13, 19, 3, 32, 26, 5, 4, 19, 20, 32, 32, 9, 7, 9, 4, 18, 25, 18, 7, 31, 15, 28, 18, 12, 15, 10, 31, 28, 12, 18, 33, 16, 2, 18, 26, 12}
			for i := 0; i < 44; i += 2 {
				localKey :=  /*line aY_FbPC5.go:1*/ byte(i) +  /*line aY_FbPC5.go:1*/ byte(positions[i]^positions[i+1]) + 159
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line aY_FbPC5.go:1*/ string(data)
		}())
	}

	if  /*line FXaTzYuh.go:1*/ len(eBmFb7v9) <= 0 {
		return  /*line klEjop0b.go:1*/ shim.Error( /*line SnVSoLwC.go:1*/ func() string {
			key :=  /*line SnVSoLwC.go:1*/ []byte("n\x1f\xb3\x87q)k\xe5de\xc9Ϊ~Pvj\x1c\x88\xdb\xee\xc9\x13w\xc1ꈝ\xde\xe0\f\xcd0J\x15\x9d\xbco\xc9")
			data :=  /*line SnVSoLwC.go:1*/ []byte("\xe1S\xb4\xda\xfd@\x0f|\x10\x04\xa6\xa0v\xef%\xfd\n\x04ڊ2\x98\r\xf7\xae\x84\xa5ȏ\x90h\xac\xf0)_խ\xff\x9e")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line SnVSoLwC.go:1*/ string(data)
		}())
	}

	if  /*line I7vceH00.go:1*/ len(dwbLZnck) <= 0 {
		return  /*line Iqk_dyPE.go:1*/ shim.Error( /*line ALMiRROA.go:1*/ func() string {
			data :=  /*line ALMiRROA.go:1*/ []byte("\x01\be\tn\xde1\xf1 of\tp\x80\xa5\xea1\xa7\x18r\xb4m\xed\x03\x05\xfb\xa5\xf1\xf3\x9e\xb1\n\x82n}b\xe9pty \xf3t\x17\x8d\x9bg")
			positions := [...]byte{0, 45, 20, 25, 44, 35, 34, 30, 11, 26, 6, 20, 22, 32, 31, 18, 29, 34, 31, 35, 44, 16, 30, 3, 27, 7, 29, 25, 1, 30, 20, 30, 28, 6, 24, 22, 15, 43, 23, 1, 41, 35, 0, 1, 44, 16, 13, 11, 15, 30, 5, 0, 20, 17, 14, 11, 44, 36}
			for i := 0; i < 58; i += 2 {
				localKey :=  /*line ALMiRROA.go:1*/ byte(i) +  /*line ALMiRROA.go:1*/ byte(positions[i]^positions[i+1]) + 64
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line ALMiRROA.go:1*/ string(data)
		}())
	}

	if  /*line UVRd2v08.go:1*/ len(nIZZng5T) <= 0 {
		return  /*line sXqGNZwQ.go:1*/ shim.Error( /*line tFOhP0qv.go:1*/ func() string {
			fullData :=  /*line tFOhP0qv.go:1*/ []byte("\x85C\x92\x10\x04\xe8W\x9fU\x9d'\x8b\x13^\xdc#W\a!\xef\x10\xfb\xca#\x04\x16\x99\x91\xdb`\x8dP6du\x10\xd9\xef/\xdb-o\xf7>\x9d\x81J~\x8a\xab\x9e(Ȉ\x87\tw3?\x93\x8d\x030\xc6)\u0096\x17p\xf7z\x90V\t\xbd\x9b\x10\x9b'\x06\x81L\xf9\xe8j\xbdR\x1e\xf0\f\xa2\xb9 \x03")
			var data []byte
			data =  /*line tFOhP0qv.go:1*/ append(data, fullData[58]+fullData[25], fullData[40]^fullData[13], fullData[20]-fullData[49], fullData[43]^fullData[81], fullData[14]+fullData[2], fullData[39]-fullData[70], fullData[38]-fullData[65], fullData[28]+fullData[48], fullData[74]^fullData[9], fullData[54]+fullData[5], fullData[27]^fullData[42], fullData[12]^fullData[57], fullData[92]+fullData[1], fullData[34]-fullData[79], fullData[69]+fullData[56], fullData[22]-fullData[6], fullData[78]^fullData[86], fullData[72]+fullData[67], fullData[8]+fullData[3], fullData[76]-fullData[50], fullData[60]+fullData[59], fullData[85]-fullData[31], fullData[66]-fullData[18], fullData[89]-fullData[26], fullData[19]^fullData[75], fullData[17]^fullData[10], fullData[51]^fullData[46], fullData[87]-fullData[91], fullData[23]^fullData[93], fullData[36]+fullData[53], fullData[47]+fullData[90], fullData[37]-fullData[45], fullData[88]-fullData[80], fullData[84]+fullData[24], fullData[71]+fullData[44], fullData[63]+fullData[7], fullData[83]^fullData[0], fullData[21]^fullData[11], fullData[82]^fullData[30], fullData[68]^fullData[55], fullData[62]-fullData[35], fullData[4]+fullData[41], fullData[15]^fullData[16], fullData[77]-fullData[64], fullData[73]^fullData[29], fullData[32]-fullData[52], fullData[33]^fullData[61])
			return  /*line tFOhP0qv.go:1*/ string(data)
		}())
	}

	eGdk4lUz :=  /*line HtQuizaA.go:1*/ func() string {
		key :=  /*line HtQuizaA.go:1*/ []byte("\xbe\xab\x9c\xb9\x9ce\xeb\x8a/\xdb\x1f\x19\xe2\xcd")
		data :=  /*line HtQuizaA.go:1*/ []byte("\xcc\xce\xea\xd6\xf7\x00\x8f\xdf\\\xbemR\xa7\x94")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line HtQuizaA.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line gTvon_FL.go:1*/ func() string {
		seed :=  /*line gTvon_FL.go:1*/ byte(157)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line gTvon_FL.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line gTvon_FL.go:1*/ fnc(178)
		return  /*line gTvon_FL.go:1*/ string(data)
	}() + dwbLZnck +  /*line tj_oR0C7.go:1*/ func() string {
		seed :=  /*line tj_oR0C7.go:1*/ byte(71)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line tj_oR0C7.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line tj_oR0C7.go:1*/ fnc(104)
		return  /*line tj_oR0C7.go:1*/ string(data)
	}() + nIZZng5T
	HnqggbWJ := LoggerRevokedAccessCounter{}
	sdHID17m := -1
	oykk_osq, xREbPiGc :=  /*line qihJ2D7l.go:1*/ fOOTaXmK.GetState(eGdk4lUz)	                                             
	if xREbPiGc != nil {
		w1e7h2B3 =  /*line MqggWep3.go:1*/ func() string {
			seed :=  /*line MqggWep3.go:1*/ byte(130)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line MqggWep3.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/  /*line MqggWep3.go:1*/ fnc(253)(161)(101)(247)(238)(217)(181)(26)(76)(128)(36)(99)(206)(159)(55)(109)(150)(128)(251)(167)(149)(40)(95)(106)(38)(63)(143)(23)(42)(78)(155)(242)(37)(76)(152)(50)(114)(228)(117)(50)(101)(212)(169)(77)(157)(65)(41)(152)(57)(117)(152)
			return  /*line MqggWep3.go:1*/ string(data)
		}() + eGdk4lUz +  /*line ZuzdaDQR.go:1*/ func() string {
			data :=  /*line ZuzdaDQR.go:1*/ []byte("G}")
			positions := [...]byte{0, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line ZuzdaDQR.go:1*/ byte(i) +  /*line ZuzdaDQR.go:1*/ byte(positions[i]^positions[i+1]) + 37
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line ZuzdaDQR.go:1*/ string(data)
		}()
		return  /*line oR3CKuTA.go:1*/ shim.Error(w1e7h2B3)
	} else if oykk_osq == nil {
		sdHID17m = 0
		dGu9RrHS := &LoggerRevokedAccessCounter{cKWTKAnj, cXmPo8Kd, sdHID17m}
		oykk_osq, xREbPiGc =  /*line Jrz0cTWd.go:1*/ json.Marshal(dGu9RrHS)
		if xREbPiGc != nil {
			return  /*line Hoz7tzPU.go:1*/ shim.Error( /*line EyDXOE4I.go:1*/ xREbPiGc.Error())
		}

	} else {
		xREbPiGc =  /*line hGQ9yMJ3.go:1*/ json.Unmarshal(oykk_osq, &HnqggbWJ)	                               
		if xREbPiGc != nil {
			return  /*line cW4a9urO.go:1*/ shim.Error( /*line lux_RX6_.go:1*/ xREbPiGc.Error())
		}
		sdHID17m = HnqggbWJ.Count + 1
		HnqggbWJ.Count = sdHID17m
		oykk_osq, _ =  /*line Ec1Oto6p.go:1*/ json.Marshal(HnqggbWJ)
	}

	                                               
	xREbPiGc =  /*line tfnAI_Nr.go:1*/ fOOTaXmK.PutState(eGdk4lUz, oykk_osq)
	if xREbPiGc != nil {
		return  /*line rZBz7_hJ.go:1*/ shim.Error( /*line CWBgzHFq.go:1*/ xREbPiGc.Error())
	}

	r3ovXd_R := &LoggerRevokedAccess{_0uGwnuQ, cXmPo8Kd, dwbLZnck, p5F1lbum, clszPHCk, nIZZng5T, eBmFb7v9, xrSSF6uG, beT6BA9K, g3ka7rbx, sdHID17m}
	pF2KGYpq, xREbPiGc :=  /*line TuI0Ap70.go:1*/ json.Marshal(r3ovXd_R)
	if xREbPiGc != nil {
		return  /*line J7Gksqfp.go:1*/ shim.Error( /*line Gx1wTQwf.go:1*/ xREbPiGc.Error())
	}
	iPrzUCH_ :=  /*line tb00HcvZ.go:1*/ func() string {
		seed :=  /*line tb00HcvZ.go:1*/ byte(34)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line tb00HcvZ.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line tb00HcvZ.go:1*/  /*line tb00HcvZ.go:1*/  /*line tb00HcvZ.go:1*/  /*line tb00HcvZ.go:1*/  /*line tb00HcvZ.go:1*/  /*line tb00HcvZ.go:1*/  /*line tb00HcvZ.go:1*/  /*line tb00HcvZ.go:1*/  /*line tb00HcvZ.go:1*/  /*line tb00HcvZ.go:1*/  /*line tb00HcvZ.go:1*/ fnc(80)(243)(17)(249)(252)(250)(255)(241)(30)(242)(13)
		return  /*line tb00HcvZ.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line FHY0Xwpx.go:1*/ func() string {
		var data []byte
		i := 0
		decryptKey := 242
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				i = 1
				for y := range data {
					data[y] = data[y] ^  /*line FHY0Xwpx.go:1*/ byte(decryptKey^y)
				}
			case 0:
				i = 2
				data =  /*line FHY0Xwpx.go:1*/ append(data, 223)
			}
		}
		return  /*line FHY0Xwpx.go:1*/ string(data)
	}() + dwbLZnck +  /*line SmcjLTYN.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 26
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				data =  /*line SmcjLTYN.go:1*/ append(data, 20)
				i = 1
			case 1:
				for y := range data {
					data[y] = data[y] +  /*line SmcjLTYN.go:1*/ byte(decryptKey^y)
				}
				i = 0
			}
		}
		return  /*line SmcjLTYN.go:1*/ string(data)
	}() + nIZZng5T +  /*line SHGw5Ev7.go:1*/ func() string {
		data :=  /*line SHGw5Ev7.go:1*/ []byte("\x04")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line SHGw5Ev7.go:1*/ byte(i) +  /*line SHGw5Ev7.go:1*/ byte(positions[i]^positions[i+1]) + 43
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line SHGw5Ev7.go:1*/ string(data)
	}() +  /*line f2awE0F7.go:1*/ strconv.Itoa(sdHID17m)
	                                       
	xREbPiGc =  /*line dXj_m2Th.go:1*/ fOOTaXmK.PutState(iPrzUCH_, pF2KGYpq)
	if xREbPiGc != nil {
		return  /*line RQDIRy8I.go:1*/ shim.Error( /*line CI20Z8ps.go:1*/ xREbPiGc.Error())
	}

	bxYNgpVn :=  /*line ba4AD2D_.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 223
		for counter := 0; i != 11; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 16:
				data =  /*line ba4AD2D_.go:1*/ append(data, "7?I\xf6"...,
				)
				i = 13
			case 10:
				for y := range data {
					data[y] = data[y] -  /*line ba4AD2D_.go:1*/ byte(decryptKey^y)
				}
				i = 11
			case 7:
				i = 4
				data =  /*line ba4AD2D_.go:1*/ append(data, 57)
			case 13:
				data =  /*line ba4AD2D_.go:1*/ append(data, "95_K"...,
				)
				i = 15
			case 14:
				data =  /*line ba4AD2D_.go:1*/ append(data, ">\xf9"...,
				)
				i = 7
			case 0:
				i = 3
				data =  /*line ba4AD2D_.go:1*/ append(data, 70)
			case 8:
				data =  /*line ba4AD2D_.go:1*/ append(data, 77)
				i = 17
			case 2:
				data =  /*line ba4AD2D_.go:1*/ append(data, "<>"...,
				)
				i = 18
			case 15:
				i = 8
				data =  /*line ba4AD2D_.go:1*/ append(data, 92)
			case 12:
				data =  /*line ba4AD2D_.go:1*/ append(data, "F\xec2("...,
				)
				i = 9
			case 5:
				i = 14
				data =  /*line ba4AD2D_.go:1*/ append(data, "4/@"...,
				)
			case 6:
				data =  /*line ba4AD2D_.go:1*/ append(data, "ABO"...,
				)
				i = 0
			case 18:
				i = 12
				data =  /*line ba4AD2D_.go:1*/ append(data, 64)
			case 17:
				i = 10
				data =  /*line ba4AD2D_.go:1*/ append(data, "c(\r"...,
				)
			case 4:
				i = 6
				data =  /*line ba4AD2D_.go:1*/ append(data, 66)
			case 3:
				data =  /*line ba4AD2D_.go:1*/ append(data, 242)
				i = 16
			case 9:
				data =  /*line ba4AD2D_.go:1*/ append(data, "\xe12,<"...,
				)
				i = 5
			case 1:
				data =  /*line ba4AD2D_.go:1*/ append(data, "\x133<"...,
				)
				i = 2
			}
		}
		return  /*line ba4AD2D_.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line W32T5miZ.go:1*/ func() string {
		fullData :=  /*line W32T5miZ.go:1*/ []byte("\xdd\xcb=\f\tH,\x8f\xee\xa0\x1cꜞ>V\x18\x92|\xee")
		var data []byte
		data =  /*line W32T5miZ.go:1*/ append(data, fullData[13]^fullData[8], fullData[12]^fullData[19], fullData[0]+fullData[17], fullData[14]^fullData[5], fullData[4]-fullData[9], fullData[18]-fullData[16], fullData[7]^fullData[11], fullData[2]-fullData[1], fullData[15]-fullData[10], fullData[3]^fullData[6])
		return  /*line W32T5miZ.go:1*/ string(data)
	}() + dwbLZnck +  /*line HS_oMdwd.go:1*/ func() string {
		key :=  /*line HS_oMdwd.go:1*/ []byte("߶m\fL\xa9\x00\xeb\x9bM")
		data :=  /*line HS_oMdwd.go:1*/ []byte("\x84\xb9\x01g)\xc4e\x87\x9f\xd3")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line HS_oMdwd.go:1*/ string(data)
	}() + dwbLZnck +  /*line ZMaOzjCT.go:1*/ func() string {
		seed :=  /*line ZMaOzjCT.go:1*/ byte(150)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line ZMaOzjCT.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line ZMaOzjCT.go:1*/  /*line ZMaOzjCT.go:1*/  /*line ZMaOzjCT.go:1*/  /*line ZMaOzjCT.go:1*/  /*line ZMaOzjCT.go:1*/  /*line ZMaOzjCT.go:1*/  /*line ZMaOzjCT.go:1*/ fnc(205)(12)(6)(249)(6)(198)(230)
		return  /*line ZMaOzjCT.go:1*/ string(data)
	}() +  /*line XWIplBsR.go:1*/ strconv.Itoa(sdHID17m)
	pPXX7o_t :=  /*line hrXTJExC.go:1*/ []byte(bxYNgpVn)
	return  /*line nzZJPVU2.go:1*/ shim.Success(pPXX7o_t)

}

func (hslgON00 *EMhVAN9S) m2P5xAJz(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	dwbLZnck, xREbPiGc :=  /*line xjuup4dQ.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line KT5riYoV.go:1*/ shim.Error( /*line b9ODh25f.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line umh051Ak.go:1*/ fmt.Sprintf( /*line AGlAKa71.go:1*/ func() string {
		key :=  /*line AGlAKa71.go:1*/ []byte("\x8b\xcd\a~\xe9\x16 ѝ\x8a$\f~\x12\x04\xf6OY+\xa5\xcalڢQ\x96\xefn\u007f\xefo\xc9v\xe28\xe6H\xba\xd4w\x99\xe3\xdc\xfa!\xaa\xa5µ\x91_A\rV\x81C\x85\x83\x89\x93\xdf%\xacA\x12z߉\xc7t\x97\u007f")
		data :=  /*line AGlAKa71.go:1*/ []byte("\xf0\xeft\x1b\x85sC\xa5\xf2\xf8\x066\x050`\x99,\rRկN\xe0\x80\x1d\xf9\x88\t\x1a\x9d=\xac\x00\x8dS\x83,\xfb\xb7\x14\xfc\x90\xaf\xd8\r\x88б\xd0\xe31 `3\xce%\xd5\xf1\xe6\xe5\xb6A\xc930@\xfd\xac\xb4V\xea\x02")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line AGlAKa71.go:1*/ string(data)
	}(), dwbLZnck)
	toHNB1kz, xREbPiGc :=  /*line Hf0fXJDU.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line IYzw_F0r.go:1*/ shim.Error( /*line Q3_bdJXs.go:1*/ xREbPiGc.Error())
	}
	return  /*line AkKzMkV4.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) dc08LXlW(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	nIZZng5T, xREbPiGc :=  /*line BbwYqC_g.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line PPav8xFO.go:1*/ shim.Error( /*line In7EAeqh.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line A9iwP7Db.go:1*/ fmt.Sprintf( /*line ADUyryIn.go:1*/ func() string {
		data :=  /*line ADUyryIn.go:1*/ []byte("\xa0\x97߷xe\xe7to\xad֕{\xc2doc\xc8\xcd}\x87x\xca\xcbLo\x95Ne\xe3hq\x04o ed\x8d\xb3c\x9as|\",\"\x88zF\x17\xffa^\x87\xaa\x89CUn\xe3u\x93\v\x1e\x19:\x8a%N\x8a\x9e\xe5")
		positions := [...]byte{32, 13, 34, 48, 3, 17, 1, 50, 63, 34, 69, 66, 1, 31, 2, 10, 70, 0, 52, 11, 26, 37, 29, 22, 34, 46, 19, 23, 13, 62, 54, 4, 13, 30, 11, 37, 18, 6, 48, 59, 49, 68, 9, 18, 11, 57, 31, 71, 30, 4, 27, 47, 42, 6, 53, 20, 22, 50, 63, 40, 64, 37, 61, 48, 10, 34, 38, 55, 21, 57, 47, 68, 62, 55, 64, 31}
		for i := 0; i < 76; i += 2 {
			localKey :=  /*line ADUyryIn.go:1*/ byte(i) +  /*line ADUyryIn.go:1*/ byte(positions[i]^positions[i+1]) + 135
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line ADUyryIn.go:1*/ string(data)
	}(), nIZZng5T)
	toHNB1kz, xREbPiGc :=  /*line a89x_sk5.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line yldV0y3j.go:1*/ shim.Error( /*line FYO7naZO.go:1*/ xREbPiGc.Error())
	}
	return  /*line _vTSVxy8.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) lFK9ocLr(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line iE46CTYc.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line TTgCptCO.go:1*/ shim.Error( /*line ENqPOzTj.go:1*/ func() string {
			fullData :=  /*line ENqPOzTj.go:1*/ []byte("\u07bb\xe9*U7\x8f*\xd9\x1c[O\xe6p\x1f\xd8#י\x94\x98\x91\x8b\x00\x85\x98=\xe4\x8d2t\\9\x19\\\x8f\xfa\xc4\xe2P(M\xae\xf7x\xf7.\xce\x13A\x8c\x05˴\xb7\x1a\frڒ\x9e\xdbŖue\xa2<\xef\xd4]\x1eՆ\x06d\xa8\xa0r\xa5ݙ#~N%\x94\x97\xaa$\xa0O\xf3\x02")
			var data []byte
			data =  /*line ENqPOzTj.go:1*/ append(data, fullData[23]-fullData[54], fullData[53]^fullData[58], fullData[0]+fullData[20], fullData[5]+fullData[7], fullData[6]-fullData[16], fullData[27]^fullData[28], fullData[63]-fullData[29], fullData[25]-fullData[44], fullData[18]^fullData[15], fullData[67]^fullData[84], fullData[45]+fullData[13], fullData[83]+fullData[43], fullData[60]^fullData[92], fullData[41]^fullData[40], fullData[55]^fullData[30], fullData[9]-fullData[76], fullData[1]+fullData[65], fullData[85]-fullData[17], fullData[87]^fullData[38], fullData[8]+fullData[86], fullData[74]+fullData[31], fullData[89]+fullData[49], fullData[21]-fullData[14], fullData[59]-fullData[75], fullData[4]+fullData[52], fullData[2]+fullData[34], fullData[72]-fullData[70], fullData[64]-fullData[51], fullData[35]-fullData[3], fullData[73]+fullData[80], fullData[91]-fullData[61], fullData[93]-fullData[81], fullData[19]^fullData[36], fullData[56]+fullData[10], fullData[32]^fullData[33], fullData[88]^fullData[47], fullData[37]^fullData[79], fullData[68]+fullData[24], fullData[48]^fullData[78], fullData[39]+fullData[82], fullData[66]-fullData[26], fullData[90]+fullData[69], fullData[42]+fullData[57], fullData[50]-fullData[71], fullData[46]^fullData[11], fullData[12]^fullData[22], fullData[77]+fullData[62])
			return  /*line ENqPOzTj.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	dwbLZnck, xREbPiGc :=  /*line HuM7T4xf.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line DDDxXkni.go:1*/ shim.Error( /*line QhpLoz3B.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line p71sHMGz.go:1*/ fmt.Sprintf( /*line _Ob48u3f.go:1*/ func() string {
		key :=  /*line _Ob48u3f.go:1*/ []byte("\xc8\xf9\x91\x83]\x9c\xff\xb1\x15\xb0-\x1e'\xc0\x8c\xa4\xcb0\x9f^\xc0j\xf47\x15\xb3\xc1\xf28iqm\x96.)\xfb\xf6\xb2n8Ο\xf1OP\x8etOt\xe41\x98\xd1\xf8~K\xbd\x84\xcdZz\x1e\xb9\x14\xf5\xeb|M)\xb1&\x0e\x0eXQ\xd1A\x11ʭ\xb1\xe8G)\x8e\xfe\x8e")
		data :=  /*line _Ob48u3f.go:1*/ []byte("\xb3\xdb\xe2\xe61\xf9\x9c\xc5z\xc2\x0f$\\\xe2\xe8˨d\xe6.\xa5H\xce\x15Yܦ\x95]\x1b#\b\xe0AB\x9e\x92\xf3\r[\xab\xec\x82m|\xac\x01<\x11\x96_\xf9\xbc\x9d1-\xed\xf6\xa2,\x13z\xdcf\xd7\xd1^hZ\x93\n,j9%\xb02t\xbe\x8f\x8b\xcabZ\xac\x83\xf3")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line _Ob48u3f.go:1*/ string(data)
	}(), dwbLZnck, kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line Z_Z8X6WJ.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line TdkdkeeB.go:1*/ shim.Error( /*line bHmRA_yK.go:1*/ xREbPiGc.Error())
	}
	return  /*line I5E65zLV.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) zsYocOg1(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line cUFhHM36.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line ZqIznuhU.go:1*/ shim.Error( /*line AUemesfo.go:1*/ func() string {
			var data []byte
			i := 11
			decryptKey := 148
			for counter := 0; i != 18; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					data =  /*line AUemesfo.go:1*/ append(data, "\xa0\xb3\x96\xa6"...,
					)
					i = 8
				case 8:
					i = 15
					data =  /*line AUemesfo.go:1*/ append(data, 83)
				case 2:
					data =  /*line AUemesfo.go:1*/ append(data, "\x88\x91"...,
					)
					i = 13
				case 12:
					data =  /*line AUemesfo.go:1*/ append(data, 41)
					i = 22
				case 10:
					data =  /*line AUemesfo.go:1*/ append(data, 123)
					i = 3
				case 3:
					i = 21
					data =  /*line AUemesfo.go:1*/ append(data, "q\xa7\xa1"...,
					)
				case 17:
					i = 23
					data =  /*line AUemesfo.go:1*/ append(data, 135)
				case 11:
					data =  /*line AUemesfo.go:1*/ append(data, 98)
					i = 2
				case 13:
					i = 17
					data =  /*line AUemesfo.go:1*/ append(data, "}\x89"...,
					)
				case 19:
					i = 18
					for y := range data {
						data[y] = data[y] +  /*line AUemesfo.go:1*/ byte(decryptKey^y)
					}
				case 23:
					data =  /*line AUemesfo.go:1*/ append(data, "\x83@R"...,
					)
					i = 9
				case 9:
					i = 4
					data =  /*line AUemesfo.go:1*/ append(data, 132)
				case 5:
					data =  /*line AUemesfo.go:1*/ append(data, ">!G"...,
					)
					i = 1
				case 14:
					data =  /*line AUemesfo.go:1*/ append(data, "\x82{\x85\x8c"...,
					)
					i = 12
				case 16:
					data =  /*line AUemesfo.go:1*/ append(data, 156)
					i = 19
				case 15:
					data =  /*line AUemesfo.go:1*/ append(data, 162)
					i = 7
				case 20:
					data =  /*line AUemesfo.go:1*/ append(data, "os\x81"...,
					)
					i = 5
				case 4:
					i = 14
					data =  /*line AUemesfo.go:1*/ append(data, "z\x89"...,
					)
				case 7:
					i = 16
					data =  /*line AUemesfo.go:1*/ append(data, "\x96\xa3"...,
					)
				case 21:
					i = 0
					data =  /*line AUemesfo.go:1*/ append(data, "[\xa0\x9e"...,
					)
				case 1:
					i = 10
					data =  /*line AUemesfo.go:1*/ append(data, "{tji"...,
					)
				case 0:
					i = 6
					data =  /*line AUemesfo.go:1*/ append(data, 178)
				case 22:
					data =  /*line AUemesfo.go:1*/ append(data, "X\x80y"...,
					)
					i = 20
				}
			}
			return  /*line AUemesfo.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	nIZZng5T, xREbPiGc :=  /*line HbS6Q7YH.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line BulQge6c.go:1*/ shim.Error( /*line dhI1j53u.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line Y9vCBqMK.go:1*/ fmt.Sprintf( /*line dtxdDnZi.go:1*/ func() string {
		seed :=  /*line dtxdDnZi.go:1*/ byte(239)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line dtxdDnZi.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/  /*line dtxdDnZi.go:1*/ fnc(140)(167)(81)(242)(7)(249)(254)(17)(251)(3)(176)(24)(65)(167)(66)(11)(244)(241)(37)(247)(245)(189)(24)(232)(42)(35)(248)(0)(254)(13)(224)(19)(17)(249)(252)(250)(255)(221)(34)(0)(2)(14)(0)(175)(10)(246)(83)(254)(242)(13)(252)(243)(12)(248)(234)(23)(221)(44)(255)(5)(2)(248)(248)(13)(176)(24)(232)(3)(78)(175)(10)(246)(66)(253)(19)(237)(18)(242)(15)(174)(24)(232)(3)(78)(175)(91)(0)
		return  /*line dtxdDnZi.go:1*/ string(data)
	}(), nIZZng5T, kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line oZTo3oFU.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line CxrAZLyb.go:1*/ shim.Error( /*line hPYihMF1.go:1*/ xREbPiGc.Error())
	}
	return  /*line c_XmMiV6.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) zdj2xVZd(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	var xREbPiGc error
	var w1e7h2B3 string
	if  /*line fsWiO0U5.go:1*/ len(eFz_TaL1) != 11 {
		return  /*line Meir7m5a.go:1*/ shim.Error( /*line yiN1z613.go:1*/ func() string {
			data :=  /*line yiN1z613.go:1*/ []byte("IҐ\x84\x17d\x99K\rn^\xe9\xaau3OjH\xef\x01\xfc0og\x91YY\x1d\x88\v\xdb\xe6oW moniej\b\x14n\x1aߊ\x9d e7sF\xf5e\x00\xdfGޝs \xc9\xf9\x16eztU\xd5g\xd71\xc1\x041\xad\xf0\x90\xaa\x0e")
			positions := [...]byte{73, 66, 15, 66, 11, 45, 16, 62, 24, 3, 5, 29, 47, 78, 6, 26, 76, 7, 59, 9, 68, 73, 75, 58, 15, 12, 53, 27, 33, 24, 2, 11, 7, 15, 80, 44, 47, 75, 68, 15, 15, 14, 62, 25, 57, 28, 55, 19, 17, 26, 2, 29, 48, 1, 10, 41, 20, 79, 31, 52, 53, 59, 6, 56, 39, 40, 1, 46, 69, 9, 78, 29, 77, 29, 68, 19, 69, 48, 75, 63, 18, 64, 2, 26, 52, 55, 8, 79, 59, 14, 64, 66, 78, 6, 20, 27, 15, 21, 5, 30, 74, 1, 46, 71, 57, 50, 50, 45, 55, 4, 5, 50, 45, 16, 10, 42, 71, 8}
			for i := 0; i < 118; i += 2 {
				localKey :=  /*line yiN1z613.go:1*/ byte(i) +  /*line yiN1z613.go:1*/ byte(positions[i]^positions[i+1]) + 187
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line yiN1z613.go:1*/ string(data)
		}())
	}
	cXmPo8Kd := eFz_TaL1[0]
	dwbLZnck := eFz_TaL1[1]
	nIZZng5T := eFz_TaL1[2]
	eBmFb7v9 := eFz_TaL1[3]
	hHY_irCR := eFz_TaL1[4]
	fvUjTcZo := eFz_TaL1[5]
	xrSSF6uG := eFz_TaL1[6]
	gBYxQKrs := eFz_TaL1[7]
	g3ka7rbx :=  /*line bEOR1uX_.go:1*/ strings.Split(eFz_TaL1[8],  /*line sZ1_HJ7y.go:1*/ func() string {
		data :=  /*line sZ1_HJ7y.go:1*/ []byte("\xc7")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line sZ1_HJ7y.go:1*/ byte(i) +  /*line sZ1_HJ7y.go:1*/ byte(positions[i]^positions[i+1]) + 155
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line sZ1_HJ7y.go:1*/ string(data)
	}())
	p5F1lbum := eFz_TaL1[9]
	clszPHCk := eFz_TaL1[10]
	_0uGwnuQ :=  /*line WQtT3KME.go:1*/ func() string {
		key :=  /*line WQtT3KME.go:1*/ []byte("\x8c 7z\xd6'l\x0eh\x12\xc3Q")
		data :=  /*line WQtT3KME.go:1*/ []byte("\xc0OP\x1d\xb3U-m\vw\xb0\"")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line WQtT3KME.go:1*/ string(data)
	}()
	cKWTKAnj :=  /*line z7vEEXYc.go:1*/ func() string {
		data :=  /*line z7vEEXYc.go:1*/ []byte("gbgKerg\u007frgFsi\x89K")
		positions := [...]byte{12, 3, 14, 8, 1, 9, 13, 6, 6, 10, 3, 8, 0, 3, 10, 7, 7, 1}
		for i := 0; i < 18; i += 2 {
			localKey :=  /*line z7vEEXYc.go:1*/ byte(i) +  /*line z7vEEXYc.go:1*/ byte(positions[i]^positions[i+1]) + 241
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line z7vEEXYc.go:1*/ string(data)
	}()

	if  /*line Byaih0jO.go:1*/ len(cXmPo8Kd) <= 0 {
		return  /*line YjY7z6le.go:1*/ shim.Error( /*line PwhLECQ2.go:1*/ func() string {
			seed :=  /*line PwhLECQ2.go:1*/ byte(191)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line PwhLECQ2.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/  /*line PwhLECQ2.go:1*/ fnc(251)(219)(225)(23)(254)(238)(13)(166)(65)(24)(246)(15)(170)(86)(239)(89)(179)(165)(68)(1)(1)(93)(168)(24)(253)(254)(241)(89)(161)(7)(8)(235)(3)(23)
			return  /*line PwhLECQ2.go:1*/ string(data)
		}())
	}

	if  /*line Hd0vBhiO.go:1*/ len(eBmFb7v9) <= 0 {
		return  /*line C0A0Zosd.go:1*/ shim.Error( /*line FU5yo_8e.go:1*/ func() string {
			var data []byte
			i := 12
			decryptKey := 251
			for counter := 0; i != 18; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 15:
					i = 18
					for y := range data {
						data[y] = data[y] -  /*line FU5yo_8e.go:1*/ byte(decryptKey^y)
					}
				case 2:
					i = 5
					data =  /*line FU5yo_8e.go:1*/ append(data, 166)
				case 11:
					data =  /*line FU5yo_8e.go:1*/ append(data, 185)
					i = 10
				case 10:
					data =  /*line FU5yo_8e.go:1*/ append(data, "u\xae\xbb"...,
					)
					i = 4
				case 5:
					data =  /*line FU5yo_8e.go:1*/ append(data, 102)
					i = 8
				case 4:
					i = 9
					data =  /*line FU5yo_8e.go:1*/ append(data, "\xbf\xc0ƒ"...,
					)
				case 13:
					data =  /*line FU5yo_8e.go:1*/ append(data, 196)
					i = 17
				case 16:
					data =  /*line FU5yo_8e.go:1*/ append(data, "\xc0ζ\xce"...,
					)
					i = 13
				case 14:
					data =  /*line FU5yo_8e.go:1*/ append(data, "\xb7\xb2\xc4"...,
					)
					i = 16
				case 0:
					data =  /*line FU5yo_8e.go:1*/ append(data, "~\xcc"...,
					)
					i = 7
				case 7:
					i = 6
					data =  /*line FU5yo_8e.go:1*/ append(data, "\xd1\xd0"...,
					)
				case 6:
					data =  /*line FU5yo_8e.go:1*/ append(data, "\xb6c\xa2"...,
					)
					i = 2
				case 8:
					i = 11
					data =  /*line FU5yo_8e.go:1*/ append(data, "\xa8d\xb3\xb9"...,
					)
				case 12:
					i = 14
					data =  /*line FU5yo_8e.go:1*/ append(data, "\xa1\xc5"...,
					)
				case 3:
					i = 1
					data =  /*line FU5yo_8e.go:1*/ append(data, 223)
				case 1:
					data =  /*line FU5yo_8e.go:1*/ append(data, "\xe5\xdb"...,
					)
					i = 15
				case 17:
					data =  /*line FU5yo_8e.go:1*/ append(data, "\xc7\xc7"...,
					)
					i = 0
				case 9:
					data =  /*line FU5yo_8e.go:1*/ append(data, "\xe6\xe4\xe3"...,
					)
					i = 3
				}
			}
			return  /*line FU5yo_8e.go:1*/ string(data)
		}())
	}

	if  /*line JkHsCeMR.go:1*/ len(dwbLZnck) <= 0 {
		return  /*line Cw1bjJHe.go:1*/ shim.Error( /*line fUyTa_bw.go:1*/ func() string {
			var data []byte
			i := 11
			decryptKey := 188
			for counter := 0; i != 7; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 14:
					i = 7
					for y := range data {
						data[y] = data[y] -  /*line fUyTa_bw.go:1*/ byte(decryptKey^y)
					}
				case 0:
					i = 18
					data =  /*line fUyTa_bw.go:1*/ append(data, "\xd0\xd4\xe0\x89"...,
					)
				case 8:
					data =  /*line fUyTa_bw.go:1*/ append(data, 186)
					i = 14
				case 16:
					i = 0
					data =  /*line fUyTa_bw.go:1*/ append(data, "\xe8\xd6"...,
					)
				case 12:
					i = 1
					data =  /*line fUyTa_bw.go:1*/ append(data, "\xc3\xc6"...,
					)
				case 11:
					i = 10
					data =  /*line fUyTa_bw.go:1*/ append(data, "\xd2\xef\xe4\xf0"...,
					)
				case 9:
					i = 3
					data =  /*line fUyTa_bw.go:1*/ append(data, "ߕ\xe3\xdd"...,
					)
				case 5:
					i = 15
					data =  /*line fUyTa_bw.go:1*/ append(data, 224)
				case 2:
					data =  /*line fUyTa_bw.go:1*/ append(data, 226)
					i = 16
				case 6:
					data =  /*line fUyTa_bw.go:1*/ append(data, "\xc7\xcb\xc8"...,
					)
					i = 13
				case 4:
					data =  /*line fUyTa_bw.go:1*/ append(data, "\x84\xc9"...,
					)
					i = 19
				case 18:
					i = 5
					data =  /*line fUyTa_bw.go:1*/ append(data, 213)
				case 19:
					data =  /*line fUyTa_bw.go:1*/ append(data, "ˁ\xc1\x83"...,
					)
					i = 17
				case 17:
					data =  /*line fUyTa_bw.go:1*/ append(data, "\xd0\xccʌ"...,
					)
					i = 12
				case 15:
					data =  /*line fUyTa_bw.go:1*/ append(data, "\xdd\xd9"...,
					)
					i = 4
				case 1:
					i = 6
					data =  /*line fUyTa_bw.go:1*/ append(data, "\xc8\xcf\xd3u"...,
					)
				case 13:
					data =  /*line fUyTa_bw.go:1*/ append(data, "\xba\xbe"...,
					)
					i = 8
				case 10:
					data =  /*line fUyTa_bw.go:1*/ append(data, "\xe7\xd9\xe8"...,
					)
					i = 9
				case 3:
					data =  /*line fUyTa_bw.go:1*/ append(data, "\x96\xe1\xe2"...,
					)
					i = 2
				}
			}
			return  /*line fUyTa_bw.go:1*/ string(data)
		}())
	}

	if  /*line ldzCrxJh.go:1*/ len(nIZZng5T) <= 0 {
		return  /*line e3GgOMaA.go:1*/ shim.Error( /*line OFiz6nzm.go:1*/ func() string {
			key :=  /*line OFiz6nzm.go:1*/ []byte("\x85\xfdL\u007f{S\xfa%*\xad\xcc\x11\xdd>\x9aH\xb8[\x88\x8dJ\x9bz!\xa9\xa7\x88\x87\x14\xd1M\xac\x9d\x15\x99=\x0fd\xcbP\xa23'\xac\xcc/\xaf")
			data :=  /*line OFiz6nzm.go:1*/ []byte("\xdap\xb1\xf1\xe9\xb4g\x8aJ\x1c21@\xad\b\xbb-\xc8\xed\xffj\b\xef\x94\x1d\xc7\xea\xec42m\x1a\f\x83Ƣ|\xd4?\xc9¦\x9b\x1e5\x9d\x16")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line OFiz6nzm.go:1*/ string(data)
		}())
	}

	if  /*line RkAzARWx.go:1*/ len(fvUjTcZo) <= 0 {
		return  /*line Kuh0fvaI.go:1*/ shim.Error( /*line IDc_NhL1.go:1*/ func() string {
			fullData :=  /*line IDc_NhL1.go:1*/ []byte("\xfb\xd3\xedh\xb9\xebI\x97\x0f \xef\x18\u007f\xcf\vSͻ\x05\xf7A\xb1y~\xa8R\xd9\xd1~\xa1\xe3\xcdܼ\xbbx\n\x98\x87\xe8f\xe9s\xa65\xb3\x18V\xd5;h&`\xa06\b>\x13K\xbb\x94ӛ\x13e\xe1")
			var data []byte
			data =  /*line IDc_NhL1.go:1*/ append(data, fullData[17]-fullData[50], fullData[0]-fullData[38], fullData[8]+fullData[25], fullData[14]-fullData[7], fullData[58]^fullData[56], fullData[53]+fullData[61], fullData[13]^fullData[10], fullData[47]^fullData[49], fullData[12]-fullData[36], fullData[62]^fullData[39], fullData[4]+fullData[59], fullData[24]+fullData[35], fullData[34]^fullData[26], fullData[20]-fullData[32], fullData[27]-fullData[21], fullData[22]-fullData[46], fullData[5]+fullData[44], fullData[64]-fullData[19], fullData[33]+fullData[45], fullData[42]-fullData[18], fullData[15]-fullData[51], fullData[30]-fullData[28], fullData[63]-fullData[43], fullData[3]^fullData[11], fullData[6]-fullData[48], fullData[37]+fullData[65], fullData[2]-fullData[16], fullData[57]+fullData[52], fullData[60]-fullData[9], fullData[29]^fullData[1], fullData[54]-fullData[31], fullData[40]+fullData[55], fullData[23]+fullData[41])
			return  /*line IDc_NhL1.go:1*/ string(data)
		}())
	}

	cR3vs4Ox :=  /*line OePswuhR.go:1*/ func() string {
		var data []byte
		i := 3
		decryptKey := 17
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 3:
				data =  /*line OePswuhR.go:1*/ append(data, "[\\_`"...,
				)
				i = 0
			case 1:
				data =  /*line OePswuhR.go:1*/ append(data, "8O"...,
				)
				i = 4
			case 4:
				for y := range data {
					data[y] = data[y] +  /*line OePswuhR.go:1*/ byte(decryptKey^y)
				}
				i = 5
			case 0:
				i = 2
				data =  /*line OePswuhR.go:1*/ append(data, "qpU"...,
				)
			case 2:
				i = 1
				data =  /*line OePswuhR.go:1*/ append(data, "rWc?"...,
				)
			}
		}
		return  /*line OePswuhR.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line mhBr_u2V.go:1*/ func() string {
		data :=  /*line mhBr_u2V.go:1*/ []byte("\xca")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line mhBr_u2V.go:1*/ byte(i) +  /*line mhBr_u2V.go:1*/ byte(positions[i]^positions[i+1]) + 155
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line mhBr_u2V.go:1*/ string(data)
	}() + dwbLZnck +  /*line xUIFYbBx.go:1*/ func() string {
		fullData :=  /*line xUIFYbBx.go:1*/ []byte("\xc1\xf0")
		var data []byte
		data =  /*line xUIFYbBx.go:1*/ append(data, fullData[1]-fullData[0])
		return  /*line xUIFYbBx.go:1*/ string(data)
	}() + nIZZng5T
	S8BnBXR2 := LoggerAccessCounter{}
	sdHID17m := -1
	nZKTLZaR, xREbPiGc :=  /*line f5MozSkA.go:1*/ fOOTaXmK.GetState(cR3vs4Ox)	                                             
	if xREbPiGc != nil {
		w1e7h2B3 =  /*line Gw9npNlt.go:1*/ func() string {
			var data []byte
			i := 19
			decryptKey := 42
			for counter := 0; i != 13; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 20:
					data =  /*line Gw9npNlt.go:1*/ append(data, 20)
					i = 10
				case 19:
					data =  /*line Gw9npNlt.go:1*/ append(data, 1)
					i = 11
				case 8:
					data =  /*line Gw9npNlt.go:1*/ append(data, 14)
					i = 7
				case 18:
					i = 0
					data =  /*line Gw9npNlt.go:1*/ append(data, "\xff\x02\xff"...,
					)
				case 15:
					i = 14
					data =  /*line Gw9npNlt.go:1*/ append(data, "\xf2\xf2"...,
					)
				case 5:
					i = 1
					data =  /*line Gw9npNlt.go:1*/ append(data, "!\xcc"...,
					)
				case 7:
					data =  /*line Gw9npNlt.go:1*/ append(data, "\xb8\x01"...,
					)
					i = 21
				case 16:
					i = 6
					data =  /*line Gw9npNlt.go:1*/ append(data, 243)
				case 6:
					data =  /*line Gw9npNlt.go:1*/ append(data, "\xf7\xed\xed"...,
					)
					i = 3
				case 9:
					i = 18
					data =  /*line Gw9npNlt.go:1*/ append(data, "\xf9\xf8\x04\xb1"...,
					)
				case 17:
					data =  /*line Gw9npNlt.go:1*/ append(data, "\x1c\xc0\a\x1d"...,
					)
					i = 5
				case 10:
					i = 17
					data =  /*line Gw9npNlt.go:1*/ append(data, 20)
				case 22:
					i = 9
					data =  /*line Gw9npNlt.go:1*/ append(data, "\x03\xb5"...,
					)
				case 21:
					i = 20
					data =  /*line Gw9npNlt.go:1*/ append(data, "\x0f\x1a\x18"...,
					)
				case 0:
					data =  /*line Gw9npNlt.go:1*/ append(data, "\x02\r"...,
					)
					i = 8
				case 3:
					i = 4
					data =  /*line Gw9npNlt.go:1*/ append(data, 182)
				case 14:
					data =  /*line Gw9npNlt.go:1*/ append(data, "\xa3ȱ\xd2"...,
					)
					i = 12
				case 11:
					i = 2
					data =  /*line Gw9npNlt.go:1*/ append(data, "\xa9\xc9"...,
					)
				case 12:
					i = 16
					data =  /*line Gw9npNlt.go:1*/ append(data, 238)
				case 1:
					i = 13
					for y := range data {
						data[y] = data[y] -  /*line Gw9npNlt.go:1*/ byte(decryptKey^y)
					}
				case 2:
					data =  /*line Gw9npNlt.go:1*/ append(data, "\xf7\xf4"...,
					)
					i = 15
				case 4:
					data =  /*line Gw9npNlt.go:1*/ append(data, 11)
					i = 22
				}
			}
			return  /*line Gw9npNlt.go:1*/ string(data)
		}() + cR3vs4Ox +  /*line Lb33y_HE.go:1*/ func() string {
			key :=  /*line Lb33y_HE.go:1*/ []byte("\x94\x9d")
			data :=  /*line Lb33y_HE.go:1*/ []byte("\xb6\x1a")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line Lb33y_HE.go:1*/ string(data)
		}()
		return  /*line JwwZlr1O.go:1*/ shim.Error(w1e7h2B3)
	} else if nZKTLZaR == nil {
		sdHID17m = 0
		dGu9RrHS := &LoggerAccessCounter{cKWTKAnj, cXmPo8Kd, sdHID17m}
		nZKTLZaR, xREbPiGc =  /*line Ff5LzDlB.go:1*/ json.Marshal(dGu9RrHS)
		if xREbPiGc != nil {
			return  /*line HuYeoj86.go:1*/ shim.Error( /*line v3AhiqSz.go:1*/ xREbPiGc.Error())
		}

	} else {
		xREbPiGc =  /*line t_dm4nSH.go:1*/ json.Unmarshal(nZKTLZaR, &S8BnBXR2)	                               
		if xREbPiGc != nil {
			return  /*line ro_wThFY.go:1*/ shim.Error( /*line LSOQlaGj.go:1*/ xREbPiGc.Error())
		}
		sdHID17m = S8BnBXR2.Count + 1
		S8BnBXR2.Count = sdHID17m
		nZKTLZaR, _ =  /*line HsXDzQCE.go:1*/ json.Marshal(S8BnBXR2)
	}

	                                       
	xREbPiGc =  /*line GZzMhyAg.go:1*/ fOOTaXmK.PutState(cR3vs4Ox, nZKTLZaR)
	if xREbPiGc != nil {
		return  /*line Ozf0m_lh.go:1*/ shim.Error( /*line JCYTO7fa.go:1*/ xREbPiGc.Error())
	}

	BogxjCz3 := &LoggerAccess{_0uGwnuQ, cXmPo8Kd, dwbLZnck, p5F1lbum, clszPHCk, nIZZng5T, eBmFb7v9, xrSSF6uG, gBYxQKrs, g3ka7rbx, fvUjTcZo, hHY_irCR, sdHID17m}
	EugKyXCJ, xREbPiGc :=  /*line OuZWlT8S.go:1*/ json.Marshal(BogxjCz3)
	if xREbPiGc != nil {
		return  /*line uejnqa_1.go:1*/ shim.Error( /*line RP6HI2n2.go:1*/ xREbPiGc.Error())
	}

	iPrzUCH_ :=  /*line MFQz7GyD.go:1*/ func() string {
		key :=  /*line MFQz7GyD.go:1*/ []byte("Ҕ*\xcfY\xae\xfcR\\\t")
		data :=  /*line MFQz7GyD.go:1*/ []byte("3\xf7\x8d4\xcc!Q\xc5\xc1{")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line MFQz7GyD.go:1*/ string(data)
	}() + cXmPo8Kd +  /*line x91UwBfA.go:1*/ func() string {
		seed :=  /*line x91UwBfA.go:1*/ byte(33)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line x91UwBfA.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line x91UwBfA.go:1*/ fnc(14)
		return  /*line x91UwBfA.go:1*/ string(data)
	}() + dwbLZnck +  /*line jpJ95rOv.go:1*/ func() string {
		data :=  /*line jpJ95rOv.go:1*/ []byte("\xf7")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line jpJ95rOv.go:1*/ byte(i) +  /*line jpJ95rOv.go:1*/ byte(positions[i]^positions[i+1]) + 216
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line jpJ95rOv.go:1*/ string(data)
	}() + nIZZng5T +  /*line Q9O_QGNP.go:1*/ strconv.Itoa(sdHID17m)
	                        
	xREbPiGc =  /*line FaXcPwS1.go:1*/ fOOTaXmK.PutState(iPrzUCH_, EugKyXCJ)
	if xREbPiGc != nil {
		return  /*line TC9bMWvl.go:1*/ shim.Error( /*line cFWX0Rm8.go:1*/ xREbPiGc.Error())
	}

	auAEXfA1 :=  /*line z9dipY3x.go:1*/ func() string {
		seed :=  /*line z9dipY3x.go:1*/ byte(218)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line z9dipY3x.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/  /*line z9dipY3x.go:1*/ fnc(29)(99)(191)(122)(5)(1)(7)(7)(199)(182)(147)(44)(86)(158)(247)(67)(132)(250)(1)(176)(161)(68)(136)(18)(50)(100)(117)(48)(105)(213)(88)(244)(229)(221)(167)(96)(178)(115)(172)(62)
		return  /*line z9dipY3x.go:1*/ string(data)
	}() + cXmPo8Kd
	pPXX7o_t :=  /*line SdY3zkoc.go:1*/ []byte(auAEXfA1)
	return  /*line dAMAr3iF.go:1*/ shim.Success(pPXX7o_t)

}

func (hslgON00 *EMhVAN9S) ctiri1Q3(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	dwbLZnck, xREbPiGc :=  /*line xyx50H1q.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line SLLvN58k.go:1*/ shim.Error( /*line wJiIUkKP.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line DTAgOKLP.go:1*/ fmt.Sprintf( /*line m_AH7Dbz.go:1*/ func() string {
		var data []byte
		i := 20
		decryptKey := 82
		for counter := 0; i != 21; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 14
				data =  /*line m_AH7Dbz.go:1*/ append(data, 27)
			case 7:
				data =  /*line m_AH7Dbz.go:1*/ append(data, 118)
				i = 16
			case 9:
				i = 15
				data =  /*line m_AH7Dbz.go:1*/ append(data, "\x17Q\tL"...,
				)
			case 18:
				data =  /*line m_AH7Dbz.go:1*/ append(data, "/\"t}"...,
				)
				i = 12
			case 16:
				data =  /*line m_AH7Dbz.go:1*/ append(data, 32)
				i = 18
			case 11:
				i = 6
				data =  /*line m_AH7Dbz.go:1*/ append(data, "h:d"...,
				)
			case 3:
				i = 5
				data =  /*line m_AH7Dbz.go:1*/ append(data, "P[Z"...,
				)
			case 1:
				i = 4
				data =  /*line m_AH7Dbz.go:1*/ append(data, "ztt"...,
				)
			case 4:
				i = 24
				data =  /*line m_AH7Dbz.go:1*/ append(data, "l=&"...,
				)
			case 22:
				data =  /*line m_AH7Dbz.go:1*/ append(data, "MEW\x11"...,
				)
				i = 19
			case 10:
				i = 1
				data =  /*line m_AH7Dbz.go:1*/ append(data, "Gfzd"...,
				)
			case 13:
				i = 7
				data =  /*line m_AH7Dbz.go:1*/ append(data, "ebw"...,
				)
			case 15:
				i = 22
				data =  /*line m_AH7Dbz.go:1*/ append(data, "FUc"...,
				)
			case 0:
				data =  /*line m_AH7Dbz.go:1*/ append(data, "NF"...,
				)
				i = 23
			case 2:
				i = 9
				data =  /*line m_AH7Dbz.go:1*/ append(data, "UA]\x0e"...,
				)
			case 24:
				i = 11
				data =  /*line m_AH7Dbz.go:1*/ append(data, "??"...,
				)
			case 14:
				for y := range data {
					data[y] = data[y] ^  /*line m_AH7Dbz.go:1*/ byte(decryptKey^y)
				}
				i = 21
			case 5:
				data =  /*line m_AH7Dbz.go:1*/ append(data, "_IyZ"...,
				)
				i = 13
			case 17:
				data =  /*line m_AH7Dbz.go:1*/ append(data, "ckf"...,
				)
				i = 8
			case 20:
				data =  /*line m_AH7Dbz.go:1*/ append(data, "]\x05W@"...,
				)
				i = 0
			case 23:
				data =  /*line m_AH7Dbz.go:1*/ append(data, 67)
				i = 2
			case 12:
				data =  /*line m_AH7Dbz.go:1*/ append(data, "j~"...,
				)
				i = 17
			case 19:
				data =  /*line m_AH7Dbz.go:1*/ append(data, "\n\x13r"...,
				)
				i = 3
			case 8:
				data =  /*line m_AH7Dbz.go:1*/ append(data, "mFp"...,
				)
				i = 10
			}
		}
		return  /*line m_AH7Dbz.go:1*/ string(data)
	}(), dwbLZnck)
	toHNB1kz, xREbPiGc :=  /*line T0ntTfM6.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line jFPCJ7Nv.go:1*/ shim.Error( /*line CucC3_Wt.go:1*/ xREbPiGc.Error())
	}
	return  /*line WPEAm_Cu.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) eEC45Wzm(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	nIZZng5T, xREbPiGc :=  /*line vdgvHdo4.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line huVflZU4.go:1*/ shim.Error( /*line iZe_Dgiy.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line YRRZfjEm.go:1*/ fmt.Sprintf( /*line iepGAbhp.go:1*/ func() string {
		fullData :=  /*line iepGAbhp.go:1*/ []byte("\xd9\xd8c\xba\xa3\vHԎ\xe8Ow\x8b\x8e;QŦr0\xd57[Y\x19\xf8\x14D\xac\xfa\x92\x8aGì\xf8\x0f\xb2H\x95%\xb6)\xc3ѿ\r\xb0Ea8\xb7\x1c\xe9\xc5\xfa\xd5i >\x1fڒR\xea\xf5Î\xee`]\xba\xaa\xb3R/\xd9\xd2KP7\xe4\x95\x1f\x9d\xd2\x10Ya>gc=\xfd\x1d\xec\xce\x03\x86\x12\x93w\x87$d\r\x03r\x01\rn\xa9\x9e׃.\xab9\xe4\x14\xa1\xaa\"\x95J\t\x84 b\x89")
		var data []byte
		data =  /*line iepGAbhp.go:1*/ append(data, fullData[126]-fullData[125], fullData[12]-fullData[57], fullData[31]+fullData[53], fullData[123]-fullData[19], fullData[115]+fullData[89], fullData[113]^fullData[37], fullData[121]-fullData[32], fullData[122]+fullData[74], fullData[105]-fullData[112], fullData[65]-fullData[114], fullData[14]^fullData[24], fullData[117]^fullData[106], fullData[111]+fullData[85], fullData[60]-fullData[93], fullData[108]+fullData[2], fullData[61]+fullData[39], fullData[41]^fullData[56], fullData[87]^fullData[109], fullData[3]^fullData[33], fullData[26]^fullData[104], fullData[96]^fullData[116], fullData[34]^fullData[8], fullData[28]+fullData[13], fullData[86]+fullData[99], fullData[62]+fullData[71], fullData[22]-fullData[95], fullData[88]-fullData[29], fullData[49]-fullData[55], fullData[25]^fullData[84], fullData[81]+fullData[67], fullData[129]-fullData[38], fullData[101]-fullData[119], fullData[118]^fullData[102], fullData[45]+fullData[17], fullData[47]^fullData[43], fullData[27]+fullData[75], fullData[54]-fullData[4], fullData[42]+fullData[97], fullData[76]-fullData[51], fullData[70]-fullData[9], fullData[50]^fullData[78], fullData[16]-fullData[69], fullData[35]-fullData[98], fullData[59]^fullData[79], fullData[64]+fullData[11], fullData[82]+fullData[1], fullData[48]^fullData[58], fullData[63]^fullData[94], fullData[7]+fullData[30], fullData[91]^fullData[127], fullData[46]+fullData[128], fullData[21]+fullData[80], fullData[103]+fullData[10], fullData[23]+fullData[52], fullData[66]+fullData[72], fullData[5]^fullData[110], fullData[0]-fullData[90], fullData[15]+fullData[44], fullData[6]^fullData[107], fullData[92]^fullData[83], fullData[124]-fullData[40], fullData[120]+fullData[77], fullData[20]-fullData[73], fullData[36]^fullData[18], fullData[100]^fullData[68])
		return  /*line iepGAbhp.go:1*/ string(data)
	}(), nIZZng5T)
	toHNB1kz, xREbPiGc :=  /*line MGGB00Fj.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line d0DD9tav.go:1*/ shim.Error( /*line l9PKLIzC.go:1*/ xREbPiGc.Error())
	}
	return  /*line McuvHKzG.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) fNGkK4OH(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line kFC5M1WE.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line GhZ31nDG.go:1*/ shim.Error( /*line zVblCMfO.go:1*/ func() string {
			fullData :=  /*line zVblCMfO.go:1*/ []byte("\xfa\xa4\x8eZ\x16\x92Qv\xec>8$\xb0\x80\xe9*Z\xb7\x03\xa8\xa3Q\x83\x8c\u007f/\xbf\xf1*V\xb3\xafQ\x8ak\xe2\xcad4\xac\xf6\xf0\xacP\x93B\xeeT\\!\xbbI\xcd\xf1i\xe2\xfd\uf16c\xec(\xb5\x84\xb7\xc875\xf8t\xddn\xedK4\xb6\xc8\r\x17\x96\xde\x10DS\xa5\xe1YE\x1d ,$;\xd6")
			var data []byte
			data =  /*line zVblCMfO.go:1*/ append(data, fullData[46]-fullData[84], fullData[32]+fullData[88], fullData[8]+fullData[33], fullData[47]^fullData[67], fullData[65]+fullData[1], fullData[93]^fullData[26], fullData[40]+fullData[71], fullData[75]^fullData[79], fullData[43]+fullData[27], fullData[17]-fullData[87], fullData[35]^fullData[58], fullData[18]-fullData[2], fullData[81]-fullData[20], fullData[12]+fullData[62], fullData[64]-fullData[51], fullData[36]-fullData[29], fullData[61]+fullData[68], fullData[44]+fullData[50], fullData[28]+fullData[73], fullData[89]-fullData[30], fullData[9]+fullData[11], fullData[49]+fullData[82], fullData[80]^fullData[42], fullData[45]+fullData[60], fullData[69]+fullData[39], fullData[90]^fullData[54], fullData[91]-fullData[59], fullData[53]+fullData[24], fullData[6]^fullData[74], fullData[57]-fullData[23], fullData[21]-fullData[70], fullData[14]+fullData[13], fullData[4]-fullData[19], fullData[76]^fullData[31], fullData[52]+fullData[83], fullData[85]+fullData[22], fullData[56]+fullData[37], fullData[41]+fullData[63], fullData[86]^fullData[10], fullData[48]^fullData[25], fullData[0]+fullData[34], fullData[5]+fullData[55], fullData[77]-fullData[72], fullData[38]^fullData[3], fullData[7]^fullData[78], fullData[66]^fullData[16], fullData[92]+fullData[15])
			return  /*line zVblCMfO.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	dwbLZnck, xREbPiGc :=  /*line DAxzB4nd.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line zjloHpDg.go:1*/ shim.Error( /*line GoVD0pQA.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line zozdK4oF.go:1*/ fmt.Sprintf( /*line wODy6M28.go:1*/ func() string {
		data :=  /*line wODy6M28.go:1*/ []byte("{\"selectږ\"~1na\xb5I{y\x9d\xe2\x91\x15`Lo\x11\x03\x8crA\xdc\x02\xaa\x9e8\x04,\"\xdc:\x15r\ra\x9f/\x82O\xaer\xa8\ai\xbb,\x16yS\x8b\xa5\x1d\"\xd4\x10\xd8a\xfe\xc7s\xe3\x8a\":\x1d\x91\xba\xbf\xf7}")
		positions := [...]byte{15, 64, 46, 12, 77, 48, 78, 16, 34, 20, 17, 22, 61, 64, 31, 48, 40, 55, 45, 19, 59, 70, 67, 8, 77, 52, 34, 35, 65, 43, 56, 27, 63, 11, 51, 68, 52, 76, 11, 14, 39, 70, 32, 51, 33, 36, 14, 19, 67, 74, 13, 57, 49, 71, 76, 60, 54, 65, 39, 64, 23, 48, 58, 57, 77, 9, 48, 20, 41, 45, 28, 36, 26, 63, 35, 75, 21, 58, 13, 47, 78, 63}
		for i := 0; i < 82; i += 2 {
			localKey :=  /*line wODy6M28.go:1*/ byte(i) +  /*line wODy6M28.go:1*/ byte(positions[i]^positions[i+1]) + 48
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line wODy6M28.go:1*/ string(data)
	}(), dwbLZnck, kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line Inont5_7.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line DZ82CHa0.go:1*/ shim.Error( /*line IpBmw3bU.go:1*/ xREbPiGc.Error())
	}
	return  /*line elABzDJI.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) uSPEGxLQ(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line mpbm0Q05.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line ACftckq7.go:1*/ shim.Error( /*line HpB0QSpH.go:1*/ func() string {
			seed :=  /*line HpB0QSpH.go:1*/ byte(128)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line HpB0QSpH.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/  /*line HpB0QSpH.go:1*/ fnc(201)(37)(8)(235)(11)(253)(251)(188)(33)(49)(245)(14)(248)(248)(9)(6)(172)(46)(39)(248)(245)(3)(13)(188)(242)(37)(51)(248)(245)(254)(17)(245)(5)(249)(185)(68)(253)(19)(237)(18)(242)(15)(172)(78)(243)(12)(248)
			return  /*line HpB0QSpH.go:1*/ string(data)
		}())
	}
	kiJzWA_D := eFz_TaL1[0]
	nIZZng5T, xREbPiGc :=  /*line G1AsEyyk.go:1*/ hslgON00.raycgBaa(fOOTaXmK)
	if xREbPiGc != nil {
		return  /*line sf3SJVhz.go:1*/ shim.Error( /*line afnyIs7B.go:1*/ xREbPiGc.Error())
	}
	                  
	fcEHFI_2 :=  /*line Zlh2VJQq.go:1*/ fmt.Sprintf( /*line Jlzy65TX.go:1*/ func() string {
		seed :=  /*line Jlzy65TX.go:1*/ byte(154)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Jlzy65TX.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/  /*line Jlzy65TX.go:1*/ fnc(225)(167)(81)(242)(7)(249)(254)(17)(251)(3)(176)(24)(65)(167)(66)(11)(244)(241)(37)(247)(245)(189)(24)(232)(42)(35)(248)(0)(254)(13)(207)(34)(0)(2)(14)(0)(175)(10)(246)(83)(254)(242)(13)(252)(243)(12)(248)(234)(23)(221)(44)(255)(5)(2)(248)(248)(13)(176)(24)(232)(3)(78)(175)(10)(246)(66)(253)(19)(237)(18)(242)(15)(174)(24)(232)(3)(78)(175)(91)(0)
		return  /*line Jlzy65TX.go:1*/ string(data)
	}(), nIZZng5T, kiJzWA_D)
	toHNB1kz, xREbPiGc :=  /*line NsDVrtVi.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line q9IPkftK.go:1*/ shim.Error( /*line Ek3S2Dxz.go:1*/ xREbPiGc.Error())
	}
	return  /*line QZiYmFgm.go:1*/ shim.Success(toHNB1kz)
}
