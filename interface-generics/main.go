package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Serializable interface {
	Encode() string
	Decode(data string)
}

type StoreItem struct {
	StoredAt       time.Time
	DataSerialized string
	DataType       string
}

// ItemOne -----------------------------
type ItemOne struct {
	Name    string
	Surname string
}

func (this *ItemOne) Encode() string {
	return fmt.Sprintf("%s--%s", this.Name, this.Surname)
}

func (this *ItemOne) Decode(data string) {
	this.Name = strings.Split(data, "--")[0]
	this.Surname = strings.Split(data, "--")[1]
}

func NewItemOneStoreItem(name, surname string) StoreItem {
	item := ItemOne{Name: name, Surname: surname}
	return StoreItem{
		StoredAt:       time.Now(),
		DataSerialized: item.Encode(),
		DataType:       "i1",
	}
}

// ItemTwo -----------------------------
type ItemTwo struct {
	Code string
	Type string
}

func (this *ItemTwo) Encode() string {
	return fmt.Sprintf("%s+++%s", this.Code, this.Type)
}

func (this *ItemTwo) Decode(data string) {
	this.Code = strings.Split(data, "+++")[0]
	this.Type = strings.Split(data, "+++")[1]
}

func NewItemTwoStoreItem(code, typ string) StoreItem {
	item := ItemTwo{Type: typ, Code: code}
	return StoreItem{
		StoredAt:       time.Now(),
		DataSerialized: item.Encode(),
		DataType:       "i2",
	}
}

func main() {

	// s := ItemOne{"ciao", "mona"}
	// itemValue := reflect.ValueOf(s)
	// itemType := reflect.TypeOf(s)
	// fmt.Println(itemValue)
	// fmt.Printf("%v\n", s)
	// fmt.Println(itemType)
	// sp := &s
	// itemValue = reflect.ValueOf(sp)
	// itemType = reflect.TypeOf(s)
	// fmt.Printf("%v\n", itemValue.Addr())
	// fmt.Printf("%p\n", sp)
	// fmt.Println(itemType)

	storeItems := []StoreItem{}

	storeItems = append(storeItems, NewItemOneStoreItem("paul", "newman"))
	storeItems = append(storeItems, NewItemTwoStoreItem("ABC", "Prod"))

	item := storeItems[0]
	itemone, _ := DecodeItem[ItemOne](item)
	fmt.Println(itemone)

	item = storeItems[1]
	itemtwo, _ := DecodeItem[ItemTwo](item)
	fmt.Println(itemtwo)
}

func DecodeItem[T any](item StoreItem) (T, bool) {
	var target T
	if item.DataType == "i1" {
		var i1 Serializable = &ItemOne{}
		i1.Decode(item.DataSerialized)
		if result, ok := i1.(T); ok {
			fmt.Println(result)
			return result, true
		}
		return target, false
	}
	if item.DataType == "i2" {
		var i2 Serializable = &ItemTwo{}
		i2.Decode(item.DataSerialized)
		if result, ok := i2.(T); ok {
			return result, true
		}
		return target, false
	}

	return target, false
}

func xDecodeItem(item StoreItem) Serializable {

	if item.DataType == "i1" {
		var i1 ItemOne
		i1.Decode(item.DataSerialized)
		return &i1
	}
	if item.DataType == "i2" {
		var i2 ItemTwo
		i2.Decode(item.DataSerialized)
		return &i2
	}

	return nil
}

func castAs[T any](item Serializable) (T, bool) {
	var target T

	if item == nil {
		return target, false
	}

	result, ok := item.(T)
	if ok {
		return result, ok
	}

	// pointerTest := reflect.ValueOf(item)
	// if pointerTest.Kind() == reflect.Ptr {
	// 	value := &item
	// 	result, ok := value.(T)
	// 	if ok {
	// 		return result, ok
	// 	}
	// }
	return target, false
}

func xcastAs[T any](item Serializable) (T, error) {
	var target T

	if item == nil {
		return target, errors.New("item must be a value and not nil")
	}

	itemValue := reflect.ValueOf(item)
	itemType := reflect.TypeOf(item)
	castType := reflect.TypeOf(target)

	if itemType.Kind() == reflect.Ptr {
		itemValue = itemValue.Elem()
		itemType = itemValue.Type()
	}

	if itemType == nil {
		return target, errors.New("item unknown")
	}

	if !itemType.ConvertibleTo(castType) {
		return target, errors.New("no casting is possible")
	}

	obj := itemValue.Convert(castType).Interface().(T)
	return obj, nil
}
