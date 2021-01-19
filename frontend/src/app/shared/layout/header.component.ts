import { Component, OnInit } from '@angular/core';

import { User, UserService } from '../../core';

@Component({
  selector: 'app-layout-header',
  templateUrl: './header.component.html'
})
export class HeaderComponent implements OnInit {
  constructor(
    private userService: UserService,
  ) {}

  currentUser: User;
  isModeAdmin: boolean = false;

  ngOnInit() {
    this.userService.currentUser.subscribe(
      (userData) => {
        this.currentUser = userData;
      }
    );
  }

  loginClient() {
    console.log('loginClient');
    
    // this.userService
    // .attemptAuth(this.currentUser)
    // .subscribe(
    //   data => {
    //     console.log(data);
    //     this.isModeAdmin = true;
    //   },
    //   err => console.log(err)
    // );

    this.isModeAdmin=false;
  }

  loginAdmin() {
    console.log('loginAdmin');
    
    // this.userService
    // .adminAttemptAuth(this.currentUser)
    // .subscribe(
    //   data => {
    //     console.log(data);
    //     this.isModeAdmin = true;
    //   },
    //   err => console.log(err)
    // );

    this.isModeAdmin=true;
  }
}
