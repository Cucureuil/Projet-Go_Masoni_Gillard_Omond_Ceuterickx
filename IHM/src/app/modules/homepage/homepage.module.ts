import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomepageComponent } from './components/homepage/homepage.component';
import {CommunModule} from '../commun/commun.module';
import {HomepageRouter} from './homepage-routing.module';
import { MinMaxMesuresComponent } from './components/min-max-mesures/min-max-mesures.component';

@NgModule({
  declarations: [HomepageComponent, MinMaxMesuresComponent],
  imports: [
    CommonModule,
    CommunModule,
    HomepageRouter,
  ],
})
export class HomepageModule { }
