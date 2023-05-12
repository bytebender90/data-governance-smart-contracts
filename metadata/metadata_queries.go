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
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (kuuXUPaZ *JPyzqDxq) leLg2iBJ(vhZhSyYv shim.ChaincodeStubInterface) pb.Response {
	aarIVocK :=  /*line mlrd9KrB.go:1*/ func() string {
		key :=  /*line mlrd9KrB.go:1*/ []byte("\xdf+\x9b\xa4\xd6[%s\x1bG\x8a,\xfd΅\xe6\x8a\x19k\xfe|\xcchM\x89\xbb\u007f>8b\xaf\x89\x1d\x94\xf1")
		data :=  /*line mlrd9KrB.go:1*/ []byte("ZM\x0e\tB\xc0\x88犹\xacfx\xf0\xe9U\xedm\xe4n\xe1\xee\xa2o\xd6 \xf3\x9f\x9c\xc3#\xea?\x11n")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line mlrd9KrB.go:1*/ string(data)
	}()

	c7Phsypr, xonUbr_n :=  /*line u12Dn6Wh.go:1*/ _IAEAzGw(vhZhSyYv, aarIVocK)
	if xonUbr_n != nil {
		return  /*line _jH6O721.go:1*/ shim.Error( /*line JowVU_My.go:1*/ xonUbr_n.Error())
	}
	return  /*line s7MPyUjZ.go:1*/ shim.Success(c7Phsypr)

}

func (kuuXUPaZ *JPyzqDxq) jddDqpYa(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	var gtLeVnpr string
	var xonUbr_n error

	if  /*line JH0EnSr8.go:1*/ len(y5yvGZQA) != 1 {
		return  /*line J6Hg5Q8x.go:1*/ shim.Error( /*line FplVfpIb.go:1*/ func() string {
			var data []byte
			i := 6
			decryptKey := 97
			for counter := 0; i != 4; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					data =  /*line FplVfpIb.go:1*/ append(data, 152)
					i = 14
				case 11:
					data =  /*line FplVfpIb.go:1*/ append(data, "\xae\xb4"...,
					)
					i = 15
				case 20:
					i = 31
					data =  /*line FplVfpIb.go:1*/ append(data, "\xf1\xb7\xa5"...,
					)
				case 21:
					i = 11
					data =  /*line FplVfpIb.go:1*/ append(data, 163)
				case 27:
					i = 7
					data =  /*line FplVfpIb.go:1*/ append(data, "\x8bσ\x8b"...,
					)
				case 1:
					i = 22
					data =  /*line FplVfpIb.go:1*/ append(data, "\xa6\xbc\xa3\xad"...,
					)
				case 18:
					i = 30
					data =  /*line FplVfpIb.go:1*/ append(data, 160)
				case 17:
					data =  /*line FplVfpIb.go:1*/ append(data, "\x9bی\x96"...,
					)
					i = 23
				case 30:
					data =  /*line FplVfpIb.go:1*/ append(data, "\xb7\xbe\xb6"...,
					)
					i = 8
				case 26:
					data =  /*line FplVfpIb.go:1*/ append(data, "\xbe\xeb"...,
					)
					i = 1
				case 29:
					i = 19
					data =  /*line FplVfpIb.go:1*/ append(data, "\x98\x9a\x93"...,
					)
				case 15:
					i = 16
					data =  /*line FplVfpIb.go:1*/ append(data, "\xb5\xa1"...,
					)
				case 9:
					i = 17
					data =  /*line FplVfpIb.go:1*/ append(data, "\x97\x93\x95\x81"...,
					)
				case 13:
					data =  /*line FplVfpIb.go:1*/ append(data, 188)
					i = 5
				case 23:
					i = 2
					data =  /*line FplVfpIb.go:1*/ append(data, "ގ\x89"...,
					)
				case 3:
					data =  /*line FplVfpIb.go:1*/ append(data, 191)
					i = 0
				case 24:
					i = 12
					data =  /*line FplVfpIb.go:1*/ append(data, "\x8e\x8a\x82"...,
					)
				case 25:
					data =  /*line FplVfpIb.go:1*/ append(data, 173)
					i = 10
				case 16:
					data =  /*line FplVfpIb.go:1*/ append(data, 166)
					i = 26
				case 8:
					data =  /*line FplVfpIb.go:1*/ append(data, 173)
					i = 25
				case 0:
					data =  /*line FplVfpIb.go:1*/ append(data, 242)
					i = 13
				case 5:
					i = 20
					data =  /*line FplVfpIb.go:1*/ append(data, 182)
				case 10:
					i = 29
					data =  /*line FplVfpIb.go:1*/ append(data, "\xf1\xfc"...,
					)
				case 12:
					data =  /*line FplVfpIb.go:1*/ append(data, "ʅ\x89\x84"...,
					)
					i = 27
				case 28:
					i = 4
					for y := range data {
						data[y] = data[y] ^  /*line FplVfpIb.go:1*/ byte(decryptKey^y)
					}
				case 6:
					i = 21
					data =  /*line FplVfpIb.go:1*/ append(data, "\x8b\xad"...,
					)
				case 14:
					i = 28
					data =  /*line FplVfpIb.go:1*/ append(data, "\xf0\xfa"...,
					)
				case 7:
					i = 9
					data =  /*line FplVfpIb.go:1*/ append(data, "Ҟ\x95\x85"...,
					)
				case 19:
					data =  /*line FplVfpIb.go:1*/ append(data, "\x85\x82\x92"...,
					)
					i = 24
				case 22:
					i = 3
					data =  /*line FplVfpIb.go:1*/ append(data, 169)
				case 31:
					i = 18
					data =  /*line FplVfpIb.go:1*/ append(data, 179)
				}
			}
			return  /*line FplVfpIb.go:1*/ string(data)
		}())
	}

	pzKghUuG := y5yvGZQA[0]
	f1MtQil9, xonUbr_n :=  /*line fmzLtqYU.go:1*/ vhZhSyYv.GetState(pzKghUuG)
	if xonUbr_n != nil {
		gtLeVnpr =  /*line HjE3BdgF.go:1*/ func() string {
			key :=  /*line HjE3BdgF.go:1*/ []byte("\xe4\xe2\n\xcc5f\xc3\x00\x86\x14\xb2\x1a\xdd\xee\x0e\xf4\xf7;\xfe\\\xed\xdbd\x95\xcb\xd5\x04o\x95'\xa0\x93\xd7H")
			data :=  /*line HjE3BdgF.go:1*/ []byte("\x9f\xc0O\xbeG\t\xb1\"\xbc6\xf4{\xb4\x82k\x90\xd7O\x91|\x8a\xbe\x10\xb5\xb8\xa1e\x1b\xf0\a\xc6\xfc\xa5h")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line HjE3BdgF.go:1*/ string(data)
		}() + pzKghUuG +  /*line cJHaLRKV.go:1*/ func() string {
			key :=  /*line cJHaLRKV.go:1*/ []byte("\xd1\x10")
			data :=  /*line cJHaLRKV.go:1*/ []byte("\xf3m")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line cJHaLRKV.go:1*/ string(data)
		}()
		return  /*line YIlT8I81.go:1*/ shim.Error(gtLeVnpr)
	} else if f1MtQil9 == nil {
		gtLeVnpr =  /*line kObzOSU7.go:1*/ func() string {
			data :=  /*line kObzOSU7.go:1*/ []byte("5#(rr)r\xdf\xd4\xdbM\xd9\xf99\xd7GP\xe6\xddd\xdc!\xdb n*\x1b\xad\xc2x\x1b\xa2\xfd\xf9 ")
			positions := [...]byte{8, 11, 20, 26, 9, 18, 27, 14, 7, 17, 2, 12, 7, 22, 33, 5, 16, 32, 31, 11, 14, 28, 30, 30, 11, 1, 33, 16, 14, 26, 0, 0, 21, 17, 15, 26, 25, 28, 13, 5}
			for i := 0; i < 40; i += 2 {
				localKey :=  /*line kObzOSU7.go:1*/ byte(i) +  /*line kObzOSU7.go:1*/ byte(positions[i]^positions[i+1]) + 156
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line kObzOSU7.go:1*/ string(data)
		}() + pzKghUuG +  /*line Y3XzIFfy.go:1*/ func() string {
			seed :=  /*line Y3XzIFfy.go:1*/ byte(93)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line Y3XzIFfy.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line Y3XzIFfy.go:1*/  /*line Y3XzIFfy.go:1*/ fnc(197)(91)
			return  /*line Y3XzIFfy.go:1*/ string(data)
		}()
		return  /*line P5BD2Osj.go:1*/ shim.Error(gtLeVnpr)
	}

	return  /*line f6RCnB4q.go:1*/ shim.Success(f1MtQil9)

}

func (kuuXUPaZ *JPyzqDxq) fKDeNQZA(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	if  /*line dYJwuG45.go:1*/ len(y5yvGZQA) != 1 {
		return  /*line sKNV6jgj.go:1*/ shim.Error( /*line d3NwaX9g.go:1*/ func() string {
			key :=  /*line d3NwaX9g.go:1*/ []byte("\xd20\xfc\f\x11\x96\x12\x19\x8f\xd9\xc9\xfb\x1f\xcf\xf0\x0f\tN\xff\xb4C\x9b\x12\x8bn\x8d%j\xa9\\S\xde\xc4\xd1\xf5*\xb7@\xae\x88\rԲ")
			data :=  /*line d3NwaX9g.go:1*/ []byte("\x1b\x9e_{\x83\bw|\x03\xf9*m\x86D]tw\xc2r\xe2c\xe0\x8a\xfb\xd3\xf0\x99\xd3\x17\xc3sB%EV}&\xb5 \xebr\x1d\xf6")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line d3NwaX9g.go:1*/ string(data)
		}())
	}
	aifJoMiz := y5yvGZQA[0]

	aarIVocK :=  /*line _PQ6zWpv.go:1*/ fmt.Sprintf( /*line EWKjZEfV.go:1*/ func() string {
		fullData :=  /*line EWKjZEfV.go:1*/ []byte(":\xff\x94\xc8E\n\xc5\xeahW\xb5(\xa8\x1f\xb4\xd9%kaq\x1fb\x9c\x8d,\xaa\xcf\xed8\xf3k\xc2\x0e\x856J\xdb\xf0\\%\xdb.\x91\xadu\xe42$X4\xe83\b_\x13;.\t\x82\xa0N\x83n\xb3S\x99\x15\xaf,\xceL\xd0Z\b\xf6ƛ\xbdن\xb6\xf7\xa8Rd^\x16=q'U\xf0\\U\xc8g\xaaƾ\xa1\x10\x8a\xbe\xb6\xa7\x99\xc7\xf5S\x82OV\xeb\xb9x\x15\x10\x02\x9c\x00\xf2\xf73\xfa/\xbb+ԡ-]E(#\xb4}u\xfc\x8a\x17\xcd\xd9\xdba\x98\xf9~\x1e\xd54")
		var data []byte
		data =  /*line EWKjZEfV.go:1*/ append(data, fullData[111]+fullData[39], fullData[9]^fullData[136], fullData[42]-fullData[147], fullData[34]+fullData[124], fullData[30]-fullData[1], fullData[96]-fullData[131], fullData[53]-fullData[137], fullData[66]-fullData[128], fullData[0]^fullData[90], fullData[58]^fullData[37], fullData[130]+fullData[6], fullData[14]+fullData[79], fullData[33]-fullData[5], fullData[148]^fullData[121], fullData[31]-fullData[85], fullData[16]^fullData[35], fullData[107]+fullData[62], fullData[83]+fullData[117], fullData[56]-fullData[10], fullData[61]^fullData[29], fullData[38]+fullData[57], fullData[74]+fullData[68], fullData[46]^fullData[52], fullData[138]-fullData[8], fullData[11]-fullData[40], fullData[102]^fullData[142], fullData[115]^fullData[18], fullData[113]+fullData[82], fullData[28]^fullData[92], fullData[141]-fullData[114], fullData[105]^fullData[27], fullData[134]-fullData[64], fullData[95]+fullData[125], fullData[103]-fullData[101], fullData[3]-fullData[12], fullData[59]^fullData[109], fullData[70]^fullData[132], fullData[55]^fullData[72], fullData[63]^fullData[106], fullData[4]^fullData[47], fullData[108]+fullData[119], fullData[80]^fullData[15], fullData[99]-fullData[24], fullData[143]^fullData[54], fullData[94]+fullData[76], fullData[122]-fullData[69], fullData[146]+fullData[26], fullData[77]-fullData[48], fullData[50]^fullData[118], fullData[91]+fullData[88], fullData[110]^fullData[126], fullData[104]^fullData[97], fullData[73]-fullData[2], fullData[32]-fullData[43], fullData[127]^fullData[123], fullData[144]-fullData[49], fullData[116]^fullData[19], fullData[84]+fullData[100], fullData[89]-fullData[75], fullData[145]^fullData[25], fullData[135]+fullData[120], fullData[7]-fullData[44], fullData[78]+fullData[65], fullData[51]-fullData[71], fullData[98]^fullData[36], fullData[129]-fullData[45], fullData[112]^fullData[67], fullData[86]^fullData[149], fullData[139]+fullData[133], fullData[23]-fullData[17], fullData[81]+fullData[41], fullData[60]^fullData[87], fullData[140]+fullData[93], fullData[22]-fullData[13], fullData[21]^fullData[20])
		return  /*line EWKjZEfV.go:1*/ string(data)
	}(), aifJoMiz)

	c7Phsypr, xonUbr_n :=  /*line UGD7H2pR.go:1*/ _IAEAzGw(vhZhSyYv, aarIVocK)
	if xonUbr_n != nil {
		return  /*line axRZGQXq.go:1*/ shim.Error( /*line MjYn59dF.go:1*/ xonUbr_n.Error())
	}
	return  /*line yKATd1EP.go:1*/ shim.Success(c7Phsypr)

}

func (kuuXUPaZ *JPyzqDxq) wYYPiY8F(vhZhSyYv shim.ChaincodeStubInterface) pb.Response {
	dAMBvbjT, xonUbr_n :=  /*line InFr8hqB.go:1*/ kuuXUPaZ.glBEweME(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line IGyUJADK.go:1*/ shim.Error( /*line x6gUX6wU.go:1*/ xonUbr_n.Error())
	}

	aarIVocK :=  /*line tNdXMw6z.go:1*/ fmt.Sprintf( /*line ZzJ9TpjU.go:1*/ func() string {
		fullData :=  /*line ZzJ9TpjU.go:1*/ []byte("\xc1L\xf4\u007f\xb8v\xc1\xaa}\x00`\xcf\xcb\x16a\x02C\xf7{\xcd>g\x16\x1c$\xe4e{0\xed\xdfȂ\x9b\xbc%\x9d*\xbe\xd4_\xc2\xf2<\xc4\xe8H[\x01\x9e#”#O`\xe6\ued0a}\xa3\xb3\x197x\x87`\x17\x1d\xe8\x89\xc6MZ|>\x9bC\x9f\xef\x83\xf2\xee\xa1\x1a\xb3\x93\x05r\x1e\"s\xc6\x11\xee\xee\x1e\xb1\x8e\x95\x994\xccvsY\xf9\xe6\x94?\x06\xb5\xd7\xd2\x0e\x01y\xdd>\xc1\x9e^")
		var data []byte
		data =  /*line ZzJ9TpjU.go:1*/ append(data, fullData[112]^fullData[61], fullData[10]-fullData[77], fullData[76]+fullData[17], fullData[21]^fullData[15], fullData[109]^fullData[60], fullData[62]^fullData[73], fullData[24]-fullData[121], fullData[48]+fullData[106], fullData[33]^fullData[2], fullData[120]-fullData[104], fullData[82]^fullData[85], fullData[67]+fullData[63], fullData[88]+fullData[71], fullData[25]^fullData[94], fullData[32]+fullData[51], fullData[8]+fullData[42], fullData[119]^fullData[38], fullData[75]^fullData[116], fullData[0]-fullData[46], fullData[96]^fullData[49], fullData[18]^fullData[98], fullData[81]^fullData[19], fullData[39]^fullData[84], fullData[114]-fullData[113], fullData[54]+fullData[37], fullData[70]^fullData[66], fullData[100]-fullData[86], fullData[64]-fullData[4], fullData[29]^fullData[72], fullData[91]+fullData[79], fullData[87]+fullData[6], fullData[20]^fullData[40], fullData[16]^fullData[14], fullData[1]^fullData[56], fullData[43]^fullData[23], fullData[122]^fullData[34], fullData[123]+fullData[69], fullData[30]+fullData[110], fullData[102]-fullData[103], fullData[74]+fullData[35], fullData[11]+fullData[80], fullData[68]+fullData[117], fullData[115]-fullData[26], fullData[55]+fullData[13], fullData[108]-fullData[7], fullData[57]+fullData[52], fullData[95]+fullData[111], fullData[9]+fullData[90], fullData[107]+fullData[22], fullData[89]^fullData[93], fullData[45]-fullData[3], fullData[58]+fullData[105], fullData[36]+fullData[31], fullData[101]-fullData[50], fullData[78]-fullData[118], fullData[44]+fullData[5], fullData[28]+fullData[83], fullData[97]^fullData[12], fullData[41]+fullData[99], fullData[53]-fullData[27], fullData[47]+fullData[92], fullData[59]-fullData[65])
		return  /*line ZzJ9TpjU.go:1*/ string(data)
	}(), dAMBvbjT)

	c7Phsypr, xonUbr_n :=  /*line PfRzaIYw.go:1*/ _IAEAzGw(vhZhSyYv, aarIVocK)
	if xonUbr_n != nil {
		return  /*line FFJub2UZ.go:1*/ shim.Error( /*line nVSXedsc.go:1*/ xonUbr_n.Error())
	}
	return  /*line olfJTGx_.go:1*/ shim.Success(c7Phsypr)

}

func (kuuXUPaZ *JPyzqDxq) z0P02F2H(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	if  /*line NsM2_9_9.go:1*/ len(y5yvGZQA) != 1 {
		return  /*line v5uo5V5R.go:1*/ shim.Error( /*line FtmuVOBI.go:1*/ func() string {
			data :=  /*line FtmuVOBI.go:1*/ []byte("\xaa\xac\xad\x93`gf\x9agva\x94\x9b\xb7m[\x8cLt.jABfe\xde}\xd3hg\xaeMa\x17[s\x06ryn\xb2m\x8e")
			positions := [...]byte{1, 13, 30, 31, 38, 21, 16, 6, 40, 42, 27, 7, 37, 8, 15, 0, 9, 31, 33, 28, 2, 4, 31, 11, 3, 12, 2, 15, 0, 9, 23, 15, 20, 7, 34, 22, 28, 38, 37, 5, 0, 16, 25, 30, 9, 36, 0, 2, 15, 1, 4, 40, 7, 26, 18, 31, 21, 17, 16, 2}
			for i := 0; i < 60; i += 2 {
				localKey :=  /*line FtmuVOBI.go:1*/ byte(i) +  /*line FtmuVOBI.go:1*/ byte(positions[i]^positions[i+1]) + 205
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line FtmuVOBI.go:1*/ string(data)
		}())
	}
	fdBHw0R_ := y5yvGZQA[0]

	dAMBvbjT, xonUbr_n :=  /*line kyti2qFi.go:1*/ kuuXUPaZ.glBEweME(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line QF_xOTzE.go:1*/ shim.Error( /*line zjwzFzwy.go:1*/ xonUbr_n.Error())
	}

	aarIVocK :=  /*line QtB1_2l0.go:1*/ fmt.Sprintf( /*line AeDyRvfg.go:1*/ func() string {
		data :=  /*line AeDyRvfg.go:1*/ []byte("e\"\x0fele\xbcPo`$9H\xcf\x18o\xecY\x9fap\xeb\x0f\"Met\u007fd2tL\x053 T\feern1\x96E\x02f\xa5\xdeov|d0rK\x1a\x8e%x[Lh\x04\xa5>]\ba\xc4RN2me\r:<%rv}}")
		positions := [...]byte{54, 65, 42, 6, 63, 55, 44, 65, 22, 66, 23, 21, 9, 67, 31, 31, 46, 63, 64, 13, 46, 65, 16, 21, 14, 58, 17, 35, 43, 74, 79, 54, 10, 35, 61, 78, 68, 64, 46, 22, 47, 56, 74, 58, 63, 79, 64, 43, 66, 58, 65, 12, 18, 31, 50, 20, 69, 62, 66, 17, 71, 29, 76, 0, 10, 32, 27, 6, 19, 2, 46, 42, 7, 78, 55, 52, 43, 66, 42, 60, 68, 17, 59, 36, 19, 78, 14, 12, 22, 32, 37, 21, 41, 16, 12, 33, 11, 52, 74, 66}
		for i := 0; i < 100; i += 2 {
			localKey :=  /*line AeDyRvfg.go:1*/ byte(i) +  /*line AeDyRvfg.go:1*/ byte(positions[i]^positions[i+1]) + 189
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line AeDyRvfg.go:1*/ string(data)
	}(), dAMBvbjT, fdBHw0R_)

	c7Phsypr, xonUbr_n :=  /*line ZP6CYjkZ.go:1*/ _IAEAzGw(vhZhSyYv, aarIVocK)
	if xonUbr_n != nil {
		return  /*line gmiKqf1e.go:1*/ shim.Error( /*line HWubHs0N.go:1*/ xonUbr_n.Error())
	}
	return  /*line qdQ1MTM2.go:1*/ shim.Success(c7Phsypr)

}

func (kuuXUPaZ *JPyzqDxq) pxhwNbVy(vhZhSyYv shim.ChaincodeStubInterface) pb.Response {

	lohinFr6, xonUbr_n :=  /*line DBsxyJ_4.go:1*/ kuuXUPaZ.glBEweME(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line pgxHWJmy.go:1*/ shim.Error( /*line S9zgnJil.go:1*/ xonUbr_n.Error())
	}

	aarIVocK :=  /*line r3leGJfu.go:1*/ fmt.Sprintf( /*line oggvlzPp.go:1*/ func() string {
		fullData :=  /*line oggvlzPp.go:1*/ []byte("zO\xd8S\xed\x17\xeb\x98ӹy\xa3\xba!\x82I/\x80V\x1c\xa4y\x88m\xeb5\xc6O\x15V\xa4\x14T3\a)0>?\x93\x9ck|\b\xd7ǖ\xbb0C\xf3Q\xf8\x81\xc5\xf3\xa7\xda4\f\xf5\x81s;jO\x05\xe4\x1c9\xd1͓Ț\xa4\xda'\x18\xbd\x12\xe2\x88l,.~\xf7\v\xf4\xf3\xc2*\x1e>ݥ\x80ÙR\xc5d\n\xbe\xb8\x1c$\xff\x9a\xf9\xb4\xa4P`\xb9\x8a\xc9")
		var data []byte
		data =  /*line oggvlzPp.go:1*/ append(data, fullData[22]+fullData[90], fullData[94]^fullData[106], fullData[48]+fullData[49], fullData[51]+fullData[31], fullData[114]-fullData[89], fullData[46]^fullData[50], fullData[5]-fullData[111], fullData[98]-fullData[27], fullData[3]+fullData[19], fullData[34]+fullData[41], fullData[6]^fullData[117], fullData[24]+fullData[1], fullData[115]+fullData[91], fullData[105]^fullData[74], fullData[57]+fullData[116], fullData[28]^fullData[0], fullData[33]^fullData[113], fullData[8]+fullData[61], fullData[58]-fullData[47], fullData[99]-fullData[35], fullData[84]-fullData[45], fullData[82]+fullData[109], fullData[2]^fullData[81], fullData[108]-fullData[95], fullData[12]+fullData[39], fullData[40]^fullData[110], fullData[9]^fullData[71], fullData[63]-fullData[76], fullData[72]^fullData[87], fullData[30]^fullData[54], fullData[21]-fullData[66], fullData[79]+fullData[75], fullData[55]+fullData[16], fullData[107]-fullData[52], fullData[101]-fullData[96], fullData[11]-fullData[53], fullData[60]^fullData[97], fullData[92]+fullData[15], fullData[36]+fullData[25], fullData[83]^fullData[93], fullData[100]-fullData[67], fullData[88]^fullData[64], fullData[4]^fullData[17], fullData[59]-fullData[56], fullData[44]^fullData[7], fullData[14]-fullData[68], fullData[70]+fullData[86], fullData[37]+fullData[69], fullData[42]^fullData[80], fullData[43]^fullData[23], fullData[73]-fullData[18], fullData[103]+fullData[78], fullData[10]-fullData[38], fullData[102]+fullData[104], fullData[29]^fullData[62], fullData[77]^fullData[32], fullData[26]-fullData[112], fullData[13]-fullData[20], fullData[65]+fullData[85])
		return  /*line oggvlzPp.go:1*/ string(data)
	}(), lohinFr6)

	c7Phsypr, xonUbr_n :=  /*line P3YSN9MG.go:1*/ _IAEAzGw(vhZhSyYv, aarIVocK)
	if xonUbr_n != nil {
		return  /*line BFIz1fSK.go:1*/ shim.Error( /*line AvmkIep3.go:1*/ xonUbr_n.Error())
	}
	return  /*line vuaNC_id.go:1*/ shim.Success(c7Phsypr)

}

func (kuuXUPaZ *JPyzqDxq) dwIdWa90(vhZhSyYv shim.ChaincodeStubInterface) pb.Response {
	var kgjcHTNY string

	dHsNol6l, xonUbr_n :=  /*line mvgN7qs9.go:1*/ kuuXUPaZ.yq7QUwKQ(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line TuEjJIKO.go:1*/ shim.Error( /*line gSfTLFI8.go:1*/ xonUbr_n.Error())
	}
	if dHsNol6l !=  /*line jdszESjm.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 207
		for counter := 0; i != 3; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				for y := range data {
					data[y] = data[y] -  /*line jdszESjm.go:1*/ byte(decryptKey^y)
				}
				i = 3
			case 4:
				i = 0
				data =  /*line jdszESjm.go:1*/ append(data, 60)
			case 6:
				data =  /*line jdszESjm.go:1*/ append(data, 50)
				i = 4
			case 5:
				i = 1
				data =  /*line jdszESjm.go:1*/ append(data, 47)
			case 2:
				i = 5
				data =  /*line jdszESjm.go:1*/ append(data, 43)
			case 1:
				data =  /*line jdszESjm.go:1*/ append(data, 53)
				i = 6
			}
		}
		return  /*line jdszESjm.go:1*/ string(data)
	}() {
		return  /*line Er8KhOuI.go:1*/ shim.Error( /*line UqFeUGtk.go:1*/ func() string {
			seed :=  /*line UqFeUGtk.go:1*/ byte(248)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line UqFeUGtk.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/  /*line UqFeUGtk.go:1*/ fnc(71)(173)(88)(189)(33)(145)(37)(63)(120)(253)(245)(251)(221)(205)(143)(36)(71)(64)(193)(133)(19)(34)(73)(68)(203)(148)(53)(28)(137)(22)(28)(69)(145)(201)(225)(197)(127)(248)(253)(245)(251)(221)(205)(143)(36)(71)(129)(13)(206)(235)(222)(179)(93)(199)(60)(197)(130)(19)(19)(41)(79)(177)(79)
			return  /*line UqFeUGtk.go:1*/ string(data)
		}())
	}
	                         
	jVqGRYFt, xonUbr_n :=  /*line zjysfVTC.go:1*/ kuuXUPaZ.iiTHmzLy(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line erYDUvYt.go:1*/ shim.Error( /*line nw1RzsB0.go:1*/ xonUbr_n.Error())
	}
	if jVqGRYFt ==  /*line Tbq_t0VD.go:1*/ func() string {
		data :=  /*line Tbq_t0VD.go:1*/ []byte(":E13\x00")
		positions := [...]byte{3, 1, 2, 4, 0, 2}
		for i := 0; i < 6; i += 2 {
			localKey :=  /*line Tbq_t0VD.go:1*/ byte(i) +  /*line Tbq_t0VD.go:1*/ byte(positions[i]^positions[i+1]) + 44
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line Tbq_t0VD.go:1*/ string(data)
	}() {
		kgjcHTNY, xonUbr_n =  /*line CO7uWFAb.go:1*/ kuuXUPaZ.c_afFRlN(vhZhSyYv)
		if xonUbr_n != nil {
			return  /*line v4Wmr3fm.go:1*/ shim.Error( /*line spOXyzdF.go:1*/ xonUbr_n.Error())
		}

	} else if jVqGRYFt ==  /*line REcyuZHy.go:1*/ func() string {
		key :=  /*line REcyuZHy.go:1*/ []byte("\xb3l\x8c\x0e")
		data :=  /*line REcyuZHy.go:1*/ []byte("\xc1\x06\xe9W")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line REcyuZHy.go:1*/ string(data)
	}() {
		kgjcHTNY, xonUbr_n =  /*line JLGqPevp.go:1*/ kuuXUPaZ.a8t6_82N(vhZhSyYv)
		if xonUbr_n != nil {
			return  /*line WJS38uUp.go:1*/ shim.Error( /*line EVvGzG09.go:1*/ xonUbr_n.Error())
		}
	} else {
		return  /*line C2LQt0vl.go:1*/ shim.Error( /*line bh5F0Ezh.go:1*/ func() string {
			seed :=  /*line bh5F0Ezh.go:1*/ byte(13)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line bh5F0Ezh.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/  /*line bh5F0Ezh.go:1*/ fnc(88)(11)(27)(229)(31)(248)(233)(80)(169)(7)(4)(91)(170)(1)(14)(168)(69)(6)(30)(235)
			return  /*line bh5F0Ezh.go:1*/ string(data)
		}())
	}

	aarIVocK :=  /*line VT76H3EJ.go:1*/ fmt.Sprintf( /*line sfkF6RVe.go:1*/ func() string {
		fullData :=  /*line sfkF6RVe.go:1*/ []byte("\xba-\xe5\xe1`\xbd\x96\x8a鉮=LXf_zF~R\xaa\xaa\xa7/\x8f\x9b\xc8\xfe\x8dg\xfe\x00\x8cG\xf1;\ue171g\r\x1dQ\xe8\x05lL\xa7\xd5s}\x144t)\xf8t\xbd(\xb8\x03pe-}\xc9\x01ퟌv\x1b\xb0bV$~\x1c:'\xaf\t\xd3\xe3\x12QE\xdc\x00\uf730K\xa8\x93/\x81\xa83\x92\x9d\x18S\x81\xaf\xb9MN")
		var data []byte
		data =  /*line sfkF6RVe.go:1*/ append(data, fullData[29]^fullData[77], fullData[57]-fullData[25], fullData[105]+fullData[0], fullData[42]+fullData[51], fullData[36]+fullData[18], fullData[84]+fullData[102], fullData[66]+fullData[73], fullData[93]^fullData[87], fullData[83]+fullData[69], fullData[86]-fullData[82], fullData[50]^fullData[15], fullData[8]-fullData[104], fullData[37]^fullData[30], fullData[80]-fullData[28], fullData[94]-fullData[23], fullData[11]^fullData[19], fullData[40]-fullData[21], fullData[59]+fullData[90], fullData[55]+fullData[96], fullData[81]+fullData[39], fullData[4]+fullData[44], fullData[107]^fullData[45], fullData[38]+fullData[9], fullData[53]^fullData[74], fullData[106]-fullData[88], fullData[71]^fullData[76], fullData[34]-fullData[64], fullData[68]^fullData[27], fullData[92]^fullData[95], fullData[63]^fullData[46], fullData[58]+fullData[12], fullData[97]-fullData[33], fullData[85]^fullData[49], fullData[56]^fullData[13], fullData[5]^fullData[100], fullData[99]-fullData[61], fullData[52]+fullData[35], fullData[6]-fullData[75], fullData[16]^fullData[41], fullData[14]^fullData[54], fullData[103]+fullData[2], fullData[47]^fullData[43], fullData[79]-fullData[91], fullData[3]^fullData[24], fullData[31]+fullData[62], fullData[22]^fullData[48], fullData[98]+fullData[89], fullData[7]+fullData[72], fullData[101]^fullData[78], fullData[26]^fullData[67], fullData[70]-fullData[60], fullData[10]^fullData[32], fullData[17]-fullData[65], fullData[20]-fullData[1])
		return  /*line sfkF6RVe.go:1*/ string(data)
	}(), kgjcHTNY)

	c7Phsypr, xonUbr_n :=  /*line kMbvFEdg.go:1*/ _IAEAzGw(vhZhSyYv, aarIVocK)
	if xonUbr_n != nil {
		return  /*line Bz9q536y.go:1*/ shim.Error( /*line uCl3ZliH.go:1*/ xonUbr_n.Error())
	}
	return  /*line QSD_C9Un.go:1*/ shim.Success(c7Phsypr)

}

func (kuuXUPaZ *JPyzqDxq) cuM598CU(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {

	fdBHw0R_ := y5yvGZQA[0]
	kgjcHTNY, xonUbr_n :=  /*line yps10Gsm.go:1*/ kuuXUPaZ.c_afFRlN(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line HzpQFZ6M.go:1*/ shim.Error( /*line RpKKP5DO.go:1*/ xonUbr_n.Error())
	}

	aarIVocK :=  /*line kpdlQ7qC.go:1*/ fmt.Sprintf( /*line vOjFqgqI.go:1*/ func() string {
		fullData :=  /*line vOjFqgqI.go:1*/ []byte("\xae=\xe6\x8cv\x8ez\xd6BS\x18\xd4j\x0e7\x89\x90\xdb\\d\x03\xbb\x8d\xed\x19\xf2\xf5h}\x99\x8f冏\xaft\xe26G\xaa\x1e\x8b<\xf0݈\x13\nDj\xca\f\xfb\x85\x96\xe4#\x99\xec\xbe\xf6\x1c[\xa1\x81\x95V\xee ^=\x9c\xbc\xa7u\x96\xaa\xd8x\xb4\xde\xc5q4\xfa\x11ށ\xcdt:M\xfd\x14\xe5oT\xa8En\x9f\xa9\x9e\xea1F\x1fϦfs\x119\xe3=\x04\xa7\t┲|ؠ\vU@\xf9\x88\xb6\xa4+.2:\xc2o\xbedD\x14\xfa\xa1\xbb\x0f\xf4\xce\x03")
		var data []byte
		data =  /*line vOjFqgqI.go:1*/ append(data, fullData[27]+fullData[46], fullData[40]+fullData[115], fullData[104]+fullData[8], fullData[66]+fullData[144], fullData[44]-fullData[82], fullData[16]^fullData[26], fullData[32]-fullData[56], fullData[10]+fullData[18], fullData[107]^fullData[123], fullData[28]-fullData[124], fullData[39]+fullData[78], fullData[108]+fullData[119], fullData[14]-fullData[72], fullData[97]+fullData[6], fullData[142]^fullData[81], fullData[141]-fullData[41], fullData[86]+fullData[53], fullData[120]-fullData[69], fullData[114]^fullData[48], fullData[111]-fullData[63], fullData[87]^fullData[55], fullData[51]^fullData[132], fullData[77]^fullData[118], fullData[146]^fullData[58], fullData[145]-fullData[73], fullData[59]+fullData[116], fullData[33]+fullData[94], fullData[0]-fullData[91], fullData[43]-fullData[3], fullData[42]-fullData[17], fullData[143]-fullData[38], fullData[84]-fullData[57], fullData[109]-fullData[139], fullData[130]^fullData[45], fullData[64]+fullData[100], fullData[29]^fullData[21], fullData[22]+fullData[36], fullData[2]-fullData[35], fullData[135]-fullData[62], fullData[80]-fullData[30], fullData[50]-fullData[138], fullData[65]-fullData[105], fullData[5]^fullData[127], fullData[122]+fullData[75], fullData[12]+fullData[52], fullData[74]+fullData[92], fullData[1]+fullData[31], fullData[113]-fullData[101], fullData[79]+fullData[99], fullData[11]-fullData[34], fullData[49]+fullData[117], fullData[88]+fullData[125], fullData[126]-fullData[140], fullData[7]-fullData[129], fullData[76]-fullData[128], fullData[23]^fullData[15], fullData[98]+fullData[61], fullData[47]-fullData[54], fullData[95]-fullData[13], fullData[4]-fullData[20], fullData[89]^fullData[85], fullData[102]^fullData[103], fullData[133]^fullData[121], fullData[25]+fullData[136], fullData[96]^fullData[112], fullData[134]+fullData[131], fullData[67]+fullData[83], fullData[70]-fullData[147], fullData[37]-fullData[93], fullData[106]^fullData[90], fullData[68]^fullData[9], fullData[137]+fullData[19], fullData[24]-fullData[71], fullData[110]-fullData[60])
		return  /*line vOjFqgqI.go:1*/ string(data)
	}(), kgjcHTNY, fdBHw0R_)

	c7Phsypr, xonUbr_n :=  /*line ue2QBq4L.go:1*/ _IAEAzGw(vhZhSyYv, aarIVocK)
	if xonUbr_n != nil {
		return  /*line YN5v8B1C.go:1*/ shim.Error( /*line dJ7iWBix.go:1*/ xonUbr_n.Error())
	}
	return  /*line e_XT2yAY.go:1*/ shim.Success(c7Phsypr)

}

func (kuuXUPaZ *JPyzqDxq) yxMJzmq_(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	if  /*line MLt40tMu.go:1*/ len(y5yvGZQA) != 1 {
		return  /*line GIk2nsbS.go:1*/ shim.Error( /*line zbSt7z9b.go:1*/ func() string {
			fullData :=  /*line zbSt7z9b.go:1*/ []byte("\x14\xd0N\x06X\x93@Y^\nA)\xbe\xa7\xbc8\x00\x87?[Q\xb3,\xe3\b\x8d\xf2\x84Oru\xd3 \xf4\xfb\xdc\bኟ\xd1\x16\\\xe2+\xee\xb2\xdb\xe5n\xdd7;Oj_\x006\xd3\xefX\x95\x1b=U.A\xfa\x87\xcd\x18?Br9Ŏ\x86wz)Yi\x9f\xe1j")
			var data []byte
			data =  /*line zbSt7z9b.go:1*/ append(data, fullData[13]-fullData[8], fullData[49]-fullData[16], fullData[35]+fullData[17], fullData[10]^fullData[65], fullData[47]-fullData[82], fullData[51]+fullData[52], fullData[55]-fullData[67], fullData[36]+fullData[19], fullData[64]-fullData[37], fullData[23]+fullData[63], fullData[66]+fullData[32], fullData[56]-fullData[76], fullData[11]^fullData[2], fullData[45]+fullData[68], fullData[42]-fullData[59], fullData[34]+fullData[85], fullData[18]-fullData[40], fullData[69]-fullData[7], fullData[60]+fullData[62], fullData[20]+fullData[50], fullData[21]^fullData[5], fullData[27]-fullData[71], fullData[72]+fullData[57], fullData[79]-fullData[9], fullData[1]+fullData[61], fullData[53]^fullData[22], fullData[74]-fullData[75], fullData[80]+fullData[6], fullData[84]+fullData[25], fullData[3]-fullData[83], fullData[38]-fullData[54], fullData[4]-fullData[33], fullData[41]^fullData[78], fullData[26]^fullData[77], fullData[58]-fullData[73], fullData[48]-fullData[29], fullData[12]-fullData[81], fullData[39]-fullData[44], fullData[24]+fullData[70], fullData[46]+fullData[14], fullData[30]^fullData[0], fullData[28]-fullData[43], fullData[15]-fullData[31])
			return  /*line zbSt7z9b.go:1*/ string(data)
		}())
	}
	fdBHw0R_ := y5yvGZQA[0]
	kgjcHTNY, xonUbr_n :=  /*line O7LKKrKW.go:1*/ kuuXUPaZ.c_afFRlN(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line ENPbSzCc.go:1*/ shim.Error( /*line E8agseyQ.go:1*/ xonUbr_n.Error())
	}

	lohinFr6, xonUbr_n :=  /*line JZkVN4lV.go:1*/ kuuXUPaZ.glBEweME(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line ytkI7JNp.go:1*/ shim.Error( /*line TWkatYJ1.go:1*/ xonUbr_n.Error())
	}

	aarIVocK :=  /*line aw7UdQLo.go:1*/ fmt.Sprintf( /*line GLRKPNPL.go:1*/ func() string {
		fullData :=  /*line GLRKPNPL.go:1*/ []byte("c\x11\x8b\xc5\x1f,?՝<\xaf\xc1K\x93\x80;\xa2\xcd\xc7\x11/\xfa\xd30\x8a\x91\xe7\x97\x18P݃\x95F3\xa7\x95\xad\xfc\x1ek\xf6\x1f\x1b?\uf804\xd00\x967~\xd0\xd8\xfaH\xac\xf6?\x90~\xd0\xff\xf1\x85\xb21ߚ\x14\r1\xdbl\xf9\xe3\x12\xa34Mz\x1fB\t,\xceP\xa1\xea\xa8\xcb\xfeO\xf7phb\xc7\xdb\x0eR\x18\xfa\xa8\xf2\xb7|\xddQk\xfb\xb5h\x8eP\x95h\xf2+&X\xabS\x15\x88\xc9J\xb8,r\x8c\x03\xbc\x11U\xcb}YS\xf2<\xf6\x1aM\x98\xd7{\x0e\xbb\\\xdd^^j\xae\x93\xae\xb1\x8c\xf0\x13<\xf5\xdf\x19\x8a\xb6\xaa\aR$0\xd0\xfe\xc63P1\xd1_\xb5?[\xe0\xa6\xdd@\xf2\xcc\xe9\x99q&\x03\x85")
		var data []byte
		data =  /*line GLRKPNPL.go:1*/ append(data, fullData[107]+fullData[63], fullData[71]+fullData[124], fullData[50]+fullData[30], fullData[178]+fullData[79], fullData[159]^fullData[184], fullData[167]+fullData[10], fullData[11]-fullData[152], fullData[175]+fullData[157], fullData[25]^fullData[92], fullData[174]-fullData[131], fullData[110]+fullData[106], fullData[195]+fullData[112], fullData[135]+fullData[193], fullData[137]-fullData[183], fullData[68]^fullData[149], fullData[9]-fullData[17], fullData[77]+fullData[109], fullData[84]-fullData[181], fullData[145]-fullData[4], fullData[173]^fullData[46], fullData[116]-fullData[23], fullData[158]^fullData[13], fullData[99]-fullData[88], fullData[148]+fullData[70], fullData[14]-fullData[176], fullData[41]^fullData[156], fullData[192]+fullData[194], fullData[89]^fullData[2], fullData[64]^fullData[36], fullData[19]+fullData[177], fullData[27]+fullData[151], fullData[182]^fullData[153], fullData[130]-fullData[87], fullData[187]^fullData[74], fullData[72]^fullData[1], fullData[103]^fullData[54], fullData[108]^fullData[104], fullData[136]^fullData[128], fullData[18]-fullData[97], fullData[133]^fullData[86], fullData[35]+fullData[98], fullData[143]^fullData[147], fullData[52]^fullData[161], fullData[33]+fullData[42], fullData[155]-fullData[180], fullData[134]-fullData[122], fullData[102]+fullData[51], fullData[171]+fullData[139], fullData[141]^fullData[101], fullData[127]^fullData[20], fullData[93]-fullData[186], fullData[15]^fullData[165], fullData[59]+fullData[111], fullData[58]+fullData[85], fullData[190]+fullData[162], fullData[191]-fullData[120], fullData[170]+fullData[62], fullData[90]^fullData[47], fullData[6]^fullData[82], fullData[166]-fullData[96], fullData[69]-fullData[119], fullData[125]^fullData[55], fullData[16]+fullData[3], fullData[39]+fullData[67], fullData[91]^fullData[37], fullData[189]+fullData[31], fullData[65]+fullData[118], fullData[43]+fullData[123], fullData[32]^fullData[160], fullData[80]^fullData[44], fullData[138]+fullData[126], fullData[34]-fullData[75], fullData[5]+fullData[142], fullData[105]^fullData[146], fullData[73]-fullData[117], fullData[66]^fullData[60], fullData[164]+fullData[144], fullData[95]-fullData[29], fullData[172]+fullData[140], fullData[154]^fullData[100], fullData[169]-fullData[185], fullData[56]+fullData[129], fullData[38]^fullData[8], fullData[113]-fullData[163], fullData[21]+fullData[40], fullData[28]+fullData[150], fullData[132]+fullData[12], fullData[45]^fullData[114], fullData[48]-fullData[0], fullData[26]+fullData[61], fullData[81]-fullData[121], fullData[24]-fullData[115], fullData[188]^fullData[53], fullData[83]+fullData[76], fullData[78]-fullData[49], fullData[94]^fullData[7], fullData[22]+fullData[168], fullData[57]+fullData[179])
		return  /*line GLRKPNPL.go:1*/ string(data)
	}(), lohinFr6, kgjcHTNY, fdBHw0R_)

	c7Phsypr, xonUbr_n :=  /*line BAi6yGln.go:1*/ _IAEAzGw(vhZhSyYv, aarIVocK)
	if xonUbr_n != nil {
		return  /*line mfuNCYjZ.go:1*/ shim.Error( /*line x1p1hPQG.go:1*/ xonUbr_n.Error())
	}
	return  /*line HAefpm23.go:1*/ shim.Success(c7Phsypr)

}

func (kuuXUPaZ *JPyzqDxq) gaj_brT8(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	if  /*line ep0t9Eqp.go:1*/ len(y5yvGZQA) != 1 {
		return  /*line JvN2Mm4P.go:1*/ shim.Error( /*line wLHuXn0i.go:1*/ func() string {
			seed :=  /*line wLHuXn0i.go:1*/ byte(180)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line wLHuXn0i.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/  /*line wLHuXn0i.go:1*/ fnc(149)(37)(245)(12)(3)(0)(243)(254)(17)(172)(65)(17)(245)(14)(248)(248)(9)(6)(255)(187)(242)(37)(51)(248)(245)(254)(17)(245)(5)(249)(185)(73)(251)(1)(9)(6)(245)(253)(3)(252)(13)
			return  /*line wLHuXn0i.go:1*/ string(data)
		}())
	}
	xwCUJUQo := y5yvGZQA[0]

	aarIVocK :=  /*line ptEm84jy.go:1*/ fmt.Sprintf( /*line fND0I6aT.go:1*/ func() string {
		seed :=  /*line fND0I6aT.go:1*/ byte(87)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line fND0I6aT.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/  /*line fND0I6aT.go:1*/ fnc(210)(75)(231)(192)(135)(7)(12)(41)(77)(157)(234)(236)(25)(217)(244)(243)(218)(165)(111)(213)(159)(251)(14)(4)(51)(126)(11)(3)(9)(15)(49)(79)(95)(200)(132)(10)(91)(177)(99)(207)(164)(61)(119)(241)(222)(201)(66)(156)(32)(67)(212)(87)(9)(18)
		return  /*line fND0I6aT.go:1*/ string(data)
	}(), xwCUJUQo)

	c7Phsypr, xonUbr_n :=  /*line SknvEwBa.go:1*/ _IAEAzGw(vhZhSyYv, aarIVocK)
	if xonUbr_n != nil {
		return  /*line Ij_Rmbi_.go:1*/ shim.Error( /*line XkRpUYh5.go:1*/ xonUbr_n.Error())
	}
	return  /*line Qumfj88v.go:1*/ shim.Success(c7Phsypr)

}

func (kuuXUPaZ *JPyzqDxq) kbgtOUnn(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	if  /*line FL5tGup3.go:1*/ len(y5yvGZQA) != 1 {
		return  /*line BfwtL_3W.go:1*/ shim.Error( /*line QR0Os8rq.go:1*/ func() string {
			fullData :=  /*line QR0Os8rq.go:1*/ []byte("|\x91\nñ\xb7\x12\xcd\x0f\x99\x01\xdb\xd4\x0f\xef:\x97T\xc1\xe1\x04\xaa:b`\x1e$\xb2\xac\xdbT\x13\xba\x84ֈ\x8e\xa0\xadL\x9cŭ+\x12\x9c\x8a\xaa\x81\xaeOFů\x948/\xc7.\xab\xba\x81E\x02\n\x92\xb6\xbc\xe49\xbf\xeb\x04$C\\\xb1'\xbc!y/3\xaa\u007f\xc9")
			var data []byte
			data =  /*line QR0Os8rq.go:1*/ append(data, fullData[40]+fullData[38], fullData[59]^fullData[52], fullData[65]-fullData[81], fullData[82]^fullData[75], fullData[24]^fullData[6], fullData[17]+fullData[25], fullData[85]+fullData[45], fullData[27]+fullData[76], fullData[63]-fullData[36], fullData[56]-fullData[8], fullData[13]-fullData[49], fullData[7]^fullData[70], fullData[69]+fullData[58], fullData[2]^fullData[84], fullData[21]^fullData[57], fullData[47]-fullData[62], fullData[0]^fullData[44], fullData[29]+fullData[9], fullData[78]+fullData[5], fullData[53]-fullData[61], fullData[26]^fullData[20], fullData[54]-fullData[50], fullData[80]^fullData[10], fullData[48]+fullData[14], fullData[37]^fullData[41], fullData[46]-fullData[77], fullData[60]-fullData[51], fullData[30]-fullData[71], fullData[4]-fullData[74], fullData[35]-fullData[79], fullData[12]+fullData[39], fullData[43]-fullData[32], fullData[1]+fullData[68], fullData[19]^fullData[33], fullData[28]-fullData[22], fullData[3]+fullData[66], fullData[73]-fullData[72], fullData[64]-fullData[16], fullData[34]-fullData[23], fullData[15]+fullData[55], fullData[31]-fullData[83], fullData[42]+fullData[18], fullData[67]^fullData[11])
			return  /*line QR0Os8rq.go:1*/ string(data)
		}())
	}
	aarIVocK := y5yvGZQA[0]
	c7Phsypr, xonUbr_n :=  /*line AK_MCnsP.go:1*/ _IAEAzGw(vhZhSyYv, aarIVocK)
	if xonUbr_n != nil {
		return  /*line vprtztMY.go:1*/ shim.Error( /*line C1hkJVVz.go:1*/ xonUbr_n.Error())
	}
	return  /*line LGZ8H014.go:1*/ shim.Success(c7Phsypr)

}
