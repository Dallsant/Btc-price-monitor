import { backend_url } from '../../app.component';
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import * as moment from 'moment';

@Injectable({
  providedIn: 'root'
})
export class CryptoPriceLoggingService {

  private subject = new Subject<any>();
  subject$ = this.subject.asObservable();

  constructor(
    private httpClient: HttpClient
  ) { }

  getCryptoPrice(): Observable<any> {
    return this.httpClient.get<any>(backend_url + '/crypto-price')
      .pipe();
  }

  getCryptoPrices(): Observable<any> {
      return this.httpClient.get<any>(backend_url + '/crypto-prices')
      .pipe();
  }

}
