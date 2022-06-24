// Code generated by gocode.Generate; DO NOT EDIT.

package pkg1

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
)

var cuegenvalMyStruct = cuegenMake("MyStruct", &MyStruct{})

// Validate validates x.
func (x *MyStruct) Validate() error {
	return cuegenCodec.Validate(cuegenvalMyStruct, x)
}

// Complete completes x.
func (x *MyStruct) Complete() error {
	return cuegenCodec.Complete(cuegenvalMyStruct, x)
}

var cuegenvalOtherStruct = cuegenMake("OtherStruct", &OtherStruct{})

// Validate validates x.
func (x *OtherStruct) Validate() error {
	return cuegenCodec.Validate(cuegenvalOtherStruct, x)
}

var cuegenvalString = cuegenMake("String", nil)

// ValidateCUE validates x.
func (x String) ValidateCUE() error {
	return cuegenCodec.Validate(cuegenvalString, x)
}

var cuegenvalSpecialString = cuegenMake("SpecialString", nil)

// ValidateSpecialString validates x.
func ValidateSpecialString(x string) error {
	return cuegenCodec.Validate(cuegenvalSpecialString, x)
}

var cuegenvalPtr = cuegenMake("Ptr", Ptr(nil))

// ValidatePtr validates x.
func ValidatePtr(x Ptr) error {
	return cuegenCodec.Validate(cuegenvalPtr, x)
}

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	f, err := cuegenInstance.LookupField(name)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 533 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\x94R]k\x13M\x14\x9e\xb3\xcd\v\xef\x1e\xaa\xe0\x0f\x10\u01b9\xcaJ\xbb\xf9\x00\x11\x96\xaeZk\x85\\\xa4\tFE\x10/\xc6\xc9d3d3\xb3\xecN\xc4P\x1bPk\xed\x9f\xf4\u059f\u0455\xfdj\xa3w\u075b9\x9c\xdd\xe7c\xe7y\xee\xe4?\x1dp\xf2K\x02\xf97B\x1e\xe7_w\x00v\x95\xce,\xd7B\xbe\xe0\x96\x17{\u0601\xd6+c,8\x04Zcn\xe7\xb0K\u0fd7*\x96\x19\xe4\x17\x84\x90\xfb\xf9\x0f\a\xe0\xee\xfb\x0fb%\xfd\x99\x8ak\xe4\x05\x81\xfc\x9c\x90v\xfe}\a\xe0\xff\x9b\xfd9\x01\aZ'|)\v\xa2V\xb9DB\u0215\xf3+\xbf$\x0e\x00\uc255\x8c\xb9\x8e|\x93F\x9d\xc8t\xa4\x16f\xaat1\v3\x95\x1d+3;\xe5\x96w\x92E\xd4\x03\x80{\xc5\xd9i|\xfbb%\xe1\n~'\\,x$i\xf1\x12Q-\x13\x93Z\xdaF\x97\u0742\xbd\xcf\xd0eKn\xe7\u0159\xd9T\xe9(c\xe8!\x0e\xd7\x13\x9b\xae\x84\r\xe8)\xba\x87\x01\xa5\aa\xaf\x8b\xee\xf3\x80\xd2p\xc3\x04\xb7\x8c~\xa1\x0f\xd9\xd4D\f\xdd\xd1\u04c0\x8e\xec\\\xa6\x15\x06\xddA@\v[}\u007fP\xba\x1aJ<\xa3\xcf\"\xd3\xde\x13f\x99\xc4\xd2\xca\xf0\xa8\x1e<\xdc\x026b\xb5\x11\xff\xc8h\u02d5\xce\x0e\xf5\xba\xcd\xde1\x0f\xddqP\xf1\x8e\x95X\x14\xac8)?\rh\xfd<\b\x19k\xe6R\xf0\x13\x8f\u0554[\x19\xbe\xad\x87\xa37\xc7\x1eN\x12)\x14\x8f\x1bp\xb8aY\xb5a\x15\u02ae\x13\x19V.<\x1cD\u06a4\xf2\xf5\\e\xa5L\xb8a3c\x18\x8e\x96\xca^\xebR\xaa\xb4-\xb1\xfb\x1e\x9e\x18}\xfcYe\xb6\xe4>-/\xad\xe2\xaaoa\xdf\u00d91A\x01\xc1\xb1M\x9b\xbf.b\xf0\x87\xab\u062a$\x96\xa3Y\xbb\xd7\xf5\xf0\f\tqnS\x97~]\x97\xfe\xdfu\xe1[e\xe9_\x97\xe5&ql\x82j\xcc\x1c\xf4\xba\xdd-\xe7\xffd\xc1?\nV\x98\xabb\b\xe8\x93GH\u021f\x00\x00\x00\xff\xffv\xf2MLm\x03\x00\x00")
