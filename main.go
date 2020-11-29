package main

import (
"errors"
"fmt"
xerrors "github.com/pkg/errors"
)

var NoSqlRow = errors.New("no sql row")

func dao()(err error){
	//查数据库sql
	//if err != nil{
	return xerrors.Wrap(NoSqlRow,"查数据库sql报错")
	//}
	//return
}

func biz()(err error){
	err = dao()
	return
}

func main() {
	err := biz()
	fmt.Printf("error : %+v \n",err)

}
