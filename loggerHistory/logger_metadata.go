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

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (hslgON00 *EMhVAN9S) nnVEQSHY(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line oKl4Yzwq.go:1*/ len(eFz_TaL1) != 7 {
		return  /*line JGhj7qsp.go:1*/ shim.Error( /*line F8Q5J3pt.go:1*/ func() string {
			var data []byte
			i := 16
			decryptKey := 207
			for counter := 0; i != 3; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 15:
					i = 13
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\xcbz"...,
					)
				case 16:
					i = 11
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\x8c\xb2\xa4\xb1"...,
					)
				case 18:
					data =  /*line F8Q5J3pt.go:1*/ append(data, 186)
					i = 12
				case 12:
					i = 2
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\xaa\xa9\xbfl"...,
					)
				case 17:
					i = 15
					data =  /*line F8Q5J3pt.go:1*/ append(data, "Ľ\xc2\xc1"...,
					)
				case 1:
					for y := range data {
						data[y] = data[y] +  /*line F8Q5J3pt.go:1*/ byte(decryptKey^y)
					}
					i = 3
				case 10:
					i = 8
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\xbc\xb6ź"...,
					)
				case 7:
					i = 1
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\xdb\xe2\xe6"...,
					)
				case 9:
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\xc8Ā"...,
					)
					i = 4
				case 0:
					i = 19
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\xd8\xca\xd0\xce"...,
					)
				case 8:
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\xb3\xc1"...,
					)
					i = 9
				case 14:
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\xd8\xcd\xc3\xc6"...,
					)
					i = 0
				case 2:
					data =  /*line F8Q5J3pt.go:1*/ append(data, 170)
					i = 10
				case 6:
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\xdf\xdc\xd5"...,
					)
					i = 7
				case 11:
					i = 18
					data =  /*line F8Q5J3pt.go:1*/ append(data, 185)
				case 5:
					i = 6
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\x86\xcc\xde\xd0"...,
					)
				case 19:
					i = 5
					data =  /*line F8Q5J3pt.go:1*/ append(data, "\x88\x9c"...,
					)
				case 13:
					i = 14
					data =  /*line F8Q5J3pt.go:1*/ append(data, 196)
				case 4:
					data =  /*line F8Q5J3pt.go:1*/ append(data, "w\xa4"...,
					)
					i = 17
				}
			}
			return  /*line F8Q5J3pt.go:1*/ string(data)
		}())
	}
	cXmPo8Kd := eFz_TaL1[0]
	c8qKfNsz := eFz_TaL1[1]
	eBmFb7v9 := eFz_TaL1[2]
	izyfYrYF := []string{eFz_TaL1[3]}
	z63NukZ2 := eFz_TaL1[4]
	fGOn2IzA := eFz_TaL1[5]
	clXFz22n := eFz_TaL1[6]
	_0uGwnuQ :=  /*line U7o7sO2U.go:1*/ func() string {
		seed :=  /*line U7o7sO2U.go:1*/ byte(6)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line U7o7sO2U.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/  /*line U7o7sO2U.go:1*/ fnc(82)(199)(134)(12)(22)(57)(77)(178)(115)(211)(169)(79)(177)(79)
		return  /*line U7o7sO2U.go:1*/ string(data)
	}()

	if  /*line w2WhvWdE.go:1*/ len(cXmPo8Kd) <= 0 {
		return  /*line cesWahKm.go:1*/ shim.Error( /*line sTRSPbhx.go:1*/ func() string {
			fullData :=  /*line sTRSPbhx.go:1*/ []byte("1\x81\x00\x97\xf5\xbc\x05\x92{<S\x82\x98\xc5\x1a\xf3\xba\b\xdf.\v3\xae\xec\xe7\xac͉\x96\xc6\xd3\xf7T\xdbC\xfd\xf1\xe9kA9\xbe\xa03k\x8aH6ɵ\xf0\"\xdbV\x94/\xc5D \x02\xc8R\xbesbY8=")
			var data []byte
			data =  /*line sTRSPbhx.go:1*/ append(data, fullData[50]+fullData[32], fullData[25]+fullData[49], fullData[18]-fullData[44], fullData[33]^fullData[16], fullData[41]^fullData[26], fullData[42]+fullData[56], fullData[29]-fullData[61], fullData[45]+fullData[28], fullData[14]+fullData[10], fullData[67]-fullData[60], fullData[8]^fullData[17], fullData[46]^fullData[9], fullData[48]^fullData[37], fullData[65]-fullData[31], fullData[11]^fullData[24], fullData[51]^fullData[59], fullData[19]+fullData[21], fullData[27]+fullData[3], fullData[38]-fullData[35], fullData[40]+fullData[47], fullData[63]-fullData[6], fullData[52]-fullData[22], fullData[36]^fullData[54], fullData[62]^fullData[30], fullData[43]^fullData[34], fullData[15]+fullData[1], fullData[13]^fullData[5], fullData[58]^fullData[2], fullData[57]+fullData[55], fullData[23]^fullData[12], fullData[0]+fullData[39], fullData[20]^fullData[64], fullData[53]^fullData[66], fullData[4]^fullData[7])
			return  /*line sTRSPbhx.go:1*/ string(data)
		}())
	}

	if  /*line BdXn5LRF.go:1*/ len(eBmFb7v9) <= 0 {
		return  /*line icKUg06E.go:1*/ shim.Error( /*line zQCYUwpr.go:1*/ func() string {
			seed :=  /*line zQCYUwpr.go:1*/ byte(233)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line zQCYUwpr.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/  /*line zQCYUwpr.go:1*/ fnc(166)(253)(235)(22)(227)(25)(243)(29)(237)(239)(26)(225)(80)(173)(24)(246)(15)(170)(86)(239)(89)(179)(165)(68)(1)(1)(93)(168)(24)(253)(254)(241)(89)(161)(7)(8)(235)(3)(23)
			return  /*line zQCYUwpr.go:1*/ string(data)
		}())
	}

	if  /*line UuXHdLNl.go:1*/ len(c8qKfNsz) <= 0 {
		return  /*line gh2ASYfE.go:1*/ shim.Error( /*line z_TxGUOv.go:1*/ func() string {
			data :=  /*line z_TxGUOv.go:1*/ []byte("+s\xf9\xdd\x13\xf7\xf1\x1b\xfbmust\xbbbe\xb4\xe8 \x17on+\xfc:pt\x198s\x05\xddi\xc8\xc2")
			positions := [...]byte{13, 30, 23, 28, 17, 17, 33, 34, 34, 16, 8, 6, 22, 5, 19, 33, 2, 13, 24, 19, 23, 34, 16, 30, 6, 31, 34, 17, 0, 7, 3, 13, 6, 27, 17, 4}
			for i := 0; i < 36; i += 2 {
				localKey :=  /*line z_TxGUOv.go:1*/ byte(i) +  /*line z_TxGUOv.go:1*/ byte(positions[i]^positions[i+1]) + 23
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line z_TxGUOv.go:1*/ string(data)
		}())
	}

	VzKzn1aa := &LoggerMetadata{_0uGwnuQ, cXmPo8Kd, c8qKfNsz, eBmFb7v9, izyfYrYF, fGOn2IzA, clXFz22n}
	EOAX9wsO, xREbPiGc :=  /*line J6kYYE2t.go:1*/ json.Marshal(VzKzn1aa)
	if xREbPiGc != nil {
		return  /*line UWZTliNh.go:1*/ shim.Error( /*line ZDn4ccj0.go:1*/ xREbPiGc.Error())
	}

	aeCTnc39 := cXmPo8Kd + z63NukZ2

	                        
	xREbPiGc =  /*line OzYvdNTo.go:1*/ fOOTaXmK.PutState(aeCTnc39, EOAX9wsO)
	if xREbPiGc != nil {
		return  /*line YW7oIRdJ.go:1*/ shim.Error( /*line BHdc_qyw.go:1*/ xREbPiGc.Error())
	}

	auAEXfA1 :=  /*line CQOZoi07.go:1*/ func() string {
		fullData :=  /*line CQOZoi07.go:1*/ []byte("\x9ap\xf7|\xfa\x85\xf0\xa7R\x820O\x17'§G\xa9\x8f\x81B?\x9b:\x06\x92T\xa9ρ1\xa2\x1a!P\xa2\a=\x80\t\x9b\xe0\x10\xee\x97k\xb9\x06\xfc\xffKM\xb89\x89\x13\x9d\xa7\xc36\x86W\t\x15\xe5\xf3\xca\xe1s,\xfa\x11{`")
		var data []byte
		data =  /*line CQOZoi07.go:1*/ append(data, fullData[26]^fullData[12], fullData[63]-fullData[17], fullData[45]+fullData[4], fullData[71]+fullData[34], fullData[72]-fullData[62], fullData[61]-fullData[43], fullData[57]-fullData[53], fullData[56]+fullData[66], fullData[23]^fullData[32], fullData[21]-fullData[2], fullData[47]-fullData[44], fullData[3]-fullData[36], fullData[8]^fullData[33], fullData[7]-fullData[20], fullData[35]^fullData[9], fullData[68]-fullData[24], fullData[31]-fullData[37], fullData[58]-fullData[11], fullData[29]+fullData[41], fullData[0]-fullData[59], fullData[67]^fullData[38], fullData[70]-fullData[60], fullData[69]^fullData[51], fullData[10]^fullData[42], fullData[50]-fullData[64], fullData[13]-fullData[52], fullData[46]-fullData[16], fullData[54]^fullData[27], fullData[49]^fullData[22], fullData[25]+fullData[28], fullData[1]-fullData[48], fullData[6]-fullData[18], fullData[55]^fullData[73], fullData[14]^fullData[15], fullData[65]+fullData[19], fullData[39]+fullData[30], fullData[5]+fullData[40])
		return  /*line CQOZoi07.go:1*/ string(data)
	}() + cXmPo8Kd
	pPXX7o_t :=  /*line BXaJFEMv.go:1*/ []byte(auAEXfA1)
	return  /*line QTbTIC05.go:1*/ shim.Success(pPXX7o_t)

}

func (hslgON00 *EMhVAN9S) pHnSOSDM(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {
	                  
	fcEHFI_2 :=  /*line TD4btZdy.go:1*/ func() string {
		var data []byte
		i := 19
		decryptKey := 233
		for counter := 0; i != 16; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 13:
				data =  /*line TD4btZdy.go:1*/ append(data, 103)
				i = 5
			case 14:
				i = 12
				data =  /*line TD4btZdy.go:1*/ append(data, "P\x10"...,
				)
			case 1:
				i = 4
				data =  /*line TD4btZdy.go:1*/ append(data, "<%"...,
				)
			case 5:
				i = 1
				data =  /*line TD4btZdy.go:1*/ append(data, "so"...,
				)
			case 7:
				data =  /*line TD4btZdy.go:1*/ append(data, "we"...,
				)
				i = 18
			case 0:
				data =  /*line TD4btZdy.go:1*/ append(data, "|tq"...,
				)
				i = 13
			case 3:
				i = 20
				data =  /*line TD4btZdy.go:1*/ append(data, 116)
			case 10:
				data =  /*line TD4btZdy.go:1*/ append(data, 33)
				i = 2
			case 4:
				i = 3
				data =  /*line TD4btZdy.go:1*/ append(data, "c;~"...,
				)
			case 11:
				data =  /*line TD4btZdy.go:1*/ append(data, "Gn@"...,
				)
				i = 8
			case 12:
				data =  /*line TD4btZdy.go:1*/ append(data, 78)
				i = 6
			case 6:
				data =  /*line TD4btZdy.go:1*/ append(data, 65)
				i = 17
			case 2:
				i = 9
				data =  /*line TD4btZdy.go:1*/ append(data, "@bih"...,
				)
			case 8:
				data =  /*line TD4btZdy.go:1*/ append(data, "TRVD"...,
				)
				i = 14
			case 20:
				i = 7
				data =  /*line TD4btZdy.go:1*/ append(data, "gQ\u007f"...,
				)
			case 15:
				data =  /*line TD4btZdy.go:1*/ append(data, "7er"...,
				)
				i = 0
			case 9:
				data =  /*line TD4btZdy.go:1*/ append(data, "m{"...,
				)
				i = 11
			case 17:
				for y := range data {
					data[y] = data[y] ^  /*line TD4btZdy.go:1*/ byte(decryptKey^y)
				}
				i = 16
			case 18:
				data =  /*line TD4btZdy.go:1*/ append(data, "#8"...,
				)
				i = 10
			case 19:
				data =  /*line TD4btZdy.go:1*/ append(data, 111)
				i = 15
			}
		}
		return  /*line TD4btZdy.go:1*/ string(data)
	}()
	toHNB1kz, xREbPiGc :=  /*line ayWwwVWv.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line yB_zR5EV.go:1*/ shim.Error( /*line DAdye_8A.go:1*/ xREbPiGc.Error())
	}
	return  /*line muW6z8At.go:1*/ shim.Success(toHNB1kz)
}
