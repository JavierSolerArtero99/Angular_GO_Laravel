import { THIS_EXPR } from '@angular/compiler/src/output/output_ast';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { User, UserService } from '../../core';

@Component({
  selector: 'app-layout-header',
  templateUrl: './header.component.html'
})
export class HeaderComponent implements OnInit {
  constructor(
    private userService: UserService,
  ) {}

  router: Router;
  currentUser: User;
  isAuthed: boolean = false;
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
    console.log(this.currentUser);
    
    
    this.userService
    .attemptAuth("login", this.currentUser)
    .subscribe(
      data => {
        console.log(data);
        this.isModeAdmin = true;
      },
      err => console.log(err)
    );

    this.isModeAdmin=false;
  }

  loginAdmin() {
    console.log('loginAdmin');
    
    this.isAuthed = true;

    this.userService
    .adminAttemptAuth(this.currentUser)
    .subscribe(
      data => {
        console.log(data);
        this.isModeAdmin = true;
      },
      err => console.log(err)
    );

    this.isModeAdmin=true;
  }
}
