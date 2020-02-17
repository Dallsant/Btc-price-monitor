import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AppMaterialModule } from '../../app-material/app-material.module';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { AlertModule } from 'ngx-alerts';
import { MatPaginatorModule } from '@angular/material/paginator';

import { BtcPriceLoggingRoutingModule } from './btc-price-logging-routing.module';
import { ListBtcPricesComponent } from './list-prices/list-prices.component';
import { MatDatepickerModule } from '@angular/material/datepicker';

@NgModule({
  declarations: [ListBtcPricesComponent],
  imports: [
    CommonModule,
    BtcPriceLoggingRoutingModule,
    AppMaterialModule,
    ReactiveFormsModule,
    HttpClientModule,
    MatPaginatorModule,
    MatDatepickerModule,
    AlertModule.forRoot({ maxMessages: 5, timeout: 5000, position: 'right' }),
  ],
  entryComponents: []
})
export class BtcPriceLoggingModule { }
