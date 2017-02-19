import { Component, NgZone } from '@angular/core';
import { Router } from '@angular/router';

declare const gapi: any;

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent {
  googleLoginButtonId = 'google-login-button';

  constructor( private _zone: NgZone,private router:Router) {}

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

      //console.log(email);
      this.router.navigate(['/members']);

      //console.log(localStorage.getItem('google_token'));
    });
  }

  private onLoginFailure = (error: any) => {
    console.log(error); }
  }