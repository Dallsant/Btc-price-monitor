import { backend_url } from '../../app.component';
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import * as moment from 'moment';

@Injectable({
  providedIn: 'root'
})
export class BsPriceLoggingService {

  private subject = new Subject<any>();
  subject$ = this.subject.asObservable();

  constructor(
    private httpClient: HttpClient
  ) { }

  getBsPrice(): Observable<any> {
    return this.httpClient.get<any>(backend_url + '/bs-price')
      .pipe();
  }

  getBsPrices(): Observable<any> {
      return this.httpClient.get<any>(backend_url + '/bs-prices')
      .pipe();
  }

}
