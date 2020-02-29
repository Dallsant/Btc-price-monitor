import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


const routes: Routes = [
  {
    path: 'btc',
    loadChildren: './modules/btc-price-logging/btc-price-logging.module#BtcPriceLoggingModule'
  },
  {
    path: 'bs',
    loadChildren: './modules/bs-price-logging/bs-price-logging.module#BsPriceLoggingModule'
  }
  ,
  {
    path: 'crypto',
    loadChildren: './modules/crypto-price-logging/crypto-price-logging.module#CryptoPriceLoggingModule'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
