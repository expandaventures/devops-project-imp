# SinTrafico API

DevOps Project Aplicacion de Rastreo.

## Descripcion

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

### Dependencias

* go version go1.15.6 linux/amd64
* Linux Ubuntu 20.04.1 LTS
* Docker 19.03.8
* Kubernetes v1.20
* git version 2.25.1
* IDE GoLand 2020.3.2
 
### Instalacion

* Descargar con el comando git clone https://github.com/userIbmTest/sintrafico.git

### Ejecutar el programa

* Ejecutar API con el lenguaje Go
```
cd sintrafico/src
go mod init test/exam
go test -v
go build main.go
./main
Validar en el navegador http://localhost:2021/
```
* Ejecutar API con Docker
```
cd sintrafico
docker build -t test/exam:latest .
docker run -p 127.0.0.1:2021:2021/tcp test/exam 
Validar en el navegador http://127.0.0.1:2021/
```
* Ejecutar API con Kubernetes
```
cd sintrafico
docker build -t test/exam:latest .
kubectl create deployment --image=test/exam server
kubectl edit deployments server --> update  imagePullPolicy: IfNotPresent 
kubectl expose deployment server --port=2021 --name=server-service
kubectl get svc server-service --> get the CLUSTER-IP
Validar en el navegador http://<CLUSTER-IP>:2021/
```
* Ejecutar end-point para mostrar todos los carros metodo GET
```
http://localhost:2021/cars
```
* Ejecutar end-point para mostrar un solo carro metodo GET
```
http://localhost:2021/car/1
```
* Ejecutar end-point para agregar un carro methodo POST
```
http://localhost:2021/car

Cuerpo del Jason
{
    "Id": "10", 
    "Brand": "Honda", 
    "Plate": "HKL2020", 
    "Location": "51 24.2078, 2 10.4418" 
}
```
* Ejecutar end-point para borrar carro metodo DELETE
```
http://localhost:2021/card/2
```
* Ejecutar end-point para agregar un carro methodo PUT
```
http://localhost:2021/caru/5

Cuerpo del Jason
{
    "Id": "5", 
    "Brand": "m", 
    "Plate": "IGH786", 
    "Location": "51 24.2078, 2 10." 
}
```

## Ayuda

En caso de tener errores validar el contenedor con el comando indicado abajo y ejecutar ./exam dentro del contenedor.
Debe de aparecer el mensaje - Rest API v1.0 - DevOps Project Aplicacion de Rastreo 
```
docker run -t -i --rm test/exam bash
```
En caso de necesitar entrar al contenedor ejecutar el comando docker ps para obtener el nombre del contenedor  <nombre del contenedor> 
y validar si la aplicacion esta ejecutandose con el comando pd -fea debe de ver el mensaje siguiente:

```
docker exec -it <container name> /bin/bash
root@6bdbbda6b4f1:/app# ps -fea 
UID          PID    PPID  C STIME TTY          TIME CMD
root           1       0  0 01:25 ?        00:00:00 ./exam
root          11       0  0 01:31 pts/0    00:00:00 /bin/bash
root          18      11  0 01:32 pts/0    00:00:00 ps -fea
```
## Autor

Marco Antonio Gomez Aguilar
https://linkedin.com/in/pmpmg

## Historial de la version

* 1.0
    * Codigo inicial

