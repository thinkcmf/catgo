//利用 reflect 技术把结构体的可 export 值复制到 dst 中，dst 必须是相似结构体的指针。
//copy the exported value of a struct to dst , with reflect.
package catgo

import (
	"errors"
	"reflect"
)

//StructCopy copy the exported value of a struct to a likely struct , with reflect.
func StructCopy(src, dst interface{}) error {
	srcV, err := srcFilter(src)
	if err != nil {
		return err
	}
	dstV, err := dstFilter(dst)
	if err != nil {
		return err
	}
	srcKeys := make(map[string]string)
	for i := 0; i < srcV.NumField(); i++ {
		srcKeys[srcV.Type().Field(i).Name] = srcV.Type().Field(i).Type.Name()
	}
	for i := 0; i < dstV.Elem().NumField(); i++ {
		fName := dstV.Elem().Type().Field(i).Name
		fTypeName := dstV.Elem().Type().Field(i).Type.Name()
		if srcFTypeName, ok := srcKeys[fName]; ok && fTypeName == srcFTypeName {
			v := srcV.FieldByName(dstV.Elem().Type().Field(i).Name)
			if v.CanInterface() {
				dstV.Elem().Field(i).Set(v)
			}
		}
	}

	return nil
}

func StructFill(src, dst interface{}) error {
	srcV, err := srcFilter(src)
	if err != nil {
		return err
	}
	dstV, err := dstFilter(dst)
	if err != nil {
		return err
	}
	srcKeys := make(map[string]string)
	for i := 0; i < srcV.NumField(); i++ {
		srcKeys[srcV.Type().Field(i).Name] = srcV.Type().Field(i).Type.Name()
	}

	for i := 0; i < dstV.Elem().NumField(); i++ {
		fName := dstV.Elem().Type().Field(i).Name
		fTypeName := dstV.Elem().Type().Field(i).Type.Name()
		if srcFTypeName, ok := srcKeys[fName]; ok && fTypeName == srcFTypeName {
			v := srcV.FieldByName(dstV.Elem().Type().Field(i).Name)
			//v2 := dstV.Elem().Field(i)
			if v.CanInterface() && !v.IsZero() {
				dstV.Elem().Field(i).Set(v)
			}
		}
	}

	return nil
}

func srcFilter(src interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), errors.New("src type error: not a struct or a pointer to struct")
	}
	return v, nil
}

func dstFilter(src interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Type().Kind() != reflect.Ptr {
		return reflect.Zero(v.Type()), errors.New("src type error: not a pointer to struct")
	}
	if v.Elem().Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), errors.New("src type error: not point to struct")
	}
	return v, nil
}
