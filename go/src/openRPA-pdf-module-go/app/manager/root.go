package manager

import (
	"openRPA-basic-module/models"

	"reflect"

	"github.com/fatih/structs"
	"github.com/pkg/errors"
	"github.com/stoewer/go-strcase"
)

func RetrieveEndPoints(JSON *models.JSON) (endPoints []string) {

	for _, v := range JSON.Roots {
		endPoints = append(endPoints, v.EndPoint)
	}

	return endPoints
}

func EndPointIndex(JSON *models.JSON, endPoint string) (endPointIndex int) {

	for i, v := range JSON.Roots {
		if v.EndPoint == endPoint {
			endPointIndex = i
		}
	}

	return endPointIndex
}

func RetrieveArgumentFromContents(JSON *models.JSON, endPoint string, argument string) (arguments []string, err error) {

	for _, v := range JSON.Roots {
		if v.EndPoint == endPoint {
			contents := v.Arguments[argument].Contents
			return contents, nil
		}
	}

	return nil, errors.New("指定したArgumentが見つかりません")
}

func RetrieveDataFromPointer(JSON *models.JSON, endPoint string) error {

	// 指定されたargumentsのpointerを参照して、contentに値をぶち込む

	var index int
	var root models.Root

	for i, v := range JSON.Roots {
		if v.EndPoint == endPoint {
			index = i
			root = v
		}
	}

	if index == 0 {
		return errors.New("指定されたendPointは存在しません")
	}

	// JSONの中の、Argument配列を読み込む
	for i, arg := range root.Arguments {
		// argumentsのpointerを参照し、配列だったら、contentsに、単体だったら、contentに格納する

		pointer := parsePointer(arg.Pointer)

		// JSON.(Input or Exports).(Exportsだったら、for回して、エンドポイントが当てはまるところの、argumentsを取り出す)
		// 1. Input or Exports

		if pointer.InputOrExports == "input" {
			// 1.1. Inputだったら、指定されたargumentsのobjectを取り出す
			in := structs.Map(JSON.Input)

			targetArgumentsValues, ok := retrieveTargetArgumentValues(in, pointer.Keys)
			if !ok {
				return errors.New("pointerの取り出しに失敗しました")
			}

			if len(targetArgumentsValues) < 2 {
				arg.Content = targetArgumentsValues[0]
			} else {
				arg.Contents = targetArgumentsValues
			}

		} else if pointer.InputOrExports == "exports" {
			// 1.2. Exportsだったら、for回して、エンドポイントが当てはまるところの、argumentsを取り出す)

			export, ok := retrieveTargetExport(JSON, pointer)
			if !ok {
				return errors.New("targetのend pointが存在しません")
			}

			e := structs.Map(export)

			targetArgumentsValues, ok := retrieveTargetArgumentValues(e, pointer.Keys)
			if !ok {
				return errors.New("pointerの取り出しに失敗しました")
			}

			if len(targetArgumentsValues) < 2 {
				arg.Content = targetArgumentsValues[0]
			} else {
				arg.Contents = targetArgumentsValues
			}

		}

		root.Arguments[i] = arg
	}

	return nil
}

func parsePointer(oldPointer []string) (pointer models.Pointer) {
	isArray := false
	for i, p := range oldPointer {
		if i == 0 {
			pointer.InputOrExports = p
		} else if i == 1 && pointer.InputOrExports == "exports" {
			pointer.EndPoint = p
		} else if (i > 0 && pointer.InputOrExports == "input") || (i > 1 && pointer.InputOrExports == "exports") {
			if p == "map(n)" {
				isArray = true
			} else {
				key := models.Key{
					Name:    p,
					IsArray: isArray,
				}
				pointer.Keys = append(pointer.Keys, key)
				isArray = false
			}
		}
	}
	return
}

func retrieveTargetExport(JSON *models.JSON, pointer models.Pointer) (export models.Export, ok bool) {
	for _, v := range JSON.Exports {
		if v.FromEndPoint == pointer.EndPoint {
			return v, true
		}
	}

	return models.Export{}, false
}

func retrieveTargetArgumentValues(args interface{}, keys []models.Key) (result []string, ok bool) {
	iv := reflect.ValueOf(args)

	out := map[string]interface{}{}
	ov := reflect.ValueOf(&out).Elem()

	// for文で、[]Keyを回して、models.Key.Nameを取り出し、そのnameのvalueを抜き取り、argsに代入して行く
	for i, v := range keys {

		// 1.IsArray = trueのとき。  args.key.nameを取得。key: name, value: args.key.name
		// 2.IsArray = falseのとき。 args.key.nameを取得。配列を作って、そこに、key: name, value: args.key.name

		// 1. nameに対応するvalueを抜き出してみる
		//for _, k := range rv.MapKeys() {
		//	spew.Dump(k)
		//}

		if v.IsArray {

			var o []interface{}  // [{"pass": "12345"},{"pass": "12345"},{"pass": "12345"}]が入る
			var o2 []interface{} // ["12345", "22222", "33333"]が入る

			for _, v2 := range out { // このoutは、map[string]interface {}で、"Data": interface{}
				value2 := v2.([]interface{})
				o = value2 //このoには、　[{"pass": "12345"},{"pass": "12345"},{"pass": "12345"}]が入る
			}

			for _, v3 := range o { // このoは、[]interface{}で、　[{"pass": "12345"},{"pass": "12345"},{"pass": "12345"}]
				value3 := v3.(map[string]interface{})
				for k, v4 := range value3 {
					if k == v.Name {
						o2 = append(o2, v4) // このo2には、 ["12345", "22222", "33333"]が入る
					}
				}
			}

			out = map[string]interface{}{}
			ov := reflect.ValueOf(&out).Elem()
			ov.SetMapIndex(reflect.ValueOf(v.Name), reflect.ValueOf(o2)) // keyをpassにして、valueを["12345","22222"]にする

		} else {

			if v.Name == "data" {
				v.Name = strcase.UpperCamelCase(v.Name)
			}

			if v.Name == "files" {
				v.Name = strcase.UpperCamelCase(v.Name)
			}

			out = map[string]interface{}{}
			ov := reflect.ValueOf(&out).Elem()
			ov.SetMapIndex(reflect.ValueOf(v.Name), iv.MapIndex(reflect.ValueOf(v.Name)))
		}
		iv = ov

		if i == len(keys)-1 && len(keys) > 0 {
			for _, v := range out {
				value3 := v.([]interface{})

				for _, v := range value3 {
					str, _ := v.(string)
					result = append(result, str)
				}

			}
			return result, true
		}
	}
	return nil, false
}
