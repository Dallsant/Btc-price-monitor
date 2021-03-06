import { Component, ViewChild, OnInit, HostListener } from '@angular/core';
import { MatMenuTrigger } from '@angular/material';
export const backend_url = 'http://localhost:8000/';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';


export  let isLogged: Boolean = false;

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  isLogged:boolean = false
  constructor(
    private httpClient: HttpClient,
    private router: Router,
  ) {
  }

  title = 'Crypto Price Record';
  @ViewChild(MatMenuTrigger, { static: false }) menu: MatMenuTrigger;
  
  ngOnInit() {
    localStorage.setItem('Authorization', '');
  }
}
