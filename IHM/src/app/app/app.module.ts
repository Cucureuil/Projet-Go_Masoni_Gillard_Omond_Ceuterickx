import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {AppComponent} from './component/app.component';
import {AppRoutingModule} from './app-routing.module';
import {CommunModule} from '../modules/commun/commun.module';



@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    CommunModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
