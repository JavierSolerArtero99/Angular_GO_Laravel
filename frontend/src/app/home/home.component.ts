import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { ArticleListConfig, TagsService, UserService } from '../core';
import { StatsService } from '../core/services/stats.service';
import { Product } from '../product/shared/product.model';
import { ProductService } from '../product/shared/product.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  constructor(
    private router: Router,
    private userService: UserService,
    private statsService: StatsService,
    private productsService: ProductService,
  ) { }

  currentUsers: number = 0;
  product: Product;
  isAuthenticated: boolean;
  listConfig: ArticleListConfig = {
    type: 'all',
    filters: {}
  };

  ngOnInit() {
    this.userService.isAuthenticated.subscribe(
      (authenticated) => {
        this.isAuthenticated = authenticated;

        // set the article list accordingly
        if (authenticated) {
          this.setListTo('feed');
        } else {
          this.setListTo('all');
        }
      }
    );

    this.statsService.getCurrentUsersCache().subscribe((data) => {
      this.currentUsers = data.users;
    });

    this.statsService.getValoredProducts().subscribe((data) => {
      this.product = data.buys.sort(function (a, b) {
        if (a.TimesBuyed < b.TimesBuyed) {
          return 1;
        }
        if (a.TimesBuyed > b.TimesBuyed) {
          return -1;
        }
        return 0;
      })[0];
    });
  }

  setListTo(type: string = '', filters: Object = {}) {
    // If feed is requested but user is not authenticated, redirect to login
    if (type === 'feed' && !this.isAuthenticated) {
      this.router.navigateByUrl('/login');
      return;
    }

    // Otherwise, set the list object
    this.listConfig = { type: type, filters: filters };
  }
}
