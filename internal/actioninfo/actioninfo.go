package actioninfo

import "log"

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	if len(dataset) == 0 {
		return
	}

	for _, v := range dataset {
		err := dp.Parse(v)

		if err != nil {
			log.Print(err)
			continue
		}
	}
	str, err := dp.ActionInfo()

	if err != nil {
		log.Print(err)

		return
	}

	log.Print(str)
}
