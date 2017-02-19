import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { Http, Response, Headers } from '@angular/http';

import { Observable } from 'rxjs/Observable';

import 'rxjs/add/operator/map';

@Injectable()
export class AuthGuard implements CanActivate {

  private tokenUrl = 'http://localhost:3000/get-token'

  constructor(private http: Http, private router: Router) { }

  canActivate(): boolean {
    if (localStorage.getItem('email')) {
      return true;
    } else {
      this.router.navigate(['/login']);
      return false;
    }
  }

  getToken() {

    console.log("asks for token");

    var body = localStorage.getItem('google_token');
    var headers = new Headers();
    //headers.append('')

    this.http.post(this.tokenUrl, body, { headers: headers })
      .map(response => response.json());
  }
}
