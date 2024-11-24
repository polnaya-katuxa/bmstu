domains
 surname, name, phone, city, street, bank, account = symbol.
 building, flat, size, floors = integer.
 model, color, number = symbol.
 price, sum = integer.
 
 address = address(city, street, building, flat).
 
 property = car(name, color, number, price);
 	    building(name, floors, price);
 	    land(name, size, price);
 	    water_vehicle(name, color, number, price).

predicates
 nondeterm phone_note(surname, phone, address).
 nondeterm ownership(surname, property).
 nondeterm bank_note(surname, bank, account, sum).
 
 nondeterm all_properties(surname, name).
 nondeterm all_properties_prices(surname, name, price).
 
 nondeterm car_price(surname, price).
 nondeterm building_price(surname, price).
 nondeterm land_price(surname, price).
 nondeterm water_vehicle_price(surname, price).
 nondeterm sum_properties_price(surname, price).

clauses
 phone_note("Eliza", "111111", address("New-York", "One", 1, 2)).
 phone_note("Dabby", "222222", address("London", "Two", 3, 4)).
 phone_note("Eliza", "333333", address("Paris", "Three", 5, 6)).
 phone_note("Donnie", "444444", address("Oslo", "Four", 7, 8)).
 phone_note("Darwin", "555555", address("Minsk", "Five", 7, 8)).
 
 bank_note("Eliza", "Sber", "usual", 1000).
 bank_note("Donnie", "Tinkoff", "credit", 1500).
 bank_note("Donnie", "VTB", "credit", 2000).
 bank_note("Darwin", "Sber", "usual", 1000).
 bank_note("Dabby", "VTB", "credit", 2000).
 
 ownership("Eliza", car("BMW", "black", "AAA111", 1000)).
 ownership("Donnie", car("Ford", "yellow", "BBB222", 1500)).
 ownership("Donnie", water_vehicle("Yacht", "pink", "CCC333", 2000)).
 ownership("Darwin", building("Empire State", 57, 1000)).
 ownership("Dabby", land("Dacha", 500, 2000)).
 ownership("Dabby", building("Green Palace", 4, 5000)).
 ownership("Dabby", car("Ford", "pink", "GGG777", 1500)).
 
 all_properties(Surname, Name) :- ownership(Surname, car(Name, _, _, _)); 
 				  ownership(Surname, building(Name, _, _)); 
 				  ownership(Surname, land(Name, _, _)); 
 				  ownership(Surname, water_vehicle(Name, _, _, _)).
 
 all_properties_prices(Surname, Name, Price) :- ownership(Surname, car(Name, _, _, Price)); 
 				  ownership(Surname, building(Name, _, Price)); 
 				  ownership(Surname, land(Name, _, Price));
 				  ownership(Surname, water_vehicle(Name, _, _, Price)).
 				  
 car_price(Surname, Price) :- ownership(Surname, car(_, _, _, Price)), !; Price = 0.
 building_price(Surname, Price) :- ownership(Surname, building(_, _, Price)), !; Price = 0.
 land_price(Surname, Price) :- ownership(Surname, land(_, _, Price)), !; Price = 0.
 water_vehicle_price(Surname, Price) :- ownership(Surname, water_vehicle(_, _, _, Price)), !; Price = 0.
 				  
 sum_properties_price(Surname, S) :- car_price(Surname, P1),
 				  building_price(Surname, P2),
 				  land_price(Surname, P3),
 				  water_vehicle_price(Surname, P4),
 				  S = P1 + P2 + P3 + P4.
 
 goal
  %all_properties("Dabby", Name).
  %sum_properties_price("Donnie", SumPrice).
  all_properties_prices("Donnie", Name, Price).

  
  