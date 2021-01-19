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
  ) {}

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
      this.currentUsers = data.current_users
    });

    // this.statsService.getProductCache().subscribe((data) => {
      this.productsService.getSingleProduct("Camiseta").subscribe((data) => {
        this.product = data.product;
      });
    // });
  }

  setListTo(type: string = '', filters: Object = {}) {
    // If feed is requested but user is not authenticated, redirect to login
    if (type === 'feed' && !this.isAuthenticated) {
      this.router.navigateByUrl('/login');
      return;
    }

    // Otherwise, set the list object
    this.listConfig = {type: type, filters: filters};
  }
}
