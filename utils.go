package utils

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//示例
//	type Animal struct {
//		Age int `default:"1"`  简单的设置默认值
//		Time int `default:",NOW_SECOND()"`   设置函数默认值， NOW_SECOND()=秒 NOW_NANO()=毫秒
//	}
//为实现了IEntity的struct设置默认值
func Default(entity interface{}) error {
	if reflect.ValueOf(entity).Kind() != reflect.Ptr {
		errors.New("传入参数必须是指针类型")
	}
	return setDefault(reflect.ValueOf(entity).Elem(), reflect.TypeOf(entity).Elem())
}

//设置默认值函数
func setDefault(_value reflect.Value, _type reflect.Type) error {
	//不能设置基础类型的指针默认值
	if _type.Kind() != reflect.Struct {
		errors.New("设置默认值的类型必须是struct")
	}
	for i := 0; i < _value.NumField(); i++ {
		if _value.Field(i).Kind() == reflect.Struct {
			setDefault(_value.Field(i), _value.Field(i).Type())
			continue
		}
		if _value.Field(i).Kind() == reflect.Ptr {
			if _value.Field(i).IsNil() {
				errors.New("指针类型的属性不能为nil")
			}
			setDefault(_value.Field(i).Elem(), _value.Field(i).Type().Elem())
			continue
		}
		_default := _type.Field(i).Tag.Get("default")
		if _default == "" {
			continue
		}
		_defaultArr := strings.Split(_default, ",")
		if len(_defaultArr) == 1 {
			_default = _defaultArr[0]
		} else {
			//如果设置了预制函数则设置 NOW_SECOND()=当前时间秒数 NOW_NANO()=当前时间毫秒数
			switch _defaultArr[1] {
			case "NOW_SECOND()":
				_value.Field(i).SetInt(time.Now().Unix())
				break
			case "NOW_NANO()":
				_value.Field(i).SetInt(time.Now().UnixNano())
				break
			}
		}

		switch _value.Field(i).Kind() {
		case reflect.Bool:
			d, err := strconv.ParseBool(_default)
			if err != nil {
				return err
			}
			_value.Field(i).SetBool(d)
			break
		case reflect.Int:
			d, err := strconv.ParseInt(_default, 10, 32)
			if err != nil {
				return err
			}
			_value.Field(i).SetInt(d)
			break
		case reflect.Int8:
			d, err := strconv.ParseInt(_default, 10, 8)
			if err != nil {
				return err
			}
			_value.Field(i).SetInt(d)
			break
		case reflect.Int16:
			d, err := strconv.ParseInt(_default, 10, 16)
			if err != nil {
				return err
			}
			_value.Field(i).SetInt(d)
			break
		case reflect.Int32:
			d, err := strconv.ParseInt(_default, 10, 32)
			if err != nil {
				return err
			}
			_value.Field(i).SetInt(d)
			break
		case reflect.Int64:
			d, err := strconv.ParseInt(_default, 10, 64)
			if err != nil {
				return err
			}
			_value.Field(i).SetInt(d)
			break
		case reflect.Uint:
			d, err := strconv.ParseUint(_default, 10, 32)
			if err != nil {
				return err
			}
			_value.Field(i).SetUint(d)
			break
		case reflect.Uint8:
			d, err := strconv.ParseUint(_default, 10, 8)
			if err != nil {
				return err
			}
			_value.Field(i).SetUint(d)
			break
		case reflect.Uint16:
			d, err := strconv.ParseUint(_default, 10, 16)
			if err != nil {
				return err
			}
			_value.Field(i).SetUint(d)
			break
		case reflect.Uint32:
			d, err := strconv.ParseUint(_default, 10, 32)
			if err != nil {
				return err
			}
			_value.Field(i).SetUint(d)
			break
		case reflect.Uint64:
			d, err := strconv.ParseUint(_default, 10, 64)
			if err != nil {
				return err
			}
			_value.Field(i).SetUint(d)
			break
		case reflect.Float32:
			d, err := strconv.ParseFloat(_default, 10)
			if err != nil {
				return err
			}
			_value.Field(i).SetFloat(d)
			break
		case reflect.Float64:
			d, err := strconv.ParseFloat(_default, 10)
			if err != nil {
				return err
			}
			_value.Field(i).SetFloat(d)
			break
		case reflect.String:
			_value.Field(i).SetString(_default)
			break
		}
	}
	return nil
}

//将长id转化为短id 类似于000000000000000062361472->62361472
func LongIdToSortId(longId string) (int64, error) {
	str := strings.TrimLeft(longId, "0")
	return strconv.ParseInt(str, 10, 64)
}

func SortIdToLongId(sortId int64) string {
	return fmt.Sprintf("%024s", sortId)
}

//字符串转化为objectId 如果字符串长度不够则自动补足24位
func StringToObjectId(strId string) (*primitive.ObjectID, error) {
	longId := fmt.Sprintf("%024s", strId)
	objId, err := primitive.ObjectIDFromHex(longId)
	if err != nil {
		return nil, err
	}
	return &objId, nil
}
