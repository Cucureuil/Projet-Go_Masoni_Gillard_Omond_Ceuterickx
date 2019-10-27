import {RouterModule, Routes} from '@angular/router';
import {ModuleWithProviders} from '@angular/compiler/src/core';
import {AirportPageComponent} from './components/airport-page/airport-page.component';

const AIRPORT_ROUTER: Routes = [
    {path: '', component: AirportPageComponent, data: {}}
];

export const AirportRouter: ModuleWithProviders = RouterModule.forChild(AIRPORT_ROUTER);
