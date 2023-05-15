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

type DatasetMetadataResponse struct {
	ObjectType	string		`json:"docType"`
	                           
	UsernameOfProvider	string		`json:"usernameOfProvider"`
	OrgOfProvider	string		`json:"orgOfProvider"`
	EmailOfProvider	string		`json:"emailOfProvider"`
	UsernameOfOwner	string		`json:"usernameOfOwner"`
	OrgOfOwner	string		`json:"orgOfOwner"`
	                        
	DatasetName	string		`json:"datasetName"`
	Identifier	string		`json:"identifier"`
	DatasetDescription	string		`json:"datasetDescription"`
	ReleaseDateTime	string		`json:"releaseDateTime"`
	ModificationDateTime	string		`json:"modificationDateTime"`
	                        
	ThemeCategory	string		`json:"themeCategory"`
	KeywordTag	string		`json:"keywordTag"`
	Language	string		`json:"language"`
	Distribution	string		`json:"distribution"`
	DataVelocity	int		`json:"dataVelocity"`
	SpatialGeographicCoverage	string		`json:"spatialGeographicCoverage"`
	TemporalCoverageStart	string		`json:"temporalCoverageStart"`
	TemporalCoverageEnd	string		`json:"temporalCoverageEnd"`
	IndustryDomain	string		`json:"industryDomain"`
	DataVolume	int		`json:"dataVolume"`
	Comments	string		`json:"comments"`
	                           
	AccessRights	[]string		`json:"accessRights"`
	ContractAgreementHash	string		`json:"contractAgreementHash"`
	ContractAgreementURL	string		`json:"contractAgreementURL"`
	TermsConditionsHash	string		`json:"termsConditionsHash"`
	TermsConditionsURL	string		`json:"termsConditionsURL"`
	CopyrightHash	string		`json:"copyrightHash"`
	CopyrightURL	string		`json:"copyrightsURL"`
	CustomAccessRights	[]string		`json:"customAccessRights"`
	                                
	Endpoint	string		`json:"endpoint"`
	Blockchain	string		`json:"blockchain"`
	Version	int		`json:"version"`
	HasAgent	string		`json:"hasAgent"`
	                      
	DataSourceMetadata	*DataSourceMetadata		`json:"dataSourceMetadata"`
}

type DataSourceMetadata struct {
	DataSourceID	string		`json:"dataSourceID"`
	DataSourceType	string		`json:"dataSourceType"`
	DataModels	*DataModelsStruct		`json:"dataModels"`
	DataProvided	*DataProvidedStruct		`json:"dataProvided"`
	Attributes	string		`json:"attributes"`
	Service	string		`json:"service"`
	ServicePath	string		`json:"servicePath"`
	Mapping	string		`json:"mapping"`
	DataportsDataModelandFormat	bool		`json:"isDataPortsCompatible"`
}

type DataModelsStruct struct {
	Type	string		`json:"type"`
	Value	string		`json:"value"`
	Metadata	string		`json:"metadata"`
}

type DataProvidedStruct struct {
	Type	string		`json:"type"`
	Value	string		`json:"value"`
	Metadata	string		`json:"metadata"`
}

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            
func _urwHavD(yZcJ7Fli []string, ibFxPbWb []string) []string {
	                                                    
	                                                                  
	z18U8WBV :=  /*line Zmqs6y7Q.go:1*/ make(map[string]bool)

	                                                                  
	for _, lFkgHiji := range yZcJ7Fli {
		if lFkgHiji != "" && lFkgHiji !=  /*line tuWMXWrz.go:1*/ func() string {
			seed :=  /*line tuWMXWrz.go:1*/ byte(13)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line tuWMXWrz.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line tuWMXWrz.go:1*/ fnc(45)
			return  /*line tuWMXWrz.go:1*/ string(data)
		}() {
			z18U8WBV[lFkgHiji] = true
		}
	}

	                                                                   
	                                                             
	for _, lFkgHiji := range ibFxPbWb {
		if lFkgHiji != "" && lFkgHiji !=  /*line kFJW8UvS.go:1*/ func() string {
			key :=  /*line kFJW8UvS.go:1*/ []byte("\xb9")
			data :=  /*line kFJW8UvS.go:1*/ []byte("\x99")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line kFJW8UvS.go:1*/ string(data)
		}() {
			z18U8WBV[lFkgHiji] = true
		}
	}

	                                                          
	                                                                               
	kOuQCykX :=  /*line GRz8jlNk.go:1*/ make([]string, 0,  /*line GWZsjJ4O.go:1*/ len(z18U8WBV))

	                                          
	for okjuZ7mN := range z18U8WBV {
		kOuQCykX =  /*line yx1wlaEh.go:1*/ append(kOuQCykX, okjuZ7mN)
	}

	                          
	return kOuQCykX
}

func nSRaDwQg(yZcJ7Fli []string, ibFxPbWb []string) bool {
	 /*line n_on6qZB.go:1*/ fmt.Println( /*line azW0Y8Zz.go:1*/ len(yZcJ7Fli))
	if  /*line DvLLMMWR.go:1*/ len(yZcJ7Fli) == 0 || yZcJ7Fli[0] == "" || yZcJ7Fli[0] ==  /*line bnB5zOe5.go:1*/ func() string {
		fullData :=  /*line bnB5zOe5.go:1*/ []byte("\x1b\xfb")
		var data []byte
		data =  /*line bnB5zOe5.go:1*/ append(data, fullData[0]-fullData[1])
		return  /*line bnB5zOe5.go:1*/ string(data)
	}() {
		return true
	}

	for _, j5VzWmWS := range yZcJ7Fli {
		bdOF7kqY := false
		for _, r5Nn_BDE := range ibFxPbWb {
			if j5VzWmWS == r5Nn_BDE {
				bdOF7kqY = true
				break
			}
		}
		if !bdOF7kqY {
			 /*line yN2bPB6E.go:1*/ fmt.Println( /*line PziXMFXu.go:1*/ func() string {
				fullData :=  /*line PziXMFXu.go:1*/ []byte("B73\xa8\xcbQ\xa8\xcd\xfe\xb2")
				var data []byte
				data =  /*line PziXMFXu.go:1*/ append(data, fullData[3]^fullData[8], fullData[9]-fullData[5], fullData[1]-fullData[4], fullData[2]+fullData[0], fullData[7]^fullData[6])
				return  /*line PziXMFXu.go:1*/ string(data)
			}(), j5VzWmWS,  /*line p9oJkTu6.go:1*/ func() string {
				seed :=  /*line p9oJkTu6.go:1*/ byte(22)
				var data []byte
				type decFunc func(byte) decFunc
				var fnc decFunc
				fnc = func(x byte) decFunc { data =  /*line p9oJkTu6.go:1*/ append(data, x^seed); seed += x; return fnc }
				 /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/  /*line p9oJkTu6.go:1*/ fnc(127)(251)(176)(33)(19)(6)(27)(236)(161)(80)(23)(248)(244)(16)(246)(15)(239)(29)(196)(51)(234)(31)(226)(11)(163)(66)(7)(10)(10)(163)(72)(1)(27)(170)(81)(253)(235)(30)(255)(170)(93)(255)(176)(37)(29)(235)(30)(255)(239)(29)(196)(51)(234)(31)(226)(11)
				return  /*line p9oJkTu6.go:1*/ string(data)
			}())
			return false
		}
	}

	return true
}

func bszIYEJU(dRPhXngW []string, fiBYbHVs []string) []string {
	                                                         
	bGI4x_39 :=  /*line HeeFPcaS.go:1*/ make(map[string]bool)
	for _, bT2slGjF := range fiBYbHVs {
		bGI4x_39[bT2slGjF] = true
	}

	                                                                                                       
	p3gcBS0_ :=  /*line Bp7Fp0Ez.go:1*/ make([]string, 0)
	for _, bT2slGjF := range dRPhXngW {
		if _, bdOF7kqY := bGI4x_39[bT2slGjF]; !bdOF7kqY {
			p3gcBS0_ =  /*line aiBlKpv1.go:1*/ append(p3gcBS0_, bT2slGjF)
		}
	}

	return p3gcBS0_
}

                                                    
func (g4rnrSNn *G1Y_7pz_) lu_3VIbC(n4c7mRtG shim.ChaincodeStubInterface) (string, error) {
	sM8l5NjA, v2IS02UJ, cSW1wSO3 :=  /*line vEAgLcMk.go:1*/ cid.GetAttributeValue(n4c7mRtG,  /*line LJFoqb_t.go:1*/ func() string {
		data :=  /*line LJFoqb_t.go:1*/ []byte("\x97s\x15\f\xe2ew\x92\x05<Use\x89")
		positions := [...]byte{4, 6, 2, 9, 3, 13, 3, 8, 7, 4, 0, 6, 8, 4, 4, 4, 6, 0}
		for i := 0; i < 18; i += 2 {
			localKey :=  /*line LJFoqb_t.go:1*/ byte(i) +  /*line LJFoqb_t.go:1*/ byte(positions[i]^positions[i+1]) + 108
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
		}
		return  /*line LJFoqb_t.go:1*/ string(data)
	}())

	if cSW1wSO3 != nil {
		return "", cSW1wSO3
	}

	if !v2IS02UJ {
		return "",  /*line CsXZ_m4q.go:1*/ errors.New( /*line xOoEBWia.go:1*/ func() string {
			key :=  /*line xOoEBWia.go:1*/ []byte("DU`?>\b#\xe8\xe2\xb2\x1d7h\xfb\xb0d\x05\\\x96\xe3\x00*\xe4¿a\x1cN\xf18w\x0fkN\xf1")
			data :=  /*line xOoEBWia.go:1*/ []byte("\xadȥ\xb7\xb2m\x95VC\x1er\xaa\xcdm\xd0\xc5y\xd0\bLb\x9fX'\xdfʏn^\xa1\xea\x82ԼX")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line xOoEBWia.go:1*/ string(data)
		}())
	}

	return sM8l5NjA, nil
}

                                                   
func (g4rnrSNn *G1Y_7pz_) rae8UYAQ(n4c7mRtG shim.ChaincodeStubInterface) (string, error) {
	xctzndN9, v2IS02UJ, cSW1wSO3 :=  /*line cG1cP4e4.go:1*/ cid.GetAttributeValue(n4c7mRtG,  /*line T6Ug0AW2.go:1*/ func() string {
		fullData :=  /*line T6Ug0AW2.go:1*/ []byte("\xccn\xfaa\xcc\xee-\x03\f@vP8ex.R\xe0\x00g`$")
		var data []byte
		data =  /*line T6Ug0AW2.go:1*/ append(data, fullData[0]-fullData[19], fullData[5]-fullData[10], fullData[11]^fullData[21], fullData[12]+fullData[6], fullData[17]-fullData[1], fullData[15]+fullData[9], fullData[18]+fullData[3], fullData[8]+fullData[20], fullData[16]-fullData[7], fullData[2]+fullData[14], fullData[4]-fullData[13])
		return  /*line T6Ug0AW2.go:1*/ string(data)
	}())

	if cSW1wSO3 != nil {
		return "", cSW1wSO3
	}

	if !v2IS02UJ {
		return "",  /*line wXDAEkLo.go:1*/ errors.New( /*line GOxK8MjM.go:1*/ func() string {
			fullData :=  /*line GOxK8MjM.go:1*/ []byte("\xc1E\xccޒ\xd4\x1b\xf8\"\x98B\x88;\xd3\x1e#=!̷V0\x143\xc3\x11\xd0\x15\rc:\xa5\xf1H\xa4O\x9b\x8b\xc0\xac\xb8g\x9fu\x9bBʧ\xad%^C\xbd\x9d\xc17CEZ\xaf\xb9\x1d \x1e")
			var data []byte
			data =  /*line GOxK8MjM.go:1*/ append(data, fullData[36]+fullData[46], fullData[19]+fullData[54], fullData[20]+fullData[14], fullData[4]+fullData[13], fullData[27]^fullData[41], fullData[43]^fullData[6], fullData[51]+fullData[63], fullData[0]^fullData[48], fullData[28]^fullData[10], fullData[22]+fullData[50], fullData[44]+fullData[18], fullData[45]-fullData[8], fullData[9]-fullData[55], fullData[47]-fullData[23], fullData[25]+fullData[29], fullData[59]-fullData[16], fullData[33]^fullData[17], fullData[61]+fullData[1], fullData[26]^fullData[31], fullData[62]-fullData[39], fullData[49]-fullData[38], fullData[56]-fullData[15], fullData[2]+fullData[53], fullData[37]^fullData[7], fullData[58]-fullData[30], fullData[21]-fullData[24], fullData[32]-fullData[11], fullData[40]-fullData[57], fullData[42]+fullData[5], fullData[34]-fullData[12], fullData[52]-fullData[35], fullData[3]^fullData[60])
			return  /*line GOxK8MjM.go:1*/ string(data)
		}())
	}

	return xctzndN9, nil
}

                                      
func (g4rnrSNn *G1Y_7pz_) y4dHt0L7(n4c7mRtG shim.ChaincodeStubInterface) (string, error) {
	cVKu8ZN2, v2IS02UJ, cSW1wSO3 :=  /*line dpFr5rl4.go:1*/ cid.GetAttributeValue(n4c7mRtG,  /*line yx_LQQ5Y.go:1*/ func() string {
		data :=  /*line yx_LQQ5Y.go:1*/ []byte("n\t\x15e")
		positions := [...]byte{2, 1, 1, 2}
		for i := 0; i < 4; i += 2 {
			localKey :=  /*line yx_LQQ5Y.go:1*/ byte(i) +  /*line yx_LQQ5Y.go:1*/ byte(positions[i]^positions[i+1]) + 208
			data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
		}
		return  /*line yx_LQQ5Y.go:1*/ string(data)
	}())

	if cSW1wSO3 != nil {
		return "", cSW1wSO3
	}

	if !v2IS02UJ {
		return "",  /*line mvllahgz.go:1*/ errors.New( /*line qkHazjWd.go:1*/ func() string {
			fullData :=  /*line qkHazjWd.go:1*/ []byte("kV\x90\xadBlU`\x10\xf6\x01:\x8d=\xff\xe8\xb6о\xf3\xe1\x14\xe4\xd41\x1e\x03d\xfd;\xd8\xd6\x01\x96\xb0Ӂk\x92\x1e\u007fy\xaf\xd9\xec#\x8b;Ϝ")
			var data []byte
			data =  /*line qkHazjWd.go:1*/ append(data, fullData[43]-fullData[37], fullData[40]^fullData[25], fullData[18]^fullData[35], fullData[17]-fullData[0], fullData[9]-fullData[31], fullData[28]-fullData[49], fullData[21]+fullData[7], fullData[39]+fullData[1], fullData[42]-fullData[13], fullData[22]^fullData[12], fullData[48]^fullData[3], fullData[47]+fullData[11], fullData[36]+fullData[19], fullData[32]+fullData[27], fullData[2]^fullData[34], fullData[33]^fullData[14], fullData[24]+fullData[4], fullData[26]^fullData[45], fullData[5]+fullData[10], fullData[6]-fullData[44], fullData[15]+fullData[46], fullData[38]^fullData[20], fullData[8]^fullData[41], fullData[30]^fullData[16], fullData[29]-fullData[23])
			return  /*line qkHazjWd.go:1*/ string(data)
		}())
	}

	return cVKu8ZN2, nil
}

                                         
func (g4rnrSNn *G1Y_7pz_) hDKHUUS5(n4c7mRtG shim.ChaincodeStubInterface) (string, error) {
	gN6Bx3KM, v2IS02UJ, cSW1wSO3 :=  /*line _9oEdJ0v.go:1*/ cid.GetAttributeValue(n4c7mRtG,  /*line zK2M03NY.go:1*/ func() string {
		var data []byte
		i := 6
		decryptKey := 12
		for counter := 0; i != 0; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 2:
				data =  /*line zK2M03NY.go:1*/ append(data, 89)
				i = 3
			case 6:
				i = 7
				data =  /*line zK2M03NY.go:1*/ append(data, 93)
			case 5:
				data =  /*line zK2M03NY.go:1*/ append(data, 85)
				i = 4
			case 7:
				i = 8
				data =  /*line zK2M03NY.go:1*/ append(data, 94)
			case 4:
				i = 0
				for y := range data {
					data[y] = data[y] +  /*line zK2M03NY.go:1*/ byte(decryptKey^y)
				}
			case 1:
				i = 5
				data =  /*line zK2M03NY.go:1*/ append(data, 90)
			case 8:
				data =  /*line zK2M03NY.go:1*/ append(data, 94)
				i = 2
			case 3:
				data =  /*line zK2M03NY.go:1*/ append(data, 79)
				i = 1
			}
		}
		return  /*line zK2M03NY.go:1*/ string(data)
	}())

	if cSW1wSO3 != nil {
		return "", cSW1wSO3
	}

	if !v2IS02UJ {
		return "",  /*line VoHia3QQ.go:1*/ errors.New( /*line GAd1S3PJ.go:1*/ func() string {
			key :=  /*line GAd1S3PJ.go:1*/ []byte("-\rLk\xe0?\xc2Jw\x1f\xf2\x12\xf7\xcdH\x90\xf0\x02\xf1\xb8|\xf61`&\xa9\x8c\xad")
			data :=  /*line GAd1S3PJ.go:1*/ []byte("\xa0\x82\xbe\xd9A\xac'jؓf\x84`/\xbd\x04U\"Z+\x9cc\x9aә\x12\xfa\x14")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line GAd1S3PJ.go:1*/ string(data)
		}())
	}

	return gN6Bx3KM, nil
}

                                              
func (g4rnrSNn *G1Y_7pz_) vNURlL_g(n4c7mRtG shim.ChaincodeStubInterface) (string, error) {
	dJwgW2jL, v2IS02UJ, cSW1wSO3 :=  /*line NOTJPNWh.go:1*/ cid.GetAttributeValue(n4c7mRtG,  /*line CP3DYeLa.go:1*/ func() string {
		var data []byte
		i := 0
		decryptKey := 217
		for counter := 0; i != 3; counter++ {
			decryptKey ^= i * counter
			switch i {
			case 6:
				i = 5
				data =  /*line CP3DYeLa.go:1*/ append(data, "\x9e\x9e"...,
				)
			case 2:
				data =  /*line CP3DYeLa.go:1*/ append(data, "\x9d\x85"...,
				)
				i = 1
			case 0:
				data =  /*line CP3DYeLa.go:1*/ append(data, "\x94\x98"...,
				)
				i = 4
			case 1:
				data =  /*line CP3DYeLa.go:1*/ append(data, "\xa1\x97"...,
				)
				i = 6
			case 4:
				data =  /*line CP3DYeLa.go:1*/ append(data, "\x8e\x89\x8f\x8b"...,
				)
				i = 2
			case 5:
				i = 3
				for y := range data {
					data[y] = data[y] +  /*line CP3DYeLa.go:1*/ byte(decryptKey^y)
				}
			}
		}
		return  /*line CP3DYeLa.go:1*/ string(data)
	}())

	if cSW1wSO3 != nil {
		return "", cSW1wSO3
	}

	if !v2IS02UJ {
		return "",  /*line zfoqYpGr.go:1*/ errors.New( /*line QxSJi6Fo.go:1*/ func() string {
			fullData :=  /*line QxSJi6Fo.go:1*/ []byte("[\xf24_U3\xb9\x1d\x1ew>=\xb1\x15_Z\x0e\x06ZX\x1e\x9d\xb0\xf1\x87\xd7\xfd\xae\x175U>\xcf[m\xa5\xcb\xff\xdd/\x0fB\x81MG)+>\x90Ir:в\x8b\x1a\xb5\x9c\x06\x0fp0\xbc7\xa9x")
			var data []byte
			data =  /*line QxSJi6Fo.go:1*/ append(data, fullData[4]^fullData[51], fullData[42]-fullData[40], fullData[43]+fullData[55], fullData[52]^fullData[12], fullData[62]+fullData[53], fullData[18]-fullData[23], fullData[32]^fullData[56], fullData[26]-fullData[57], fullData[3]+fullData[13], fullData[2]+fullData[29], fullData[28]^fullData[65], fullData[45]^fullData[44], fullData[11]-fullData[7], fullData[0]+fullData[58], fullData[54]^fullData[37], fullData[25]+fullData[21], fullData[35]-fullData[5], fullData[24]-fullData[8], fullData[64]+fullData[6], fullData[63]+fullData[10], fullData[14]^fullData[46], fullData[41]-fullData[38], fullData[27]+fullData[50], fullData[59]+fullData[15], fullData[30]+fullData[20], fullData[48]-fullData[60], fullData[47]+fullData[39], fullData[33]+fullData[16], fullData[36]-fullData[19], fullData[34]+fullData[17], fullData[9]+fullData[1], fullData[31]+fullData[61], fullData[22]-fullData[49])
			return  /*line QxSJi6Fo.go:1*/ string(data)
		}())
	}

	return dJwgW2jL, nil
}

                                          
func (g4rnrSNn *G1Y_7pz_) xI9p85Jy(n4c7mRtG shim.ChaincodeStubInterface) (string, error) {
	igyLs9co, v2IS02UJ, cSW1wSO3 :=  /*line PqS1_nWz.go:1*/ cid.GetAttributeValue(n4c7mRtG,  /*line f_lQd3lZ.go:1*/ func() string {
		fullData :=  /*line f_lQd3lZ.go:1*/ []byte("\x95\xc1J\xf5\x9b\xd6)p\x04a\xac\\\x02i\xf4\xb9")
		var data []byte
		data =  /*line f_lQd3lZ.go:1*/ append(data, fullData[5]-fullData[9], fullData[2]+fullData[6], fullData[15]+fullData[10], fullData[12]^fullData[7], fullData[3]^fullData[4], fullData[0]^fullData[14], fullData[8]+fullData[13], fullData[1]-fullData[11])
		return  /*line f_lQd3lZ.go:1*/ string(data)
	}())

	if cSW1wSO3 != nil {
		return "", cSW1wSO3
	}

	if !v2IS02UJ {
		return "",  /*line dJsAigEp.go:1*/ errors.New( /*line VigQ8ayW.go:1*/ func() string {
			fullData :=  /*line VigQ8ayW.go:1*/ []byte("Q\x06\x81\xdb\n\x92\xbe\x00\xffeU\xe7\xb43\xd8\xea\xe5\xab\xc3\xe7g\x8c\x9c7\x85\xb3ǝ,\x0eրQ\xa9\x14\xfdj\x8b$݂\x93gGB\x04e\x8f\x9a\bB\x8f\xba\xa1\tq \xf2")
			var data []byte
			data =  /*line VigQ8ayW.go:1*/ append(data, fullData[13]+fullData[44], fullData[21]+fullData[11], fullData[15]-fullData[24], fullData[37]+fullData[19], fullData[25]^fullData[39], fullData[35]^fullData[22], fullData[9]^fullData[49], fullData[8]^fullData[48], fullData[53]^fullData[2], fullData[4]-fullData[33], fullData[23]-fullData[18], fullData[40]+fullData[57], fullData[31]-fullData[29], fullData[10]+fullData[34], fullData[14]^fullData[52], fullData[32]^fullData[38], fullData[16]+fullData[51], fullData[30]+fullData[47], fullData[0]^fullData[55], fullData[1]-fullData[27], fullData[54]+fullData[36], fullData[20]-fullData[43], fullData[5]+fullData[3], fullData[17]-fullData[50], fullData[41]-fullData[56], fullData[26]^fullData[12], fullData[45]+fullData[46], fullData[28]-fullData[6], fullData[42]-fullData[7])
			return  /*line VigQ8ayW.go:1*/ string(data)
		}())
	}

	return igyLs9co, nil
}

func (g4rnrSNn *G1Y_7pz_) dHINeXgX(n4c7mRtG shim.ChaincodeStubInterface) (string, error) {
	hGBwBh29, v2IS02UJ, cSW1wSO3 :=  /*line ChNzUUkM.go:1*/ cid.GetAttributeValue(n4c7mRtG,  /*line KFwnh4Ob.go:1*/ func() string {
		key :=  /*line KFwnh4Ob.go:1*/ []byte("ko\xd5\xe1T\xdc\xd2\xf6\xdc8W\x82\xc4")
		data :=  /*line KFwnh4Ob.go:1*/ []byte("\x0f\x0e\xa1\x80$\xb3\xa0\x82\xafj8\xee\xa1")
		for i, b := range key {
			data[i] = data[i] ^ b
		}
		return  /*line KFwnh4Ob.go:1*/ string(data)
	}())

	if cSW1wSO3 != nil {
		return "", cSW1wSO3
	}

	if !v2IS02UJ {
		return "",  /*line z7DdGb_o.go:1*/ errors.New( /*line dzhP_CAm.go:1*/ func() string {
			var data []byte
			i := 7
			decryptKey := 11
			for counter := 0; i != 10; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 4:
					i = 10
					for y := range data {
						data[y] = data[y] -  /*line dzhP_CAm.go:1*/ byte(decryptKey^y)
					}
				case 1:
					data =  /*line dzhP_CAm.go:1*/ append(data, "\xdc\xd4"...,
					)
					i = 13
				case 2:
					i = 0
					data =  /*line dzhP_CAm.go:1*/ append(data, "\xe8\xeb\xec"...,
					)
				case 8:
					data =  /*line dzhP_CAm.go:1*/ append(data, 219)
					i = 4
				case 7:
					i = 9
					data =  /*line dzhP_CAm.go:1*/ append(data, "\xde\xdc\xda"...,
					)
				case 6:
					data =  /*line dzhP_CAm.go:1*/ append(data, "\xd6\xce"...,
					)
					i = 12
				case 11:
					data =  /*line dzhP_CAm.go:1*/ append(data, "\x9d\xeb"...,
					)
					i = 2
				case 0:
					i = 8
					data =  /*line dzhP_CAm.go:1*/ append(data, "\xe3\xe9"...,
					)
				case 3:
					data =  /*line dzhP_CAm.go:1*/ append(data, "\xca\xde\xdf"...,
					)
					i = 6
				case 13:
					data =  /*line dzhP_CAm.go:1*/ append(data, "Ƃ\xcc"...,
					)
					i = 5
				case 9:
					data =  /*line dzhP_CAm.go:1*/ append(data, "Ԉ"...,
					)
					i = 3
				case 5:
					data =  /*line dzhP_CAm.go:1*/ append(data, 239)
					i = 11
				case 12:
					i = 1
					data =  /*line dzhP_CAm.go:1*/ append(data, 200)
				}
			}
			return  /*line dzhP_CAm.go:1*/ string(data)
		}())
	}

	return hGBwBh29, nil
}

func (g4rnrSNn *G1Y_7pz_) zSzRgzTF(n4c7mRtG shim.ChaincodeStubInterface) (string, error) {
	hGBwBh29, v2IS02UJ, cSW1wSO3 :=  /*line ocI8Uuub.go:1*/ cid.GetAttributeValue(n4c7mRtG,  /*line RKFhKHEN.go:1*/ func() string {
		key :=  /*line RKFhKHEN.go:1*/ []byte("\xe9\xc0\u007f\x98\x8c")
		data :=  /*line RKFhKHEN.go:1*/ []byte("N-\xe0\x01\xf8")
		for i, b := range key {
			data[i] = data[i] - b
		}
		return  /*line RKFhKHEN.go:1*/ string(data)
	}())

	if cSW1wSO3 != nil {
		return "", cSW1wSO3
	}

	if !v2IS02UJ {
		return "",  /*line z9dv8VuK.go:1*/ errors.New( /*line fCaDdkIw.go:1*/ func() string {
			var data []byte
			i := 6
			decryptKey := 200
			for counter := 0; i != 0; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 9
					data =  /*line fCaDdkIw.go:1*/ append(data, "9-"...,
					)
				case 5:
					i = 2
					data =  /*line fCaDdkIw.go:1*/ append(data, "6\xeb"...,
					)
				case 10:
					i = 0
					for y := range data {
						data[y] = data[y] -  /*line fCaDdkIw.go:1*/ byte(decryptKey^y)
					}
				case 11:
					i = 8
					data =  /*line fCaDdkIw.go:1*/ append(data, "%\xe1"...,
					)
				case 1:
					data =  /*line fCaDdkIw.go:1*/ append(data, 62)
					i = 10
				case 7:
					i = 4
					data =  /*line fCaDdkIw.go:1*/ append(data, "JC"...,
					)
				case 12:
					i = 3
					data =  /*line fCaDdkIw.go:1*/ append(data, 58)
				case 6:
					i = 5
					data =  /*line fCaDdkIw.go:1*/ append(data, "3<-6"...,
					)
				case 9:
					i = 11
					data =  /*line fCaDdkIw.go:1*/ append(data, "'77"...,
					)
				case 4:
					data =  /*line fCaDdkIw.go:1*/ append(data, "NKBD"...,
					)
					i = 1
				case 8:
					i = 7
					data =  /*line fCaDdkIw.go:1*/ append(data, "GR\xfc"...,
					)
				case 2:
					data =  /*line fCaDdkIw.go:1*/ append(data, ")="...,
					)
					i = 12
				}
			}
			return  /*line fCaDdkIw.go:1*/ string(data)
		}())
	}

	return hGBwBh29, nil
}

func rZSPla5G(n4c7mRtG shim.ChaincodeStubInterface, qVH0siY7 string) ([]byte, error) {

	 /*line HJjBHoEP.go:1*/ fmt.Printf( /*line VGqOvtkJ.go:1*/ func() string {
		seed :=  /*line VGqOvtkJ.go:1*/ byte(90)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line VGqOvtkJ.go:1*/ append(data, x+seed); seed += x; return fnc }
		 /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/  /*line VGqOvtkJ.go:1*/ fnc(211)(243)(71)(254)(15)(221)(36)(240)(13)(7)(217)(19)(14)(2)(247)(8)(210)(41)(3)(223)(36)(240)(13)(7)(218)(33)(254)(247)(5)(249)(185)(81)(4)(240)(13)(7)(218)(33)(254)(247)(5)(249)(211)(208)(27)(78)(151)
		return  /*line VGqOvtkJ.go:1*/ string(data)
	}(), qVH0siY7)

	sTqtwvtR, cSW1wSO3 :=  /*line Xwth5pcb.go:1*/ n4c7mRtG.GetQueryResult(qVH0siY7)
	if cSW1wSO3 != nil {
		return nil, cSW1wSO3
	}
	defer  /*line ykr41L2H.go:1*/ sTqtwvtR.Close()

	hZwbBCnD, cSW1wSO3 :=  /*line EpNfI5J9.go:1*/ erPVaOf5(sTqtwvtR)
	if cSW1wSO3 != nil {
		return nil, cSW1wSO3
	}

	 /*line J2PZBFVl.go:1*/ fmt.Printf( /*line LcxmWg7j.go:1*/ func() string {
		fullData :=  /*line LcxmWg7j.go:1*/ []byte("\n\t\xf6rw!JD\xc4BgxRf\xabū\x8bc\xe4\x9c\xdc\xfa\xc7\xe1m\xefͫ\x11\x1dC\x99\xcd\u007f`\r\xbe\n\xc0\xa6\xff\xef\xf8;\x87\v\xa8\nT\\3\xf4\x0eFB\xa2w&\xa7UW\x15\x8f\x87\xf6/h~\xa5\xa6\x95\xadO\xf49\xe0q\x85́\x04=\x83ݙ8\xea\x1as\xa4\xce\xda\xc3")
		var data []byte
		data =  /*line LcxmWg7j.go:1*/ append(data, fullData[39]+fullData[25], fullData[45]^fullData[59], fullData[65]-fullData[63], fullData[21]-fullData[4], fullData[78]+fullData[42], fullData[36]+fullData[7], fullData[62]+fullData[35], fullData[88]^fullData[34], fullData[41]+fullData[89], fullData[22]-fullData[80], fullData[43]-fullData[40], fullData[47]^fullData[33], fullData[12]^fullData[5], fullData[60]-fullData[76], fullData[23]^fullData[28], fullData[44]+fullData[75], fullData[27]-fullData[64], fullData[85]^fullData[2], fullData[13]-fullData[74], fullData[57]+fullData[92], fullData[9]-fullData[79], fullData[38]-fullData[69], fullData[91]+fullData[90], fullData[29]+fullData[67], fullData[50]-fullData[1], fullData[37]-fullData[6], fullData[15]+fullData[72], fullData[54]^fullData[66], fullData[16]-fullData[82], fullData[77]-fullData[0], fullData[14]-fullData[17], fullData[19]^fullData[71], fullData[68]^fullData[46], fullData[58]^fullData[31], fullData[73]-fullData[84], fullData[83]-fullData[48], fullData[18]+fullData[26], fullData[93]+fullData[56], fullData[10]-fullData[52], fullData[51]+fullData[55], fullData[86]^fullData[49], fullData[30]+fullData[61], fullData[70]^fullData[20], fullData[81]^fullData[53], fullData[24]^fullData[8], fullData[32]^fullData[87], fullData[3]^fullData[11])
		return  /*line LcxmWg7j.go:1*/ string(data)
	}(),  /*line k0yjW8EG.go:1*/ hZwbBCnD.String())

	return  /*line RrCzTYAm.go:1*/ hZwbBCnD.Bytes(), nil
}

func erPVaOf5(sTqtwvtR shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	                                                 
	var hZwbBCnD bytes.Buffer
	 /*line Xb0bcApr.go:1*/ hZwbBCnD.WriteString( /*line JeZv0BhS.go:1*/ func() string {
		key :=  /*line JeZv0BhS.go:1*/ []byte("n")
		data :=  /*line JeZv0BhS.go:1*/ []byte("\xed")
		for i, b := range key {
			data[i] = data[i] + b
		}
		return  /*line JeZv0BhS.go:1*/ string(data)
	}())

	b8G2sCOe := false
	for  /*line fIwXWcO4.go:1*/ sTqtwvtR.HasNext() {
		z7JsnEVe, cSW1wSO3 :=  /*line W_Yd9PlM.go:1*/ sTqtwvtR.Next()
		if cSW1wSO3 != nil {
			return nil, cSW1wSO3
		}
		                                                                           
		if b8G2sCOe {
			 /*line DWI0CTCb.go:1*/ hZwbBCnD.WriteString( /*line ZTJCF4_w.go:1*/ func() string {
				fullData :=  /*line ZTJCF4_w.go:1*/ []byte("\xc5\xe9")
				var data []byte
				data =  /*line ZTJCF4_w.go:1*/ append(data, fullData[1]^fullData[0])
				return  /*line ZTJCF4_w.go:1*/ string(data)
			}())
		}
		 /*line mRrtM62L.go:1*/ hZwbBCnD.WriteString( /*line B8RXGLS4.go:1*/ func() string {
			key :=  /*line B8RXGLS4.go:1*/ []byte("\x9dz\x92\xe2\"f\x9f")
			data :=  /*line B8RXGLS4.go:1*/ []byte("ި\xb9\x83W\xbc\x9b")
			for i, b := range key {
				data[i] = data[i] + b
			}
			return  /*line B8RXGLS4.go:1*/ string(data)
		}())
		 /*line g2URZLY5.go:1*/ hZwbBCnD.WriteString( /*line LWUr0uYo.go:1*/ func() string {
			fullData :=  /*line LWUr0uYo.go:1*/ []byte("\xd9\xfb")
			var data []byte
			data =  /*line LWUr0uYo.go:1*/ append(data, fullData[1]^fullData[0])
			return  /*line LWUr0uYo.go:1*/ string(data)
		}())
		 /*line HMTstO30.go:1*/ hZwbBCnD.WriteString(z7JsnEVe.Key)
		 /*line hsepJt77.go:1*/ hZwbBCnD.WriteString( /*line Do6WeOyz.go:1*/ func() string {
			fullData :=  /*line Do6WeOyz.go:1*/ []byte("\x11\xef")
			var data []byte
			data =  /*line Do6WeOyz.go:1*/ append(data, fullData[0]-fullData[1])
			return  /*line Do6WeOyz.go:1*/ string(data)
		}())

		 /*line IK0RY0kS.go:1*/ hZwbBCnD.WriteString( /*line W3zoGJbn.go:1*/ func() string {
			fullData :=  /*line W3zoGJbn.go:1*/ []byte("J\x0e4h\xf8)\x1dd\xf8\x81\xef\xfc\xcb\xe4Y(\xb5R\xf6\xe9{\xa4")
			var data []byte
			data =  /*line W3zoGJbn.go:1*/ append(data, fullData[4]+fullData[2], fullData[8]+fullData[15], fullData[12]^fullData[19], fullData[18]-fullData[21], fullData[13]+fullData[9], fullData[5]^fullData[0], fullData[6]+fullData[17], fullData[7]+fullData[1], fullData[3]+fullData[11], fullData[14]^fullData[20], fullData[10]-fullData[16])
			return  /*line W3zoGJbn.go:1*/ string(data)
		}())
		                                             
		 /*line WMzhdGKJ.go:1*/ hZwbBCnD.WriteString( /*line Y5LJwD_B.go:1*/ string(z7JsnEVe.Value))
		 /*line nNz6XrTO.go:1*/ hZwbBCnD.WriteString( /*line EX7R0T6e.go:1*/ func() string {
			seed :=  /*line EX7R0T6e.go:1*/ byte(95)
			var data []byte
			type decFunc func(byte) decFunc
			var fnc decFunc
			fnc = func(x byte) decFunc { data =  /*line EX7R0T6e.go:1*/ append(data, x-seed); seed += x; return fnc }
			 /*line EX7R0T6e.go:1*/ fnc(220)
			return  /*line EX7R0T6e.go:1*/ string(data)
		}())
		b8G2sCOe = true
	}
	 /*line nTxJF0oT.go:1*/ hZwbBCnD.WriteString( /*line pY9InH2u.go:1*/ func() string {
		seed :=  /*line pY9InH2u.go:1*/ byte(83)
		var data []byte
		type decFunc func(byte) decFunc
		var fnc decFunc
		fnc = func(x byte) decFunc { data =  /*line pY9InH2u.go:1*/ append(data, x-seed); seed += x; return fnc }
		 /*line pY9InH2u.go:1*/ fnc(176)
		return  /*line pY9InH2u.go:1*/ string(data)
	}())

	return &hZwbBCnD, nil
}

func (g4rnrSNn *G1Y_7pz_) cBvtAt_z(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var hSHnto8d, w5VhCLka string
	var cSW1wSO3 error

	if  /*line EL0lNydN.go:1*/ len(ktsi1_SQ) != 1 {
		return  /*line ojib9wUk.go:1*/ shim.Error( /*line zFJS1Uf6.go:1*/ func() string {
			fullData :=  /*line zFJS1Uf6.go:1*/ []byte("\xb2\x88\x9c\x91\xfb\xa3\xf7\x90\x87\x14\xc42Lz\xfd\xe4\xec^L\xb0\xe7\x8b\xc7W\x85\xe28\xba7ff!\x01d\xa7\xe3^\x83($\xdat\xad)\xb8\xa8Xn\xc0\x1d/\x8b`\x16B<*\r\x18\xc1\xb4Q\x19\xaa\xa1\xdd\xf3\x00p\xf6i\x14\xfc\xb1\x97%\x83\xab\x96\xff\xbe_{\xb3Y\xc5L\xac\v\xf9\x81\x17\x17\x87\xb2\x92\x99ٴa%\xfb,\x8b\xa1\t\xfdBC`\xfd\xd6\xda6\n\x0e=$")
			var data []byte
			data =  /*line zFJS1Uf6.go:1*/ append(data, fullData[54]^fullData[88], fullData[57]+fullData[99], fullData[106]+fullData[29], fullData[66]^fullData[2], fullData[9]^fullData[30], fullData[20]+fullData[103], fullData[85]-fullData[52], fullData[1]-fullData[75], fullData[38]-fullData[98], fullData[113]-fullData[53], fullData[58]-fullData[63], fullData[94]-fullData[116], fullData[3]^fullData[72], fullData[114]+fullData[46], fullData[80]-fullData[84], fullData[96]+fullData[97], fullData[90]^fullData[104], fullData[71]^fullData[82], fullData[68]+fullData[69], fullData[105]^fullData[43], fullData[32]^fullData[109], fullData[83]^fullData[59], fullData[10]+fullData[5], fullData[81]^fullData[56], fullData[31]^fullData[18], fullData[51]+fullData[112], fullData[40]^fullData[60], fullData[86]^fullData[26], fullData[15]^fullData[74], fullData[55]-fullData[115], fullData[34]-fullData[93], fullData[110]^fullData[44], fullData[4]-fullData[37], fullData[62]^fullData[70], fullData[17]-fullData[89], fullData[13]-fullData[92], fullData[6]^fullData[76], fullData[79]^fullData[78], fullData[67]^fullData[47], fullData[42]+fullData[27], fullData[21]^fullData[77], fullData[64]-fullData[50], fullData[8]^fullData[25], fullData[73]+fullData[48], fullData[0]^fullData[22], fullData[107]-fullData[65], fullData[35]+fullData[7], fullData[39]-fullData[19], fullData[12]+fullData[14], fullData[111]^fullData[95], fullData[41]+fullData[87], fullData[108]^fullData[28], fullData[91]-fullData[45], fullData[100]+fullData[101], fullData[16]+fullData[24], fullData[117]+fullData[61], fullData[11]^fullData[23], fullData[36]^fullData[102], fullData[33]^fullData[49])
			return  /*line zFJS1Uf6.go:1*/ string(data)
		}())
	}

	hSHnto8d = ktsi1_SQ[0]
	mGZijNgp, cSW1wSO3 :=  /*line pT0G1PvX.go:1*/ n4c7mRtG.GetState(hSHnto8d)	                                      
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line AeteDdh3.go:1*/ func() string {
			key :=  /*line AeteDdh3.go:1*/ []byte("\xb3S!\xc5\xf4Υ\xe7\xb1 \x8c\x86W?\x8f7\xfb\xb4\xe0\xc5\v\xb8D\x9a\x89uA\b\xc8\xdam7hC")
			data :=  /*line AeteDdh3.go:1*/ []byte(".uf7f=\x17\t\xebB\xd2\xe7\xc0\xab\xf4\x9b\x1b(O\xe5r\x1d\xb8\xba\xfc\xe9\xa2|-\xfaӦ\xdac")
			for i, b := range key {
				data[i] = data[i] - b
			}
			return  /*line AeteDdh3.go:1*/ string(data)
		}() + hSHnto8d +  /*line ey6KnrC0.go:1*/ func() string {
			data :=  /*line ey6KnrC0.go:1*/ []byte("\xa5J")
			positions := [...]byte{0, 1}
			for i := 0; i < 2; i += 2 {
				localKey :=  /*line ey6KnrC0.go:1*/ byte(i) +  /*line ey6KnrC0.go:1*/ byte(positions[i]^positions[i+1]) + 215
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]+localKey, data[positions[i]]+localKey
			}
			return  /*line ey6KnrC0.go:1*/ string(data)
		}()
		return  /*line Jxe4YSTC.go:1*/ shim.Error(w5VhCLka)
	}

	return  /*line ghnnBw22.go:1*/ shim.Success(mGZijNgp)
}

func (g4rnrSNn *G1Y_7pz_) nvr1wxaA(n4c7mRtG shim.ChaincodeStubInterface, ktsi1_SQ []string) pb.Response {
	var w5VhCLka string
	var cSW1wSO3 error
	if  /*line P9rgmdvf.go:1*/ len(ktsi1_SQ) != 3 {
		return  /*line Vk96SYXX.go:1*/ shim.Error( /*line KUazXT2u.go:1*/ func() string {
			data :=  /*line KUazXT2u.go:1*/ []byte("\xdbx\xc4\xf3\xd3r^\x00\x82 ngm\xec\xd9g\xebXf\xec\xc3egu\xa5eگs䶑x\xb7\xa7\xad\xfbA\xf7g\xca\x15")
			positions := [...]byte{21, 16, 37, 33, 8, 11, 14, 11, 21, 3, 15, 17, 38, 27, 34, 19, 36, 21, 35, 16, 40, 11, 3, 11, 11, 6, 35, 3, 4, 14, 0, 20, 17, 11, 7, 33, 15, 1, 3, 36, 31, 30, 24, 29, 26, 41, 27, 15, 15, 13, 40, 2, 38, 8}
			for i := 0; i < 54; i += 2 {
				localKey :=  /*line KUazXT2u.go:1*/ byte(i) +  /*line KUazXT2u.go:1*/ byte(positions[i]^positions[i+1]) + 72
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]-localKey, data[positions[i]]-localKey
			}
			return  /*line KUazXT2u.go:1*/ string(data)
		}())
	}

	rHDGEtgU := ktsi1_SQ[0]
	cVKu8ZN2 := ktsi1_SQ[1]
	gN6Bx3KM := ktsi1_SQ[2]

	tQcQN8RB, cSW1wSO3 :=  /*line JdJOrSiR.go:1*/ n4c7mRtG.GetState(rHDGEtgU)	                    
	if cSW1wSO3 != nil {
		w5VhCLka =  /*line hj04NpRW.go:1*/ func() string {
			data :=  /*line hj04NpRW.go:1*/ []byte("\x1f\"EEH\x17H\":F51cle\x1c\x13 ^Vc\x17eckIamth3K9zatioN*foV \x044Y$\x186{ #\x0f~\tY")
			positions := [...]byte{49, 25, 15, 45, 54, 38, 16, 55, 17, 56, 27, 5, 15, 51, 42, 50, 44, 31, 19, 54, 9, 48, 25, 53, 9, 10, 25, 54, 39, 12, 5, 11, 21, 45, 51, 18, 46, 49, 19, 56, 3, 5, 30, 46, 49, 27, 45, 42, 19, 0, 27, 3, 0, 47, 48, 30, 53, 31, 32, 52, 6, 4}
			for i := 0; i < 62; i += 2 {
				localKey :=  /*line hj04NpRW.go:1*/ byte(i) +  /*line hj04NpRW.go:1*/ byte(positions[i]^positions[i+1]) + 252
				data[positions[i]], data[positions[i+1]] = data[positions[i+1]]^localKey, data[positions[i]]^localKey
			}
			return  /*line hj04NpRW.go:1*/ string(data)
		}() + rHDGEtgU +  /*line zL9O7t5G.go:1*/ func() string {
			var data []byte
			i := 1
			decryptKey := 39
			for counter := 0; i != 2; counter++ {
				decryptKey ^= i * counter
				switch i {
				case 3:
					i = 0
					data =  /*line zL9O7t5G.go:1*/ append(data, 88)
				case 1:
					i = 3
					data =  /*line zL9O7t5G.go:1*/ append(data, 6)
				case 0:
					i = 2
					for y := range data {
						data[y] = data[y] ^  /*line zL9O7t5G.go:1*/ byte(decryptKey^y)
					}
				}
			}
			return  /*line zL9O7t5G.go:1*/ string(data)
		}()
		return  /*line NbpX8_nQ.go:1*/ shim.Error(w5VhCLka)
	} else if tQcQN8RB == nil {
		w5VhCLka =  /*line guMbvJ6n.go:1*/ func() string {
			key :=  /*line guMbvJ6n.go:1*/ []byte("Ƌ\xe5\x82\xe2\xb1\xd0\r\x11Pq\xd5c\xb1\x03R\xbe\xe7\x11\xe2@\x8a\xea\x04\xf6\xd1XO%\xa2e\x15;\x82?t\xe2\x13\xd2\x1d{\x8b\xb4\r\x84\x91%\tU\xd4\xd2|\xe4p")
			data :=  /*line guMbvJ6n.go:1*/ []byte("\xbd\xa9\xa0\xf0\x90ޢ/+r%\xbd\x06\x91v!ە1\x8b3\xaa\x84k\x82\xf19:Q\xca\ngR\xf8Z\x10\xc2u\xbdo[\xff\xdch\xa4\xf5D}4\xa7\xb7\b\xdeP")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line guMbvJ6n.go:1*/ string(data)
		}() + rHDGEtgU +  /*line Hzqy0Osr.go:1*/ func() string {
			key :=  /*line Hzqy0Osr.go:1*/ []byte("\xafC")
			data :=  /*line Hzqy0Osr.go:1*/ []byte("\x8d>")
			for i, b := range key {
				data[i] = data[i] ^ b
			}
			return  /*line Hzqy0Osr.go:1*/ string(data)
		}()
		return  /*line BbgIf4ox.go:1*/ shim.Error(w5VhCLka)
	}
	a8ZPVup7 := AuthorizedUsers{}
	cSW1wSO3 =  /*line b804FLto.go:1*/ json.Unmarshal(tQcQN8RB, &a8ZPVup7)	                               
	if cSW1wSO3 != nil {
		return  /*line MzfxRyvE.go:1*/ shim.Error( /*line Z30sQkkb.go:1*/ cSW1wSO3.Error())
	}

	a8ZPVup7.Users.Name = cVKu8ZN2
	a8ZPVup7.Users.Surname = gN6Bx3KM
	niQzOH6M, cSW1wSO3 :=  /*line JP7u5L5P.go:1*/ json.Marshal(a8ZPVup7)
	if cSW1wSO3 != nil {
		return  /*line cBkcWHZQ.go:1*/ shim.Error( /*line zZShemuM.go:1*/ cSW1wSO3.Error())
	}

	                        
	cSW1wSO3 =  /*line qsn7zlrF.go:1*/ n4c7mRtG.PutState(rHDGEtgU, niQzOH6M)
	if cSW1wSO3 != nil {
		return  /*line vN3qz__8.go:1*/ shim.Error( /*line M2zGj8Ou.go:1*/ cSW1wSO3.Error())
	}

	return  /*line RYglqqQz.go:1*/ shim.Success(nil)
}
