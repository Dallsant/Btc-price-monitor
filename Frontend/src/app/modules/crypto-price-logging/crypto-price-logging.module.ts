import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AppMaterialModule } from '../../app-material/app-material.module';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { AlertModule } from 'ngx-alerts';
import { MatPaginatorModule } from '@angular/material/paginator';

import { BsPriceLoggingRoutingModule } from './crypto-price-logging-routing.module';
import { ListCryptoPricesComponent } from './list-prices/list-crypto-prices.component';
import { MatDatepickerModule } from '@angular/material/datepicker';
import {MatGridListModule} from '@angular/material'

@NgModule({
  declarations: [ListCryptoPricesComponent],
  imports: [
    CommonModule,
    MatGridListModule,
    BsPriceLoggingRoutingModule,
    AppMaterialModule,
    ReactiveFormsModule,
    HttpClientModule,
    MatPaginatorModule,
    MatDatepickerModule,
    AlertModule.forRoot({ maxMessages: 5, timeout: 5000, position: 'right' }),
  ],
  entryComponents: []
})
export class CryptoPriceLoggingModule { }
