package errortool

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	s:=`codehuozheMSG`
	err:=fmt.Errorf(s)
	if HasErr(err) {
		fmt.Println(GetErrInfo(err))
	}
	fmt.Println(GetErrMap(t1("MCM2000","MSGXX")))
}

func t1(code,msg string) error {
	return NewError(code,msg)
}
