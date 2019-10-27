import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {ApiService} from '../../../commun/services/api.service';

@Component({
  selector: 'app-airport-page',
  templateUrl: './airport-page.component.html',
  styleUrls: ['./airport-page.component.scss']
})
export class AirportPageComponent implements OnInit {
  airportId : string;

  constructor(private apiServ: ApiService, private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.params.subscribe( data => {
      this.airportId = data.id;
      this.loadData();
    });
  }

  private loadData() {
    this.apiServ.doGetRequest(ApiService.API_GET_AIRPORT_DATA + this.airportId).subscribe( data => {
      console.log(data);
    });
  }
}
