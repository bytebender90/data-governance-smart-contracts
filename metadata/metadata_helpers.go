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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	pb "github.com/hyperledger/fabric/protos/peer"
)

                                                    
func (kuuXUPaZ *JPyzqDxq) iiTHmzLy(vhZhSyYv shim.ChaincodeStubInterface) (string, error) {
	jVqGRYFt, r0gMkWZ6, xonUbr_n :=  /*line cUHf90Kr.go:1*/ cid.GetAttributeValue(vhZhSyYv,  /*line lutlWO1z.go:1*/ func() string {
		seed :=  /*line lutlWO1z.go:1*/ byte(98)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line lutlWO1z.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/  /*line lutlWO1z.go:1*/ fnc(203)(160)(18)(87)(170)(69)(151)(42)(71)(153)(27)(84)(154)(65)
		return  /*line lutlWO1z.go:1*/ string(data)
	}())

	if xonUbr_n != nil {
		return "", xonUbr_n
	}

	if !r0gMkWZ6 {
		return "",  /*line kFW4K92i.go:1*/ errors.New( /*line U0oI0snU.go:1*/ func() string {
			seed :=  /*line U0oI0snU.go:1*/ byte(34)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line U0oI0snU.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/  /*line U0oI0snU.go:1*/ fnc(139)(32)(18)(87)(170)(69)(151)(42)(71)(153)(27)(84)(154)(65)(48)(161)(85)(170)(82)(155)(47)(113)(225)(179)(33)(139)(32)(237)(39)(74)(158)(60)(110)(225)(187)
			return  /*line U0oI0snU.go:1*/ string(data)
		}())
	}

	return jVqGRYFt, nil
}

                                                   
func (kuuXUPaZ *JPyzqDxq) a8t6_82N(vhZhSyYv shim.ChaincodeStubInterface) (string, error) {
	kFtED_RC, r0gMkWZ6, xonUbr_n :=  /*line kKQy9ZZR.go:1*/ cid.GetAttributeValue(vhZhSyYv,  /*line jtY6QQIb.go:1*/ func() string {
		data :=  /*line jtY6QQIb.go:1*/ []byte("WBkPrHTIOP5")
		positions := [...]byte{2, 1, 6, 3, 7, 0, 5, 10, 9, 5, 0, 2}
		for i := 0; i < 12; i += 2 {
			localKey :=  /*line jtY6QQIb.go:1*/ byte(i) +  /*line jtY6QQIb.go:1*/ byte(positions[i]^positions[i+1]) + 10
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line jtY6QQIb.go:1*/ string(data)
	}())

	if xonUbr_n != nil {
		return "", xonUbr_n
	}

	if !r0gMkWZ6 {
		return "",  /*line HMEnUD5M.go:1*/ errors.New( /*line pt0UyrQ7.go:1*/ func() string {
			data :=  /*line pt0UyrQ7.go:1*/ []byte("\xedx\xeae\x9fHů\xfdr\nv\x8a\xda\tr\x03\xaa״.\xde5\xa1\xb2'i\xfe\xb2i\x9d\x04")
			positions := [...]byte{14, 30, 25, 25, 23, 30, 22, 5, 24, 5, 20, 28, 6, 4, 24, 7, 22, 23, 4, 5, 10, 27, 12, 28, 17, 11, 8, 18, 16, 22, 28, 2, 19, 31, 31, 12, 20, 23, 12, 21, 13, 17, 0, 6, 14, 28, 30, 14}
			for i := 0; i < 48; i += 2 {
				localKey :=  /*line pt0UyrQ7.go:1*/ byte(i) +  /*line pt0UyrQ7.go:1*/ byte(positions[i]^positions[i+1]) + 68
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line pt0UyrQ7.go:1*/ string(data)
		}())
	}

	return kFtED_RC, nil
}

func (kuuXUPaZ *JPyzqDxq) yq7QUwKQ(vhZhSyYv shim.ChaincodeStubInterface) (string, error) {
	tu6TYtHu, r0gMkWZ6, xonUbr_n :=  /*line UOJo2EPg.go:1*/ cid.GetAttributeValue(vhZhSyYv,  /*line TNgwCpuT.go:1*/ func() string {
		fullData :=  /*line TNgwCpuT.go:1*/ []byte("\xf91\x89\xac\xb5\x9d}k\xb8\x94\t0\xa2C-\x9f\xa6\xd7]>\xe6\xa3ͻ!\xd5")
		var data []byte
		data =  /*line TNgwCpuT.go:1*/ append(data, fullData[7]+fullData[0], fullData[3]^fullData[22], fullData[15]+fullData[25], fullData[23]+fullData[16], fullData[14]+fullData[13], fullData[8]^fullData[17], fullData[12]-fullData[11], fullData[6]^fullData[10], fullData[9]-fullData[24], fullData[5]+fullData[4], fullData[20]+fullData[2], fullData[18]^fullData[1], fullData[21]-fullData[19])
		return  /*line TNgwCpuT.go:1*/ string(data)
	}())

	if xonUbr_n != nil {
		return "", xonUbr_n
	}

	if !r0gMkWZ6 {
		return "",  /*line B_tpMFOe.go:1*/ errors.New( /*line eltxc5a7.go:1*/ func() string {
			data :=  /*line eltxc5a7.go:1*/ []byte("l\xfaD\xf7 a\xf1t\xf1\xc0\xf1m\xfb`\xf1\xdcZ\xf4m\xbdOX\"n\xe9")
			positions := [...]byte{22, 13, 0, 6, 6, 2, 8, 12, 22, 20, 15, 1, 21, 8, 22, 3, 24, 8, 11, 16, 11, 3, 14, 13, 10, 1, 11, 16, 21, 9, 6, 10, 17, 19}
			for i := 0; i < 34; i += 2 {
				localKey :=  /*line eltxc5a7.go:1*/ byte(i) +  /*line eltxc5a7.go:1*/ byte(positions[i]^positions[i+1]) + 123
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line eltxc5a7.go:1*/ string(data)
		}())
	}

	return tu6TYtHu, nil
}

func (kuuXUPaZ *JPyzqDxq) glBEweME(vhZhSyYv shim.ChaincodeStubInterface) (string, error) {
	taAS4TM5, r0gMkWZ6, xonUbr_n :=  /*line hBLdhj3j.go:1*/ cid.GetAttributeValue(vhZhSyYv,  /*line FUxfEaBE.go:1*/ func() string {
		seed :=  /*line FUxfEaBE.go:1*/ byte(171)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line FUxfEaBE.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line FUxfEaBE.go:1*/  /*line FUxfEaBE.go:1*/  /*line FUxfEaBE.go:1*/  /*line FUxfEaBE.go:1*/  /*line FUxfEaBE.go:1*/  /*line FUxfEaBE.go:1*/  /*line FUxfEaBE.go:1*/  /*line FUxfEaBE.go:1*/ fnc(202)(254)(242)(13)(252)(243)(12)(248)
		return  /*line FUxfEaBE.go:1*/ string(data)
	}())

	if xonUbr_n != nil {
		return "", xonUbr_n
	}

	if !r0gMkWZ6 {
		return "",  /*line tbFgCeuU.go:1*/ errors.New( /*line mzhCDtTz.go:1*/ func() string {
			seed :=  /*line mzhCDtTz.go:1*/ byte(232)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line mzhCDtTz.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/  /*line mzhCDtTz.go:1*/ fnc(93)(184)(98)(209)(158)(47)(106)(204)(83)(231)(225)(194)(130)(251)(239)(241)(225)(179)(33)(139)(32)(237)(39)(74)(158)(60)(110)(225)(187)
			return  /*line mzhCDtTz.go:1*/ string(data)
		}())
	}

	return taAS4TM5, nil
}

func (kuuXUPaZ *JPyzqDxq) mUd4z6cw(vhZhSyYv shim.ChaincodeStubInterface, y5yvGZQA []string) pb.Response {
	var xonUbr_n error
	bXQWJgv9 := y5yvGZQA[0]
	pzKghUuG := y5yvGZQA[1]

	                                         
	j9GLBFOm, xonUbr_n :=  /*line FxhTBoKo.go:1*/ vhZhSyYv.GetState(pzKghUuG)
	if xonUbr_n != nil {
		return  /*line l2a7hqoJ.go:1*/ shim.Error( /*line H8XEQNum.go:1*/ func() string {
			key :=  /*line H8XEQNum.go:1*/ []byte("!\xb8\xb2\x91\f\x11\xf1\xbfɢ<\xad\x96\xfa\xff\x81`Oem\x18~c\xb7\f\xf9\x8f\b\xe6I\xceQk\xc0\x19\x98\x83\x8asB\x06?\xad9k)\x10\xd3")
			data :=  /*line H8XEQNum.go:1*/ []byte("g\xd9\xdb\xfdiu\xd1˦\x82_\xc5\U000d9521\t)E\x19p\x1bC\xd9m\x94\xea(\x89/\xee<\x0e\xb4x\xfc\xe2\xfe\x12bcG\xc4J\x1fZ*\xf3")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line H8XEQNum.go:1*/ string(data)
		}() +  /*line rbg2S6xi.go:1*/ xonUbr_n.Error())
	} else if j9GLBFOm == nil {
		return  /*line tGY_75lo.go:1*/ shim.Error( /*line Eud8zdXP.go:1*/ func() string {
			fullData :=  /*line Eud8zdXP.go:1*/ []byte("\x91\xae\xa1\xa3\f&\xfcw\xd0'\xe6LQŴ\xd1\x1c\x84\xed}[>\xdd\xee\xc73b\x8a\x802\xc1qI\x9b\x16\x17\x8d\xffL\xee*\x94\x19\xcc\xcfT\xdfjSQ8\xfd\xc0rX\x81\xee\xc2ִ\xfc\xa4k\x8b(\x8eo;#\x18\xdf\b")
			var data []byte
			data =  /*line Eud8zdXP.go:1*/ append(data, fullData[54]^fullData[16], fullData[43]+fullData[3], fullData[44]-fullData[47], fullData[11]+fullData[9], fullData[25]+fullData[18], fullData[17]-fullData[34], fullData[39]^fullData[55], fullData[42]+fullData[20], fullData[22]^fullData[51], fullData[19]-fullData[69], fullData[35]^fullData[66], fullData[61]+fullData[13], fullData[59]^fullData[24], fullData[26]-fullData[23], fullData[7]+fullData[6], fullData[64]^fullData[71], fullData[70]-fullData[53], fullData[41]+fullData[15], fullData[50]^fullData[38], fullData[60]-fullData[33], fullData[5]-fullData[57], fullData[63]+fullData[58], fullData[28]-fullData[4], fullData[14]-fullData[48], fullData[30]-fullData[2], fullData[52]-fullData[32], fullData[46]+fullData[27], fullData[68]+fullData[12], fullData[40]+fullData[21], fullData[49]^fullData[31], fullData[37]^fullData[0], fullData[56]-fullData[36], fullData[67]+fullData[29], fullData[8]-fullData[62], fullData[45]+fullData[10], fullData[65]^fullData[1])
			return  /*line Eud8zdXP.go:1*/ string(data)
		}() + pzKghUuG)
	}
	ls5l3RtF := DatasetMetadata{}
	xonUbr_n =  /*line So4xm2D5.go:1*/ json.Unmarshal(j9GLBFOm, &ls5l3RtF)	                               
	if xonUbr_n != nil {
		return  /*line PU94vWaG.go:1*/ shim.Error( /*line k7J7q7Qo.go:1*/ xonUbr_n.Error())
	}

	qKQytQy6, xonUbr_n :=  /*line eOMAXENM.go:1*/ kuuXUPaZ.glBEweME(vhZhSyYv)
	if xonUbr_n != nil {
		return  /*line YgFpmKFv.go:1*/ shim.Error( /*line VUrf_pJr.go:1*/ xonUbr_n.Error())
	}

	if ls5l3RtF.UsernameOfProvider != qKQytQy6 {
		return  /*line _DR5GNjK.go:1*/ shim.Error( /*line xPZVl0Lz.go:1*/ func() string {
			key :=  /*line xPZVl0Lz.go:1*/ []byte("'\xa4\xce\x14]\xea\xf80ݳf:x\xf5\x10\x12\x8f\x80\xa3\xb6\xd9%\x83\x10:\xecq'{F\xb7{K\xaa\xafUw\x0e\x10D1\x85\x06:k")
			data :=  /*line xPZVl0Lz.go:1*/ []byte("hʢm}\x9e\x90U\xfd\xd7\aN\x19\xd5``\xe0\xf6\xcaҼW\xa3s[\x82QR\v\"\xd6\x0f.\x8a\xdb=\x12.t%E\xe4u_\x1f")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line xPZVl0Lz.go:1*/ string(data)
		}())
	}

	ls5l3RtF.EmailOfProvider = bXQWJgv9

	Qk2Ir6Pj, _ :=  /*line UzX0tatU.go:1*/ json.Marshal(ls5l3RtF)
	xonUbr_n =  /*line BbwUkoyP.go:1*/ vhZhSyYv.PutState(pzKghUuG, Qk2Ir6Pj)
	if xonUbr_n != nil {
		return  /*line q7bWPv1y.go:1*/ shim.Error( /*line za8eClJE.go:1*/ xonUbr_n.Error())
	}

	return  /*line I8gunRyT.go:1*/ shim.Success(nil)

}

func _IAEAzGw(vhZhSyYv shim.ChaincodeStubInterface, aarIVocK string) ([]byte, error) {

	 /*line NG8NniO9.go:1*/ fmt.Printf( /*line TFJ77Doq.go:1*/ func() string {
		seed :=  /*line TFJ77Doq.go:1*/ byte(139)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line TFJ77Doq.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/  /*line TFJ77Doq.go:1*/ fnc(166)(17)(37)(2)(29)(215)(40)(224)(23)(5)(211)(49)(246)(14)(229)(26)(206)(57)(253)(221)(28)(224)(23)(5)(210)(39)(8)(235)(3)(23)(167)(95)(248)(224)(23)(5)(210)(39)(8)(235)(3)(23)(189)(78)(183)(58)(137)
		return  /*line TFJ77Doq.go:1*/ string(data)
	}(), aarIVocK)

	pS4pXzLH, xonUbr_n :=  /*line _oaYAC6u.go:1*/ vhZhSyYv.GetQueryResult(aarIVocK)
	if xonUbr_n != nil {
		return nil, xonUbr_n
	}
	defer  /*line UzRVEXAA.go:1*/ pS4pXzLH.Close()

	h0MRRAn3, xonUbr_n :=  /*line lKhkxxY9.go:1*/ j4_oP90A(pS4pXzLH)
	if xonUbr_n != nil {
		return nil, xonUbr_n
	}

	 /*line D7j2J4VR.go:1*/ fmt.Printf( /*line N0Ao_qRi.go:1*/ func() string {
		data :=  /*line N0Ao_qRi.go:1*/ []byte("-ag+<QZ$QOQeNKlMFQrg|iryS7rih\\dq\x04/>.-H\xd9u\\t:]9si")
		positions := [...]byte{46, 21, 38, 38, 15, 44, 30, 3, 34, 8, 34, 13, 3, 21, 7, 1, 36, 44, 32, 3, 9, 19, 33, 38, 10, 34, 8, 17, 4, 46, 30, 34, 28, 12, 15, 37, 38, 37, 40, 38, 20, 40, 25, 33, 43, 35, 12, 29, 38, 30, 40, 3, 13, 6, 17, 37}
		for i := 0; i < 56; i += 2 {
			localKey :=  /*line N0Ao_qRi.go:1*/ byte(i) +  /*line N0Ao_qRi.go:1*/ byte(positions[i]^positions[i+1]) + 240
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line N0Ao_qRi.go:1*/ string(data)
	}(),  /*line MwcfEJd8.go:1*/ h0MRRAn3.String())

	return  /*line MpN9NJtV.go:1*/ h0MRRAn3.Bytes(), nil
}

func j4_oP90A(pS4pXzLH shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	                                                 
	var h0MRRAn3 bytes.Buffer
	 /*line tlvv1tQc.go:1*/ h0MRRAn3.WriteString( /*line tZxcFGoF.go:1*/ func() string {
		fullData :=  /*line tZxcFGoF.go:1*/ []byte("ʑ")
		var data []byte
		data =  /*line tZxcFGoF.go:1*/ append(data, fullData[1]^fullData[0])
		return  /*line tZxcFGoF.go:1*/ string(data)
	}())

	taHKR7Ox := false
	for  /*line AWK2ZkjM.go:1*/ pS4pXzLH.HasNext() {
		xzkLgU7l, xonUbr_n :=  /*line BMHHlqQc.go:1*/ pS4pXzLH.Next()
		if xonUbr_n != nil {
			return nil, xonUbr_n
		}
		                                                                           
		if taHKR7Ox {
			 /*line Ic2BL5C0.go:1*/ h0MRRAn3.WriteString( /*line p5IfOvxa.go:1*/ func() string {
				var data []byte
				i := 0
				decryptKey := 170
				for counter := 0; i != 1; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 2:
						for y := range data {
							data[y] = data[y] -  /*line p5IfOvxa.go:1*/ byte(decryptKey^y)
						}
						i = 1
					case 0:
						i = 2
						data =  /*line p5IfOvxa.go:1*/ append(data, 212)
					}
				}
				return  /*line p5IfOvxa.go:1*/ string(data)
			}())
		}
		 /*line nFObOzIM.go:1*/ h0MRRAn3.WriteString( /*line UxHebZJO.go:1*/ func() string {
			data :=  /*line UxHebZJO.go:1*/ []byte("\xa0]K\x1ey\xf1:")
			positions := [...]byte{0, 5, 3, 0, 1, 0, 3, 0}
			for i := 0; i < 8; i += 2 {
				localKey :=  /*line UxHebZJO.go:1*/ byte(i) +  /*line UxHebZJO.go:1*/ byte(positions[i]^positions[i+1]) + 125
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line UxHebZJO.go:1*/ string(data)
		}())
		 /*line XtOIBJRX.go:1*/ h0MRRAn3.WriteString( /*line osq195lG.go:1*/ func() string {
			key :=  /*line osq195lG.go:1*/ []byte("N")
			data :=  /*line osq195lG.go:1*/ []byte("l")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line osq195lG.go:1*/ string(data)
		}())
		 /*line niw110rd.go:1*/ h0MRRAn3.WriteString(xzkLgU7l.Key)
		 /*line tagkjrLm.go:1*/ h0MRRAn3.WriteString( /*line gj8Y4Y2p.go:1*/ func() string {
			key :=  /*line gj8Y4Y2p.go:1*/ []byte("!")
			data :=  /*line gj8Y4Y2p.go:1*/ []byte("\x01")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line gj8Y4Y2p.go:1*/ string(data)
		}())

		 /*line oMSG91as.go:1*/ h0MRRAn3.WriteString( /*line jvo5mGia.go:1*/ func() string {
			key :=  /*line jvo5mGia.go:1*/ []byte("\x87\xfca\x88\xb0\xd7<\xf4\x1b\xc1%")
			data :=  /*line jvo5mGia.go:1*/ []byte("\xb3\x1c\x83\xda\x15:\xabf\u007f\xe3_")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line jvo5mGia.go:1*/ string(data)
		}())
		                                             
		 /*line ntsihHlw.go:1*/ h0MRRAn3.WriteString( /*line QXwl6U1r.go:1*/ string(xzkLgU7l.Value))
		 /*line ncnff5ZT.go:1*/ h0MRRAn3.WriteString( /*line Q8wBFKMx.go:1*/ func() string {
			fullData :=  /*line Q8wBFKMx.go:1*/ []byte("\xb1\xcc")
			var data []byte
			data =  /*line Q8wBFKMx.go:1*/ append(data, fullData[0]+fullData[1])
			return  /*line Q8wBFKMx.go:1*/ string(data)
		}())
		taHKR7Ox = true
	}
	 /*line KOwwlfEO.go:1*/ h0MRRAn3.WriteString( /*line P_g0MAaO.go:1*/ func() string {
		key :=  /*line P_g0MAaO.go:1*/ []byte("\x9e")
		data :=  /*line P_g0MAaO.go:1*/ []byte("\xc3")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line P_g0MAaO.go:1*/ string(data)
	}())

	return &h0MRRAn3, nil
}
