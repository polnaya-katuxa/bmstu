domains
 name, phone, city, street = symbol.
 building, flat = integer.
 address = address(city, street, building, flat).
 model, color, number = symbol.
 price = integer.

predicates
 nondeterm phone_note(name, phone, address).
 nondeterm car_note(name, model, color, price, number).
 nondeterm note_by_car(model, color, name, phone, city).

clauses
 phone_note("Eliza", "111111", address("New-York", "One", 1, 2)).
 phone_note("Dabby", "222222", address("London", "Two", 3, 4)).
 phone_note("Eliza", "333333", address("Paris", "Three", 5, 6)).
 phone_note("Donnie", "444444", address("Oslo", "Four", 7, 8)).
 phone_note("Darwin", "555555", address("Minsk", "Five", 7, 8)).
 
 car_note("Eliza", "BMW", "black", 1000, "AAA111").
 car_note("Donnie", "Ford", "yellow", 1500, "BBB222").
 car_note("Donnie", "Mercedes", "pink", 2000, "CCC333").
 car_note("Darwin", "BMW", "black", 1000, "DDD444").
 car_note("Dabby", "Mercedes", "white", 2000, "EEE555").
 car_note("Dabby", "Lada", "black", 500, "FFF666").
 car_note("Dabby", "Ford", "pink", 1500, "GGG777").
 
 note_by_car(Model, Color, Name, Phone, City) :- car_note(Name, Model, Color, _, _), phone_note(Name, Phone, address(City, _, _, _)).
 
goal
 %car_note("Eliza", Model, Color, Price, Number).
 %phone_note(Name, Phone, address("Oslo", Street, House, Room)).
 %note_by_car("BMW", "black", Name, Phone, City).
 %phone_note(Name, "444444", address("Oslo", "Four", 7, 8)).
 %phone_note(Name, Phone, address(City, Street, House, Room)).
 %phone_note("Dabby", _, address(City, Street, House, Room)).
 %car_note("Dabby", _, address(City, Street, House, Room)).
 phone_note(_, "444444", address(City, _, _, _)).