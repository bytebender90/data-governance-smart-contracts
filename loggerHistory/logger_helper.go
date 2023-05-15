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
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
)

func bXRfsdhG(fOOTaXmK shim.ChaincodeStubInterface, fcEHFI_2 string) ([]byte, error) {

	 /*line DjkHpj0i.go:1*/ fmt.Printf( /*line S4dqIjLz.go:1*/ func() string {
		fullData :=  /*line S4dqIjLz.go:1*/ []byte("tɴo\x0e9<RN\xa5\x8eHNNGe\x8e\xe3::\xba\xde;\xfb\xd3_\be\a\xca+Q_e\x99\xe0\x96XF\xc0s\xf1#\xbc6GKp\xdbS\xdc\xdd\xd4kM\xab\x04g\x00\x0e\n\xc9\xc8,\xa7\fn{\xb6/\xe1*rS9\xe72s;\x87\xb8q \x9bs\x96\v\xbdV2\xdcG\xc2\x02")
		var data []byte
		data =  /*line S4dqIjLz.go:1*/ append(data, fullData[88]^fullData[67], fullData[30]^fullData[86], fullData[82]+fullData[45], fullData[58]^fullData[33], fullData[75]-fullData[77], fullData[25]^fullData[59], fullData[5]+fullData[6], fullData[29]-fullData[27], fullData[78]-fullData[61], fullData[40]^fullData[60], fullData[90]^fullData[16], fullData[9]+fullData[39], fullData[91]+fullData[63], fullData[46]+fullData[71], fullData[10]+fullData[21], fullData[92]-fullData[13], fullData[24]+fullData[84], fullData[57]^fullData[26], fullData[47]+fullData[93], fullData[51]+fullData[0], fullData[8]^fullData[22], fullData[44]+fullData[69], fullData[42]^fullData[31], fullData[43]+fullData[87], fullData[37]+fullData[23], fullData[38]^fullData[76], fullData[18]^fullData[11], fullData[28]^fullData[66], fullData[2]+fullData[20], fullData[79]^fullData[35], fullData[68]^fullData[36], fullData[7]-fullData[70], fullData[14]^fullData[89], fullData[80]-fullData[73], fullData[54]-fullData[48], fullData[17]+fullData[85], fullData[32]-fullData[65], fullData[55]+fullData[1], fullData[12]-fullData[50], fullData[49]^fullData[19], fullData[64]-fullData[74], fullData[53]-fullData[56], fullData[62]+fullData[72], fullData[34]+fullData[81], fullData[41]^fullData[52], fullData[4]-fullData[83], fullData[15]^fullData[3])
		return  /*line S4dqIjLz.go:1*/ string(data)
	}(), fcEHFI_2)

	a77C4uwl, xREbPiGc :=  /*line RhmV24bh.go:1*/ fOOTaXmK.GetQueryResult(fcEHFI_2)
	if xREbPiGc != nil {
		return nil, xREbPiGc
	}
	defer  /*line XpyuaVck.go:1*/ a77C4uwl.Close()

	qJFgybXV, xREbPiGc :=  /*line ETyhJamr.go:1*/ a39SAHUm(a77C4uwl)
	if xREbPiGc != nil {
		return nil, xREbPiGc
	}

	 /*line H4S_eW7U.go:1*/ fmt.Printf( /*line MNSBhDAN.go:1*/ func() string {
		seed :=  /*line MNSBhDAN.go:1*/ byte(119)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line MNSBhDAN.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/  /*line MNSBhDAN.go:1*/ fnc(164)(59)(189)(120)(255)(219)(218)(164)(85)(177)(59)(137)(32)(66)(123)(254)(206)(197)(141)(249)(22)(28)(69)(145)(252)(25)(48)(87)(179)(95)(119)(63)(130)(244)(245)(241)(187)(137)(32)(66)(123)(254)(194)(84)(195)(212)(63)
		return  /*line MNSBhDAN.go:1*/ string(data)
	}(),  /*line Dix3QHpe.go:1*/ qJFgybXV.String())

	return  /*line _ys7hPLh.go:1*/ qJFgybXV.Bytes(), nil
}

func a39SAHUm(a77C4uwl shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	                                                 
	var qJFgybXV bytes.Buffer
	 /*line mjvh39KD.go:1*/ qJFgybXV.WriteString( /*line iHzQmQP_.go:1*/ func() string {
		fullData :=  /*line iHzQmQP_.go:1*/ []byte("\xf2i")
		var data []byte
		data =  /*line iHzQmQP_.go:1*/ append(data, fullData[0]+fullData[1])
		return  /*line iHzQmQP_.go:1*/ string(data)
	}())

	ptH_RCSJ := false
	for  /*line zln4wMJS.go:1*/ a77C4uwl.HasNext() {
		aYd88xGv, xREbPiGc :=  /*line xorsEYdb.go:1*/ a77C4uwl.Next()
		if xREbPiGc != nil {
			return nil, xREbPiGc
		}
		                                                                           
		if ptH_RCSJ {
			 /*line NBo5WeUq.go:1*/ qJFgybXV.WriteString( /*line AO6GG8Cr.go:1*/ func() string {
				data :=  /*line AO6GG8Cr.go:1*/ []byte("\x14")
				positions := [...]byte{0, 0}
				for i := 0; i < 2; i += 2 {
					localKey :=  /*line AO6GG8Cr.go:1*/ byte(i) +  /*line AO6GG8Cr.go:1*/ byte(positions[i]^positions[i+1]) + 56
					data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
				}
				return  /*line AO6GG8Cr.go:1*/ string(data)
			}())
		}
		 /*line CL7M31Ti.go:1*/ qJFgybXV.WriteString( /*line U4mzkGJo.go:1*/ func() string {
			var data []byte
			i := 2
			decryptKey := 229
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 8:
					data =  /*line U4mzkGJo.go:1*/ append(data, 149)
					i = 4
				case 0:
					i = 6
					data =  /*line U4mzkGJo.go:1*/ append(data, 101)
				case 3:
					data =  /*line U4mzkGJo.go:1*/ append(data, 80)
					i = 7
				case 6:
					for y := range data {
						data[y] = data[y] +  /*line U4mzkGJo.go:1*/ byte(decryptKey^y)
					}
					i = 1
				case 4:
					i = 5
					data =  /*line U4mzkGJo.go:1*/ append(data, 162)
				case 5:
					data =  /*line U4mzkGJo.go:1*/ append(data, 76)
					i = 0
				case 7:
					i = 8
					data =  /*line U4mzkGJo.go:1*/ append(data, 122)
				case 2:
					i = 3
					data =  /*line U4mzkGJo.go:1*/ append(data, 168)
				}
			}
			return  /*line U4mzkGJo.go:1*/ string(data)
		}())
		 /*line I4udKpWf.go:1*/ qJFgybXV.WriteString( /*line dzwA1aK9.go:1*/ func() string {
			seed :=  /*line dzwA1aK9.go:1*/ byte(169)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line dzwA1aK9.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line dzwA1aK9.go:1*/ fnc(139)
			return  /*line dzwA1aK9.go:1*/ string(data)
		}())
		 /*line s3uwyBMp.go:1*/ qJFgybXV.WriteString(aYd88xGv.Key)
		 /*line uh1sjR66.go:1*/ qJFgybXV.WriteString( /*line CDOpphhn.go:1*/ func() string {
			seed :=  /*line CDOpphhn.go:1*/ byte(66)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line CDOpphhn.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line CDOpphhn.go:1*/ fnc(96)
			return  /*line CDOpphhn.go:1*/ string(data)
		}())

		 /*line imQBwrjc.go:1*/ qJFgybXV.WriteString( /*line q_S5uZKz.go:1*/ func() string {
			fullData :=  /*line q_S5uZKz.go:1*/ []byte("1\x1a\x17\xf1P{4\x97x\xec\xb2h\xec\a{[R\x88Y@7K")
			var data []byte
			data =  /*line q_S5uZKz.go:1*/ append(data, fullData[9]+fullData[19], fullData[5]-fullData[15], fullData[14]^fullData[18], fullData[13]+fullData[21], fullData[20]^fullData[16], fullData[7]-fullData[6], fullData[8]^fullData[2], fullData[1]^fullData[11], fullData[4]-fullData[12], fullData[3]+fullData[0], fullData[17]^fullData[10])
			return  /*line q_S5uZKz.go:1*/ string(data)
		}())
		                                             
		 /*line zbRRTweh.go:1*/ qJFgybXV.WriteString( /*line g5E2Kf9e.go:1*/ string(aYd88xGv.Value))
		 /*line LZK_Cvjc.go:1*/ qJFgybXV.WriteString( /*line BVu8cXy9.go:1*/ func() string {
			data :=  /*line BVu8cXy9.go:1*/ []byte("m")
			positions := [...]byte{0, 0}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line BVu8cXy9.go:1*/ byte(i) +  /*line BVu8cXy9.go:1*/ byte(positions[i]^positions[i+1]) + 16
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line BVu8cXy9.go:1*/ string(data)
		}())
		ptH_RCSJ = true
	}
	 /*line SLF0wj5P.go:1*/ qJFgybXV.WriteString( /*line yfQ9ddsV.go:1*/ func() string {
		fullData :=  /*line yfQ9ddsV.go:1*/ []byte("p-")
		var data []byte
		data =  /*line yfQ9ddsV.go:1*/ append(data, fullData[0]^fullData[1])
		return  /*line yfQ9ddsV.go:1*/ string(data)
	}())

	return &qJFgybXV, nil
}

                                          
func (hslgON00 *EMhVAN9S) raycgBaa(fOOTaXmK shim.ChaincodeStubInterface) (string, error) {
	c8qKfNsz, uGGsVrBp, xREbPiGc :=  /*line c9yJMGYg.go:1*/ cid.GetAttributeValue(fOOTaXmK,  /*line C_dPdqRQ.go:1*/ func() string {
		data :=  /*line C_dPdqRQ.go:1*/ []byte("\x9a\x93{rn\x8am\x8b")
		positions := [...]byte{7, 0, 5, 0, 1, 7, 2, 2, 1, 7}
		for i := 0; i < 10; i += 2 {
			localKey :=  /*line C_dPdqRQ.go:1*/ byte(i) +  /*line C_dPdqRQ.go:1*/ byte(positions[i]^positions[i+1]) + 228
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line C_dPdqRQ.go:1*/ string(data)
	}())

	if xREbPiGc != nil {
		return "", xREbPiGc
	}

	if !uGGsVrBp {
		return "",  /*line HCg7kFTm.go:1*/ errors.New( /*line MwBfCihs.go:1*/ func() string {
			key :=  /*line MwBfCihs.go:1*/ []byte("7\x87\x01@\fFV\xfb\xa7e5\x8f~\xb4\xca\x18E\x9d\xa5~\xbf\xd3\xe5A\x17\x97\x83\xd8\x01")
			data :=  /*line MwBfCihs.go:1*/ []byte("\xac\xfaf\xb2z\xa7\xc3`\xc7Ʃ\x03\xf0\x1d,\x8d\xb9\x02\xc5\xe72\xf3R\xaa\x8a\n\xecFh")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line MwBfCihs.go:1*/ string(data)
		}())
	}

	return c8qKfNsz, nil
}

                                              
func (hslgON00 *EMhVAN9S) vFrawvRh(fOOTaXmK shim.ChaincodeStubInterface) (string, error) {
	eBmFb7v9, uGGsVrBp, xREbPiGc :=  /*line srDrA_sv.go:1*/ cid.GetAttributeValue(fOOTaXmK,  /*line Ss5jLTjF.go:1*/ func() string {
		seed :=  /*line Ss5jLTjF.go:1*/ byte(143)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Ss5jLTjF.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/  /*line Ss5jLTjF.go:1*/ fnc(224)(29)(235)(22)(227)(25)(243)(29)(237)(239)(26)(225)
		return  /*line Ss5jLTjF.go:1*/ string(data)
	}())

	if xREbPiGc != nil {
		return "", xREbPiGc
	}

	if !uGGsVrBp {
		return "",  /*line Db1KmAz5.go:1*/ errors.New( /*line qiUnMCvy.go:1*/ func() string {
			fullData :=  /*line qiUnMCvy.go:1*/ []byte("\"ـ\xddHj+\x13\x85\x82\xfbu\tu#q\xb3\x1fTwq\r3\xed!<@\xd6H\xa2\x80\xfbz\x01\\\xe7\xda\xf1\xbd\xac\xef\xde\x02\x81\xfa\awF5\xc1\xe9\xffm\v:\xf4ʚ_\x9eB\xf9\xec/S\x8b")
			var data []byte
			data =  /*line qiUnMCvy.go:1*/ append(data, fullData[9]^fullData[23], fullData[4]^fullData[54], fullData[44]+fullData[52], fullData[27]+fullData[65], fullData[56]-fullData[34], fullData[14]+fullData[47], fullData[37]-fullData[19], fullData[31]-fullData[57], fullData[48]-fullData[49], fullData[50]+fullData[30], fullData[32]-fullData[53], fullData[17]^fullData[15], fullData[61]-fullData[1], fullData[42]+fullData[58], fullData[33]^fullData[13], fullData[11]+fullData[51], fullData[10]+fullData[46], fullData[36]-fullData[20], fullData[8]^fullData[35], fullData[64]+fullData[0], fullData[55]^fullData[2], fullData[18]-fullData[40], fullData[60]+fullData[41], fullData[39]+fullData[38], fullData[29]-fullData[63], fullData[7]+fullData[21], fullData[43]+fullData[62], fullData[28]^fullData[24], fullData[22]+fullData[26], fullData[5]+fullData[12], fullData[45]-fullData[59], fullData[16]^fullData[3], fullData[6]+fullData[25])
			return  /*line qiUnMCvy.go:1*/ string(data)
		}())
	}

	return eBmFb7v9, nil
}
