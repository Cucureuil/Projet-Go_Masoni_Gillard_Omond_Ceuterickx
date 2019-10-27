import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomepageComponent } from './components/homepage/homepage.component';
import {CommunModule} from '../commun/commun.module';
import {HomepageRouter} from './homepage-routing.module';

@NgModule({
  declarations: [HomepageComponent],
  imports: [
    CommonModule,
    CommunModule,
    HomepageRouter,
  ],
})
export class HomepageModule { }
