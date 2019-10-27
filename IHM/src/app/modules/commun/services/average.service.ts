import { Injectable } from '@angular/core';
import {ApiService} from './api.service';
import TsMap from 'ts-map';

@Injectable({
  providedIn: 'root'
})
export class AverageService {
  airportsAverageList: TsMap<string, number> = new TsMap<any, any>([]);
  minMaxList: number[][] = [[9999, -9999], [9999, -9999], [9999, -9999]];

  constructor(private apiServ: ApiService) {
    apiServ.doGetRequest(ApiService.API_GET_ALL_AIRPORTS).subscribe(data => {
      const container = JSON.parse(JSON.stringify(data));
      this.airportsAverageList = new TsMap<any, any>([]);

      for (const airportData of container) {
        this.airportsAverageList.set(airportData.IdAirport, airportData.Averages);

        for (let i = 0; i < 3; i++) {
          const value = airportData.Averages[i].Average;

          if (value < this.minMaxList[i][0]) {
            this.minMaxList[i][0] = value;
          }

          if (value > this.minMaxList[i][1]) {
            this.minMaxList[i][1] = value;
          }
        }
      }
    });
  }
}
