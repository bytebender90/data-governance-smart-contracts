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
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (g4rnrSNn *G1Y_7pz_) i0TROu9J(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	if  /*line _Klm4nDl.go:1*/ len(ktsi1_SQ) != 3 {
		return  /*line AMjUCYAb.go:1*/ shim.Error( /*line JiyOtm0z.go:1*/ func() string {
			data :=  /*line JiyOtm0z.go:1*/ []byte("H^#ol=3vA\x82guTb6t vO\xf5SeBk\xddEtp[c\x84epX1) d\r?\xdbs[}\x84n+m֍\x8f\xfa Sa\"\x8b\f3b \x04ey")
			positions := [...]byte{40, 1, 20, 40, 48, 33, 5, 53, 6, 53, 32, 30, 42, 28, 61, 57, 59, 18, 46, 57, 56, 13, 24, 19, 39, 17, 57, 38, 19, 22, 51, 33, 22, 0, 26, 32, 7, 1, 48, 1, 43, 20, 9, 28, 2, 23, 63, 13, 50, 12, 33, 49, 22, 3, 59, 17, 55, 33, 48, 1, 55, 28, 63, 28, 14, 44}
			for i := 0; i < 66; i += 2 {
				localKey :=  /*line JiyOtm0z.go:1*/ byte(i) +  /*line JiyOtm0z.go:1*/ byte(positions[i]^positions[i+1]) + 180
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line JiyOtm0z.go:1*/ string(data)
		}())
	}
	mzjt9rLK :=  /*line IsMKRC5f.go:1*/ func() string {
		data :=  /*line IsMKRC5f.go:1*/ []byte("\x86a\xed}s\xd2\xc2rimi\xed`")
		positions := [...]byte{3, 6, 7, 0, 7, 6, 12, 3, 7, 5, 3, 5, 12, 0, 2, 11}
		for i := 0; i < 16; i += 2 {
			localKey :=  /*line IsMKRC5f.go:1*/ byte(i) +  /*line IsMKRC5f.go:1*/ byte(positions[i]^positions[i+1]) + 112
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line IsMKRC5f.go:1*/ string(data)
	}()
	eviEA3sp := ktsi1_SQ[0]

	if  /*line _sZeDnKj.go:1*/ len(eviEA3sp) <= 0 {
		return  /*line DEI61NHn.go:1*/ shim.Error( /*line bYCgURUn.go:1*/ func() string {
			fullData :=  /*line bYCgURUn.go:1*/ []byte("RF%sr\x14\xa7GВ\x1adڞ\xe9\u007fy`p\b\xbbT\x90\x00\xe7<\x87Áf\x01O(\xef\xc9r\xba3\x19\x8a4\x00칲\"\x82l\xd8p\xe3\x93\xf3\x95\"\xf1\xc0\x8d\xa3\xdea\x8d\x0e\xec\x17\f\x95V\n\x86+\xe1\x1f\xbb\x93x\xf5]")
			var data []byte
			data =  /*line bYCgURUn.go:1*/ append(data, fullData[75]-fullData[40], fullData[63]^fullData[61], fullData[52]-fullData[15], fullData[43]^fullData[48], fullData[41]^fullData[3], fullData[2]-fullData[56], fullData[71]+fullData[51], fullData[29]-fullData[1], fullData[11]+fullData[68], fullData[20]^fullData[12], fullData[66]-fullData[32], fullData[42]+fullData[16], fullData[6]-fullData[26], fullData[47]^fullData[30], fullData[36]+fullData[73], fullData[59]+fullData[53], fullData[67]^fullData[54], fullData[57]+fullData[74], fullData[50]-fullData[28], fullData[77]+fullData[19], fullData[9]^fullData[44], fullData[55]+fullData[49], fullData[14]-fullData[34], fullData[65]-fullData[13], fullData[60]+fullData[62], fullData[22]-fullData[45], fullData[76]+fullData[70], fullData[21]-fullData[33], fullData[69]-fullData[38], fullData[39]-fullData[10], fullData[37]^fullData[7], fullData[25]-fullData[27], fullData[35]-fullData[0], fullData[58]^fullData[8], fullData[17]+fullData[5], fullData[23]^fullData[4], fullData[24]+fullData[46], fullData[31]+fullData[72], fullData[18]^fullData[64])
			return  /*line bYCgURUn.go:1*/ string(data)
		}())
	}

	eBDS4r4c := ktsi1_SQ[1]

	if  /*line GYgLp51Q.go:1*/ len(eBDS4r4c) <= 0 {
		return  /*line CoX0fR4E.go:1*/ shim.Error( /*line ujo0zCvB.go:1*/ func() string {
			fullData :=  /*line ujo0zCvB.go:1*/ []byte("b\xf3\u007f\xe7\xfc\x8f\x8a\x12\x80\x91\x12\x9e\xd40\"\xda\x02^҂\x06\xebS\x95\xb1\xce\xcf\xf2\x83\xc7-\x14\xb3\x9ft\xc0\xae\xeeďr\xf9K\xff\x1ck\xae~S$o\xb0\xdb\x04\xafu\x90\xa7\xcbe\x1e\x1b\xfe\xb8V\u007fej\xe0\xd6\xf1\x8do\x80}\x1e\xfbl")
			var data []byte
			data =  /*line ujo0zCvB.go:1*/ append(data, fullData[38]+fullData[73], fullData[55]^fullData[31], fullData[68]-fullData[77], fullData[52]+fullData[6], fullData[58]^fullData[21], fullData[40]+fullData[4], fullData[74]^fullData[44], fullData[18]-fullData[66], fullData[63]-fullData[22], fullData[3]^fullData[29], fullData[72]-fullData[16], fullData[48]+fullData[14], fullData[30]^fullData[17], fullData[41]^fullData[71], fullData[45]^fullData[42], fullData[70]+fullData[47], fullData[35]+fullData[36], fullData[37]^fullData[25], fullData[67]^fullData[75], fullData[57]^fullData[26], fullData[69]^fullData[32], fullData[39]-fullData[50], fullData[23]^fullData[1], fullData[43]-fullData[56], fullData[62]+fullData[34], fullData[7]^fullData[65], fullData[59]^fullData[53], fullData[5]-fullData[61], fullData[19]-fullData[0], fullData[28]^fullData[24], fullData[10]+fullData[60], fullData[15]+fullData[64], fullData[20]^fullData[13], fullData[11]+fullData[9], fullData[51]+fullData[8], fullData[33]^fullData[46], fullData[76]^fullData[12], fullData[54]-fullData[2], fullData[49]-fullData[27])
			return  /*line ujo0zCvB.go:1*/ string(data)
		}())
	}
	gZATGwlS := ktsi1_SQ[2]
	if  /*line Dm4rdfVj.go:1*/ len(gZATGwlS) <= 0 {
		gZATGwlS =  /*line aJ4yjHXe.go:1*/ func() string {
			var data []byte
			i := 0
			decryptKey := 211
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					i = 1
					data =  /*line aJ4yjHXe.go:1*/ append(data, 2)
				case 1:
					for y := range data {
						data[y] = data[y] -  /*line aJ4yjHXe.go:1*/ byte(decryptKey^y)
					}
					i = 2
				}
			}
			return  /*line aJ4yjHXe.go:1*/ string(data)
		}()
	}

	tAfFZezC, cSW1wSO3 :=  /*line NGwHItg2.go:1*/ strconv.Atoi(gZATGwlS)
	if cSW1wSO3 != nil {
		 /*line l8WVeATV.go:1*/ fmt.Println(cSW1wSO3)

	}

	w8EOizcx, cSW1wSO3 :=  /*line ZRZRnsct.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line SBDgu6qn.go:1*/ shim.Error( /*line Y22osJtJ.go:1*/ cSW1wSO3.Error())
	}

	                                            
	qFOnXzGg, cSW1wSO3 :=  /*line _qJmMB5Y.go:1*/ fvIQVpVM(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line zRfa7Fhs.go:1*/ shim.Error( /*line hIGSD6Ii.go:1*/ cSW1wSO3.Error())
	}
	if w8EOizcx != qFOnXzGg {
		w5VhCLka =  /*line vZtcLyPy.go:1*/ func() string {
			var data []byte
			i := 2
			decryptKey := 121
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 11:
					data =  /*line vZtcLyPy.go:1*/ append(data, "D\x0fc"...,
					)
					i = 12
				case 4:
					for y := range data {
						data[y] = data[y] -  /*line vZtcLyPy.go:1*/ byte(decryptKey^y)
					}
					i = 6
				case 2:
					i = 0
					data =  /*line vZtcLyPy.go:1*/ append(data, "r\x18"...,
					)
				case 12:
					data =  /*line vZtcLyPy.go:1*/ append(data, 96)
					i = 1
				case 0:
					i = 7
					data =  /*line vZtcLyPy.go:1*/ append(data, 58)
				case 3:
					data =  /*line vZtcLyPy.go:1*/ append(data, "\\F"...,
					)
					i = 11
				case 13:
					i = 4
					data =  /*line vZtcLyPy.go:1*/ append(data, 8)
				case 5:
					data =  /*line vZtcLyPy.go:1*/ append(data, "ko\x1aZ"...,
					)
					i = 8
				case 1:
					data =  /*line vZtcLyPy.go:1*/ append(data, "Q]$\t"...,
					)
					i = 13
				case 10:
					i = 5
					data =  /*line vZtcLyPy.go:1*/ append(data, "\x129 K"...,
					)
				case 8:
					data =  /*line vZtcLyPy.go:1*/ append(data, "m["...,
					)
					i = 9
				case 9:
					data =  /*line vZtcLyPy.go:1*/ append(data, "NTVL"...,
					)
					i = 3
				case 7:
					i = 10
					data =  /*line vZtcLyPy.go:1*/ append(data, "feac"...,
					)
				}
			}
			return  /*line vZtcLyPy.go:1*/ string(data)
		}() + w8EOizcx +  /*line LiIvfvFj.go:1*/ func() string {
			key :=  /*line LiIvfvFj.go:1*/ []byte("_R\xd4\xcdVB\xd1N`\xb4-\x84\x15")
			data :=  /*line LiIvfvFj.go:1*/ []byte("\xc5\xc1F\xed\xba\xa3E\xaf\xd3\x19\xa1\xbe5")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line LiIvfvFj.go:1*/ string(data)
		}() + eviEA3sp +  /*line KTNTtWgr.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 156
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 1
					data =  /*line KTNTtWgr.go:1*/ append(data, 187)
				case 1:
					data =  /*line KTNTtWgr.go:1*/ append(data, 229)
					i = 2
				case 2:
					for y := range data {
						data[y] = data[y] ^  /*line KTNTtWgr.go:1*/ byte(decryptKey^y)
					}
					i = 0
				}
			}
			return  /*line KTNTtWgr.go:1*/ string(data)
		}()
		return  /*line hHE5akd4.go:1*/ shim.Error(w5VhCLka)
	}

	okjuZ7mN := eviEA3sp +  /*line ChSgfO5z.go:1*/ func() string {
		fullData :=  /*line ChSgfO5z.go:1*/ []byte("\xd4\xfcl\xef\x90\b;˸2z\xbd")
		var data []byte
		data =  /*line ChSgfO5z.go:1*/ append(data, fullData[4]^fullData[1], fullData[3]+fullData[10], fullData[9]+fullData[6], fullData[0]^fullData[11], fullData[5]+fullData[2], fullData[7]^fullData[8])
		return  /*line ChSgfO5z.go:1*/ string(data)
	}()
	b2ym12bm := DatasetLimits{mzjt9rLK, eviEA3sp, w8EOizcx, eBDS4r4c, tAfFZezC}
	u0tEk6z4, cSW1wSO3 :=  /*line nsZaABJB.go:1*/ json.Marshal(b2ym12bm)
	if cSW1wSO3 != nil {
		return  /*line km49czsW.go:1*/ shim.Error( /*line Ez3nAXcJ.go:1*/ cSW1wSO3.Error())
	}

	                        
	cSW1wSO3 =  /*line uXdUWq_H.go:1*/ n4c7mRtG.PutState(okjuZ7mN, u0tEk6z4)
	if cSW1wSO3 != nil {
		return  /*line BtteVQap.go:1*/ shim.Error( /*line zbjtrMi1.go:1*/ cSW1wSO3.Error())
	}

	xHTOVzVs :=  /*line MEVRh20d.go:1*/ func() string {
		seed :=  /*line MEVRh20d.go:1*/ byte(121)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line MEVRh20d.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/  /*line MEVRh20d.go:1*/ fnc(207)(25)(21)(239)(187)(66)(3)(0)(9)(178)(83)(242)(15)(172)(76)(253)(4)(252)(11)(255)(173)(70)(9)(3)(174)(68)(253)(19)(237)(18)(242)(15)(198)(230)
		return  /*line MEVRh20d.go:1*/ string(data)
	}() + eviEA3sp
	q8ZF8zaz :=  /*line EK8avee4.go:1*/ []byte(xHTOVzVs)
	return  /*line kdDcJkkh.go:1*/ shim.Success(q8ZF8zaz)

}

func (g4rnrSNn *G1Y_7pz_) bPYhnH_X(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	qVH0siY7 :=  /*line zAJCWiTO.go:1*/ func() string {
		key :=  /*line zAJCWiTO.go:1*/ []byte("\xcc8\x162\t\xbf\xc8\x13a\xc8!\x8bѝ\xd0@\xa4\xa8\x17\\\x9d\xab;E\xb3\xed1\xc9R M\xf9\x179\xb7È\x87A\xad")
		data :=  /*line zAJCWiTO.go:1*/ []byte("\xaf\xea]3c\xa6\x9ba\x0e\xaa\x01\xaf\xaa\x85\x94/\xbf\xacb\x14\xc8w\xffݑtC\x98!E'SR4\xb2\xb1\xeb\x9b<\xd0")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line zAJCWiTO.go:1*/ string(data)
	}()

	bxM55Mmo, cSW1wSO3 :=  /*line g861fLdG.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line Q9KrJF_y.go:1*/ shim.Error( /*line K3GlOrQj.go:1*/ cSW1wSO3.Error())
	}
	return  /*line GOlTE3aQ.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) g7MUqdhN(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	lWKzd4Pe, cSW1wSO3 :=  /*line ICxD_HIC.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line bHD20PHJ.go:1*/ shim.Error( /*line AEUGnN2z.go:1*/ cSW1wSO3.Error())
	}

	qVH0siY7 :=  /*line dbNReHZA.go:1*/ fmt.Sprintf( /*line gBNKEm_f.go:1*/ func() string {
		fullData :=  /*line gBNKEm_f.go:1*/ []byte("Y\xcaT\xd8\x0e\xaes|6\x1f\xc3!\xad\b\xe8V\xb2\xaf\u007f\xd8\xc9\xd8f\xa81G0=ڣ\xe8\xa02z\xfc\v\x82\x18\xe3\b\xfa]\xbc\xe6T\xe6\xces\xb5\x82ϱ|\x00\xf5\x8dt,\xf8\n\xde\xdec\xb2\xacb\xab;U|hE\v^\x0f4\xaadJ\x85en(\xab݇\x0e|\xf7\x13\x89f\xb4\xb4\xff\x920s\xac=\u05ff\xf0\xd2KRaS\xaf\xe7RF\u007f\x16\xe4\xe9\xfa\x0e; x*\xef\xbe\xf0\x84\x18p")
		var data []byte
		data =  /*line gBNKEm_f.go:1*/ append(data, fullData[17]-fullData[75], fullData[1]^fullData[14], fullData[84]^fullData[5], fullData[71]+fullData[119], fullData[111]^fullData[121], fullData[82]-fullData[10], fullData[16]+fullData[51], fullData[113]^fullData[65], fullData[25]-fullData[21], fullData[0]-fullData[109], fullData[63]+fullData[127], fullData[23]-fullData[81], fullData[47]-fullData[58], fullData[9]^fullData[27], fullData[45]^fullData[49], fullData[67]^fullData[2], fullData[114]^fullData[85], fullData[57]^fullData[120], fullData[18]+fullData[40], fullData[54]-fullData[79], fullData[55]^fullData[30], fullData[90]^fullData[83], fullData[110]^fullData[70], fullData[59]+fullData[37], fullData[102]^fullData[93], fullData[95]-fullData[24], fullData[100]-fullData[62], fullData[92]+fullData[12], fullData[66]^fullData[19], fullData[115]-fullData[125], fullData[91]+fullData[86], fullData[69]-fullData[96], fullData[31]^fullData[20], fullData[103]-fullData[80], fullData[15]+fullData[89], fullData[123]-fullData[78], fullData[11]^fullData[105], fullData[34]^fullData[61], fullData[118]-fullData[74], fullData[22]+fullData[42], fullData[43]-fullData[36], fullData[4]+fullData[107], fullData[3]^fullData[98], fullData[64]+fullData[48], fullData[56]+fullData[94], fullData[101]^fullData[28], fullData[87]^fullData[39], fullData[44]^fullData[72], fullData[38]-fullData[6], fullData[108]-fullData[99], fullData[77]^fullData[35], fullData[52]+fullData[116], fullData[106]+fullData[13], fullData[76]^fullData[46], fullData[68]-fullData[124], fullData[29]+fullData[50], fullData[73]^fullData[7], fullData[126]-fullData[60], fullData[112]-fullData[41], fullData[122]+fullData[8], fullData[97]+fullData[53], fullData[26]-fullData[117], fullData[104]+fullData[32], fullData[88]-fullData[33])
		return  /*line gBNKEm_f.go:1*/ string(data)
	}(), lWKzd4Pe)

	bxM55Mmo, cSW1wSO3 :=  /*line wDP56IOi.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line itqFXymJ.go:1*/ shim.Error( /*line XQFMIECL.go:1*/ cSW1wSO3.Error())
	}
	return  /*line TuY0_zb0.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) pgrPWzHl(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var cVKu8ZN2, w5VhCLka string
	var cSW1wSO3 error

	if  /*line BIXrN_yX.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line m5lHRsvw.go:1*/ shim.Error( /*line cZ8hgKup.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 161
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 20:
					i = 23
					data =  /*line cZ8hgKup.go:1*/ append(data, "r\xc1\xb1\xbe"...,
					)
				case 4:
					data =  /*line cZ8hgKup.go:1*/ append(data, 186)
					i = 22
				case 2:
					data =  /*line cZ8hgKup.go:1*/ append(data, "\x91\xd7\xe9"...,
					)
					i = 24
				case 23:
					i = 1
					data =  /*line cZ8hgKup.go:1*/ append(data, 187)
				case 3:
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xab\xd1\xc3"...,
					)
					i = 15
				case 6:
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xd6\xde\xdb\xd1"...,
					)
					i = 16
				case 21:
					i = 17
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xcc;\xd0"...,
					)
				case 11:
					i = 9
					data =  /*line cZ8hgKup.go:1*/ append(data, 106)
				case 10:
					i = 6
					data =  /*line cZ8hgKup.go:1*/ append(data, "ދ"...,
					)
				case 8:
					i = 18
					data =  /*line cZ8hgKup.go:1*/ append(data, "º\xb3"...,
					)
				case 18:
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xa5\xa4"...,
					)
					i = 4
				case 22:
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xb0\xb2"...,
					)
					i = 26
				case 17:
					i = 14
					data =  /*line cZ8hgKup.go:1*/ append(data, 216)
				case 0:
					i = 10
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xc9\xc8"...,
					)
				case 19:
					i = 20
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xaf±\xc1"...,
					)
				case 26:
					data =  /*line cZ8hgKup.go:1*/ append(data, 172)
					i = 11
				case 9:
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xaf\xa9\xbd"...,
					)
					i = 19
				case 13:
					i = 7
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xe6\xed\xf1"...,
					)
				case 16:
					data =  /*line cZ8hgKup.go:1*/ append(data, 209)
					i = 25
				case 14:
					i = 5
					for y := range data {
						data[y] = data[y] -  /*line cZ8hgKup.go:1*/ byte(decryptKey^y)
					}
				case 7:
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xad\x9c"...,
					)
					i = 8
				case 12:
					i = 21
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xc4z"...,
					)
				case 25:
					i = 2
					data =  /*line cZ8hgKup.go:1*/ append(data, "ߒ\xe2\xd6"...,
					)
				case 1:
					data =  /*line cZ8hgKup.go:1*/ append(data, "w\xc8"...,
					)
					i = 12
				case 15:
					i = 0
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xd0\xd8\xd9"...,
					)
				case 24:
					data =  /*line cZ8hgKup.go:1*/ append(data, "\xdb\xea\xe7\xe0"...,
					)
					i = 13
				}
			}
			return  /*line cZ8hgKup.go:1*/ string(data)
		}())
	}

	cVKu8ZN2 = ktsi1_SQ[0]
	okjuZ7mN := cVKu8ZN2 +  /*line jTyWb9nE.go:1*/ func() string {
		key :=  /*line jTyWb9nE.go:1*/ []byte("}w\x13\xd2\xc5J")
		data :=  /*line jTyWb9nE.go:1*/ []byte("\x11\x1e~\xbb\xb19")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line jTyWb9nE.go:1*/ string(data)
	}()
	mGZijNgp, cSW1wSO3 :=  /*line XXLkRpik.go:1*/ n4c7mRtG.GetState(okjuZ7mN)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line Bq0dvap0.go:1*/ func() string {
			seed :=  /*line Bq0dvap0.go:1*/ byte(110)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line Bq0dvap0.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/  /*line Bq0dvap0.go:1*/ fnc(233)(121)(21)(87)(174)(89)(181)(26)(76)(128)(36)(99)(206)(159)(55)(109)(150)(128)(251)(167)(149)(40)(95)(106)(39)(79)(139)(41)(67)(65)(200)(153)(53)(24)
			return  /*line Bq0dvap0.go:1*/ string(data)
		}() + okjuZ7mN +  /*line s6P2mOOh.go:1*/ func() string {
			fullData :=  /*line s6P2mOOh.go:1*/ []byte("=\x8a@\x98")
			var data []byte
			data =  /*line s6P2mOOh.go:1*/ append(data, fullData[1]+fullData[3], fullData[0]+fullData[2])
			return  /*line s6P2mOOh.go:1*/ string(data)
		}()
		return  /*line XLWWKpQw.go:1*/ shim.Error(w5VhCLka)
	} else if mGZijNgp == nil {
		w5VhCLka =  /*line AFE_Lhsm.go:1*/ func() string {
			data :=  /*line AFE_Lhsm.go:1*/ []byte(",!\xcbwr \x84\"\x8c\x9a\x87\xe4\x92arl4 eb\xabxt\x85\x99\xa4o#\x83\xb9\x8aoyEe\x9autm~f")
			positions := [...]byte{25, 5, 36, 8, 25, 2, 32, 28, 39, 30, 28, 29, 27, 38, 39, 21, 25, 39, 8, 33, 33, 29, 11, 9, 38, 24, 19, 10, 37, 16, 3, 2, 20, 3, 40, 8, 29, 9, 35, 38, 29, 0, 8, 37, 37, 39, 12, 0, 33, 1, 3, 3, 18, 15, 23, 14, 14, 6, 23, 12}
			for i := 0; i < 60; i += 2 {
				localKey :=  /*line AFE_Lhsm.go:1*/ byte(i) +  /*line AFE_Lhsm.go:1*/ byte(positions[i]^positions[i+1]) + 175
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line AFE_Lhsm.go:1*/ string(data)
		}() + okjuZ7mN +  /*line vJ5e9XVK.go:1*/ func() string {
			data :=  /*line vJ5e9XVK.go:1*/ []byte("\x81&")
			positions := [...]byte{0, 1}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line vJ5e9XVK.go:1*/ byte(i) +  /*line vJ5e9XVK.go:1*/ byte(positions[i]^positions[i+1]) + 3
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line vJ5e9XVK.go:1*/ string(data)
		}()
		return  /*line I_71ShkM.go:1*/ shim.Error(w5VhCLka)
	}

	return  /*line is_7d87w.go:1*/ shim.Success(mGZijNgp)
}

func (g4rnrSNn *G1Y_7pz_) ndkxu5Fu(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	if  /*line kKVMoelu.go:1*/ len(ktsi1_SQ) < 1 {
		return  /*line OdQBDyEz.go:1*/ shim.Error( /*line JSo7Bx8H.go:1*/ func() string {
			fullData :=  /*line JSo7Bx8H.go:1*/ []byte("\x02i\xbc\xbf\xd7F\xb4\xf3\x19d\xa6\x89H\x9cvm\x91\xf6w\x13YƢ;\x87\xe7\x92X$\x18ˊi@\xf2\xbd\xeah\xf2\x8c\x05K\x04\xea\xee\x9e<Ǽ\x15G\xad\xb5\xbd\xcb\xf5\xcb\f\xc1o\nǷ^TQ\x02g\x81{\xa4\xd0\xe8\xa07\x9fc\xcd \x10\xeaM\x94\x05\xfe\x1bF\x89\xe3\xd4\rdt\x0e|\x86?\xa7tdG\xa2\xa1<\xe1\x97\x1c\xa9\xce\xef\xd4`B1\xb6L\x97\xc9\xf2]\x8aN\xf1\xb1\"\x9e\x90(tP\xecWVO\xd6\xfd#\x8bO\xf2#\xa3o,\x14\x94\x0f\xf4\xe2뜿\\\xe9ZX\xd6\x17\x84\xf6\xb0ܰf\xdc\x00jdv\xa2\x10\x17")
			var data []byte
			data =  /*line JSo7Bx8H.go:1*/ append(data, fullData[46]-fullData[7], fullData[9]+fullData[60], fullData[70]^fullData[47], fullData[114]-fullData[50], fullData[15]+fullData[40], fullData[68]+fullData[122], fullData[32]^fullData[57], fullData[49]+fullData[121], fullData[62]+fullData[53], fullData[156]^fullData[17], fullData[117]^fullData[97], fullData[155]-fullData[88], fullData[8]^fullData[98], fullData[0]-fullData[73], fullData[23]-fullData[134], fullData[87]-fullData[157], fullData[99]+fullData[48], fullData[52]-fullData[5], fullData[159]^fullData[126], fullData[151]^fullData[75], fullData[95]^fullData[25], fullData[44]^fullData[150], fullData[1]-fullData[66], fullData[153]^fullData[13], fullData[161]-fullData[59], fullData[90]+fullData[27], fullData[116]+fullData[4], fullData[80]-fullData[14], fullData[104]+fullData[26], fullData[162]^fullData[45], fullData[20]+fullData[61], fullData[119]+fullData[72], fullData[120]^fullData[34], fullData[108]+fullData[22], fullData[11]+fullData[164], fullData[67]-fullData[42], fullData[128]+fullData[165], fullData[51]+fullData[2], fullData[168]^fullData[29], fullData[65]-fullData[43], fullData[18]-fullData[131], fullData[94]-fullData[41], fullData[107]-fullData[69], fullData[112]-fullData[124], fullData[92]+fullData[71], fullData[118]+fullData[142], fullData[129]+fullData[28], fullData[135]+fullData[167], fullData[100]-fullData[89], fullData[166]^fullData[146], fullData[105]-fullData[140], fullData[170]+fullData[79], fullData[127]+fullData[86], fullData[91]^fullData[83], fullData[10]^fullData[54], fullData[6]-fullData[138], fullData[145]+fullData[39], fullData[30]^fullData[169], fullData[81]^fullData[136], fullData[113]+fullData[109], fullData[103]-fullData[77], fullData[19]-fullData[102], fullData[132]+fullData[93], fullData[141]^fullData[21], fullData[163]^fullData[144], fullData[74]^fullData[171], fullData[84]-fullData[31], fullData[149]+fullData[158], fullData[35]+fullData[76], fullData[101]^fullData[58], fullData[115]+fullData[106], fullData[154]^fullData[96], fullData[133]-fullData[130], fullData[85]-fullData[160], fullData[110]^fullData[147], fullData[38]^fullData[82], fullData[56]-fullData[152], fullData[63]^fullData[143], fullData[111]^fullData[33], fullData[125]^fullData[139], fullData[148]^fullData[137], fullData[24]^fullData[36], fullData[123]-fullData[12], fullData[78]^fullData[64], fullData[37]-fullData[55], fullData[3]^fullData[16])
			return  /*line JSo7Bx8H.go:1*/ string(data)
		}())
	}

	eviEA3sp := ktsi1_SQ[0]
	pN03xyVt := eviEA3sp +  /*line JoUZQRVZ.go:1*/ func() string {
		key :=  /*line JoUZQRVZ.go:1*/ []byte("\\\a#GA]")
		data :=  /*line JoUZQRVZ.go:1*/ []byte("\xc8p\x90\xb0\xb5\xd0")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line JoUZQRVZ.go:1*/ string(data)
	}()

	wnS7ZxFz := DatasetLimits{}
	dLsP7EsR, cSW1wSO3 :=  /*line H7G2W5v7.go:1*/ n4c7mRtG.GetState(pN03xyVt)	                            
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line LYnzolyf.go:1*/ func() string {
			data :=  /*line LYnzolyf.go:1*/ []byte("\x14^EMrohR\x19pFaw#wu G\xad\x9b[\xf6h;EerLl.m\xb4t\x96?fo8AdataH\xdfumw")
			positions := [...]byte{37, 44, 27, 13, 26, 43, 37, 13, 46, 24, 23, 21, 20, 13, 6, 26, 31, 45, 47, 29, 6, 3, 46, 19, 1, 0, 27, 38, 6, 22, 8, 7, 15, 24, 1, 9, 1, 20, 14, 8, 34, 9, 8, 17, 33, 24, 12, 20, 45, 18, 33, 46}
			for i := 0; i < 52; i += 2 {
				localKey :=  /*line LYnzolyf.go:1*/ byte(i) +  /*line LYnzolyf.go:1*/ byte(positions[i]^positions[i+1]) + 202
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line LYnzolyf.go:1*/ string(data)
		}() + eviEA3sp +  /*line JWzzj0PC.go:1*/ func() string {
			data :=  /*line JWzzj0PC.go:1*/ []byte("M\xf2")
			positions := [...]byte{1, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line JWzzj0PC.go:1*/ byte(i) +  /*line JWzzj0PC.go:1*/ byte(positions[i]^positions[i+1]) + 47
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line JWzzj0PC.go:1*/ string(data)
		}()
		return  /*line sXpC_HRZ.go:1*/ shim.Error(w5VhCLka)
	} else if dLsP7EsR == nil {
		return  /*line JxsjzNzX.go:1*/ shim.Success(nil)
	}

	cSW1wSO3 =  /*line C7tfiyev.go:1*/ json.Unmarshal(dLsP7EsR, &wnS7ZxFz)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line E86cBpGh.go:1*/ func() string {
			var data []byte
			i := 5
			decryptKey := 125
			for counter := 0; i != 4; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 10:
					i = 9
					data =  /*line E86cBpGh.go:1*/ append(data, 226)
				case 12:
					i = 11
					data =  /*line E86cBpGh.go:1*/ append(data, "\xd2Ԍ"...,
					)
				case 0:
					data =  /*line E86cBpGh.go:1*/ append(data, "\xc8\xc3\xd0"...,
					)
					i = 12
				case 5:
					data =  /*line E86cBpGh.go:1*/ append(data, "\xf1\x99\xb9"...,
					)
					i = 8
				case 1:
					data =  /*line E86cBpGh.go:1*/ append(data, "\xc2\xde\xe3\xe7"...,
					)
					i = 3
				case 8:
					data =  /*line E86cBpGh.go:1*/ append(data, "\xe7\xe4\xe2"...,
					)
					i = 10
				case 13:
					data =  /*line E86cBpGh.go:1*/ append(data, "\xdbӅ\xc6"...,
					)
					i = 0
				case 6:
					i = 4
					for y := range data {
						data[y] = data[y] -  /*line E86cBpGh.go:1*/ byte(decryptKey^y)
					}
				case 7:
					i = 1
					data =  /*line E86cBpGh.go:1*/ append(data, 161)
				case 3:
					i = 13
					data =  /*line E86cBpGh.go:1*/ append(data, "\xdd݆"...,
					)
				case 11:
					data =  /*line E86cBpGh.go:1*/ append(data, "\xb7\xbd\xba"...,
					)
					i = 2
				case 9:
					i = 7
					data =  /*line E86cBpGh.go:1*/ append(data, "\x93\xb8"...,
					)
				case 2:
					i = 6
					data =  /*line E86cBpGh.go:1*/ append(data, 182)
				}
			}
			return  /*line E86cBpGh.go:1*/ string(data)
		}() + pN03xyVt +  /*line T9NxZ6oG.go:1*/ func() string {
			fullData :=  /*line T9NxZ6oG.go:1*/ []byte("\x99ډ\xa3")
			var data []byte
			data =  /*line T9NxZ6oG.go:1*/ append(data, fullData[0]+fullData[2], fullData[1]+fullData[3])
			return  /*line T9NxZ6oG.go:1*/ string(data)
		}()
		return  /*line FrWinlnq.go:1*/ shim.Error(w5VhCLka)
	}
	eBDS4r4c := wnS7ZxFz.DateLimit
	wzZlVbNo :=  /*line o0dd1UYv.go:1*/ strings.Split(eBDS4r4c,  /*line xwvlGhhr.go:1*/ func() string {
		data :=  /*line xwvlGhhr.go:1*/ []byte("h")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line xwvlGhhr.go:1*/ byte(i) +  /*line xwvlGhhr.go:1*/ byte(positions[i]^positions[i+1]) + 57
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line xwvlGhhr.go:1*/ string(data)
	}())
	c10w_gM6, cSW1wSO3 :=  /*line fKlrZXNO.go:1*/ strconv.Atoi(wzZlVbNo[0])
	if cSW1wSO3 != nil {
		 /*line C0IXLNGH.go:1*/ fmt.Println(cSW1wSO3)

	}
	uMPOamgi, cSW1wSO3 :=  /*line caPFlAQp.go:1*/ strconv.Atoi(wzZlVbNo[1])
	if cSW1wSO3 != nil {
		 /*line g7mL0pCu.go:1*/ fmt.Println(cSW1wSO3)

	}
	zsfPk1At, cSW1wSO3 :=  /*line Dp27nTMb.go:1*/ strconv.Atoi(wzZlVbNo[2])
	if cSW1wSO3 != nil {
		 /*line LqZaVwLt.go:1*/ fmt.Println(cSW1wSO3)

	}

	                  
	icCLZb33 :=  /*line j2Tgzn6q.go:1*/ time.Now()
	wC9xEX22 :=  /*line nOvzeR3q.go:1*/ icCLZb33.Format( /*line FlmzGACY.go:1*/ func() string {
		data :=  /*line FlmzGACY.go:1*/ []byte("1C0A/EM\xb1>\xac")
		positions := [...]byte{5, 3, 3, 8, 6, 8, 0, 3, 3, 1, 7, 9}
		for i := 0; i < 12; i += 2 {
			localKey :=  /*line FlmzGACY.go:1*/ byte(i) +  /*line FlmzGACY.go:1*/ byte(positions[i]^positions[i+1]) + 107
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line FlmzGACY.go:1*/ string(data)
	}())
	g2McJT39 :=  /*line EDiQiE3P.go:1*/ strings.Split(wC9xEX22,  /*line _zq71mI3.go:1*/ func() string {
		key :=  /*line _zq71mI3.go:1*/ []byte("\xf6")
		data :=  /*line _zq71mI3.go:1*/ []byte("9")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line _zq71mI3.go:1*/ string(data)
	}())
	hb2gZ8g6, cSW1wSO3 :=  /*line putDuKdx.go:1*/ strconv.Atoi(g2McJT39[0])
	if cSW1wSO3 != nil {
		 /*line uy9pPFrw.go:1*/ fmt.Println(cSW1wSO3)

	}
	lUmzLbXy, cSW1wSO3 :=  /*line CIvA5xAw.go:1*/ strconv.Atoi(g2McJT39[1])
	if cSW1wSO3 != nil {
		 /*line nnk9x0U6.go:1*/ fmt.Println(cSW1wSO3)

	}
	sDlsLiwx, cSW1wSO3 :=  /*line AyMLIclA.go:1*/ strconv.Atoi(g2McJT39[2])
	if cSW1wSO3 != nil {
		 /*line B9CNJEWz.go:1*/ fmt.Println(cSW1wSO3)

	}
	iehcvFck :=  /*line Dk0afGR6.go:1*/ QqEt7o0Z(c10w_gM6, uMPOamgi, zsfPk1At)

	cWzoB_u_ :=  /*line KUmRfZjQ.go:1*/ QqEt7o0Z(hb2gZ8g6, lUmzLbXy, sDlsLiwx)
	r3WdRUIM :=  /*line j5zO7pP2.go:1*/  /*line j5zO7pP2.go:1*/ cWzoB_u_.Sub(iehcvFck).Hours() / 24
	if r3WdRUIM <= 0 {
		w5VhCLka =  /*line jZwc8tCl.go:1*/ func() string {
			key :=  /*line jZwc8tCl.go:1*/ []byte("\x83bI\xe8\"O\xdfG\x17m\x80\xbf\xd5n\xc8ӆ3\xd3\x18u")
			data :=  /*line jZwc8tCl.go:1*/ []byte("\xf8@\f\x9aP \xade-O\xd4װN\xac\xb2\xf2R\xa0}\x01")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line jZwc8tCl.go:1*/ string(data)
		}() + eviEA3sp +  /*line iCN7MMU6.go:1*/ func() string {
			data :=  /*line iCN7MMU6.go:1*/ []byte("\xd0t\x8fn\fh a\xa4ai\xb1\xfbȤe\xa7yet")
			positions := [...]byte{14, 16, 2, 2, 0, 0, 4, 13, 4, 16, 5, 11, 0, 11, 4, 11, 16, 1, 12, 16, 8, 13}
			for i := 0; i < 22; i += 2 {
				localKey :=  /*line iCN7MMU6.go:1*/ byte(i) +  /*line iCN7MMU6.go:1*/ byte(positions[i]^positions[i+1]) + 173
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line iCN7MMU6.go:1*/ string(data)
		}() +  /*line GfLpQiqY.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 163
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 0:
					i = 1
					data =  /*line GfLpQiqY.go:1*/ append(data, 221)
				case 3:
					data =  /*line GfLpQiqY.go:1*/ append(data, 129)
					i = 0
				case 1:
					i = 2
					for y := range data {
						data[y] = data[y] +  /*line GfLpQiqY.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return  /*line GfLpQiqY.go:1*/ string(data)
		}()
		return  /*line lixSPthZ.go:1*/ shim.Error(w5VhCLka)
	}
	return  /*line ENqWiA_f.go:1*/ shim.Success(nil)
}

func QqEt7o0Z(c10w_gM6, zsfPk1At, uMPOamgi int) time.Time {
	return  /*line z_OT9ns5.go:1*/ time.Date(c10w_gM6,  /*line GkGJOU0P.go:1*/ time.Month(zsfPk1At), uMPOamgi, 0, 0, 0, 0, time.UTC)

}
