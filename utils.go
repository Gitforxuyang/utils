package utils

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//示例
//	type Animal struct {
//		Age int `default:"1"`  简单的设置默认值
//		Time int `default:",NOW_SECOND()"`   设置函数默认值， NOW_SECOND()=秒 NOW_MILLI()=毫秒
//	}
//为实现了IEntity的struct设置默认值
func Default(entity interface{}) error {
	if reflect.ValueOf(entity).Kind() != reflect.Ptr {
		return errors.New("传入参数必须是指针类型")
	}
	return setDefault(reflect.ValueOf(entity).Elem(), reflect.TypeOf(entity).Elem())
}

//设置默认值函数
func setDefault(_value reflect.Value, _type reflect.Type) error {
	//不能设置基础类型的指针默认值
	if _type.Kind() != reflect.Struct {
		return errors.New("设置默认值的类型必须是struct")
	}
	for i := 0; i < _value.NumField(); i++ {
		if _value.Field(i).Kind() == reflect.Struct {
			setDefault(_value.Field(i), _value.Field(i).Type())
			continue
		}
		if _value.Field(i).Kind() == reflect.Ptr {
			if _value.Field(i).IsNil() {
				return errors.New("指针类型的属性不能为nil")
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
			//如果设置了预制函数则设置 NOW_SECOND()=当前时间秒数 NOW_MILLI()=当前时间毫秒数
			switch _defaultArr[1] {
			case "NOW_SECOND()":
				_value.Field(i).SetInt(time.Now().Unix())
				break
			case "NOW_MILLI()":
				_value.Field(i).SetInt(time.Now().UnixNano() / 1e6)
				break
			default:
				return errors.New("未知的内置默认函数")
				break
			}
			continue
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
	return fmt.Sprintf("%024d", sortId)
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

//只能支持基础类型数组或切片使用 不支持基础类型的指针
func Includes(arr interface{}, value interface{}) (bool, error) {
	arrKind := reflect.TypeOf(arr).Kind()
	valueKind := reflect.TypeOf(value).Kind()
	//如果是传递过来的是指针 则拿到指针的具体类型
	if arrKind == reflect.Ptr || valueKind == reflect.Ptr {
		return false, errors.New("不能传入指针类型")
	}
	//_value := reflect.ValueOf(value)
	_arr := reflect.ValueOf(arr)
	switch arrKind {
	case reflect.Array:
	case reflect.Slice:
		for i := 0; i < _arr.Len(); i++ {
			if reflect.DeepEqual(value, _arr.Index(i).Interface()) {
				return true, nil
			}
		}
		break
	default:
		return false, nil
	}
	return false, nil
}

func GetUUIDStr() string {
	uuid := uuid.NewV4()
	return uuid.String()
}

//获取没有 破折号的 UUID
func GetNoDashUUIDStr() string {
	uuid := uuid.NewV4()
	str := strings.ReplaceAll(uuid.String(), "-", "")
	return str
}

func GetFixed(value interface{}) (float64, error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.ParseFloat(fmt.Sprintf("%d", value), 64)
	case reflect.Float32, reflect.Float64:
		return strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	}
	return 0, errors.New("不支持的类型")
}
