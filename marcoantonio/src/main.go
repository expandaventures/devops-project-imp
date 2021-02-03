package main

import (
	"encoding/json"
	_ "errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var page = `<html>
<head>
    <title>DevOps Project Aplicacion de Rastreo</title>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
  <script>
        $(document).ready(function(){
        $("button").click(function(){
        var user = '';
        var streetName = '';
        var streetNumber = '';
        var city = '';
        var state = '';
        var country = '';
        var username = '';
        var contadorLocation = 0;
        var contadorLogin = 0;
                
            $.ajax({
                url: 'https://randomuser.me/api/',
                dataType: 'json',
                success: function(data) {
                for (i in data.results) {
                console.log(data.results);
                  for (j in data.results[i].name) {
                    user += data.results[i].name[j];
                    user += ' '
                  }
                for (j in data.results[i].location) {
                    while (contadorLocation <= 0)
                    {
                        streetName += data.results[i].location[j].name
                        streetNumber += data.results[i].location[j].number
                        city += data.results[i].location.city;
                        state += data.results[i].location.state;
                        country += data.results[i].location.country;
                        contadorLocation++;
                    }
                }
                for (j in data.results[i].login) {
                    while (contadorLogin <= 0)
                    {
                        username += data.results[i].login.username;
                        contadorLogin++;
                    }

                }
              }
            $('#name').html('<h2>Nombre: ' + user + '</h2>');
            $('#location').html('<h2>Domicilio: ' + streetName + ' ' +streetNumber + ' in ' +  city +  ', ' +  state +  ', ' +  country +  '.' + '</h2>');
            $('#username').html('<h2>Logeado con usuario: ' + username + '</h2>');
                }
            });
        });
    });
  </script>
</head>
<body>
<div id="div1"><h1>Bienvenido Aplicacion de Rastreo!</h1></div>
<button>Generar Usuario</button>
<div id="name"></div>
<div id="location"></div>
<div id="username"></div>

</body>
</html>`

type Car struct {
	Id string `json:"Id"`
	Brand string `json:"Brand"`
	Plate string `json:"Plate"`
	Location string `json:"Location"`
}

var Cars []Car

func homePage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, page)
	fmt.Println("Endpoint Hit: user")
}

func returnAllCars(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllCars")
	json.NewEncoder(w).Encode(Cars)
}

func returnSingleCar(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnSingleCar")
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: " + key)
	for _, car := range Cars {
		if car.Id == key {
			json.NewEncoder(w).Encode(car)
		}
	}
}

func createNewCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewCar")
	respBytes, _ := ioutil.ReadAll(r.Body)
	var car Car
	json.Unmarshal(respBytes, &car)
	Cars = append(Cars, car)
	json.NewEncoder(w).Encode(car)

}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteCar")
	vars := mux.Vars(r)
	id := vars["id"]
	for index, car := range Cars {
		if car.Id == id {
			Cars = append(Cars[:index], Cars[index+1:]...)
		}
	}

}

func updateCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateCar")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Reading Value for Key :", id)
	respBytes, _ := ioutil.ReadAll(r.Body)
	var aut Car
	fmt.Println(aut)
	for index, aut := range Cars {
		if aut.Id == id {
			var car Car
			json.Unmarshal(respBytes, &car)
			Cars[index] = car
			json.NewEncoder(w).Encode(car)
			fmt.Println("Request :", Cars[index])
			fmt.Println("Current :", aut)
			fmt.Println("Updated :", Cars[index])
		}
	}


}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/cars", returnAllCars)
	myRouter.HandleFunc("/car/{id}", returnSingleCar)
	myRouter.HandleFunc("/car", createNewCar).Methods("POST")
	myRouter.HandleFunc("/card/{id}", deleteCar).Methods("DELETE")
	myRouter.HandleFunc("/caru/{id}", updateCar).Methods("PUT")
	log.Fatal(http.ListenAndServe(":2021", myRouter))

	//fmt.Println("Rest API v1.0 - IBM SRE Random Users")
	//http.HandleFunc("/user", handler)
	//log.Fatal(http.ListenAndServe(":2000", nil))
}

func main() {
	fmt.Println("Rest API v1.0 - DevOps Project Aplicacion de Rastreo")
	Cars = []Car{
		{Id: "1", Brand: "Porche", Plate: "MKT-2015", Location: "41°24'12.2 N 2°10'26.5E"},
		{Id: "2", Brand: "BMW", Plate: "ZCY-2021", Location: "41 24.2028, 2 10.4418"},
		{Id: "3", Brand: "Mercedes", Plate: "AMD-2018", Location: "49 24.2048, 2°10'26.5E"},
		{Id: "4", Brand: "Honda", Plate: "AAZ-2020", Location: "51 24.2078, 2 10.4418"},
		{Id: "5", Brand: "Mazda", Plate: "GKV-2019", Location: "61 24.2038, 2°10'26.5E"},
	}
	handleRequests()
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}