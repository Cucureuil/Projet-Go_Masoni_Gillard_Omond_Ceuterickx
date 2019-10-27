import { Injectable } from '@angular/core';
import {ApiService} from './api.service';
import TsMap from 'ts-map';
import {BehaviorSubject} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AverageService {
  private airportsAverageSource = new BehaviorSubject<TsMap<string, number> >(new TsMap<any, any>([]));
  airportsAverageList = this.airportsAverageSource.asObservable();

  minMaxList: number[][] = [[9999, -9999], [9999, -9999], [9999, -9999]];

  constructor(private apiServ: ApiService) {
    apiServ.doGetRequest(ApiService.API_GET_ALL_AIRPORTS).subscribe(data => {
      const container = JSON.parse(JSON.stringify(data));
      const list = new TsMap<string, number>([]);

      for (const airportData of container) {
        list.set(airportData.IdAirport, airportData.Averages);

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
      this.updateCurrentAirportList(list);
    });
  }

  private updateCurrentAirportList(params: TsMap<string, number>) {
    this.airportsAverageSource.next(params);
  }
}
