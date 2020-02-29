import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ListCryptoPricesComponent } from './list-prices/list-crypto-prices.component';

const routes: Routes = [
  {
    path: 'list',
    component: ListCryptoPricesComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BsPriceLoggingRoutingModule { }
