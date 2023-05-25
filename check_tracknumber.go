package main

import "fmt"

func main_info(sign string) string {
	if sign[0] == 'R' {
		return "Регистрируемое отправление (до 2 кг)"
	} else if sign[0] == 'L' {
		return "Отслеживоемое письмо"
	} else if sign[0] == 'V' {
		return "Письмо с объявленной ценностью"
	} else if sign[0] == 'C' {
		return "Международная посылка (более 2 кг)"
	} else if sign[0] == 'E' {
		return "Экспресс отправление (EMS)"
	} else if sign[0] == 'U' {
		return "Нерегистрируемое и неотслеживаемое отправление"
	} else if sign[0] == 'Z' {
		return "Простой регистрируемый пакет (SRM)"
	} else {
		return sign
	}
}

func is_num(s string) bool {
	var result bool = true
	
	for _, c := range s {
		if c < '0' || c > '9' {
			result = false
			break
		}
	}
	
	return result
}

func main() {
	var tracknum string
	
	fmt.Print("Введите трек-номер: ")
	fmt.Scanln(&tracknum)
	
	if len(tracknum) == 13 && is_num(tracknum) == false {
		sign := tracknum[:2]
		iso := tracknum[11:]
		
		fmt.Println(main_info(sign))
		fmt.Println("Место отправления ", iso)
	} else if len(tracknum) == 13 {
		fmt.Println("Внутренний трек-номер")
	} else if len(tracknum) == 14 {
		fmt.Println("Индекс отправителя:", string(tracknum[:6]))
	} else {
		fmt.Println("Неизвестный тип трек-номера")
	}
}
