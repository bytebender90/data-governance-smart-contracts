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

func (g4rnrSNn *G1Y_7pz_) wKBsTKzx(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	var iECyqXpw []string
	if  /*line JDxnGsbs.go:1*/ len(ktsi1_SQ) != 4 {
		return  /*line kaX4QIuF.go:1*/ shim.Error( /*line nGbvNK9Y.go:1*/ func() string {
			key :=  /*line nGbvNK9Y.go:1*/ []byte("\xb3\xb5\xa1J\x03\x9f\xa1\xc7\b\x12Ԛqѳ\x00\x1f\xd5%\xc4֬\xf2C\xc7\xf3y\x1d\x16\xe6n\x8a!\x8d/@=\x96\x19Oö>\x9f&N\x8b/sQ\x14\x9a\x11 \xd3v\xa72\xfb2\xef/<\xdb^rS\xc6\x06\xbdݲ\x90\xf3\xc6ǍdBcEp\xc1\xab\xc6i\x92ߥ`\x13\xd2:\x8b\xdfRQ\xccI\xbf\xaf\xa1")
			data :=  /*line nGbvNK9Y.go:1*/ []byte("\xfc#\x04\xb9u\x11\x06*|2B\x0f\xde3\x18r?D\x8b\xe47\x1eY\xb84X瑉\x14\x8eϙ\xfd\x94\xa3\xb1\xff\x87\xb6\xe3\xeal\xbfj\xaf\xff\x90消\xba\u007f\x81@\xdb\xd3RK\x97a\x9c\xa5N\xd1\xdb\xc242\xdd '\x03g54\xad\xa5\xa5ƪ\xe34\xcb\x18\xd2\xf9G\x19\xd33\xf8Z\xffH\xbf\xb6?\xbd \x1c\x11")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line nGbvNK9Y.go:1*/ string(data)
		}())
	}

	eviEA3sp := ktsi1_SQ[0]
	olYMzSWK := ktsi1_SQ[1]
	tIeFhn0D := ktsi1_SQ[2]
	ezgSKu0t :=  /*line f4mJnvvQ.go:1*/ strings.Split(ktsi1_SQ[3],  /*line qJSkKFhY.go:1*/ func() string {
		seed :=  /*line qJSkKFhY.go:1*/ byte(0)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line qJSkKFhY.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line qJSkKFhY.go:1*/ fnc(44)
		return  /*line qJSkKFhY.go:1*/ string(data)
	}())

	b4dMB1zL, cSW1wSO3 :=  /*line IT7P6Wm2.go:1*/ jmSY84co(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line gFhB2gON.go:1*/ shim.Error( /*line yY0lBabS.go:1*/ cSW1wSO3.Error())
	}

	                                                                  
	zzhdvyqZ := &DatasetMetadataResponse{}
	cSW1wSO3 =  /*line JXsjGtR3.go:1*/ json.Unmarshal( /*line kGSoUfD_.go:1*/ []byte(b4dMB1zL), zzhdvyqZ)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line akYcfz0z.go:1*/ func() string {
			fullData :=  /*line akYcfz0z.go:1*/ []byte("\x1cwHQ\xaaD\x93\x027\t6\xb1\xd1;\xf4\x8aJ`i\xa0\xb6W'\x8a\x9f\xc5\xed\x18\xa4.\xea1\x1e\xc3_\xba\xfd\x0f(\xaa\x05HQol>\xec\n\x05\b\x94\xf8\a\xee\xe7jϪ\xe4V]\xca")
			var data []byte
			data =  /*line akYcfz0z.go:1*/ append(data, fullData[20]+fullData[25], fullData[9]-fullData[54], fullData[45]+fullData[52], fullData[16]+fullData[38], fullData[24]^fullData[26], fullData[2]+fullData[22], fullData[51]^fullData[15], fullData[60]-fullData[13], fullData[18]+fullData[12], fullData[6]^fullData[11], fullData[7]+fullData[5], fullData[21]+fullData[47], fullData[3]+fullData[27], fullData[43]+fullData[36], fullData[40]-fullData[19], fullData[57]+fullData[35], fullData[30]-fullData[61], fullData[8]-fullData[33], fullData[48]+fullData[55], fullData[23]^fullData[39], fullData[0]+fullData[41], fullData[42]-fullData[46], fullData[50]+fullData[56], fullData[37]^fullData[17], fullData[44]^fullData[49], fullData[53]+fullData[1], fullData[59]-fullData[10], fullData[29]-fullData[58], fullData[14]+fullData[34], fullData[32]+fullData[31], fullData[4]+fullData[28])
			return  /*line akYcfz0z.go:1*/ string(data)
		}() + eviEA3sp +  /*line JO8ybRhq.go:1*/ func() string {
			key :=  /*line JO8ybRhq.go:1*/ []byte("\x96\xaf")
			data :=  /*line JO8ybRhq.go:1*/ []byte("\x8c\xce")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line JO8ybRhq.go:1*/ string(data)
		}()
		return  /*line XN6YtiU8.go:1*/ shim.Error(w5VhCLka)
	}

	if  /*line pyuV6dwW.go:1*/ nSRaDwQg(ezgSKu0t, zzhdvyqZ.CustomAccessRights) {
		 /*line okjCqjqR.go:1*/ fmt.Println( /*line LsTvL6DD.go:1*/ func() string {
			seed :=  /*line LsTvL6DD.go:1*/ byte(243)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line LsTvL6DD.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/  /*line LsTvL6DD.go:1*/ fnc(52)(151)(46)(94)(185)(115)(232)(199)(134)(31)(47)(25)(117)(252)(246)(237)(213)(168)(3)(71)(144)(32)(66)(146)(36)(245)(60)(111)(220)(185)(126)(251)
			return  /*line LsTvL6DD.go:1*/ string(data)
		}())
	} else {
		w5VhCLka =  /*line NdCfK9S_.go:1*/ func() string {
			key :=  /*line NdCfK9S_.go:1*/ []byte("\x98\xd9R}\xae'\xa7L\x89YB(\xc3\x1b6\xb9J\xd1v\xf2G%\n$`\xd4\x13ʅ\x80x'i\x97\xb8\xbb\x82\x88")
			data :=  /*line NdCfK9S_.go:1*/ []byte("\xe3I\xf3\xf5\xc4H\xcbֱ\xc9\x15J\xacS1g\x19\xa4\xfd\x82(H\x16=\x03\x8fR\xa9\xee\xa0\xfaB\xfeѼ\xb8\xa0\xf5")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line NdCfK9S_.go:1*/ string(data)
		}()
		return  /*line nkTIhq29.go:1*/ shim.Error(w5VhCLka)
	}

	                                            
	lGAVkzbs, cSW1wSO3 :=  /*line Gi2I3uqq.go:1*/ n4c7mRtG.GetState(eviEA3sp)
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line MPg3KSmx.go:1*/ func() string {
			key :=  /*line MPg3KSmx.go:1*/ []byte("\xe0\x12\xbdIvn\xb98\xdfL\n\f-\x10b\x19\xd6q\x93\xfe\xa4\xa4W-%\x00N(\x8cwe\xe4Ԃ\xbab9\xeaE\xa7\xf7~Q\xed\xc2\xe5\xce\xc8m\xde\xd9+R\x9f\xf0K\x1c")
			data :=  /*line MPg3KSmx.go:1*/ []byte("\x9b0\xf8;\x04\x01\xcb\x1a\xe5nLmD|\a}\xf6\x05\xfc\xde\xc7\xcc2NN /]\xf8\x1f\n\x96\xbd\xf8\xdb\x16P\x85+\x87\x91\x11#Ͷ\x8d\xab\xe8\t\xbf\xadJ!\xfa\x84q<")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line MPg3KSmx.go:1*/ string(data)
		}() + eviEA3sp +  /*line sp7nMnu4.go:1*/ func() string {
			key :=  /*line sp7nMnu4.go:1*/ []byte("X\f")
			data :=  /*line sp7nMnu4.go:1*/ []byte("\xcaq")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line sp7nMnu4.go:1*/ string(data)
		}()
		return  /*line U5F3qgiM.go:1*/ shim.Error(w5VhCLka)
	} else if (lGAVkzbs == nil) && (olYMzSWK == "" || olYMzSWK ==  /*line zWPnYFKM.go:1*/ func() string {
		data :=  /*line zWPnYFKM.go:1*/ []byte("s")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line zWPnYFKM.go:1*/ byte(i) +  /*line zWPnYFKM.go:1*/ byte(positions[i]^positions[i+1]) + 83
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line zWPnYFKM.go:1*/ string(data)
	}()) {
		w5VhCLka =  /*line nrwqblE2.go:1*/ func() string {
			seed :=  /*line nrwqblE2.go:1*/ byte(73)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line nrwqblE2.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/  /*line nrwqblE2.go:1*/ fnc(50)(167)(35)(45)(0)(253)(3)(176)(24)(232)(50)(20)(253)(187)(80)(245)(13)(251)(252)(10)(0)(246)(6)(255)(178)(73)(10)(173)(78)(1)(5)(172)(68)(1)(1)(3)(5)(247)(255)(202)(242)(38)(27)(8)(3)(249)(255)(188)(84)(251)(177)(70)(3)(5)(246)(188)(73)(5)(251)(11)(172)(80)(5)(237)(10)(253)(250)(189)(68)(253)(19)(237)(18)(242)(15)(198)(230)
			return  /*line nrwqblE2.go:1*/ string(data)
		}() + eviEA3sp +  /*line FaGJgUWe.go:1*/ func() string {
			seed :=  /*line FaGJgUWe.go:1*/ byte(175)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line FaGJgUWe.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line FaGJgUWe.go:1*/  /*line FaGJgUWe.go:1*/ fnc(115)(91)
			return  /*line FaGJgUWe.go:1*/ string(data)
		}()
		return  /*line IGZKFru4.go:1*/ shim.Error(w5VhCLka)
	} else if lGAVkzbs != nil {
		pukA0X3v := PublicDataset{}
		cSW1wSO3 =  /*line hkRAChTD.go:1*/ json.Unmarshal(lGAVkzbs, &pukA0X3v)	                               
		if cSW1wSO3 != nil {
			return  /*line RLhBy5yl.go:1*/ shim.Error( /*line YtfBetYB.go:1*/ cSW1wSO3.Error())
		}
		if olYMzSWK == "" || olYMzSWK ==  /*line GQ4aloMF.go:1*/ func() string {
			seed :=  /*line GQ4aloMF.go:1*/ byte(63)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line GQ4aloMF.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line GQ4aloMF.go:1*/ fnc(225)
			return  /*line GQ4aloMF.go:1*/ string(data)
		}() {
			iECyqXpw = pukA0X3v.Permission
		}
		 /*line RuCmso0t.go:1*/ fmt.Println( /*line zX91QgFB.go:1*/ func() string {
			seed :=  /*line zX91QgFB.go:1*/ byte(237)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line zX91QgFB.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/  /*line zX91QgFB.go:1*/ fnc(84)(47)(0)(245)(9)(246)(188)(67)(18)(254)(1)(251)(254)(179)(65)(2)(0)(2)(14)(0)(173)(82)(247)(254)(1)(12)(255)
			return  /*line zX91QgFB.go:1*/ string(data)
		}())
		if  /*line H7I2xt2T.go:1*/ len(ezgSKu0t) > 0 {
			ezgSKu0t =  /*line U4SpyUsH.go:1*/ _urwHavD(ezgSKu0t, pukA0X3v.CustomAccessRights)
		}
	}

	w8EOizcx, cSW1wSO3 :=  /*line BtNY6e8p.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line mIqFaVeX.go:1*/ shim.Error( /*line uqEkiX0P.go:1*/ cSW1wSO3.Error())
	}

	                                            
	qFOnXzGg, cSW1wSO3 :=  /*line f2HWLQjI.go:1*/ fvIQVpVM(n4c7mRtG, []string{eviEA3sp})
	if cSW1wSO3 != nil {
		return  /*line jfvZ5B9f.go:1*/ shim.Error( /*line FZJ6qtGW.go:1*/ cSW1wSO3.Error())
	}
	if w8EOizcx != qFOnXzGg {
		w5VhCLka =  /*line IARR6_BT.go:1*/ func() string {
			key :=  /*line IARR6_BT.go:1*/ []byte("c\xef<O\xb1\x8c\xa0\x9f7\xbe\xe3\xe9\x86Vy\xfdq.\xb4\x1a.&\xd1ӓv@E\xbc\xb5\xef6")
			data :=  /*line IARR6_BT.go:1*/ []byte("\x183\t#\xc1\xe3҃\x03dk\x86\xee\xca\xe8x\x03:\xbbX;T\x94\x91\x8d\xff3 \xb6\x851\xea")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line IARR6_BT.go:1*/ string(data)
		}() + w8EOizcx +  /*line HAXS4JVD.go:1*/ func() string {
			fullData :=  /*line HAXS4JVD.go:1*/ []byte("ȗ\xbbs\xeao\xbf؋\x9c\xde?\xf3\x95\x98Z\x02\x85x\x11w\xa9\xc8\xc4r\xae")
			var data []byte
			data =  /*line HAXS4JVD.go:1*/ append(data, fullData[12]+fullData[3], fullData[7]+fullData[1], fullData[23]+fullData[25], fullData[14]-fullData[18], fullData[22]+fullData[9], fullData[10]^fullData[6], fullData[17]-fullData[19], fullData[20]+fullData[4], fullData[2]^fullData[0], fullData[15]^fullData[11], fullData[16]+fullData[24], fullData[21]-fullData[5], fullData[8]+fullData[13])
			return  /*line HAXS4JVD.go:1*/ string(data)
		}() + eviEA3sp +  /*line O9wywVN7.go:1*/ func() string {
			fullData :=  /*line O9wywVN7.go:1*/ []byte("\x8d k]")
			var data []byte
			data =  /*line O9wywVN7.go:1*/ append(data, fullData[0]-fullData[2], fullData[3]^fullData[1])
			return  /*line O9wywVN7.go:1*/ string(data)
		}()
		return  /*line HmK5ft5t.go:1*/ shim.Error(w5VhCLka)
	}

	if olYMzSWK ==  /*line l_nGVFij.go:1*/ func() string {
		key :=  /*line l_nGVFij.go:1*/ []byte("\xc3\xef\x9f$")
		data :=  /*line l_nGVFij.go:1*/ []byte("\xb1\x8a\xfe@")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line l_nGVFij.go:1*/ string(data)
	}() {
		iECyqXpw = []string{ /*line CBhkx1MZ.go:1*/ func() string {
			fullData :=  /*line CBhkx1MZ.go:1*/ []byte("\xc3\xf0\x94\x16â\\\xa2")
			var data []byte
			data =  /*line CBhkx1MZ.go:1*/ append(data, fullData[6]+fullData[3], fullData[7]+fullData[0], fullData[5]^fullData[4], fullData[2]^fullData[1])
			return  /*line CBhkx1MZ.go:1*/ string(data)
		}()}
	} else if olYMzSWK ==  /*line mPMDj8qX.go:1*/ func() string {
		data :=  /*line mPMDj8qX.go:1*/ []byte("\xa6\xa5\x8bi\xa3y")
		positions := [...]byte{4, 4, 2, 1, 4, 0, 2, 4}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line mPMDj8qX.go:1*/ byte(i) +  /*line mPMDj8qX.go:1*/ byte(positions[i]^positions[i+1]) + 23
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line mPMDj8qX.go:1*/ string(data)
	}() {
		iECyqXpw = []string{ /*line k6FLgP2f.go:1*/ func() string {
			key :=  /*line k6FLgP2f.go:1*/ []byte("Շ+@")
			data :=  /*line k6FLgP2f.go:1*/ []byte("\xa7\xe2J$")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line k6FLgP2f.go:1*/ string(data)
		}(),  /*line k45skitC.go:1*/ func() string {
			data :=  /*line k45skitC.go:1*/ []byte("mױ\xdbڜ")
			positions := [...]byte{3, 4, 5, 2, 3, 1, 3, 4}
			for i := 0; i < 8; i += 2 {
				localKey :=  /*line k45skitC.go:1*/ byte(i) +  /*line k45skitC.go:1*/ byte(positions[i]^positions[i+1]) + 47
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line k45skitC.go:1*/ string(data)
		}()}
	} else if olYMzSWK ==  /*line kYMZ66RP.go:1*/ func() string {
		key :=  /*line kYMZ66RP.go:1*/ []byte("3T֡\x89\xff\xaf")
		data :=  /*line kYMZ66RP.go:1*/ []byte("\xa3\xb9H\x14\xf2r#")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line kYMZ66RP.go:1*/ string(data)
	}() {
		iECyqXpw = []string{ /*line pqUcOV25.go:1*/ func() string {
			key :=  /*line pqUcOV25.go:1*/ []byte("\x16o\x05\xa1")
			data :=  /*line pqUcOV25.go:1*/ []byte("\x88\xd4f\x05")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line pqUcOV25.go:1*/ string(data)
		}(),  /*line Czgb7XyG.go:1*/ func() string {
			fullData :=  /*line Czgb7XyG.go:1*/ []byte("\x9a\x80\U000b8db0\xfe\xcd\xed\xaca\xd1")
			var data []byte
			data =  /*line Czgb7XyG.go:1*/ append(data, fullData[8]^fullData[1], fullData[10]-fullData[2], fullData[0]^fullData[6], fullData[11]^fullData[3], fullData[5]+fullData[4], fullData[7]+fullData[9])
			return  /*line Czgb7XyG.go:1*/ string(data)
		}(),  /*line aRGlavDI.go:1*/ func() string {
			fullData :=  /*line aRGlavDI.go:1*/ []byte("\xec\x1bX\bM\x10\x9f\x9e]N\xbd\xd6\x1b\xa2")
			var data []byte
			data =  /*line aRGlavDI.go:1*/ append(data, fullData[10]-fullData[4], fullData[3]+fullData[8], fullData[5]-fullData[7], fullData[0]^fullData[6], fullData[12]+fullData[9], fullData[1]+fullData[2], fullData[13]^fullData[11])
			return  /*line aRGlavDI.go:1*/ string(data)
		}()}
	}

	                                          
	mzjt9rLK :=  /*line o7Oa1Rz8.go:1*/ func() string {
		var data []byte
		i := 8
		decryptKey := 86
		for counter := 0; i != 2; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 4:
				i = 5
				data =  /*line o7Oa1Rz8.go:1*/ append(data, 18)
			case 1:
				data =  /*line o7Oa1Rz8.go:1*/ append(data, "\x03\x17"...,
				)
				i = 7
			case 3:
				data =  /*line o7Oa1Rz8.go:1*/ append(data, 1)
				i = 0
			case 6:
				data =  /*line o7Oa1Rz8.go:1*/ append(data, 25)
				i = 1
			case 5:
				data =  /*line o7Oa1Rz8.go:1*/ append(data, "\x19="...,
				)
				i = 6
			case 8:
				i = 4
				data =  /*line o7Oa1Rz8.go:1*/ append(data, "/\v\x1f\x10"...,
				)
			case 7:
				data =  /*line o7Oa1Rz8.go:1*/ append(data, "\x06\x11\a"...,
				)
				i = 3
			case 0:
				i = 2
				for y := range data {
					data[y] = data[y] ^  /*line o7Oa1Rz8.go:1*/ byte(decryptKey^y)
				}
			}
		}
		return  /*line o7Oa1Rz8.go:1*/ string(data)
	}()
	l5TiVBv5 := &PublicDataset{mzjt9rLK, eviEA3sp, w8EOizcx, iECyqXpw, ezgSKu0t}
	niQzOH6M, cSW1wSO3 :=  /*line sWWtClFZ.go:1*/ json.Marshal(l5TiVBv5)
	if cSW1wSO3 != nil {
		return  /*line e3a3f2uj.go:1*/ shim.Error( /*line HE2ix1M3.go:1*/ cSW1wSO3.Error())
	}
	syHhHGpa := eviEA3sp
	                        
	cSW1wSO3 =  /*line c1s428KM.go:1*/ n4c7mRtG.PutState(syHhHGpa, niQzOH6M)
	if cSW1wSO3 != nil {
		return  /*line fQNzPj08.go:1*/ shim.Error( /*line EEFQOoQh.go:1*/ cSW1wSO3.Error())
	}

	               
	xwK9eDsS, cSW1wSO3 :=  /*line iWuuplAE.go:1*/ _YWFxzXg(n4c7mRtG, []string{eviEA3sp, w8EOizcx, olYMzSWK, tIeFhn0D,  /*line Ia3dgT4S.go:1*/ strings.Join(ezgSKu0t,  /*line TozMuww0.go:1*/ func() string {
		fullData :=  /*line TozMuww0.go:1*/ []byte("Jf")
		var data []byte
		data =  /*line TozMuww0.go:1*/ append(data, fullData[1]^fullData[0])
		return  /*line TozMuww0.go:1*/ string(data)
	}())})
	if cSW1wSO3 != nil {
		return  /*line rh3fKAFd.go:1*/ shim.Error( /*line FZ1i5rI4.go:1*/ cSW1wSO3.Error())
	}
	 /*line zGox0AIe.go:1*/ fmt.Println(xwK9eDsS)

	xHTOVzVs :=  /*line detwmrpH.go:1*/ func() string {
		key :=  /*line detwmrpH.go:1*/ []byte("\x83I\x05\x03\xbc\xafx\xd4\xe4\xfe\xbb\xb9\xbf")
		data :=  /*line detwmrpH.go:1*/ []byte("\xd1\x1f`\x1d\xa8\xb2\xfc\x8d\x8fg\xb9\x81a")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line detwmrpH.go:1*/ string(data)
	}() + eviEA3sp +  /*line IREAcJRG.go:1*/ func() string {
		key :=  /*line IREAcJRG.go:1*/ []byte("\xe3\x15ߟ\x0f\xaa~\xfd\xb5,\xb2\x15\xe1\xfd\xd9\xf2\x1b8\xcc\xf3\x90\xefr;[\x85")
		data :=  /*line IREAcJRG.go:1*/ []byte("\x03~R\xbf\u007f\x1f\xe0i\x1e\x8fҌJqA\x12\x8b\x9d>`\xf9b\xe5\xa4\xca\xf3")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line IREAcJRG.go:1*/ string(data)
	}() + olYMzSWK
	q8ZF8zaz :=  /*line uDeyo2dU.go:1*/ []byte(xHTOVzVs)

	return  /*line FrEXHXWC.go:1*/ shim.Success(q8ZF8zaz)
}

func (g4rnrSNn *G1Y_7pz_) j5pqkIRD(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	var cSW1wSO3 error
	if  /*line XPnL_Jan.go:1*/ len(ktsi1_SQ) != 4 {
		return  /*line jx_cdDoX.go:1*/ shim.Error( /*line ZqvGLtxv.go:1*/ func() string {
			seed :=  /*line ZqvGLtxv.go:1*/ byte(55)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line ZqvGLtxv.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/  /*line ZqvGLtxv.go:1*/ fnc(128)(37)(63)(138)(23)(46)(79)(156)(73)(62)(202)(155)(46)(81)(165)(87)(92)(7)(5)(196)(201)(163)(59)(132)(0)(248)(249)(248)(239)(153)(36)(109)(13)(18)(25)(48)(113)(215)(179)(95)(119)(2)
			return  /*line ZqvGLtxv.go:1*/ string(data)
		}())
	}

	eviEA3sp := ktsi1_SQ[0]
	olYMzSWK := ktsi1_SQ[1]
	tIeFhn0D := ktsi1_SQ[2]
	ezgSKu0t :=  /*line jNey9egB.go:1*/ strings.Split(ktsi1_SQ[3],  /*line kcVTfs3X.go:1*/ func() string {
		key :=  /*line kcVTfs3X.go:1*/ []byte("\xbb")
		data :=  /*line kcVTfs3X.go:1*/ []byte("\xe7")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line kcVTfs3X.go:1*/ string(data)
	}())

	w8EOizcx, cSW1wSO3 :=  /*line WBQqgscV.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line D1B2Fqbt.go:1*/ shim.Error( /*line yycxjU_S.go:1*/ cSW1wSO3.Error())
	}

	tQcQN8RB, cSW1wSO3 :=  /*line T7CawKap.go:1*/ n4c7mRtG.GetState(eviEA3sp)	                    
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line uh3ZTyjh.go:1*/ func() string {
			var data []byte
			i := 1
			decryptKey := 0
			for counter := 0; i != 10; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 15:
					data =  /*line uh3ZTyjh.go:1*/ append(data, 237)
					i = 6
				case 19:
					i = 10
					for y := range data {
						data[y] = data[y] +  /*line uh3ZTyjh.go:1*/ byte(decryptKey^y)
					}
				case 3:
					data =  /*line uh3ZTyjh.go:1*/ append(data, "\xe9D73"...,
					)
					i = 15
				case 5:
					i = 2
					data =  /*line uh3ZTyjh.go:1*/ append(data, "\x11[_["...,
					)
				case 0:
					i = 13
					data =  /*line uh3ZTyjh.go:1*/ append(data, "SQ\x14g"...,
					)
				case 14:
					data =  /*line uh3ZTyjh.go:1*/ append(data, "[n"...,
					)
					i = 18
				case 20:
					data =  /*line uh3ZTyjh.go:1*/ append(data, 7)
					i = 21
				case 2:
					data =  /*line uh3ZTyjh.go:1*/ append(data, "Xg\x1b"...,
					)
					i = 14
				case 18:
					i = 11
					data =  /*line uh3ZTyjh.go:1*/ append(data, "tgmo"...,
					)
				case 16:
					data =  /*line uh3ZTyjh.go:1*/ append(data, "SZVX"...,
					)
					i = 20
				case 11:
					data =  /*line uh3ZTyjh.go:1*/ append(data, "-="...,
					)
					i = 7
				case 6:
					data =  /*line uh3ZTyjh.go:1*/ append(data, "84"...,
					)
					i = 9
				case 17:
					data =  /*line uh3ZTyjh.go:1*/ append(data, "Y["...,
					)
					i = 0
				case 8:
					data =  /*line uh3ZTyjh.go:1*/ append(data, "\xe52:<"...,
					)
					i = 3
				case 1:
					i = 16
					data =  /*line uh3ZTyjh.go:1*/ append(data, "_\x05'"...,
					)
				case 9:
					i = 4
					data =  /*line uh3ZTyjh.go:1*/ append(data, "F2K"...,
					)
				case 12:
					i = 8
					data =  /*line uh3ZTyjh.go:1*/ append(data, "5164"...,
					)
				case 7:
					data =  /*line uh3ZTyjh.go:1*/ append(data, 35)
					i = 12
				case 13:
					i = 5
					data =  /*line uh3ZTyjh.go:1*/ append(data, 97)
				case 4:
					data =  /*line uh3ZTyjh.go:1*/ append(data, "<J\x0f\xfc"...,
					)
					i = 19
				case 21:
					i = 17
					data =  /*line uh3ZTyjh.go:1*/ append(data, "&\r0J"...,
					)
				}
			}
			return  /*line uh3ZTyjh.go:1*/ string(data)
		}() + eviEA3sp +  /*line GBnnAJ2B.go:1*/ func() string {
			var data []byte
			i := 2
			decryptKey := 98
			for counter := 0; i != 1; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 2:
					data =  /*line GBnnAJ2B.go:1*/ append(data, 190)
					i = 0
				case 3:
					for y := range data {
						data[y] = data[y] +  /*line GBnnAJ2B.go:1*/ byte(decryptKey^y)
					}
					i = 1
				case 0:
					data =  /*line GBnnAJ2B.go:1*/ append(data, 24)
					i = 3
				}
			}
			return  /*line GBnnAJ2B.go:1*/ string(data)
		}()
		return  /*line J17vYlQ6.go:1*/ shim.Error(w5VhCLka)
	} else if tQcQN8RB == nil {
		w5VhCLka =  /*line uIO2wLNm.go:1*/ func() string {
			var data []byte
			i := 3
			decryptKey := 21
			for counter := 0; i != 15; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 12:
					i = 5
					data =  /*line uIO2wLNm.go:1*/ append(data, "0Uqr"...,
					)
				case 8:
					data =  /*line uIO2wLNm.go:1*/ append(data, "sq"...,
					)
					i = 7
				case 10:
					i = 1
					data =  /*line uIO2wLNm.go:1*/ append(data, "\x90_F"...,
					)
				case 5:
					i = 11
					data =  /*line uIO2wLNm.go:1*/ append(data, "vpp"...,
					)
				case 2:
					data =  /*line uIO2wLNm.go:1*/ append(data, "\x8d\x80"...,
					)
					i = 10
				case 0:
					i = 8
					data =  /*line uIO2wLNm.go:1*/ append(data, 122)
				case 16:
					i = 13
					data =  /*line uIO2wLNm.go:1*/ append(data, "\x82\x80"...,
					)
				case 7:
					data =  /*line uIO2wLNm.go:1*/ append(data, "u&G"...,
					)
					i = 12
				case 4:
					data =  /*line uIO2wLNm.go:1*/ append(data, "\x8a\x868"...,
					)
					i = 14
				case 9:
					i = 2
					data =  /*line uIO2wLNm.go:1*/ append(data, 122)
				case 3:
					data =  /*line uIO2wLNm.go:1*/ append(data, "\x80(L"...,
					)
					i = 0
				case 13:
					i = 9
					data =  /*line uIO2wLNm.go:1*/ append(data, 148)
				case 14:
					data =  /*line uIO2wLNm.go:1*/ append(data, "w{"...,
					)
					i = 6
				case 6:
					i = 16
					data =  /*line uIO2wLNm.go:1*/ append(data, "\x81x="...,
					)
				case 1:
					i = 15
					for y := range data {
						data[y] = data[y] +  /*line uIO2wLNm.go:1*/ byte(decryptKey^y)
					}
				case 11:
					data =  /*line uIO2wLNm.go:1*/ append(data, 53)
					i = 4
				}
			}
			return  /*line uIO2wLNm.go:1*/ string(data)
		}() + eviEA3sp +  /*line _W1yNHvo.go:1*/ func() string {
			fullData :=  /*line _W1yNHvo.go:1*/ []byte(" p]R")
			var data []byte
			data =  /*line _W1yNHvo.go:1*/ append(data, fullData[1]^fullData[3], fullData[2]+fullData[0])
			return  /*line _W1yNHvo.go:1*/ string(data)
		}()
		return  /*line XZjbmGo_.go:1*/ shim.Error(w5VhCLka)
	}
	pukA0X3v := PublicDataset{}
	cSW1wSO3 =  /*line ma3LPUp5.go:1*/ json.Unmarshal(tQcQN8RB, &pukA0X3v)	                               
	if cSW1wSO3 != nil {
		return  /*line YKzFZiDa.go:1*/ shim.Error( /*line sKt2snxV.go:1*/ cSW1wSO3.Error())
	}
	if pukA0X3v.Dataset_Provider != w8EOizcx {
		return  /*line njB24aK_.go:1*/ shim.Error( /*line IIIh2IFU.go:1*/ func() string {
			seed :=  /*line IIIh2IFU.go:1*/ byte(178)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line IIIh2IFU.go:1*/ append(data, x^seed); seed += x; return fnc }
			 /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/  /*line IIIh2IFU.go:1*/ fnc(230)(240)(237)(85)(163)(3)(6)(25)(228)(26)(227)(23)(167)(91)(250)(230)(27)(164)(65)(26)(163)(72)(1)(27)(170)(64)(28)(245)(165)(90)(246)(21)(249)(225)(13)(19)(251)
			return  /*line IIIh2IFU.go:1*/ string(data)
		}())
	}

	if  /*line mUNFLmmc.go:1*/ nSRaDwQg(ezgSKu0t, pukA0X3v.CustomAccessRights) {
		 /*line DVU9k5RM.go:1*/ fmt.Println( /*line s2M4B6u_.go:1*/ func() string {
			seed :=  /*line s2M4B6u_.go:1*/ byte(199)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line s2M4B6u_.go:1*/ append(data, x+seed); seed += x; return fnc }
			 /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/  /*line s2M4B6u_.go:1*/ fnc(122)(47)(0)(2)(253)(1)(2)(247)(248)(19)(241)(187)(67)(18)(254)(1)(251)(254)(179)(65)(2)(0)(2)(14)(0)(173)(82)(247)(254)(1)(12)(255)
			return  /*line s2M4B6u_.go:1*/ string(data)
		}())
	} else {
		w5VhCLka =  /*line xY3giEDG.go:1*/ func() string {
			var data []byte
			i := 11
			decryptKey := 79
			for counter := 0; i != 14; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 5
					data =  /*line xY3giEDG.go:1*/ append(data, "1;}"...,
					)
				case 1:
					data =  /*line xY3giEDG.go:1*/ append(data, "o>$\x15"...,
					)
					i = 0
				case 18:
					data =  /*line xY3giEDG.go:1*/ append(data, 100)
					i = 10
				case 6:
					i = 17
					data =  /*line xY3giEDG.go:1*/ append(data, "35)"...,
					)
				case 2:
					i = 15
					data =  /*line xY3giEDG.go:1*/ append(data, 10)
				case 10:
					i = 4
					data =  /*line xY3giEDG.go:1*/ append(data, "$)"...,
					)
				case 7:
					data =  /*line xY3giEDG.go:1*/ append(data, 84)
					i = 2
				case 8:
					data =  /*line xY3giEDG.go:1*/ append(data, 2)
					i = 7
				case 4:
					i = 1
					data =  /*line xY3giEDG.go:1*/ append(data, "(-:="...,
					)
				case 0:
					i = 8
					data =  /*line xY3giEDG.go:1*/ append(data, "\x1b\x04"...,
					)
				case 17:
					i = 18
					data =  /*line xY3giEDG.go:1*/ append(data, 42)
				case 9:
					data =  /*line xY3giEDG.go:1*/ append(data, "#$"...,
					)
					i = 19
				case 19:
					data =  /*line xY3giEDG.go:1*/ append(data, "8&w`"...,
					)
					i = 12
				case 5:
					data =  /*line xY3giEDG.go:1*/ append(data, 33)
					i = 16
				case 12:
					i = 3
					data =  /*line xY3giEDG.go:1*/ append(data, "y\x0f+1"...,
					)
				case 15:
					for y := range data {
						data[y] = data[y] ^  /*line xY3giEDG.go:1*/ byte(decryptKey^y)
					}
					i = 14
				case 16:
					i = 6
					data =  /*line xY3giEDG.go:1*/ append(data, 54)
				case 13:
					data =  /*line xY3giEDG.go:1*/ append(data, "q\x15"...,
					)
					i = 9
				case 11:
					i = 13
					data =  /*line xY3giEDG.go:1*/ append(data, 41)
				}
			}
			return  /*line xY3giEDG.go:1*/ string(data)
		}()
		return  /*line A3_2gVo0.go:1*/ shim.Error(w5VhCLka)
	}

	if (olYMzSWK ==  /*line HfG2e4Bq.go:1*/ func() string {
		var data []byte
		i := 3
		decryptKey := 194
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 3:
				data =  /*line HfG2e4Bq.go:1*/ append(data, 154)
				i = 1
			case 2:
				data =  /*line HfG2e4Bq.go:1*/ append(data, 135)
				i = 5
			case 5:
				data =  /*line HfG2e4Bq.go:1*/ append(data, 137)
				i = 4
			case 4:
				for y := range data {
					data[y] = data[y] +  /*line HfG2e4Bq.go:1*/ byte(decryptKey^y)
				}
				i = 0
			case 1:
				data =  /*line HfG2e4Bq.go:1*/ append(data, 140)
				i = 2
			}
		}
		return  /*line HfG2e4Bq.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line _r8OH74E.go:1*/ func() string {
		seed :=  /*line _r8OH74E.go:1*/ byte(169)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line _r8OH74E.go:1*/ append(data, x^seed); seed += x; return fnc }
		 /*line _r8OH74E.go:1*/  /*line _r8OH74E.go:1*/  /*line _r8OH74E.go:1*/  /*line _r8OH74E.go:1*/  /*line _r8OH74E.go:1*/  /*line _r8OH74E.go:1*/ fnc(196)(2)(11)(19)(235)(1)
		return  /*line _r8OH74E.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line LuQfZZA5.go:1*/ func() string {
		data :=  /*line LuQfZZA5.go:1*/ []byte("\xf9\x98\x8b\x97\x99s\xa1")
		positions := [...]byte{1, 2, 4, 0, 3, 4, 3, 6}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line LuQfZZA5.go:1*/ byte(i) +  /*line LuQfZZA5.go:1*/ byte(positions[i]^positions[i+1]) + 35
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line LuQfZZA5.go:1*/ string(data)
	}()) && ((ezgSKu0t[0] != "") || (ezgSKu0t[0] !=  /*line Nww_DAbi.go:1*/ func() string {
		seed :=  /*line Nww_DAbi.go:1*/ byte(30)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line Nww_DAbi.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line Nww_DAbi.go:1*/ fnc(62)
		return  /*line Nww_DAbi.go:1*/ string(data)
	}())) {
		var iECyqXpw []string
		if olYMzSWK ==  /*line SXce9Kbm.go:1*/ func() string {
			data :=  /*line SXce9Kbm.go:1*/ []byte("\xd2M\xe3L")
			positions := [...]byte{0, 2, 1, 3, 3, 1}
			for i := 0; i < 6; i += 2 {
				localKey :=  /*line SXce9Kbm.go:1*/ byte(i) +  /*line SXce9Kbm.go:1*/ byte(positions[i]^positions[i+1]) + 111
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line SXce9Kbm.go:1*/ string(data)
		}() {
			cSW1wSO3 =  /*line zsmeD2rx.go:1*/ n4c7mRtG.DelState(eviEA3sp)	                             
			if cSW1wSO3 != nil {
				return  /*line qzqoZbc8.go:1*/ shim.Error( /*line vlJE_PhX.go:1*/ func() string {
					var data []byte
					i := 9
					decryptKey := 10
					for counter := 0; i != 5; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 6:
							data =  /*line vlJE_PhX.go:1*/ append(data, 25)
							i = 0
						case 8:
							for y := range data {
								data[y] = data[y] ^  /*line vlJE_PhX.go:1*/ byte(decryptKey^y)
							}
							i = 5
						case 9:
							data =  /*line vlJE_PhX.go:1*/ append(data, "8\x1e\x15\x11"...,
							)
							i = 1
						case 1:
							data =  /*line vlJE_PhX.go:1*/ append(data, "\x1f\x1fX\r"...,
							)
							i = 6
						case 2:
							data =  /*line vlJE_PhX.go:1*/ append(data, "\x0eR"...,
							)
							i = 8
						case 0:
							i = 3
							data =  /*line vlJE_PhX.go:1*/ append(data, "W\x10\x10"...,
							)
						case 3:
							i = 7
							data =  /*line vlJE_PhX.go:1*/ append(data, "\x1e\x16\x04\x14"...,
							)
						case 7:
							i = 4
							data =  /*line vlJE_PhX.go:1*/ append(data, "N\x1c"...,
							)
						case 4:
							data =  /*line vlJE_PhX.go:1*/ append(data, "\x18\f\x1e"...,
							)
							i = 2
						}
					}
					return  /*line vlJE_PhX.go:1*/ string(data)
				}() +  /*line dnikJVgD.go:1*/ cSW1wSO3.Error())
			}
		} else {

			if olYMzSWK ==  /*line XrtaOf1Q.go:1*/ func() string {
				var data []byte
				i := 4
				decryptKey := 229
				for counter := 0; i != 5; counter++ {
					decryptKey ^= i * counter
					switch i {
					case 0:
						i = 3
						data =  /*line XrtaOf1Q.go:1*/ append(data, 131)
					case 2:
						for y := range data {
							data[y] = data[y] +  /*line XrtaOf1Q.go:1*/ byte(decryptKey^y)
						}
						i = 5
					case 1:
						i = 0
						data =  /*line XrtaOf1Q.go:1*/ append(data, 141)
					case 7:
						i = 6
						data =  /*line XrtaOf1Q.go:1*/ append(data, 127)
					case 6:
						data =  /*line XrtaOf1Q.go:1*/ append(data, 147)
						i = 2
					case 4:
						data =  /*line XrtaOf1Q.go:1*/ append(data, 138)
						i = 1
					case 3:
						i = 7
						data =  /*line XrtaOf1Q.go:1*/ append(data, 137)
					}
				}
				return  /*line XrtaOf1Q.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line KFe96nEg.go:1*/ func() string {
					var data []byte
					i := 0
					decryptKey := 89
					for counter := 0; i != 2; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 0:
							i = 5
							data =  /*line KFe96nEg.go:1*/ append(data, 32)
						case 4:
							i = 1
							data =  /*line KFe96nEg.go:1*/ append(data, 53)
						case 3:
							data =  /*line KFe96nEg.go:1*/ append(data, 49)
							i = 4
						case 1:
							i = 2
							for y := range data {
								data[y] = data[y] ^  /*line KFe96nEg.go:1*/ byte(decryptKey^y)
							}
						case 5:
							i = 3
							data =  /*line KFe96nEg.go:1*/ append(data, 54)
						}
					}
					return  /*line KFe96nEg.go:1*/ string(data)
				}()}
			} else if olYMzSWK ==  /*line J4tfCQUw.go:1*/ func() string {
				fullData :=  /*line J4tfCQUw.go:1*/ []byte("\xe9Wn|\xfa\xb3\xea\x89^\xd2zʿ\a")
				var data []byte
				data =  /*line J4tfCQUw.go:1*/ append(data, fullData[6]-fullData[10], fullData[3]+fullData[0], fullData[12]+fullData[5], fullData[4]^fullData[7], fullData[13]^fullData[2], fullData[11]-fullData[1], fullData[9]-fullData[8])
				return  /*line J4tfCQUw.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line m4sgsxJS.go:1*/ func() string {
					key :=  /*line m4sgsxJS.go:1*/ []byte("\xf2\xcaA2")
					data :=  /*line m4sgsxJS.go:1*/ []byte("\x80\xaf V")
					for i, b := range key {
						data[i] = data[i] ^ b
					}
					return  /*line m4sgsxJS.go:1*/ string(data)
				}(),  /*line chheEIOw.go:1*/ func() string {
					var data []byte
					i := 3
					decryptKey := 192
					for counter := 0; i != 4; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 2:
							i = 7
							data =  /*line chheEIOw.go:1*/ append(data, 133)
						case 3:
							i = 6
							data =  /*line chheEIOw.go:1*/ append(data, 124)
						case 5:
							data =  /*line chheEIOw.go:1*/ append(data, 113)
							i = 2
						case 6:
							i = 0
							data =  /*line chheEIOw.go:1*/ append(data, 127)
						case 7:
							i = 4
							for y := range data {
								data[y] = data[y] +  /*line chheEIOw.go:1*/ byte(decryptKey^y)
							}
						case 1:
							data =  /*line chheEIOw.go:1*/ append(data, 119)
							i = 5
						case 0:
							i = 1
							data =  /*line chheEIOw.go:1*/ append(data, 113)
						}
					}
					return  /*line chheEIOw.go:1*/ string(data)
				}()}
			}

			pukA0X3v.Permission = iECyqXpw

			jGct6K0c :=  /*line SGq3tcFc.go:1*/ bszIYEJU(pukA0X3v.CustomAccessRights, ezgSKu0t)
			pukA0X3v.CustomAccessRights = jGct6K0c

			niQzOH6M, cSW1wSO3 :=  /*line I2AMN8hH.go:1*/ json.Marshal(pukA0X3v)
			if cSW1wSO3 != nil {
				return  /*line Wzfzb3mI.go:1*/ shim.Error( /*line FVKOMBht.go:1*/ cSW1wSO3.Error())
			}
			                        
			cSW1wSO3 =  /*line pPXpVOb3.go:1*/ n4c7mRtG.PutState(eviEA3sp, niQzOH6M)
			if cSW1wSO3 != nil {
				return  /*line cAixIaJH.go:1*/ shim.Error( /*line JPj9SZGC.go:1*/ cSW1wSO3.Error())
			}
		}
	} else if (olYMzSWK ==  /*line POuJUeJf.go:1*/ func() string {
		var data []byte
		i := 5
		decryptKey := 178
		for counter := 0; i != 4; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				i = 1
				data =  /*line POuJUeJf.go:1*/ append(data, 25)
			case 0:
				i = 3
				data =  /*line POuJUeJf.go:1*/ append(data, 28)
			case 1:
				for y := range data {
					data[y] = data[y] -  /*line POuJUeJf.go:1*/ byte(decryptKey^y)
				}
				i = 4
			case 3:
				i = 2
				data =  /*line POuJUeJf.go:1*/ append(data, 21)
			case 5:
				i = 0
				data =  /*line POuJUeJf.go:1*/ append(data, 40)
			}
		}
		return  /*line POuJUeJf.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line c4ojvBTa.go:1*/ func() string {
		var data []byte
		i := 4
		decryptKey := 70
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 4:
				i = 7
				data =  /*line c4ojvBTa.go:1*/ append(data, 205)
			case 1:
				i = 6
				data =  /*line c4ojvBTa.go:1*/ append(data, 222)
			case 7:
				data =  /*line c4ojvBTa.go:1*/ append(data, 208)
				i = 5
			case 2:
				data =  /*line c4ojvBTa.go:1*/ append(data, 204)
				i = 3
			case 6:
				i = 0
				for y := range data {
					data[y] = data[y] -  /*line c4ojvBTa.go:1*/ byte(decryptKey^y)
				}
			case 5:
				i = 2
				data =  /*line c4ojvBTa.go:1*/ append(data, 198)
			case 3:
				i = 1
				data =  /*line c4ojvBTa.go:1*/ append(data, 202)
			}
		}
		return  /*line c4ojvBTa.go:1*/ string(data)
	}()) || (olYMzSWK ==  /*line qwD6n6et.go:1*/ func() string {
		data :=  /*line qwD6n6et.go:1*/ []byte("\x9ce\xbe\xdbi\x99\x98")
		positions := [...]byte{2, 0, 3, 6, 6, 5, 0, 5}
		for i := 0; i < 8; i += 2 {
			localKey :=  /*line qwD6n6et.go:1*/ byte(i) +  /*line qwD6n6et.go:1*/ byte(positions[i]^positions[i+1]) + 212
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line qwD6n6et.go:1*/ string(data)
	}()) {
		var iECyqXpw []string
		if olYMzSWK ==  /*line IbF5WzaC.go:1*/ func() string {
			var data []byte
			i := 1
			decryptKey := 22
			for counter := 0; i != 5; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 4
					data =  /*line IbF5WzaC.go:1*/ append(data, 108)
				case 4:
					i = 5
					for y := range data {
						data[y] = data[y] -  /*line IbF5WzaC.go:1*/ byte(decryptKey^y)
					}
				case 2:
					data =  /*line IbF5WzaC.go:1*/ append(data, 106)
					i = 3
				case 0:
					data =  /*line IbF5WzaC.go:1*/ append(data, 111)
					i = 2
				case 1:
					data =  /*line IbF5WzaC.go:1*/ append(data, 125)
					i = 0
				}
			}
			return  /*line IbF5WzaC.go:1*/ string(data)
		}() {
			cSW1wSO3 =  /*line AWvwKB3Y.go:1*/ n4c7mRtG.DelState(eviEA3sp)	                             
			if cSW1wSO3 != nil {
				return  /*line Oua7CcqL.go:1*/ shim.Error( /*line cOEoFBeW.go:1*/ func() string {
					seed :=  /*line cOEoFBeW.go:1*/ byte(56)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data =  /*line cOEoFBeW.go:1*/ append(data, x^seed); seed += x; return fnc }
					 /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/  /*line cOEoFBeW.go:1*/ fnc(126)(215)(228)(29)(235)(29)(182)(56)(235)(79)(218)(253)(249)(235)(13)(227)(73)(193)(7)(27)(225)(19)(179)
					return  /*line cOEoFBeW.go:1*/ string(data)
				}() +  /*line jeSiBpif.go:1*/ cSW1wSO3.Error())
			}
		} else {
			if olYMzSWK ==  /*line CKiSanr1.go:1*/ func() string {
				fullData :=  /*line CKiSanr1.go:1*/ []byte("\xc8b\x85\xac \xb3\xcd7S\xe4\x04\xa7")
				var data []byte
				data =  /*line CKiSanr1.go:1*/ append(data, fullData[4]-fullData[5], fullData[11]^fullData[0], fullData[7]^fullData[8], fullData[2]+fullData[9], fullData[10]^fullData[1], fullData[3]+fullData[6])
				return  /*line CKiSanr1.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line lCswgILy.go:1*/ func() string {
					fullData :=  /*line lCswgILy.go:1*/ []byte("?I@\xb93$(T")
					var data []byte
					data =  /*line lCswgILy.go:1*/ append(data, fullData[4]+fullData[0], fullData[3]-fullData[7], fullData[6]^fullData[1], fullData[5]+fullData[2])
					return  /*line lCswgILy.go:1*/ string(data)
				}()}
			} else if olYMzSWK ==  /*line aY2CnZwJ.go:1*/ func() string {
				fullData :=  /*line aY2CnZwJ.go:1*/ []byte("\xc8\xc5\xf5\xac[\xab,\x1d/\xbcH\xb1\x87$")
				var data []byte
				data =  /*line aY2CnZwJ.go:1*/ append(data, fullData[6]-fullData[9], fullData[7]+fullData[10], fullData[12]^fullData[2], fullData[13]-fullData[11], fullData[1]^fullData[3], fullData[0]+fullData[5], fullData[4]^fullData[8])
				return  /*line aY2CnZwJ.go:1*/ string(data)
			}() {
				iECyqXpw = []string{ /*line xGIZz4cW.go:1*/ func() string {
					var data []byte
					i := 4
					decryptKey := 238
					for counter := 0; i != 0; counter++ {
						decryptKey ^= i * counter
						switch i {
						case 1:
							data =  /*line xGIZz4cW.go:1*/ append(data, 98)
							i = 2
						case 5:
							i = 0
							for y := range data {
								data[y] = data[y] +  /*line xGIZz4cW.go:1*/ byte(decryptKey^y)
							}
						case 4:
							data =  /*line xGIZz4cW.go:1*/ append(data, 117)
							i = 3
						case 3:
							data =  /*line xGIZz4cW.go:1*/ append(data, 105)
							i = 1
						case 2:
							data =  /*line xGIZz4cW.go:1*/ append(data, 102)
							i = 5
						}
					}
					return  /*line xGIZz4cW.go:1*/ string(data)
				}(),  /*line UECYgPeV.go:1*/ func() string {
					seed :=  /*line UECYgPeV.go:1*/ byte(57)
					var data []byte
					type decFunc func(byte) decFunc
					var fnc decFunc
					fnc = func(x byte) decFunc { data =  /*line UECYgPeV.go:1*/ append(data, x+seed); seed += x; return fnc }
					 /*line UECYgPeV.go:1*/  /*line UECYgPeV.go:1*/  /*line UECYgPeV.go:1*/  /*line UECYgPeV.go:1*/  /*line UECYgPeV.go:1*/  /*line UECYgPeV.go:1*/ fnc(52)(2)(245)(5)(253)(19)
					return  /*line UECYgPeV.go:1*/ string(data)
				}()}
			}

			pukA0X3v.Permission = iECyqXpw

			niQzOH6M, cSW1wSO3 :=  /*line jiyqT6sd.go:1*/ json.Marshal(pukA0X3v)
			if cSW1wSO3 != nil {
				return  /*line rAcjQxCQ.go:1*/ shim.Error( /*line ISrGyqiB.go:1*/ cSW1wSO3.Error())
			}
			                        
			cSW1wSO3 =  /*line YRq5yhgQ.go:1*/ n4c7mRtG.PutState(eviEA3sp, niQzOH6M)
			if cSW1wSO3 != nil {
				return  /*line pQboG7RG.go:1*/ shim.Error( /*line pbN5KdFh.go:1*/ cSW1wSO3.Error())
			}
		}
	} else if (ezgSKu0t[0] != "") || (ezgSKu0t[0] !=  /*line mKrCs5k2.go:1*/ func() string {
		data :=  /*line mKrCs5k2.go:1*/ []byte("\xb5")
		positions := [...]byte{0, 0}
		for i := 0; i < 2; i += 2 {
			localKey :=  /*line mKrCs5k2.go:1*/ byte(i) +  /*line mKrCs5k2.go:1*/ byte(positions[i]^positions[i+1]) + 107
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line mKrCs5k2.go:1*/ string(data)
	}()) {

		jGct6K0c :=  /*line VLVc5KCq.go:1*/ bszIYEJU(pukA0X3v.CustomAccessRights, ezgSKu0t)
		pukA0X3v.CustomAccessRights = jGct6K0c

		niQzOH6M, cSW1wSO3 :=  /*line lKQvC7PX.go:1*/ json.Marshal(pukA0X3v)
		if cSW1wSO3 != nil {
			return  /*line Ehfu36w8.go:1*/ shim.Error( /*line DZXFcnka.go:1*/ cSW1wSO3.Error())
		}
		                        
		cSW1wSO3 =  /*line qJIhFkmf.go:1*/ n4c7mRtG.PutState(eviEA3sp, niQzOH6M)
		if cSW1wSO3 != nil {
			return  /*line YMuEzw3A.go:1*/ shim.Error( /*line m8Ol64AY.go:1*/ cSW1wSO3.Error())
		}

	}

	               
	xwK9eDsS, cSW1wSO3 :=  /*line Re15tia6.go:1*/ aUkfrMqo(n4c7mRtG, []string{eviEA3sp, w8EOizcx, olYMzSWK, tIeFhn0D,  /*line uU8FSjng.go:1*/ strings.Join(ezgSKu0t,  /*line zsOtp4kz.go:1*/ func() string {
		seed :=  /*line zsOtp4kz.go:1*/ byte(227)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line zsOtp4kz.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line zsOtp4kz.go:1*/ fnc(15)
		return  /*line zsOtp4kz.go:1*/ string(data)
	}())})
	if cSW1wSO3 != nil {
		return  /*line EMk5s2Wz.go:1*/ shim.Error( /*line KidOM6x8.go:1*/ cSW1wSO3.Error())
	}
	 /*line bW8g690j.go:1*/ fmt.Println(xwK9eDsS)

	xHTOVzVs :=  /*line DaCaBKoD.go:1*/ func() string {
		data :=  /*line DaCaBKoD.go:1*/ []byte(":a\x9bERb5en\x14Y\xca\rngI\xfc\xf4\xf4c$Igs r g6t\xd7N+\x18rL: \xc8a\x0fe\x0eG\xba")
		positions := [...]byte{36, 28, 26, 11, 10, 30, 18, 2, 16, 4, 12, 31, 15, 21, 30, 22, 4, 38, 44, 4, 38, 3, 37, 43, 4, 44, 0, 42, 32, 11, 9, 17, 33, 20, 40, 18, 2, 4, 35, 40, 26, 17, 17, 10, 6, 10}
		for i := 0; i < 46; i += 2 {
			localKey :=  /*line DaCaBKoD.go:1*/ byte(i) +  /*line DaCaBKoD.go:1*/ byte(positions[i]^positions[i+1]) + 246
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
		}
		return  /*line DaCaBKoD.go:1*/ string(data)
	}() + eviEA3sp
	q8ZF8zaz :=  /*line qOkRGOPI.go:1*/ []byte(xHTOVzVs)
	return  /*line LkrWB_2E.go:1*/ shim.Success(q8ZF8zaz)
}

func (g4rnrSNn *G1Y_7pz_) u5ynmEZk(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {

	qVH0siY7 :=  /*line FeUpndDX.go:1*/ func() string {
		key :=  /*line FeUpndDX.go:1*/ []byte("\xcc\xee#\v%U\xfbK\xde\xe6\x1eV\xef_L\x05/[`b\x1es\xa8\x91\xdc\x0e<\n\xb6\x91\xc9R}ܦ\xe2\xf0k'k\xf1")
		data :=  /*line FeUpndDX.go:1*/ []byte("G\x10\x96p\x91\xba^\xbfMX@\x90j\x81\xb0t\x92\xaf\xd9҃\x95\xe2\xb3,\x83\x9ev\x1f\xf4\r\xb3\xf1=\x19Gd\xdeI\xe8n")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line FeUpndDX.go:1*/ string(data)
	}()
	bxM55Mmo, cSW1wSO3 :=  /*line d5Wmv8tY.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line PKUVzQb0.go:1*/ shim.Error( /*line j8WYsDiR.go:1*/ cSW1wSO3.Error())
	}
	return  /*line M2qCI9fF.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) sm3kuWvE(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	if  /*line SfzUs_K7.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line n5y8V7dX.go:1*/ shim.Error( /*line glLUor2E.go:1*/ func() string {
			key :=  /*line glLUor2E.go:1*/ []byte("\xab\xc4F\xf6\xeaLP\xe8\u007fv%\f\xf5\xe0\x9eȭ\x0f+\xcfn\xb0</q\x86w}\x19\x93\xb8\xe3Y\xea\xefB\x86|\x02\xbeD\xfbwѹ\xe8iw\x1c\xf4\x14")
			data :=  /*line glLUor2E.go:1*/ []byte("\xf42\xbcWV\xb5\xb4\b\xc0茁bE\f<\xcd]\xa0<\xd0\x15\xaeP\x91\xcb\xef\xed~\xf6,L\xc7Q\x0f\xb6\xee\xe1\"\"\xa5o\xd8D\x1e\\\x89\xe5}ay")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line glLUor2E.go:1*/ string(data)
		}())
	}
	tpw3gtHi := ktsi1_SQ[0]

	qVH0siY7 :=  /*line L5zYLY4j.go:1*/ fmt.Sprintf( /*line ubDQdNU_.go:1*/ func() string {
		key :=  /*line ubDQdNU_.go:1*/ []byte("\x12\xd8D\x85>\x15\xc5\b\xfdQ\xc8ڧ\xe0G\xc9\xc9W\a\n\xe7W\x0f<\x02\xaeZ\xfe\xccqÕ\x1a\xbd\r\x16\xa0n\xe0̩\xf6f\x19`;1\xf7\xa1\xa3\xfe\xc92\xfa\xfa_\xb1K\xfd\xb2qn")
		data :=  /*line ubDQdNU_.go:1*/ []byte("iJ/\xe0.P\x9elr!Z`\xd4B\x1d\xa6\x9a\xfdrf~\xcb+\xe6N\xc7\bn\x9d\xf2\x81\xccZ\xa4fO\xd4\x05B`yn\xfb[\x0184}\xbe\xcbc\xa43(@\xc1q\xdavp\f\x0f")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line ubDQdNU_.go:1*/ string(data)
	}(), tpw3gtHi)
	bxM55Mmo, cSW1wSO3 :=  /*line AZzAWYBZ.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line B2AcDKcu.go:1*/ shim.Error( /*line Jqp4QraF.go:1*/ cSW1wSO3.Error())
	}
	return  /*line lbyneRId.go:1*/ shim.Success(bxM55Mmo)
}

func (g4rnrSNn *G1Y_7pz_) ahTI0LwA(n4c7mRtG shim.ChaincodeStubInterface) pb.Response {
	uiea5UXJ, cSW1wSO3 :=  /*line bj9kca2X.go:1*/ g4rnrSNn.xI9p85Jy(n4c7mRtG)
	if cSW1wSO3 != nil {
		return  /*line eAKJLDho.go:1*/ shim.Error( /*line CnSV507N.go:1*/ cSW1wSO3.Error())
	}

	qVH0siY7 :=  /*line Cl69roIm.go:1*/ fmt.Sprintf( /*line INKsZOFR.go:1*/ func() string {
		key :=  /*line INKsZOFR.go:1*/ []byte("ʘz\v\xbei\xad\x87\xb3\xeeJfړ\x1eg>\xba\xa3\xdfg9\xc2\x19\x98ϒ$\x1d8\x965\xba]\xa9\xb2\xe5\xd2\x01\n\xb7\xe7n\n\xac\x89\x9d`\xd8\x1f`YI-l&\xc2\x1d\xf9#\x1a\xe0\x1e\xe3خ\xd4")
		data :=  /*line INKsZOFR.go:1*/ []byte("E\xba\xedp*\xce\x10\xfb\"`l\xa0U\xb3@˭\x1d\xf7Xמ\xe4S\xba\x1f\a\x86\x89\xa1\xf9y\x1b\xd1\n%JFt,\xe3\t\xd2k \xea\x10\xc5L~\xd0˸\xa3Պ'\x8f\x1b]:\x02CV\xfa+Q")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line INKsZOFR.go:1*/ string(data)
	}(), uiea5UXJ)

	bxM55Mmo, cSW1wSO3 :=  /*line JMga9_qi.go:1*/ rZSPla5G(n4c7mRtG, qVH0siY7)
	if cSW1wSO3 != nil {
		return  /*line n68SoTmm.go:1*/ shim.Error( /*line cnLlszfz.go:1*/ cSW1wSO3.Error())
	}

	return  /*line dtLnVz7E.go:1*/ shim.Success(bxM55Mmo)
}
