import { Component, NgZone } from '@angular/core';
import { Router } from '@angular/router';

import { AuthGuard } from '../auth.service';

declare const gapi: any;

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  providers: [ AuthGuard ]
})

export class LoginComponent {
  googleLoginButtonId = 'google-login-button';

  constructor( private _zone: NgZone,private router:Router, private authGuard: AuthGuard) {}

  ngAfterViewInit() { // Converts the Google login button stub to an actual button.
    gapi.signin2.render(this.googleLoginButtonId, {
      'onsuccess': this.onLoginSuccess,
      'onfailure': this.onLoginFailure,
      'scope': 'profile email',
      'theme': 'dark',
      'width': 240,
      'height': 50,
      'longtitle': true
    });

  }

  private onLoginSuccess = (user: any) => {
    this._zone.run(() => {
      var email = user.getBasicProfile().getEmail();
      var auth_token = user.getAuthResponse().id_token;

      // Store token in local storage for authentication
      localStorage.setItem('email', email);
      localStorage.setItem('google_token', auth_token);

      this.authGuard.getToken().subscribe(
        data => {
          console.log(data._body);
          localStorage.setItem('my_token', data._body);
          this.router.navigate(['/members']);
        },
        error => console.error("error retrieving data", error)
        );

      //console.log(email);


      //console.log(localStorage.getItem('google_token'));
    });
  }

  private onLoginFailure = (error: any) => {
    console.log(error); }
  }