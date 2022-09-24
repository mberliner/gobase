Servicios para las entidades del proyecto

Persona

Capas:

Domain
	Entidades de negocio débiles (sin lógica)
Service
	Resuelve la lógica del negocio
Repository
	Acceso a datos

Patrones:

MVC
	domain, controller
Inversión de Control e Inyección de dependencias
	Config.go en cada package

Llamadas para probar microservicios:


curl -X GET -H "Content-Type: application/json"  http://localhost:8080/personas

curl -X POST -H "Content-Type: application/json" -d '{"Nombre": "Levon","Apellido": "Nacarian","FechaNacimiento": "12-02-2234"}'  http://localhost:8080/personas

curl -X GET -H "Content-Type: application/json"  http://localhost:8080/personas/1

curl -X PUT -H "Content-Type: application/json" -d '{"Nombre": "Levon1","Apellido": "Nacarian1","FechaNacimiento": "12-02-2234"}'  http://localhost:8080/personas/1

curl -X PATCH -H "Content-Type: application/json" -d '{"Nombre": "Levon1", "Apellido": "Nxxxxaaa"}'  http://localhost:8080/personas/1

curl -X DELETE -H "Content-Type: application/json"  http://localhost:8080/personas/1
