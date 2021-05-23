package missingo

import (
	"fmt"
)

type defaultContainer struct {
	defaultValue interface{}
}

func Default(defaultValue interface{}) *defaultContainer {
	e := &defaultContainer{defaultValue}
	fmt.Println("Default:", e)
	return e
}

// 1. try
// cannot use values (type []string) as type []interface {} in argument to missingo.Default("default value").IfNotSet
// ----
// func (e *defaultContainer) IfNotSet(param []interface{}) interface{} {
// 	fmt.Printf("Param: %q, %T\n", param, param)
// 	return 1
// }
// ----

// 2. try
// invalid argument param (type interface {}) for len
// ----
// func (e *defaultContainer) IfNotSet(param interface{}) interface{} {
// 	fmt.Printf("Param: %q, %T\n", param, param)
// 	len(param)
// 	return 1
// }
// ----

// 3. try

// 	vof := reflect.ValueOf(param)
// 	fmt.Printf("Reflect value: %q, %T\n", vof, vof)
// 	fmt.Println("Length:", vof.Len())

// 	if vof.Len() == 0 {
// 		return e.defaultValue
// 	} else {
// 		return vof.Index(1)
// 	}

// 	// fmt.Printf("Param: %q, %T, %d\n", param, param, len(param))

// 	// val := param[0]
// 	// fmt.Printf("Got value: %q, %T\n", val, val)

// 	// res := func(a []interface{}) interface{} {
// 	// 	return a[0]
// 	// }(param)
// 	// fmt.Printf("Got value: %q, %T\n", res, res)

// 	// return sliceVal[0]
// 	// We need to copy the values in the interface slice: []interface{}
// 	// realVal := make([]interface{}, 1)
// 	// for i, val := range sliceVal {
// 	// 	realVal[i] = val
// 	// }
// 	// realVal[0] = sliceVal[0]
// 	// return realVal[0]

// 	// val := [1]interface{}{param[0]}
// 	// v2 := val[0]
// 	// fmt.Printf("First value: %v, %q, %T\n", v2, v2, v2)
// 	// // return v2

// 	// switch len(val) {
// 	// case 0:
// 	// 	return e.defaultValue
// 	// case 1:
// 	// 	return val[0]
// 	// default:
// 	// 	panic("Expected exactly 0 (for default) or 1 parameters!")
// 	// }

// 	// switch val := param[0].(type) {
// 	// default:
// 	// 	for _, v := range val {
// 	// 		return v
// 	// 	}
// 	// }

// 	// switch val := param.(type) {
// 	// case interface{}:
// 	// 	return e.defaultValue
// 	// default:
// 	// 	if len(val) == 1 {
// 	// 		return val[0]
// 	// 	} else {
// 	// 		panic("Expected exactly 0 (for default) or 1 parameters!")
// 	// 	}
// 	// }

// 	// if _, empty := param[0].(interface{}); empty {
// 	// 	return e.defaultValue
// 	// } else if len(param[0]) == 1 {
// 	// 	return param[0]
// 	// } else {
// 	// 	panic("Expected exactly 0 (for default) or 1 parameters!")
// 	// }
// 	return 5
// }
