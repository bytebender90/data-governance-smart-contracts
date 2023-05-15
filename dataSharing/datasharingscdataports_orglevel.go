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

func (g4rnrSNn *G1Y_7pz_) ry2azAMI(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var cSW1wSO3 error
	var w5VhCLka string
	if  /*line oXf5f8vT.go:1*/ len(ktsi1_SQ) != 5 {
		return  /*line GHNieb_d.go:1*/ shim.Error( /*line X61YStKL.go:1*/ func() string {
			data :=  /*line X61YStKL.go:1*/ []byte("\xe1n\\\x05*i| \x0e\x0eg9m\a1tl\x01uuf\x18lW;Q\xf6\n/uhctAoC(r:qu\x15D-\xfb=cm\x02ud\x87O:g.d\x95\"\x04e\xa6t\xd4O{\x9a^")
			positions := [...]byte{65, 61, 48, 19, 24, 45, 38, 21, 14, 33, 0, 27, 13, 67, 48, 65, 44, 59, 51, 65, 28, 36, 35, 4, 11, 19, 59, 38, 63, 58, 26, 58, 58, 42, 36, 3, 43, 41, 3, 63, 16, 27, 9, 16, 53, 17, 2, 47, 35, 22, 13, 42, 50, 28, 6, 49, 64, 27, 16, 8, 41, 6, 63, 35, 45, 53, 63, 26, 43, 33, 30, 4, 57, 23, 13, 66, 22, 24, 25, 24, 45, 56, 3, 20, 67, 66, 63, 23}
			for i := 0; i < 88; i += 2 {
				localKey :=  /*line X61YStKL.go:1*/ byte(i) +  /*line X61YStKL.go:1*/ byte(positions[i]^positions[i+1]) + 156
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line X61YStKL.go:1*/ string(data)
		}())
	}
	cVj_Mj6s := ktsi1_SQ[0]
	eviEA3sp := ktsi1_SQ[1]
	olYMzSWK := ktsi1_SQ[2]
	pbLLcfzS := ktsi1_SQ[3]
	ezgSKu0t :=  /*line I5ysVivx.go:1*/ strings.Split(ktsi1_SQ[4],  /*line TJPNFXUf.go:1*/ func() string {
		key :=  /*line TJPNFXUf.go:1*/ []byte("\xc7")
		data :=  /*line TJPNFXUf.go:1*/ []byte("e")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line TJPNFXUf.go:1*/ string(data)
	}())
	mzjt9rLK :=  /*line jsf_0nzq.go:1*/ func() string {
		key :=  /*line jsf_0nzq.go:1*/ []byte("8ӯ\x93\x03\x9d\xccAײT\xacЗ\x9f4@\xc3")
		data :=  /*line jsf_0nzq.go:1*/ []byte("\x8a8 \bh\x10@\x82:\x15\xb9\x1fC\xd9\x18\x83\xb2*")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line jsf_0nzq.go:1*/ string(data)
	}()

	sM8l5NjA, cSW1wSO3 :=  /*line SQoMpyG0.go:1*/ g4rnrSNn.lu_3VIbC(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line DqrtWIYQ.go:1*/ shim.Error( /*line ZYpiZLXw.go:1*/ cSW1wSO3.Error())
	}
	if sM8l5NjA ==  /*line l67CGt_K.go:1*/ func() string {
		fullData :=  /*line l67CGt_K.go:1*/ []byte("\x97\fY{\xb2\x06%&")
		var data []byte
		data =  /*line l67CGt_K.go:1*/ append(data, fullData[7]-fullData[4], fullData[0]-fullData[6], fullData[3]-fullData[5], fullData[2]+fullData[1])
		return  /*line l67CGt_K.go:1*/ string(data)
	}() {
		return  /*line HWH9jrB2.go:1*/ shim.Error( /*line xRVnHRnf.go:1*/ func() string {
			key :=  /*line xRVnHRnf.go:1*/ []byte("\xb2\x13M\x1f\xb8Ȯ\x9fВ\xfc\x1f\xd8\xdff\x11\x12\v,lu\x17c\xa4\xb6,\xcbˈ\xb5\x94\xfer\xefG\xf8P m\xb2\x86\xea\xc4I\xed0^\x169\xbem")
			data :=  /*line xRVnHRnf.go:1*/ []byte("\xf7\x8b\xc1\x84*6\x0f\v\xf0\ao\x84JR\x86tsy\x9a\xdb\xe97\xd5\t'\xa10>\xfc\xd5\x03p\xd9P\xb5aʁ\xe1\x1b\xf5X%\xb5\r\x91\xc1y\x9e1\xe0")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line xRVnHRnf.go:1*/ string(data)
		}())
	}

	nTXZkyTM, cSW1wSO3 :=  /*line zAVxkpGB.go:1*/ g4rnrSNn.dHINeXgX(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line nCT5xoCq.go:1*/ shim.Error( /*line Etzojxtc.go:1*/ cSW1wSO3.Error())
	}
	if nTXZkyTM !=  /*line jk7RW1Kz.go:1*/ func() string {
		key :=  /*line jk7RW1Kz.go:1*/ []byte("\xc0T\xd0\x01\xaa")
		data :=  /*line jk7RW1Kz.go:1*/ []byte("\xa1\x10\x9dh\xc4")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line jk7RW1Kz.go:1*/ string(data)
	}() {
		return  /*line HjI5h8Hj.go:1*/ shim.Error( /*line CmUdxXKS.go:1*/ func() string {
			fullData :=  /*line CmUdxXKS.go:1*/ []byte(">&\xa5\x91\x179BRĞO\xb1\xb1ǐ˧\x89U\xe6*\b\xdc\x05\xd6\xeb\f\\\x0f\xe7\xb4\xea;\x8e\x96HZ\xb8-sUS\xfb\xe03\x06Yk\x19%\xff\xe7\xf3n\x12\x93\x9e}\xc2\f\xa7\x0e\"]L^#\n\xae^\xb8U\xeb[\x19\x1e4\x981џy\xf35\x1b\xc0\x86\xbe\x02hTL?\xfb\xb2\x18\x19\xf9jz\x85\xee7Wi}æHl")
			var data []byte
			data =  /*line CmUdxXKS.go:1*/ append(data, fullData[89]+fullData[29], fullData[83]^fullData[73], fullData[90]+fullData[95], fullData[87]^fullData[13], fullData[106]+fullData[63], fullData[97]^fullData[34], fullData[21]+fullData[98], fullData[69]^fullData[5], fullData[45]^fullData[1], fullData[59]+fullData[40], fullData[7]-fullData[101], fullData[24]-fullData[104], fullData[33]^fullData[51], fullData[46]-fullData[25], fullData[27]+fullData[8], fullData[84]-fullData[37], fullData[103]+fullData[67], fullData[39]-fullData[23], fullData[62]^fullData[88], fullData[28]^fullData[109], fullData[96]^fullData[47], fullData[16]-fullData[6], fullData[64]^fullData[38], fullData[49]-fullData[12], fullData[42]^fullData[9], fullData[41]-fullData[44], fullData[55]+fullData[22], fullData[74]-fullData[60], fullData[105]+fullData[31], fullData[56]^fullData[50], fullData[36]^fullData[76], fullData[17]+fullData[43], fullData[48]-fullData[80], fullData[26]+fullData[18], fullData[61]^fullData[99], fullData[92]+fullData[20], fullData[57]^fullData[54], fullData[2]-fullData[102], fullData[66]-fullData[58], fullData[30]-fullData[108], fullData[11]^fullData[3], fullData[82]+fullData[53], fullData[100]^fullData[19], fullData[52]-fullData[14], fullData[68]^fullData[15], fullData[79]-fullData[65], fullData[75]+fullData[71], fullData[77]^fullData[70], fullData[0]^fullData[91], fullData[72]-fullData[86], fullData[85]-fullData[10], fullData[107]-fullData[78], fullData[4]-fullData[94], fullData[32]^fullData[35], fullData[81]+fullData[93])
			return  /*line CmUdxXKS.go:1*/ string(data)
		}())
	}

	vWNUn3N4, cSW1wSO3 :=  /*line AoXmdnA6.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line nhuLv8nu.go:1*/ shim.Error( /*line szQLxqiW.go:1*/ cSW1wSO3.Error())
	}

	yRT2gJYk := eviEA3sp +  /*line IUd2ug5e.go:1*/ func() string {
		fullData :=  /*line IUd2ug5e.go:1*/ []byte("\a(")
		var data []byte
		data =  /*line IUd2ug5e.go:1*/ append(data, fullData[1]+fullData[0])
		return  /*line IUd2ug5e.go:1*/ string(data)
	}() + vWNUn3N4
	                                
	dSS2AaxE, cSW1wSO3 :=  /*line HdVVVGB3.go:1*/ n4c7mRtG.GetState(yRT2gJYk)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line zz9Ps348.go:1*/ func() string {
			seed :=  /*line zz9Ps348.go:1*/ byte(157)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line zz9Ps348.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/  /*line zz9Ps348.go:1*/ fnc(222)(167)(35)(45)(0)(253)(3)(176)(24)(232)(36)(27)(8)(3)(249)(255)(188)(84)(251)(177)(67)(5)(253)(254)(8)(181)(65)(20)(255)(244)(7)(3)(247)(17)(231)(19)(245)(6)(255)(178)(70)(9)(3)(174)(84)(244)(253)(187)(68)(253)(19)(237)(18)(242)(15)(198)(230)
			return  /*line zz9Ps348.go:1*/ string(data)
		}() + eviEA3sp +  /*line KtImTKKs.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 65
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 2
					data =  /*line KtImTKKs.go:1*/ append(data, 223)
				case 0:
					for y := range data {
						data[y] = data[y] +  /*line KtImTKKs.go:1*/ byte(decryptKey^y)
					}
					i = 1
				case 2:
					i = 0
					data =  /*line KtImTKKs.go:1*/ append(data, 59)
				}
			}
			return  /*line KtImTKKs.go:1*/ string(data)
		}()
		return  /*line cMNLzry0.go:1*/ shim.Error(w5VhCLka)
	} else if (dSS2AaxE == nil) && (olYMzSWK == "" || olYMzSWK ==  /*line k3JTrZ8C.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 107
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				i = 1
				for y := range data {
					data[y] = data[y] ^  /*line k3JTrZ8C.go:1*/ byte(decryptKey^y)
				}
			case 2:
				data =  /*line k3JTrZ8C.go:1*/ append(data, 75)
				i = 0
			}
		}
		return  /*line k3JTrZ8C.go:1*/ string(data)
	}()) {
		w5VhCLka =  /*line xbEuC8pK.go:1*/ func() string {
			data :=  /*line xbEuC8pK.go:1*/ []byte("\xcd\"\x99\x0fror\xfc\x9a\xdf\xe7\xc0\xb5g\xe3_\x99gX\xb2\xc8i\x8b\tTis.\xa4\x96\xd4 d\xc9f\x89NW\xefБDaQ(e\x9e3t|\xcafQ8\xf7 inҖ\aR\u007f[ed\xef \xa3\x97z8\xf2xd v\x81~ \xb8\xcet\xc3s\xb2t\x87\x8f")
			positions := [...]byte{38, 47, 59, 85, 3, 3, 7, 18, 23, 83, 29, 60, 40, 72, 18, 33, 28, 66, 47, 76, 20, 58, 61, 73, 80, 62, 77, 12, 2, 46, 9, 29, 66, 8, 74, 16, 62, 10, 30, 69, 69, 71, 65, 54, 37, 27, 70, 16, 22, 81, 38, 22, 39, 37, 14, 0, 52, 70, 13, 71, 44, 66, 29, 69, 0, 62, 24, 76, 50, 11, 12, 78, 68, 83, 17, 73, 40, 17, 11, 17, 29, 19, 7, 36, 28, 59, 71, 33, 73, 15, 36, 59, 47, 65, 16, 73, 87, 72, 41, 50, 88, 76, 70, 22, 66, 46, 30, 49, 71, 53, 52, 78, 28, 38, 39, 43, 35, 59, 41, 63}
			for i := 0; i < 120; i += 2 {
				localKey :=  /*line xbEuC8pK.go:1*/ byte(i) +  /*line xbEuC8pK.go:1*/ byte(positions[i]^positions[i+1]) + 95
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line xbEuC8pK.go:1*/ string(data)
		}() + eviEA3sp +  /*line VuZybsln.go:1*/ func() string {
			data :=  /*line VuZybsln.go:1*/ []byte("\xbdb")
			positions := [...]byte{0, 1}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line VuZybsln.go:1*/ byte(i) +  /*line VuZybsln.go:1*/ byte(positions[i]^positions[i+1]) + 191
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line VuZybsln.go:1*/ string(data)
		}()
		return  /*line HpaO0EJX.go:1*/ shim.Error(w5VhCLka)
	}

	b4dMB1zL, cSW1wSO3 :=  /*line eubTJW21.go:1*/ jmSY84co(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line Wx1aaV6n.go:1*/ shim.Error( /*line KBrq7R14.go:1*/ cSW1wSO3.Error())
	}

	zzhdvyqZ := &DatasetMetadataResponse{}
	cSW1wSO3 =  /*line Rq3oXDpL.go:1*/ json.Unmarshal( /*line EO45BFZs.go:1*/ []byte(b4dMB1zL), zzhdvyqZ)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line EFCQFy_B.go:1*/ func() string {
			var data []byte
			i := 13
			decryptKey := 203
			for counter := 0; i != 14; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 5:
					data =  /*line EFCQFy_B.go:1*/ append(data, "_gb`"...,
					)
					i = 4
				case 2:
					i = 12
					data =  /*line EFCQFy_B.go:1*/ append(data, "oge@"...,
					)
				case 12:
					i = 7
					data =  /*line EFCQFy_B.go:1*/ append(data, 147)
				case 0:
					data =  /*line EFCQFy_B.go:1*/ append(data, "\x80\x80"...,
					)
					i = 10
				case 15:
					i = 6
					data =  /*line EFCQFy_B.go:1*/ append(data, ")L"...,
					)
				case 6:
					data =  /*line EFCQFy_B.go:1*/ append(data, "fm"...,
					)
					i = 2
				case 13:
					data =  /*line EFCQFy_B.go:1*/ append(data, 139)
					i = 3
				case 8:
					data =  /*line EFCQFy_B.go:1*/ append(data, 124)
					i = 1
				case 3:
					data =  /*line EFCQFy_B.go:1*/ append(data, "1S\u007f~"...,
					)
					i = 9
				case 4:
					for y := range data {
						data[y] = data[y] +  /*line EFCQFy_B.go:1*/ byte(decryptKey^y)
					}
					i = 14
				case 11:
					i = 15
					data =  /*line EFCQFy_B.go:1*/ append(data, 66)
				case 10:
					data =  /*line EFCQFy_B.go:1*/ append(data, "}\x88"...,
					)
					i = 8
				case 1:
					i = 5
					data =  /*line EFCQFy_B.go:1*/ append(data, "|6"...,
					)
				case 9:
					data =  /*line EFCQFy_B.go:1*/ append(data, "z|+"...,
					)
					i = 11
				case 7:
					data =  /*line EFCQFy_B.go:1*/ append(data, "\x8d="...,
					)
					i = 0
				}
			}
			return  /*line EFCQFy_B.go:1*/ string(data)
		}() + eviEA3sp +  /*line QNz8z87o.go:1*/ func() string {
			data :=  /*line QNz8z87o.go:1*/ []byte("E}")
			positions := [...]byte{0, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line QNz8z87o.go:1*/ byte(i) +  /*line QNz8z87o.go:1*/ byte(positions[i]^positions[i+1]) + 35
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line QNz8z87o.go:1*/ string(data)
		}()
		return  /*line ynzQSLIt.go:1*/ shim.Error(w5VhCLka)
	}

	                                                                  
	if  /*line nhohvVt2.go:1*/ nSRaDwQg(ezgSKu0t, zzhdvyqZ.CustomAccessRights) {
		 /*line DlNFE1zH.go:1*/ fmt.Println( /*line GArEzkMr.go:1*/ func() string {
			key :=  /*line GArEzkMr.go:1*/ []byte("\b%\xf28Q\x99\x86a?\x9c \xa9\xc0\xf2Z\x8c\xe3-\x89$i,,\x86\xbeI\xcb.\x89\b\x03\b")
			data :=  /*line GArEzkMr.go:1*/ []byte("I\x95b\xaa\xc0\t\xf8ʠ\x10\x85\xc9#g\xcd\x00R\x9a\xa9\x85̏\x91\xf91i=\x97\xf0pw{")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line GArEzkMr.go:1*/ string(data)
		}())
	} else {
		w5VhCLka =  /*line UllAEaO_.go:1*/ func() string {
			fullData :=  /*line UllAEaO_.go:1*/ []byte("!a\x93E\xe4]\xe6\x9d\xff\xda\xeb^\xa7\x13\x81ҙ\x1b\nM\x1fi\xca\xcc\xd4L|\x06\x1a\xdd\xc8\xdf\xecWƙ~\xbcmr&R\x8bl\xfda;\x1eUō\uf126}l\xf9\xe0\x96\xd9F\xa1\x16\xeb-\x03\xa1\xfc\xefɍc\" x\x00")
			var data []byte
			data =  /*line UllAEaO_.go:1*/ append(data, fullData[9]^fullData[61], fullData[75]+fullData[72], fullData[36]^fullData[46], fullData[7]^fullData[51], fullData[35]+fullData[59], fullData[43]^fullData[65], fullData[1]^fullData[13], fullData[3]+fullData[29], fullData[20]+fullData[17], fullData[34]^fullData[4], fullData[55]+fullData[63], fullData[42]^fullData[56], fullData[10]-fullData[26], fullData[15]^fullData[37], fullData[6]+fullData[14], fullData[24]+fullData[25], fullData[23]-fullData[21], fullData[74]+fullData[44], fullData[64]^fullData[11], fullData[48]^fullData[0], fullData[12]^fullData[30], fullData[70]^fullData[57], fullData[2]+fullData[50], fullData[18]+fullData[33], fullData[54]^fullData[47], fullData[49]^fullData[53], fullData[45]-fullData[67], fullData[16]-fullData[40], fullData[68]+fullData[52], fullData[38]-fullData[19], fullData[73]+fullData[41], fullData[8]^fullData[58], fullData[60]-fullData[31], fullData[66]^fullData[69], fullData[27]^fullData[39], fullData[62]+fullData[5], fullData[32]-fullData[22], fullData[28]+fullData[71])
			return  /*line UllAEaO_.go:1*/ string(data)
		}()
		return  /*line APXEExah.go:1*/ shim.Error(w5VhCLka)
	}

	_ra7bbD0, cSW1wSO3 :=  /*line BRqU9yzX.go:1*/ g4rnrSNn.y4dHt0L7(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line knRK9K3f.go:1*/ shim.Error( /*line BPFzBXVw.go:1*/ cSW1wSO3.Error())
	}

	b2UcdGWD, cSW1wSO3 :=  /*line Gb6hXS8T.go:1*/ g4rnrSNn.hDKHUUS5(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line p6jWY1ev.go:1*/ shim.Error( /*line M7ge_b1d.go:1*/ cSW1wSO3.Error())
	}

	aZUvqwea, cSW1wSO3 :=  /*line WYfcUnSV.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line z45Jj28b.go:1*/ shim.Error( /*line iRw8KRqh.go:1*/ cSW1wSO3.Error())
	}

	                                                   
	qFOnXzGg, cSW1wSO3 :=  /*line R8DJgfv4.go:1*/ fvIQVpVM(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line J57EIX1Q.go:1*/ shim.Error( /*line nTizV8jj.go:1*/ cSW1wSO3.Error())
	}

	                                          
	lYVRCpWS, cSW1wSO3 :=  /*line YheHe5_1.go:1*/ grz7gdfX(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line LGCZrui0.go:1*/ shim.Error( /*line IQbqa0lm.go:1*/ cSW1wSO3.Error())
	}

	                                            
	w0Qu5bcv, cSW1wSO3 :=  /*line XDz__kle.go:1*/ imxpv0H5(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line qlGfVNPX.go:1*/ shim.Error( /*line N8W7MSld.go:1*/ cSW1wSO3.Error())
	}

	aZuYW6KY := RequestAccessByOrg{mzjt9rLK, cVj_Mj6s, eviEA3sp, qFOnXzGg, w0Qu5bcv,
		lYVRCpWS, olYMzSWK, ezgSKu0t, vWNUn3N4, pbLLcfzS, aZUvqwea, _ra7bbD0, b2UcdGWD, nTXZkyTM}
	hvFCuRzl, cSW1wSO3 :=  /*line G3kqWVJn.go:1*/ json.Marshal(aZuYW6KY)
	if cSW1wSO3 != nil {
		return  /*line _rzDuE3L.go:1*/ shim.Error( /*line BT3AIC7E.go:1*/ cSW1wSO3.Error())
	}

	                        
	cSW1wSO3 =  /*line OtPDkk78.go:1*/ n4c7mRtG.PutState(cVj_Mj6s, hvFCuRzl)
	if cSW1wSO3 != nil {
		return  /*line FTh0CuQW.go:1*/ shim.Error( /*line JZITJM87.go:1*/ cSW1wSO3.Error())
	}

	xHTOVzVs :=  /*line vhR6ePUz.go:1*/ func() string {
		fullData :=  /*line vhR6ePUz.go:1*/ []byte("\\@\x18\x98\xd3\x00\xe4\x99'\xa4\xd5\xec\x9bK5\x1a_ߵms\x89]}\xf2+\xbe\x0fu\xe3\xec\xd0.d\x82V\xa6\xf7[:\x89\xdb\xf6\x95\a|\x81\x05\xba\xd8")
		var data []byte
		data =  /*line vhR6ePUz.go:1*/ append(data, fullData[35]^fullData[27], fullData[4]-fullData[33], fullData[14]+fullData[1], fullData[19]+fullData[47], fullData[15]^fullData[39], fullData[28]^fullData[44], fullData[6]+fullData[46], fullData[31]-fullData[16], fullData[30]^fullData[7], fullData[0]-fullData[37], fullData[36]^fullData[10], fullData[17]+fullData[43], fullData[23]^fullData[22], fullData[2]-fullData[9], fullData[48]+fullData[18], fullData[24]+fullData[32], fullData[21]+fullData[41], fullData[29]^fullData[34], fullData[12]-fullData[8], fullData[40]+fullData[49], fullData[20]^fullData[5], fullData[38]-fullData[42], fullData[3]^fullData[11], fullData[26]+fullData[45], fullData[13]-fullData[25])
		return  /*line vhR6ePUz.go:1*/ string(data)
	}() + eviEA3sp +  /*line wchpumft.go:1*/ func() string {
		data :=  /*line wchpumft.go:1*/ []byte("\x94ha\x8d \x9fG\xa1\x00`\xe3\x01bm\x1c\x1a\x9ee\x0e")
		positions := [...]byte{16, 10, 0, 6, 11, 14, 10, 18, 3, 18, 16, 9, 11, 9, 8, 15, 9, 5, 6, 16, 7, 9, 3, 18}
		for i := 0; i < 24; i += 2 {
			localKey :=  /*line wchpumft.go:1*/ byte(i) +  /*line wchpumft.go:1*/ byte(positions[i]^positions[i+1]) + 95
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line wchpumft.go:1*/ string(data)
	}()
	q8ZF8zaz :=  /*line D0FNOPdB.go:1*/ []byte(xHTOVzVs)
	return  /*line Gsf5eW3B.go:1*/ shim.Success(q8ZF8zaz)
}

func (g4rnrSNn *G1Y_7pz_) zJkLod3G(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	if  /*line rrHLPkDD.go:1*/ len(ktsi1_SQ) != 5 {
		return  /*line uN5PQxTC.go:1*/ shim.Error( /*line VCII6IEd.go:1*/ func() string {
			data :=  /*line VCII6IEd.go:1*/ []byte("\x80\x83O\x93li\xf4¸r\x99\x9bLv\xb2t\u007fNu\xb4\x1cAr\x9d\x03\xecM\xf2<uncQio\x94\xa7We\xa6,es\xce\xe9O\xe8\"\xfaYۭc\x8dsy\n\x0f\xd4咹\xd2]x\b\x81\u007f\xb0iF\xc7\xee\x1d")
			positions := [...]byte{28, 62, 0, 25, 26, 26, 71, 20, 67, 63, 16, 57, 6, 55, 27, 40, 37, 43, 23, 49, 21, 12, 73, 3, 68, 62, 21, 72, 25, 32, 53, 67, 46, 53, 36, 65, 7, 60, 25, 59, 7, 37, 12, 71, 16, 62, 28, 10, 64, 45, 49, 8, 48, 57, 36, 1, 2, 6, 64, 21, 56, 70, 28, 23, 37, 6, 19, 70, 47, 58, 24, 61, 39, 28, 66, 35, 39, 72, 11, 51, 23, 60, 44, 43, 70, 12, 19, 10, 28, 7, 8, 26, 10, 66, 13, 66, 3, 14, 50, 43, 43, 58, 21, 27, 28, 27, 70, 45, 48, 61}
			for i := 0; i < 110; i += 2 {
				localKey :=  /*line VCII6IEd.go:1*/ byte(i) +  /*line VCII6IEd.go:1*/ byte(positions[i]^positions[i+1]) + 66
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line VCII6IEd.go:1*/ string(data)
		}())
	}
	cVj_Mj6s := ktsi1_SQ[0]
	eviEA3sp := ktsi1_SQ[1]
	olYMzSWK := ktsi1_SQ[2]
	pbLLcfzS := ktsi1_SQ[3]
	ezgSKu0t :=  /*line E967sfh3.go:1*/ strings.Split(ktsi1_SQ[4],  /*line Gzv5o6o4.go:1*/ func() string {
		data :=  /*line Gzv5o6o4.go:1*/ []byte("\x9a")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line Gzv5o6o4.go:1*/ byte(i) +  /*line Gzv5o6o4.go:1*/ byte(positions[i]^positions[i+1]) + 146
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line Gzv5o6o4.go:1*/ string(data)
	}())
	mzjt9rLK :=  /*line H46bvgp1.go:1*/ func() string {
		seed :=  /*line H46bvgp1.go:1*/ byte(87)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line H46bvgp1.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/  /*line H46bvgp1.go:1*/ fnc(169)(101)(214)(176)(80)(174)(93)(122)(18)(55)(127)(247)(234)(206)(155)(19)(72)(144)(34)(82)(164)(23)(101)(160)(99)(187)
		return  /*line H46bvgp1.go:1*/ string(data)
	}()

	hGBwBh29, cSW1wSO3 :=  /*line z6tOpzEA.go:1*/ g4rnrSNn.dHINeXgX(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line HozzznQq.go:1*/ shim.Error( /*line _Flma2bL.go:1*/ cSW1wSO3.Error())
	}
	if hGBwBh29 !=  /*line UKhA33GS.go:1*/ func() string {
		key :=  /*line UKhA33GS.go:1*/ []byte("\x16\xcf\b-x")
		data :=  /*line UKhA33GS.go:1*/ []byte("K\x95e<\xf6")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line UKhA33GS.go:1*/ string(data)
	}() {
		return  /*line DOGNmw7A.go:1*/ shim.Error( /*line EvAawhmV.go:1*/ func() string {
			var data []byte
			i := 2
			decryptKey := 215
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 10:
					i = 15
					data =  /*line EvAawhmV.go:1*/ append(data, "\xa8\xe8\xed"...,
					)
				case 1:
					i = 11
					data =  /*line EvAawhmV.go:1*/ append(data, "\xcb\xdb\xce"...,
					)
				case 13:
					data =  /*line EvAawhmV.go:1*/ append(data, "ʂ\xc2"...,
					)
					i = 12
				case 8:
					data =  /*line EvAawhmV.go:1*/ append(data, "\xcfɄ"...,
					)
					i = 5
				case 2:
					data =  /*line EvAawhmV.go:1*/ append(data, "\xe5\xc5\xc4"...,
					)
					i = 0
				case 5:
					data =  /*line EvAawhmV.go:1*/ append(data, "\xc6\xdb"...,
					)
					i = 4
				case 11:
					data =  /*line EvAawhmV.go:1*/ append(data, "\xc9\xd8\xc1"...,
					)
					i = 7
				case 7:
					i = 3
					data =  /*line EvAawhmV.go:1*/ append(data, "ǐ\xc3"...,
					)
				case 4:
					i = 1
					data =  /*line EvAawhmV.go:1*/ append(data, "\u0558"...,
					)
				case 14:
					i = 6
					for y := range data {
						data[y] = data[y] ^  /*line EvAawhmV.go:1*/ byte(decryptKey^y)
					}
				case 15:
					data =  /*line EvAawhmV.go:1*/ append(data, "\xec\xe9\xfe\xf1"...,
					)
					i = 14
				case 0:
					i = 13
					data =  /*line EvAawhmV.go:1*/ append(data, "Ў\xc0\xde"...,
					)
				case 12:
					i = 8
					data =  /*line EvAawhmV.go:1*/ append(data, "\xc4\xcc"...,
					)
				case 3:
					i = 9
					data =  /*line EvAawhmV.go:1*/ append(data, "\xd3\xc1"...,
					)
				case 9:
					i = 10
					data =  /*line EvAawhmV.go:1*/ append(data, "\xdb\xde\xef\xef"...,
					)
				}
			}
			return  /*line EvAawhmV.go:1*/ string(data)
		}())
	}

	dJwgW2jL, cSW1wSO3 :=  /*line QyjcnNem.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line S5ldxNtN.go:1*/ shim.Error( /*line tXuP09pB.go:1*/ cSW1wSO3.Error())
	}

	yRT2gJYk := eviEA3sp +  /*line mVgCEIKg.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 62
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				data =  /*line mVgCEIKg.go:1*/ append(data, 17)
				i = 0
			case 0:
				for y := range data {
					data[y] = data[y] ^  /*line mVgCEIKg.go:1*/ byte(decryptKey^y)
				}
				i = 1
			}
		}
		return  /*line mVgCEIKg.go:1*/ string(data)
	}() + dJwgW2jL
	                                
	dSS2AaxE, cSW1wSO3 :=  /*line crAl803e.go:1*/ n4c7mRtG.GetState(yRT2gJYk)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line BDYvqbvI.go:1*/ func() string {
			data :=  /*line BDYvqbvI.go:1*/ []byte("\x17\xe8\xf2\xe0\x1c\xberE:\xe4\xd9\xe3\xfd\xea+\xe6\x02t\x01\xf8c7/\x04k K\x05\x9dh:\x0f\x02\xd1\\\x8d\xd5Q\xec\xbaP\n\xc2 \xb5De \aaB\xc9s\xba\xfdҴ")
			positions := [...]byte{53, 35, 44, 55, 18, 35, 5, 31, 41, 33, 15, 28, 38, 4, 48, 32, 53, 51, 13, 15, 44, 30, 39, 9, 56, 2, 36, 16, 11, 10, 27, 37, 21, 54, 4, 41, 11, 36, 42, 7, 31, 14, 28, 34, 45, 26, 35, 0, 16, 2, 1, 7, 53, 37, 35, 15, 12, 21, 7, 3, 9, 19, 1, 23, 1, 55, 40, 50, 22, 21}
			for i := 0; i < 70; i += 2 {
				localKey :=  /*line BDYvqbvI.go:1*/ byte(i) +  /*line BDYvqbvI.go:1*/ byte(positions[i]^positions[i+1]) + 128
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line BDYvqbvI.go:1*/ string(data)
		}() + eviEA3sp +  /*line x0z3Qsjz.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 131
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 0
					data =  /*line x0z3Qsjz.go:1*/ append(data, 163)
				case 0:
					i = 1
					data =  /*line x0z3Qsjz.go:1*/ append(data, 253)
				case 1:
					for y := range data {
						data[y] = data[y] -  /*line x0z3Qsjz.go:1*/ byte(decryptKey^y)
					}
					i = 2
				}
			}
			return  /*line x0z3Qsjz.go:1*/ string(data)
		}()
		return  /*line izlnzWgv.go:1*/ shim.Error(w5VhCLka)
	} else if dSS2AaxE != nil {
		hnzizKM2 := AuthorizedOrgs{}
		cSW1wSO3 =  /*line fz4su771.go:1*/ json.Unmarshal(dSS2AaxE, &hnzizKM2)	                               
		if cSW1wSO3 != nil {
			return  /*line QWa7xeuT.go:1*/ shim.Error( /*line _h6OA1CS.go:1*/ cSW1wSO3.Error())
		}
		if  /*line Uhy_1qIs.go:1*/ nSRaDwQg(ezgSKu0t, hnzizKM2.CustomAccessRights) {
			 /*line AQjVmJxF.go:1*/ fmt.Println( /*line HqYZqnCQ.go:1*/ func() string {
				var data []byte
				i := 9
				decryptKey := 177
				for counter := 0; i != 11; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 2:
						data =  /*line HqYZqnCQ.go:1*/ append(data, "\x93\x90"...,
						)
						i = 14
					case 15:
						i = 16
						data =  /*line HqYZqnCQ.go:1*/ append(data, 163)
					case 4:
						i = 12
						data =  /*line HqYZqnCQ.go:1*/ append(data, 138)
					case 5:
						data =  /*line HqYZqnCQ.go:1*/ append(data, "\xa8\x9e\x9d\xaf"...,
						)
						i = 3
					case 14:
						i = 0
						data =  /*line HqYZqnCQ.go:1*/ append(data, 66)
					case 9:
						data =  /*line HqYZqnCQ.go:1*/ append(data, "u\xa3\xa2"...,
						)
						i = 6
					case 16:
						i = 2
						data =  /*line HqYZqnCQ.go:1*/ append(data, "\xb4\xb1\xb1"...,
						)
					case 10:
						data =  /*line HqYZqnCQ.go:1*/ append(data, 151)
						i = 8
					case 13:
						data =  /*line HqYZqnCQ.go:1*/ append(data, "\x98\x9fK\x9c"...,
						)
						i = 7
					case 8:
						i = 1
						data =  /*line HqYZqnCQ.go:1*/ append(data, "\xa2\xa0"...,
						)
					case 6:
						data =  /*line HqYZqnCQ.go:1*/ append(data, "\xa3\xa7\xa7"...,
						)
						i = 5
					case 1:
						for y := range data {
							data[y] = data[y] +  /*line HqYZqnCQ.go:1*/ byte(decryptKey^y)
						}
						i = 11
					case 12:
						i = 13
						data =  /*line HqYZqnCQ.go:1*/ append(data, 139)
					case 0:
						data =  /*line HqYZqnCQ.go:1*/ append(data, "\x82\x8b"...,
						)
						i = 4
					case 3:
						i = 15
						data =  /*line HqYZqnCQ.go:1*/ append(data, "\x9fY"...,
						)
					case 7:
						i = 10
						data =  /*line HqYZqnCQ.go:1*/ append(data, "\x92\x97"...,
						)
					}
				}
				return  /*line HqYZqnCQ.go:1*/ string(data)
			}())
		} else {
			w5VhCLka =  /*line CooZU8_o.go:1*/ func() string {
				seed :=  /*line CooZU8_o.go:1*/ byte(91)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data =  /*line CooZU8_o.go:1*/ append(data, x^seed); seed += x; return fnc }
				 /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/  /*line CooZU8_o.go:1*/ fnc(32)(89)(145)(23)(14)(229)(29)(174)(0)(24)(5)(37)(19)(225)(23)(167)(77)(14)(250)(247)(21)(226)(81)(163)(6)(8)(22)(250)(240)(83)(180)(19)(234)(31)(226)(11)(161)(89)
				return  /*line CooZU8_o.go:1*/ string(data)
			}()
			return  /*line E8HazQEj.go:1*/ shim.Error(w5VhCLka)
		}
	}

	_ra7bbD0, cSW1wSO3 :=  /*line upmD6quy.go:1*/ g4rnrSNn.y4dHt0L7(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line zrxjkpd0.go:1*/ shim.Error( /*line G0vqapp4.go:1*/ cSW1wSO3.Error())
	}

	b2UcdGWD, cSW1wSO3 :=  /*line lVuKccUh.go:1*/ g4rnrSNn.hDKHUUS5(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line V3oYYp6Z.go:1*/ shim.Error( /*line E_uq7K5G.go:1*/ cSW1wSO3.Error())
	}

	aZUvqwea, cSW1wSO3 :=  /*line U5B_VQXK.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line ivAvu2hV.go:1*/ shim.Error( /*line BwzVBlAB.go:1*/ cSW1wSO3.Error())
	}

	j0BZ3Inj := AuthorizedOrgs{}
	syHhHGpa := eviEA3sp +  /*line I1tb6Vxc.go:1*/ func() string {
		data :=  /*line I1tb6Vxc.go:1*/ []byte("\xc7")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line I1tb6Vxc.go:1*/ byte(i) +  /*line I1tb6Vxc.go:1*/ byte(positions[i]^positions[i+1]) + 104
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line I1tb6Vxc.go:1*/ string(data)
	}() + dJwgW2jL
	ze01ef7p, cSW1wSO3 :=  /*line NIjw3yzN.go:1*/ n4c7mRtG.GetState(syHhHGpa)
	if cSW1wSO3 != nil {
		return  /*line OQNfbV4l.go:1*/ shim.Error( /*line MywuEcuJ.go:1*/ func() string {
			key :=  /*line MywuEcuJ.go:1*/ []byte("\xa2\xf0\xddMֶ\x11L\xe2\x17h=\xda\xdcy\x8d\xe14kk\x1ct\xd2+\xf4\b\x8a\x85\x1d\xb5\x85?N\x85")
			data :=  /*line MywuEcuJ.go:1*/ []byte("䑴!\xb3\xd218\x8d7\x0fX\xae\xfc\x16\xff\x86U\x05\x02f\x15\xa6B\x9bf\xaa\xe4~\xd6\xe0L=\xbf")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line MywuEcuJ.go:1*/ string(data)
		}() +  /*line G5C6s8PN.go:1*/ cSW1wSO3.Error())
	} else if ze01ef7p == nil {
		return  /*line BA6ULxTq.go:1*/ shim.Error( /*line VWP3BgF9.go:1*/ func() string {
			key :=  /*line VWP3BgF9.go:1*/ []byte("\v7̪c\xe5\xac\xd3>@\x96(\xf5\x94\x02\xc9\xf3\xd3k\x8f\x04\u139a\x1d\x9e\xd6ѻ\x120")
			data :=  /*line VWP3BgF9.go:1*/ []byte("I1\x99v\f\x8d\xbb\x8e0)\xe49\u007f\xd5m\xa5-\x96\b\x91j\x8e\xe6\x86Jԋ\x9d\xb9S4")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line VWP3BgF9.go:1*/ string(data)
		}())
	}

	cSW1wSO3 =  /*line LgAwC16y.go:1*/ json.Unmarshal(ze01ef7p, &j0BZ3Inj)	                               
	if cSW1wSO3 != nil {
		return  /*line KWMdM9IH.go:1*/ shim.Error( /*line UTCXkiYt.go:1*/ cSW1wSO3.Error())
	}

	if dJwgW2jL != j0BZ3Inj.Organization {
		return  /*line tRh4WlG1.go:1*/ shim.Error( /*line R9k_Vnqv.go:1*/ func() string {
			seed :=  /*line R9k_Vnqv.go:1*/ byte(176)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line R9k_Vnqv.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/  /*line R9k_Vnqv.go:1*/ fnc(165)(25)(243)(20)(255)(244)(7)(3)(247)(17)(235)(255)(188)(79)(3)(245)(250)(13)(251)(17)(231)(19)(245)(6)(255)
			return  /*line R9k_Vnqv.go:1*/ string(data)
		}())
	}

	zJkLod3G := &RequestRevokeAccessByOrg{mzjt9rLK, cVj_Mj6s, eviEA3sp, j0BZ3Inj.Dataset_Provider, olYMzSWK, ezgSKu0t, dJwgW2jL, pbLLcfzS, aZUvqwea, _ra7bbD0, b2UcdGWD, hGBwBh29}
	hIj4xsdd, cSW1wSO3 :=  /*line wa6JSzQX.go:1*/ json.Marshal(zJkLod3G)
	if cSW1wSO3 != nil {
		return  /*line KnWoNhcm.go:1*/ shim.Error( /*line Ds0AaOjG.go:1*/ cSW1wSO3.Error())
	}

	                                       
	cSW1wSO3 =  /*line G1IhBhnZ.go:1*/ n4c7mRtG.PutState(cVj_Mj6s, hIj4xsdd)
	if cSW1wSO3 != nil {
		return  /*line Z35utB49.go:1*/ shim.Error( /*line CwXctQ2U.go:1*/ cSW1wSO3.Error())
	}
	xHTOVzVs :=  /*line hr8rYeTa.go:1*/ func() string {
		data :=  /*line hr8rYeTa.go:1*/ []byte("Yr\xed\xf0 \x89-\xc2o\xa6e 3\xd1q3's\xc4\x1e\x1f-\xc7\xf5 idֆ")
		positions := [...]byte{27, 5, 9, 27, 13, 19, 22, 28, 1, 7, 2, 21, 18, 19, 12, 22, 1, 16, 7, 5, 9, 5, 3, 21, 27, 23, 6, 12, 15, 23, 20, 21}
		for i := 0; i < 32; i += 2 {
			localKey :=  /*line hr8rYeTa.go:1*/ byte(i) +  /*line hr8rYeTa.go:1*/ byte(positions[i]^positions[i+1]) + 151
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line hr8rYeTa.go:1*/ string(data)
	}() + cVj_Mj6s +  /*line PwsLomyw.go:1*/ func() string {
		key :=  /*line PwsLomyw.go:1*/ []byte("+z\xd6\x15g\xb8\xf6\x12\x8c\xda\x19\xee")
		data :=  /*line PwsLomyw.go:1*/ []byte("_\x15\xf6q\x06̗a\xe9\xae#\xce")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line PwsLomyw.go:1*/ string(data)
	}() + eviEA3sp +  /*line pAfPq1f4.go:1*/ func() string {
		key :=  /*line pAfPq1f4.go:1*/ []byte("\xf8\xa6\xc0\xd3h\xf0\xf4}q\xb3s\x9dm\xf5\r&\xa15C")
		data :=  /*line pAfPq1f4.go:1*/ []byte("\x18\x0e!F\x88RY\xe2\xdf\xd3\xe6\x12\xcfbv\x9a\x15\x9a\xa7")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line pAfPq1f4.go:1*/ string(data)
	}()
	q8ZF8zaz :=  /*line isJuEi7f.go:1*/ []byte(xHTOVzVs)
	return  /*line EYQFCmxq.go:1*/ shim.Success(q8ZF8zaz)

}

func (g4rnrSNn *G1Y_7pz_) yftWTcK6(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	if  /*line rJLVS1cR.go:1*/ len(ktsi1_SQ) != 2 {
		return  /*line WZjElzC2.go:1*/ shim.Error( /*line JclAxCrp.go:1*/ func() string {
			data :=  /*line JclAxCrp.go:1*/ []byte("\xaf\xe9SFrrb\x01t \xe1]\x18\x9de\xc8\b\x15\xf2 arg\xa4\xcce\r\xac\xf1\x93w\x97\x0ep\xea[\xaai\xe6\xfdp/\x0f\xaaR\x88w\xe1\xdfs\x14\xaf\x9e\xf2\xe2ﲗ\xee\v`\xee\xed4")
			positions := [...]byte{23, 47, 54, 56, 38, 30, 59, 56, 12, 38, 50, 41, 6, 45, 61, 1, 13, 42, 63, 7, 29, 3, 0, 2, 0, 27, 62, 6, 43, 46, 59, 45, 57, 32, 45, 18, 17, 23, 10, 34, 52, 53, 41, 7, 16, 11, 38, 24, 32, 17, 26, 35, 51, 60, 39, 58, 48, 15, 47, 24, 56, 51, 28, 2, 42, 36, 55, 35, 7, 60, 26, 3, 40, 30, 63, 30, 52, 41, 53, 52, 46, 0, 54, 31}
			for i := 0; i < 84; i += 2 {
				localKey :=  /*line JclAxCrp.go:1*/ byte(i) +  /*line JclAxCrp.go:1*/ byte(positions[i]^positions[i+1]) + 54
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line JclAxCrp.go:1*/ string(data)
		}())
	}

	hSHnto8d := ktsi1_SQ[0]
	pbLLcfzS := ktsi1_SQ[1]

	w8EOizcx, cSW1wSO3 :=  /*line jz1TynaB.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line Ddyv5pqk.go:1*/ shim.Error( /*line IDQ88Fdj.go:1*/ cSW1wSO3.Error())
	}

	mzjt9rLK :=  /*line u_YvjDQD.go:1*/ func() string {
		fullData :=  /*line u_YvjDQD.go:1*/ []byte("\xd6H[\ah\\4\xc0\xb2\x9c\x06j\x0f!bY\xeck\f\xca\xf9\x81@\xb8\x11\x1eb}")
		var data []byte
		data =  /*line u_YvjDQD.go:1*/ append(data, fullData[20]-fullData[23], fullData[21]-fullData[18], fullData[25]^fullData[11], fullData[19]-fullData[26], fullData[2]^fullData[6], fullData[8]-fullData[22], fullData[27]+fullData[16], fullData[15]+fullData[13], fullData[17]-fullData[10], fullData[7]-fullData[5], fullData[1]+fullData[3], fullData[0]+fullData[9], fullData[12]^fullData[4], fullData[24]+fullData[14])
		return  /*line u_YvjDQD.go:1*/ string(data)
	}()
	qo8GOXik := RequestAccessByOrg{}
	mGZijNgp, cSW1wSO3 :=  /*line MSVnv7Yq.go:1*/ n4c7mRtG.GetState(hSHnto8d)	                                      
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line kbaIggCB.go:1*/ func() string {
			fullData :=  /*line kbaIggCB.go:1*/ []byte("\xf4|\xa3\xfd,\x95H$\x89+\x11\x82*\x8d\xcej\xc3\x17\xa1Μ\xed\x8a\xcd7y\xe0b\xcf8\xa9y\xa7vMo\xa1\xd5|l\f\xad\x15\xa1{\xc7\xe8\x90=\b\xa8Y\x9b\x86\x9b\x8e{\xe5\xf6\x97\x1f\xff;\xc4r6\xacV\xec\xf6\xee;\x9bv\x85\xa5\xd66W\x04\xe6\xcc<\x01\x1f\xab")
			var data []byte
			data =  /*line kbaIggCB.go:1*/ append(data, fullData[39]^fullData[17], fullData[55]^fullData[66], fullData[16]+fullData[11], fullData[38]+fullData[69], fullData[53]+fullData[68], fullData[14]^fullData[43], fullData[20]-fullData[12], fullData[8]^fullData[85], fullData[30]-fullData[35], fullData[15]^fullData[6], fullData[70]^fullData[50], fullData[74]-fullData[7], fullData[75]-fullData[82], fullData[34]+fullData[60], fullData[33]-fullData[10], fullData[54]^fullData[61], fullData[0]+fullData[4], fullData[45]+fullData[41], fullData[57]^fullData[22], fullData[72]-fullData[56], fullData[52]+fullData[81], fullData[26]-fullData[44], fullData[83]-fullData[13], fullData[73]-fullData[67], fullData[2]+fullData[28], fullData[59]+fullData[19], fullData[65]+fullData[71], fullData[40]^fullData[25], fullData[62]-fullData[76], fullData[5]^fullData[80], fullData[42]-fullData[36], fullData[21]-fullData[23], fullData[49]+fullData[51], fullData[32]^fullData[63], fullData[79]-fullData[18], fullData[47]-fullData[9], fullData[48]+fullData[77], fullData[37]-fullData[27], fullData[78]-fullData[24], fullData[84]^fullData[31], fullData[64]+fullData[3], fullData[58]+fullData[1], fullData[46]+fullData[29])
			return  /*line kbaIggCB.go:1*/ string(data)
		}() + hSHnto8d +  /*line EzLzVgpa.go:1*/ func() string {
			var data []byte
			i := 0
			decryptKey := 180
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					i = 3
					data =  /*line EzLzVgpa.go:1*/ append(data, 204)
				case 3:
					for y := range data {
						data[y] = data[y] +  /*line EzLzVgpa.go:1*/ byte(decryptKey^y)
					}
					i = 1
				case 0:
					data =  /*line EzLzVgpa.go:1*/ append(data, 114)
					i = 2
				}
			}
			return  /*line EzLzVgpa.go:1*/ string(data)
		}()
		return  /*line In7NEOsy.go:1*/ shim.Error(w5VhCLka)
	} else if mGZijNgp == nil {
		w5VhCLka =  /*line CAQE2sBC.go:1*/ func() string {
			fullData :=  /*line CAQE2sBC.go:1*/ []byte("\x15\xf4\\'ā\xcbDX\xe0\x93\xb1\xec\xccbH!\\\xe0\x8d\x9f8D\xf7\x86?@\xe5\xf7\xfdV\xf4\x84\x8a\xcb\xd3ғ\u007f\n \xc2\x0f\xaa\xe1\xa8x\xd4d\x9bG\xc1d\xf9\xf6d\x1c\x05b\xe8\xed\x03O`\x16K\x9f\xeaC\xbb\x87\nE\xd5\xd6\xe4r?o\x8e\x88`\xfa\x18_\x10\x95\x98]4\x18(Ħ\xd8\xe3\xf3k\xb6\x8a\x80\x15 ̥\x8b*\xca\xf6i\x19\xa1<\x9b\x85\xa9B\x05\a\xed")
			var data []byte
			data =  /*line CAQE2sBC.go:1*/ append(data, fullData[59]+fullData[37], fullData[47]^fullData[108], fullData[88]-fullData[90], fullData[85]+fullData[58], fullData[74]-fullData[52], fullData[29]-fullData[79], fullData[119]^fullData[66], fullData[95]+fullData[25], fullData[111]^fullData[49], fullData[39]^fullData[91], fullData[68]^fullData[118], fullData[7]-fullData[73], fullData[2]-fullData[23], fullData[41]+fullData[11], fullData[106]-fullData[71], fullData[54]^fullData[87], fullData[80]-fullData[110], fullData[94]-fullData[55], fullData[17]+fullData[4], fullData[75]^fullData[5], fullData[84]^fullData[3], fullData[1]-fullData[105], fullData[28]-fullData[32], fullData[70]^fullData[96], fullData[89]+fullData[77], fullData[20]-fullData[38], fullData[22]-fullData[36], fullData[63]+fullData[117], fullData[112]-fullData[34], fullData[0]+fullData[81], fullData[27]-fullData[100], fullData[97]^fullData[83], fullData[103]-fullData[8], fullData[14]-fullData[116], fullData[109]-fullData[61], fullData[62]+fullData[102], fullData[92]^fullData[98], fullData[113]^fullData[69], fullData[15]+fullData[56], fullData[35]-fullData[76], fullData[57]+fullData[78], fullData[31]-fullData[10], fullData[65]^fullData[21], fullData[82]-fullData[86], fullData[53]^fullData[19], fullData[18]+fullData[26], fullData[40]-fullData[115], fullData[42]-fullData[93], fullData[99]-fullData[64], fullData[16]+fullData[50], fullData[51]^fullData[44], fullData[45]+fullData[107], fullData[114]+fullData[9], fullData[101]^fullData[48], fullData[6]-fullData[30], fullData[43]-fullData[72], fullData[60]+fullData[24], fullData[33]+fullData[67], fullData[104]^fullData[12], fullData[13]+fullData[46])
			return  /*line CAQE2sBC.go:1*/ string(data)
		}() + hSHnto8d +  /*line bc40lScv.go:1*/ func() string {
			data :=  /*line bc40lScv.go:1*/ []byte("ǘ")
			positions := [...]byte{0, 1}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line bc40lScv.go:1*/ byte(i) +  /*line bc40lScv.go:1*/ byte(positions[i]^positions[i+1]) + 185
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line bc40lScv.go:1*/ string(data)
		}()
		return  /*line PzE2qytf.go:1*/ shim.Error(w5VhCLka)
	}

	cSW1wSO3 =  /*line McxiS9Am.go:1*/ json.Unmarshal(mGZijNgp, &qo8GOXik)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line mzMhKUvF.go:1*/ func() string {
			data :=  /*line mzMhKUvF.go:1*/ []byte("\x8a\"\x8er\x95Br\xd2\xecc\x14\xea\x03\xee\xc1\xf5btG }\x0ec\xe5d\xf7 ݐ\xa5\xee")
			positions := [...]byte{16, 5, 2, 18, 13, 10, 4, 12, 5, 20, 21, 21, 23, 27, 2, 5, 9, 28, 4, 0, 9, 5, 29, 7, 16, 9, 4, 20, 23, 11, 16, 4, 18, 25, 0, 30, 12, 15, 14, 8, 28, 11}
			for i := 0; i < 42; i += 2 {
				localKey :=  /*line mzMhKUvF.go:1*/ byte(i) +  /*line mzMhKUvF.go:1*/ byte(positions[i]^positions[i+1]) + 77
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line mzMhKUvF.go:1*/ string(data)
		}() + hSHnto8d +  /*line ll0JW4os.go:1*/ func() string {
			data :=  /*line ll0JW4os.go:1*/ []byte("0\xd5")
			positions := [...]byte{0, 1}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line ll0JW4os.go:1*/ byte(i) +  /*line ll0JW4os.go:1*/ byte(positions[i]^positions[i+1]) + 178
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line ll0JW4os.go:1*/ string(data)
		}()
		return  /*line KPQ5ErFl.go:1*/ shim.Error(w5VhCLka)
	}
	eviEA3sp := qo8GOXik.Dataset_Name
	hcl4L2y6 := qo8GOXik.Dataset_Provider
	olYMzSWK := qo8GOXik.Permission
	ezgSKu0t := qo8GOXik.CustomAccessRights
	dJwgW2jL := qo8GOXik.Organization
	hOgpzUCn := qo8GOXik.Username
	ipXT6jpu := qo8GOXik.Name
	wSxUCaJ0 := qo8GOXik.Surname
	v5vcwK31 := qo8GOXik.Role
	deTJzFce := qo8GOXik.Description
	e5loS483 := qo8GOXik.DatasetProviderOrg

	                                            
	qFOnXzGg, cSW1wSO3 :=  /*line CPhfPFbZ.go:1*/ fvIQVpVM(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line IZR_DrE2.go:1*/ shim.Error( /*line sHakchNO.go:1*/ cSW1wSO3.Error())
	}
	if w8EOizcx != qFOnXzGg {
		w5VhCLka =  /*line SU3r2MBU.go:1*/ func() string {
			seed :=  /*line SU3r2MBU.go:1*/ byte(217)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line SU3r2MBU.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/  /*line SU3r2MBU.go:1*/ fnc(162)(89)(145)(23)(14)(229)(29)(174)(0)(24)(28)(1)(27)(170)(85)(252)(241)(30)(251)(253)(229)(11)(25)(241)(166)(89)(246)(30)(235)(190)(98)(132)
			return  /*line SU3r2MBU.go:1*/ string(data)
		}() + w8EOizcx +  /*line LJ01Gap0.go:1*/ func() string {
			data :=  /*line LJ01Gap0.go:1*/ []byte("for\r\x9c\\\x04as\x01\xdd\xc7X")
			positions := [...]byte{12, 4, 6, 5, 11, 10, 3, 10, 12, 4, 3, 6, 6, 9}
			for i := 0; i < 14; i += 2 {
				localKey :=  /*line LJ01Gap0.go:1*/ byte(i) +  /*line LJ01Gap0.go:1*/ byte(positions[i]^positions[i+1]) + 88
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line LJ01Gap0.go:1*/ string(data)
		}() + eviEA3sp +  /*line tUPYA81T.go:1*/ func() string {
			seed :=  /*line tUPYA81T.go:1*/ byte(64)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line tUPYA81T.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line tUPYA81T.go:1*/  /*line tUPYA81T.go:1*/ fnc(226)(91)
			return  /*line tUPYA81T.go:1*/ string(data)
		}()
		return  /*line EvLrpV8w.go:1*/ shim.Error(w5VhCLka)
	}
	syHhHGpa := eviEA3sp +  /*line z25geM1s.go:1*/ func() string {
		fullData :=  /*line z25geM1s.go:1*/ []byte("\xe4\xcb")
		var data []byte
		data =  /*line z25geM1s.go:1*/ append(data, fullData[0]^fullData[1])
		return  /*line z25geM1s.go:1*/ string(data)
	}() + dJwgW2jL

	b4dMB1zL, cSW1wSO3 :=  /*line yEzDvSr0.go:1*/ jmSY84co(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line E9Xg5cWb.go:1*/ shim.Error( /*line Ec_KgHDA.go:1*/ cSW1wSO3.Error())
	}

	zzhdvyqZ := &DatasetMetadataResponse{}
	cSW1wSO3 =  /*line rxq8wF__.go:1*/ json.Unmarshal( /*line Gtf_0xZV.go:1*/ []byte(b4dMB1zL), zzhdvyqZ)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line Ze9feIzB.go:1*/ func() string {
			seed :=  /*line Ze9feIzB.go:1*/ byte(25)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line Ze9feIzB.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/  /*line Ze9feIzB.go:1*/ fnc(98)(167)(35)(45)(0)(253)(3)(176)(24)(232)(36)(27)(8)(3)(249)(255)(188)(84)(251)(177)(68)(1)(254)(12)(245)(1)(187)(42)(9)(252)(255)
			return  /*line Ze9feIzB.go:1*/ string(data)
		}() + eviEA3sp +  /*line WVZ4ycAO.go:1*/ func() string {
			seed :=  /*line WVZ4ycAO.go:1*/ byte(93)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line WVZ4ycAO.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line WVZ4ycAO.go:1*/  /*line WVZ4ycAO.go:1*/ fnc(127)(161)
			return  /*line WVZ4ycAO.go:1*/ string(data)
		}()
		return  /*line s5K7grbg.go:1*/ shim.Error(w5VhCLka)
	}

	                                                                  
	if  /*line m2ua9_ol.go:1*/ nSRaDwQg(ezgSKu0t, zzhdvyqZ.CustomAccessRights) {
		 /*line N8oBLjU_.go:1*/ fmt.Println( /*line S0YWUyQM.go:1*/ func() string {
			data :=  /*line S0YWUyQM.go:1*/ []byte("Ap\xdfr\x85p<iace|hbstqYb\xceccrsdysigv\u007f_")
			positions := [...]byte{11, 19, 9, 30, 11, 2, 4, 2, 4, 4, 13, 24, 4, 26, 25, 26, 30, 6, 26, 31, 30, 17, 26, 18, 2, 12, 16, 29, 19, 22, 18, 17, 17, 30}
			for i := 0; i < 34; i += 2 {
				localKey :=  /*line S0YWUyQM.go:1*/ byte(i) +  /*line S0YWUyQM.go:1*/ byte(positions[i]^positions[i+1]) + 242
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line S0YWUyQM.go:1*/ string(data)
		}())
	} else {
		w5VhCLka =  /*line I7W7Sdgx.go:1*/ func() string {
			var data []byte
			i := 13
			decryptKey := 48
			for counter := 0; i != 15; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 14:
					i = 6
					data =  /*line I7W7Sdgx.go:1*/ append(data, "\xf7\x13\x15\x15"...,
					)
				case 6:
					i = 5
					data =  /*line I7W7Sdgx.go:1*/ append(data, 11)
				case 8:
					data =  /*line I7W7Sdgx.go:1*/ append(data, "-,"...,
					)
					i = 11
				case 2:
					data =  /*line I7W7Sdgx.go:1*/ append(data, "\x1d0+-"...,
					)
					i = 8
				case 3:
					i = 9
					data =  /*line I7W7Sdgx.go:1*/ append(data, "\xd7&\x1e\xf1"...,
					)
				case 7:
					for y := range data {
						data[y] = data[y] -  /*line I7W7Sdgx.go:1*/ byte(decryptKey^y)
					}
					i = 15
				case 10:
					i = 7
					data =  /*line I7W7Sdgx.go:1*/ append(data, "\xfc\xb0\f"...,
					)
				case 5:
					data =  /*line I7W7Sdgx.go:1*/ append(data, 197)
					i = 2
				case 1:
					data =  /*line I7W7Sdgx.go:1*/ append(data, "\x1e\x1e"...,
					)
					i = 12
				case 12:
					data =  /*line I7W7Sdgx.go:1*/ append(data, "\xcf\xdc\xc5"...,
					)
					i = 14
				case 0:
					data =  /*line I7W7Sdgx.go:1*/ append(data, "\x15$)"...,
					)
					i = 3
				case 9:
					i = 10
					data =  /*line I7W7Sdgx.go:1*/ append(data, "\xf3\xfc"...,
					)
				case 13:
					i = 4
					data =  /*line I7W7Sdgx.go:1*/ append(data, "%\xcd\xed"...,
					)
				case 4:
					data =  /*line I7W7Sdgx.go:1*/ append(data, "\x1b "...,
					)
					i = 1
				case 11:
					i = 0
					data =  /*line I7W7Sdgx.go:1*/ append(data, "\xdc\x1e\x15\x16"...,
					)
				}
			}
			return  /*line I7W7Sdgx.go:1*/ string(data)
		}()
		return  /*line zk5wLSw0.go:1*/ shim.Error(w5VhCLka)
	}

	var iECyqXpw []string
	if olYMzSWK ==  /*line k6nuiOQA.go:1*/ func() string {
		var data []byte
		i := 4
		decryptKey := 104
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				data =  /*line k6nuiOQA.go:1*/ append(data, 7)
				i = 3
			case 5:
				data =  /*line k6nuiOQA.go:1*/ append(data, 4)
				i = 0
			case 0:
				i = 1
				for y := range data {
					data[y] = data[y] ^  /*line k6nuiOQA.go:1*/ byte(decryptKey^y)
				}
			case 3:
				i = 5
				data =  /*line k6nuiOQA.go:1*/ append(data, 0)
			case 4:
				data =  /*line k6nuiOQA.go:1*/ append(data, 17)
				i = 2
			}
		}
		return  /*line k6nuiOQA.go:1*/ string(data)
	}() {
		iECyqXpw = []string{ /*line KKG7BGwR.go:1*/ func() string {
			var data []byte
			i := 4
			decryptKey := 100
			for counter := 0; i != 3; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					data =  /*line KKG7BGwR.go:1*/ append(data, 254)
					i = 0
				case 5:
					i = 3
					for y := range data {
						data[y] = data[y] +  /*line KKG7BGwR.go:1*/ byte(decryptKey^y)
					}
				case 2:
					i = 5
					data =  /*line KKG7BGwR.go:1*/ append(data, 237)
				case 0:
					i = 1
					data =  /*line KKG7BGwR.go:1*/ append(data, 240)
				case 1:
					i = 2
					data =  /*line KKG7BGwR.go:1*/ append(data, 235)
				}
			}
			return  /*line KKG7BGwR.go:1*/ string(data)
		}()}
	} else if olYMzSWK ==  /*line KVGjkps4.go:1*/ func() string {
		seed :=  /*line KVGjkps4.go:1*/ byte(156)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line KVGjkps4.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line KVGjkps4.go:1*/  /*line KVGjkps4.go:1*/  /*line KVGjkps4.go:1*/  /*line KVGjkps4.go:1*/  /*line KVGjkps4.go:1*/  /*line KVGjkps4.go:1*/ fnc(9)(20)(29)(63)(123)(9)
		return  /*line KVGjkps4.go:1*/ string(data)
	}() {
		iECyqXpw = []string{ /*line rgU5X6MW.go:1*/ func() string {
			data :=  /*line rgU5X6MW.go:1*/ []byte("rUad")
			positions := [...]byte{1, 1, 1, 1}
			for i := 0; i < 4; i += 2 {
				localKey :=  /*line rgU5X6MW.go:1*/ byte(i) +  /*line rgU5X6MW.go:1*/ byte(positions[i]^positions[i+1]) + 135
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line rgU5X6MW.go:1*/ string(data)
		}(),  /*line Gdz7PWb5.go:1*/ func() string {
			data :=  /*line Gdz7PWb5.go:1*/ []byte("m\x03\xb9^\xc7y")
			positions := [...]byte{3, 1, 3, 2, 1, 1, 1, 4}
			for i := 0; i < 8; i += 2 {
				localKey :=  /*line Gdz7PWb5.go:1*/ byte(i) +  /*line Gdz7PWb5.go:1*/ byte(positions[i]^positions[i+1]) + 77
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line Gdz7PWb5.go:1*/ string(data)
		}()}
	} else if olYMzSWK ==  /*line IzxeboHZ.go:1*/ func() string {
		seed :=  /*line IzxeboHZ.go:1*/ byte(153)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line IzxeboHZ.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line IzxeboHZ.go:1*/  /*line IzxeboHZ.go:1*/  /*line IzxeboHZ.go:1*/  /*line IzxeboHZ.go:1*/  /*line IzxeboHZ.go:1*/  /*line IzxeboHZ.go:1*/  /*line IzxeboHZ.go:1*/ fnc(215)(245)(13)(1)(246)(10)(1)
		return  /*line IzxeboHZ.go:1*/ string(data)
	}() {
		iECyqXpw = []string{ /*line TRG7KhXq.go:1*/ func() string {
			seed :=  /*line TRG7KhXq.go:1*/ byte(72)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line TRG7KhXq.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line TRG7KhXq.go:1*/  /*line TRG7KhXq.go:1*/  /*line TRG7KhXq.go:1*/  /*line TRG7KhXq.go:1*/ fnc(58)(231)(8)(21)
			return  /*line TRG7KhXq.go:1*/ string(data)
		}(),  /*line Azwefydv.go:1*/ func() string {
			var data []byte
			i := 1
			decryptKey := 180
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					data =  /*line Azwefydv.go:1*/ append(data, 44)
					i = 4
				case 4:
					i = 2
					for y := range data {
						data[y] = data[y] -  /*line Azwefydv.go:1*/ byte(decryptKey^y)
					}
				case 5:
					i = 0
					data =  /*line Azwefydv.go:1*/ append(data, 24)
				case 7:
					i = 5
					data =  /*line Azwefydv.go:1*/ append(data, 38)
				case 1:
					data =  /*line Azwefydv.go:1*/ append(data, 35)
					i = 7
				case 6:
					data =  /*line Azwefydv.go:1*/ append(data, 24)
					i = 3
				case 0:
					data =  /*line Azwefydv.go:1*/ append(data, 30)
					i = 6
				}
			}
			return  /*line Azwefydv.go:1*/ string(data)
		}(),  /*line axRI171W.go:1*/ func() string {
			var data []byte
			i := 0
			decryptKey := 121
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 7:
					i = 1
					data =  /*line axRI171W.go:1*/ append(data, 19)
				case 6:
					data =  /*line axRI171W.go:1*/ append(data, 1)
					i = 5
				case 5:
					data =  /*line axRI171W.go:1*/ append(data, 11)
					i = 3
				case 1:
					data =  /*line axRI171W.go:1*/ append(data, 17)
					i = 4
				case 4:
					i = 2
					for y := range data {
						data[y] = data[y] +  /*line axRI171W.go:1*/ byte(decryptKey^y)
					}
				case 8:
					data =  /*line axRI171W.go:1*/ append(data, 8)
					i = 7
				case 0:
					i = 6
					data =  /*line axRI171W.go:1*/ append(data, 11)
				case 3:
					data =  /*line axRI171W.go:1*/ append(data, 13)
					i = 8
				}
			}
			return  /*line axRI171W.go:1*/ string(data)
		}()}
	} else {
		 /*line wzJB6ziv.go:1*/ fmt.Println( /*line PzqsL7KR.go:1*/ func() string {
			seed :=  /*line PzqsL7KR.go:1*/ byte(87)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line PzqsL7KR.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/  /*line PzqsL7KR.go:1*/ fnc(165)(107)(219)(98)(24)(36)(69)(69)(223)(188)(122)(224)(203)(74)(228)(189)(135)(9)(14)(38)(76)(142)(34)(67)
			return  /*line PzqsL7KR.go:1*/ string(data)
		}())
	}

	tR32qygH, cSW1wSO3 :=  /*line CmodPwQK.go:1*/ n4c7mRtG.GetState(syHhHGpa)	                    
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line j2LfqhUn.go:1*/ func() string {
			var data []byte
			i := 0
			decryptKey := 16
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 21:
					data =  /*line j2LfqhUn.go:1*/ append(data, "\x9e\xcf\xd8\xc1"...,
					)
					i = 16
				case 18:
					i = 4
					data =  /*line j2LfqhUn.go:1*/ append(data, 242)
				case 12:
					data =  /*line j2LfqhUn.go:1*/ append(data, "\x97\x99\x9e"...,
					)
					i = 22
				case 0:
					i = 3
					data =  /*line j2LfqhUn.go:1*/ append(data, "\x91\xc9"...,
					)
				case 23:
					i = 7
					data =  /*line j2LfqhUn.go:1*/ append(data, "\xb1\xa9"...,
					)
				case 19:
					data =  /*line j2LfqhUn.go:1*/ append(data, 130)
					i = 15
				case 5:
					data =  /*line j2LfqhUn.go:1*/ append(data, "\xba\xa8\xe7"...,
					)
					i = 18
				case 20:
					i = 6
					data =  /*line j2LfqhUn.go:1*/ append(data, "\xa4\xac"...,
					)
				case 15:
					i = 23
					data =  /*line j2LfqhUn.go:1*/ append(data, "\x9f\x9b\x87\xa3"...,
					)
				case 7:
					data =  /*line j2LfqhUn.go:1*/ append(data, "\xbd\xa7\xa0\xa2"...,
					)
					i = 2
				case 17:
					i = 12
					data =  /*line j2LfqhUn.go:1*/ append(data, "\x8f\x97ٝ"...,
					)
				case 14:
					i = 11
					data =  /*line j2LfqhUn.go:1*/ append(data, 161)
				case 3:
					i = 21
					data =  /*line j2LfqhUn.go:1*/ append(data, "\xad\x9b\x9c\x80"...,
					)
				case 9:
					i = 13
					data =  /*line j2LfqhUn.go:1*/ append(data, "\x8f\x8b\x81"...,
					)
				case 10:
					i = 9
					data =  /*line j2LfqhUn.go:1*/ append(data, 128)
				case 1:
					i = 5
					data =  /*line j2LfqhUn.go:1*/ append(data, "\xac\xb8\xad"...,
					)
				case 16:
					i = 10
					data =  /*line j2LfqhUn.go:1*/ append(data, 166)
				case 22:
					i = 19
					data =  /*line j2LfqhUn.go:1*/ append(data, "\x99ӑ\x84"...,
					)
				case 13:
					data =  /*line j2LfqhUn.go:1*/ append(data, "\x81\xda"...,
					)
					i = 17
				case 4:
					for y := range data {
						data[y] = data[y] ^  /*line j2LfqhUn.go:1*/ byte(decryptKey^y)
					}
					i = 8
				case 6:
					data =  /*line j2LfqhUn.go:1*/ append(data, "\xb2Ჯ"...,
					)
					i = 14
				case 2:
					i = 20
					data =  /*line j2LfqhUn.go:1*/ append(data, 237)
				case 11:
					i = 1
					data =  /*line j2LfqhUn.go:1*/ append(data, "徺"...,
					)
				}
			}
			return  /*line j2LfqhUn.go:1*/ string(data)
		}() + syHhHGpa +  /*line e30e0vtz.go:1*/ func() string {
			key :=  /*line e30e0vtz.go:1*/ []byte("\xfb\xc8")
			data :=  /*line e30e0vtz.go:1*/ []byte("\x1dE")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line e30e0vtz.go:1*/ string(data)
		}()
		return  /*line oJKDPfCc.go:1*/ shim.Error(w5VhCLka)
	} else if tR32qygH == nil && (olYMzSWK == "" || olYMzSWK ==  /*line IQ8o71QS.go:1*/ func() string {
		fullData :=  /*line IQ8o71QS.go:1*/ []byte("\xb4\x94")
		var data []byte
		data =  /*line IQ8o71QS.go:1*/ append(data, fullData[0]-fullData[1])
		return  /*line IQ8o71QS.go:1*/ string(data)
	}()) {
		w5VhCLka =  /*line _zzmuPdg.go:1*/ func() string {
			var data []byte
			i := 23
			decryptKey := 222
			for counter := 0; i != 14; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					data =  /*line _zzmuPdg.go:1*/ append(data, 7)
					i = 1
				case 21:
					i = 11
					data =  /*line _zzmuPdg.go:1*/ append(data, "/6)"...,
					)
				case 27:
					data =  /*line _zzmuPdg.go:1*/ append(data, 182)
					i = 17
				case 25:
					i = 20
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xea\xec"...,
					)
				case 26:
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xff\xff"...,
					)
					i = 29
				case 23:
					data =  /*line _zzmuPdg.go:1*/ append(data, "\x06\xae\xce"...,
					)
					i = 24
				case 6:
					i = 5
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xd7(*"...,
					)
				case 19:
					i = 21
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xea31A"...,
					)
				case 20:
					i = 0
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xa8\xf5\xeb\r"...,
					)
				case 15:
					data =  /*line _zzmuPdg.go:1*/ append(data, 44)
					i = 19
				case 7:
					i = 13
					data =  /*line _zzmuPdg.go:1*/ append(data, "\x1d\x15\x11\xdc"...,
					)
				case 2:
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xea\x02\v\x13"...,
					)
					i = 22
				case 32:
					for y := range data {
						data[y] = data[y] +  /*line _zzmuPdg.go:1*/ byte(decryptKey^y)
					}
					i = 14
				case 29:
					i = 25
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xb0\xbd\xa6\xd5"...,
					)
				case 22:
					i = 16
					data =  /*line _zzmuPdg.go:1*/ append(data, "\r\t"...,
					)
				case 10:
					data =  /*line _zzmuPdg.go:1*/ append(data, "\x12\t"...,
					)
					i = 30
				case 5:
					i = 15
					data =  /*line _zzmuPdg.go:1*/ append(data, "\x1875"...,
					)
				case 12:
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xfc\xe7"...,
					)
					i = 32
				case 0:
					i = 28
					data =  /*line _zzmuPdg.go:1*/ append(data, "\t\x02"...,
					)
				case 30:
					data =  /*line _zzmuPdg.go:1*/ append(data, "\f\f\xb3"...,
					)
					i = 3
				case 31:
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xde\x1c\""...,
					)
					i = 9
				case 18:
					data =  /*line _zzmuPdg.go:1*/ append(data, 38)
					i = 6
				case 1:
					i = 27
					data =  /*line _zzmuPdg.go:1*/ append(data, 9)
				case 11:
					i = 12
					data =  /*line _zzmuPdg.go:1*/ append(data, 53)
				case 16:
					i = 8
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xc6/+\xd9"...,
					)
				case 13:
					data =  /*line _zzmuPdg.go:1*/ append(data, 195)
					i = 2
				case 9:
					i = 18
					data =  /*line _zzmuPdg.go:1*/ append(data, 26)
				case 3:
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xfd\x04\xb2\x05"...,
					)
					i = 4
				case 8:
					data =  /*line _zzmuPdg.go:1*/ append(data, " (.!"...,
					)
					i = 31
				case 24:
					data =  /*line _zzmuPdg.go:1*/ append(data, "\xfc\x01"...,
					)
					i = 26
				case 28:
					data =  /*line _zzmuPdg.go:1*/ append(data, 13)
					i = 10
				case 17:
					data =  /*line _zzmuPdg.go:1*/ append(data, "\x0f\x11\x0f\x13"...,
					)
					i = 7
				}
			}
			return  /*line _zzmuPdg.go:1*/ string(data)
		}() + eviEA3sp +  /*line xJfzs0tp.go:1*/ func() string {
			data :=  /*line xJfzs0tp.go:1*/ []byte("3l")
			positions := [...]byte{1, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line xJfzs0tp.go:1*/ byte(i) +  /*line xJfzs0tp.go:1*/ byte(positions[i]^positions[i+1]) + 77
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line xJfzs0tp.go:1*/ string(data)
		}()
		return  /*line SXaMLna2.go:1*/ shim.Error(w5VhCLka)
	} else if tR32qygH != nil {
		hnzizKM2 := AuthorizedOrgs{}
		cSW1wSO3 =  /*line kRicpr8p.go:1*/ json.Unmarshal(tR32qygH, &hnzizKM2)	                               
		if cSW1wSO3 != nil {
			return  /*line HfNGl5Gi.go:1*/ shim.Error( /*line VikXo3jG.go:1*/ cSW1wSO3.Error())
		}

		if olYMzSWK == "" || olYMzSWK ==  /*line kVmKzzNX.go:1*/ func() string {
			key :=  /*line kVmKzzNX.go:1*/ []byte("?")
			data :=  /*line kVmKzzNX.go:1*/ []byte("\xe1")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line kVmKzzNX.go:1*/ string(data)
		}() {
			iECyqXpw = hnzizKM2.Permission
		}

		 /*line DUFM2tmU.go:1*/ fmt.Println( /*line yrno4oXq.go:1*/ func() string {
			fullData :=  /*line yrno4oXq.go:1*/ []byte("ѹ\x84F\xfe\xf8S\x93c\xffT\x01\xa8h7\x1c\x96\f\vq\xc3o\tV&\x9b<M\xcb\x11c3V\xfevN>\x8e\x0e£\x0f\xd1讅4x\xeb\x96\x1a\x83Oj")
			var data []byte
			data =  /*line yrno4oXq.go:1*/ append(data, fullData[34]+fullData[28], fullData[46]+fullData[26], fullData[50]+fullData[23], fullData[27]-fullData[43], fullData[11]-fullData[7], fullData[17]^fullData[13], fullData[12]+fullData[47], fullData[1]-fullData[32], fullData[3]^fullData[31], fullData[5]-fullData[45], fullData[24]+fullData[35], fullData[48]^fullData[2], fullData[36]-fullData[0], fullData[44]-fullData[37], fullData[6]+fullData[38], fullData[53]^fullData[22], fullData[33]-fullData[25], fullData[10]+fullData[29], fullData[19]-fullData[4], fullData[39]-fullData[52], fullData[51]^fullData[40], fullData[41]+fullData[8], fullData[9]^fullData[16], fullData[42]+fullData[49], fullData[18]^fullData[30], fullData[14]-fullData[20], fullData[15]^fullData[21])
			return  /*line yrno4oXq.go:1*/ string(data)
		}())
		if  /*line GhHnwvwA.go:1*/ len(ezgSKu0t) > 0 {
			ezgSKu0t =  /*line ly7uz51n.go:1*/ _urwHavD(ezgSKu0t, hnzizKM2.CustomAccessRights)
		}

	}

	l5TiVBv5 := &AuthorizedOrgs{mzjt9rLK, eviEA3sp, hcl4L2y6, e5loS483, deTJzFce, iECyqXpw, ezgSKu0t, dJwgW2jL,
		hOgpzUCn, ipXT6jpu, wSxUCaJ0, v5vcwK31}
	niQzOH6M, cSW1wSO3 :=  /*line Y73bZGzB.go:1*/ json.Marshal(l5TiVBv5)
	if cSW1wSO3 != nil {
		return  /*line n14XzI6T.go:1*/ shim.Error( /*line KjJMxTPg.go:1*/ cSW1wSO3.Error())
	}

	                        
	cSW1wSO3 =  /*line N2S765aX.go:1*/ n4c7mRtG.PutState(syHhHGpa, niQzOH6M)
	if cSW1wSO3 != nil {
		return  /*line Csr40lBJ.go:1*/ shim.Error( /*line gixtSZ6H.go:1*/ cSW1wSO3.Error())
	}

	cSW1wSO3 =  /*line e_zh7vTa.go:1*/ n4c7mRtG.DelState(hSHnto8d)	                                                                
	if cSW1wSO3 != nil {
		return  /*line QzU_6n3V.go:1*/ shim.Error( /*line kimXLndK.go:1*/ func() string {
			var data []byte
			i := 2
			decryptKey := 249
			for counter := 0; i != 7; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 10:
					i = 1
					data =  /*line kimXLndK.go:1*/ append(data, 52)
				case 2:
					data =  /*line kimXLndK.go:1*/ append(data, "\x1c8="...,
					)
					i = 3
				case 17:
					data =  /*line kimXLndK.go:1*/ append(data, 64)
					i = 4
				case 3:
					data =  /*line kimXLndK.go:1*/ append(data, "A77\xf0"...,
					)
					i = 16
				case 4:
					data =  /*line kimXLndK.go:1*/ append(data, 66)
					i = 12
				case 8:
					i = 5
					data =  /*line kimXLndK.go:1*/ append(data, 113)
				case 0:
					i = 10
					data =  /*line kimXLndK.go:1*/ append(data, 44)
				case 21:
					i = 11
					data =  /*line kimXLndK.go:1*/ append(data, "@L>\xe6"...,
					)
				case 11:
					data =  /*line kimXLndK.go:1*/ append(data, "66,"...,
					)
					i = 6
				case 18:
					i = 13
					data =  /*line kimXLndK.go:1*/ append(data, "/C"...,
					)
				case 13:
					i = 0
					data =  /*line kimXLndK.go:1*/ append(data, "5<8"...,
					)
				case 15:
					data =  /*line kimXLndK.go:1*/ append(data, "ccp"...,
					)
					i = 8
				case 12:
					data =  /*line kimXLndK.go:1*/ append(data, 70)
					i = 21
				case 6:
					i = 14
					data =  /*line kimXLndK.go:1*/ append(data, "#1)"...,
					)
				case 5:
					i = 9
					data =  /*line kimXLndK.go:1*/ append(data, 98)
				case 20:
					i = 15
					data =  /*line kimXLndK.go:1*/ append(data, "Zef\x10"...,
					)
				case 9:
					data =  /*line kimXLndK.go:1*/ append(data, "mo2"...,
					)
					i = 19
				case 16:
					data =  /*line kimXLndK.go:1*/ append(data, "EM\xff"...,
					)
					i = 17
				case 14:
					i = 18
					data =  /*line kimXLndK.go:1*/ append(data, 59)
				case 19:
					i = 7
					for y := range data {
						data[y] = data[y] -  /*line kimXLndK.go:1*/ byte(decryptKey^y)
					}
				case 1:
					i = 20
					data =  /*line kimXLndK.go:1*/ append(data, "\xe9WZW"...,
					)
				}
			}
			return  /*line kimXLndK.go:1*/ string(data)
		}() +  /*line NnSHmEK9.go:1*/ cSW1wSO3.Error())
	}

	               
	xwK9eDsS, cSW1wSO3 :=  /*line hzMlWznz.go:1*/ o8Z_JkaH(n4c7mRtG, []string{eviEA3sp, w8EOizcx, dJwgW2jL, hSHnto8d,  /*line Of3lRK00.go:1*/ func() string {
		key :=  /*line Of3lRK00.go:1*/ []byte("\xf1\xe5\x06\x96\xc5.")
		data :=  /*line Of3lRK00.go:1*/ []byte("2Hi\xfb5\xa2")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line Of3lRK00.go:1*/ string(data)
	}(), pbLLcfzS, hOgpzUCn, ipXT6jpu, wSxUCaJ0, v5vcwK31, olYMzSWK,  /*line IUmPzra_.go:1*/ strings.Join(ezgSKu0t,  /*line _au1fXE2.go:1*/ func() string {
		seed :=  /*line _au1fXE2.go:1*/ byte(239)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line _au1fXE2.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line _au1fXE2.go:1*/ fnc(27)
		return  /*line _au1fXE2.go:1*/ string(data)
	}()), e5loS483, deTJzFce})
	if cSW1wSO3 != nil {
		return  /*line U3To1huE.go:1*/ shim.Error( /*line Uc0v_Zpj.go:1*/ cSW1wSO3.Error())
	}
	 /*line RDJRgugL.go:1*/ fmt.Println(xwK9eDsS)

	wt_3C8vB :=  /*line czv_zh_I.go:1*/ func() string {
		var data []byte
		i := 8
		decryptKey := 114
		for counter := 0; i != 9; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 10:
				i = 1
				data =  /*line czv_zh_I.go:1*/ append(data, "\xa6\xb0"...,
				)
			case 5:
				i = 0
				data =  /*line czv_zh_I.go:1*/ append(data, "\xa2¸\xbb"...,
				)
			case 1:
				data =  /*line czv_zh_I.go:1*/ append(data, "\xac\xba"...,
				)
				i = 5
			case 6:
				data =  /*line czv_zh_I.go:1*/ append(data, "ɱ\xc1\xb1"...,
				)
				i = 2
			case 0:
				i = 7
				data =  /*line czv_zh_I.go:1*/ append(data, "\xbbȯ\xa9"...,
				)
			case 3:
				data =  /*line czv_zh_I.go:1*/ append(data, "\xb9\xab"...,
				)
				i = 10
			case 4:
				for y := range data {
					data[y] = data[y] -  /*line czv_zh_I.go:1*/ byte(decryptKey^y)
				}
				i = 9
			case 2:
				i = 4
				data =  /*line czv_zh_I.go:1*/ append(data, "\xbe\xc3"...,
				)
			case 8:
				i = 3
				data =  /*line czv_zh_I.go:1*/ append(data, 181)
			case 7:
				i = 6
				data =  /*line czv_zh_I.go:1*/ append(data, "\xbd\xb7ʹ"...,
				)
			}
		}
		return  /*line czv_zh_I.go:1*/ string(data)
	}()
	pc2yhaw_, cSW1wSO3 :=  /*line Oa7lZ8gt.go:1*/ n4c7mRtG.CreateCompositeKey(wt_3C8vB, []string{dJwgW2jL, eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line I5r6qMxu.go:1*/ shim.Error( /*line tKv80csA.go:1*/ cSW1wSO3.Error())
	}
	lFkgHiji :=  /*line u1bzPSkX.go:1*/ func() []byte {
		var data []byte
		i := 0
		decryptKey := 208
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				i = 2
				data =  /*line u1bzPSkX.go:1*/ append(data, 210)
			case 2:
				for y := range data {
					data[y] = data[y] -  /*line u1bzPSkX.go:1*/ byte(decryptKey^y)
				}
				i = 1
			}
		}
		return data
	}()
	cSW1wSO3 =  /*line QzpRxit6.go:1*/ n4c7mRtG.PutState(pc2yhaw_, lFkgHiji)
	if cSW1wSO3 != nil {
		return  /*line gQGL4KDx.go:1*/ shim.Error( /*line ujwzireT.go:1*/ cSW1wSO3.Error())
	}

	xHTOVzVs :=  /*line tAHthHhQ.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 101
		for counter := 0; i != 4; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 5:
				for y := range data {
					data[y] = data[y] ^  /*line tAHthHhQ.go:1*/ byte(decryptKey^y)
				}
				i = 4
			case 0:
				data =  /*line tAHthHhQ.go:1*/ append(data, 37)
				i = 6
			case 7:
				data =  /*line tAHthHhQ.go:1*/ append(data, "3%"...,
				)
				i = 3
			case 8:
				data =  /*line tAHthHhQ.go:1*/ append(data, "< %"...,
				)
				i = 0
			case 6:
				i = 5
				data =  /*line tAHthHhQ.go:1*/ append(data, 108)
			case 3:
				i = 1
				data =  /*line tAHthHhQ.go:1*/ append(data, "\"*,<"...,
				)
			case 2:
				data =  /*line tAHthHhQ.go:1*/ append(data, 15)
				i = 7
			case 1:
				data =  /*line tAHthHhQ.go:1*/ append(data, 38)
				i = 8
			}
		}
		return  /*line tAHthHhQ.go:1*/ string(data)
	}() + dJwgW2jL +  /*line Ix_mziU5.go:1*/ func() string {
		data :=  /*line Ix_mziU5.go:1*/ []byte("H\x06\xe4\xf5 4\x11\x14nN4u\xef0o*\xf3>\xf5\xd1%")
		positions := [...]byte{2, 17, 1, 9, 17, 20, 17, 15, 18, 12, 19, 5, 17, 17, 6, 7, 5, 16, 12, 16, 7, 10, 18, 5, 0, 3, 6, 13}
		for i := 0; i < 28; i += 2 {
			localKey :=  /*line Ix_mziU5.go:1*/ byte(i) +  /*line Ix_mziU5.go:1*/ byte(positions[i]^positions[i+1]) + 16
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line Ix_mziU5.go:1*/ string(data)
	}() +  /*line kvxTrbtH.go:1*/ func() string {
		fullData :=  /*line kvxTrbtH.go:1*/ []byte("e\xf9\xef\\\x992\xe8\xc8\x00\x85f\v\x9b\xfc\x98\xa8ia\xc5=\xba&")
		var data []byte
		data =  /*line kvxTrbtH.go:1*/ append(data, fullData[9]+fullData[2], fullData[19]+fullData[5], fullData[7]-fullData[15], fullData[18]-fullData[17], fullData[13]-fullData[12], fullData[16]+fullData[11], fullData[1]-fullData[14], fullData[4]-fullData[21], fullData[0]-fullData[8], fullData[3]-fullData[6], fullData[10]+fullData[20])
		return  /*line kvxTrbtH.go:1*/ string(data)
	}() + eviEA3sp
	q8ZF8zaz :=  /*line gxnbMXzi.go:1*/ []byte(xHTOVzVs)

	return  /*line ZoE4BrUP.go:1*/ shim.Success(q8ZF8zaz)
}

func (g4rnrSNn *G1Y_7pz_) ePZzyUt5(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	var cSW1wSO3 error
	if  /*line MSYaTaOY.go:1*/ len(ktsi1_SQ) != 5 {
		return  /*line qZ0Xm5Rr.go:1*/ shim.Error( /*line zxlA5jaJ.go:1*/ func() string {
			var data []byte
			i := 5
			decryptKey := 49
			for counter := 0; i != 20; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 12:
					for y := range data {
						data[y] = data[y] ^  /*line zxlA5jaJ.go:1*/ byte(decryptKey^y)
					}
					i = 20
				case 18:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xc6\xc1ٌ"...,
					)
					i = 15
				case 22:
					data =  /*line zxlA5jaJ.go:1*/ append(data, 251)
					i = 10
				case 9:
					i = 27
					data =  /*line zxlA5jaJ.go:1*/ append(data, 194)
				case 13:
					i = 30
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xd0\xd9"...,
					)
				case 2:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xe1\xf0\xd5"...,
					)
					i = 1
				case 28:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "ʘ\xdd\xd5"...,
					)
					i = 7
				case 7:
					i = 3
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xf7\xa4\xe1"...,
					)
				case 16:
					i = 21
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xdb\xc4\xca\xce"...,
					)
				case 26:
					i = 18
					data =  /*line zxlA5jaJ.go:1*/ append(data, 210)
				case 10:
					i = 29
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xb1\xbe"...,
					)
				case 0:
					i = 26
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xc4\xc9\xd3"...,
					)
				case 21:
					i = 23
					data =  /*line zxlA5jaJ.go:1*/ append(data, "ؕ"...,
					)
				case 24:
					i = 25
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xeb\xe2"...,
					)
				case 27:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xd4\xc7"...,
					)
					i = 13
				case 6:
					data =  /*line zxlA5jaJ.go:1*/ append(data, 175)
					i = 17
				case 8:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\x86\xc0\xd6"...,
					)
					i = 12
				case 1:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xf4\xf5\xf4"...,
					)
					i = 19
				case 17:
					i = 2
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xfc\xec\xfe\xe4"...,
					)
				case 3:
					i = 24
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xf3\xef\xe3\xf7"...,
					)
				case 29:
					data =  /*line zxlA5jaJ.go:1*/ append(data, 220)
					i = 11
				case 4:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\x96\xd0"...,
					)
					i = 9
				case 14:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\x86\x90\x8e\x88"...,
					)
					i = 8
				case 19:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xe3\xe0\xdd\xef"...,
					)
					i = 22
				case 11:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xe0\xeb\xff"...,
					)
					i = 14
				case 15:
					data =  /*line zxlA5jaJ.go:1*/ append(data, 193)
					i = 16
				case 5:
					i = 0
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xec\xca"...,
					)
				case 25:
					data =  /*line zxlA5jaJ.go:1*/ append(data, 226)
					i = 6
				case 23:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xdb\xd1"...,
					)
					i = 4
				case 30:
					data =  /*line zxlA5jaJ.go:1*/ append(data, "\xd1\xca"...,
					)
					i = 28
				}
			}
			return  /*line zxlA5jaJ.go:1*/ string(data)
		}())
	}

	eviEA3sp := ktsi1_SQ[0]
	olYMzSWK := ktsi1_SQ[1]
	dJwgW2jL := ktsi1_SQ[2]
	pbLLcfzS := ktsi1_SQ[3]
	ezgSKu0t :=  /*line _p3e8puX.go:1*/ strings.Split(ktsi1_SQ[4],  /*line kdOXjazQ.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 17
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				for y := range data {
					data[y] = data[y] ^  /*line kdOXjazQ.go:1*/ byte(decryptKey^y)
				}
				i = 2
			case 1:
				data =  /*line kdOXjazQ.go:1*/ append(data, 61)
				i = 0
			}
		}
		return  /*line kdOXjazQ.go:1*/ string(data)
	}())

	w8EOizcx, cSW1wSO3 :=  /*line IM2sBJ3Y.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line Eyczfj2i.go:1*/ shim.Error( /*line jzi4PnQd.go:1*/ cSW1wSO3.Error())
	}

	b4dMB1zL, cSW1wSO3 :=  /*line q_NM3ZOy.go:1*/ jmSY84co(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line BCDhz52N.go:1*/ shim.Error( /*line JYHK6aoB.go:1*/ cSW1wSO3.Error())
	}

	zzhdvyqZ := &DatasetMetadataResponse{}
	cSW1wSO3 =  /*line uupEUvus.go:1*/ json.Unmarshal( /*line Wc3KkaZk.go:1*/ []byte(b4dMB1zL), zzhdvyqZ)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line hNU4ODx6.go:1*/ func() string {
			fullData :=  /*line hNU4ODx6.go:1*/ []byte("C\x05\x8e\nH\xccPee:KY'\xdeT\xe1\x91kYk\x86a\x8d\x15\xa6l\x18+\x95\xe0\x954\xaci\xf0\x1bw\xde\xe1\x87Z_\xe3\xcb\x06\xbc\xaf\xff\x90\xe9?\x05_\xc2ޥ\x19\xed9\xbf\x80\xe4")
			var data []byte
			data =  /*line hNU4ODx6.go:1*/ append(data, fullData[59]+fullData[45], fullData[39]-fullData[8], fullData[55]^fullData[29], fullData[36]-fullData[1], fullData[22]-fullData[35], fullData[37]+fullData[16], fullData[18]^fullData[27], fullData[9]^fullData[26], fullData[24]-fullData[25], fullData[49]^fullData[43], fullData[15]+fullData[7], fullData[17]^fullData[3], fullData[53]-fullData[11], fullData[46]-fullData[0], fullData[52]+fullData[44], fullData[33]-fullData[51], fullData[5]+fullData[14], fullData[21]-fullData[57], fullData[38]+fullData[2], fullData[58]-fullData[56], fullData[6]^fullData[31], fullData[34]^fullData[28], fullData[42]-fullData[60], fullData[12]^fullData[4], fullData[54]+fullData[20], fullData[50]^fullData[40], fullData[19]^fullData[10], fullData[41]^fullData[23], fullData[47]-fullData[32], fullData[61]-fullData[30], fullData[13]-fullData[48])
			return  /*line hNU4ODx6.go:1*/ string(data)
		}() + eviEA3sp +  /*line ewbt77Bg.go:1*/ func() string {
			seed :=  /*line ewbt77Bg.go:1*/ byte(255)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line ewbt77Bg.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line ewbt77Bg.go:1*/  /*line ewbt77Bg.go:1*/ fnc(35)(91)
			return  /*line ewbt77Bg.go:1*/ string(data)
		}()
		return  /*line RG_96_s9.go:1*/ shim.Error(w5VhCLka)
	}

	                                                                  
	if  /*line piQmERzZ.go:1*/ nSRaDwQg(ezgSKu0t, zzhdvyqZ.CustomAccessRights) {
		 /*line Z1LGJQBP.go:1*/ fmt.Println( /*line vBYXom0t.go:1*/ func() string {
			seed :=  /*line vBYXom0t.go:1*/ byte(9)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line vBYXom0t.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/  /*line vBYXom0t.go:1*/ fnc(72)(33)(2)(6)(21)(255)(252)(227)(12)(13)(227)(73)(209)(246)(10)(247)(21)(226)(81)(163)(6)(8)(22)(250)(240)(83)(180)(19)(234)(31)(226)(11)
			return  /*line vBYXom0t.go:1*/ string(data)
		}())
	} else {
		w5VhCLka =  /*line AlExcBHS.go:1*/ func() string {
			fullData :=  /*line AlExcBHS.go:1*/ []byte("\x15\xf7\xbf\x87\xf7\xb5b7\u007f\xb2ک+R\x14\x06{\xd5K\xb1\x99?\xed\x86G2<YY\x04\x1c\v$\x87\x9b\xe1Ѹ\xa9P'c\x9c\xebx\xf1\t)\xc7\xcdX\rd\xe0\xd0O\xd7\xdbȤ\xde\xc5cy\xb2\xd9\x1e\x9c\xb7\xa2\x03\n(Kq\xcd")
			var data []byte
			data =  /*line AlExcBHS.go:1*/ append(data, fullData[25]-fullData[68], fullData[24]+fullData[57], fullData[47]+fullData[30], fullData[44]-fullData[15], fullData[40]+fullData[18], fullData[74]^fullData[66], fullData[41]-fullData[45], fullData[17]^fullData[1], fullData[28]^fullData[62], fullData[34]+fullData[3], fullData[52]-fullData[51], fullData[65]+fullData[20], fullData[10]^fullData[5], fullData[11]^fullData[48], fullData[33]^fullData[53], fullData[22]^fullData[49], fullData[19]+fullData[9], fullData[8]^fullData[71], fullData[73]+fullData[72], fullData[13]-fullData[60], fullData[58]-fullData[27], fullData[14]^fullData[63], fullData[35]+fullData[21], fullData[6]^fullData[70], fullData[0]-fullData[64], fullData[31]+fullData[50], fullData[43]-fullData[23], fullData[26]^fullData[55], fullData[56]+fullData[67], fullData[29]^fullData[32], fullData[16]-fullData[46], fullData[75]+fullData[42], fullData[39]^fullData[7], fullData[2]+fullData[38], fullData[59]^fullData[54], fullData[69]+fullData[36], fullData[4]+fullData[12], fullData[61]+fullData[37])
			return  /*line AlExcBHS.go:1*/ string(data)
		}()
		return  /*line GpZqVyR1.go:1*/ shim.Error(w5VhCLka)
	}

	                                            
	qFOnXzGg, cSW1wSO3 :=  /*line Gh_CyPGy.go:1*/ fvIQVpVM(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line xp2Cxn7b.go:1*/ shim.Error( /*line mczjVghC.go:1*/ cSW1wSO3.Error())
	}
	if qFOnXzGg != w8EOizcx {
		return  /*line ybsLGYER.go:1*/ shim.Error( /*line AkmRQ3lW.go:1*/ func() string {
			fullData :=  /*line AkmRQ3lW.go:1*/ []byte("\x81t\xb5}\xe8\xe8\x11\xe7E\x86\xfa\x11|\xe9\bq\xf3\x92pg\xcb\xef\xef\x831[\xff\xe9\xfe\xeb\x91\xe8b\bݠPƾoq8R\xed\xe0\xdb*\xfd\x15z\xe6\xd2\xf0?\x85\x10\xb8:\v+qF")
			var data []byte
			data =  /*line AkmRQ3lW.go:1*/ append(data, fullData[61]^fullData[33], fullData[12]+fullData[16], fullData[21]+fullData[54], fullData[59]^fullData[58], fullData[0]^fullData[44], fullData[32]-fullData[43], fullData[53]-fullData[20], fullData[57]^fullData[42], fullData[47]^fullData[17], fullData[15]-fullData[26], fullData[10]+fullData[39], fullData[30]+fullData[13], fullData[1]^fullData[6], fullData[48]^fullData[40], fullData[24]+fullData[22], fullData[25]-fullData[7], fullData[8]^fullData[46], fullData[41]+fullData[5], fullData[51]^fullData[35], fullData[4]+fullData[3], fullData[50]-fullData[18], fullData[27]^fullData[9], fullData[49]^fullData[11], fullData[28]+fullData[19], fullData[14]-fullData[31], fullData[60]^fullData[55], fullData[56]^fullData[45], fullData[38]^fullData[34], fullData[36]-fullData[29], fullData[2]^fullData[37], fullData[52]^fullData[23])
			return  /*line AkmRQ3lW.go:1*/ string(data)
		}())
	}
	rHDGEtgU := eviEA3sp +  /*line JfdwVsbk.go:1*/ func() string {
		var data []byte
		i := 2
		decryptKey := 88
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 0
				for y := range data {
					data[y] = data[y] -  /*line JfdwVsbk.go:1*/ byte(decryptKey^y)
				}
			case 2:
				data =  /*line JfdwVsbk.go:1*/ append(data, 136)
				i = 1
			}
		}
		return  /*line JfdwVsbk.go:1*/ string(data)
	}() + dJwgW2jL
	tQcQN8RB, cSW1wSO3 :=  /*line sr3pXXTX.go:1*/ n4c7mRtG.GetState(rHDGEtgU)	                    
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line G2aYXfJ3.go:1*/ func() string {
			var data []byte
			i := 22
			decryptKey := 63
			for counter := 0; i != 26; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					i = 5
					data =  /*line G2aYXfJ3.go:1*/ append(data, "\xa2\xa2S"...,
					)
				case 20:
					data =  /*line G2aYXfJ3.go:1*/ append(data, 163)
					i = 9
				case 8:
					data =  /*line G2aYXfJ3.go:1*/ append(data, 145)
					i = 6
				case 4:
					i = 17
					data =  /*line G2aYXfJ3.go:1*/ append(data, "\x8d\xa2\x9e"...,
					)
				case 22:
					data =  /*line G2aYXfJ3.go:1*/ append(data, "\xb1Yy\xa7"...,
					)
					i = 3
				case 19:
					data =  /*line G2aYXfJ3.go:1*/ append(data, "\x9d\x9dF"...,
					)
					i = 23
				case 16:
					i = 10
					data =  /*line G2aYXfJ3.go:1*/ append(data, "\x85\x84"...,
					)
				case 0:
					i = 16
					data =  /*line G2aYXfJ3.go:1*/ append(data, 139)
				case 24:
					i = 21
					data =  /*line G2aYXfJ3.go:1*/ append(data, "xfu"...,
					)
				case 9:
					i = 19
					data =  /*line G2aYXfJ3.go:1*/ append(data, 167)
				case 18:
					data =  /*line G2aYXfJ3.go:1*/ append(data, "\x83}9"...,
					)
					i = 11
				case 12:
					for y := range data {
						data[y] = data[y] -  /*line G2aYXfJ3.go:1*/ byte(decryptKey^y)
					}
					i = 26
				case 11:
					i = 24
					data =  /*line G2aYXfJ3.go:1*/ append(data, "jh"...,
					)
				case 21:
					i = 15
					data =  /*line G2aYXfJ3.go:1*/ append(data, 104)
				case 5:
					data =  /*line G2aYXfJ3.go:1*/ append(data, "xa\x82\x9e"...,
					)
					i = 20
				case 17:
					data =  /*line G2aYXfJ3.go:1*/ append(data, "\x93\x97\x9b\u007f"...,
					)
					i = 8
				case 10:
					i = 1
					data =  /*line G2aYXfJ3.go:1*/ append(data, 153)
				case 3:
					i = 2
					data =  /*line G2aYXfJ3.go:1*/ append(data, 164)
				case 14:
					i = 13
					data =  /*line G2aYXfJ3.go:1*/ append(data, 126)
				case 1:
					i = 4
					data =  /*line G2aYXfJ3.go:1*/ append(data, 79)
				case 6:
					i = 7
					data =  /*line G2aYXfJ3.go:1*/ append(data, "u\x89"...,
					)
				case 23:
					i = 0
					data =  /*line G2aYXfJ3.go:1*/ append(data, "\x9b\x93E\x85"...,
					)
				case 13:
					data =  /*line G2aYXfJ3.go:1*/ append(data, "1\x84\x8e"...,
					)
					i = 25
				case 25:
					i = 18
					data =  /*line G2aYXfJ3.go:1*/ append(data, "\x8e=\x8e"...,
					)
				case 15:
					i = 12
					data =  /*line G2aYXfJ3.go:1*/ append(data, "t;."...,
					)
				case 7:
					i = 14
					data =  /*line G2aYXfJ3.go:1*/ append(data, "{\x82"...,
					)
				}
			}
			return  /*line G2aYXfJ3.go:1*/ string(data)
		}() + rHDGEtgU +  /*line LOTcnXmX.go:1*/ func() string {
			key :=  /*line LOTcnXmX.go:1*/ []byte("\xa0\xeb")
			data :=  /*line LOTcnXmX.go:1*/ []byte("\xc2h")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line LOTcnXmX.go:1*/ string(data)
		}()
		return  /*line o1PpToia.go:1*/ shim.Error(w5VhCLka)
	} else if tQcQN8RB == nil {
		w5VhCLka =  /*line tG5Mz3aK.go:1*/ func() string {
			data :=  /*line tG5Mz3aK.go:1*/ []byte("\x03\xe1κ\r\x11p\" \xf1Dat2\x9c\xbf'\xaf\x16\xd0es to\x8e\xc6ex\x18\xcb\x02:\xee")
			positions := [...]byte{25, 14, 6, 13, 1, 29, 1, 23, 1, 14, 26, 30, 29, 33, 17, 33, 25, 8, 19, 16, 23, 15, 2, 9, 13, 4, 26, 3, 16, 23, 18, 30, 15, 0, 4, 31, 14, 5}
			for i := 0; i < 38; i += 2 {
				localKey :=  /*line tG5Mz3aK.go:1*/ byte(i) +  /*line tG5Mz3aK.go:1*/ byte(positions[i]^positions[i+1]) + 51
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line tG5Mz3aK.go:1*/ string(data)
		}() + rHDGEtgU +  /*line iXzibUql.go:1*/ func() string {
			data :=  /*line iXzibUql.go:1*/ []byte("\xb0U")
			positions := [...]byte{1, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line iXzibUql.go:1*/ byte(i) +  /*line iXzibUql.go:1*/ byte(positions[i]^positions[i+1]) + 50
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line iXzibUql.go:1*/ string(data)
		}()
		return  /*line Sqq6fQ8_.go:1*/ shim.Error(w5VhCLka)
	}
	hnzizKM2 := AuthorizedOrgs{}
	cSW1wSO3 =  /*line enumLc_Q.go:1*/ json.Unmarshal(tQcQN8RB, &hnzizKM2)	                               
	if cSW1wSO3 != nil {
		return  /*line JLt8edDG.go:1*/ shim.Error( /*line Cawhz4gG.go:1*/ cSW1wSO3.Error())
	}

	if (olYMzSWK ==  /*line U4OSCzyz.go:1*/ func() string {
		var data []byte
		i := 3
		decryptKey := 144
		for counter := 0; i != 1; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				i = 1
				for y := range data {
					data[y] = data[y] ^  /*line U4OSCzyz.go:1*/ byte(decryptKey^y)
				}
			case 2:
				data =  /*line U4OSCzyz.go:1*/ append(data, 252)
				i = 0
			case 3:
				i = 5
				data =  /*line U4OSCzyz.go:1*/ append(data, 233)
			case 4:
				data =  /*line U4OSCzyz.go:1*/ append(data, 248)
				i = 2
			case 5:
				i = 4
				data =  /*line U4OSCzyz.go:1*/ append(data, 255)
			}
		}
		return  /*line U4OSCzyz.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line WLbt_bKm.go:1*/ func() string {
		data :=  /*line WLbt_bKm.go:1*/ []byte("\xac\rd\x1c\xff\xaf")
		positions := [...]byte{4, 4, 5, 0, 5, 3, 0, 1}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line WLbt_bKm.go:1*/ byte(i) +  /*line WLbt_bKm.go:1*/ byte(positions[i]^positions[i+1]) + 153
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line WLbt_bKm.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line e2rga1_C.go:1*/ func() string {
		var data []byte
		i := 5
		decryptKey := 36
		for counter := 0; i != 6; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 1:
				i = 7
				data =  /*line e2rga1_C.go:1*/ append(data, 130)
			case 3:
				i = 6
				for y := range data {
					data[y] = data[y] -  /*line e2rga1_C.go:1*/ byte(decryptKey^y)
				}
			case 2:
				data =  /*line e2rga1_C.go:1*/ append(data, 119)
				i = 1
			case 8:
				i = 2
				data =  /*line e2rga1_C.go:1*/ append(data, 124)
			case 7:
				data =  /*line e2rga1_C.go:1*/ append(data, 128)
				i = 3
			case 5:
				i = 4
				data =  /*line e2rga1_C.go:1*/ append(data, 122)
			case 4:
				data =  /*line e2rga1_C.go:1*/ append(data, 112)
				i = 0
			case 0:
				i = 8
				data =  /*line e2rga1_C.go:1*/ append(data, 122)
			}
		}
		return  /*line e2rga1_C.go:1*/ string(data)
	}()) && ((ezgSKu0t[0] != "") || (ezgSKu0t[0] !=  /*line lYHBfbIE.go:1*/ func() string {
		data :=  /*line lYHBfbIE.go:1*/ []byte("v")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line lYHBfbIE.go:1*/ byte(i) +  /*line lYHBfbIE.go:1*/ byte(positions[i]^positions[i+1]) + 86
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line lYHBfbIE.go:1*/ string(data)
	}())) {
		var iECyqXpw []string
		if olYMzSWK ==  /*line kciLoE52.go:1*/ func() string {
			seed :=  /*line kciLoE52.go:1*/ byte(160)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line kciLoE52.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line kciLoE52.go:1*/  /*line kciLoE52.go:1*/  /*line kciLoE52.go:1*/  /*line kciLoE52.go:1*/ fnc(18)(23)(42)(87)
			return  /*line kciLoE52.go:1*/ string(data)
		}() {
			cSW1wSO3 =  /*line EfZoBO53.go:1*/ n4c7mRtG.DelState(rHDGEtgU)	                             
			if cSW1wSO3 != nil {
				return  /*line EnjKQtzu.go:1*/ shim.Error( /*line W5_HrEHk.go:1*/ func() string {
					fullData :=  /*line W5_HrEHk.go:1*/ []byte("\xc6@\xeaB\xa0E\xf3/\xe6\xf3\x1f\xbet9\r\x8f\x86\x9a'C҇\xed\xb2>!p\b3l\xc5x\x13(\xd1\xde\xd2-D\xbd\xa0v7M\x82\xe3")
					var data []byte
					data =  /*line W5_HrEHk.go:1*/ append(data, fullData[12]+fullData[36], fullData[34]-fullData[26], fullData[9]+fullData[41], fullData[38]+fullData[33], fullData[18]+fullData[24], fullData[6]-fullData[15], fullData[8]^fullData[0], fullData[19]^fullData[42], fullData[37]+fullData[3], fullData[20]-fullData[23], fullData[25]-fullData[39], fullData[30]+fullData[40], fullData[16]^fullData[2], fullData[35]+fullData[21], fullData[27]+fullData[29], fullData[31]-fullData[32], fullData[14]-fullData[22], fullData[1]^fullData[28], fullData[7]+fullData[5], fullData[10]-fullData[11], fullData[43]^fullData[13], fullData[45]+fullData[44], fullData[17]^fullData[4])
					return  /*line W5_HrEHk.go:1*/ string(data)
				}() +  /*line W4Lwzvoz.go:1*/ cSW1wSO3.Error())
			}
		} else {
			if olYMzSWK ==  /*line Fv3A9afx.go:1*/ func() string {
				var data []byte
				i := 3
				decryptKey := 125
				for counter := 0; i != 7; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 3:
						data =  /*line Fv3A9afx.go:1*/ append(data, 216)
						i = 0
					case 2:
						i = 4
						data =  /*line Fv3A9afx.go:1*/ append(data, 231)
					case 6:
						data =  /*line Fv3A9afx.go:1*/ append(data, 209)
						i = 5
					case 0:
						i = 1
						data =  /*line Fv3A9afx.go:1*/ append(data, 217)
					case 4:
						i = 7
						for y := range data {
							data[y] = data[y] -  /*line Fv3A9afx.go:1*/ byte(decryptKey^y)
						}
					case 5:
						i = 2
						data =  /*line Fv3A9afx.go:1*/ append(data, 213)
					case 1:
						i = 6
						data =  /*line Fv3A9afx.go:1*/ append(data, 205)
					}
				}
				return  /*line Fv3A9afx.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line UVc6MnhC.go:1*/ func() string {
					data :=  /*line UVc6MnhC.go:1*/ []byte("r=A\x16")
					positions := [...]byte{1, 3, 1, 2}
					for i := 0; i < 4; i += 2 {
						localKey :=  /*line UVc6MnhC.go:1*/ byte(i) +  /*line UVc6MnhC.go:1*/ byte(positions[i]^positions[i+1]) + 215
						data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
					}
					return  /*line UVc6MnhC.go:1*/ string(data)
				}()}
			} else if olYMzSWK ==  /*line D9PAXoFz.go:1*/ func() string {
				var data []byte
				i := 2
				decryptKey := 49
				for counter := 0; i != 6; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 7:
						i = 6
						for y := range data {
							data[y] = data[y] -  /*line D9PAXoFz.go:1*/ byte(decryptKey^y)
						}
					case 2:
						i = 1
						data =  /*line D9PAXoFz.go:1*/ append(data, 146)
					case 0:
						data =  /*line D9PAXoFz.go:1*/ append(data, 154)
						i = 8
					case 8:
						i = 7
						data =  /*line D9PAXoFz.go:1*/ append(data, 152)
					case 4:
						i = 0
						data =  /*line D9PAXoFz.go:1*/ append(data, 143)
					case 3:
						data =  /*line D9PAXoFz.go:1*/ append(data, 148)
						i = 4
					case 1:
						data =  /*line D9PAXoFz.go:1*/ append(data, 136)
						i = 5
					case 5:
						data =  /*line D9PAXoFz.go:1*/ append(data, 146)
						i = 3
					}
				}
				return  /*line D9PAXoFz.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line G9BdVMW9.go:1*/ func() string {
					seed :=  /*line G9BdVMW9.go:1*/ byte(243)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data =  /*line G9BdVMW9.go:1*/ append(data, x+seed); seed += x; return fnc }
					 /*line G9BdVMW9.go:1*/  /*line G9BdVMW9.go:1*/  /*line G9BdVMW9.go:1*/  /*line G9BdVMW9.go:1*/ fnc(127)(243)(252)(3)
					return  /*line G9BdVMW9.go:1*/ string(data)
				}(),  /*line zhSCMdv9.go:1*/ func() string {
					key :=  /*line zhSCMdv9.go:1*/ []byte("\x03*\x0f\xfc\xd4\\")
					data :=  /*line zhSCMdv9.go:1*/ []byte("nEk\x95\xb2%")
					for i, b := range key {
						data[i] = data[i] ^ b
					}
					return  /*line zhSCMdv9.go:1*/ string(data)
				}()}
			}

			hnzizKM2.Permission = iECyqXpw

			jGct6K0c :=  /*line pB2QnkJB.go:1*/ bszIYEJU(hnzizKM2.CustomAccessRights, ezgSKu0t)
			hnzizKM2.CustomAccessRights = jGct6K0c

			niQzOH6M, cSW1wSO3 :=  /*line C4fbTJqi.go:1*/ json.Marshal(hnzizKM2)
			if cSW1wSO3 != nil {
				return  /*line GBUMFsrW.go:1*/ shim.Error( /*line R5Gm5tNz.go:1*/ cSW1wSO3.Error())
			}

			                        
			cSW1wSO3 =  /*line fNKPmP1d.go:1*/ n4c7mRtG.PutState(rHDGEtgU, niQzOH6M)
			if cSW1wSO3 != nil {
				return  /*line qvgDIP0e.go:1*/ shim.Error( /*line DVmedXwA.go:1*/ cSW1wSO3.Error())
			}
		}
	} else if (olYMzSWK ==  /*line rS8VOgbL.go:1*/ func() string {
		seed :=  /*line rS8VOgbL.go:1*/ byte(39)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line rS8VOgbL.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line rS8VOgbL.go:1*/  /*line rS8VOgbL.go:1*/  /*line rS8VOgbL.go:1*/  /*line rS8VOgbL.go:1*/ fnc(85)(25)(244)(237)
		return  /*line rS8VOgbL.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line zPWdojmq.go:1*/ func() string {
		data :=  /*line zPWdojmq.go:1*/ []byte("\xb6\x95\x96\xa3f\xee")
		positions := [...]byte{5, 3, 3, 1, 1, 0, 2, 1}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line zPWdojmq.go:1*/ byte(i) +  /*line zPWdojmq.go:1*/ byte(positions[i]^positions[i+1]) + 208
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line zPWdojmq.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line zMzmp0pj.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 160
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 0:
				i = 6
				data =  /*line zMzmp0pj.go:1*/ append(data, 242)
			case 4:
				data =  /*line zMzmp0pj.go:1*/ append(data, 254)
				i = 8
			case 5:
				data =  /*line zMzmp0pj.go:1*/ append(data, 252)
				i = 4
			case 8:
				i = 2
				for y := range data {
					data[y] = data[y] -  /*line zMzmp0pj.go:1*/ byte(decryptKey^y)
				}
			case 6:
				data =  /*line zMzmp0pj.go:1*/ append(data, 0)
				i = 7
			case 3:
				data =  /*line zMzmp0pj.go:1*/ append(data, 241)
				i = 5
			case 7:
				i = 3
				data =  /*line zMzmp0pj.go:1*/ append(data, 2)
			case 1:
				i = 0
				data =  /*line zMzmp0pj.go:1*/ append(data, 252)
			}
		}
		return  /*line zMzmp0pj.go:1*/ string(data)
	}()) {
		var iECyqXpw []string
		if olYMzSWK ==  /*line Oc5xGWWm.go:1*/ func() string {
			fullData :=  /*line Oc5xGWWm.go:1*/ []byte("\u007f\xd9\x14\u05f8\xe3<f")
			var data []byte
			data =  /*line Oc5xGWWm.go:1*/ append(data, fullData[7]^fullData[2], fullData[6]-fullData[3], fullData[1]^fullData[4], fullData[5]-fullData[0])
			return  /*line Oc5xGWWm.go:1*/ string(data)
		}() {
			cSW1wSO3 =  /*line yUKAcWYq.go:1*/ n4c7mRtG.DelState(rHDGEtgU)	                             
			if cSW1wSO3 != nil {
				return  /*line eat2v3DI.go:1*/ shim.Error( /*line VfolfJYK.go:1*/ func() string {
					key :=  /*line VfolfJYK.go:1*/ []byte("\xb7\xccwN\xa8[<\xa4\xc2vV\x84\xd8\xc4m|\x8b\xa1\xf9\u007f,ԕ")
					data :=  /*line VfolfJYK.go:1*/ []byte("\xf1\xad\x1e\"\xcd?\x1cЭV2ᴡ\x19\x19\xabҍ\x1eX\xb1\xaf")
					for i, b := range key {
						data[i] = data[i] ^ b
					}
					return  /*line VfolfJYK.go:1*/ string(data)
				}() +  /*line wxYfRs7I.go:1*/ cSW1wSO3.Error())
			}
		} else {
			if olYMzSWK ==  /*line oZbeWrSJ.go:1*/ func() string {
				fullData :=  /*line oZbeWrSJ.go:1*/ []byte("\x16T\x8b\xa3\xcfNC\x1b6\xf1Ƣ")
				var data []byte
				data =  /*line oZbeWrSJ.go:1*/ append(data, fullData[4]^fullData[11], fullData[7]+fullData[1], fullData[5]+fullData[0], fullData[10]+fullData[3], fullData[9]-fullData[2], fullData[8]+fullData[6])
				return  /*line oZbeWrSJ.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line QyzcfLHO.go:1*/ func() string {
					seed :=  /*line QyzcfLHO.go:1*/ byte(67)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data =  /*line QyzcfLHO.go:1*/ append(data, x^seed); seed += x; return fnc }
					 /*line QyzcfLHO.go:1*/  /*line QyzcfLHO.go:1*/  /*line QyzcfLHO.go:1*/  /*line QyzcfLHO.go:1*/ fnc(49)(17)(228)(13)
					return  /*line QyzcfLHO.go:1*/ string(data)
				}()}
			} else if olYMzSWK ==  /*line DdTv2S8V.go:1*/ func() string {
				fullData :=  /*line DdTv2S8V.go:1*/ []byte("\x82/a\x83\xed\x82\xca\x15\xa0\xeb\xd3u\xf4\x06")
				var data []byte
				data =  /*line DdTv2S8V.go:1*/ append(data, fullData[4]+fullData[3], fullData[1]-fullData[6], fullData[12]-fullData[5], fullData[13]^fullData[11], fullData[9]^fullData[0], fullData[8]^fullData[10], fullData[7]^fullData[2])
				return  /*line DdTv2S8V.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line flMFJqkp.go:1*/ func() string {
					data :=  /*line flMFJqkp.go:1*/ []byte(">)=8")
					positions := [...]byte{3, 2, 0, 1}
					for i := 0; i < 4; i += 2 {
						localKey :=  /*line flMFJqkp.go:1*/ byte(i) +  /*line flMFJqkp.go:1*/ byte(positions[i]^positions[i+1]) + 88
						data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
					}
					return  /*line flMFJqkp.go:1*/ string(data)
				}(),  /*line yVrsKP0j.go:1*/ func() string {
					key :=  /*line yVrsKP0j.go:1*/ []byte("\xfd\xa6<V\v\xd8")
					data :=  /*line yVrsKP0j.go:1*/ []byte("j\x15\xa0\xbfqQ")
					for i, b := range key {
						data[i] = data[i] - b
					}
					return  /*line yVrsKP0j.go:1*/ string(data)
				}()}
			}

			hnzizKM2.Permission = iECyqXpw

			niQzOH6M, cSW1wSO3 :=  /*line fDZuT4ww.go:1*/ json.Marshal(hnzizKM2)
			if cSW1wSO3 != nil {
				return  /*line bzT3nBQQ.go:1*/ shim.Error( /*line nriWfJPz.go:1*/ cSW1wSO3.Error())
			}

			                        
			cSW1wSO3 =  /*line ooYTz1ss.go:1*/ n4c7mRtG.PutState(rHDGEtgU, niQzOH6M)
			if cSW1wSO3 != nil {
				return  /*line diIXp2mZ.go:1*/ shim.Error( /*line p7Mgx4br.go:1*/ cSW1wSO3.Error())
			}
		}

	} else if (ezgSKu0t[0] != "") || (ezgSKu0t[0] !=  /*line O6ObHMCz.go:1*/ func() string {
		seed :=  /*line O6ObHMCz.go:1*/ byte(134)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line O6ObHMCz.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line O6ObHMCz.go:1*/ fnc(154)
		return  /*line O6ObHMCz.go:1*/ string(data)
	}()) {

		jGct6K0c :=  /*line Hk3BVXnC.go:1*/ bszIYEJU(hnzizKM2.CustomAccessRights, ezgSKu0t)
		hnzizKM2.CustomAccessRights = jGct6K0c

		niQzOH6M, cSW1wSO3 :=  /*line CsHotjSm.go:1*/ json.Marshal(hnzizKM2)
		if cSW1wSO3 != nil {
			return  /*line YYSRtLKd.go:1*/ shim.Error( /*line FjDlbno0.go:1*/ cSW1wSO3.Error())
		}

		                        
		cSW1wSO3 =  /*line J417miMJ.go:1*/ n4c7mRtG.PutState(rHDGEtgU, niQzOH6M)
		if cSW1wSO3 != nil {
			return  /*line EV8l1xzv.go:1*/ shim.Error( /*line Omi2oSmJ.go:1*/ cSW1wSO3.Error())
		}
	}

	                                    
	y5UnTQUP, cSW1wSO3 :=  /*line X96X17tC.go:1*/ baNjNl2z(n4c7mRtG, []string{eviEA3sp, w8EOizcx, dJwgW2jL, olYMzSWK, pbLLcfzS,
		hnzizKM2.Username, hnzizKM2.Name, hnzizKM2.Surname, hnzizKM2.Role,  /*line qtOkWZXk.go:1*/ strings.Join(ezgSKu0t,  /*line sjMasMqt.go:1*/ func() string {
			fullData :=  /*line sjMasMqt.go:1*/ []byte("Jf")
			var data []byte
			data =  /*line sjMasMqt.go:1*/ append(data, fullData[1]^fullData[0])
			return  /*line sjMasMqt.go:1*/ string(data)
		}()), hnzizKM2.DatasetProviderOrg, hnzizKM2.Description})
	if cSW1wSO3 != nil {
		return  /*line CzOJJN4Q.go:1*/ shim.Error( /*line bt1heV01.go:1*/ cSW1wSO3.Error())
	}
	 /*line zFzrPIPt.go:1*/ fmt.Println(y5UnTQUP)

	xHTOVzVs :=  /*line nJ8KV5z5.go:1*/ func() string {
		key :=  /*line nJ8KV5z5.go:1*/ []byte("+ئ\xa8\a\xb9*\x0f\x1b\xf5,\x13\xfa\b\xda\xefY\x89\xe04\x84\xcd\x12\xe9j\v\x0e\x97U\xa8Q\xfc\xb2\xd1)\xa90\x1fgeߍ\xcc$\xa6\xdf")
		data :=  /*line nJ8KV5z5.go:1*/ []byte("\u007f@\v\xc8v+\x91p\x89^\xa6tnqI]y\xf1A\xa7\xa4/wN\xd8+\x80\xfc\xcb\x17\xbca\x16\xf1\x8f\x18\xa2?\xcb\xc6S\xee?\x89\x1a\xff")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line nJ8KV5z5.go:1*/ string(data)
	}() + eviEA3sp
	q8ZF8zaz :=  /*line wt0otFYq.go:1*/ []byte(xHTOVzVs)
	return  /*line R3Rz1DGm.go:1*/ shim.Success(q8ZF8zaz)

}

func (g4rnrSNn *G1Y_7pz_) kWyKFD4c(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	if  /*line B30Vzznd.go:1*/ len(ktsi1_SQ) != 2 {
		return  /*line fs8dxwF3.go:1*/ shim.Error( /*line RcQMnEKw.go:1*/ func() string {
			key :=  /*line RcQMnEKw.go:1*/ []byte("\x1cu\x0f8\x913s\xcf\x0e\x8elf\x19\xd0\u0084\xccW\vͮ\x1d\x1bŊ\xfb\x1eb\xb3\xfd\xc4\xe9\xaf\x12T\xcf \x81\xe3\x90'@l%\xa4 f\x15\x18x\xd2n\xbe\xaa<\x02\x1e\xfc\b/\x90H")
			data :=  /*line RcQMnEKw.go:1*/ []byte("e\xe3r\xa7\x03\xa5\xd82\x82\xae\xdaۆ2'\xf6\xec\xc6q\xed\x0f\x8f\x82:\xf7`\x8c\xd6&+\xe4.'\x82\xb92\x94\xeaQ\xf7G\xb2і\x19\x85ى8\xe16\x8e\xe4ʠc\x92a|\x98\xfd\xad")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line RcQMnEKw.go:1*/ string(data)
		}())
	}
	hSHnto8d := ktsi1_SQ[0]
	pbLLcfzS := ktsi1_SQ[1]

	lWKzd4Pe, cSW1wSO3 :=  /*line ISTCgDMC.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line J38MfqBj.go:1*/ shim.Error( /*line GgbzzQWY.go:1*/ cSW1wSO3.Error())
	}
	qo8GOXik := RequestAccessByOrg{}
	mGZijNgp, cSW1wSO3 :=  /*line _isQi87b.go:1*/ n4c7mRtG.GetState(hSHnto8d)	                                      
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line Iw6xeiCV.go:1*/ func() string {
			data :=  /*line Iw6xeiCV.go:1*/ []byte("A\xd1\xf6rg\x01r:)vF\xc0\xfel$kb$\x9f ghA,rD0\xfeH2t\x1aLc\fHf\x9b\xf1f\xe9r \xd6eR7es\x9a&I4\x8b")
			positions := [...]byte{49, 7, 15, 45, 9, 17, 21, 49, 11, 32, 4, 36, 50, 7, 8, 1, 37, 28, 53, 27, 38, 31, 26, 8, 25, 16, 18, 46, 23, 44, 40, 28, 32, 5, 26, 21, 12, 16, 28, 34, 29, 43, 14, 22, 29, 32, 23, 46, 50, 46, 0, 15, 46, 2, 40, 38, 22, 21, 28, 27, 22, 52, 46, 35, 28, 50, 7, 5}
			for i := 0; i < 68; i += 2 {
				localKey :=  /*line Iw6xeiCV.go:1*/ byte(i) +  /*line Iw6xeiCV.go:1*/ byte(positions[i]^positions[i+1]) + 226
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line Iw6xeiCV.go:1*/ string(data)
		}() + hSHnto8d +  /*line PoKZ5H45.go:1*/ func() string {
			var data []byte
			i := 2
			decryptKey := 123
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 1
					for y := range data {
						data[y] = data[y] ^  /*line PoKZ5H45.go:1*/ byte(decryptKey^y)
					}
				case 2:
					data =  /*line PoKZ5H45.go:1*/ append(data, 95)
					i = 0
				case 0:
					data =  /*line PoKZ5H45.go:1*/ append(data, 1)
					i = 3
				}
			}
			return  /*line PoKZ5H45.go:1*/ string(data)
		}()
		return  /*line WJOMaF8Q.go:1*/ shim.Error(w5VhCLka)
	} else if mGZijNgp == nil {
		w5VhCLka =  /*line q2Ek_LA0.go:1*/ func() string {
			key :=  /*line q2Ek_LA0.go:1*/ []byte("\xde\xeb4T2`Af\xf7\xb1Wa\xa2\xee\x17\x87\x81\xde\xd9\x1c\x97\xcc3X\x12ú\x01\xf1JGWqc")
			data :=  /*line q2Ek_LA0.go:1*/ []byte("\x9d7\x11\x1e@\x0f1\xbcCq\xed\x0eÅ\t\xe7\xee\x96GI\xe1\x9d@\x1ca]\xb8d\x80+\x1e\x1c\x03\xbd")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line q2Ek_LA0.go:1*/ string(data)
		}() + hSHnto8d +  /*line BbyWwtzD.go:1*/ func() string {
			seed :=  /*line BbyWwtzD.go:1*/ byte(59)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line BbyWwtzD.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line BbyWwtzD.go:1*/  /*line BbyWwtzD.go:1*/ fnc(231)(91)
			return  /*line BbyWwtzD.go:1*/ string(data)
		}()
		return  /*line sGC0Jot3.go:1*/ shim.Error(w5VhCLka)

	}

	cSW1wSO3 =  /*line wSysuxPg.go:1*/ json.Unmarshal(mGZijNgp, &qo8GOXik)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line tM1JTrVR.go:1*/ func() string {
			seed :=  /*line tM1JTrVR.go:1*/ byte(188)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line tM1JTrVR.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/  /*line tM1JTrVR.go:1*/ fnc(55)(21)(77)(199)(142)(25)(53)(26)(76)(128)(36)(99)(206)(159)(55)(109)(150)(128)(251)(167)(146)(37)(72)(156)(45)(91)(113)(12)(33)(62)(123)
			return  /*line tM1JTrVR.go:1*/ string(data)
		}() + hSHnto8d +  /*line gdZxzVvo.go:1*/ func() string {
			key :=  /*line gdZxzVvo.go:1*/ []byte("\x1c}")
			data :=  /*line gdZxzVvo.go:1*/ []byte(">\x00")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line gdZxzVvo.go:1*/ string(data)
		}()
		return  /*line f6uGfyXN.go:1*/ shim.Error(w5VhCLka)
	}

	if lWKzd4Pe != qo8GOXik.Dataset_Provider {
		return  /*line RV7LhzbU.go:1*/ shim.Error( /*line Momd_tvu.go:1*/ func() string {
			fullData :=  /*line Momd_tvu.go:1*/ []byte("\xe5=\xe3g\xfe\xb3k\x15\xaa\x029f\xd8I\xb198\xa8!\xfd\x17w\xc1{\x0e\x92tv\xb3\xb3\xf4hv \x19\xc7B\x87\x86i\xd0ΰ\xb7\x95\rOCĊ\xb0|\xf8O\x8f\xb2\xe3\xbc\xcda\xfc&\x0f\xff")
			var data []byte
			data =  /*line Momd_tvu.go:1*/ append(data, fullData[63]-fullData[14], fullData[61]-fullData[43], fullData[48]^fullData[42], fullData[39]-fullData[13], fullData[20]^fullData[32], fullData[47]-fullData[41], fullData[5]^fullData[35], fullData[26]+fullData[30], fullData[49]+fullData[0], fullData[28]^fullData[22], fullData[21]-fullData[24], fullData[29]-fullData[10], fullData[31]+fullData[19], fullData[56]^fullData[37], fullData[8]+fullData[27], fullData[62]^fullData[23], fullData[33]^fullData[46], fullData[52]^fullData[12], fullData[11]+fullData[4], fullData[34]^fullData[51], fullData[57]+fullData[50], fullData[38]^fullData[2], fullData[44]-fullData[18], fullData[40]-fullData[6], fullData[25]^fullData[55], fullData[53]^fullData[1], fullData[3]^fullData[9], fullData[15]+fullData[16], fullData[36]-fullData[58], fullData[45]-fullData[17], fullData[54]^fullData[60], fullData[59]^fullData[7])
			return  /*line Momd_tvu.go:1*/ string(data)
		}())

	}
	eviEA3sp := qo8GOXik.Dataset_Name
	dJwgW2jL := qo8GOXik.Organization
	hOgpzUCn := qo8GOXik.Username
	ipXT6jpu := qo8GOXik.Name
	wSxUCaJ0 := qo8GOXik.Surname
	v5vcwK31 := qo8GOXik.Role

	if qo8GOXik.ObjectType !=  /*line ziszZPNA.go:1*/ func() string {
		fullData :=  /*line ziszZPNA.go:1*/ []byte("\x9eH\xf2\xc1H\xc3\xf1\xd5\xf3\xb0\xb9\xe7\x1a\a$\x96\xea\xb0\x1f\xff\xec\x98\xe7\xb3b\xb2kD\xa6W+\x85N!\xd6\xf2b`\x9eå\x16Z .פ\xf5\x86#.\x90")
		var data []byte
		data =  /*line ziszZPNA.go:1*/ append(data, fullData[11]+fullData[26], fullData[36]^fullData[13], fullData[49]-fullData[25], fullData[15]-fullData[33], fullData[39]^fullData[28], fullData[4]+fullData[30], fullData[42]^fullData[44], fullData[12]^fullData[50], fullData[27]^fullData[41], fullData[29]-fullData[35], fullData[31]+fullData[6], fullData[45]+fullData[21], fullData[14]-fullData[10], fullData[46]^fullData[3], fullData[32]-fullData[16], fullData[23]^fullData[2], fullData[51]^fullData[8], fullData[5]-fullData[37], fullData[17]^fullData[7], fullData[34]^fullData[40], fullData[47]^fullData[48], fullData[24]-fullData[43], fullData[38]^fullData[22], fullData[19]-fullData[9], fullData[20]^fullData[0], fullData[18]+fullData[1])
		return  /*line ziszZPNA.go:1*/ string(data)
	}() {
		               
		xwK9eDsS, cSW1wSO3 :=  /*line D4NKuSpF.go:1*/ o8Z_JkaH(n4c7mRtG, []string{eviEA3sp, lWKzd4Pe, dJwgW2jL, hSHnto8d,  /*line NZfkMz_s.go:1*/ func() string {
			var data []byte
			i := 4
			decryptKey := 87
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 5:
					i = 1
					data =  /*line NZfkMz_s.go:1*/ append(data, 30)
				case 3:
					i = 5
					data =  /*line NZfkMz_s.go:1*/ append(data, 20)
				case 4:
					i = 2
					data =  /*line NZfkMz_s.go:1*/ append(data, 236)
				case 2:
					data =  /*line NZfkMz_s.go:1*/ append(data, 12)
					i = 3
				case 1:
					i = 0
					for y := range data {
						data[y] = data[y] +  /*line NZfkMz_s.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return  /*line NZfkMz_s.go:1*/ string(data)
		}(), pbLLcfzS, hOgpzUCn, ipXT6jpu, wSxUCaJ0, v5vcwK31, qo8GOXik.Permission,  /*line h5UDblOe.go:1*/ strings.Join(qo8GOXik.CustomAccessRights,  /*line HsTs4P2z.go:1*/ func() string {
			key :=  /*line HsTs4P2z.go:1*/ []byte("\x97")
			data :=  /*line HsTs4P2z.go:1*/ []byte("\xbb")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line HsTs4P2z.go:1*/ string(data)
		}()), qo8GOXik.DatasetProviderOrg, qo8GOXik.Description})
		if cSW1wSO3 != nil {
			return  /*line opkeC121.go:1*/ shim.Error( /*line nl8zfcOt.go:1*/ cSW1wSO3.Error())
		}
		 /*line LmOJoGyG.go:1*/ fmt.Println(xwK9eDsS)

	}

	cSW1wSO3 =  /*line AQc0jNAZ.go:1*/ n4c7mRtG.DelState(hSHnto8d)	                                                                
	if cSW1wSO3 != nil {
		return  /*line gZrs1hSI.go:1*/ shim.Error( /*line wPmhFpLT.go:1*/ func() string {
			key :=  /*line wPmhFpLT.go:1*/ []byte("\xbc\x9bhn>W\x97N\xf3\x05\xf9)klo\xfau}\x95\x00\x05\xf6l\x06v&S\xbd&\x9d\x9f\xde")
			data :=  /*line wPmhFpLT.go:1*/ []byte("\xfa\xfa\x01\x02[3\xb7:\x9c%\x9dL\a\t\x1b\x9fU\x0f\xf0qp\x93\x1frVG0\xdeC\xee\xec\xe4")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line wPmhFpLT.go:1*/ string(data)
		}() +  /*line zd3lpmSW.go:1*/ cSW1wSO3.Error())
	}

	xHTOVzVs :=  /*line Xb935Lx8.go:1*/ func() string {
		var data []byte
		i := 1
		decryptKey := 230
		for counter := 0; i != 5; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 10:
				data =  /*line Xb935Lx8.go:1*/ append(data, 113)
				i = 2
			case 0:
				data =  /*line Xb935Lx8.go:1*/ append(data, 109)
				i = 3
			case 9:
				i = 0
				data =  /*line Xb935Lx8.go:1*/ append(data, 106)
			case 8:
				i = 10
				data =  /*line Xb935Lx8.go:1*/ append(data, 113)
			case 4:
				data =  /*line Xb935Lx8.go:1*/ append(data, 19)
				i = 7
			case 6:
				data =  /*line Xb935Lx8.go:1*/ append(data, 95)
				i = 9
			case 2:
				i = 4
				data =  /*line Xb935Lx8.go:1*/ append(data, 54)
			case 1:
				data =  /*line Xb935Lx8.go:1*/ append(data, 77)
				i = 6
			case 7:
				i = 5
				for y := range data {
					data[y] = data[y] -  /*line Xb935Lx8.go:1*/ byte(decryptKey^y)
				}
			case 3:
				i = 8
				data =  /*line Xb935Lx8.go:1*/ append(data, 100)
			}
		}
		return  /*line Xb935Lx8.go:1*/ string(data)
	}() + qo8GOXik.RequestID +  /*line Wa7Y5d36.go:1*/ func() string {
		data :=  /*line Wa7Y5d36.go:1*/ []byte("\xe5{\x9cr\x02\x97Eʵd\x8fta\u007fe\x95 ")
		positions := [...]byte{15, 1, 0, 13, 8, 7, 7, 10, 4, 0, 8, 15, 13, 0, 5, 2, 13, 0, 6, 0}
		for i := 0; i < 20; i += 2 {
			localKey :=  /*line Wa7Y5d36.go:1*/ byte(i) +  /*line Wa7Y5d36.go:1*/ byte(positions[i]^positions[i+1]) + 195
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line Wa7Y5d36.go:1*/ string(data)
	}() + eviEA3sp +  /*line iuFp5JmS.go:1*/ func() string {
		data :=  /*line iuFp5JmS.go:1*/ []byte("\x0f\xcfr \x9c ty\xbd ߴg\x95\x95ϑ\xaftio\xac ")
		positions := [...]byte{3, 11, 0, 11, 10, 13, 8, 3, 17, 11, 7, 7, 1, 15, 1, 3, 14, 15, 7, 0, 17, 16, 4, 7, 21, 13}
		for i := 0; i < 26; i += 2 {
			localKey :=  /*line iuFp5JmS.go:1*/ byte(i) +  /*line iuFp5JmS.go:1*/ byte(positions[i]^positions[i+1]) + 27
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line iuFp5JmS.go:1*/ string(data)
	}() + dJwgW2jL +  /*line u2iy5oE0.go:1*/ func() string {
		key :=  /*line u2iy5oE0.go:1*/ []byte("\x83\xb4l\xba\xd8KԆ\xd6N\x9c\x9b}ձ)\xac\xb1")
		data :=  /*line u2iy5oE0.go:1*/ []byte("\xa3\x1c\xcd-\xf8\xad9\xebDn\x00\x00\xe9:%\x8e\x10\xd1")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line u2iy5oE0.go:1*/ string(data)
	}()
	q8ZF8zaz :=  /*line s3GSMnz4.go:1*/ []byte(xHTOVzVs)
	return  /*line zezPnyye.go:1*/ shim.Success(q8ZF8zaz)
}

func (g4rnrSNn *G1Y_7pz_) awFrSkXv(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	lWKzd4Pe, cSW1wSO3 :=  /*line WUXoeob1.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line z5oxYu25.go:1*/ shim.Error( /*line EJtUcok7.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line Yw04Ug5b.go:1*/ fmt.Sprintf( /*line GLVzumiU.go:1*/ func() string {
		key :=  /*line GLVzumiU.go:1*/ []byte(">\xedtf\\M\xb6Sf\x89CU\xe6a\x12\x80\vF3\xab\xdd#\x9cX\xaf7\xaa~\x92Ӗ\xbb\xe4Ws\x90N\xef\"4\x11C\x1c\x01d\xb3\x90o\t\x00\xbc'/\xf8Yx\xfd\x84\x10'ڟy\x02k\x9e\x91\xfd\xfe")
		data :=  /*line GLVzumiU.go:1*/ []byte("=5\xff\xff\x10\x18\xad!\t\xe9\xdf\xe5\x95\xc1R\xefX\x0eFň\xff\x9eʣ.\xc7\xf7Ӡކ\u007f\f\xf2\xe3%SW\x1ba$\x06+\xbe\xb1\xd1\x05Xs\xa9M0x\x19\xf7y\xe5T>\x98\x83\xc1 \xbaՑ\x80\u007f")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line GLVzumiU.go:1*/ string(data)
	}(), lWKzd4Pe)
	bxM55Mmo, cSW1wSO3 :=  /*line sPTxcoMz.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line RCYYT4IY.go:1*/ shim.Error( /*line V4XEtptx.go:1*/ cSW1wSO3.Error())
	}

	return  /*line WzX8qGzR.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) xdPw2gZb(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line glkLXOvO.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line t7sqHvqJ.go:1*/ shim.Error( /*line dv9x8OrH.go:1*/ func() string {
			fullData :=  /*line dv9x8OrH.go:1*/ []byte("\xcaA\xad\x91\xa2`!\xe3\x06g\xca\xdc\xf5\xa6\x9bB\xcc\r\xfe\xcf@\x11$\x16\xfc\xc2\n\x11\xd8N\xcc?\x16۪\xa9hU\xb4\xe5\xe8\xcfgo\xb0Ֆ\\ez\x01:\x85\x1a\xb3\xa8\xc6\x11\xc8ůCQh\x01\v\xb7K\x93\xb3\xad\x9f_G\xb1\xb8>ѰO\xaaľ!TX\x80\x99HM\xda\x1e,\x9b")
			var data []byte
			data =  /*line dv9x8OrH.go:1*/ append(data, fullData[80]+fullData[71], fullData[87]+fullData[45], fullData[85]+fullData[91], fullData[70]^fullData[30], fullData[88]+fullData[22], fullData[13]^fullData[41], fullData[48]-fullData[50], fullData[36]+fullData[75], fullData[11]-fullData[93], fullData[89]-fullData[33], fullData[24]^fullData[14], fullData[56]^fullData[69], fullData[76]-fullData[77], fullData[84]+fullData[57], fullData[44]-fullData[15], fullData[39]^fullData[3], fullData[26]+fullData[32], fullData[49]-fullData[92], fullData[0]-fullData[37], fullData[52]+fullData[40], fullData[27]+fullData[62], fullData[17]^fullData[63], fullData[25]^fullData[78], fullData[31]-fullData[21], fullData[5]-fullData[20], fullData[51]-fullData[12], fullData[53]-fullData[4], fullData[6]+fullData[79], fullData[35]^fullData[16], fullData[7]-fullData[86], fullData[59]+fullData[60], fullData[58]-fullData[72], fullData[28]+fullData[46], fullData[47]+fullData[65], fullData[1]-fullData[83], fullData[54]+fullData[74], fullData[23]+fullData[67], fullData[82]^fullData[10], fullData[8]^fullData[9], fullData[81]^fullData[66], fullData[19]^fullData[34], fullData[29]-fullData[90], fullData[73]^fullData[42], fullData[43]^fullData[64], fullData[2]+fullData[38], fullData[68]^fullData[18], fullData[55]-fullData[61])
			return  /*line dv9x8OrH.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]
	lWKzd4Pe, cSW1wSO3 :=  /*line I5Dry85L.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line DkVJNmvV.go:1*/ shim.Error( /*line qd9LpI__.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line uIUqDxrr.go:1*/ fmt.Sprintf( /*line D6fqTLpF.go:1*/ func() string {
		fullData :=  /*line D6fqTLpF.go:1*/ []byte("#۵$\x89\x86\x13\xc8\xc9і\x9b\x94\x10\xb3\x87\x9a\uf5f2\x12\x05\xc6R\xec.\x10OV3\xea\xf5\x14\x13\x9do\xad\xfa\x90]\u007f\xbc\x91-\x9cx]%\xf7\x98\xb5\x8eX\xb4xp9\xd7Z\x97ꊡ\xc6@B_\xe9gK\x97\x981\xf2q18\x90\xf3M\xde9:\fď\x12\xaaO;\xa1Z\x84\xf7\xc0K\xab\xe9(\v\xddӺ\xc8\xe4\xcb\"O\xa2\xfdEn\xa7xwȬ\xa3d\xed\xca\x1f\xd3@\x01b\x0e\xc3|\x17\xc0a\xb0b\x02\x14\xf9\x15\xd8`81\x19]wT\xc2 ^\xa3\xfb\b\xd6`\xa8a|^\xc4*\x1eC\xecxG\xed\x8cx\x1a\x120.\x02\nxה\x14")
		var data []byte
		data =  /*line D6fqTLpF.go:1*/ append(data, fullData[162]^fullData[59], fullData[30]-fullData[115], fullData[20]^fullData[155], fullData[145]-fullData[17], fullData[134]^fullData[111], fullData[112]^fullData[146], fullData[176]-fullData[75], fullData[43]+fullData[164], fullData[53]-fullData[110], fullData[106]-fullData[132], fullData[101]+fullData[88], fullData[68]^fullData[143], fullData[51]+fullData[165], fullData[11]+fullData[15], fullData[82]^fullData[148], fullData[71]^fullData[93], fullData[144]^fullData[32], fullData[19]-fullData[157], fullData[84]+fullData[2], fullData[38]-fullData[147], fullData[10]-fullData[72], fullData[35]+fullData[14], fullData[160]-fullData[104], fullData[0]-fullData[124], fullData[127]^fullData[42], fullData[33]+fullData[23], fullData[163]+fullData[136], fullData[105]+fullData[87], fullData[52]-fullData[78], fullData[49]+fullData[1], fullData[122]-fullData[66], fullData[140]-fullData[48], fullData[128]^fullData[121], fullData[54]-fullData[137], fullData[63]^fullData[149], fullData[8]^fullData[102], fullData[150]+fullData[174], fullData[170]+fullData[169], fullData[57]+fullData[108], fullData[16]-fullData[69], fullData[37]+fullData[45], fullData[22]^fullData[90], fullData[99]+fullData[129], fullData[166]-fullData[153], fullData[125]+fullData[130], fullData[114]^fullData[6], fullData[27]^fullData[25], fullData[41]^fullData[103], fullData[81]+fullData[98], fullData[44]+fullData[175], fullData[64]^fullData[47], fullData[118]+fullData[13], fullData[21]+fullData[91], fullData[65]+fullData[171], fullData[167]^fullData[173], fullData[28]^fullData[56], fullData[24]+fullData[61], fullData[156]+fullData[119], fullData[131]-fullData[109], fullData[31]-fullData[77], fullData[9]^fullData[117], fullData[4]^fullData[96], fullData[100]+fullData[39], fullData[58]+fullData[7], fullData[126]-fullData[67], fullData[76]^fullData[95], fullData[74]-fullData[107], fullData[154]+fullData[92], fullData[159]^fullData[151], fullData[60]-fullData[5], fullData[158]+fullData[34], fullData[139]^fullData[135], fullData[79]+fullData[177], fullData[123]^fullData[29], fullData[73]^fullData[70], fullData[161]+fullData[141], fullData[113]-fullData[142], fullData[168]-fullData[116], fullData[18]+fullData[120], fullData[26]+fullData[46], fullData[89]-fullData[152], fullData[3]-fullData[172], fullData[86]-fullData[138], fullData[94]+fullData[133], fullData[55]+fullData[50], fullData[36]^fullData[80], fullData[62]-fullData[40], fullData[97]+fullData[12], fullData[83]-fullData[85])
		return  /*line D6fqTLpF.go:1*/ string(data)
	}(), lWKzd4Pe, eviEA3sp)
	bxM55Mmo, cSW1wSO3 :=  /*line yHolns7c.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line omfGbWCa.go:1*/ shim.Error( /*line AIhnzvJF.go:1*/ cSW1wSO3.Error())
	}

	return  /*line pI9qgYeG.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) hst_3p9W(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	fGjbAOzG, cSW1wSO3 :=  /*line fPnbzbKm.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line kZrnmIaz.go:1*/ shim.Error( /*line ozfE0Wj2.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line AYbPztZb.go:1*/ fmt.Sprintf( /*line jzSjxEeI.go:1*/ func() string {
		seed :=  /*line jzSjxEeI.go:1*/ byte(211)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line jzSjxEeI.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/  /*line jzSjxEeI.go:1*/ fnc(168)(167)(81)(242)(7)(249)(254)(17)(251)(3)(176)(24)(65)(167)(66)(11)(244)(241)(37)(247)(245)(189)(24)(232)(48)(19)(12)(4)(240)(14)(1)(205)(34)(0)(2)(14)(0)(207)(55)(214)(35)(245)(187)(10)(246)(77)(3)(245)(250)(13)(251)(17)(231)(19)(245)(6)(255)(180)(24)(232)(3)(78)(175)(91)(0)
		return  /*line jzSjxEeI.go:1*/ string(data)
	}(), fGjbAOzG)
	bxM55Mmo, cSW1wSO3 :=  /*line mtOB5_hz.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line kPAOKwZ_.go:1*/ shim.Error( /*line fEzbkQaP.go:1*/ cSW1wSO3.Error())
	}

	return  /*line E1xzsH6Q.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) loBDwPCe(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line dhZ3ICyC.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line GAtB6vzZ.go:1*/ shim.Error( /*line yVpFXtIR.go:1*/ func() string {
			fullData :=  /*line yVpFXtIR.go:1*/ []byte("(\xb2c*R)\xc3\xd6\x1f\x19g!\xeaD_\t4C\x1a\xfb\aH\x1d|\x9d\xc5q^\x86t@\x8a\xd8\xe6K\xfa\x8bP\x16\xd4\xc6v\xab\x18\x9e\x02pL3\xc6\xddisX\fiY\x14\x9b\xc2rz\xb1\xfe\x853\xe6\xccv\xeb\u007f\x01\x9c\xe2\xc4T\xf89\xf5\x1e\xe6\xbbVĊ\xc6_\x0f\xb7\xde\xfeV\xa7\x88")
			var data []byte
			data =  /*line yVpFXtIR.go:1*/ append(data, fullData[2]+fullData[80], fullData[55]-fullData[19], fullData[17]+fullData[65], fullData[40]+fullData[58], fullData[88]-fullData[34], fullData[27]-fullData[78], fullData[89]+fullData[28], fullData[5]-fullData[15], fullData[85]-fullData[64], fullData[74]-fullData[4], fullData[8]+fullData[21], fullData[66]-fullData[26], fullData[67]-fullData[86], fullData[91]^fullData[48], fullData[20]^fullData[51], fullData[25]^fullData[62], fullData[54]+fullData[57], fullData[32]-fullData[84], fullData[36]+fullData[12], fullData[11]^fullData[47], fullData[44]+fullData[83], fullData[23]^fullData[9], fullData[52]-fullData[71], fullData[0]-fullData[35], fullData[46]-fullData[37], fullData[18]^fullData[14], fullData[81]^fullData[6], fullData[72]+fullData[39], fullData[69]+fullData[61], fullData[60]-fullData[87], fullData[31]-fullData[38], fullData[59]^fullData[42], fullData[68]^fullData[43], fullData[33]-fullData[70], fullData[29]^fullData[75], fullData[7]^fullData[1], fullData[13]+fullData[22], fullData[82]-fullData[73], fullData[63]-fullData[24], fullData[3]^fullData[56], fullData[50]+fullData[93], fullData[30]+fullData[16], fullData[79]-fullData[90], fullData[41]+fullData[76], fullData[77]^fullData[53], fullData[92]+fullData[49], fullData[45]^fullData[10])
			return  /*line yVpFXtIR.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]
	fGjbAOzG, cSW1wSO3 :=  /*line rPBsk6cy.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line pxz2YEnT.go:1*/ shim.Error( /*line Zzl2Ggey.go:1*/ cSW1wSO3.Error())
	}
	                  
	qVH0siY7 :=  /*line Cs_O_FP2.go:1*/ fmt.Sprintf( /*line Pz3lZhm6.go:1*/ func() string {
		data :=  /*line Pz3lZhm6.go:1*/ []byte("{\xd1\xf1G/\nc\xdc?r\r\a{\xd3do\xaf\xfaype\xfd%\"\xf8e\x97\r\x1b\x92\xd5Ac\xafe\u07bcFy38g\"\\\xd1o\xc4!anizatrYn^\xa3\x16%s\"\xf75d\x93R\xc0\xc8e\xdd\xde\xf4\xac\xcce\r:\x96%sn}}")
		positions := [...]byte{68, 22, 59, 74, 7, 22, 57, 26, 75, 75, 59, 75, 43, 69, 28, 7, 24, 47, 8, 37, 3, 46, 82, 33, 67, 63, 33, 66, 39, 27, 5, 74, 54, 7, 28, 69, 21, 58, 77, 5, 55, 72, 21, 28, 33, 11, 29, 47, 21, 2, 73, 28, 63, 66, 4, 16, 40, 33, 44, 1, 63, 79, 67, 17, 79, 37, 22, 13, 64, 7, 7, 79, 35, 71, 30, 72, 3, 22, 10, 36, 47, 17, 66, 57, 37, 4}
		for i := 0; i < 86; i += 2 {
			localKey :=  /*line Pz3lZhm6.go:1*/ byte(i) +  /*line Pz3lZhm6.go:1*/ byte(positions[i]^positions[i+1]) + 234
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line Pz3lZhm6.go:1*/ string(data)
	}(), fGjbAOzG, eviEA3sp)
	bxM55Mmo, cSW1wSO3 :=  /*line xPKn4W6o.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line Ioztexem.go:1*/ shim.Error( /*line AXL1Ixoa.go:1*/ cSW1wSO3.Error())
	}

	return  /*line lUgZEeAG.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) iWHhTpxb(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	fGjbAOzG, cSW1wSO3 :=  /*line SY6VOVi7.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line Jr6sMakR.go:1*/ shim.Error( /*line aRc7hfJg.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line XgsGqx1O.go:1*/ fmt.Sprintf( /*line tdqpS1Yq.go:1*/ func() string {
		seed :=  /*line tdqpS1Yq.go:1*/ byte(151)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line tdqpS1Yq.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/  /*line tdqpS1Yq.go:1*/ fnc(236)(161)(87)(30)(245)(235)(26)(231)(21)(253)(174)(0)(65)(89)(176)(235)(12)(47)(211)(13)(239)(91)(238)(224)(227)(240)(1)(30)(251)(253)(229)(11)(25)(241)(201)(61)(235)(4)(89)(248)(238)(213)(253)(235)(22)(227)(25)(243)(29)(237)(239)(26)(225)(82)(248)(152)(119)(186)(161)(89)(0)
		return  /*line tdqpS1Yq.go:1*/ string(data)
	}(), fGjbAOzG)
	bxM55Mmo, cSW1wSO3 :=  /*line _RLShhSI.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line xWfmd0OL.go:1*/ shim.Error( /*line v2e7P26g.go:1*/ cSW1wSO3.Error())
	}

	return  /*line e1KmeCwZ.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) v8QuhfD6(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line RqIYnNoH.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line En8bmL1l.go:1*/ shim.Error( /*line DnN2tpTg.go:1*/ func() string {
			seed :=  /*line DnN2tpTg.go:1*/ byte(29)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line DnN2tpTg.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/  /*line DnN2tpTg.go:1*/ fnc(102)(241)(234)(191)(137)(15)(25)(238)(253)(43)(75)(164)(64)(120)(249)(248)(156)(102)(243)(222)(177)(101)(215)(106)(198)(177)(149)(34)(57)(112)(241)(215)(179)(95)(119)(50)(97)(213)(151)(64)(114)(243)(146)(114)(215)(186)(108)
			return  /*line DnN2tpTg.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]
	fGjbAOzG, cSW1wSO3 :=  /*line UrHH8Ggq.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line T_6fwXoI.go:1*/ shim.Error( /*line Q5euyStz.go:1*/ cSW1wSO3.Error())
	}
	                  
	qVH0siY7 :=  /*line VeK_lmAi.go:1*/ fmt.Sprintf( /*line N1AXFBNA.go:1*/ func() string {
		fullData :=  /*line N1AXFBNA.go:1*/ []byte("g&1\xe8\vӀ\x9a7\xda\f\x14\x8dD_\xae\xa9\xb8\xfb\x15\x97-\x1a\xb9\x9eږ\f\xe5\xac\x01\x00\x041\x97y\xe6(\xc7\r\xfbΪ\xdf/\xe5D(Z\xb3\xac)Yɬ纴Y\xf2\x0fA`\xa4R\xf6$#\xd71\xae\xc5\xc5鬧o\xb8\xd7\xe7[e\xa6\xb1\x8e\xe24\x87\x00\xbdix\xd2s \xab\x1d\x91[I\xf7\x0fb\x83\xb3\x93\xe5ytc;2cE,a\x9d\x13\xa1\xc546\xe6l\xe2\xc7\xfe\x15Q{\x89I-\b;\x9e\xb6\xfao\x8dGC\x06\xfe\xa3\xa3\xf9\xccAx\x118\x0e\x1b]\xd5g\a\x13\xf6'L")
		var data []byte
		data =  /*line N1AXFBNA.go:1*/ append(data, fullData[51]^fullData[64], fullData[52]+fullData[53], fullData[22]-fullData[75], fullData[6]^fullData[106], fullData[49]-fullData[140], fullData[121]+fullData[44], fullData[71]-fullData[102], fullData[156]+fullData[39], fullData[111]^fullData[154], fullData[2]^fullData[141], fullData[41]-fullData[54], fullData[38]+fullData[93], fullData[9]-fullData[14], fullData[29]^fullData[84], fullData[48]-fullData[65], fullData[101]+fullData[62], fullData[109]+fullData[88], fullData[100]-fullData[144], fullData[72]+fullData[57], fullData[34]-fullData[160], fullData[43]^fullData[56], fullData[1]^fullData[32], fullData[63]^fullData[135], fullData[60]+fullData[117], fullData[143]-fullData[89], fullData[158]-fullData[24], fullData[118]-fullData[21], fullData[112]-fullData[40], fullData[138]^fullData[31], fullData[45]^fullData[20], fullData[81]^fullData[27], fullData[95]-fullData[69], fullData[86]+fullData[33], fullData[161]^fullData[37], fullData[116]^fullData[92], fullData[127]-fullData[145], fullData[18]+fullData[123], fullData[7]^fullData[73], fullData[58]^fullData[129], fullData[150]-fullData[28], fullData[130]-fullData[0], fullData[139]^fullData[85], fullData[87]-fullData[19], fullData[15]+fullData[23], fullData[42]-fullData[131], fullData[148]-fullData[5], fullData[128]-fullData[3], fullData[67]-fullData[16], fullData[59]-fullData[97], fullData[74]-fullData[151], fullData[96]^fullData[108], fullData[79]-fullData[91], fullData[113]-fullData[78], fullData[105]^fullData[83], fullData[159]+fullData[46], fullData[98]^fullData[35], fullData[114]+fullData[146], fullData[125]+fullData[50], fullData[152]+fullData[11], fullData[133]^fullData[66], fullData[107]^fullData[80], fullData[132]+fullData[8], fullData[94]^fullData[61], fullData[124]^fullData[26], fullData[110]-fullData[25], fullData[70]+fullData[119], fullData[90]^fullData[10], fullData[17]^fullData[147], fullData[77]^fullData[55], fullData[36]-fullData[149], fullData[136]^fullData[68], fullData[126]+fullData[76], fullData[4]-fullData[82], fullData[155]-fullData[104], fullData[30]^fullData[134], fullData[153]+fullData[157], fullData[13]^fullData[115], fullData[12]+fullData[122], fullData[47]-fullData[142], fullData[137]+fullData[103], fullData[120]+fullData[99])
		return  /*line N1AXFBNA.go:1*/ string(data)
	}(), fGjbAOzG, eviEA3sp)
	bxM55Mmo, cSW1wSO3 :=  /*line PlAU9sTW.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line y68RAFTY.go:1*/ shim.Error( /*line yhFGrgF9.go:1*/ cSW1wSO3.Error())
	}

	return  /*line wNLYu9nj.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) wtNtF9Dv(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	lWKzd4Pe, cSW1wSO3 :=  /*line lrfKOpEK.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line r0RxGscM.go:1*/ shim.Error( /*line XAAtjSgT.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line dcwJ3pPK.go:1*/ fmt.Sprintf( /*line CyLc8znB.go:1*/ func() string {
		seed :=  /*line CyLc8znB.go:1*/ byte(190)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line CyLc8znB.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/  /*line CyLc8znB.go:1*/ fnc(197)(161)(87)(30)(245)(235)(26)(231)(21)(253)(174)(0)(65)(89)(176)(235)(12)(47)(211)(13)(239)(91)(238)(224)(227)(240)(1)(30)(251)(253)(229)(11)(25)(241)(201)(61)(235)(4)(89)(248)(238)(222)(249)(229)(23)(254)(238)(13)(217)(47)(252)(229)(25)(225)(13)(19)(251)(166)(16)(24)(119)(186)(161)(89)(0)
		return  /*line CyLc8znB.go:1*/ string(data)
	}(), lWKzd4Pe)
	bxM55Mmo, cSW1wSO3 :=  /*line HHLVDMxE.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line I4Qb6e3Y.go:1*/ shim.Error( /*line wFdjwFcY.go:1*/ cSW1wSO3.Error())
	}

	return  /*line p9frJQoO.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) aqKW3Kcg(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line BsoVtus8.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line dAc5eout.go:1*/ shim.Error( /*line r6TuzlD6.go:1*/ func() string {
			data :=  /*line r6TuzlD6.go:1*/ []byte("\x8env\xf4l\x96d )~Lc\xfben\x89\xb6\xaeumb\xa3\xbd| zϴec\x8blngB\x91\x80t\xe1\x9a]> \x8f\xb4\xcae")
			positions := [...]byte{44, 16, 38, 0, 22, 26, 12, 21, 27, 0, 3, 23, 8, 0, 36, 10, 23, 0, 22, 17, 21, 5, 15, 43, 3, 23, 30, 5, 3, 27, 39, 45, 44, 39, 44, 5, 34, 23, 9, 36, 16, 23, 31, 40, 25, 41, 35, 41, 0, 11, 36, 40, 9, 23}
			for i := 0; i < 54; i += 2 {
				localKey :=  /*line r6TuzlD6.go:1*/ byte(i) +  /*line r6TuzlD6.go:1*/ byte(positions[i]^positions[i+1]) + 171
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line r6TuzlD6.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]
	lWKzd4Pe, cSW1wSO3 :=  /*line NqbQo9xA.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line iiyAuix6.go:1*/ shim.Error( /*line qMxxYNdZ.go:1*/ cSW1wSO3.Error())
	}
	                  
	qVH0siY7 :=  /*line mOiajz8E.go:1*/ fmt.Sprintf( /*line pmM_ZAKb.go:1*/ func() string {
		key :=  /*line pmM_ZAKb.go:1*/ []byte("\x12\xb8\x17^c'\x8d\xfd\x80\xe8\xb2\x01z\xa0#]\x11#\xf8\x18[\x00\x05\xbb>T\x8785u\x01\x9e.'\x06\xe2\x16q\xc0\xecaL\b֝{\xac}8ę!\xc2\xe3J}\xa7\x8e\v\xcf\xeeK\x0fI\x97/;\x9cZ\x93J\xd6\x12|\xa0\xd6R\x9ev\xfb[\xe10\x00\x8b")
		data :=  /*line pmM_ZAKb.go:1*/ []byte("\x8dڊ\xc3ό\xf0q\xefZ\xd4;\xf5\u0087\xcctwq\x88\xc0\"?\xdd\u007f\xc9\xfb\xa0\xa4\xe7j\x18\x93\x8bUT}\xe4\xe2\x18\x83\xb0iJ\xfe\xee\x11\xf1\x974\v\x908L\xae\xe2\x19\xb0E\xf1\x13\xbe1u\xb9\x93\x9c\x10\xbb\x06\xafJq\xea\x01C\xb7\xc0\xb0\x1d\x80TR}\b")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line pmM_ZAKb.go:1*/ string(data)
	}(), lWKzd4Pe, eviEA3sp)
	bxM55Mmo, cSW1wSO3 :=  /*line IDq9ePRW.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line H70nNeLi.go:1*/ shim.Error( /*line KvTXs9bn.go:1*/ cSW1wSO3.Error())
	}

	return  /*line S6L0ztlf.go:1*/ shim.Success(bxM55Mmo)
}

                                   

func (g4rnrSNn *G1Y_7pz_) v4t5SrhO(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	fGjbAOzG, cSW1wSO3 :=  /*line mcYAdd8T.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line K2AUl27h.go:1*/ shim.Error( /*line F2oox62J.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line FWWM2P29.go:1*/ fmt.Sprintf( /*line gDnxzX6D.go:1*/ func() string {
		key :=  /*line gDnxzX6D.go:1*/ []byte("O\xcfR\x8c\x85\xf3\x90\xb18\x1a\t\x867C6<\b3\x83\x97\x8c\xa8+\xa2;X\xd0\xcfŨ\xae\xb1߶\x0e\v\x11\x13\x1e\x16!\x02\x9e.D\x94֟\x85F\xb3'O3`\xc8h\xb8\t0\x1c\xfa\xde\xfd\xf6t\xe6\xc7`\b\xa6!*")
		data :=  /*line gDnxzX6D.go:1*/ []byte("4\xed!\xe9\xe9\x96\xf3\xc5Wh+\xbcLaRSkg\xfa\xe7\xe9\x8a\x11\x80i=\xa1\xba\xa0\xdbڅ\x8d\xd3xdzvzWBa\xfb]7֯\xd0\xf7!\x91\vm\\\x12\xaf\t\xd6`J}\x8e\xb7\x92\x98V\xdc\xe5E{\x84\\W")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line gDnxzX6D.go:1*/ string(data)
	}(), fGjbAOzG)
	bxM55Mmo, cSW1wSO3 :=  /*line mgHYhkDN.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line WN32UHxk.go:1*/ shim.Error( /*line BctqgjXM.go:1*/ cSW1wSO3.Error())
	}

	return  /*line nsB4wiTS.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) je7g236F(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line Cfw9hq9z.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line VLssNbZ6.go:1*/ shim.Error( /*line W3D82QIr.go:1*/ func() string {
			seed :=  /*line W3D82QIr.go:1*/ byte(22)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line W3D82QIr.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/  /*line W3D82QIr.go:1*/ fnc(95)(27)(230)(23)(225)(7)(17)(166)(109)(235)(227)(18)(20)(232)(27)(228)(84)(134)(59)(228)(15)(25)(231)(82)(238)(249)(205)(242)(17)(230)(31)(227)(3)(23)(167)(74)(25)(229)(23)(254)(238)(13)(166)(66)(15)(16)(232)
			return  /*line W3D82QIr.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]
	fGjbAOzG, cSW1wSO3 :=  /*line I3CcVcLr.go:1*/ g4rnrSNn.vNURlL_g(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line bTr_GVG2.go:1*/ shim.Error( /*line G4THf4cs.go:1*/ cSW1wSO3.Error())
	}
	                  
	qVH0siY7 :=  /*line Cdbk33Ep.go:1*/ fmt.Sprintf( /*line bm7B0lOZ.go:1*/ func() string {
		key :=  /*line bm7B0lOZ.go:1*/ []byte("8##\xa9\xe5}$\x8d\x1b#F\x10\x83k\x19b\v\xcf͐N\x8f6\x94r\x1d\xe6\x0e\xe7{벓\x92\x17\xa9=^?S3\x97\xfeg\xbdx5D\x83\xec\rUi\x0fCV\x84\x14\xc0&\xabn\x06f\v\x00_\xd2<X!\xbd\xae\x1b\xc6\xf4\x9a\x1931\x00\x8aS?V\xf7\x9b\x97\xb0D\x8f\xa9\xc4")
		data :=  /*line bm7B0lOZ.go:1*/ []byte("\xb3E\x96\x0eQ\xe2\x87\x01\x8a\x95hJ\xfe\x8d}\xd1n#F\x00\xb3\xb1p\xb6ĂW\x83L\xee_\xe6\xe5\xf7\x8d\x18\xa8ã\x94\x96\xfac\xda0\xba\xae\x93\xf5S/\x81\x8b~\xb5\xbd\xe5\x82)\xa0\f\xe2o\xd5y\"\x99\xf4a\xcbC\xe9\xd0\u007f'h\xfb\x8c\x98\xa5_\xf8\xb4\xac\xbb\x19չշ\xb1&A")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line bm7B0lOZ.go:1*/ string(data)
	}(), fGjbAOzG, eviEA3sp)
	bxM55Mmo, cSW1wSO3 :=  /*line hf0n5FQY.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line IbS23pI5.go:1*/ shim.Error( /*line cpg7izvU.go:1*/ cSW1wSO3.Error())
	}

	return  /*line mFy3T4T7.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) cQ4WuEb4(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	lWKzd4Pe, cSW1wSO3 :=  /*line PphNz1V6.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line F4Ai6Deh.go:1*/ shim.Error( /*line NsR0QjCL.go:1*/ cSW1wSO3.Error())
	}

	                  
	qVH0siY7 :=  /*line Ii976uly.go:1*/ fmt.Sprintf( /*line de1oGxsP.go:1*/ func() string {
		data :=  /*line de1oGxsP.go:1*/ []byte("{\xb3T\xaa\x95\x87cfo!6v{\"=a\u007fKK\x88f\xdd\x1b\"Rg\x87q\xac\x92\x89C\xb06voke&zcNus\xe8B\xf6\x01\x8ag\"\x93;:a\x98og͡\xf8p{\\\xfa\xa8der\"\xe8\"\xb9s\"ޞ")
		positions := [...]byte{31, 47, 52, 59, 56, 18, 17, 27, 63, 20, 38, 64, 57, 29, 14, 72, 7, 42, 31, 10, 46, 11, 76, 58, 32, 25, 31, 22, 9, 30, 59, 56, 64, 53, 41, 65, 30, 22, 53, 38, 53, 65, 58, 33, 3, 15, 9, 53, 27, 5, 16, 56, 38, 25, 18, 2, 41, 21, 33, 39, 57, 57, 18, 52, 44, 75, 64, 60, 58, 14, 62, 48, 4, 41, 26, 19, 1, 48, 11, 33, 28, 57, 51, 41, 18, 15, 17, 14, 70, 33, 17, 55}
		for i := 0; i < 92; i += 2 {
			localKey :=  /*line de1oGxsP.go:1*/ byte(i) +  /*line de1oGxsP.go:1*/ byte(positions[i]^positions[i+1]) + 196
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line de1oGxsP.go:1*/ string(data)
	}(), lWKzd4Pe)
	bxM55Mmo, cSW1wSO3 :=  /*line bMSzJxEb.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line AN9zZmh3.go:1*/ shim.Error( /*line O2LtTZat.go:1*/ cSW1wSO3.Error())
	}

	return  /*line BcKpD6Z4.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) zRIYep9j(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line CBrNjyDo.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line d_Zu7vMU.go:1*/ shim.Error( /*line SjL_Pcyz.go:1*/ func() string {
			var data []byte
			i := 7
			decryptKey := 243
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 19:
					data =  /*line SjL_Pcyz.go:1*/ append(data, "K_\x02"...,
					)
					i = 10
				case 17:
					data =  /*line SjL_Pcyz.go:1*/ append(data, "ov"...,
					)
					i = 2
				case 2:
					i = 14
					data =  /*line SjL_Pcyz.go:1*/ append(data, "f1~"...,
					)
				case 10:
					i = 9
					data =  /*line SjL_Pcyz.go:1*/ append(data, "\x03gYP"...,
					)
				case 15:
					data =  /*line SjL_Pcyz.go:1*/ append(data, 77)
					i = 5
				case 6:
					i = 15
					data =  /*line SjL_Pcyz.go:1*/ append(data, "EQ"...,
					)
				case 18:
					i = 19
					data =  /*line SjL_Pcyz.go:1*/ append(data, "EM"...,
					)
				case 1:
					data =  /*line SjL_Pcyz.go:1*/ append(data, "S[@"...,
					)
					i = 4
				case 12:
					for y := range data {
						data[y] = data[y] ^  /*line SjL_Pcyz.go:1*/ byte(decryptKey^y)
					}
					i = 8
				case 16:
					i = 17
					data =  /*line SjL_Pcyz.go:1*/ append(data, "j|"...,
					)
				case 13:
					data =  /*line SjL_Pcyz.go:1*/ append(data, "Y\x1cr@"...,
					)
					i = 3
				case 11:
					data =  /*line SjL_Pcyz.go:1*/ append(data, "9|~"...,
					)
					i = 16
				case 4:
					i = 18
					data =  /*line SjL_Pcyz.go:1*/ append(data, "\vd\\"...,
					)
				case 3:
					i = 1
					data =  /*line SjL_Pcyz.go:1*/ append(data, "VEZ"...,
					)
				case 5:
					i = 11
					data =  /*line SjL_Pcyz.go:1*/ append(data, "u}"...,
					)
				case 9:
					data =  /*line SjL_Pcyz.go:1*/ append(data, 66)
					i = 6
				case 7:
					i = 0
					data =  /*line SjL_Pcyz.go:1*/ append(data, "rT"...,
					)
				case 0:
					i = 13
					data =  /*line SjL_Pcyz.go:1*/ append(data, "OYSW"...,
					)
				case 14:
					data =  /*line SjL_Pcyz.go:1*/ append(data, "v{p"...,
					)
					i = 12
				}
			}
			return  /*line SjL_Pcyz.go:1*/ string(data)
		}())
	}
	eviEA3sp := ktsi1_SQ[0]
	lWKzd4Pe, cSW1wSO3 :=  /*line w1hT6ZJL.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line ERF2Ruuw.go:1*/ shim.Error( /*line sCRk26v5.go:1*/ cSW1wSO3.Error())
	}
	                  
	qVH0siY7 :=  /*line hA3s0zFF.go:1*/ fmt.Sprintf( /*line HUxNlIBS.go:1*/ func() string {
		key :=  /*line HUxNlIBS.go:1*/ []byte("\x98\x18\xc65\xcc7כN\x15{Ğ\xa9\xd5y¡ܪ\xe9\xa2F\xa0\xa0i\x05~;U\xc8\xe2\x8c\xc6᪠\xa7\xdf@\x11\xa4\xaa\x91\xa5N~6\xc54\x82\xabf\x03=J\\\x1c\x92\xdc\xc0\v\x9bߞU\xab\xd2ݡ\x8e\x15\x01\xae\x02I5\xe0\xa4o.D$\x83*E\vp\xed\xabY\x9aК\xb5VB")
		data :=  /*line HUxNlIBS.go:1*/ []byte("\x13:9\x9a8\x9c:\x0f\xbd\x87\x9d\xfe\x19\xcb9\xe8%\xf5U\x1aNĀ\xc2\xf2\xcev\xf3\xa0\xc8<\x16\xde+W\x19\v\fC\x81t\a\x0f\x04\x18\x90\xf7\x857\x9b\xa4\u05c8g\x9e\xbe\xbd\x8f\xf7P\x1f{\rN\x14\xbe\x0f7O\xc3\xc87&!$uWD\x05㏷\x89\xf7\x89\xb3l\xddR͓\xbc\xf5\r\xd7ӿ")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line HUxNlIBS.go:1*/ string(data)
	}(), lWKzd4Pe, eviEA3sp)
	bxM55Mmo, cSW1wSO3 :=  /*line X88_iGsH.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line iyubX_UA.go:1*/ shim.Error( /*line jEU6oc89.go:1*/ cSW1wSO3.Error())
	}

	return  /*line Gu0TE4Wb.go:1*/ shim.Success(bxM55Mmo)
}
