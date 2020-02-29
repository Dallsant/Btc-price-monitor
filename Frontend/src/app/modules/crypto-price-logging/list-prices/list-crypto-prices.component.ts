import { Component, OnInit, ViewChild, Inject, ChangeDetectorRef } from '@angular/core';
import { MatPaginator, MatTableDataSource } from '@angular/material';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { AlertService } from 'ngx-alerts';
import { CryptoPriceLoggingService } from '../crypto-price-logging.service';
import { backend_url } from '../../../app.component';
import { Chart } from 'chart.js';


export interface CryptoPriceRecord {
  ID: number;
  eth: number;
  ltc: number;
  xrp: number;
  eos: number;
  xlm: number;
  bnb: number;
  updated: string;
  timestamp: number;
}


@Component({
  selector: 'app-list-prices',
  templateUrl: './list-crypto-prices.component.html',
  styleUrls: ['./list-crypto-prices.component.css']
})
export class ListCryptoPricesComponent implements OnInit {
  chart = [];
  @ViewChild('cryptoPricesPaginator', { static: false }) cryptoPricesPaginator: MatPaginator;

  cryptoPricedisplayedColumns: string[] = ['ID', 'value', 'updated'];
  LastCryptoPrice: CryptoPriceRecord
  CryptoPricesDataSource: MatTableDataSource<CryptoPriceRecord[]>;
  constructor(
    public dialog: MatDialog,
    public CryptoPriceLoggingService: CryptoPriceLoggingService,
    private alertService: AlertService,
  ) { }

  listCryptoPrices() {
    this.CryptoPriceLoggingService.getCryptoPrices().subscribe(data => {
      this.CryptoPricesDataSource = new MatTableDataSource(data.reverse());
      this.CryptoPricesDataSource.paginator = this.cryptoPricesPaginator;
    })
  }

  getCryptoPrice() {
    this.CryptoPriceLoggingService.getCryptoPrice().subscribe(data => {
      this.LastCryptoPrice = data;
    })
  }

  ngOnInit() {
    this.listCryptoPrices();
    this.getCryptoPrice();
  }
}