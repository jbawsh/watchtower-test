import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

declare const gapi: any;

@Component({
  selector: 'app-members',
  templateUrl: './members.component.html',
  styleUrls: ['./members.component.css']
})
export class MembersComponent {

  email: string;

  constructor(private router: Router) { }

  logout() {
    let that = this;
    var auth2 = gapi.auth2.getAuthInstance();
    auth2.signOut().then(function () {
      console.log('User signed out');
      that.router.navigateByUrl('/login');
      localStorage.removeItem('email');
      localStorage.removeItem('google_token')
    });
    //console.log(gapi.auth2);

  }

  ngAfterViewInit() {
    //this.logout();
  }

}
