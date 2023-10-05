package task15

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	return strings.Repeat("s", size)

}

func someFunc() string {
	v := createHugeString(1 << 10)
	// делаем justString локальной переменной, т.к локальные переменные, в отличии от глобальных, помещаются в стэк, а не в кучу
	// Т.к работа со стэком быстрее, то и скорость программы быстрее
	var justString string
	// размер v может быть меньше 100 символов
	if len(v) < 100 {
		justString = v
	} else {
		justString = v[:100]
	}
	fmt.Println(v)
	fmt.Println(justString)
	return justString

}
func Task15() {
	someFunc()

}

// при justString = v или justString = v[:100], происходит копирование данных, а не создание новыХ, тоесть переменная ссылается на память выделенную под v.
// v занимает много памяти и не используется после выполнение somefunc(), однако сборщик мусора не может удалить переменную v потому что juststring ссылается на неё
//todo разобраьбся
