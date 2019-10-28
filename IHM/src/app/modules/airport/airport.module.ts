import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AirportPageComponent } from './components/airport-page/airport-page.component';
import {AirportRouter} from './airport-routing.module';
import {CommunModule} from '../commun/commun.module';

@NgModule({
  declarations: [AirportPageComponent],
  imports: [
    CommonModule,
    CommunModule,
    AirportRouter
  ]
})
export class AirportModule { }
