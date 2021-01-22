import { Component, Input, OnInit } from "@angular/core";
import { UserService } from "../../core";
import { ProductService } from "../shared/product.service";

import { Product } from "./../shared/product.model";

@Component({
  selector: "app-product-preview",
  templateUrl: "./product-preview.component.html",
})
export class ProductPreviewComponent implements OnInit {
  @Input() product: Product;

  constructor(
    private productsService: ProductService,
    private userService: UserService
  ) {}

  ngOnInit(): void {
    console.log(this.product);
  }

  likeProduct() {
    const isLiked = this.product.LikesList.some(
      (like) => like.UserID === this.userService.getCurrentUser().id
    );

    console.log(this.productsService);

    if (!isLiked) {
      // dandole like al producto
      this.productsService
        .likeProduct(this.product, this.userService.getCurrentUser().id)
        .subscribe((data) => {
          this.product.Likes++;
          this.product.LikesList.push({
            ProductID: this.product.Id,
            UserID: this.userService.getCurrentUser().id,
          });
        });
    } else {
      // dandole like al producto

      this.product.Likes--;
      this.product.LikesList = this.product.LikesList.filter(
        (like) => like.ProductID !== this.product.Id
      );
    }
  }
}
