import { Component, OnInit } from '@angular/core';
import {AverageService} from '../../../commun/services/average.service';

@Component({
  selector: 'app-min-max-mesures',
  templateUrl: './min-max-mesures.component.html',
  styleUrls: ['./min-max-mesures.component.scss']
})
export class MinMaxMesuresComponent implements OnInit {
  minMaxList = [[0, 0], [0, 0], [0, 0]];
  constructor(private averageServ: AverageService) { }

  ngOnInit() {
    this.minMaxList = this.averageServ.minMaxList;
  }

}
