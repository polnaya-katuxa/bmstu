-- a = 2.0
-- b = 2

-- if b < 3 then 
--     if b == 2 then
--         print(a)
--         print("pupi")
--     else
--         print("b != 2")
--     end
--     print("www")
-- else
--     print("else if")
-- end

a = 0

print("start loop")

for i = 10, 0, -3 do
    print(i)
end

print("end loop")

-- print(a == b)
-- print(a == 5.0)
-- print(b == 2)
-- print(true == false)
-- print(false == false)
-- print("false" == "false")
-- print("ok" == "kk")

-- print(a ~= b)
-- print(a ~= 5.0)
-- print(b ~= 2)
-- print(true ~= false)
-- print(false ~= false)
-- print("false" ~= "false")
-- print("ok" ~= "kk")

-- print(a > b)
-- print(a > 5.0)
-- print(b > 2)
-- print(false > false)
-- print("false" > "false")
-- print("ok" > "kk")

-- print(a >= b)
-- print(a >= 5.0)
-- print(b >= 2)
-- print(false >= false)
-- print("false" >= "false")
-- print("ok" >= "kk")
-- print("ok" >= "pk")

-- print(a < b)
-- print(a < 5.0)
-- print(b < 2)
-- print(true < false)
-- print(false < false)
-- print("false" < "false")
-- print("ok" < "kk")

-- print(a <= b)
-- print(a <= 5.0)
-- print(b <= 2)
-- print(true <= false)
-- print(false <= false)
-- print("false" <= "false")
-- print("ok" <= "kk")

-- print(a ^ b)
-- print(1 ^ 2)
-- print(5 ^ 1)
-- print(1.2 ^ 2.0)
-- print(50.8 ^ 0)
-- print(25 ^ 0.5)
-- print(25 ^ -0.5)


-- Таблица общего вида:
-- hello = "world"
-- bebebe = "bababa"

-- result = hello.." "..bebebe

-- print(result)
-- print(#result + #"pivo")

-- c = nil
-- print(c)

-- print(true or false and false or false)
-- print(true and true)
-- print(false and false)

-- print(true or false)
-- print(true or true)
-- print(false or false)

-- print(5 * 8)
-- a = 6.0
-- b = 5.0
-- print((a - 5) * (b - 1))
-- print(5 * -8)

-- print(5 / 8)
-- print((a - 5) / (b - 1))
-- print(5 / 0.0)
-- print(5 / -8)

-- print(9 // 8)
-- print(5 // 8)
-- print((a - 5) // (b - 1))
-- print(5 // 0)
-- print(9 // -8)

-- print(9 % 8)
-- print(5 % 8)
-- print((a - 5) % (b - 1))
-- print(5 % 0)
-- print(9 % -8)


-- a, b = 5, 6.0
-- a = 5
-- b = 6.0

-- print(b + a)
-- print(a + b)
-- print(result)

-- a, b = b, a
-- print("bebebe")
-- print(a)
-- print(b)

-- a = -3.5
-- b = -6
-- print(a)
-- print(b)

-- a = false
-- print(a)
-- print(not false)

-- print("-----")
-- b = not a
-- print(not a)

-- c = a + b * 10

-- empty = {} -- Пустая таблица
-- empty[1] = "первый"        -- Добавление элемента с целым индексом
-- empty[3] = "второй"        -- Добавление элемента с целым индексом
-- empty["третий"] = "третий" -- Добавление элемента со строковым индексом
-- empty[1] = nil             -- Удаление элемента из таблицы  

-- empty[2], empty[3] = empty[3], empty["третий"]

-- -- Классический массив - строки индексируются по умолчанию целыми числами, начиная с 1
-- days1 = {"понедельник", "вторник", "среда", "четверг", "пятница", "суббота", "воскресенье"}

-- -- Массив с произвольной индексацией
-- days2 = {[0]="воскресенье", [1]="понедельник", [2]="вторник", [3]="среда", [4]="четверг", [5]="пятница", [6]="суббота"}

-- -- Запись (структура) - значения различных типов индексируются литералами
-- person = {tabnum = 123342,                   -- Табельный номер
--           fio = "Иванов Степан Васильевич",  -- Ф.И.О.
--           post = "слесарь-инструментальщик", -- Должность
--           salary = 25800.45,                 -- Оклад
--           sdate = "23.10.2013",              -- Дата приёма на работу
--           bdate = "08.08.1973"}              -- Дата рождения 

-- pfio = person.fio --Обращение к элементу структуры.

-- Множество - индексы используются для хранения значений
-- workDays = {["понедельник"]=true, ["вторник"]=true, ["среда"]=true, ["четверг"]=true, ["пятница"]=true}
-- workDays["суббота"] = true -- Добавление субботы в число рабочих дней
-- workDays["среда"] = nil    -- По средам больше не работаем
-- workDays.hello = false
-- workDays[true] = true
-- workDays[1+1] = true

-- d = "понедельник"
-- -- Проверка, является ли d рабочим днём
-- if workDays[d] then 
--   print (d.." - рабочий день")
-- else
--   print (d.." - выходной день")
-- end

-- for k, v in workDays do
--   print (k, v)
-- end