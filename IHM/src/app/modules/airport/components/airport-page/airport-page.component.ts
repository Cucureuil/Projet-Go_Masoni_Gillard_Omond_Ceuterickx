import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {ApiService} from '../../../commun/services/api.service';
import {SensorModel} from "../../../commun/models/sensor.model";

@Component({
  selector: 'app-airport-page',
  templateUrl: './airport-page.component.html',
  styleUrls: ['./airport-page.component.scss']
})
export class AirportPageComponent implements OnInit {
  airportId: string;
  windDataList: SensorModel[] = [];
  tempDataList: SensorModel[] = [];
  pressDataList: SensorModel[] = [];
  startDate: string;
  endDate: string;

  constructor(private apiServ: ApiService, private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.params.subscribe( data => {
      this.airportId = data.id;
      this.loadData();
    });
  }

  private loadData(dates: string = '') {
    this.apiServ.doGetRequest(ApiService.API_GET_AIRPORT_DATA + this.airportId + dates).subscribe( res => {
      const container = JSON.parse(JSON.stringify(res)).Data;
      this.windDataList = [];
      this.tempDataList = [];
      this.pressDataList = [];

      let data: any[] = container[0].Data;
      if (data !== undefined && data != null) {
        for (const sensor of data) {
          this.windDataList.push(Object.assign(new SensorModel(), {
            _idSensor: sensor.IdSensor,
            _idAirport: sensor.IdAirport,
            _value: sensor.Value,
            _dateMeasure: sensor.DateMeasure
          }));
        }
      }

      data = container[1].Data;
      if (data !== undefined && data != null) {
        for (const sensor of data) {
          this.tempDataList.push(Object.assign(new SensorModel(), {
            _idSensor: sensor.IdSensor,
            _idAirport: sensor.IdAirport,
            _value: sensor.Value,
            _dateMeasure: sensor.DateMeasure
          }));
        }
      }

      data = container[2].Data;
      if (data !== undefined && data != null) {
        for (const sensor of data) {
          this.pressDataList.push(Object.assign(new SensorModel(), {
            _idSensor: sensor.IdSensor,
            _idAirport: sensor.IdAirport,
            _value: sensor.Value,
            _dateMeasure: sensor.DateMeasure
          }));
        }
      }
    });
  }

  dateToString(date: Date) {
    return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
  }


  onFilterByDate() {
    if ((this.endDate === undefined || this.endDate == null) && (this.startDate === undefined || this.startDate == null)) {
      this.loadData();
    } else {
      let startDate = this.startDate;
      let endDate = this.endDate;
      if (startDate === undefined || startDate == null) {
        startDate = '1950-01-01';
      }
      if (endDate === undefined || endDate == null) {
        endDate = '2500-01-01';
      }
      this.loadData('/' + startDate + '-00-00-00' + '/' + endDate + '-00-00-00');
    }
  }
}
