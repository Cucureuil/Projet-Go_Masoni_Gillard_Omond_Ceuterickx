# Projet-Go_Masoni_Gillard_Omond_Ceuterickx

# Projet-Go_Masoni_Gillard_Omond_Ceuterickx

<h1>INSTALL</h1>
<p>
  OPEN CMD OR BASH : <br/>
go get "github.com/gorilla/mux"<br/>
go get "github.com/eclipse/paho.mqtt.golang"<br/>
</p>

<h1>RUN</h1>

<h3>1 - RUN MOSQUITTO SERVER</h3>
<p>
  OPEN CMD :<br/>
mosquitto -v <br/>
</p>

<h3>2 - API</h3> 
<p>
  RUN api.exe
</p>

<h3>3 - SUBSCRIBER</h3>
<p>
  RUN AirportA.exe<br/>
  RUN AirportB.exe<br/>
</p>

<h3>4 - PUBLISHER</h3>
<p>
  RUN WindSensor.exe<br/>
  RUN PressureSensor.exe<br/>
  RUN TemperatureSensor.exe<br/>
</p>

<h1>API ROUTING</h1>

<p>
/airports -- list of all airport<br/>
/airport/{id} -- data for one airport<br/>
/sensor/{id} -- data for one sensor<br/>
/date/{date} -- data for one date (YYYY-MM-DD-hh-mm-ss)<br/>
/date/{startDate}/{endDate} -- data between two dates (YYYY-MM-DD-hh-mm-ss)<br/>
/average -- averages for each types of sensor<br/>
</p>

