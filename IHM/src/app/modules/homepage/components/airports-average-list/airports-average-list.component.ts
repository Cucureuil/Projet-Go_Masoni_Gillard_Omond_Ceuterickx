import { Component, OnInit } from '@angular/core';
import {AverageService} from '../../../commun/services/average.service';
import TsMap from 'ts-map';

@Component({
  selector: 'app-airports-average-list',
  templateUrl: './airports-average-list.component.html',
  styleUrls: ['./airports-average-list.component.scss']
})
export class AirportsAverageListComponent implements OnInit {
  airportsList = new TsMap();

  constructor(private averageServ: AverageService) { }

  ngOnInit() {
    this.averageServ.airportsAverageList.subscribe( data => {
      this.airportsList = data;
    });
  }
}
