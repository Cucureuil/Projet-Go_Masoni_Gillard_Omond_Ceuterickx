import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AirportPageComponent } from './components/airport-page/airport-page.component';
import {AirportRouter} from './airport-routing.module';
import {CommunModule} from '../commun/commun.module';
import {FormsModule} from "@angular/forms";

@NgModule({
  declarations: [AirportPageComponent],
    imports: [
        CommonModule,
        CommunModule,
        AirportRouter,
        FormsModule
    ]
})
export class AirportModule { }
