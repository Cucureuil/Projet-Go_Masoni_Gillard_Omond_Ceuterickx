import { Component, OnInit } from '@angular/core';
import {ApiService} from '../../../commun/services/api.service';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent implements OnInit {
  airportsList = [];

  constructor(private apiServ: ApiService) { }

  ngOnInit() {
    this.apiServ.doGetRequest(ApiService.API_GET_ALL_AIRPORTS).subscribe( data => {
    });
  }
}
