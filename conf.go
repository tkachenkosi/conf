package conf

import (
	"bufio"
	"fmt"
	"os"
	// "slices"
	"reflect"
	"strconv"
	"strings"
)

type conf struct {
	secName  string
	fileName string
	key      string
	value    string
	result   map[string]string
	// err      error
}

func NewConf(fileName string) (*conf, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil, fmt.Errorf("файл конфигурации не найден: %s", fileName)
	}
	return &conf{
		fileName: fileName,
		secName:  "main",
		result:   make(map[string]string),
	}, nil
}

// exempl: Read("[main]", &s)
func (c *conf) Read(secName string, s any) error {
	c.secName = secName

	if err := c.parser(); err == nil {
		// fmt.Println("read", c.secName)
		// Получаем значение структуры через рефлексию
		val := reflect.ValueOf(s).Elem()
		typ := val.Type()

		// Итерируемся по полям структуры
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)                           // Значение поля
			fieldName := strings.ToLower(typ.Field(i).Name) // Имя поля

			// Проверяем, есть ли ключ в map с таким же именем
			if mapValue, ok := c.result[fieldName]; ok {
				// Преобразуем значение из map в тип поля структуры
				switch field.Kind() {
				case reflect.String:
					field.SetString(mapValue)
				case reflect.Int:
					intValue, err := strconv.Atoi(mapValue)
					if err != nil {
						return fmt.Errorf("ошибка преобразования значения для поля %s: %v", fieldName, err)
					}
					field.SetInt(int64(intValue))
				case reflect.Bool:
					boolValue, err := strconv.ParseBool(mapValue)
					if err != nil {
						return fmt.Errorf("ошибка преобразования значения для поля %s: %v", fieldName, err)
					}
					field.SetBool(boolValue)
				case reflect.Float64:
					floatValue, err := strconv.ParseFloat(mapValue, 64)
					if err != nil {
						return fmt.Errorf("ошибка преобразования значения для поля %s: %v", fieldName, err)
					}
					field.SetFloat(floatValue)
				default:
					return fmt.Errorf("неподдерживаемый тип поля: %s", field.Kind())
				}
			}
		}
	} else {
		return err
	}

	return nil
}

func (c *conf) parser() error {
	var ok bool

	f, err := os.Open(c.fileName)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer f.Close()

	// каждый раз чистим map
	c.result = make(map[string]string, 2)

	buf := bufio.NewScanner(f)

	for buf.Scan() {
		line := strings.TrimSpace(buf.Text())

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		if !ok && strings.Contains(line, c.secName) {
			ok = true
			continue
		}

		if ok && strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			break
		}

		if ok && strings.Contains(line, "=") {
			fl := strings.SplitN(line, "=", 2)
			c.key = strings.TrimSpace(fl[0])
			c.value = strings.TrimSpace(fl[1])

			if len(c.key) > 0 && len(c.value) > 0 {
				c.result[c.key] = c.value
			}
		}
	}

	if err := buf.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return nil
}
