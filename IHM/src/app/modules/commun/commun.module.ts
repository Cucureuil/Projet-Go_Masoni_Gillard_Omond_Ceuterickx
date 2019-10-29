import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HeaderComponent } from './components/header/header.component';
import {ApiService} from './services/api.service';
import {HttpClientModule} from '@angular/common/http';

@NgModule({
    declarations: [HeaderComponent],
    exports: [
        HeaderComponent,
    ],
    imports: [
        CommonModule,
        HttpClientModule,
    ],
    providers: [
        ApiService
    ]
})
export class CommunModule { }
