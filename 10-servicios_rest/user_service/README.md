Servicios para las entidades del proyecto

User

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

curl -X POST -H "Content-Type: application/json"  -d '{"nombre": "Luis", "apellido": "Rolix", "usuario": "LRolix", "edad": "12", "password": "pass"}'     http://localhost:8080/users

curl -X POST -H "Content-Type: application/json"  -d '{"usuario": "LRolix", "password": "pass"}'     http://localhost:8080/users/login


