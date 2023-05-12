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

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (hslgON00 *EMhVAN9S) jhiv8o7v(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	var w1e7h2B3 string
	var xREbPiGc error
	if  /*line KO1RW8tl.go:1*/ len(eFz_TaL1) != 4 {
		return  /*line DakoTey3.go:1*/ shim.Error( /*line kLsuNd2I.go:1*/ func() string {
			var data []byte
			i := 33
			decryptKey := 140
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 24:
					i = 5
					for y := range data {
						data[y] = data[y] -  /*line kLsuNd2I.go:1*/ byte(decryptKey^y)
					}
				case 3:
					i = 35
					data =  /*line kLsuNd2I.go:1*/ append(data, "\x00\xfe\xf6"...,
					)
				case 26:
					i = 32
					data =  /*line kLsuNd2I.go:1*/ append(data, "nd"...,
					)
				case 14:
					data =  /*line kLsuNd2I.go:1*/ append(data, "\x13\xcf\xc6\xf3"...,
					)
					i = 36
				case 11:
					i = 0
					data =  /*line kLsuNd2I.go:1*/ append(data, 245)
				case 17:
					data =  /*line kLsuNd2I.go:1*/ append(data, 85)
					i = 29
				case 33:
					i = 23
					data =  /*line kLsuNd2I.go:1*/ append(data, "\xfb!\x13"...,
					)
				case 21:
					i = 17
					data =  /*line kLsuNd2I.go:1*/ append(data, 79)
				case 31:
					data =  /*line kLsuNd2I.go:1*/ append(data, "\x11\x10\x1a\xc9"...,
					)
					i = 4
				case 18:
					i = 14
					data =  /*line kLsuNd2I.go:1*/ append(data, 23)
				case 22:
					i = 10
					data =  /*line kLsuNd2I.go:1*/ append(data, "\x18."...,
					)
				case 6:
					data =  /*line kLsuNd2I.go:1*/ append(data, 245)
					i = 2
				case 29:
					i = 24
					data =  /*line kLsuNd2I.go:1*/ append(data, "FY"...,
					)
				case 7:
					data =  /*line kLsuNd2I.go:1*/ append(data, "\xb0\xfe\x05"...,
					)
					i = 27
				case 28:
					data =  /*line kLsuNd2I.go:1*/ append(data, "\xe8\xef\xab"...,
					)
					i = 6
				case 34:
					data =  /*line kLsuNd2I.go:1*/ append(data, "+%4"...,
					)
					i = 8
				case 1:
					i = 3
					data =  /*line kLsuNd2I.go:1*/ append(data, "\x01\a\x05\xbf"...,
					)
				case 15:
					i = 20
					data =  /*line kLsuNd2I.go:1*/ append(data, "fT\x10V"...,
					)
				case 13:
					data =  /*line kLsuNd2I.go:1*/ append(data, 67)
					i = 21
				case 35:
					data =  /*line kLsuNd2I.go:1*/ append(data, "\xe4\xa0\xf4"...,
					)
					i = 11
				case 0:
					data =  /*line kLsuNd2I.go:1*/ append(data, "\xfc\xf6"...,
					)
					i = 28
				case 8:
					i = 18
					data =  /*line kLsuNd2I.go:1*/ append(data, ")\"\x10"...,
					)
				case 32:
					data =  /*line kLsuNd2I.go:1*/ append(data, "f`"...,
					)
					i = 19
				case 25:
					i = 30
					data =  /*line kLsuNd2I.go:1*/ append(data, 2)
				case 10:
					i = 34
					data =  /*line kLsuNd2I.go:1*/ append(data, "\xdb\x19"...,
					)
				case 2:
					i = 25
					data =  /*line kLsuNd2I.go:1*/ append(data, 238)
				case 19:
					data =  /*line kLsuNd2I.go:1*/ append(data, "\x1e3\x1cs"...,
					)
					i = 13
				case 36:
					i = 31
					data =  /*line kLsuNd2I.go:1*/ append(data, "\x13\f"...,
					)
				case 16:
					data =  /*line kLsuNd2I.go:1*/ append(data, 5)
					i = 7
				case 27:
					i = 12
					data =  /*line kLsuNd2I.go:1*/ append(data, "\x05\xfd"...,
					)
				case 4:
					data =  /*line kLsuNd2I.go:1*/ append(data, "\x01\xf2\xcc"...,
					)
					i = 9
				case 12:
					i = 1
					data =  /*line kLsuNd2I.go:1*/ append(data, "\t\t\r"...,
					)
				case 23:
					i = 22
					data =  /*line kLsuNd2I.go:1*/ append(data, " ()\x19"...,
					)
				case 20:
					data =  /*line kLsuNd2I.go:1*/ append(data, "ngYX"...,
					)
					i = 26
				case 9:
					i = 16
					data =  /*line kLsuNd2I.go:1*/ append(data, "\x13\x01"...,
					)
				case 30:
					i = 15
					data =  /*line kLsuNd2I.go:1*/ append(data, "\xf0\xf0\xee"...,
					)
				}
			}
			return  /*line kLsuNd2I.go:1*/ string(data)
		}())
	}

	oB4gJZbe := eFz_TaL1[0]
	fvUjTcZo := eFz_TaL1[1]
	izyfYrYF := []string{eFz_TaL1[2]}
	xrSSF6uG := eFz_TaL1[3]
	_0uGwnuQ :=  /*line LEiSmo6o.go:1*/ func() string {
		seed :=  /*line LEiSmo6o.go:1*/ byte(5)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line LEiSmo6o.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/  /*line LEiSmo6o.go:1*/ fnc(65)(39)(25)(231)(62)(196)(26)(251)(231)(14)(52)(200)(1)(23)(233)(23)(249)(231)
		return  /*line LEiSmo6o.go:1*/ string(data)
	}()
	cKWTKAnj :=  /*line gInchMn0.go:1*/ func() string {
		var data []byte
		i := 0
		decryptKey := 98
		for counter := 0; i != 6; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 10:
				data =  /*line gInchMn0.go:1*/ append(data, "\x80e~"...,
				)
				i = 3
			case 0:
				i = 9
				data =  /*line gInchMn0.go:1*/ append(data, "Vt\x84r"...,
				)
			case 9:
				i = 7
				data =  /*line gInchMn0.go:1*/ append(data, "i\x86\x89\x87"...,
				)
			case 8:
				i = 1
				data =  /*line gInchMn0.go:1*/ append(data, 128)
			case 1:
				i = 5
				data =  /*line gInchMn0.go:1*/ append(data, 128)
			case 5:
				i = 11
				data =  /*line gInchMn0.go:1*/ append(data, "~v"...,
				)
			case 2:
				for y := range data {
					data[y] = data[y] -  /*line gInchMn0.go:1*/ byte(decryptKey^y)
				}
				i = 6
			case 3:
				data =  /*line gInchMn0.go:1*/ append(data, 146)
				i = 8
			case 4:
				i = 2
				data =  /*line gInchMn0.go:1*/ append(data, "f\u007f"...,
				)
			case 7:
				i = 10
				data =  /*line gInchMn0.go:1*/ append(data, 125)
			case 11:
				data =  /*line gInchMn0.go:1*/ append(data, "dK"...,
				)
				i = 4
			}
		}
		return  /*line gInchMn0.go:1*/ string(data)
	}()

	if  /*line iuORl2QK.go:1*/ len(oB4gJZbe) <= 0 {
		return  /*line ttWJKkWO.go:1*/ shim.Error( /*line V1KmXjRw.go:1*/ func() string {
			data :=  /*line V1KmXjRw.go:1*/ []byte("\xf2it)Sm\xf5\xddde \x04t\"58t6 m\bst.b\x05\ta n\xf6-1\vpCtc\x1a\x17a%;n_")
			positions := [...]byte{32, 31, 7, 26, 6, 3, 39, 0, 26, 7, 11, 42, 37, 26, 41, 42, 20, 7, 26, 39, 3, 38, 31, 35, 3, 17, 31, 34, 27, 44, 7, 10, 25, 40, 30, 23, 34, 10, 13, 0, 14, 15, 44, 1, 5, 35, 7, 5, 40, 38, 33, 32, 30, 8}
			for i := 0; i < 54; i += 2 {
				localKey :=  /*line V1KmXjRw.go:1*/ byte(i) +  /*line V1KmXjRw.go:1*/ byte(positions[i]^positions[i+1]) + 171
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line V1KmXjRw.go:1*/ string(data)
		}())
	}

	if  /*line LFoUbTMl.go:1*/ len(fvUjTcZo) <= 0 {
		return  /*line eJp4W7vx.go:1*/ shim.Error( /*line KM2J7by7.go:1*/ func() string {
			var data []byte
			i := 9
			decryptKey := 11
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 8:
					i = 0
					data =  /*line KM2J7by7.go:1*/ append(data, "\xc9\xc5¼"...,
					)
				case 0:
					data =  /*line KM2J7by7.go:1*/ append(data, "\xc0\xd4"...,
					)
					i = 10
				case 12:
					data =  /*line KM2J7by7.go:1*/ append(data, "\xc8\xca\xc9"...,
					)
					i = 1
				case 10:
					i = 6
					for y := range data {
						data[y] = data[y] -  /*line KM2J7by7.go:1*/ byte(decryptKey^y)
					}
				case 3:
					i = 7
					data =  /*line KM2J7by7.go:1*/ append(data, "\xb0\xc2"...,
					)
				case 4:
					data =  /*line KM2J7by7.go:1*/ append(data, "\xcĕ\xbd"...,
					)
					i = 12
				case 9:
					i = 3
					data =  /*line KM2J7by7.go:1*/ append(data, "\xa0\xc0"...,
					)
				case 7:
					data =  /*line KM2J7by7.go:1*/ append(data, "\xbe\xbbk\xb7"...,
					)
					i = 5
				case 2:
					data =  /*line KM2J7by7.go:1*/ append(data, "}\xca"...,
					)
					i = 4
				case 1:
					data =  /*line KM2J7by7.go:1*/ append(data, "\xcdw"...,
					)
					i = 8
				case 5:
					i = 11
					data =  /*line KM2J7by7.go:1*/ append(data, "\xba\xb7\xbbf"...,
					)
				case 11:
					data =  /*line KM2J7by7.go:1*/ append(data, "\xa3\xa5c\xa3"...,
					)
					i = 2
				}
			}
			return  /*line KM2J7by7.go:1*/ string(data)
		}())
	}

	if  /*line y0_Zjf9t.go:1*/ len(xrSSF6uG) <= 0 {
		return  /*line LZWJrHEr.go:1*/ shim.Error( /*line pRpzwE4j.go:1*/ func() string {
			data :=  /*line pRpzwE4j.go:1*/ []byte("Tv\x00e\xbatG\x01B $\x1b7\xf0\xee\x1a\xaf=#W\f*\x13\xb3\xbfmpt\v\xf1str?ng")
			positions := [...]byte{10, 13, 14, 22, 1, 28, 18, 20, 28, 2, 23, 4, 4, 17, 33, 6, 21, 22, 11, 24, 17, 16, 12, 7, 1, 12, 28, 15, 17, 29, 23, 23, 18, 11, 10, 2, 8, 24, 15, 28, 19, 10, 14, 18}
			for i := 0; i < 44; i += 2 {
				localKey :=  /*line pRpzwE4j.go:1*/ byte(i) +  /*line pRpzwE4j.go:1*/ byte(positions[i]^positions[i+1]) + 169
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line pRpzwE4j.go:1*/ string(data)
		}())
	}

	faMQ13o3 :=  /*line oAToX53U.go:1*/ func() string {
		data :=  /*line oAToX53U.go:1*/ []byte("dMt\x81S\xf3\x00\xf7ɓߗtz\x05a$aK\x02\xf8")
		positions := [...]byte{1, 6, 11, 16, 6, 19, 8, 10, 9, 11, 1, 13, 5, 11, 20, 1, 13, 5, 3, 3, 16, 9, 14, 7, 13, 3}
		for i := 0; i < 26; i += 2 {
			localKey :=  /*line oAToX53U.go:1*/ byte(i) +  /*line oAToX53U.go:1*/ byte(positions[i]^positions[i+1]) + 116
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line oAToX53U.go:1*/ string(data)
	}() + oB4gJZbe
	MNfTuYHi := DataSourceMetadataCounter{}
	sdHID17m := -1
	dWLU1lja, xREbPiGc :=  /*line J4_xZuSn.go:1*/ fOOTaXmK.GetState(faMQ13o3)
	if xREbPiGc != nil {
		w1e7h2B3 =  /*line DpFRAVXA.go:1*/ func() string {
			fullData :=  /*line DpFRAVXA.go:1*/ []byte("4\x8f\x1d.}\x1f\x0e\xb1$\xa3\xd1\x02^\"\x04ѫ=\xa8$XdGp\r\xddPCY\x04\x93\\\xa1\x91\x99u\xa6\xda\xf0\x15\xc1\x03\xae\xd1rQ\\p\xe20\x06\xc6\u007f\xe9\xcb_X/\xc3L\xe0\x95\xad~\xa2\xfe\x9f/q\xc0b\x13\xc0z,\xeb\x13\x83P\xf8e\x1d\xbbΒ7DƘ!\xb0_\xb8U\xe6z\u00827z\xdb\xd7L2Y\xb0\xfe\x1b\xe9\x0f\"\xf9\xae\xa3")
			var data []byte
			data =  /*line DpFRAVXA.go:1*/ append(data, fullData[102]-fullData[10], fullData[19]-fullData[11], fullData[65]+fullData[22], fullData[52]-fullData[24], fullData[84]+fullData[60], fullData[47]^fullData[5], fullData[26]+fullData[110], fullData[95]-fullData[56], fullData[97]^fullData[92], fullData[41]^fullData[89], fullData[42]+fullData[88], fullData[30]-fullData[103], fullData[12]^fullData[85], fullData[77]+fullData[108], fullData[20]^fullData[17], fullData[69]-fullData[46], fullData[1]+fullData[33], fullData[8]-fullData[90], fullData[55]-fullData[38], fullData[35]+fullData[16], fullData[72]-fullData[104], fullData[106]-fullData[34], fullData[57]-fullData[82], fullData[98]+fullData[53], fullData[9]+fullData[40], fullData[28]-fullData[79], fullData[36]+fullData[83], fullData[2]+fullData[86], fullData[94]^fullData[87], fullData[73]+fullData[111], fullData[61]+fullData[37], fullData[78]-fullData[100], fullData[13]-fullData[105], fullData[32]^fullData[96], fullData[101]-fullData[44], fullData[109]^fullData[67], fullData[39]-fullData[18], fullData[54]^fullData[112], fullData[99]-fullData[50], fullData[29]^fullData[80], fullData[0]+fullData[49], fullData[63]-fullData[81], fullData[75]^fullData[66], fullData[62]-fullData[59], fullData[45]^fullData[68], fullData[14]+fullData[21], fullData[113]+fullData[51], fullData[64]+fullData[15], fullData[58]+fullData[7], fullData[31]+fullData[76], fullData[71]+fullData[91], fullData[74]^fullData[93], fullData[3]-fullData[6], fullData[4]^fullData[107], fullData[43]-fullData[70], fullData[48]-fullData[23], fullData[25]+fullData[27])
			return  /*line DpFRAVXA.go:1*/ string(data)
		}() + faMQ13o3 +  /*line lkgkFU5u.go:1*/ func() string {
			key :=  /*line lkgkFU5u.go:1*/ []byte("kH")
			data :=  /*line lkgkFU5u.go:1*/ []byte("\x8d\xc5")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line lkgkFU5u.go:1*/ string(data)
		}()
		return  /*line DYirEQAY.go:1*/ shim.Error(w1e7h2B3)
	} else if dWLU1lja == nil {
		sdHID17m = 0
		oB4gJZbe := &LoggerRevokedAccessCounter{cKWTKAnj, oB4gJZbe, sdHID17m}
		dWLU1lja, xREbPiGc =  /*line ecinlOkk.go:1*/ json.Marshal(oB4gJZbe)
		if xREbPiGc != nil {
			return  /*line sxEvwqpe.go:1*/ shim.Error( /*line sBtGDbzv.go:1*/ xREbPiGc.Error())
		}

	} else {
		xREbPiGc =  /*line HzgpRU9y.go:1*/ json.Unmarshal(dWLU1lja, &MNfTuYHi)	                               
		if xREbPiGc != nil {
			return  /*line RJQyDlIJ.go:1*/ shim.Error( /*line uou9DuDu.go:1*/ xREbPiGc.Error())
		}
		sdHID17m = MNfTuYHi.Count + 1
		MNfTuYHi.Count = sdHID17m
		dWLU1lja, _ =  /*line CKEfKBR9.go:1*/ json.Marshal(MNfTuYHi)
	}

	                                                     
	xREbPiGc =  /*line vPICNcbj.go:1*/ fOOTaXmK.PutState(faMQ13o3, dWLU1lja)
	if xREbPiGc != nil {
		return  /*line zoH8u1I1.go:1*/ shim.Error( /*line Dr_X_pzA.go:1*/ xREbPiGc.Error())
	}

	r3ovXd_R := &DataSourceMetadata{_0uGwnuQ, oB4gJZbe, fvUjTcZo, izyfYrYF, xrSSF6uG}
	pF2KGYpq, xREbPiGc :=  /*line uL5fCtzI.go:1*/ json.Marshal(r3ovXd_R)
	if xREbPiGc != nil {
		return  /*line aX4ERd49.go:1*/ shim.Error( /*line teZCCmu4.go:1*/ xREbPiGc.Error())
	}
	iPrzUCH_ :=  /*line yrz7Rl9i.go:1*/ func() string {
		fullData :=  /*line yrz7Rl9i.go:1*/ []byte("f}\x06\x02J\xbd5\x8bj\xe0\xd1\xd0\xc7g\xd8]3*R\x1dR\xfa\xd6_\xcfAo\xaa\x0f(#\x8b߁\"\t\xdc3")
		var data []byte
		data =  /*line yrz7Rl9i.go:1*/ append(data, fullData[0]^fullData[3], fullData[33]+fullData[9], fullData[4]+fullData[17], fullData[2]^fullData[13], fullData[7]^fullData[14], fullData[19]+fullData[18], fullData[32]^fullData[27], fullData[24]-fullData[15], fullData[16]-fullData[11], fullData[23]-fullData[21], fullData[34]^fullData[26], fullData[28]^fullData[8], fullData[1]-fullData[35], fullData[29]-fullData[12], fullData[6]-fullData[10], fullData[22]+fullData[31], fullData[25]+fullData[37], fullData[36]^fullData[5], fullData[20]-fullData[30])
		return  /*line yrz7Rl9i.go:1*/ string(data)
	}() + oB4gJZbe +  /*line VYX0eksx.go:1*/ func() string {
		fullData :=  /*line VYX0eksx.go:1*/ []byte("a\x90")
		var data []byte
		data =  /*line VYX0eksx.go:1*/ append(data, fullData[1]-fullData[0])
		return  /*line VYX0eksx.go:1*/ string(data)
	}() +  /*line I01pzY1c.go:1*/ strconv.Itoa(sdHID17m)
	                                       
	xREbPiGc =  /*line jzqKGq95.go:1*/ fOOTaXmK.PutState(iPrzUCH_, pF2KGYpq)
	if xREbPiGc != nil {
		return  /*line zd6pRfg8.go:1*/ shim.Error( /*line _y5VrqnK.go:1*/ xREbPiGc.Error())
	}

	return  /*line CumEEcNa.go:1*/ shim.Success(nil)

}

func (hslgON00 *EMhVAN9S) ghDiHhnv(fOOTaXmK shim.ChaincodeStubInterface) pb.Response {

	fcEHFI_2 :=  /*line Lk3fvuiz.go:1*/ fmt.Sprintf( /*line C1EIeldj.go:1*/ func() string {
		seed :=  /*line C1EIeldj.go:1*/ byte(97)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line C1EIeldj.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/  /*line C1EIeldj.go:1*/ fnc(220)(95)(15)(16)(39)(71)(140)(41)(77)(157)(234)(236)(25)(217)(244)(243)(218)(165)(111)(213)(159)(251)(14)(4)(42)(113)(245)(215)(160)(92)(190)(121)(227)(200)(120)(8)(31)(43)(89)(175)(113)(207)(95)(25)(50)
		return  /*line C1EIeldj.go:1*/ string(data)
	}())
	toHNB1kz, xREbPiGc :=  /*line Dz5Oz3se.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line SlCBtEma.go:1*/ shim.Error( /*line JXos_Q0F.go:1*/ xREbPiGc.Error())
	}
	return  /*line fGr0jji9.go:1*/ shim.Success(toHNB1kz)
}

func (hslgON00 *EMhVAN9S) gMvMnZkl(fOOTaXmK shim.ChaincodeStubInterface, eFz_TaL1 []string) pb.Response {
	if  /*line a8Zit0pr.go:1*/ len(eFz_TaL1) != 1 {
		return  /*line jndCwOCB.go:1*/ shim.Error( /*line OZDbd54v.go:1*/ func() string {
			seed :=  /*line OZDbd54v.go:1*/ byte(70)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line OZDbd54v.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/  /*line OZDbd54v.go:1*/ fnc(3)(37)(8)(235)(11)(253)(251)(188)(33)(49)(245)(14)(248)(248)(9)(6)(172)(46)(39)(248)(245)(3)(13)(188)(242)(37)(51)(248)(245)(254)(17)(245)(5)(249)(185)(68)(253)(19)(237)(18)(252)(6)(253)(241)(2)(187)(77)(248)(15)(237)(3)(253)(19)(237)(191)(78)(243)(12)(248)
			return  /*line OZDbd54v.go:1*/ string(data)
		}())
	}
	oB4gJZbe := eFz_TaL1[0]

	                  
	fcEHFI_2 :=  /*line wvFekkN9.go:1*/ fmt.Sprintf( /*line pOWJq_mz.go:1*/ func() string {
		var data []byte
		i := 20
		decryptKey := 155
		for counter := 0; i != 10; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 19:
				i = 28
				data =  /*line pOWJq_mz.go:1*/ append(data, 104)
			case 9:
				data =  /*line pOWJq_mz.go:1*/ append(data, "_\t"...,
				)
				i = 8
			case 16:
				data =  /*line pOWJq_mz.go:1*/ append(data, "\f\b"...,
				)
				i = 9
			case 14:
				i = 30
				data =  /*line pOWJq_mz.go:1*/ append(data, "G^\x18"...,
				)
			case 4:
				i = 15
				data =  /*line pOWJq_mz.go:1*/ append(data, "))9<"...,
				)
			case 27:
				i = 18
				data =  /*line pOWJq_mz.go:1*/ append(data, " \x1d\x04\x02"...,
				)
			case 7:
				data =  /*line pOWJq_mz.go:1*/ append(data, 41)
				i = 12
			case 8:
				data =  /*line pOWJq_mz.go:1*/ append(data, 87)
				i = 29
			case 1:
				i = 22
				data =  /*line pOWJq_mz.go:1*/ append(data, "XC"...,
				)
			case 11:
				i = 21
				data =  /*line pOWJq_mz.go:1*/ append(data, "2&"...,
				)
			case 6:
				data =  /*line pOWJq_mz.go:1*/ append(data, "\n\x1c\b\x14"...,
				)
				i = 14
			case 25:
				i = 1
				data =  /*line pOWJq_mz.go:1*/ append(data, "\f\x1e"...,
				)
			case 30:
				data =  /*line pOWJq_mz.go:1*/ append(data, "@\x05\x0f"...,
				)
				i = 31
			case 0:
				i = 27
				data =  /*line pOWJq_mz.go:1*/ append(data, "\x01\x15"...,
				)
			case 15:
				data =  /*line pOWJq_mz.go:1*/ append(data, "\x152\""...,
				)
				i = 32
			case 29:
				data =  /*line pOWJq_mz.go:1*/ append(data, 84)
				i = 3
			case 3:
				i = 10
				for y := range data {
					data[y] = data[y] ^  /*line pOWJq_mz.go:1*/ byte(decryptKey^y)
				}
			case 13:
				data =  /*line pOWJq_mz.go:1*/ append(data, ")?+-"...,
				)
				i = 7
			case 26:
				data =  /*line pOWJq_mz.go:1*/ append(data, "\t\a\x0f"...,
				)
				i = 6
			case 31:
				data =  /*line pOWJq_mz.go:1*/ append(data, "\x1c*"...,
				)
				i = 24
			case 12:
				data =  /*line pOWJq_mz.go:1*/ append(data, "3'"...,
				)
				i = 23
			case 20:
				data =  /*line pOWJq_mz.go:1*/ append(data, "\x14L\x1e"...,
				)
				i = 26
			case 28:
				data =  /*line pOWJq_mz.go:1*/ append(data, "a& 4"...,
				)
				i = 5
			case 23:
				data =  /*line pOWJq_mz.go:1*/ append(data, 103)
				i = 19
			case 18:
				data =  /*line pOWJq_mz.go:1*/ append(data, ",+\x00"...,
				)
				i = 13
			case 22:
				i = 17
				data =  /*line pOWJq_mz.go:1*/ append(data, "Z3"...,
				)
			case 17:
				data =  /*line pOWJq_mz.go:1*/ append(data, 23)
				i = 0
			case 2:
				data =  /*line pOWJq_mz.go:1*/ append(data, 48)
				i = 11
			case 24:
				i = 25
				data =  /*line pOWJq_mz.go:1*/ append(data, 4)
			case 5:
				data =  /*line pOWJq_mz.go:1*/ append(data, ">\r2"...,
				)
				i = 4
			case 32:
				i = 2
				data =  /*line pOWJq_mz.go:1*/ append(data, 52)
			case 21:
				i = 16
				data =  /*line pOWJq_mz.go:1*/ append(data, "0r\x15"...,
				)
			}
		}
		return  /*line pOWJq_mz.go:1*/ string(data)
	}(), oB4gJZbe)
	toHNB1kz, xREbPiGc :=  /*line c2bDlm3e.go:1*/ bXRfsdhG(fOOTaXmK, fcEHFI_2)
	if xREbPiGc != nil {
		return  /*line Azqudfns.go:1*/ shim.Error( /*line JECLIsES.go:1*/ xREbPiGc.Error())
	}
	return  /*line F7pwhaS6.go:1*/ shim.Success(toHNB1kz)
}
