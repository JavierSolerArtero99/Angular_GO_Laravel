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
  loading: boolean = true;

  constructor(
    private productService: ProductService,
    private route: ActivatedRoute,
    // private articlesService: ArticlesService,
    // private commentsService: CommentsService,
    private router: Router
  ) // private userService: UserService
  {}

  ngOnInit() {
    // slicepara cojer el nombre del producto 
    let productName = this.router.url.substr(9)
    
    // obteniendo los datos del producto
    this.productService.getSingleProduct(productName).subscribe((data) => {
      this.product = data.product;
      this.loading = false
      console.log(this.product);
    });
  }
}
