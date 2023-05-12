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

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

                     
func (u3xzONWW *ZohimWzR) iYOlV2HM(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	var dTakwHTF error
	var ggIJA7_C string

	iztzjRXX := lsLFSQCl[0]
	lkgLredg := lsLFSQCl[1]

	lw5scISX, dTakwHTF :=  /*line AKD90AMn.go:1*/ u3xzONWW.nyOeNqDh(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line OjXbdKXc.go:1*/ shim.Error( /*line FeHPy37G.go:1*/ dTakwHTF.Error())
	}

	cVwY5UTY, dTakwHTF :=  /*line xfrLfX8p.go:1*/ u3xzONWW.vz2OhzsG(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line TzjTJJ9n.go:1*/ shim.Error( /*line jsQKXnNQ.go:1*/ dTakwHTF.Error())
	}

	_br98sk5, dTakwHTF :=  /*line rbDuGiC7.go:1*/ u3xzONWW.qbaWnNO0(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line DlFFxvqW.go:1*/ shim.Error( /*line ReBuKsUe.go:1*/ dTakwHTF.Error())
	}

	mBkwo0i6, dTakwHTF :=  /*line U0fO6Brr.go:1*/ ipEgxdqE.GetState(iztzjRXX)
	if dTakwHTF != nil {
		ggIJA7_C =  /*line oQyX9wD_.go:1*/ func() string {
			key :=  /*line oQyX9wD_.go:1*/ []byte("\xde>\x04D\x9c\xb8\xf4\xe3\xef\x98\xfbm\x14\x15\xc6}\x00\x19G\xbf\x93ɚb\xe6\x03\xa0\xa47\xba#7X\xf2")
			data :=  /*line oQyX9wD_.go:1*/ []byte("Y`I\xb6\x0e'f\x05)\xbaA\xce}\x81+\xe1 \x8d\xb6\xdf\xfa.\x0e\x82Yw\x01\x18\x9cډ\xa6\xca\x12")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line oQyX9wD_.go:1*/ string(data)
		}() + iztzjRXX +  /*line a8zVNWvS.go:1*/ func() string {
			seed :=  /*line a8zVNWvS.go:1*/ byte(187)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line a8zVNWvS.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line a8zVNWvS.go:1*/  /*line a8zVNWvS.go:1*/ fnc(221)(21)
			return  /*line a8zVNWvS.go:1*/ string(data)
		}()
		return  /*line OLhU2XgY.go:1*/ shim.Error(ggIJA7_C)
	} else if mBkwo0i6 == nil {
		 /*line gdZAP7ny.go:1*/ fmt.Println( /*line pLeynlH9.go:1*/ func() string {
			var data []byte
			i := 5
			decryptKey := 19
			for counter := 0; i != 8; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 13:
					i = 4
					data =  /*line pLeynlH9.go:1*/ append(data, "\xa7p"...,
					)
				case 11:
					i = 6
					data =  /*line pLeynlH9.go:1*/ append(data, 148)
				case 0:
					i = 2
					data =  /*line pLeynlH9.go:1*/ append(data, 163)
				case 7:
					data =  /*line pLeynlH9.go:1*/ append(data, "E\x8c\x96\x9f"...,
					)
					i = 9
				case 9:
					i = 1
					data =  /*line pLeynlH9.go:1*/ append(data, "\xac\\\xa9"...,
					)
				case 1:
					i = 12
					data =  /*line pLeynlH9.go:1*/ append(data, "\xad\xb1`"...,
					)
				case 6:
					i = 7
					data =  /*line pLeynlH9.go:1*/ append(data, "\x8f\x85\x90\x8b"...,
					)
				case 4:
					i = 3
					data =  /*line pLeynlH9.go:1*/ append(data, 85)
				case 10:
					data =  /*line pLeynlH9.go:1*/ append(data, "\x95\x9eN\xa2"...,
					)
					i = 0
				case 3:
					for y := range data {
						data[y] = data[y] +  /*line pLeynlH9.go:1*/ byte(decryptKey^y)
					}
					i = 8
				case 12:
					i = 13
					data =  /*line pLeynlH9.go:1*/ append(data, "\xa4\xaa\x9a\xa7"...,
					)
				case 5:
					i = 10
					data =  /*line pLeynlH9.go:1*/ append(data, "~\x91"...,
					)
				case 2:
					data =  /*line pLeynlH9.go:1*/ append(data, 148)
					i = 11
				}
			}
			return  /*line pLeynlH9.go:1*/ string(data)
		}() + iztzjRXX)
		return  /*line s_VP0Fib.go:1*/ shim.Error( /*line GfYd3U6N.go:1*/ func() string {
			key :=  /*line GfYd3U6N.go:1*/ []byte("^-\xa2\xa9\xac\xf7L\x82\x13\xadi\xf0\xbd\x9fr\xee\xca,\x1a\xba&\x9c\x92\xd3\xfas]\x04b#")
			data :=  /*line GfYd3U6N.go:1*/ []byte("\nE\xcbڌ\x82?\xe7a\xc3\b\x9dؿ\x16\x81\xaf_:\xd4I貶\x82\x1a.pX\x03")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line GfYd3U6N.go:1*/ string(data)
		}() + iztzjRXX)
	}

	biXWDyeJ := UserCredentials{}
	dTakwHTF =  /*line UXWlGIXB.go:1*/ json.Unmarshal(mBkwo0i6, &biXWDyeJ)	                               
	if dTakwHTF != nil {
		return  /*line VvLKRlU1.go:1*/ shim.Error( /*line DnEVEbNz.go:1*/ dTakwHTF.Error())
	}

	skhCclRZ, dTakwHTF :=  /*line HNf1iJ7e.go:1*/ ipEgxdqE.GetState(biXWDyeJ.UserId)
	if dTakwHTF != nil {
		ggIJA7_C =  /*line zQzEAQA_.go:1*/ func() string {
			key :=  /*line zQzEAQA_.go:1*/ []byte("\x84\xba\xa2\xc0\xcee\xb1\xab?R\x93\xfb\xc45\x9fA\xec\xf9\xd7\"\xe8\xa4\x03V\xbc!\xee8\xf0\xe4U\xf8~\x00")
			data :=  /*line zQzEAQA_.go:1*/ []byte("\xff\xdc\xe72@\xd4#\xcdyt\xd9\\-\xa1\x04\xa5\fmFBO\twv/\x95O\xacU\x04\xbbg\xf0 ")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line zQzEAQA_.go:1*/ string(data)
		}() + biXWDyeJ.UserId +  /*line CH4VGgQc.go:1*/ func() string {
			data :=  /*line CH4VGgQc.go:1*/ []byte("\"\xbb")
			positions := [...]byte{1, 1}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line CH4VGgQc.go:1*/ byte(i) +  /*line CH4VGgQc.go:1*/ byte(positions[i]^positions[i+1]) + 194
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line CH4VGgQc.go:1*/ string(data)
		}()
		return  /*line tZz9g9KH.go:1*/ shim.Error(ggIJA7_C)
	} else if skhCclRZ == nil {
		 /*line vjYVJp7H.go:1*/ fmt.Println( /*line g_p8Q1jv.go:1*/ func() string {
			fullData :=  /*line g_p8Q1jv.go:1*/ []byte("\xff\x0e\x98\xf3\xfc5uE\f\x80\"\x04\xeb-e\x8c\xb7\tF\x17IF$\xf2\xff\x01\x1dieԂ\x8cl\xf8mN\x9b\x94&\xa3ef\x05;\x02i\x80g\xb9>\x95X")
			var data []byte
			data =  /*line g_p8Q1jv.go:1*/ append(data, fullData[48]-fullData[28], fullData[0]+fullData[45], fullData[35]+fullData[19], fullData[23]+fullData[9], fullData[33]+fullData[34], fullData[15]+fullData[37], fullData[42]^fullData[32], fullData[1]-fullData[36], fullData[41]-fullData[21], fullData[47]^fullData[17], fullData[20]+fullData[38], fullData[7]^fullData[14], fullData[13]^fullData[51], fullData[24]-fullData[31], fullData[30]-fullData[26], fullData[40]-fullData[3], fullData[5]+fullData[12], fullData[44]^fullData[6], fullData[29]+fullData[50], fullData[46]-fullData[8], fullData[39]-fullData[43], fullData[22]^fullData[11], fullData[25]-fullData[2], fullData[18]^fullData[10], fullData[4]+fullData[49], fullData[27]+fullData[16])
			return  /*line g_p8Q1jv.go:1*/ string(data)
		}() + biXWDyeJ.UserId)
		return  /*line KK0dJ8jV.go:1*/ shim.Error( /*line oNuDG9_D.go:1*/ func() string {
			var data []byte
			i := 1
			decryptKey := 12
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 10:
					data =  /*line oNuDG9_D.go:1*/ append(data, "\vAZ\x06"...,
					)
					i = 11
				case 6:
					i = 2
					for y := range data {
						data[y] = data[y] ^  /*line oNuDG9_D.go:1*/ byte(decryptKey^y)
					}
				case 4:
					i = 9
					data =  /*line oNuDG9_D.go:1*/ append(data, 73)
				case 3:
					data =  /*line oNuDG9_D.go:1*/ append(data, 23)
					i = 6
				case 0:
					data =  /*line oNuDG9_D.go:1*/ append(data, "Q]\f"...,
					)
					i = 3
				case 5:
					data =  /*line oNuDG9_D.go:1*/ append(data, "UIR\x1b"...,
					)
					i = 0
				case 7:
					data =  /*line oNuDG9_D.go:1*/ append(data, 72)
					i = 5
				case 1:
					data =  /*line oNuDG9_D.go:1*/ append(data, "zG"...,
					)
					i = 4
				case 11:
					data =  /*line oNuDG9_D.go:1*/ append(data, "IK\x05W"...,
					)
					i = 8
				case 9:
					data =  /*line oNuDG9_D.go:1*/ append(data, "_O"...,
					)
					i = 10
				case 8:
					data =  /*line oNuDG9_D.go:1*/ append(data, "PES\x1e"...,
					)
					i = 7
				}
			}
			return  /*line oNuDG9_D.go:1*/ string(data)
		}() + biXWDyeJ.UserId)
	}

	jC3LR1TY := User{}
	dTakwHTF =  /*line CT4zIOEF.go:1*/ json.Unmarshal(skhCclRZ, &jC3LR1TY)	                               
	if dTakwHTF != nil {
		return  /*line A_oBR9bi.go:1*/ shim.Error( /*line p39IJcyF.go:1*/ dTakwHTF.Error())
	}

	jC3LR1TY.Status = lkgLredg

	if cVwY5UTY ==  /*line CO2rom_e.go:1*/ func() string {
		fullData :=  /*line CO2rom_e.go:1*/ []byte("\x9e\xb6\xb7\x8c\xa3\xe8\a\xcb\x0eo")
		var data []byte
		data =  /*line CO2rom_e.go:1*/ append(data, fullData[8]^fullData[9], fullData[3]^fullData[5], fullData[1]+fullData[2], fullData[6]-fullData[0], fullData[7]+fullData[4])
		return  /*line CO2rom_e.go:1*/ string(data)
	}() {
		if _br98sk5 != biXWDyeJ.Organization || jC3LR1TY.ExternalOrg != lw5scISX {
			return  /*line SV2kIybM.go:1*/ shim.Error( /*line GlsISzti.go:1*/ func() string {
				key :=  /*line GlsISzti.go:1*/ []byte("\xaa2\xc7\xc0\x11O\xca\xc2\xd4\xe7h\x13bJ\xc2,WX!\x06\x9c~\x8bHג\x8a\xc8\x14\xc5,_o\xbeÙ\x06rtg\xff\xe7\x11\x1dK\x84\xc2&\xac\x8cC\xffVe\xf9\xfc\xd4\xf3?\x8bF\x1er\x1c\x81\xaa[k\xe4\xbc\xfd\xc0")
				data :=  /*line GlsISzti.go:1*/ []byte("\xfe\x9a,\xe0\x80\xc11\xe25K\xd5|\xd0j%\x8d\xc5x\x94k\x10\x9e\xf4\xb68\xf5\xfe1\x8a*L\xc0\xd2!2\x0et\xe6\xe7\x87nU}\x96k\xea1\x98\xcc\x00\xabdv\xd4pj9W_\xfa\xb8\x85ӊ\xea$\xbc\xdfM+k3")
				for i, b := range key {
					data[i] = data[i] - b
				}
				return  /*line GlsISzti.go:1*/ string(data)
			}())
		}

		if biXWDyeJ.Type !=  /*line mciNWeEy.go:1*/ func() string {
			fullData :=  /*line mciNWeEy.go:1*/ []byte("92@\xf2:\x97\xee\x87")
			var data []byte
			data =  /*line mciNWeEy.go:1*/ append(data, fullData[7]+fullData[6], fullData[4]+fullData[0], fullData[5]^fullData[3], fullData[1]+fullData[2])
			return  /*line mciNWeEy.go:1*/ string(data)
		}() {
			return  /*line JgSL3BvW.go:1*/ shim.Error( /*line NNVhWszr.go:1*/ func() string {
				seed :=  /*line NNVhWszr.go:1*/ byte(29)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data =  /*line NNVhWszr.go:1*/ append(data, x^seed); seed += x; return fnc }
				 /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/  /*line NNVhWszr.go:1*/ fnc(73)(14)(17)(165)(69)(29)(235)(87)(175)(25)(251)(248)(231)(80)(163)(2)(11)(80)(164)(1)(4)(10)(7)(19)(251)(233)(5)(19)(169)(93)(225)(28)(245)(161)(87)(10)(230)(27)(164)(73)(18)(224)(12)(26)(231)(4)
				return  /*line NNVhWszr.go:1*/ string(data)
			}())
		}
	} else if cVwY5UTY ==  /*line JTjFKDQ7.go:1*/ func() string {
		fullData :=  /*line JTjFKDQ7.go:1*/ []byte("\xdfa\xb7t)\xd6\x11l")
		var data []byte
		data =  /*line JTjFKDQ7.go:1*/ append(data, fullData[5]-fullData[1], fullData[0]-fullData[7], fullData[3]^fullData[6], fullData[4]-fullData[2])
		return  /*line JTjFKDQ7.go:1*/ string(data)
	}() {
		return  /*line QPqQTixB.go:1*/ shim.Error( /*line idSc2ZFN.go:1*/ func() string {
			fullData :=  /*line idSc2ZFN.go:1*/ []byte("\xeeC\xbe3\x12\x1a\x00\xe8\xd4~E\\\xad\xc2\xfb\x8ct4\xac\xfb\b\\\xe5\v-HP\xcc^9q\xabMm*O(\xee\a\x97\xa2pR\x19\xb1\x89\x1e\x87ukUd:MٵA0\xb7\x95\x8d;y\x19\x1e\x05\b!!{")
			var data []byte
			data =  /*line idSc2ZFN.go:1*/ append(data, fullData[8]+fullData[69], fullData[62]-fullData[23], fullData[57]^fullData[11], fullData[9]+fullData[14], fullData[18]+fullData[16], fullData[40]-fullData[56], fullData[7]^fullData[15], fullData[52]+fullData[3], fullData[47]^fullData[37], fullData[2]-fullData[26], fullData[67]+fullData[42], fullData[48]+fullData[31], fullData[46]+fullData[10], fullData[50]^fullData[17], fullData[32]+fullData[68], fullData[36]^fullData[20], fullData[51]+fullData[6], fullData[0]-fullData[45], fullData[25]+fullData[63], fullData[66]^fullData[49], fullData[22]-fullData[30], fullData[35]+fullData[5], fullData[19]^fullData[60], fullData[4]-fullData[44], fullData[43]^fullData[33], fullData[61]+fullData[34], fullData[59]^fullData[55], fullData[1]+fullData[64], fullData[21]+fullData[38], fullData[39]+fullData[27], fullData[13]+fullData[12], fullData[41]+fullData[65], fullData[58]^fullData[54], fullData[29]^fullData[53], fullData[28]^fullData[24])
			return  /*line idSc2ZFN.go:1*/ string(data)
		}())
	} else if cVwY5UTY ==  /*line NlS1DcD0.go:1*/ func() string {
		seed :=  /*line NlS1DcD0.go:1*/ byte(207)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line NlS1DcD0.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/  /*line NlS1DcD0.go:1*/ fnc(61)(113)(241)(229)(194)(135)(7)(209)(213)(173)(99)(194)(137)
		return  /*line NlS1DcD0.go:1*/ string(data)
	}() && _br98sk5 != biXWDyeJ.Organization {

		return  /*line nyEgJI7h.go:1*/ shim.Error( /*line B4pVxWPM.go:1*/ func() string {
			seed :=  /*line B4pVxWPM.go:1*/ byte(203)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line B4pVxWPM.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/  /*line B4pVxWPM.go:1*/ fnc(25)(73)(161)(69)(130)(7)(7)(195)(199)(145)(43)(82)(169)(87)(91)(249)(240)(237)(140)(92)(185)(110)(222)(205)(143)(43)(65)(149)(27)(241)(35)(72)(144)(44)(94)(181)(112)(223)(107)(37)(73)(144)(45)(1)(72)(153)(53)(24)(132)(252)(245)(238)(229)(120)(57)(119)(244)(217)(191)(122)(231)(217)(102)(27)(57)(103)(200)(157)(53)(123)(221)(205)(143)(36)(71)
			return  /*line B4pVxWPM.go:1*/ string(data)
		}())

	}

	nQa48sqX, _ :=  /*line xG6arpXc.go:1*/ json.Marshal(jC3LR1TY)
	dTakwHTF =  /*line OdvRk1Xw.go:1*/ ipEgxdqE.PutState(biXWDyeJ.UserId, nQa48sqX)	                       
	if dTakwHTF != nil {
		return  /*line mGvRkr75.go:1*/ shim.Error( /*line eWGf1CMZ.go:1*/ dTakwHTF.Error())
	}

	gd7Gzym9, _ :=  /*line nLkZKdAD.go:1*/ json.Marshal(biXWDyeJ)
	dTakwHTF =  /*line KMoPhVQK.go:1*/ ipEgxdqE.PutState(iztzjRXX, gd7Gzym9)	                             
	if dTakwHTF != nil {
		return  /*line z1BLidMT.go:1*/ shim.Error( /*line tWhu4t3R.go:1*/ dTakwHTF.Error())
	}

	return  /*line F3pwUZlq.go:1*/ shim.Success(nil)
}

                 
func (u3xzONWW *ZohimWzR) _IMfdl2d(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	var dTakwHTF error
	var ggIJA7_C string

	iztzjRXX := lsLFSQCl[0]
	lkgLredg := lsLFSQCl[1]

	cVwY5UTY, dTakwHTF :=  /*line XHiBoe4T.go:1*/ u3xzONWW.vz2OhzsG(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line axxtRFg_.go:1*/ shim.Error( /*line ZJryTsSB.go:1*/ dTakwHTF.Error())
	}
	_br98sk5, dTakwHTF :=  /*line xhwAAQSt.go:1*/ u3xzONWW.qbaWnNO0(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line xm4bZ01K.go:1*/ shim.Error( /*line NzSzdLhC.go:1*/ dTakwHTF.Error())
	}

	mBkwo0i6, dTakwHTF :=  /*line w8rUnszr.go:1*/ ipEgxdqE.GetState(iztzjRXX)
	if dTakwHTF != nil {
		ggIJA7_C =  /*line hFPG2BvC.go:1*/ func() string {
			var data []byte
			i := 10
			decryptKey := 70
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 8:
					data =  /*line hFPG2BvC.go:1*/ append(data, "\f\x04\xbe\x03"...,
					)
					i = 3
				case 7:
					i = 5
					data =  /*line hFPG2BvC.go:1*/ append(data, 9)
				case 9:
					i = 7
					data =  /*line hFPG2BvC.go:1*/ append(data, "\x00\xb0\xfe\xfb"...,
					)
				case 5:
					data =  /*line hFPG2BvC.go:1*/ append(data, "\xb4\x0e\x0e\xfa"...,
					)
					i = 8
				case 10:
					i = 6
					data =  /*line hFPG2BvC.go:1*/ append(data, "\xfe\xa4"...,
					)
				case 4:
					i = 9
					data =  /*line hFPG2BvC.go:1*/ append(data, "\xb3\x06"...,
					)
				case 11:
					data =  /*line hFPG2BvC.go:1*/ append(data, "\xf8\xfa\xf2\xf0"...,
					)
					i = 4
				case 6:
					i = 0
					data =  /*line hFPG2BvC.go:1*/ append(data, "\xc6\xf2"...,
					)
				case 3:
					i = 12
					data =  /*line hFPG2BvC.go:1*/ append(data, "\v\x15\xc2"...,
					)
				case 2:
					data =  /*line hFPG2BvC.go:1*/ append(data, "Ŭ\xcf\xe9"...,
					)
					i = 11
				case 12:
					for y := range data {
						data[y] = data[y] -  /*line hFPG2BvC.go:1*/ byte(decryptKey^y)
					}
					i = 1
				case 0:
					data =  /*line hFPG2BvC.go:1*/ append(data, "\xf9\xf5\xf7\xa6"...,
					)
					i = 2
				}
			}
			return  /*line hFPG2BvC.go:1*/ string(data)
		}() + iztzjRXX +  /*line i8PSVaR0.go:1*/ func() string {
			key :=  /*line i8PSVaR0.go:1*/ []byte("\x91\xdf")
			data :=  /*line i8PSVaR0.go:1*/ []byte("\xb3\\")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line i8PSVaR0.go:1*/ string(data)
		}()
		return  /*line HbYpqNR9.go:1*/ shim.Error(ggIJA7_C)
	} else if mBkwo0i6 == nil {
		 /*line JLqrxWFa.go:1*/ fmt.Println( /*line I4zlN_Nx.go:1*/ func() string {
			key :=  /*line I4zlN_Nx.go:1*/ []byte("<\x0fG\xf3\xcb\x00\xfaTF\xca\xdf\x04p^\x04\xa4\xb4\xc6\xc1\xf3G\xf2e\xc5\xc5\xc0l\xc0\xc4\t")
			data :=  /*line I4zlN_Nx.go:1*/ []byte("\x90w\xb0f\xebum\xb9\xb88@q\xd5~h\x13\x199\xe1a\xb6f\x85*=)\xdf4\xfe)")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line I4zlN_Nx.go:1*/ string(data)
		}() + iztzjRXX)
		return  /*line uIWN_mGK.go:1*/ shim.Error( /*line JYSK6gPh.go:1*/ func() string {
			seed :=  /*line JYSK6gPh.go:1*/ byte(244)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line JYSK6gPh.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/  /*line JYSK6gPh.go:1*/ fnc(160)(252)(249)(250)(163)(83)(10)(230)(27)(234)(15)(16)(232)(85)(174)(23)(234)(10)(163)(72)(1)(27)(170)(81)(253)(235)(30)(255)(176)(26)
			return  /*line JYSK6gPh.go:1*/ string(data)
		}() + iztzjRXX)
	}

	biXWDyeJ := UserCredentials{}
	dTakwHTF =  /*line NZQj8aCm.go:1*/ json.Unmarshal(mBkwo0i6, &biXWDyeJ)	                               
	if dTakwHTF != nil {
		return  /*line v343KyRw.go:1*/ shim.Error( /*line auiH1yRi.go:1*/ dTakwHTF.Error())
	}

	skhCclRZ, dTakwHTF :=  /*line EsZjfiK2.go:1*/ ipEgxdqE.GetState(biXWDyeJ.UserId)
	if dTakwHTF != nil {
		ggIJA7_C =  /*line ZvHpIMM_.go:1*/ func() string {
			fullData :=  /*line ZvHpIMM_.go:1*/ []byte("rh\xb8\x88\xb9\x10LrьX\xeaF\xac\xa3\xbb\x1f\xb5D\\\x00\xfe\xbc\xaf\xde\u007f\x13]\xe5(\x91pJ\u070f&W\xc6D@z|\xbec0}\xdb\x05=\xde\xea\x03\x84\xa8\x88\x94bм\x04\x85[\xe2\x83\x17H\xe5p")
			var data []byte
			data =  /*line ZvHpIMM_.go:1*/ append(data, fullData[35]^fullData[27], fullData[24]-fullData[58], fullData[21]^fullData[15], fullData[64]+fullData[61], fullData[20]^fullData[7], fullData[8]^fullData[42], fullData[5]^fullData[56], fullData[10]^fullData[40], fullData[34]^fullData[17], fullData[60]-fullData[43], fullData[2]-fullData[0], fullData[47]+fullData[19], fullData[11]^fullData[63], fullData[57]^fullData[22], fullData[39]-fullData[46], fullData[13]-fullData[65], fullData[38]+fullData[33], fullData[29]+fullData[6], fullData[41]^fullData[26], fullData[45]+fullData[14], fullData[49]^fullData[4], fullData[12]+fullData[16], fullData[18]^fullData[44], fullData[3]-fullData[1], fullData[67]^fullData[51], fullData[31]^fullData[59], fullData[25]+fullData[62], fullData[28]^fullData[30], fullData[32]-fullData[66], fullData[53]-fullData[54], fullData[50]-fullData[52], fullData[37]-fullData[36], fullData[23]-fullData[48], fullData[9]+fullData[55])
			return  /*line ZvHpIMM_.go:1*/ string(data)
		}() + biXWDyeJ.UserId +  /*line CARlmPAr.go:1*/ func() string {
			fullData :=  /*line CARlmPAr.go:1*/ []byte("Wu\x9d\xe0")
			var data []byte
			data =  /*line CARlmPAr.go:1*/ append(data, fullData[0]^fullData[1], fullData[3]+fullData[2])
			return  /*line CARlmPAr.go:1*/ string(data)
		}()
		return  /*line HNjdlqlG.go:1*/ shim.Error(ggIJA7_C)
	} else if skhCclRZ == nil {
		 /*line H7AKvgd3.go:1*/ fmt.Println( /*line WM9VDYZS.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 59
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 6:
					i = 7
					data =  /*line WM9VDYZS.go:1*/ append(data, "\xa4M"...,
					)
				case 4:
					i = 9
					data =  /*line WM9VDYZS.go:1*/ append(data, "_D"...,
					)
				case 1:
					i = 11
					data =  /*line WM9VDYZS.go:1*/ append(data, "V\xa6"...,
					)
				case 3:
					i = 5
					data =  /*line WM9VDYZS.go:1*/ append(data, "\x91\xa4\xa4"...,
					)
				case 0:
					data =  /*line WM9VDYZS.go:1*/ append(data, "\x91H\x94\x8e"...,
					)
					i = 4
				case 10:
					data =  /*line WM9VDYZS.go:1*/ append(data, "\xadU\xa2\xa6"...,
					)
					i = 1
				case 12:
					data =  /*line WM9VDYZS.go:1*/ append(data, "\x98\xa2"...,
					)
					i = 0
				case 5:
					data =  /*line WM9VDYZS.go:1*/ append(data, "\xb0\x9eX"...,
					)
					i = 8
				case 9:
					for y := range data {
						data[y] = data[y] -  /*line WM9VDYZS.go:1*/ byte(decryptKey^y)
					}
					i = 2
				case 8:
					i = 10
					data =  /*line WM9VDYZS.go:1*/ append(data, 164)
				case 7:
					i = 12
					data =  /*line WM9VDYZS.go:1*/ append(data, 163)
				case 11:
					data =  /*line WM9VDYZS.go:1*/ append(data, "\xa3\x98"...,
					)
					i = 6
				}
			}
			return  /*line WM9VDYZS.go:1*/ string(data)
		}() + biXWDyeJ.UserId)
		return  /*line BqASe5iG.go:1*/ shim.Error( /*line Z3G34w8p.go:1*/ func() string {
			data :=  /*line Z3G34w8p.go:1*/ []byte("T\x97e\xdeg \x8bs no\x96u\x84.\x99{wax) ed:\x9a")
			positions := [...]byte{18, 1, 6, 13, 15, 11, 25, 18, 14, 22, 20, 11, 3, 20, 15, 19, 25, 20, 22, 3, 22, 13, 3, 16, 18, 4, 1, 1}
			for i := 0; i < 28; i += 2 {
				localKey :=  /*line Z3G34w8p.go:1*/ byte(i) +  /*line Z3G34w8p.go:1*/ byte(positions[i]^positions[i+1]) + 224
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line Z3G34w8p.go:1*/ string(data)
		}() + biXWDyeJ.UserId)
	}

	jC3LR1TY := User{}
	dTakwHTF =  /*line rKgteAhZ.go:1*/ json.Unmarshal(skhCclRZ, &jC3LR1TY)	                               
	if dTakwHTF != nil {
		return  /*line AlPDLQPy.go:1*/ shim.Error( /*line PztsYF0m.go:1*/ dTakwHTF.Error())
	}

	jC3LR1TY.Status = lkgLredg

	lw5scISX, dTakwHTF :=  /*line cb4s2JhN.go:1*/ u3xzONWW.nyOeNqDh(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line XRN1hkJE.go:1*/ shim.Error( /*line LR3cx3hl.go:1*/ dTakwHTF.Error())
	}

	if cVwY5UTY ==  /*line HsSMqA0K.go:1*/ func() string {
		data :=  /*line HsSMqA0K.go:1*/ []byte("aA=\xd9n")
		positions := [...]byte{1, 2, 3, 2, 1, 1}
		for i := 0; i < 6; i += 2 {
			localKey :=  /*line HsSMqA0K.go:1*/ byte(i) +  /*line HsSMqA0K.go:1*/ byte(positions[i]^positions[i+1]) + 105
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line HsSMqA0K.go:1*/ string(data)
	}() {
		if _br98sk5 != biXWDyeJ.Organization || jC3LR1TY.ExternalOrg != lw5scISX {
			return  /*line d2c1YLAb.go:1*/ shim.Error( /*line mXgeEMKN.go:1*/ func() string {
				seed :=  /*line mXgeEMKN.go:1*/ byte(222)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data =  /*line mXgeEMKN.go:1*/ append(data, x+seed); seed += x; return fnc }
				 /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/  /*line mXgeEMKN.go:1*/ fnc(118)(20)(253)(187)(79)(3)(245)(185)(65)(3)(9)(252)(5)(178)(67)(254)(13)(178)(83)(242)(15)(172)(65)(2)(17)(245)(13)(239)(187)(65)(2)(0)(12)(6)(249)(6)(255)(173)(79)(255)(254)(13)(167)(70)(9)(3)(174)(84)(244)(253)(187)(79)(8)(247)(247)(255)(188)(79)(3)(245)(250)(13)(251)(17)(231)(19)(245)(6)(255)(5)
				return  /*line mXgeEMKN.go:1*/ string(data)
			}())
		}

		if biXWDyeJ.Type !=  /*line eH3eZA7s.go:1*/ func() string {
			seed :=  /*line eH3eZA7s.go:1*/ byte(16)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line eH3eZA7s.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line eH3eZA7s.go:1*/  /*line eH3eZA7s.go:1*/  /*line eH3eZA7s.go:1*/  /*line eH3eZA7s.go:1*/ fnc(101)(254)(242)(13)
			return  /*line eH3eZA7s.go:1*/ string(data)
		}() {
			return  /*line GSdmrSao.go:1*/ shim.Error( /*line hGZH8ym4.go:1*/ func() string {
				var data []byte
				i := 4
				decryptKey := 106
				for counter := 0; i != 0; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 3:
						i = 7
						data =  /*line hGZH8ym4.go:1*/ append(data, "\x1c\xd6 "...,
						)
					case 16:
						data =  /*line hGZH8ym4.go:1*/ append(data, "\xcf\x1d"...,
						)
						i = 18
					case 18:
						i = 15
						data =  /*line hGZH8ym4.go:1*/ append(data, "\x17\x14$"...,
						)
					case 1:
						data =  /*line hGZH8ym4.go:1*/ append(data, "\t\x15\t\x19"...,
						)
						i = 10
					case 7:
						i = 9
						data =  /*line hGZH8ym4.go:1*/ append(data, "\"\x1a\xd2\x1e"...,
						)
					case 11:
						data =  /*line hGZH8ym4.go:1*/ append(data, "''\xd8\x1e"...,
						)
						i = 14
					case 2:
						i = 17
						data =  /*line hGZH8ym4.go:1*/ append(data, "\f\x11\r"...,
						)
					case 6:
						data =  /*line hGZH8ym4.go:1*/ append(data, "\xf1\xf6\xf5"...,
						)
						i = 2
					case 15:
						i = 8
						data =  /*line hGZH8ym4.go:1*/ append(data, "\xca\n"...,
						)
					case 10:
						i = 16
						data =  /*line hGZH8ym4.go:1*/ append(data, "\x03!\x11"...,
						)
					case 9:
						data =  /*line hGZH8ym4.go:1*/ append(data, " ,"...,
						)
						i = 11
					case 8:
						data =  /*line hGZH8ym4.go:1*/ append(data, 7)
						i = 13
					case 5:
						for y := range data {
							data[y] = data[y] -  /*line hGZH8ym4.go:1*/ byte(decryptKey^y)
						}
						i = 0
					case 4:
						i = 3
						data =  /*line hGZH8ym4.go:1*/ append(data, "\t\x1c"...,
						)
					case 14:
						data =  /*line hGZH8ym4.go:1*/ append(data, "\x1b\x13\xc4\b"...,
						)
						i = 1
					case 13:
						data =  /*line hGZH8ym4.go:1*/ append(data, 252)
						i = 12
					case 12:
						data =  /*line hGZH8ym4.go:1*/ append(data, "\b\xb1"...,
						)
						i = 6
					case 17:
						data =  /*line hGZH8ym4.go:1*/ append(data, 18)
						i = 5
					}
				}
				return  /*line hGZH8ym4.go:1*/ string(data)
			}())
		}
	} else if cVwY5UTY ==  /*line BHRKEEX6.go:1*/ func() string {
		data :=  /*line BHRKEEX6.go:1*/ []byte("\xe2\xf1\xf2\xf2")
		positions := [...]byte{0, 2, 1, 3}
		for i := 0; i < 4; i += 2 {
			localKey :=  /*line BHRKEEX6.go:1*/ byte(i) +  /*line BHRKEEX6.go:1*/ byte(positions[i]^positions[i+1]) + 123
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line BHRKEEX6.go:1*/ string(data)
	}() {
		return  /*line AonP3JpS.go:1*/ shim.Error( /*line S4I_GQ4l.go:1*/ func() string {
			fullData :=  /*line S4I_GQ4l.go:1*/ []byte("W0\u007fU.\xe5\x053ʅĢ\xf6nY_9\x8a\x18\x88~|\xab|\xe1\xbc\x1c\x9dK\xed\xde\xdd\x1b\xc8elv\nu\xd9\xcc-B=\x97\x80o\x9f\xca\xc2\x00\xa7\xeaK\xa08\bߦ\xa4\xb42\xa7\x02\x9c\x14")
			var data []byte
			data =  /*line S4I_GQ4l.go:1*/ append(data, fullData[26]+fullData[7], fullData[54]-fullData[61], fullData[4]^fullData[42], fullData[62]^fullData[30], fullData[14]-fullData[16], fullData[27]+fullData[10], fullData[49]^fullData[58], fullData[1]+fullData[43], fullData[15]+fullData[37], fullData[59]+fullData[8], fullData[22]+fullData[33], fullData[47]-fullData[2], fullData[41]-fullData[48], fullData[31]-fullData[23], fullData[50]+fullData[13], fullData[35]+fullData[60], fullData[53]-fullData[52], fullData[57]-fullData[21], fullData[36]^fullData[63], fullData[6]-fullData[64], fullData[20]^fullData[56], fullData[12]^fullData[44], fullData[46]^fullData[32], fullData[5]^fullData[45], fullData[9]-fullData[34], fullData[38]^fullData[65], fullData[17]+fullData[39], fullData[51]+fullData[25], fullData[0]+fullData[18], fullData[19]+fullData[29], fullData[11]^fullData[40], fullData[3]-fullData[24], fullData[55]^fullData[28])
			return  /*line S4I_GQ4l.go:1*/ string(data)
		}())
	} else if cVwY5UTY ==  /*line WU7zafsk.go:1*/ func() string {
		var data []byte
		i := 3
		decryptKey := 254
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 8:
				i = 4
				data =  /*line WU7zafsk.go:1*/ append(data, "DP"...,
				)
			case 5:
				data =  /*line WU7zafsk.go:1*/ append(data, 94)
				i = 6
			case 6:
				i = 8
				data =  /*line WU7zafsk.go:1*/ append(data, "Z\x1cB"...,
				)
			case 3:
				data =  /*line WU7zafsk.go:1*/ append(data, "WM_"...,
				)
				i = 1
			case 1:
				i = 0
				data =  /*line WU7zafsk.go:1*/ append(data, 97)
			case 4:
				i = 7
				data =  /*line WU7zafsk.go:1*/ append(data, "KS"...,
				)
			case 7:
				i = 2
				for y := range data {
					data[y] = data[y] -  /*line WU7zafsk.go:1*/ byte(decryptKey^y)
				}
			case 0:
				data =  /*line WU7zafsk.go:1*/ append(data, 92)
				i = 5
			}
		}
		return  /*line WU7zafsk.go:1*/ string(data)
	}() && _br98sk5 != biXWDyeJ.Organization {

		return  /*line akiXh3vz.go:1*/ shim.Error( /*line crST9FXy.go:1*/ func() string {
			fullData :=  /*line crST9FXy.go:1*/ []byte("\xa9r\xe7|\xb9\x8e\x05\u007fu\xc2_\xae\u007fe\rl\x9f\x13l䥽,\"\xbb\xf8\xf1}m}@\x95\xb1\x10\xfa'n\x1f\xfe\x8d\t\x03\x9b\x80\x93\x1c\x06BM\x10\xbb\xfa\xccp\xec\u007f\xdd6G\xad\xd0\xe1\xede~\x06=\xd5\xffY\xf5\x9f\xa6Q\x05\x17\xe4F\x04T=&\xd92N\xbdv\x9e\xf0\x12\x94:\xab|\xd5\xfd\xcb\xfe\xc4\xeb\x88\u05ff\x93\xe4\x1a\xc6\xc4\x15\x84\x9b.\xcfV\xb1E\xb6%c\b\xd7\xf9{\f\xad6\xfdL\xa0\xf2@L\x01%\x9aeqQ)\xa8k\x8fK^p\x1e\x8dW\xf4\x0f")
			var data []byte
			data =  /*line crST9FXy.go:1*/ append(data, fullData[78]-fullData[116], fullData[26]^fullData[90], fullData[91]-fullData[106], fullData[62]^fullData[134], fullData[8]+fullData[51], fullData[81]+fullData[127], fullData[9]+fullData[0], fullData[110]^fullData[24], fullData[118]+fullData[38], fullData[37]+fullData[115], fullData[17]-fullData[72], fullData[75]^fullData[64], fullData[20]^fullData[96], fullData[86]+fullData[95], fullData[73]+fullData[112], fullData[12]+fullData[104], fullData[41]+fullData[143], fullData[22]+fullData[47], fullData[105]+fullData[46], fullData[149]-fullData[92], fullData[131]-fullData[2], fullData[133]-fullData[107], fullData[32]-fullData[84], fullData[31]^fullData[61], fullData[142]^fullData[23], fullData[120]+fullData[16], fullData[11]-fullData[48], fullData[76]-fullData[53], fullData[97]^fullData[42], fullData[5]-fullData[36], fullData[89]-fullData[114], fullData[147]+fullData[123], fullData[144]-fullData[14], fullData[57]^fullData[69], fullData[101]+fullData[87], fullData[135]+fullData[40], fullData[59]^fullData[82], fullData[125]+fullData[80], fullData[55]-fullData[10], fullData[29]+fullData[129], fullData[126]^fullData[103], fullData[94]^fullData[4], fullData[1]-fullData[121], fullData[56]-fullData[85], fullData[39]-fullData[35], fullData[145]+fullData[137], fullData[15]+fullData[65], fullData[49]+fullData[33], fullData[70]+fullData[7], fullData[54]+fullData[93], fullData[130]+fullData[117], fullData[119]-fullData[71], fullData[34]-fullData[100], fullData[63]+fullData[50], fullData[79]^fullData[66], fullData[30]^fullData[111], fullData[25]-fullData[109], fullData[21]+fullData[139], fullData[136]-fullData[68], fullData[140]^fullData[74], fullData[88]-fullData[141], fullData[28]-fullData[132], fullData[43]^fullData[128], fullData[122]+fullData[148], fullData[6]-fullData[44], fullData[99]+fullData[3], fullData[52]^fullData[124], fullData[98]-fullData[113], fullData[67]-fullData[18], fullData[13]+fullData[108], fullData[27]^fullData[45], fullData[77]^fullData[83], fullData[19]^fullData[146], fullData[102]^fullData[60], fullData[138]^fullData[58])
			return  /*line crST9FXy.go:1*/ string(data)
		}())

	}

	nQa48sqX, _ :=  /*line l80viNv8.go:1*/ json.Marshal(jC3LR1TY)
	dTakwHTF =  /*line Fj_cwaGZ.go:1*/ ipEgxdqE.PutState(biXWDyeJ.UserId, nQa48sqX)	                       
	if dTakwHTF != nil {
		return  /*line comZk4Qf.go:1*/ shim.Error( /*line mjAjw6Rh.go:1*/ dTakwHTF.Error())
	}

	gd7Gzym9, _ :=  /*line EBGNu3Sz.go:1*/ json.Marshal(biXWDyeJ)
	dTakwHTF =  /*line dlWoLwAe.go:1*/ ipEgxdqE.PutState(iztzjRXX, gd7Gzym9)	                             
	if dTakwHTF != nil {
		return  /*line ZsFTrBvM.go:1*/ shim.Error( /*line HzZj7ERu.go:1*/ dTakwHTF.Error())
	}

	return  /*line Wiokj3VB.go:1*/ shim.Success(nil)
}

                
func (u3xzONWW *ZohimWzR) iiQph8om(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	var dTakwHTF error
	var ggIJA7_C string

	iztzjRXX := lsLFSQCl[0]
	eXE7UdH0 := lsLFSQCl[1]

	cVwY5UTY, dTakwHTF :=  /*line KyvdfBoP.go:1*/ u3xzONWW.vz2OhzsG(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line d2_czQ5G.go:1*/ shim.Error( /*line D6BCTw9B.go:1*/ dTakwHTF.Error())
	}

	if cVwY5UTY ==  /*line XPjX6p1d.go:1*/ func() string {
		key :=  /*line XPjX6p1d.go:1*/ []byte("\xf1\xd9\xe1;")
		data :=  /*line XPjX6p1d.go:1*/ []byte("\x84\xaa\x84I")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line XPjX6p1d.go:1*/ string(data)
	}() {
		return  /*line pSKlcu6h.go:1*/ shim.Error( /*line lGpdnIoF.go:1*/ func() string {
			var data []byte
			i := 0
			decryptKey := 58
			for counter := 0; i != 7; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 13:
					i = 11
					data =  /*line lGpdnIoF.go:1*/ append(data, "\x9b\x9d\x8e"...,
					)
				case 1:
					i = 16
					data =  /*line lGpdnIoF.go:1*/ append(data, "m\x8c\x8e"...,
					)
				case 3:
					data =  /*line lGpdnIoF.go:1*/ append(data, "\x89{"...,
					)
					i = 8
				case 11:
					i = 7
					for y := range data {
						data[y] = data[y] +  /*line lGpdnIoF.go:1*/ byte(decryptKey^y)
					}
				case 9:
					data =  /*line lGpdnIoF.go:1*/ append(data, ";\x8e\x888"...,
					)
					i = 3
				case 12:
					data =  /*line lGpdnIoF.go:1*/ append(data, "tx#c"...,
					)
					i = 1
				case 15:
					i = 2
					data =  /*line lGpdnIoF.go:1*/ append(data, "\x81\x91"...,
					)
				case 16:
					i = 4
					data =  /*line lGpdnIoF.go:1*/ append(data, 149)
				case 10:
					i = 14
					data =  /*line lGpdnIoF.go:1*/ append(data, "\u007f\u007f+k"...,
					)
				case 8:
					data =  /*line lGpdnIoF.go:1*/ append(data, "\x88y\x872"...,
					)
					i = 15
				case 14:
					data =  /*line lGpdnIoF.go:1*/ append(data, 123)
					i = 5
				case 5:
					i = 12
					data =  /*line lGpdnIoF.go:1*/ append(data, "m't"...,
					)
				case 4:
					i = 6
					data =  /*line lGpdnIoF.go:1*/ append(data, 130)
				case 2:
					i = 13
					data =  /*line lGpdnIoF.go:1*/ append(data, "\xa2\xa1\xa4"...,
					)
				case 0:
					data =  /*line lGpdnIoF.go:1*/ append(data, "e\x82s"...,
					)
					i = 10
				case 6:
					i = 9
					data =  /*line lGpdnIoF.go:1*/ append(data, 128)
				}
			}
			return  /*line lGpdnIoF.go:1*/ string(data)
		}())
	}

	eiKUiRpK, dTakwHTF :=  /*line nvBK81Os.go:1*/ u3xzONWW.rjGOlCPK(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line QFc3ID_q.go:1*/ shim.Error( /*line izVdwMcJ.go:1*/ dTakwHTF.Error())
	}

	mBkwo0i6, dTakwHTF :=  /*line ynqaHO0_.go:1*/ ipEgxdqE.GetState(iztzjRXX)
	if dTakwHTF != nil {
		ggIJA7_C =  /*line KrfVIRPz.go:1*/ func() string {
			fullData :=  /*line KrfVIRPz.go:1*/ []byte("ϭ\v$\\\x03[\x94?\xef-\t{.A\xf4\xbeO\xea\x8b\xe3JVl;\xaa/\xb2=j,\U0007538a\u007f\t\x05\x99۽\xbd\x1f}\x9e\xfe\xe5\x97\xcb*\xf4\x1f\x915\xa0\u007f''\xe6\x1b\xdb=`\xae\xb9\xf3C\xb6")
			var data []byte
			data =  /*line KrfVIRPz.go:1*/ append(data, fullData[7]^fullData[9], fullData[46]+fullData[61], fullData[59]+fullData[49], fullData[42]-fullData[1], fullData[28]^fullData[17], fullData[54]+fullData[0], fullData[12]^fullData[36], fullData[45]+fullData[3], fullData[13]-fullData[50], fullData[32]^fullData[47], fullData[51]+fullData[56], fullData[2]-fullData[25], fullData[57]-fullData[16], fullData[66]^fullData[26], fullData[60]+fullData[34], fullData[38]+fullData[48], fullData[65]+fullData[10], fullData[43]^fullData[11], fullData[64]+fullData[67], fullData[21]^fullData[29], fullData[40]-fullData[22], fullData[37]+fullData[62], fullData[52]+fullData[20], fullData[63]^fullData[33], fullData[27]-fullData[8], fullData[14]^fullData[53], fullData[41]-fullData[4], fullData[44]^fullData[18], fullData[58]+fullData[35], fullData[6]-fullData[24], fullData[39]+fullData[19], fullData[23]+fullData[5], fullData[31]-fullData[55], fullData[15]+fullData[30])
			return  /*line KrfVIRPz.go:1*/ string(data)
		}() + iztzjRXX +  /*line KVRGEvEJ.go:1*/ func() string {
			fullData :=  /*line KVRGEvEJ.go:1*/ []byte("\x9d\xe0\xee\x10")
			var data []byte
			data =  /*line KVRGEvEJ.go:1*/ append(data, fullData[3]-fullData[2], fullData[1]^fullData[0])
			return  /*line KVRGEvEJ.go:1*/ string(data)
		}()
		return  /*line IlcbQBcS.go:1*/ shim.Error(ggIJA7_C)
	} else if mBkwo0i6 == nil {
		 /*line Sywigzqw.go:1*/ fmt.Println( /*line EH63c711.go:1*/ func() string {
			var data []byte
			i := 7
			decryptKey := 39
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					i = 1
					data =  /*line EH63c711.go:1*/ append(data, "\xbc\xc8"...,
					)
				case 1:
					i = 8
					data =  /*line EH63c711.go:1*/ append(data, "\xab\xbad"...,
					)
				case 3:
					for y := range data {
						data[y] = data[y] -  /*line EH63c711.go:1*/ byte(decryptKey^y)
					}
					i = 0
				case 8:
					data =  /*line EH63c711.go:1*/ append(data, "\xb3\xb1\xb7`"...,
					)
					i = 4
				case 9:
					i = 5
					data =  /*line EH63c711.go:1*/ append(data, "\xb6\xd0\xcd"...,
					)
				case 10:
					i = 3
					data =  /*line EH63c711.go:1*/ append(data, "\xc1\x84k"...,
					)
				case 6:
					data =  /*line EH63c711.go:1*/ append(data, "\xc8r\xc8\xc3"...,
					)
					i = 9
				case 4:
					data =  /*line EH63c711.go:1*/ append(data, "\xa6Ƹ\xbf"...,
					)
					i = 10
				case 7:
					data =  /*line EH63c711.go:1*/ append(data, "\xaa\xbf\xbd"...,
					)
					i = 6
				case 5:
					data =  /*line EH63c711.go:1*/ append(data, "\xbdʿ{"...,
					)
					i = 2
				}
			}
			return  /*line EH63c711.go:1*/ string(data)
		}() + iztzjRXX)
		return  /*line EncuM2Fp.go:1*/ shim.Error( /*line umAp9myg.go:1*/ func() string {
			var data []byte
			i := 12
			decryptKey := 232
			for counter := 0; i != 6; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 11:
					i = 13
					data =  /*line umAp9myg.go:1*/ append(data, 120)
				case 4:
					data =  /*line umAp9myg.go:1*/ append(data, 126)
					i = 10
				case 5:
					i = 0
					data =  /*line umAp9myg.go:1*/ append(data, 121)
				case 13:
					i = 3
					data =  /*line umAp9myg.go:1*/ append(data, "\x82s\x80"...,
					)
				case 14:
					data =  /*line umAp9myg.go:1*/ append(data, "\x89\x92"...,
					)
					i = 1
				case 2:
					i = 15
					data =  /*line umAp9myg.go:1*/ append(data, 131)
				case 9:
					i = 6
					for y := range data {
						data[y] = data[y] +  /*line umAp9myg.go:1*/ byte(decryptKey^y)
					}
				case 0:
					i = 4
					data =  /*line umAp9myg.go:1*/ append(data, "},p"...,
					)
				case 1:
					data =  /*line umAp9myg.go:1*/ append(data, 58)
					i = 8
				case 8:
					i = 2
					data =  /*line umAp9myg.go:1*/ append(data, "\x8e\x8f\x80\x88"...,
					)
				case 10:
					data =  /*line umAp9myg.go:1*/ append(data, "n{"...,
					)
					i = 7
				case 15:
					i = 11
					data =  /*line umAp9myg.go:1*/ append(data, "y\x84w1"...,
					)
				case 7:
					data =  /*line umAp9myg.go:1*/ append(data, "{<!"...,
					)
					i = 9
				case 3:
					i = 5
					data =  /*line umAp9myg.go:1*/ append(data, "0}"...,
					)
				case 12:
					data =  /*line umAp9myg.go:1*/ append(data, "r\x85"...,
					)
					i = 14
				}
			}
			return  /*line umAp9myg.go:1*/ string(data)
		}() + iztzjRXX)
	}

	biXWDyeJ := UserCredentials{}
	dTakwHTF =  /*line r8RkBD87.go:1*/ json.Unmarshal(mBkwo0i6, &biXWDyeJ)	                               
	if dTakwHTF != nil {
		return  /*line nP1YsYG2.go:1*/ shim.Error( /*line vniZfevG.go:1*/ dTakwHTF.Error())
	}
	iLqMgi9G :=  /*line iz3pr31r.go:1*/ eCTdxTGW( /*line vfjze6Ur.go:1*/ string(eXE7UdH0))

	biXWDyeJ.Password = iLqMgi9G
	biXWDyeJ.Created_by = eiKUiRpK

	dS0w8e4b, dTakwHTF :=  /*line FXMpMaNz.go:1*/ u3xzONWW.qbaWnNO0(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line rjwRVPUZ.go:1*/ shim.Error( /*line zrpsGfzM.go:1*/ dTakwHTF.Error())
	}

	if (biXWDyeJ.Organization != dS0w8e4b) && (cVwY5UTY ==  /*line ihUrCwat.go:1*/ func() string {
		seed :=  /*line ihUrCwat.go:1*/ byte(92)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line ihUrCwat.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line ihUrCwat.go:1*/  /*line ihUrCwat.go:1*/  /*line ihUrCwat.go:1*/  /*line ihUrCwat.go:1*/  /*line ihUrCwat.go:1*/ fnc(5)(3)(9)(252)(5)
		return  /*line ihUrCwat.go:1*/ string(data)
	}()) {
		return  /*line zzQ04PVj.go:1*/ shim.Error( /*line cokUCDxT.go:1*/ func() string {
			var data []byte
			i := 6
			decryptKey := 16
			for counter := 0; i != 12; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 16:
					data =  /*line cokUCDxT.go:1*/ append(data, 19)
					i = 1
				case 2:
					i = 0
					data =  /*line cokUCDxT.go:1*/ append(data, "hgne"...,
					)
				case 3:
					i = 11
					data =  /*line cokUCDxT.go:1*/ append(data, "B3=\xe8"...,
					)
				case 13:
					i = 19
					data =  /*line cokUCDxT.go:1*/ append(data, "F4;"...,
					)
				case 18:
					data =  /*line cokUCDxT.go:1*/ append(data, "24@;"...,
					)
					i = 8
				case 19:
					i = 5
					data =  /*line cokUCDxT.go:1*/ append(data, ",*"...,
					)
				case 10:
					i = 14
					data =  /*line cokUCDxT.go:1*/ append(data, 98)
				case 7:
					data =  /*line cokUCDxT.go:1*/ append(data, "Q\xf99G"...,
					)
					i = 13
				case 14:
					i = 9
					data =  /*line cokUCDxT.go:1*/ append(data, "JV"...,
					)
				case 20:
					i = 12
					for y := range data {
						data[y] = data[y] -  /*line cokUCDxT.go:1*/ byte(decryptKey^y)
					}
				case 9:
					data =  /*line cokUCDxT.go:1*/ append(data, "\a[T"...,
					)
					i = 21
				case 8:
					i = 22
					data =  /*line cokUCDxT.go:1*/ append(data, "K\xfc"...,
					)
				case 22:
					i = 7
					data =  /*line cokUCDxT.go:1*/ append(data, 72)
				case 6:
					data =  /*line cokUCDxT.go:1*/ append(data, "$F>\xf6"...,
					)
					i = 18
				case 1:
					i = 17
					data =  /*line cokUCDxT.go:1*/ append(data, "Xln\x1f"...,
					)
				case 21:
					data =  /*line cokUCDxT.go:1*/ append(data, "EU"...,
					)
					i = 15
				case 15:
					i = 20
					data =  /*line cokUCDxT.go:1*/ append(data, 85)
				case 17:
					i = 10
					data =  /*line cokUCDxT.go:1*/ append(data, "fbk*"...,
					)
				case 4:
					data =  /*line cokUCDxT.go:1*/ append(data, "?1"...,
					)
					i = 3
				case 0:
					i = 16
					data =  /*line cokUCDxT.go:1*/ append(data, "cT"...,
					)
				case 5:
					i = 4
					data =  /*line cokUCDxT.go:1*/ append(data, "\xe142\xe2"...,
					)
				case 11:
					data =  /*line cokUCDxT.go:1*/ append(data, ";+"...,
					)
					i = 2
				}
			}
			return  /*line cokUCDxT.go:1*/ string(data)
		}())
	}

	gd7Gzym9, _ :=  /*line BtzYoGuG.go:1*/ json.Marshal(biXWDyeJ)
	dTakwHTF =  /*line AnsH6PMh.go:1*/ ipEgxdqE.PutState(iztzjRXX, gd7Gzym9)	                             
	if dTakwHTF != nil {
		return  /*line G0Zlhuit.go:1*/ shim.Error( /*line UVOpSDNE.go:1*/ dTakwHTF.Error())
	}

	oEXQD2hD :=  /*line q85V2zCy.go:1*/ func() string {
		data :=  /*line q85V2zCy.go:1*/ []byte("\xe6\xc9\xce\x1d\xdezK\xfd\xcdh-sͣl\xb4\x96\xcfge\xd3ase\xda")
		positions := [...]byte{7, 6, 13, 13, 14, 20, 16, 8, 16, 17, 1, 5, 19, 18, 10, 7, 7, 17, 6, 4, 2, 1, 12, 13, 22, 18, 20, 24, 4, 3, 18, 19, 15, 4, 0, 12}
		for i := 0; i < 36; i += 2 {
			localKey :=  /*line q85V2zCy.go:1*/ byte(i) +  /*line q85V2zCy.go:1*/ byte(positions[i]^positions[i+1]) + 152
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line q85V2zCy.go:1*/ string(data)
	}()
	iZLjZ2UK :=  /*line H4OhlXOF.go:1*/ []byte(oEXQD2hD)
	return  /*line W6uu0Kiz.go:1*/ shim.Success(iZLjZ2UK)
}

               
func (u3xzONWW *ZohimWzR) dnDmY7x_(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	var ggIJA7_C string
	if  /*line pzQXbfWs.go:1*/ len(lsLFSQCl) < 2 {
		return  /*line uxomfRUB.go:1*/ shim.Error( /*line E4d0eG2n.go:1*/ func() string {
			data :=  /*line E4d0eG2n.go:1*/ []byte("\xfdn\xe1\xacrh\xa793\x10z\xfemM\x06r \rf\xa0\x14\x92\x0fH*\fn,s.\xfeE\xb4\xe59*\xc0\x93\xd7g\x97Z")
			positions := [...]byte{22, 21, 23, 23, 32, 34, 19, 10, 37, 2, 27, 8, 41, 30, 35, 9, 41, 9, 3, 13, 24, 21, 2, 38, 14, 30, 7, 2, 30, 14, 40, 6, 13, 33, 27, 5, 20, 6, 14, 27, 11, 38, 36, 0, 25, 11, 21, 10, 5, 17, 22, 9, 7, 37, 3, 7, 10, 34}
			for i := 0; i < 58; i += 2 {
				localKey :=  /*line E4d0eG2n.go:1*/ byte(i) +  /*line E4d0eG2n.go:1*/ byte(positions[i]^positions[i+1]) + 59
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line E4d0eG2n.go:1*/ string(data)
		}())
	}

	iztzjRXX := lsLFSQCl[0]
	eXE7UdH0 := lsLFSQCl[1]

	cVwY5UTY, dTakwHTF :=  /*line VgSzGTCF.go:1*/ u3xzONWW.vz2OhzsG(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line QoJjhiUW.go:1*/ shim.Error( /*line I69x4A37.go:1*/ dTakwHTF.Error())
	}

	               
	q3K86eOa :=  /*line GOjY3gN8.go:1*/ eCTdxTGW( /*line L_TF37cp.go:1*/ string(eXE7UdH0))
	lW3UMwus := iztzjRXX
	                
	qK0_FqHv, dTakwHTF :=  /*line CNZyl4Yu.go:1*/ ipEgxdqE.GetState(lW3UMwus)
	if dTakwHTF != nil {
		return  /*line vaO7UjHq.go:1*/ shim.Error( /*line UlSn4MAw.go:1*/ func() string {
			key :=  /*line UlSn4MAw.go:1*/ []byte("\xd9\x05ޛ\xe2\xe6\\\x99E\xbcr枮")
			data :=  /*line UlSn4MAw.go:1*/ []byte(".xC\r\x02T\xcb\re\"\xe1[\f\x12")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line UlSn4MAw.go:1*/ string(data)
		}())
	}
	                   
	var pxcCIdMn UserCredentials
	dTakwHTF =  /*line GXmvfDY6.go:1*/ json.Unmarshal(qK0_FqHv, &pxcCIdMn)
	if dTakwHTF != nil {
		ggIJA7_C =  /*line LNRmrKsZ.go:1*/ func() string {
			key :=  /*line LNRmrKsZ.go:1*/ []byte("\x12\t\x89\xeb&lu\x17{v?\x12\xd4\xcf\x05q\xdbB\x88\x18\"d\xa9\xf8\x12\xb9\xc9vG\x04\x0fq#\x8f\xbd\\\xb6\xe8\xe3a?~\x85\xaf")
			data :=  /*line LNRmrKsZ.go:1*/ []byte("i\x19\xbc\x87L\x03\xfd\v\xbf\xac\x0f]L\xa3`\xf6\x8e1\xecMP\x01\xbb(c\xba\x9c\xfc\xd9sZ\x03E\x91\xb8\x17\xaf\x8a\x8b\x00.\xe7\xb5q")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line LNRmrKsZ.go:1*/ string(data)
		}() + iztzjRXX +  /*line Qy7V6qe5.go:1*/ func() string {
			seed :=  /*line Qy7V6qe5.go:1*/ byte(80)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line Qy7V6qe5.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line Qy7V6qe5.go:1*/  /*line Qy7V6qe5.go:1*/ fnc(210)(91)
			return  /*line Qy7V6qe5.go:1*/ string(data)
		}()
		return  /*line eGKSAbQa.go:1*/ shim.Error(ggIJA7_C)
	}

	iARVJR0B, dTakwHTF :=  /*line NkeXZUI8.go:1*/ ipEgxdqE.GetState(pxcCIdMn.Organization)
	if dTakwHTF != nil {
		return  /*line MwjrfCPV.go:1*/ shim.Error( /*line yA9_TLpp.go:1*/ func() string {
			data :=  /*line yA9_TLpp.go:1*/ []byte("Orӳn\xb8za\xad\x15\xe5\xd6ƹ\xa2l\xc0\x1bou3\xe3")
			positions := [...]byte{11, 12, 20, 17, 9, 2, 12, 17, 13, 20, 16, 15, 2, 14, 9, 11, 21, 21, 10, 14, 8, 5, 21, 3, 13, 14, 12, 20}
			for i := 0; i < 28; i += 2 {
				localKey :=  /*line yA9_TLpp.go:1*/ byte(i) +  /*line yA9_TLpp.go:1*/ byte(positions[i]^positions[i+1]) + 35
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line yA9_TLpp.go:1*/ string(data)
		}())
	}
	                   
	var eqLGKena Org
	dTakwHTF =  /*line BUJcFzfM.go:1*/ json.Unmarshal(iARVJR0B, &eqLGKena)
	if dTakwHTF != nil {
		ggIJA7_C =  /*line mV45xWn9.go:1*/ func() string {
			key :=  /*line mV45xWn9.go:1*/ []byte("\u007f\x9b\x05@s\\\x1d\xe9\x9dN\xa3L\xdb^8j*\x9c\xb2n\xd3\x1ds\x04\xb6\xb2\x9d\n\x1a\n\\\xdav\xd0\xed\xa5\xf2b8")
			data :=  /*line mV45xWn9.go:1*/ []byte("\x04\xb9@2\x013o˧l\xed#\xaf~J\x0fM\xf5\xc1\x1a\xb6o\x16`\x96\xdd\xefm{d5\xa0\x17\xa4\x84ʜX\x18")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line mV45xWn9.go:1*/ string(data)
		}() + pxcCIdMn.Organization +  /*line F_nmb45y.go:1*/ func() string {
			data :=  /*line F_nmb45y.go:1*/ []byte("G\xec")
			positions := [...]byte{0, 1}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line F_nmb45y.go:1*/ byte(i) +  /*line F_nmb45y.go:1*/ byte(positions[i]^positions[i+1]) + 201
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line F_nmb45y.go:1*/ string(data)
		}()
		return  /*line WPJKKxsb.go:1*/ shim.Error(ggIJA7_C)
	}

	jZBkzcx4, dTakwHTF :=  /*line DHeiBezs.go:1*/ u3xzONWW.mRLMxx35(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line rkWRqC1O.go:1*/ shim.Error( /*line X5oDVbUR.go:1*/ dTakwHTF.Error())
	}
	lw5scISX, dTakwHTF :=  /*line y3cICGzB.go:1*/ u3xzONWW.nyOeNqDh(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line baQIBZQh.go:1*/ shim.Error( /*line bsDFbZJS.go:1*/ dTakwHTF.Error())
	}

	                          
	var ecerSq1k string
	                             
	if (iztzjRXX == pxcCIdMn.Username) && (q3K86eOa == pxcCIdMn.Password) && (pxcCIdMn.Username != pxcCIdMn.Created_by) {
		ecerSq1k =  /*line Dns18BxH.go:1*/ fmt.Sprintf( /*line bwnjkObI.go:1*/ func() string {
			var data []byte
			i := 35
			decryptKey := 54
			for counter := 0; i != 7; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 27:
					i = 19
					data =  /*line bwnjkObI.go:1*/ append(data, "\x99{\xa1"...,
					)
				case 18:
					data =  /*line bwnjkObI.go:1*/ append(data, 70)
					i = 3
				case 30:
					i = 12
					data =  /*line bwnjkObI.go:1*/ append(data, 110)
				case 1:
					i = 26
					data =  /*line bwnjkObI.go:1*/ append(data, "\xa7\x99\xa7\x8f"...,
					)
				case 32:
					i = 14
					data =  /*line bwnjkObI.go:1*/ append(data, "\xde\xe2\xe2\xd5"...,
					)
				case 24:
					data =  /*line bwnjkObI.go:1*/ append(data, "\xd7\xe5\xd9\xe2"...,
					)
					i = 43
				case 29:
					data =  /*line bwnjkObI.go:1*/ append(data, "\xd3ӆ\xa1"...,
					)
					i = 40
				case 35:
					i = 48
					data =  /*line bwnjkObI.go:1*/ append(data, "\xc8n\xbb"...,
					)
				case 12:
					data =  /*line bwnjkObI.go:1*/ append(data, "}~\xd1\xcd"...,
					)
					i = 20
				case 28:
					i = 4
					data =  /*line bwnjkObI.go:1*/ append(data, "So"...,
					)
				case 44:
					i = 8
					data =  /*line bwnjkObI.go:1*/ append(data, 143)
				case 46:
					i = 33
					data =  /*line bwnjkObI.go:1*/ append(data, "\xe8\xea\xe2\xdb"...,
					)
				case 21:
					i = 6
					data =  /*line bwnjkObI.go:1*/ append(data, "\x99\xa9"...,
					)
				case 22:
					i = 27
					data =  /*line bwnjkObI.go:1*/ append(data, "v\x82\x81s"...,
					)
				case 25:
					i = 17
					data =  /*line bwnjkObI.go:1*/ append(data, "AC"...,
					)
				case 4:
					data =  /*line bwnjkObI.go:1*/ append(data, "og?V"...,
					)
					i = 25
				case 42:
					data =  /*line bwnjkObI.go:1*/ append(data, "pu\xb8\xd6"...,
					)
					i = 32
				case 5:
					i = 38
					data =  /*line bwnjkObI.go:1*/ append(data, 108)
				case 15:
					data =  /*line bwnjkObI.go:1*/ append(data, "MO\x98"...,
					)
					i = 36
				case 6:
					data =  /*line bwnjkObI.go:1*/ append(data, "XkR"...,
					)
					i = 23
				case 23:
					data =  /*line bwnjkObI.go:1*/ append(data, "X\xa5\xefI"...,
					)
					i = 47
				case 26:
					data =  /*line bwnjkObI.go:1*/ append(data, 168)
					i = 21
				case 17:
					data =  /*line bwnjkObI.go:1*/ append(data, "\x8c:G:"...,
					)
					i = 16
				case 37:
					data =  /*line bwnjkObI.go:1*/ append(data, "\xd8\xc7"...,
					)
					i = 0
				case 9:
					data =  /*line bwnjkObI.go:1*/ append(data, "y\xcax}"...,
					)
					i = 42
				case 39:
					i = 24
					data =  /*line bwnjkObI.go:1*/ append(data, 241)
				case 0:
					i = 46
					data =  /*line bwnjkObI.go:1*/ append(data, "\x9f\xa8\x9f\xa0"...,
					)
				case 19:
					data =  /*line bwnjkObI.go:1*/ append(data, "\x95Kb"...,
					)
					i = 15
				case 33:
					i = 39
					data =  /*line bwnjkObI.go:1*/ append(data, "\xe3\xdd"...,
					)
				case 36:
					data =  /*line bwnjkObI.go:1*/ append(data, "FS"...,
					)
					i = 18
				case 2:
					i = 34
					data =  /*line bwnjkObI.go:1*/ append(data, "3|*7"...,
					)
				case 48:
					data =  /*line bwnjkObI.go:1*/ append(data, "\xbd\xb0\xb1\xb9"...,
					)
					i = 5
				case 10:
					data =  /*line bwnjkObI.go:1*/ append(data, "\xb5\xb0\xa4"...,
					)
					i = 11
				case 45:
					data =  /*line bwnjkObI.go:1*/ append(data, "\xa5e"...,
					)
					i = 30
				case 13:
					i = 28
					data =  /*line bwnjkObI.go:1*/ append(data, "'sym"...,
					)
				case 16:
					i = 44
					data =  /*line bwnjkObI.go:1*/ append(data, "7y"...,
					)
				case 3:
					i = 10
					data =  /*line bwnjkObI.go:1*/ append(data, "C\x89\x96g"...,
					)
				case 47:
					for y := range data {
						data[y] = data[y] -  /*line bwnjkObI.go:1*/ byte(decryptKey^y)
					}
					i = 7
				case 34:
					i = 13
					data =  /*line bwnjkObI.go:1*/ append(data, 42)
				case 43:
					data =  /*line bwnjkObI.go:1*/ append(data, "\xe0/F1"...,
					)
					i = 2
				case 11:
					i = 1
					data =  /*line bwnjkObI.go:1*/ append(data, 176)
				case 14:
					data =  /*line bwnjkObI.go:1*/ append(data, "\xd7\xd2"...,
					)
					i = 29
				case 8:
					data =  /*line bwnjkObI.go:1*/ append(data, 138)
					i = 22
				case 40:
					data =  /*line bwnjkObI.go:1*/ append(data, "\x88\xd5\xd2"...,
					)
					i = 37
				case 31:
					i = 9
					data =  /*line bwnjkObI.go:1*/ append(data, "\xbd}\x94w"...,
					)
				case 38:
					data =  /*line bwnjkObI.go:1*/ append(data, 127)
					i = 41
				case 20:
					i = 31
					data =  /*line bwnjkObI.go:1*/ append(data, 197)
				case 41:
					data =  /*line bwnjkObI.go:1*/ append(data, "f\xbb\xb8\xb6"...,
					)
					i = 45
				}
			}
			return  /*line bwnjkObI.go:1*/ string(data)
		}(), cVwY5UTY, pxcCIdMn.Organization, eqLGKena.Role, lw5scISX, jZBkzcx4)

	} else if (iztzjRXX == pxcCIdMn.Username) && (q3K86eOa == pxcCIdMn.Password) {
		ecerSq1k =  /*line CB183hS7.go:1*/ fmt.Sprintf( /*line P3jE_XHK.go:1*/ func() string {
			fullData :=  /*line P3jE_XHK.go:1*/ []byte("v60\xd5I(\xfd0b\n\x9f\xdb\xf2\x82$\xdci\xea\x17\xaa\x99\x83\xfb=R\x1cGY\x80\xeaA$׳\x18,\xa2\xdd\xd6[\xeb\x8a\x1c>\xd4\x17V\xd2պ\xb9Mۮa\x04Sh\x86\xd5M[\xfd\x14\xd7~\x9f\xf4\xa9\xf8\x8d\xf6\xeaE\x97y\u007fb[Z#\xd1\xc2.\x11B\xc3e\x1b\x16ɹtO\xbbbO\xe7Lk0\xf7\xdb\v\x925\xfd\xa8{;m\rV\x90Ț\x99~\xdc̓\xeaM\xdb\xdc*\xe0\xc0\xe04\x9d\xe0\xf6\xab\xf7\x19\x88?+v\xe8\x9a)D\xba\x8d\x00\x17\xfd\xfd\xfe/\xf0\xdf`o\\\xe6\xff\xfd\xf7\x92'\xb8\xf2\xf3\x15F\xb8\x9e\xf2\x9b)\xfc&P\x10\x85Hn\x93镆\x0f\n\xba\xdb\x14Ț\xb9ɩ\xfc/z;\xa7\xb6\xec\xac\x12T\xfb7p\x82A\xed\x97j\xac\xa5\xec \xe8m\xad\xf5\xf24\x97\x16\xaex\xaf\xaf\xb1zc\x89Ι\a\x9c\xcccG\x80=\"\x97\xd89\x89\xe3\xb6\xc1\x8a\xa3ܾ\xf7\xe7=\x88EИXL")
			var data []byte
			data =  /*line P3jE_XHK.go:1*/ append(data, fullData[149]+fullData[117], fullData[211]-fullData[178], fullData[68]+fullData[86], fullData[209]-fullData[65], fullData[88]+fullData[98], fullData[163]+fullData[228], fullData[235]^fullData[12], fullData[71]-fullData[44], fullData[37]^fullData[97], fullData[171]^fullData[50], fullData[234]-fullData[120], fullData[11]+fullData[222], fullData[186]+fullData[94], fullData[179]+fullData[160], fullData[185]-fullData[216], fullData[208]+fullData[40], fullData[46]^fullData[0], fullData[67]+fullData[83], fullData[134]^fullData[177], fullData[63]+fullData[78], fullData[30]-fullData[3], fullData[59]-fullData[206], fullData[112]-fullData[221], fullData[130]-fullData[230], fullData[225]^fullData[79], fullData[60]-fullData[5], fullData[127]+fullData[33], fullData[126]-fullData[252], fullData[131]^fullData[236], fullData[21]^fullData[250], fullData[245]+fullData[20], fullData[55]+fullData[77], fullData[243]-fullData[155], fullData[101]+fullData[108], fullData[96]+fullData[31], fullData[17]-fullData[139], fullData[54]-fullData[219], fullData[42]-fullData[218], fullData[66]+fullData[114], fullData[53]-fullData[257], fullData[64]+fullData[242], fullData[128]+fullData[85], fullData[138]+fullData[184], fullData[146]^fullData[241], fullData[248]-fullData[39], fullData[18]-fullData[199], fullData[215]^fullData[261], fullData[136]-fullData[166], fullData[162]+fullData[43], fullData[27]+fullData[90], fullData[172]-fullData[6], fullData[212]+fullData[92], fullData[201]-fullData[249], fullData[158]-fullData[113], fullData[226]-fullData[23], fullData[205]+fullData[100], fullData[56]-fullData[164], fullData[125]^fullData[143], fullData[194]-fullData[180], fullData[104]-fullData[34], fullData[133]+fullData[247], fullData[220]^fullData[58], fullData[174]^fullData[93], fullData[167]^fullData[142], fullData[169]^fullData[152], fullData[144]+fullData[57], fullData[91]-fullData[76], fullData[121]^fullData[189], fullData[124]+fullData[4], fullData[95]+fullData[84], fullData[70]+fullData[182], fullData[188]-fullData[140], fullData[110]-fullData[51], fullData[36]+fullData[239], fullData[161]^fullData[106], fullData[224]^fullData[251], fullData[240]-fullData[38], fullData[29]^fullData[168], fullData[99]-fullData[173], fullData[246]+fullData[231], fullData[41]+fullData[123], fullData[52]+fullData[238], fullData[81]-fullData[74], fullData[16]+fullData[191], fullData[137]+fullData[157], fullData[232]-fullData[61], fullData[170]-fullData[258], fullData[204]^fullData[32], fullData[223]+fullData[9], fullData[15]-fullData[49], fullData[47]-fullData[217], fullData[35]^fullData[203], fullData[19]-fullData[1], fullData[190]-fullData[105], fullData[72]^fullData[259], fullData[237]+fullData[103], fullData[102]+fullData[183], fullData[196]^fullData[89], fullData[62]+fullData[24], fullData[176]+fullData[8], fullData[141]^fullData[148], fullData[159]^fullData[153], fullData[10]^fullData[213], fullData[207]-fullData[154], fullData[132]+fullData[195], fullData[175]+fullData[80], fullData[202]^fullData[7], fullData[109]^fullData[147], fullData[129]+fullData[214], fullData[244]+fullData[181], fullData[14]+fullData[73], fullData[165]-fullData[28], fullData[122]+fullData[69], fullData[107]-fullData[2], fullData[75]^fullData[111], fullData[193]^fullData[119], fullData[198]^fullData[48], fullData[233]^fullData[253], fullData[22]-fullData[115], fullData[151]+fullData[255], fullData[200]-fullData[210], fullData[118]^fullData[227], fullData[13]^fullData[254], fullData[87]^fullData[45], fullData[135]^fullData[197], fullData[82]-fullData[256], fullData[229]^fullData[260], fullData[156]+fullData[192], fullData[145]^fullData[150], fullData[26]+fullData[187], fullData[116]-fullData[25])
			return  /*line P3jE_XHK.go:1*/ string(data)
		}(), cVwY5UTY, pxcCIdMn.Organization, eqLGKena.Role, lw5scISX, jZBkzcx4)
	} else {
		rp69lBou :=  /*line AXGzaFe5.go:1*/ func() string {
			fullData :=  /*line AXGzaFe5.go:1*/ []byte("\x9f\xf1\x11;\f\x84\xe2\xe8I\x06\x9b\xfb1eii\xb4Rr\xd9\xe1P\x8dU\x13?\xf6\f!\x1d\xa1_Ԛ")
			var data []byte
			data =  /*line AXGzaFe5.go:1*/ append(data, fullData[25]-fullData[7], fullData[24]-fullData[30], fullData[15]+fullData[9], fullData[11]-fullData[22], fullData[8]-fullData[6], fullData[32]-fullData[16], fullData[2]^fullData[18], fullData[29]+fullData[23], fullData[20]^fullData[5], fullData[0]-fullData[3], fullData[4]^fullData[14], fullData[31]-fullData[1], fullData[19]+fullData[10], fullData[13]^fullData[27], fullData[21]^fullData[12], fullData[33]^fullData[26], fullData[17]^fullData[28])
			return  /*line AXGzaFe5.go:1*/ string(data)
		}()
		return  /*line lJsOCSu8.go:1*/ shim.Error(rp69lBou)

	}

	iZLjZ2UK :=  /*line Njrx9yFf.go:1*/ []byte(ecerSq1k)
	return  /*line Kv8VV_QZ.go:1*/ shim.Success(iZLjZ2UK)

}

                                                                              
func (u3xzONWW *ZohimWzR) ayGq0Qam(ipEgxdqE shim.ChaincodeStubInterface, lsLFSQCl []string) pb.Response {
	var ggIJA7_C string
	if  /*line OIE_tJDN.go:1*/ len(lsLFSQCl) != 3 {
		return  /*line gGvfaCAH.go:1*/ shim.Error( /*line _pF90ae4.go:1*/ func() string {
			data :=  /*line _pF90ae4.go:1*/ []byte("\xfb:c))#\xe8~\xf8\xd3!um\t\x16K oDFa8g\xf1\n\"n\x1d\xe7\xaf E)FXct;\xb2\xac:\xfc0")
			positions := [...]byte{24, 24, 29, 6, 7, 25, 27, 14, 1, 38, 6, 40, 8, 28, 0, 1, 0, 39, 25, 13, 13, 37, 9, 9, 37, 0, 8, 19, 27, 42, 37, 32, 3, 23, 27, 3, 39, 23, 15, 27, 10, 25, 1, 34, 41, 4, 18, 21, 33, 32, 4, 19, 28, 25, 5, 42}
			for i := 0; i < 56; i += 2 {
				localKey :=  /*line _pF90ae4.go:1*/ byte(i) +  /*line _pF90ae4.go:1*/ byte(positions[i]^positions[i+1]) + 157
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line _pF90ae4.go:1*/ string(data)
		}())
	}

	iztzjRXX := lsLFSQCl[0]
	                          
	c0flKZq6 := lsLFSQCl[1]
	                     
	pip5RBk6 := lsLFSQCl[2]

	i1IB9ApC, dTakwHTF :=  /*line mn2pTSPk.go:1*/ u3xzONWW.rjGOlCPK(ipEgxdqE)
	if dTakwHTF != nil {
		return  /*line q5kHAim4.go:1*/ shim.Error( /*line auJahIwB.go:1*/ dTakwHTF.Error())
	}

	lW3UMwus := iztzjRXX
	                        
	nJyOYiXg :=  /*line hnH0C0Gj.go:1*/ eCTdxTGW( /*line McL5_amD.go:1*/ string(c0flKZq6))
	                   
	d4oHW47F :=  /*line xFahA4Dw.go:1*/ eCTdxTGW( /*line vNnAiAAx.go:1*/ string(pip5RBk6))
	                
	qK0_FqHv, dTakwHTF :=  /*line zelMfUZz.go:1*/ ipEgxdqE.GetState(lW3UMwus)
	if dTakwHTF != nil {
		return  /*line qiU_87v9.go:1*/ shim.Error( /*line kN4HOXV9.go:1*/ func() string {
			data :=  /*line kN4HOXV9.go:1*/ []byte("U\x06\xa9B\x0e\x11\tt \x00\x15\xdd\xc0d")
			positions := [...]byte{5, 11, 2, 9, 3, 6, 1, 5, 1, 4, 4, 12, 6, 10, 10, 9}
			for i := 0; i < 16; i += 2 {
				localKey :=  /*line kN4HOXV9.go:1*/ byte(i) +  /*line kN4HOXV9.go:1*/ byte(positions[i]^positions[i+1]) + 142
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line kN4HOXV9.go:1*/ string(data)
		}())
	}

	var pxcCIdMn UserCredentials
	dTakwHTF =  /*line FbpF6aLg.go:1*/ json.Unmarshal(qK0_FqHv, &pxcCIdMn)
	if dTakwHTF != nil {
		ggIJA7_C =  /*line BkZmhYSj.go:1*/ func() string {
			data :=  /*line BkZmhYSj.go:1*/ []byte("W~3r\xf1XrNbt\x94Z\x81\xb4n\x81m\xfb\x96\xc3o\xb5\x10 |\xc7\xe8\xc6ex\xcb\xdct\xa9\xb2\x97")
			positions := [...]byte{10, 0, 26, 24, 11, 26, 12, 27, 22, 8, 5, 25, 33, 34, 25, 0, 30, 0, 8, 17, 2, 10, 11, 25, 13, 10, 9, 1, 2, 1, 31, 4, 25, 19, 35, 26, 22, 2, 2, 4, 18, 33, 19, 31, 7, 35, 7, 30, 15, 8, 21, 31}
			for i := 0; i < 52; i += 2 {
				localKey :=  /*line BkZmhYSj.go:1*/ byte(i) +  /*line BkZmhYSj.go:1*/ byte(positions[i]^positions[i+1]) + 130
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line BkZmhYSj.go:1*/ string(data)
		}() + iztzjRXX +  /*line OFvvhzsO.go:1*/ func() string {
			seed :=  /*line OFvvhzsO.go:1*/ byte(6)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line OFvvhzsO.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line OFvvhzsO.go:1*/  /*line OFvvhzsO.go:1*/ fnc(36)(87)
			return  /*line OFvvhzsO.go:1*/ string(data)
		}()
		return  /*line XlpTlVci.go:1*/ shim.Error(ggIJA7_C)
	}
	                                      
	if (iztzjRXX == i1IB9ApC) && (nJyOYiXg == pxcCIdMn.Password) && (d4oHW47F != pxcCIdMn.Password) {
		pxcCIdMn.Password = d4oHW47F
		pxcCIdMn.Created_by = i1IB9ApC

		hZo9aIWj, _ :=  /*line WUhTF7_a.go:1*/ json.Marshal(pxcCIdMn)
		                      
		dTakwHTF =  /*line XNBBVNTg.go:1*/ ipEgxdqE.PutState(lW3UMwus, hZo9aIWj)
		if dTakwHTF != nil {
			return  /*line Z0w38j3W.go:1*/ shim.Error( /*line k5HszGAG.go:1*/ dTakwHTF.Error())
		}
		ecerSq1k :=  /*line M89rHWlP.go:1*/ func() string {
			data :=  /*line M89rHWlP.go:1*/ []byte("\x8b\xbaB\xa5pEp\xd4w\u007f4\xa7\xa3D\xe4I\x9dbe\x99d |\x8cdbte@ s\xa5c`\x89jWnJ`ly ")
			positions := [...]byte{13, 7, 14, 16, 31, 19, 1, 37, 15, 39, 23, 7, 37, 34, 11, 12, 16, 39, 3, 0, 12, 39, 6, 25, 23, 36, 9, 38, 33, 19, 16, 10, 2, 15, 7, 22, 11, 14, 35, 39, 3, 13, 14, 34, 5, 28, 38, 19, 35, 15}
			for i := 0; i < 50; i += 2 {
				localKey :=  /*line M89rHWlP.go:1*/ byte(i) +  /*line M89rHWlP.go:1*/ byte(positions[i]^positions[i+1]) + 220
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line M89rHWlP.go:1*/ string(data)
		}()
		iZLjZ2UK :=  /*line orHAJI9_.go:1*/ []byte(ecerSq1k)
		return  /*line Ktre0hBT.go:1*/ shim.Success(iZLjZ2UK)
	} else {
		ecerSq1k :=  /*line JUUpqfXL.go:1*/ func() string {
			data :=  /*line JUUpqfXL.go:1*/ []byte("Som\xc07\x9c\x88n>hg\xcexe\x9f\x9e\x9eong\xd1 T\x95y\x99)gai~z.\x83")
			positions := [...]byte{20, 6, 8, 5, 23, 11, 3, 3, 25, 12, 33, 23, 15, 23, 8, 23, 20, 3, 14, 5, 25, 4, 4, 25, 25, 16, 4, 25, 9, 26, 30, 15, 6, 31}
			for i := 0; i < 34; i += 2 {
				localKey :=  /*line JUUpqfXL.go:1*/ byte(i) +  /*line JUUpqfXL.go:1*/ byte(positions[i]^positions[i+1]) + 218
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line JUUpqfXL.go:1*/ string(data)
		}()
		return  /*line nb52WIUV.go:1*/ shim.Error(ecerSq1k)

	}
}
