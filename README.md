# avito-internship-test

Задание находится [здесь](https://github.com/timb418/internship_backend_2022)
 
Для запуска: cd в корень проекта, docker-compose up --build -d

Эндпоинты:
/profit - Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.  

/user-balance - Метод получения баланса пользователя. Принимает id пользователя.  

/reserve - Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.  

/acknowledge - Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. Принимает id пользователя, ИД услуги, ИД заказа, сумму.

Формат ответа - 200 / 4xx коды

Примеры запросов:  
/profit  
    {
    "UserId": "timur10",
    "MoneyAmount": "1000"
    }

/user-balance  
    {
    "UserId": "timur10"
    }

/reserve  
{
"UserId":      "timur10",
"ServiceId":   "456",
"OrderId"   :  "sdvbidsvbgruer",
"MoneyAmount": "50.01"
}


/acknowledge  
{
"UserId":      "timur10",
"ServiceId":   "888",
"OrderId"   :  "sdvbidsvbgruer",
"MoneyAmount": "90.51"
}
