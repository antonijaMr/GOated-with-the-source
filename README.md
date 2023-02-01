Baza podataka koristenjem skladistenja key-value (osnovna mapa u golangu)

Nasa key-value mapa se sastoji od kljuca koji je string, on je indirektno povezan sa Itemima koji sadrze genericki tip content i listu svojih kljuceva jer jedan item moze imati vise kljuceva.
Implementacija ove mape je ostvarena pomocu standardne mape u kojoj je kljuc string, vrijednost je vezana za strukturu map_value koja sadrzi dva slice-a.(Jedan slice sa indexima oslobodjenih pozicija, drugi slice zauzetih indexa)
Indexi nam sluze da pristupimo slicu Item_list koji sadrzi sve Iteme

Nasa baza sadrzi CRUD API tj. sadrzi create, read, update i delete funkcije.
Sadrzi web aplikaciju preko koje mozemo koristiti bazu podataka, te sve podatke mozemo spremiti u CSV file.

Use case:
Sama poenta baze podataka jeste da se koristi kao "search engine", te kao takva idealna je za baze podataka koje medjusobno dijele osobine (tag.ove/keys)
Jedan primjer takve baze bi bila biblioteka clanaka gdje jedan clanak moze da ima vise aplikabilnih tema i da jedna tema ima vise clanaka.
Omogoceno koristenje tagova koji pomazu pri pretrazi i sortiranju data entries
Jos jedan primjer je povezevinaje zanrova sa filmovima koje pripadaju.
Svaki od opisa i tagova/zanrova mozemo urediti ili sveukupno clanak/film izbrisati.
