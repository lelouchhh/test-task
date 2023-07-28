## Конкретизация ТЗ

1. Пулл задач, на которые можно получить решение, заранее известен.
2. Решение задачи является исходным код (альтернативой мог бы быть вывод программы по представляемым параметрам. Однако это имеет мало смысла, ведь это не проверяет знание алгоритомов)
3. Оплатить можно только наличнымы. (После чего админ руками меняет уровень долга пользователя)
4. Пользователь может получить список задач, которые предоставляет сервис

## Доработка тз алгоритма
```
Дан неотсортированный массив из N чисел от 1 до N,
при этом несколько чисел из диапазона [1, N] пропущено,
а некоторые присутствуют дважды.

Найти все пропущенные числа.

Дополнение:
1. 1 < N < 2**32
2. output: []int
```
## Запуск
```
 1. cd shop && docker-compose up --build -d 
 2. cd stock-service && docker-compose up --build -d 
```
## EndPoints
``` 
[GIN-debug] GET    /algorithm

[GIN-debug] POST   /user
	FirstName  string  `json:"first_name" validate:"required,gt=2"`
	SecondName string  `json:"second_name" validate:"required,gt=2"`
	LastName   string  `json:"last_name" validate:"required,gt=2"`
	Email      string  `json:"email" validate:"required,email"`
[GIN-debug] PUT    /increase_debt   
	Id   string  `json:"id" validate:"required"`
	Debt float32 `json:"debt" validate:"required"`     
[GIN-debug] PUT    /decrease_debt        
	Id   string  `json:"id" validate:"required"`
	Debt float32 `json:"debt" validate:"required"`
[GIN-debug] POST   /purchase       
	UserId string `json:"user_id"  validate:"required"`
	AlgoId string `json:"algo_id"  validate:"required"`      
```

## Что можно было бы добавить

1. Микросервис для емейл рассылки
2. Добавить JWT
3. Добавить больше функционала в виде новых ендпоинтов и таблиц в базе