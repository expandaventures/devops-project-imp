# DevOps Project
Dev Ops Engineering Project

## Aplicación de Rastreo

El objetivo es construir una applicación web donde se administre una flotilla de vehiculos y se pueda ver su posición.
Cada vehiculo tiene los siguientes datos:

1. Id de Vehiculo
2. Placas
3. Ultima posición conocida (lat,lon)

## Requerimientos

1. Construir un API HTTP Rest con  en la que se pueda 
 - Insertar un vehiculo.
 - Actualizar un vehiculo.
 - Borrar cada Vehículo.
2. La API deben de contar con autenticación de usuario. (e.g. usuario/contrasena, API Key, Oauth2, etc)
3. Cada Usuario solo debe de poder interactuar con los vehiculos creados por él mismo.
4. Levantar el API en un servicio en la nube como Amazon Web Services o Google Cloud o Heroku. 


## Entregables

1. Crea un repositorio publico en tu cuenta de github.
2. Proveer instrucciones para instalar y levantar en ambiente la nube.
3. Proveer ejemplos con CURL para Insertar, actualizar y borrar cada vehículo.

 - Te recomendamos que uses Flask o Django Rest Framework para hacer tu app
 - Bonus points si agregas Unit Tests. 
 - Bonus points si usas alguna Herramienta como Docker, CloudFormation, Terraform, Helm o cualquier otra herramienta de Infrastructure-as-Code.
 - Bonus points si existe la posibiliadad de hacer deployment local.

La evaluación tendra la siguiente forma:

1. Back End: 35%
   1. CRUD 20%
   2. Auth 15%
2. Deployment en la Nube: 35%
3. Estructura y legilibilidad del código, incluyendo el uso de buenas prácticas: 30%
4. Bonus points: 30% extra (10% unit tests, 10% deployment con Infrastructure-as-Code, 10% deployment local).

Suerte! 
