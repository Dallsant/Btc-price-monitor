import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AppMaterialModule } from '../../app-material/app-material.module';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { AlertModule } from 'ngx-alerts';
import { MatPaginatorModule } from '@angular/material/paginator';

import { BsPriceLoggingRoutingModule } from './bs-price-logging-routing.module';
import { ListBsPricesComponent } from './list-prices/list-bs-prices.component';
import { MatDatepickerModule } from '@angular/material/datepicker';

@NgModule({
  declarations: [ListBsPricesComponent],
  imports: [
    CommonModule,
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
export class BsPriceLoggingModule { }
