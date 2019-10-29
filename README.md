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
/airport/{id}/{sensorType} -- data for one airport by sensorType<br/>
/airport/{id}/{startDate}/{endDate} -- data for one airport between two dates<br/>
/airport/{id}/{sensorType}/{startDate}/{endDate} -- data for one airport by sensorType and between two dates<br/>  
<br/>  
/sensor/{id} -- data for one sensor<br/>
/sensor/{startDate}/{endDate} -- data for one sensor between two dates<br/>
<br/>
/average/airport/{id}/{sensorType} -- averages by sensor type and airport<br/>
/average/airport/{id}/{sensorType}/{startDate}/{endDate} -- averages by sensor type and airport and between two dates<br/>
/average/sensor/{id} -- average of a sensor<br/>
/average/sensor/{id}/{startDate}/{endDate} -- average of a sensor between two dates<br/>
/average/type/{sensorType} -- average by sensor type<br/>
/average/type/{sensorType}/{startDate}/{endDate} -- average by sensor type between two dates<br/>  
</p>

