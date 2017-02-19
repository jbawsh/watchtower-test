import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { Http, Response, Headers } from '@angular/http';

import { Observable } from 'rxjs/Observable';

import 'rxjs/add/operator/map';
//import 'rxjs/add/operator/do';

@Injectable()
export class AuthGuard implements CanActivate {

  private tokenUrl = 'http://localhost:3000/get-token';

  private dogsUrl = 'http://localhost:3000/dogs';

  constructor(private http: Http, private router: Router) { }

  canActivate(): boolean {
    if (localStorage.getItem('email')) {
      return true;
    } else {
      this.router.navigate(['/login']);
      return false;
    }
  }

  getToken(): Observable<any> {

    console.log("asks for token");

    var body = localStorage.getItem('google_token');

    console.log(JSON.stringify(body));

    body = JSON.stringify(body);

    var headers = new Headers();
    //headers.append('')

    return this.http.post(this.tokenUrl, body, { headers: headers });
      //.map(response => response.json())
      //.do(response => console.log(response.json()));
  }

  getData(): Observable<any> {
    console.log("asks for data");

    var token = localStorage.getItem("my_token");
    var tokenString = token;

    var headers = new Headers();
    headers.append('Authorization', tokenString);

    return this.http.get(this.dogsUrl, {headers: headers})
      .map(response => response.json());
  }
}
