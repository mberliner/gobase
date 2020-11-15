Servicios para las entidades del proyecto

User
Persona

Capas:

Domain
	Entidades de negocio débiles (sin lógica)
Repository
	Acceso a datos
Service
	Resuelve la lógica del negocio



Llamadas para probar por fuera:

curl -X POST -H "Content-Type: application/json"  -d '{"nombre": "Luis", "apellido": "Rolix", "usuario": "LRolix", "edad": "12", "password": "pass"}'     http://localhost:8080/user

curl -X POST -H "Content-Type: application/json"  -d '{"usuario": "LRolix", "password": "pass"}'     http://localhost:8080/user/login
