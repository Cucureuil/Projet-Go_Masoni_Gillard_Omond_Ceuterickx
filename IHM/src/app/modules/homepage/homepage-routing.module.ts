import {RouterModule, Routes} from '@angular/router';
import {ModuleWithProviders} from '@angular/compiler/src/core';
import {HomepageComponent} from './components/homepage/homepage.component';

const HOMEPAGE_ROUTER: Routes = [
    {path: '', component: HomepageComponent, data: {}}
];

export const HomepageRouter: ModuleWithProviders = RouterModule.forChild(HOMEPAGE_ROUTER);
