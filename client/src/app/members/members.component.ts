import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthGuard } from '../auth.service';

declare const gapi: any;

@Component({
  selector: 'app-members',
  templateUrl: './members.component.html',
  styleUrls: ['./members.component.css'],
  providers: [ AuthGuard ]
})
export class MembersComponent {

  email: string;

  public data: any;

  constructor(private router: Router, private authGuard: AuthGuard) { }

  logout() {
    let that = this;
    var auth2 = gapi.auth2.getAuthInstance();
    auth2.signOut().then(function () {
      console.log('User signed out');
      that.router.navigateByUrl('/login');
      localStorage.removeItem('email');
      localStorage.removeItem('google_token');
      localStorage.removeItem('my_token');
    });
    //console.log(gapi.auth2);

  }

  testData() {
    this.authGuard.getData().subscribe(
      data => {
        console.log(data);
        this.data = data;
      },
      error => console.log("error"),
    )
  }

  ngAfterViewInit() {
    //this.logout();
  }

}
