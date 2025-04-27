days1 = {"понедельник", "вторник", "среда", "четверг", "пятница", "суббота", "воскресенье"}
workDays = {["понедельник"]=true, ["вторник"]=true, ["среда"]=true, ["четверг"]=true, ["пятница"]=true, ["суббота"]=false, ["воскресенье"]=false}

i = 0
while i < 7 do
    if workDays[days1[i]] then
        days1[i] = days1[i].." - рабочий"
    else
        days1[i] = days1[i].." - выходной"
    end
    i = i + 1
end

print(days1)

days1 = {"понедельник", "вторник", "среда", "четверг", "пятница", "суббота", "воскресенье"}

i = i - 1

repeat
    if workDays[days1[i]] and workDays[days1[i - 1]] then
        print("с "..days1[i].." по "..days1[i - 1].." можно выспаться")
    else
        print("с "..days1[i].." по "..days1[i - 1].." нельзя выспаться")
    end
    i = i - 1
until i > 0