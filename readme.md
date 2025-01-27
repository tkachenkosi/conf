Программа предназначена для работы с конфигурационными файлами в формате, похожем на INI, который содержит секции и ключи. Она считывает информацию из определенной секции, разделяя её на ключи и значения, и сохраняет эти значения в соответствующие поля структуры, которые соответствуют названиям ключей в файле конфигурации.

Пример испоьзования:

[postgres]
host=192.168.1.2
port=5432
db=books
conns=5
user=user1


    c := struct {
    Host  string
 	Db    string
 	Conns int
    }{}

    ini, err := conf.NewConf("app.ini")
    if err != nil {
	fmt.Println("Ошибка:", err)
	return
    }

    if err := ini.Read("[postgres]", &c); err == nil {
	fmt.Println("Заполненная структура секции postgres:")
        fmt.Println("Host:  ", c.Host)
        fmt.Println("Db:    ", c.Db)
        fmt.Println("Conns: ", c.Conns)
    }
