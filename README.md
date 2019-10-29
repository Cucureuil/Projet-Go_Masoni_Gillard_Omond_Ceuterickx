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
/average/type/{sensorType} -- average by sensor type -- <b>On ne fait pas</b> <br/> 
/average/type/{sensorType}/{startDate}/{endDate} -- average by sensor type between two dates -- <b>On ne fait pas</b> <br/>  
</p>

<h1>REDIS</h1>

<p>
  Installation : https://redis.io/<br/>
  Lancer le serveur :
  <ul>
    <li>
      Ouvrir terminal à la racine du dossier d'installation
    </li>
    <li>
      Taper "redis-server"
    </li>
  </ul>
  
  Une fois le serveur lancé, ouvrir un deuxième terminal à la racine et taper "redis-cli" pour accéder au terminal de redis.<br/>
  Dans ce terminal, ajouter les données de tests (copy paste la liste suivant dans le terminal) :<br/>
  <br/>
  sadd idAirports AAA1 BBB1<br/>
  <br/>
  zadd "AAA1:Atmospheric pressure" 1571855613 1:1:25.02<br/>
  zadd "AAA1:Atmospheric pressure" 1571855623 2:1:521.32<br/>
  zadd "AAA1:Temperature" 1571855633 3:2:32.52<br/>
  zadd "AAA1:Temperature" 1571855643 4:2:5.3<br/>
  zadd "AAA1:Wind speed" 1571855643 5:3:25.3<br/>
  zadd "BBB1:Atmospheric pressure" 1571855653 6:4:648.5<br/>
  zadd "BBB1:Atmospheric pressure" 1571855663 7:4:985.30<br/>
  zadd "BBB1:Temperature" 1571855673 8:5:74.9<br/>
</p>

<h1>IHM</h1>

<p>
  Installation : angular
  Lancer le serveur :
  <ul>
    <li>
      Ouvrir terminal à la racine du dossier d'IHM
    </li>
    <li>
      Taper "ng serve --o"
    </li>
  </ul>
  
  Une fois le serveur lancé, le navigateur va s'ouvrir sur la page de tableau de bord
</p>
