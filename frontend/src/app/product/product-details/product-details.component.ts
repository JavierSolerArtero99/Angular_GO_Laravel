import { Component, OnInit } from "@angular/core";
import { FormControl } from "@angular/forms";
import { ActivatedRoute, Router } from "@angular/router";
import { User } from "../../core";
import { Product } from "../shared/product.model";
import { ProductService } from "../shared/product.service";

@Component({
  selector: "app-product-details",
  templateUrl: "./product-details.component.html",
})
export class ProductDetailsComponent implements OnInit {
  products: Product[];
  product: Product;
  currentUser: User;
  canModify: boolean;
  comments: Comment[];
  commentControl = new FormControl();
  commentFormErrors = {};
  isSubmitting = false;
  isDeleting = false;

  constructor(
    private productService: ProductService,
    private route: ActivatedRoute,
    // private articlesService: ArticlesService,
    // private commentsService: CommentsService,
    private router: Router
  ) // private userService: UserService
  {}

  ngOnInit() {
    this.productService.getProductList(null).subscribe((data) => {
      this.products = data.products;
      console.log(data);
    });
  }
}
