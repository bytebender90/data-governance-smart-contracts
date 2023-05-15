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
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	pb "github.com/hyperledger/fabric/protos/peer"
)

               
func woUuOwcA(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) (string, error) {
	if  /*line IyTFz7tE.go:1*/ len(y5yvGZQA) != 7 {
		return "",  /*line cpdYfYA7.go:1*/ fmt.Errorf( /*line qoOdZIjo.go:1*/ func() string {
			var data []byte
			i := 22
			decryptKey := 206
			for counter := 0; i != 7; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 17:
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xd3\xd7"...,
					)
					i = 26
				case 29:
					i = 13
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xbb\xcb"...,
					)
				case 15:
					data =  /*line qoOdZIjo.go:1*/ append(data, "zk"...,
					)
					i = 30
				case 2:
					i = 25
					data =  /*line qoOdZIjo.go:1*/ append(data, "ƺǾ"...,
					)
				case 6:
					data =  /*line qoOdZIjo.go:1*/ append(data, 228)
					i = 27
				case 31:
					data =  /*line qoOdZIjo.go:1*/ append(data, "\x9d\x85\xda"...,
					)
					i = 11
				case 23:
					data =  /*line qoOdZIjo.go:1*/ append(data, "ޘ\xc3"...,
					)
					i = 21
				case 14:
					for y := range data {
						data[y] = data[y] -  /*line qoOdZIjo.go:1*/ byte(decryptKey^y)
					}
					i = 7
				case 16:
					data =  /*line qoOdZIjo.go:1*/ append(data, "h\xb6\xb8\xa9"...,
					)
					i = 28
				case 28:
					i = 5
					data =  /*line qoOdZIjo.go:1*/ append(data, 169)
				case 22:
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xa8\xcc\xc0\xcb"...,
					)
					i = 19
				case 1:
					data =  /*line qoOdZIjo.go:1*/ append(data, 146)
					i = 14
				case 26:
					data =  /*line qoOdZIjo.go:1*/ append(data, "χ"...,
					)
					i = 31
				case 0:
					i = 9
					data =  /*line qoOdZIjo.go:1*/ append(data, 204)
				case 21:
					i = 24
					data =  /*line qoOdZIjo.go:1*/ append(data, 229)
				case 4:
					i = 17
					data =  /*line qoOdZIjo.go:1*/ append(data, 223)
				case 9:
					i = 29
					data =  /*line qoOdZIjo.go:1*/ append(data, 190)
				case 18:
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xb5\xaf\x9f\xe7"...,
					)
					i = 10
				case 32:
					i = 18
					data =  /*line qoOdZIjo.go:1*/ append(data, 98)
				case 11:
					i = 1
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xc4\xce\xd6\xc5"...,
					)
				case 12:
					data =  /*line qoOdZIjo.go:1*/ append(data, 182)
					i = 2
				case 20:
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xd5\xe7"...,
					)
					i = 3
				case 27:
					i = 20
					data =  /*line qoOdZIjo.go:1*/ append(data, 145)
				case 24:
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xdc\xdb\xd8"...,
					)
					i = 6
				case 25:
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xb5\xbd\xc2\xc0"...,
					)
					i = 15
				case 10:
					i = 8
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xeb\xf2"...,
					)
				case 5:
					i = 32
					data =  /*line qoOdZIjo.go:1*/ append(data, 181)
				case 19:
					i = 0
					data =  /*line qoOdZIjo.go:1*/ append(data, 205)
				case 13:
					i = 12
					data =  /*line qoOdZIjo.go:1*/ append(data, 118)
				case 3:
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xde\xd2\xcf"...,
					)
					i = 4
				case 30:
					i = 16
					data =  /*line qoOdZIjo.go:1*/ append(data, "\x93\xb7"...,
					)
				case 8:
					i = 23
					data =  /*line qoOdZIjo.go:1*/ append(data, "\xea\xe5"...,
					)
				}
			}
			return  /*line qoOdZIjo.go:1*/ string(data)
		}())
	}
	fdBHw0R_ := y5yvGZQA[0]
	taAS4TM5 := y5yvGZQA[1]
	tu6TYtHu := y5yvGZQA[2]
	tvV5RxQz := y5yvGZQA[3]
	zrdHgEga := y5yvGZQA[4]
	zkXVyJDc := y5yvGZQA[5]
	mP0iLz9a := y5yvGZQA[6]

	bAMSTiaS := []string{ /*line rVPLZRrp.go:1*/ func() string {
		data :=  /*line rVPLZRrp.go:1*/ []byte("l\xd8lo\xcdʵ\xc8pl\xcdaxM\x1c\xdbld\xd3\xcdg")
		positions := [...]byte{18, 14, 0, 0, 7, 0, 6, 15, 19, 4, 5, 12, 1, 18, 16, 1, 20, 3, 2, 15, 12, 15, 20, 3, 2, 16, 10, 15, 1, 5}
		for i := 0; i < 30; i += 2 {
			localKey :=  /*line rVPLZRrp.go:1*/ byte(i) +  /*line rVPLZRrp.go:1*/ byte(positions[i]^positions[i+1]) + 154
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line rVPLZRrp.go:1*/ string(data)
	}(), fdBHw0R_, taAS4TM5, tu6TYtHu, tvV5RxQz, zrdHgEga, zkXVyJDc, mP0iLz9a}
	bthOyi2E :=  /*line mxKGkMHM.go:1*/ make([][]byte,  /*line QyvBzjQm.go:1*/ len(bAMSTiaS))
	for cBTQZ_Tb, gCsvJoOL := range bAMSTiaS {
		bthOyi2E[cBTQZ_Tb] =  /*line BszlmGFu.go:1*/ []byte(gCsvJoOL)
	}

	k9jvpio9 :=  /*line EBj2mNmM.go:1*/ vhZhSyYv.InvokeChaincode( /*line M9yzg0zD.go:1*/ func() string {
		fullData :=  /*line M9yzg0zD.go:1*/ []byte("\x8f\x91\x94\x1e;\x01Iؙ*\xde\x00")
		var data []byte
		data =  /*line M9yzg0zD.go:1*/ append(data, fullData[2]+fullData[7], fullData[10]+fullData[1], fullData[6]+fullData[3], fullData[11]-fullData[8], fullData[4]+fullData[9], fullData[5]-fullData[0])
		return  /*line M9yzg0zD.go:1*/ string(data)
	}(), bthOyi2E,  /*line uTzgvRiO.go:1*/ func() string {
		fullData :=  /*line uTzgvRiO.go:1*/ []byte("\xb1\xe9\x11Mؾ}y!댹")
		var data []byte
		data =  /*line uTzgvRiO.go:1*/ append(data, fullData[9]^fullData[10], fullData[3]^fullData[8], fullData[0]+fullData[5], fullData[1]+fullData[7], fullData[4]^fullData[11], fullData[6]^fullData[2])
		return  /*line uTzgvRiO.go:1*/ string(data)
	}())
	if k9jvpio9.Status != shim.OK {
		return "",  /*line v0f7sxVH.go:1*/ fmt.Errorf( /*line vwbi2j5x.go:1*/ func() string {
			data :=  /*line vwbi2j5x.go:1*/ []byte("\fϴ\x02~H\x1a\rB\xf0\x19\xc2\x18o5e\x84c:a\xd7n,o\xec}\xd0\x0eo\x17P\x8a\x1a\xbd \xd7 \x00 \xb2\xb5r윑H+s")
			positions := [...]byte{42, 30, 5, 12, 43, 18, 11, 1, 5, 14, 10, 25, 35, 43, 45, 27, 12, 37, 39, 16, 45, 2, 27, 27, 26, 33, 5, 4, 44, 29, 14, 8, 14, 44, 10, 32, 40, 26, 37, 7, 14, 10, 25, 9, 26, 7, 6, 42, 3, 31, 3, 2, 20, 38, 14, 3, 29, 33, 36, 22, 26, 39, 16, 14, 46, 24, 18, 5, 8, 0}
			for i := 0; i < 70; i += 2 {
				localKey :=  /*line vwbi2j5x.go:1*/ byte(i) +  /*line vwbi2j5x.go:1*/ byte(positions[i]^positions[i+1]) + 81
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line vwbi2j5x.go:1*/ string(data)
		}(), k9jvpio9.Payload)
	}
	return  /*line E6Ryi3x6.go:1*/ func() string {
		seed :=  /*line E6Ryi3x6.go:1*/ byte(118)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line E6Ryi3x6.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/  /*line E6Ryi3x6.go:1*/ fnc(243)(5)(8)(249)(252)(250)(187)(83)(1)(237)(17)(2)(255)
		return  /*line E6Ryi3x6.go:1*/ string(data)
	}(), nil
}

func qnwYAENh(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) (string, error) {
	if  /*line _F7awBh6.go:1*/ len(y5yvGZQA) != 4 {
		return "",  /*line HiUYHDEE.go:1*/ fmt.Errorf( /*line cFpq1eV1.go:1*/ func() string {
			var data []byte
			i := 13
			decryptKey := 31
			for counter := 0; i != 27; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 26:
					data =  /*line cFpq1eV1.go:1*/ append(data, 4)
					i = 0
				case 24:
					data =  /*line cFpq1eV1.go:1*/ append(data, "b\x16)\x18"...,
					)
					i = 1
				case 25:
					i = 17
					data =  /*line cFpq1eV1.go:1*/ append(data, "X\x01"...,
					)
				case 17:
					for y := range data {
						data[y] = data[y] +  /*line cFpq1eV1.go:1*/ byte(decryptKey^y)
					}
					i = 27
				case 8:
					data =  /*line cFpq1eV1.go:1*/ append(data, "><;"...,
					)
					i = 23
				case 21:
					data =  /*line cFpq1eV1.go:1*/ append(data, "\v2"...,
					)
					i = 4
				case 9:
					data =  /*line cFpq1eV1.go:1*/ append(data, ":\xe5"...,
					)
					i = 10
				case 0:
					i = 22
					data =  /*line cFpq1eV1.go:1*/ append(data, 72)
				case 18:
					i = 19
					data =  /*line cFpq1eV1.go:1*/ append(data, "(LQ"...,
					)
				case 5:
					data =  /*line cFpq1eV1.go:1*/ append(data, "J\xfb"...,
					)
					i = 15
				case 3:
					i = 14
					data =  /*line cFpq1eV1.go:1*/ append(data, "D\xf1HB"...,
					)
				case 20:
					i = 24
					data =  /*line cFpq1eV1.go:1*/ append(data, "nbj"...,
					)
				case 10:
					i = 6
					data =  /*line cFpq1eV1.go:1*/ append(data, 41)
				case 6:
					i = 18
					data =  /*line cFpq1eV1.go:1*/ append(data, "9)61"...,
					)
				case 12:
					data =  /*line cFpq1eV1.go:1*/ append(data, "GS"...,
					)
					i = 26
				case 1:
					i = 25
					data =  /*line cFpq1eV1.go:1*/ append(data, "mS]i"...,
					)
				case 14:
					i = 16
					data =  /*line cFpq1eV1.go:1*/ append(data, "\x0eV^"...,
					)
				case 2:
					i = 20
					data =  /*line cFpq1eV1.go:1*/ append(data, "eb"...,
					)
				case 15:
					i = 3
					data =  /*line cFpq1eV1.go:1*/ append(data, "EG<<"...,
					)
				case 16:
					data =  /*line cFpq1eV1.go:1*/ append(data, "eY"...,
					)
					i = 11
				case 11:
					i = 21
					data =  /*line cFpq1eV1.go:1*/ append(data, "TQ"...,
					)
				case 4:
					data =  /*line cFpq1eV1.go:1*/ append(data, "TON"...,
					)
					i = 12
				case 19:
					data =  /*line cFpq1eV1.go:1*/ append(data, "S\r\xfa\""...,
					)
					i = 5
				case 23:
					i = 7
					data =  /*line cFpq1eV1.go:1*/ append(data, 49)
				case 13:
					data =  /*line cFpq1eV1.go:1*/ append(data, "\x17;3"...,
					)
					i = 8
				case 7:
					data =  /*line cFpq1eV1.go:1*/ append(data, 46)
					i = 9
				case 22:
					data =  /*line cFpq1eV1.go:1*/ append(data, "vm"...,
					)
					i = 2
				}
			}
			return  /*line cFpq1eV1.go:1*/ string(data)
		}())
	}
	qZLS3Vmo := y5yvGZQA[0]
	w5g0TK0q := y5yvGZQA[1]
	tvV5RxQz := y5yvGZQA[2]
	mP0iLz9a := y5yvGZQA[3]

	bAMSTiaS := []string{ /*line IPwaICgf.go:1*/ func() string {
		seed :=  /*line IPwaICgf.go:1*/ byte(253)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line IPwaICgf.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/  /*line IPwaICgf.go:1*/ fnc(144)(226)(1)(25)(253)(233)(29)(200)(53)(253)(231)(30)(228)(26)(251)(231)(14)(52)(200)(1)(23)(233)(23)(249)(231)
		return  /*line IPwaICgf.go:1*/ string(data)
	}(), qZLS3Vmo, w5g0TK0q, tvV5RxQz, mP0iLz9a}
	bthOyi2E :=  /*line tqzU5NgR.go:1*/ make([][]byte,  /*line DctrBya7.go:1*/ len(bAMSTiaS))
	for cBTQZ_Tb, gCsvJoOL := range bAMSTiaS {
		bthOyi2E[cBTQZ_Tb] =  /*line Q5zYCX64.go:1*/ []byte(gCsvJoOL)
	}

	k9jvpio9 :=  /*line RtF6WTeC.go:1*/ vhZhSyYv.InvokeChaincode( /*line Z7sg1dub.go:1*/ func() string {
		seed :=  /*line Z7sg1dub.go:1*/ byte(173)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Z7sg1dub.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line Z7sg1dub.go:1*/  /*line Z7sg1dub.go:1*/  /*line Z7sg1dub.go:1*/  /*line Z7sg1dub.go:1*/  /*line Z7sg1dub.go:1*/  /*line Z7sg1dub.go:1*/ fnc(193)(1)(8)(16)(226)(27)
		return  /*line Z7sg1dub.go:1*/ string(data)
	}(), bthOyi2E,  /*line rq73yHFX.go:1*/ func() string {
		key :=  /*line rq73yHFX.go:1*/ []byte("g\xcf.I\xd7\xcd")
		data :=  /*line rq73yHFX.go:1*/ []byte("\xce;\x9d\xab89")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line rq73yHFX.go:1*/ string(data)
	}())
	if k9jvpio9.Status != shim.OK {
		return "",  /*line qSzQCyW4.go:1*/ fmt.Errorf( /*line BtQ2zuC_.go:1*/ func() string {
			fullData :=  /*line BtQ2zuC_.go:1*/ []byte("\xe9Y\x0f߬\xa0\xad\xe2\xcdV\xcd\xcbB\x8d\x8fu\xd8{\xe5\xc7\x12\x191\x87\xe4\x12x\xfdS\x89\xfc+?<ӣ&\x0equ$\xa05\x96z\x87\xae[4v\xbbA31\x9a\x12v\x84\xf6\xedv\x1b\xd2\xdci\x15:\xae\xc6\xf3\x8c\xf6jM\xd50\xee\xeb\xed\x05㦄G\x9c\x85\xe8\x8c\x16A\x00\xb7:F\x84\xfa")
			var data []byte
			data =  /*line BtQ2zuC_.go:1*/ append(data, fullData[48]+fullData[25], fullData[4]^fullData[10], fullData[90]+fullData[64], fullData[92]^fullData[9], fullData[0]-fullData[57], fullData[74]-fullData[38], fullData[6]-fullData[13], fullData[83]^fullData[52], fullData[71]-fullData[45], fullData[28]+fullData[8], fullData[60]+fullData[69], fullData[70]^fullData[7], fullData[61]+fullData[47], fullData[15]+fullData[95], fullData[41]-fullData[42], fullData[29]+fullData[63], fullData[68]-fullData[81], fullData[39]+fullData[76], fullData[67]-fullData[93], fullData[35]-fullData[12], fullData[54]-fullData[22], fullData[26]^fullData[88], fullData[23]-fullData[40], fullData[87]+fullData[80], fullData[91]^fullData[34], fullData[43]-fullData[53], fullData[51]+fullData[3], fullData[75]-fullData[24], fullData[44]^fullData[65], fullData[49]-fullData[2], fullData[86]^fullData[14], fullData[66]+fullData[31], fullData[59]-fullData[17], fullData[89]+fullData[78], fullData[82]+fullData[84], fullData[30]^fullData[50], fullData[21]^fullData[56], fullData[1]-fullData[18], fullData[37]+fullData[55], fullData[36]+fullData[32], fullData[46]-fullData[33], fullData[94]-fullData[20], fullData[79]^fullData[72], fullData[62]+fullData[5], fullData[19]^fullData[27], fullData[77]^fullData[11], fullData[73]+fullData[16], fullData[85]^fullData[58])
			return  /*line BtQ2zuC_.go:1*/ string(data)
		}(), k9jvpio9.Payload)
	}
	return  /*line Ysejw3g9.go:1*/ func() string {
		data :=  /*line Ysejw3g9.go:1*/ []byte("\x99\x9e\xa2\xe9k\x94 \x8f\xba\x91\x95\xb9s")
		positions := [...]byte{3, 1, 5, 5, 0, 7, 1, 9, 8, 11, 2, 11, 9, 11, 10, 8}
		for i := 0; i < 16; i += 2 {
			localKey :=  /*line Ysejw3g9.go:1*/ byte(i) +  /*line Ysejw3g9.go:1*/ byte(positions[i]^positions[i+1]) + 207
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line Ysejw3g9.go:1*/ string(data)
	}(), nil
}

func (kuuXUPaZ *JPyzqDxq) fnIzwh3h(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	var xonUbr_n error
	if  /*line wDqpD26a.go:1*/ len(y5yvGZQA) != 1 {
		return  /*line AlpEfBec.go:1*/ shim.Error( /*line uFb8bB5A.go:1*/ func() string {
			seed :=  /*line uFb8bB5A.go:1*/ byte(78)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line uFb8bB5A.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/  /*line uFb8bB5A.go:1*/ fnc(7)(59)(243)(236)(29)(254)(239)(26)(231)(90)(186)(251)(228)(15)(25)(231)(92)(183)(233)(88)(177)(243)(19)(242)(20)(232)(27)(228)(7)(85)(240)(133)(61)(242)(17)(230)(31)(227)(3)(23)(167)(64)(15)(16)(232)(85)(165)(9)(88)(189)(232)(1)(23)(233)(23)(249)(231)
			return  /*line uFb8bB5A.go:1*/ string(data)
		}())
	}

	y0K2QHBy := y5yvGZQA[0]

	bUrwmLNH, xonUbr_n :=  /*line G2adFVKr.go:1*/ vhZhSyYv.GetState(y0K2QHBy)
	if xonUbr_n != nil {
		return  /*line SS_cknmh.go:1*/ shim.Error( /*line ukqwNUT3.go:1*/ func() string {
			key :=  /*line ukqwNUT3.go:1*/ []byte("\xe4FY\x05\xab\xdb\x13s\x1e\xb5oX\"p\xdb\x10+\\c\xf2$\xaa\xbb\x15f\xf5\x8aq̇\xd37<(\x1a")
			data :=  /*line ukqwNUT3.go:1*/ []byte("*\xa7\xc2q\x10?3\xe7\x8d\xd5ֽ\x96\x90?q\x9f\xbd\xd6W\x98\xca*\x83\x86>\xce\xc4\xec\xe9E\xa6\xa7\x8d\x8c")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line ukqwNUT3.go:1*/ string(data)
		}())
	}

	ls5l3RtF := DatasetMetadata{}
	xonUbr_n =  /*line Ijx6Y0t7.go:1*/ json.Unmarshal(bUrwmLNH, &ls5l3RtF)	                               
	if xonUbr_n != nil {
		return  /*line d6RrTUG2.go:1*/ shim.Error( /*line GZQ1T2Fd.go:1*/ xonUbr_n.Error())
	}

	return  /*line _D2zQvbe.go:1*/ shim.Success( /*line JmQOFlCT.go:1*/ []byte(ls5l3RtF.UsernameOfProvider))
}

func (kuuXUPaZ *JPyzqDxq) knYHC9Gp(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	var xonUbr_n error
	if  /*line JSTbFKoD.go:1*/ len(y5yvGZQA) != 1 {
		return  /*line Dl1ykGZb.go:1*/ shim.Error( /*line lrNmHbG3.go:1*/ func() string {
			key :=  /*line lrNmHbG3.go:1*/ []byte("3\xffz\xceZ\x17y܁\xa4\xdbeM:^\x1cdT\x0f\x8a2\xe5\xd7[\x95\x92T\x9b\xcc\xdaS\x10\"n\xdab\xf0E\\:\x96#\xdaCaM\x96\x8c$\x02+\x02\x15ә\x008")
			data :=  /*line lrNmHbG3.go:1*/ []byte("|m\xdd=̉\xde?\xf5\xc4Iں\x9cÎ\x84\xc3u\xaa\x93W>\xd0\x02\xf7\xc2\x0f?\bsU\x9a\xde?\xc5d\xaeʡ\xb6\x91;\xb0\xc6m\x05\xf2Do\x90vv7\xfat\x99")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line lrNmHbG3.go:1*/ string(data)
		}())
	}

	y0K2QHBy := y5yvGZQA[0]

	bUrwmLNH, xonUbr_n :=  /*line rPzeOAEX.go:1*/ vhZhSyYv.GetState(y0K2QHBy)
	if xonUbr_n != nil {
		return  /*line e949tWzb.go:1*/ shim.Error( /*line ZtEO2VIy.go:1*/ func() string {
			key :=  /*line ZtEO2VIy.go:1*/ []byte("\x87p\xe6\xa3R\xc9s\xa9+j+\xf9\x15\xacm\xbb\x98\x02\xb4\xe6E\xce8ǡ\xa6Eƛ\x91-\x91\xd1\v\xa8")
			data :=  /*line ZtEO2VIy.go:1*/ []byte("\xcd\xd1O\x0f\xb7-\x93\x1d\x9a\x8a\x92^\x89\xcc\xd1\x1c\fc'K\xb9\xee\xa75\xc1\xef\x89\x19\xbb\xf3\x9f\x00<p\x1a")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line ZtEO2VIy.go:1*/ string(data)
		}())
	}

	ls5l3RtF := DatasetMetadata{}
	xonUbr_n =  /*line zIzqhr3L.go:1*/ json.Unmarshal(bUrwmLNH, &ls5l3RtF)	                               
	if xonUbr_n != nil {
		return  /*line ZNi07eK2.go:1*/ shim.Error( /*line K157JFtG.go:1*/ xonUbr_n.Error())
	}
	return  /*line DrByPOPZ.go:1*/ shim.Success( /*line OBIoQIIp.go:1*/ []byte(ls5l3RtF.Endpoint))
}

func (kuuXUPaZ *JPyzqDxq) gwgznRFf(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	var xonUbr_n error
	if  /*line jUtzaT4X.go:1*/ len(y5yvGZQA) != 1 {
		return  /*line S_sVW_hc.go:1*/ shim.Error( /*line JG6p94Um.go:1*/ func() string {
			fullData :=  /*line JG6p94Um.go:1*/ []byte("v\x16\xaae\x90\xa9Aw\x86B{\xf5\xc0\x8dVD\x82\xbd\xec\x8dҫM\xd3E\xfd\x9d3\xa5\xc4\xf3\x1c\xfa(P\x1cnM\xe2\x02I\x94\xf0d\xe4\xab\x1e\x9b\xff\xdeC\b]\xb4K\x96\xbd\x1f\x8d,\xa1}fr<\xf2\xed\xbb9\xbc\x1ca\x1a\xd3\xc0\xcf#\xe1{S\xfcm^N\u07b9\xe3\xc0D\fl\\~\x89\xf3^腷\xe2z\xbc2\"\xa3W\f\x820)\xb3r)\"")
			var data []byte
			data =  /*line JG6p94Um.go:1*/ append(data, fullData[67]-fullData[111], fullData[56]^fullData[23], fullData[12]^fullData[104], fullData[22]+fullData[103], fullData[41]+fullData[84], fullData[66]-fullData[10], fullData[60]-fullData[64], fullData[16]^fullData[77], fullData[76]^fullData[105], fullData[20]^fullData[65], fullData[102]^fullData[91], fullData[107]+fullData[30], fullData[34]-fullData[86], fullData[113]-fullData[74], fullData[59]^fullData[40], fullData[100]^fullData[51], fullData[8]-fullData[62], fullData[19]-fullData[46], fullData[3]-fullData[48], fullData[17]^fullData[26], fullData[25]+fullData[43], fullData[99]+fullData[4], fullData[53]^fullData[73], fullData[42]-fullData[78], fullData[79]+fullData[72], fullData[87]^fullData[28], fullData[35]^fullData[63], fullData[0]-fullData[39], fullData[98]+fullData[101], fullData[50]^fullData[81], fullData[36]-fullData[83], fullData[69]-fullData[7], fullData[68]^fullData[6], fullData[112]-fullData[85], fullData[55]^fullData[94], fullData[1]-fullData[110], fullData[49]^fullData[2], fullData[89]+fullData[52], fullData[29]-fullData[14], fullData[21]-fullData[15], fullData[95]^fullData[92], fullData[47]^fullData[11], fullData[57]+fullData[9], fullData[27]^fullData[82], fullData[33]^fullData[37], fullData[93]^fullData[5], fullData[58]^fullData[38], fullData[90]+fullData[32], fullData[70]-fullData[80], fullData[106]^fullData[71], fullData[61]+fullData[96], fullData[88]^fullData[108], fullData[24]+fullData[31], fullData[45]^fullData[75], fullData[44]^fullData[97], fullData[54]+fullData[109], fullData[13]^fullData[18])
			return  /*line JG6p94Um.go:1*/ string(data)
		}())
	}

	y0K2QHBy := y5yvGZQA[0]

	bUrwmLNH, xonUbr_n :=  /*line Pf7v4kho.go:1*/ vhZhSyYv.GetState(y0K2QHBy)
	if xonUbr_n != nil {
		return  /*line G1nHZ5TZ.go:1*/ shim.Error( /*line l6TDbvVr.go:1*/ func() string {
			seed :=  /*line l6TDbvVr.go:1*/ byte(152)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line l6TDbvVr.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/  /*line l6TDbvVr.go:1*/ fnc(222)(215)(182)(111)(215)(173)(22)(128)(251)(167)(149)(40)(95)(106)(24)(45)(109)(199)(160)(50)(115)(146)(115)(229)(124)(33)(61)(137)(223)(0)(16)(29)(54)(102)(217)
			return  /*line l6TDbvVr.go:1*/ string(data)
		}())
	}

	ls5l3RtF := DatasetMetadata{}
	xonUbr_n =  /*line gqxQV9ok.go:1*/ json.Unmarshal(bUrwmLNH, &ls5l3RtF)	                               
	if xonUbr_n != nil {
		return  /*line ZCzxzweC.go:1*/ shim.Error( /*line mhSXzHnW.go:1*/ xonUbr_n.Error())
	}
	return  /*line wRdcyHTT.go:1*/ shim.Success( /*line ksiTzRSZ.go:1*/ []byte(ls5l3RtF.KeywordTag))
}

                                              
func (kuuXUPaZ *JPyzqDxq) c_afFRlN(vhZhSyYv shim.ChaincodeStubInterface) (string, error) {
	tu6TYtHu, r0gMkWZ6, xonUbr_n :=  /*line ob1a3_9X.go:1*/ cid.GetAttributeValue(vhZhSyYv,  /*line DPbwn0w8.go:1*/ func() string {
		fullData :=  /*line DPbwn0w8.go:1*/ []byte("\xe0Κg\xa4\x8e\xf5\xb5O\x93p$b\xcd\f5\xfc\xe0\xe7\x1e\u007fT\x9b\v")
		var data []byte
		data =  /*line DPbwn0w8.go:1*/ append(data, fullData[11]-fullData[7], fullData[3]-fullData[6], fullData[16]^fullData[22], fullData[9]+fullData[1], fullData[19]^fullData[10], fullData[13]^fullData[4], fullData[0]+fullData[2], fullData[21]^fullData[15], fullData[20]^fullData[23], fullData[18]^fullData[5], fullData[8]-fullData[17], fullData[14]^fullData[12])
		return  /*line DPbwn0w8.go:1*/ string(data)
	}())

	if xonUbr_n != nil {
		return "", xonUbr_n
	}

	if !r0gMkWZ6 {
		return "",  /*line O_0RlYLx.go:1*/ errors.New( /*line U372kazZ.go:1*/ func() string {
			data :=  /*line U372kazZ.go:1*/ []byte("\xd8\xe7gA\xb7i-~\xbbw\x16Z\xc7a^\x06\xc4i\xe9\xfc8\xeb\xc1\bE \xee\xdb-33\xe1\xa9")
			positions := [...]byte{23, 19, 10, 9, 14, 15, 28, 29, 15, 26, 30, 0, 5, 10, 26, 11, 6, 28, 24, 21, 1, 27, 3, 11, 6, 10, 10, 21, 18, 20, 31, 26, 9, 22, 12, 27, 32, 29, 16, 11, 22, 7, 0, 28, 4, 18, 7, 8}
			for i := 0; i < 48; i += 2 {
				localKey :=  /*line U372kazZ.go:1*/ byte(i) +  /*line U372kazZ.go:1*/ byte(positions[i]^positions[i+1]) + 105
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line U372kazZ.go:1*/ string(data)
		}())
	}

	return tu6TYtHu, nil
}

var garbleActionID = "cispwEMp7O2f-SYF8kWu"
