import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ListBtcPricesComponent } from './list-prices/list-prices.component';

const routes: Routes = [
  {
    path: 'list',
    component: ListBtcPricesComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BtcPriceLoggingRoutingModule { }
