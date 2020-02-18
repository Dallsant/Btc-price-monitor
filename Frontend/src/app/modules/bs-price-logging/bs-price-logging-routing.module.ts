import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ListBsPricesComponent } from './list-prices/list-bs-prices.component';

const routes: Routes = [
  {
    path: 'list',
    component: ListBsPricesComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BsPriceLoggingRoutingModule { }
