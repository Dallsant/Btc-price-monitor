import { backend_url } from '../../app.component';
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import * as moment from 'moment';

@Injectable({
  providedIn: 'root'
})
export class BtcPriceLoggingService {

  private subject = new Subject<any>();
  subject$ = this.subject.asObservable();

  constructor(
    private httpClient: HttpClient
  ) { }

  getBtcPrice(): Observable<any> {
    return this.httpClient.get<any>(backend_url + '/btc-price')
      .pipe();
  }

  getBtcPrices(): Observable<any> {
      return this.httpClient.get<any>(backend_url + '/btc-prices')
      .pipe();
  }

}
