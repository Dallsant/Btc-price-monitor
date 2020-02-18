import { Component, OnInit, ViewChild, Inject, ChangeDetectorRef  } from '@angular/core';
import { MatPaginator, MatTableDataSource } from '@angular/material';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { AlertService } from 'ngx-alerts';
import { BsPriceLoggingService } from '../bs-price-logging.service';
import { backend_url } from '../../../app.component';
import { Chart } from 'chart.js';

export interface ChartAxis {
  x: string;
  y: number;
}

export interface BsPriceRecord {
  ID: number;
  value: number;
  updated: string;
  timestamp: number;
}

export interface BsPriceElement {
  chartData: ChartAxis[];
  tableData: BsPriceRecord[];
}

export interface LastBsPriceElement {
  value: number;
  updated: string;
  timestamp: number;
}

@Component({
  selector: 'app-list-prices',
  templateUrl: './list-bs-prices.component.html',
  styleUrls: ['./list-bs-prices.component.css']
})
export class ListBsPricesComponent implements OnInit {
  chart = [];
  @ViewChild('bsPricesPaginator', { static: false }) bsPricesPaginator: MatPaginator;

  bsPricedisplayedColumns: string[] = ['ID', 'value', 'updated'];
  LastBtcPrice:LastBsPriceElement
  BsPricesDataSource: MatTableDataSource<BsPriceElement>;
  constructor(
    public dialog: MatDialog,
    public bsPriceLoggingService: BsPriceLoggingService,
    private alertService: AlertService,
  ) {}

  listBsPrices() {
    this.bsPriceLoggingService.getBsPrices().subscribe(data => {
      this.BsPricesDataSource =  new MatTableDataSource( data.tableData.reverse()); 
      this.BsPricesDataSource.paginator = this.bsPricesPaginator;
      this.formatChart(data.chartData)
    })
  }

  getBsPrice() {
    this.bsPriceLoggingService.getBsPrice().subscribe(data => {
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
    this.listBsPrices();
    this.getBsPrice();
  }
}