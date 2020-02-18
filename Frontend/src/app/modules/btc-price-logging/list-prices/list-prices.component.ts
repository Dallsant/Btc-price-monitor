import { Component, OnInit, ViewChild, Inject, ChangeDetectorRef  } from '@angular/core';
import { MatPaginator, MatTableDataSource } from '@angular/material';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { AlertService } from 'ngx-alerts';
import { BtcPriceLoggingService } from '../btc-price-logging.service';
import { backend_url } from '../../../app.component';
import { Chart } from 'chart.js';

export interface ChartAxis {
  x: string;
  y: number;
}

export interface BtcPriceRecord {
  ID: number;
  value: number;
  updated: string;
  timestamp: number;
}

export interface BtcPriceElement {
  chartData: ChartAxis[];
  tableData: BtcPriceRecord[];
}

export interface LastBtcPriceElement {
  value: number;
  updated: string;
  timestamp: number;
}

@Component({
  selector: 'app-list-prices',
  templateUrl: './list-prices.component.html',
  styleUrls: ['./list-prices.component.css']
})
export class ListBtcPricesComponent implements OnInit {
  chart = [];
  @ViewChild('btcPricesPaginator', { static: false }) btcPricesPaginator: MatPaginator;

  btcPricedisplayedColumns: string[] = ['ID', 'value', 'updated'];
  LastBtcPrice:LastBtcPriceElement
  BtcPricesDataSource: MatTableDataSource<BtcPriceElement>;
  constructor(
    public dialog: MatDialog,
    public btcPriceLoggingService: BtcPriceLoggingService,
    private alertService: AlertService,
  ) {}

  listBtcPrices() {
    this.btcPriceLoggingService.getBtcPrices().subscribe(data => {
      this.BtcPricesDataSource =  new MatTableDataSource( data.tableData.reverse()); 
      console.log(data.tableData.length)
      this.BtcPricesDataSource.paginator = this.btcPricesPaginator;
      this.formatChart(data.chartData)
    })
  }

  getBtcPrice() {
    this.btcPriceLoggingService.getBtcPrice().subscribe(data => {
      this.LastBtcPrice = data; 
    })
  }

  formatChart(dataset:ChartAxis[]){
    const labels = dataset.map(item=>item.x)
    this.chart = new Chart('canvas', {
      type: 'line',
      data: {
        labels: labels,
        datasets: [
          { 
            data: dataset,
            borderColor: "#3cba9f",
            borderWidth: 1
          }
        ]
      },
      options: {
        legend: {
          display: false
        },
        scales: {
          xAxes: [{
            display: true
          }],
          yAxes: [{
            display: true
          }],
        }
      }
    });
  }
  ngOnInit() {
    this.listBtcPrices();
    this.getBtcPrice();
  }
}